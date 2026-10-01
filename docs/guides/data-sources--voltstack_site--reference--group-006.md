---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-2030232122221311-1023110023321001-2110001220110310-3120212113121310-2310332323101203-1331203130331231-2021020023101132-0331103101112023"></a>

## custom_storage_config.storage_class_list.storage_classes — storage_classes / 203032033033 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333)
- custom_storage_config.storage_class_list.storage_classes

<a id="canonical-3133213000010220-2101310031202100-2122022021122323-2130301100112211-2220233133323332-2311033322113321-3200130120013212-3303211302031022"></a>

Type: `"list"`. Computed.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2213221311112213-3213333103030101-1311120031101330-1213020321232212-2333321323231213-2232030313103100-1110133112001332-3012323200012030"></a>

## Direct properties — storage_classes / 203032033033 / 3

<a id="canonical-0332223122301100-0310330022122023-3330301311233101-1223130233320333-1223112011110311-2012200110113300-1331301210021312-0320023112333002"></a>

<a id="canonical-1132032203101212-1303312031012013-0023103121211322-1000120113002003-2031102012331103-1230123000110100-2220121122011212-2010302111300302"></a>

## advanced_storage_parameters property — storage_classes / 203032033033 / 4

Type: `["map", "string"]`. Computed.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

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
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-1022012233003221-3201030211301001-3112321200330112-3331022112223021-1210312113011000-2003022303201100-0000012131210300-0322011110113211"></a>

<a id="canonical-0113220303130122-2001211023332213-2021111233333321-0010202222120213-0230112130330222-1101003312122033-3111113111101001-3030221021213103"></a>

## allow_volume_expansion property — storage_classes / 203032033033 / 5

Type: `"bool"`. Computed.

Allow Volume Expansion. Allow volume expansion.

Upstream description:

Allow volume expansion.

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

- [custom_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1322100112301000-3320130211010212-1211322301011211-1201121020333033-1132123112232220-3201231101201010-2102133333021022-1132203203221322): complete subsection reference.

<a id="canonical-3201030002222002-1212003311201302-0122032331132310-3003003312322200-3101120211112323-2121020031113033-1321112312322331-0011210221300313"></a>

<a id="canonical-2012011120200110-1321320313330013-3222212102301110-3203200123221323-3330000230030231-0303211000132201-1221302221013230-0122213113112312"></a>

## default_storage_class property — storage_classes / 203032033033 / 6

Type: `"bool"`. Computed.

Make this storage class default storage class for the K8s cluster.

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

<a id="canonical-1312233201013031-2130302311023231-1120102032013103-2202311303231322-3030130222221132-3133330101130303-3230111032001221-1301203130012331"></a>

<a id="canonical-2002221022012220-2330021301221011-3030110110023111-3002312111031010-3023302131001022-0310232130203110-0330231301122021-2013021333320103"></a>

## description_spec property — storage_classes / 203032033033 / 7

Type: `"string"`. Computed.

Storage Class Description. Description for this storage class.

- [hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1001010133313201-1021212203222132-1023231101320002-0011003123210100-2131032021320320-3303023111020310-2312213010202322-3210102210222120): complete subsection reference.

- [netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-3132023112030331-2303213311311100-0121132331313303-0011212123323222-2111203302010101-2130023010322133-0233200011002002-0022010221113111): complete subsection reference.

- [pure_service_orchestrator](data-sources--voltstack_site--reference--group-006.md#canonical-2213301022312101-0220001232311302-3103321123100103-2333332331303311-1213112123211302-0323330130321032-0333202133301210-1223202222332031): complete subsection reference.

<a id="canonical-3300131012321300-1300312121330030-2112101022120333-3103203122332031-3020303130032200-2231231103201313-1130103213330301-1202033221001300"></a>

<a id="canonical-0201300112323110-2121213333110111-1223311323200110-2321313303321010-3221113211303131-1003131222023123-2000221311203111-1031331312031030"></a>

## reclaim_policy property — storage_classes / 203032033033 / 8

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Reclaim Policy.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-3123310312102101-0212122121222103-3222210131313323-2211201011120230-1002123231012030-0112120320030020-1022310310310333-1030203200120020"></a>

<a id="canonical-3322311222200223-2030020010022301-3233012230120012-3220111013210230-2301002203301022-0220311331202221-2213221203213020-1310303320201123"></a>

## storage_class_name property — storage_classes / 203032033033 / 9

Type: `"string"`. Computed.

Name of the storage class as it will appear in K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0111130301322303-0033220123122221-3333301232120132-3012123220203213-2202012201010120-2011113313200032-2232230021220310-1210332102013022"></a>

<a id="canonical-1111222333123110-1112233231020222-2100311221113032-1310022232313322-0101032211123123-2022101212300213-0311221201312102-2201023020230111"></a>

## storage_device property — storage_classes / 203032033033 / 10

Type: `"string"`. Computed.

Storage device that this class will use. The Device name defined at previous step.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64,
      "min": 1
    },
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0210012023221013-0123112201031221-0003320100121003-1112221112122232-0223331233300233-2310322200102213-0111312311013331-3233323311030123"></a>

## Next pages — storage_classes / 203032033033 / 11

- [custom_storage_config.storage_class_list.storage_classes.custom_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1322100112301000-3320130211010212-1211322301011211-1201121020333033-1132123112232220-3201231101201010-2102133333021022-1132203203221322)
- [custom_storage_config.storage_class_list.storage_classes.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1001010133313201-1021212203222132-1023231101320002-0011003123210100-2131032021320320-3303023111020310-2312213010202322-3210102210222120)
- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-3132023112030331-2303213311311100-0121132331313303-0011212123323222-2111203302010101-2130023010322133-0233200011002002-0022010221113111)
- [custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator](data-sources--voltstack_site--reference--group-006.md#canonical-2213301022312101-0220001232311302-3103321123100103-2333332331303311-1213112123211302-0323330130321032-0333202133301210-1223202222332031)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1322100112301000-3320130211010212-1211322301011211-1201121020333033-1132123112232220-3201231101201010-2102133333021022-1132203203221322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101323130300100-0012323322222200-3001110110213033-2112032202022021-1302232213022003-3311012200110311-2302203003201211-3110002001203212"></a>

## custom_storage_config.storage_class_list.storage_classes.custom_storage — custom_storage / 232022010303 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332)
- custom_storage_config.storage_class_list.storage_classes.custom_storage

<a id="canonical-1211221012233210-2110032322123030-0323230231331332-3323200212200000-1022102020000110-1121320102233102-1111301012130000-1132301102000230"></a>

Type: `"single"`. Computed.

Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into
given site.

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

<a id="canonical-3230332103311220-2231331101032011-1320131212233210-1023320020330203-0220320221103011-1311231232332232-0033201112030113-2331330331212123"></a>

## Direct properties — custom_storage / 232022010303 / 3

<a id="canonical-2220223121222001-3011202002200022-3111221020132133-2000200231123212-0303213202321001-1022223231003212-3020001003212031-0322212032220010"></a>

<a id="canonical-0123113321323220-1122330303023321-2201101030023012-2303000213031113-0000031001202310-2332001210320013-3231123031101030-0202211330311313"></a>

## yaml property — custom_storage / 232022010303 / 4

Type: `"string"`. Computed.

Storage Class YAML. K8s YAML for StorageClass.

Upstream description:

K8s YAML for StorageClass.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3100202022000020-3123211202123311-0222212301300213-2231020102320321-2313302121112332-3001321202201001-2210203121030321-2310211101212230"></a>

## Next pages — custom_storage / 232022010303 / 5

- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1001010133313201-1021212203222132-1023231101320002-0011003123210100-2131032021320320-3303023111020310-2312213010202322-3210102210222120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120002130333230-0002001022232203-0033010201100120-3320120133120012-3002310231201120-1032031333213102-2301123300313022-3100231002000333"></a>

## custom_storage_config.storage_class_list.storage_classes.hpe_storage — hpe_storage / 313131211013 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332)
- custom_storage_config.storage_class_list.storage_classes.hpe_storage

<a id="canonical-1120231012313213-2220312202220212-2120302032223220-1332331120133220-0313230330102211-2233110110100303-1301232203023321-2103203303020001"></a>

Type: `"single"`. Computed.

Storage class Device configuration for HPE Storage.

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

<a id="canonical-0311301010000012-2331220113032030-1021022331313030-1033120320031203-3002301323112121-0020201202310100-1032001223121010-2210112220022230"></a>

## Direct properties — hpe_storage / 313131211013 / 3

<a id="canonical-1001123202312133-1031131123311223-0332033330011210-2232012110100311-0203022232030333-0030332201010323-3002213302010313-3231023311002121"></a>

<a id="canonical-2322011011030123-3112131122211001-3031312232101330-0310321322231330-2200023032032210-2102312003310020-2031100221013330-2121023311032031"></a>

## allow_mutations property — hpe_storage / 313131211013 / 4

Type: `"string"`. Computed.

Mutation can override specified parameters.

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

<a id="canonical-1331223121032202-1223002021021303-1211201231132022-3012020033202112-2200330001301221-3231102033203203-0323323302000113-0013300012110201"></a>

<a id="canonical-2131201211231230-0323032231320020-1121132013223100-1301020030330021-2312221021132010-3103131202300113-1331002231111310-2011022132230020"></a>

## allow_overrides property — hpe_storage / 313131211013 / 5

Type: `"string"`. Computed.

AllowOverrides. PVC can override specified parameters.

Upstream description:

PVC can override specified parameters.

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

<a id="canonical-2300111121120201-0211212310013022-0122120031122122-2323203112222302-1223120211321210-1102201120330112-2000202303211221-0001031333330312"></a>

<a id="canonical-1031332132011110-0001110302000102-3323300200213003-2131001333321222-1233323013202133-3331213032332221-3000312202123220-1120323003312231"></a>

## dedupe_enabled property — hpe_storage / 313131211013 / 6

Type: `"bool"`. Computed.

Indicates that the volume should enable deduplication.

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

<a id="canonical-1102211003310010-2102003030120203-1023222123112313-1312330021023220-1321213302022122-3330020100100103-2101302011201000-2022102111201001"></a>

<a id="canonical-2032121132301010-3310323010032322-3102213133013320-0203320023213030-0120330223303003-1130322003020022-1300131223210233-0312331321032123"></a>

## description_spec property — hpe_storage / 313131211013 / 7

Type: `"string"`. Computed.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

<a id="canonical-3331130203312001-1120233332201223-1023131123300233-0033012200320323-2113321312213331-0103322311301221-3300110213310131-2102132101313101"></a>

<a id="canonical-3132120122213310-2330101121131301-3231022300320102-0331101213021213-0102123123021312-3133130313112113-3312203330201131-2223123111233102"></a>

## destroy_on_delete property — hpe_storage / 313131211013 / 8

Type: `"bool"`. Computed.

Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is
deleted.

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

<a id="canonical-1310031321211221-1313030200103221-2020213113210032-2302110211331032-0000133122330123-2033013123023232-3320301131111231-3000302010110130"></a>

<a id="canonical-1310133101203220-0132002111120101-3203022110303010-2113121300000010-1201330020020113-0010121233301201-3130201303210003-3212203213010030"></a>

## encrypted property — hpe_storage / 313131211013 / 9

Type: `"bool"`. Computed.

Indicates that the volume should be encrypted.

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

<a id="canonical-1331232202222101-2132032302113233-1121313212021000-0323222120322121-1012331330022102-2313202332132011-3321211132102003-1332323000231230"></a>

<a id="canonical-0230332000102311-3220231313330121-2023112103233130-2022233333023101-2300302211130200-0113332031200002-2201030230031111-2232232330310011"></a>

## folder property — hpe_storage / 313131211013 / 10

Type: `"string"`. Computed.

The name of the folder in which to place the volume.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1100223232330030-0301300323302301-1101330233102033-3201231120112112-2312022301000110-0121120033233021-2111012132111002-2223133212123011"></a>

<a id="canonical-3202232100321000-2112203123303113-2010322311023200-0330221031331133-0001132312033111-2332113001021222-1300002212102110-0323011000113011"></a>

## limit_iops property — hpe_storage / 313131211013 / 11

Type: `"string"`. Computed.

LimitIops. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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
  }
}
```

<a id="canonical-1001331232332310-3333210303200021-2011232323222222-1010130103220110-2001110210032201-1012010200010022-1231322023001032-2122100233212132"></a>

<a id="canonical-0103020021313033-3320232133112033-1202111101103203-3202121202301311-0211022220103132-0011123022313211-0120320023132031-0203022021311000"></a>

## limit_mbps property — hpe_storage / 313131211013 / 12

Type: `"string"`. Computed.

LimitMbps. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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
  }
}
```

