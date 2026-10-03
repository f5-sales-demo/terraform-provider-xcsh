---
page_title: "xcsh_forward_proxy_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy reference."
---

# xcsh_forward_proxy_policy reference

<a id="canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121321011002221-3211211030322322-1021021131023132-0302021203323001-3033332321101211-2202301333001111-0000132330301123-0323103332021102"></a>

## Property reference — Property reference / 302333103020 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- Property reference

<a id="canonical-2232003132103002-0033132202132331-3223121023131222-1031321322210021-0220223203313302-0110121000002222-2230301103013221-3312030223211001"></a>

## Direct properties — Property reference / 302333103020 / 3

- [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2013011230133310-0110112322303323-1102232112233312-0332123113232113-1131213101211320-3211310122131000-0030202322111201-0103202332111313): complete subsection reference.

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301): complete subsection reference.

<a id="canonical-3332302313200021-2310100332200010-0003311112223112-3032131130222311-0210321011231033-3131222322320101-3220221302233301-1323132203301203"></a>

<a id="canonical-0232103100302002-2100133122101301-2031300011200331-3301022202302102-0131311132301322-0301001322130321-2233133013102100-2000320202032020"></a>

## annotations property — Property reference / 302333103020 / 4

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

- [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0301302122312100-0303121321221121-2233002013020312-1211203203103312-1033012321032120-1121133311231131-1221332120300202-1103321303300103): complete subsection reference.

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032): complete subsection reference.

<a id="canonical-2010032211002230-0010030311032231-2300320300000130-0113012111022332-1030103303303031-0120303012110302-3113302110111031-0022322132201223"></a>

<a id="canonical-1311001322312313-1201021101133311-2032110232022323-2222021112300020-3103102132113112-2030201102030233-2023213010332020-3221332103110303"></a>

## description property — Property reference / 302333103020 / 5

Type: `"string"`. Computed.

Description of the ForwardProxyPolicy.

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

- [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0301032221101120-3103200113001310-3223232331220200-1332133303133312-1323311130131310-1001333230111322-0230221210331302-3200332121101102): complete subsection reference.

<a id="canonical-1312130103301321-2320312332000213-3303230110120213-3222232232212023-3121301203230102-1002322311330003-0203122233012303-1321210131201003"></a>

<a id="canonical-2101030202210210-3123031012122303-3321332031300213-0303000231011002-1001322230012323-1331022032031110-1303333101120203-1020233111131323"></a>

## ID property — Property reference / 302333103020 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1012001101201222-1012022131203020-3203200201202223-2103323221221102-1022123030101113-2031111001231000-2330233231220220-2220021330323000"></a>

<a id="canonical-3301132300131120-0221030201031211-2123031221023020-3203121132300313-0132033330001321-1230211201100113-0130020111011031-0202010003032032"></a>

## labels property — Property reference / 302333103020 / 7

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

<a id="canonical-2001101330212233-0303013320213032-2112220012212021-1001211111022003-2001021311030210-0131113130000211-2311232301201112-2111111332203323"></a>

<a id="canonical-3203022102320030-1312211212202233-3231112303320323-0121120112322030-1002102232221220-3230020013213203-2323113303233002-1202201111022300"></a>

## name property — Property reference / 302333103020 / 8

Type: `"string"`. Required.

Name of the ForwardProxyPolicy.

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

<a id="canonical-2023100322011231-0020021321002032-0301210102032303-0323320022020223-3302000132022232-3100113330130333-3110103303113211-1102312302013000"></a>

<a id="canonical-0011200221032021-2212022131333200-2111211231002313-0331221131202013-3330213003102101-0211303003321203-3013210321101203-3322322233313211"></a>

## namespace property — Property reference / 302333103020 / 9

Type: `"string"`. Required.

Namespace where the ForwardProxyPolicy exists.

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

- [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3003120302033030-0022133033303133-1021310202213331-0113203001333310-0100202002013200-2202322230231323-0010302133310002-3313303112221211): complete subsection reference.

- [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1032301033033313-1101300232301332-0110322320310030-1032321212200203-0212320010031030-1100000123210022-0031113133102112-0301111301311103): complete subsection reference.

- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232): complete subsection reference.

<a id="canonical-2020022121322030-2212032101312331-2332212211333013-3322221203013313-3033313323112200-3201231221022110-0011022133220222-3221000220303010"></a>

