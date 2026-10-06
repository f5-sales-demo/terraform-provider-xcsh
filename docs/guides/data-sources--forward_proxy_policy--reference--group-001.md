---
page_title: "xcsh_forward_proxy_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy reference."
---

# xcsh_forward_proxy_policy reference

<a id="canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- Property reference

<a id="canonical-1121321011002221-3211211030322322-1021021131023132-0302021203323001-3033332321101211-2202301333001111-0000132330301123-0323103332021102"></a>

### Direct properties for `xcsh_forward_proxy_policy`

- [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2013011230133310-0110112322303323-1102232112233312-0332123113232113-1131213101211320-3211310122131000-0030202322111201-0103202332111313): complete subsection reference.

- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301): complete subsection reference.

<a id="canonical-3332302313200021-2310100332200010-0003311112223112-3032131130222311-0210321011231033-3131222322320101-3220221302233301-1323132203301203"></a>

<a id="canonical-2232003132103002-0033132202132331-3223121023131222-1031321322210021-0220223203313302-0110121000002222-2230301103013221-3312030223211001"></a>

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

- [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0301302122312100-0303121321221121-2233002013020312-1211203203103312-1033012321032120-1121133311231131-1221332120300202-1103321303300103): complete subsection reference.

- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032): complete subsection reference.

<a id="canonical-2010032211002230-0010030311032231-2300320300000130-0113012111022332-1030103303303031-0120303012110302-3113302110111031-0022322132201223"></a>

<a id="canonical-0232103100302002-2100133122101301-2031300011200331-3301022202302102-0131311132301322-0301001322130321-2233133013102100-2000320202032020"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the ForwardProxyPolicy.

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

- [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0301032221101120-3103200113001310-3223232331220200-1332133303133312-1323311130131310-1001333230111322-0230221210331302-3200332121101102): complete subsection reference.

<a id="canonical-1312130103301321-2320312332000213-3303230110120213-3222232232212023-3121301203230102-1002322311330003-0203122233012303-1321210131201003"></a>

<a id="canonical-1311001322312313-1201021101133311-2032110232022323-2222021112300020-3103102132113112-2030201102030233-2023213010332020-3221332103110303"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1012001101201222-1012022131203020-3203200201202223-2103323221221102-1022123030101113-2031111001231000-2330233231220220-2220021330323000"></a>

<a id="canonical-2101030202210210-3123031012122303-3321332031300213-0303000231011002-1001322230012323-1331022032031110-1303333101120203-1020233111131323"></a>

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

<a id="canonical-2001101330212233-0303013320213032-2112220012212021-1001211111022003-2001021311030210-0131113130000211-2311232301201112-2111111332203323"></a>

<a id="canonical-3301132300131120-0221030201031211-2123031221023020-3203121132300313-0132033330001321-1230211201100113-0130020111011031-0202010003032032"></a>

#### `name` property

Type: `"string"`. Required.

Name of the ForwardProxyPolicy.

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

<a id="canonical-2023100322011231-0020021321002032-0301210102032303-0323320022020223-3302000132022232-3100113330130333-3110103303113211-1102312302013000"></a>

<a id="canonical-3203022102320030-1312211212202233-3231112303320323-0121120112322030-1002102232221220-3230020013213203-2323113303233002-1202201111022300"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the ForwardProxyPolicy exists.

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

- [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3003120302033030-0022133033303133-1021310202213331-0113203001333310-0100202002013200-2202322230231323-0010302133310002-3313303112221211): complete subsection reference.

- [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1032301033033313-1101300232301332-0110322320310030-1032321212200203-0212320010031030-1100000123210022-0031113133102112-0301111301311103): complete subsection reference.

- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232): complete subsection reference.

<a id="canonical-0011200221032021-2212022131333200-2111211231002313-0331221131202013-3330213003102101-0211303003321203-3013210321101203-3322322233313211"></a>