<a id="canonical-1201230022012313-3022201331133010-0230033212022200-2121230230323122-0322101110010003-3330231121110210-1302320231321032-2021023300200323"></a>

<a id="canonical-0232133233221103-0030232103111010-2133132223300010-3302321120122103-2122000121130012-0112330033112031-2022223002331311-2301122232222112"></a>

## performance_policy property — hpe_storage / 313131211013 / 13

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

The name of the performance policy to assign to the volume.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2100122202100230-0202321122100132-0001103233101021-2123132121023020-2101022111332010-3100010213031122-3031221113130130-1333312001111331"></a>

<a id="canonical-1033020120002301-0213221310333313-0030112103233300-0310323133002330-1333031133130301-0110212023201000-2330032202010300-1010031202121101"></a>

## pool property — hpe_storage / 313131211013 / 14

Type: `"string"`. Computed.

The name of the pool in which to place the volume.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1100012012023202-3231212110132223-2212112310132030-0000002303102321-1002232211312231-1230232310223231-1112113101303003-2000120302300110"></a>

<a id="canonical-1212030302213103-0023132121232102-2112200222020332-3103102100000033-2233033223013321-2201120200000010-3301111112011231-2231132032201212"></a>

## protection_template property — hpe_storage / 313131211013 / 15

Type: `"string"`. Computed.

The name of the performance policy to assign to the volume.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3132233201123012-0031212310311311-2230202033133000-0333023320321020-0100132133322320-0323003233330133-2123232223011020-3131212232121103"></a>

<a id="canonical-2111011303123213-2001132112000311-1000103100021113-0022000212322210-1111131030323112-1003230313002331-0013000100202031-3030332220231323"></a>

## secret_name property — hpe_storage / 313131211013 / 16

Type: `"string"`. Computed.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

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

<a id="canonical-1031333233120033-1130101130101001-2113003103312202-0310223000130120-1132221211212302-1001313233111113-3310010022321000-3103331323211233"></a>

<a id="canonical-3320212220113330-0220303312011213-1030133221101220-0212100121000330-3211132012011003-0211003111013032-1200100113311110-2213012031013131"></a>

## secret_namespace property — hpe_storage / 313131211013 / 17

Type: `"string"`. Computed.

The SecretNamespace parameter is used to identify name of namespace where secret resides.

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

<a id="canonical-3331212120310213-2022122013321103-3331322332121201-0321023303123033-0202100122020112-2122121013220313-1103220213233113-1320222321002011"></a>

<a id="canonical-2301000332310231-3031022331200311-3331301133210301-0320132030111132-2103331301012322-2103020231010133-2132300003010133-2320122021222301"></a>

## sync_on_detach property — hpe_storage / 313131211013 / 18

Type: `"bool"`. Computed.

Indicates that a snapshot of the volume should be synced to the replication partner each time it is
detached from a node.

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

<a id="canonical-2320001011102003-3132011230230301-3202310223122231-0120120202213103-1212102003131033-2301122303212303-3201013312123123-3213023223202232"></a>

<a id="canonical-2130211023200010-0103103313313001-2301231030030020-0123121331302111-3000201311103331-2310032120322311-2101033302033333-0333312333312310"></a>

## thick property — hpe_storage / 313131211013 / 19

Type: `"bool"`. Computed.

Indicates that the volume should be thick provisioned.

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

<a id="canonical-2231332300302202-2133020003121320-1302120122200302-1033222131021121-1121203101120222-1322333220120212-0312222211323002-0220003110211312"></a>

## Next pages — hpe_storage / 313131211013 / 20

- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3132023112030331-2303213311311100-0121132331313303-0011212123323222-2111203302010101-2130023010322133-0233200011002002-0022010221113111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002322332322203-2310201020211112-0132122220223133-0311112112333021-0212030320202132-0233022232323333-2123121333230011-0113002113102130"></a>

## custom_storage_config.storage_class_list.storage_classes.netapp_trident — netapp_trident / 230131333130 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident

<a id="canonical-2122232300201103-2320002132011110-3002001031122120-2230321221303122-3031023031310110-2011312122321301-1011031023122200-3100122212122023"></a>

Type: `"single"`. Computed.

Storage class Device configuration for NetApp Trident.

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

<a id="canonical-0131132112321000-1003303223002032-3223131222222321-0313311311203011-3003321222331122-2023110012022313-3221321202212230-0103212110032201"></a>

## Direct properties — netapp_trident / 230131333130 / 3

- [selector](data-sources--voltstack_site--reference--group-006.md#canonical-3121111003211222-1120131012130133-0030001233023233-0032221113332301-2030113021113011-3210103300303333-2113120001002223-0031201013221312): complete subsection reference.

<a id="canonical-3312112232333210-2012021031120131-3220121213331032-2122232102010011-2032111302220100-2330031330003002-2310002010310130-2122111200113010"></a>

<a id="canonical-2013011001001111-3223200221132300-2303212201311332-3003020303300311-1332030130330332-0332011211021222-2121130031123131-0301122031302210"></a>

## storage_pools property — netapp_trident / 230131333130 / 4

Type: `"string"`. Computed.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-0200000310132002-2110121013333332-3221301122001311-1311131012013311-1030002101013021-3322313112001323-2000302203003231-3211022211203300"></a>

## Next pages — netapp_trident / 230131333130 / 5

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector](data-sources--voltstack_site--reference--group-006.md#canonical-3121111003211222-1120131012130133-0030001233023233-0032221113332301-2030113021113011-3210103300303333-2113120001002223-0031201013221312)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3121111003211222-1120131012130133-0030001233023233-0032221113332301-2030113021113011-3210103300303333-2113120001002223-0031201013221312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301113110120200-3120022011213323-0031322001130303-1311311102332122-2103103233312301-0012130023210301-1132332312110220-3311111100232120"></a>

## custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector — selector / 333221021222 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332)
- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-3132023112030331-2303213311311100-0121132331313303-0011212123323222-2111203302010101-2130023010322133-0233200011002002-0022010221113111)
- custom_storage_config.storage_class_list.storage_classes.netapp_trident.selector

<a id="canonical-1020311312111332-3313330021022132-1110303101112332-2003122202301123-0333020122033000-3012012310013100-3032103122220111-2232133111003311"></a>

Type: `"single"`. Computed.

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

Upstream description:

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

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

<a id="canonical-2231013210233231-3230133113012332-1333300001030033-3320320203201110-3132230012203032-0311020232230221-3130031302020102-0003322010212300"></a>

## Direct properties — selector / 333221021222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301220112102121-1311101010211021-2131303032212001-0323232031320331-3032323311223100-2303003103232233-2000231002313033-3212213201112332"></a>

## Next pages — selector / 333221021222 / 4

- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-3132023112030331-2303213311311100-0121132331313303-0011212123323222-2111203302010101-2130023010322133-0233200011002002-0022010221113111)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2213301022312101-0220001232311302-3103321123100103-2333332331303311-1213112123211302-0323330130321032-0333202133301210-1223202222332031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001030110330232-1030113021030101-2122133323300103-2033110023010200-0313103211200022-1203121100012002-3202113301210102-1013323032202020"></a>

## custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator — pure_service_orchestrator / 012121032221 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333)
- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332)
- custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator

<a id="canonical-3213200231210123-3220102013323032-0332101202002133-3211003331102312-1012222113213130-1232211030032220-2201102332020021-1322023000303102"></a>

Type: `"single"`. Computed.

Storage class Device configuration for Pure Service Orchestrator.

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

<a id="canonical-2323031102113113-3300122313123103-2100313211033031-2310223001200131-2021010331331112-0102122223002201-3320220233202033-1032021203102303"></a>

## Direct properties — pure_service_orchestrator / 012121032221 / 3

<a id="canonical-3233101111032132-3123302133323032-3112012123020311-1220310032123203-2013013010111210-1223213312300233-0303232222300220-1011231322002120"></a>

<a id="canonical-0110110233121232-0033320330233212-0131102121112200-2111101223210303-3112130323311132-2131101211310123-0230202331001302-0303322333311311"></a>

## backend property — pure_service_orchestrator / 012121032221 / 4

Type: `"string"`. Computed.

\[Enum: block|file\] Defines type of Pure storage backend block or file. The volume will have the
aspects defined in the chosen virtual pool. Possible values are \`block\`, \`file\`.

Upstream description:

Defines type of Pure storage backend block or file. The volume will have the aspects defined in the
chosen virtual pool.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "block",
    "file"
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
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"block\\\",\\\"file\\\"]"
  }
}
```

<a id="canonical-2312132302100312-1103010020201030-1110301133332103-0310122002313101-0132013001012031-0331220201002321-1203222110112111-3010110101323323"></a>

<a id="canonical-1302300303223211-2013122313021313-0002112302000213-3011233112121230-0303022123222213-2122300121001111-2201022102121030-3121031322332230"></a>

## bandwidth_limit property — pure_service_orchestrator / 012121032221 / 5

Type: `"string"`. Computed.

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Upstream description:

It must be between 1 MB/s and 512 GB/s. Enter the size as a number (bytes must be multiple of 512)
or number with a single character unit symbol. Valid unit symbols are K, M, G, representing KiB,
MiB, and GiB.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 12,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 12,
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
    "ves.io.schema.rules.string.max_len": "12"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "12"
  }
}
```

<a id="canonical-0313133332320020-3200103120133123-2221330233133201-2313313113311332-0233113132003220-1010332130203221-0210200232223322-0013101203232233"></a>

<a id="canonical-3023022210012030-1120323133200033-0100330300130001-1022312013130220-2111111001213132-1303201210303332-1202013120010332-3201100113310222"></a>

## iops_limit property — pure_service_orchestrator / 012121032221 / 6

Type: `"number"`. Computed.

Enable IOPS limitation. It must be between 100 and 100 million. If value is 0, IOPS limit is not
defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100000000,
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
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,100-100000000"
  }
}
```

<a id="canonical-3202332123101121-0010333033032311-3132221231230203-2301113011220100-0202022331120103-2203202033222020-0321013221012013-0302123331003312"></a>

## Next pages — pure_service_orchestrator / 012121032221 / 7