## All schema paths — Property reference / 302333103020 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2220010203033123-1110321130230310-3023220022020123-2201003102133013-3330100232323231-2031021200012211-0323121203331011-2113031103331212) |
| `allow_list` | [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1131021322232311-1013030302000130-0100233010200032-3222300200230310-2310111103303033-3003220000012102-3210000312000223-1112313030320123) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2022122121202213-3223323003033031-1223233302001210-0311322121231330-2220322313000010-2223021012321233-2100310212011003-3130211221101103) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3030100300212120-1000233011133121-0032302022330313-0130000221313222-2010212200131003-2220032023333303-2110033223213232-0320320131212330) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2010310213310011-0112301002210230-0130321233033311-0330313012221012-2202211302020012-1132032311212202-1001301030022200-0122110203220032) |
| `allow_list.dest_list` | [allow_list.dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3213201020300211-2123023211301310-1121012133222102-1022101230133000-3231102212113031-1023130302323022-3103202321012211-2220021030033013) |
| `allow_list.dest_list.ipv6_prefixes` | [allow_list.dest_list.ipv6_prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3123303120212202-0312212231030033-0301003230313031-2113012211102221-0311301333311112-0202311221012310-1031211031221223-2031132301331211) |
| `allow_list.dest_list.port_ranges` | [allow_list.dest_list.port_ranges](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1220300203003101-3120310303023113-0203213101230303-1033021212313121-0233131123221022-2210033332131212-2211120222322233-1220020120013322) |
| `allow_list.dest_list.prefixes` | [allow_list.dest_list.prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3231322001003123-1212111213133100-2100332113223313-3033233303230033-2232313200003033-3101302202121311-0103012310310312-2201020101312310) |
| `allow_list.http_list` | [allow_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3000203003021202-2211013203331301-1330110033302100-2112133303112221-0013302320330110-1210220123033031-2230212100031122-2101201012303011) |
| `allow_list.http_list.any_path` | [allow_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1313032120321232-3120020232300320-0312031332102322-1013232212103311-0102003333231130-0233303131210010-3210032231010320-3210100210132102) |
| `allow_list.http_list.exact_value` | [allow_list.http_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3322110132312230-0232233022122131-2013223103023310-2230101001310130-3020132122321212-1030232023032221-1111300000232322-2122213213313331) |
| `allow_list.http_list.path_exact_value` | [allow_list.http_list.path_exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1132131320100203-2202201021331301-1020320021023113-1213021212331100-2021330302131002-1212313201031031-2001001113123310-1210020223203030) |
| `allow_list.http_list.path_prefix_value` | [allow_list.http_list.path_prefix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1022231130033231-0110222032301210-0211313203320001-3213221111002313-1023301230202103-1330333113013010-3111113211111320-1002232023203333) |
| `allow_list.http_list.path_regex_value` | [allow_list.http_list.path_regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1031322021022022-1232132013220131-0323010222123232-1333123302023330-3111230213203320-3222003222101012-3331010133003101-3220303023003302) |
| `allow_list.http_list.regex_value` | [allow_list.http_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3211311133000030-0301331303232311-1221022202020102-1001311113131012-2022210231101013-0312100313203231-1202231221211112-1321103113331313) |
| `allow_list.http_list.suffix_value` | [allow_list.http_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2322123031203333-2330310021200300-1222031301002132-2112323203210200-1031203300002012-0012233001003201-0010101130021211-1122301131000031) |
| `allow_list.tls_list` | [allow_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1310033300013021-0303211310233301-1333201112312230-1020111020333000-3032230113022100-3221231011302033-1232313210111013-1311210021321222) |
| `allow_list.tls_list.exact_value` | [allow_list.tls_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2303011022320211-1133031222301231-2311313030113101-1201030220130001-0312032211012022-1101232000212132-2030310221010331-0132101022113301) |
| `allow_list.tls_list.regex_value` | [allow_list.tls_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3032002112232013-0122321031213001-0013331012103003-2320111132232221-2010332121211322-0330120213303122-0213123213132231-1312002130310123) |
| `allow_list.tls_list.suffix_value` | [allow_list.tls_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3322022102122330-1231001021312311-1303011111102001-1312031210000333-3231222021213232-0033013221320131-3122313320101231-3121212130221213) |
| `annotations` | [annotations](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3332302313200021-2310100332200010-0003311112223112-3032131130222311-0210321011231033-3131222322320101-3220221302233301-1323132203301203) |
| `any_proxy` | [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3123313021211132-2130110232333032-0303111023110033-3112321221232113-3211111022110111-2013331232011111-0032333110001313-3002032103013321) |
| `deny_list` | [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1001323031020030-0323022102120112-0100301120113130-2133213212210110-0212222011312233-0032322201311210-2203332013010002-1320313230312301) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3202123133231120-3101322131023113-1123322322131321-0130100100233133-3021322023330010-0310121002321033-3123110211202120-0110003111203300) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2112111020202300-2222202333232012-1022323012033111-0030230131103010-2112211121120322-2013321023210300-3200122001111102-0300303301030332) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3313011032001133-0311320311311111-2131321213230021-3302013201121103-1131001131231221-1332231010202103-2031302211130121-3213313202020113) |
| `deny_list.dest_list` | [deny_list.dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2331003033203130-1302130022131313-3312031100210302-2233131133332231-3200232200213002-2321111202110330-3012300203213133-1003302211003312) |
| `deny_list.dest_list.ipv6_prefixes` | [deny_list.dest_list.ipv6_prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2330010031030121-3331021223202330-1200330222120201-3001310332111220-3123131203212112-0130231312102002-2111020202222123-1002331302201201) |
| `deny_list.dest_list.port_ranges` | [deny_list.dest_list.port_ranges](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1120312030113121-3303211131320001-0012130201122201-3200223103312300-2231100020001321-1002231003313320-2122301100223102-0200301020311300) |
| `deny_list.dest_list.prefixes` | [deny_list.dest_list.prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0101323223130201-2203121332132113-3320232221023332-1231023302332320-1010203200333311-3320233031300010-3010231211222222-1212030010203031) |
| `deny_list.http_list` | [deny_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3332030102031203-0322222033321303-0022302000331031-3003201030233013-0131122203222102-3222210031012032-0211203203113000-3233123021230211) |
| `deny_list.http_list.any_path` | [deny_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3230132322031011-2133220301300220-2012122102202221-3200111331031203-2211310032233301-0011203030130120-0203231311212332-3210111123303223) |
| `deny_list.http_list.exact_value` | [deny_list.http_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0012230302020123-0330301232311322-3330003032122112-0131201121211311-2201020133100120-1322133021113032-3111033213133313-0312212200101212) |
| `deny_list.http_list.path_exact_value` | [deny_list.http_list.path_exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0211020133310020-0131333020120021-3032322211001223-2323233322122303-0022202010231001-3000202130022103-2223222001032020-1321323133011231) |
| `deny_list.http_list.path_prefix_value` | [deny_list.http_list.path_prefix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3020331321013002-0012130110023033-1212210101030202-2230021133020021-1332313220102032-0133021013203303-3101002230130121-2013032333302233) |
| `deny_list.http_list.path_regex_value` | [deny_list.http_list.path_regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1103223221022112-3012223233300331-0021313223330230-0200331110223000-2302300233123323-2213300031101003-0002211330123233-3002110010302223) |
| `deny_list.http_list.regex_value` | [deny_list.http_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1012112322021113-0023001033230232-0131133121310233-2131220002021111-1310222030121131-3033122200313031-0301121312103213-2301000300013103) |
| `deny_list.http_list.suffix_value` | [deny_list.http_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2010322031123131-0231020202311312-2303112332233133-2321101203211130-2210320213221311-1111133023102233-2221331333133333-0330201112310112) |
| `deny_list.tls_list` | [deny_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1000130231032301-3133031001131333-2331322231330023-3331131322033322-2231110111010132-1021221122331200-1100230331222101-3303103021121013) |
| `deny_list.tls_list.exact_value` | [deny_list.tls_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2200321232000100-1203230131131332-1301301230020232-1003231001131313-1132120113331121-2030011033213231-2202133130131212-2203110230100321) |
| `deny_list.tls_list.regex_value` | [deny_list.tls_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1101333332133120-0311301001333310-0332201220121120-0332203313022231-3203330013022211-3203323233132123-3020112203101120-0100020010231102) |
| `deny_list.tls_list.suffix_value` | [deny_list.tls_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3313012111113221-2110330230202032-1130301331020233-3330121010122013-3312120100002003-3001311131321302-1101323222222122-2202210223310123) |
| `description` | [description](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2010032211002230-0010030311032231-2300320300000130-0113012111022332-1030103303303031-0120303012110302-3113302110111031-0022322132201223) |
| `drp_http_connect` | [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1232303201032011-2232231121322221-2301300110332023-1032001113103110-3110111022311231-0232330102320221-3220013322023130-3203010220300301) |
| `id` | [ID](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1312130103301321-2320312332000213-3303230110120213-3222232232212023-3121301203230102-1002322311330003-0203122233012303-1321210131201003) |
| `labels` | [labels](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1012001101201222-1012022131203020-3203200201202223-2103323221221102-1022123030101113-2031111001231000-2330233231220220-2220021330323000) |
| `name` | [name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2001101330212233-0303013320213032-2112220012212021-1001211111022003-2001021311030210-0131113130000211-2311232301201112-2111111332203323) |
| `namespace` | [namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2023100322011231-0020021321002032-0301210102032303-0323320022020223-3302000132022232-3100113330130333-3110103303113211-1102312302013000) |
| `network_connector` | [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3121013223110112-3301132223320301-3111002130110031-3133111102233212-3002211032112332-0123211112320311-1111302121022100-3210222210003131) |
| `network_connector.name` | [network_connector.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2003113321102020-3010121213103312-3031033201303331-0313231000121112-0232203233330233-2032111202212003-3031322300210301-1203021101021222) |
| `network_connector.namespace` | [network_connector.namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3001003103103132-0100122102320212-1322201123202103-0102123201312231-3201023213121303-0323322132203322-2003132202010323-3222021230313000) |
| `network_connector.tenant` | [network_connector.tenant](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2300322032101212-2102011102022232-1123320022202301-0232110132313232-2102300220002003-2112303223102033-2111222121323110-2233110200123221) |
| `proxy_label_selector` | [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3320221102222111-2312332233123123-0022022013122310-0231213333221113-3301223112111020-0120030021112201-1210012120013200-1210110111112213) |
| `proxy_label_selector.expressions` | [proxy_label_selector.expressions](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1300023203010112-2003122212123121-0320231300323302-0312320001120013-2122100132003011-2212000122231010-3133130330220312-0311103002312001) |
| `rule_list` | [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3300003111010003-1022200013220321-2120312030000021-0322110310133010-2300202200330301-2102222012123212-3322012103010130-2332110210313311) |
| `rule_list.rules` | [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3100302032212022-2021123222100320-3213331000002321-2003220103311020-0101020103332002-1313332302120213-0321120011110223-2110233133303132) |
| `rule_list.rules.action` | [rule_list.rules.action](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1131000023102020-1231010222000020-3001331200220123-0203303331330312-1031123113111032-1000301210102321-2102313200023203-1200130222231333) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3202113122323120-0131313120111220-3331020322313230-3012320022131230-3333301200321111-2210213110123333-1330333230301321-2002120302021203) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0232322210231213-1331210222033021-1312001113030130-2111300303102332-2121333301322131-2302331100020102-1210101111321300-3032213222012221) |
| `rule_list.rules.dst_asn_list` | [rule_list.rules.dst_asn_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0303030201113121-1103123211313320-2021122332033320-3110030310233020-3022311312010321-0002213022211101-2320012313111133-0132122022111120) |
| `rule_list.rules.dst_asn_list.as_numbers` | [rule_list.rules.dst_asn_list.as_numbers](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2003230122023011-2121213210123031-1113031023122221-2023230133300111-3123122013003332-2131330132000112-0031101033020121-1131233311303333) |
| `rule_list.rules.dst_asn_set` | [rule_list.rules.dst_asn_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3111311212203313-3311102310322000-2032122321103130-2302311032021331-1101223222020021-1221122022032002-0310111021133330-3110122110213322) |
| `rule_list.rules.dst_asn_set.name` | [rule_list.rules.dst_asn_set.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1101221330120033-1333020323112103-1323002010222002-2102032223122320-1132110312103222-2200012123201213-3112232111211030-1210132333120311) |
| `rule_list.rules.dst_asn_set.namespace` | [rule_list.rules.dst_asn_set.namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2120023233120332-2313201203332121-3221303013000302-3313110123232132-3222003033113030-3112012021003122-3003202012203030-2221013220020202) |
| `rule_list.rules.dst_asn_set.tenant` | [rule_list.rules.dst_asn_set.tenant](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1122030031232123-1012003211300232-1033023221232233-3000333032030121-1133000013010303-1323213322122323-3201110332302110-0211000113131201) |
| `rule_list.rules.dst_ip_prefix_set` | [rule_list.rules.dst_ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1103131002311232-1320220001232303-1312231231010331-3322012200212232-3223223122313132-2032031102120010-3322321001210302-2101003302223312) |
| `rule_list.rules.dst_ip_prefix_set.name` | [rule_list.rules.dst_ip_prefix_set.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2322102303003010-2022232312113301-1331313022032220-1131321010233223-1300303320133011-3231313011201333-0020111332302310-1132121021332323) |
| `rule_list.rules.dst_ip_prefix_set.namespace` | [rule_list.rules.dst_ip_prefix_set.namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1330211321130102-3133131322230002-3231200232110233-1010310030122001-1312311213001131-2211112210021303-2333021103132022-0020101222301320) |
| `rule_list.rules.dst_ip_prefix_set.tenant` | [rule_list.rules.dst_ip_prefix_set.tenant](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2232311300100111-3031232132100333-1033202333201013-3012320212003300-0311313133332212-0202323320231222-3120131210002123-0221113222323012) |
| `rule_list.rules.dst_label_selector` | [rule_list.rules.dst_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201201320303023-3100103321201122-1012310023331330-2231213211022200-1110312301211321-1002000012322203-1121022022022300-2000320212022123) |
| `rule_list.rules.dst_label_selector.expressions` | [rule_list.rules.dst_label_selector.expressions](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2312310323122132-1130022111033331-1322302131230201-2021323031003133-3010223222000120-1001330130221100-2210212100300220-0132211111233031) |
| `rule_list.rules.dst_prefix_list` | [rule_list.rules.dst_prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0223030022201003-2200200200212311-1002332201330230-0333233231033333-1022231312203330-3210013000132011-1110132233323131-2212022301232230) |
| `rule_list.rules.dst_prefix_list.prefixes` | [rule_list.rules.dst_prefix_list.prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3011333011112000-0301131132120001-2232312232221323-0233023333111321-3203230023021320-1300111201302213-1122210113213333-0012322011223032) |
| `rule_list.rules.http_list` | [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0313322022003312-1011021132002233-2233113003012002-3000231031132201-2330222203212330-2232331110203021-0132310310120212-2113011020012122) |
| `rule_list.rules.http_list.http_list` | [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2332313220202000-3312100330020212-3112003002100213-3103003112132222-3313133221001310-3112333121203012-1230330210012213-1330013121333122) |
| `rule_list.rules.http_list.http_list.any_path` | [rule_list.rules.http_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1000221323210102-2001232000102021-2022332131133200-3300211133331011-3312130111202030-1023022200210103-0011112210031123-3122032033202131) |
| `rule_list.rules.http_list.http_list.exact_value` | [rule_list.rules.http_list.http_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3111112333130221-1303131331030032-1201221220120131-0313200330220210-2122220120012211-1012121203203220-2133211012103330-3222132032203203) |
| `rule_list.rules.http_list.http_list.path_exact_value` | [rule_list.rules.http_list.http_list.path_exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3202020112112301-0130301032331013-3223031020012102-2221222111000222-3131102031122003-3110003202330003-0232201311303011-2001222202320000) |
| `rule_list.rules.http_list.http_list.path_prefix_value` | [rule_list.rules.http_list.http_list.path_prefix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0203302233233303-3101333323220301-3122101122123001-2131332212312103-2110111331131123-0230030030203131-1010302022300202-0312302103121033) |
| `rule_list.rules.http_list.http_list.path_regex_value` | [rule_list.rules.http_list.http_list.path_regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2101002331233101-2121123312030213-1212133022102300-1003221121320210-1313231012130201-1122023123030130-2330231333003100-2022021211212023) |
| `rule_list.rules.http_list.http_list.regex_value` | [rule_list.rules.http_list.http_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2231033013310301-1300010312031130-3103331320013211-0332101033001032-2213312312232303-2301232003320300-0210113011202003-2122311321121102) |
| `rule_list.rules.http_list.http_list.suffix_value` | [rule_list.rules.http_list.http_list.suffix_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2222033202030201-0103312203123121-1130330031223122-3231321110323103-3201330223123021-1322032300300130-0000130001102111-1012212032120010) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2221331210221003-2132123310203232-0200013110031023-0002101033031131-1221132113133023-3022000030320031-3023000332220103-2022031021310013) |
| `rule_list.rules.ip_prefix_set.name` | [rule_list.rules.ip_prefix_set.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1230000211031033-1020001022123233-1012031230103223-2031203301201211-0020233121223212-3221202220022001-1230002323220221-2003330022033332) |
| `rule_list.rules.ip_prefix_set.namespace` | [rule_list.rules.ip_prefix_set.namespace](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3231013322121303-3211023310222012-1212300102133111-0333313213331301-0121132232133202-2103321011331112-3123300233201132-1211030321020113) |
| `rule_list.rules.ip_prefix_set.tenant` | [rule_list.rules.ip_prefix_set.tenant](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0123201102220123-2130032032012111-3233022021022102-1312131023210020-3003132201332311-3012220202301203-1202330010320311-3320333321030031) |
| `rule_list.rules.label_selector` | [rule_list.rules.label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1232121312200123-3023221110001301-1102132213030122-0023311123012201-0231100200223121-2233212231230032-0320102101301123-1033320303203202) |
| `rule_list.rules.label_selector.expressions` | [rule_list.rules.label_selector.expressions](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3330332011201323-1221230021010310-2332112102112030-1123222232210223-2121002033103331-3011232302122230-0013310022221131-1231021130120120) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1210001310130000-3210012311233133-3203210131221230-2122331100231103-1232002322201110-3312010130102220-0001101001020211-0112012302211310) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3022123321130201-0032231300012322-3003302202102003-3300010122103121-0311233211231032-2320102111100012-0032210002131133-3003213030131220) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2200322132332303-2323301312032032-1123321320220013-1031301233201230-2020131130311130-2013200130002311-1111312110323011-2312120023120021) |
| `rule_list.rules.no_http_connect_port` | [rule_list.rules.no_http_connect_port](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1100003212003111-2323300223301233-1031230030301310-3212211233102103-3011002303203300-0112210230132020-3010220233131302-0233121302100032) |
| `rule_list.rules.port_matcher` | [rule_list.rules.port_matcher](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0010111210103020-0201131031233012-3022203303022110-3320030321330122-1300200033301312-0013313002212230-0103213311310311-0302211222100201) |
| `rule_list.rules.port_matcher.invert_matcher` | [rule_list.rules.port_matcher.invert_matcher](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201020021223012-2301010332213221-0102121212210210-2120330110231032-2003233113322032-0121223011011120-1012200323203231-3000132232033233) |
| `rule_list.rules.port_matcher.ports` | [rule_list.rules.port_matcher.ports](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2310013313132102-1102331133322010-0223312113301331-1113002323122333-1222023100020223-1211001012200122-2300023031310131-1030222200333210) |
| `rule_list.rules.prefix_list` | [rule_list.rules.prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0011311212302311-1032132100233321-3030201131123332-0303230101110013-2301332123200111-2303313331112220-2231020112312331-1132213222233333) |
| `rule_list.rules.prefix_list.prefixes` | [rule_list.rules.prefix_list.prefixes](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0003033122013012-1312321111102100-2111321330300023-0111213310330331-1233012132332232-1300222012221331-0200000311031133-3133120133121102) |
| `rule_list.rules.tls_list` | [rule_list.rules.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0310303133133202-3123020233103333-0330303332213232-3102210130001301-0203320323021232-1333013233022222-2023200032210202-2023003003301330) |
| `rule_list.rules.tls_list.tls_list` | [rule_list.rules.tls_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2221221222120112-0230021231220123-3301300121011210-0203100102003010-2231102212311222-2220120122210122-3001211332213332-3022002102200321) |
| `rule_list.rules.tls_list.tls_list.exact_value` | [rule_list.rules.tls_list.tls_list.exact_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1231232302320130-2023322210002011-2320030032332032-1011123111333213-0333310013131311-3210100331302133-3211133103003322-3331331100000321) |
| `rule_list.rules.tls_list.tls_list.regex_value` | [rule_list.rules.tls_list.tls_list.regex_value](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3131121023120210-0013333111312111-2322322202212121-0022122110003230-1101133001020130-1131012220221333-1000012033303111-0100311333222213) |
| `rule_list.rules.tls_list.tls_list.suffix_value` | [rule_list.rules.tls_list.tls_list.suffix_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0323013202223203-1131011033132200-0032201033002102-2011120320133203-1201121313211301-0000013202131003-0302033120113001-3210300002010233) |
| `rule_list.rules.url_category_list` | [rule_list.rules.url_category_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3211012133002021-0031221310323002-1213030330010223-2320300332101200-2021130012032312-2120110323122333-3031003001021132-1121323213231210) |
| `rule_list.rules.url_category_list.url_categories` | [rule_list.rules.url_category_list.url_categories](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2131221100000212-2130132000023302-2131000020200320-1031112200021103-0002233003031312-0131003300330230-0132211220220100-1231300133212100) |

