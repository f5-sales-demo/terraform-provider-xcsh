---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022032232213312-3323333021303101-2101020210011123-0113032001332322-2001311211230230-3110020122003002-1310313210022320-0303232332322232"></a>

## Property reference — Property reference / 111013232123 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- Property reference

<a id="canonical-2032303112312321-1111322023022201-3101202331331101-0133302020032332-0232223013322221-1123302100220112-3320001210322202-3021020201310023"></a>

## Direct properties — Property reference / 111013232123 / 3

<a id="canonical-1213313012313112-1021101020132223-1112032233233313-3113130121021212-1020222100320332-1000102331201322-3210101311000310-1132231121000211"></a>

<a id="canonical-1213000011232012-3200320113133013-3200101102130101-2200331211202131-1113133032211003-3003312012113212-1320301002002332-2321000222033122"></a>

## address property — Property reference / 111013232123 / 4

Type: `"string"`. Computed.

Site's geographical address that can be used to determine its latitude and longitude.

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

<a id="canonical-1132032312013111-0010020123030333-0033303100202133-0201030101333301-0032013022321332-0102021210000313-3211012230312203-0122022210030320"></a>

<a id="canonical-2020210120101031-0101211100111322-0233333321123010-2303322322313130-3120012133001321-0321231111233013-0102000003101203-1022203311301011"></a>

## annotations property — Property reference / 111013232123 / 5

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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

- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-0230212222332032-2133213333223020-2203212003102330-2322011323032333-1322302002100331-3032210130333303-0102011202012330-1010333331021220): complete subsection reference.

- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-2010013132202033-2110133330011130-1310012103113012-3100032310101130-3003301111111022-3001111033212101-0321233100220322-1323121313313301): complete subsection reference.

- [coordinates](data-sources--securemesh_site--reference--group-001.md#canonical-1011033201210123-3012222111003002-1033110130112121-0220000011220021-0121202312300300-0303330130211101-2021110123002101-3330010310012223): complete subsection reference.

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302): complete subsection reference.