- [custom_storage_config.storage_class_list.storage_classes](data-sources--voltstack_site--reference--group-005.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030221123011101-1120000031311212-2320023310222231-3310320012030333-2332030332002031-3322301223223221-2301201321323012-3021113021122111"></a>

## custom_storage_config.storage_device_list — storage_device_list / 020222022003 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- custom_storage_config.storage_device_list

<a id="canonical-2110121303013103-2013320002321000-2222023302212213-0231130220321220-1211302020321103-3003100131032203-2000130020021313-2001022210030112"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this fleet.

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

<a id="canonical-0301012103231031-3323002330310313-1333010103301201-0232332301001201-2132222130302101-3110031033000112-3212322001033312-2022322000132212"></a>

## Direct properties — storage_device_list / 020222022003 / 3

- [storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320): complete subsection reference.

<a id="canonical-1033133131203130-2132020133220212-2020110020312302-0223222031030210-1013130213212310-0233110232302313-0330302223121021-2232031000112211"></a>

## Next pages — storage_device_list / 020222022003 / 4

- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033000312222131-1033123000330132-3203021202130000-1033131322102332-2233222000103300-0110031032312331-0103232023112131-2301331000123330"></a>

## custom_storage_config.storage_device_list.storage_devices — storage_devices / 220233202303 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- custom_storage_config.storage_device_list.storage_devices

<a id="canonical-3113101302221132-1332120022310323-2322103121333011-3000010202030322-2112320032130003-0303131312031310-0312031110020113-1331023010323300"></a>

Type: `"list"`. Computed.

List of Storage Devices. List of custom storage devices.

Upstream description:

List of custom storage devices.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2310322302222123-2330302022222022-3231130013202133-0333211303110320-1133110303022122-1323231131203301-2031111210332223-1130103011120002"></a>

## Direct properties — storage_devices / 220233202303 / 3

<a id="canonical-3111113302032030-1303302330302320-0202031322201202-1102100131331122-2103322303220011-0331013201320123-3000010111232112-1322323002130333"></a>

<a id="canonical-0010021121003300-0022132003022032-1320311232030002-0300113030302301-1000300132001331-3321113000303223-0223020323012021-0022210112200202"></a>

## advanced_advanced_parameters property — storage_devices / 220233202303 / 4

Type: `["map", "string"]`. Computed.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

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
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [custom_storage](data-sources--voltstack_site--reference--group-006.md#canonical-2103323311302212-3012301121020112-3302002000001223-2222301321213320-0233231321301312-3133102321223231-0201330123200320-3320300020223101): complete subsection reference.

- [hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303): complete subsection reference.

- [netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213): complete subsection reference.

- [pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220): complete subsection reference.

<a id="canonical-2013213203030033-0021020012333212-2321130200110100-0131100010030310-1001332033002222-0201103320111210-1132203300231111-3303100212202232"></a>

<a id="canonical-3012202131101211-0122220201103300-1203130012120331-1101321113133201-0111223212220121-0233021220122220-3310202022333302-0102032000310302"></a>

## storage_device property — storage_devices / 220233202303 / 5

Type: `"string"`. Computed.

Storage Device. Storage device and device unit.

Upstream description:

Storage device and device unit.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0233311002311321-1233130320312232-3122330130331210-2320231303221123-0010321230103312-3213003232201121-1231020302311232-0332103102130022"></a>

## Next pages — storage_devices / 220233202303 / 6

- [custom_storage_config.storage_device_list.storage_devices.custom_storage](data-sources--voltstack_site--reference--group-006.md#canonical-2103323311302212-3012301121020112-3302002000001223-2222301321213320-0233231321301312-3133102321223231-0201330123200320-3320300020223101)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2103323311302212-3012301121020112-3302002000001223-2222301321213320-0233231321301312-3133102321223231-0201330123200320-3320300020223101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002203232002130-0301121222200131-1221322033103321-3102002323133221-3111313002022330-0021311103303013-0212030220103212-2223311020323321"></a>

## custom_storage_config.storage_device_list.storage_devices.custom_storage — custom_storage / 022303100122 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- custom_storage_config.storage_device_list.storage_devices.custom_storage

<a id="canonical-2332031032021202-1111330332012301-3202133110202113-3311103301303322-0032003102200211-3131330120023130-3222100211101203-1100102211333012"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for custom storage.

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

<a id="canonical-0132023121111301-2322032113232003-3133330231301311-0313333012002032-2120210101022110-0333112011311021-1002330101010330-2131201320321000"></a>

## Direct properties — custom_storage / 022303100122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000233001002112-1332133333332000-3222113030103303-3331223321122130-0003330202331001-0102313232010322-1022123011023331-1122002022131223"></a>

## Next pages — custom_storage / 022303100122 / 4

- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223212121030133-0203312320132122-1131122012020023-0123032033202021-0320333300022032-2302302021130331-3120323121211031-0112212021222223"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage — hpe_storage / 221033230311 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage

<a id="canonical-1023120032130102-2302130001011320-3120321113221100-2331311103220211-3103033110300103-2212010203322211-3012230111311132-2223102232020111"></a>

Type: `"single"`. Computed.

Configuration parameter for hpe storage.

Upstream description:

Device configuration for HPE Storage.

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

<a id="canonical-2200121222030113-2212110333021100-0122201002021201-0102011233303302-2230323121331220-0111013223210000-3112201030322223-0231203202300033"></a>

## Direct properties — hpe_storage / 221033230311 / 3

<a id="canonical-3033011222032021-2330331100201332-2120210022023333-0303213111220201-3100203012021030-0311012100031113-0123202132133212-1213233121101101"></a>

<a id="canonical-1131111230132320-0121023113333313-2330320221312011-3130232220312222-2003222101211321-0022120001011323-1213331303103020-1030110203022303"></a>

## api_server_port property — hpe_storage / 221033230311 / 4

Type: `"number"`. Computed.

Storage server Port. Enter Storage Server Port.

Upstream description:

Enter Storage Server Port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-0122300323023331-0203100332331313-1111221301330031-0012232301000000-0212021313012301-0120200230200320-1201220232201223-0311312323332133): complete subsection reference.

<a id="canonical-2323303012333102-1232320212033102-2312031133220203-0021330331112130-1113031310120230-0201202021103202-0130211032220033-3322222331230120"></a>

<a id="canonical-2312321233200201-3212003232112120-2022020022011130-1302033113101123-3112111213110121-3210331130212103-1001001201211012-3102003223111230"></a>

## iscsi_chap_user property — hpe_storage / 221033230311 / 5

Type: `"string"`. Computed.

Chap Username to connect to the HPE storage.

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

- [password](data-sources--voltstack_site--reference--group-006.md#canonical-1322030132122002-2231302102210211-2121133010323132-1203111123113000-2233311303011303-3111303113221012-2000211130030321-1300303230333130): complete subsection reference.

<a id="canonical-2100323023022220-2221303120302200-3000201000200222-2132021003123222-1222231223212031-1220212020321013-3122021230330123-1033133232330021"></a>

<a id="canonical-0300211322231000-3203121200200212-3222311232322210-1211020320032002-2113230203331232-0221232320213130-1020112323120300-2030123201131301"></a>

## storage_server_ip_address property — hpe_storage / 221033230311 / 6

Type: `"string"`. Computed.

Storage Server IP address. Enter storage server IP address.

Upstream description:

Enter storage server IP address.

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

<a id="canonical-3203013123301010-0333103220031311-0230131301032202-0032120012102111-0233321102032133-0100023102122233-3311031032003321-1322000232133302"></a>

<a id="canonical-1030001021331030-2321011013030232-1131111003103323-0120332303113021-0222133033310323-0202300223110222-2321211202302121-2102120122102213"></a>

## storage_server_name property — hpe_storage / 221033230311 / 7

Type: `"string"`. Computed.

Storage Server Name. Enter storage server Name.

Upstream description:

Enter storage server Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3313120130130002-1222022101200020-1132033011331113-3000102201323021-2013202110111110-0303131102230111-1133330101203210-2212231303330033"></a>

<a id="canonical-3221033233330313-2131312103311100-0131331030301313-1313222022113201-0223022121100300-1102022011302231-1220201211032011-2331302130011313"></a>

## username property — hpe_storage / 221033230311 / 8

Type: `"string"`. Computed.

Username to connect to the HPE storage management IP.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3312233302322211-3030032232100230-2312113023131312-1132203323032121-3220000132021330-3202120012302303-2302210231332233-1233321003120200"></a>

## Next pages — hpe_storage / 221033230311 / 9

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-0122300323023331-0203100332331313-1111221301330031-0012232301000000-0212021313012301-0120200230200320-1201220232201223-0311312323332133)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-1322030132122002-2231302102210211-2121133010323132-1203111123113000-2233311303011303-3111303113221012-2000211130030321-1300303230333130)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0122300323023331-0203100332331313-1111221301330031-0012232301000000-0212021313012301-0120200230200320-1201220232201223-0311312323332133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120203200123302-0131120231301312-1031033130331022-0222330302202112-3003032302102003-2122010021021013-1122103301001210-0133013121033013"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password — iscsi_chap_password / 310123231210 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

<a id="canonical-1122231331032321-1202313231320220-0031313110021301-2110231330230232-2200011222120100-2213221303331310-3231322203231101-2122300220311331"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-3111302021212022-3323012100331302-0102233123023030-1221120222022233-2112333032003301-1320312232013321-2302322001201310-0131030031013031"></a>

## Direct properties — iscsi_chap_password / 310123231210 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-1221132222333010-2011021011023133-3311133112033130-2103031033122033-2212201323131213-3032131132032023-1333103230221303-1000312131101310): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-1131030220222003-0031032333001021-3102212310202311-0321110213120120-1210321200302102-2210232131302012-0013312110033022-2230313323010132): complete subsection reference.

<a id="canonical-3202121130133112-2203031312021231-1021120132331111-3111122301023233-0112122122231323-0311002330033003-3100132221033112-2132032011321311"></a>

## Next pages — iscsi_chap_password / 310123231210 / 4

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-1221132222333010-2011021011023133-3311133112033130-2103031033122033-2212201323131213-3032131132032023-1333103230221303-1000312131101310)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-1131030220222003-0031032333001021-3102212310202311-0321110213120120-1210321200302102-2210232131302012-0013312110033022-2230313323010132)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1221132222333010-2011021011023133-3311133112033130-2103031033122033-2212201323131213-3032131132032023-1333103230221303-1000312131101310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130311220003001-0022212022321223-2202230002201303-2201110300231103-1131332332122113-1032311003030311-3223303301221301-1332230131123331"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info — blindfold_secret_info / 031223122131 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-0122300323023331-0203100332331313-1111221301330031-0012232301000000-0212021313012301-0120200230200320-1201220232201223-0311312323332133)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info

<a id="canonical-0033030110233021-3311002221331120-1300122100133333-0321332230320112-2301121202011020-2022201301130001-2131123213011022-1311021113332222"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-2110211322223012-0100210023001200-1021112333023312-2220331011032312-1122203001210312-1312312102213011-2003030012213202-1122220331121113"></a>

## Direct properties — blindfold_secret_info / 031223122131 / 3

<a id="canonical-2332320221232111-2032000033213133-2203322333312121-3123100331021212-1132221323110023-0011322033231301-3112120010231111-0230300132022233"></a>

<a id="canonical-2011033130120323-0330202213300200-2132131101101031-1001132100311011-0102323211132303-0032021201322302-2022000210312132-1112310330020331"></a>

## decryption_provider property — blindfold_secret_info / 031223122131 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
  }
}
```

<a id="canonical-0332111212012023-1301123100012130-1000002102233211-3032102102003002-1330022010021020-1102300322113332-2201111123301000-3122010221320112"></a>

<a id="canonical-3033101021011300-1023010202033121-0301032223021133-0131213323203122-3213210210101230-1320122101030210-0222330132230003-0320233231023110"></a>

## location property — blindfold_secret_info / 031223122131 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1120011002130313-1303322133210132-0101131003331010-0103112230311023-3321223311103212-3120001201113030-2033021000331303-3222120322330121"></a>

<a id="canonical-3033301211333010-1323010030233133-2222132001123220-0130032330100201-1130300102202032-1212033302313331-3330231331003322-1000323213010213"></a>

## store_provider property — blindfold_secret_info / 031223122131 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
  }
}
```

<a id="canonical-2300132203330102-3021030011120233-1212230321212312-0133223002311023-2013131023000002-3301003023033210-1203330213332331-3322031121111203"></a>

## Next pages — blindfold_secret_info / 031223122131 / 7

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-0122300323023331-0203100332331313-1111221301330031-0012232301000000-0212021313012301-0120200230200320-1201220232201223-0311312323332133)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1131030220222003-0031032333001021-3102212310202311-0321110213120120-1210321200302102-2210232131302012-0013312110033022-2230313323010132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111133311200020-3333012123023123-3111333301113201-2323122010330031-3332320030301012-0111221111221222-2203100222231212-3300120200222011"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info — clear_secret_info / 200322002013 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-0122300323023331-0203100332331313-1111221301330031-0012232301000000-0212021313012301-0120200230200320-1201220232201223-0311312323332133)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-2021101221320002-3031001120003323-1222021320121232-1322010003321311-1212011000220203-3330031313023131-3330201020212331-0331012130321331"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-2121012211012232-0112302230023210-2202000203103130-1210101302013211-0302313312322001-2323111023120332-3222030102321023-3333022221121123"></a>

## Direct properties — clear_secret_info / 200322002013 / 3

<a id="canonical-1133210203333310-0131030323330032-3321100230132322-2213021101311013-0101132330322220-1021133322220011-3030002301102230-3231031112302001"></a>

<a id="canonical-3303112100022100-2230103131300102-1133211112322233-1200000130011111-1201300110013021-1230121311330231-3330212120033133-2030230233120011"></a>

## provider_ref property — clear_secret_info / 200322002013 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2111120211310210-3021031222123223-3310211112033331-1220200203002213-0212201121201323-1232120320101010-3033113322221011-0010120123320022"></a>

<a id="canonical-0333323112221212-2231123100232021-2012223322000133-0222231120301023-0011310031033132-2211000210230233-0312330323323210-1332102232232332"></a>

## URL property — clear_secret_info / 200322002013 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0231123212201000-0332331303222132-1100210120312212-0022110222201332-0011212331301321-2200111102110133-3010012012213003-0022201103203222"></a>

## Next pages — clear_secret_info / 200322002013 / 6

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--voltstack_site--reference--group-006.md#canonical-0122300323023331-0203100332331313-1111221301330031-0012232301000000-0212021313012301-0120200230200320-1201220232201223-0311312323332133)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1322030132122002-2231302102210211-2121133010323132-1203111123113000-2233311303011303-3111303113221012-2000211130030321-1300303230333130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020010122032021-2202332213013010-1322331123011011-2200312111333230-2131300101322322-0133303200323011-2121120023122001-3120010232011030"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password — password / 013213333030 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-0020013100220301-3032102103002233-3311000130213132-3212013330323201-3212200201203010-0023220030331002-1230133122322310-2001222222323020"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-0313032110011110-1032000121221110-2200203221010322-1200320301011030-0312021013123122-1110132131101223-0212112321333132-2010333011320033"></a>