<a id="canonical-3232331133331233-0302212221030133-2120310333332120-2212200032033032-1000321132111232-1001302330120331-1010121133320003-0011201013221013"></a>

## Next pages — Property reference / 302333103020 / 11

- [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2013011230133310-0110112322303323-1102232112233312-0332123113232113-1131213101211320-3211310122131000-0030202322111201-0103202332111313)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0301302122312100-0303121321221121-2233002013020312-1211203203103312-1033012321032120-1121133311231131-1221332120300202-1103321303300103)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0301032221101120-3103200113001310-3223232331220200-1332133303133312-1323311130131310-1001333230111322-0230221210331302-3200332121101102)
- [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3003120302033030-0022133033303133-1021310202213331-0113203001333310-0100202002013200-2202322230231323-0010302133310002-3313303112221211)
- [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1032301033033313-1101300232301332-0110322320310030-1032321212200203-0212320010031030-1100000123210022-0031113133102112-0301111301311103)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-2013011230133310-0110112322303323-1102232112233312-0332123113232113-1131213101211320-3211310122131000-0030202322111201-0103202332111313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320311311320032-2211101212313231-1311001312032322-1113331210210222-3201231331022021-3113113122312332-1021302313321220-2101101202031321"></a>

## allow_all — allow_all / 313002321320 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- allow_all

<a id="canonical-2220010203033123-1110321130230310-3023220022020123-2201003102133013-3330100232323231-2031021200012211-0323121203331011-2113031103331212"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all, allow\_list, deny\_list, rule\_list\] Enable this option

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

OneOf alternatives in this subsection:

- [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2220010203033123-1110321130230310-3023220022020123-2201003102133013-3330100232323231-2031021200012211-0323121203331011-2113031103331212)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1131021322232311-1013030302000130-0100233010200032-3222300200230310-2310111103303033-3003220000012102-3210000312000223-1112313030320123)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1001323031020030-0323022102120112-0100301120113130-2133213212210110-0212222011312233-0032322201311210-2203332013010002-1320313230312301)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3300003111010003-1022200013220321-2120312030000021-0322110310133010-2300202200330301-2102222012123212-3322012103010130-2332110210313311)

Select alternatives according to the provider validators above.

<a id="canonical-3120221333313232-1031220012103331-1013010213002300-3213213313322003-2012001303102232-3220120023003103-1131022101000303-2310310201030101"></a>

## Direct properties — allow_all / 313002321320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320112331223001-0113332212223102-1330012313331132-0022311303230103-2010220133310330-1331211332200033-3031133000322230-3122002210203110"></a>

## Next pages — allow_all / 313002321320 / 4

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311313221312121-0033321013313332-1233333123212133-2032132131200200-2003003202121310-2102003333101200-0232303310012202-2302302101123310"></a>

## allow_list — allow_list / 202123221122 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- allow_list

<a id="canonical-1131021322232311-1013030302000130-0100233010200032-3222300200230310-2310111103303033-3003220000012102-3210000312000223-1112313030320123"></a>

Type: `"single"`. Computed.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

<a id="canonical-0321222002133323-3320130323122110-2031322331123113-3031103121203011-0013101311232200-1302102222210213-0220201113133011-1113221223212230"></a>

## Direct properties — allow_list / 202123221122 / 3

- [default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0030303121002013-0030230300323302-0331211020221310-2031122301233132-2302323223020300-1111032132203212-2110232020320101-2222103003013032): complete subsection reference.

- [default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3113302312100001-2023231111123213-2001132012213201-3201232313122321-1303331212012032-0000213333223112-0332000331031330-3200202312201302): complete subsection reference.

- [default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0033122023312122-3323023000110110-2233110321213003-3032320103323031-1301233130322022-1321202110021120-2221023211021333-1231012012012121): complete subsection reference.

- [dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1312200232331133-0311010303232332-0221203000300221-3030003302210100-0321011202310111-3201210103311212-3203032110002002-0031020320121203): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2003210313333300-1210322233003121-2111232031231120-3330220111130030-3303312221122102-3032333222111200-1220013320222331-2322233331321302): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0130102331212330-2221222233301120-3102100211102202-2313010000230021-0201201101130132-0323322310303120-3203003310013103-2330103233003213): complete subsection reference.

<a id="canonical-1202231202211002-1021303130211210-0220303322222320-0031233132023232-0111201032011030-3033030333203301-1300303210330311-0120332310233211"></a>

## Next pages — allow_list / 202123221122 / 4

- [allow_list.default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0030303121002013-0030230300323302-0331211020221310-2031122301233132-2302323223020300-1111032132203212-2110232020320101-2222103003013032)
- [allow_list.default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3113302312100001-2023231111123213-2001132012213201-3201232313122321-1303331212012032-0000213333223112-0332000331031330-3200202312201302)
- [allow_list.default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0033122023312122-3323023000110110-2233110321213003-3032320103323031-1301233130322022-1321202110021120-2221023211021333-1231012012012121)
- [allow_list.dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1312200232331133-0311010303232332-0221203000300221-3030003302210100-0321011202310111-3201210103311212-3203032110002002-0031020320121203)
- [allow_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2003210313333300-1210322233003121-2111232031231120-3330220111130030-3303312221122102-3032333222111200-1220013320222331-2322233331321302)
- [allow_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0130102331212330-2221222233301120-3102100211102202-2313010000230021-0201201101130132-0323322310303120-3203003310013103-2330103233003213)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0030303121002013-0030230300323302-0331211020221310-2031122301233132-2302323223020300-1111032132203212-2110232020320101-2222103003013032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323202103300311-3302302202033013-1033213203321310-1002011303302022-2213323310331212-1200011022233213-1331023003203001-2101010223310310"></a>

## allow_list.default_action_allow — default_action_allow / 212302200113 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.default_action_allow

<a id="canonical-2022122121202213-3223323003033031-1223233302001210-0311322121231330-2220322313000010-2223021012321233-2100310212011003-3130211221101103"></a>

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

<a id="canonical-3301220233111222-2223232003110022-1202331203012002-2023203001133202-0132331210002313-0033001333021100-1210103230210133-3203232221320323"></a>

## Direct properties — default_action_allow / 212302200113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223013101223101-0103032032321211-1202332011101231-3023133331023022-0012320031002232-3120300022131122-3223213311323330-3122021023301121"></a>

## Next pages — default_action_allow / 212302200113 / 4

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3113302312100001-2023231111123213-2001132012213201-3201232313122321-1303331212012032-0000213333223112-0332000331031330-3200202312201302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322100001200020-0002223210032222-1202333223133000-2303213010002232-0001120201103232-2323020333320013-2201223110222103-3220221111030300"></a>

## allow_list.default_action_deny — default_action_deny / 021203020032 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.default_action_deny

<a id="canonical-3030100300212120-1000233011133121-0032302022330313-0130000221313222-2010212200131003-2220032023333303-2110033223213232-0320320131212330"></a>

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

<a id="canonical-2200121201302023-0321303222021332-3021001133301110-1313101300100000-2122000100323213-0011012303001020-0202331303211103-1232102001321313"></a>

## Direct properties — default_action_deny / 021203020032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111112110021321-0132312203111211-1331102130300312-1200013111012000-1210320130232312-1231000111220032-3233322111003033-0023223011021021"></a>

## Next pages — default_action_deny / 021203020032 / 4

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0033122023312122-3323023000110110-2233110321213003-3032320103323031-1301233130322022-1321202110021120-2221023211021333-1231012012012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201113103132002-1130033210000301-3033301122111330-3232201022330231-2303131221103112-1100001010101121-2311213101312322-0000001000221100"></a>

## allow_list.default_action_next_policy — default_action_next_policy / 330120132303 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.default_action_next_policy

<a id="canonical-2010310213310011-0112301002210230-0130321233033311-0330313012221012-2202211302020012-1132032311212202-1001301030022200-0122110203220032"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-3033200101222220-0212032230322113-2321112213010021-3110203211012210-1210001330320323-1012003212112121-2100211211033110-1312230323230311"></a>

## Direct properties — default_action_next_policy / 330120132303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202113231003233-2303321121031112-3000232031020020-0003322300133232-1000321111220001-3001003313330133-1211212232220112-2222300223310332"></a>

## Next pages — default_action_next_policy / 330120132303 / 4

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-1312200232331133-0311010303232332-0221203000300221-3030003302210100-0321011202310111-3201210103311212-3203032110002002-0031020320121203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131032101022210-1223221213133331-0000213001130222-2332231313021330-3300310011131301-2303232300330011-1133203013311102-1322100212333020"></a>

## allow_list.dest_list — dest_list / 003033332002 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.dest_list

<a id="canonical-3213201020300211-2123023211301310-1121012133222102-1022101230133000-3231102212113031-1023130302323022-3103202321012211-2220021030033013"></a>

Type: `"list"`. Computed.

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

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

<a id="canonical-0301203230031113-2003221212031303-3221231313102301-3221111022113302-0003322320031022-0211232113300301-0121212123331210-1101122230333310"></a>

## Direct properties — dest_list / 003033332002 / 3

<a id="canonical-3123303120212202-0312212231030033-0301003230313031-2113012211102221-0311301333311112-0202311221012310-1031211031221223-2031132301331211"></a>

<a id="canonical-1000122110221022-1333330223032123-2213032302323210-3101000131333113-3232331101030210-3223201233322113-2210111213203232-1132330030032023"></a>

## ipv6_prefixes property — dest_list / 003033332002 / 4

Type: `["list", "string"]`. Computed.

IPv6 Prefixes. Destination IPv6 prefixes.

Upstream description:

Destination IPv6 prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1220300203003101-3120310303023113-0203213101230303-1033021212313121-0233131123221022-2210033332131212-2211120222322233-1220020120013322"></a>

<a id="canonical-1200212332123102-3133311011230030-0010031313031230-1130232200020332-0120200012322111-2220221200212200-0310031002221113-0303303122202020"></a>

## port_ranges property — dest_list / 003033332002 / 5

Type: `"string"`. Computed.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Upstream description:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-3231322001003123-1212111213133100-2100332113223313-3033233303230033-2232313200003033-3101302202121311-0103012310310312-2201020101312310"></a>

<a id="canonical-3101223202111103-0010133210133000-1021000003211223-2110320210101020-2002012333133311-3331332010221200-3312232033022301-0032233322023212"></a>

## prefixes property — dest_list / 003033332002 / 6

Type: `["list", "string"]`. Computed.

IPv4 Prefixes. Destination IPv4 prefixes.

Upstream description:

Destination IPv4 prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2133320030330130-1021202130033011-2312303000313030-2311321231103220-2031331030203000-3203212111111321-2223300203131310-0330210322300201"></a>

## Next pages — dest_list / 003033332002 / 7

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-2003210313333300-1210322233003121-2111232031231120-3330220111130030-3303312221122102-3032333222111200-1220013320222331-2322233331321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232230033101013-2331320302330031-3320223100032303-3023300332211221-1113103033333313-1020132212330000-2002333302201213-1031003110102000"></a>

## allow_list.http_list — http_list / 303030031202 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.http_list

<a id="canonical-3000203003021202-2211013203331301-1330110033302100-2112133303112221-0013302320330110-1210220123033031-2230212100031122-2101201012303011"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

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

<a id="canonical-0310203032210011-1112010211003313-1023021223210102-0211111301303211-1231100323012120-0300332130001303-3130233113320020-1232331333202300"></a>

## Direct properties — http_list / 303030031202 / 3

- [any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0112333302030223-1012002210301010-3003120220201120-2111211021212231-0222030110113232-0101203023222230-3320102110000233-0230130110202023): complete subsection reference.

<a id="canonical-3322110132312230-0232233022122131-2013223103023310-2230101001310130-3020132122321212-1030232023032221-1111300000232322-2122213213313331"></a>

<a id="canonical-0112301132003020-3021210032313011-2320120002001221-3031231310232001-0332011310103121-0120311223022023-2112233113003330-3203213233221200"></a>

## exact_value property — http_list / 303030031202 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1132131320100203-2202201021331301-1020320021023113-1213021212331100-2021330302131002-1212313201031031-2001001113123310-1210020223203030"></a>

<a id="canonical-0323233302221202-1332001101112202-2002121321031012-3212321211013211-1210132113321000-3132222031030130-0223131300311322-0032210332021323"></a>

