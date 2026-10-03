---
page_title: "xcsh_bgp reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp reference."
---

# xcsh_bgp reference

<a id="canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001333331113003-2033223123213103-2330302023311220-0300332323203132-3120330110122223-1001100023222210-2231312210331002-3023123221300122"></a>

## Property reference — Property reference / 222202130012 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- Property reference

<a id="canonical-1313212022011033-0332312013323333-0001112133012011-0311323132113113-2330201100013121-1200311201130332-3022003321123310-2331022323123300"></a>

## Direct properties — Property reference / 222202130012 / 3

<a id="canonical-0333023303010102-1201222003000203-0203320033200110-0301313133222320-3232131130212123-2131010000121111-0211220301331013-3002322023100112"></a>

<a id="canonical-3303020103121232-3121211130013320-2331231112102112-2011332111002110-0010312110213120-2102112310012202-2003131311311313-1322022321103222"></a>

## annotations property — Property reference / 222202130012 / 4

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

- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011): complete subsection reference.

<a id="canonical-0103222222321021-2102011021322131-1102030001130233-3333103221130320-1111323220301322-0122132022322313-1023303221333212-3201221123222010"></a>

<a id="canonical-0110112222113010-2001110202012101-1010130131220033-3022031213020223-2011023020002331-1233333210232211-1031311322131330-0321131231333321"></a>

## description property — Property reference / 222202130012 / 5

Type: `"string"`. Computed.

Description of the BGP.

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

<a id="canonical-0121002032313321-3020222310313233-2123030203222010-3212111103113023-3022030310313011-2302333123331023-2121111020100200-3223033330121112"></a>

<a id="canonical-2000031230002112-0113222333021102-1010332122002110-1130303311213303-1013331113011323-3001232000110232-3130023313210131-2312122230012111"></a>

## ID property — Property reference / 222202130012 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1130302202202012-0311131203310221-1032223001102133-3310011103321312-2003230032313331-3131203020100003-3100203032213131-2300303121011123"></a>

<a id="canonical-2010020213101101-0213020332323113-0332111232310211-2011221231210302-3101001312200321-2102312113220113-1211213220020222-2000232102200001"></a>

## labels property — Property reference / 222202130012 / 7

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

<a id="canonical-1103331230312120-2121322022231033-3022210013310031-2112003212332030-3023000023202022-2033322020131012-1000100232302322-3330100000300211"></a>

<a id="canonical-0111200313020221-1132020233023310-3302002202031033-1003103001202120-2302211113211212-1232110032130021-1000323101230310-0023123301213231"></a>

## name property — Property reference / 222202130012 / 8

Type: `"string"`. Required.

Name of the BGP.

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

<a id="canonical-1132201002020113-2301232223101221-2010101013333332-0210011101032312-3131112220033023-1330123011121302-0130310233323100-3232302213220321"></a>

<a id="canonical-3112222330103112-1232102000010111-3221200231330120-2111102012133131-3301322302001111-1020122020011311-1212110130213130-0003323003321000"></a>

## namespace property — Property reference / 222202130012 / 9

Type: `"string"`. Required.

Namespace where the BGP exists.

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

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203): complete subsection reference.

- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120): complete subsection reference.

<a id="canonical-1202220203113032-3023222133120132-1111213302312112-3232100230022023-2022323332001022-3333220022322222-2302333032213002-2121211330101312"></a>