## Direct properties — password / 013213333030 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-3112000030002000-1301133113201202-2211300313302230-3001032002213022-2322313320303221-1231300301133332-0301332103132211-3331130303312130): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-0322023013022231-0311101030322121-0112310222033233-2113110302101003-1021332000212320-0232210221313231-2303122103303132-3311302322012331): complete subsection reference.

<a id="canonical-3023332302132010-3130103102021320-2002113303022320-0333032302102122-1100001330023010-1231030201312032-2111133300010002-3232012011110233"></a>

## Next pages — password / 013213333030 / 4

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-3112000030002000-1301133113201202-2211300313302230-3001032002213022-2322313320303221-1231300301133332-0301332103132211-3331130303312130)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-0322023013022231-0311101030322121-0112310222033233-2113110302101003-1021332000212320-0232210221313231-2303122103303132-3311302322012331)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3112000030002000-1301133113201202-2211300313302230-3001032002213022-2322313320303221-1231300301133332-0301332103132211-3331130303312130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333110110232303-2101112201332301-2030111302331022-0212313113202020-0200033100301112-0020133333030220-3202230301322311-0211312100121203"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info — blindfold_secret_info / 332332121222 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-1322030132122002-2231302102210211-2121133010323132-1203111123113000-2233311303011303-3111303113221012-2000211130030321-1300303230333130)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-3013002031003103-2331233010330220-2323230223211311-0310212002101303-2020213033330321-0202122030010322-3330013033121213-0223001312020133"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-3121003323032313-3233201130022002-0101212200213012-1003210221332022-1311022032223122-2232232311232313-3211001231213010-3322331330132132"></a>

## Direct properties — blindfold_secret_info / 332332121222 / 3

<a id="canonical-1300011123321200-1210333101013311-1200223002211120-1013110323023332-3020202331310301-0101221201033011-1310011203122313-1120131002011332"></a>

<a id="canonical-1022023000200302-3201331333102233-3210012002223313-0203002213032012-1103231132031103-0102123020002122-3010001223110233-0201133111120321"></a>

## decryption_provider property — blindfold_secret_info / 332332121222 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
  }
}
```

<a id="canonical-2230222302301123-1311113020133122-0232122122210330-0222133201132320-0313110103203333-1210320231303003-1120210112333302-1133212330320013"></a>

<a id="canonical-3121033130003030-3313132012110120-2111120333332320-2112323223013312-2200020131311300-1001130201232033-0010102111110233-2313232312132231"></a>

## location property — blindfold_secret_info / 332332121222 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0203333213032133-0202200222222100-0110102130202022-1102313013021202-0001323001111131-3111121030032220-0013212120022233-1012302101331321"></a>

<a id="canonical-2200013121201230-2201303223233103-3130003012030003-3212201203000210-0103113300333223-1102011330221230-0201233012131002-0021210303010302"></a>

## store_provider property — blindfold_secret_info / 332332121222 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
  }
}
```

<a id="canonical-2321212213130123-3022322322121231-1321232303232103-0023231101032120-1001013012201120-3201331213221101-2331010231223121-1031331333000313"></a>

## Next pages — blindfold_secret_info / 332332121222 / 7

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-1322030132122002-2231302102210211-2121133010323132-1203111123113000-2233311303011303-3111303113221012-2000211130030321-1300303230333130)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0322023013022231-0311101030322121-0112310222033233-2113110302101003-1021332000212320-0232210221313231-2303122103303132-3311302322012331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233111223200302-2012223220223122-3303213101002312-2311321223130211-0100221310322301-0121112300120211-3231132301130022-1001300131212311"></a>

## custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info — clear_secret_info / 331312033333 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage](data-sources--voltstack_site--reference--group-006.md#canonical-1213312310123323-0323000212110011-2001112020113322-1123113022001120-3130323323200332-2331200213213113-1312001232223110-0000301011213303)
- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-1322030132122002-2231302102210211-2121133010323132-1203111123113000-2233311303011303-3111303113221012-2000211130030321-1300303230333130)
- custom_storage_config.storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-1033230000220103-0101133321200113-3202203322110222-3322301201030112-0103122100303322-1112012133102310-1130203330323022-2310221233001320"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-2032103202231201-0020100020311202-1002003001302222-0201002110021003-1213103320033022-3110212203023102-0201231102213332-2213131201131030"></a>

## Direct properties — clear_secret_info / 331312033333 / 3

<a id="canonical-3022300221211321-0303000123331222-2110001310001221-3021030033022103-0332220031320211-0310120313031031-0111021110231330-0311021022013000"></a>

<a id="canonical-1322222000321110-3201011103321023-3231112032221211-3210110321232202-0103213111013030-1122111003210203-0302021230230033-3301002220111231"></a>

## provider_ref property — clear_secret_info / 331312033333 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0312332130033002-1133033010023231-1032001110020132-1203012030120303-0223331330003121-3231330203003032-3301121023310120-0103100302203100"></a>

<a id="canonical-3023000103210100-2310103103310310-0311330212102231-2321102020312233-1300220121300311-2201323221330031-1302223330233310-3233013130301010"></a>

## URL property — clear_secret_info / 331312033333 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0313213011233001-1020233113103300-1132023111321203-2212231200311133-1021031323122013-1302211203220210-3320022100122211-1020222120023133"></a>

## Next pages — clear_secret_info / 331312033333 / 6

- [custom_storage_config.storage_device_list.storage_devices.hpe_storage.password](data-sources--voltstack_site--reference--group-006.md#canonical-1322030132122002-2231302102210211-2121133010323132-1203111123113000-2233311303011303-3111303113221012-2000211130030321-1300303230333130)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010310330003233-3110021312222002-2212230112322213-2231110210300020-3100322011120002-2310122313222120-0300131033031102-3222133330121303"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident — netapp_trident / 231311111101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident

<a id="canonical-0022113233230111-1001211223123102-2103001131023001-1110222103231331-1132223133013111-0210313202100333-0021320101033222-3303023323131230"></a>

Type: `"single"`. Computed.

Device configuration for NetApp Trident Storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-backend_choice": "[\"netapp_backend_ontap_nas\",\"netapp_backend_ontap_san\"]"
}
```

<a id="canonical-1312210310222022-3133231221001310-1202011133013231-0321111022223301-3022320233201033-3210001110232102-3031102022333200-3200311203012110"></a>

## Direct properties — netapp_trident / 231311111101 / 3

- [netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121): complete subsection reference.

- [netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103): complete subsection reference.

<a id="canonical-2202223323110213-1312322102132202-0212303303231311-1301010130001200-0322331113223200-2330101332222122-1312303021300013-2301003132312133"></a>

## Next pages — netapp_trident / 231311111101 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312011000131232-2222113012130202-2211301122120023-2211002033311032-0331232021332130-3013313131133333-3102223313033231-3111120313021331"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas — netapp_backend_ontap_nas / 210011203002 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-1120232211022313-3130031333030033-3121231112211133-1303330011112302-1320323211223321-0201100222010211-1323330330300322-0000203313200311"></a>

Type: `"single"`. Computed.

Configuration of storage backend for NetApp ONTAP NAS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

<a id="canonical-3231330221310332-0121220112210123-2310131301332320-3002002020003222-1001100011033201-1221331310212221-1220003011233023-1331101332120113"></a>

## Direct properties — netapp_backend_ontap_nas / 210011203002 / 3

- [auto_export_cidrs](data-sources--voltstack_site--reference--group-006.md#canonical-3023110102133001-1001102011231011-3321021103112213-3300113033231022-3000123123032031-1321101002032022-1221002103100310-3101300303031302): complete subsection reference.

<a id="canonical-1331002302320122-1222132323001131-2030122031033320-1313213113313012-0313122211200203-0031230313031110-3300300020202133-0321121023321333"></a>

<a id="canonical-2321310111120233-3002300021012102-0331020103312023-3303102123223021-3131312110321332-0022312230203303-2010103300213310-0202323112002320"></a>

## auto_export_policy property — netapp_backend_ontap_nas / 210011203002 / 4

Type: `"bool"`. Computed.

Policy configuration for this feature.

Upstream description:

Enable automatic export policy creation and updating.

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

<a id="canonical-1120031121113021-3210010212021021-3013033130012131-1022030203012100-2133301010003131-2230222303102221-0132330210102300-2023130032003200"></a>

<a id="canonical-2211010302010322-2123132030221130-3032020000330332-2123030330003333-3230231021021220-2201201101111131-2311122001210030-1003130033112211"></a>

## backend_name property — netapp_backend_ontap_nas / 210011203002 / 5

Type: `"string"`. Computed.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Upstream description:

Configuration of Backend Name. Driver is name + "\_" + dataLIF.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 50,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3032121201302031-1022223213202112-2103332130101303-2110102312323110-0020300023320030-0201333133113012-3130330322112301-1213200002121100"></a>

<a id="canonical-1113220220120021-1130132233121012-0321213113202110-2103103213112123-0232322201313203-1220121220200313-1203231123210112-3312102302222001"></a>

## client_certificate property — netapp_backend_ontap_nas / 210011203002 / 6

Type: `"string"`. Computed.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-3033220200230020-1210023212323100-1002221121233220-1100211111133031-1301312313020200-0202312331112100-3022013031103013-1310332110301100): complete subsection reference.

<a id="canonical-1010100321121331-1223001322012122-0231120130132031-1013223310031302-2012301103112321-1122120303020210-3103031110032302-1322211322012030"></a>

<a id="canonical-2330320023203110-0020212210203310-2210213131222113-2330223011010012-3333112010200130-0001211313221111-3310202303022331-2111232333132211"></a>

## data_lif_dns_name property — netapp_backend_ontap_nas / 210011203002 / 7

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1302230233311031-3332101011003121-3001310111023301-3320333013013313-3301332133110310-0103111013331222-3003101333321210-3112033221201122"></a>

<a id="canonical-2312032022003322-3011100100033212-1000213210010113-1100322101320310-2302223221320321-1003312113210322-3020031302001121-0220232100000210"></a>

## data_lif_ip property — netapp_backend_ontap_nas / 210011203002 / 8

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

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

<a id="canonical-0233031101112100-1121302021100023-1311003010330020-1233010201310213-0233201111230233-0003230001012100-3211111220030132-3001001022210313"></a>

<a id="canonical-2212210301130231-1312221200103130-3110130322321031-3000122120311133-2323130313023003-2302103101233333-0031321330002333-1322133300100113"></a>

## labels property — netapp_backend_ontap_nas / 210011203002 / 9

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

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
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2131223032123311-3301121202223133-1002110201122332-1111111333311032-0122233223100213-2021130311320331-2032130221223111-1313103313303021"></a>

<a id="canonical-1322123331211300-1312222202020001-1302131213102020-3002122230310021-1313122002331321-0233200221212112-0230232113331030-1011223113220002"></a>

## limit_aggregate_usage property — netapp_backend_ontap_nas / 210011203002 / 10

Type: `"string"`. Computed.

Fail provisioning if usage is above this percentage. Not enforced by default.

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
  }
}
```

<a id="canonical-2322022310323013-0323121101202220-1100212113001321-3202320012301111-2130303020010003-1012312213103301-1303332210121311-3001210333121210"></a>

<a id="canonical-2021020332000103-2133221300020133-2231200131001202-1213200131320023-3033131303220000-3321213303121102-1323033221021110-1323111210323201"></a>

## limit_volume_size property — netapp_backend_ontap_nas / 210011203002 / 11

Type: `"string"`. Computed.

Fail provisioning if requested volume size is above this value. Not enforced by default.

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
  }
}
```

<a id="canonical-0330311320030211-3322330102321112-3102011110310012-1323132211210000-2120330013203021-0322131013003331-3230020221132132-2001022221121212"></a>

<a id="canonical-3130201202013022-1300223121312001-1120332010112113-3000300210201010-3210132322110202-0123202130312023-0201312011331222-2123030012013312"></a>

## management_lif_dns_name property — netapp_backend_ontap_nas / 210011203002 / 12

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2022111031212101-0003121032203303-1322323200300223-0132121023312101-2313021003323311-2210212030231330-1200013032202003-2000201021101331"></a>

<a id="canonical-3203330032320232-3103231200023202-1333002001231132-3221013200210310-1231320321123013-0033030102101103-2000112113322323-1031020231200302"></a>

## management_lif_ip property — netapp_backend_ontap_nas / 210011203002 / 13

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

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

<a id="canonical-0010202321003330-3233132013212112-1313003222230121-0221133231113121-1123323233213002-2313211110312023-1022013010002321-1232013222312313"></a>

<a id="canonical-0022102301303103-0120230121031323-3121022002033210-2032323312232103-0000002332232130-3323123000222131-2100333311312110-1220311020223232"></a>

## nfs_mount_options property — netapp_backend_ontap_nas / 210011203002 / 14

Type: `"string"`. Computed.

Comma-separated list of NFS mount OPTIONS. Not enforced by default.

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
  }
}
```