- [default_blocked_services](data-sources--securemesh_site--reference--group-004.md#canonical-1332002011330322-3112302133232021-3020303330123333-0121332303210202-2011220212000322-3013133110313113-2300320313113130-0311122221021323): complete subsection reference.

- [default_network_config](data-sources--securemesh_site--reference--group-004.md#canonical-1032132001130300-2232230303121130-2322233123010303-1320223100112113-0002203210303003-1010323303201032-2221222230312013-0301020013000030): complete subsection reference.

<a id="canonical-1020202211300103-1113100020130322-3223101010321231-0301003120203130-1312121101230132-2232203303302322-2102310223233321-0213222032010321"></a>

<a id="canonical-3331201210302222-0100213223221222-2310222333113101-2202330203120131-3201210300013110-0313210102200011-0030111101132301-3322111002130002"></a>

## description property — Property reference / 111013232123 / 6

Type: `"string"`. Computed.

Description of the SecuremeshSite.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3000213122000303-2031300013032002-2000320323020130-1321022210201313-3133133211300103-0002322302223212-0032223333003213-2233203332203221"></a>

<a id="canonical-1101200213233330-3133301323302113-0300131033200102-0311210221233322-1332022212101300-1301021232001333-1001033133313022-0133012001321113"></a>

## ID property — Property reference / 111013232123 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2003312033033001-0003322123103200-0003210200302311-2323013312221021-1032032302211212-2033100222113033-3311011230311102-0221123310333312): complete subsection reference.

<a id="canonical-1100320301231103-0221203211221213-3130200300002311-0311003002201231-1131031031010213-0222232022311110-0003213021123001-0313310330131023"></a>

<a id="canonical-2233220002312002-0321132202101113-3313100211332001-3010322120330032-2003131023311001-0032233213312113-0321012012013102-1223321020230123"></a>

## labels property — Property reference / 111013232123 / 8

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

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

- [log_receiver](data-sources--securemesh_site--reference--group-004.md#canonical-0211220000103000-2211130101211031-0220003100131103-2300012120232311-0013332232122023-1233120111001231-1112013003212021-3201020203022120): complete subsection reference.

- [logs_streaming_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-2030232101301321-3022120220023011-2330231031223312-0201232010301230-3330312010210010-0011010130130130-3212310310100011-2332213030110233): complete subsection reference.

- [master_node_configuration](data-sources--securemesh_site--reference--group-004.md#canonical-3313100220202200-1000302100010131-1301231302121331-1300120222103101-0133123100113323-2120331223001002-2101030221002332-2323223232331200): complete subsection reference.

<a id="canonical-3123223031003332-0331100130321123-3332223232231101-2320103320330012-2131132032131112-2212300112330000-2020133113031202-3212202030222321"></a>

<a id="canonical-0001332120220013-3202201323313030-0311001031223033-1131032032330330-2231210030333033-1131123201111220-0313121333322121-3021231222010123"></a>

## name property — Property reference / 111013232123 / 9

Type: `"string"`. Required.

Name of the SecuremeshSite.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2010231111032120-0112222231103113-1131303023013113-1021001132102001-1011223231133213-1131300222220030-2320110231201130-1332211323101212"></a>

<a id="canonical-2131300200202222-0201130113211123-3221221213010222-2123122121122101-0121023301302021-3100330121121202-1011013100213020-3202003120022023"></a>

## namespace property — Property reference / 111013232123 / 10

Type: `"string"`. Required.

Namespace where the SecuremeshSite exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
  }
}
```

- [no_bond_devices](data-sources--securemesh_site--reference--group-004.md#canonical-1003113100300311-2222313300001010-2021310302002030-1112332132120233-1002033303013130-2231330222100030-0102321101101202-1321023010332110): complete subsection reference.

- [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2020103213000102-1110301232330123-2322321023321200-2131023113310313-2233322222100111-2032312323132100-0321122012303013-0320311231012313): complete subsection reference.

- [os](data-sources--securemesh_site--reference--group-004.md#canonical-3012121031011130-1233001032220331-3130122211200032-2131001233320133-0031212331013332-2032112211110320-0203323010102313-1032213120331101): complete subsection reference.

- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331): complete subsection reference.

- [sw](data-sources--securemesh_site--reference--group-004.md#canonical-1101033231030033-1231112223010232-0121200330001132-1331111222212213-3023302310111332-0331233101000132-0230010020103302-3320133222123131): complete subsection reference.

<a id="canonical-1200001331121323-2000312231301022-1021131121133231-2203101233230130-2002023302203220-3300202330010102-3001000032303102-1022301010212002"></a>

<a id="canonical-3322311120101110-1310032323003231-3101112133130211-2331132201103103-3222010030221112-2200210212313102-2232232323001233-1123131222301122"></a>

## volterra_certified_hw property — Property reference / 111013232123 / 11

Type: `"string"`. Computed.

Name for generic server certified hardware to form this Secure Mesh site.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "minLength": 1,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

- [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-0301131103113000-3303232313310330-0122211113001113-3130222002332001-1310200111021022-2330131102030223-3102231232233231-3301133132020110): complete subsection reference.

<a id="canonical-1310131033332013-0202121123323330-0102113222101230-0102212013321312-2232110300310023-3323113211023021-3230000011311301-0203221321312013"></a>

<a id="canonical-3130311320101120-0033102220132002-3120033233120002-0212211112210203-3312233021332023-2022112031313123-2023303221030232-2023020003033301"></a>

## worker_nodes property — Property reference / 111013232123 / 12

Type: `["list", "string"]`. Computed.

Worker Nodes. Names of worker nodes.

Upstream description:

Names of worker nodes.

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

<a id="canonical-0223011303112132-0300101120121102-3231220200332313-1103320300213123-3112100100212301-0321131130320011-1211323001121113-0023100023312101"></a>

## All schema paths — Property reference / 111013232123 / 13

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--securemesh_site--reference--group-001.md#canonical-1213313012313112-1021101020132223-1112032233233313-3113130121021212-1020222100320332-1000102331201322-3210101311000310-1132231121000211) |
| `annotations` | [annotations](data-sources--securemesh_site--reference--group-001.md#canonical-1132032312013111-0010020123030333-0033303100202133-0201030101333301-0032013022321332-0102021210000313-3211012230312203-0122022210030320) |
| `blocked_services` | [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-2032103203120130-3220033100222121-2102023221121102-3220133021022303-2130033131221111-2233000121213101-1132202212331333-3122101200303103) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-3033120300133112-2112112020312200-3330332201200303-2021301000011301-0113322310122322-2210013310013113-3333311220222132-3122230310202321) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](data-sources--securemesh_site--reference--group-001.md#canonical-3333112331020021-3203232320322011-0020211131230231-0212100113313020-0102311121323230-3101111223011132-1302310020313321-3211111231001013) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](data-sources--securemesh_site--reference--group-001.md#canonical-3300121033020031-2321321222231130-1030130011211030-1312120303311101-3001331231103131-3202013130321332-0223130013222210-1101101100233233) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](data-sources--securemesh_site--reference--group-001.md#canonical-0113102000201120-0211013131300323-3210333313230121-1022223331110202-3211131103100310-1211110001120300-3003132000231231-2322230001233031) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](data-sources--securemesh_site--reference--group-001.md#canonical-2200333102320321-2323003001111323-3023121331213303-2232323130311220-2220120001331031-2200330221202233-3030332033321333-3200203301331211) |
| `bond_device_list` | [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-2013100011013002-1322321012232320-3131331223201021-1233321301001010-0302132332033003-3122030000322130-3012320010232210-1313213002003321) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-3120001323121310-2000002020130202-3203230020222022-3023100012033200-3123022332303010-2322101121013112-0322212302120111-3333203313111102) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](data-sources--securemesh_site--reference--group-001.md#canonical-3000322220121222-2001002212321031-2110321331201311-3113021323312020-1332232221220011-3011013131222320-1221233002021222-1232330333123001) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](data-sources--securemesh_site--reference--group-001.md#canonical-0103011200023021-0322102223220233-1303112232210123-2310012322103310-1323333112110301-0102220123211103-1313212021002000-1332301231213323) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](data-sources--securemesh_site--reference--group-001.md#canonical-3220002101223012-2132212011311121-3132030001001320-1202310133120330-1013201321302312-1001333300102111-3212310100033111-3031022313000300) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](data-sources--securemesh_site--reference--group-001.md#canonical-3032012230300202-3223112332320202-0232030232131123-1221200020300100-1313020233313200-1113312002103111-2123211300021012-1222021222110032) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](data-sources--securemesh_site--reference--group-001.md#canonical-3031103203212002-2113301102021103-0313113200103021-3301333111121031-1211300220123230-2221313122122211-3110313112032020-1122111001331023) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](data-sources--securemesh_site--reference--group-001.md#canonical-3222211123020122-0032013030132233-0003031102020121-3332320112100010-3211310102211011-1000001110312030-2333113330121201-1311113203010031) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](data-sources--securemesh_site--reference--group-001.md#canonical-2231230101123012-1232332001210131-3332132203111032-2132233303302110-3303212200230002-0122013320211111-3101312032002121-0112330000233311) |
| `coordinates` | [coordinates](data-sources--securemesh_site--reference--group-001.md#canonical-0310002233301332-0203132332010333-3221203003210002-0030100102320021-2323200302012321-2303331201310022-2033000313321221-3022101002231313) |
| `coordinates.latitude` | [coordinates.latitude](data-sources--securemesh_site--reference--group-001.md#canonical-3322213320332223-1101123103012303-3023113312332333-2310230230120313-3131330123223010-3222122031132133-3323231311113032-1210303310301010) |
| `coordinates.longitude` | [coordinates.longitude](data-sources--securemesh_site--reference--group-001.md#canonical-0311002313021301-1132301012103333-0212110113122101-2233303211132211-1201223002221013-1122322032211030-2030102321230233-2331123300120132) |
| `custom_network_config` | [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-2233320021012133-3133111313320122-2033222313102002-1133023232002203-1233333201211020-0101301313301031-3010033010333012-2323223323102013) |
| `custom_network_config.active_enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2231221221332213-2233233321332130-3332323020032013-0111333202023220-3321133123221023-1313233003133200-1222201330202202-2000032130212033) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-0200231000333120-0322311003031032-2122211012012303-2212220031333300-0333320221112200-1013211202203311-2000300330020300-1001100101311321) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name](data-sources--securemesh_site--reference--group-001.md#canonical-1310021232303200-2123033100133221-1212200213221012-0231213033113331-2332223221331120-2023113111102132-3101313233213023-1310211110220231) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](data-sources--securemesh_site--reference--group-001.md#canonical-2211210212120020-2103010030231213-3103321330323132-0332323023101220-2022230201312022-2131230002120100-3303231113132033-0223013120002230) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](data-sources--securemesh_site--reference--group-001.md#canonical-3100123332101210-0001002012231113-1002003220122230-2202020302033022-1213323332110200-1300213200301132-0021223333303201-2320233011301010) |
| `custom_network_config.active_forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-1211000032103020-3302102123103323-2033010013032223-1312001102021012-3021102001112332-1301103020012132-1110110303320313-3012213303211130) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-1302010221203110-3022121113333223-3031221222332333-3203201321102133-1303123333302123-0323002301330021-1223321011332330-2303230020202333) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name](data-sources--securemesh_site--reference--group-001.md#canonical-3011023332033013-2333120211031110-1110223020312211-2010131130333102-1112103013120331-1233030000312012-2100210120321003-1130303330133231) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--securemesh_site--reference--group-001.md#canonical-3012001013211231-0102312201131200-2331120131233113-0233110120312021-3003020301120033-3133202100100302-1110123123131212-0331221311030310) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--securemesh_site--reference--group-001.md#canonical-1302331021021013-2201211301032323-2201231122321130-0123211132233000-0221222020000232-0013203233132323-1030333100322231-2110102112221122) |
| `custom_network_config.active_network_policies` | [custom_network_config.active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2310131011010330-0030132311203103-1230232212011103-3222121112332122-3200001233033120-0330030301321222-3010213021102302-1220112110013332) |
| `custom_network_config.active_network_policies.network_policies` | [custom_network_config.active_network_policies.network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-3310120021212322-0312331312220121-0332222221203021-0222333310010232-1210313103230232-2220232313221112-0302212000031102-0113223101213011) |
| `custom_network_config.active_network_policies.network_policies.name` | [custom_network_config.active_network_policies.network_policies.name](data-sources--securemesh_site--reference--group-001.md#canonical-2131322000303022-0210030120030331-1131301213322123-0100103331010213-1000320102020021-2031032221000223-0101011020310212-1210303220333003) |
| `custom_network_config.active_network_policies.network_policies.namespace` | [custom_network_config.active_network_policies.network_policies.namespace](data-sources--securemesh_site--reference--group-001.md#canonical-0023110213001123-1222333320200001-2003030112303022-0323311321030023-3000103123110123-2121211231012210-2120030320010110-2311213313033032) |
| `custom_network_config.active_network_policies.network_policies.tenant` | [custom_network_config.active_network_policies.network_policies.tenant](data-sources--securemesh_site--reference--group-001.md#canonical-2302202110321300-0101202001132012-3132330200300022-0312031020102321-3211132211221212-3100303212113331-0022200302003230-3223112032121122) |
| `custom_network_config.default_config` | [custom_network_config.default_config](data-sources--securemesh_site--reference--group-001.md#canonical-0333311101111230-0021121321102033-2210121231023110-2201330330002101-0300221102323232-2020311100323110-2301011310223103-2002233331213300) |
| `custom_network_config.default_interface_config` | [custom_network_config.default_interface_config](data-sources--securemesh_site--reference--group-002.md#canonical-2232302102110131-2300101300331130-1110233201020003-3103003123311333-1222301310332130-2300202311022311-1221330232110201-1100212331333110) |
| `custom_network_config.default_sli_config` | [custom_network_config.default_sli_config](data-sources--securemesh_site--reference--group-002.md#canonical-1333232020032320-1302020130120320-1302012103211021-0300002232003310-3112111211330102-0101011020303131-2202311233323010-3003212030112333) |
| `custom_network_config.forward_proxy_allow_all` | [custom_network_config.forward_proxy_allow_all](data-sources--securemesh_site--reference--group-002.md#canonical-1300223103320320-1300212111122103-1102103313301112-2211122123233301-0001210032310011-2033302012101031-3221120112023231-3232021233002200) |
| `custom_network_config.global_network_list` | [custom_network_config.global_network_list](data-sources--securemesh_site--reference--group-002.md#canonical-3023310221220300-1302002010133212-3210311301321133-2022300122203110-1110030311012220-1101300331211312-0202231212303313-3211011101100300) |
| `custom_network_config.global_network_list.global_network_connections` | [custom_network_config.global_network_list.global_network_connections](data-sources--securemesh_site--reference--group-002.md#canonical-1110200103322021-1210111323300102-2330330020323322-0103110033112110-2122311202202100-2231101230131310-2031301012332120-0022032020222313) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--securemesh_site--reference--group-002.md#canonical-2110033010023120-2130100332011210-1023301133103010-0213212101320221-2130003321220023-1222033103021013-0020122010122131-2213212002302330) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--securemesh_site--reference--group-002.md#canonical-1021112022332212-1003231221233332-3032331311211000-2222112333031202-0221201033102331-3301303321032213-2031220311011331-1002102313103210) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](data-sources--securemesh_site--reference--group-002.md#canonical-3033100330130220-1001312020132223-2233200231320330-0012031113121113-0022311112232031-0132332313220111-1001010111221100-1200331000101121) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](data-sources--securemesh_site--reference--group-002.md#canonical-3233313110023013-2022132100312013-1203113100300330-2310330013322220-0331320223330130-2112331203130031-0230210123301231-3023211131312331) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](data-sources--securemesh_site--reference--group-002.md#canonical-3210303102103121-1012331033013303-2202103313003133-3313011330301112-0230112002001100-3011103300132122-2231331322020323-0302321220032132) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--securemesh_site--reference--group-002.md#canonical-2223132301121023-3310130301301200-2100100200323021-0103022020223111-0313330200102222-0323033031213222-2031030002000220-3230212222332231) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--securemesh_site--reference--group-002.md#canonical-1010202331031111-3030032320202321-1130111030101211-0321301003112011-0332133211322132-2213032032333013-1332221022023211-2131122112313022) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](data-sources--securemesh_site--reference--group-002.md#canonical-2120232311011003-0022322232322130-2022032313320303-3203221101212102-2301021203321233-1011101232313330-2333331122321032-3230032032323221) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](data-sources--securemesh_site--reference--group-002.md#canonical-0110203323111223-0312003032212200-3233033101213023-2320012021230313-3030203110312103-0030322111133021-2103220201230030-3103320133110000) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](data-sources--securemesh_site--reference--group-002.md#canonical-0011132313122120-2003030232230011-3032323321013000-2033110232132313-1230102032003030-2033211213313113-2323211333233303-2111310211123211) |
| `custom_network_config.interface_list` | [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3112322013210300-2310210212202233-1233001131103233-0120201310210033-0012103332220321-2210203011021323-1220221221300110-0121122113011032) |
| `custom_network_config.interface_list.interfaces` | [custom_network_config.interface_list.interfaces](data-sources--securemesh_site--reference--group-002.md#canonical-0113222210001322-3332033202322010-2212133102001223-1000033320232031-1230320220020313-2220323021313212-3111303200002220-2331233222211103) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](data-sources--securemesh_site--reference--group-002.md#canonical-0002121102222003-0023010103003102-2131022220211201-3230001110012302-2110232322302201-0010022000202223-0113332123101201-2130303320300002) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](data-sources--securemesh_site--reference--group-002.md#canonical-2220232202201003-2102333303321201-1310131131012110-2333211103112201-0220121311101232-3110321000000230-0303113200313302-1012202123312332) |
| `custom_network_config.interface_list.interfaces.dedicated_interface` | [custom_network_config.interface_list.interfaces.dedicated_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3022133301323021-1010030202011303-1212112212202132-1311213312303030-1222331230200133-1220311231313120-2032223201123123-1102122111303133) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_interface.cluster](data-sources--securemesh_site--reference--group-002.md#canonical-0222323011301031-0002300201330213-1301023233112213-2201202123302113-2213210030303230-2322330022110302-0313021033132021-2023020320132233) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_interface.device](data-sources--securemesh_site--reference--group-002.md#canonical-0112301303222103-1332120321001210-0332101033013322-0332231311112202-0313313300002022-1112211111230021-2002021111012130-3323012022302033) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.is_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.is_primary](data-sources--securemesh_site--reference--group-002.md#canonical-2101100330123300-0102323203313300-3220001122020332-1111002331031120-1132312301201023-0312200020232331-2313002220102320-2032303330220032) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor](data-sources--securemesh_site--reference--group-002.md#canonical-2221020323202013-1111103210132021-0102211223333102-3000310202113332-3213033103203013-0220101330100112-1321103030120123-0300201121230103) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled](data-sources--securemesh_site--reference--group-002.md#canonical-2333223212023231-2111300102213310-0200020120322012-3023011322113113-1332322303123123-3013003110321110-2310300003312313-0002210010130001) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_interface.mtu](data-sources--securemesh_site--reference--group-002.md#canonical-3211001331121133-3130333023201032-1202110201012211-1120031212002320-2232113012202101-3303011232330202-1122012301103322-3232030201021011) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_interface.node](data-sources--securemesh_site--reference--group-002.md#canonical-3131331121313202-3222110111200230-2000122221302230-1113212133023303-2310012321210120-1320313132212013-1332231230103102-0322100320030121) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.not_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.not_primary](data-sources--securemesh_site--reference--group-002.md#canonical-3222100232030213-2300110201311301-0300310232023310-0003100121231301-1213112322122223-3011323201213032-1031001232312311-1032120011312033) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.priority` | [custom_network_config.interface_list.interfaces.dedicated_interface.priority](data-sources--securemesh_site--reference--group-002.md#canonical-1023221313301021-0231003212111021-0212033022232102-0002220220311013-1202101120221213-0200302223222332-2011202223323222-3033333313301100) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface` | [custom_network_config.interface_list.interfaces.dedicated_management_interface](data-sources--securemesh_site--reference--group-002.md#canonical-3210221002221233-1110210113003103-1103231100323023-3203121310130322-0320011032111102-2033201210122100-3210213103102123-3310023311013013) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster](data-sources--securemesh_site--reference--group-002.md#canonical-3101020103233111-3031231201200300-2010313101113002-0321303200101020-3310030202033013-0231113312321032-2013130022331020-3312031231200000) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.device](data-sources--securemesh_site--reference--group-002.md#canonical-3111122311232223-1031333102033123-0110331230110312-2132001303201103-1113222201311020-0132030031123212-3221021303220021-2333120132100021) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu](data-sources--securemesh_site--reference--group-002.md#canonical-2322222002311103-0130312221000130-3132132101022203-3133013200313020-1201313033200023-2212121100331002-2212112331331200-0012122200300212) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.node](data-sources--securemesh_site--reference--group-002.md#canonical-0123130010310023-1021020210322300-1323012232120013-3310311212201010-1031303112132330-1033333110221233-1303002022013122-0010001011101233) |
| `custom_network_config.interface_list.interfaces.description_spec` | [custom_network_config.interface_list.interfaces.description_spec](data-sources--securemesh_site--reference--group-002.md#canonical-1201200232222320-0022013233021303-3102333112213120-2003120120312133-2321320232203202-1211102103333023-3200133020333302-0122032202102111) |
| `custom_network_config.interface_list.interfaces.ethernet_interface` | [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--securemesh_site--reference--group-002.md#canonical-1022212232030301-3232133210323302-1120332033001212-1203203031122321-0000211223301030-1210120130203012-1021332122020022-3110202300012220) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.cluster` | [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](data-sources--securemesh_site--reference--group-002.md#canonical-1021020202020010-2330203313103103-3132231332033012-1301102100221020-1232120130123232-0223100202101131-0211322313110033-2011213100010310) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.device` | [custom_network_config.interface_list.interfaces.ethernet_interface.device](data-sources--securemesh_site--reference--group-002.md#canonical-3110130011020022-3031110231301200-1233233211313012-1211322323221332-3130203023212333-0220033211202111-1223012213020002-1002210133032201) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](data-sources--securemesh_site--reference--group-002.md#canonical-3003333233310300-1301012213030330-2213301210201221-2131311021201302-1201301013300101-3111100110323212-1101323330202222-1032012030121211) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--securemesh_site--reference--group-002.md#canonical-1030122312331223-2222221320101001-3130131321212103-0201100223033110-0100331333203222-3030330031131221-2111023003210303-0221230313330332) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](data-sources--securemesh_site--reference--group-002.md#canonical-3002101301123131-1112120232013102-1033102033033331-1202121320033033-0101130231123032-3331133010101110-3312000311201322-3023320303332121) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](data-sources--securemesh_site--reference--group-002.md#canonical-3030302010032112-2112313010301301-2313111113301331-2223203322021013-2100203103022013-2013010233212110-3230022322311120-3221020211112200) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--securemesh_site--reference--group-002.md#canonical-2032011320132321-2003310211000232-3003213220303031-2311001220011312-3211121313320310-1210321100222310-0133332123103202-0022111032033222) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address](data-sources--securemesh_site--reference--group-002.md#canonical-1232232211122200-1332101332233321-3123101120033011-3003101320303012-3213203012230330-1131322223211211-3233202102112032-0011023121111033) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address](data-sources--securemesh_site--reference--group-002.md#canonical-2323302120132122-3221113211330220-2332103132201221-1003313102311201-3002022212333232-3321330300032210-2301211223101112-2210322122103133) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site--reference--group-002.md#canonical-1123031302312121-0321113030000023-1322220323331301-1231023120002110-2333122333310333-0022330130120202-1220233000022013-1030033313223023) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site--reference--group-002.md#canonical-3112230120011013-2033311201013221-2331332033232122-2301000111012332-3210300013022213-0021120123313312-0130030110033013-1110230223020300) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix](data-sources--securemesh_site--reference--group-002.md#canonical-3323230022032310-1002022213020302-0220022330313111-1222112032113232-0210211001103010-2111200322303013-0210112302030323-0200033211000333) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings](data-sources--securemesh_site--reference--group-002.md#canonical-1210002003210120-2200133201002030-2012211201012103-1000211033133320-2100031223323113-3022300012121110-0000003121320301-1322110321111101) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site--reference--group-002.md#canonical-0021021012300302-1100213220123322-2210221311012120-0003030231303013-3110333323201133-2301132233013301-2313210322221311-2011131302110020) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](data-sources--securemesh_site--reference--group-002.md#canonical-0233231231001331-0221332010300312-2011120301233110-3310233122103130-0323230213011233-1333203103133101-2102310210132323-2002100100232001) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](data-sources--securemesh_site--reference--group-002.md#canonical-0221023020100021-2032120231233312-3213023333131010-3311121231103010-2332230310033210-2230233101121322-2032110031301321-3323220132211121) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](data-sources--securemesh_site--reference--group-002.md#canonical-0302312332000213-2112313033030323-3030303132302120-0322312222301310-2300323111101012-1120231123030211-1020011300032220-3021331230321220) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site--reference--group-002.md#canonical-3122331322132022-2300023233211002-0003212123023312-1122123232233222-1021323320112212-3320011230100323-2323322232332131-2333010322030222) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag](data-sources--securemesh_site--reference--group-002.md#canonical-1102313031323313-0303111303011023-1031202302233120-2223210202321120-2223210313231231-1212103022132022-1003231223122312-2220011030230203) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map](data-sources--securemesh_site--reference--group-002.md#canonical-2112320022210130-0132213322321322-0300031010102020-2122120311212230-1330123112112230-1113011210031213-3321103001000201-1202132132132210) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](data-sources--securemesh_site--reference--group-002.md#canonical-2033001112110333-0201220012122113-2231210303311313-0201201133202111-2021000100300332-0121301302212012-0323132232211213-2323120333221132) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](data-sources--securemesh_site--reference--group-002.md#canonical-0311133330010101-2031303132020220-3320012002320301-1010313022100212-1020123023131021-0233331020101011-0022323332213310-0131212100011003) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--securemesh_site--reference--group-002.md#canonical-1000032202212113-0321333133032113-1221023333002001-3211232021310203-1312122012102303-2221200013312301-2301310332313102-2301330013331301) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](data-sources--securemesh_site--reference--group-002.md#canonical-0333023020211333-0331000111012100-2203203123201303-2333022302222113-1123210103332123-0000221110020022-2210332020322213-3130220102230133) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--securemesh_site--reference--group-002.md#canonical-3200323310203232-0030201222233020-0333232011101012-1003031022021223-3133003200313123-3221000202131121-1223310320013022-1001322323323320) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--securemesh_site--reference--group-002.md#canonical-0130230330230320-0120203010100201-2013013102103101-1200303203330111-1020103020032132-3131010230103310-1102111202220302-1232003032022212) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site--reference--group-002.md#canonical-2201120333231100-2122023211212223-3131030232311302-1333011030230010-2032331212001303-3012003121202101-3322003013120033-3011222333013001) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](data-sources--securemesh_site--reference--group-002.md#canonical-2100330220112213-3330331002320230-2222223002232112-3102230000030301-2020002001031310-0212033130300323-0201022000112030-3013110210021220) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site--reference--group-002.md#canonical-0013300010223331-0012223210230011-1301300002332333-2032130033220321-1331311333320013-1323103220303303-2112212113320121-3333102311331030) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](data-sources--securemesh_site--reference--group-002.md#canonical-3330103230200130-3331032213131102-3222012002130332-2030010312223002-0123231003030113-3230021100120231-0231132223110331-2021200100022200) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site--reference--group-002.md#canonical-0203110231201222-1322020112313320-2301033200031132-2220113030221220-0213310002311012-0333323111332001-2220321213331321-0010303230133013) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site--reference--group-002.md#canonical-2103113221102023-2001022110333301-0130221132211233-3203101030300211-3020002203330331-2333012020002320-0220113323202012-3002310320010132) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix](data-sources--securemesh_site--reference--group-002.md#canonical-2220323221000330-2000012211103133-0011110302011113-1131323310222213-1303012101233201-1202121223221011-3133213201103023-2303101013220332) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--securemesh_site--reference--group-002.md#canonical-1203132001100132-1333001221003131-0300213022313123-0013010323023211-3102012311200021-0023222020001033-1231031131321032-1100032320011202) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--securemesh_site--reference--group-002.md#canonical-3321211103322323-2030310113320021-3101002123212122-0133220130032121-1311232011230111-0112013330320220-3210101111102031-2230321202102320) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--securemesh_site--reference--group-002.md#canonical-0013301011023133-2121202232103202-1032331220133020-0330012333321033-3331301023103222-3312200301102222-0212201112133232-0210001010323032) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site--reference--group-003.md#canonical-1323312213010023-1113023330133013-3013300003201203-2233202013331323-3021333311231101-2113320020221331-0303323031330323-2102213210033010) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](data-sources--securemesh_site--reference--group-003.md#canonical-2311213123000010-1120322323132333-2100132130021101-2113332210020012-2312031212120233-0330321132330110-0120313121311233-1002032232222120) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](data-sources--securemesh_site--reference--group-003.md#canonical-3133313223230011-1302212220200020-0202121212132123-0130223332131003-2200021223231120-0102103223313301-0311113122132023-2222001232332012) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--securemesh_site--reference--group-003.md#canonical-1333000131323031-2101310123010020-1331010001313121-0001113223031030-2022112210112231-1101332120113102-2322030013121103-0120231321213300) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](data-sources--securemesh_site--reference--group-003.md#canonical-2333331000032131-3202332123310111-3230000023322012-0212300233033211-3201222011012302-0332030201112123-3113200331111232-0012010033302222) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](data-sources--securemesh_site--reference--group-003.md#canonical-2203230312311122-0203332033233330-1121213113311133-3223200121013203-1313201201033201-1213100210122310-1232111003122132-3112002001311231) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](data-sources--securemesh_site--reference--group-002.md#canonical-1103030322133112-1200230312122111-2030110310133032-3213132320001210-0230322023123232-3031220311003303-2100010330013011-2002322010322021) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--securemesh_site--reference--group-003.md#canonical-0302131131031330-2300333001112021-2230332300331111-0300003103222213-3023103111203122-1212211113220310-1213302101300223-3102231313033222) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](data-sources--securemesh_site--reference--group-003.md#canonical-1221030200112322-2101012230230010-3002012021310200-1132113233010223-0203223323310132-1211033211103232-0020103300232331-2003021302011113) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.is_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](data-sources--securemesh_site--reference--group-003.md#canonical-3121220030212331-3122312203202111-0231313222110313-2300330132031022-1023200111313312-0301113203330132-3330110221203102-0312301313311111) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](data-sources--securemesh_site--reference--group-003.md#canonical-2003211312301332-3132220001212230-2230300223312031-1032320001201113-0122231132133121-1111033013313001-1203301310211222-3021203313322203) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](data-sources--securemesh_site--reference--group-003.md#canonical-1033213023213311-3121211313310311-3300312333312121-2013000301133222-2321333103133210-1300310221013100-2112123110203101-0012312021110310) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.mtu` | [custom_network_config.interface_list.interfaces.ethernet_interface.mtu](data-sources--securemesh_site--reference--group-002.md#canonical-1233100002131323-2021030101131221-3110212103223111-1220011232301012-0312200012003202-2310022003323000-3232030133100030-0231202302011321) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-3203323101200331-2231200030221112-0210002220000232-1223021111032030-0321333203313311-3223001123230301-3113011332122002-3331232311122002) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.node` | [custom_network_config.interface_list.interfaces.ethernet_interface.node](data-sources--securemesh_site--reference--group-002.md#canonical-2120123031030110-2321220332310303-1331122002320101-1201113122020100-0013003020023020-1023001102220101-0022213311320330-0123201312030323) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.not_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](data-sources--securemesh_site--reference--group-003.md#canonical-2232102323120102-0012302320122231-0020003021032322-2201213021201203-2233133332223310-2331312131130023-3031122302303110-1222220013000033) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.priority` | [custom_network_config.interface_list.interfaces.ethernet_interface.priority](data-sources--securemesh_site--reference--group-002.md#canonical-0213212002331110-2222010221220303-0201001201230202-0232112001232223-0320003310203332-3012232333201320-3202031310303300-3021000112121110) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](data-sources--securemesh_site--reference--group-003.md#canonical-3223131023011013-2002133303330001-3003202020010120-1021101133002001-3111321223131122-3023300021001123-3010311123223311-3222030231333321) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](data-sources--securemesh_site--reference--group-003.md#canonical-1120103332113103-1230033332113021-0310122022301101-0202033021121013-0133112233211111-0220123320210002-0212033320201313-3201132112331222) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-0332212332202103-2032113102321301-0021302111313300-0021223211201310-3200300002202113-1212332033322020-2333302003102032-0213100131111312) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-2223200300303121-3001021110321112-3113131000201332-0201321013323122-0323030020220021-1301003322013003-1220303000102103-3220002223202003) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](data-sources--securemesh_site--reference--group-003.md#canonical-0211113131311222-3030112323203213-0302322100230001-2100211323231132-0200321333212132-1201230332333321-1231003102102112-0033122302330133) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-3021132310001110-3212303202220312-0121123120302131-2130010210232113-0321112210310330-1231011002333222-0130031300123012-0113232312121301) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw](data-sources--securemesh_site--reference--group-003.md#canonical-0231101102111110-0230202220002320-1222031112100321-3131002330103301-3311011101010211-3321212112131011-0112231332012320-3220123020232000) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server](data-sources--securemesh_site--reference--group-003.md#canonical-3120202322311110-0201131002210222-0103031311000003-1100003303320222-1123203101231100-2301132101132320-1301233311203133-2111011031210303) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-2131313223223213-2101022222130122-1230132233200200-0111331210203112-3232210233030302-0002330032021121-0313220122031321-3202203011210202) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--securemesh_site--reference--group-003.md#canonical-1312030021101203-2221322310022020-3030120321323131-1030211333000123-1223333123301331-2310211322012113-3312202100313303-3211332001200322) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-1100031010302103-1303010010313231-3303111301221230-1312100122311221-0113021132102202-0230310212111123-0313233220103223-3030130210010223) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](data-sources--securemesh_site--reference--group-003.md#canonical-1221000031121220-0101301331202111-3032320100322211-3311231103121300-3303232000020302-2011332232030100-1000002330003130-0303033132200331) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](data-sources--securemesh_site--reference--group-003.md#canonical-2201132020311010-1233210100322101-1320121112332223-2001103100000132-2123213231330033-2210311133112311-2330303023231000-3203330200210102) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw](data-sources--securemesh_site--reference--group-003.md#canonical-1020031120133120-2202000302311230-0102012120012212-3111000321212212-2011032231221221-0022023103203110-2203333032003211-0313023333122233) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server](data-sources--securemesh_site--reference--group-003.md#canonical-1220022201311002-1323130132121130-3120122333211322-0003010213001220-2211322203200030-3133211133302131-0131313222213100-1303113321233122) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-1120121003000331-1202331000213133-2023022332221223-0130031031002003-3103012323221321-3220203031303322-3231101332100100-0202111213000202) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.storage_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](data-sources--securemesh_site--reference--group-003.md#canonical-1022100303122012-0113121323103320-0232110132001322-2220301230203210-2122302022031232-1011123223330310-3313301000310231-1210320202023002) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.untagged` | [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](data-sources--securemesh_site--reference--group-003.md#canonical-2130000133101322-1021033312120211-2010122313020132-1203201320032021-0001011231012110-1321100212211120-2021031211000322-0130303331321000) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id` | [custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id](data-sources--securemesh_site--reference--group-002.md#canonical-2321113020112322-1221233230032013-0223013223231003-1312233201133021-2211210113311221-3032320000003323-0000322303211313-0202012301132311) |
| `custom_network_config.interface_list.interfaces.labels` | [custom_network_config.interface_list.interfaces.labels](data-sources--securemesh_site--reference--group-002.md#canonical-0103122323222030-3202322220231100-2333212223230300-3301131122121212-3123333020123321-0101223010200232-3331330112110203-0100121133011133) |
| `custom_network_config.no_forward_proxy` | [custom_network_config.no_forward_proxy](data-sources--securemesh_site--reference--group-003.md#canonical-3002103003330320-2131230000033032-2322102012023310-0313013311312032-1220000310202121-3301331331132001-0220313213323023-3231012202103213) |
| `custom_network_config.no_global_network` | [custom_network_config.no_global_network](data-sources--securemesh_site--reference--group-003.md#canonical-1111330110231302-2333020021330122-1133202012223120-2202231311301133-3331131300001310-3312202103033323-1222130233120112-3112230230100201) |
| `custom_network_config.no_network_policy` | [custom_network_config.no_network_policy](data-sources--securemesh_site--reference--group-003.md#canonical-3301231103303223-3023033332232001-1223010023210301-3133310220311032-3121222222323000-1020211010203131-1333110033110213-0320132302211112) |
| `custom_network_config.sli_config` | [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-2231020131223023-2310221002020202-2033232322303321-1230123330232230-0220223213033322-3313020230211112-2201312013110103-0001232100102322) |
| `custom_network_config.sli_config.dc_cluster_group` | [custom_network_config.sli_config.dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-3100202303332203-2000302303301220-2311022013132030-3203003200033220-3203202231112111-0101312312121123-1232212021021131-1111301323231012) |
| `custom_network_config.sli_config.dc_cluster_group.name` | [custom_network_config.sli_config.dc_cluster_group.name](data-sources--securemesh_site--reference--group-003.md#canonical-1303233011311132-3103232201120233-2131203311322121-3321323131010302-0032223301223232-3321221123312022-3313031000113100-2313231001111122) |
| `custom_network_config.sli_config.dc_cluster_group.namespace` | [custom_network_config.sli_config.dc_cluster_group.namespace](data-sources--securemesh_site--reference--group-003.md#canonical-3033311110013010-1012302223232100-2210212301223313-1103022222333103-0223210013310102-3102120023033033-2132022103211211-1110321333310110) |
| `custom_network_config.sli_config.dc_cluster_group.tenant` | [custom_network_config.sli_config.dc_cluster_group.tenant](data-sources--securemesh_site--reference--group-003.md#canonical-2212003300210333-1030201133223202-0300021121122123-3333322031002322-3100213020112001-2032330220311023-1313321223222322-3231031000211311) |
| `custom_network_config.sli_config.labels` | [custom_network_config.sli_config.labels](data-sources--securemesh_site--reference--group-003.md#canonical-0020232022012222-1331113231131031-0310203331102000-3031000211201102-3300210010133101-1202333131010002-1031121023313221-1320032111232303) |
| `custom_network_config.sli_config.nameserver` | [custom_network_config.sli_config.nameserver](data-sources--securemesh_site--reference--group-003.md#canonical-1331123212201131-2102011002323312-3123323200010332-0012232233103333-3210332003033233-1202210022101303-3131102232103322-0323220310210221) |
| `custom_network_config.sli_config.no_dc_cluster_group` | [custom_network_config.sli_config.no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-1332302013221103-0200201310332201-1121031320221112-2223100202020231-0022201321021300-2312103211120031-2100323333222332-2330303113012023) |
| `custom_network_config.sli_config.no_static_routes` | [custom_network_config.sli_config.no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3230310013132213-2002201233330311-1030020213233111-2301012111331323-2222132322030220-0023312232113303-0213101012002023-3131212131210202) |
| `custom_network_config.sli_config.no_v6_static_routes` | [custom_network_config.sli_config.no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-1200200111011100-0101312303120211-1213313102212320-1011023132212002-1120133211103013-3133220202113210-3320213112333221-3133023030120122) |
| `custom_network_config.sli_config.static_routes` | [custom_network_config.sli_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2231231200033032-3000233021002020-2032220320010111-3102232203323312-3212131022033203-2021012200003232-0303313223003331-1300232221330203) |
| `custom_network_config.sli_config.static_routes.static_routes` | [custom_network_config.sli_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3221103033130222-1020231203122332-1101230320200200-2002120131102031-0111111130330330-3112020201330231-2313313002230230-3330012031302003) |
| `custom_network_config.sli_config.static_routes.static_routes.attrs` | [custom_network_config.sli_config.static_routes.static_routes.attrs](data-sources--securemesh_site--reference--group-003.md#canonical-0103131211212203-3312301113212012-0230102131322002-1223231233332233-0011310301322221-0221133000102020-1211300303031011-3202103312001132) |
| `custom_network_config.sli_config.static_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-1233013322333232-2112311321223331-3031012220311123-0112310130322112-3321122302322102-0232013120332230-1000031221233120-0302020212202103) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_routes.static_routes.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-0112230100301123-2322032012203323-0021211010302311-0201232030013203-0123310220312001-3103311212100222-2022312111303001-1120202321212130) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_routes.static_routes.ip_prefixes](data-sources--securemesh_site--reference--group-003.md#canonical-3313011323231212-3322330333323231-2033212013121001-0000032312303021-3133221022301223-1000303210220231-0333310133101230-1033312101230012) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-1031010033021031-0231302212001012-1122003321233222-1333232322120000-1013301233003222-0000123122122131-1223030001122023-1011201012333301) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-2330013130122333-2123333021221301-2223002333201013-3223301111030231-2332212321202320-0030021333300000-1230102310102300-1201233023321110) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-2331033110302220-3300032333220302-1331031000312320-1331112231023230-2230321030230200-0311201011221231-3133001332321212-2322102211132302) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--reference--group-003.md#canonical-0100313212123300-0010121120311223-0222013322130102-3111313223201201-3033302110002102-0133113213223021-1010110123132133-1213102021113220) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--reference--group-003.md#canonical-3011002232001130-3013110020123321-2232302200201213-0230321302303202-0012030111000303-2003211021031002-0233030211313323-0021313201321113) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--reference--group-003.md#canonical-3001332020231303-0003100130323132-2202131103023213-1101310333312023-3323010021020011-2113303311203333-2020022122132002-1110033120013022) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--reference--group-003.md#canonical-1212320111013001-2102303330301003-0020232321220133-0313323032132102-0201033101202121-2012211133232331-2013022230201203-3123332210113120) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--reference--group-003.md#canonical-3033111021020020-1300302223301002-0110303021021001-2230322201303111-2112233210300031-0021302211220332-2302331110110321-3011302031021321) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--reference--group-003.md#canonical-2301203302200231-3003021103021121-2002330010112213-3302333220200021-0303102310301333-0100200301131231-2210121322202122-2303032313121300) |
| `custom_network_config.sli_config.static_v6_routes` | [custom_network_config.sli_config.static_v6_routes](data-sources--securemesh_site--reference--group-003.md#canonical-1011312002112110-0130130332021012-3113221003200012-2131203023101232-2210132233201133-1032031320002021-3300313303231333-2232221331100311) |
| `custom_network_config.sli_config.static_v6_routes.static_routes` | [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-1133333030101312-1133202100031013-2003012022130030-2220202020332211-2213111020011011-3231323001100011-0102333201122221-0102032120131301) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.attrs` | [custom_network_config.sli_config.static_v6_routes.static_routes.attrs](data-sources--securemesh_site--reference--group-003.md#canonical-2210100103121300-1233011331301210-0220132330232332-0102312002233213-1000030201300120-1200120112231020-1011310102002103-2123233012210332) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-2133331301002003-2301321203010013-2123131002332302-0320001202323133-0030120203020302-2000330232232103-0122103200012303-0231202123303120) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-2101320010231020-3201210202133222-2120212003311000-1023102113300320-2232010003313010-2103021111213013-3233103220132202-2331003121323213) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes](data-sources--securemesh_site--reference--group-003.md#canonical-3213010123211112-1103332010010020-2132013312100103-2301322130213222-3021101212012321-0012130210110101-3121020112120002-1333001002012202) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-2103002313232112-1313131031213213-2201302300202320-1112001303020233-1310021233230113-2110332232323231-0013210133333232-2312210120200020) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-3232211231021112-3200031110201232-1031222232330332-3231201023100110-2133110013002333-3221210102102031-3201010203130221-1023001312122112) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-1203110011313231-0203121232121101-1010021310130330-1323130032132310-2110202131201231-2322013322133012-0233203020331101-1023231001113102) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--reference--group-003.md#canonical-3220330230031000-3111131023200310-1321210233012010-3102333313111212-1132301320023313-1033302213100022-3332201121312223-2033313100003331) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--reference--group-003.md#canonical-1120323321313012-1311103000211032-0303322201133000-0213030113322300-1203022001202203-2230201211323221-0022031213100102-2130322120000213) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--reference--group-003.md#canonical-0311303010100123-0123003333313112-0210212330212202-2103001131132031-2130013322023223-1232323222123203-0133313031103021-0222002302131201) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--reference--group-003.md#canonical-0130131302303003-1302320211101211-2213233201031321-0202100102123033-0213010233232130-2130021003311113-3021302202201030-2021033202333322) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--reference--group-003.md#canonical-0221302202122021-3130231020301023-0203222120030021-3102013131332113-0133002203232232-3020103022221330-3302331303111113-1123330201332312) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--reference--group-003.md#canonical-3220202121010322-3322310030322221-0022301301232310-2112220202231323-2123210201100001-1201133233132303-0020102311322030-0222223122201103) |
| `custom_network_config.sli_config.vip` | [custom_network_config.sli_config.vip](data-sources--securemesh_site--reference--group-003.md#canonical-0100111223313012-0313210321233023-1132301322012123-1023022130231201-0022223031211113-0312110322320222-3320000000133330-2321211033013032) |
| `custom_network_config.slo_config` | [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-3203230011213303-1200213003033022-0312002200201320-1301113021001202-1202030213233131-2103313030131111-3200310231211010-3120233132212000) |
| `custom_network_config.slo_config.dc_cluster_group` | [custom_network_config.slo_config.dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-2112133333311322-0020211102020101-1003010110102313-0231323232131011-3322131312301331-3013331021200232-0100231230013120-1111221211313001) |
| `custom_network_config.slo_config.dc_cluster_group.name` | [custom_network_config.slo_config.dc_cluster_group.name](data-sources--securemesh_site--reference--group-003.md#canonical-2022112201232222-3131100111132130-1300211012000011-0232232112022201-0020210130131121-0203100012210323-2331012103332003-3002002112110303) |
| `custom_network_config.slo_config.dc_cluster_group.namespace` | [custom_network_config.slo_config.dc_cluster_group.namespace](data-sources--securemesh_site--reference--group-003.md#canonical-2032302231220201-0301331131221111-2333010221313003-1202331113132103-2333120313211012-3020313003310230-2210233022113330-0321031232211112) |
| `custom_network_config.slo_config.dc_cluster_group.tenant` | [custom_network_config.slo_config.dc_cluster_group.tenant](data-sources--securemesh_site--reference--group-003.md#canonical-2001011230300211-1222200003013122-0303100003322310-1012200000112102-2011032113301320-2300200031010031-0121011311300221-2110111202032301) |
| `custom_network_config.slo_config.labels` | [custom_network_config.slo_config.labels](data-sources--securemesh_site--reference--group-003.md#canonical-2210302203212202-1222110322010001-0110013321002222-3132021133030301-0211123001113121-1321111202031032-2330302301032022-0021131100102133) |
| `custom_network_config.slo_config.nameserver` | [custom_network_config.slo_config.nameserver](data-sources--securemesh_site--reference--group-003.md#canonical-2221003221331011-3320132222300020-0100021300131332-1111300003000103-3311011230111033-0301030300132221-1233103211203303-3313211330302321) |
| `custom_network_config.slo_config.no_dc_cluster_group` | [custom_network_config.slo_config.no_dc_cluster_group](data-sources--securemesh_site--reference--group-003.md#canonical-0132312211030001-1331133323030321-1112220112320311-1210112133321021-0330201101110210-1332130300103212-3320233200221320-0201230310022132) |
| `custom_network_config.slo_config.no_static_routes` | [custom_network_config.slo_config.no_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3332113130332130-3032332301301200-1102000113322310-3022202300211330-1111221101331320-0232111132110232-3330020201033233-0002321232303302) |
| `custom_network_config.slo_config.no_v6_static_routes` | [custom_network_config.slo_config.no_v6_static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-1221000212123021-2130332100220221-2303233201303231-2203030023213111-1211002320231020-0022111112033123-0313000301030020-0103202112112003) |
| `custom_network_config.slo_config.static_routes` | [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-2210011211232111-1333210323302323-3100022003031013-0322323220313130-1231011202202010-2022330320223121-0131213000030300-3013220003103020) |
| `custom_network_config.slo_config.static_routes.static_routes` | [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3321332102212130-1332122021312020-3120012232332212-0312311130201202-0000102122010002-0331100102013201-3121021012302232-2333213003121133) |
| `custom_network_config.slo_config.static_routes.static_routes.attrs` | [custom_network_config.slo_config.static_routes.static_routes.attrs](data-sources--securemesh_site--reference--group-003.md#canonical-1132133002032133-2211333100000321-0031231232133213-0103311232322132-2303003131010130-2112010331123313-0332131222130021-2231323121300011) |
| `custom_network_config.slo_config.static_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-003.md#canonical-3321101231101020-3010010021322020-1121322101230021-0300313213110020-1322233331131210-1023223311111013-1013103112302210-1013022131113111) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_routes.static_routes.ip_address](data-sources--securemesh_site--reference--group-003.md#canonical-3133302001003230-2221003220131131-1101031122011212-3311022320023113-1322103203013130-0112020133301221-3211331332023200-1313200100010013) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_routes.static_routes.ip_prefixes](data-sources--securemesh_site--reference--group-003.md#canonical-0001112002313223-2213203203102300-1301333233333302-3123102003111202-3001130230013133-1130003303011032-2313221230323112-2230313013301302) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-0311223002220100-1223212022301101-3110321321330320-1011010200122323-2200321031332332-2231133221031212-3123133010301321-2010000322332001) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-2220023102321031-2121321211111311-0130001332231202-0121333013311031-0011120300221320-2032013021221232-2013213332033312-2331300112000112) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-003.md#canonical-1211212133113102-0023112010301203-1021010111002333-2313202322132321-3331302133120101-2122022323032132-3010210120101020-0323230102331123) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--reference--group-003.md#canonical-2300102000201103-3310311302122211-1120301132300312-0022110231210002-0000131203000321-0222311312010110-0323011110120200-2111003022210001) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--reference--group-003.md#canonical-0330201010332002-2121131101233310-3323301103310112-3002320033313210-0312230211000201-2003231012103231-1133032123102033-2201000102031201) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--reference--group-004.md#canonical-0320001221323213-1020213321002300-3231000210023021-3302312010232312-1220021211111000-3321103122110311-2230033020133001-2113100311313020) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--reference--group-004.md#canonical-3212202321020010-2301111131030120-3111302133103010-0103032023112300-1023013203023111-3023122131123021-3321313132232323-0133331002332021) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--reference--group-004.md#canonical-2203312211320011-3131010323030110-3320011032102223-2332022100101310-2021120023112231-0230122332323313-0231112131210113-3332321233320000) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--reference--group-003.md#canonical-0323313301310110-1223112101312131-1231212202312231-2003220321233223-0222033231003020-3012003321210113-0100130203212200-1103132102222210) |
| `custom_network_config.slo_config.static_v6_routes` | [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1212021001113201-2122021300221120-1000333120100121-2030130032022213-3213000202020201-2101013021122322-1331310002011210-0211012230310102) |
| `custom_network_config.slo_config.static_v6_routes.static_routes` | [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1212001302010000-0210013303012030-1313013113110223-2001110232013332-3200210003000111-1320322301313131-3312013333331302-1012131120122010) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.attrs` | [custom_network_config.slo_config.static_v6_routes.static_routes.attrs](data-sources--securemesh_site--reference--group-004.md#canonical-1302110022332131-3211230330032011-2213032303230232-2203000022103110-2031120130231012-3130011133033311-0231201000330003-3222231021201232) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-004.md#canonical-3022000203031331-0322233331100123-3301300012132303-0233202332133031-1220332333001003-2033202332121233-0230021230223201-3123332222233001) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_address](data-sources--securemesh_site--reference--group-004.md#canonical-3122300222121200-3112223131011333-3212301020303020-2302333111022300-2223023131202210-3223010010220111-1323100302321002-0032312222003012) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes](data-sources--securemesh_site--reference--group-004.md#canonical-2102202003021330-3330130121313323-0020030033330102-2223321233213100-1331302201120122-0121010132001112-3112331320020321-2300320332222230) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-004.md#canonical-3132212311001322-1132003312220233-3013032021000113-1200003003320312-0010303022110012-1032330133300302-3230123030301123-0000210330002120) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-004.md#canonical-1213312003230022-1002111310010103-0120103123303020-0212323023222022-1212301202313103-3021120103012212-3300220003123131-0133103323030120) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-004.md#canonical-2302101220311111-1321231311230213-2122322332200133-0312301323331031-2221111233123100-1323323023033233-1320022200223033-3331331132101310) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind](data-sources--securemesh_site--reference--group-004.md#canonical-2223021223232303-3203120332003231-0231032133332333-0011023303031223-3030120323301022-1003020013211002-1132002210312120-2302330100213232) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name](data-sources--securemesh_site--reference--group-004.md#canonical-1003201120021112-3331223330322323-3321311220022120-1322221120220303-1201233022130311-3002200012020013-0002110201011232-1232002200232233) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](data-sources--securemesh_site--reference--group-004.md#canonical-3133303020320101-3212030112002310-2302023123031203-1230022010333032-1110023200130133-2113222020220103-0110133022203332-0120321203303112) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](data-sources--securemesh_site--reference--group-004.md#canonical-0110230000022231-2101222011022122-1022032310102031-3311231111210313-1010312122130111-1220021022310321-0230021003203222-3332321213201233) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid](data-sources--securemesh_site--reference--group-004.md#canonical-3313031210110312-2021122121303121-0003102223222120-3312033122301203-2332222113110121-3033021230121121-1232230312112033-3322333330000033) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node](data-sources--securemesh_site--reference--group-004.md#canonical-3301230333222303-2132331133123122-0131011002032103-3311332120221110-3022102203200231-1200131111312131-2322000130102311-2002223002330131) |
| `custom_network_config.slo_config.vip` | [custom_network_config.slo_config.vip](data-sources--securemesh_site--reference--group-003.md#canonical-3210021312300201-2310101333123103-2323233222221030-3301211022311111-0001133202023331-2320013031003330-1031231000323130-0221301111212332) |
| `custom_network_config.sm_connection_public_ip` | [custom_network_config.sm_connection_public_ip](data-sources--securemesh_site--reference--group-004.md#canonical-3332101202300013-1323013311020031-1121022330000133-0132231033113030-3323301332301312-3230201120110002-0123201002112222-0330030212331003) |
| `custom_network_config.sm_connection_pvt_ip` | [custom_network_config.sm_connection_pvt_ip](data-sources--securemesh_site--reference--group-004.md#canonical-3020133301223333-2110301201020200-3003003211233011-1210031202111313-0003012210120331-2120102322321023-1112030033031222-3022323023010201) |
| `custom_network_config.tunnel_dead_timeout` | [custom_network_config.tunnel_dead_timeout](data-sources--securemesh_site--reference--group-001.md#canonical-0232233001131112-1230122221121223-0002012123111112-1323113321212212-2230111323222220-0232010300021103-0333021213211301-1112012231122031) |
| `custom_network_config.vip_vrrp_mode` | [custom_network_config.vip_vrrp_mode](data-sources--securemesh_site--reference--group-001.md#canonical-3223013213030120-3221033302130002-1311120130133110-0223302203201302-3233223220323202-0133321302312022-2132021233012332-1021331122031231) |
| `default_blocked_services` | [default_blocked_services](data-sources--securemesh_site--reference--group-004.md#canonical-0222313322110202-3222012232011223-0203132322203021-1303023113330102-3200230321333002-2300333011303130-0231222333102031-0120211113002321) |
| `default_network_config` | [default_network_config](data-sources--securemesh_site--reference--group-004.md#canonical-2122210302010011-3212100213220132-2330130311001112-2011021211102302-2103113231321100-3301332130302322-2031303113101321-2302100012130111) |
| `description` | [description](data-sources--securemesh_site--reference--group-001.md#canonical-1020202211300103-1113100020130322-3223101010321231-0301003120203130-1312121101230132-2232203303302322-2102310223233321-0213222032010321) |
| `id` | [ID](data-sources--securemesh_site--reference--group-001.md#canonical-3000213122000303-2031300013032002-2000320323020130-1321022210201313-3133133211300103-0002322302223212-0032223333003213-2233203332203221) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-0210130110302003-1323303232123313-1003012301010312-0110212302012222-3300012002232211-3113113201323101-0312211031230120-0002200102133233) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-0200320320011133-1013200002032332-1203003303021130-2011103233231121-1332033333200103-1203312030322100-1121322211200332-3330222001020323) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-1303210300033023-0211111222111330-2110102132311233-2313133223010023-2020210203123123-0232321322322133-0222233331023232-3110323110130111) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2101121203233030-3233211222200201-2333220333213111-1112212323233100-2231311330033112-2303122110300103-3210320201102233-2311030023000013) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](data-sources--securemesh_site--reference--group-004.md#canonical-3113111033333320-2313232231333311-1003113213232233-1131021203312231-3133230312211003-3101001132301302-1103111220100231-3011111111221202) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](data-sources--securemesh_site--reference--group-004.md#canonical-1232011112330201-2211013133303221-2211220200030310-1322323032301311-3011133320201200-0022231123100333-0123300212231033-1223200320333202) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](data-sources--securemesh_site--reference--group-004.md#canonical-1302031310102230-2120101033020211-2322111011110011-2220300002120213-2323113001231012-0313033301310112-0221123030010300-0312001322331003) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2321121233112220-1303313023012111-2002030102013122-2200120022131031-3010102013121301-3132202300030103-0222222011313312-1001101233300033) |
| `labels` | [labels](data-sources--securemesh_site--reference--group-001.md#canonical-1100320301231103-0221203211221213-3130200300002311-0311003002201231-1131031031010213-0222232022311110-0003213021123001-0313310330131023) |
| `log_receiver` | [log_receiver](data-sources--securemesh_site--reference--group-004.md#canonical-0102131320131032-2101231003111113-2031221232312101-1200101123013210-3000313023310220-2111010331000113-2203120313023102-1030202021202023) |
| `log_receiver.name` | [log_receiver.name](data-sources--securemesh_site--reference--group-004.md#canonical-2132013330000302-1030232220023313-0231200221002311-2231322013201010-3213332003133010-1100122012101330-3213330003333101-1311102120013010) |
| `log_receiver.namespace` | [log_receiver.namespace](data-sources--securemesh_site--reference--group-004.md#canonical-2213013121110112-1200202331022131-2003020231330032-3332210121001310-1300112233130230-0113110102133212-3221300110322210-2032031033303010) |
| `log_receiver.tenant` | [log_receiver.tenant](data-sources--securemesh_site--reference--group-004.md#canonical-1301220211333132-0002202021321333-0120033212203330-1223120212010132-3211310213223011-1020220230121110-3222110131110113-2233211003313220) |
| `logs_streaming_disabled` | [logs_streaming_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-0332333310211122-2213330231333323-0120012320213233-0023230002030030-2020002103301010-3200213101010323-0233111023320031-1313211021030121) |
| `master_node_configuration` | [master_node_configuration](data-sources--securemesh_site--reference--group-004.md#canonical-2230020233110202-2201330322320011-0303301122202110-1022320121131032-1033112133022203-1001000213112231-1000023203102100-0023321330011213) |
| `master_node_configuration.name` | [master_node_configuration.name](data-sources--securemesh_site--reference--group-004.md#canonical-2110013133030011-1023101101020022-0031312320300221-1213102300022012-1332210301302022-2310222132003200-1001221033003030-0031031000122122) |
| `master_node_configuration.public_ip` | [master_node_configuration.public_ip](data-sources--securemesh_site--reference--group-004.md#canonical-3202110112223130-0030231332133011-3203333301122100-1122013210201310-2310222322030321-1221131010233103-1303012132332310-2300332333021311) |
| `name` | [name](data-sources--securemesh_site--reference--group-001.md#canonical-3123223031003332-0331100130321123-3332223232231101-2320103320330012-2131132032131112-2212300112330000-2020133113031202-3212202030222321) |
| `namespace` | [namespace](data-sources--securemesh_site--reference--group-001.md#canonical-2010231111032120-0112222231103113-1131303023013113-1021001132102001-1011223231133213-1131300222220030-2320110231201130-1332211323101212) |
| `no_bond_devices` | [no_bond_devices](data-sources--securemesh_site--reference--group-004.md#canonical-2213003202010300-2102012200213123-2322022103020000-1020020123123221-1230331011321230-0331220100231301-0230110030232233-3132321222002122) |
| `offline_survivability_mode` | [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2330312322333212-3232020001213131-2111111121000303-2133200333233303-3112331203023332-3112300310000302-3102001112330030-2010132100312213) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-3122230113320022-2313310211310210-2203002021223322-2010303101010301-2320303003212113-0333332322102300-1211031132013321-2221201320220020) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-1231133110100020-2211231012200221-1013122123211120-0203330231230011-2303003123010013-0102101310200101-1122130031013231-2203121333222032) |
| `os` | [os](data-sources--securemesh_site--reference--group-004.md#canonical-0330232302010030-3230130021221032-0310020122222321-1301301322103031-0313211331012012-1313002103013332-3111030003211221-2331333002303002) |
| `os.default_os_version` | [os.default_os_version](data-sources--securemesh_site--reference--group-004.md#canonical-1122011120133032-1313020121210133-3021011022202333-3212221120001221-3312021202301211-3030213222221122-0113312131102123-0232201112011123) |
| `os.operating_system_version` | [os.operating_system_version](data-sources--securemesh_site--reference--group-004.md#canonical-1312231230132320-3202210103230302-0100232030221131-2232320102123000-1302312201110233-1232100022333320-0310232121300221-0213112301223133) |
| `performance_enhancement_mode` | [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2220211101203300-1331123022300032-3023223032213031-1300220112003022-1021132202112331-2103202312112321-1312033132112010-2033330103020020) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-2032111302222332-2032030030200300-3232230200222223-3301222311012332-3013331233321212-1213113010023311-3201332023002123-3023333132313030) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--securemesh_site--reference--group-004.md#canonical-3102130121131230-2102010201021121-1233223011312322-2321023033300223-3012203310310130-1210030030313120-1012332233021311-1202202201120120) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--securemesh_site--reference--group-004.md#canonical-2031200302011233-1311133202120102-1303312300230010-2003002303111032-3300320132000303-2013010223122310-3130100112303211-1223121211111331) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-1332200111030101-3110330210101333-2322321022212310-1313130330310013-0300113123221332-1322031023102322-0323213123233003-3302312022210113) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-3231222133032010-3311120001311132-2331212131312333-1021331020330201-2112330100230130-0100130212333220-1231001320212332-0011131212102321) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--securemesh_site--reference--group-004.md#canonical-1113021133323323-0013333110230331-1332003111302233-3020113321023003-1230312001332021-0030203032332031-2312223313001223-2022033133231210) |
| `sw` | [sw](data-sources--securemesh_site--reference--group-004.md#canonical-0031120033012022-2031213033001211-2103110033322020-2310001210120001-3201013202113103-0102313321033031-1303113132330001-2221201130201031) |
| `sw.default_sw_version` | [sw.default_sw_version](data-sources--securemesh_site--reference--group-004.md#canonical-3210231311203130-3231111231311130-0102211200322000-1311121221011211-2300202303001302-1020122202023221-3221110302231103-1320131031101031) |
| `sw.volterra_software_version` | [sw.volterra_software_version](data-sources--securemesh_site--reference--group-004.md#canonical-3112113222112330-3131123000331302-2113101132230321-2030023333030123-2033233230100202-2121003230211133-1023202001023123-1033331210201321) |
| `volterra_certified_hw` | [volterra_certified_hw](data-sources--securemesh_site--reference--group-001.md#canonical-1200001331121323-2000312231301022-1021131121133231-2203101233230130-2002023302203220-3300202330010102-3001000032303102-1022301010212002) |
| `waf_signatures` | [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-2113100111102311-1232323133020231-0002221002002112-3312133210103211-3221032113022331-1121312103302102-1130023311203103-1003312033313130) |
| `waf_signatures.automatic` | [waf_signatures.automatic](data-sources--securemesh_site--reference--group-004.md#canonical-3123320331231112-0021013113100132-2000323020112012-0020133012031110-1233313003331030-2321322203021022-1212102020320211-0203200300133331) |
| `waf_signatures.manual` | [waf_signatures.manual](data-sources--securemesh_site--reference--group-004.md#canonical-0031200001221223-3322031022111201-2112110313030123-1300120012202210-0103222301302202-3133012233033112-2023002130112320-3202011230111201) |
| `worker_nodes` | [worker_nodes](data-sources--securemesh_site--reference--group-001.md#canonical-1310131033332013-0202121123323330-0102113222101230-0102212013321312-2232110300310023-3323113211023021-3230000011311301-0203221321312013) |

<a id="canonical-0213132232322322-2022220322110202-3132130221032221-1021100311002112-3313130120100023-1033311222322312-1220321320223113-3302230011000232"></a>

## Next pages — Property reference / 111013232123 / 14

- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-0230212222332032-2133213333223020-2203212003102330-2322011323032333-1322302002100331-3032210130333303-0102011202012330-1010333331021220)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-2010013132202033-2110133330011130-1310012103113012-3100032310101130-3003301111111022-3001111033212101-0321233100220322-1323121313313301)
- [coordinates](data-sources--securemesh_site--reference--group-001.md#canonical-1011033201210123-3012222111003002-1033110130112121-0220000011220021-0121202312300300-0303330130211101-2021110123002101-3330010310012223)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [default_blocked_services](data-sources--securemesh_site--reference--group-004.md#canonical-1332002011330322-3112302133232021-3020303330123333-0121332303210202-2011220212000322-3013133110313113-2300320313113130-0311122221021323)
- [default_network_config](data-sources--securemesh_site--reference--group-004.md#canonical-1032132001130300-2232230303121130-2322233123010303-1320223100112113-0002203210303003-1010323303201032-2221222230312013-0301020013000030)
- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2003312033033001-0003322123103200-0003210200302311-2323013312221021-1032032302211212-2033100222113033-3311011230311102-0221123310333312)
- [log_receiver](data-sources--securemesh_site--reference--group-004.md#canonical-0211220000103000-2211130101211031-0220003100131103-2300012120232311-0013332232122023-1233120111001231-1112013003212021-3201020203022120)
- [logs_streaming_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-2030232101301321-3022120220023011-2330231031223312-0201232010301230-3330312010210010-0011010130130130-3212310310100011-2332213030110233)
- [master_node_configuration](data-sources--securemesh_site--reference--group-004.md#canonical-3313100220202200-1000302100010131-1301231302121331-1300120222103101-0133123100113323-2120331223001002-2101030221002332-2323223232331200)
- [no_bond_devices](data-sources--securemesh_site--reference--group-004.md#canonical-1003113100300311-2222313300001010-2021310302002030-1112332132120233-1002033303013130-2231330222100030-0102321101101202-1321023010332110)
- [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2020103213000102-1110301232330123-2322321023321200-2131023113310313-2233322222100111-2032312323132100-0321122012303013-0320311231012313)
- [os](data-sources--securemesh_site--reference--group-004.md#canonical-3012121031011130-1233001032220331-3130122211200032-2131001233320133-0031212331013332-2032112211110320-0203323010102313-1032213120331101)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331)
- [sw](data-sources--securemesh_site--reference--group-004.md#canonical-1101033231030033-1231112223010232-0121200330001132-1331111222212213-3023302310111332-0331233101000132-0230010020103302-3320133222123131)
- [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-0301131103113000-3303232313310330-0122211113001113-3130222002332001-1310200111021022-2330131102030223-3102231232233231-3301133132020110)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0230212222332032-2133213333223020-2203212003102330-2322011323032333-1322302002100331-3032210130333303-0102011202012330-1010333331021220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113023121303133-0132013310021221-2122020131202231-2000201332332321-2211131312021012-1030032010231010-3021332211202232-1112100210303021"></a>

## blocked_services — blocked_services / 012302103211 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- blocked_services

<a id="canonical-2032103203120130-3220033100222121-2102023221121102-3220133021022303-2130033131221111-2233000121213101-1132202212331333-3122101200303103"></a>

Type: `"single"`. Computed.

\[OneOf: blocked\_services, default\_blocked\_services; Default: default\_blocked\_services\]
Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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

- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-2032103203120130-3220033100222121-2102023221121102-3220133021022303-2130033131221111-2233000121213101-1132202212331333-3122101200303103)
- [default_blocked_services](data-sources--securemesh_site--reference--group-004.md#canonical-0222313322110202-3222012232011223-0203132322203021-1303023113330102-3200230321333002-2300333011303130-0231222333102031-0120211113002321)

Select alternatives according to the provider validators above.

<a id="canonical-1021001311321323-1122333123202031-1301310332102020-3101020102131013-0030031333022301-0233012123200213-3212310231213132-2032310013220331"></a>

## Direct properties — blocked_services / 012302103211 / 3

- [blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111): complete subsection reference.

<a id="canonical-1021030032222121-2310213321233123-2222312111132331-3322003332221131-1203022102312102-3300000021211020-3030131231311313-1102201330112000"></a>

## Next pages — blocked_services / 012302103211 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210210213222031-0012311022133201-0333233221232101-2023101222010220-1101113331001231-3031233010223002-1103102011213203-3210113312302310"></a>

## blocked_services.blocked_service — blocked_service / 120230311023 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-0230212222332032-2133213333223020-2203212003102330-2322011323032333-1322302002100331-3032210130333303-0102011202012330-1010333331021220)
- blocked_services.blocked_service

<a id="canonical-3033120300133112-2112112020312200-3330332201200303-2021301000011301-0113322310122322-2210013310013113-3333311220222132-3122230310202321"></a>

Type: `"list"`. Computed.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1231203023303230-1321322300231210-2131220212222113-2130223313210131-1020132320032301-1132210302300221-3231223102221322-1101033133002023"></a>

## Direct properties — blocked_service / 120230311023 / 3

- [DNS](data-sources--securemesh_site--reference--group-001.md#canonical-1131021102112331-3032331332323223-3122122202211322-0213122000332000-0120111132323032-1103330010003001-2100122320102103-2102032213310133): complete subsection reference.

<a id="canonical-3300121033020031-2321321222231130-1030130011211030-1312120303311101-3001331231103131-3202013130321332-0223130013222210-1101101100233233"></a>

<a id="canonical-0000021300213333-2113011002310331-3023001002000221-1020200221121130-1032210311133010-0123200130123130-3313110313130021-1320201103333023"></a>

## network_type property — blocked_service / 120230311023 / 4

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [SSH](data-sources--securemesh_site--reference--group-001.md#canonical-0232101031313012-2330312230001023-1332032110333332-0200103320131330-1022303231323301-3132031211113122-2232303321133132-0003100233000111): complete subsection reference.

- [web_user_interface](data-sources--securemesh_site--reference--group-001.md#canonical-0311301232112313-3021021211223123-1211221301233221-0222313322332120-0031320031113223-0130032220010112-3222120202230001-1321103111003200): complete subsection reference.

<a id="canonical-2023001022030013-3310001322231212-0003323203110202-1210030132002333-1220110110202202-0133003102202313-0130313320222332-3112112030113222"></a>

## Next pages — blocked_service / 120230311023 / 5

- [blocked_services.blocked_service.dns](data-sources--securemesh_site--reference--group-001.md#canonical-1131021102112331-3032331332323223-3122122202211322-0213122000332000-0120111132323032-1103330010003001-2100122320102103-2102032213310133)
- [blocked_services.blocked_service.ssh](data-sources--securemesh_site--reference--group-001.md#canonical-0232101031313012-2330312230001023-1332032110333332-0200103320131330-1022303231323301-3132031211113122-2232303321133132-0003100233000111)
- [blocked_services.blocked_service.web_user_interface](data-sources--securemesh_site--reference--group-001.md#canonical-0311301232112313-3021021211223123-1211221301233221-0222313322332120-0031320031113223-0130032220010112-3222120202230001-1321103111003200)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-0230212222332032-2133213333223020-2203212003102330-2322011323032333-1322302002100331-3032210130333303-0102011202012330-1010333331021220)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1131021102112331-3032331332323223-3122122202211322-0213122000332000-0120111132323032-1103330010003001-2100122320102103-2102032213310133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323113223302022-0323300231332231-0311131311322100-1302012020221130-3122012233112220-0210233010100232-1022233020100030-3300230331030301"></a>

## blocked_services.blocked_service.DNS — DNS / 001301122312 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-0230212222332032-2133213333223020-2203212003102330-2322011323032333-1322302002100331-3032210130333303-0102011202012330-1010333331021220)
- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111)
- blocked_services.blocked_service.DNS

<a id="canonical-3333112331020021-3203232320322011-0020211131230231-0212100113313020-0102311121323230-3101111223011132-1302310020313321-3211111231001013"></a>

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

<a id="canonical-2103223003323103-1010111331301033-0002303213302203-2213233232011320-3233103301013001-0112021103213123-3221001231121122-1012033130110033"></a>

## Direct properties — DNS / 001301122312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223213002013312-3313001310113313-0011031130330102-0123021011203332-2310210130203303-1102232303320021-1310221031110310-0001113220233110"></a>

## Next pages — DNS / 001301122312 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0232101031313012-2330312230001023-1332032110333332-0200103320131330-1022303231323301-3132031211113122-2232303321133132-0003100233000111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033311100132002-2030123002032202-3112121033322231-1203001131300132-3333302201323222-1100222222330132-1031330000033211-3021122230132302"></a>

## blocked_services.blocked_service.SSH — SSH / 201222031021 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-0230212222332032-2133213333223020-2203212003102330-2322011323032333-1322302002100331-3032210130333303-0102011202012330-1010333331021220)
- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111)
- blocked_services.blocked_service.SSH

<a id="canonical-0113102000201120-0211013131300323-3210333313230121-1022223331110202-3211131103100310-1211110001120300-3003132000231231-2322230001233031"></a>

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

<a id="canonical-0320232300020032-1201302103131321-2231331123211103-3111321110302321-3033133330022123-0001332200233003-3301110320010303-1330131231203012"></a>

## Direct properties — SSH / 201222031021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110233302021131-2133232323313101-3301311012232020-3033021222232313-0213201023333203-1110012131221133-1301321200003231-0130131112202302"></a>

## Next pages — SSH / 201222031021 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0311301232112313-3021021211223123-1211221301233221-0222313322332120-0031320031113223-0130032220010112-3222120202230001-1321103111003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303312103012102-2320030032011110-0133321130323002-1211031213001203-1333003221021002-1031313103122313-3300102330231130-0233030122301202"></a>

## blocked_services.blocked_service.web_user_interface — web_user_interface / 321000012132 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [blocked_services](data-sources--securemesh_site--reference--group-001.md#canonical-0230212222332032-2133213333223020-2203212003102330-2322011323032333-1322302002100331-3032210130333303-0102011202012330-1010333331021220)
- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-2200333102320321-2323003001111323-3023121331213303-2232323130311220-2220120001331031-2200330221202233-3030332033321333-3200203301331211"></a>

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

<a id="canonical-2222311121113031-3132322031333320-2110012023012231-0120131022213213-2232001302132033-3233003302313311-1012112032320301-3003222222201103"></a>

## Direct properties — web_user_interface / 321000012132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202213203233200-0112302023110220-0020100101300111-2212312011222230-0011003013331223-1313313223022103-1121103031202323-0031112300002131"></a>

## Next pages — web_user_interface / 321000012132 / 4

- [blocked_services.blocked_service](data-sources--securemesh_site--reference--group-001.md#canonical-2332311030120320-2033123303323230-3223233232303313-2002301313321213-2211310122011321-1222201120122312-3232302211021302-0033020222011111)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2010013132202033-2110133330011130-1310012103113012-3100032310101130-3003301111111022-3001111033212101-0321233100220322-1323121313313301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203210313332011-3322131132111222-3100213120132030-1033011222000210-3233030122322233-0131002201203030-3320332011102233-3001131112122220"></a>

## bond_device_list — bond_device_list / 221313130020 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- bond_device_list

<a id="canonical-2013100011013002-1322321012232320-3131331223201021-1233321301001010-0302132332033003-3122030000322130-3012320010232210-1313213002003321"></a>

Type: `"single"`. Computed.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

Upstream description:

List of bond devices for this fleet.

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

- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-2013100011013002-1322321012232320-3131331223201021-1233321301001010-0302132332033003-3122030000322130-3012320010232210-1313213002003321)
- [no_bond_devices](data-sources--securemesh_site--reference--group-004.md#canonical-2213003202010300-2102012200213123-2322022103020000-1020020123123221-1230331011321230-0331220100231301-0230110030232233-3132321222002122)

Select alternatives according to the provider validators above.

<a id="canonical-1001011231220202-2132311233233130-1023210032212321-3311203000103222-3200022201122000-2303323120303320-2131233112102311-0330033221013101"></a>

## Direct properties — bond_device_list / 221313130020 / 3

- [bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-0211000112232223-1200320332313222-3003332031311010-0313012123303323-1313110123023210-3230102000112310-0202203321310320-2033203100320033): complete subsection reference.

<a id="canonical-0222102232100113-1031031211001211-1220001023032333-0300333000333331-1212032202303322-2103232330102121-3322332002101230-3000021201132333"></a>

## Next pages — bond_device_list / 221313130020 / 4

- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-0211000112232223-1200320332313222-3003332031311010-0313012123303323-1313110123023210-3230102000112310-0202203321310320-2033203100320033)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0211000112232223-1200320332313222-3003332031311010-0313012123303323-1313110123023210-3230102000112310-0202203321310320-2033203100320033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030033211303012-0230113102020100-3111221032232330-2320211120333102-1231210002201012-1112112131023320-2133121301033301-2231131230221020"></a>

## bond_device_list.bond_devices — bond_devices / 103001103000 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-2010013132202033-2110133330011130-1310012103113012-3100032310101130-3003301111111022-3001111033212101-0321233100220322-1323121313313301)
- bond_device_list.bond_devices

<a id="canonical-3120001323121310-2000002020130202-3203230020222022-3023100012033200-3123022332303010-2322101121013112-0322212302120111-3333203313111102"></a>

Type: `"list"`. Computed.

Bond Devices. List of bond devices.

Upstream description:

List of bond devices.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1020310032312200-2020002311323223-0020012330131302-1210301331122022-0220310332103203-2032201311210022-1131102033212220-3200130333112300"></a>

## Direct properties — bond_devices / 103001103000 / 3

- [active_backup](data-sources--securemesh_site--reference--group-001.md#canonical-1102111301023302-1031122213000132-3112323230002001-2223321110331313-0121300033001032-3002300132321102-3213033312310011-0311031212000103): complete subsection reference.

<a id="canonical-0103011200023021-0322102223220233-1303112232210123-2310012322103310-1323333112110301-0102220123211103-1313212021002000-1332301231213323"></a>

<a id="canonical-1213212223013222-1010231200202111-2033220210102322-0022130313233201-0001011303312112-3032222023033223-3331002123132023-2001302030110301"></a>

## devices property — bond_devices / 103001103000 / 4

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

- [lacp](data-sources--securemesh_site--reference--group-001.md#canonical-3130112203023000-2130231101312022-1103112001033330-2202313331100321-0113313301133223-1122021331212031-3132032320210100-0332013103323010): complete subsection reference.

<a id="canonical-3031103203212002-2113301102021103-0313113200103021-3301333111121031-1211300220123230-2221313122122211-3110313112032020-1122111001331023"></a>

<a id="canonical-3031032313312021-0312320021011301-3101100030032300-0200212013302331-1021110011232223-0313022222311303-3313130313210210-1300312002213032"></a>

## link_polling_interval property — bond_devices / 103001103000 / 5

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

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

<a id="canonical-3222211123020122-0032013030132233-0003031102020121-3332320112100010-3211310102211011-1000001110312030-2333113330121201-1311113203010031"></a>

<a id="canonical-1022200131032211-0223102131002233-0003011301021211-2301102232213001-2101022203120231-0320011303332102-1111332203100022-2130210000123110"></a>

## link_up_delay property — bond_devices / 103001103000 / 6

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

<a id="canonical-2231230101123012-1232332001210131-3332132203111032-2132233303302110-3303212200230002-0122013320211111-3101312032002121-0112330000233311"></a>

<a id="canonical-0310222103302220-1021113230313021-0023330101003301-2232312120122323-1312323021320033-3200023021122333-1122200303030010-0131133011032233"></a>

## name property — bond_devices / 103001103000 / 7

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

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

<a id="canonical-2302101133012011-0021012010303331-3331223312120030-1212233122032202-3300213102332110-1201330201023321-0230310313013311-1233222000230020"></a>

## Next pages — bond_devices / 103001103000 / 8

- [bond_device_list.bond_devices.active_backup](data-sources--securemesh_site--reference--group-001.md#canonical-1102111301023302-1031122213000132-3112323230002001-2223321110331313-0121300033001032-3002300132321102-3213033312310011-0311031212000103)
- [bond_device_list.bond_devices.lacp](data-sources--securemesh_site--reference--group-001.md#canonical-3130112203023000-2130231101312022-1103112001033330-2202313331100321-0113313301133223-1122021331212031-3132032320210100-0332013103323010)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-2010013132202033-2110133330011130-1310012103113012-3100032310101130-3003301111111022-3001111033212101-0321233100220322-1323121313313301)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1102111301023302-1031122213000132-3112323230002001-2223321110331313-0121300033001032-3002300132321102-3213033312310011-0311031212000103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220110211003033-3201312211203013-3032112001332232-0320012101000320-2330010331330123-3210311010202120-2331113331202221-1002230030322210"></a>

## bond_device_list.bond_devices.active_backup — active_backup / 031033303313 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-2010013132202033-2110133330011130-1310012103113012-3100032310101130-3003301111111022-3001111033212101-0321233100220322-1323121313313301)
- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-0211000112232223-1200320332313222-3003332031311010-0313012123303323-1313110123023210-3230102000112310-0202203321310320-2033203100320033)
- bond_device_list.bond_devices.active_backup

<a id="canonical-3000322220121222-2001002212321031-2110321331201311-3113021323312020-1332232221220011-3011013131222320-1221233002021222-1232330333123001"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2303000310322320-0012312320332101-3302213212111022-0100033330200030-2232201000333313-1313012232132033-3202001130132013-0011033213131223"></a>

## Direct properties — active_backup / 031033303313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312310311123123-1030203112321132-0302322112222210-3033213233221013-1021032332002121-3200130213300220-2021321033320021-1033212303223003"></a>

## Next pages — active_backup / 031033303313 / 4

- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-0211000112232223-1200320332313222-3003332031311010-0313012123303323-1313110123023210-3230102000112310-0202203321310320-2033203100320033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3130112203023000-2130231101312022-1103112001033330-2202313331100321-0113313301133223-1122021331212031-3132032320210100-0332013103323010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121123022013012-1313313001111030-3211232313022013-1011221111221300-2222302012331201-1030131322220030-1020130112120320-1311200030330213"></a>

## bond_device_list.bond_devices.lacp — lacp / 003330322103 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [bond_device_list](data-sources--securemesh_site--reference--group-001.md#canonical-2010013132202033-2110133330011130-1310012103113012-3100032310101130-3003301111111022-3001111033212101-0321233100220322-1323121313313301)
- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-0211000112232223-1200320332313222-3003332031311010-0313012123303323-1313110123023210-3230102000112310-0202203321310320-2033203100320033)
- bond_device_list.bond_devices.lacp

<a id="canonical-3220002101223012-2132212011311121-3132030001001320-1202310133120330-1013201321302312-1001333300102111-3212310100033111-3031022313000300"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

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

<a id="canonical-0333302223313321-0102230313130203-3331130232022131-1331302101303102-2001330301230113-0101030132303310-0203002001102211-1322020121000322"></a>

## Direct properties — lacp / 003330322103 / 3

<a id="canonical-3032012230300202-3223112332320202-0232030232131123-1221200020300100-1313020233313200-1113312002103111-2123211300021012-1222021222110032"></a>

<a id="canonical-2323012031002232-2331232012331012-2210010100202221-0000121213212221-1111311132003001-1300123222031033-0021121330023132-3133113301022300"></a>

## rate property — lacp / 003330322103 / 4

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

<a id="canonical-3000020223330032-2110033010232121-3122323202200330-0122123233323121-0301030010210222-0310331021223220-1312030132011233-2012001210130123"></a>

## Next pages — lacp / 003330322103 / 5

- [bond_device_list.bond_devices](data-sources--securemesh_site--reference--group-001.md#canonical-0211000112232223-1200320332313222-3003332031311010-0313012123303323-1313110123023210-3230102000112310-0202203321310320-2033203100320033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1011033201210123-3012222111003002-1033110130112121-0220000011220021-0121202312300300-0303330130211101-2021110123002101-3330010310012223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101120302231221-1313133130201200-2303000111320123-3123002130110321-0113232132301021-2010002132311222-3003132023130330-0312223103310000"></a>

## coordinates — coordinates / 133000301312 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- coordinates

<a id="canonical-0310002233301332-0203132332010333-3221203003210002-0030100102320021-2323200302012321-2303331201310022-2033000313321221-3022101002231313"></a>

Type: `"single"`. Computed.

Coordinates of the site which provides the site physical location.

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

<a id="canonical-2032122012110222-2210201122212311-3222221120303102-0010312033100132-0200333213221301-0123312311023220-0100103211330320-0312313310331223"></a>

## Direct properties — coordinates / 133000301312 / 3

<a id="canonical-3322213320332223-1101123103012303-3023113312332333-2310230230120313-3131330123223010-3222122031132133-3323231311113032-1210303310301010"></a>

<a id="canonical-1033301233213223-0300233112112002-3231310120022123-1110203121211110-2010233010222301-3321301321213032-3030202302303212-3003103130013110"></a>

## latitude property — coordinates / 133000301312 / 4

Type: `"number"`. Computed.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-0311002313021301-1132301012103333-0212110113122101-2233303211132211-1201223002221013-1122322032211030-2030102321230233-2331123300120132"></a>

<a id="canonical-1203002233000320-2020231232223312-3130323300332110-0031112111322003-3030133311100101-0001220332010132-0233121102232302-2332310031123112"></a>

## longitude property — coordinates / 133000301312 / 5

Type: `"number"`. Computed.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-0213132003122002-2211123131123321-1023320211120303-3011331033120131-2122331322221321-2331323211110213-3331002021010212-1122030330323001"></a>

## Next pages — coordinates / 133000301312 / 6

- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220110311130220-2313122002212323-2301221222131130-3000202123223330-2210123222122022-2013023221232111-1203012200133312-2010032222000032"></a>

## custom_network_config — custom_network_config / 012212032303 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- custom_network_config

<a id="canonical-2233320021012133-3133111313320122-2033222313102002-1133023232002203-1233333201211020-0101301313301031-3010033010333012-2323223323102013"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_network\_config, default\_network\_config; Default: default\_network\_config\]
SmsNetworkConfiguration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-interface_choice": "[\"default_interface_config\",\"interface_list\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

OneOf alternatives in this subsection:

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-2233320021012133-3133111313320122-2033222313102002-1133023232002203-1233333201211020-0101301313301031-3010033010333012-2323223323102013)
- [default_network_config](data-sources--securemesh_site--reference--group-004.md#canonical-2122210302010011-3212100213220132-2330130311001112-2011021211102302-2103113231321100-3301332130302322-2031303113101321-2302100012130111)

Select alternatives according to the provider validators above.

<a id="canonical-3103032000310222-2123123331223330-3032013102303312-3303113013133322-0323212133133110-0332331032003132-0232103323210110-1111211220023203"></a>

## Direct properties — custom_network_config / 012212032303 / 3

- [active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2121211033231322-2203220213000332-2202233212321213-0130121222023200-3311013231122020-0231000222212011-3020201030020002-0313223323032033): complete subsection reference.

- [active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-1133131032301110-1111213131002033-1212103011002311-0322211003112121-0023301012130222-3322123132312112-1030310022302011-2223001302001033): complete subsection reference.

- [active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2233302301301023-1033312232031023-0130313233310323-1030010202322132-0130012202331122-2310303231101102-3023000233010220-0211230331203321): complete subsection reference.

- [default_config](data-sources--securemesh_site--reference--group-001.md#canonical-0213010111103030-2232130120020312-2232212233012303-0130022232321200-2210311331132301-3230130121221200-3003012201333332-3010000023301103): complete subsection reference.

- [default_interface_config](data-sources--securemesh_site--reference--group-001.md#canonical-0221331032223311-1132302203121312-0021102302131133-3322101113132303-1300102233223222-1000010333010222-3231122200220010-0021330001132102): complete subsection reference.

- [default_sli_config](data-sources--securemesh_site--reference--group-002.md#canonical-3020220113203132-3201033313310101-2303223130323211-1312212103100221-1033222131230102-2211300031022203-3010220131212010-2100113013003232): complete subsection reference.

- [forward_proxy_allow_all](data-sources--securemesh_site--reference--group-002.md#canonical-0021110230123011-0002322300203323-2221012133322223-1230013203032231-0311333131023102-1000033021012331-1102030212000203-3033322220123220): complete subsection reference.

- [global_network_list](data-sources--securemesh_site--reference--group-002.md#canonical-1231123132132331-2223121303300333-2000322321111311-3130320102120130-1001133101023002-0220333313131302-2232221010203001-1012133002033103): complete subsection reference.

- [interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120): complete subsection reference.

- [no_forward_proxy](data-sources--securemesh_site--reference--group-003.md#canonical-2211132100230232-0332032112001113-2331331231012012-3320322002220323-0120102222300321-3113101202102323-2322232121101001-2230333201203121): complete subsection reference.

- [no_global_network](data-sources--securemesh_site--reference--group-003.md#canonical-3203231233310210-3312303100312211-0212030103112311-2211012220123012-0310031332022133-2222302312121122-2022301033230003-3123301021221321): complete subsection reference.

- [no_network_policy](data-sources--securemesh_site--reference--group-003.md#canonical-0211120021231100-0200211311001101-0221332130221012-3033122221023220-0003033302132322-2132311332201133-2020001113203320-1003003232002322): complete subsection reference.

- [sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033): complete subsection reference.

- [slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001): complete subsection reference.

- [sm_connection_public_ip](data-sources--securemesh_site--reference--group-004.md#canonical-3202030221310100-2332312320000312-2311123323010011-3123113311220311-0110233233302131-3330002323231000-1132213111220103-0102232030213320): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--securemesh_site--reference--group-004.md#canonical-2020320120002130-2233001122130003-0032113133002121-1310301132212321-2221222033132221-1323303312320211-0313232000130213-3301100230311111): complete subsection reference.

<a id="canonical-0232233001131112-1230122221121223-0002012123111112-1323113321212212-2230111323222220-0232010300021103-0333021213211301-1112012231122031"></a>

<a id="canonical-0000230320231210-0301021123230003-1200233311231121-3220220001301312-2013230120333001-2213231213120021-3310001120101032-1220302221300100"></a>

## tunnel_dead_timeout property — custom_network_config / 012212032303 / 4

Type: `"number"`. Computed.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Upstream description:

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180000,
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
    "ves.io.schema.rules.uint32.lte": "180000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  }
}
```

<a id="canonical-3223013213030120-3221033302130002-1311120130133110-0223302203201302-3233223220323202-0133321302312022-2132021233012332-1021331122031231"></a>

<a id="canonical-0210113320312123-3321330202133031-2123000213110110-3011231311330230-2322010203030103-1301022003210110-2322320202310213-1122023110010222"></a>

## vip_vrrp_mode property — custom_network_config / 012212032303 / 5

Type: `"string"`. Computed.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Upstream description:

VRRP advertisement mode for VIP

Invalid VRRP mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIP_VRRP_INVALID",
  "enum": [
    "VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023213113000002-3332013112010123-2331131210122033-1120201011122232-1111120330213033-2210111102102202-1103023221100100-3312233233310000"></a>

## Next pages — custom_network_config / 012212032303 / 6

- [custom_network_config.active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2121211033231322-2203220213000332-2202233212321213-0130121222023200-3311013231122020-0231000222212011-3020201030020002-0313223323032033)
- [custom_network_config.active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-1133131032301110-1111213131002033-1212103011002311-0322211003112121-0023301012130222-3322123132312112-1030310022302011-2223001302001033)
- [custom_network_config.active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2233302301301023-1033312232031023-0130313233310323-1030010202322132-0130012202331122-2310303231101102-3023000233010220-0211230331203321)
- [custom_network_config.default_config](data-sources--securemesh_site--reference--group-001.md#canonical-0213010111103030-2232130120020312-2232212233012303-0130022232321200-2210311331132301-3230130121221200-3003012201333332-3010000023301103)
- [custom_network_config.default_interface_config](data-sources--securemesh_site--reference--group-001.md#canonical-0221331032223311-1132302203121312-0021102302131133-3322101113132303-1300102233223222-1000010333010222-3231122200220010-0021330001132102)
- [custom_network_config.default_sli_config](data-sources--securemesh_site--reference--group-002.md#canonical-3020220113203132-3201033313310101-2303223130323211-1312212103100221-1033222131230102-2211300031022203-3010220131212010-2100113013003232)
- [custom_network_config.forward_proxy_allow_all](data-sources--securemesh_site--reference--group-002.md#canonical-0021110230123011-0002322300203323-2221012133322223-1230013203032231-0311333131023102-1000033021012331-1102030212000203-3033322220123220)
- [custom_network_config.global_network_list](data-sources--securemesh_site--reference--group-002.md#canonical-1231123132132331-2223121303300333-2000322321111311-3130320102120130-1001133101023002-0220333313131302-2232221010203001-1012133002033103)
- [custom_network_config.interface_list](data-sources--securemesh_site--reference--group-002.md#canonical-3103221231221110-0233010032301102-3331303332300233-3130201203333213-3000102101320231-1132310313012232-3212133313120012-0033201122232120)
- [custom_network_config.no_forward_proxy](data-sources--securemesh_site--reference--group-003.md#canonical-2211132100230232-0332032112001113-2331331231012012-3320322002220323-0120102222300321-3113101202102323-2322232121101001-2230333201203121)
- [custom_network_config.no_global_network](data-sources--securemesh_site--reference--group-003.md#canonical-3203231233310210-3312303100312211-0212030103112311-2211012220123012-0310031332022133-2222302312121122-2022301033230003-3123301021221321)
- [custom_network_config.no_network_policy](data-sources--securemesh_site--reference--group-003.md#canonical-0211120021231100-0200211311001101-0221332130221012-3033122221023220-0003033302132322-2132311332201133-2020001113203320-1003003232002322)
- [custom_network_config.sli_config](data-sources--securemesh_site--reference--group-003.md#canonical-3103313002121330-2310210032130122-0020030012032313-3231200213021302-1023120230300033-0330321311213001-0130211012103120-3321323300103033)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.sm_connection_public_ip](data-sources--securemesh_site--reference--group-004.md#canonical-3202030221310100-2332312320000312-2311123323010011-3123113311220311-0110233233302131-3330002323231000-1132213111220103-0102232030213320)
- [custom_network_config.sm_connection_pvt_ip](data-sources--securemesh_site--reference--group-004.md#canonical-2020320120002130-2233001122130003-0032113133002121-1310301132212321-2221222033132221-1323303312320211-0313232000130213-3301100230311111)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2121211033231322-2203220213000332-2202233212321213-0130121222023200-3311013231122020-0231000222212011-3020201030020002-0313223323032033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210312032331233-1231221030333001-1110002101331011-2323033111222231-2101011130003031-0301231230222332-2332032120001320-3122010013220021"></a>

## custom_network_config.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 222321003222 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.active_enhanced_firewall_policies

<a id="canonical-2231221221332213-2233233321332130-3332323020032013-0111333202023220-3321133123221023-1313233003133200-1222201330202202-2000032130212033"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

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

<a id="canonical-3122102320311311-2203002223012313-1030110330003121-3130120010123000-2222313321133321-2311200212333002-0211111302031121-0200300312213123"></a>

## Direct properties — active_enhanced_firewall_policies / 222321003222 / 3

- [enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-3103022213102111-2231030100021203-3032011000001113-3223221103232003-1023303021310111-2013133132323110-0130220122221211-3330020213001100): complete subsection reference.

<a id="canonical-0300110103222122-3120010103133021-1133123031133233-2133322031220003-0301211320321102-0033332113322223-1020332012311030-0210113120300131"></a>

## Next pages — active_enhanced_firewall_policies / 222321003222 / 4

- [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-3103022213102111-2231030100021203-3032011000001113-3223221103232003-1023303021310111-2013133132323110-0130220122221211-3330020213001100)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3103022213102111-2231030100021203-3032011000001113-3223221103232003-1023303021310111-2013133132323110-0130220122221211-3330020213001100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321331202103330-2310010130123100-2301010301030203-3103012110133203-2012311233312100-3132110303300332-0000310020320032-3023303012300010"></a>

## custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 200333100012 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2121211033231322-2203220213000332-2202233212321213-0130121222023200-3311013231122020-0231000222212011-3020201030020002-0313223323032033)
- custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-0200231000333120-0322311003031032-2122211012012303-2212220031333300-0333320221112200-1013211202203311-2000300330020300-1001100101311321"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3313131322210110-3111322222333321-0203031103333101-0332202000300203-2023201230101320-3223313002221102-0201102312230022-2130310000333222"></a>

## Direct properties — enhanced_firewall_policies / 200333100012 / 3

<a id="canonical-1310021232303200-2123033100133221-1212200213221012-0231213033113331-2332223221331120-2023113111102132-3101313233213023-1310211110220231"></a>

<a id="canonical-0322213101311230-3222102003231331-3311320132021030-3230301230302133-1310013012223300-2002103000102231-0203323001021032-2001230321310021"></a>

## name property — enhanced_firewall_policies / 200333100012 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2211210212120020-2103010030231213-3103321330323132-0332323023101220-2022230201312022-2131230002120100-3303231113132033-0223013120002230"></a>

<a id="canonical-2033311201302213-0123110313230221-0322120122020322-3313101121311113-2330303023220303-2311223010212022-1002001131120013-1233223320203333"></a>

## namespace property — enhanced_firewall_policies / 200333100012 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3100123332101210-0001002012231113-1002003220122230-2202020302033022-1213323332110200-1300213200301132-0021223333303201-2320233011301010"></a>

<a id="canonical-3312110133203333-1111103123303301-1210102021000223-1030020013120302-2222102323030132-2033231003232201-0021023233322210-0211013012000112"></a>

## tenant property — enhanced_firewall_policies / 200333100012 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3332322223122023-3130103212112303-1233133012220313-2022011211310200-1013113121113122-0132211123111233-0233111210221212-1332330032210323"></a>

## Next pages — enhanced_firewall_policies / 200333100012 / 7

- [custom_network_config.active_enhanced_firewall_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2121211033231322-2203220213000332-2202233212321213-0130121222023200-3311013231122020-0231000222212011-3020201030020002-0313223323032033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1133131032301110-1111213131002033-1212103011002311-0322211003112121-0023301012130222-3322123132312112-1030310022302011-2223001302001033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210301111023121-0332013212022333-2222232320221022-3201300220223301-2203332132322112-2101131300002310-0331303300310011-0032002301300332"></a>

## custom_network_config.active_forward_proxy_policies — active_forward_proxy_policies / 022213331232 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.active_forward_proxy_policies

<a id="canonical-1211000032103020-3302102123103323-2033010013032223-1312001102021012-3021102001112332-1301103020012132-1110110303320313-3012213303211130"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

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

<a id="canonical-3020001103101110-2122011202213101-1310213031031212-2002130011100003-1012031302321322-1223220103122330-2033111110220000-3221212002030012"></a>

## Direct properties — active_forward_proxy_policies / 022213331232 / 3

- [forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2322102011033330-3311122322232122-3110210323002020-2101102312132100-2200320210201012-2330002312320103-0322200022312333-2010233032111131): complete subsection reference.

<a id="canonical-3223111133001023-0011211030310011-0112112320101331-0013302023321311-0233321202232312-0231201313031211-1031113211121021-2213333122321210"></a>

## Next pages — active_forward_proxy_policies / 022213331232 / 4

- [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2322102011033330-3311122322232122-3110210323002020-2101102312132100-2200320210201012-2330002312320103-0322200022312333-2010233032111131)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2322102011033330-3311122322232122-3110210323002020-2101102312132100-2200320210201012-2330002312320103-0322200022312333-2010233032111131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302213221120213-0111002011000201-1323232010021212-1311113220130233-0331203001010210-3020223200021331-3222130003301111-0112320022330123"></a>

## custom_network_config.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 101101031323 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-1133131032301110-1111213131002033-1212103011002311-0322211003112121-0023301012130222-3322123132312112-1030310022302011-2223001302001033)
- custom_network_config.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-1302010221203110-3022121113333223-3031221222332333-3203201321102133-1303123333302123-0323002301330021-1223321011332330-2303230020202333"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0013311121101031-0313023103232033-0230201002301231-1310223001023212-0310320201010131-2203001130232103-0231303101333013-1233311112313132"></a>

## Direct properties — forward_proxy_policies / 101101031323 / 3

<a id="canonical-3011023332033013-2333120211031110-1110223020312211-2010131130333102-1112103013120331-1233030000312012-2100210120321003-1130303330133231"></a>

<a id="canonical-0112222130031323-1132210001103212-1113012103321021-1202120010001002-3000300111020211-0333010313123312-1332123313203023-1022102031112210"></a>

## name property — forward_proxy_policies / 101101031323 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3012001013211231-0102312201131200-2331120131233113-0233110120312021-3003020301120033-3133202100100302-1110123123131212-0331221311030310"></a>

<a id="canonical-1031201200100210-0030000223311012-0220101311011213-3001332102001230-2133000110302003-3203020301000322-2031321012100012-3102000033223123"></a>

## namespace property — forward_proxy_policies / 101101031323 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1302331021021013-2201211301032323-2201231122321130-0123211132233000-0221222020000232-0013203233132323-1030333100322231-2110102112221122"></a>

<a id="canonical-1101001021223011-2311021030333222-0113303021111301-0032110311110232-3231003202013301-0100203300130102-1122003330110310-0013202300133330"></a>

## tenant property — forward_proxy_policies / 101101031323 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1322132220000202-2002330030323032-2233003333200211-0113320023133003-0020003203133220-1021020020103113-0011223011330011-3221020111030303"></a>

## Next pages — forward_proxy_policies / 101101031323 / 7

- [custom_network_config.active_forward_proxy_policies](data-sources--securemesh_site--reference--group-001.md#canonical-1133131032301110-1111213131002033-1212103011002311-0322211003112121-0023301012130222-3322123132312112-1030310022302011-2223001302001033)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2233302301301023-1033312232031023-0130313233310323-1030010202322132-0130012202331122-2310303231101102-3023000233010220-0211230331203321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331110131333323-3103202302022033-3322123020221110-2112100313120101-1123232011131020-3203023102030203-2130220000023103-3033130213023121"></a>

## custom_network_config.active_network_policies — active_network_policies / 310031302300 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.active_network_policies

<a id="canonical-2310131011010330-0030132311203103-1230232212011103-3222121112332122-3200001233033120-0330030301321222-3010213021102302-1220112110013332"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

<a id="canonical-2110320202230203-1301131313133210-0203203102032320-3210210021122312-3200101110131210-0211203101331310-2332033032032311-2133330103101310"></a>

## Direct properties — active_network_policies / 310031302300 / 3

- [network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-3132321221110302-3103123011012013-1023332300330111-2233101131033002-2322222300000203-1110112213112200-2222213003032332-2021132333123313): complete subsection reference.

<a id="canonical-3320121100130101-1321111130113123-1313301122333200-1311102001301022-1310233213032201-0230022013031210-3111312023032113-3333112232023032"></a>

## Next pages — active_network_policies / 310031302300 / 4

- [custom_network_config.active_network_policies.network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-3132321221110302-3103123011012013-1023332300330111-2233101131033002-2322222300000203-1110112213112200-2222213003032332-2021132333123313)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3132321221110302-3103123011012013-1023332300330111-2233101131033002-2322222300000203-1110112213112200-2222213003032332-2021132333123313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310211332113311-3312023100232330-1122001221210133-0323020221311213-0011123031330320-1010021001230011-1102213013123021-3110020111323110"></a>

## custom_network_config.active_network_policies.network_policies — network_policies / 131330220311 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2233302301301023-1033312232031023-0130313233310323-1030010202322132-0130012202331122-2310303231101102-3023000233010220-0211230331203321)
- custom_network_config.active_network_policies.network_policies

<a id="canonical-3310120021212322-0312331312220121-0332222221203021-0222333310010232-1210313103230232-2220232313221112-0302212000031102-0113223101213011"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1300303022203322-0103112022120002-2113302100021201-0111100110302232-1132131002221300-2133130330022301-1023331212321012-2012310330232011"></a>

## Direct properties — network_policies / 131330220311 / 3

<a id="canonical-2131322000303022-0210030120030331-1131301213322123-0100103331010213-1000320102020021-2031032221000223-0101011020310212-1210303220333003"></a>

<a id="canonical-1023233333201232-1130000013123003-1330022132101031-3222010023000103-0022200030102122-0133230111320332-2201212221302010-0101322302232223"></a>

## name property — network_policies / 131330220311 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0023110213001123-1222333320200001-2003030112303022-0323311321030023-3000103123110123-2121211231012210-2120030320010110-2311213313033032"></a>

<a id="canonical-0112011302322001-2102032333233333-1133330103023232-2212131020000312-3133011222230220-0032021103223203-0231111000302120-0103211210020000"></a>

## namespace property — network_policies / 131330220311 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2302202110321300-0101202001132012-3132330200300022-0312031020102321-3211132211221212-3100303212113331-0022200302003230-3223112032121122"></a>

<a id="canonical-1222122122303020-0230301101232232-0233202131021333-2101121122200300-2300001011213200-0122330231110321-1213233102333201-1331232221032223"></a>

## tenant property — network_policies / 131330220311 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3222212201133120-1100101103110001-1201322201310233-2333023102110310-3202322231332200-3310021030113012-3111000011310223-2012222310300331"></a>

## Next pages — network_policies / 131330220311 / 7

- [custom_network_config.active_network_policies](data-sources--securemesh_site--reference--group-001.md#canonical-2233302301301023-1033312232031023-0130313233310323-1030010202322132-0130012202331122-2310303231101102-3023000233010220-0211230331203321)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0213010111103030-2232130120020312-2232212233012303-0130022232321200-2210311331132301-3230130121221200-3003012201333332-3010000023301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121311300123312-3301032013003212-3231233332032312-0213212102112003-1023220311132000-0302232023100322-0112302000322213-3122033010032201"></a>

## custom_network_config.default_config — default_config / 330322212002 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.default_config

<a id="canonical-0333311101111230-0021121321102033-2210121231023110-2201330330002101-0300221102323232-2020311100323110-2301011310223103-2002233331213300"></a>

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

<a id="canonical-3120100200323012-2020101121323300-0111313021230300-0210102033321331-2211300232332132-3313033303033333-0113101010121233-1013211013200230"></a>

## Direct properties — default_config / 330322212002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101123103310123-0210200111012311-0333023200100130-1122320110221102-0131330233002213-3123223331122032-2213122100331101-1312303233321130"></a>

## Next pages — default_config / 330322212002 / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0221331032223311-1132302203121312-0021102302131133-3322101113132303-1300102233223222-1000010333010222-3231122200220010-0021330001132102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