## All schema paths — Property reference / 222202130012 / 10

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
| `peers.external.interface.tenant` | [peers.external.interface.tenant](data-sources--bgp--reference--group-001.md#canonical-2010302211213103-3013002221020001-3332112312311130-0130302200213211-3313300010320012-3321312021013003-3020320302330112-3122120320101211) |
| `peers.external.interface_list` | [peers.external.interface_list](data-sources--bgp--reference--group-001.md#canonical-1131003102033211-0333221032220223-0023200032231100-2013300112301130-0031303020113131-2010322132133300-1012200303023331-3131102012332300) |
| `peers.external.interface_list.interfaces` | [peers.external.interface_list.interfaces](data-sources--bgp--reference--group-001.md#canonical-0110222120101112-0312332110020303-1233001231211312-2011000133033220-1200203201231211-0133010013110113-1111210002300013-3133233300223320) |
| `peers.external.interface_list.interfaces.name` | [peers.external.interface_list.interfaces.name](data-sources--bgp--reference--group-001.md#canonical-1010031202002133-3313203100130222-0330201013102133-3031332223233202-3132211020110213-1210332120130013-0232322312021001-2323021112313303) |
| `peers.external.interface_list.interfaces.namespace` | [peers.external.interface_list.interfaces.namespace](data-sources--bgp--reference--group-001.md#canonical-2323113102032032-2221020100133310-0133010002111130-1222300222302212-3201123212100301-2332331203000113-3013322232132321-0020130332113212) |
| `peers.external.interface_list.interfaces.tenant` | [peers.external.interface_list.interfaces.tenant](data-sources--bgp--reference--group-001.md#canonical-2212312010113133-3333030010201100-2313310233320222-2103000213010102-1011203222231212-0333022010102102-3221223100331033-3011010022232222) |
| `peers.external.md5_auth_key` | [peers.external.md5_auth_key](data-sources--bgp--reference--group-001.md#canonical-1022323110021200-0330120000303001-1012032232001011-3022133113213030-0102321230012231-1202223211331020-0012002320210121-3030100212232303) |
| `peers.external.no_authentication` | [peers.external.no_authentication](data-sources--bgp--reference--group-001.md#canonical-3010203300113023-1231010032331030-3213303023310212-2312220222312332-2012023013111300-1011231132301000-0211223031231302-1123031030321122) |
| `peers.external.port` | [peers.external.port](data-sources--bgp--reference--group-001.md#canonical-1313033220213200-1021023010131200-1102300113011113-1333331330233300-0233201331311033-0322120202013200-0232221300330301-1132031301323333) |
| `peers.external.subnet_begin_offset` | [peers.external.subnet_begin_offset](data-sources--bgp--reference--group-001.md#canonical-2331120100013221-0023022333021102-1310023113232103-2132110031112022-3022002213130330-0313223000200212-2130212200313012-1331331013100201) |
| `peers.external.subnet_begin_offset_v6` | [peers.external.subnet_begin_offset_v6](data-sources--bgp--reference--group-001.md#canonical-3121032111233233-0321022012233100-1010201311203210-0123312021112011-0331123222120202-2313010120030331-2330000100011321-0201101332030103) |
| `peers.external.subnet_end_offset` | [peers.external.subnet_end_offset](data-sources--bgp--reference--group-001.md#canonical-3132312130000000-0103020001303010-2312202020000122-1210131322223213-0113130220011111-0303230310023111-0210333000001000-3123103323102131) |
| `peers.external.subnet_end_offset_v6` | [peers.external.subnet_end_offset_v6](data-sources--bgp--reference--group-001.md#canonical-2102303120202222-1101122200300331-2321131322103300-0311300212313110-3303110233212130-3230132132232221-1001320232031231-1330213021001311) |
| `peers.label` | [peers.label](data-sources--bgp--reference--group-001.md#canonical-0123121322103020-0223020301300330-3211123132213213-0022323023121220-0233230123220203-3001022322331020-1210110333123222-1111311331003032) |
| `peers.metadata` | [peers.metadata](data-sources--bgp--reference--group-001.md#canonical-0323110232012220-2323021101313302-1131100021310013-2300111000230232-0220333233310311-2210133023310010-2322001020321033-2000302001210002) |
| `peers.metadata.description_spec` | [peers.metadata.description_spec](data-sources--bgp--reference--group-001.md#canonical-2223311023010213-3121010210313020-2030220320232011-1200132130213032-0322313212333123-0320122132122203-1322103120101210-1031031231120232) |
| `peers.metadata.name` | [peers.metadata.name](data-sources--bgp--reference--group-001.md#canonical-2302013311232220-1132332100102310-3323223320113000-2000310023001102-0300130112031333-2021013310233013-0130311133130201-1303102322103210) |
| `peers.passive_mode_disabled` | [peers.passive_mode_disabled](data-sources--bgp--reference--group-001.md#canonical-2003303132030221-3021102113031322-3201232230221211-0232032133210320-2002211032030020-3020202121302320-0332212123132202-0113001321121222) |
| `peers.passive_mode_enabled` | [peers.passive_mode_enabled](data-sources--bgp--reference--group-001.md#canonical-1120310213202313-2103130203003131-1220330133013100-0203002310202010-0302102131313121-2111323012030323-3212030222032012-1210112221231110) |
| `peers.routing_policies` | [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-0303322023121133-1100203002202001-0220122012203033-3113131201321322-0203132102321222-3232121332030323-3031221322011110-2303100013303102) |
| `peers.routing_policies.route_policy` | [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-1212222011200010-0211021023032113-0310303323230131-2022133121130021-2110310322233300-1233132303322002-3201202302003110-0331133131312320) |
| `peers.routing_policies.route_policy.all_nodes` | [peers.routing_policies.route_policy.all_nodes](data-sources--bgp--reference--group-001.md#canonical-1311112220200302-0101132330010231-3302020331130011-3000211222010120-3130222112113010-0312323131130201-3033310131333121-0020332300013131) |
| `peers.routing_policies.route_policy.inbound` | [peers.routing_policies.route_policy.inbound](data-sources--bgp--reference--group-001.md#canonical-1323200103330211-1213010211022210-3110123201112132-2133111300222332-3113002310011321-2121013101110202-1021221233122312-0331223331023231) |
| `peers.routing_policies.route_policy.node_name` | [peers.routing_policies.route_policy.node_name](data-sources--bgp--reference--group-001.md#canonical-2201013133111121-1020201010313021-3033132103120103-1021220020230210-2232303321111030-1323301310011212-0123231313213322-2001323303210311) |
| `peers.routing_policies.route_policy.node_name.node` | [peers.routing_policies.route_policy.node_name.node](data-sources--bgp--reference--group-001.md#canonical-2103021021302002-2111000313103212-2222233313312333-2020222330220001-2213021332112033-0022321033022200-1202013130103131-2030333321302010) |
| `peers.routing_policies.route_policy.object_refs` | [peers.routing_policies.route_policy.object_refs](data-sources--bgp--reference--group-001.md#canonical-2111012301313303-0112231230100100-0223033001211232-0022123311202023-0332330322221123-2033131122113120-0000010113121133-3222123313020332) |
| `peers.routing_policies.route_policy.object_refs.kind` | [peers.routing_policies.route_policy.object_refs.kind](data-sources--bgp--reference--group-001.md#canonical-0202332003120331-1113023102311312-0103212303121131-1012310313320202-1011132023012220-2120003210122002-1322102100121311-3022011022103120) |
| `peers.routing_policies.route_policy.object_refs.name` | [peers.routing_policies.route_policy.object_refs.name](data-sources--bgp--reference--group-001.md#canonical-3301230133111313-3201102121033202-3201233321320301-1311310120112023-3131023033013331-2330021110003022-2101230112133220-0101331233032110) |
| `peers.routing_policies.route_policy.object_refs.namespace` | [peers.routing_policies.route_policy.object_refs.namespace](data-sources--bgp--reference--group-001.md#canonical-0212022032123033-1233310233300301-0120112213211032-3120221310033120-1310122012202333-2001312001101202-2332121020002011-3032222023000031) |
| `peers.routing_policies.route_policy.object_refs.tenant` | [peers.routing_policies.route_policy.object_refs.tenant](data-sources--bgp--reference--group-001.md#canonical-0020033130233233-0221110123131301-1211322223030033-1131031323131331-3202131313122310-1010031010110300-3100111301111113-1101122131302110) |
| `peers.routing_policies.route_policy.object_refs.uid` | [peers.routing_policies.route_policy.object_refs.uid](data-sources--bgp--reference--group-001.md#canonical-1203011331312213-2200211033202231-0310103203320222-1000133001210233-3220122332331200-0122123310211232-3020132321330032-1230331030312102) |
| `peers.routing_policies.route_policy.outbound` | [peers.routing_policies.route_policy.outbound](data-sources--bgp--reference--group-001.md#canonical-2230131211311312-3310032303030110-2133222212333031-3322220311230132-0233201213210322-1213313033331032-2313100201003102-3121113333001002) |
| `where` | [where](data-sources--bgp--reference--group-001.md#canonical-1032200123222212-3232232020300023-0212233033002010-1202033301101323-3132330001120313-2323310330203233-3001220010112023-3031110023033022) |
| `where.site` | [where.site](data-sources--bgp--reference--group-001.md#canonical-0130100031103010-0103100200123001-1333221132210021-1013113302101321-2111022303100303-1131223001303132-3300133301121000-2001012100223110) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-2111130212202122-2010122132311213-1030211313221313-1020023231022331-0323221100113300-0100100113132122-0330332130122332-2131031100333222) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-1223123112103332-2113010300103112-0321103331333231-1000022023300032-0032332112003312-0201011322212202-1203223301232121-1223103133003101) |
| `where.site.network_type` | [where.site.network_type](data-sources--bgp--reference--group-001.md#canonical-0013201030023013-3013120320223303-1303231311013213-0212223132021321-0130222020213212-0323010112233013-0102233123012011-0010313023313302) |
| `where.site.ref` | [where.site.ref](data-sources--bgp--reference--group-001.md#canonical-3213230032233033-1331131121320133-0210121110113120-1132322223032022-2222310300010111-2132033131221311-1233023021213132-3010032222031300) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--bgp--reference--group-001.md#canonical-3310321001030322-0113113110022211-0233211231101120-2013221111322203-1103303200201230-3001230212302103-3221203101020231-1122000332003311) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--bgp--reference--group-001.md#canonical-1211120301233121-0322122020030002-3210301223221000-3210002033200200-0333220121303112-3023120310033320-0111210222003220-2301113123000133) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--bgp--reference--group-001.md#canonical-1110310020320133-1023120201223012-0210002202303031-3030032022203021-2111301001031122-1331232133133333-2233313332112313-3200010302320312) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--bgp--reference--group-001.md#canonical-3103132103131100-0323131023000023-1120220120133313-0121032322102310-1103011313330321-2020112132200023-3112133313300132-2111013303320020) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--bgp--reference--group-001.md#canonical-3133002332002232-0322211230012010-3133023133200323-1103203232222332-3223211110123231-2100223303020201-3212010222011011-2332013232232323) |
| `where.virtual_site` | [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-3320101110123322-1011312002232332-3112322123131231-0301223333010302-0302133212210222-3033203023211103-1030203211121231-0313310101103321) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-3122313321211100-0331010002300012-1322002101211302-1000330021230021-1220322221321123-2312203011311023-3020010330202031-1233330030201001) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-2211312233303220-2131111113212001-2312100331023211-0321120203022312-1231231132302000-3312013200222010-3320302311101100-0120020123021212) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--bgp--reference--group-001.md#canonical-2102131220223300-2101111232022231-1000223032013211-1003113123322001-1131111030113000-1321333310331003-3012221122103331-2103201030023011) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--bgp--reference--group-001.md#canonical-0110323222102130-2100212211222200-3013310010121220-2001201231032003-2123310330202102-2300100220333032-2012312321132010-3023331221311012) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--bgp--reference--group-001.md#canonical-3111030300032222-2332210030021322-1021310012201320-0231232003020011-3012031123130321-0321111103213302-3000120130311313-3223000022132111) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--bgp--reference--group-001.md#canonical-1313001302322031-3322100032323301-0031030220220230-0202333233021210-1323112312303223-0231013132022022-2302100113202312-2312000320210212) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--bgp--reference--group-001.md#canonical-1320121200322031-0003231231001031-3231032202301230-3331201322111231-3222121133202033-0003213120101032-2012322001232001-3201332002212012) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--bgp--reference--group-001.md#canonical-0033302331230021-1010030233100303-3032323132303033-3320233303000323-0201010303213210-3331330230211322-2312332013030101-0012103333110101) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--bgp--reference--group-001.md#canonical-0021313313213030-2101211032000022-1301103021312311-0011201112303311-0003301213021213-0311102333131121-0013301021313322-3210231111111013) |

<a id="canonical-0031323003220111-0012210332103323-1313230020323013-0102130102001210-0010111000103311-2000212112022120-0220330112021302-2103210110022031"></a>

## Next pages — Property reference / 222202130012 / 11

- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323033033300321-2120331033211132-3201210223113320-2131322201231213-2223210311211021-1031003200221203-3011233230212011-0332200302310202"></a>

## bgp_parameters — bgp_parameters / 011113302323 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- bgp_parameters

<a id="canonical-1331102210001001-3233333012232210-2133213322320300-0122301322301212-0331120111212111-2000211122130131-1200213000230022-2312103201120222"></a>

Type: `"single"`. Computed.

Configuration parameter for bgp parameters.

Upstream description:

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

<a id="canonical-0231003330321202-1313123320230113-0111011301001011-1313210313111112-2003020311033203-1303120202231131-1223312003222130-0022131331331022"></a>

## Direct properties — bgp_parameters / 011113302323 / 3

<a id="canonical-0303133002131301-0213301031332131-1031221331211030-1013010320220113-3031121222312130-2112333001311223-1130202022232220-2021010022302321"></a>

<a id="canonical-3312211321122220-2000010213032123-0112001111300030-3101122200333231-3112312103333320-3330133212030220-0202220103200211-3130201220210313"></a>

## asn property — bgp_parameters / 011113302323 / 4

Type: `"number"`. Computed.

ASN. Autonomous System Number.

Upstream description:

Autonomous System Number.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0230121230201123-2123130210010001-3013133020020032-3030122322333032-1211310223111110-1322012133313033-1333222203010222-1130130101220030"></a>

## ip_address property — bgp_parameters / 011113302323 / 5

Type: `"string"`. Computed.

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Upstream description:

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

- [local_address](data-sources--bgp--reference--group-001.md#canonical-3111011321003313-1121200210332311-2321030331002223-2020202110220122-1132011030011110-0011030002211231-1101012321300323-0121122203002102): complete subsection reference.

<a id="canonical-2001022113012330-3103112332202212-0220132022011321-1012333301203330-1323102012122112-1232103102223221-2123022132023111-2232033311032313"></a>

## Next pages — bgp_parameters / 011113302323 / 6

- [bgp_parameters.from_site](data-sources--bgp--reference--group-001.md#canonical-0022310012030221-1213011121130033-1003200300022300-3132203211031211-0232111302120130-1121332200113312-3010031203211310-2211102103113101)
- [bgp_parameters.local_address](data-sources--bgp--reference--group-001.md#canonical-3111011321003313-1121200210332311-2321030331002223-2020202110220122-1132011030011110-0011030002211231-1101012321300323-0121122203002102)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0022310012030221-1213011121130033-1003200300022300-3132203211031211-0232111302120130-1121332200113312-3010031203211310-2211102103113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202111300131210-3203302232131111-3101032321203131-3002010112222010-2213130221323112-0132113230111223-0223111012031230-1010000022022311"></a>

## bgp_parameters.from_site — from_site / 321220133001 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011)
- bgp_parameters.from_site

<a id="canonical-3220223122020320-3211101131121003-1211111213321021-0010232021301202-3201021111301202-0033233220322102-3032010010302110-2310103100312112"></a>

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

<a id="canonical-1022022312320330-2020232320311002-0302332022123133-3312321232103201-2132220310310112-2110220133223032-2212001200321200-1013133110013103"></a>

## Direct properties — from_site / 321220133001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330101312223000-0302002331321233-0320330321332132-2330231031000033-1300120211130230-2010232111031133-3233100211011321-2230200322221110"></a>

## Next pages — from_site / 321220133001 / 4

- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-3111011321003313-1121200210332311-2321030331002223-2020202110220122-1132011030011110-0011030002211231-1101012321300323-0121122203002102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220230321012203-3000033332013133-0330010301013321-1021232121230201-1011031112323332-1331102031201202-0232132320203211-0030001300313200"></a>

## bgp_parameters.local_address — local_address / 320210110232 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011)
- bgp_parameters.local_address

<a id="canonical-3003311203202311-2133333121121312-3201223331032303-2012001200101020-0203112121330130-2300122132232312-0023232321321311-0320311302221223"></a>

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

<a id="canonical-2212002323122030-0003033332230323-3022233032112323-3030121022310103-2122100332301313-3111130101310121-3332201220212302-0323022133313031"></a>

## Direct properties — local_address / 320210110232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233011320032331-2132302112132331-3012010202223322-2323030020211113-2103020030231221-1103133230031101-0131300230123220-3011212202110123"></a>

## Next pages — local_address / 320210110232 / 4

- [bgp_parameters](data-sources--bgp--reference--group-001.md#canonical-1333100302301220-1213100230321130-1223021312221103-3121023012300002-0123222300312200-0012102122302300-3011220000310312-1112210021202011)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201310103133001-1220133113131211-3121233022232312-3002022103012033-1213110003321220-2000310233232011-3102130320202222-3223321232002130"></a>

## peers — peers / 123321201032 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- peers

<a id="canonical-0103320030001203-2210113302313322-2102201203303201-3220122113223211-1320132330302102-2212221010233321-3030230223112103-1020203200020231"></a>

Type: `"list"`. Computed.

Peers. List of peers.

Upstream description:

List of peers.

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

<a id="canonical-3121100300031000-2323123201301303-3202302230102210-1033313013121020-2112330122120203-2220010121132130-0212200131132030-0122313121203002"></a>

## Direct properties — peers / 123321201032 / 3

- [bfd_disabled](data-sources--bgp--reference--group-001.md#canonical-0333102131200232-0212101230300102-0120213222211232-1221030212211333-2011321012232211-0020030102113110-2031300103321002-3123101301312330): complete subsection reference.

- [bfd_enabled](data-sources--bgp--reference--group-001.md#canonical-0101313101233210-3201102200021201-0210020233311333-3010133210220310-3211312111020311-1323220303202223-0113132322000221-3301133030301130): complete subsection reference.

- [disable_spec](data-sources--bgp--reference--group-001.md#canonical-0222100230333202-3031200021030321-0023031002311302-1112213023120022-2331203113112212-0020100302211212-2323133023103320-0230102202000003): complete subsection reference.

- [ebgp_multihop_disabled](data-sources--bgp--reference--group-001.md#canonical-0033331103223012-3322132002320203-2131031123113112-2031130022302232-2311223023232020-3301220102313030-0221003000013223-0110212200100020): complete subsection reference.

- [ebgp_multihop_enabled](data-sources--bgp--reference--group-001.md#canonical-2011212023110130-0102112211203100-1011230333102231-2303022201230010-3100320310221030-0003220201231030-2302023023312232-1133320100323013): complete subsection reference.

- [external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302): complete subsection reference.

<a id="canonical-0123121322103020-0223020301300330-3211123132213213-0022323023121220-0233230123220203-3001022322331020-1210110333123222-1111311331003032"></a>

<a id="canonical-0332200010222311-1031311331210023-1332012231322211-0131130210311322-3033323002232331-3013011232320333-2303000301122031-2001023110120031"></a>

## label property — peers / 123321201032 / 4

Type: `"string"`. Computed.

Label. Specify whether this peer should be.

Upstream description:

Specify whether this peer should be.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [metadata](data-sources--bgp--reference--group-001.md#canonical-1210203212232330-2312201101103002-3021033131220003-3100323101002011-2333003323332203-2332113013012033-3302202221133313-1231022220113331): complete subsection reference.

- [passive_mode_disabled](data-sources--bgp--reference--group-001.md#canonical-3330222003020013-3002312001203013-3102312030323313-0223103232321233-1203310122113212-3222022000002200-3300330030300020-0011203212322220): complete subsection reference.

- [passive_mode_enabled](data-sources--bgp--reference--group-001.md#canonical-3111202111113022-2232110120120100-0211312103003231-2022332313131333-2020120311312200-0310213133111102-0030021111133110-1110203031113002): complete subsection reference.

- [routing_policies](data-sources--bgp--reference--group-001.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321): complete subsection reference.

<a id="canonical-0321300310211212-2221103310233211-1322232022011313-1203033203231223-0311333120300130-0332302201130313-3301211222303330-0012113222130221"></a>

## Next pages — peers / 123321201032 / 5

- [peers.bfd_disabled](data-sources--bgp--reference--group-001.md#canonical-0333102131200232-0212101230300102-0120213222211232-1221030212211333-2011321012232211-0020030102113110-2031300103321002-3123101301312330)
- [peers.bfd_enabled](data-sources--bgp--reference--group-001.md#canonical-0101313101233210-3201102200021201-0210020233311333-3010133210220310-3211312111020311-1323220303202223-0113132322000221-3301133030301130)
- [peers.disable_spec](data-sources--bgp--reference--group-001.md#canonical-0222100230333202-3031200021030321-0023031002311302-1112213023120022-2331203113112212-0020100302211212-2323133023103320-0230102202000003)
- [peers.ebgp_multihop_disabled](data-sources--bgp--reference--group-001.md#canonical-0033331103223012-3322132002320203-2131031123113112-2031130022302232-2311223023232020-3301220102313030-0221003000013223-0110212200100020)
- [peers.ebgp_multihop_enabled](data-sources--bgp--reference--group-001.md#canonical-2011212023110130-0102112211203100-1011230333102231-2303022201230010-3100320310221030-0003220201231030-2302023023312232-1133320100323013)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [peers.metadata](data-sources--bgp--reference--group-001.md#canonical-1210203212232330-2312201101103002-3021033131220003-3100323101002011-2333003323332203-2332113013012033-3302202221133313-1231022220113331)
- [peers.passive_mode_disabled](data-sources--bgp--reference--group-001.md#canonical-3330222003020013-3002312001203013-3102312030323313-0223103232321233-1203310122113212-3222022000002200-3300330030300020-0011203212322220)
- [peers.passive_mode_enabled](data-sources--bgp--reference--group-001.md#canonical-3111202111113022-2232110120120100-0211312103003231-2022332313131333-2020120311312200-0310213133111102-0030021111133110-1110203031113002)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0333102131200232-0212101230300102-0120213222211232-1221030212211333-2011321012232211-0020030102113110-2031300103321002-3123101301312330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230013133300210-1003221133221323-1121101032212323-3121222023120113-2321312010212130-3210120103202330-3300010032203322-3113321123122003"></a>

## peers.bfd_disabled — bfd_disabled / 213131333231 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.bfd_disabled

<a id="canonical-0300221110021101-1012301311120321-2203013110210213-3023010220313321-3021311222020031-1110300012020213-1131323131213002-1103331213012123"></a>

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

<a id="canonical-0112322030101123-3222202213230330-1032022232111223-3101102003132331-0223223301123302-1333131332312211-3203203010323320-1131121222302301"></a>

## Direct properties — bfd_disabled / 213131333231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021232002011111-3103001030023021-1201220023211310-1210312213111122-1000011211130002-1210312211221031-3230000300230223-2303201132222003"></a>

## Next pages — bfd_disabled / 213131333231 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0101313101233210-3201102200021201-0210020233311333-3010133210220310-3211312111020311-1323220303202223-0113132322000221-3301133030301130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032100311032321-2033003020230201-3123300012022100-3201333232313022-3201111110212333-0133030000203001-0211110333010000-0311202221013330"></a>

## peers.bfd_enabled — bfd_enabled / 331312100200 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.bfd_enabled

<a id="canonical-3000003210321332-2111213301021333-0202021220032121-3302212331100010-0030202111103110-3320330310320002-3223203213212331-2113311321001102"></a>

Type: `"single"`. Computed.

BFD. BFD parameters.

Upstream description:

BFD parameters.

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

<a id="canonical-1122130012013122-0310323031131203-1202330302320210-2121022013012030-0132311323013231-1122002031300331-2032320010130333-2023011013121022"></a>

## Direct properties — bfd_enabled / 331312100200 / 3

<a id="canonical-0320301111200211-3002313200232102-3122001200012030-0012121011101123-2232100323031211-3222302312103200-2212030310120300-3032223111213120"></a>

<a id="canonical-2300301113302221-0331330300301112-2122213333120221-1030123332332032-2120210202013030-3300210110130322-2112203121332231-0331000113202102"></a>

## multiplier property — bfd_enabled / 331312100200 / 4

Type: `"number"`. Computed.

Specify Number of missed packets to bring session down'.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2310110323203223-1121001101111120-0203201331323212-2123033013222001-1120201000011131-2130303311330000-0012301102112130-3201112222331333"></a>

## receive_interval_milliseconds property — bfd_enabled / 331312100200 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1020233121221323-0021130300023331-3013103010232210-3321132302111213-1032220323100033-1311221303301012-3201023223322221-3202303002300323"></a>

## transmit_interval_milliseconds property — bfd_enabled / 331312100200 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0320321021201301-2031201103022313-1013332221311131-0121021130303200-2203203003020331-3002220132331323-3310330000011122-2211122310213122"></a>

## Next pages — bfd_enabled / 331312100200 / 7

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0222100230333202-3031200021030321-0023031002311302-1112213023120022-2331203113112212-0020100302211212-2323133023103320-0230102202000003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221131231333011-0332021113021212-0202102322100100-3322332031111100-3001103011020312-1021133001110323-3312033012330103-3103121301302033"></a>

## peers.disable_spec — disable_spec / 302221330102 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.disable_spec

<a id="canonical-0203021110002310-2201313303232131-1222220022021201-0231020100122210-1000233300020301-0201021231321211-3122011201122330-2222030322230120"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-0333213022210200-0202223310221121-0023110303121130-2131002312113231-2203012100023321-2133130110111221-0213330011302303-2201020223122212"></a>

## Direct properties — disable_spec / 302221330102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312333123123001-3323121023213111-0002033222111201-1313212301331120-2113223000022012-0011322021302112-3213111310232211-2013132131130022"></a>

## Next pages — disable_spec / 302221330102 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0033331103223012-3322132002320203-2131031123113112-2031130022302232-2311223023232020-3301220102313030-0221003000013223-0110212200100020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110012201022131-0002021111310122-2132110222220301-1302022112020330-1031210300313302-2130321120203113-0132002001110230-0203002033003321"></a>

## peers.ebgp_multihop_disabled — ebgp_multihop_disabled / 020212031010 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.ebgp_multihop_disabled

<a id="canonical-2013322032310313-2200222232011221-1201021232113031-2213003001023302-3330013301212322-1003011103110133-0301010030201033-2200211220232021"></a>

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

<a id="canonical-0002330213001032-2310033121103003-3333133320032211-2303021233010322-2130200210102121-0013330332121000-0313100332210211-0303202323132030"></a>

## Direct properties — ebgp_multihop_disabled / 020212031010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311123123000233-0123233302110013-3011303021031302-0212023130303232-1112332123002002-0322211220021232-1212000312012230-3203312202103021"></a>

## Next pages — ebgp_multihop_disabled / 020212031010 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2011212023110130-0102112211203100-1011230333102231-2303022201230010-3100320310221030-0003220201231030-2302023023312232-1133320100323013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211013203001101-3301301320233312-2331313200110303-3100001233332032-3000012022220131-2021313101223332-3210010010031023-1303303103020033"></a>

## peers.ebgp_multihop_enabled — ebgp_multihop_enabled / 302132313122 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.ebgp_multihop_enabled

<a id="canonical-2331033000012203-2012132301332310-2212222230221131-0102311301012321-0320221213212033-3202303311133120-3331310330001123-3111013102101322"></a>

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

<a id="canonical-3331322312100102-1032212210022110-0013313102323010-2313101012023112-2220123033111001-2123202202201112-0333322232032303-3211312232002132"></a>

## Direct properties — ebgp_multihop_enabled / 302132313122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323120010200131-2121230101100112-3200302000331012-1323021210302121-0002110201003030-3112333331202032-1032210303131131-0200033010001333"></a>

## Next pages — ebgp_multihop_enabled / 302132313122 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303103023013131-2332032123220223-1310221222232021-3022110023233102-3133001033123030-0010233220201122-2333003111202213-0222133000231120"></a>

## peers.external — external / 101020232231 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.external

<a id="canonical-3320231030122312-3230203312212100-3333203132111222-3311213103001021-2022203123012203-1011331220003231-2010031220022321-1320320010223003"></a>

Type: `"single"`. Computed.

External BGP Peer. External BGP Peer parameters.

Upstream description:

External BGP Peer parameters.

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

<a id="canonical-1200201321310132-3211301012001203-2321032111223003-2002321113301111-3230121022300213-2121102110030222-0022212302321320-3120212221132330"></a>

## Direct properties — external / 101020232231 / 3

<a id="canonical-1121220230230023-3213011210120231-0102010300130112-0303131301332020-1230230112102101-1132110000321003-3101132301012302-1101112213021321"></a>

<a id="canonical-1033330301000101-0033003021031200-3323121021330133-1203310313231013-3332030313330320-1203100323133311-2031202113003113-2310302022130302"></a>

## address property — external / 101020232231 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

Upstream description:

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

<a id="canonical-1102220231332201-3320203031033103-1031201310101230-0131110031121301-2000111332122002-0032232010003210-1321303223011302-2312232320011032"></a>

<a id="canonical-2121202233132010-1312203031102102-1111011332131230-0120110121100210-0100310311222100-0311031132113202-1223101211032031-2030203103020002"></a>

## address_ipv6 property — external / 101020232231 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

Upstream description:

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

<a id="canonical-1021222233001232-3220033123003030-1122221202010133-1323031323033030-1202122301332312-0222201201310012-1311330012120330-2020123033130331"></a>

<a id="canonical-1011132323302201-2111032022002332-1323203102013133-1132211133223231-1123130312030232-0002121303101213-1220212222231023-0101223112122032"></a>

## asn property — external / 101020232231 / 6

Type: `"number"`. Computed.

ASN. Autonomous System Number for BGP peer.

Upstream description:

Autonomous System Number for BGP peer.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [interface_list](data-sources--bgp--reference--group-001.md#canonical-0332112000321201-2012130032100313-3003023100112021-1221001212101031-3222213330302233-1032331013323132-1312222321313302-1202133211002223): complete subsection reference.

<a id="canonical-1022323110021200-0330120000303001-1012032232001011-3022133113213030-0102321230012231-1202223211331020-0012002320210121-3030100212232303"></a>

<a id="canonical-1220211102303321-2133032303321013-2133123210223001-1132103331001232-3213220021033033-0102130023123022-0130333201133113-0123312110110331"></a>

## md5_auth_key property — external / 101020232231 / 7

Type: `"string"`. Computed.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

Upstream description:

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385)

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

- [no_authentication](data-sources--bgp--reference--group-001.md#canonical-1120030012320211-2222321211313311-2102010233321200-2013312131010001-1321330100100030-3300110231320011-3010023322032023-0111033230310112): complete subsection reference.

<a id="canonical-1313033220213200-1021023010131200-1102300113011113-1333331330233300-0233201331311033-0322120202013200-0232221300330301-1132031301323333"></a>

<a id="canonical-3202302013330300-3202030001001133-1233130020012112-1101021012223230-3101211210132023-2013030320031220-1210233232022130-2003233032332222"></a>

## port property — external / 101020232231 / 8

Type: `"number"`. Computed.

Peer Port. Peer TCP port number.

Upstream description:

Peer TCP port number.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2331120100013221-0023022333021102-1310023113232103-2132110031112022-3022002213130330-0313223000200212-2130212200313012-1331331013100201"></a>

<a id="canonical-0002111023113033-1200022210100123-3201113333323303-0101332020202131-3001103123011330-0020122200013312-1332332231222133-1002310221131231"></a>

## subnet_begin_offset property — external / 101020232231 / 9

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3200231103310132-3321113110322233-3021321221101303-1100320303200102-0212002332023132-2122121100213123-3000223210201120-2210311211003232"></a>

## subnet_begin_offset_v6 property — external / 101020232231 / 10

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1131222010112200-0220021330023000-0122010123333132-2222313202201102-3021110033013003-2332221123222010-3010122230030033-1221133301230112"></a>

## subnet_end_offset property — external / 101020232231 / 11

Type: `"number"`. Computed.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0202010331313302-3110330212211201-2332133010010033-2021000102201112-0021102101321300-3200010323320001-1330213202011111-0112111212030031"></a>

## subnet_end_offset_v6 property — external / 101020232231 / 12

Type: `"number"`. Computed.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0002100101000120-3000022002333000-1220200213030220-1110103120233000-1231303330201113-0213130123330013-3120211220012101-2202011313031022"></a>

## Next pages — external / 101020232231 / 13

- [peers.external.default_gateway](data-sources--bgp--reference--group-001.md#canonical-3302111002200333-2232131201012220-3012131031221203-0231033330123033-0312303122301332-2030003200003033-3103100223231301-2121203020011020)
- [peers.external.default_gateway_v6](data-sources--bgp--reference--group-001.md#canonical-2002320231121031-3100100212113030-0230121333202013-1323223233323022-0131021013202001-0010023003120312-2003102220233200-2321100213130002)
- [peers.external.disable_spec](data-sources--bgp--reference--group-001.md#canonical-2031302223232031-3332022333300230-3232333112123211-2201223032303231-3322002212022102-0200022231100301-1301203212110233-2301232231103002)
- [peers.external.disable_v6](data-sources--bgp--reference--group-001.md#canonical-0021110022103030-1111002111033331-1201033111231011-0322330321011013-3332210111100103-0332031311332332-0030111113113210-1201020022013330)
- [peers.external.external_connector](data-sources--bgp--reference--group-001.md#canonical-2322223132011310-1220013012220232-3121311110023132-2333123201301333-0122201202231200-3303102021221130-3200300331032213-3301121313301300)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030)
- [peers.external.from_site](data-sources--bgp--reference--group-001.md#canonical-2000201023230011-1131222320032101-0211021122331001-3211313212301012-3230211211100130-2220222213320110-3311003212020111-2221133201020221)
- [peers.external.from_site_v6](data-sources--bgp--reference--group-001.md#canonical-2330211123232010-0122030211321121-1030113211210123-0233323001233222-0302103102003332-3223201002013231-0103012003131213-3010020320120111)
- [peers.external.interface](data-sources--bgp--reference--group-001.md#canonical-1202300022003132-3311301221023101-2320321113301022-1223131323301322-2221113331121331-3110002032220221-1230201213100311-1330112310021023)
- [peers.external.interface_list](data-sources--bgp--reference--group-001.md#canonical-0332112000321201-2012130032100313-3003023100112021-1221001212101031-3222213330302233-1032331013323132-1312222321313302-1202133211002223)
- [peers.external.no_authentication](data-sources--bgp--reference--group-001.md#canonical-1120030012320211-2222321211313311-2102010233321200-2013312131010001-1321330100100030-3300110231320011-3010023322032023-0111033230310112)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-3302111002200333-2232131201012220-3012131031221203-0231033330123033-0312303122301332-2030003200003033-3103100223231301-2121203020011020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112011012203232-1300221300203111-3330200120101020-2312230121112013-2122323210310333-3031132112033231-3112320011221301-0300231222012321"></a>

## peers.external.default_gateway — default_gateway / 131231101211 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.default_gateway

<a id="canonical-0021121311303000-3010033321120203-0010010301131203-0112201030220313-1110220021022303-0203022300303211-1111000130201231-2131033222220202"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2103121311232323-1200222302101003-3331121013332232-1000303023133000-3100202021331103-0321312300032022-3112010011331022-1313301123022020"></a>

## Direct properties — default_gateway / 131231101211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010021032202201-3012103332223213-1230032030320011-0322211320132213-2301231033213203-3002101030020230-2022323230233011-0212123230033200"></a>

## Next pages — default_gateway / 131231101211 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2002320231121031-3100100212113030-0230121333202013-1323223233323022-0131021013202001-0010023003120312-2003102220233200-2321100213130002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303033110232121-0101111130002332-0033112211033221-3100311122122201-2010103223033003-1301222331320011-1230321001323131-3012003311200201"></a>

## peers.external.default_gateway_v6 — default_gateway_v6 / 312003212032 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.default_gateway_v6

<a id="canonical-0230120223123220-0100311333320220-1133321230112230-1230320010103221-2033233332121321-1031101113331011-2223230020321000-3111000010010000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway v6.

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

<a id="canonical-2300313302313100-1032123012221300-0022220013213013-0033201110000131-0233313303010201-3110030213133301-0110023020002030-1013103213131201"></a>

## Direct properties — default_gateway_v6 / 312003212032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030011023213032-2100103202002302-2130131022300223-3331212331001202-0010330222103213-2013100112230331-0012220223111001-3202302201120322"></a>

## Next pages — default_gateway_v6 / 312003212032 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2031302223232031-3332022333300230-3232333112123211-2201223032303231-3322002212022102-0200022231100301-1301203212110233-2301232231103002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031221003130333-3010223331303231-1001033123203130-1120131221010201-3102111220010203-3201023213103012-2131023133012300-2220010312320121"></a>

## peers.external.disable_spec — disable_spec / 012221101200 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.disable_spec

<a id="canonical-3133330231313202-3100112210013201-2101333233330211-0100032220012232-2000013313332032-0200320232300022-1101232103032210-2301031203101211"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-0103303112300203-0033030113001102-1323013212122023-3310313323100022-0312032213301211-2121002322032022-1133230113031333-0233110203330333"></a>

## Direct properties — disable_spec / 012221101200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303021103320030-1312203302320331-1320222003333013-1302013233310100-0032233032313333-1310211013211222-0133202312213123-2020130211003111"></a>

## Next pages — disable_spec / 012221101200 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0021110022103030-1111002111033331-1201033111231011-0322330321011013-3332210111100103-0332031311332332-0030111113113210-1201020022013330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113013123120010-2221021312320320-1310210311032001-2003100112103202-2331031101231113-0020112110032301-0230103003302032-2322310302011221"></a>

## peers.external.disable_v6 — disable_v6 / 233020121020 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.disable_v6

<a id="canonical-0321223012111320-0031213300123313-1101323110233032-3303111020121023-1032023301312131-2210120330030103-3213302120013000-2222013222331312"></a>

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

<a id="canonical-2330332033202032-2311211122022331-2231223130032100-0013032212222223-2210210012322213-2200231033030222-0220321102330312-2303020331233000"></a>

## Direct properties — disable_v6 / 233020121020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033003210213313-1013201201331301-2112121302113230-2322303311232310-2200021131001002-2102332001022221-3112123202113333-1331231220032302"></a>

## Next pages — disable_v6 / 233020121020 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2322223132011310-1220013012220232-3121311110023132-2333123201301333-0122201202231200-3303102021221130-3200300331032213-3301121313301300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022032202330202-2011312002201302-3300021332201120-1112003102122013-3033110302023031-1132123203003323-1102001331230323-2131220210120023"></a>

## peers.external.external_connector — external_connector / 222311101110 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.external_connector

<a id="canonical-2100122021212133-2310011032311121-0102303333202023-1312110013301302-3222003313303312-2131033011021210-2223002023203330-2230332021100022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for external connector.

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

<a id="canonical-2110121013222302-3133121223211030-1030133320320110-0013002021330023-2203031112101213-1013032002330000-2103102133113002-0333121301312032"></a>

## Direct properties — external_connector / 222311101110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203103103011230-1201100032022031-1132013013021123-0121111313120211-3301002100220200-3320233133133001-3301121220121020-1313231212111213"></a>

## Next pages — external_connector / 222311101110 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033302112212320-2213130211011032-3113331213213320-3311022330121313-3333302030010121-3200003200212220-3233222211111311-1210031021120231"></a>

## peers.external.family_inet — family_inet / 121130103123 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.family_inet

<a id="canonical-2000203102221113-2101022100011023-0203212133231212-1111221012312330-1212302100213231-1203111013200120-0101011132101011-2320021130201323"></a>

Type: `"single"`. Computed.

Configuration parameter for family inet.

Upstream description:

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

<a id="canonical-0221122013301233-2123301111113331-1130021103301033-0100331221102332-1201203210200223-1302020212301101-1122110331213210-3132001010332222"></a>

## Direct properties — family_inet / 121130103123 / 3

- [disable_spec](data-sources--bgp--reference--group-001.md#canonical-1103102203233201-1000030232002020-3020300023300300-1201301123103121-0320312313212021-1210101023120212-1100103123121332-1312202122321302): complete subsection reference.

- [enable](data-sources--bgp--reference--group-001.md#canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132): complete subsection reference.

<a id="canonical-2000023232111312-2001111201331222-1220103310223003-2132122300102030-0313110323131210-2202011033133100-1110220020113203-0323101000303201"></a>

## Next pages — family_inet / 121130103123 / 4

- [peers.external.family_inet.disable_spec](data-sources--bgp--reference--group-001.md#canonical-1103102203233201-1000030232002020-3020300023300300-1201301123103121-0320312313212021-1210101023120212-1100103123121332-1312202122321302)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1103102203233201-1000030232002020-3020300023300300-1201301123103121-0320312313212021-1210101023120212-1100103123121332-1312202122321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223331323131313-2112111032031203-0121033003011322-3102210311201011-0111200330321310-0212321311301012-3222323222110033-2313221210101320"></a>

## peers.external.family_inet.disable_spec — disable_spec / 201003202320 / 2

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

<a id="canonical-0211302321011021-3312203122322313-1203311312133313-3130133100210301-1322212101220100-2301223012231302-0230001310133101-1102312320330311"></a>

## Direct properties — disable_spec / 201003202320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031112001121031-0011320331033000-3222113001220110-1300112233003313-1112103020202322-2202022201111320-1203213120211313-1233332212310212"></a>

## Next pages — disable_spec / 201003202320 / 4

- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132132201030302-1121121202211223-2130213131011302-0031201101131311-0233002311030133-3131210212321210-1012321223233113-3302013000030231"></a>

## peers.external.family_inet.enable — enable / 003030132320 / 2

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

Upstream description:

IPv4 Unicast.

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

<a id="canonical-0311132211323233-2331313213113021-1002101210011222-2323012100221220-2022222013112113-0310320213322223-2312323111110333-0031112200213123"></a>

## Direct properties — enable / 003030132320 / 3

- [aggregation](data-sources--bgp--reference--group-001.md#canonical-0302321100202233-0231213322311003-3030221203212302-3100323132212200-2120102003121311-2133210022223132-2310332021103113-2202223020221302): complete subsection reference.

<a id="canonical-3230123202321201-3030320213013330-2211121123320231-3212102321010122-2022123312012002-3233022320120212-0000203322032312-0000023223302003"></a>

## Next pages — enable / 003030132320 / 4

- [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-0302321100202233-0231213322311003-3030221203212302-3100323132212200-2120102003121311-2133210022223132-2310332021103113-2202223020221302)
- [peers.external.family_inet](data-sources--bgp--reference--group-001.md#canonical-3233033203221203-2023022102020230-2121122323230102-0213221302313323-2301011130322321-3303102220230022-1001231111331110-0201223330323030)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0302321100202233-0231213322311003-3030221203212302-3100323132212200-2120102003121311-2133210022223132-2310332021103113-2202223020221302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301103110313120-0212120303233120-0303230201320223-1303103002133321-1332132330101113-2300110330223132-2302321130222232-2023320121102212"></a>

## peers.external.family_inet.enable.aggregation — aggregation / 123311030002 / 2

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

<a id="canonical-0212210130320123-3220331320313312-1110232032101222-1011213131031021-3202331321233211-2333301003322131-1321000111233001-1011100200321012"></a>

## Direct properties — aggregation / 123311030002 / 3

<a id="canonical-2113101301030013-2200231010322033-1303310112230033-1312220301132320-2200103110201332-0002311002113333-0033032103011101-3111102300220111"></a>

<a id="canonical-0000223233230300-2201110210131031-1213300031131020-3012330231312013-0300111322231123-1131130331300222-2000120022301103-0332310013311323"></a>

## ip_prefix property — aggregation / 123311030002 / 4

Type: `"string"`. Computed.

IP Prefix. Specify IPv4 subnet for aggregation.

Upstream description:

Specify IPv4 subnet for aggregation.

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

<a id="canonical-0011202231000133-0000313201120202-1022001333213330-3333232123101112-1000332222123100-0012111113322302-1213113121021232-2122210320232101"></a>

## Next pages — aggregation / 123311030002 / 5

- [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--reference--group-001.md#canonical-1013013330322010-1201100003203212-2020223121313320-1233121112122021-3123030203103021-0211123113202323-3320211201101133-0133311123111331)
- [peers.external.family_inet.enable](data-sources--bgp--reference--group-001.md#canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1013013330322010-1201100003203212-2020223121313320-1233121112122021-3123030203103021-0211123113202323-3320211201101133-0133311123111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323203201011220-2231030130113133-1010113201023211-3030133101332013-2200303132333220-1322130013023131-3300312220330030-3203323030002110"></a>

## peers.external.family_inet.enable.aggregation.options — options / 112300231213 / 2

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

Upstream description:

Configuration parameter for options

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

<a id="canonical-3131013232332201-2311010030113102-3321020233103102-1301223112022332-1323303030031103-2332233032123021-0122022011013033-0022133212031210"></a>

## Direct properties — options / 112300231213 / 3

- [summary_only](data-sources--bgp--reference--group-001.md#canonical-2112132320203310-2103233321210210-1203102031323320-0200030213203021-1131321210103322-1333003103000332-3003033001020110-1333220002231232): complete subsection reference.

<a id="canonical-0003300132023213-2223311231332301-1001202211130333-1320031232101002-0121203000221130-2323203330221221-3012113003010010-2123330010301033"></a>

## Next pages — options / 112300231213 / 4

- [peers.external.family_inet.enable.aggregation.options.summary_only](data-sources--bgp--reference--group-001.md#canonical-2112132320203310-2103233321210210-1203102031323320-0200030213203021-1131321210103322-1333003103000332-3003033001020110-1333220002231232)
- [peers.external.family_inet.enable.aggregation](data-sources--bgp--reference--group-001.md#canonical-0302321100202233-0231213322311003-3030221203212302-3100323132212200-2120102003121311-2133210022223132-2310332021103113-2202223020221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2112132320203310-2103233321210210-1203102031323320-0200030213203021-1131321210103322-1333003103000332-3003033001020110-1333220002231232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222303230221312-1201000333113013-0133003001013122-1010002233230312-1131200101101133-1212132311123010-3030030220001120-0322333330231322"></a>

## peers.external.family_inet.enable.aggregation.options.summary_only — summary_only / 233323133001 / 2

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

<a id="canonical-0022001110300013-3210110213030230-3021332020331213-2032312012133111-1101023332011032-1031013112213021-3232220303033331-2010322122131232"></a>

## Direct properties — summary_only / 233323133001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332302030331211-3122101222121002-0203101202211103-1032122023133323-2213200103123222-0332221303103102-1211031030301000-2300033132003320"></a>

## Next pages — summary_only / 233323133001 / 4

- [peers.external.family_inet.enable.aggregation.options](data-sources--bgp--reference--group-001.md#canonical-1013013330322010-1201100003203212-2020223121313320-1233121112122021-3123030203103021-0211123113202323-3320211201101133-0133311123111331)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2000201023230011-1131222320032101-0211021122331001-3211313212301012-3230211211100130-2220222213320110-3311003212020111-2221133201020221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200111032332120-2102200301213012-1010223320331003-0003130322031333-1222103003123232-1100313312121331-3111021221020011-2333310111233130"></a>

## peers.external.from_site — from_site / 330121211131 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.from_site

<a id="canonical-2002331301301011-0003200103023021-0231312233311312-0203010210323212-1123231320002032-2331001223332102-0301020233331230-3301302011031132"></a>

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

<a id="canonical-3023111320000312-3013301222203010-0211202023023103-3221013130331102-3323313113003100-3123000312231130-3132010010330002-0001130130303101"></a>

## Direct properties — from_site / 330121211131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113021022303002-1322121021130331-3320003020001130-3103013223103031-3110303231312301-2302323110322232-3133011010121331-0220130322201110"></a>

## Next pages — from_site / 330121211131 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2330211123232010-0122030211321121-1030113211210123-0233323001233222-0302103102003332-3223201002013231-0103012003131213-3010020320120111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031102233122013-3012321112000333-3113021113231302-1203100313310003-1032001203021200-0123023322130023-0313333300231313-1100322320100303"></a>

## peers.external.from_site_v6 — from_site_v6 / 320321202032 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.from_site_v6

<a id="canonical-2102020101131000-0201302113220000-3211131132002002-3032302112200321-0310302230330331-1200331033232313-3220332123202331-2022030202320211"></a>

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

<a id="canonical-1022301200003210-0132021020031121-0132332231233301-2103031100001122-1311011022223113-0330300121321010-3321023032121301-3310313222223332"></a>

## Direct properties — from_site_v6 / 320321202032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100320010211103-3110122202310003-0210213000021302-1332033210231033-2332223120303220-3030003232310313-2203002200201231-0110203223033032"></a>

## Next pages — from_site_v6 / 320321202032 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1202300022003132-3311301221023101-2320321113301022-1223131323301322-2221113331121331-3110002032220221-1230201213100311-1330112310021023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331031121332031-3120131322220231-3031002220012312-0013113133320330-0330122002213003-3213221112312103-0102113100310121-1023332320223031"></a>

## peers.external.interface — interface / 121330332202 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.interface

<a id="canonical-0311231232312012-1133023122013130-2323110300021310-3331130021212330-0301332310033231-1033030333021030-3202131222100233-3311100302310033"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0303101132100212-3332002022311120-0231032120223302-1131011102232211-3223100221311202-0120302320300312-0331032031111310-3303210120322103"></a>

## Direct properties — interface / 121330332202 / 3

<a id="canonical-2331322113222030-3123110310312212-1203320112130302-0203013011221133-1003100302022311-2003112003131000-0312113110323013-3120010003011310"></a>

<a id="canonical-3120110311230233-1012100230332300-0323322132002030-3301303321232002-0320132132302230-0112231210023330-0203332102321111-1031200322202123"></a>

## name property — interface / 121330332202 / 4

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

<a id="canonical-3213202110102213-0221220010330000-0001331233033220-3310201301031103-0000321213232023-2231333033011212-3210301121200112-0133120200100201"></a>

<a id="canonical-2022320022013023-0211031122132320-2303113102213322-1031000211032221-2223320113222321-3133022120223031-2031310023232230-0323130133032320"></a>

## namespace property — interface / 121330332202 / 5

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

<a id="canonical-2010302211213103-3013002221020001-3332112312311130-0130302200213211-3313300010320012-3321312021013003-3020320302330112-3122120320101211"></a>

<a id="canonical-3333300331102311-2211311132131032-3221111120323221-2121310212012222-0232120310023212-3303002333031113-0203233222101110-2121201101223131"></a>

## tenant property — interface / 121330332202 / 6

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

<a id="canonical-0102221133321111-0223211021322022-0032332211013311-1233110303322300-2031112222021200-2322110220010100-2110002323032303-0201210103210200"></a>

## Next pages — interface / 121330332202 / 7

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0332112000321201-2012130032100313-3003023100112021-1221001212101031-3222213330302233-1032331013323132-1312222321313302-1202133211002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002231032131311-0310213320100323-2101332303032233-3212232012132123-2012132130323122-3110332222003110-3003103012001011-0002022300221301"></a>

## peers.external.interface_list — interface_list / 201233230300 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.interface_list

<a id="canonical-1131003102033211-0333221032220223-0023200032231100-2013300112301130-0031303020113131-2010322132133300-1012200303023331-3131102012332300"></a>

Type: `"single"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

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

<a id="canonical-2223022202231031-3113233200233103-1202201233031230-0023132031030030-2112021100022222-2131330200121233-3112233202002300-1111110122130003"></a>

## Direct properties — interface_list / 201233230300 / 3

- [interfaces](data-sources--bgp--reference--group-001.md#canonical-1330200232221112-0233331010013003-0133020022210001-0001002000131013-3122110233112112-0130321300320101-3213100313302113-3203002330022323): complete subsection reference.

<a id="canonical-3332030122321311-3303203331021303-0220023112132211-1013200012030213-0330003130211010-3002302333101003-0223323132333210-2023301031102322"></a>

## Next pages — interface_list / 201233230300 / 4

- [peers.external.interface_list.interfaces](data-sources--bgp--reference--group-001.md#canonical-1330200232221112-0233331010013003-0133020022210001-0001002000131013-3122110233112112-0130321300320101-3213100313302113-3203002330022323)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1330200232221112-0233331010013003-0133020022210001-0001002000131013-3122110233112112-0130321300320101-3213100313302113-3203002330022323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122102322121023-2101011021122303-2021200213312113-2300313202033300-1031020201101130-3122222210331120-0011130033311201-0322211122013113"></a>

## peers.external.interface_list.interfaces — interfaces / 211303212312 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [peers.external.interface_list](data-sources--bgp--reference--group-001.md#canonical-0332112000321201-2012130032100313-3003023100112021-1221001212101031-3222213330302233-1032331013323132-1312222321313302-1202133211002223)
- peers.external.interface_list.interfaces

<a id="canonical-0110222120101112-0312332110020303-1233001231211312-2011000133033220-1200203201231211-0133010013110113-1111210002300013-3133233300223320"></a>

Type: `"list"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

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

<a id="canonical-2231212033202200-0310112200321321-0311310123022221-1002323310111032-2301332122232313-2010200201320310-3021200332121302-2310332020202120"></a>

## Direct properties — interfaces / 211303212312 / 3

<a id="canonical-1010031202002133-3313203100130222-0330201013102133-3031332223233202-3132211020110213-1210332120130013-0232322312021001-2323021112313303"></a>

<a id="canonical-1002232232330320-2121231332220221-0300121310300121-0110231001003102-0100111113031000-1333212010301033-0213120010322123-1013210032300110"></a>

## name property — interfaces / 211303212312 / 4

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

<a id="canonical-2323113102032032-2221020100133310-0133010002111130-1222300222302212-3201123212100301-2332331203000113-3013322232132321-0020130332113212"></a>

<a id="canonical-2332112322001221-2321001102102202-1133202231220101-0203223313033322-0032322133211322-1100012100201011-1313310012111331-3103100313333000"></a>

## namespace property — interfaces / 211303212312 / 5

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

<a id="canonical-2212312010113133-3333030010201100-2313310233320222-2103000213010102-1011203222231212-0333022010102102-3221223100331033-3011010022232222"></a>

<a id="canonical-3201232201201110-0322303101322320-0202202031112111-2122131330321102-2121301211132300-1020232200123111-0100030002310123-3023120220101031"></a>

## tenant property — interfaces / 211303212312 / 6

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

<a id="canonical-0003133110220102-2213032331110012-3133100032010121-1223012102012233-2212302122200103-0331311123013100-0013002121302022-2310111220201003"></a>

## Next pages — interfaces / 211303212312 / 7

- [peers.external.interface_list](data-sources--bgp--reference--group-001.md#canonical-0332112000321201-2012130032100313-3003023100112021-1221001212101031-3222213330302233-1032331013323132-1312222321313302-1202133211002223)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1120030012320211-2222321211313311-2102010233321200-2013312131010001-1321330100100030-3300110231320011-3010023322032023-0111033230310112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001023100031321-1031303331012110-1320212100112211-0013100001211013-1121330122112320-2202221120302131-3031132022331303-0111221013333123"></a>

## peers.external.no_authentication — no_authentication / 202101033233 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- peers.external.no_authentication

<a id="canonical-3010203300113023-1231010032331030-3213303023310212-2312220222312332-2012023013111300-1011231132301000-0211223031231302-1123031030321122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no authentication.

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

<a id="canonical-0023231203130323-2212133220313003-3320332020312030-0131102233210312-1213301033311321-1130013321131102-2121302100303232-0031332111332202"></a>

## Direct properties — no_authentication / 202101033233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200020330201102-1233032233003200-0010210031123133-3232032233010031-1302231312313321-0231333201212033-2103111233031333-2023223212311311"></a>

## Next pages — no_authentication / 202101033233 / 4

- [peers.external](data-sources--bgp--reference--group-001.md#canonical-1100310013200210-1203020110333122-1021132031333130-2101223233233213-1131110231302303-0103022103010102-1213030130001120-2231321011221302)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1210203212232330-2312201101103002-3021033131220003-3100323101002011-2333003323332203-2332113013012033-3302202221133313-1231022220113331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333001012233302-0111103302032010-2301003231013203-3211321012123013-1011300323331311-3222110130330130-3032102033102300-3000130312300312"></a>

## peers.metadata — metadata / 311301312103 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.metadata

<a id="canonical-0323110232012220-2323021101313302-1131100021310013-2300111000230232-0220333233310311-2210133023310010-2322001020321033-2000302001210002"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-2133332301302101-2232320110110020-2001300110103222-1000210330201203-0210131101002132-3322003031011012-1130222223331312-0313310223221323"></a>

## Direct properties — metadata / 311301312103 / 3

<a id="canonical-2223311023010213-3121010210313020-2030220320232011-1200132130213032-0322313212333123-0320122132122203-1322103120101210-1031031231120232"></a>

<a id="canonical-3033233032123310-2222203122212301-3133030332323031-1121032033001101-1203231323312200-3022023001022130-1211113123310232-2232031302231220"></a>

## description_spec property — metadata / 311301312103 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2302013311232220-1132332100102310-3323223320113000-2000310023001102-0300130112031333-2021013310233013-0130311133130201-1303102322103210"></a>

<a id="canonical-0322003133210022-3120112303012212-1033333131010331-0303010303132202-2210202110011200-2031110303222233-2120233113100000-2311130012032200"></a>

## name property — metadata / 311301312103 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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

<a id="canonical-0323233111221201-3211002321121220-1202032130123210-0001331131211120-1033031310202131-1032130303220012-0130000312120331-1111130301130122"></a>

## Next pages — metadata / 311301312103 / 6

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-3330222003020013-3002312001203013-3102312030323313-0223103232321233-1203310122113212-3222022000002200-3300330030300020-0011203212322220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103103302300202-2331011313103003-0323102223133310-1111003333130123-3333133013322000-2311200130010112-2133002232132002-2112120110122333"></a>

## peers.passive_mode_disabled — passive_mode_disabled / 120021333102 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.passive_mode_disabled

<a id="canonical-2003303132030221-3021102113031322-3201232230221211-0232032133210320-2002211032030020-3020202121302320-0332212123132202-0113001321121222"></a>

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

<a id="canonical-1312332011222100-2330100233130333-3233312302203323-2312322101233333-0033213203222323-2102000312222020-2000001003023113-2323131223010310"></a>

## Direct properties — passive_mode_disabled / 120021333102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123223321312300-0122212323233221-1210021231001200-1120233022131030-2021122210021132-2201132212203212-3013120010322300-0212122023102011"></a>

## Next pages — passive_mode_disabled / 120021333102 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-3111202111113022-2232110120120100-0211312103003231-2022332313131333-2020120311312200-0310213133111102-0030021111133110-1110203031113002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211230000222023-1323202121303101-2211001131023032-0222123211103331-3223200211313013-3322200001211311-1223011121232210-0021231333202033"></a>

## peers.passive_mode_enabled — passive_mode_enabled / 202010332302 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.passive_mode_enabled

<a id="canonical-1120310213202313-2103130203003131-1220330133013100-0203002310202010-0302102131313121-2111323012030323-3212030222032012-1210112221231110"></a>

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

<a id="canonical-1013320132032302-2101031233122230-1120220330010102-3311020120101331-1032311122332002-0222222112230220-2033001110033222-3032231233230021"></a>

## Direct properties — passive_mode_enabled / 202010332302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022332212033110-0013031112323300-2210220123302021-3330132220301100-0011210121203332-0303222130110102-0003123312011002-1100233122132311"></a>

## Next pages — passive_mode_enabled / 202010332302 / 4

- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312230123101233-0231122220330203-3031000020110222-2310303021133011-1103012320002223-3012201203121133-2130223133133131-2222123131333332"></a>

## peers.routing_policies — routing_policies / 233132033002 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- peers.routing_policies

<a id="canonical-0303322023121133-1100203002202001-0220122012203033-3113131201321322-0203132102321222-3232121332030323-3031221322011110-2303100013303102"></a>

Type: `"single"`. Computed.

List of rules which can be applied on all or particular nodes.

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

<a id="canonical-3013000011000221-1123121002210022-3121301030321030-1032312333221311-2123130302232203-2023102003330003-1022212233313331-2233103123211331"></a>

## Direct properties — routing_policies / 233132033002 / 3

- [route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213): complete subsection reference.

<a id="canonical-1213032032033133-0103111322103323-1202110220000210-1133230311012102-3302333301212302-2211133122202331-0203320131101102-2202020332131100"></a>

## Next pages — routing_policies / 233132033002 / 4

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131013011223233-2023123032323003-2332130031133020-0320001213303303-1332213111321300-1201101203132323-2112033111100131-1130100203302031"></a>

## peers.routing_policies.route_policy — route_policy / 201303311233 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- peers.routing_policies.route_policy

<a id="canonical-1212222011200010-0211021023032113-0310303323230131-2022133121130021-2110310322233300-1233132303322002-3201202302003110-0331133131312320"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-2112333201101332-1131210012033112-2032312110123200-3332133031103002-0212132301031032-1202002321131113-3101222022031300-1230320210012301"></a>

## Direct properties — route_policy / 201303311233 / 3

- [all_nodes](data-sources--bgp--reference--group-001.md#canonical-1220302232032111-3231033031212122-2313203103110102-2021011101033322-2000233220133301-3223011201222103-3211020013233022-0232313021301111): complete subsection reference.

- [inbound](data-sources--bgp--reference--group-001.md#canonical-1232333103020021-1131222111101323-0002120202331110-0320123110230203-0302321323200120-2212113320302320-3213022323211110-0010220110221213): complete subsection reference.

- [node_name](data-sources--bgp--reference--group-001.md#canonical-3130312021232102-1102212310221310-0111003332211020-0112323103100110-0021103221220320-2202321201013013-1333301200000213-1303311003021121): complete subsection reference.

- [object_refs](data-sources--bgp--reference--group-001.md#canonical-1002032323113111-1211032132333310-1312111000010132-0321312013001231-2123000222012221-3002031100311132-2103312112331032-2010201313013100): complete subsection reference.

- [outbound](data-sources--bgp--reference--group-001.md#canonical-2332211331120011-1303021331222323-3132200032031301-0222302012100303-2113103131110121-0313230121332221-1023133020001100-1321221221331102): complete subsection reference.

<a id="canonical-3320310213310032-0102231130030222-1323031301310110-0110102000012232-0311321313323003-2300300010321110-2202102112123210-0021233111220003"></a>

## Next pages — route_policy / 201303311233 / 4

- [peers.routing_policies.route_policy.all_nodes](data-sources--bgp--reference--group-001.md#canonical-1220302232032111-3231033031212122-2313203103110102-2021011101033322-2000233220133301-3223011201222103-3211020013233022-0232313021301111)
- [peers.routing_policies.route_policy.inbound](data-sources--bgp--reference--group-001.md#canonical-1232333103020021-1131222111101323-0002120202331110-0320123110230203-0302321323200120-2212113320302320-3213022323211110-0010220110221213)
- [peers.routing_policies.route_policy.node_name](data-sources--bgp--reference--group-001.md#canonical-3130312021232102-1102212310221310-0111003332211020-0112323103100110-0021103221220320-2202321201013013-1333301200000213-1303311003021121)
- [peers.routing_policies.route_policy.object_refs](data-sources--bgp--reference--group-001.md#canonical-1002032323113111-1211032132333310-1312111000010132-0321312013001231-2123000222012221-3002031100311132-2103312112331032-2010201313013100)
- [peers.routing_policies.route_policy.outbound](data-sources--bgp--reference--group-001.md#canonical-2332211331120011-1303021331222323-3132200032031301-0222302012100303-2113103131110121-0313230121332221-1023133020001100-1321221221331102)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1220302232032111-3231033031212122-2313203103110102-2021011101033322-2000233220133301-3223011201222103-3211020013233022-0232313021301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300311030303110-3212311112300320-1301002332300230-3130033112323000-3333010120310300-3132330030102032-2332023231011332-0120302310313111"></a>

## peers.routing_policies.route_policy.all_nodes — all_nodes / 101213200112 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.all_nodes

<a id="canonical-1311112220200302-0101132330010231-3302020331130011-3000211222010120-3130222112113010-0312323131130201-3033310131333121-0020332300013131"></a>

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

<a id="canonical-3202310333103333-3121010302121302-3323002300022301-3222101231113110-3310302110122311-3211333000231010-3112203201013231-3131333300110302"></a>

## Direct properties — all_nodes / 101213200112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330230120111021-3232001022302231-2221033203110131-0000021103231002-2033021110230101-1221333221331100-2102123113032101-2302123021030020"></a>

## Next pages — all_nodes / 101213200112 / 4

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1232333103020021-1131222111101323-0002120202331110-0320123110230203-0302321323200120-2212113320302320-3213022323211110-0010220110221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011300100231311-3212233333113110-2230132020111302-0100320020323320-1133031313213001-2310302012203112-1230111200011323-0000221330322201"></a>

## peers.routing_policies.route_policy.inbound — inbound / 311231002001 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.inbound

<a id="canonical-1323200103330211-1213010211022210-3110123201112132-2133111300222332-3113002310011321-2121013101110202-1021221233122312-0331223331023231"></a>

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

<a id="canonical-1233131320020201-3213002110111030-1101120310023322-3031110130130303-0200200212222323-0021011112313122-3313132103203122-1331113123102132"></a>

## Direct properties — inbound / 311231002001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321311220003123-3233232122231330-1032030223011032-0202313330321131-0121230223200230-2022333311213310-0300330312220312-0201310312032300"></a>

## Next pages — inbound / 311231002001 / 4

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-3130312021232102-1102212310221310-0111003332211020-0112323103100110-0021103221220320-2202321201013013-1333301200000213-1303311003021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233010232031010-0002031103213122-2011203102013311-1002332212300030-2020012300023300-1201112311221000-2220030202011033-1133030102223320"></a>

## peers.routing_policies.route_policy.node_name — node_name / 320023203203 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.node_name

<a id="canonical-2201013133111121-1020201010313021-3033132103120103-1021220020230210-2232303321111030-1323301310011212-0123231313213322-2001323303210311"></a>

Type: `"single"`. Computed.

List of nodes on which BGP routing policy has to be applied.

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

<a id="canonical-3102220023013212-2222200133210023-3303100303103102-3313100030322322-3101211230000021-1002103302223030-1330331212130310-3010320033012101"></a>

## Direct properties — node_name / 320023203203 / 3

<a id="canonical-2103021021302002-2111000313103212-2222233313312333-2020222330220001-2213021332112033-0022321033022200-1202013130103131-2030333321302010"></a>

<a id="canonical-3031231220321122-0010131002012310-3323123113301003-2213010120200023-1313001131233010-3100322212310010-0000033101231203-2210221030320311"></a>

## node property — node_name / 320023203203 / 4

Type: `["list", "string"]`. Computed.

Select BGP Session on which policy will be applied.

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

<a id="canonical-3013210202013100-1112221303130230-2302332031322223-1033133302003020-1121323023132321-1111023320030113-3311300201001233-1301310113201033"></a>

## Next pages — node_name / 320023203203 / 5

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1002032323113111-1211032132333310-1312111000010132-0321312013001231-2123000222012221-3002031100311132-2103312112331032-2010201313013100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311002023320322-0130232001213113-3020303303012232-1102232101322302-1113023210200021-3311121320220320-0012031320213311-3332122010003330"></a>

## peers.routing_policies.route_policy.object_refs — object_refs / 210023312202 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.object_refs

<a id="canonical-2111012301313303-0112231230100100-0223033001211232-0022123311202023-0332330322221123-2033131122113120-0000010113121133-3222123313020332"></a>

Type: `"list"`. Computed.

BGP routing policy. Select route policy to apply.

Upstream description:

Select route policy to apply.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3303331303311200-0002010233222013-3311123121021110-0000102100200123-0032102331111333-1230100112123233-1112221010112011-1201323301032233"></a>

## Direct properties — object_refs / 210023312202 / 3

<a id="canonical-0202332003120331-1113023102311312-0103212303121131-1012310313320202-1011132023012220-2120003210122002-1322102100121311-3022011022103120"></a>

<a id="canonical-1032113013320301-1312301020232222-1233133012201112-0131322233112011-3211102112111030-0000320032203303-2020221322202033-2201130130332302"></a>

## kind property — object_refs / 210023312202 / 4

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

<a id="canonical-3301230133111313-3201102121033202-3201233321320301-1311310120112023-3131023033013331-2330021110003022-2101230112133220-0101331233032110"></a>

<a id="canonical-3012202102322210-0113121212321020-1022232033133021-1321323020320020-3202213022103213-2302332212211322-0122220031031133-1322211102300022"></a>

## name property — object_refs / 210023312202 / 5

Type: `"string"`. Computed.

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

<a id="canonical-0212022032123033-1233310233300301-0120112213211032-3120221310033120-1310122012202333-2001312001101202-2332121020002011-3032222023000031"></a>

<a id="canonical-3032301302310101-2032310003010320-0311221213210012-1101322223212111-2221013213310313-0122202302311211-2212201030012031-2221113210332032"></a>

## namespace property — object_refs / 210023312202 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0020033130233233-0221110123131301-1211322223030033-1131031323131331-3202131313122310-1010031010110300-3100111301111113-1101122131302110"></a>

<a id="canonical-2302030102023123-2023133320221332-0230330331303213-2330212111123100-2001101322232301-0302110223323111-2021100020230112-0310221123001313"></a>

## tenant property — object_refs / 210023312202 / 7

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

<a id="canonical-1203011331312213-2200211033202231-0310103203320222-1000133001210233-3220122332331200-0122123310211232-3020132321330032-1230331030312102"></a>

<a id="canonical-3123232130233222-1312222201121120-3122102013231211-1031100221230132-2023030123223310-3222121201330101-0123200333330002-2302302333310133"></a>

## uid property — object_refs / 210023312202 / 8

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

<a id="canonical-2112003220303320-1212303133331011-1330210333130012-3210320032321011-1332200023022300-1020103303221200-3020220223103002-1113021210201323"></a>

## Next pages — object_refs / 210023312202 / 9

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2332211331120011-1303021331222323-3132200032031301-0222302012100303-2113103131110121-0313230121332221-1023133020001100-1321221221331102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213213010100001-1011212302230032-0003012001302300-1222233122120230-0021021120032021-3001221032133032-2233102012332230-1221331302031022"></a>

## peers.routing_policies.route_policy.outbound — outbound / 310132312010 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [peers](data-sources--bgp--reference--group-001.md#canonical-2101012020132223-0200301120200112-1021132231322112-2320020022131312-1220110313323122-2033302131332110-2211111110210022-1201330210213203)
- [peers.routing_policies](data-sources--bgp--reference--group-001.md#canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321)
- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- peers.routing_policies.route_policy.outbound

<a id="canonical-2230131211311312-3310032303030110-2133222212333031-3322220311230132-0233201213210322-1213313033331032-2313100201003102-3121113333001002"></a>

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

<a id="canonical-2121302213103130-2222102323333133-3220213101221300-1213102000132120-3013332101220003-1230003133223300-2011110232112033-1131012031233013"></a>

## Direct properties — outbound / 310132312010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020010301100223-3330133122233132-3202301103120220-2021113132221122-0333131100032333-0023223101311121-0133322122212101-2322231300323211"></a>

## Next pages — outbound / 310132312010 / 4

- [peers.routing_policies.route_policy](data-sources--bgp--reference--group-001.md#canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003312111232110-2102333200023130-1233321230203212-1122201000200233-0011130331112132-1132221121310301-3203312220313032-3303310132201032"></a>

## where — where / 230121010120 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- where

<a id="canonical-1032200123222212-3232232020300023-0212233033002010-1202033301101323-3132330001120313-2323310330203233-3001220010112023-3031110023033022"></a>

Type: `"single"`. Computed.

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Upstream description:

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-2103220002123313-2332323121122331-2231201320132101-2211203103333212-1001023302321111-0203023231320301-2000121330030311-1020210323103021"></a>

## Direct properties — where / 230121010120 / 3

- [site](data-sources--bgp--reference--group-001.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231): complete subsection reference.

- [virtual_site](data-sources--bgp--reference--group-001.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110): complete subsection reference.

<a id="canonical-0013301123103001-3200031103233133-0303220230203022-1013121021311223-3020301223302111-2110233003320030-1303032123133110-2333020132313200"></a>

## Next pages — where / 230121010120 / 4

- [where.site](data-sources--bgp--reference--group-001.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202100122002013-3230210302012031-3101320223333121-2313211230221101-2312221200012102-2221103321123300-0223132202230132-3023203000010203"></a>

## where.site — site / 211212232001 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- where.site

<a id="canonical-0130100031103010-0103100200123001-1333221132210021-1013113302101321-2111022303100303-1131223001303132-3300133301121000-2001012100223110"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-0020202212333003-2322101223210010-3122101333232112-3231203120301230-1102122313021122-3321121102223031-1300130023111303-0002313011130213"></a>

## Direct properties — site / 211212232001 / 3

- [disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-3122023133003302-3023202213330003-0231200001213223-2231221223020101-3212031112123311-2313120302110201-2230313302210330-1012032020032320): complete subsection reference.

- [enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-1111201203213021-3301110111233310-2200022323201221-0102112120331310-0302300200123303-1210001110022232-1023002002230103-2200301321220133): complete subsection reference.

<a id="canonical-0013201030023013-3013120320223303-1303231311013213-0212223132021321-0130222020213212-0323010112233013-0102233123012011-0010313023313302"></a>

<a id="canonical-0012202110131301-0323313103112230-3011332011230110-0231032123331322-1102001332023021-3112032333120001-3233012323101311-3210020230332110"></a>

## network_type property — site / 211212232001 / 4

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

- [ref](data-sources--bgp--reference--group-001.md#canonical-2032131210203002-0232030220020231-0311222200212231-2112302020010131-2132010132112002-3321332212213133-1313032231031213-0120311333001101): complete subsection reference.

<a id="canonical-3220122012110313-1311313113332013-0010021110033121-1331230212311321-3233320211223201-0100202132202202-1310131223330021-0333230033221200"></a>

## Next pages — site / 211212232001 / 5

- [where.site.disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-3122023133003302-3023202213330003-0231200001213223-2231221223020101-3212031112123311-2313120302110201-2230313302210330-1012032020032320)
- [where.site.enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-1111201203213021-3301110111233310-2200022323201221-0102112120331310-0302300200123303-1210001110022232-1023002002230103-2200301321220133)
- [where.site.ref](data-sources--bgp--reference--group-001.md#canonical-2032131210203002-0232030220020231-0311222200212231-2112302020010131-2132010132112002-3321332212213133-1313032231031213-0120311333001101)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-3122023133003302-3023202213330003-0231200001213223-2231221223020101-3212031112123311-2313120302110201-2230313302210330-1012032020032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023110132020003-3111223033033223-0211112212301102-2331203133032021-2232310103231020-3133301211221012-3223331123200112-2103323012320133"></a>

## where.site.disable_internet_vip — disable_internet_vip / 113213220322 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.site](data-sources--bgp--reference--group-001.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- where.site.disable_internet_vip

<a id="canonical-2111130212202122-2010122132311213-1030211313221313-1020023231022331-0323221100113300-0100100113132122-0330332130122332-2131031100333222"></a>

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

<a id="canonical-2102121311330111-0331322312232122-0113022302110211-3032130123232131-3002210211133310-3223230023323032-0112010020323330-0233313322120311"></a>

## Direct properties — disable_internet_vip / 113213220322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302132302012330-2122311012100300-1012100321012012-2310212311302211-3133233223321113-3003100330313331-1112310001221032-0031002311023332"></a>

## Next pages — disable_internet_vip / 113213220322 / 4

- [where.site](data-sources--bgp--reference--group-001.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1111201203213021-3301110111233310-2200022323201221-0102112120331310-0302300200123303-1210001110022232-1023002002230103-2200301321220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022202323021013-0133210331120232-1123132233310313-2012301232022013-2111311110003031-1110313211311320-0111021120131203-1330113123102010"></a>

## where.site.enable_internet_vip — enable_internet_vip / 131113311023 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.site](data-sources--bgp--reference--group-001.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- where.site.enable_internet_vip

<a id="canonical-1223123112103332-2113010300103112-0321103331333231-1000022023300032-0032332112003312-0201011322212202-1203223301232121-1223103133003101"></a>

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

<a id="canonical-1233132103303032-1023222022032230-2102331020210023-3302103100022110-0233302323122003-0230110012000133-1223011023330022-0320122133112100"></a>

## Direct properties — enable_internet_vip / 131113311023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033230202010033-1023230233131023-3101100221232011-0120201111230013-2320210220312131-2323333231211110-0011133221112210-2313002123313120"></a>

## Next pages — enable_internet_vip / 131113311023 / 4

- [where.site](data-sources--bgp--reference--group-001.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2032131210203002-0232030220020231-0311222200212231-2112302020010131-2132010132112002-3321332212213133-1313032231031213-0120311333001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031213310200033-3330300300020002-2200201130212013-2020201000230302-1303133123200132-3020111331030231-2020323332122203-3112210132223301"></a>

## where.site.ref — ref / 102323021201 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.site](data-sources--bgp--reference--group-001.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- where.site.ref

<a id="canonical-3213230032233033-1331131121320133-0210121110113120-1132322223032022-2222310300010111-2132033131221311-1233023021213132-3010032222031300"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0222131233102002-3230233210033313-0121102033132122-3333311103230303-3333132212011100-3000230223002331-1122323330111001-2032131002130000"></a>

## Direct properties — ref / 102323021201 / 3

<a id="canonical-3310321001030322-0113113110022211-0233211231101120-2013221111322203-1103303200201230-3001230212302103-3221203101020231-1122000332003311"></a>

<a id="canonical-1110213210030313-0033213230211033-0223311103003232-0022010200122330-0222003031111310-3122321310122001-0233212330013220-3201230322000103"></a>

## kind property — ref / 102323021201 / 4

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

<a id="canonical-1211120301233121-0322122020030002-3210301223221000-3210002033200200-0333220121303112-3023120310033320-0111210222003220-2301113123000133"></a>

<a id="canonical-2110223013310201-0112211013101133-0020112121013213-3112031303323102-1203003121200301-3101233013001010-0022221201130122-0300221213310111"></a>

## name property — ref / 102323021201 / 5

Type: `"string"`. Computed.

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

<a id="canonical-1110310020320133-1023120201223012-0210002202303031-3030032022203021-2111301001031122-1331232133133333-2233313332112313-3200010302320312"></a>

<a id="canonical-1211310300320100-3002013330032112-0213203333333203-0002110130332122-1023301111113111-2113321233020030-2330031031220312-0112001111321220"></a>

## namespace property — ref / 102323021201 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3103132103131100-0323131023000023-1120220120133313-0121032322102310-1103011313330321-2020112132200023-3112133313300132-2111013303320020"></a>

<a id="canonical-3030210033131313-1331331311110001-0203320023133231-3002201213213123-1321030210021323-3020331331020122-0131122130230330-1130302302123123"></a>

## tenant property — ref / 102323021201 / 7

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

<a id="canonical-3133002332002232-0322211230012010-3133023133200323-1103203232222332-3223211110123231-2100223303020201-3212010222011011-2332013232232323"></a>

<a id="canonical-1033021132110233-3223222213212332-3311002213020333-2301232102323123-2010223110220123-2312323300130302-0312021330022001-1000313323212030"></a>

## uid property — ref / 102323021201 / 8

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

<a id="canonical-0322210332201211-3132223212013112-0001220023233012-2213102311220011-2332213133303222-1331033211011320-2111112210002201-1233111113013330"></a>

## Next pages — ref / 102323021201 / 9

- [where.site](data-sources--bgp--reference--group-001.md#canonical-0322322033232122-2303320112011223-0113201323122302-3012203220333320-0232122330220202-3032121230010233-3211132123232001-3300210123011231)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222001202000120-0133233013213013-2322020302122311-3133030130030211-1302321302210213-0003222003201333-2131103032132000-0323032202312223"></a>

## where.virtual_site — virtual_site / 002231321230 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- where.virtual_site

<a id="canonical-3320101110123322-1011312002232332-3112322123131231-0301223333010302-0302133212210222-3033203023211103-1030203211121231-0313310101103321"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-0331310311233110-0312323013221031-1120311030131001-1331211202303112-0323032013320231-1201112032013230-0131303330223013-0120323310110231"></a>

## Direct properties — virtual_site / 002231321230 / 3

- [disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-0003322330321310-0321121132103301-3001300220121020-0133002112133020-1323200033212220-0323331020002032-3202320112232122-0033333320130323): complete subsection reference.

- [enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-1031022120323101-1230331131031223-0010130110120102-3002031333203320-2231001012012212-1011332202111132-0030200133231123-3133131303023213): complete subsection reference.

<a id="canonical-2102131220223300-2101111232022231-1000223032013211-1003113123322001-1131111030113000-1321333310331003-3012221122103331-2103201030023011"></a>

<a id="canonical-1311130332303013-0120133011330320-0132130221212323-3203020123320101-2112313002212312-2301233002021313-2211322120222230-2231011010221212"></a>

## network_type property — virtual_site / 002231321230 / 4

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

- [ref](data-sources--bgp--reference--group-001.md#canonical-2031111210002031-1022301331231023-3001212221103321-2200012000110022-2322013032303131-2033321130332102-1103231300011332-1310313221103001): complete subsection reference.

<a id="canonical-2122232230210011-1211330133122203-3323132032101022-0011322301320020-3233001333033113-0213220123331212-1001111013201332-3201011021223102"></a>

## Next pages — virtual_site / 002231321230 / 5

- [where.virtual_site.disable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-0003322330321310-0321121132103301-3001300220121020-0133002112133020-1323200033212220-0323331020002032-3202320112232122-0033333320130323)
- [where.virtual_site.enable_internet_vip](data-sources--bgp--reference--group-001.md#canonical-1031022120323101-1230331131031223-0010130110120102-3002031333203320-2231001012012212-1011332202111132-0030200133231123-3133131303023213)
- [where.virtual_site.ref](data-sources--bgp--reference--group-001.md#canonical-2031111210002031-1022301331231023-3001212221103321-2200012000110022-2322013032303131-2033321130332102-1103231300011332-1310313221103001)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-0003322330321310-0321121132103301-3001300220121020-0133002112133020-1323200033212220-0323331020002032-3202320112232122-0033333320130323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133222100222223-1100023203331133-0210233330033013-0330331310301331-1012300023102222-2102321233200323-3012302312010331-2112002022030012"></a>

## where.virtual_site.disable_internet_vip — disable_internet_vip / 001323011322 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- where.virtual_site.disable_internet_vip

<a id="canonical-3122313321211100-0331010002300012-1322002101211302-1000330021230021-1220322221321123-2312203011311023-3020010330202031-1233330030201001"></a>

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

<a id="canonical-0231131222102020-1223230110303031-2001332303232302-3203313213330310-0323322012300020-2312022113021313-1020021100130010-2330211222011220"></a>

## Direct properties — disable_internet_vip / 001323011322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111022121311320-3131020023221202-2102122103232203-2100000333230003-2323321130203223-3031132113121222-1133233130012323-2300210222223011"></a>

## Next pages — disable_internet_vip / 001323011322 / 4

- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-1031022120323101-1230331131031223-0010130110120102-3002031333203320-2231001012012212-1011332202111132-0030200133231123-3133131303023213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312113300111033-3211133233010020-2003020101222222-2031313230021303-0322203101001120-1012023221112012-2031301233313130-3210323013320123"></a>

## where.virtual_site.enable_internet_vip — enable_internet_vip / 223211302110 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- where.virtual_site.enable_internet_vip

<a id="canonical-2211312233303220-2131111113212001-2312100331023211-0321120203022312-1231231132302000-3312013200222010-3320302311101100-0120020123021212"></a>

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

<a id="canonical-0011113120320001-3300221032033100-0300133002201323-1010031233020122-2013203320011103-3122130010321132-0223232111001302-2012232021101220"></a>

## Direct properties — enable_internet_vip / 223211302110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213112310303113-0102321033102332-2331113302203301-2130012310012222-2121031101023022-2231032123233023-1301310210003023-2211331231110322"></a>

## Next pages — enable_internet_vip / 223211302110 / 4

- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)

<a id="canonical-2031111210002031-1022301331231023-3001212221103321-2200012000110022-2322013032303131-2033321130332102-1103231300011332-1310313221103001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033111333220332-0122221332020201-1332332122123101-2033230002023111-3121303221123312-0231221321301220-2221023201223212-0223222311102022"></a>

## where.virtual_site.ref — ref / 320331330001 / 2

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Property reference](data-sources--bgp--reference--group-001.md#canonical-2133002300232122-2112123312313302-3002130002300113-2210220333320001-3000003311122210-2013212232010312-0200311202002111-2023102231001120)
- [where](data-sources--bgp--reference--group-001.md#canonical-0300332133213020-3212230310013301-3212003102003220-2331123121213021-0200021121002323-3333331230031120-3032210031203331-1030000212002120)
- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- where.virtual_site.ref

<a id="canonical-0110323222102130-2100212211222200-3013310010121220-2001201231032003-2123310330202102-2300100220333032-2012312321132010-3023331221311012"></a>

Type: `"list"`. Computed.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3312012213011031-1002322300311201-3333311033322312-2233003020212112-1313200232002322-1202031033320021-0132311211213011-2301100312232030"></a>

## Direct properties — ref / 320331330001 / 3

<a id="canonical-3111030300032222-2332210030021322-1021310012201320-0231232003020011-3012031123130321-0321111103213302-3000120130311313-3223000022132111"></a>

<a id="canonical-3332320000313201-1331213333123211-1022233133222220-1000323012130112-3212231110020331-1233232012203330-2110021323010021-3101331330021223"></a>

## kind property — ref / 320331330001 / 4

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

<a id="canonical-1313001302322031-3322100032323301-0031030220220230-0202333233021210-1323112312303223-0231013132022022-2302100113202312-2312000320210212"></a>

<a id="canonical-0321021033313221-3031223102231013-3010333013013312-2230211013030001-3020302021022100-2231133101203213-0320121300330303-1000031000330012"></a>

## name property — ref / 320331330001 / 5

Type: `"string"`. Computed.

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

<a id="canonical-1320121200322031-0003231231001031-3231032202301230-3331201322111231-3222121133202033-0003213120101032-2012322001232001-3201332002212012"></a>

<a id="canonical-1030023323132123-0020032333132231-0332022333002111-0122311313100301-2321312123003002-3111031030012302-0100020023000121-2130133330201110"></a>

## namespace property — ref / 320331330001 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0033302331230021-1010030233100303-3032323132303033-3320233303000323-0201010303213210-3331330230211322-2312332013030101-0012103333110101"></a>

<a id="canonical-2032322322022233-0011013020013010-3133033222201131-0102023032200030-0302330213303223-3231210223331311-3222203001011212-3120322212232300"></a>

## tenant property — ref / 320331330001 / 7

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

<a id="canonical-0021313313213030-2101211032000022-1301103021312311-0011201112303311-0003301213021213-0311102333131121-0013301021313322-3210231111111013"></a>

<a id="canonical-2111110112212211-2331231101203030-2123310111123203-3201021330221032-1132123023110310-3001230033321132-0030033323011102-3111031112022121"></a>

## uid property — ref / 320331330001 / 8

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

<a id="canonical-0003313033312203-2200122213333303-3011111112333001-2201032010201210-1321032102210021-0031331031120031-3300213200332232-0330232131211320"></a>

## Next pages — ref / 320331330001 / 9

- [where.virtual_site](data-sources--bgp--reference--group-001.md#canonical-0001123113232320-2133112103332320-3121101122131111-3032121101310020-3002331132301230-3231303333120301-1120100023033233-2010330120222110)
- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