## path_exact_value property — http_list / 303030031202 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1022231130033231-0110222032301210-0211313203320001-3213221111002313-1023301230202103-1330333113013010-3111113211111320-1002232023203333"></a>

<a id="canonical-3321233313002120-1020001203110311-1311103212013103-1012203002031113-1301100333332302-2011101311010133-3232310203233123-1220213020022112"></a>

## path_prefix_value property — http_list / 303030031202 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1031322021022022-1232132013220131-0323010222123232-1333123302023330-3111230213203320-3222003222101012-3331010133003101-3220303023003302"></a>

<a id="canonical-1112301102131133-1230020103130320-2321111301211111-1110230212331101-0313002013100212-0110103033013022-3232313323101210-0302232012203303"></a>

## path_regex_value property — http_list / 303030031202 / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3211311133000030-0301331303232311-1221022202020102-1001311113131012-2022210231101013-0312100313203231-1202231221211112-1321103113331313"></a>

<a id="canonical-0201101030110130-0133022201120033-1123120323220323-3001322232020200-0002102222211010-2030311122013120-2123322210220123-1200321322202322"></a>

## regex_value property — http_list / 303030031202 / 8

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2322123031203333-2330310021200300-1222031301002132-2112323203210200-1031203300002012-0012233001003201-0010101130021211-1122301131000031"></a>

<a id="canonical-3212110312200130-0203331323003222-0222330232221020-1113031000331323-0001333221210133-3220323322112333-1322010032331103-1202133111132200"></a>

## suffix_value property — http_list / 303030031202 / 9

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2113010330111203-0030011012031223-3022122123010031-0003200121122001-1013103000201220-3133203200203202-0112122311210010-3131031011011300"></a>

## Next pages — http_list / 303030031202 / 10

- [allow_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0112333302030223-1012002210301010-3003120220201120-2111211021212231-0222030110113232-0101203023222230-3320102110000233-0230130110202023)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0112333302030223-1012002210301010-3003120220201120-2111211021212231-0222030110113232-0101203023222230-3320102110000233-0230130110202023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200033010121221-1100203322022210-0031213121320000-2230232021102002-2120311111023332-1221110030212111-0231331003230000-3201012021312200"></a>

## allow_list.http_list.any_path — any_path / 030000020032 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- [allow_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2003210313333300-1210322233003121-2111232031231120-3330220111130030-3303312221122102-3032333222111200-1220013320222331-2322233331321302)
- allow_list.http_list.any_path

<a id="canonical-1313032120321232-3120020232300320-0312031332102322-1013232212103311-0102003333231130-0233303131210010-3210032231010320-3210100210132102"></a>

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

<a id="canonical-3202301110120212-1131300001323033-0113013202003002-1303110320331213-1201130132322002-3223013233330210-0221330120232332-2031223023132010"></a>

## Direct properties — any_path / 030000020032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102330102322031-2121330322102330-3233112010002113-1201230100122330-0232300111203023-2131102001000100-2120232320011022-2201330300220102"></a>

## Next pages — any_path / 030000020032 / 4

- [allow_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2003210313333300-1210322233003121-2111232031231120-3330220111130030-3303312221122102-3032333222111200-1220013320222331-2322233331321302)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0130102331212330-2221222233301120-3102100211102202-2313010000230021-0201201101130132-0323322310303120-3203003310013103-2330103233003213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131310103230131-3123310032033120-1203232202010230-0200010333323013-0133002020101202-3312103232131102-2333001023212003-2133301013031111"></a>

## allow_list.tls_list — tls_list / 122212032032 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.tls_list

<a id="canonical-1310033300013021-0303211310233301-1333201112312230-1020111020333000-3032230113022100-3221231011302033-1232313210111013-1311210021321222"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

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

<a id="canonical-3111032032313302-3002113230202332-0133130320111300-0101022221020201-3230102012003133-2333121012110000-3100303320310220-0122213102021002"></a>

## Direct properties — tls_list / 122212032032 / 3

<a id="canonical-2303011022320211-1133031222301231-2311313030113101-1201030220130001-0312032211012022-1101232000212132-2030310221010331-0132101022113301"></a>

<a id="canonical-0121131210333333-3030301031120200-0011001312231100-0022312333130301-1233221200233322-2130302002200233-3213023122031032-1232030222310010"></a>

## exact_value property — tls_list / 122212032032 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3032002112232013-0122321031213001-0013331012103003-2320111132232221-2010332121211322-0330120213303122-0213123213132231-1312002130310123"></a>

<a id="canonical-3021323333300303-2133130023001222-0101013001213130-0220320010122201-0020130113200132-3011012212121131-2131032223133031-3001133333303123"></a>

## regex_value property — tls_list / 122212032032 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3322022102122330-1231001021312311-1303011111102001-1312031210000333-3231222021213232-0033013221320131-3122313320101231-3121212130221213"></a>

<a id="canonical-2332101331212203-3101030030021131-0021323230101000-3121201030011230-2123023130213122-3120210310232101-1031320031332023-3130123200020231"></a>

## suffix_value property — tls_list / 122212032032 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0102112201123223-3102333131333322-3320111200301202-3113103022033313-2110010323130333-0112312302212020-2001201320123031-1032323031023230"></a>

## Next pages — tls_list / 122212032032 / 7

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0301302122312100-0303121321221121-2233002013020312-1211203203103312-1033012321032120-1121133311231131-1221332120300202-1103321303300103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101021023231031-2113130220101333-2132120322201232-3130110021311320-1331003302110220-1133320223200332-2131233302233113-2021331221322113"></a>

## any_proxy — any_proxy / 211010212311 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- any_proxy

<a id="canonical-3123313021211132-2130110232333032-0303111023110033-3112321221232113-3211111022110111-2013331232011111-0032333110001313-3002032103013321"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_proxy, drp\_http\_connect, network\_connector, proxy\_label\_selector\] Enable this
option

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

OneOf alternatives in this subsection:

- [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3123313021211132-2130110232333032-0303111023110033-3112321221232113-3211111022110111-2013331232011111-0032333110001313-3002032103013321)
- [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1232303201032011-2232231121322221-2301300110332023-1032001113103110-3110111022311231-0232330102320221-3220013322023130-3203010220300301)
- [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3121013223110112-3301132223320301-3111002130110031-3133111102233212-3002211032112332-0123211112320311-1111302121022100-3210222210003131)
- [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3320221102222111-2312332233123123-0022022013122310-0231213333221113-3301223112111020-0120030021112201-1210012120013200-1210110111112213)

Select alternatives according to the provider validators above.

<a id="canonical-1033300313310203-2210210312030002-2302311132110123-2222100202020123-0031222231021231-0231001013220203-1232120303130131-3010212000203330"></a>

## Direct properties — any_proxy / 211010212311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111101313031202-3231303001113322-3133130302131121-2213032033301122-1231002222130331-3002211232003322-0211123322002131-2111200202003231"></a>

## Next pages — any_proxy / 211010212311 / 4

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203233110032031-0212300331032002-1221132331030310-2013132232111112-3232300030310221-1311101011321103-3313230131200013-3031030113302323"></a>

## deny_list — deny_list / 332010221133 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- deny_list

<a id="canonical-1001323031020030-0323022102120112-0100301120113130-2133213212210110-0212222011312233-0032322201311210-2203332013010002-1320313230312301"></a>

Type: `"single"`. Computed.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

<a id="canonical-0122103233013230-3132110211212021-3231221023002130-0333310121133032-3110330013132211-3232301011133302-1202130101201221-2322102323000013"></a>

## Direct properties — deny_list / 332010221133 / 3

- [default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2320023020013103-2020321103323012-0111100023132313-1300030130011212-1100013300221331-1122003113221010-3023301200120322-2103003103223210): complete subsection reference.

- [default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1321033032102020-2223102003211213-2002133030321300-1231211110303320-2220302021100203-0131303313131022-0110030231122211-2112120222320210): complete subsection reference.

- [default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1211233001331303-0113112300220112-3013130022201000-2003023332212223-0023233311310321-1030032112211001-3232223121112030-0120322122033330): complete subsection reference.

- [dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3121022111323330-2120112103213211-0021311202003332-2320031121110013-1230010313311233-3213010111011323-3313312310120221-3020003130112323): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0303231023303013-2000101121012313-0330202210132020-1231300133103021-2013033113212023-1103013023330023-2023023313001213-2111003101313132): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2113232131200211-1332200130321213-3202223120213122-1231110313001033-0110120003233212-0233311012330332-1233102100022213-2221100112210232): complete subsection reference.

<a id="canonical-0133112123111310-0300212301022221-3221003300011103-0101012312000123-3230013110030320-3012331022301221-1013203132032232-3323120220302103"></a>

## Next pages — deny_list / 332010221133 / 4

- [deny_list.default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2320023020013103-2020321103323012-0111100023132313-1300030130011212-1100013300221331-1122003113221010-3023301200120322-2103003103223210)
- [deny_list.default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1321033032102020-2223102003211213-2002133030321300-1231211110303320-2220302021100203-0131303313131022-0110030231122211-2112120222320210)
- [deny_list.default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1211233001331303-0113112300220112-3013130022201000-2003023332212223-0023233311310321-1030032112211001-3232223121112030-0120322122033330)
- [deny_list.dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3121022111323330-2120112103213211-0021311202003332-2320031121110013-1230010313311233-3213010111011323-3313312310120221-3020003130112323)
- [deny_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0303231023303013-2000101121012313-0330202210132020-1231300133103021-2013033113212023-1103013023330023-2023023313001213-2111003101313132)
- [deny_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2113232131200211-1332200130321213-3202223120213122-1231110313001033-0110120003233212-0233311012330332-1233102100022213-2221100112210232)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-2320023020013103-2020321103323012-0111100023132313-1300030130011212-1100013300221331-1122003113221010-3023301200120322-2103003103223210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130103022123012-1303320320101312-3032300012333103-0230300011220132-1030133203303123-3021202003321121-1000132322031212-1222213113012100"></a>

## deny_list.default_action_allow — default_action_allow / 211300023313 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.default_action_allow

<a id="canonical-3202123133231120-3101322131023113-1123322322131321-0130100100233133-3021322023330010-0310121002321033-3123110211202120-0110003111203300"></a>

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

<a id="canonical-1220303000300000-3233112211312323-1322021031303121-2303030002033000-0023302002113212-1021032111332321-1332022030302220-2123323123113013"></a>

## Direct properties — default_action_allow / 211300023313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031122201300011-1222113212000022-1303003030220203-2012203013223022-0321221201203132-2121330013312002-1332100001030100-3220120000001232"></a>

## Next pages — default_action_allow / 211300023313 / 4

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-1321033032102020-2223102003211213-2002133030321300-1231211110303320-2220302021100203-0131303313131022-0110030231122211-2112120222320210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221302220113300-0013133331003210-2300103301002110-3300312020120013-3233021130202013-1200002203232000-1210033102223131-3312200231100312"></a>

## deny_list.default_action_deny — default_action_deny / 111113001310 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.default_action_deny

<a id="canonical-2112111020202300-2222202333232012-1022323012033111-0030230131103010-2112211121120322-2013321023210300-3200122001111102-0300303301030332"></a>

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

<a id="canonical-1131000122300212-3101212130200330-1121313100131211-0200311112013002-1203003311221101-2232102013222112-0232011132310211-3032303320111220"></a>

## Direct properties — default_action_deny / 111113001310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011301211130303-3303003101320322-1313211120120202-2220011033222032-0221323321123313-3100331233133231-2023032031112030-3100231311000310"></a>

## Next pages — default_action_deny / 111113001310 / 4

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-1211233001331303-0113112300220112-3013130022201000-2003023332212223-0023233311310321-1030032112211001-3232223121112030-0120322122033330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123031013021121-2022010100220313-2330121303000323-1030331232231231-3122012132132021-2132231230313233-0011121322311220-1321022133131300"></a>

## deny_list.default_action_next_policy — default_action_next_policy / 120202321211 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.default_action_next_policy

<a id="canonical-3313011032001133-0311320311311111-2131321213230021-3302013201121103-1131001131231221-1332231010202103-2031302211130121-3213313202020113"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-1302102313210303-0232220130111031-2231123230102201-3202021033031322-0301211030200201-0223220120200020-1111232000331131-0210012022013311"></a>

## Direct properties — default_action_next_policy / 120202321211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213200300310212-3112110110023211-3302321211113330-0131322103313102-3110101102111321-1002001312002201-0331012321330031-0121120232311203"></a>

## Next pages — default_action_next_policy / 120202321211 / 4

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3121022111323330-2120112103213211-0021311202003332-2320031121110013-1230010313311233-3213010111011323-3313312310120221-3020003130112323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331313022001122-1310231011233012-1130031302332230-3313112230030123-2122121131031233-0020202332032031-0121023132223200-3013202320101311"></a>

## deny_list.dest_list — dest_list / 230202201303 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.dest_list