- [password](data-sources--voltstack_site--reference--group-006.md#canonical-1002310022110011-2103123221130210-3333002311233120-1212302011123111-0301200011000101-3333030303011301-1112030202220022-0321120301201210): complete subsection reference.

<a id="canonical-2211003101112310-1320320020311230-3122303200122032-3310000323022103-2012012033302200-1300203001003301-1012021301001111-2131312330303223"></a>

<a id="canonical-3120111002013213-3120001221220002-1201212301133033-2012331201122321-0333321210300311-3323201303231111-2311131330230103-2132113300000003"></a>

## region property — netapp_backend_ontap_nas / 210011203002 / 15

Type: `"string"`. Computed.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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
  }
}
```

- [storage](data-sources--voltstack_site--reference--group-006.md#canonical-2311303031330130-3033222013211211-0323033223223103-2032323002133220-3111032223203030-0022222022100212-0130111323101220-2211301031011101): complete subsection reference.

<a id="canonical-3210012310033311-1200100310100222-0230232023121100-1313202103333112-3101333011133131-2322133211221132-3221102313220100-0132120010022020"></a>

<a id="canonical-1323111311211013-0112232221132313-3223003133100120-3013122112312233-3031221303122112-2212211200012223-2210101313322102-3202323223333122"></a>

## storage_driver_name property — netapp_backend_ontap_nas / 210011203002 / 16

Type: `"string"`. Computed.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-1303233013132231-1323333303222030-3311223231323132-3011230312313111-1311211101120333-3023131303103021-2122003201020122-3210101000111001"></a>

<a id="canonical-3310321233213323-2323203112233121-0012121122030130-3000323203230211-0202012320020012-3103112330210111-3002332112310332-1222112223310031"></a>

## storage_prefix property — netapp_backend_ontap_nas / 210011203002 / 17

Type: `"string"`. Computed.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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
  }
}
```

<a id="canonical-0110032200333323-1100230001211303-2001023223210031-3011333112001013-1132033110000032-3233330231021212-2222200012323233-0310020231030301"></a>

<a id="canonical-3310033032013001-2303321313013103-1231201123220322-1213113210132012-3212011011032330-2020111003202132-2312222233132001-2012120002010322"></a>

## svm property — netapp_backend_ontap_nas / 210011203002 / 18

Type: `"string"`. Computed.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1020132332112111-0212002113230030-3023322233133021-2223320013112030-1031013130023101-1111131113300201-2123232001302331-1032232012201130"></a>

<a id="canonical-0312320202212303-0201002321201201-2223101313212303-1111332133323012-0131223311131213-1232321013233100-1033011321210031-2020111000111021"></a>

## trusted_ca_certificate property — netapp_backend_ontap_nas / 210011203002 / 19

Type: `"string"`. Computed.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

<a id="canonical-2031020121111103-2031301031211110-0133122221123111-2211210011022331-3013123021212103-2230113221113133-2230331330030230-0313110300330312"></a>

<a id="canonical-0321102201310300-1031132102301230-1213111123302002-2012101233102102-2120021222221021-0203012301002023-0201200303202321-3012312102230212"></a>

## username property — netapp_backend_ontap_nas / 210011203002 / 20

Type: `"string"`. Computed.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-3302310011222212-2300320333332333-1213011212311032-0212030321122120-0200212122323331-1203113212003212-2213010220123323-1012221230210323): complete subsection reference.

<a id="canonical-2003323103302333-2010111002001033-1121030021021012-2101102022101321-0302320323210220-3312302122321202-2120111310312231-1220321232202132"></a>

## Next pages — netapp_backend_ontap_nas / 210011203002 / 21

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](data-sources--voltstack_site--reference--group-006.md#canonical-3023110102133001-1001102011231011-3321021103112213-3300113033231022-3000123123032031-1321101002032022-1221002103100310-3101300303031302)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-3033220200230020-1210023212323100-1002221121233220-1100211111133031-1301312313020200-0202312331112100-3022013031103013-1310332110301100)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-1002310022110011-2103123221130210-3333002311233120-1212302011123111-0301200011000101-3333030303011301-1112030202220022-0321120301201210)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--voltstack_site--reference--group-006.md#canonical-2311303031330130-3033222013211211-0323033223223103-2032323002133220-3111032223203030-0022222022100212-0130111323101220-2211301031011101)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-3302310011222212-2300320333332333-1213011212311032-0212030321122120-0200212122323331-1203113212003212-2213010220123323-1012221230210323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3023110102133001-1001102011231011-3321021103112213-3300113033231022-3000123123032031-1321101002032022-1221002103100310-3101300303031302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213112002332220-0023010320202311-2102303032213120-0320011300012213-2321310032023032-1113010232330200-0033322323020231-0312122213113213"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs — auto_export_cidrs / 131012300112 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-3330300323303032-0231031000201233-1212011232213123-1223000100321112-3032201312110003-1011232010201000-2331203212011111-1213012010232021"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-0120021321010301-1231220121123322-2310120312102222-3023121110103112-1212213231223010-0223323300202232-0012313313030302-1311122232002123"></a>

## Direct properties — auto_export_cidrs / 131012300112 / 3

<a id="canonical-3320021320031120-3221033201313212-1033300333212000-2202222131111030-3011030201201232-1322313031132231-1103213003223200-2032310202122030"></a>

<a id="canonical-2311303021022211-0320223033011121-0032303023003313-2031102021122132-2122011112120032-3022130001213032-3010203220201331-0200302033311111"></a>

## prefixes property — auto_export_cidrs / 131012300112 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3103032000201012-3201022113233002-3211012032331032-0333132303000312-2233200213101313-3110233130201020-2130113100103332-0211331212010010"></a>

## Next pages — auto_export_cidrs / 131012300112 / 5

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3033220200230020-1210023212323100-1002221121233220-1100211111133031-1301312313020200-0202312331112100-3022013031103013-1310332110301100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232030301030133-2310332321302021-0000222230132011-1003203203010030-1230203133101033-3202031322303233-3220133223102003-3221133120303033"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key — client_private_key / 133301312003 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-1112003030121331-0223210010223200-2330002213033301-0233231330003323-1323032333030033-0131103212223101-3123300022222020-0022002300332232"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-1230133213031030-3002102321313003-0212012122230202-2030320132330133-0123102333232022-3121003332013220-3112302011130022-3331323231011211"></a>

## Direct properties — client_private_key / 133301312003 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-3113121131321201-0331302001103311-2311230011100010-3011300231020112-3011033001031310-1201220321020231-3120113011332210-0130202222223130): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-2212301320212020-1110331332031201-2122113323202022-2333121231210311-2120103022211031-2222230331103100-0123120323101310-2320110330033120): complete subsection reference.

<a id="canonical-2031113131313223-0023001200200012-2311220103122033-1221213131013000-0012313131002313-0313023201201001-0310023122311100-1112311231301213"></a>

## Next pages — client_private_key / 133301312003 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-3113121131321201-0331302001103311-2311230011100010-3011300231020112-3011033001031310-1201220321020231-3120113011332210-0130202222223130)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-2212301320212020-1110331332031201-2122113323202022-2333121231210311-2120103022211031-2222230331103100-0123120323101310-2320110330033120)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3113121131321201-0331302001103311-2311230011100010-3011300231020112-3011033001031310-1201220321020231-3120113011332210-0130202222223130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330103020220312-3200121002213323-3133211013102121-0103212002231002-1330223312001323-2112010020323033-1032212132003320-2001202322220323"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info — blindfold_secret_info / 130222212311 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-3033220200230020-1210023212323100-1002221121233220-1100211111133031-1301312313020200-0202312331112100-3022013031103013-1310332110301100)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-3310121310303303-1210303111222330-1322131220220023-3111132122010122-1333003203112100-3211222133101201-2133313201130322-1201002031102233"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-1132303331000111-3013230110023320-3133332210113210-0121233000220301-0330210101100123-0230100330300303-3030232012233310-2203303102131222"></a>

## Direct properties — blindfold_secret_info / 130222212311 / 3

<a id="canonical-1302202112011012-2012233221332200-1131102203201301-3230100103231231-2311032311313301-1131220102002331-0002210332103300-0013033102033031"></a>

<a id="canonical-3030100120101023-0203112312220112-2122121110001123-2030313102233313-2013000302332122-3202113312221111-2202100102023201-2013132113221302"></a>

## decryption_provider property — blindfold_secret_info / 130222212311 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
  }
}
```

<a id="canonical-0313232332223332-3000311032121001-0320201011010022-3331321303032232-2300113232322231-0233211130212203-1311023103023000-3232232133102130"></a>

<a id="canonical-0202022322102210-2103032331322110-0100331301200200-0121322100111133-3103023113023111-2212222331121113-1231010032233201-2203002021310022"></a>

## location property — blindfold_secret_info / 130222212311 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2122031121010012-3210221232131333-1222222300123000-0013212032323233-0133032210022320-2322310001221231-2332331301323300-2013201313221121"></a>

<a id="canonical-3222101023010302-0233210223120031-3003110020102130-3321202211230022-1210202333131121-2031321111221331-0133322331130113-1300201232220021"></a>

## store_provider property — blindfold_secret_info / 130222212311 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
  }
}
```

<a id="canonical-2223133011332230-1121110313320030-1330003232213101-0212321202212012-2121213130020212-0101222101131032-1210202333203211-1000122301023211"></a>

## Next pages — blindfold_secret_info / 130222212311 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-3033220200230020-1210023212323100-1002221121233220-1100211111133031-1301312313020200-0202312331112100-3022013031103013-1310332110301100)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2212301320212020-1110331332031201-2122113323202022-2333121231210311-2120103022211031-2222230331103100-0123120323101310-2320110330033120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331022222321231-0132301210101112-1111230120300333-3230100300303010-1103012233110300-1103232132302113-2120230120323132-2211332232101100"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info — clear_secret_info / 310030030330 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-3033220200230020-1210023212323100-1002221121233220-1100211111133031-1301312313020200-0202312331112100-3022013031103013-1310332110301100)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-0011112110133123-1223210023020100-2212020202211033-0132133233002332-0322002222110211-0012201213221113-2032302123310122-2031220322332302"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-2331002032332133-3012131230333020-3030330011333002-1302232101201020-0002033200123012-1230123332312102-0302203123130300-1330121311220011"></a>

## Direct properties — clear_secret_info / 310030030330 / 3

<a id="canonical-0310100012331110-3220132301021213-2203122212002001-2033322033121312-1203300201000322-3123301023312023-0301212310210301-0020131111121322"></a>

<a id="canonical-0220230303031111-2130323110212202-3232211323201023-3112211201120132-3201333321100003-2332210113121102-0300002101100312-2233330021200030"></a>

## provider_ref property — clear_secret_info / 310030030330 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3120302321302111-2203303310003310-2301033102121320-0123223200120222-2113320020002322-0032201000013030-1310030312133331-1303033003303000"></a>

<a id="canonical-3331021203233310-3303322131032212-0330301312311030-2022032322011232-0030131211103011-0001131102121010-1113231032012000-3211102323013311"></a>

## URL property — clear_secret_info / 310030030330 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2001211230003120-1110331201320011-3223001031312003-3330322300120232-0022322202200011-1122020130130233-0332122202102303-1220311113001121"></a>

## Next pages — clear_secret_info / 310030030330 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-3033220200230020-1210023212323100-1002221121233220-1100211111133031-1301312313020200-0202312331112100-3022013031103013-1310332110301100)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1002310022110011-2103123221130210-3333002311233120-1212302011123111-0301200011000101-3333030303011301-1112030202220022-0321120301201210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031301131010303-1113120020302302-0131300132233031-1112123321313131-2120230210333331-3033320322311121-0203013211133210-2031032201333210"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password — password / 202321213030 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-2030130100132323-3030230331103333-2203232320313110-2312132122012210-3221221203011100-3102300122030330-3332031123301201-1023210323033113"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-3120103213031220-1131211202332203-1123130323112122-3012113310332331-3031012130213113-0130113201012022-3322331110302101-0032321120312321"></a>