### All schema paths for `xcsh_forward_proxy_policy`

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
| `network_connector.tenant` | [network_connector.tenant](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2300322032101212-2102011102022232-1123320022202301-0232110132313232-2102300220002003-2112303223102033-2111222121323110-2233110200123221) |
| `proxy_label_selector` | [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3320221102222111-2312332233123123-0022022013122310-0231213333221113-3301223112111020-0120030021112201-1210012120013200-1210110111112213) |
| `proxy_label_selector.expressions` | [proxy_label_selector.expressions](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1300023203010112-2003122212123121-0320231300323302-0312320001120013-2122100132003011-2212000122231010-3133130330220312-0311103002312001) |
| `rule_list` | [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3300003111010003-1022200013220321-2120312030000021-0322110310133010-2300202200330301-2102222012123212-3322012103010130-2332110210313311) |
| `rule_list.rules` | [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3100302032212022-2021123222100320-3213331000002321-2003220103311020-0101020103332002-1313332302120213-0321120011110223-2110233133303132) |
| `rule_list.rules.action` | [rule_list.rules.action](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1131000023102020-1231010222000020-3001331200220123-0203303331330312-1031123113111032-1000301210102321-2102313200023203-1200130222231333) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3202113122323120-0131313120111220-3331020322313230-3012320022131230-3333301200321111-2210213110123333-1330333230301321-2002120302021203) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0232322210231213-1331210222033021-1312001113030130-2111300303102332-2121333301322131-2302331100020102-1210101111321300-3032213222012221) |
| `rule_list.rules.dst_asn_list` | [rule_list.rules.dst_asn_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0303030201113121-1103123211313320-2021122332033320-3110030310233020-3022311312010321-0002213022211101-2320012313111133-0132122022111120) |
| `rule_list.rules.dst_asn_list.as_numbers` | [rule_list.rules.dst_asn_list.as_numbers](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2003230122023011-2121213210123031-1113031023122221-2023230133300111-3123122013003332-2131330132000112-0031101033020121-1131233311303333) |
| `rule_list.rules.dst_asn_set` | [rule_list.rules.dst_asn_set](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3111311212203313-3311102310322000-2032122321103130-2302311032021331-1101223222020021-1221122022032002-0310111021133330-3110122110213322) |
| `rule_list.rules.dst_asn_set.name` | [rule_list.rules.dst_asn_set.name](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1101221330120033-1333020323112103-1323002010222002-2102032223122320-1132110312103222-2200012123201213-3112232111211030-1210132333120311) |
| `rule_list.rules.dst_asn_set.namespace` | [rule_list.rules.dst_asn_set.namespace](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2120023233120332-2313201203332121-3221303013000302-3313110123232132-3222003033113030-3112012021003122-3003202012203030-2221013220020202) |
| `rule_list.rules.dst_asn_set.tenant` | [rule_list.rules.dst_asn_set.tenant](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1122030031232123-1012003211300232-1033023221232233-3000333032030121-1133000013010303-1323213322122323-3201110332302110-0211000113131201) |
| `rule_list.rules.dst_ip_prefix_set` | [rule_list.rules.dst_ip_prefix_set](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1103131002311232-1320220001232303-1312231231010331-3322012200212232-3223223122313132-2032031102120010-3322321001210302-2101003302223312) |
| `rule_list.rules.dst_ip_prefix_set.name` | [rule_list.rules.dst_ip_prefix_set.name](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2322102303003010-2022232312113301-1331313022032220-1131321010233223-1300303320133011-3231313011201333-0020111332302310-1132121021332323) |
| `rule_list.rules.dst_ip_prefix_set.namespace` | [rule_list.rules.dst_ip_prefix_set.namespace](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1330211321130102-3133131322230002-3231200232110233-1010310030122001-1312311213001131-2211112210021303-2333021103132022-0020101222301320) |
| `rule_list.rules.dst_ip_prefix_set.tenant` | [rule_list.rules.dst_ip_prefix_set.tenant](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2232311300100111-3031232132100333-1033202333201013-3012320212003300-0311313133332212-0202323320231222-3120131210002123-0221113222323012) |
| `rule_list.rules.dst_label_selector` | [rule_list.rules.dst_label_selector](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0201201320303023-3100103321201122-1012310023331330-2231213211022200-1110312301211321-1002000012322203-1121022022022300-2000320212022123) |
| `rule_list.rules.dst_label_selector.expressions` | [rule_list.rules.dst_label_selector.expressions](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2312310323122132-1130022111033331-1322302131230201-2021323031003133-3010223222000120-1001330130221100-2210212100300220-0132211111233031) |
| `rule_list.rules.dst_prefix_list` | [rule_list.rules.dst_prefix_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0223030022201003-2200200200212311-1002332201330230-0333233231033333-1022231312203330-3210013000132011-1110132233323131-2212022301232230) |
| `rule_list.rules.dst_prefix_list.prefixes` | [rule_list.rules.dst_prefix_list.prefixes](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3011333011112000-0301131132120001-2232312232221323-0233023333111321-3203230023021320-1300111201302213-1122210113213333-0012322011223032) |
| `rule_list.rules.http_list` | [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0313322022003312-1011021132002233-2233113003012002-3000231031132201-2330222203212330-2232331110203021-0132310310120212-2113011020012122) |
| `rule_list.rules.http_list.http_list` | [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2332313220202000-3312100330020212-3112003002100213-3103003112132222-3313133221001310-3112333121203012-1230330210012213-1330013121333122) |
| `rule_list.rules.http_list.http_list.any_path` | [rule_list.rules.http_list.http_list.any_path](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1000221323210102-2001232000102021-2022332131133200-3300211133331011-3312130111202030-1023022200210103-0011112210031123-3122032033202131) |
| `rule_list.rules.http_list.http_list.exact_value` | [rule_list.rules.http_list.http_list.exact_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3111112333130221-1303131331030032-1201221220120131-0313200330220210-2122220120012211-1012121203203220-2133211012103330-3222132032203203) |
| `rule_list.rules.http_list.http_list.path_exact_value` | [rule_list.rules.http_list.http_list.path_exact_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3202020112112301-0130301032331013-3223031020012102-2221222111000222-3131102031122003-3110003202330003-0232201311303011-2001222202320000) |
| `rule_list.rules.http_list.http_list.path_prefix_value` | [rule_list.rules.http_list.http_list.path_prefix_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0203302233233303-3101333323220301-3122101122123001-2131332212312103-2110111331131123-0230030030203131-1010302022300202-0312302103121033) |
| `rule_list.rules.http_list.http_list.path_regex_value` | [rule_list.rules.http_list.http_list.path_regex_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2101002331233101-2121123312030213-1212133022102300-1003221121320210-1313231012130201-1122023123030130-2330231333003100-2022021211212023) |
| `rule_list.rules.http_list.http_list.regex_value` | [rule_list.rules.http_list.http_list.regex_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2231033013310301-1300010312031130-3103331320013211-0332101033001032-2213312312232303-2301232003320300-0210113011202003-2122311321121102) |
| `rule_list.rules.http_list.http_list.suffix_value` | [rule_list.rules.http_list.http_list.suffix_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2222033202030201-0103312203123121-1130330031223122-3231321110323103-3201330223123021-1322032300300130-0000130001102111-1012212032120010) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2221331210221003-2132123310203232-0200013110031023-0002101033031131-1221132113133023-3022000030320031-3023000332220103-2022031021310013) |
| `rule_list.rules.ip_prefix_set.name` | [rule_list.rules.ip_prefix_set.name](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1230000211031033-1020001022123233-1012031230103223-2031203301201211-0020233121223212-3221202220022001-1230002323220221-2003330022033332) |
| `rule_list.rules.ip_prefix_set.namespace` | [rule_list.rules.ip_prefix_set.namespace](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3231013322121303-3211023310222012-1212300102133111-0333313213331301-0121132232133202-2103321011331112-3123300233201132-1211030321020113) |
| `rule_list.rules.ip_prefix_set.tenant` | [rule_list.rules.ip_prefix_set.tenant](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0123201102220123-2130032032012111-3233022021022102-1312131023210020-3003132201332311-3012220202301203-1202330010320311-3320333321030031) |
| `rule_list.rules.label_selector` | [rule_list.rules.label_selector](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1232121312200123-3023221110001301-1102132213030122-0023311123012201-0231100200223121-2233212231230032-0320102101301123-1033320303203202) |
| `rule_list.rules.label_selector.expressions` | [rule_list.rules.label_selector.expressions](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3330332011201323-1221230021010310-2332112102112030-1123222232210223-2121002033103331-3011232302122230-0013310022221131-1231021130120120) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1210001310130000-3210012311233133-3203210131221230-2122331100231103-1232002322201110-3312010130102220-0001101001020211-0112012302211310) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3022123321130201-0032231300012322-3003302202102003-3300010122103121-0311233211231032-2320102111100012-0032210002131133-3003213030131220) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2200322132332303-2323301312032032-1123321320220013-1031301233201230-2020131130311130-2013200130002311-1111312110323011-2312120023120021) |
| `rule_list.rules.no_http_connect_port` | [rule_list.rules.no_http_connect_port](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1100003212003111-2323300223301233-1031230030301310-3212211233102103-3011002303203300-0112210230132020-3010220233131302-0233121302100032) |
| `rule_list.rules.port_matcher` | [rule_list.rules.port_matcher](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0010111210103020-0201131031233012-3022203303022110-3320030321330122-1300200033301312-0013313002212230-0103213311310311-0302211222100201) |
| `rule_list.rules.port_matcher.invert_matcher` | [rule_list.rules.port_matcher.invert_matcher](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0201020021223012-2301010332213221-0102121212210210-2120330110231032-2003233113322032-0121223011011120-1012200323203231-3000132232033233) |
| `rule_list.rules.port_matcher.ports` | [rule_list.rules.port_matcher.ports](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2310013313132102-1102331133322010-0223312113301331-1113002323122333-1222023100020223-1211001012200122-2300023031310131-1030222200333210) |
| `rule_list.rules.prefix_list` | [rule_list.rules.prefix_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0011311212302311-1032132100233321-3030201131123332-0303230101110013-2301332123200111-2303313331112220-2231020112312331-1132213222233333) |
| `rule_list.rules.prefix_list.prefixes` | [rule_list.rules.prefix_list.prefixes](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0003033122013012-1312321111102100-2111321330300023-0111213310330331-1233012132332232-1300222012221331-0200000311031133-3133120133121102) |
| `rule_list.rules.tls_list` | [rule_list.rules.tls_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0310303133133202-3123020233103333-0330303332213232-3102210130001301-0203320323021232-1333013233022222-2023200032210202-2023003003301330) |
| `rule_list.rules.tls_list.tls_list` | [rule_list.rules.tls_list.tls_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2221221222120112-0230021231220123-3301300121011210-0203100102003010-2231102212311222-2220120122210122-3001211332213332-3022002102200321) |
| `rule_list.rules.tls_list.tls_list.exact_value` | [rule_list.rules.tls_list.tls_list.exact_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1231232302320130-2023322210002011-2320030032332032-1011123111333213-0333310013131311-3210100331302133-3211133103003322-3331331100000321) |
| `rule_list.rules.tls_list.tls_list.regex_value` | [rule_list.rules.tls_list.tls_list.regex_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3131121023120210-0013333111312111-2322322202212121-0022122110003230-1101133001020130-1131012220221333-1000012033303111-0100311333222213) |
| `rule_list.rules.tls_list.tls_list.suffix_value` | [rule_list.rules.tls_list.tls_list.suffix_value](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0323013202223203-1131011033132200-0032201033002102-2011120320133203-1201121313211301-0000013202131003-0302033120113001-3210300002010233) |
| `rule_list.rules.url_category_list` | [rule_list.rules.url_category_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3211012133002021-0031221310323002-1213030330010223-2320300332101200-2021130012032312-2120110323122333-3031003001021132-1121323213231210) |
| `rule_list.rules.url_category_list.url_categories` | [rule_list.rules.url_category_list.url_categories](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2131221100000212-2130132000023302-2131000020200320-1031112200021103-0002233003031312-0131003300330230-0132211220220100-1231300133212100) |

