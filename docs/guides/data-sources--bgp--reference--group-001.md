---
page_title: "xcsh_bgp reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp reference."
---

# xcsh_bgp reference

<a id="canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- Property reference

<a id="canonical-3001333331113003-2033223123213103-2330302023311220-0300332323203132-3120330110122223-1001100023222210-2231312210331002-3023123221300122"></a>

### Direct properties for `xcsh_bgp`

<a id="canonical-0333023303010102-1201222003000203-0203320033200110-0301313133222320-3232131130212123-2131010000121111-0211220301331013-3002322023100112"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011): complete subsection reference.

<a id="canonical-0103222222321021-2102011021322131-1102030001130233-3333103221130320-1111323220301322-0122132022322313-1023303221333212-3201221123222010"></a>

<a id="canonical-1313212022011033-0332312013323333-0001112133012011-0311323132113113-2330201100013121-1200311201130332-3022003321123310-2331022323123300"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the BGP.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0121002032313321-3020222310313233-2123030203222010-3212111103113023-3022030310313011-2302333123331023-2121111020100200-3223033330121112"></a>

<a id="canonical-3303020103121232-3121211130013320-2331231112102112-2011332111002110-0010312110213120-2102112310012202-2003131311311313-1322022321103222"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1130302202202012-0311131203310221-1032223001102133-3310011103321312-2003230032313331-3131203020100003-3100203032213131-2300303121011123"></a>

<a id="canonical-0110112222113010-2001110202012101-1010130131220033-3022031213020223-2011023020002331-1233333210232211-1031311322131330-0321131231333321"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-1103331230312120-2121322022231033-3022210013310031-2112003212332030-3023000023202022-2033322020131012-1000100232302322-3330100000300211"></a>

<a id="canonical-2000031230002112-0113222333021102-1010332122002110-1130303311213303-1013331113011323-3001232000110232-3130023313210131-2312122230012111"></a>

#### `name` property

Type: `"string"`. Required.

Name of the BGP.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1132201002020113-2301232223101221-2010101013333332-0210011101032312-3131112220033023-1330123011121302-0130310233323100-3232302213220321"></a>

<a id="canonical-2010020213101101-0213020332323113-0332111232310211-2011221231210302-3101001312200321-2102312113220113-1211213220020222-2000232102200001"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the BGP exists.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203): complete subsection reference.

- [where](data-sources--bgp--reference--group-002.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120): complete subsection reference.

<a id="canonical-0111200313020221-1132020233023310-3302002202031033-1003103001202120-2302211113211212-1232110032130021-1000323101230310-0023123301213231"></a>