<a id="canonical-2331003033203130-1302130022131313-3312031100210302-2233131133332231-3200232200213002-2321111202110330-3012300203213133-1003302211003312"></a>

Type: `"list"`. Computed.

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

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

<a id="canonical-3300201002311210-0230110202233031-0121312331221032-3101031032130003-3011133132231200-2021102313222101-0021131111013122-1110010202301113"></a>

## Direct properties — dest_list / 230202201303 / 3

<a id="canonical-2330010031030121-3331021223202330-1200330222120201-3001310332111220-3123131203212112-0130231312102002-2111020202222123-1002331302201201"></a>

<a id="canonical-3121211012321200-1203032323133111-0223111330222001-2212330030303022-3101213010132021-2103113033333013-1112021232133003-1110032032323033"></a>

## ipv6_prefixes property — dest_list / 230202201303 / 4

Type: `["list", "string"]`. Computed.

IPv6 Prefixes. Destination IPv6 prefixes.

Upstream description:

Destination IPv6 prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1120312030113121-3303211131320001-0012130201122201-3200223103312300-2231100020001321-1002231003313320-2122301100223102-0200301020311300"></a>

<a id="canonical-1113301011032022-2130003120203103-0100100221111021-1101321103021222-0201231223012210-2020322023321131-1103013102203311-1220013211233303"></a>

## port_ranges property — dest_list / 230202201303 / 5

Type: `"string"`. Computed.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Upstream description:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-0101323223130201-2203121332132113-3320232221023332-1231023302332320-1010203200333311-3320233031300010-3010231211222222-1212030010203031"></a>

<a id="canonical-3031130312331221-0322111230320233-0301102220213011-1333203003031022-0232321320320111-2222311333330002-3301002122213323-1111323002232013"></a>

## prefixes property — dest_list / 230202201303 / 6

Type: `["list", "string"]`. Computed.

IPv4 Prefixes. Destination IPv4 prefixes.

Upstream description:

Destination IPv4 prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1001320131002213-3121033222030220-0121210003201321-0303133021120021-3333132031222301-1131101313203033-2222113331133300-2001002330313310"></a>

## Next pages — dest_list / 230202201303 / 7

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0303231023303013-2000101121012313-0330202210132020-1231300133103021-2013033113212023-1103013023330023-2023023313001213-2111003101313132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301302332333133-3303112133031131-1312223002113223-2200301331330122-0111122021330110-0123333023031222-0201312112130203-0020021112130300"></a>

## deny_list.http_list — http_list / 223111013330 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.http_list

<a id="canonical-3332030102031203-0322222033321303-0022302000331031-3003201030233013-0131122203222102-3222210031012032-0211203203113000-3233123021230211"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

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

<a id="canonical-3320100021000113-2311111103320103-1213112201201313-2102131233210320-2330200103212020-1220310212210333-0232002112022301-1222120320301122"></a>

## Direct properties — http_list / 223111013330 / 3

- [any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2113310132020112-0200202022310230-2221330231311012-2231221100213331-0302201312201010-1321002101330110-3320222232213123-3131302301011300): complete subsection reference.

<a id="canonical-0012230302020123-0330301232311322-3330003032122112-0131201121211311-2201020133100120-1322133021113032-3111033213133313-0312212200101212"></a>

<a id="canonical-1232221221100301-1332030311220121-0210131210301231-1010131021031321-0001120113010130-1033232231312312-3011213333203102-1000323020320330"></a>

## exact_value property — http_list / 223111013330 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0211020133310020-0131333020120021-3032322211001223-2323233322122303-0022202010231001-3000202130022103-2223222001032020-1321323133011231"></a>

<a id="canonical-2210101222032312-1211001330221001-3333212110132233-0012103322123200-3002310232102302-0023121101021130-0300301111012233-1133310302122111"></a>

## path_exact_value property — http_list / 223111013330 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3020331321013002-0012130110023033-1212210101030202-2230021133020021-1332313220102032-0133021013203303-3101002230130121-2013032333302233"></a>

<a id="canonical-2003130213303013-0012021232201023-0001112210203100-3100303111310001-1123033001312313-3331113031122302-0320120201333300-3133130331220202"></a>

## path_prefix_value property — http_list / 223111013330 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1103223221022112-3012223233300331-0021313223330230-0200331110223000-2302300233123323-2213300031101003-0002211330123233-3002110010302223"></a>

<a id="canonical-0030321311130232-1301302312101132-0111011331311321-0102023333321202-2211233023332100-0033030230220003-2233322222121002-3300022130002013"></a>

## path_regex_value property — http_list / 223111013330 / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1012112322021113-0023001033230232-0131133121310233-2131220002021111-1310222030121131-3033122200313031-0301121312103213-2301000300013103"></a>

<a id="canonical-0032300231100230-3033223211333203-3011111311021022-3311100302002313-3103310310031322-0200120202220310-0213030303232130-0332303312022302"></a>

## regex_value property — http_list / 223111013330 / 8

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2010322031123131-0231020202311312-2303112332233133-2321101203211130-2210320213221311-1111133023102233-2221331333133333-0330201112310112"></a>

<a id="canonical-1320221033001133-2212030023221132-2332212303321312-3033003031312121-0013300012321131-3103103121020320-3032122030132130-1201211013013200"></a>

## suffix_value property — http_list / 223111013330 / 9

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2300331100301010-0132033331133320-3311203333123221-1232113132322221-2202100211202201-0100320213003212-1030131200303332-3032213012331312"></a>

## Next pages — http_list / 223111013330 / 10

- [deny_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2113310132020112-0200202022310230-2221330231311012-2231221100213331-0302201312201010-1321002101330110-3320222232213123-3131302301011300)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-2113310132020112-0200202022310230-2221330231311012-2231221100213331-0302201312201010-1321002101330110-3320222232213123-3131302301011300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123233203033110-2321001103003312-2133101312000333-3110320122213033-3201013123331300-0230031231231010-1121301303231020-0211023223330313"></a>

## deny_list.http_list.any_path — any_path / 201000332033 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- [deny_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0303231023303013-2000101121012313-0330202210132020-1231300133103021-2013033113212023-1103013023330023-2023023313001213-2111003101313132)
- deny_list.http_list.any_path

<a id="canonical-3230132322031011-2133220301300220-2012122102202221-3200111331031203-2211310032233301-0011203030130120-0203231311212332-3210111123303223"></a>

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

<a id="canonical-3332331031122132-0230003010003212-2201210210303331-0032033032112000-2211120100223120-1101110111210000-2230303222301300-0011203210032023"></a>

## Direct properties — any_path / 201000332033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112311011033231-3003102303313001-1131330113331111-1011302122220200-3131012032302000-2111203133012310-2230232102010003-1110313302112100"></a>

## Next pages — any_path / 201000332033 / 4

- [deny_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0303231023303013-2000101121012313-0330202210132020-1231300133103021-2013033113212023-1103013023330023-2023023313001213-2111003101313132)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-2113232131200211-1332200130321213-3202223120213122-1231110313001033-0110120003233212-0233311012330332-1233102100022213-2221100112210232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030102232000213-1321211213320002-0211303131111300-1110310010312033-2101232010330023-1112200330220101-3222000010201033-2103222002232031"></a>

## deny_list.tls_list — tls_list / 102100202031 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.tls_list

<a id="canonical-1000130231032301-3133031001131333-2331322231330023-3331131322033322-2231110111010132-1021221122331200-1100230331222101-3303103021121013"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

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

<a id="canonical-2112332210130303-2333120120313103-2121221233031110-3123033013310010-2321221002320202-0111230311221222-3323330122023323-0022303312322121"></a>

## Direct properties — tls_list / 102100202031 / 3

<a id="canonical-2200321232000100-1203230131131332-1301301230020232-1003231001131313-1132120113331121-2030011033213231-2202133130131212-2203110230100321"></a>

<a id="canonical-2210122300233202-1002120312231103-3311222120010301-0133303010100310-2331131323301223-2011331301022012-1200313020131300-3003222000031313"></a>

## exact_value property — tls_list / 102100202031 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1101333332133120-0311301001333310-0332201220121120-0332203313022231-3203330013022211-3203323233132123-3020112203101120-0100020010231102"></a>

<a id="canonical-3312123120302003-2201222021001312-2132302102323013-3311332221121023-3013201323203313-2023032223222223-1133001030130210-1122313211020222"></a>

## regex_value property — tls_list / 102100202031 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3313012111113221-2110330230202032-1130301331020233-3330121010122013-3312120100002003-3001311131321302-1101323222222122-2202210223310123"></a>

<a id="canonical-3222120212202312-2323123111012211-0011023132322012-3223222101232200-2013020111300322-3101133303000213-1022113003230320-0313201011321303"></a>

## suffix_value property — tls_list / 102100202031 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2330101110021003-1032103131221331-1011112231023303-3032102200012122-0110131102133123-1110231122032202-2000300122321213-2210323112021003"></a>

## Next pages — tls_list / 102100202031 / 7

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0301032221101120-3103200113001310-3223232331220200-1332133303133312-1323311130131310-1001333230111322-0230221210331302-3200332121101102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123100332212121-0023122103011121-0213322023301131-1000030332023223-3023021311111001-1313303323101022-2313001222321210-0022003122303313"></a>

## drp_http_connect — drp_http_connect / 201322012023 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- drp_http_connect

<a id="canonical-1232303201032011-2232231121322221-2301300110332023-1032001113103110-3110111022311231-0232330102320221-3220013322023130-3203010220300301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for drp http connect.

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

<a id="canonical-1211113113202232-0331123031320223-2303331001002201-3023032312310030-2232303222212131-1333103223121023-2030003222320330-2003330323123210"></a>

## Direct properties — drp_http_connect / 201322012023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323221310122321-0310032321031031-2231032302333022-1022313203122023-3031101313010203-3011223111101112-0021013313013331-1002232211320212"></a>

## Next pages — drp_http_connect / 201322012023 / 4

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3003120302033030-0022133033303133-1021310202213331-0113203001333310-0100202002013200-2202322230231323-0010302133310002-3313303112221211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212211213222122-2212233000302031-1231120032021012-3200200222130120-3133312223131221-2332213101321122-3010210102030223-3010203000001201"></a>

## network_connector — network_connector / 331303223121 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- network_connector

<a id="canonical-3121013223110112-3301132223320301-3111002130110031-3133111102233212-3002211032112332-0123211112320311-1111302121022100-3210222210003131"></a>

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

<a id="canonical-0020332333203301-2021310003101213-3320300012002303-0310102300321113-3030210300310120-1330001133003332-2023331302203022-2123021322002233"></a>

## Direct properties — network_connector / 331303223121 / 3

<a id="canonical-2003113321102020-3010121213103312-3031033201303331-0313231000121112-0232203233330233-2032111202212003-3031322300210301-1203021101021222"></a>

<a id="canonical-2223110222303222-3023133122311123-0320330103021002-0133000100130330-0132230132330101-3101202321130020-0001232021333210-3322322301123013"></a>

## name property — network_connector / 331303223121 / 4

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

<a id="canonical-3001003103103132-0100122102320212-1322201123202103-0102123201312231-3201023213121303-0323322132203322-2003132202010323-3222021230313000"></a>

<a id="canonical-1110220110310112-2313230003003021-0230001212033213-3311311232330011-0132112221333003-2103001121310223-1312113223003102-0130130212211032"></a>

## namespace property — network_connector / 331303223121 / 5

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

<a id="canonical-2300322032101212-2102011102022232-1123320022202301-0232110132313232-2102300220002003-2112303223102033-2111222121323110-2233110200123221"></a>

<a id="canonical-1023323122030220-1312230331101030-1002231311302110-1102302223011301-2013332203011123-0123323020322100-3030033022003013-3311333230302303"></a>

## tenant property — network_connector / 331303223121 / 6

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

<a id="canonical-1033012330231203-1102313023020320-2332033210133111-3212210032320022-2230103003020022-2213230201332211-3133311310100332-2033222331301320"></a>

## Next pages — network_connector / 331303223121 / 7

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-1032301033033313-1101300232301332-0110322320310030-1032321212200203-0212320010031030-1100000123210022-0031113133102112-0301111301311103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102102213213222-0232300100102113-3211210123020002-3332032021233120-2133110103222122-1033000030212130-2233102103121310-0103311130322233"></a>

## proxy_label_selector — proxy_label_selector / 101312232212 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- proxy_label_selector

<a id="canonical-3320221102222111-2312332233123123-0022022013122310-0231213333221113-3301223112111020-0120030021112201-1210012120013200-1210110111112213"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-2212113011111320-0223031221033031-2033233023330101-2310221233003123-1122023233133332-3210002003313022-1103301033010203-1030332002321300"></a>

## Direct properties — proxy_label_selector / 101312232212 / 3