## Direct properties — password / 202321213030 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-3021323132123321-1010032100230300-3331011321133320-2002202031322131-0013011020101233-2012110232333112-2012100011301020-0323112021330333): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-2333223221321303-3013003200212322-1132001210112203-0231203322112311-1313230112111132-2001303022320222-1113000010121320-3111301031332322): complete subsection reference.

<a id="canonical-3020221011133323-3013011010000030-3333300231033133-1333133201112210-1033231301202232-1231201012301201-0113003011120031-1033202001111001"></a>

## Next pages — password / 202321213030 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-3021323132123321-1010032100230300-3331011321133320-2002202031322131-0013011020101233-2012110232333112-2012100011301020-0323112021330333)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-2333223221321303-3013003200212322-1132001210112203-0231203322112311-1313230112111132-2001303022320222-1113000010121320-3111301031332322)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3021323132123321-1010032100230300-3331011321133320-2002202031322131-0013011020101233-2012110232333112-2012100011301020-0323112021330333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311113110122111-1101232200320200-3013013301232313-0323002321123022-1300131332223133-1330200111202332-2201211210213300-0020001010010031"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info — blindfold_secret_info / 022100333223 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-1002310022110011-2103123221130210-3333002311233120-1212302011123111-0301200011000101-3333030303011301-1112030202220022-0321120301201210)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-3011303110221103-0330312233012002-3101011220221122-2002112213001211-1200132030023002-0312330210212201-1222013010000302-3120312002100031"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-0232200320302301-0300101212002031-0233002032132033-3120232213222011-3013222222310031-1030332202133030-3021221320233233-1322310020221321"></a>

## Direct properties — blindfold_secret_info / 022100333223 / 3

<a id="canonical-2030332122131231-2211211021300010-1110323020001020-0130220223021033-3122013203211231-2310022020230113-2130313021332123-2100321020003003"></a>

<a id="canonical-2322232321203111-2233103002023022-3102331011302021-1122332332200312-3131211023110201-3322012331233210-2222023321211002-0022301303310210"></a>

## decryption_provider property — blindfold_secret_info / 022100333223 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
  }
}
```

<a id="canonical-1333133220233300-1111023200011201-1302011212022102-2113331333130323-0331012111003010-1101121302301131-2030220033310222-1211010212111311"></a>

<a id="canonical-0202231302232203-1131103033213021-1020030300013231-2223102003333221-0200200311023322-1110232010100211-3021132211201322-1133330100012111"></a>

## location property — blindfold_secret_info / 022100333223 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3103112131230110-2223320232013131-2212200132333330-1001223322033233-0112333000222110-1012201112232301-2001302032322001-1320013301220001"></a>

<a id="canonical-1031321313102020-0131131111022302-0221101033000200-1032131011302331-2032313223131332-3200032331211232-3002112301130312-2121221002132000"></a>

## store_provider property — blindfold_secret_info / 022100333223 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
  }
}
```

<a id="canonical-0023201011133101-3133203032000323-0033133320230012-2020203031103232-3230221122320033-2203301231103000-0300313230313232-0122002012301310"></a>

## Next pages — blindfold_secret_info / 022100333223 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-1002310022110011-2103123221130210-3333002311233120-1212302011123111-0301200011000101-3333030303011301-1112030202220022-0321120301201210)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2333223221321303-3013003200212322-1132001210112203-0231203322112311-1313230112111132-2001303022320222-1113000010121320-3111301031332322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133301322211223-0303311200123230-2113030232023210-1001122300332020-0113212030310301-0323321030312032-2222220023233101-3002033232331321"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info — clear_secret_info / 130000311300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-1002310022110011-2103123221130210-3333002311233120-1212302011123111-0301200011000101-3333030303011301-1112030202220022-0321120301201210)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-0202001003131302-1012011230233300-2213213132300302-0030233013321322-1123303110332203-1213131022303232-3302302310011311-0202231333132133"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-0132132322223233-2031010300122301-0011202303233013-1111020312011221-0010120322331113-1301133221210331-2330300303322320-1022102323231303"></a>

## Direct properties — clear_secret_info / 130000311300 / 3

<a id="canonical-1030231032003322-3001230123013311-3013302032103330-1000233113201121-0220100212200332-2200022101203230-1330300030103213-2101002132130323"></a>

<a id="canonical-3033212130133001-2132011010330210-1213132102323130-1220102212000020-3013331133300003-0203113322320121-1022123021233000-2100321100030122"></a>

## provider_ref property — clear_secret_info / 130000311300 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0312123322132133-3020120202312022-2101232121231000-3330222222312331-1012033122132300-0022302222003011-2321012312123012-2120302021021133"></a>

<a id="canonical-2020022303001200-0130022200003313-3310223203102330-0110021020200321-2130213223231222-0133032310030131-1331012100310323-0331310012112330"></a>

## URL property — clear_secret_info / 130000311300 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0030112120130021-0312223333210200-0113233210001110-3301230031012320-2202311311220331-3101200302232213-1213120123211130-1003203020001221"></a>

## Next pages — clear_secret_info / 130000311300 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--voltstack_site--reference--group-006.md#canonical-1002310022110011-2103123221130210-3333002311233120-1212302011123111-0301200011000101-3333030303011301-1112030202220022-0321120301201210)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2311303031330130-3033222013211211-0323033223223103-2032323002133220-3111032223203030-0022222022100212-0130111323101220-2211301031011101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013210102032301-3301110002203100-3201033300333113-0222020001321313-1030131300010133-0202022320233211-1030021200120110-0013102320132020"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage — storage / 322103101131 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-2232231110103130-3310232122211120-3200123123332030-3220022222131330-1202233022132103-3332111212223232-0211030111210010-3131300233321121"></a>

Type: `"list"`. Computed.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1201030231000112-0301002010102120-1312102012112213-3330022302333223-0123130130333030-0233002012333120-2011023112232220-0100203110012121"></a>

## Direct properties — storage / 322103101131 / 3

<a id="canonical-3003020230121221-1211032211200330-3232332333300310-2031002201012221-2102121232331102-1211233111202100-3110120323130223-1300301121232203"></a>

<a id="canonical-0223020203102330-2231032022032000-0321323012332313-0322313122223123-0012322131003212-3322022311233201-1120232330012213-3103022103002032"></a>

## labels property — storage / 322103101131 / 4

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

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
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-2213113020020123-0222221331302120-1230132212213301-2232000123212223-0222311312212211-2101010212112123-0310301102100023-2303023123312310): complete subsection reference.

<a id="canonical-0132233121132213-0321110313030010-1102202102002232-1212311202212210-2112013021323211-3011312131303221-0230213120010213-1031303012101010"></a>

<a id="canonical-1112231312120210-0232010211112333-3021001230021233-1010311123232212-2100211003322100-0120212230333003-1310332123002133-1102201220033331"></a>

## zone property — storage / 322103101131 / 5

Type: `"string"`. Computed.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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
  }
}
```

<a id="canonical-2022311201323122-0333230231300211-2311000233311100-3021112003332131-3331310210311011-3020001203102002-3003130202211233-0032020230332321"></a>

## Next pages — storage / 322103101131 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-2213113020020123-0222221331302120-1230132212213301-2232000123212223-0222311312212211-2101010212112123-0310301102100023-2303023123312310)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2213113020020123-0222221331302120-1230132212213301-2232000123212223-0222311312212211-2101010212112123-0310301102100023-2303023123312310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033202130303132-0110310000213210-1213200122010320-3132023132221231-3112330221231330-1131103002021010-0303001231121013-3112133310012033"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults — volume_defaults / 111221013312 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--voltstack_site--reference--group-006.md#canonical-2311303031330130-3033222013211211-0323033223223103-2032323002133220-3111032223203030-0022222022100212-0130111323101220-2211301031011101)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-2202122323211000-3323123320133233-1121030213022321-0211333122310323-1122220000200101-0232013303120210-1333022201223300-3211133202331210"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

<a id="canonical-2322220302121123-1230111302311121-3002032003222212-0102112210201112-1323302312331022-3022313333023002-1332311232023102-0313312030201331"></a>

## Direct properties — volume_defaults / 111221013312 / 3

<a id="canonical-1233202013100113-0101331210100122-2301331213202012-3210020112120330-3220202103212330-0320021030033030-2222022313301220-0203033322001210"></a>

<a id="canonical-0121303233311320-0212321112230001-3030323020320131-0133012020221120-3233100102030222-3012002201302012-1330021220030303-0130331132103320"></a>

## adaptive_qos_policy property — volume_defaults / 111221013312 / 4

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1201200231013202-2001211030330203-2102222200113022-2133203321013130-2301000310331211-2213131133212101-3130033132211111-1003130030021232"></a>

<a id="canonical-2030303003003301-1020223120232011-0022222130213310-1033132112030030-1302021312203330-1130200132023010-3330023211321201-1011332221233213"></a>

## encryption property — volume_defaults / 111221013312 / 5

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

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

<a id="canonical-3300121112021212-0212232123223232-2102010310220201-0232212023010201-1223231203232121-2221030323212222-0102021122330112-2310311210213203"></a>

<a id="canonical-3221130220121013-3131001232010020-1210132030102033-3121213323020002-0132331211121303-2011120213212003-0111302021013303-3112231001020133"></a>

## export_policy property — volume_defaults / 111221013312 / 6

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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
  }
}
```

- [no_qos](data-sources--voltstack_site--reference--group-006.md#canonical-0011130321331333-3233112102123003-0012121131032212-3013010233010202-1101332132213310-1310220133301111-0000100112312223-3132320301031212): complete subsection reference.

<a id="canonical-2210021120113010-0122032011211121-3113221212122033-0222111021302101-0330302210322023-3332303131202133-2031100321212302-1212330302223112"></a>

<a id="canonical-1331002233210301-1000222001301203-1021213222211322-2031103220322110-3103000331013320-2033131121001100-3230333003032222-3212122101101013"></a>

## qos_policy property — volume_defaults / 111221013312 / 7

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3033003021031302-0233011210301211-0030112210000122-3303311121032201-0313212003212300-3230131113231002-2210313000211121-2330321232211133"></a>

<a id="canonical-0002213300313330-2000231003323001-1320332113031023-3203330210300101-3320213021210103-2120103022100102-1303201033321203-0113231203300023"></a>

## security_style property — volume_defaults / 111221013312 / 8

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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
  }
}
```

<a id="canonical-1210233131222220-2101100302121111-1111223321130030-2300013032320032-0112033030230132-2132000323301131-3233131313030100-0210001212330122"></a>

<a id="canonical-2102120222002022-1020232123132200-3123321122303311-3001133220333220-1100011000000212-3101103131113321-2222110113133300-1300123221303031"></a>

## snapshot_dir property — volume_defaults / 111221013312 / 9

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

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

<a id="canonical-0311201130232312-2030102010300130-0331102310123033-2232212202030100-2020022032020321-0122021020022032-0201000210131101-3333330232113010"></a>

<a id="canonical-2322010203233100-2101010310122023-1021132103233010-1002221032211120-2233312121123233-1131013300202113-1233001300332013-2322331323323332"></a>

## snapshot_policy property — volume_defaults / 111221013312 / 10

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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
  }
}
```

<a id="canonical-0311011201102123-1223332303032130-2300123120233100-3022001200021010-3133320023110031-2232232113103023-1010222221222103-3322232122201231"></a>

<a id="canonical-2021332232133132-0021130230303032-0130233323212321-2220321011221100-3332323232331200-1102030320223000-2130112323201010-2101303302133032"></a>

## snapshot_reserve property — volume_defaults / 111221013312 / 11

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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
  }
}
```

<a id="canonical-1130000032212312-3020012212123101-0112101303113212-0222011032202232-2212221231001300-2121023132221110-1221303121022123-3321322232010211"></a>

<a id="canonical-0031221303100013-0213230301222122-2303330333101133-3202102103322202-3023210002002033-3213031033121232-0103100212132002-1321222120321100"></a>

## space_reserve property — volume_defaults / 111221013312 / 12

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-1210123101100110-3020230332110231-0103102001101312-0001133100011202-0131300113303031-1132000102103321-1231011311333233-0223020133112010"></a>

<a id="canonical-0132000301002102-0221213301001121-3230023131010221-0201013020120311-1210000332103313-1232333032302210-1213221203332102-3112000021332320"></a>

## split_on_clone property — volume_defaults / 111221013312 / 13

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

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

<a id="canonical-0311210030112013-3213211123301103-1020123212320001-3223131101010100-2203033230220313-2130202213311212-1133333200233311-2003320233230131"></a>

<a id="canonical-3031020030200121-3031100220211002-2133200121302212-1121113210210102-1030213132303012-3110322130110330-2201110322211133-1132122210323120"></a>

## tiering_policy property — volume_defaults / 111221013312 / 14

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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
  }
}
```