<a id="canonical-2013011230133310-0110112322303323-1102232112233312-0332123113232113-1131213101211320-3211310122131000-0030202322111201-0103202332111313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_all` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- allow_all

<a id="canonical-2220010203033123-1110321130230310-3023220022020123-2201003102133013-3330100232323231-2031021200012211-0323121203331011-2113031103331212"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all, allow\_list, deny\_list, rule\_list\] Enable this option

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

OneOf alternatives in this subsection:

- [allow_all](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2220010203033123-1110321130230310-3023220022020123-2201003102133013-3330100232323231-2031021200012211-0323121203331011-2113031103331212)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1131021322232311-1013030302000130-0100233010200032-3222300200230310-2310111103303033-3003220000012102-3210000312000223-1112313030320123)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1001323031020030-0323022102120112-0100301120113130-2133213212210110-0212222011312233-0032322201311210-2203332013010002-1320313230312301)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3300003111010003-1022200013220321-2120312030000021-0322110310133010-2300202200330301-2102222012123212-3322012103010130-2332110210313311)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- allow_list

<a id="canonical-1131021322232311-1013030302000130-0100233010200032-3222300200230310-2310111103303033-3003220000012102-3210000312000223-1112313030320123"></a>

Type: `"single"`. Computed.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

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

<a id="canonical-1311313221312121-0033321013313332-1233333123212133-2032132131200200-2003003202121310-2102003333101200-0232303310012202-2302302101123310"></a>