<a id="canonical-1300023203010112-2003122212123121-0320231300323302-0312320001120013-2122100132003011-2212000122231010-3133130330220312-0311103002312001"></a>

<a id="canonical-3201221023102231-1103013100212032-3222302330230102-0130311032332111-3130230001032000-3032222213222030-0113121100011321-0132020003221331"></a>

## expressions property — proxy_label_selector / 101312232212 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3212300232001223-2123011012131020-3223020201230330-0202302223120313-0312300033332013-3301130201203221-3131000103021022-1021111301110321"></a>

## Next pages — proxy_label_selector / 101312232212 / 5

- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131103020121130-2231320103311321-3021310333212030-0210020132313111-0202211203211120-3333320233113230-3332113232021010-0230330310223032"></a>

## rule_list — rule_list / 200331002122 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- rule_list

<a id="canonical-3300003111010003-1022200013220321-2120312030000021-0322110310133010-2300202200330301-2102222012123212-3322012103010130-2332110210313311"></a>

Type: `"single"`. Computed.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

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

<a id="canonical-2101200020212221-0013303030001333-2022223120300103-3012112002122002-2113212322133223-3013313113113301-0100210233003123-1131300012331020"></a>

## Direct properties — rule_list / 200331002122 / 3

- [rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003): complete subsection reference.

<a id="canonical-2123112112132212-3301301220202313-3132220233201323-0110033000300101-2111013212232300-0233213300010111-3110222221301133-1133301232321102"></a>

## Next pages — rule_list / 200331002122 / 4

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333202101113131-1322003211033310-1000011220000220-0103303023120310-1321213101302202-3301211203323012-2002203010000200-0313220121111032"></a>

## rule_list.rules — rules / 132200233220 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- rule_list.rules

<a id="canonical-3100302032212022-2021123222100320-3213331000002321-2003220103311020-0101020103332002-1313332302120213-0321120011110223-2110233133303132"></a>

Type: `"list"`. Computed.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-3102200130310101-1320033222211333-2132320211100233-2203303023132313-2112021012011133-2223110012231002-2331201300210122-2202231122202003"></a>

## Direct properties — rules / 132200233220 / 3

<a id="canonical-1131000023102020-1231010222000020-3001331200220123-0203303331330312-1031123113111032-1000301210102321-2102313200023203-1200130222231333"></a>

<a id="canonical-2303200321110022-2112220203110231-2231333132213333-1201030230330221-1230313130211201-0333230132112230-2013102030221021-2133301230100331"></a>