<a id="canonical-1113133223321201-0331312030213131-1302213112311111-0300130121200110-2332221302110221-3032310223210112-3103232030222120-0211102222001213"></a>

<a id="canonical-0131111020302101-0323320030023131-1332110011330221-3200323111210101-3333331220213100-0113133221122101-1200310122331011-0321032233012233"></a>

## unix_permissions property — volume_defaults / 111221013312 / 15

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

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

<a id="canonical-2132322001212131-0223130312112200-0212103013133212-3230001123110333-3223202310221130-2003312203312133-0213213220112121-3222022022121201"></a>

## Next pages — volume_defaults / 111221013312 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](data-sources--voltstack_site--reference--group-006.md#canonical-0011130321331333-3233112102123003-0012121131032212-3013010233010202-1101332132213310-1310220133301111-0000100112312223-3132320301031212)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--voltstack_site--reference--group-006.md#canonical-2311303031330130-3033222013211211-0323033223223103-2032323002133220-3111032223203030-0022222022100212-0130111323101220-2211301031011101)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0011130321331333-3233112102123003-0012121131032212-3013010233010202-1101332132213310-1310220133301111-0000100112312223-3132320301031212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113220332011120-1030122222220011-0320210221102223-1210220313212110-0131001233230333-3030222022100310-3313030002210020-2201232220333331"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos — no_qos / 230203300000 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--voltstack_site--reference--group-006.md#canonical-2311303031330130-3033222013211211-0323033223223103-2032323002133220-3111032223203030-0022222022100212-0130111323101220-2211301031011101)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-2213113020020123-0222221331302120-1230132212213301-2232000123212223-0222311312212211-2101010212112123-0310301102100023-2303023123312310)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-1122002330301231-3332030032122223-3010301002103302-2120023112110333-0332303001300033-0320220120110103-3030200102213132-3102303300231322"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0022120232022023-1231013211112333-1213301222011321-3123210323013203-3122011312202332-3301103211323202-1111301110032331-0012333021303001"></a>

## Direct properties — no_qos / 230203300000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001021032203132-2101223000212232-2102112223222310-0032100101102222-0333103102222000-0233303330200133-2312132222202113-3202302330230021"></a>

## Next pages — no_qos / 230203300000 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-2213113020020123-0222221331302120-1230132212213301-2232000123212223-0222311312212211-2101010212112123-0310301102100023-2303023123312310)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3302310011222212-2300320333332333-1213011212311032-0212030321122120-0200212122323331-1203113212003212-2213010220123323-1012221230210323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220301111001300-2021133111102101-2202213101321000-2301223302021121-2312023220313320-2122122113311113-2010330001020221-3313230300020121"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults — volume_defaults / 001133203032 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-2013003120030312-2020320020311231-1012303121303303-3000103220201222-0202021310121313-1233110222123120-3211031320122132-0111321202300313"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

<a id="canonical-3310320322211321-1202303022031112-0320003003303211-0221031330020312-2200201201031320-1222121321233330-3033313121211312-1200022133200110"></a>

## Direct properties — volume_defaults / 001133203032 / 3

<a id="canonical-0332012323012113-3331331303301111-2312030003000303-2103022101131312-0211203013103221-2012113211211030-0211330002130123-0101332230031301"></a>

<a id="canonical-1112130300330213-1123002131333131-3211331223231032-1220132122310203-2200130031012100-3331120003110301-0202303302233023-1201333301121110"></a>

## adaptive_qos_policy property — volume_defaults / 001133203032 / 4

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2330012310331211-0333220333231131-0213222331212230-3011301112232011-1020313302023210-1131323133310201-0033023032130223-2321233331321212"></a>

<a id="canonical-3020232202111110-0100333230031200-3023222221123103-2021131112231012-2111233310211103-1302102310030323-3312111110231103-1312130311230123"></a>

## encryption property — volume_defaults / 001133203032 / 5

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

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

<a id="canonical-2330103300222332-3122101033213330-1302023230222032-0003111122323300-1013111330100312-0310322303000113-2113030121302220-0120201302133223"></a>

<a id="canonical-3220310323001121-0333222023231123-2001110122310113-0210321201020021-1102323303032120-3331310023222012-3100213213003030-3313212200030002"></a>

## export_policy property — volume_defaults / 001133203032 / 6

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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
  }
}
```

- [no_qos](data-sources--voltstack_site--reference--group-006.md#canonical-1323033212221013-0311330203300301-1010002223231322-3101113302120300-1030020110031010-3131311233000330-2233320223231133-3033033230332332): complete subsection reference.

<a id="canonical-0231220222020012-0130021010220330-0311033322333011-3132002133311311-1001033032112010-1331302312322302-1302201000232312-0232111210230011"></a>

<a id="canonical-3213000110030111-0202333001103121-1122303313202323-1330131323101333-3301320302212000-1013012321221111-3322202301032210-0132032103330220"></a>

## qos_policy property — volume_defaults / 001133203032 / 7

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0102233131031020-3333030200323302-1311033213321132-1021133220000322-2333103112023203-0311011301100310-1032213221212012-2122120231132102"></a>

<a id="canonical-3033023031121011-2133110332302201-1102320003021113-1303221121132022-0203130202003221-0220233110203013-2312033301101201-0033101013212131"></a>

## security_style property — volume_defaults / 001133203032 / 8

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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
  }
}
```

<a id="canonical-1313012020300223-0202333010210233-3331212010211322-0023332321112332-0320033331101231-3021113220101210-1200013100100123-1330210120222332"></a>

<a id="canonical-2120301233203012-0123113102220203-1311321313021211-1310100003022002-0323331032000121-0223330210013100-0222232130222011-0303200011130310"></a>

## snapshot_dir property — volume_defaults / 001133203032 / 9

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

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

<a id="canonical-0320010232311000-0210100110201233-2111013023003220-0220132130310321-2131003202202011-1203121130030122-0203113133213320-1011232232020022"></a>

<a id="canonical-3202233212023300-3210333022010202-0211012023200123-0102132001131312-0100330320133133-1120313120300231-0121031003322102-3223010030022121"></a>

## snapshot_policy property — volume_defaults / 001133203032 / 10

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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
  }
}
```

<a id="canonical-2322010013202111-3213302011131011-0323231022333203-3003202303220133-3303213300123200-0002200100132200-3213220012221100-3030311002131212"></a>

<a id="canonical-3220311033011230-2333303230133300-2321132000210110-1322123322323120-3220223223110322-2220112223330123-3022020211100303-2220220311312320"></a>

## snapshot_reserve property — volume_defaults / 001133203032 / 11

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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
  }
}
```

<a id="canonical-3112220123012101-1322022320210311-1320320223322103-0033012311112102-1230203233022033-2330311331331030-0130312201201321-2321103203103310"></a>

<a id="canonical-3011031001013221-1012303133203232-2331222302330123-1101230032032312-2010331202330201-1230232130230020-3033120311023113-2200221303011121"></a>

## space_reserve property — volume_defaults / 001133203032 / 12

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
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
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-0033321023332322-1032101222000001-1102233033002302-0023030202320030-2013022020333202-0102313122131202-0233023111000033-2120300230023210"></a>

<a id="canonical-0021321133110020-1313011102123011-3013000330030031-0012332222203321-0102212330300110-1021320110131100-0201212312220221-1001301121333013"></a>

## split_on_clone property — volume_defaults / 001133203032 / 13

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

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

<a id="canonical-3012213031002233-2321132102313332-3033101301332013-2013101202002020-2210002221113311-2301021021103111-3331100330123120-0222102001331321"></a>

<a id="canonical-3101332022130111-1020230311121303-0020320333002112-3301003033300331-3101003232202323-0203312331113032-0023031330231213-0032031002203133"></a>

## tiering_policy property — volume_defaults / 001133203032 / 14

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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
  }
}
```

<a id="canonical-1133100103031323-2133001301222203-3331320303013111-2133023231111323-1312033012300130-0213113110123332-0021233133213222-1311323122021031"></a>

<a id="canonical-2102210110032211-3311230011200102-3002211230331013-0011332211220202-3133220031200301-3322331123131330-2111101003223010-1120132013112013"></a>

## unix_permissions property — volume_defaults / 001133203032 / 15

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

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

<a id="canonical-2030000210022021-1203113333022001-1331131103331312-0322201101113202-2121022301011112-2232322023132020-0011112100122311-0333230320110012"></a>

## Next pages — volume_defaults / 001133203032 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](data-sources--voltstack_site--reference--group-006.md#canonical-1323033212221013-0311330203300301-1010002223231322-3101113302120300-1030020110031010-3131311233000330-2233320223231133-3033033230332332)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1323033212221013-0311330203300301-1010002223231322-3101113302120300-1030020110031010-3131311233000330-2233320223231133-3033033230332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310331331222212-0303210223123223-3013303122330313-0311211003210300-3303300222303320-3101331323310113-0202102102012201-3022030213032232"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos — no_qos / 122333332013 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--voltstack_site--reference--group-006.md#canonical-3120312100213201-1323131132100130-0323322310230110-3312231330122121-0011130122301332-2210321112032213-2233221020103233-0012223202032121)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-3302310011222212-2300320333332333-1213011212311032-0212030321122120-0200212122323331-1203113212003212-2213010220123323-1012221230210323)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos

<a id="canonical-3213013022231030-3103033103332033-0031201113132312-1103321313203013-0131021300133011-0010120101233200-1132323120222012-1320031333113221"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0213203321331022-0221100220011220-1222302311332210-1112000131221211-0113101201223211-0023333311131330-3202103312112322-3301221011110030"></a>

## Direct properties — no_qos / 122333332013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301030333330101-3132103300023220-3300301131012011-1212233233202101-2002220132123300-2221011100213233-1000030030012030-0223303211213132"></a>

## Next pages — no_qos / 122333332013 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--voltstack_site--reference--group-006.md#canonical-3302310011222212-2300320333332333-1213011212311032-0212030321122120-0200212122323331-1203113212003212-2213010220123323-1012221230210323)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313001213120222-1301030010121032-1130022211232102-2233112203013022-1110113301223211-0212200303230210-1213110300231101-1113032232223102"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san — netapp_backend_ontap_san / 001022031100 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="canonical-3022033300123200-0132011120311211-1212112222123300-0203033132101022-0022113013121021-1100003322310303-2322112130220113-3321102032312030"></a>

Type: `"single"`. Computed.

Configuration of storage backend for NetApp ONTAP SAN.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-chap_choice": "[\"no_chap\",\"use_chap\"]",
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

<a id="canonical-1320000132320012-2100031112321322-2120331020213032-2021002112021030-2112303022030322-0131101112302031-2132123212002233-2133031103012012"></a>

## Direct properties — netapp_backend_ontap_san / 001022031100 / 3

<a id="canonical-2031120012020313-1111232001320133-3123021021110301-1211221222101111-1321320020121300-0332203121203031-2213111022112112-2131301000010223"></a>

<a id="canonical-2200132022300212-0121223012233130-3032013333321102-1223223100110201-3320022030110231-1103022002310313-0202133002311323-1210313322003131"></a>

## client_certificate property — netapp_backend_ontap_san / 001022031100 / 4

Type: `"string"`. Computed.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0031221320012122-2002233213121101-3003200313031331-2212223302113100-2313302033333122-2032231231331013-0221000231001003-2110023331211220): complete subsection reference.

<a id="canonical-1320233031203300-2230112320023213-0031011113020212-2333220202130110-0001101302203310-2123210312300313-1332133110103333-0130121012111323"></a>

<a id="canonical-0323200220232000-2232120032330323-1132301221233213-1221330313222112-2330211211112002-1133330230212111-2303021120331233-1323312301302103"></a>

## data_lif_dns_name property — netapp_backend_ontap_san / 001022031100 / 5

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0013313130222223-1302002233221212-2132022123322320-1030301132120303-0103330312010333-0122113021010323-0212023113311223-3332213322212223"></a>

<a id="canonical-0301010031231113-0330100120022110-2211030201332232-0021231100130012-2030110101001020-1321232011123311-2131212223122031-1132222113112310"></a>

## data_lif_ip property — netapp_backend_ontap_san / 001022031100 / 6

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

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

<a id="canonical-2320003333212011-1231121131132103-3112211013020031-3101321112001222-2022313200321001-2102213012213101-0130201322200322-2232010033011203"></a>

<a id="canonical-2211132023113033-2202323320101021-3223233031013202-1320023011220113-0313213203311113-3331203021221212-2022031103222320-0331123213332123"></a>

## igroup_name property — netapp_backend_ontap_san / 001022031100 / 7

Type: `"string"`. Computed.

Name of the igroup for SAN volumes to use.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1023220320301233-1220233102222003-2103321211103103-1230000220032120-1031221023210030-3120313201310221-0031122101102112-0312230330203120"></a>

<a id="canonical-2200113011030312-2302203103010202-1113331000030230-1211201300031333-3322123032210111-0113211333132102-1332113001321011-0023332203331303"></a>

## labels property — netapp_backend_ontap_san / 001022031100 / 8

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

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
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3012221001031331-2232033312201113-1102030302133223-3302000323000121-0012022023223310-1032231221132032-3021032021020303-3022002110011013"></a>

<a id="canonical-0321312331032132-2330233201112310-1211300001033323-1101130112301323-0030333013103032-2221210323313131-3200232000301110-3013332011213101"></a>

## limit_aggregate_usage property — netapp_backend_ontap_san / 001022031100 / 9

Type: `"number"`. Computed.

Fail provisioning if usage is above this percentage. Not enforced by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-2210302321032123-3130212312023212-2212032320221112-3133301123333311-1232011330131320-2123131222000203-3200220213013133-2333312011031123"></a>

<a id="canonical-3213133121223021-0303101233323332-0333311100132032-1321232323321121-1110223023231111-1102122033103212-0111033223023313-3030302332330011"></a>

## limit_volume_size property — netapp_backend_ontap_san / 001022031100 / 10

Type: `"number"`. Computed.

Fail provisioning if requested volume size in GBi is above this value. Not enforced by default.

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

<a id="canonical-2012200003300132-2232110220213220-0213212033021333-1122330122002123-3331112323110101-2322132011002112-1202303133301123-3003202110230131"></a>

<a id="canonical-2230101302100003-2301321233301130-1102222133021213-1202111120123021-0032320120332233-2100111122013211-3011113100022321-1200002022130312"></a>

## management_lif_dns_name property — netapp_backend_ontap_san / 001022031100 / 11

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1222112332213330-0013100223020013-2001330300302113-1330230302322230-1222333312230311-2203333123031132-0220130031103023-3201100311013301"></a>

<a id="canonical-0022203200011121-0201100010033331-2330000011000033-1032133103023012-2310002011020321-0321312022122313-0003030322301203-1333031223300300"></a>

## management_lif_ip property — netapp_backend_ontap_san / 001022031100 / 12

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

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

- [no_chap](data-sources--voltstack_site--reference--group-007.md#canonical-2123020023031200-2333220022121210-1120020130313012-0313132120312202-2233312231113231-2122223330311101-1312011233231233-1310013123333331): complete subsection reference.

- [password](data-sources--voltstack_site--reference--group-007.md#canonical-2002230133022000-3011203023322232-2010110321231200-0013321303301320-2200010220303333-0123330233021231-1213023311113322-3001232310123311): complete subsection reference.

<a id="canonical-1032003333113013-3210230002113133-0331300132233023-3103203312112113-1012310102023133-3330221222003120-1131112211213130-3023022110202322"></a>

<a id="canonical-2003130131121101-2012222323301002-3131102313101020-1310203131120111-0232231120311121-2033223030132101-2202132001310202-3120313202320310"></a>

## region property — netapp_backend_ontap_san / 001022031100 / 13

Type: `"string"`. Computed.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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
  }
}
```

