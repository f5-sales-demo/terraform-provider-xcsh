---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- Property reference

<a id="canonical-0123212013020010-3023110133203130-2200122012100033-1220112231002032-1331313203131202-3003023131031211-1301300201103131-3323322301031220"></a>

### Direct properties for `xcsh_route`

<a id="canonical-3020120333200033-2300323130121123-2331101331323112-3232210113030120-2032232000033331-2102102200001331-2302312203331132-2330312130323323"></a>

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

<a id="canonical-0332303001321231-1021113033132130-3020102030303131-3020033032033112-0222221330313123-1321012121113102-3122113133000200-0333111301321331"></a>

<a id="canonical-1121201102031232-0232011232201321-2200132203322023-1310131122011011-0122300323123023-0301120221202133-0330031300321220-0123101001332001"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Route.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0033200301122023-0333022010121320-2013222100033003-2332112302011221-0030231203011030-1301221023331003-3232230230322103-2331112232023010"></a>

<a id="canonical-3110322211300201-1223312220131030-0202311333130010-3002301312032123-2322230310012010-0333033131030100-0330000310313233-1011323000323012"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0321011303110330-1310022312231103-0023301122333230-3220002121011002-3223131011031200-2133030002332332-0001232003033033-1013120101232211"></a>

<a id="canonical-1031001233310232-0210002212000203-3032221303220232-2003002312000220-2122131032320210-0232010200220013-2212103300000312-2012320333033130"></a>

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

<a id="canonical-1323133102102313-2333132030131021-3002331332021100-3121301003230212-3222032302132112-2122323330213201-3010133002121123-1133013011003031"></a>

<a id="canonical-3013333011000130-2300232012300021-3011332230230200-0212331101003232-1332330021231132-3322222331001132-3333302000232333-3323120313022021"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Route.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1010103202111321-3333033120130111-3312132223220222-3220330121012122-2210101331033211-3003320013001001-1300313200002011-2100112221200011"></a>