## action property — rules / 132200233220 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW|NEXT\_POLICY\] The rule action determines the disposition of the input request
API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds
forward. If it matches a rule with a DENY action, the processing of the request is terminated and an
appropriate message/code returned to.. Possible values are \`DENY\`, \`ALLOW\`, \`NEXT\_POLICY\`.
Defaults to \`DENY\`.

Upstream description:

The rule action determines the disposition of the input request API. If a policy matches a rule with
an ALLOW action, the processing of the request proceeds forward. If it matches a rule with a DENY
action, the processing of the request is terminated and an appropriate message/code returned to the
originator. If it matches a rule with a NEXT\_POLICY\_SET action, evaluation of the current policy
set terminates and evaluation of the next policy set in the chain begins.

&#8203;- DENY: DENY

Deny the request. &#8203;- ALLOW: ALLOW

Allow the request to proceed. &#8203;- NEXT\_POLICY\_SET: NEXT\_POLICY\_SET

Terminate evaluation of the current policy set and begin evaluating the next policy set in the
chain. Note that the evaluation of any remaining policies in the current policy set is skipped.
&#8203;- NEXT\_POLICY: NEXT\_POLICY

Terminate evaluation of the current policy and begin evaluating the next policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
LAST\_POLICY: LAST\_POLICY

Terminate evaluation of the current policy and begin evaluating the last policy in the policy set.
Note that the evaluation of any remaining rules in the current policy is skipped. &#8203;-
GOTO\_POLICY: GOTO\_POLICY

Terminate evaluation of the current policy and begin evaluating a specific policy in the policy set.
The policy is specified using the goto\_policy field in the rule and must be after the current
policy in the policy set.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW",
    "NEXT_POLICY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [all_destinations](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0132013321333203-2123012010320331-2230021013321013-1130111030022220-1132232330000022-2311021020311230-3133132103021322-3113330232312103): complete subsection reference.

- [all_sources](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1033231212331302-0311032231200320-1111232330003011-2021232302220313-1200220323300203-2323112301323222-3001111212012001-3300332012333313): complete subsection reference.

- [dst_asn_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1300121121133123-1331331113003123-2123002310303113-2213203323123203-2322233201100012-2333000202102110-2233223130210112-3013002000222300): complete subsection reference.

- [dst_asn_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0333101311200103-1201232200123231-1312310331322313-0330220231113033-1023301323021201-2233333012002321-0200302211110200-2303223023100212): complete subsection reference.

- [dst_ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3313033010201213-1303212323211203-3223031032231111-1000013301232331-2032030123132131-2222301002313201-2232022003332321-3221120203213211): complete subsection reference.

- [dst_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0321111022030032-2322302230031301-2330333002000223-1223111132313113-0013121312101130-2001323322322232-3010000201300030-2311113000302123): complete subsection reference.

- [dst_prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1302021011333221-1202321123122020-2132230322001012-2021003030013130-3003132133103002-0232330130210222-0130211121032322-1300232013202233): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211): complete subsection reference.

- [ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2212103130223310-3330130302022032-0203030031113130-1003332132223122-1023111013210012-0030311012030221-2111023131032223-2232213101213000): complete subsection reference.

- [label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0121013023233311-1203120030030033-1332300301332300-0321210320000300-1112203010322331-1133332312311203-2212003212000211-0102321303223132): complete subsection reference.

- [metadata](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3223102201230333-0133323300113202-0020231011110213-1103132313011100-2303122321311003-3120230013203320-0203112302310020-0011101020213312): complete subsection reference.

- [no_http_connect_port](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3322110302231022-2313001321321311-0011112313123332-0331303200302032-1300033123310301-3131201202203220-0023301312310322-2323311123323123): complete subsection reference.

- [port_matcher](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3303223031202122-2130021033020311-3000130130333300-0233223101103033-0220002000332013-0100230323032120-1303100210111331-2311331133310002): complete subsection reference.

- [prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3212130112120111-0303230130202111-3022100232223233-3002201320311201-2201332023302332-2211231010010220-3022300313111010-2321121102331231): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3130031113132122-1313310010332323-3012313301202331-1020032123210331-2033230333200121-0311201313100230-3313331323102333-1023320301301321): complete subsection reference.

- [url_category_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0233202033222111-2122021223212100-2100002000120303-2133131211221023-2302211312302313-1031130223233300-0112000112200222-3200310020303232): complete subsection reference.

<a id="canonical-2113331110003021-0032010310322133-2333332222002010-2002112033100310-3231233121032310-3230311030021320-2031223311131330-1201020311331100"></a>

## Next pages — rules / 132200233220 / 5

- [rule_list.rules.all_destinations](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0132013321333203-2123012010320331-2230021013321013-1130111030022220-1132232330000022-2311021020311230-3133132103021322-3113330232312103)
- [rule_list.rules.all_sources](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1033231212331302-0311032231200320-1111232330003011-2021232302220313-1200220323300203-2323112301323222-3001111212012001-3300332012333313)
- [rule_list.rules.dst_asn_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1300121121133123-1331331113003123-2123002310303113-2213203323123203-2322233201100012-2333000202102110-2233223130210112-3013002000222300)
- [rule_list.rules.dst_asn_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0333101311200103-1201232200123231-1312310331322313-0330220231113033-1023301323021201-2233333012002321-0200302211110200-2303223023100212)
- [rule_list.rules.dst_ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3313033010201213-1303212323211203-3223031032231111-1000013301232331-2032030123132131-2222301002313201-2232022003332321-3221120203213211)
- [rule_list.rules.dst_label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0321111022030032-2322302230031301-2330333002000223-1223111132313113-0013121312101130-2001323322322232-3010000201300030-2311113000302123)
- [rule_list.rules.dst_prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1302021011333221-1202321123122020-2132230322001012-2021003030013130-3003132133103002-0232330130210222-0130211121032322-1300232013202233)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211)
- [rule_list.rules.ip_prefix_set](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2212103130223310-3330130302022032-0203030031113130-1003332132223122-1023111013210012-0030311012030221-2111023131032223-2232213101213000)
- [rule_list.rules.label_selector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0121013023233311-1203120030030033-1332300301332300-0321210320000300-1112203010322331-1133332312311203-2212003212000211-0102321303223132)
- [rule_list.rules.metadata](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3223102201230333-0133323300113202-0020231011110213-1103132313011100-2303122321311003-3120230013203320-0203112302310020-0011101020213312)
- [rule_list.rules.no_http_connect_port](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3322110302231022-2313001321321311-0011112313123332-0331303200302032-1300033123310301-3131201202203220-0023301312310322-2323311123323123)
- [rule_list.rules.port_matcher](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3303223031202122-2130021033020311-3000130130333300-0233223101103033-0220002000332013-0100230323032120-1303100210111331-2311331133310002)
- [rule_list.rules.prefix_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3212130112120111-0303230130202111-3022100232223233-3002201320311201-2201332023302332-2211231010010220-3022300313111010-2321121102331231)
- [rule_list.rules.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3130031113132122-1313310010332323-3012313301202331-1020032123210331-2033230333200121-0311201313100230-3313331323102333-1023320301301321)
- [rule_list.rules.url_category_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0233202033222111-2122021223212100-2100002000120303-2133131211221023-2302211312302313-1031130223233300-0112000112200222-3200310020303232)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0132013321333203-2123012010320331-2230021013321013-1130111030022220-1132232330000022-2311021020311230-3133132103021322-3113330232312103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312033032210232-2332302302322022-2231132101333000-1020112321231330-1101123332213303-0002230310131123-2010310030202300-1103132132200012"></a>

## rule_list.rules.all_destinations — all_destinations / 002231332123 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.all_destinations

<a id="canonical-3202113122323120-0131313120111220-3331020322313230-3012320022131230-3333301200321111-2210213110123333-1330333230301321-2002120302021203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all destinations.

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

<a id="canonical-3231330201002220-2102310302013022-3302333322223222-3030210033112203-0221202000130022-3232303231130303-1333120200100022-0321113132011202"></a>

## Direct properties — all_destinations / 002231332123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301112130131030-3121333301213221-3123313312332231-0130221012203302-3001313230223113-1121213311303311-1301131231101122-3323320101302113"></a>

## Next pages — all_destinations / 002231332123 / 4

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-1033231212331302-0311032231200320-1111232330003011-2021232302220313-1200220323300203-2323112301323222-3001111212012001-3300332012333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213012313031101-1130021130232021-0222013100200333-3231031033121133-3310332011211330-2013213323122230-0112003022233313-0000312101230100"></a>

## rule_list.rules.all_sources — all_sources / 122103111233 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.all_sources

<a id="canonical-0232322210231213-1331210222033021-1312001113030130-2111300303102332-2121333301322131-2302331100020102-1210101111321300-3032213222012221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all sources.

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

<a id="canonical-1002111220202201-0132221302023313-1002332013221313-0331303200212203-1113210011332330-3232223310130233-3131123113101323-0200310310233100"></a>

## Direct properties — all_sources / 122103111233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201220113322132-1202123011123313-3200032221230102-0033023201323001-1130102210302101-2111230231130231-3113321013100231-3121202331212333"></a>

## Next pages — all_sources / 122103111233 / 4

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-1300121121133123-1331331113003123-2123002310303113-2213203323123203-2322233201100012-2333000202102110-2233223130210112-3013002000222300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111333312011231-2222232133021202-2301302121213223-3323122010302021-0320303313201200-3002320022102302-1323131101322312-3100211101022211"></a>

## rule_list.rules.dst_asn_list — dst_asn_list / 032033013121 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_asn_list

<a id="canonical-0303030201113121-1103123211313320-2021122332033320-3110030310233020-3022311312010321-0002213022211101-2320012313111133-0132122022111120"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-2223300303222101-1313012123312020-0031133203303301-0323120112102310-1131113012202201-2113333330220110-2111000233222123-2332000110231112"></a>

## Direct properties — dst_asn_list / 032033013121 / 3

<a id="canonical-2003230122023011-2121213210123031-1113031023122221-2023230133300111-3123122013003332-2131330132000112-0031101033020121-1131233311303333"></a>

<a id="canonical-3133011023111312-3232313130332200-1103002332030103-3211123211132033-3212032120310100-2020001311313031-3021313230122002-3221102020011200"></a>

## as_numbers property — dst_asn_list / 032033013121 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-1132000110120000-1201210113013302-1320313210023011-2133211013031020-0311220021031321-1213011020200210-2022313103310310-1123213300121203"></a>

## Next pages — dst_asn_list / 032033013121 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0333101311200103-1201232200123231-1312310331322313-0330220231113033-1023301323021201-2233333012002321-0200302211110200-2303223023100212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123311302223000-2103010031232022-3231011021220202-3030031310120231-0301232123102111-2123230332103221-3302133331322100-2021212021310113"></a>

## rule_list.rules.dst_asn_set — dst_asn_set / 321233222230 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_asn_set

<a id="canonical-3111311212203313-3311102310322000-2032122321103130-2302311032021331-1101223222020021-1221122022032002-0310111021133330-3110122110213322"></a>

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

<a id="canonical-3012302200103203-1121212122021230-1200032111231321-2001003310312110-2110300133012120-3330003111100202-2002022023202010-3023230103313021"></a>

## Direct properties — dst_asn_set / 321233222230 / 3

<a id="canonical-1101221330120033-1333020323112103-1323002010222002-2102032223122320-1132110312103222-2200012123201213-3112232111211030-1210132333120311"></a>

<a id="canonical-1131333211321331-3330100303003312-0111011022023213-3302022312032101-3233320120103003-3220032321213210-1131022212231220-1310030012010012"></a>

## name property — dst_asn_set / 321233222230 / 4

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

<a id="canonical-2120023233120332-2313201203332121-3221303013000302-3313110123232132-3222003033113030-3112012021003122-3003202012203030-2221013220020202"></a>

<a id="canonical-0320212203211003-3231230230023201-1210211231101210-1133323011332230-0011031120030110-3313021312131303-1031113210232322-2303021000222230"></a>

## namespace property — dst_asn_set / 321233222230 / 5

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

<a id="canonical-1122030031232123-1012003211300232-1033023221232233-3000333032030121-1133000013010303-1323213322122323-3201110332302110-0211000113131201"></a>

<a id="canonical-0320321213313312-0310132030011333-1123311031103210-0012303033203302-0032002023231002-2031013000220010-3223111110200232-1300211113312203"></a>

## tenant property — dst_asn_set / 321233222230 / 6

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

<a id="canonical-1002313120213232-0231203232210313-0203120132210001-2020003013102223-0220122303101333-0232312203022031-0321003110231213-3102003330230303"></a>

## Next pages — dst_asn_set / 321233222230 / 7

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3313033010201213-1303212323211203-3223031032231111-1000013301232331-2032030123132131-2222301002313201-2232022003332321-3221120203213211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032200220230332-3020330300211231-0211010322121110-0203222331313331-2011331333011012-2020100230201323-0331132131022121-0223332012030013"></a>

## rule_list.rules.dst_ip_prefix_set — dst_ip_prefix_set / 133000230312 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_ip_prefix_set

<a id="canonical-1103131002311232-1320220001232303-1312231231010331-3322012200212232-3223223122313132-2032031102120010-3322321001210302-2101003302223312"></a>

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

<a id="canonical-0123123110233000-0232230203123101-0002033101020121-3221101023220021-1032331312023231-1120311100123030-0032322023122332-0212103113203223"></a>

## Direct properties — dst_ip_prefix_set / 133000230312 / 3

<a id="canonical-2322102303003010-2022232312113301-1331313022032220-1131321010233223-1300303320133011-3231313011201333-0020111332302310-1132121021332323"></a>

<a id="canonical-3221210000230031-1033010021200020-3022203003021233-3020012202001133-3130311021100022-3203310203323012-0101003330001311-1222031223013030"></a>

## name property — dst_ip_prefix_set / 133000230312 / 4

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

<a id="canonical-1330211321130102-3133131322230002-3231200232110233-1010310030122001-1312311213001131-2211112210021303-2333021103132022-0020101222301320"></a>

<a id="canonical-3323203213101023-2100221332121302-0003013220332233-0033033123330330-3223031112322301-1110122113331333-0033003112020100-0101320010003102"></a>

## namespace property — dst_ip_prefix_set / 133000230312 / 5

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

<a id="canonical-2232311300100111-3031232132100333-1033202333201013-3012320212003300-0311313133332212-0202323320231222-3120131210002123-0221113222323012"></a>

<a id="canonical-3312203333132120-2113132202211203-0112211003311013-0000033220012333-1130313100133131-3300202211333213-1011130021132323-3000121133023102"></a>

## tenant property — dst_ip_prefix_set / 133000230312 / 6

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

<a id="canonical-2212202332020310-3013213102221331-1113323033013100-3233200033010003-0120303130031023-3112002023332302-0130330310321300-2233123333002012"></a>

## Next pages — dst_ip_prefix_set / 133000230312 / 7

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0321111022030032-2322302230031301-2330333002000223-1223111132313113-0013121312101130-2001323322322232-3010000201300030-2311113000302123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201220230001003-0003311330133011-2212332322202223-3302302121102110-0221223011213033-1302333222113022-3100103022321202-2313003131220033"></a>

## rule_list.rules.dst_label_selector — dst_label_selector / 211320031231 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_label_selector

<a id="canonical-0201201320303023-3100103321201122-1012310023331330-2231213211022200-1110312301211321-1002000012322203-1121022022022300-2000320212022123"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-1330021300223030-3213302133121032-1030121011211210-2030103103013223-0200003233113001-0221132213332102-0200220132211000-3122313100300223"></a>

## Direct properties — dst_label_selector / 211320031231 / 3

<a id="canonical-2312310323122132-1130022111033331-1322302131230201-2021323031003133-3010223222000120-1001330130221100-2210212100300220-0132211111233031"></a>

<a id="canonical-2231003312111000-0301120311320102-1020301000001122-3031123232311133-1020122322120122-3112013312030120-1113201132331312-1332211132322122"></a>

## expressions property — dst_label_selector / 211320031231 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2233220302102232-1032223100113130-3201310022122222-3301221132102012-3103132023323220-2103001232003313-2231211110032233-2130333112313132"></a>

## Next pages — dst_label_selector / 211320031231 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-1302021011333221-1202321123122020-2132230322001012-2021003030013130-3003132133103002-0232330130210222-0130211121032322-1300232013202233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032023123012220-1031321303213132-1230121323331333-3010123110011210-2131231211232201-1203323213022203-2203233010333110-0111230013330231"></a>

## rule_list.rules.dst_prefix_list — dst_prefix_list / 030033022313 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_prefix_list

<a id="canonical-0223030022201003-2200200200212311-1002332201330230-0333233231033333-1022231312203330-3210013000132011-1110132233323131-2212022301232230"></a>

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

<a id="canonical-0320013323322303-3102103301020002-2030131011202120-0002203302230322-2301223031332121-1112102113122112-0230333031300211-0211013323232133"></a>

## Direct properties — dst_prefix_list / 030033022313 / 3

<a id="canonical-3011333011112000-0301131132120001-2232312232221323-0233023333111321-3203230023021320-1300111201302213-1122210113213333-0012322011223032"></a>

<a id="canonical-3120302223311121-1213213120222120-0132320123110022-2000222210001231-1203323031123303-2330212232011312-2022002312303101-2101202320231210"></a>

## prefixes property — dst_prefix_list / 030033022313 / 4

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

<a id="canonical-1011112012223223-2333332021002012-1321300021021300-3111301123113332-0031223121303203-0331201123123022-1133320113132123-2030202312303203"></a>

## Next pages — dst_prefix_list / 030033022313 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121102111033322-1122313003321003-1103103211332003-2313002003100231-2330113213211311-0331021231232130-1130032132002132-0133133032223331"></a>

## rule_list.rules.http_list — http_list / 332231112333 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.http_list

<a id="canonical-0313322022003312-1011021132002233-2233113003012002-3000231031132201-2330222203212330-2232331110203021-0132310310120212-2113011020012122"></a>

Type: `"single"`. Computed.

URLListType.

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

<a id="canonical-1132020222130112-1330303321311233-1122022330303221-0322303003210121-3322022213003023-0310003333220202-1221103031002313-1021000203211201"></a>

## Direct properties — http_list / 332231112333 / 3

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1003321321330001-0303131032030213-0232323233130012-2203312203321133-1321010000012031-2111113211232030-2120302312102011-3301121212202203): complete subsection reference.

<a id="canonical-3200030231311020-0331212232003022-0311111000302230-2202011013012232-2010211222231223-0011033333301333-2333330033310313-1212312230302321"></a>

## Next pages — http_list / 332231112333 / 4

- [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1003321321330001-0303131032030213-0232323233130012-2203312203321133-1321010000012031-2111113211232030-2120302312102011-3301121212202203)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-1003321321330001-0303131032030213-0232323233130012-2203312203321133-1321010000012031-2111113211232030-2120302312102011-3301121212202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212110001203133-2023222031112010-2111200201032113-2012322222330000-0103033211023030-1231032322332300-3212110010013321-1113130011303211"></a>

## rule_list.rules.http_list.http_list — http_list / 023103121022 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211)
- rule_list.rules.http_list.http_list

<a id="canonical-2332313220202000-3312100330020212-3112003002100213-3103003112132222-3313133221001310-3112333121203012-1230330210012213-1330013121333122"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

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

<a id="canonical-0311021133333022-0030100223233131-2012112113302232-0321013213331101-3133231123111122-1121320032202311-2002131013300333-0200231220200132"></a>

## Direct properties — http_list / 023103121022 / 3

- [any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3333022102011120-3232120200301210-1001022300001130-0322221310302030-1333303212210200-0213321232313103-3223113020301220-2230030322330313): complete subsection reference.

<a id="canonical-3111112333130221-1303131331030032-1201221220120131-0313200330220210-2122220120012211-1012121203203220-2133211012103330-3222132032203203"></a>

<a id="canonical-1311132312002033-0221000302223313-3102103122220133-0302220010103200-0121000312210120-0123203100232113-2202121203131221-3221323011313103"></a>

## exact_value property — http_list / 023103121022 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3202020112112301-0130301032331013-3223031020012102-2221222111000222-3131102031122003-3110003202330003-0232201311303011-2001222202320000"></a>

<a id="canonical-3103320202020133-3130101201130012-1212111321321222-1020132320210133-3233012101120311-0002231030000321-3321001301100030-0112001312213120"></a>

## path_exact_value property — http_list / 023103121022 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0203302233233303-3101333323220301-3122101122123001-2131332212312103-2110111331131123-0230030030203131-1010302022300202-0312302103121033"></a>

<a id="canonical-2212321210201002-2231032130303312-1011111002302210-0001122200003213-2203332023020030-0003221121330021-3133301321223201-0320022002332002"></a>

## path_prefix_value property — http_list / 023103121022 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2101002331233101-2121123312030213-1212133022102300-1003221121320210-1313231012130201-1122023123030130-2330231333003100-2022021211212023"></a>

<a id="canonical-0010100313130122-0011213020133202-2100222011023230-2323101221230222-3231212010230110-1213201123323133-2302332320120311-1022120232100112"></a>

## path_regex_value property — http_list / 023103121022 / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2231033013310301-1300010312031130-3103331320013211-0332101033001032-2213312312232303-2301232003320300-0210113011202003-2122311321121102"></a>

<a id="canonical-0112133112311310-2333020030312123-2303201023123210-2023310121101020-1202323222202322-1033313020032210-3022303122123031-1203131111313122"></a>

## regex_value property — http_list / 023103121022 / 8

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2222033202030201-0103312203123121-1130330031223122-3231321110323103-3201330223123021-1322032300300130-0000130001102111-1012212032120010"></a>

<a id="canonical-2112310103131003-2031203102031330-0130020110310232-1203320301321003-2202201301112301-1300120030213032-0110230312230023-2220300113201031"></a>

## suffix_value property — http_list / 023103121022 / 9

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0303232121130121-0311013322212110-2113111002120100-3020222021033001-0201022331212330-3113011223012333-0123132023332200-0203220211130323"></a>

## Next pages — http_list / 023103121022 / 10

- [rule_list.rules.http_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3333022102011120-3232120200301210-1001022300001130-0322221310302030-1333303212210200-0213321232313103-3223113020301220-2230030322330313)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3333022102011120-3232120200301210-1001022300001130-0322221310302030-1333303212210200-0213321232313103-3223113020301220-2230030322330313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033021123123231-2022001130333301-2001322123132213-1223221001221121-2231230232231331-1012330311113031-2221112331320031-3002200030330130"></a>

## rule_list.rules.http_list.http_list.any_path — any_path / 002111201123 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211)
- [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1003321321330001-0303131032030213-0232323233130012-2203312203321133-1321010000012031-2111113211232030-2120302312102011-3301121212202203)
- rule_list.rules.http_list.http_list.any_path

<a id="canonical-1000221323210102-2001232000102021-2022332131133200-3300211133331011-3312130111202030-1023022200210103-0011112210031123-3122032033202131"></a>

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

<a id="canonical-1122332003233303-2102210301003221-1310011033120321-1030132230221101-0100103332233222-0132310121230032-1002023031233301-3321222102331133"></a>

## Direct properties — any_path / 002111201123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033010100310012-0302221123200233-3010033333220130-0012100311031112-3230310112023112-0021233330200321-2311333232120210-3233123300003223"></a>

## Next pages — any_path / 002111201123 / 4

- [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1003321321330001-0303131032030213-0232323233130012-2203312203321133-1321010000012031-2111113211232030-2120302312102011-3301121212202203)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-2212103130223310-3330130302022032-0203030031113130-1003332132223122-1023111013210012-0030311012030221-2111023131032223-2232213101213000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133302302110032-1131133110032100-2202021133233311-0333223332013003-3333230203002130-3220003122111300-3303022220321333-2122113312102013"></a>

## rule_list.rules.ip_prefix_set — ip_prefix_set / 133032311113 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.ip_prefix_set

<a id="canonical-2221331210221003-2132123310203232-0200013110031023-0002101033031131-1221132113133023-3022000030320031-3023000332220103-2022031021310013"></a>

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

<a id="canonical-3133032233303201-0013210203233321-2331003133011001-2002210231312230-1210133110330022-2112202112100310-1320001203110213-1013013212020133"></a>

## Direct properties — ip_prefix_set / 133032311113 / 3

<a id="canonical-1230000211031033-1020001022123233-1012031230103223-2031203301201211-0020233121223212-3221202220022001-1230002323220221-2003330022033332"></a>

<a id="canonical-2032220022221302-3023121333222332-0031210330100221-1321221221323311-0110030330130000-2030330131031110-3331232103312103-1202302323133320"></a>

## name property — ip_prefix_set / 133032311113 / 4

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

<a id="canonical-3231013322121303-3211023310222012-1212300102133111-0333313213331301-0121132232133202-2103321011331112-3123300233201132-1211030321020113"></a>

<a id="canonical-0103112020113103-0311011311121300-2120101213001232-0231221310320231-0032333122031202-1002000021233012-0122333100322210-0333330020330302"></a>

## namespace property — ip_prefix_set / 133032311113 / 5

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

<a id="canonical-0123201102220123-2130032032012111-3233022021022102-1312131023210020-3003132201332311-3012220202301203-1202330010320311-3320333321030031"></a>

<a id="canonical-3122011212302110-2021231322300201-1312011123333122-2211113003010300-0121323202233312-3333002002102211-2313311120001220-3120321331020332"></a>

## tenant property — ip_prefix_set / 133032311113 / 6

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

<a id="canonical-0230221203301001-2232110003222101-0112222210330303-2321131113022233-0220202211303103-1021021223232203-3300323311321332-2031113332123201"></a>

## Next pages — ip_prefix_set / 133032311113 / 7

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-0121013023233311-1203120030030033-1332300301332300-0321210320000300-1112203010322331-1133332312311203-2212003212000211-0102321303223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332030011220132-2220301233333223-1113033303022122-1111131021333200-2122121331013020-3111122123213332-3211010301031101-3333003200203322"></a>

## rule_list.rules.label_selector — label_selector / 111030233212 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.label_selector

<a id="canonical-1232121312200123-3023221110001301-1102132213030122-0023311123012201-0231100200223121-2233212231230032-0320102101301123-1033320303203202"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-0130311023031133-2230002213322313-3301201200130321-1021100110003310-3212302300312213-3311301101030212-0301031222132013-0233330213100131"></a>

## Direct properties — label_selector / 111030233212 / 3

<a id="canonical-3330332011201323-1221230021010310-2332112102112030-1123222232210223-2121002033103331-3011232302122230-0013310022221131-1231021130120120"></a>

<a id="canonical-1220001001303333-2131033000332210-1132011323130030-2100022223132100-0323113230203133-2113121011002120-2212233300023303-2131212132222120"></a>

## expressions property — label_selector / 111030233212 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3103130021013002-3301331111120032-0022212000113322-0200111120100211-2223201301002211-3330302003313133-1032300231000031-1001311133301020"></a>

## Next pages — label_selector / 111030233212 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3223102201230333-0133323300113202-0020231011110213-1103132313011100-2303122321311003-3120230013203320-0203112302310020-0011101020213312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131221002102321-1003231000230030-2331302322320311-3121101011120300-2103012031231210-1102111012000211-0200220321311003-1301100030130030"></a>

## rule_list.rules.metadata — metadata / 320033231300 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.metadata

<a id="canonical-1210001310130000-3210012311233133-3203210131221230-2122331100231103-1232002322201110-3312010130102220-0001101001020211-0112012302211310"></a>

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

<a id="canonical-0300022301021012-3222200131120112-1331103330102222-1132310331310221-3223310231012033-2200210313123330-0322300020111312-0122201003230311"></a>

## Direct properties — metadata / 320033231300 / 3

<a id="canonical-3022123321130201-0032231300012322-3003302202102003-3300010122103121-0311233211231032-2320102111100012-0032210002131133-3003213030131220"></a>

<a id="canonical-3331100003023230-0303011031302102-3121301112223330-2321022031213312-3020001111323033-0221302321312330-2313023032103201-3110001121122100"></a>

## description_spec property — metadata / 320033231300 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2200322132332303-2323301312032032-1123321320220013-1031301233201230-2020131130311130-2013200130002311-1111312110323011-2312120023120021"></a>

<a id="canonical-2113102233123033-0112001220131003-1223121011303232-1320100033100203-3111120211110132-3220012012210333-3312021033323330-0213303112130123"></a>

## name property — metadata / 320033231300 / 5

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

<a id="canonical-1320022320320122-1013221030302310-0313233313322021-2201002221102101-3312123321211311-3331013213122111-0030013132121303-2023003121100010"></a>

## Next pages — metadata / 320033231300 / 6

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3322110302231022-2313001321321311-0011112313123332-0331303200302032-1300033123310301-3131201202203220-0023301312310322-2323311123323123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313113310300110-2333111331032101-0201112110133000-2113021133123133-1023031310231103-0123303323020111-3100101222033301-2200213001100231"></a>

## rule_list.rules.no_http_connect_port — no_http_connect_port / 213303002212 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.no_http_connect_port

<a id="canonical-1100003212003111-2323300223301233-1031230030301310-3212211233102103-3011002303203300-0112210230132020-3010220233131302-0233121302100032"></a>

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

<a id="canonical-2103020322312312-2331231302031213-3113201221111210-2233003100132230-0121021212332130-2020301223231301-1121222021030231-2230031023102113"></a>

## Direct properties — no_http_connect_port / 213303002212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313121322023012-0131231002132021-3222133300212311-3213010133113130-2103232031110332-2320121210201223-0201212032222130-1200222101210213"></a>

## Next pages — no_http_connect_port / 213303002212 / 4

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3303223031202122-2130021033020311-3000130130333300-0233223101103033-0220002000332013-0100230323032120-1303100210111331-2311331133310002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112232023130103-1201230121101032-3100131011331313-1113233211120221-2120232311113021-0030302201223320-3332332230311222-1132201033310231"></a>

## rule_list.rules.port_matcher — port_matcher / 010333230122 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.port_matcher

<a id="canonical-0010111210103020-0201131031233012-3022203303022110-3320030321330122-1300200033301312-0013313002212230-0103213311310311-0302211222100201"></a>

Type: `"single"`. Computed.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

Upstream description:

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

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

<a id="canonical-3303301220132130-1010000201003311-2112221033023011-1213231022013102-1013211012322133-0110202301212120-2112310211111021-0302203303112301"></a>

## Direct properties — port_matcher / 010333230122 / 3

<a id="canonical-0201020021223012-2301010332213221-0102121212210210-2120330110231032-2003233113322032-0121223011011120-1012200323203231-3000132232033233"></a>

<a id="canonical-2210333021021302-2111131010122323-2110012111131302-2301110300123023-1033202231110210-3231011000330332-1230001213000132-0320102011020223"></a>

## invert_matcher property — port_matcher / 010333230122 / 4

Type: `"bool"`. Computed.

Invert Port Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-2310013313132102-1102331133322010-0223312113301331-1113002323122333-1222023100020223-1211001012200122-2300023031310131-1030222200333210"></a>

<a id="canonical-0200213032030303-3003133233132003-1031031302023131-2210332201313103-2203203321303213-1132310313330012-3231230231222221-1310023301001302"></a>

## ports property — port_matcher / 010333230122 / 5

Type: `["list", "string"]`. Computed.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Upstream description:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-". The start and end values are considered to be part of the range.

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
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3103013130232122-3203310231123130-3312012302321200-2233113022303332-3033230010302031-3201223020230211-0210120230233000-0210222022323002"></a>

## Next pages — port_matcher / 010333230122 / 6

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3212130112120111-0303230130202111-3022100232223233-3002201320311201-2201332023302332-2211231010010220-3022300313111010-2321121102331231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133133232023122-0333101020300313-3132100021132002-0310011111123132-1120021102302223-3223303313323112-2113331112301111-1012022332331302"></a>

## rule_list.rules.prefix_list — prefix_list / 210310231010 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.prefix_list

<a id="canonical-0011311212302311-1032132100233321-3030201131123332-0303230101110013-2301332123200111-2303313331112220-2231020112312331-1132213222233333"></a>

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

<a id="canonical-1033000211012101-0230002121123103-3030033120223102-3300330201012213-1011120220212033-0003021001002212-1331003313230233-1010231332023011"></a>

## Direct properties — prefix_list / 210310231010 / 3

<a id="canonical-0003033122013012-1312321111102100-2111321330300023-0111213310330331-1233012132332232-1300222012221331-0200000311031133-3133120133121102"></a>

<a id="canonical-2031103100003321-2322233313313020-2021123202330231-1332130110120202-3131203231131022-2122032321133102-2001333123231230-3032233201032113"></a>

## prefixes property — prefix_list / 210310231010 / 4

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

<a id="canonical-2123112123103131-0201132210130213-2221000302020113-1311333003112200-2232023210031013-1233003310333030-2330313303232023-3223210211033200"></a>

## Next pages — prefix_list / 210310231010 / 5

- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3130031113132122-1313310010332323-3012313301202331-1020032123210331-2033230333200121-0311201313100230-3313331323102333-1023320301301321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011320012013222-2313333310203131-1313102232100323-1100032333200000-1023032110032030-2003223001333012-1033131302113202-0233010211323213"></a>

## rule_list.rules.tls_list — tls_list / 212320333302 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.tls_list

<a id="canonical-0310303133133202-3123020233103333-0330303332213232-3102210130001301-0203320323021232-1333013233022222-2023200032210202-2023003003301330"></a>

Type: `"single"`. Computed.

DomainListType.

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

<a id="canonical-1021133311202030-1333123231303013-3230103122210001-1023202210212032-1123003312133122-3031300322023330-3300200012011000-3232203320213110"></a>

## Direct properties — tls_list / 212320333302 / 3

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3010320131320220-1022001113232320-3102131330323313-0200102002102232-3133121213313032-1023122302010301-3211303001020120-1312123113010233): complete subsection reference.

<a id="canonical-0003230202233211-3020230233013311-2112303322103011-1321110223021213-2320223003122121-1113221322013122-0310312221301321-0013333030001010"></a>

## Next pages — tls_list / 212320333302 / 4

- [rule_list.rules.tls_list.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3010320131320220-1022001113232320-3102131330323313-0200102002102232-3133121213313032-1023122302010301-3211303001020120-1312123113010233)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)

<a id="canonical-3010320131320220-1022001113232320-3102131330323313-0200102002102232-3133121213313032-1023122302010301-3211303001020120-1312123113010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203301102131023-1302012010213030-0001020131200110-2303332111323120-1001311133321012-3021122010033011-3300200320222001-0001101020113131"></a>

## rule_list.rules.tls_list.tls_list — tls_list / 311111003333 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [rule_list.rules.tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3130031113132122-1313310010332323-3012313301202331-1020032123210331-2033230333200121-0311201313100230-3313331323102333-1023320301301321)
- rule_list.rules.tls_list.tls_list

<a id="canonical-2221221222120112-0230021231220123-3301300121011210-0203100102003010-2231102212311222-2220120122210122-3001211332213332-3022002102200321"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

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

<a id="canonical-2220221100312100-2122121230220012-2210320102211001-0312011221003113-1101313100322213-2123203011131032-2223002131102310-1331103230113132"></a>

## Direct properties — tls_list / 311111003333 / 3

<a id="canonical-1231232302320130-2023322210002011-2320030032332032-1011123111333213-0333310013131311-3210100331302133-3211133103003322-3331331100000321"></a>

<a id="canonical-0023022121300001-2020211302213222-1300213003311201-2100100133011333-2230012103132120-0301130112000023-1002331011201221-1311002220322200"></a>

## exact_value property — tls_list / 311111003333 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3131121023120210-0013333111312111-2322322202212121-0022122110003230-1101133001020130-1131012220221333-1000012033303111-0100311333222213"></a>
