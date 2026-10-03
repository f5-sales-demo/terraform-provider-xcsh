---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322012003023202-3223310300321223-3312101032331020-0330333222200130-3213323131001013-1312000122110320-2211121011110220-0222102311201223"></a>

## Property reference — Property reference / 330130300032 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- Property reference

<a id="canonical-0302231313310200-2022122003311100-3120131201301320-2011223103200030-1023322231002032-2330133231302323-2230123001000331-3131122301032231"></a>

## Direct properties — Property reference / 330130300032 / 3

<a id="canonical-2312111200030001-3233120030233311-1233032113220311-3213123110101331-0321111111003020-2332122223331201-0210313320110101-3022110311100321"></a>

<a id="canonical-2222301112323303-2112301302113302-2131222211022231-1202231022033011-3002111123331200-0221021220130011-0112123012030032-0012022132330323"></a>

## address property — Property reference / 330130300032 / 4

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

- [allow_all_usb](data-sources--voltstack_site--reference--group-003.md#canonical-0111302202332011-0112122333123302-1101032220130223-0232232033112203-0332320320312331-0100020031232322-0202331122230211-2001303230221001): complete subsection reference.

<a id="canonical-1003102001120310-2000210102213301-0111310011131310-0211021312022301-0233133321032323-0212102321201102-2202113031233313-1110123200013110"></a>

<a id="canonical-0333201312022301-2131001031230123-3322232120031210-2020321100331213-1333313130000100-3330111032011201-0102031301030110-3322112320332121"></a>

## annotations property — Property reference / 330130300032 / 5

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

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

- [blocked_services](data-sources--voltstack_site--reference--group-003.md#canonical-1320032310313331-3022123233030331-2113013102303221-1102302013031322-3112012101222212-2322201032032023-1000232212120103-2222220211221331): complete subsection reference.

- [bond_device_list](data-sources--voltstack_site--reference--group-003.md#canonical-2221000333033013-3101020231120210-1203330110312111-2232113313123020-1212333332233022-1123312130301022-3122012202003103-0010302130011021): complete subsection reference.

- [coordinates](data-sources--voltstack_site--reference--group-003.md#canonical-0120100301310332-3321102221032322-3313312210320231-1230000302000320-0220120120321122-0322232023312223-3122020223200100-1323202230122303): complete subsection reference.

- [custom_dns](data-sources--voltstack_site--reference--group-003.md#canonical-3112321101202333-3003313103032102-1030303202301211-3022013012030312-1320113311313012-0023210300332103-3303003201100112-1123020201113201): complete subsection reference.

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201): complete subsection reference.

- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020): complete subsection reference.

- [default_blocked_services](data-sources--voltstack_site--reference--group-008.md#canonical-3213101201031113-3231303201013301-0231002333210112-1212331331322010-0231102213222122-1331321110330131-0100113322130111-3112310031310310): complete subsection reference.

- [default_network_config](data-sources--voltstack_site--reference--group-008.md#canonical-2131230020322123-1303110032220233-0323021011312130-2022321132301211-2212210322021202-3222312320032321-0313130320333323-2113222132113330): complete subsection reference.

- [default_sriov_interface](data-sources--voltstack_site--reference--group-008.md#canonical-0030112211201131-0233032203330233-3012211220221210-1101021313232031-2202102032101020-1213210233030113-2330202220220200-1203113120131202): complete subsection reference.

- [default_storage_config](data-sources--voltstack_site--reference--group-008.md#canonical-0021033110303103-0003030222103213-0103111210211111-3220133201312030-3103021232313120-1031020022023221-3102223131032022-2133202231130231): complete subsection reference.

- [deny_all_usb](data-sources--voltstack_site--reference--group-008.md#canonical-0102023001301202-2201311011312032-0331020330222220-1022123133210333-1111000022131312-0020010200333130-0312111120133202-0331021201333110): complete subsection reference.

<a id="canonical-2210200303003020-3201111101330022-3300122123301332-2002121012230113-3130102100100333-2003033031212130-3121321300233123-1210123131331322"></a>

<a id="canonical-0313320031102212-3103000123312103-0023312311103031-1211130000223101-0111321132032131-1201133030200100-0322110133202122-1102120122112110"></a>

## description property — Property reference / 330130300032 / 6

Type: `"string"`. Computed.

Description of the VoltstackSite.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [disable_gpu](data-sources--voltstack_site--reference--group-008.md#canonical-2013133233301023-3002101131132333-0132130032032231-1201022321001011-2201310323103301-3212311322221321-3312032120220131-0232123003003001): complete subsection reference.

- [disable_vm](data-sources--voltstack_site--reference--group-008.md#canonical-1033000201210201-3132201233301200-1212311330301013-3030133000212021-3320323323311230-1211202021132213-3011000201300121-0321301200302120): complete subsection reference.

- [enable_gpu](data-sources--voltstack_site--reference--group-008.md#canonical-3301023003030020-3333310331323132-3301311002332122-0011132321032203-1003232133133122-1220020221013232-2022011123101120-1230033021010322): complete subsection reference.

- [enable_vgpu](data-sources--voltstack_site--reference--group-008.md#canonical-1131330000323001-2000332111122333-3330122310233213-2112201320320113-3011213113023220-1211302100220022-2031013022001131-1002210211033323): complete subsection reference.

- [enable_vm](data-sources--voltstack_site--reference--group-009.md#canonical-2120033220230201-3201323001000303-2300010003321030-1322013303121300-1311012110311000-2112222220331213-2322312333333022-2023330301130102): complete subsection reference.

<a id="canonical-0120001012211001-2132003111030330-3110213113121033-2103112023111013-2132013232300212-0122200121103031-0312221002313111-1123133333133021"></a>

<a id="canonical-1012100103202112-2022323122131300-2211331312220213-3203113111322322-2020133003100111-2332102213120323-1021200000102033-1303120000131001"></a>

## ID property — Property reference / 330130300032 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster](data-sources--voltstack_site--reference--group-009.md#canonical-1132230233133203-3131223100302000-2120313313123100-3220310332231001-0231212333120223-3101313210021001-1002131331120000-1311031032021230): complete subsection reference.

- [kubernetes_upgrade_drain](data-sources--voltstack_site--reference--group-009.md#canonical-3320131330201111-2130101200102000-2332212122311310-2211301113232033-3130213303132011-2010111113113320-3322320132110020-1012320321211130): complete subsection reference.

<a id="canonical-2210112203231323-0233121313333102-0103012221211323-2102232100220332-0233102332100201-2020213122032230-2323201121321300-3230333303000231"></a>

<a id="canonical-2010231212020322-2212033232020031-2113232330321312-3112330230310323-3332222031031330-2031020200322130-2002122202100222-3212033003332011"></a>

## labels property — Property reference / 330130300032 / 8

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

- [local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-0013103332313000-3213203202333213-2320232030232220-1310031130101133-1101202000301222-0131131223112332-1322201131002100-1033110110022000): complete subsection reference.

- [log_receiver](data-sources--voltstack_site--reference--group-009.md#canonical-1100301300011210-1122120223131022-0032330322112300-1102133312212303-2031101220113112-1233113010003231-0222131002303000-0111231120132313): complete subsection reference.

- [logs_streaming_disabled](data-sources--voltstack_site--reference--group-009.md#canonical-1100220121332220-0132031112133000-1301301112301011-2112212331211203-0122110000030122-2312103221213312-2022032120012312-1133010222332101): complete subsection reference.

- [master_node_configuration](data-sources--voltstack_site--reference--group-009.md#canonical-1003223011333012-3011003222112223-1001002232011322-1312002320233201-0300202321101310-1133031313320121-1320013310302222-1103131023001123): complete subsection reference.

<a id="canonical-1123120011000021-0231320221333121-1031133312300121-3031302212102203-3232220333310312-2002310001322312-3213323312131122-0012313003112320"></a>

<a id="canonical-1303312012210221-1303213133002303-1333233110100310-3212120122202011-3321332323223022-0210112032131010-0323330223323220-3122202101013322"></a>

## name property — Property reference / 330130300032 / 9

Type: `"string"`. Required.

Name of the VoltstackSite.

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

<a id="canonical-3110001303012032-0322132232112013-1232202132013330-1001031321112210-1033103122202333-3021330333331101-1223211003213310-2330121121322210"></a>

<a id="canonical-0231300111001211-2312020010002012-2011121233112112-0301221321232022-1103103102231210-2323303112311101-1112020022210013-2023231001123132"></a>

## namespace property — Property reference / 330130300032 / 10

Type: `"string"`. Required.

Namespace where the VoltstackSite exists.

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

- [no_bond_devices](data-sources--voltstack_site--reference--group-009.md#canonical-0120020123011223-1311000213300223-2220032023023232-0032303201302330-0123302133032113-3033112202120311-0320232003212212-3311033312103210): complete subsection reference.

- [no_k8s_cluster](data-sources--voltstack_site--reference--group-009.md#canonical-2111131212322100-1100210202033210-0300330213220310-1223033122312231-3201223203323210-0320122322220222-1211023320230222-1310012211130330): complete subsection reference.

- [no_local_control_plane](data-sources--voltstack_site--reference--group-009.md#canonical-1321303301011331-3111230333030121-1023310203120102-1032331222030111-1112320330012000-1133103002222310-2000121010031030-0122132302233222): complete subsection reference.

- [offline_survivability_mode](data-sources--voltstack_site--reference--group-009.md#canonical-3012201203013310-2202301221032113-3211212303232023-3001230312211120-0223122330211122-3321300212212230-1123030231222122-0310112232022022): complete subsection reference.

- [os](data-sources--voltstack_site--reference--group-010.md#canonical-3100102213131302-3311133131010110-2212201213202300-0103003312120103-3133113233030313-2213101010303002-2212003321123131-0123323330123023): complete subsection reference.

- [sriov_interfaces](data-sources--voltstack_site--reference--group-010.md#canonical-1103213303213211-0311100031033111-3302230130111311-2303232333310320-0133310202311332-2311222201011323-3120221112120033-0011312001223100): complete subsection reference.

- [sw](data-sources--voltstack_site--reference--group-010.md#canonical-3133230301232221-1223023111221301-1202112211133303-2202023113310010-3011203311032120-0303120202211331-2113302331010032-2112301303323312): complete subsection reference.

- [usb_policy](data-sources--voltstack_site--reference--group-010.md#canonical-1223032330333221-1120211003302212-3132123311030210-2021022313010323-1200101121303220-0221231301121030-0203332322210210-1120030001133101): complete subsection reference.

<a id="canonical-2113220120023223-1013222000301233-2320200300323223-0330202310113222-1111320320113030-2323300030220210-0013212213030033-0101023203013310"></a>

<a id="canonical-0032333132212311-2130302021003000-1000231323332010-3022002031202213-2310131300103233-3032033002103133-2331203220232103-1100122103101332"></a>

## volterra_certified_hw property — Property reference / 330130300032 / 11

Type: `"string"`. Computed.

Name for generic server certified hardware to form this App Stack site.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [waf_signatures](data-sources--voltstack_site--reference--group-010.md#canonical-1131011312303133-1023302002122003-1333133200020121-3021223030101331-1222202302102321-1033021231101223-0223110202000230-0231001012033213): complete subsection reference.

<a id="canonical-3223332113220012-3100321111133222-3122031101323223-2123133223311332-3033100003300030-1311031203313221-2320302110232022-0301233233030331"></a>

<a id="canonical-3102233323103230-2101222002000131-2032310013110300-0202311223032312-1022100111331212-0211011210010231-2322131221302300-0200121312132201"></a>

## worker_nodes property — Property reference / 330130300032 / 12

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