### All schema paths for `xcsh_bgp`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--bgp--reference--group-001.md#canonical-0333023303010102-1201222003000203-0203320033200110-0301313133222320-3232131130212123-2131010000121111-0211220301331013-3002322023100112) |
| `bgp_parameters` | [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1331102210001001-3233333012232210-2133213322320300-0122301322301212-0331120111212111-2000211122130131-1200213000230022-2312103201120222) |
| `bgp_parameters.asn` | [bgp_parameters.asn](data-sources--bgp--reference--group-001.md#canonical-0303133002131301-0213301031332131-1031221331211030-1013010320220113-3031121222312130-2112333001311223-1130202022232220-2021010022302321) |
| `bgp_parameters.from_site` | [bgp_parameters.from_site](data-sources--bgp--reference--group-001.md#canonical-3220223122020320-3211101131121003-1211111213321021-0010232021301202-3201021111301202-0033233220322102-3032010010302110-2310103100312112) |
| `bgp_parameters.ip_address` | [bgp_parameters.ip_address](data-sources--bgp--reference--group-001.md#canonical-1231103313020232-3122103313123022-1002021313302233-3310103201003223-0303332133131211-1011220102023103-1120223223122100-0322012231222110) |
| `bgp_parameters.local_address` | [bgp_parameters.local_address](data-sources--bgp--reference--group-001.md#canonical-3003311203202311-2133333121121312-3201223331032303-2012001200101020-0203112121330130-2300122132232312-0023232321321311-0320311302221223) |
| `description` | [description](data-sources--bgp--reference--group-001.md#canonical-0103222222321021-2102011021322131-1102030001130233-3333103221130320-1111323220301322-0122132022322313-1023303221333212-3201221123222010) |
| `id` | [ID](data-sources--bgp--reference--group-001.md#canonical-0121002032313321-3020222310313233-2123030203222010-3212111103113023-3022030310313011-2302333123331023-2121111020100200-3223033330121112) |
| `labels` | [labels](data-sources--bgp--reference--group-001.md#canonical-1130302202202012-0311131203310221-1032223001102133-3310011103321312-2003230032313331-3131203020100003-3100203032213131-2300303121011123) |
| `name` | [name](data-sources--bgp--reference--group-001.md#canonical-1103331230312120-2121322022231033-3022210013310031-2112003212332030-3023000023202022-2033322020131012-1000100232302322-3330100000300211) |
| `namespace` | [namespace](data-sources--bgp--reference--group-001.md#canonical-1132201002020113-2301232223101221-2010101013333332-0210011101032312-3131112220033023-1330123011121302-0130310233323100-3232302213220321) |
| `peers` | [peers](data-sources--bgp--reference--group-001.md#canonical-0103320030001203-2210113302313322-2102201203303201-3220122113223211-1320132330302102-2212221010233321-3030230223112103-1020203200020231) |
| `peers.bfd_disabled` | [peers.bfd_disabled](data-sources--bgp--reference--group-001.md#canonical-0300221110021101-1012301311120321-2203013110210213-3023010220313321-3021311222020031-1110300012020213-1131323131213002-1103331213012123) |
| `peers.bfd_enabled` | [peers.bfd_enabled](data-sources--bgp--reference--group-001.md#canonical-3000003210321332-2111213301021333-0202021220032121-3302212331100010-0030202111103110-3320330310320002-3223203213212331-2113311321001102) |
| `peers.bfd_enabled.multiplier` | [peers.bfd_enabled.multiplier](data-sources--bgp--reference--group-001.md#canonical-0320301111200211-3002313200232102-3122001200012030-0012121011101123-2232100323031211-3222302312103200-2212030310120300-3032223111213120) |
| `peers.bfd_enabled.receive_interval_milliseconds` | [peers.bfd_enabled.receive_interval_milliseconds](data-sources--bgp--reference--group-001.md#canonical-3333301022110111-1232010123231231-3213012111010120-3223003201130122-2103302100130022-0210100220012010-0231130333300113-1302320011110103) |
| `peers.bfd_enabled.transmit_interval_milliseconds` | [peers.bfd_enabled.transmit_interval_milliseconds](data-sources--bgp--reference--group-001.md#canonical-3112112313103220-0100033332121321-1011001222213021-3320012020330121-2203222323022020-2001113332002212-3221222201130330-2103220113122200) |
| `peers.disable_spec` | [peers.disable_spec](data-sources--bgp--reference--group-001.md#canonical-0203021110002310-2201313303232131-1222220022021201-0231020100122210-1000233300020301-0201021231321211-3122011201122330-2222030322230120) |
| `peers.ebgp_multihop_disabled` | [peers.ebgp_multihop_disabled](data-sources--bgp--reference--group-001.md#canonical-2013322032310313-2200222232011221-1201021232113031-2213003001023302-3330013301212322-1003011103110133-0301010030201033-2200211220232021) |
| `peers.ebgp_multihop_enabled` | [peers.ebgp_multihop_enabled](data-sources--bgp--reference--group-001.md#canonical-2331033000012203-2012132301332310-2212222230221131-0102311301012321-0320221213212033-3202303311133120-3331310330001123-3111013102101322) |
| `peers.external` | [peers.external](data-sources--bgp--reference--group-001.md#canonical-3320231030122312-3230203312212100-3333203132111222-3311213103001021-2022203123012203-1011331220003231-2010031220022321-1320320010223003) |
| `peers.external.address` | [peers.external.address](data-sources--bgp--reference--group-001.md#canonical-1121220230230023-3213011210120231-0102010300130112-0303131301332020-1230230112102101-1132110000321003-3101132301012302-1101112213021321) |
| `peers.external.address_ipv6` | [peers.external.address_ipv6](data-sources--bgp--reference--group-001.md#canonical-1102220231332201-3320203031033103-1031201310101230-0131110031121301-2000111332122002-0032232010003210-1321303223011302-2312232320011032) |
| `peers.external.asn` | [peers.external.asn](data-sources--bgp--reference--group-001.md#canonical-1021222233001232-3220033123003030-1122221202010133-1323031323033030-1202122301332312-0222201201310012-1311330012120330-2020123033130331) |
| `peers.external.default_gateway` | [peers.external.default_gateway](data-sources--bgp--reference--group-001.md#canonical-0021121311303000-3010033321120203-0010010301131203-0112201030220313-1110220021022303-0203022300303211-1111000130201231-2131033222220202) |
| `peers.external.default_gateway_v6` | [peers.external.default_gateway_v6](data-sources--bgp--reference--group-001.md#canonical-0230120223123220-0100311333320220-1133321230112230-1230320010103221-2033233332121321-1031101113331011-2223230020321000-3111000010010000) |
| `peers.external.disable_spec` | [peers.external.disable_spec](data-sources--bgp--reference--group-001.md#canonical-3133330231313202-3100112210013201-2101333233330211-0100032220012232-2000013313332032-0200320232300022-1101232103032210-2301031203101211) |
| `peers.external.disable_v6` | [peers.external.disable_v6](data-sources--bgp--reference--group-001.md#canonical-0321223012111320-0031213300123313-1101323110233032-3303111020121023-1032023301312131-2210120330030103-3213302120013000-2222013222331312) |
| `peers.external.external_connector` | [peers.external.external_connector](data-sources--bgp--reference--group-001.md#canonical-2100122021212133-2310011032311121-0102303333202023-1312110013301302-3222003313303312-2131033011021210-2223002023203330-2230332021100022) |
| `peers.external.family_inet` | [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-2000203102221113-2101022100011023-0203212133231212-1111221012312330-1212302100213231-1203111013200120-0101011132101011-2320021130201323) |
| `peers.external.family_inet.disable_spec` | [peers.external.family_inet.disable_spec](data-sources--bgp--reference--group-001.md#canonical-3222132212233022-0200232331200221-3213011303021233-0012001001213330-3321320201333131-2203233123002110-2321233232102330-1332332210130311) |
| `peers.external.family_inet.enable` | [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-1011233203311103-2020131001310103-0210130323030331-3220330200002123-1221213323200323-0200231301132023-2212222020321312-3301030233013100) |
| `peers.external.family_inet.enable.aggregation` | [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-0032203012112013-1031220302033333-1310222002203230-2131030130311120-3220302022023331-3100210203001123-0122311330120002-3121111331320120) |
| `peers.external.family_inet.enable.aggregation.ip_prefix` | [peers.external.family_inet.enable.aggregation.ip_prefix](data-sources--bgp--reference--group-001.md#canonical-2113101301030013-2200231010322033-1303310112230033-1312220301132320-2200103110201332-0002311002113333-0033032103011101-3111102300220111) |
| `peers.external.family_inet.enable.aggregation.options` | [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--reference--group-001.md#canonical-0203121121000323-0102031121310021-0322112211221023-2103333131211300-0330022033123303-1310232222312000-1021211220012023-1031012002101112) |
| `peers.external.family_inet.enable.aggregation.options.summary_only` | [peers.external.family_inet.enable.aggregation.options.summary_only](data-sources--bgp--reference--group-001.md#canonical-0330031011011330-1102302011021100-3133222210133302-1021330100133230-3302221220220100-3320000201331033-2230001112111331-3213233102101033) |
| `peers.external.from_site` | [peers.external.from_site](data-sources--bgp--reference--group-001.md#canonical-2002331301301011-0003200103023021-0231312233311312-0203010210323212-1123231320002032-2331001223332102-0301020233331230-3301302011031132) |
| `peers.external.from_site_v6` | [peers.external.from_site_v6](data-sources--bgp--reference--group-001.md#canonical-2102020101131000-0201302113220000-3211131132002002-3032302112200321-0310302230330331-1200331033232313-3220332123202331-2022030202320211) |
| `peers.external.interface` | [peers.external.interface](data-sources--bgp--reference--group-001.md#canonical-0311231232312012-1133023122013130-2323110300021310-3331130021212330-0301332310033231-1033030333021030-3202131222100233-3311100302310033) |
| `peers.external.interface.name` | [peers.external.interface.name](data-sources--bgp--reference--group-001.md#canonical-2331322113222030-3123110310312212-1203320112130302-0203013011221133-1003100302022311-2003112003131000-0312113110323013-3120010003011310) |
| `peers.external.interface.namespace` | [peers.external.interface.namespace](data-sources--bgp--reference--group-001.md#canonical-3213202110102213-0221220010330000-0001331233033220-3310201301031103-0000321213232023-2231333033011212-3210301121200112-0133120200100201) |
| `peers.external.interface.tenant` | [peers.external.interface.tenant](data-sources--bgp--reference--group-002.md#canonical-2010302211213103-3013002221020001-3332112312311130-0130302200213211-3313300010320012-3321312021013003-3020320302330112-3122120320101211) |
| `peers.external.interface_list` | [peers.external.interface_list](data-sources--bgp--reference--group-002.md#canonical-1131003102033211-0333221032220223-0023200032231100-2013300112301130-0031303020113131-2010322132133300-1012200303023331-3131102012332300) |
| `peers.external.interface_list.interfaces` | [peers.external.interface_list.interfaces](data-sources--bgp--reference--group-002.md#canonical-0110222120101112-0312332110020303-1233001231211312-2011000133033220-1200203201231211-0133010013110113-1111210002300013-3133233300223320) |
| `peers.external.interface_list.interfaces.name` | [peers.external.interface_list.interfaces.name](data-sources--bgp--reference--group-002.md#canonical-1010031202002133-3313203100130222-0330201013102133-3031332223233202-3132211020110213-1210332120130013-0232322312021001-2323021112313303) |
| `peers.external.interface_list.interfaces.namespace` | [peers.external.interface_list.interfaces.namespace](data-sources--bgp--reference--group-002.md#canonical-2323113102032032-2221020100133310-0133010002111130-1222300222302212-3201123212100301-2332331203000113-3013322232132321-0020130332113212) |
| `peers.external.interface_list.interfaces.tenant` | [peers.external.interface_list.interfaces.tenant](data-sources--bgp--reference--group-002.md#canonical-2212312010113133-3333030010201100-2313310233320222-2103000213010102-1011203222231212-0333022010102102-3221223100331033-3011010022232222) |
| `peers.external.md5_auth_key` | [peers.external.md5_auth_key](data-sources--bgp--reference--group-001.md#canonical-1022323110021200-0330120000303001-1012032232001011-3022133113213030-0102321230012231-1202223211331020-0012002320210121-3030100212232303) |
| `peers.external.no_authentication` | [peers.external.no_authentication](data-sources--bgp--reference--group-002.md#canonical-3010203300113023-1231010032331030-3213303023310212-2312220222312332-2012023013111300-1011231132301000-0211223031231302-1123031030321122) |
| `peers.external.port` | [peers.external.port](data-sources--bgp--reference--group-001.md#canonical-1313033220213200-1021023010131200-1102300113011113-1333331330233300-0233201331311033-0322120202013200-0232221300330301-1132031301323333) |
| `peers.external.subnet_begin_offset` | [peers.external.subnet_begin_offset](data-sources--bgp--reference--group-001.md#canonical-2331120100013221-0023022333021102-1310023113232103-2132110031112022-3022002213130330-0313223000200212-2130212200313012-1331331013100201) |
| `peers.external.subnet_begin_offset_v6` | [peers.external.subnet_begin_offset_v6](data-sources--bgp--reference--group-001.md#canonical-3121032111233233-0321022012233100-1010201311203210-0123312021112011-0331123222120202-2313010120030331-2330000100011321-0201101332030103) |
| `peers.external.subnet_end_offset` | [peers.external.subnet_end_offset](data-sources--bgp--reference--group-001.md#canonical-3132312130000000-0103020001303010-2312202020000122-1210131322223213-0113130220011111-0303230310023111-0210333000001000-3123103323102131) |
| `peers.external.subnet_end_offset_v6` | [peers.external.subnet_end_offset_v6](data-sources--bgp--reference--group-001.md#canonical-2102303120202222-1101122200300331-2321131322103300-0311300212313110-3303110233212130-3230132132232221-1001320232031231-1330213021001311) |
| `peers.label` | [peers.label](data-sources--bgp--reference--group-001.md#canonical-0123121322103020-0223020301300330-3211123132213213-0022323023121220-0233230123220203-3001022322331020-1210110333123222-1111311331003032) |
| `peers.metadata` | [peers.metadata](data-sources--bgp--reference--group-002.md#canonical-0323110232012220-2323021101313302-1131100021310013-2300111000230232-0220333233310311-2210133023310010-2322001020321033-2000302001210002) |
| `peers.metadata.description_spec` | [peers.metadata.description_spec](data-sources--bgp--reference--group-002.md#canonical-2223311023010213-3121010210313020-2030220320232011-1200132130213032-0322313212333123-0320122132122203-1322103120101210-1031031231120232) |
| `peers.metadata.name` | [peers.metadata.name](data-sources--bgp--reference--group-002.md#canonical-2302013311232220-1132332100102310-3323223320113000-2000310023001102-0300130112031333-2021013310233013-0130311133130201-1303102322103210) |
| `peers.passive_mode_disabled` | [peers.passive_mode_disabled](data-sources--bgp--reference--group-002.md#canonical-2003303132030221-3021102113031322-3201232230221211-0232032133210320-2002211032030020-3020202121302320-0332212123132202-0113001321121222) |
| `peers.passive_mode_enabled` | [peers.passive_mode_enabled](data-sources--bgp--reference--group-002.md#canonical-1120310213202313-2103130203003131-1220330133013100-0203002310202010-0302102131313121-2111323012030323-3212030222032012-1210112221231110) |
| `peers.routing_policies` | [peers.routing_policies](data-sources--bgp--reference--group-002.md#canonical-0303322023121133-1100203002202001-0220122012203033-3113131201321322-0203132102321222-3232121332030323-3031221322011110-2303100013303102) |
| `peers.routing_policies.route_policy` | [peers.routing_policies.route_policy](data-sources--bgp--reference--group-002.md#canonical-1212222011200010-0211021023032113-0310303323230131-2022133121130021-2110310322233300-1233132303322002-3201202302003110-0331133131312320) |
| `peers.routing_policies.route_policy.all_nodes` | [peers.routing_policies.route_policy.all_nodes](data-sources--bgp--reference--group-002.md#canonical-1311112220200302-0101132330010231-3302020331130011-3000211222010120-3130222112113010-0312323131130201-3033310131333121-0020332300013131) |
| `peers.routing_policies.route_policy.inbound` | [peers.routing_policies.route_policy.inbound](data-sources--bgp--reference--group-002.md#canonical-1323200103330211-1213010211022210-3110123201112132-2133111300222332-3113002310011321-2121013101110202-1021221233122312-0331223331023231) |
| `peers.routing_policies.route_policy.node_name` | [peers.routing_policies.route_policy.node_name](data-sources--bgp--reference--group-002.md#canonical-2201013133111121-1020201010313021-3033132103120103-1021220020230210-2232303321111030-1323301310011212-0123231313213322-2001323303210311) |
| `peers.routing_policies.route_policy.node_name.node` | [peers.routing_policies.route_policy.node_name.node](data-sources--bgp--reference--group-002.md#canonical-2103021021302002-2111000313103212-2222233313312333-2020222330220001-2213021332112033-0022321033022200-1202013130103131-2030333321302010) |
| `peers.routing_policies.route_policy.object_refs` | [peers.routing_policies.route_policy.object_refs](data-sources--bgp--reference--group-002.md#canonical-2111012301313303-0112231230100100-0223033001211232-0022123311202023-0332330322221123-2033131122113120-0000010113121133-3222123313020332) |
| `peers.routing_policies.route_policy.object_refs.kind` | [peers.routing_policies.route_policy.object_refs.kind](data-sources--bgp--reference--group-002.md#canonical-0202332003120331-1113023102311312-0103212303121131-1012310313320202-1011132023012220-2120003210122002-1322102100121311-3022011022103120) |
| `peers.routing_policies.route_policy.object_refs.name` | [peers.routing_policies.route_policy.object_refs.name](data-sources--bgp--reference--group-002.md#canonical-3301230133111313-3201102121033202-3201233321320301-1311310120112023-3131023033013331-2330021110003022-2101230112133220-0101331233032110) |
| `peers.routing_policies.route_policy.object_refs.namespace` | [peers.routing_policies.route_policy.object_refs.namespace](data-sources--bgp--reference--group-002.md#canonical-0212022032123033-1233310233300301-0120112213211032-3120221310033120-1310122012202333-2001312001101202-2332121020002011-3032222023000031) |
| `peers.routing_policies.route_policy.object_refs.tenant` | [peers.routing_policies.route_policy.object_refs.tenant](data-sources--bgp--reference--group-002.md#canonical-0020033130233233-0221110123131301-1211322223030033-1131031323131331-3202131313122310-1010031010110300-3100111301111113-1101122131302110) |
| `peers.routing_policies.route_policy.object_refs.uid` | [peers.routing_policies.route_policy.object_refs.uid](data-sources--bgp--reference--group-002.md#canonical-1203011331312213-2200211033202231-0310103203320222-1000133001210233-3220122332331200-0122123310211232-3020132321330032-1230331030312102) |
| `peers.routing_policies.route_policy.outbound` | [peers.routing_policies.route_policy.outbound](data-sources--bgp--reference--group-002.md#canonical-2230131211311312-3310032303030110-2133222212333031-3322220311230132-0233201213210322-1213313033331032-2313100201003102-3121113333001002) |
| `where` | [where](data-sources--bgp--reference--group-002.md#canonical-1032200123222212-3232232020300023-0212233033002010-1202033301101323-3132330001120313-2323310330203233-3001220010112023-3031110023033022) |
| `where.site` | [where.site](data-sources--bgp--reference--group-002.md#canonical-0130100031103010-0103100200123001-1333221132210021-1013113302101321-2111022303100303-1131223001303132-3300133301121000-2001012100223110) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--bgp--reference--group-002.md#canonical-2111130212202122-2010122132311213-1030211313221313-1020023231022331-0323221100113300-0100100113132122-0330332130122332-2131031100333222) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--bgp--reference--group-002.md#canonical-1223123112103332-2113010300103112-0321103331333231-1000022023300032-0032332112003312-0201011322212202-1203223301232121-1223103133003101) |
| `where.site.network_type` | [where.site.network_type](data-sources--bgp--reference--group-002.md#canonical-0013201030023013-3013120320223303-1303231311013213-0212223132021321-0130222020213212-0323010112233013-0102233123012011-0010313023313302) |
| `where.site.ref` | [where.site.ref](data-sources--bgp--reference--group-002.md#canonical-3213230032233033-1331131121320133-0210121110113120-1132322223032022-2222310300010111-2132033131221311-1233023021213132-3010032222031300) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--bgp--reference--group-002.md#canonical-3310321001030322-0113113110022211-0233211231101120-2013221111322203-1103303200201230-3001230212302103-3221203101020231-1122000332003311) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--bgp--reference--group-002.md#canonical-1211120301233121-0322122020030002-3210301223221000-3210002033200200-0333220121303112-3023120310033320-0111210222003220-2301113123000133) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--bgp--reference--group-002.md#canonical-1110310020320133-1023120201223012-0210002202303031-3030032022203021-2111301001031122-1331232133133333-2233313332112313-3200010302320312) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--bgp--reference--group-002.md#canonical-3103132103131100-0323131023000023-1120220120133313-0121032322102310-1103011313330321-2020112132200023-3112133313300132-2111013303320020) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--bgp--reference--group-002.md#canonical-3133002332002232-0322211230012010-3133023133200323-1103203232222332-3223211110123231-2100223303020201-3212010222011011-2332013232232323) |
| `where.virtual_site` | [where.virtual_site](data-sources--bgp--reference--group-002.md#canonical-3320101110123322-1011312002232332-3112322123131231-0301223333010302-0302133212210222-3033203023211103-1030203211121231-0313310101103321) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--bgp--reference--group-002.md#canonical-3122313321211100-0331010002300012-1322002101211302-1000330021230021-1220322221321123-2312203011311023-3020010330202031-1233330030201001) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--bgp--reference--group-002.md#canonical-2211312233303220-2131111113212001-2312100331023211-0321120203022312-1231231132302000-3312013200222010-3320302311101100-0120020123021212) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--bgp--reference--group-002.md#canonical-2102131220223300-2101111232022231-1000223032013211-1003113123322001-1131111030113000-1321333310331003-3012221122103331-2103201030023011) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--bgp--reference--group-002.md#canonical-0110323222102130-2100212211222200-3013310010121220-2001201231032003-2123310330202102-2300100220333032-2012312321132010-3023331221311012) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--bgp--reference--group-002.md#canonical-3111030300032222-2332210030021322-1021310012201320-0231232003020011-3012031123130321-0321111103213302-3000120130311313-3223000022132111) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--bgp--reference--group-002.md#canonical-1313001302322031-3322100032323301-0031030220220230-0202333233021210-1323112312303223-0231013132022022-2302100113202312-2312000320210212) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--bgp--reference--group-002.md#canonical-1320121200322031-0003231231001031-3231032202301230-3331201322111231-3222121133202033-0003213120101032-2012322001232001-3201332002212012) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--bgp--reference--group-002.md#canonical-0033302331230021-1010030233100303-3032323132303033-3320233303000323-0201010303213210-3331330230211322-2312332013030101-0012103333110101) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--bgp--reference--group-002.md#canonical-0021313313213030-2101211032000022-1301103021312311-0011201112303311-0003301213021213-0311102333131121-0013301021313322-3210231111111013) |

<a id="canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bgp_parameters` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- bgp_parameters

<a id="canonical-1331102210001001-3233333012232210-2133213322320300-0122301322301212-0331120111212111-2000211122130131-1200213000230022-2312103201120222"></a>

Type: `"single"`. Computed.

Configuration parameter for bgp parameters.

Additional upstream details:

BGP parameters for the local site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-router_id_choice": "[\"from_site\",\"ip_address\",\"local_address\"]"
}
```

<a id="canonical-2323033033300321-2120331033211132-3201210223113320-2131322201231213-2223210311211021-1031003200221203-3011233230212011-0332200302310202"></a>

### Direct properties for `bgp_parameters`

<a id="canonical-0303133002131301-0213301031332131-1031221331211030-1013010320220113-3031121222312130-2112333001311223-1130202022232220-2021010022302321"></a>

#### `bgp_parameters.asn` property

Type: `"number"`. Computed.

ASN. Autonomous System Number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [from_site](data-sources--bgp--reference--group-001.md#canonical-0022310012030221-1213011121130033-1003200300022300-3132203211031211-0232111302120130-1121332200113312-3010031203211310-2211102103113101): complete subsection reference.

<a id="canonical-1231103313020232-3122103313123022-1002021313302233-3310103201003223-0303332133131211-1011220102023103-1120223223122100-0322012231222110"></a>

<a id="canonical-0231003330321202-1313123320230113-0111011301001011-1313210313111112-2003020311033203-1303120202231131-1223312003222130-0022131331331022"></a>

#### `bgp_parameters.ip_address` property

Type: `"string"`. Computed.

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [local_address](data-sources--bgp--reference--group-001.md#canonical-3111011321003313-1121200210332311-2321030331002223-2020202110220122-1132011030011110-0011030002211231-1101012321300323-0121122203002102): complete subsection reference.

<a id="canonical-0022310012030221-1213011121130033-1003200300022300-3132203211031211-0232111302120130-1121332200113312-3010031203211310-2211102103113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bgp_parameters.from_site` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011)
- bgp_parameters.from_site

<a id="canonical-3220223122020320-3211101131121003-1211111213321021-0010232021301202-3201021111301202-0033233220322102-3032010010302110-2310103100312112"></a>

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

<a id="canonical-3111011321003313-1121200210332311-2321030331002223-2020202110220122-1132011030011110-0011030002211231-1101012321300323-0121122203002102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bgp_parameters.local_address` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011)
- bgp_parameters.local_address

<a id="canonical-3003311203202311-2133333121121312-3201223331032303-2012001200101020-0203112121330130-2300122132232312-0023232321321311-0320311302221223"></a>

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

<a id="canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- peers

<a id="canonical-0103320030001203-2210113302313322-2102201203303201-3220122113223211-1320132330302102-2212221010233321-3030230223112103-1020203200020231"></a>

Type: `"list"`. Computed.

Peers. List of peers.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-3201310103133001-1220133113131211-3121233022232312-3002022103012033-1213110003321220-2000310233232011-3102130320202222-3223321232002130"></a>

### Direct properties for `peers`

- [bfd_disabled](data-sources--bgp--reference--group-001.md#canonical-0333102131200232-0212101230300102-0120213222211232-1221030212211333-2011321012232211-0020030102113110-2031300103321002-3123101301312330): complete subsection reference.

- [bfd_enabled](data-sources--bgp--reference--group-001.md#canonical-0101313101233210-3201102200021201-0210020233311333-3010133210220310-3211312111020311-1323220303202223-0113132322000221-3301133030301130): complete subsection reference.

- [disable_spec](data-sources--bgp--reference--group-001.md#canonical-0222100230333202-3031200021030321-0023031002311302-1112213023120022-2331203113112212-0020100302211212-2323133023103320-0230102202000003): complete subsection reference.

- [ebgp_multihop_disabled](data-sources--bgp--reference--group-001.md#canonical-0033331103223012-3322132002320203-2131031123113112-2031130022302232-2311223023232020-3301220102313030-0221003000013223-0110212200100020): complete subsection reference.

- [ebgp_multihop_enabled](data-sources--bgp--reference--group-001.md#canonical-2011212023110130-0102112211203100-1011230333102231-2303022201230010-3100320310221030-0003220201231030-2302023023312232-1133320100323013): complete subsection reference.

- [external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302): complete subsection reference.

<a id="canonical-0123121322103020-0223020301300330-3211123132213213-0022323023121220-0233230123220203-3001022322331020-1210110333123222-1111311331003032"></a>

<a id="canonical-3121100300031000-2323123201301303-3202302230102210-1033313013121020-2112330122120203-2220010121132130-0212200131132030-0122313121203002"></a>

#### `peers.label` property

Type: `"string"`. Computed.

Label. Specify whether this peer should be.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "labeling",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--bgp--reference--group-002.md#canonical-1210203212232330-2312201101103002-3021033131220003-3100323101002011-2333003323332203-2332113013012033-3302202221133313-1231022220113331): complete subsection reference.

- [passive_mode_disabled](data-sources--bgp--reference--group-002.md#canonical-3330222003020013-3002312001203013-3102312030323313-0223103232321233-1203310122113212-3222022000002200-3300330030300020-0011203212322220): complete subsection reference.

- [passive_mode_enabled](data-sources--bgp--reference--group-002.md#canonical-3111202111113022-2232110120120100-0211312103003231-2022332313131333-2020120311312200-0310213133111102-0030021111133110-1110203031113002): complete subsection reference.

- [routing_policies](data-sources--bgp--reference--group-002.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321): complete subsection reference.

<a id="canonical-0333102131200232-0212101230300102-0120213222211232-1221030212211333-2011321012232211-0020030102113110-2031300103321002-3123101301312330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.bfd_disabled` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.bfd_disabled

<a id="canonical-0300221110021101-1012301311120321-2203013110210213-3023010220313321-3021311222020031-1110300012020213-1131323131213002-1103331213012123"></a>

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

<a id="canonical-0101313101233210-3201102200021201-0210020233311333-3010133210220310-3211312111020311-1323220303202223-0113132322000221-3301133030301130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.bfd_enabled` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.bfd_enabled

<a id="canonical-3000003210321332-2111213301021333-0202021220032121-3302212331100010-0030202111103110-3320330310320002-3223203213212331-2113311321001102"></a>

Type: `"single"`. Computed.

BFD. BFD parameters.

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

<a id="canonical-2032100311032321-2033003020230201-3123300012022100-3201333232313022-3201111110212333-0133030000203001-0211110333010000-0311202221013330"></a>

### Direct properties for `peers.bfd_enabled`

<a id="canonical-0320301111200211-3002313200232102-3122001200012030-0012121011101123-2232100323031211-3222302312103200-2212030310120300-3032223111213120"></a>

#### `peers.bfd_enabled.multiplier` property

Type: `"number"`. Computed.

Specify Number of missed packets to bring session down'.

Additional upstream details:

Specify Number of missed packets to bring session down"

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-3333301022110111-1232010123231231-3213012111010120-3223003201130122-2103302100130022-0210100220012010-0231130333300113-1302320011110103"></a>

<a id="canonical-1122130012013122-0310323031131203-1202330302320210-2121022013012030-0132311323013231-1122002031300331-2032320010130333-2023011013121022"></a>

#### `peers.bfd_enabled.receive_interval_milliseconds` property

Type: `"number"`. Computed.

BFD receive interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-3112112313103220-0100033332121321-1011001222213021-3320012020330121-2203222323022020-2001113332002212-3221222201130330-2103220113122200"></a>

<a id="canonical-2300301113302221-0331330300301112-2122213333120221-1030123332332032-2120210202013030-3300210110130322-2112203121332231-0331000113202102"></a>

#### `peers.bfd_enabled.transmit_interval_milliseconds` property

Type: `"number"`. Computed.

BFD transmit interval timer, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-0222100230333202-3031200021030321-0023031002311302-1112213023120022-2331203113112212-0020100302211212-2323133023103320-0230102202000003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.disable_spec` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.disable_spec

<a id="canonical-0203021110002310-2201313303232131-1222220022021201-0231020100122210-1000233300020301-0201021231321211-3122011201122330-2222030322230120"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033331103223012-3322132002320203-2131031123113112-2031130022302232-2311223023232020-3301220102313030-0221003000013223-0110212200100020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.ebgp_multihop_disabled` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.ebgp_multihop_disabled

<a id="canonical-2013322032310313-2200222232011221-1201021232113031-2213003001023302-3330013301212322-1003011103110133-0301010030201033-2200211220232021"></a>

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

<a id="canonical-2011212023110130-0102112211203100-1011230333102231-2303022201230010-3100320310221030-0003220201231030-2302023023312232-1133320100323013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.ebgp_multihop_enabled` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.ebgp_multihop_enabled

<a id="canonical-2331033000012203-2012132301332310-2212222230221131-0102311301012321-0320221213212033-3202303311133120-3331310330001123-3111013102101322"></a>

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

<a id="canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.external

<a id="canonical-3320231030122312-3230203312212100-3333203132111222-3311213103001021-2022203123012203-1011331220003231-2010031220022321-1320320010223003"></a>

Type: `"single"`. Computed.

External BGP Peer. External BGP Peer parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"address\",\"default_gateway\",\"disable\",\"external_connector\",\"from_site\",\"subnet_begin_offset\",\"subnet_end_offset\"]",
  "x-ves-oneof-field-address_choice_v6": "[\"address_ipv6\",\"default_gateway_v6\",\"disable_v6\",\"from_site_v6\",\"subnet_begin_offset_v6\",\"subnet_end_offset_v6\"]",
  "x-ves-oneof-field-auth_choice": "[\"md5_auth_key\",\"no_authentication\"]",
  "x-ves-oneof-field-interface_choice": "[\"interface\",\"interface_list\"]"
}
```

<a id="canonical-0303103023013131-2332032123220223-1310221222232021-3022110023233102-3133001033123030-0010233220201122-2333003111202213-0222133000231120"></a>

### Direct properties for `peers.external`

<a id="canonical-1121220230230023-3213011210120231-0102010300130112-0303131301332020-1230230112102101-1132110000321003-3101132301012302-1101112213021321"></a>

#### `peers.external.address` property

Type: `"string"`. Computed.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1102220231332201-3320203031033103-1031201310101230-0131110031121301-2000111332122002-0032232010003210-1321303223011302-2312232320011032"></a>

<a id="canonical-1200201321310132-3211301012001203-2321032111223003-2002321113301111-3230121022300213-2121102110030222-0022212302321320-3120212221132330"></a>

#### `peers.external.address_ipv6` property

Type: `"string"`. Computed.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1021222233001232-3220033123003030-1122221202010133-1323031323033030-1202122301332312-0222201201310012-1311330012120330-2020123033130331"></a>

<a id="canonical-1033330301000101-0033003021031200-3323121021330133-1203310313231013-3332030313330320-1203100323133311-2031202113003113-2310302022130302"></a>

#### `peers.external.asn` property

Type: `"number"`. Computed.

ASN. Autonomous System Number for BGP peer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [default_gateway](data-sources--bgp--reference--group-001.md#canonical-3302111002200333-2232131201012220-3012131031221203-0231033330123033-0312303122301332-2030003200003033-3103100223231301-2121203020011020): complete subsection reference.

- [default_gateway_v6](data-sources--bgp--reference--group-001.md#canonical-2002320231121031-3100100212113030-0230121333202013-1323223233323022-0131021013202001-0010023003120312-2003102220233200-2321100213130002): complete subsection reference.

- [disable_spec](data-sources--bgp--reference--group-001.md#canonical-2031302223232031-3332022333300230-3232333112123211-2201223032303231-3322002212022102-0200022231100301-1301203212110233-2301232231103002): complete subsection reference.

- [disable_v6](data-sources--bgp--reference--group-001.md#canonical-0021110022103030-1111002111033331-1201033111231011-0322330321011013-3332210111100103-0332031311332332-0030111113113210-1201020022013330): complete subsection reference.

- [external_connector](data-sources--bgp--reference--group-001.md#canonical-2322223132011310-1220013012220232-3121311110023132-2333123201301333-0122201202231200-3303102021221130-3200300331032213-3301121313301300): complete subsection reference.

- [family_inet](data-sources--bgp--reference--group-001.md#canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030): complete subsection reference.

- [from_site](data-sources--bgp--reference--group-001.md#canonical-2000201023230011-1131222320032101-0211021122331001-3211313212301012-3230211211100130-2220222213320110-3311003212020111-2221133201020221): complete subsection reference.

- [from_site_v6](data-sources--bgp--reference--group-001.md#canonical-2330211123232010-0122030211321121-1030113211210123-0233323001233222-0302103102003332-3223201002013231-0103012003131213-3010020320120111): complete subsection reference.

- [interface](data-sources--bgp--reference--group-001.md#canonical-1202300022003132-3311301221023101-2320321113301022-1223131323301322-2221113331121331-3110002032220221-1230201213100311-1330112310021023): complete subsection reference.

- [interface_list](data-sources--bgp--reference--group-002.md#canonical-0332112000321201-2012130032100313-3003023100112021-1221001212101031-3222213330302233-1032331013323132-1312222321313302-1202133211002223): complete subsection reference.

<a id="canonical-1022323110021200-0330120000303001-1012032232001011-3022133113213030-0102321230012231-1202223211331020-0012002320210121-3030100212232303"></a>

<a id="canonical-2121202233132010-1312203031102102-1111011332131230-0120110121100210-0100310311222100-0311031132113202-1223101211032031-2030203103020002"></a>

#### `peers.external.md5_auth_key` property

Type: `"string"`. Computed.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [no_authentication](data-sources--bgp--reference--group-002.md#canonical-1120030012320211-2222321211313311-2102010233321200-2013312131010001-1321330100100030-3300110231320011-3010023322032023-0111033230310112): complete subsection reference.

<a id="canonical-1313033220213200-1021023010131200-1102300113011113-1333331330233300-0233201331311033-0322120202013200-0232221300330301-1132031301323333"></a>

<a id="canonical-1011132323302201-2111032022002332-1323203102013133-1132211133223231-1123130312030232-0002121303101213-1220212222231023-0101223112122032"></a>

#### `peers.external.port` property

Type: `"number"`. Computed.

Peer Port. Peer TCP port number.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2331120100013221-0023022333021102-1310023113232103-2132110031112022-3022002213130330-0313223000200212-2130212200313012-1331331013100201"></a>

<a id="canonical-1220211102303321-2133032303321013-2133123210223001-1132103331001232-3213220021033033-0102130023123022-0130333201133113-0123312110110331"></a>

#### `peers.external.subnet_begin_offset` property

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3121032111233233-0321022012233100-1010201311203210-0123312021112011-0331123222120202-2313010120030331-2330000100011321-0201101332030103"></a>

<a id="canonical-3202302013330300-3202030001001133-1233130020012112-1101021012223230-3101211210132023-2013030320031220-1210233232022130-2003233032332222"></a>

#### `peers.external.subnet_begin_offset_v6` property

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3132312130000000-0103020001303010-2312202020000122-1210131322223213-0113130220011111-0303230310023111-0210333000001000-3123103323102131"></a>

<a id="canonical-0002111023113033-1200022210100123-3201113333323303-0101332020202131-3001103123011330-0020122200013312-1332332231222133-1002310221131231"></a>

#### `peers.external.subnet_end_offset` property

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2102303120202222-1101122200300331-2321131322103300-0311300212313110-3303110233212130-3230132132232221-1001320232031231-1330213021001311"></a>

<a id="canonical-3200231103310132-3321113110322233-3021321221101303-1100320303200102-0212002332023132-2122121100213123-3000223210201120-2210311211003232"></a>

#### `peers.external.subnet_end_offset_v6` property

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3302111002200333-2232131201012220-3012131031221203-0231033330123033-0312303122301332-2030003200003033-3103100223231301-2121203020011020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.default_gateway` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.default_gateway

<a id="canonical-0021121311303000-3010033321120203-0010010301131203-0112201030220313-1110220021022303-0203022300303211-1111000130201231-2131033222220202"></a>

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

<a id="canonical-2002320231121031-3100100212113030-0230121333202013-1323223233323022-0131021013202001-0010023003120312-2003102220233200-2321100213130002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.default_gateway_v6` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.default_gateway_v6

<a id="canonical-0230120223123220-0100311333320220-1133321230112230-1230320010103221-2033233332121321-1031101113331011-2223230020321000-3111000010010000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway v6.

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

<a id="canonical-2031302223232031-3332022333300230-3232333112123211-2201223032303231-3322002212022102-0200022231100301-1301203212110233-2301232231103002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.disable_spec` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.disable_spec

<a id="canonical-3133330231313202-3100112210013201-2101333233330211-0100032220012232-2000013313332032-0200320232300022-1101232103032210-2301031203101211"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021110022103030-1111002111033331-1201033111231011-0322330321011013-3332210111100103-0332031311332332-0030111113113210-1201020022013330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.disable_v6` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.disable_v6

<a id="canonical-0321223012111320-0031213300123313-1101323110233032-3303111020121023-1032023301312131-2210120330030103-3213302120013000-2222013222331312"></a>

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

<a id="canonical-2322223132011310-1220013012220232-3121311110023132-2333123201301333-0122201202231200-3303102021221130-3200300331032213-3301121313301300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.external_connector` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.external_connector

<a id="canonical-2100122021212133-2310011032311121-0102303333202023-1312110013301302-3222003313303312-2131033011021210-2223002023203330-2230332021100022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for external connector.

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

<a id="canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.family_inet

<a id="canonical-2000203102221113-2101022100011023-0203212133231212-1111221012312330-1212302100213231-1203111013200120-0101011132101011-2320021130201323"></a>

Type: `"single"`. Computed.

Configuration parameter for family inet.

Additional upstream details:

Parameters for inet family.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-3033302112212320-2213130211011032-3113331213213320-3311022330121313-3333302030010121-3200003200212220-3233222211111311-1210031021120231"></a>

### Direct properties for `peers.external.family_inet`

- [disable_spec](data-sources--bgp--reference--group-001.md#canonical-1103102203233201-1000030232002020-3020300023300300-1201301123103121-0320312313212021-1210101023120212-1100103123121332-1312202122321302): complete subsection reference.

- [enable](data-sources--bgp--reference--group-001.md#canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132): complete subsection reference.

<a id="canonical-1103102203233201-1000030232002020-3020300023300300-1201301123103121-0320312313212021-1210101023120212-1100103123121332-1312202122321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.disable_spec` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030)
- peers.external.family_inet.disable_spec

<a id="canonical-3222132212233022-0200232331200221-3213011303021233-0012001001213330-3321320201333131-2203233123002110-2321233232102330-1332332210130311"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.enable` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030)
- peers.external.family_inet.enable

<a id="canonical-1011233203311103-2020131001310103-0210130323030331-3220330200002123-1221213323200323-0200231301132023-2212222020321312-3301030233013100"></a>

Type: `"single"`. Computed.

Unicast IPv4. IPv4 Unicast.

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

<a id="canonical-2132132201030302-1121121202211223-2130213131011302-0031201101131311-0233002311030133-3131210212321210-1012321223233113-3302013000030231"></a>

### Direct properties for `peers.external.family_inet.enable`

- [aggregation](data-sources--bgp--reference--group-001.md#canonical-0302321100202233-0231213322311003-3030221203212302-3100323132212200-2120102003121311-2133210022223132-2310332021103113-2202223020221302): complete subsection reference.

<a id="canonical-0302321100202233-0231213322311003-3030221203212302-3100323132212200-2120102003121311-2133210022223132-2310332021103113-2202223020221302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.enable.aggregation` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132)
- peers.external.family_inet.enable.aggregation

<a id="canonical-0032203012112013-1031220302033333-1310222002203230-2131030130311120-3220302022023331-3100210203001123-0122311330120002-3121111331320120"></a>

Type: `"list"`. Computed.

BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take
effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing
table and applies to outbound advertisements.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1301103110313120-0212120303233120-0303230201320223-1303103002133321-1332132330101113-2300110330223132-2302321130222232-2023320121102212"></a>

### Direct properties for `peers.external.family_inet.enable.aggregation`

<a id="canonical-2113101301030013-2200231010322033-1303310112230033-1312220301132320-2200103110201332-0002311002113333-0033032103011101-3111102300220111"></a>

#### `peers.external.family_inet.enable.aggregation.ip_prefix` property

Type: `"string"`. Computed.

IP Prefix. Specify IPv4 subnet for aggregation.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [options](data-sources--bgp--reference--group-001.md#canonical-1013013330322010-1201100003203212-2020223121313320-1233121112122021-3123030203103021-0211123113202323-3320211201101133-0133311123111331): complete subsection reference.

<a id="canonical-1013013330322010-1201100003203212-2020223121313320-1233121112122021-3123030203103021-0211123113202323-3320211201101133-0133311123111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.enable.aggregation.options` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132)
- [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-0302321100202233-0231213322311003-3030221203212302-3100323132212200-2120102003121311-2133210022223132-2310332021103113-2202223020221302)
- peers.external.family_inet.enable.aggregation.options

<a id="canonical-0203121121000323-0102031121310021-0322112211221023-2103333131211300-0330022033123303-1310232222312000-1021211220012023-1031012002101112"></a>

Type: `"list"`. Computed.

Aggregation OPTIONS. Configuration parameter for options

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3323203201011220-2231030130113133-1010113201023211-3030133101332013-2200303132333220-1322130013023131-3300312220330030-3203323030002110"></a>

### Direct properties for `peers.external.family_inet.enable.aggregation.options`

- [summary_only](data-sources--bgp--reference--group-001.md#canonical-2112132320203310-2103233321210210-1203102031323320-0200030213203021-1131321210103322-1333003103000332-3003033001020110-1333220002231232): complete subsection reference.

<a id="canonical-2112132320203310-2103233321210210-1203102031323320-0200030213203021-1131321210103322-1333003103000332-3003033001020110-1333220002231232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.enable.aggregation.options.summary_only` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132)
- [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-0302321100202233-0231213322311003-3030221203212302-3100323132212200-2120102003121311-2133210022223132-2310332021103113-2202223020221302)
- [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--reference--group-001.md#canonical-1013013330322010-1201100003203212-2020223121313320-1233121112122021-3123030203103021-0211123113202323-3320211201101133-0133311123111331)
- peers.external.family_inet.enable.aggregation.options.summary_only

<a id="canonical-0330031011011330-1102302011021100-3133222210133302-1021330100133230-3302221220220100-3320000201331033-2230001112111331-3213233102101033"></a>

Type: `"single"`. Computed.

Configuration parameter for summary only.

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

<a id="canonical-2000201023230011-1131222320032101-0211021122331001-3211313212301012-3230211211100130-2220222213320110-3311003212020111-2221133201020221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.from_site` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.from_site

<a id="canonical-2002331301301011-0003200103023021-0231312233311312-0203010210323212-1123231320002032-2331001223332102-0301020233331230-3301302011031132"></a>

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

<a id="canonical-2330211123232010-0122030211321121-1030113211210123-0233323001233222-0302103102003332-3223201002013231-0103012003131213-3010020320120111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.from_site_v6` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.from_site_v6

<a id="canonical-2102020101131000-0201302113220000-3211131132002002-3032302112200321-0310302230330331-1200331033232313-3220332123202331-2022030202320211"></a>

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

<a id="canonical-1202300022003132-3311301221023101-2320321113301022-1223131323301322-2221113331121331-3110002032220221-1230201213100311-1330112310021023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.interface` properties

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.interface

<a id="canonical-0311231232312012-1133023122013130-2323110300021310-3331130021212330-0301332310033231-1033030333021030-3202131222100233-3311100302310033"></a>

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

<a id="canonical-2331031121332031-3120131322220231-3031002220012312-0013113133320330-0330122002213003-3213221112312103-0102113100310121-1023332320223031"></a>

### Direct properties for `peers.external.interface`

<a id="canonical-2331322113222030-3123110310312212-1203320112130302-0203013011221133-1003100302022311-2003112003131000-0312113110323013-3120010003011310"></a>

#### `peers.external.interface.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3213202110102213-0221220010330000-0001331233033220-3310201301031103-0000321213232023-2231333033011212-3210301121200112-0133120200100201"></a>