### Direct properties for `allow_list`

- [default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0030303121002013-0030230300323302-0331211020221310-2031122301233132-2302323223020300-1111032132203212-2110232020320101-2222103003013032): complete subsection reference.

- [default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3113302312100001-2023231111123213-2001132012213201-3201232313122321-1303331212012032-0000213333223112-0332000331031330-3200202312201302): complete subsection reference.

- [default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0033122023312122-3323023000110110-2233110321213003-3032320103323031-1301233130322022-1321202110021120-2221023211021333-1231012012012121): complete subsection reference.

- [dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1312200232331133-0311010303232332-0221203000300221-3030003302210100-0321011202310111-3201210103311212-3203032110002002-0031020320121203): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2003210313333300-1210322233003121-2111232031231120-3330220111130030-3303312221122102-3032333222111200-1220013320222331-2322233331321302): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0130102331212330-2221222233301120-3102100211102202-2313010000230021-0201201101130132-0323322310303120-3203003310013103-2330103233003213): complete subsection reference.

<a id="canonical-0030303121002013-0030230300323302-0331211020221310-2031122301233132-2302323223020300-1111032132203212-2110232020320101-2222103003013032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.default_action_allow` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.default_action_allow

<a id="canonical-2022122121202213-3223323003033031-1223233302001210-0311322121231330-2220322313000010-2223021012321233-2100310212011003-3130211221101103"></a>

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