<a id="canonical-2333102122021211-1231132122211123-1213120103111111-1133211203232020-0332131021113032-1013232103113203-2023103322220320-2033213301133112"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Route exists.

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
  }
}
```

- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021): complete subsection reference.

<a id="canonical-3123313223111132-1012233332000333-0113203320003121-1312001010220222-1200331121033023-1330301122301313-3203130301123022-2331212320203222"></a>

### All schema paths for `xcsh_route`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--route--reference--group-001.md#canonical-3020120333200033-2300323130121123-2331101331323112-3232210113030120-2032232000033331-2102102200001331-2302312203331132-2330312130323323) |
| `description` | [description](data-sources--route--reference--group-001.md#canonical-0332303001321231-1021113033132130-3020102030303131-3020033032033112-0222221330313123-1321012121113102-3122113133000200-0333111301321331) |
| `id` | [ID](data-sources--route--reference--group-001.md#canonical-0033200301122023-0333022010121320-2013222100033003-2332112302011221-0030231203011030-1301221023331003-3232230230322103-2331112232023010) |
| `labels` | [labels](data-sources--route--reference--group-001.md#canonical-0321011303110330-1310022312231103-0023301122333230-3220002121011002-3223131011031200-2133030002332332-0001232003033033-1013120101232211) |
| `name` | [name](data-sources--route--reference--group-001.md#canonical-1323133102102313-2333132030131021-3002331332021100-3121301003230212-3222032302132112-2122323330213201-3010133002121123-1133013011003031) |
| `namespace` | [namespace](data-sources--route--reference--group-001.md#canonical-1010103202111321-3333033120130111-3312132223220222-3220330121012122-2210101331033211-3003320013001001-1300313200002011-2100112221200011) |
| `routes` | [routes](data-sources--route--reference--group-001.md#canonical-0320012033211300-0233121012310301-0320121012020311-3003231220302330-3121322100330222-2033320013310023-0301031010300221-0122333021323232) |
| `routes.bot_defense_javascript_injection` | [routes.bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-0021032101233101-3323232033321002-2323001110301313-2030030301033320-3021033010012302-0111031332120230-3210031000323200-1111032112032212) |
| `routes.bot_defense_javascript_injection.javascript_location` | [routes.bot_defense_javascript_injection.javascript_location](data-sources--route--reference--group-001.md#canonical-3201222233321221-2001032000100233-3232100020000303-1333112022223203-0101030132320100-3102233130111233-0302213133103300-3101103120210230) |
| `routes.bot_defense_javascript_injection.javascript_tags` | [routes.bot_defense_javascript_injection.javascript_tags](data-sources--route--reference--group-001.md#canonical-1131200010232100-0231320132322302-2303020203123232-0101023012100122-1022131033230032-3010000230112133-2333111113220010-2313000312110323) |
| `routes.bot_defense_javascript_injection.javascript_tags.javascript_url` | [routes.bot_defense_javascript_injection.javascript_tags.javascript_url](data-sources--route--reference--group-001.md#canonical-0311333322120220-1130113003102123-2222311121132102-2233110131220130-0201122331112213-1003011210202122-3322013120313212-3302102111322122) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes](data-sources--route--reference--group-001.md#canonical-3132333132102112-0022101033131332-3332122002012100-0321332000200323-1003233010212032-0013000202320330-1101033233210220-3233000021132232) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag](data-sources--route--reference--group-001.md#canonical-1331003230133230-0203132020200310-3302331312313221-3103210120100330-0102332111312303-1011320020000011-1010331122120003-2322330030200222) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value](data-sources--route--reference--group-001.md#canonical-2130320221122202-2212210203110210-0322110202030132-1120000010232301-0001333110101321-0301300332200200-3123303111002220-1311113331030123) |
| `routes.disable_location_add` | [routes.disable_location_add](data-sources--route--reference--group-001.md#canonical-1321310200023001-2332033133130021-3031320133013231-2010000001000203-1100303301012321-2010030201211110-0031322102002302-0233230103313220) |
| `routes.inherited_bot_defense_javascript_injection` | [routes.inherited_bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-3223200301103331-3023232320321133-1122331213230113-0312202233121010-3331332123003221-2032103211300311-3232130213330300-3213200333033122) |
| `routes.inherited_waf_exclusion` | [routes.inherited_waf_exclusion](data-sources--route--reference--group-001.md#canonical-0002221303113200-3221213302332021-3202213312032222-2201223023231320-3221203001320122-0201212310231010-3310300101321231-3031121123031121) |
| `routes.match` | [routes.match](data-sources--route--reference--group-001.md#canonical-0013130210101202-2031231200311303-0112132131220101-0032301000312111-2321230213130321-2230010011222120-3101021201112300-0210233330020000) |
| `routes.match.headers` | [routes.match.headers](data-sources--route--reference--group-001.md#canonical-3323230031312220-2200333123130100-3103323003322212-0132302322011133-3002100222303231-3313331313010320-3132333200033202-2322123321030221) |
| `routes.match.headers.exact` | [routes.match.headers.exact](data-sources--route--reference--group-001.md#canonical-2310311221020233-1222020220110320-2133300322003323-3322333222101123-3303213031333012-1221211320122113-1211213222331112-1212331123103133) |
| `routes.match.headers.invert_match` | [routes.match.headers.invert_match](data-sources--route--reference--group-001.md#canonical-0101213120100120-3230003220303003-0131123000033202-0310130320022002-1220213211312012-3231023302232212-0202232122302303-0222021112302033) |
| `routes.match.headers.name` | [routes.match.headers.name](data-sources--route--reference--group-001.md#canonical-0032013013131033-0122213302303032-3210200003102310-2213000112130010-2101102212131132-1023032130022011-2202311131313100-1112121100133300) |
| `routes.match.headers.presence` | [routes.match.headers.presence](data-sources--route--reference--group-001.md#canonical-0130301021233311-2332132320002331-3132031002130201-3112230320231303-2121231222230300-0123303213312212-2001310011013233-1031332321131333) |
| `routes.match.headers.regex` | [routes.match.headers.regex](data-sources--route--reference--group-001.md#canonical-3222333223001120-3010011201303332-1201333013312313-1003303131120012-2022333222001113-0220110132012120-2223101011232132-2213210013113203) |
| `routes.match.http_method` | [routes.match.http_method](data-sources--route--reference--group-001.md#canonical-2232300131002221-2202101130132111-3131111330011213-1011230121001300-0210203231312020-3120203202333232-0123312330122202-1311200112000312) |
| `routes.match.incoming_port` | [routes.match.incoming_port](data-sources--route--reference--group-001.md#canonical-3020231022012021-3011123331001312-2012203112303122-3313212132120201-2311311221123210-3131312103303303-1030223203021123-1030123002332122) |
| `routes.match.incoming_port.no_port_match` | [routes.match.incoming_port.no_port_match](data-sources--route--reference--group-001.md#canonical-1213112211110031-1032230031310322-2003322003121121-2102200230320011-3111100020222320-1112210233101001-0031013132022223-1223323000023112) |
| `routes.match.incoming_port.port` | [routes.match.incoming_port.port](data-sources--route--reference--group-001.md#canonical-3202302213032302-2303131203103323-3031330301211312-0102211132010130-3221032132030030-0220321120220313-1013111112001311-3323233202331121) |
| `routes.match.incoming_port.port_ranges` | [routes.match.incoming_port.port_ranges](data-sources--route--reference--group-001.md#canonical-2132133312222012-0100310313322112-0332130212001122-2233233121102320-1321032300222033-1111210221202312-0323232110203021-0022200331330102) |
| `routes.match.path` | [routes.match.path](data-sources--route--reference--group-001.md#canonical-3002112200132320-1123033113020333-1320211330313202-3002300002201010-3131133221132011-3301200330010321-0312102222020211-0002100333101203) |
| `routes.match.path.path` | [routes.match.path.path](data-sources--route--reference--group-001.md#canonical-0020110020001231-0012301101011231-0001001123310233-0112223113032200-0123020320232100-1213212220320123-0302220022332230-1023013010033101) |
| `routes.match.path.prefix` | [routes.match.path.prefix](data-sources--route--reference--group-001.md#canonical-3101312232311100-1300110100020312-3232103012020101-3223210010321311-0300120031023010-2210021102333122-3131310201001331-2001003003220003) |
| `routes.match.path.regex` | [routes.match.path.regex](data-sources--route--reference--group-001.md#canonical-0113233220320010-0132221113332130-0333302313112131-0123122003321021-0122200312220321-1001023213032102-0201212231310123-2112233103312222) |
| `routes.match.query_params` | [routes.match.query_params](data-sources--route--reference--group-001.md#canonical-1202202201102123-3320301132300112-3322331022310212-3121030231331300-1232130001230003-2311203130302223-0310300332210113-1031333203022013) |
| `routes.match.query_params.exact` | [routes.match.query_params.exact](data-sources--route--reference--group-001.md#canonical-2010232202021031-0321023112120020-1300230112102033-0020212132222201-0102131313231322-0011111012313020-0101222320001010-2131030320133202) |
| `routes.match.query_params.key` | [routes.match.query_params.key](data-sources--route--reference--group-001.md#canonical-3300122321000130-1313001023302210-1121002320211010-2003001012132112-0313310330022312-3120303330223022-0233213313332031-1212011330123210) |
| `routes.match.query_params.regex` | [routes.match.query_params.regex](data-sources--route--reference--group-001.md#canonical-3320201203203202-1310222123110001-1331331031103101-0010033201100032-2321100003233002-2003000112200101-3113112320212330-3302220221232223) |
| `routes.request_cookies_to_add` | [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-1223121012233330-3231123230130033-0311021022223133-2023213011131122-2310310020302220-0200323133031311-2302000111001021-3001203022333220) |
| `routes.request_cookies_to_add.name` | [routes.request_cookies_to_add.name](data-sources--route--reference--group-001.md#canonical-1012110131222200-2333121000112002-3013322100313301-2301331322201133-2321331003023300-3312232110320002-3111120120302011-3322300122330100) |
| `routes.request_cookies_to_add.overwrite` | [routes.request_cookies_to_add.overwrite](data-sources--route--reference--group-001.md#canonical-2113200121312011-2122212302313323-3331331301110032-2230132213013121-3311022020023221-3313201121222021-0201123300123223-3133311320111332) |
| `routes.request_cookies_to_add.secret_value` | [routes.request_cookies_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-3303132222213001-3103200013311031-0012233130013322-1303013020012300-2101220003113312-1133032100112203-1003330211322321-3302213311022301) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-001.md#canonical-2203111113222230-1000203001203202-0331001102331103-3020021002231030-0103202212122321-0212131003212123-2303332330013311-2000211300230323) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--route--reference--group-001.md#canonical-1030000301013130-0113101231100132-3002012331303131-2020132012330011-0032031331332003-2120023033033102-3101333002131323-2131102123120320) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--route--reference--group-001.md#canonical-3120102230233230-2222312102221012-0322103331233132-2021013311222010-3033220330023131-2213310030022030-0130310220320322-3220123313200231) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--route--reference--group-001.md#canonical-1032330223212221-3202222131311123-1331011011211030-0113322003313030-0122132033200111-2231223020110011-3000221100201102-3130103013202111) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info` | [routes.request_cookies_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-001.md#canonical-1231022201203102-2201022221231330-1331102232122310-1310023310012322-2003330201012210-0312020112311122-2112000303011013-3332312200330102) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--route--reference--group-001.md#canonical-0112002312332131-2132312212203003-3332110112022022-1211300121313023-0323000002123022-0331200323210123-3210332111321102-3113332212020330) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.url` | [routes.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--route--reference--group-001.md#canonical-0232233300300022-1010210121033022-0002332111112203-2212322122320323-3031220110222103-0132013231122002-1032323310122121-2021213030003033) |
| `routes.request_cookies_to_add.value` | [routes.request_cookies_to_add.value](data-sources--route--reference--group-001.md#canonical-0022021021001130-0001131223000102-1202210022102001-3310003320130300-3112121301203010-0031310101130223-2102111300232031-0231310003221131) |
| `routes.request_cookies_to_remove` | [routes.request_cookies_to_remove](data-sources--route--reference--group-001.md#canonical-3033022302301101-1023101032231000-2132023011100330-0301320202033302-3311210302210033-0121213201300000-1202102333103011-1112130111221132) |
| `routes.request_headers_to_add` | [routes.request_headers_to_add](data-sources--route--reference--group-002.md#canonical-3003023133133313-2311002201102312-1303313131123220-2130233111202333-0211202201202300-2020102303201032-1320112201113131-2312221010222200) |
| `routes.request_headers_to_add.append` | [routes.request_headers_to_add.append](data-sources--route--reference--group-002.md#canonical-3220133111130133-2202301000232330-3133221102102320-2010002232311312-0203330322113312-3111202032123102-2022201323011331-1320120232011301) |
| `routes.request_headers_to_add.name` | [routes.request_headers_to_add.name](data-sources--route--reference--group-002.md#canonical-1033321013023132-1120002130323302-0312303112200131-2021322002022003-0032113320010030-3203113101213322-3331301001003001-2300301103300032) |
| `routes.request_headers_to_add.secret_value` | [routes.request_headers_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-3320233012100023-2300202310101300-2032102123001211-2313002301123002-2230122230221131-2022032013301032-3123331033013103-2031113033232321) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info` | [routes.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-002.md#canonical-1000121202221302-0230031232230112-3200331313130111-0121100022303033-3110311122102002-2001000122231101-3030101101130123-0013110111212030) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--route--reference--group-002.md#canonical-3231023210130201-1123003312301331-2321021031031001-0111202002310321-0002033011131133-0222133323030200-0200211123132020-1303023332100001) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--route--reference--group-002.md#canonical-1132313222210321-3230200220313111-0310030022312313-3111311220333031-0210322213222110-0031333003230031-2201303330100303-0100031130123301) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--route--reference--group-002.md#canonical-0231021011133121-2201002100121032-3212123011313323-1131331233312031-0100333330102133-1333023210200132-2013300111201313-3311131322001312) |
| `routes.request_headers_to_add.secret_value.clear_secret_info` | [routes.request_headers_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-002.md#canonical-3221123313112322-2112011033201222-1231132311322211-0222001110303032-1032023130121023-3022312310203102-0102312221331131-2210323302211011) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--route--reference--group-002.md#canonical-3220203131213120-1211302213220111-1130211303110022-2301100231311102-3311201113123200-3102123002113200-3101021320001311-1323003223121102) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.url` | [routes.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--route--reference--group-002.md#canonical-3030202021321213-0300031113130021-2023321012311013-1323023123021013-3223001332010121-0231200332220020-1203122231201202-0103331033031302) |
| `routes.request_headers_to_add.value` | [routes.request_headers_to_add.value](data-sources--route--reference--group-002.md#canonical-2000130020132212-3001011003110202-2130322010023023-2030011012302002-0023301011122333-2330112011122231-1322230123320302-1312231313110131) |
| `routes.request_headers_to_remove` | [routes.request_headers_to_remove](data-sources--route--reference--group-001.md#canonical-2110331033213021-1001030033020203-2033201231321201-3210111101012023-3100211211122023-1320032212221311-0333023031230203-2301032112010212) |
| `routes.response_cookies_to_add` | [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3202023130203030-3021232130210002-3010232023121333-1310232103101313-1321310113112223-3210010210212202-1020203032301213-2120313322120233) |
| `routes.response_cookies_to_add.add_domain` | [routes.response_cookies_to_add.add_domain](data-sources--route--reference--group-002.md#canonical-3100022203120300-0132303223300112-2330032112001102-3203310332131202-1300121222130201-0233200112032000-3002333031300321-3023021110103202) |
| `routes.response_cookies_to_add.add_expiry` | [routes.response_cookies_to_add.add_expiry](data-sources--route--reference--group-002.md#canonical-3221323212232013-3300331313322302-3101100101111223-3311321000310301-3122312331121322-1221021203123201-1311233303213003-0302222100113030) |
| `routes.response_cookies_to_add.add_httponly` | [routes.response_cookies_to_add.add_httponly](data-sources--route--reference--group-002.md#canonical-1303132013102032-3023023130201123-0030012332200113-3131212310111303-2103121300122101-3202231021220000-3030103021330100-3020131103111020) |
| `routes.response_cookies_to_add.add_partitioned` | [routes.response_cookies_to_add.add_partitioned](data-sources--route--reference--group-002.md#canonical-1112333111223210-3003113110300011-3210332220201322-2111112213310310-1022221313002333-1201220023222323-2103122033201112-1320220033212120) |
| `routes.response_cookies_to_add.add_path` | [routes.response_cookies_to_add.add_path](data-sources--route--reference--group-002.md#canonical-0330102120000112-1212001202001313-2023312113013232-0233123003201030-1012013210033011-3331311321223021-3231321021203223-2003212110211212) |
| `routes.response_cookies_to_add.add_secure` | [routes.response_cookies_to_add.add_secure](data-sources--route--reference--group-002.md#canonical-0101013013130113-2031003112130123-1112121203323322-0332020223001301-2333103131300321-3000311320331030-3220000121212202-1230211220003333) |
| `routes.response_cookies_to_add.ignore_domain` | [routes.response_cookies_to_add.ignore_domain](data-sources--route--reference--group-002.md#canonical-3031021330200310-1012200132012200-3313320213033223-3221222033233322-0100123202310121-1302132233221310-3011333230022321-0301012121100321) |
| `routes.response_cookies_to_add.ignore_expiry` | [routes.response_cookies_to_add.ignore_expiry](data-sources--route--reference--group-002.md#canonical-0033313321123231-3233032313223320-2203010120112122-2003132131212001-3002202102210220-0133130013012000-0121212233132233-1333332023110202) |
| `routes.response_cookies_to_add.ignore_httponly` | [routes.response_cookies_to_add.ignore_httponly](data-sources--route--reference--group-002.md#canonical-3003111333130300-1311201021212100-2101020002003302-0333210212032001-1132003013311111-2112322330323113-0210313330032302-0033032321203332) |
| `routes.response_cookies_to_add.ignore_max_age` | [routes.response_cookies_to_add.ignore_max_age](data-sources--route--reference--group-002.md#canonical-3120223322221221-0113223033222003-1021222201111020-2301102331110320-2130101313030300-0331210212003030-1012022022031320-1032300123130222) |
| `routes.response_cookies_to_add.ignore_partitioned` | [routes.response_cookies_to_add.ignore_partitioned](data-sources--route--reference--group-002.md#canonical-2300132020121310-3312110102021333-0000321013333322-0332021131201303-2313133113032310-2210023103321131-0103321111011123-2313033032122113) |
| `routes.response_cookies_to_add.ignore_path` | [routes.response_cookies_to_add.ignore_path](data-sources--route--reference--group-002.md#canonical-0002212130310110-2121323102232332-1100220310201323-2121220122231210-0331003211113302-0230031010311103-1233210322100332-1033121221203111) |
| `routes.response_cookies_to_add.ignore_samesite` | [routes.response_cookies_to_add.ignore_samesite](data-sources--route--reference--group-002.md#canonical-2031011013010021-2030221023333133-1001012031030023-0213320222223310-3303320011233332-0231222112122322-2001030311021133-3130111303323023) |
| `routes.response_cookies_to_add.ignore_secure` | [routes.response_cookies_to_add.ignore_secure](data-sources--route--reference--group-002.md#canonical-3232003310300232-3013031310031211-2321101201200032-1232133332300230-3133332110103003-3221001233302331-2320231323002120-1010312322013103) |
| `routes.response_cookies_to_add.ignore_value` | [routes.response_cookies_to_add.ignore_value](data-sources--route--reference--group-002.md#canonical-1212131111303233-1100331201120033-0123332320112020-0322011131220313-2001120011113010-0220012312232011-0332013223232100-2312332021011223) |
| `routes.response_cookies_to_add.max_age_value` | [routes.response_cookies_to_add.max_age_value](data-sources--route--reference--group-002.md#canonical-3312220203003201-1202332322121102-3101202332310032-0033020112330130-2113322331022211-2012013121130311-2313032132203030-3321213130032303) |
| `routes.response_cookies_to_add.name` | [routes.response_cookies_to_add.name](data-sources--route--reference--group-002.md#canonical-3133133202210310-0333131000220312-0130002212230020-2202010212010310-0103311003301333-2213012011122003-3122220031100310-1033323331233121) |
| `routes.response_cookies_to_add.overwrite` | [routes.response_cookies_to_add.overwrite](data-sources--route--reference--group-002.md#canonical-1110300321121313-3321333231101102-3112202102312222-2101333031130333-0233320332121122-3103333122033011-2101122320303122-3221022022321320) |
| `routes.response_cookies_to_add.samesite_lax` | [routes.response_cookies_to_add.samesite_lax](data-sources--route--reference--group-002.md#canonical-3302021321231001-2322002000002223-0333113312101322-2030233101232012-1320030300223312-0120000033232303-2232032332220010-2011212011303222) |
| `routes.response_cookies_to_add.samesite_none` | [routes.response_cookies_to_add.samesite_none](data-sources--route--reference--group-002.md#canonical-2130231100213011-1111031200220311-2020210122112130-2211302003310301-2330230023313102-2020001203203212-3321312211131313-1233310333130232) |
| `routes.response_cookies_to_add.samesite_strict` | [routes.response_cookies_to_add.samesite_strict](data-sources--route--reference--group-002.md#canonical-1131213011031031-0220033311302000-1011301203100032-2013320001033031-0121030011301331-2210213200311212-3331130112232232-1311131113333323) |
| `routes.response_cookies_to_add.secret_value` | [routes.response_cookies_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-0013232213233331-2300123023120132-1303001112122331-2120313301321033-2312202313003230-0111233302000113-2120222213301221-0200002130233302) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-002.md#canonical-3331330211103113-0001221300202302-1123130230131013-3211030121123111-3331220220331313-0022130322100001-2002211022312310-2312132320323131) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--route--reference--group-002.md#canonical-3113320100320202-2311133301331303-2323330031222011-1300110121321001-1330320031211301-0032203231322112-3103331120212001-2030030332121300) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--route--reference--group-002.md#canonical-2031112221010300-0130330131122123-1230030100301333-2120301111020010-3020200202033231-0320002020230230-1033122311133130-1130030103202333) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--route--reference--group-002.md#canonical-1102300031110132-1331122223233331-1222333100320003-2022212032311130-0031103122321102-3013122231303030-3122333311232033-2100032212232020) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info` | [routes.response_cookies_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-002.md#canonical-3213102213132123-1201133011001322-3023130222130000-0330020113333021-3022112322210322-3032320000211231-2112302202212202-2033221201133332) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--route--reference--group-002.md#canonical-2023320231323210-1132133032201102-2011320132212223-2021100321132013-3111222011320111-1000211312222210-3111322330232032-2320030022332020) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.url` | [routes.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--route--reference--group-002.md#canonical-3021333022202231-2102111001221113-2120103330100000-3300321010333002-1303122322120110-3303301313233301-1303032000122320-0001130100133313) |
| `routes.response_cookies_to_add.value` | [routes.response_cookies_to_add.value](data-sources--route--reference--group-002.md#canonical-2302233101300301-0011022330231212-0101010113211333-2112211032210201-3102130313203131-3103100120101003-1131013202200120-2022223220302211) |
| `routes.response_cookies_to_remove` | [routes.response_cookies_to_remove](data-sources--route--reference--group-001.md#canonical-2122001010102022-3212130023313021-3103320112003010-0120130000010032-3210123121313301-1130010200132133-2302010223133103-1120232203202133) |
| `routes.response_headers_to_add` | [routes.response_headers_to_add](data-sources--route--reference--group-002.md#canonical-0203001121301223-3212032101312230-1230230200012310-3233330320211312-0120323333312032-1111130030201210-3323122202212311-0002200220031312) |
| `routes.response_headers_to_add.append` | [routes.response_headers_to_add.append](data-sources--route--reference--group-002.md#canonical-0100303221213101-1021103202122123-3020321120121213-1021130302231322-0230130323321131-3203220220120121-0320302130220222-0222001131130210) |
| `routes.response_headers_to_add.name` | [routes.response_headers_to_add.name](data-sources--route--reference--group-002.md#canonical-0020010132023210-0132320303312221-2322100130300100-2000120230023121-2223300012211020-0220020030311320-2231021010133231-2010330300330110) |
| `routes.response_headers_to_add.secret_value` | [routes.response_headers_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-2111300222103120-3330313130331002-0011033102011312-1102321002021202-2002101212303231-2221121012222121-1232211223103010-0020233300021021) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info` | [routes.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--route--reference--group-002.md#canonical-3303001330031120-0023013233031032-0011302310320131-0013122121223331-2333321310230013-2202110020220303-2130021203013303-2233103221221223) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--route--reference--group-002.md#canonical-3331203103223013-1303100323220122-2101333312211330-2001303033200202-1321313233101111-1332021310121330-1210010313001011-0130032332200223) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--route--reference--group-002.md#canonical-3113320010231333-1233131031100202-0213220030113112-1100022321022113-1220011331002231-2223332123323032-1223201323123302-3133102312003010) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--route--reference--group-002.md#canonical-2201222211030213-0121012221212013-1101210333002122-0222023002031300-0012000233332021-1233011201201112-3120020222210001-2011333123131313) |
| `routes.response_headers_to_add.secret_value.clear_secret_info` | [routes.response_headers_to_add.secret_value.clear_secret_info](data-sources--route--reference--group-002.md#canonical-3223231310032112-2113032003332221-0203233330221030-3101311120301312-0311121331120023-2211223132003122-1321233130111131-3302232001101310) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--route--reference--group-002.md#canonical-3233102233322112-3210021220103123-2103003031003200-0223113001021233-0002231012100100-2320333023211332-1032212103310021-2103020330122121) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.url` | [routes.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--route--reference--group-002.md#canonical-0030231031000102-1233131012031030-3003312013320022-1320323212012113-1022121132133022-0332010130000013-2013320223210322-1101231203311011) |
| `routes.response_headers_to_add.value` | [routes.response_headers_to_add.value](data-sources--route--reference--group-002.md#canonical-1101223032332231-3200231002331032-2130120310233232-1222010322113232-3001332303300011-2103302032200211-1012003120323202-3200301110300131) |
| `routes.response_headers_to_remove` | [routes.response_headers_to_remove](data-sources--route--reference--group-001.md#canonical-0100323310101123-1200301200110233-0113332200103222-1120112302030131-3022202033232021-2131001000111011-2301121301002022-0200000300023313) |
| `routes.route_destination` | [routes.route_destination](data-sources--route--reference--group-002.md#canonical-2012333103302102-1200112011010103-2321032030203303-0230113222313323-1003211133101332-3200122302022000-3111311320333223-1221111012203323) |
| `routes.route_destination.auto_host_rewrite` | [routes.route_destination.auto_host_rewrite](data-sources--route--reference--group-003.md#canonical-3320103021021333-2320302212111331-0333023310023013-2021311320021223-1102233012033200-3001223012121022-0032120021023131-2331312331131202) |
| `routes.route_destination.buffer_policy` | [routes.route_destination.buffer_policy](data-sources--route--reference--group-003.md#canonical-1302233012213012-2332130203230332-0022033321122321-0133100101210112-1312310002210131-1113101103331313-1232122213110233-2321302000032132) |
| `routes.route_destination.buffer_policy.disabled` | [routes.route_destination.buffer_policy.disabled](data-sources--route--reference--group-003.md#canonical-0002220032322303-1310330103010313-2313010133021202-1200032201132202-2321101203320221-1022100321030230-1332003311230310-0333311223011321) |
| `routes.route_destination.buffer_policy.max_request_bytes` | [routes.route_destination.buffer_policy.max_request_bytes](data-sources--route--reference--group-003.md#canonical-3112021200022010-1000103231203003-2000321330011133-2231312333020033-2022001202122221-3210232312211001-2113022212320313-0103323002000111) |
| `routes.route_destination.cors_policy` | [routes.route_destination.cors_policy](data-sources--route--reference--group-003.md#canonical-1222330011322002-0121211303010100-3332332302301010-2230013201321332-1100220320003030-3310212210030220-3013310222103221-3023003132010010) |
| `routes.route_destination.cors_policy.allow_credentials` | [routes.route_destination.cors_policy.allow_credentials](data-sources--route--reference--group-003.md#canonical-2000320200211022-2200333213120312-1101101031300230-0033313323013312-2021212213011303-3130322330332330-3210301122102123-0101211320312030) |
| `routes.route_destination.cors_policy.allow_headers` | [routes.route_destination.cors_policy.allow_headers](data-sources--route--reference--group-003.md#canonical-1232032130332330-2322331323030120-1223220211331111-0213010132302221-0032101022023323-2222231330331211-0320003320131213-3323102201220132) |
| `routes.route_destination.cors_policy.allow_methods` | [routes.route_destination.cors_policy.allow_methods](data-sources--route--reference--group-003.md#canonical-1200200212313201-3312003232031101-0323120333002110-1002213223002023-1123121120002022-1300222010102300-1112302103133010-2133321102300203) |
| `routes.route_destination.cors_policy.allow_origin` | [routes.route_destination.cors_policy.allow_origin](data-sources--route--reference--group-003.md#canonical-1331003012231210-1332013312023323-3310312012031111-0313202302002112-3133110332320223-2321002132122022-2132220020313330-3311100300232132) |
| `routes.route_destination.cors_policy.allow_origin_regex` | [routes.route_destination.cors_policy.allow_origin_regex](data-sources--route--reference--group-003.md#canonical-2022023130102200-0323232032112102-0312000030311021-2313130302321231-2121112311003300-2013111230211301-1011300231332132-1130201232330321) |
| `routes.route_destination.cors_policy.disabled` | [routes.route_destination.cors_policy.disabled](data-sources--route--reference--group-003.md#canonical-0130220300201313-3120312323303032-0220202022310132-0113023010220230-3221313101102111-1232222030333232-2203020313200120-3210320330120300) |
| `routes.route_destination.cors_policy.expose_headers` | [routes.route_destination.cors_policy.expose_headers](data-sources--route--reference--group-003.md#canonical-1202201232130132-2103221113011010-0130002022203300-3002222232001003-3312222321231032-1332012002313322-2100133303123301-1100222022330131) |
| `routes.route_destination.cors_policy.maximum_age` | [routes.route_destination.cors_policy.maximum_age](data-sources--route--reference--group-003.md#canonical-3100202122313133-1230200133310222-1100322311323303-1011221101323331-0101300021101333-2001133330020123-3300332332113331-0200021300210112) |
| `routes.route_destination.csrf_policy` | [routes.route_destination.csrf_policy](data-sources--route--reference--group-003.md#canonical-0012221000302210-2120130031101110-2012202323101031-2313320121113330-2200000231213221-1313213202031322-1011110302101130-1333231101330023) |
| `routes.route_destination.csrf_policy.all_load_balancer_domains` | [routes.route_destination.csrf_policy.all_load_balancer_domains](data-sources--route--reference--group-003.md#canonical-1032000232021032-2203220130120331-0131301313131311-1113122112330013-1303030130220023-3213321221033130-1233133133222002-2310000300110123) |
| `routes.route_destination.csrf_policy.custom_domain_list` | [routes.route_destination.csrf_policy.custom_domain_list](data-sources--route--reference--group-003.md#canonical-1310230003302211-0320032131312003-0302033020300113-1032310211003213-3020121102221123-3111233313101101-1022002112100212-0110310320332300) |
| `routes.route_destination.csrf_policy.custom_domain_list.domains` | [routes.route_destination.csrf_policy.custom_domain_list.domains](data-sources--route--reference--group-003.md#canonical-2213030300110021-1211211230230101-0012111231313220-1123333213120122-0011132033230101-1301213333331200-0002213210303013-0121310321000213) |
| `routes.route_destination.csrf_policy.disabled` | [routes.route_destination.csrf_policy.disabled](data-sources--route--reference--group-003.md#canonical-2301122332221223-2232322311213222-1122330322022022-2221330002100111-1302230213111031-3232231211221220-1131132101110002-1121311021311332) |
| `routes.route_destination.destinations` | [routes.route_destination.destinations](data-sources--route--reference--group-003.md#canonical-0123103300201300-0133220001112102-3123313021033101-2011331230332103-2003131110320020-0120031313112313-2032310300213120-1220333031112033) |
| `routes.route_destination.destinations.cluster` | [routes.route_destination.destinations.cluster](data-sources--route--reference--group-003.md#canonical-3231230112300012-3211213301201003-2110332000213333-0101013113222110-3201203101110031-0232021023213013-2203003301020202-0203200201230120) |
| `routes.route_destination.destinations.cluster.kind` | [routes.route_destination.destinations.cluster.kind](data-sources--route--reference--group-003.md#canonical-2023203023322221-1321230112011222-3011221132032203-3302010331000120-0320332303303131-2003102120210212-2101330023112130-0100302203123323) |
| `routes.route_destination.destinations.cluster.name` | [routes.route_destination.destinations.cluster.name](data-sources--route--reference--group-003.md#canonical-3112311221211210-0313322000331322-3221221300132210-2232132021013313-1220221000112000-2333233001121333-2213300321023020-2002022213302313) |
| `routes.route_destination.destinations.cluster.namespace` | [routes.route_destination.destinations.cluster.namespace](data-sources--route--reference--group-003.md#canonical-1132332131312022-0012121102011013-0303013013202103-0003020321232012-1030301332303130-0103231313233132-2201032303312133-2130120022312231) |
| `routes.route_destination.destinations.cluster.tenant` | [routes.route_destination.destinations.cluster.tenant](data-sources--route--reference--group-003.md#canonical-1312033031332331-2333123300111032-3312101222223021-0203001102102000-0213003332030022-1000132313313211-1133131312313020-0203333113111120) |
| `routes.route_destination.destinations.cluster.uid` | [routes.route_destination.destinations.cluster.uid](data-sources--route--reference--group-003.md#canonical-3100310020103031-3322011000230033-0321313032113210-1231022103033223-0033200311230123-2321003021313322-3033102131111333-2002011131010313) |
| `routes.route_destination.destinations.endpoint_subsets` | [routes.route_destination.destinations.endpoint_subsets](data-sources--route--reference--group-003.md#canonical-3221313010300123-3230231303320312-1033130310323111-3113200111331201-1130201213122301-0000220101113010-0203222023210212-2203212230320121) |
| `routes.route_destination.destinations.priority` | [routes.route_destination.destinations.priority](data-sources--route--reference--group-003.md#canonical-1103023203302233-1103002001301022-0222322021130231-1132132130320201-1310323121113320-2221333201321103-0331130022003132-3312332022233033) |
| `routes.route_destination.destinations.weight` | [routes.route_destination.destinations.weight](data-sources--route--reference--group-003.md#canonical-2022211031111231-2212022313310213-0013233222122113-0212131131202200-1033030331020132-1212312011300011-2131120020221311-1320021331131102) |
| `routes.route_destination.do_not_retract_cluster` | [routes.route_destination.do_not_retract_cluster](data-sources--route--reference--group-003.md#canonical-3300002313111001-3022013222331312-3300201203021323-1120332212313202-2030231331113330-0331132000121002-2312231120211321-1211203021322210) |
| `routes.route_destination.endpoint_subsets` | [routes.route_destination.endpoint_subsets](data-sources--route--reference--group-003.md#canonical-0111220131233220-2011102323133313-1112213011000231-1212031321203230-2101000321020020-0113220131200313-2211002023110200-3210022323112030) |
| `routes.route_destination.hash_policy` | [routes.route_destination.hash_policy](data-sources--route--reference--group-003.md#canonical-0101203232313012-0132111110103110-1210323003221112-0231201021123032-0131030302333120-2031100233303130-1203221210132131-0010013120110303) |
| `routes.route_destination.hash_policy.cookie` | [routes.route_destination.hash_policy.cookie](data-sources--route--reference--group-003.md#canonical-1330312103300201-0203133022010202-0102130111003330-3333121220300030-0233032132221322-1310332022201012-1012311331011031-1010011023002321) |
| `routes.route_destination.hash_policy.cookie.add_httponly` | [routes.route_destination.hash_policy.cookie.add_httponly](data-sources--route--reference--group-003.md#canonical-2212202030102202-2300032213100111-0133023110320230-3130220102103131-2321111133123221-2202303313113203-3133320213311320-0221003103021112) |
| `routes.route_destination.hash_policy.cookie.add_secure` | [routes.route_destination.hash_policy.cookie.add_secure](data-sources--route--reference--group-003.md#canonical-0312100323202112-3130232033222231-3022333013221230-3100310231021111-1223022301123111-0113021121122201-2102120222222011-3321133012303010) |
| `routes.route_destination.hash_policy.cookie.ignore_httponly` | [routes.route_destination.hash_policy.cookie.ignore_httponly](data-sources--route--reference--group-003.md#canonical-0121200123102130-2010232111210121-2320220033300111-2123003301332311-2030323003303203-2222113313203132-0000121303001112-2033010302213101) |
| `routes.route_destination.hash_policy.cookie.ignore_samesite` | [routes.route_destination.hash_policy.cookie.ignore_samesite](data-sources--route--reference--group-003.md#canonical-1113021333132020-1312013331300031-1313100321010012-1012233213112311-1021211022023020-1012221010321113-1013223213033323-3320232033111331) |
| `routes.route_destination.hash_policy.cookie.ignore_secure` | [routes.route_destination.hash_policy.cookie.ignore_secure](data-sources--route--reference--group-003.md#canonical-1200220122311232-1100133010122202-3010312200212223-2230213033113302-2211003220132203-3223021103101232-3210030223021222-1202112013003322) |
| `routes.route_destination.hash_policy.cookie.name` | [routes.route_destination.hash_policy.cookie.name](data-sources--route--reference--group-003.md#canonical-1132232000111222-2303302022220233-3220220322013210-0210332111012131-3331310201301301-0103232033033301-0100011232301232-2130211213323102) |
| `routes.route_destination.hash_policy.cookie.path` | [routes.route_destination.hash_policy.cookie.path](data-sources--route--reference--group-003.md#canonical-3302313302130112-1001330323221022-3222023021331021-3223211221001020-0113121120212231-1332013201232000-3201330321222012-0133021032003223) |
| `routes.route_destination.hash_policy.cookie.samesite_lax` | [routes.route_destination.hash_policy.cookie.samesite_lax](data-sources--route--reference--group-003.md#canonical-1101030203132110-0030302300002311-3200031203113002-2220100020130211-0123333231031310-1212001300113030-0132202103212210-2101103331311332) |
| `routes.route_destination.hash_policy.cookie.samesite_none` | [routes.route_destination.hash_policy.cookie.samesite_none](data-sources--route--reference--group-003.md#canonical-2020310323332301-0111033321003303-3303032020132310-2023000202310013-1120222211233133-1110211113102203-1233133332301211-1012230110133210) |
| `routes.route_destination.hash_policy.cookie.samesite_strict` | [routes.route_destination.hash_policy.cookie.samesite_strict](data-sources--route--reference--group-003.md#canonical-3222201312131122-2200132013000131-0220203130232302-1330230132131011-1020201000323001-2321323313303221-1101003200020202-1110030012302011) |
| `routes.route_destination.hash_policy.cookie.ttl` | [routes.route_destination.hash_policy.cookie.ttl](data-sources--route--reference--group-003.md#canonical-3233023033100300-3030032022322130-2121003320033303-3202301333130030-2202033010010003-3111330230102323-1133232011203332-1113221110000213) |
| `routes.route_destination.hash_policy.header_name` | [routes.route_destination.hash_policy.header_name](data-sources--route--reference--group-003.md#canonical-3331100322001333-1103323120000230-2313220003312222-1303013313123111-3233301112322001-0012033330101111-2321313103233322-0200031001133103) |
| `routes.route_destination.hash_policy.source_ip` | [routes.route_destination.hash_policy.source_ip](data-sources--route--reference--group-003.md#canonical-0212101201012011-1230332123331333-0202321200213020-3021233331332113-2232303131323312-0200130112201132-1122020111330132-1231221320233020) |
| `routes.route_destination.hash_policy.terminal` | [routes.route_destination.hash_policy.terminal](data-sources--route--reference--group-003.md#canonical-1112023313102231-1202210123321230-0122133132110022-1222210102011110-0133222213021121-2020011220321210-2200022230011222-0331211031120112) |
| `routes.route_destination.host_rewrite` | [routes.route_destination.host_rewrite](data-sources--route--reference--group-003.md#canonical-2133310012230300-3121322112101113-1301130310223033-3213232333331021-2033000013121000-0321223033222022-2121013031130022-1203022101010203) |
| `routes.route_destination.mirror_policy` | [routes.route_destination.mirror_policy](data-sources--route--reference--group-003.md#canonical-3012212223031212-2310112013100223-0103302112000003-1321112330111213-2022221331131203-2320313013030310-1221323010120003-3000300130110112) |
| `routes.route_destination.mirror_policy.cluster` | [routes.route_destination.mirror_policy.cluster](data-sources--route--reference--group-003.md#canonical-2333032232110233-1231100021031203-2133320101002202-1023200020210300-3313031000032200-3313031103001111-2201032122321320-1333032232332012) |
| `routes.route_destination.mirror_policy.cluster.kind` | [routes.route_destination.mirror_policy.cluster.kind](data-sources--route--reference--group-003.md#canonical-1333330032203031-0311112100110322-0320200012203021-3232122000001310-3033131200133132-3221021331202122-3331112333331100-0113333112102002) |
| `routes.route_destination.mirror_policy.cluster.name` | [routes.route_destination.mirror_policy.cluster.name](data-sources--route--reference--group-003.md#canonical-0311212211213111-0002033132023332-3000322320123311-0130122023200220-0233200123210023-0003210331102100-1000332101213131-1221113310213321) |
| `routes.route_destination.mirror_policy.cluster.namespace` | [routes.route_destination.mirror_policy.cluster.namespace](data-sources--route--reference--group-003.md#canonical-3212000122103123-2000213232202201-0232333323212122-2210202300201320-2131330332011231-3030321223332230-3103231003220133-3120213103223130) |
| `routes.route_destination.mirror_policy.cluster.tenant` | [routes.route_destination.mirror_policy.cluster.tenant](data-sources--route--reference--group-003.md#canonical-0233202213221102-1123313031222303-2232033020320011-3130323123021131-1111220023012030-0102131101232300-0103300013222013-3101022230132321) |
| `routes.route_destination.mirror_policy.cluster.uid` | [routes.route_destination.mirror_policy.cluster.uid](data-sources--route--reference--group-003.md#canonical-2202302113211221-3232020113200333-1220100223101002-0120003333201002-2111000013213132-1030121011132323-0201110133033203-3101131110103121) |
| `routes.route_destination.mirror_policy.percent` | [routes.route_destination.mirror_policy.percent](data-sources--route--reference--group-003.md#canonical-0012112212323330-0312130300111032-1211320130330222-1132223200230113-2323320222211032-3003323203003223-1113022031201121-1101200120102323) |
| `routes.route_destination.mirror_policy.percent.denominator` | [routes.route_destination.mirror_policy.percent.denominator](data-sources--route--reference--group-003.md#canonical-3031032003122003-0220033330032222-2113031220000311-1330022322202213-2030101100303013-3003000130000210-2332203302200113-2113130103231300) |
| `routes.route_destination.mirror_policy.percent.numerator` | [routes.route_destination.mirror_policy.percent.numerator](data-sources--route--reference--group-003.md#canonical-0102022313121132-2120020311213101-0212302221012011-2220232001001032-2020302202100333-1330221313010230-3321001002202302-2110312210230001) |
| `routes.route_destination.prefix_rewrite` | [routes.route_destination.prefix_rewrite](data-sources--route--reference--group-003.md#canonical-3323002322120103-1100303323310211-1001322230333213-2303112211320333-2200131301120012-1222320003002222-2031003021122002-1113333110203322) |
| `routes.route_destination.priority` | [routes.route_destination.priority](data-sources--route--reference--group-003.md#canonical-3033021010220123-0333220312323132-2132112322230112-0322221030111300-2220001112223212-0212333211012212-3130203320220013-0110333332131131) |
| `routes.route_destination.query_params` | [routes.route_destination.query_params](data-sources--route--reference--group-003.md#canonical-3321110120313220-2013020002030030-3011001131132323-1231023120231312-1013333001131123-0333102222303113-1323321101023310-3033020133300023) |
| `routes.route_destination.query_params.remove_all_params` | [routes.route_destination.query_params.remove_all_params](data-sources--route--reference--group-003.md#canonical-2101121003203123-0223021020032113-2011330012200012-1112330210122201-0123310011221321-1021000212230312-1300311210320321-3113000211110202) |
| `routes.route_destination.query_params.replace_params` | [routes.route_destination.query_params.replace_params](data-sources--route--reference--group-003.md#canonical-0203010002012011-0333230002222010-2023232102311120-2211330233113322-1331212322311002-1103013310322120-2022231122013220-1013322101231230) |
| `routes.route_destination.query_params.retain_all_params` | [routes.route_destination.query_params.retain_all_params](data-sources--route--reference--group-003.md#canonical-3110131320031132-3132130210012030-3101203030232302-0011131012011223-2303022312202101-3213312211231331-1021213213223021-0102223332230102) |
| `routes.route_destination.regex_rewrite` | [routes.route_destination.regex_rewrite](data-sources--route--reference--group-003.md#canonical-1113333100322133-0311223100133020-1312111323022322-3112013221213100-0132012310210201-0121122001313302-2311000331013232-1322213121333000) |
| `routes.route_destination.regex_rewrite.pattern` | [routes.route_destination.regex_rewrite.pattern](data-sources--route--reference--group-003.md#canonical-0221130301220333-3203200010020132-0012131302201110-0313131221100312-1231103120310021-2022103212312202-2301120220220012-2113001333130132) |
| `routes.route_destination.regex_rewrite.substitution` | [routes.route_destination.regex_rewrite.substitution](data-sources--route--reference--group-003.md#canonical-1200223033022111-3331121302202130-3320122212220003-2323300002020100-0113231123331033-3200123033103003-0230201033222133-1021321310023201) |
| `routes.route_destination.retract_cluster` | [routes.route_destination.retract_cluster](data-sources--route--reference--group-003.md#canonical-3001033031301301-2302130030020123-1022212302202131-0113301002112223-1223321322333122-1311001322132100-3321100203130133-2321133102001113) |
| `routes.route_destination.retry_policy` | [routes.route_destination.retry_policy](data-sources--route--reference--group-003.md#canonical-3100123200200022-0301001002022331-1320200312130020-2310213003202101-1322001202303011-1023021322132313-1330113100311330-1303130112011230) |
| `routes.route_destination.retry_policy.back_off` | [routes.route_destination.retry_policy.back_off](data-sources--route--reference--group-003.md#canonical-1233313102211331-3003211003132213-1031031220002231-3032021330100013-0131133100210112-2323312123000110-1130113221302320-3132321112331111) |
| `routes.route_destination.retry_policy.back_off.base_interval` | [routes.route_destination.retry_policy.back_off.base_interval](data-sources--route--reference--group-003.md#canonical-0300121220113300-1032120312211100-0010210002303111-2301220310131312-3113112100030023-3311302230222210-2010323331131312-0022111013101020) |
| `routes.route_destination.retry_policy.back_off.max_interval` | [routes.route_destination.retry_policy.back_off.max_interval](data-sources--route--reference--group-003.md#canonical-1320130323132312-3322222223222320-0000332110302322-3332120321021021-3200222131200300-3202020033210030-3331020000303330-3203032323021302) |
| `routes.route_destination.retry_policy.num_retries` | [routes.route_destination.retry_policy.num_retries](data-sources--route--reference--group-003.md#canonical-0212213312211330-0201013320113230-3011031203221021-2102321122023302-1013022030011203-3303332312132332-0332113120100002-0303021023002030) |
| `routes.route_destination.retry_policy.per_try_timeout` | [routes.route_destination.retry_policy.per_try_timeout](data-sources--route--reference--group-003.md#canonical-2102011133012302-0130330312323333-2200322110200233-1323320330000233-2032013220221302-2311100131313032-0323302013332131-2221012011113211) |
| `routes.route_destination.retry_policy.retriable_status_codes` | [routes.route_destination.retry_policy.retriable_status_codes](data-sources--route--reference--group-003.md#canonical-0101101120322203-3300200013002221-3023012203033301-3220012031012032-3221003120003210-1102133211030212-2330003033312130-2311220321320122) |
| `routes.route_destination.retry_policy.retry_condition` | [routes.route_destination.retry_policy.retry_condition](data-sources--route--reference--group-003.md#canonical-2231110123331132-3203021031101033-0131112023322130-0303303000320130-2301012201123310-3120222200321303-0131110022000300-0221102311202020) |
| `routes.route_destination.spdy_config` | [routes.route_destination.spdy_config](data-sources--route--reference--group-003.md#canonical-2220333321120322-1012311301331321-1222310330312303-0211211333230002-1211312321113233-2121331210102023-2211301303232213-0333222022011000) |
| `routes.route_destination.spdy_config.use_spdy` | [routes.route_destination.spdy_config.use_spdy](data-sources--route--reference--group-003.md#canonical-2300000302331011-2100120201023231-0330310011312311-3302233202003231-0031322101013331-1022030122131200-3010120231013300-0310201300003323) |
| `routes.route_destination.timeout` | [routes.route_destination.timeout](data-sources--route--reference--group-003.md#canonical-1022102013100110-0100022022001233-1021301310030120-0300100021032012-0032131220012321-0110003122310132-2123022131202200-3300130330010310) |
| `routes.route_destination.web_socket_config` | [routes.route_destination.web_socket_config](data-sources--route--reference--group-003.md#canonical-2111330201102302-2111201100010002-2010103013301033-3332110123332101-1001013130200012-0302033232323010-1010103000323230-1023130113102023) |
| `routes.route_destination.web_socket_config.use_websocket` | [routes.route_destination.web_socket_config.use_websocket](data-sources--route--reference--group-003.md#canonical-1213211200100330-1201003203223231-1301100011132011-3332333010001231-1222202321000103-0133312301202333-1002002233333331-2230211303311112) |
| `routes.route_direct_response` | [routes.route_direct_response](data-sources--route--reference--group-003.md#canonical-1212322131103131-3331322002112210-3302312013303101-2022332233000021-2331331310211102-2012232111300232-1232100300131230-3102300330203330) |
| `routes.route_direct_response.response_body_encoded` | [routes.route_direct_response.response_body_encoded](data-sources--route--reference--group-003.md#canonical-1122333030303213-2302103111311320-0330203100031100-2211211113110203-2030010322221212-1322231213011213-1011202022323233-0011313231001112) |
| `routes.route_direct_response.response_code` | [routes.route_direct_response.response_code](data-sources--route--reference--group-003.md#canonical-1331003122112003-2233330230331213-0213233122301301-0103212310112113-3030300132232231-2012203111312333-3231122332311232-0030210200011012) |
| `routes.route_redirect` | [routes.route_redirect](data-sources--route--reference--group-003.md#canonical-2102022120213232-1030011300302311-0311321010220223-3200101231120130-2100103001102211-2012101031333022-3332232323210112-3300121301223120) |
| `routes.route_redirect.host_redirect` | [routes.route_redirect.host_redirect](data-sources--route--reference--group-003.md#canonical-1333112031333303-2102222201203013-1320032333212121-2103101132312333-1320303130132010-3101212013012312-1332113113101102-0100121310012130) |
| `routes.route_redirect.path_redirect` | [routes.route_redirect.path_redirect](data-sources--route--reference--group-003.md#canonical-0203021000131202-0301230010213332-2133320320301333-0123332320131101-3132213300232012-0302201202023130-0033120012000311-1200130121130130) |
| `routes.route_redirect.prefix_rewrite` | [routes.route_redirect.prefix_rewrite](data-sources--route--reference--group-003.md#canonical-0223302113020133-3311020313211320-1023233220001201-3301213001122110-0032031001220231-1300010032031003-3032002220021330-1323213102301232) |
| `routes.route_redirect.proto_redirect` | [routes.route_redirect.proto_redirect](data-sources--route--reference--group-003.md#canonical-0301210131233012-2311101233232210-1222320003203320-3230132112222223-2003310120202031-0132232130322320-2210021103001221-2001332211010213) |
| `routes.route_redirect.remove_all_params` | [routes.route_redirect.remove_all_params](data-sources--route--reference--group-003.md#canonical-2032120322132110-1000112332320112-0310031222110120-2202130021203033-2122032200112200-2233012000301212-2303221033021232-1011122110332110) |
| `routes.route_redirect.replace_params` | [routes.route_redirect.replace_params](data-sources--route--reference--group-003.md#canonical-1110232033212233-2302201030122003-3010120323123111-1323212221231220-0220223211022233-3231330233133030-1012211130301111-2133021130002133) |
| `routes.route_redirect.response_code` | [routes.route_redirect.response_code](data-sources--route--reference--group-003.md#canonical-2102333210131021-1012211112331000-0212111313222313-0122033201002133-1332113322023102-3102030000211222-3223233301203022-2121222332120001) |
| `routes.route_redirect.retain_all_params` | [routes.route_redirect.retain_all_params](data-sources--route--reference--group-003.md#canonical-2033321200101322-1032113232103113-3000112022122030-3230131201303010-1010030033011233-1121222123100323-1133313020021332-3231311103130002) |
| `routes.service_policy` | [routes.service_policy](data-sources--route--reference--group-003.md#canonical-0120211110302033-0112203233032231-3310003010012113-3123123220311133-2311331120231020-2023310002002100-1322033322100030-2211011113112032) |
| `routes.service_policy.disable_spec` | [routes.service_policy.disable_spec](data-sources--route--reference--group-003.md#canonical-0103020030002230-3110133023123232-1123201313310303-1223212303110230-0012032210122132-3301011300031101-1201321010100311-2022202212311012) |
| `routes.waf_exclusion_policy` | [routes.waf_exclusion_policy](data-sources--route--reference--group-003.md#canonical-0102021320210111-2132312231202321-3013023220302010-0003020210112022-2020200210101302-2333300233103113-1131321320110012-3110003001331232) |
| `routes.waf_exclusion_policy.name` | [routes.waf_exclusion_policy.name](data-sources--route--reference--group-003.md#canonical-1330111020203003-3313231120300213-0133202300130220-1233231202100320-1133331011211112-3122010203130312-2113030030332303-0002123321223002) |
| `routes.waf_exclusion_policy.namespace` | [routes.waf_exclusion_policy.namespace](data-sources--route--reference--group-003.md#canonical-1230223032103331-2131013110202030-1010203010332101-3321203033111002-3113222200013311-1210031301322010-1233201322223302-1330022223222001) |
| `routes.waf_exclusion_policy.tenant` | [routes.waf_exclusion_policy.tenant](data-sources--route--reference--group-003.md#canonical-0001332303212101-0002322132021222-2310330202221010-1221121012011120-1103321233231200-1120033111330313-2231210122233322-1123101111220111) |
| `routes.waf_type` | [routes.waf_type](data-sources--route--reference--group-003.md#canonical-0000203232210111-1312212321222112-1133232323000121-0223302032303213-3021033121313200-1330222203301112-2110020213030231-1033301211332320) |
| `routes.waf_type.app_firewall` | [routes.waf_type.app_firewall](data-sources--route--reference--group-003.md#canonical-0020231331220323-3031331033100200-1120333103130222-1211103132113331-3010320123323121-0313201322312212-2010031112310211-0330130310103330) |
| `routes.waf_type.app_firewall.app_firewall` | [routes.waf_type.app_firewall.app_firewall](data-sources--route--reference--group-003.md#canonical-3333133223101013-2213111111313231-0012020331031330-0223031321232131-3303022200200131-0330201332023012-1100001222032202-2202330011023331) |
| `routes.waf_type.app_firewall.app_firewall.kind` | [routes.waf_type.app_firewall.app_firewall.kind](data-sources--route--reference--group-003.md#canonical-2121213330030331-2200332111100100-3121131311211211-3010133023032333-1101300010220200-3232333101212131-1302231301032323-1131031110231221) |
| `routes.waf_type.app_firewall.app_firewall.name` | [routes.waf_type.app_firewall.app_firewall.name](data-sources--route--reference--group-003.md#canonical-3110120212000211-0221110303020000-0330132003200110-1302313121021130-1011330120000230-2121122333010111-1021321310322203-1120333002303213) |
| `routes.waf_type.app_firewall.app_firewall.namespace` | [routes.waf_type.app_firewall.app_firewall.namespace](data-sources--route--reference--group-003.md#canonical-2112212222101111-2302011323133332-1030330011000032-3312203323222110-0322230312033001-1333032110211100-1132322323130120-2132302230320110) |
| `routes.waf_type.app_firewall.app_firewall.tenant` | [routes.waf_type.app_firewall.app_firewall.tenant](data-sources--route--reference--group-003.md#canonical-0232130132102221-1230300033232101-1300023002221300-3223300100330332-0230020201323120-0331232021101021-1102021222311321-3211122332022312) |
| `routes.waf_type.app_firewall.app_firewall.uid` | [routes.waf_type.app_firewall.app_firewall.uid](data-sources--route--reference--group-003.md#canonical-2300302123331103-0220313323130323-2210230032300213-3123223001320032-1213003303311121-1132023111130221-2200123202333313-2320300033301002) |
| `routes.waf_type.disable_waf` | [routes.waf_type.disable_waf](data-sources--route--reference--group-003.md#canonical-1311301330230310-3322000023323223-0231222322120003-0133003210322311-1030011031130121-2212310322320133-3203333030201113-1313023003101221) |
| `routes.waf_type.inherit_waf` | [routes.waf_type.inherit_waf](data-sources--route--reference--group-003.md#canonical-0113313123212030-3112020123111331-0123333101202012-1330102303203232-1011233123333203-2330313020102010-0120222220011000-0213232212330100) |

<a id="canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- routes

<a id="canonical-0320012033211300-0233121012310301-0320121012020311-3003231220302330-3121322100330222-2033320013310023-0301031010300221-0122333021323232"></a>

Type: `"list"`. Computed.

List of routes to match for incoming request.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 257,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 257,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "257"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "257"
  }
}
```

<a id="canonical-2303232230111331-0023010132200323-3202002223101110-3311002011021030-3330310332210233-2332121130220002-3032310213300212-3020310211312321"></a>

### Direct properties for `routes`

- [bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-3002323020322233-1130103332320120-0230113310310331-0300223313301321-1303010011012203-1121121321123311-1013021211210002-2213313033310202): complete subsection reference.

<a id="canonical-1321310200023001-2332033133130021-3031320133013231-2010000001000203-1100303301012321-2010030201211110-0031322102002302-0233230103313220"></a>

<a id="canonical-1023001233120322-1332202113201021-0031132323013031-0313202011000120-2301021111010301-3202320021221320-1332113312203222-3313322213011121"></a>

#### `routes.disable_location_add` property

Type: `"bool"`. Computed.

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

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

- [inherited_bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-1303022010313313-2131031121103111-3323301113200323-0102212203023013-3110203122223210-0333010332212223-0310130110332311-0220100331322200): complete subsection reference.

- [inherited_waf_exclusion](data-sources--route--reference--group-001.md#canonical-3020000312033221-3002122322300202-2223000211301020-3123013233103202-3302221323011013-0200312133131211-1323123201303312-1120301031010313): complete subsection reference.

- [match](data-sources--route--reference--group-001.md#canonical-2202021113211000-0321023232131332-0023133030000100-0113121112201031-1323001231200003-1323310313022102-0010110103313023-3023113002113312): complete subsection reference.

- [request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-3332003033002121-1112131201111132-3001303023032232-2113101310100210-2012310213112313-1121123203013213-1000102220303021-3230333330312222): complete subsection reference.

<a id="canonical-3033022302301101-1023101032231000-2132023011100330-0301320202033302-3311210302210033-0121213201300000-1202102333103011-1112130111221132"></a>

<a id="canonical-0123332213010011-0001302200132001-3200213101233312-0210312223303000-1322011313101210-1012102110121021-1302333213211001-0313331033311012"></a>

#### `routes.request_cookies_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](data-sources--route--reference--group-002.md#canonical-0020121133012230-3021000322312112-2000301212110101-0232120123123121-1233000233203221-2213312130233232-3120100112111032-1021020312231033): complete subsection reference.

<a id="canonical-2110331033213021-1001030033020203-2033201231321201-3210111101012023-3100211211122023-1320032212221311-0333023031230203-2301032112010212"></a>

<a id="canonical-3332332211130133-2102020033302302-2121302122323120-0310302220012010-3333231322100121-0213321112323031-1222001333212102-3130202102331033"></a>

#### `routes.request_headers_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031): complete subsection reference.

<a id="canonical-2122001010102022-3212130023313021-3103320112003010-0120130000010032-3210123121313301-1130010200132133-2302010223133103-1120232203202133"></a>

<a id="canonical-1121111100311121-2331303110232321-3321221102322021-1202222333013211-3113102003201013-1111031122223313-0322102002030321-0323031002130232"></a>

#### `routes.response_cookies_to_remove` property

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--route--reference--group-002.md#canonical-2031132012221313-2231011210002013-3003320011222122-2113303301333231-3130312023212100-1221120320210322-1210200320102232-3230120313103301): complete subsection reference.

<a id="canonical-0100323310101123-1200301200110233-0113332200103222-1120112302030131-3022202033232021-2131001000111011-2301121301002022-0200000300023313"></a>

<a id="canonical-1120121321111121-0131303120011301-3313203113102120-2033233222101213-2011300310202333-3122031103122101-1000323101310200-0332322310132001"></a>

#### `routes.response_headers_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [route_destination](data-sources--route--reference--group-002.md#canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230): complete subsection reference.

- [route_direct_response](data-sources--route--reference--group-003.md#canonical-1120103010111221-1131011202033302-3012012120230213-0222131120131102-2101233332131203-3211033100232203-0013202020220303-2233021120011122): complete subsection reference.

- [route_redirect](data-sources--route--reference--group-003.md#canonical-3031003311233213-2113111021311003-2111031332300130-1111311130312100-2120221111012231-3331330300223001-2023020031123100-0201231213001103): complete subsection reference.

- [service_policy](data-sources--route--reference--group-003.md#canonical-0122003031122022-2312121320313011-3232203000312111-2300213223121003-3333300031112101-3031102112131312-1230313311203332-3230301311131320): complete subsection reference.

- [waf_exclusion_policy](data-sources--route--reference--group-003.md#canonical-2130023122111302-3312131310223330-3111030320001222-1300131123332221-1300033010300132-3331301213032123-3013023033221232-0223231010031303): complete subsection reference.

- [waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120): complete subsection reference.

<a id="canonical-3002323020322233-1130103332320120-0230113310310331-0300223313301321-1303010011012203-1121121321123311-1013021211210002-2213313033310202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.bot_defense_javascript_injection` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.bot_defense_javascript_injection

<a id="canonical-0021032101233101-3323232033321002-2323001110301313-2030030301033320-3021033010012302-0111031332120230-3210031000323200-1111032112032212"></a>

Type: `"single"`. Computed.

Bot Defense JavaScript Injection Configuration for inline bot defense deployments.

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

<a id="canonical-1210020001231210-1203132200311102-1012002301233221-2313012303223103-0302313032003213-3323213321222011-2113030101101211-2312220321022110"></a>

### Direct properties for `routes.bot_defense_javascript_injection`

<a id="canonical-3201222233321221-2001032000100233-3232100020000303-1333112022223203-0101030132320100-3102233130111233-0302213133103300-3101103120210230"></a>

#### `routes.bot_defense_javascript_injection.javascript_location` property

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [javascript_tags](data-sources--route--reference--group-001.md#canonical-0032020112201122-1012130210333101-0230310321112302-3001301011031320-1330333031033033-3002212331120302-3202100220110013-3210300021012002): complete subsection reference.

<a id="canonical-0032020112201122-1012130210333101-0230310321112302-3001301011031320-1330333031033033-3002212331120302-3202100220110013-3210300021012002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.bot_defense_javascript_injection.javascript_tags` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-3002323020322233-1130103332320120-0230113310310331-0300223313301321-1303010011012203-1121121321123311-1013021211210002-2213313033310202)
- routes.bot_defense_javascript_injection.javascript_tags

<a id="canonical-1131200010232100-0231320132322302-2303020203123232-0101023012100122-1022131033230032-3010000230112133-2333111113220010-2313000312110323"></a>

Type: `"list"`. Computed.

Select Add item to configure your JavaScript tag. If adding both Bot Adv and Fraud, the Bot
JavaScript should be added first.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3011132321231312-0002220131133121-0301330100203110-3120113100332132-0031223203013001-3131103120332330-1103111323112020-1220222301323033"></a>

### Direct properties for `routes.bot_defense_javascript_injection.javascript_tags`

<a id="canonical-0311333322120220-1130113003102123-2222311121132102-2233110131220130-0201122331112213-1003011210202122-3322013120313212-3302102111322122"></a>

#### `routes.bot_defense_javascript_injection.javascript_tags.javascript_url` property

Type: `"string"`. Computed.

Please enter the full URL (include domain and path), or relative path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 2048,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tag_attributes](data-sources--route--reference--group-001.md#canonical-2131210232233001-0332212010223020-3103022313003020-1331311221202031-3330101313013002-3231212212132323-0330203211012311-0321003012023313): complete subsection reference.

<a id="canonical-2131210232233001-0332212010223020-3103022313003020-1331311221202031-3330101313013002-3231212212132323-0330203211012311-0321003012023313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.bot_defense_javascript_injection](data-sources--route--reference--group-001.md#canonical-3002323020322233-1130103332320120-0230113310310331-0300223313301321-1303010011012203-1121121321123311-1013021211210002-2213313033310202)
- [routes.bot_defense_javascript_injection.javascript_tags](data-sources--route--reference--group-001.md#canonical-0032020112201122-1012130210333101-0230310321112302-3001301011031320-1330333031033033-3002212331120302-3202100220110013-3210300021012002)
- routes.bot_defense_javascript_injection.javascript_tags.tag_attributes

<a id="canonical-3132333132102112-0022101033131332-3332122002012100-0321332000200323-1003233010212032-0013000202320330-1101033233210220-3233000021132232"></a>

Type: `"list"`. Computed.

Add the tag attributes you want to include in your JavaScript tag.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3213133231212331-2020030211022021-1300213300332320-3001002023103321-0231000203231113-1310200333233103-3020332313101321-0301212022130001"></a>

### Direct properties for `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes`

<a id="canonical-1331003230133230-0203132020200310-3302331312313221-3103210120100330-0102332111312303-1011320020000011-1010331122120003-2322330030200222"></a>

#### `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag` property

Type: `"string"`. Computed.

\[Enum:
JS\_ATTR\_ID|JS\_ATTR\_CID|JS\_ATTR\_CN|JS\_ATTR\_API\_DOMAIN|JS\_ATTR\_API\_URL|JS\_ATTR\_API\_PATH|JS\_ATTR\_ASYNC|JS\_ATTR\_DEFER\]
Select from one of the predefined tag attributes. Possible values are \`JS\_ATTR\_ID\`,
\`JS\_ATTR\_CID\`, \`JS\_ATTR\_CN\`, \`JS\_ATTR\_API\_DOMAIN\`, \`JS\_ATTR\_API\_URL\`,
\`JS\_ATTR\_API\_PATH\`, \`JS\_ATTR\_ASYNC\`, \`JS\_ATTR\_DEFER\`. Defaults to \`JS\_ATTR\_ID\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "JS_ATTR_ID",
  "enum": [
    "JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2130320221122202-2212210203110210-0322110202030132-1120000010232301-0001333110101321-0301300332200200-3123303111002220-1311113331030123"></a>

<a id="canonical-2102303212110310-1020022002313211-1231212210231011-0320330132123110-3113113212013311-3011001123023100-0320031232012200-1233120232313012"></a>

#### `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value` property

Type: `"string"`. Computed.

Value. Add the tag attribute value.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-1303022010313313-2131031121103111-3323301113200323-0102212203023013-3110203122223210-0333010332212223-0310130110332311-0220100331322200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.inherited_bot_defense_javascript_injection` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.inherited_bot_defense_javascript_injection

<a id="canonical-3223200301103331-3023232320321133-1122331213230113-0312202233121010-3331332123003221-2032103211300311-3232130213330300-3213200333033122"></a>

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

<a id="canonical-3020000312033221-3002122322300202-2223000211301020-3123013233103202-3302221323011013-0200312133131211-1323123201303312-1120301031010313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.inherited_waf_exclusion` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.inherited_waf_exclusion

<a id="canonical-0002221303113200-3221213302332021-3202213312032222-2201223023231320-3221203001320122-0201212310231010-3310300101321231-3031121123031121"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherited waf exclusion.

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

<a id="canonical-2202021113211000-0321023232131332-0023133030000100-0113121112201031-1323001231200003-1323310313022102-0010110103313023-3023113002113312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.match

<a id="canonical-0013130210101202-2031231200311303-0112132131220101-0032301000312111-2321230213130321-2230010011222120-3101021201112300-0210233330020000"></a>

Type: `"list"`. Computed.

Match. Route match condition.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3232011302032111-3312321332022323-1032012213213232-3330322321120323-3201301333132330-2211111312100001-0023120230122231-1112130101310020"></a>

### Direct properties for `routes.match`

- [headers](data-sources--route--reference--group-001.md#canonical-2321232120032120-0200213000010321-2133321111300200-0223030102012020-2300113333311221-0110102303332033-3320223203213001-3322020331231112): complete subsection reference.

<a id="canonical-2232300131002221-2202101130132111-3131111330011213-1011230121001300-0210203231312020-3120203202333232-0123312330122202-1311200112000312"></a>

<a id="canonical-3010321003320110-0210303323022220-3201213303120313-2232101030132112-0013231220222222-3112022033313123-3000321102212222-0313303110333101"></a>

#### `routes.match.http_method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--route--reference--group-001.md#canonical-2031100000103321-3333121232221220-1220133121020202-2312110311203103-0020033033301030-0122103313301203-1003020212133202-1011022320213220): complete subsection reference.

- [path](data-sources--route--reference--group-001.md#canonical-1301102313212233-1120200113202122-3202022223333200-2310300311201032-0021311002021121-3033220033232211-2230212033100002-0132103320200311): complete subsection reference.

- [query_params](data-sources--route--reference--group-001.md#canonical-0213023313013322-0311023222011020-1120220323022332-3230323100313320-3202111122112100-2322331010113121-1301202323012023-3233012302120223): complete subsection reference.

<a id="canonical-2321232120032120-0200213000010321-2133321111300200-0223030102012020-2300113333311221-0110102303332033-3320223203213001-3322020331231112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.headers` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.match](data-sources--route--reference--group-001.md#canonical-2202021113211000-0321023232131332-0023133030000100-0113121112201031-1323001231200003-1323310313022102-0010110103313023-3023113002113312)
- routes.match.headers

<a id="canonical-3323230031312220-2200333123130100-3103323003322212-0132302322011133-3002100222303231-3313331313010320-3132333200033202-2322123321030221"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

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
    "minItems": 0,
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

<a id="canonical-1010220001302233-2122131201133120-1131030122300123-1303111302112020-3120311011332330-1323113233002211-1031020121003021-1213120120202303"></a>

### Direct properties for `routes.match.headers`

<a id="canonical-2310311221020233-1222020220110320-2133300322003323-3322333222101123-3303213031333012-1221211320122113-1211213222331112-1212331123103133"></a>

#### `routes.match.headers.exact` property

Type: `"string"`. Computed.

Exclusive with \[presence regular expression\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-0101213120100120-3230003220303003-0131123000033202-0310130320022002-1220213211312012-3231023302232212-0202232122302303-0222021112302033"></a>

<a id="canonical-0302033220002233-0111131121000311-0002210111222221-2311233110323213-1001110133212202-1311103311302211-0003013322332302-2211301111331332"></a>

#### `routes.match.headers.invert_match` property

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-0032013013131033-0122213302303032-3210200003102310-2213000112130010-2101102212131132-1023032130022011-2202311131313100-1112121100133300"></a>

<a id="canonical-0120303101103012-0132112132021133-3012201010322101-1213300103322032-0121203121331203-0112210330113212-1223103013122102-0323003211312302"></a>

#### `routes.match.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0130301021233311-2332132320002331-3132031002130201-3112230320231303-2121231222230300-0123303213312212-2001310011013233-1031332321131333"></a>

<a id="canonical-0031200111010201-0221112012012322-3010312210322013-1102023332331003-2231033021121302-3020332003203300-1003010203001000-3003231323131302"></a>

#### `routes.match.headers.presence` property

Type: `"bool"`. Computed.

Exclusive with \[exact regular expression\] If true, check for presence of header.

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

<a id="canonical-3222333223001120-3010011201303332-1201333013312313-1003303131120012-2022333222001113-0220110132012120-2223101011232132-2213210013113203"></a>

<a id="canonical-1200302230120003-3113200221203332-0212100300303002-0232313000320222-2101020132331222-0233223210031130-3200302333302220-1000001203321311"></a>

#### `routes.match.headers.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2031100000103321-3333121232221220-1220133121020202-2312110311203103-0020033033301030-0122103313301203-1003020212133202-1011022320213220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.incoming_port` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.match](data-sources--route--reference--group-001.md#canonical-2202021113211000-0321023232131332-0023133030000100-0113121112201031-1323001231200003-1323310313022102-0010110103313023-3023113002113312)
- routes.match.incoming_port

<a id="canonical-3020231022012021-3011123331001312-2012203112303122-3313212132120201-2311311221123210-3131312103303303-1030223203021123-1030123002332122"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-2002033111003302-3211121212312123-0112000000102132-3113313200002132-2120220211110130-3003103222001221-1211001123100013-2310303311103033"></a>

### Direct properties for `routes.match.incoming_port`

- [no_port_match](data-sources--route--reference--group-001.md#canonical-2302301103330013-2212011312320322-1020111103202001-1033011310200011-3233232132213100-3211022203102002-0113030120322330-1122302321131330): complete subsection reference.

<a id="canonical-3202302213032302-2303131203103323-3031330301211312-0102211132010130-3221032132030030-0220321120220313-1013111112001311-3323233202331121"></a>

<a id="canonical-2220102221021311-0100310231022100-3231231310323232-0120201221321101-0333220331230333-2302222011102313-3231330322311133-0332001230032333"></a>

#### `routes.match.incoming_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2132133312222012-0100310313322112-0332130212001122-2233233121102320-1321032300222033-1111210221202312-0323232110203021-0022200331330102"></a>

<a id="canonical-2130101101201031-1320220303332032-1102323000303030-3101011232203331-0110110200122110-0212322312320333-0120303003101223-3022133310002002"></a>

#### `routes.match.incoming_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-2302301103330013-2212011312320322-1020111103202001-1033011310200011-3233232132213100-3211022203102002-0113030120322330-1122302321131330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.match](data-sources--route--reference--group-001.md#canonical-2202021113211000-0321023232131332-0023133030000100-0113121112201031-1323001231200003-1323310313022102-0010110103313023-3023113002113312)
- [routes.match.incoming_port](data-sources--route--reference--group-001.md#canonical-2031100000103321-3333121232221220-1220133121020202-2312110311203103-0020033033301030-0122103313301203-1003020212133202-1011022320213220)
- routes.match.incoming_port.no_port_match

<a id="canonical-1213112211110031-1032230031310322-2003322003121121-2102200230320011-3111100020222320-1112210233101001-0031013132022223-1223323000023112"></a>

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

<a id="canonical-1301102313212233-1120200113202122-3202022223333200-2310300311201032-0021311002021121-3033220033232211-2230212033100002-0132103320200311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.path` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.match](data-sources--route--reference--group-001.md#canonical-2202021113211000-0321023232131332-0023133030000100-0113121112201031-1323001231200003-1323310313022102-0010110103313023-3023113002113312)
- routes.match.path

<a id="canonical-3002112200132320-1123033113020333-1320211330313202-3002300002201010-3131133221132011-3301200330010321-0312102222020211-0002100333101203"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-1203201023122112-1013300202202221-2302122320133102-2031331120320000-0220030011131232-3023211213201222-2121320312022321-1323033303123100"></a>

### Direct properties for `routes.match.path`

<a id="canonical-0020110020001231-0012301101011231-0001001123310233-0112223113032200-0123020320232100-1213212220320123-0302220022332230-1023013010033101"></a>

#### `routes.match.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3101312232311100-1300110100020312-3232103012020101-3223210010321311-0300120031023010-2210021102333122-3131310201001331-2001003003220003"></a>

<a id="canonical-1313012300033132-3001032222111131-3101230312303111-1120132101000031-3000102002121221-0032111011130130-0330001133031123-1010133200033320"></a>

#### `routes.match.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0113233220320010-0132221113332130-0333302313112131-0123122003321021-0122200312220321-1001023213032102-0201212231310123-2112233103312222"></a>

<a id="canonical-0003221003030322-1111012202223220-3121110020310000-0032301201002232-2112020110230311-1002332111030103-0302213202132023-2122112132030332"></a>

#### `routes.match.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0213023313013322-0311023222011020-1120220323022332-3230323100313320-3202111122112100-2322331010113121-1301202323012023-3233012302120223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.query_params` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.match](data-sources--route--reference--group-001.md#canonical-2202021113211000-0321023232131332-0023133030000100-0113121112201031-1323001231200003-1323310313022102-0010110103313023-3023113002113312)
- routes.match.query_params

<a id="canonical-1202202201102123-3320301132300112-3322331022310212-3121030231331300-1232130001230003-2311203130302223-0310300332210113-1031333203022013"></a>

Type: `"list"`. Computed.

Query Parameters. List of (key, value) query parameters.

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

<a id="canonical-2232031321021131-0122023023121330-2312120200000301-1013132113033123-1212232112310213-3321122033123310-1223213023132312-2133021203103300"></a>

### Direct properties for `routes.match.query_params`

<a id="canonical-2010232202021031-0321023112120020-1300230112102033-0020212132222201-0102131313231322-0011111012313020-0101222320001010-2131030320133202"></a>

#### `routes.match.query_params.exact` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\] Exact match value for the query parameter key.

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
  }
}
```

<a id="canonical-3300122321000130-1313001023302210-1121002320211010-2003001012132112-0313310330022312-3120303330223022-0233213313332031-1212011330123210"></a>

<a id="canonical-3113212313232011-2010222022312123-2202100031302211-2201023203022222-1111123013230201-3012320000013122-1112220120310110-1110203101330301"></a>

#### `routes.match.query_params.key` property

Type: `"string"`. Computed.

Query parameter key In the above example, assignee\_username is the key.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3320201203203202-1310222123110001-1331331031103101-0010033201100032-2321100003233002-2003000112200101-3113112320212330-3302220221232223"></a>

<a id="canonical-3331203312010203-3303311002001232-3021010023302132-2210023000233021-1022033023231012-0123322112201233-0000110030001310-1311312213230203"></a>

#### `routes.match.query_params.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact\] regular expression match value for the query parameter key.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3332003033002121-1112131201111132-3001303023032232-2113101310100210-2012310213112313-1121123203013213-1000102220303021-3230333330312222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.request_cookies_to_add

<a id="canonical-1223121012233330-3231123230130033-0311021022223133-2023213011131122-2310310020302220-0200323133031311-2302000111001021-3001203022333220"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3211330212201021-1310221021023302-3100120300032330-1313212033000303-0300303212303110-0231322100302032-0203331010310000-3011023101110221"></a>

### Direct properties for `routes.request_cookies_to_add`

<a id="canonical-1012110131222200-2333121000112002-3013322100313301-2301331322201133-2321331003023300-3312232110320002-3111120120302011-3322300122330100"></a>

#### `routes.request_cookies_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2113200121312011-2122212302313323-3331331301110032-2230132213013121-3311022020023221-3313201121222021-0201123300123223-3133311320111332"></a>

<a id="canonical-3330232323102022-1011220221310201-1331000101300112-0121102231210320-3033312303210010-0013002012130030-1202303310030011-0113200313012231"></a>

#### `routes.request_cookies_to_add.overwrite` property

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

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

- [secret_value](data-sources--route--reference--group-001.md#canonical-3202123231212200-3302310010022300-2221120133320331-1121100201331312-1311331222331110-0220110130212213-2203001222221120-3100023020210120): complete subsection reference.

<a id="canonical-0022021021001130-0001131223000102-1202210022102001-3310003320130300-3112121301203010-0031310101130223-2102111300232031-0231310003221131"></a>

<a id="canonical-0321012113133322-0033000102130001-1313333030012131-1221101020330220-2130202230212023-3002323320332022-0021233122002022-1200010021110112"></a>

#### `routes.request_cookies_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3202123231212200-3302310010022300-2221120133320331-1121100201331312-1311331222331110-0220110130212213-2203001222221120-3100023020210120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-3332003033002121-1112131201111132-3001303023032232-2113101310100210-2012310213112313-1121123203013213-1000102220303021-3230333330312222)
- routes.request_cookies_to_add.secret_value

<a id="canonical-3303132222213001-3103200013311031-0012233130013322-1303013020012300-2101220003113312-1133032100112203-1003330211322321-3302213311022301"></a>

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

<a id="canonical-3321230102233210-1000121333313200-3230032132012213-0223001223212013-1123330112000321-3312032132220330-1010112311003011-3101200112120311"></a>

### Direct properties for `routes.request_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--route--reference--group-001.md#canonical-0023311012200030-2113331020223133-3013122100331220-3131312122013123-3100031111201123-1100110123022300-3231032302311103-2210332100132110): complete subsection reference.

- [clear_secret_info](data-sources--route--reference--group-001.md#canonical-3211112030130310-3323201221232121-0300123002000220-1032023320120033-1032012323002333-3201312003232221-2013312112123011-2022200030312312): complete subsection reference.

<a id="canonical-0023311012200030-2113331020223133-3013122100331220-3131312122013123-3100031111201123-1100110123022300-3231032302311103-2210332100132110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-3332003033002121-1112131201111132-3001303023032232-2113101310100210-2012310213112313-1121123203013213-1000102220303021-3230333330312222)
- [routes.request_cookies_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-3202123231212200-3302310010022300-2221120133320331-1121100201331312-1311331222331110-0220110130212213-2203001222221120-3100023020210120)
- routes.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-2203111113222230-1000203001203202-0331001102331103-3020021002231030-0103202212122321-0212131003212123-2303332330013311-2000211300230323"></a>

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

<a id="canonical-3231222322011223-3000122002100123-1302130120032130-0331020023320013-1111331000112212-0231020322033322-2010312320302132-1101220112022223"></a>

### Direct properties for `routes.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1030000301013130-0113101231100132-3002012331303131-2020132012330011-0032031331332003-2120023033033102-3101333002131323-2131102123120320"></a>

#### `routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3120102230233230-2222312102221012-0322103331233132-2021013311222010-3033220330023131-2213310030022030-0130310220320322-3220123313200231"></a>

<a id="canonical-3300302220011312-0121112312123201-1211101313033211-3331030202321231-0001020102133113-2023123222123320-0320232110233233-0020302111021021"></a>

#### `routes.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1032330223212221-3202222131311123-1331011011211030-0113322003313030-0122132033200111-2231223020110011-3000221100201102-3130103013202111"></a>

<a id="canonical-3013212203121030-3131301110121111-0132310303033212-3230133303003230-1320033322003013-1101033201212031-1222132220222303-3110222101131322"></a>

#### `routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3211112030130310-3323201221232121-0300123002000220-1032023320120033-1032012323002333-3201312003232221-2013312112123011-2022200030312312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.request_cookies_to_add](data-sources--route--reference--group-001.md#canonical-3332003033002121-1112131201111132-3001303023032232-2113101310100210-2012310213112313-1121123203013213-1000102220303021-3230333330312222)
- [routes.request_cookies_to_add.secret_value](data-sources--route--reference--group-001.md#canonical-3202123231212200-3302310010022300-2221120133320331-1121100201331312-1311331222331110-0220110130212213-2203001222221120-3100023020210120)
- routes.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-1231022201203102-2201022221231330-1331102232122310-1310023310012322-2003330201012210-0312020112311122-2112000303011013-3332312200330102"></a>

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

<a id="canonical-1032300013110231-3133320102030231-1323103023313130-2113031322101013-1120202212202310-1033323320312301-1311002211220023-2231001131002001"></a>

### Direct properties for `routes.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0112002312332131-2132312212203003-3332110112022022-1211300121313023-0323000002123022-0331200323210123-3210332111321102-3113332212020330"></a>

#### `routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0232233300300022-1010210121033022-0002332111112203-2212322122320323-3031220110222103-0132013231122002-1032323310122121-2021213030003033"></a>

<a id="canonical-3123100022022033-1121131030232211-1030030111311112-0030100302112222-1321112103110010-1232000213003202-1031333300211022-0100132223001122"></a>

#### `routes.request_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