- [storage](data-sources--voltstack_site--reference--group-007.md#canonical-1321303001321312-0013123013303321-1200013023232333-2120313211303332-0233020303202230-0101001210010021-1033011211222210-2222201031023232): complete subsection reference.

<a id="canonical-2002213000201000-0302100113202011-3011030000302013-0112323031221220-1310331312113003-0233032123222122-1021031013303002-0232220032230011"></a>

<a id="canonical-1320332311101001-2003101031101200-1032003230111130-1010211233021220-3122021031020223-2303323133022020-3021020011130000-2232023303003101"></a>

## storage_driver_name property — netapp_backend_ontap_san / 001022031100 / 14

Type: `"string"`. Computed.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-0011023012123321-2122120331020311-2303100010021103-0321012311223321-2133130013110122-2233110000223123-2333300333000222-0201122203102120"></a>

<a id="canonical-2300223301102333-0303310212333031-1210000121122211-3123121013002100-0122220232220112-3012330211103131-3303332131021010-1103102103101020"></a>

## storage_prefix property — netapp_backend_ontap_san / 001022031100 / 15

Type: `"string"`. Computed.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 80,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 80,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2221032023133202-2232003220311323-0203200131301233-0003303112320013-0211312103023132-3220313211310002-2110121120010112-1233301212330131"></a>

<a id="canonical-1323101203220203-1013212310010210-3231320323013000-3132102110303231-3212010002311130-3013223333120130-2220113201131001-0103012000323330"></a>

## svm property — netapp_backend_ontap_san / 001022031100 / 16

Type: `"string"`. Computed.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2002311001101002-3033100302231010-3212102020002310-1231212202202121-0203330023123211-1300201323210202-1002020122132000-0231111001221132"></a>

<a id="canonical-0121300031133331-0013033231103100-1333231330111112-1002330213100021-0011222331133031-3022130212130230-0110012031033002-0103201102333011"></a>

## trusted_ca_certificate property — netapp_backend_ontap_san / 001022031100 / 17

Type: `"string"`. Computed.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221): complete subsection reference.

<a id="canonical-3003032010101313-2112011212331231-3310330302222132-0020211121103333-3310302331222312-1132211231333101-2102332030102010-2330213220301000"></a>

<a id="canonical-3213131000201233-0002133000301023-0032112131132010-3102233222000121-1201113211303100-3000230110323033-2021112111101323-1031331230302213"></a>

## username property — netapp_backend_ontap_san / 001022031100 / 18

Type: `"string"`. Computed.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-0010213121112030-2023010302202201-0032322030031231-2203123331233303-3331013031313301-2111303021000302-2003330020101200-0100231311031011): complete subsection reference.

<a id="canonical-3031321110102122-1200003133230011-0133220311121023-2002133001012002-3130121321333232-2233122233001223-1230132121131023-3333013300000131"></a>

## Next pages — netapp_backend_ontap_san / 001022031100 / 19

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0031221320012122-2002233213121101-3003200313031331-2212223302113100-2313302033333122-2032231231331013-0221000231001003-2110023331211220)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](data-sources--voltstack_site--reference--group-007.md#canonical-2123020023031200-2333220022121210-1120020130313012-0313132120312202-2233312231113231-2122223330311101-1312011233231233-1310013123333331)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-007.md#canonical-2002230133022000-3011203023322232-2010110321231200-0013321303301320-2200010220303333-0123330233021231-1213023311113322-3001232310123311)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--voltstack_site--reference--group-007.md#canonical-1321303001321312-0013123013303321-1200013023232333-2120313211303332-0233020303202230-0101001210010021-1033011211222210-2222201031023232)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-0010213121112030-2023010302202201-0032322030031231-2203123331233303-3331013031313301-2111303021000302-2003330020101200-0100231311031011)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0031221320012122-2002233213121101-3003200313031331-2212223302113100-2313302033333122-2032231231331013-0221000231001003-2110023331211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033101120323131-2120312320230320-2310233322021131-2321000212111203-3103133133122323-3112123333330001-2332302211201301-0003331220212030"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key — client_private_key / 120121330301 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-0222303201032310-0131133233310330-2031333031122333-2301203230033132-2013201230113132-3103232310111023-3220010222022033-3213132023313210"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-3200132223122313-3212023220313222-2210003300310000-3121223103020000-0313302331333301-1221013122221002-3202323233211213-1003033322202103"></a>

## Direct properties — client_private_key / 120121330301 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-1023323103301012-2313103212302323-2223102022332232-0113012100322333-1113300122112302-3020003211203221-3113130333302022-0200200103031200): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-0223330300130022-0012123333120000-1011000003112310-3110211020323233-3111311202233031-3310030221212333-2212202123032122-2010322330313223): complete subsection reference.

<a id="canonical-2031003030222221-3222101321331232-2002201312303022-0033001010023303-0010201303231011-1210012203002133-3110202220313122-0032121030311132"></a>

## Next pages — client_private_key / 120121330301 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-1023323103301012-2313103212302323-2223102022332232-0113012100322333-1113300122112302-3020003211203221-3113130333302022-0200200103031200)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](data-sources--voltstack_site--reference--group-006.md#canonical-0223330300130022-0012123333120000-1011000003112310-3110211020323233-3111311202233031-3310030221212333-2212202123032122-2010322330313223)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1023323103301012-2313103212302323-2223102022332232-0113012100322333-1113300122112302-3020003211203221-3113130333302022-0200200103031200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023001223211203-1333120201000303-0001013331220110-0131003122033321-3331012112233001-0033003000310310-3202033330133223-1301222012201232"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info — blindfold_secret_info / 130230311222 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0031221320012122-2002233213121101-3003200313031331-2212223302113100-2313302033333122-2032231231331013-0221000231001003-2110023331211220)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-1133030132032232-2203110312212102-0221303020001030-1101033231101201-3322033002030300-3200331110231311-2102331232200331-3310213131103032"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-0000110111031230-2232013213130320-3132312322300100-0212122312033133-0002201201222003-3032312130010132-3010232232300032-0022031332102321"></a>

## Direct properties — blindfold_secret_info / 130230311222 / 3

<a id="canonical-3201122132223123-1310331302112031-1012222201323111-3221333202233212-1210221310112300-0201330332201112-1000300212202022-3110103113101211"></a>

<a id="canonical-2020220332113222-0010203102310033-2131111001000321-1221213210122012-1133301030021023-2122202210302001-2333222133213021-0202322023322202"></a>

## decryption_provider property — blindfold_secret_info / 130230311222 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
  }
}
```

<a id="canonical-1221233120213213-0011301131312302-3101123310122103-3211330320211000-0231033113323020-1230331311313023-2101023311030331-3312000102322001"></a>

<a id="canonical-0100213332201332-3332020131021332-2231301133230203-1203032112231111-3203111322122310-0112021310012221-3211011200200211-3013322233233120"></a>

## location property — blindfold_secret_info / 130230311222 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1013012131221331-2221330003211103-2013110223320220-3000233303230001-1213013302131231-2011101203231213-2220323102232020-1121211030210121"></a>

<a id="canonical-3031013221100122-2100101013220112-2202320130001323-1320103123023323-0113010013331120-1232000010120302-1021033222100310-3013113103333132"></a>

## store_provider property — blindfold_secret_info / 130230311222 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
  }
}
```

<a id="canonical-2022331022023221-2022020221301211-3301332210010121-2033202301100030-3301212031333100-2011202220321213-1303223323100113-2303120032032203"></a>

## Next pages — blindfold_secret_info / 130230311222 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0031221320012122-2002233213121101-3003200313031331-2212223302113100-2313302033333122-2032231231331013-0221000231001003-2110023331211220)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0223330300130022-0012123333120000-1011000003112310-3110211020323233-3111311202233031-3310030221212333-2212202123032122-2010322330313223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210121201203222-3003213313033131-0102123013200031-0211033323223022-1221301021313032-1312110113333231-3311211101100001-1103313312323213"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info — clear_secret_info / 033023203021 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0031221320012122-2002233213121101-3003200313031331-2212223302113100-2313302033333122-2032231231331013-0221000231001003-2110023331211220)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-3333131322330023-3231210120302333-3000321010211301-0201232120312203-1121021000122223-2230212220010133-1030201312003311-2213222322010001"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-0001110000103202-0002032323032233-0021323203033023-1221112312322033-3123011301301101-3113212320133123-0321020303321003-1031310032023111"></a>

## Direct properties — clear_secret_info / 033023203021 / 3

<a id="canonical-0320300110213310-0202010300303010-0013003110121201-3312122320013023-2132220323122030-2103123021332330-2003223133023000-1223131021130130"></a>

<a id="canonical-2213101302023333-2233303123210313-1332121112132232-3312300101022113-0302010233011210-2111123030222232-0113210231302022-3303111130000123"></a>

## provider_ref property — clear_secret_info / 033023203021 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0122312232323000-0100101100220310-2310001101020200-1012310220001233-2121222033231311-1011131332311003-1201222210300120-0203131110112031"></a>