<a id="canonical-3113302312100001-2023231111123213-2001132012213201-3201232313122321-1303331212012032-0000213333223112-0332000331031330-3200202312201302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.default_action_deny` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.default_action_deny

<a id="canonical-3030100300212120-1000233011133121-0032302022330313-0130000221313222-2010212200131003-2220032023333303-2110033223213232-0320320131212330"></a>

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

<a id="canonical-0033122023312122-3323023000110110-2233110321213003-3032320103323031-1301233130322022-1321202110021120-2221023211021333-1231012012012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.default_action_next_policy` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.default_action_next_policy

<a id="canonical-2010310213310011-0112301002210230-0130321233033311-0330313012221012-2202211302020012-1132032311212202-1001301030022200-0122110203220032"></a>

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

<a id="canonical-1312200232331133-0311010303232332-0221203000300221-3030003302210100-0321011202310111-3201210103311212-3203032110002002-0031020320121203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.dest_list` properties

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131032101022210-1223221213133331-0000213001130222-2332231313021330-3300310011131301-2303232300330011-1133203013311102-1322100212333020"></a>

### Direct properties for `allow_list.dest_list`

<a id="canonical-3123303120212202-0312212231030033-0301003230313031-2113012211102221-0311301333311112-0202311221012310-1031211031221223-2031132301331211"></a>

#### `allow_list.dest_list.ipv6_prefixes` property

Type: `["list", "string"]`. Computed.

IPv6 Prefixes. Destination IPv6 prefixes.

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

<a id="canonical-0301203230031113-2003221212031303-3221231313102301-3221111022113302-0003322320031022-0211232113300301-0121212123331210-1101122230333310"></a>

#### `allow_list.dest_list.port_ranges` property

Type: `"string"`. Computed.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Additional upstream details:

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

<a id="canonical-1000122110221022-1333330223032123-2213032302323210-3101000131333113-3232331101030210-3223201233322113-2210111213203232-1132330030032023"></a>

#### `allow_list.dest_list.prefixes` property

Type: `["list", "string"]`. Computed.

IPv4 Prefixes. Destination IPv4 prefixes.

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

<a id="canonical-2003210313333300-1210322233003121-2111232031231120-3330220111130030-3303312221122102-3032333222111200-1220013320222331-2322233331321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.http_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.http_list

<a id="canonical-3000203003021202-2211013203331301-1330110033302100-2112133303112221-0013302320330110-1210220123033031-2230212100031122-2101201012303011"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3232230033101013-2331320302330031-3320223100032303-3023300332211221-1113103033333313-1020132212330000-2002333302201213-1031003110102000"></a>

### Direct properties for `allow_list.http_list`

- [any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0112333302030223-1012002210301010-3003120220201120-2111211021212231-0222030110113232-0101203023222230-3320102110000233-0230130110202023): complete subsection reference.

<a id="canonical-3322110132312230-0232233022122131-2013223103023310-2230101001310130-3020132122321212-1030232023032221-1111300000232322-2122213213313331"></a>

<a id="canonical-0310203032210011-1112010211003313-1023021223210102-0211111301303211-1231100323012120-0300332130001303-3130233113320020-1232331333202300"></a>

#### `allow_list.http_list.exact_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0112301132003020-3021210032313011-2320120002001221-3031231310232001-0332011310103121-0120311223022023-2112233113003330-3203213233221200"></a>

#### `allow_list.http_list.path_exact_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0323233302221202-1332001101112202-2002121321031012-3212321211013211-1210132113321000-3132222031030130-0223131300311322-0032210332021323"></a>

#### `allow_list.http_list.path_prefix_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3321233313002120-1020001203110311-1311103212013103-1012203002031113-1301100333332302-2011101311010133-3232310203233123-1220213020022112"></a>

#### `allow_list.http_list.path_regex_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1112301102131133-1230020103130320-2321111301211111-1110230212331101-0313002013100212-0110103033013022-3232313323101210-0302232012203303"></a>

#### `allow_list.http_list.regex_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0201101030110130-0133022201120033-1123120323220323-3001322232020200-0002102222211010-2030311122013120-2123322210220123-1200321322202322"></a>

#### `allow_list.http_list.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0112333302030223-1012002210301010-3003120220201120-2111211021212231-0222030110113232-0101203023222230-3320102110000233-0230130110202023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.http_list.any_path` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- [allow_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2003210313333300-1210322233003121-2111232031231120-3330220111130030-3303312221122102-3032333222111200-1220013320222331-2322233331321302)
- allow_list.http_list.any_path

<a id="canonical-1313032120321232-3120020232300320-0312031332102322-1013232212103311-0102003333231130-0233303131210010-3210032231010320-3210100210132102"></a>

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

<a id="canonical-0130102331212330-2221222233301120-3102100211102202-2313010000230021-0201201101130132-0323322310303120-3203003310013103-2330103233003213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_list.tls_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [allow_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3101200333003233-2232233032312210-1012001303103230-3020221312330330-0112330100123122-0333132122203233-1112103332200202-2102032231311301)
- allow_list.tls_list

<a id="canonical-1310033300013021-0303211310233301-1333201112312230-1020111020333000-3032230113022100-3221231011302033-1232313210111013-1311210021321222"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1131310103230131-3123310032033120-1203232202010230-0200010333323013-0133002020101202-3312103232131102-2333001023212003-2133301013031111"></a>

### Direct properties for `allow_list.tls_list`

<a id="canonical-2303011022320211-1133031222301231-2311313030113101-1201030220130001-0312032211012022-1101232000212132-2030310221010331-0132101022113301"></a>

#### `allow_list.tls_list.exact_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3111032032313302-3002113230202332-0133130320111300-0101022221020201-3230102012003133-2333121012110000-3100303320310220-0122213102021002"></a>

#### `allow_list.tls_list.regex_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0121131210333333-3030301031120200-0011001312231100-0022312333130301-1233221200233322-2130302002200233-3213023122031032-1232030222310010"></a>

#### `allow_list.tls_list.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0301302122312100-0303121321221121-2233002013020312-1211203203103312-1033012321032120-1121133311231131-1221332120300202-1103321303300103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_proxy` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- any_proxy

<a id="canonical-3123313021211132-2130110232333032-0303111023110033-3112321221232113-3211111022110111-2013331232011111-0032333110001313-3002032103013321"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_proxy, drp\_http\_connect, network\_connector, proxy\_label\_selector\] Enable this
option

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

OneOf alternatives in this subsection:

- [any_proxy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3123313021211132-2130110232333032-0303111023110033-3112321221232113-3211111022110111-2013331232011111-0032333110001313-3002032103013321)
- [drp_http_connect](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1232303201032011-2232231121322221-2301300110332023-1032001113103110-3110111022311231-0232330102320221-3220013322023130-3203010220300301)
- [network_connector](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3121013223110112-3301132223320301-3111002130110031-3133111102233212-3002211032112332-0123211112320311-1111302121022100-3210222210003131)
- [proxy_label_selector](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3320221102222111-2312332233123123-0022022013122310-0231213333221113-3301223112111020-0120030021112201-1210012120013200-1210110111112213)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- deny_list

<a id="canonical-1001323031020030-0323022102120112-0100301120113130-2133213212210110-0212222011312233-0032322201311210-2203332013010002-1320313230312301"></a>

Type: `"single"`. Computed.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

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

<a id="canonical-3203233110032031-0212300331032002-1221132331030310-2013132232111112-3232300030310221-1311101011321103-3313230131200013-3031030113302323"></a>

### Direct properties for `deny_list`

- [default_action_allow](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2320023020013103-2020321103323012-0111100023132313-1300030130011212-1100013300221331-1122003113221010-3023301200120322-2103003103223210): complete subsection reference.

- [default_action_deny](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1321033032102020-2223102003211213-2002133030321300-1231211110303320-2220302021100203-0131303313131022-0110030231122211-2112120222320210): complete subsection reference.

- [default_action_next_policy](data-sources--forward_proxy_policy--reference--group-001.md#canonical-1211233001331303-0113112300220112-3013130022201000-2003023332212223-0023233311310321-1030032112211001-3232223121112030-0120322122033330): complete subsection reference.

- [dest_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-3121022111323330-2120112103213211-0021311202003332-2320031121110013-1230010313311233-3213010111011323-3313312310120221-3020003130112323): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0303231023303013-2000101121012313-0330202210132020-1231300133103021-2013033113212023-1103013023330023-2023023313001213-2111003101313132): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2113232131200211-1332200130321213-3202223120213122-1231110313001033-0110120003233212-0233311012330332-1233102100022213-2221100112210232): complete subsection reference.

<a id="canonical-2320023020013103-2020321103323012-0111100023132313-1300030130011212-1100013300221331-1122003113221010-3023301200120322-2103003103223210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.default_action_allow` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.default_action_allow

<a id="canonical-3202123133231120-3101322131023113-1123322322131321-0130100100233133-3021322023330010-0310121002321033-3123110211202120-0110003111203300"></a>

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

<a id="canonical-1321033032102020-2223102003211213-2002133030321300-1231211110303320-2220302021100203-0131303313131022-0110030231122211-2112120222320210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.default_action_deny` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.default_action_deny

<a id="canonical-2112111020202300-2222202333232012-1022323012033111-0030230131103010-2112211121120322-2013321023210300-3200122001111102-0300303301030332"></a>

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

<a id="canonical-1211233001331303-0113112300220112-3013130022201000-2003023332212223-0023233311310321-1030032112211001-3232223121112030-0120322122033330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.default_action_next_policy` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.default_action_next_policy

<a id="canonical-3313011032001133-0311320311311111-2131321213230021-3302013201121103-1131001131231221-1332231010202103-2031302211130121-3213313202020113"></a>

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

<a id="canonical-3121022111323330-2120112103213211-0021311202003332-2320031121110013-1230010313311233-3213010111011323-3313312310120221-3020003130112323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.dest_list` properties

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3331313022001122-1310231011233012-1130031302332230-3313112230030123-2122121131031233-0020202332032031-0121023132223200-3013202320101311"></a>

### Direct properties for `deny_list.dest_list`

<a id="canonical-2330010031030121-3331021223202330-1200330222120201-3001310332111220-3123131203212112-0130231312102002-2111020202222123-1002331302201201"></a>

#### `deny_list.dest_list.ipv6_prefixes` property

Type: `["list", "string"]`. Computed.

IPv6 Prefixes. Destination IPv6 prefixes.

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

<a id="canonical-3300201002311210-0230110202233031-0121312331221032-3101031032130003-3011133132231200-2021102313222101-0021131111013122-1110010202301113"></a>

#### `deny_list.dest_list.port_ranges` property

Type: `"string"`. Computed.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Additional upstream details:

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

<a id="canonical-3121211012321200-1203032323133111-0223111330222001-2212330030303022-3101213010132021-2103113033333013-1112021232133003-1110032032323033"></a>

#### `deny_list.dest_list.prefixes` property

Type: `["list", "string"]`. Computed.

IPv4 Prefixes. Destination IPv4 prefixes.

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

<a id="canonical-0303231023303013-2000101121012313-0330202210132020-1231300133103021-2013033113212023-1103013023330023-2023023313001213-2111003101313132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.http_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.http_list

<a id="canonical-3332030102031203-0322222033321303-0022302000331031-3003201030233013-0131122203222102-3222210031012032-0211203203113000-3233123021230211"></a>

Type: `"list"`. Computed.

HTTP URLs. URLs for HTTP connections.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2301302332333133-3303112133031131-1312223002113223-2200301331330122-0111122021330110-0123333023031222-0201312112130203-0020021112130300"></a>

### Direct properties for `deny_list.http_list`

- [any_path](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2113310132020112-0200202022310230-2221330231311012-2231221100213331-0302201312201010-1321002101330110-3320222232213123-3131302301011300): complete subsection reference.

<a id="canonical-0012230302020123-0330301232311322-3330003032122112-0131201121211311-2201020133100120-1322133021113032-3111033213133313-0312212200101212"></a>

<a id="canonical-3320100021000113-2311111103320103-1213112201201313-2102131233210320-2330200103212020-1220310212210333-0232002112022301-1222120320301122"></a>

#### `deny_list.http_list.exact_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1232221221100301-1332030311220121-0210131210301231-1010131021031321-0001120113010130-1033232231312312-3011213333203102-1000323020320330"></a>

#### `deny_list.http_list.path_exact_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2210101222032312-1211001330221001-3333212110132233-0012103322123200-3002310232102302-0023121101021130-0300301111012233-1133310302122111"></a>

#### `deny_list.http_list.path_prefix_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2003130213303013-0012021232201023-0001112210203100-3100303111310001-1123033001312313-3331113031122302-0320120201333300-3133130331220202"></a>

#### `deny_list.http_list.path_regex_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0030321311130232-1301302312101132-0111011331311321-0102023333321202-2211233023332100-0033030230220003-2233322222121002-3300022130002013"></a>

#### `deny_list.http_list.regex_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0032300231100230-3033223211333203-3011111311021022-3311100302002313-3103310310031322-0200120202220310-0213030303232130-0332303312022302"></a>

#### `deny_list.http_list.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2113310132020112-0200202022310230-2221330231311012-2231221100213331-0302201312201010-1321002101330110-3320222232213123-3131302301011300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.http_list.any_path` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- [deny_list.http_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0303231023303013-2000101121012313-0330202210132020-1231300133103021-2013033113212023-1103013023330023-2023023313001213-2111003101313132)
- deny_list.http_list.any_path

<a id="canonical-3230132322031011-2133220301300220-2012122102202221-3200111331031203-2211310032233301-0011203030130120-0203231311212332-3210111123303223"></a>

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

<a id="canonical-2113232131200211-1332200130321213-3202223120213122-1231110313001033-0110120003233212-0233311012330332-1233102100022213-2221100112210232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_list.tls_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [deny_list](data-sources--forward_proxy_policy--reference--group-001.md#canonical-2122121121221331-3101113313231201-2130023321112032-0332002202032333-1230000011130203-3231223211110132-1011013002122011-1303213030022032)
- deny_list.tls_list

<a id="canonical-1000130231032301-3133031001131333-2331322231330023-3331131322033322-2231110111010132-1021221122331200-1100230331222101-3303103021121013"></a>

Type: `"list"`. Computed.

TLS Domains. Domains in SNI for TLS connections.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0030102232000213-1321211213320002-0211303131111300-1110310010312033-2101232010330023-1112200330220101-3222000010201033-2103222002232031"></a>

### Direct properties for `deny_list.tls_list`

<a id="canonical-2200321232000100-1203230131131332-1301301230020232-1003231001131313-1132120113331121-2030011033213231-2202133130131212-2203110230100321"></a>

#### `deny_list.tls_list.exact_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2112332210130303-2333120120313103-2121221233031110-3123033013310010-2321221002320202-0111230311221222-3323330122023323-0022303312322121"></a>

#### `deny_list.tls_list.regex_value` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2210122300233202-1002120312231103-3311222120010301-0133303010100310-2331131323301223-2011331301022012-1200313020131300-3003222000031313"></a>

#### `deny_list.tls_list.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0301032221101120-3103200113001310-3223232331220200-1332133303133312-1323311130131310-1001333230111322-0230221210331302-3200332121101102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `drp_http_connect` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- drp_http_connect

<a id="canonical-1232303201032011-2232231121322221-2301300110332023-1032001113103110-3110111022311231-0232330102320221-3220013322023130-3203010220300301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for drp http connect.

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

<a id="canonical-3003120302033030-0022133033303133-1021310202213331-0113203001333310-0100202002013200-2202322230231323-0010302133310002-3313303112221211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `network_connector` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- network_connector

<a id="canonical-3121013223110112-3301132223320301-3111002130110031-3133111102233212-3002211032112332-0123211112320311-1111302121022100-3210222210003131"></a>

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

<a id="canonical-3212211213222122-2212233000302031-1231120032021012-3200200222130120-3133312223131221-2332213101321122-3010210102030223-3010203000001201"></a>

### Direct properties for `network_connector`

<a id="canonical-2003113321102020-3010121213103312-3031033201303331-0313231000121112-0232203233330233-2032111202212003-3031322300210301-1203021101021222"></a>

#### `network_connector.name` property

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

<a id="canonical-3001003103103132-0100122102320212-1322201123202103-0102123201312231-3201023213121303-0323322132203322-2003132202010323-3222021230313000"></a>
