---
page_title: "xcsh_forward_proxy_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy reference."
---

# xcsh_forward_proxy_policy reference

<a id="canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010223113223030-0200022032020301-1130332322132112-3321312110232333-0322330201222222-1200031210232300-3102131311020112-0132220212213003"></a>

## Property reference — Property reference / 210323300001 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- Property reference

<a id="canonical-3323100201012012-1132113123311021-0212132311033220-2232000011123202-2200020100312301-0133002333222331-2222321212231302-2031030201302010"></a>

## Direct properties — Property reference / 210323300001 / 3

- [allow_all](resources--forward_proxy_policy--reference--group-001.md#canonical-2220212320111303-0320122201200113-3102201023111301-2113200301102333-3233010100233033-3322330201020030-2012231030013231-1231233310333033): complete subsection reference.

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302): complete subsection reference.

<a id="canonical-2023212020311303-2101301302220330-0202033010331321-3211033002011320-3231332200033320-0011231123301133-1002300221301303-3203210321230212"></a>

<a id="canonical-0101200222111311-2311221032232112-3001202023130002-1122301323220010-1233313110222300-2002210132023021-3323231310130001-3221232333002211"></a>

## annotations property — Property reference / 210323300001 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [any_proxy](resources--forward_proxy_policy--reference--group-001.md#canonical-2332333300301220-2020011003301331-0311130310031221-0302003113033233-1113111033022331-1120333131201021-0101002021020003-0020122103301200): complete subsection reference.

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110): complete subsection reference.

<a id="canonical-2101010223032021-2222013022222022-1101001210111111-0333021103300203-0002302133230303-3201031110023312-0121111012210020-1220110332110012"></a>

<a id="canonical-2022132321313303-0313302131022012-3212100330123213-2113211220130330-0133102131101023-3133133130102000-2010132213220310-1300210331000032"></a>

## description property — Property reference / 210323300001 / 5

Type: `"string"`. Optional.

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

<a id="canonical-1312002302010300-3102223220021302-2132203220103233-3212012233311012-0112020212330330-0301002210111012-1231301303002102-2033010223012311"></a>

<a id="canonical-3210132131323232-3102121312021332-1301302221200033-2000221020133213-3022102203112111-3010210301112032-2200323113222022-1020021213122022"></a>

## disable property — Property reference / 210323300001 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [drp_http_connect](resources--forward_proxy_policy--reference--group-001.md#canonical-0121000321310323-0220232133302100-1210330110221132-1112130320100200-1022002100133130-3210200311003310-3011210001322311-1121030003102231): complete subsection reference.

<a id="canonical-1003330020213121-3212233011220201-3211301103003310-3202133311312102-1003303132112321-2300210331311100-1022313020312330-1322203323220012"></a>

<a id="canonical-2233330311320310-0203030030002122-3212131333232322-2310210211000200-1113033332101012-3000203003003113-3020212322301213-2200111231121012"></a>

## ID property — Property reference / 210323300001 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0132002111230200-1212130233332322-1012123300210112-0022103023311012-0230121331120220-1203010333302020-2021110013232001-0230330030222111"></a>

<a id="canonical-0110323223033000-1100013333110302-2121011220213013-3330201031231220-2222230222220012-1102202213330112-3013223212330223-0030023020122213"></a>

## labels property — Property reference / 210323300001 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-3003211310103233-0101302332032103-0132013002030333-0030000112231103-1232302102223011-1032103002221033-3311230111010133-3033333010313021"></a>

<a id="canonical-2120300310202112-2303123222302232-1301110300123201-0300310231232113-3300012320022021-2313323213121121-2230003221230101-3201211121312022"></a>

## name property — Property reference / 210323300001 / 9

Type: `"string"`. Required.

Name of the Forward Proxy Policy. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

<a id="canonical-1033103221200233-1302322000130102-2221001113013022-2310313301031332-3223330210313303-3331220031123211-2220213110200023-0012333103031123"></a>

<a id="canonical-1230232020012113-3323332210132132-1313003123010211-0233312321101110-3103202033101012-1020031023033331-0233221130100030-0022013311333231"></a>

## namespace property — Property reference / 210323300001 / 10

Type: `"string"`. Required.

Namespace where the Forward Proxy Policy is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [network_connector](resources--forward_proxy_policy--reference--group-001.md#canonical-1023113102302011-2230123310331013-2321010211321003-0002231332301103-2321203320301010-2301303233002003-1220120223230121-3222022320233033): complete subsection reference.

- [proxy_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-0221230323223211-3010132100100233-3210122132012122-1223331132311301-3301002003203030-3311233232103222-1102003000313200-0033331023010010): complete subsection reference.

- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120): complete subsection reference.

- [timeouts](resources--forward_proxy_policy--reference--group-002.md#canonical-0022013123220103-0231302233112311-3331201312000023-3031022310310321-1300302233001110-2213110210212201-2331202121302302-0031230221002103): complete subsection reference.

<a id="canonical-1013310002211000-3232212323230203-2133003113001012-0020301332333012-0112002302231021-1120030002132220-0010132022333211-3220101023131301"></a>

## All schema paths — Property reference / 210323300001 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](resources--forward_proxy_policy--reference--group-001.md#canonical-0121312202302102-2231233303211011-1103033232320311-2031110203232011-2133030230003000-2310212330031323-3003302100131110-1233313102221131) |
| `allow_list` | [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-1311231332333100-2111013320332112-3313331323213200-3313103011331031-3220202302202001-0321333203011131-1100223300033202-0111200213331220) |
| `allow_list.default_action_allow` | [allow_list.default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-0023231231131223-0231203313202002-1200022302210212-0222200223300222-2211223231311101-1203201302313121-3301220000300212-2313201313132222) |
| `allow_list.default_action_deny` | [allow_list.default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-0322112211201332-2113233020122221-3013003020230121-2201203020330120-2322211010110031-2101301330233300-1033122120320020-3103103130100002) |
| `allow_list.default_action_next_policy` | [allow_list.default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-2022310213110011-2130130123130331-3101000210102001-0132010320011233-2311212201030010-3302020332323222-0011131122103222-1201230310221113) |
| `allow_list.dest_list` | [allow_list.dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3201003000222021-3002221311232131-3233120010221212-3101131030322200-2012132212032301-0000223021222311-2301322223331302-1123031312331231) |
| `allow_list.dest_list.ipv6_prefixes` | [allow_list.dest_list.ipv6_prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-2233303220113123-1321112002333333-3133200110201303-0120120332013103-2123123000112313-2200023233031033-3203320211321012-1113003113213122) |
| `allow_list.dest_list.port_ranges` | [allow_list.dest_list.port_ranges](resources--forward_proxy_policy--reference--group-001.md#canonical-2001023320111112-0223212032123232-2013111102333331-0311223302322012-3033310013113203-2200103320011212-2023003012222212-2002213211230231) |
| `allow_list.dest_list.prefixes` | [allow_list.dest_list.prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-3213200323311123-2221112021223223-3033323203300210-2232202132021100-1311003231003011-2321310122312212-1321312232210022-2000312333133111) |
| `allow_list.http_list` | [allow_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3201133302332311-0311321231023330-1032110220113102-1230011212131033-1213320001023313-0302033020112320-2102020210210110-3103031200021112) |
| `allow_list.http_list.any_path` | [allow_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-2201020110333010-0011303323202001-1232323233220011-1122211133203323-1312332123100312-1133131123313123-3113221203022330-1101320032303110) |
| `allow_list.http_list.exact_value` | [allow_list.http_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-2122210111201331-2120021303203121-1001132232030223-1233003222010113-1322213311331321-3213203232102033-1300210032300101-2030322100201321) |
| `allow_list.http_list.path_exact_value` | [allow_list.http_list.path_exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-1200022202200202-3112323103313132-2021332033102101-3010003132031322-2332013030223330-3232110302212233-1211213203310213-1322113002203313) |
| `allow_list.http_list.path_prefix_value` | [allow_list.http_list.path_prefix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-1300222202310232-3101102232000331-0322332032213112-1133111121012102-2222100030020313-3123103300122203-0202100103301101-1311132130123212) |
| `allow_list.http_list.path_regex_value` | [allow_list.http_list.path_regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-2123133311320233-1333312332001132-0211000122333032-2100103233010111-0021330210231322-3003200100121232-2122122101120233-2330122210133002) |
| `allow_list.http_list.regex_value` | [allow_list.http_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-2302333120112210-2203320230330301-1101210320033101-2020003211020213-1303133230122201-1131123310120333-2311100033002213-3203331323021021) |
| `allow_list.http_list.suffix_value` | [allow_list.http_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-2322001331233223-3112221120011033-1032232121021120-1303122230311231-3330220112320131-1130331332220021-0212332322231120-1322312221101000) |
| `allow_list.tls_list` | [allow_list.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0111312021333023-1200001213031032-3212113030020121-3131211033121010-2200131211110221-1330323231303122-1302211020013310-3210211131022001) |
| `allow_list.tls_list.exact_value` | [allow_list.tls_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-0130120303202202-2223201132001312-3203323111231323-0113000313030013-1133120311221322-2121122202121111-1310330213121210-1231030121320233) |
| `allow_list.tls_list.regex_value` | [allow_list.tls_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-1313222132221203-0200113232111300-0203231130203321-2201011133001333-2122023131102113-1020222011113022-1322311222201301-3311103321333203) |
| `allow_list.tls_list.suffix_value` | [allow_list.tls_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-1330023202021011-3321300312320310-0303003120231030-2330211111310222-0223301130300132-3132300001011203-2003030211101010-1202001010332000) |
| `annotations` | [annotations](resources--forward_proxy_policy--reference--group-001.md#canonical-2023212020311303-2101301302220330-0202033010331321-3211033002011320-3231332200033320-0011231123301133-1002300221301303-3203210321230212) |
| `any_proxy` | [any_proxy](resources--forward_proxy_policy--reference--group-001.md#canonical-3202221230111110-3022111320011013-2322033201010111-3030113322130000-1220333313001323-1313220232010300-0322110103312310-3022103112101303) |
| `deny_list` | [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0323211232130133-3331110131013022-0230220113303203-1032310112003303-0202232331231322-0030100202012110-2032021102211123-2130331222321120) |
| `deny_list.default_action_allow` | [deny_list.default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-1311330330211102-0310021230313132-2033213101033331-1001101213303333-3202321031213031-2303201300033120-0123113001003123-1100133232112210) |
| `deny_list.default_action_deny` | [deny_list.default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-0321332131332322-0231330030323110-1213112230212221-0113012020220023-2333330103302201-2101300212110022-3102300230020003-0301320021032323) |
| `deny_list.default_action_next_policy` | [deny_list.default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-1232011323123231-1322220223002012-1111330302303201-1033311223322010-1320231110311021-2201023132202202-2132102003223200-3023021011001302) |
| `deny_list.dest_list` | [deny_list.dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3321113231102220-0112322021012331-0001202201311330-2330231222111130-2223311023111301-2211121331221323-3312103001032121-0220133213102023) |
| `deny_list.dest_list.ipv6_prefixes` | [deny_list.dest_list.ipv6_prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-3202200110331333-0231113001210310-3012203000122222-3310112000311012-2213103101313010-1301323100213303-2130311022311313-2020201002111130) |
| `deny_list.dest_list.port_ranges` | [deny_list.dest_list.port_ranges](resources--forward_proxy_policy--reference--group-001.md#canonical-1331000201110230-1230023012132230-2123322201301020-1022203013020201-0031332003212232-2300013312032310-3131200302103221-1112100202020001) |
| `deny_list.dest_list.prefixes` | [deny_list.dest_list.prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-0121323130322021-2332100302013200-3113112011003310-1111000132130301-2331231121022120-3330022300211331-3101321031032222-0001102013222001) |
| `deny_list.http_list` | [deny_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020030012010222-0010310332000213-3231301123313003-0102011032332201-3123002221021303-2031011320321311-2321131120222310-2032312230332202) |
| `deny_list.http_list.any_path` | [deny_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-2200020100211301-2031101210211110-3111201232030112-3003120331203111-0000030303210023-2202321313230002-1111031233111033-3120212110212010) |
| `deny_list.http_list.exact_value` | [deny_list.http_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-1331032230310001-2100113300113201-3210220210200212-0203131231103331-3123021100331122-1022221221221021-1100221330232332-3012313232012012) |
| `deny_list.http_list.path_exact_value` | [deny_list.http_list.path_exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-2320022213233012-0322202111111030-1200120121320010-0113330311331122-1031312313221210-2313321103113303-3112000020201012-1301213301213202) |
| `deny_list.http_list.path_prefix_value` | [deny_list.http_list.path_prefix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-3132121311033121-3120301110033132-0102212120000101-3312202311331031-2111202223010120-3020203130101132-2021221301302023-3013112212103301) |
| `deny_list.http_list.path_regex_value` | [deny_list.http_list.path_regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-2021322001010223-1102122220313113-0002011031023032-0023033112032131-2310303100203000-0320002322210010-1122212322031111-3203321303311121) |
| `deny_list.http_list.regex_value` | [deny_list.http_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-0332333103100012-2312313103131021-2113221310211121-0331110130223012-0101131233023310-1220212302132313-1113130203022322-3112013311003012) |
| `deny_list.http_list.suffix_value` | [deny_list.http_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-2310300300033331-0000233013321103-3312103120232011-0311212021010133-3311112221032210-2000321131321132-1332332311011001-2333323101300121) |
| `deny_list.tls_list` | [deny_list.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0210032332022302-3302131202312221-0131210122311212-1000212112031320-3113023321023201-1122222102120103-0331023030320132-0000033130133201) |
| `deny_list.tls_list.exact_value` | [deny_list.tls_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-0310003310121110-0111121223231133-1233011121122300-0223010210010231-2203231310003200-1210112022301022-0013221020113012-3001302202203200) |
| `deny_list.tls_list.regex_value` | [deny_list.tls_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-0303320220031210-0223033312202100-0232213200310313-0030112222332131-3001213000011033-0230101000112202-2130002321030001-0220232230003323) |
| `deny_list.tls_list.suffix_value` | [deny_list.tls_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-3302231102333112-0102300003323303-0132123320113030-0033033200110312-3022022211213212-3012201203203213-1201102131333301-2113031312310112) |
| `description` | [description](resources--forward_proxy_policy--reference--group-001.md#canonical-2101010223032021-2222013022222022-1101001210111111-0333021103300203-0002302133230303-3201031110023312-0121111012210020-1220110332110012) |
| `disable` | [disable](resources--forward_proxy_policy--reference--group-001.md#canonical-1312002302010300-3102223220021302-2132203220103233-3212012233311012-0112020212330330-0301002210111012-1231301303002102-2033010223012311) |
| `drp_http_connect` | [drp_http_connect](resources--forward_proxy_policy--reference--group-001.md#canonical-2213200132201033-0010330133330103-3302333011021032-1303202212331123-2132220013311200-2312220130321232-1031202232211113-1122101231012322) |
| `id` | [ID](resources--forward_proxy_policy--reference--group-001.md#canonical-1003330020213121-3212233011220201-3211301103003310-3202133311312102-1003303132112321-2300210331311100-1022313020312330-1322203323220012) |
| `labels` | [labels](resources--forward_proxy_policy--reference--group-001.md#canonical-0132002111230200-1212130233332322-1012123300210112-0022103023311012-0230121331120220-1203010333302020-2021110013232001-0230330030222111) |
| `name` | [name](resources--forward_proxy_policy--reference--group-001.md#canonical-3003211310103233-0101302332032103-0132013002030333-0030000112231103-1232302102223011-1032103002221033-3311230111010133-3033333010313021) |
| `namespace` | [namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-1033103221200233-1302322000130102-2221001113013022-2310313301031332-3223330210313303-3331220031123211-2220213110200023-0012333103031123) |
| `network_connector` | [network_connector](resources--forward_proxy_policy--reference--group-001.md#canonical-1321130202130211-0232301301212322-2320303122132031-2011310313212200-0211110113010100-3213131112122303-0231300030003123-3302323112321201) |
| `network_connector.name` | [network_connector.name](resources--forward_proxy_policy--reference--group-001.md#canonical-2122320303201123-2322302110023012-3212033312212330-2122302330333113-3031331100111133-1332310201210320-1130132101210030-2001013012212312) |
| `network_connector.namespace` | [network_connector.namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-0312101020102200-3323200132322002-3200012302010311-2032132322021323-3033230123103212-2132112031132012-3320123020303230-2302232333300213) |
| `network_connector.tenant` | [network_connector.tenant](resources--forward_proxy_policy--reference--group-001.md#canonical-1312223200121302-1210131332133231-1000332201030203-3133320310210023-3201331230300012-3231311013020220-0130101333221302-1300130212301303) |
| `proxy_label_selector` | [proxy_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-3111002303022031-3313302320131202-0002322100332212-1022112321301020-3203320320321333-0222031133023001-2312200122230321-1032320001202331) |
| `proxy_label_selector.expressions` | [proxy_label_selector.expressions](resources--forward_proxy_policy--reference--group-001.md#canonical-3201211130202332-2202302330013100-0300022120303313-1132021123100111-0200313000200102-2313001132212200-2112010202220210-1021223303003112) |
| `rule_list` | [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3120101031021000-0311222230303232-3110212032012330-3000223220331220-2233310120133312-0001101213022230-2123233112313210-2100201222030032) |
| `rule_list.rules` | [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-3312120222023102-3011020210320310-3103002210201000-0003220102311011-3111011031221010-0103012112312202-3021233312200333-0333313230200112) |
| `rule_list.rules.action` | [rule_list.rules.action](resources--forward_proxy_policy--reference--group-001.md#canonical-1210301032202321-3321023231033121-3012222222312121-2213000223102330-3221112323212121-2220323212113001-0301111201301003-1021313123100030) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](resources--forward_proxy_policy--reference--group-001.md#canonical-3132030112213100-2310103330022200-2131003300123302-1132133201031302-2103020101332133-1302013203011222-3012230100000133-1323023210222320) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](resources--forward_proxy_policy--reference--group-001.md#canonical-3212102110303220-1013322223223200-1012020002331200-0212001203213020-2323303320021303-3322102310313321-1021033112030103-3312210203030223) |
| `rule_list.rules.dst_asn_list` | [rule_list.rules.dst_asn_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0133131032312222-3031301123203003-2313303030213010-1313332220210321-3002120022010323-0202200023303100-1030300300322002-1200331011132330) |
| `rule_list.rules.dst_asn_list.as_numbers` | [rule_list.rules.dst_asn_list.as_numbers](resources--forward_proxy_policy--reference--group-001.md#canonical-3122100001023130-1031310110202203-2320323301313322-2101212013320131-3303133021033233-0310230113000310-1322033110313122-0030033232102132) |
| `rule_list.rules.dst_asn_set` | [rule_list.rules.dst_asn_set](resources--forward_proxy_policy--reference--group-001.md#canonical-2331112120122301-0133313301032221-0012303212331010-3021321001312212-2011002333330103-1011203112230202-1012223023111230-3122202012111221) |
| `rule_list.rules.dst_asn_set.name` | [rule_list.rules.dst_asn_set.name](resources--forward_proxy_policy--reference--group-001.md#canonical-2322000130221222-1321221031303133-0032001101310020-2130212203122212-3211100001133102-1000122221100303-1112313001031313-3032303020230001) |
| `rule_list.rules.dst_asn_set.namespace` | [rule_list.rules.dst_asn_set.namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-2001131231213200-0003031233302220-0010301231230113-0233023021113333-0222111301013010-2321132202331030-0121201201113332-2332233323310122) |
| `rule_list.rules.dst_asn_set.tenant` | [rule_list.rules.dst_asn_set.tenant](resources--forward_proxy_policy--reference--group-001.md#canonical-3033301130331101-0201132001103210-3012030332110000-0031113331303210-3331302031021301-3023201223030311-0302000312031001-1010031303212302) |
| `rule_list.rules.dst_ip_prefix_set` | [rule_list.rules.dst_ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-0011232133000131-0011100200323101-3030033303023333-2022302203321221-0302332210113131-2213100122312100-0002032010322222-0203323300320333) |
| `rule_list.rules.dst_ip_prefix_set.name` | [rule_list.rules.dst_ip_prefix_set.name](resources--forward_proxy_policy--reference--group-001.md#canonical-2311223211201033-0000130323333133-1320113031333113-2303021333301121-0322023201012113-1321003022001233-2122312210302132-0212312201131203) |
| `rule_list.rules.dst_ip_prefix_set.namespace` | [rule_list.rules.dst_ip_prefix_set.namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-3232300201110110-3313012332102222-1300330111333200-0012222112102123-3120123231220133-0033002310121010-2211122013030021-3233003303032033) |
| `rule_list.rules.dst_ip_prefix_set.tenant` | [rule_list.rules.dst_ip_prefix_set.tenant](resources--forward_proxy_policy--reference--group-001.md#canonical-0113010212002230-0211133322030033-2210022102230201-0330111011313212-0010131203211201-1201222120201001-2023221113111110-3110322102132021) |
| `rule_list.rules.dst_label_selector` | [rule_list.rules.dst_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-3333031302001101-0030110031100012-3200112300222120-1330112310103100-0132330023312120-2222301223332202-1300021322022100-3323210112201323) |
| `rule_list.rules.dst_label_selector.expressions` | [rule_list.rules.dst_label_selector.expressions](resources--forward_proxy_policy--reference--group-001.md#canonical-0323111132202312-0031303221201210-3021012033322010-1310333200133110-2202123021111223-1323330200131233-0112302030000220-1003210113132111) |
| `rule_list.rules.dst_prefix_list` | [rule_list.rules.dst_prefix_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3133300003333203-3230030332132300-2132312332213112-1001310301133020-0313101113230210-1111313302213121-2123212301312212-0321233021211211) |
| `rule_list.rules.dst_prefix_list.prefixes` | [rule_list.rules.dst_prefix_list.prefixes](resources--forward_proxy_policy--reference--group-001.md#canonical-1123201213103230-0212322122321330-2123100330003321-2322212312312213-1003202210102321-0231113301131021-2333113331203202-3123302003313213) |
| `rule_list.rules.http_list` | [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0303221302121213-3300222312030333-0201113302021220-0233320100320221-3132010201031003-0121131020030220-2113330031013200-2121113112200120) |
| `rule_list.rules.http_list.http_list` | [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3233101000301203-2221302110220122-3120023122101002-2103030031301331-2102112332233320-2132233021113322-3300223022101021-1110310001131020) |
| `rule_list.rules.http_list.http_list.any_path` | [rule_list.rules.http_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-3202321113213112-0133121003330220-1110021232002330-0003032332330231-2030303330111010-0032032002123322-3231021102212313-1330020232232221) |
| `rule_list.rules.http_list.http_list.exact_value` | [rule_list.rules.http_list.http_list.exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-0110303312312100-3302332023333011-3013211102110311-1332321111223000-1211030231013022-0130233232003230-0302010002233123-1002000330312222) |
| `rule_list.rules.http_list.http_list.path_exact_value` | [rule_list.rules.http_list.http_list.path_exact_value](resources--forward_proxy_policy--reference--group-001.md#canonical-1030230101321100-1320202313303213-0103113132013002-0303123013001311-0230201220102033-0222131300323020-3001221033201333-2223323001002233) |
| `rule_list.rules.http_list.http_list.path_prefix_value` | [rule_list.rules.http_list.http_list.path_prefix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-1023000131030012-0201003310033101-0033320003332302-3031201313032022-1001201121030000-2312330012311310-1013022130120212-0100002203013231) |
| `rule_list.rules.http_list.http_list.path_regex_value` | [rule_list.rules.http_list.http_list.path_regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-2321103313110001-0220320211301003-1231110312331110-2222213123220123-3013113011312202-2133321323331223-1102220123121220-0331112123220121) |
| `rule_list.rules.http_list.http_list.regex_value` | [rule_list.rules.http_list.http_list.regex_value](resources--forward_proxy_policy--reference--group-001.md#canonical-3103123221200021-3003232300100022-2111010130321012-0023121232312321-1312133230011200-0322300122033311-3020100022021201-3222312221301311) |
| `rule_list.rules.http_list.http_list.suffix_value` | [rule_list.rules.http_list.http_list.suffix_value](resources--forward_proxy_policy--reference--group-001.md#canonical-3312001323010103-2001203101301300-3011033023130120-0111033103113213-2003321023232010-1201123221220112-3103132112103011-0231230112101223) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-3121023022033302-1031312030100132-0330222032021033-1032003012313021-1111132222010301-1201132021123313-3320303010120013-3311010121132023) |
| `rule_list.rules.ip_prefix_set.name` | [rule_list.rules.ip_prefix_set.name](resources--forward_proxy_policy--reference--group-001.md#canonical-2012010232113010-0022122301320321-2332020121012222-2012112123021011-0110010201230221-3100302001030100-3101220203020122-2322031231120023) |
| `rule_list.rules.ip_prefix_set.namespace` | [rule_list.rules.ip_prefix_set.namespace](resources--forward_proxy_policy--reference--group-001.md#canonical-0333111100230131-1132003010020103-2013122011020022-2210320330111033-3313310322023312-0233022323031230-2033211011121300-0303103022002103) |
| `rule_list.rules.ip_prefix_set.tenant` | [rule_list.rules.ip_prefix_set.tenant](resources--forward_proxy_policy--reference--group-001.md#canonical-3200333102332123-3301130310212333-3122102030101313-0300002032000321-2311210311332130-2232320100322102-1012333101323210-1302001331302021) |
| `rule_list.rules.label_selector` | [rule_list.rules.label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-3121111211111121-1122302203201112-3102332311230230-3203303120100331-3021222232232112-3133330102113323-2112203100221130-3322102323110331) |
| `rule_list.rules.label_selector.expressions` | [rule_list.rules.label_selector.expressions](resources--forward_proxy_policy--reference--group-001.md#canonical-2301003312033221-2123222132203212-3121001210320023-2103033321110233-1230021311201000-2002231310010012-3003212231231022-2313022330311321) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](resources--forward_proxy_policy--reference--group-001.md#canonical-1332230121002020-3211213221030200-2100000310310030-2010212011132131-0020012313222302-2212122211112020-2130030012301333-1312320023312112) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](resources--forward_proxy_policy--reference--group-001.md#canonical-0233103121111200-2102202012102232-1232130022221021-3011020121123011-0021131301002021-3100000210113020-2323223031000121-3320112203111311) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](resources--forward_proxy_policy--reference--group-001.md#canonical-1210211300000123-1010312231021100-1002312113332021-0211031021311120-0310310012311330-1220100012201131-1200010330123212-0223033020112330) |
| `rule_list.rules.no_http_connect_port` | [rule_list.rules.no_http_connect_port](resources--forward_proxy_policy--reference--group-001.md#canonical-0111023013123131-3100003330320311-1212021131321300-1233200332221332-3212323113123333-1123130023200200-0130023123113223-2103000121330203) |
| `rule_list.rules.port_matcher` | [rule_list.rules.port_matcher](resources--forward_proxy_policy--reference--group-002.md#canonical-2223210130031330-3322111110333000-0232120132010310-2310330231232132-3323020012121232-2033322201200330-0302302232022332-1031020232131121) |
| `rule_list.rules.port_matcher.invert_matcher` | [rule_list.rules.port_matcher.invert_matcher](resources--forward_proxy_policy--reference--group-002.md#canonical-2031230022033200-1220133000313003-3310210211031010-2312003221303330-0313231111130210-2032003311133131-0102223231332201-0322012233103232) |
| `rule_list.rules.port_matcher.ports` | [rule_list.rules.port_matcher.ports](resources--forward_proxy_policy--reference--group-002.md#canonical-3301203210020311-3231130311123200-1321231101330031-1101233123123203-1002322130031303-1010101013211122-0102111221101023-1322013330103312) |
| `rule_list.rules.prefix_list` | [rule_list.rules.prefix_list](resources--forward_proxy_policy--reference--group-002.md#canonical-1211122323021202-0133301231103221-0220121222330021-2011313231220002-1302231303213021-2132320101120110-1120033302011123-0303330333031021) |
| `rule_list.rules.prefix_list.prefixes` | [rule_list.rules.prefix_list.prefixes](resources--forward_proxy_policy--reference--group-002.md#canonical-3202103130210002-0303202322222201-3131323233020000-1032321012120302-3303222032300010-1310112310113321-1223132120323121-0021212203132213) |
| `rule_list.rules.tls_list` | [rule_list.rules.tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-1210221123213132-0203202220231001-0221211321313010-3132220223103201-3330133110311231-3313102001230210-0012003131101132-3201023102012202) |
| `rule_list.rules.tls_list.tls_list` | [rule_list.rules.tls_list.tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0320220031133313-0130201231002103-1231201120112001-1331023212011310-2033322223202032-0301311201130303-1232213000333200-0320323303110202) |
| `rule_list.rules.tls_list.tls_list.exact_value` | [rule_list.rules.tls_list.tls_list.exact_value](resources--forward_proxy_policy--reference--group-002.md#canonical-2201020223300212-1232232212110133-3203021000333022-1301110310010322-1210022220123121-2222121002120100-3020223032103200-1303231320212113) |
| `rule_list.rules.tls_list.tls_list.regex_value` | [rule_list.rules.tls_list.tls_list.regex_value](resources--forward_proxy_policy--reference--group-002.md#canonical-2333312232301233-2123321110303321-2300230212320103-1020203302000201-1000200322201031-1221131001330232-2221200101032223-1222132303101233) |
| `rule_list.rules.tls_list.tls_list.suffix_value` | [rule_list.rules.tls_list.tls_list.suffix_value](resources--forward_proxy_policy--reference--group-002.md#canonical-2131322113012020-3133313202132232-2122222001332100-3103212103122010-1103322003212310-2212321111333220-2311013330300120-2002222202001222) |
| `rule_list.rules.url_category_list` | [rule_list.rules.url_category_list](resources--forward_proxy_policy--reference--group-002.md#canonical-1101212112133130-0012123030323210-2001102332220101-0320131122131103-2002132301022320-1233000110003110-0210213223123000-0312130202013230) |
| `rule_list.rules.url_category_list.url_categories` | [rule_list.rules.url_category_list.url_categories](resources--forward_proxy_policy--reference--group-002.md#canonical-2210032102001010-0111200021312310-2023030213021110-0033210311202033-2201330003033213-1100110230031303-3101132231122103-3330303323110030) |
| `timeouts` | [timeouts](resources--forward_proxy_policy--reference--group-002.md#canonical-0312333203320031-2001323121301301-2113320111022333-1231122110212013-2112302103320203-2103010000220022-2200022121212202-0311001302331110) |
| `timeouts.create` | [timeouts.create](resources--forward_proxy_policy--reference--group-002.md#canonical-2213123131121111-2003023001310022-3233110121302133-3332021310311333-1223232223222031-2220332203230002-1030331232011310-2210102102320223) |
| `timeouts.delete` | [timeouts.delete](resources--forward_proxy_policy--reference--group-002.md#canonical-0300322100323133-0201121113223010-2110130301231110-1203030033211001-1311113313112223-2330232130231302-3310211203102123-3031013202333322) |
| `timeouts.read` | [timeouts.read](resources--forward_proxy_policy--reference--group-002.md#canonical-1303112313021102-3120111033121101-3203312232201202-3312120101300222-0022111202312021-2102001213232013-3022001023320321-3213031122111211) |
| `timeouts.update` | [timeouts.update](resources--forward_proxy_policy--reference--group-002.md#canonical-1133330220013103-1010331003022300-2013122001033321-3033022332310323-0010301233300202-0112133323331020-1231212230120213-3301131233111100) |

<a id="canonical-1220331211123112-1121232023033333-2323021222032211-3103030300202213-0013113110103120-0001131033023133-0232011211222032-0221233011110212"></a>

## Next pages — Property reference / 210323300001 / 12

- [allow_all](resources--forward_proxy_policy--reference--group-001.md#canonical-2220212320111303-0320122201200113-3102201023111301-2113200301102333-3233010100233033-3322330201020030-2012231030013231-1231233310333033)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- [any_proxy](resources--forward_proxy_policy--reference--group-001.md#canonical-2332333300301220-2020011003301331-0311130310031221-0302003113033233-1113111033022331-1120333131201021-0101002021020003-0020122103301200)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- [drp_http_connect](resources--forward_proxy_policy--reference--group-001.md#canonical-0121000321310323-0220232133302100-1210330110221132-1112130320100200-1022002100133130-3210200311003310-3011210001322311-1121030003102231)
- [network_connector](resources--forward_proxy_policy--reference--group-001.md#canonical-1023113102302011-2230123310331013-2321010211321003-0002231332301103-2321203320301010-2301303233002003-1220120223230121-3222022320233033)
- [proxy_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-0221230323223211-3010132100100233-3210122132012122-1223331132311301-3301002003203030-3311233232103222-1102003000313200-0033331023010010)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [timeouts](resources--forward_proxy_policy--reference--group-002.md#canonical-0022013123220103-0231302233112311-3331201312000023-3031022310310321-1300302233001110-2213110210212201-2331202121302302-0031230221002103)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2220212320111303-0320122201200113-3102201023111301-2113200301102333-3233010100233033-3322330201020030-2012231030013231-1231233310333033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200210301323223-2023303322100112-2212223320033323-1030232313302022-2102023201012311-1013212112213133-1312111121321301-1221033012012220"></a>

## allow_all — allow_all / 021221102100 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- allow_all

<a id="canonical-0121312202302102-2231233303211011-1103033232320311-2031110203232011-2133030230003000-2310212330031323-3003302100131110-1233313102221131"></a>

Type: `["object", {}]`. Optional.

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

- [allow_all](resources--forward_proxy_policy--reference--group-001.md#canonical-0121312202302102-2231233303211011-1103033232320311-2031110203232011-2133030230003000-2310212330031323-3003302100131110-1233313102221131)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-1311231332333100-2111013320332112-3313331323213200-3313103011331031-3220202302202001-0321333203011131-1100223300033202-0111200213331220)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0323211232130133-3331110131013022-0230220113303203-1032310112003303-0202232331231322-0030100202012110-2032021102211123-2130331222321120)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3120101031021000-0311222230303232-3110212032012330-3000223220331220-2233310120133312-0001101213022230-2123233112313210-2100201222030032)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all = {}
```

<a id="canonical-1013102100330133-0333333103232110-0210300101302202-0003012132321211-3331001020001130-3113030221000221-2101302303101130-0300000121201020"></a>

## Direct properties — allow_all / 021221102100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103231020012003-0123200323210332-1033022322133122-1013220300001102-2203030012021212-0312203301103002-3130330200302302-0031121320333220"></a>

## Next pages — allow_all / 021221102100 / 4

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101001101320230-1123020112311223-3313220123332133-0110033201320222-0113313300032032-1112133010222333-1231122210020023-1032211001201012"></a>

## allow_list — allow_list / 303322010311 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- allow_list

<a id="canonical-1311231332333100-2111013320332112-3313331323213200-3313103011331031-3220202302202001-0321333203011131-1100223300033202-0111200213331220"></a>

Type: `"object"`. single nested block, Optional.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_deny"),
  validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_next_policy"),
  validators.ConflictingObjectAttributes("default_action_deny",
    "default_action_next_policy")}
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
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

Terraform syntax:

```terraform
allow_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223030333012103-2120220320123310-3003103013231301-3110010230211302-1001103231101100-0013031003010330-3230101212312123-1033232203120100"></a>

## Direct properties — allow_list / 303322010311 / 3

- [default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-0231103023111223-1121311013222311-1321003331232120-0002332130022101-2221313233103033-0030112312212322-2022230010223220-3121130333220220): complete subsection reference.

- [default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-0303202333200221-3201101023003011-3023003012312033-0230133031130332-3220122121121212-3103321302133102-3320211313223120-0302113002033131): complete subsection reference.

- [default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-0102023110230223-1321130302013303-3233122322301133-1312313012311333-3203202301130030-1223120211132022-0333122030012002-2220033132221020): complete subsection reference.

- [dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2123133001021123-3312330013331222-1001231112320212-1022312231230113-1333210101311221-0200201221100320-0132310013032131-0003031010320120): complete subsection reference.

- [http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2310303221301111-0331133010332222-0321302332130310-1303130312022203-0013123232111121-3331330020111333-1320031200223001-0201212320121031): complete subsection reference.

- [tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2210220233212330-2121101130022310-2101231231222001-0223201021123312-3112220102330021-2120202030302111-2222203232332103-0133113122113112): complete subsection reference.

<a id="canonical-1320122002122303-3020310333230330-3203330332101102-3121210121332030-1101022233003131-1210012333131100-2330013303132003-1010323211111310"></a>

## Next pages — allow_list / 303322010311 / 4

- [allow_list.default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-0231103023111223-1121311013222311-1321003331232120-0002332130022101-2221313233103033-0030112312212322-2022230010223220-3121130333220220)
- [allow_list.default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-0303202333200221-3201101023003011-3023003012312033-0230133031130332-3220122121121212-3103321302133102-3320211313223120-0302113002033131)
- [allow_list.default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-0102023110230223-1321130302013303-3233122322301133-1312313012311333-3203202301130030-1223120211132022-0333122030012002-2220033132221020)
- [allow_list.dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2123133001021123-3312330013331222-1001231112320212-1022312231230113-1333210101311221-0200201221100320-0132310013032131-0003031010320120)
- [allow_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2310303221301111-0331133010332222-0321302332130310-1303130312022203-0013123232111121-3331330020111333-1320031200223001-0201212320121031)
- [allow_list.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2210220233212330-2121101130022310-2101231231222001-0223201021123312-3112220102330021-2120202030302111-2222203232332103-0133113122113112)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0231103023111223-1121311013222311-1321003331232120-0002332130022101-2221313233103033-0030112312212322-2022230010223220-3121130333220220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303012010022323-0130123220233130-2021001012101320-3333210100213103-2103323300100030-1003003110301203-0100312013321223-0301121001001332"></a>

## allow_list.default_action_allow — default_action_allow / 000111113202 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- allow_list.default_action_allow

<a id="canonical-0023231231131223-0231203313202002-1200022302210212-0222200223300222-2211223231311101-1203201302313121-3301220000300212-2313201313132222"></a>

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
default_action_allow = {}
```

<a id="canonical-1012203003322301-0323031001320302-2121030323110223-2131131002121120-2102012320101123-3030013023003123-3022101003303003-3221301211331313"></a>

## Direct properties — default_action_allow / 000111113202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131103030111110-0110301033003022-2320120303301121-2210310021000220-2320310232120031-2103013321010120-3013320112112120-0333012013202030"></a>

## Next pages — default_action_allow / 000111113202 / 4

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0303202333200221-3201101023003011-3023003012312033-0230133031130332-3220122121121212-3103321302133102-3320211313223120-0302113002033131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333232022132010-2112331000231100-3123233001110310-3301303122330022-2013122120220032-1110310311130123-0013321122013332-2333222320103101"></a>

## allow_list.default_action_deny — default_action_deny / 023230002131 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- allow_list.default_action_deny

<a id="canonical-0322112211201332-2113233020122221-3013003020230121-2201203020330120-2322211010110031-2101301330233300-1033122120320020-3103103130100002"></a>

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
default_action_deny = {}
```

<a id="canonical-1321210231201121-2031103212202222-3232122133301021-0303122330131203-0313223330102202-0000033302122000-2000110312200202-3011203000001303"></a>

## Direct properties — default_action_deny / 023230002131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221322321012210-2002301001311212-0030320101033001-2031113213031022-1032203203230012-3122323121330230-2100113211220332-3301100111210120"></a>

## Next pages — default_action_deny / 023230002131 / 4

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0102023110230223-1321130302013303-3233122322301133-1312313012311333-3203202301130030-1223120211132022-0333122030012002-2220033132221020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032303332331102-0030323112201321-0112102033311321-3312201332132131-2222100120002331-0122221221001232-2102200031233230-1013133032330332"></a>

## allow_list.default_action_next_policy — default_action_next_policy / 303011000231 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- allow_list.default_action_next_policy

<a id="canonical-2022310213110011-2130130123130331-3101000210102001-0132010320011233-2311212201030010-3302020332323222-0011131122103222-1201230310221113"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_action_next_policy = {}
```

<a id="canonical-3012130003003221-2301302330112200-0130320302021001-2002332202132102-3131122310112303-0311111002230012-0002110033010332-0220000231003212"></a>

## Direct properties — default_action_next_policy / 303011000231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112311130222031-2322001223330232-2222233022113312-3112213213013020-0203010233002221-2301133322320212-1332031023131012-2212230330102101"></a>

## Next pages — default_action_next_policy / 303011000231 / 4

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2123133001021123-3312330013331222-1001231112320212-1022312231230113-1333210101311221-0200201221100320-0132310013032131-0003031010320120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220320331101121-2323003003311332-2211112100210123-1021202313011233-1331002320300011-3110013132301330-0302303121031123-2331323133332321"></a>

## allow_list.dest_list — dest_list / 200113122023 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- allow_list.dest_list

<a id="canonical-3201003000222021-3002221311232131-3233120010221212-3101131030322200-2012132212032301-0000223021222311-2301322223331302-1123031312331231"></a>

Type: `"object"`. list nested block, Optional.

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("port_ranges")}
```

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
dest_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322102212111220-1322030001000012-2033000331011001-1232120133231013-0003122100202330-3233311301203000-3322331121331332-3001121210232212"></a>

## Direct properties — dest_list / 200113122023 / 3

<a id="canonical-2233303220113123-1321112002333333-3133200110201303-0120120332013103-2123123000112313-2200023233031033-3203320211321012-1113003113213122"></a>

<a id="canonical-3232123311311132-3232301132112113-1011311302021130-1001103101002322-3000110113201300-3133311202332320-2202301032010111-2223011201322321"></a>

## ipv6_prefixes property — dest_list / 200113122023 / 4

Type: `["list", "string"]`. Optional.

IPv6 Prefixes. Destination IPv6 prefixes.

Upstream description:

Destination IPv6 prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-2001023320111112-0223212032123232-2013111102333331-0311223302322012-3033310013113203-2200103320011212-2023003012222212-2002213211230231"></a>

<a id="canonical-2310021011001323-0013220213223003-0023132333323133-0133323211133012-3000101012231312-1033331101332132-0303113012030200-1221132333011200"></a>

## port_ranges property — dest_list / 200113122023 / 5

Type: `"string"`. Optional.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Upstream description:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-3213200323311123-2221112021223223-3033323203300210-2232202132021100-1311003231003011-2321310122312212-1321312232210022-2000312333133111"></a>

<a id="canonical-1000213030231022-2233011212200101-1100122031303230-3031210110330030-0133331033202101-2302301233111012-1200132310111212-0000332303311033"></a>

## prefixes property — dest_list / 200113122023 / 6

Type: `["list", "string"]`. Optional.

IPv4 Prefixes. Destination IPv4 prefixes.

Upstream description:

Destination IPv4 prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-3000131133203212-1313000001330213-3301213023211221-2130121201222210-0130322200301210-3101202201221201-1032032021133230-3111110320313212"></a>

## Next pages — dest_list / 200113122023 / 7

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2310303221301111-0331133010332222-0321302332130310-1303130312022203-0013123232111121-3331330020111333-1320031200223001-0201212320121031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222013323122132-3103020111313030-0031213212321131-0023211130033113-1202000002333311-3333010023032332-1303033211320200-2112112030213322"></a>

## allow_list.http_list — http_list / 112123102312 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- allow_list.http_list

<a id="canonical-3201133302332311-0311321231023330-1032110220113102-1230011212131033-1213320001023313-0302033020112320-2102020210210110-3103031200021112"></a>

Type: `"object"`. list nested block, Optional.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_path",
    "path_exact_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("path_prefix_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
```

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
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313212023131312-1312223102331233-2002330002130301-1200030001022300-3010030202233102-1311012312303113-0301212011233320-3302032132213312"></a>

## Direct properties — http_list / 112123102312 / 3

- [any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-1033021103202010-1132021021121131-3020213032013102-3111110120213112-1210030022112102-0101103122110321-2002321032131223-0131333233013312): complete subsection reference.

<a id="canonical-2122210111201331-2120021303203121-1001132232030223-1233003222010113-1322213311331321-3213203232102033-1300210032300101-2030322100201321"></a>

<a id="canonical-1010031221223131-2231132123000133-1212002222112100-1023021323332311-0103020123000022-2312000102312022-0103021333211133-0111102101121012"></a>

## exact_value property — http_list / 112123102312 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1200022202200202-3112323103313132-2021332033102101-3010003132031322-2332013030223330-3232110302212233-1211213203310213-1322113002203313"></a>

<a id="canonical-1333333302312032-1002010113332203-2221133133000033-1321231101320201-1120013111202203-1111313031331210-0012231101301321-0102232110233233"></a>

## path_exact_value property — http_list / 112123102312 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1300222202310232-3101102232000331-0322332032213112-1133111121012102-2222100030020313-3123103300122203-0202100103301101-1311132130123212"></a>

<a id="canonical-2120200010033100-3212210211011002-3233322131033332-0111033002322033-1103120300310011-3230321202113122-0120121101001221-0102311033000011"></a>

## path_prefix_value property — http_list / 112123102312 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2123133311320233-1333312332001132-0211000122333032-2100103233010111-0021330210231322-3003200100121232-2122122101120233-2330122210133002"></a>

<a id="canonical-2211033213003323-3310300230131212-1233011310232330-3222002301310000-3202012010033033-0110332231021311-2132211322002303-0100313112120121"></a>

## path_regex_value property — http_list / 112123102312 / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2302333120112210-2203320230330301-1101210320033101-2020003211020213-1303133230122201-1131123310120333-2311100033002213-3203331323021021"></a>

<a id="canonical-3333020222312001-3323032132220310-0112100001230200-2103332223231233-3312321020333212-3231321232130031-0100000203303320-1213331030313232"></a>

## regex_value property — http_list / 112123102312 / 8

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2322001331233223-3112221120011033-1032232121021120-1303122230311231-3330220112320131-1130331332220021-0212332322231120-1322312221101000"></a>

<a id="canonical-2023023300201130-2123300002113222-0012323332023013-0311003211001231-0210112021103131-3100003100301110-0002000230121310-2232112111321031"></a>

## suffix_value property — http_list / 112123102312 / 9

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0233133231002301-0000021311200033-1112021201213332-3030203013303332-1311200113323303-1212111321101301-0100032212300111-0211022121102312"></a>

## Next pages — http_list / 112123102312 / 10

- [allow_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-1033021103202010-1132021021121131-3020213032013102-3111110120213112-1210030022112102-0101103122110321-2002321032131223-0131333233013312)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-1033021103202010-1132021021121131-3020213032013102-3111110120213112-1210030022112102-0101103122110321-2002321032131223-0131333233013312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213022023330103-1313332100032101-1131121213301222-3213123003320011-3113002112013221-1233221320010200-2333203301213022-2233000130232223"></a>

## allow_list.http_list.any_path — any_path / 033202001310 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- [allow_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2310303221301111-0331133010332222-0321302332130310-1303130312022203-0013123232111121-3331330020111333-1320031200223001-0201212320121031)
- allow_list.http_list.any_path

<a id="canonical-2201020110333010-0011303323202001-1232323233220011-1122211133203323-1312332123100312-1133131123313123-3113221203022330-1101320032303110"></a>

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
any_path = {}
```

<a id="canonical-0103303020011121-1112230310312333-0100120100013331-3031011121012322-2012133210112013-0113021222201200-2003113102113332-1100031102130322"></a>

## Direct properties — any_path / 033202001310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100120003133123-3233023100201123-0212220131330020-1223022133321100-1323210311101323-3032030212210131-3222121231103323-2021323020303121"></a>

## Next pages — any_path / 033202001310 / 4

- [allow_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2310303221301111-0331133010332222-0321302332130310-1303130312022203-0013123232111121-3331330020111333-1320031200223001-0201212320121031)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2210220233212330-2121101130022310-2101231231222001-0223201021123312-3112220102330021-2120202030302111-2222203232332103-0133113122113112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321321133111112-1302013302200220-1013320233032213-3332033332011202-3330002222211011-2301322002203133-2100111132233220-2310102303220301"></a>

## allow_list.tls_list — tls_list / 131130320311 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- allow_list.tls_list

<a id="canonical-0111312021333023-1200001213031032-3212113030020121-3131211033121010-2200131211110221-1330323231303122-1302211020013310-3210211131022001"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
```

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
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122111231121210-0321332130130220-1223002311331223-3321012133203010-0212203302212021-2130203302122222-3313132111030022-0030120133231023"></a>

## Direct properties — tls_list / 131130320311 / 3

<a id="canonical-0130120303202202-2223201132001312-3203323111231323-0113000313030013-1133120311221322-2121122202121111-1310330213121210-1231030121320233"></a>

<a id="canonical-0001120320332212-0130110000112212-0033003103131022-2012333223123001-2313211132331030-2203100232010202-1122322021231122-0110100101233020"></a>

## exact_value property — tls_list / 131130320311 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1313222132221203-0200113232111300-0203231130203321-2201011133001333-2122023131102113-1020222011113022-1322311222201301-3311103321333203"></a>

<a id="canonical-2333110220102113-0002330232202110-0221113111221110-2221113200322313-1320201212311311-3223022322221022-2221231201110121-1031011021120232"></a>

## regex_value property — tls_list / 131130320311 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1330023202021011-3321300312320310-0303003120231030-2330211111310222-0223301130300132-3132300001011203-2003030211101010-1202001010332000"></a>

<a id="canonical-0202321111312010-3232032332300111-1201322201221331-3223032221021112-1011031100303123-1222211333001122-1102003020233003-2233113300330112"></a>

## suffix_value property — tls_list / 131130320311 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0101000022030131-0030132110002302-1130003213220003-0311333231313322-3210002111010303-0112132222030130-2300331301201320-3320011312210130"></a>

## Next pages — tls_list / 131130320311 / 7

- [allow_list](resources--forward_proxy_policy--reference--group-001.md#canonical-0322102321320320-0311022011323123-3103001210112003-2023012333100230-2232032323023213-3032031323031130-0233001010221331-3011031120333302)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2332333300301220-2020011003301331-0311130310031221-0302003113033233-1113111033022331-1120333131201021-0101002021020003-0020122103301200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231211221303322-3003020022130201-2211213311022020-2131131230223232-0003322123232203-0312113322113100-3231132001030213-3221310002023002"></a>

## any_proxy — any_proxy / 312133311111 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- any_proxy

<a id="canonical-3202221230111110-3022111320011013-2322033201010111-3030113322130000-1220333313001323-1313220232010300-0322110103312310-3022103112101303"></a>

Type: `["object", {}]`. Optional.

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

- [any_proxy](resources--forward_proxy_policy--reference--group-001.md#canonical-3202221230111110-3022111320011013-2322033201010111-3030113322130000-1220333313001323-1313220232010300-0322110103312310-3022103112101303)
- [drp_http_connect](resources--forward_proxy_policy--reference--group-001.md#canonical-2213200132201033-0010330133330103-3302333011021032-1303202212331123-2132220013311200-2312220130321232-1031202232211113-1122101231012322)
- [network_connector](resources--forward_proxy_policy--reference--group-001.md#canonical-1321130202130211-0232301301212322-2320303122132031-2011310313212200-0211110113010100-3213131112122303-0231300030003123-3302323112321201)
- [proxy_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-3111002303022031-3313302320131202-0002322100332212-1022112321301020-3203320320321333-0222031133023001-2312200122230321-1032320001202331)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_proxy = {}
```

<a id="canonical-1131101330101320-1220321100211203-0002112102220211-0313011231233303-2002123312301312-3003230001333312-1122232311211030-3210022110130111"></a>

## Direct properties — any_proxy / 312133311111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313312132333233-1223312121033231-1020101321110133-3310323133010212-3311220103320231-0230130200321331-2032200103033112-2233222213121331"></a>

## Next pages — any_proxy / 312133311111 / 4

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131200102211202-2213110231022233-2022032201312221-2213202100023300-1023021111031010-3212221120120202-3323303221130020-0031123231322302"></a>

## deny_list — deny_list / 213213113212 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- deny_list

<a id="canonical-0323211232130133-3331110131013022-0230220113303203-1032310112003303-0202232331231322-0030100202012110-2032021102211123-2130331222321120"></a>

Type: `"object"`. single nested block, Optional.

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP).

Upstream description:

URL(s) and domains policy for forward proxy for a connection type (TLS or HTTP)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_deny"),
  validators.ConflictingObjectAttributes("default_action_allow",
    "default_action_next_policy"),
  validators.ConflictingObjectAttributes("default_action_deny",
    "default_action_next_policy")}
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
  "x-ves-oneof-field-default_action_choice": "[\"default_action_allow\",\"default_action_deny\",\"default_action_next_policy\"]"
}
```

Terraform syntax:

```terraform
deny_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120200023211030-2010031201332333-2333311220131010-1312211322031021-3010031211022222-3222202110102220-2230002210012123-0310330021133203"></a>

## Direct properties — deny_list / 213213113212 / 3

- [default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-0122202113223013-1332131032301111-1333123012113303-2231320331122011-1300310122322033-3323121021113021-1330012223033122-3322301321300122): complete subsection reference.

- [default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-2231023203300130-0112131021122023-2133103322203123-3003231132300030-3021322323213202-1003310311112002-1031310130003013-0112001033322222): complete subsection reference.

- [default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-1210012221212003-3013110130210311-3333030320301320-1032121100213131-1032211322233131-1323001120212031-3003003131301321-1030020323131023): complete subsection reference.

- [dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2021023221213120-2220230112212201-0313100222120123-1100030123122122-1000300112300300-1031311201022221-0220222331203131-2221231130321032): complete subsection reference.

- [http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3202121002221232-1001233131102222-3332221121223212-0111023100312303-2220221211222132-2312230221111030-0022112301033230-3001113000001013): complete subsection reference.

- [tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2030011123202013-2020022321113210-3000123000322103-0332101233032232-3031200322302233-3320333113121232-0103133001223023-3020331222230303): complete subsection reference.

<a id="canonical-0101022002222033-0131102030312330-3002211121002001-1123200223003220-2203013001211231-0310112102021012-3133323313232021-1022131123112020"></a>

## Next pages — deny_list / 213213113212 / 4

- [deny_list.default_action_allow](resources--forward_proxy_policy--reference--group-001.md#canonical-0122202113223013-1332131032301111-1333123012113303-2231320331122011-1300310122322033-3323121021113021-1330012223033122-3322301321300122)
- [deny_list.default_action_deny](resources--forward_proxy_policy--reference--group-001.md#canonical-2231023203300130-0112131021122023-2133103322203123-3003231132300030-3021322323213202-1003310311112002-1031310130003013-0112001033322222)
- [deny_list.default_action_next_policy](resources--forward_proxy_policy--reference--group-001.md#canonical-1210012221212003-3013110130210311-3333030320301320-1032121100213131-1032211322233131-1323001120212031-3003003131301321-1030020323131023)
- [deny_list.dest_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2021023221213120-2220230112212201-0313100222120123-1100030123122122-1000300112300300-1031311201022221-0220222331203131-2221231130321032)
- [deny_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3202121002221232-1001233131102222-3332221121223212-0111023100312303-2220221211222132-2312230221111030-0022112301033230-3001113000001013)
- [deny_list.tls_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2030011123202013-2020022321113210-3000123000322103-0332101233032232-3031200322302233-3320333113121232-0103133001223023-3020331222230303)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0122202113223013-1332131032301111-1333123012113303-2231320331122011-1300310122322033-3323121021113021-1330012223033122-3322301321300122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103203333232022-0113321023122010-1110330310001000-1021112320223032-0302013301010110-2020012032001201-0233230202302121-0201222220300203"></a>

## deny_list.default_action_allow — default_action_allow / 022310031022 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- deny_list.default_action_allow

<a id="canonical-1311330330211102-0310021230313132-2033213101033331-1001101213303333-3202321031213031-2303201300033120-0123113001003123-1100133232112210"></a>

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
default_action_allow = {}
```

<a id="canonical-0022113210133231-3001121020132220-0313220020010301-1021101301123032-0313013210122111-3030012303222303-1300322212130020-0022230010111213"></a>

## Direct properties — default_action_allow / 022310031022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311013031220002-2120112001111032-3301221100012000-1120123120230222-1210330000010033-0131330002313211-3003313322223012-3120212123202002"></a>

## Next pages — default_action_allow / 022310031022 / 4

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2231023203300130-0112131021122023-2133103322203123-3003231132300030-3021322323213202-1003310311112002-1031310130003013-0112001033322222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122233320032331-0111200213003210-0021211311031033-2222320200302111-0021023003133300-0231003102332310-3301333131101022-0213011102030201"></a>

## deny_list.default_action_deny — default_action_deny / 202332131310 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- deny_list.default_action_deny

<a id="canonical-0321332131332322-0231330030323110-1213112230212221-0113012020220023-2333330103302201-2101300212110022-3102300230020003-0301320021032323"></a>

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
default_action_deny = {}
```

<a id="canonical-0312230002231222-0000123002320330-3102221201023311-3321133110030000-1312022202311321-0000020013302132-1200103323300023-3212120003202033"></a>

## Direct properties — default_action_deny / 202332131310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331300013320321-2212330120011132-2020220330110021-3013031201212031-0122102330032222-2321110113000030-1012332030311021-3313001100012031"></a>

## Next pages — default_action_deny / 202332131310 / 4

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-1210012221212003-3013110130210311-3333030320301320-1032121100213131-1032211322233131-1323001120212031-3003003131301321-1030020323131023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322200123233233-2022101310300302-3111131012203212-3131223311022112-3023100102200002-1311133000021012-2130013123210120-2120012202032101"></a>

## deny_list.default_action_next_policy — default_action_next_policy / 232323310031 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- deny_list.default_action_next_policy

<a id="canonical-1232011323123231-1322220223002012-1111330302303201-1033311223322010-1320231110311021-2201023132202202-2132102003223200-3023021011001302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_action_next_policy = {}
```

<a id="canonical-2230103210103121-1303101232320323-0131212123132020-1222011100112231-2030230032300001-3202120323201300-0320322313312233-1130220203232213"></a>

## Direct properties — default_action_next_policy / 232323310031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031203332011212-0102233220033131-2303221323210331-3010011033332220-2002120011020321-3333301132201112-2300212312312302-0001011002012210"></a>

## Next pages — default_action_next_policy / 232323310031 / 4

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2021023221213120-2220230112212201-0313100222120123-1100030123122122-1000300112300300-1031311201022221-0220222331203131-2221231130321032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221132013032213-0100300133012120-3300220231031001-2223032312213113-1212233000020123-0030103222233201-1111111201203002-1012033320303131"></a>

## deny_list.dest_list — dest_list / 311002120000 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- deny_list.dest_list

<a id="canonical-3321113231102220-0112322021012331-0001202201311330-2330231222111130-2223311023111301-2211121331221323-3312103001032121-0220133213102023"></a>

Type: `"object"`. list nested block, Optional.

L4 destinations for non-HTTP and non-TLS connections and TLS connections without SNI.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("port_ranges")}
```

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
dest_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130122130302112-2213112102210201-2032013233101201-0021333022030203-2101031223230210-1130110103303031-1233020100210122-1323111032212003"></a>

## Direct properties — dest_list / 311002120000 / 3

<a id="canonical-3202200110331333-0231113001210310-3012203000122222-3310112000311012-2213103101313010-1301323100213303-2130311022311313-2020201002111130"></a>

<a id="canonical-0200020003311120-0021321120300133-0030021123002000-2033322031331030-0123100322013130-0023212231221320-3221010112321212-3112133110213221"></a>

## ipv6_prefixes property — dest_list / 311002120000 / 4

Type: `["list", "string"]`. Optional.

IPv6 Prefixes. Destination IPv6 prefixes.

Upstream description:

Destination IPv6 prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-1331000201110230-1230023012132230-2123322201301020-1022203013020201-0031332003212232-2300013312032310-3131200302103221-1112100202020001"></a>

<a id="canonical-3030311130131302-3211031031001230-0030213221313221-3112313313121130-3333303320022003-1133302123111001-1130011111023012-3011332200130031"></a>

## port_ranges property — dest_list / 311002120000 / 5

Type: `"string"`. Optional.

String containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by '-'.

Upstream description:

A string containing a comma separated list of port ranges. Each port range consists of a single port
or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

<a id="canonical-0121323130322021-2332100302013200-3113112011003310-1111000132130301-2331231121022120-3330022300211331-3101321031032222-0001102013222001"></a>

<a id="canonical-1013223303003011-3303210012003330-3020300323122231-3300100133000233-0132312320021103-2231102110102222-0022030212303121-2130201310301113"></a>

## prefixes property — dest_list / 311002120000 / 6

Type: `["list", "string"]`. Optional.

IPv4 Prefixes. Destination IPv4 prefixes.

Upstream description:

Destination IPv4 prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-0100120333011233-3302131013201232-0121301122031203-2002331222003032-2300331320202131-2031023303210330-0120100331233030-3122103100012123"></a>

## Next pages — dest_list / 311002120000 / 7

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-3202121002221232-1001233131102222-3332221121223212-0111023100312303-2220221211222132-2312230221111030-0022112301033230-3001113000001013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011022231001221-1131233122320203-3313201323110202-0330110333031002-1133332122231003-2312101022030122-3321233123013210-3233021201200313"></a>

## deny_list.http_list — http_list / 233112201231 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- deny_list.http_list

<a id="canonical-3020030012010222-0010310332000213-3231301123313003-0102011032332201-3123002221021303-2031011320321311-2321131120222310-2032312230332202"></a>

Type: `"object"`. list nested block, Optional.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_path",
    "path_exact_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("path_prefix_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
```

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
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111213320202010-2332233022133203-2231233311031223-2322333021023310-0302121102102113-3323230332233112-1201313031200301-0013212121030130"></a>

## Direct properties — http_list / 233112201231 / 3

- [any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-0221121112200211-3020121331300210-0322123232233121-0300231231222130-0200311103230333-0221311123211303-1233113230022321-2100011111111202): complete subsection reference.

<a id="canonical-1331032230310001-2100113300113201-3210220210200212-0203131231103331-3123021100331122-1022221221221021-1100221330232332-3012313232012012"></a>

<a id="canonical-3123321000200221-1210102331303122-2213313132032332-3101310232012321-2131030310213321-3212003020320033-0123313130130300-2012011212230100"></a>

## exact_value property — http_list / 233112201231 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2320022213233012-0322202111111030-1200120121320010-0113330311331122-1031312313221210-2313321103113303-3112000020201012-1301213301213202"></a>

<a id="canonical-1110021230233033-1300032033113022-1011102011221211-2312100030123020-3130210030023132-1211230020102211-0321032312321222-3220202133232131"></a>

## path_exact_value property — http_list / 233112201231 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3132121311033121-3120301110033132-0102212120000101-3312202311331031-2111202223010120-3020203130101132-2021221301302023-3013112212103301"></a>

<a id="canonical-3230200010003303-0220231023001302-0121132331320302-1101130232201301-0311211321023111-2121113010330221-2110032132333303-0210303030101301"></a>

## path_prefix_value property — http_list / 233112201231 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2021322001010223-1102122220313113-0002011031023032-0023033112032131-2310303100203000-0320002322210010-1122212322031111-3203321303311121"></a>

<a id="canonical-2103131323110222-1032223012303322-2121332102032010-2030321320330233-3332111023111211-2012331232033301-0100023003011221-0013011302233001"></a>

## path_regex_value property — http_list / 233112201231 / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0332333103100012-2312313103131021-2113221310211121-0331110130223012-0101131233023310-1220212302132313-1113130203022322-3112013311003012"></a>

<a id="canonical-2311021033203320-1200012332003100-0300320031301233-0113001333031232-0313311200302021-2313003322011210-1303132113223202-1331132111010012"></a>

## regex_value property — http_list / 233112201231 / 8

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2310300300033331-0000233013321103-3312103120232011-0311212021010133-3311112221032210-2000321131321132-1332332311011001-2333323101300121"></a>

<a id="canonical-1213300220231212-3000131111101203-0203012231013031-1013230030030102-1312003010013021-1230311311212133-0321131212320110-2222201313331001"></a>

## suffix_value property — http_list / 233112201231 / 9

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2231222331121321-3030100212303332-0212203323313333-3232213201133210-2321200132311000-0321112322131331-1312121302202210-0032301322221302"></a>

## Next pages — http_list / 233112201231 / 10

- [deny_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-0221121112200211-3020121331300210-0322123232233121-0300231231222130-0200311103230333-0221311123211303-1233113230022321-2100011111111202)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0221121112200211-3020121331300210-0322123232233121-0300231231222130-0200311103230333-0221311123211303-1233113230022321-2100011111111202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032201001303330-3133133022111131-3011202332310301-2231221031233200-0120213132001012-3113212000132211-3231310031011013-1030001120231013"></a>

## deny_list.http_list.any_path — any_path / 013303030221 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- [deny_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3202121002221232-1001233131102222-3332221121223212-0111023100312303-2220221211222132-2312230221111030-0022112301033230-3001113000001013)
- deny_list.http_list.any_path

<a id="canonical-2200020100211301-2031101210211110-3111201232030112-3003120331203111-0000030303210023-2202321313230002-1111031233111033-3120212110212010"></a>

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
any_path = {}
```

<a id="canonical-2232230012132111-3233202113323100-3312000011210133-3110103101331300-3210113233021221-3310010323122323-1322101032201313-2310111201333030"></a>

## Direct properties — any_path / 013303030221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003000032211132-0032010310031133-0332231311213211-1202102031233220-3302333003123310-1101331030033133-3212322201300031-0221022100003310"></a>

## Next pages — any_path / 013303030221 / 4

- [deny_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3202121002221232-1001233131102222-3332221121223212-0111023100312303-2220221211222132-2312230221111030-0022112301033230-3001113000001013)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2030011123202013-2020022321113210-3000123000322103-0332101233032232-3031200322302233-3320333113121232-0103133001223023-3020331222230303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022300121231002-1203023111222322-1303021130000013-0123210331213023-1322321130231120-1301131300212321-0121200303020032-1300033321100312"></a>

## deny_list.tls_list — tls_list / 100310020012 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- deny_list.tls_list

<a id="canonical-0210032332022302-3302131202312221-0131210122311212-1000212112031320-3113023321023201-1122222102120103-0331023030320132-0000033130133201"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Upstream description:

Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
```

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
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210331301013132-3001312313312002-0313211030000201-1111203120233033-1000113123203203-3121130233300200-3331313131123021-2332203302103123"></a>

## Direct properties — tls_list / 100310020012 / 3

<a id="canonical-0310003310121110-0111121223231133-1233011121122300-0223010210010231-2203231310003200-1210112022301022-0013221020113012-3001302202203200"></a>

<a id="canonical-3311013323100131-2033120322031320-1301003210032113-1032003230323203-1133323310223203-2201332201031110-0010303312002003-0100300012032313"></a>

## exact_value property — tls_list / 100310020012 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0303320220031210-0223033312202100-0232213200310313-0030112222332131-3001213000011033-0230101000112202-2130002321030001-0220232230003323"></a>

<a id="canonical-3023230112131012-3032211301113012-2310123233322302-0013133300132211-0312132220020202-0232200001133022-2120201120010120-2311001230210132"></a>

## regex_value property — tls_list / 100310020012 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3302231102333112-0102300003323303-0132123320113030-0033033200110312-3022022211213212-3012201203203213-1201102131333301-2113031312310112"></a>

<a id="canonical-1303122100031321-1232103311110321-0311010300230013-2023031331113130-2013200331130001-1130320020223212-2131302111302110-3213221012132123"></a>

## suffix_value property — tls_list / 100310020012 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0312012003001331-2301203110333232-3113003012121220-0302232103333212-1033111300001113-1320111321303330-2011112232221121-2113000132112011"></a>

## Next pages — tls_list / 100310020012 / 7

- [deny_list](resources--forward_proxy_policy--reference--group-001.md#canonical-3020120030321321-2120102321010031-3202220001101331-0300302332020322-0012012033030332-3200110122022021-0300012113131201-0010300210310110)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0121000321310323-0220232133302100-1210330110221132-1112130320100200-1022002100133130-3210200311003310-3011210001322311-1121030003102231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130320211111203-3322232010332230-1110213222031220-1323011312333111-2013220203031021-0031030101210111-0330103232211113-3003131200011201"></a>

## drp_http_connect — drp_http_connect / 132001201101 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- drp_http_connect

<a id="canonical-2213200132201033-0010330133330103-3302333011021032-1303202212331123-2132220013311200-2312220130321232-1031202232211113-1122101231012322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
drp_http_connect = {}
```

<a id="canonical-0132313221333123-1313110213201120-3301011232012103-2203213121131101-1223121232023112-1133213221102302-0021230021112213-1212132120032020"></a>

## Direct properties — drp_http_connect / 132001201101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120021000013211-0313332203032302-3300303002000320-3111200002002221-1133212300022222-3002331232020133-3201203232020303-3203300021200310"></a>

## Next pages — drp_http_connect / 132001201101 / 4

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-1023113102302011-2230123310331013-2321010211321003-0002231332301103-2321203320301010-2301303233002003-1220120223230121-3222022320233033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032230322222032-1123331112302001-2211322013323230-2313112330000331-1231102113301233-2000311321312112-0111303010230101-2011202131020103"></a>

## network_connector — network_connector / 331011010222 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- network_connector

<a id="canonical-1321130202130211-0232301301212322-2320303122132031-2011310313212200-0211110113010100-3213131112122303-0231300030003123-3302323112321201"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
network_connector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031133001213101-2221313330112230-3022103013032220-3201132023211223-3133322310122111-0333312323323222-0031232202030313-3331031211223123"></a>

## Direct properties — network_connector / 331011010222 / 3

<a id="canonical-2122320303201123-2322302110023012-3212033312212330-2122302330333113-3031331100111133-1332310201210320-1130132101210030-2001013012212312"></a>

<a id="canonical-2112223111132323-1321310310211220-3300303010033321-1113033223300331-2110303212313321-3121201213123032-2301302323023302-0013222302333300"></a>

## name property — network_connector / 331011010222 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0312101020102200-3323200132322002-3200012302010311-2032132322021323-3033230123103212-2132112031132012-3320123020303230-2302232333300213"></a>

<a id="canonical-0222103230232121-3113030033231023-2332201111021313-3323333322110030-3320311222110211-0332032311202011-1133030003233123-1322031302033122"></a>

## namespace property — network_connector / 331011010222 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1312223200121302-1210131332133231-1000332201030203-3133320310210023-3201331230300012-3231311013020220-0130101333221302-1300130212301303"></a>

<a id="canonical-0001110232001313-0120100222310230-1220123333020102-1122322003001312-2010333223202303-2300112000132223-2023210333021232-1231132321312133"></a>

## tenant property — network_connector / 331011010222 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3100232133220310-0202201301130323-3112103233130011-3020100331131011-1100230113311111-1200230221230333-1332332002033210-0211230103322310"></a>

## Next pages — network_connector / 331011010222 / 7

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0221230323223211-3010132100100233-3210122132012122-1223331132311301-3301002003203030-3311233232103222-1102003000313200-0033331023010010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311312113310003-1023132120210200-2013013303013130-2330301313010201-3213133020131103-1223232211010322-0201031302222310-3003232022213320"></a>

## proxy_label_selector — proxy_label_selector / 033012112133 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- proxy_label_selector

<a id="canonical-3111002303022031-3313302320131202-0002322100332212-1022112321301020-3203320320321333-0222031133023001-2312200122230321-1032320001202331"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
proxy_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003102133310233-2333101230301232-3333131100313211-0023132133200100-2020130022200301-0220301220002110-1222233032031200-3331122002313023"></a>

## Direct properties — proxy_label_selector / 033012112133 / 3

<a id="canonical-3201211130202332-2202302330013100-0300022120303313-1132021123100111-0200313000200102-2313001132212200-2112010202220210-1021223303003112"></a>

<a id="canonical-0100233301133132-0303111323303303-1311312312002210-2030211112113321-1303000011032031-1211032002132213-2112112032102333-0320230310230231"></a>

## expressions property — proxy_label_selector / 033012112133 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-2302232110220222-1001111131103232-0121201001202112-0121203230011120-2113020133220011-2010330220033202-3323031231200102-3010031230212123"></a>

## Next pages — proxy_label_selector / 033012112133 / 5

- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223002031000023-1313320302122313-0300332110220013-2122303312221303-0313310233101131-3020201203003021-2131302130002032-1021322230123032"></a>

## rule_list — rule_list / 131120321131 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- rule_list

<a id="canonical-3120101031021000-0311222230303232-3110212032012330-3000223220331220-2233310120133312-0001101213022230-2123233112313210-2100201222030032"></a>

Type: `"object"`. single nested block, Optional.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101202002000202-2312133211333121-1330321211101211-1132213323100001-0032231001101323-0030011233012022-2103202312223132-2221031010101113"></a>

## Direct properties — rule_list / 131120321131 / 3

- [rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330): complete subsection reference.

<a id="canonical-1133033300210130-3022332121200130-3332323032010322-3311130330301321-2010330220322102-3221010200313103-1222212202113113-0033232112031230"></a>

## Next pages — rule_list / 131120321131 / 4

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300322033120023-2302101301222021-0021102213010021-1302000201212331-2313321233011312-2222110001123022-2310010310131302-3301103002200022"></a>

## rule_list.rules — rules / 223001230323 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- rule_list.rules

<a id="canonical-3312120222023102-3011020210320310-3103002210201000-0003220102311011-3111011031221010-0103012112312202-3021233312200333-0333313230200112"></a>

Type: `"object"`. list nested block, Optional.

Custom Rule List. List of custom rules.

Upstream description:

List of custom rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_destinations",
    "dst_asn_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_asn_set"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "http_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "tls_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("all_sources",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sources",
    "label_selector"),
  validators.ConflictingListObjectAttributes("all_sources",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_asn_set"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_asn_list",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "dst_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_asn_set",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "dst_label_selector"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_ip_prefix_set",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "dst_prefix_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_label_selector",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("dst_prefix_list",
    "http_list"),
  validators.ConflictingListObjectAttributes("dst_prefix_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("dst_prefix_list",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("http_list",
    "tls_list"),
  validators.ConflictingListObjectAttributes("http_list",
    "url_category_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("no_http_connect_port",
    "port_matcher"),
  validators.ConflictingListObjectAttributes("tls_list",
    "url_category_list")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010200031331102-3202132030233103-3003331233322123-2300210321002321-3013122201011300-1332322122221212-3301331321202301-3000201000221310"></a>

## Direct properties — rules / 223001230323 / 3

<a id="canonical-1210301032202321-3321023231033121-3012222222312121-2213000223102330-3221112323212121-2220323212113001-0301111201301003-1021313123100030"></a>

<a id="canonical-1020200222122322-3033022203031223-3100112010311002-0330000232202003-3010101101330310-2220233212322200-2323210222000022-1312032221300230"></a>

## action property — rules / 223001230323 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW",
    "NEXT_POLICY"),
}
```

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

- [all_destinations](resources--forward_proxy_policy--reference--group-001.md#canonical-3303212022323032-0022002332220301-1322021013031312-2132120202033330-2121131310233020-0103000131331110-2012330122300033-1031202211311121): complete subsection reference.

- [all_sources](resources--forward_proxy_policy--reference--group-001.md#canonical-2222002001020101-1111232232301313-0120210021110313-1122331022321320-1033330212102033-2100023110012212-2222020302330033-1303301022233201): complete subsection reference.

- [dst_asn_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2011102313313211-2131223301223322-0233232012200022-3213221220130021-0300133023100031-0302010133211122-0000302210010220-0302330123310033): complete subsection reference.

- [dst_asn_set](resources--forward_proxy_policy--reference--group-001.md#canonical-0003132101030212-3123122112331210-3222323123002210-2001112030002030-0201331012012233-1132230211211100-0300021210200213-3101023102002203): complete subsection reference.

- [dst_ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-2212102330010013-2031123132223301-0131212022001000-2233000012210232-0111000131320100-3220031130123032-3123221133100232-2303223202311211): complete subsection reference.

- [dst_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-2110220112000323-1230101213332333-3323222310110320-3030023320300210-3110111103122013-0113132230312113-0332013123321023-3023112213221112): complete subsection reference.

- [dst_prefix_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2220013131103110-0321123030130202-1211323031103321-3120101110022022-2312300220103132-3303110213302112-2320201002022211-3233200233130211): complete subsection reference.

- [http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220): complete subsection reference.

- [ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-0310221030310313-2132220021203023-2232111022112130-3332333211121113-3031033033311333-0320210103020113-3230210210002002-0321121231301103): complete subsection reference.

- [label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-2112013033001320-0321130123000130-0033123321022312-1123023302201232-3023123010010201-0202333213232300-1112213012003221-3122233032102202): complete subsection reference.

- [metadata](resources--forward_proxy_policy--reference--group-001.md#canonical-3030113001113311-1321300323302332-1203022110332230-1131331102123232-3132221000110231-0333211230303301-3323103223120032-0220113312111300): complete subsection reference.

- [no_http_connect_port](resources--forward_proxy_policy--reference--group-001.md#canonical-3310331010213110-3101213300003222-3021120021321003-0333023333133021-2032303012222011-1332203203222230-0321133031203200-3003323103232221): complete subsection reference.

- [port_matcher](resources--forward_proxy_policy--reference--group-001.md#canonical-1310001202000220-0222212030212031-3110301233213311-0331231201211302-1222330112030113-0332221131200032-0000000110301303-1000000231020321): complete subsection reference.

- [prefix_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0312131012032312-0131110323211333-0123311001012011-2103130111123002-2020111100012001-3211231333323102-3111311122001102-1233223101130020): complete subsection reference.

- [tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0321220033202311-3123011210302013-3200210310012232-1230303120121032-2003133313330030-3321320211122231-2200231330033213-2211001300113003): complete subsection reference.

- [url_category_list](resources--forward_proxy_policy--reference--group-002.md#canonical-3133120013122023-1010000233103233-3100211201111331-1132333033310100-1001321120000211-0013313110101113-0312320322323313-0031022220302232): complete subsection reference.

<a id="canonical-1303122000031220-0112101210122333-3200013031202123-0211102123110312-0200103310200133-3302122330133020-2330311323211010-0013311202031010"></a>

## Next pages — rules / 223001230323 / 5

- [rule_list.rules.all_destinations](resources--forward_proxy_policy--reference--group-001.md#canonical-3303212022323032-0022002332220301-1322021013031312-2132120202033330-2121131310233020-0103000131331110-2012330122300033-1031202211311121)
- [rule_list.rules.all_sources](resources--forward_proxy_policy--reference--group-001.md#canonical-2222002001020101-1111232232301313-0120210021110313-1122331022321320-1033330212102033-2100023110012212-2222020302330033-1303301022233201)
- [rule_list.rules.dst_asn_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2011102313313211-2131223301223322-0233232012200022-3213221220130021-0300133023100031-0302010133211122-0000302210010220-0302330123310033)
- [rule_list.rules.dst_asn_set](resources--forward_proxy_policy--reference--group-001.md#canonical-0003132101030212-3123122112331210-3222323123002210-2001112030002030-0201331012012233-1132230211211100-0300021210200213-3101023102002203)
- [rule_list.rules.dst_ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-2212102330010013-2031123132223301-0131212022001000-2233000012210232-0111000131320100-3220031130123032-3123221133100232-2303223202311211)
- [rule_list.rules.dst_label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-2110220112000323-1230101213332333-3323222310110320-3030023320300210-3110111103122013-0113132230312113-0332013123321023-3023112213221112)
- [rule_list.rules.dst_prefix_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2220013131103110-0321123030130202-1211323031103321-3120101110022022-2312300220103132-3303110213302112-2320201002022211-3233200233130211)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220)
- [rule_list.rules.ip_prefix_set](resources--forward_proxy_policy--reference--group-001.md#canonical-0310221030310313-2132220021203023-2232111022112130-3332333211121113-3031033033311333-0320210103020113-3230210210002002-0321121231301103)
- [rule_list.rules.label_selector](resources--forward_proxy_policy--reference--group-001.md#canonical-2112013033001320-0321130123000130-0033123321022312-1123023302201232-3023123010010201-0202333213232300-1112213012003221-3122233032102202)
- [rule_list.rules.metadata](resources--forward_proxy_policy--reference--group-001.md#canonical-3030113001113311-1321300323302332-1203022110332230-1131331102123232-3132221000110231-0333211230303301-3323103223120032-0220113312111300)
- [rule_list.rules.no_http_connect_port](resources--forward_proxy_policy--reference--group-001.md#canonical-3310331010213110-3101213300003222-3021120021321003-0333023333133021-2032303012222011-1332203203222230-0321133031203200-3003323103232221)
- [rule_list.rules.port_matcher](resources--forward_proxy_policy--reference--group-001.md#canonical-1310001202000220-0222212030212031-3110301233213311-0331231201211302-1222330112030113-0332221131200032-0000000110301303-1000000231020321)
- [rule_list.rules.prefix_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0312131012032312-0131110323211333-0123311001012011-2103130111123002-2020111100012001-3211231333323102-3111311122001102-1233223101130020)
- [rule_list.rules.tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0321220033202311-3123011210302013-3200210310012232-1230303120121032-2003133313330030-3321320211122231-2200231330033213-2211001300113003)
- [rule_list.rules.url_category_list](resources--forward_proxy_policy--reference--group-002.md#canonical-3133120013122023-1010000233103233-3100211201111331-1132333033310100-1001321120000211-0013313110101113-0312320322323313-0031022220302232)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-3303212022323032-0022002332220301-1322021013031312-2132120202033330-2121131310233020-0103000131331110-2012330122300033-1031202211311121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121300330032000-0023333333010012-2212220300302033-0032021112001003-3233323020223230-3013001213201231-2131312001033222-1010331332312030"></a>

## rule_list.rules.all_destinations — all_destinations / 013133033330 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.all_destinations

<a id="canonical-3132030112213100-2310103330022200-2131003300123302-1132133201031302-2103020101332133-1302013203011222-3012230100000133-1323023210222320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_destinations = {}
```

<a id="canonical-1033213230230103-2132233222030000-2122203101111133-1301033012031130-3302101312120020-0011110323010202-1121120222323110-0001303221200021"></a>

## Direct properties — all_destinations / 013133033330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110303213001120-1133110230220212-3331022211123202-0103013021223011-1303223320101022-1301013233202200-1211230300130331-1311310223022122"></a>

## Next pages — all_destinations / 013133033330 / 4

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2222002001020101-1111232232301313-0120210021110313-1122331022321320-1033330212102033-2100023110012212-2222020302330033-1303301022233201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130010012013010-2032031001132310-3122320003110002-3023132300332032-0010300020133032-2212323113132110-1332012102022233-1112220022230320"></a>

## rule_list.rules.all_sources — all_sources / 121313312200 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.all_sources

<a id="canonical-3212102110303220-1013322223223200-1012020002331200-0212001203213020-2323303320021303-3322102310313321-1021033112030103-3312210203030223"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_sources = {}
```

<a id="canonical-1322112201211130-2101130022113211-1130012032300312-1310102322322322-1121320313003232-3012001121101331-0231310123221131-2131303102223202"></a>

## Direct properties — all_sources / 121313312200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122202223113323-1131232100133322-1332202331033033-2003210233322103-3123012311230012-3202203221221031-2322012000000131-1001212201113202"></a>

## Next pages — all_sources / 121313312200 / 4

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2011102313313211-2131223301223322-0233232012200022-3213221220130021-0300133023100031-0302010133211122-0000302210010220-0302330123310033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003132301232130-1200001233002332-1203013031020132-1123001131011213-1130103221233133-0131122230120000-3200003130130133-2220022330021232"></a>

## rule_list.rules.dst_asn_list — dst_asn_list / 013123021131 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.dst_asn_list

<a id="canonical-0133131032312222-3031301123203003-2313303030213010-1313332220210321-3002120022010323-0202200023303100-1030300300322002-1200331011132330"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
dst_asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111003123201123-2302231003310012-3223230333333203-0203301023200310-1133310313320002-3012221101333131-1010232102122113-0032203100210313"></a>

## Direct properties — dst_asn_list / 013123021131 / 3

<a id="canonical-3122100001023130-1031310110202203-2320323301313322-2101212013320131-3303133021033233-0310230113000310-1322033110313122-0030033232102132"></a>

<a id="canonical-0111011011300231-3110220331112132-1003211233133021-3102002013101312-1000231212012033-0312330112021320-0121001022000021-0222003120020232"></a>

## as_numbers property — dst_asn_list / 013123021131 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-3111022310020031-1122010011203121-1030330130100003-1023232010301333-3231202133132321-1231113202213302-3221311211020313-3323202322321301"></a>

## Next pages — dst_asn_list / 013123021131 / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0003132101030212-3123122112331210-3222323123002210-2001112030002030-0201331012012233-1132230211211100-0300021210200213-3101023102002203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110113321033112-2011001302312032-0200033122020011-3012321322230020-3002120202130020-0321103023111022-3322320132233101-3212310120313012"></a>

## rule_list.rules.dst_asn_set — dst_asn_set / 302022100030 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.dst_asn_set

<a id="canonical-2331112120122301-0133313301032221-0012303212331010-3021321001312212-2011002333330103-1011203112230202-1012223023111230-3122202012111221"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
dst_asn_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210303322031023-3113003303031033-1131233030232021-3030212030211003-1322123112200033-0033300100301302-0330211232310111-2103313300113212"></a>

## Direct properties — dst_asn_set / 302022100030 / 3

<a id="canonical-2322000130221222-1321221031303133-0032001101310020-2130212203122212-3211100001133102-1000122221100303-1112313001031313-3032303020230001"></a>

<a id="canonical-3011112233132201-0032321323301231-1331203211300330-2120203233203323-3132202222132103-1220332012233110-1200313200003313-1122032101232120"></a>

## name property — dst_asn_set / 302022100030 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-2001131231213200-0003031233302220-0010301231230113-0233023021113333-0222111301013010-2321132202331030-0121201201113332-2332233323310122"></a>

<a id="canonical-2213003021012301-0211003320011110-0211320330002130-1310231232033221-3030013333302001-1011003131320023-3213100023133233-0011113103331032"></a>

## namespace property — dst_asn_set / 302022100030 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3033301130331101-0201132001103210-3012030332110000-0031113331303210-3331302031021301-3023201223030311-0302000312031001-1010031303212302"></a>

<a id="canonical-0221031200031331-0210023203322312-1101013131010301-2023010121002020-0202110223101223-1022112013020002-0310322132220100-3033102122233331"></a>

## tenant property — dst_asn_set / 302022100030 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3013003320002302-0332101330321131-2232010311113301-0130120123030120-2132030110330211-3222221312002102-3212030323133000-1020133332131202"></a>

## Next pages — dst_asn_set / 302022100030 / 7

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2212102330010013-2031123132223301-0131212022001000-2233000012210232-0111000131320100-3220031130123032-3123221133100232-2303223202311211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020001111210133-2321100111000002-3320001132301210-3010203222002300-1011221023101320-0230032311223223-2001233221013221-1311020320303301"></a>

## rule_list.rules.dst_ip_prefix_set — dst_ip_prefix_set / 330033203320 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.dst_ip_prefix_set

<a id="canonical-0011232133000131-0011100200323101-3030033303023333-2022302203321221-0302332210113131-2213100122312100-0002032010322222-0203323300320333"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
dst_ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111010032002320-3302200020221100-0301203110331303-1003031102030000-0121030310311101-3323200203001331-2301030112201321-1331231102113303"></a>

## Direct properties — dst_ip_prefix_set / 330033203320 / 3

<a id="canonical-2311223211201033-0000130323333133-1320113031333113-2303021333301121-0322023201012113-1321003022001233-2122312210302132-0212312201131203"></a>

<a id="canonical-3312102212101312-2002132002312321-1303310300103230-2031112133103110-1323031220213213-3232201310000102-0301122233010310-1231000000113000"></a>

## name property — dst_ip_prefix_set / 330033203320 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3232300201110110-3313012332102222-1300330111333200-0012222112102123-3120123231220133-0033002310121010-2211122013030021-3233003303032033"></a>

<a id="canonical-3323032013200002-3312223220212300-1032230032021102-1110213013223221-1023232210000231-3101221221013032-1100002321212232-0110011122113323"></a>

## namespace property — dst_ip_prefix_set / 330033203320 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0113010212002230-0211133322030033-2210022102230201-0330111011313212-0010131203211201-1201222120201001-2023221113111110-3110322102132021"></a>

<a id="canonical-0130013021201010-1321121201202323-3101120122202233-3213220001132123-2200323013331100-1101332301311222-1301212031212310-1010121002010101"></a>

## tenant property — dst_ip_prefix_set / 330033203320 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2023033033301333-0022311103330230-3333323132002230-0113200123313030-2012200131320003-3133122233233221-1012021211302213-2212113323112030"></a>

## Next pages — dst_ip_prefix_set / 330033203320 / 7

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2110220112000323-1230101213332333-3323222310110320-3030023320300210-3110111103122013-0113132230312113-0332013123321023-3023112213221112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130102032322232-2330003002311010-2213001033123112-0221033111220013-0123223111321013-3201222023133203-0233233032222202-3201023231312331"></a>

## rule_list.rules.dst_label_selector — dst_label_selector / 223013232031 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.dst_label_selector

<a id="canonical-3333031302001101-0030110031100012-3200112300222120-1330112310103100-0132330023312120-2222301223332202-1300021322022100-3323210112201323"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
dst_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020130012303120-0103122111000232-3032233102333312-3022121021121330-2131132113130031-3120221302111121-0113002022033223-3130303202110012"></a>

## Direct properties — dst_label_selector / 223013232031 / 3

<a id="canonical-0323111132202312-0031303221201210-3021012033322010-1310333200133110-2202123021111223-1323330200131233-0112302030000220-1003210113132111"></a>

<a id="canonical-1003310311030000-3330330222030332-0122233320031203-1023313320300213-2212233003112232-3032121101030222-1112101300102013-0213301102232330"></a>

## expressions property — dst_label_selector / 223013232031 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-0213311123012231-3100000003000011-2032312221232323-1200122233033011-1020212323200012-2021120310211011-2331233302120011-2301012023311021"></a>

## Next pages — dst_label_selector / 223013232031 / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2220013131103110-0321123030130202-1211323031103321-3120101110022022-2312300220103132-3303110213302112-2320201002022211-3233200233130211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100200021001013-3201110231320211-3011300101320221-1310301300130102-0333223003232101-0133032223200031-2111300111203230-2313330102112020"></a>

## rule_list.rules.dst_prefix_list — dst_prefix_list / 120012022001 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.dst_prefix_list

<a id="canonical-3133300003333203-3230030332132300-2132312332213112-1001310301133020-0313101113230210-1111313302213121-2123212301312212-0321233021211211"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
dst_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310120000100301-2333012133030013-3110322122100213-0132322102212331-2110002122322111-2303121223231201-3030212313310023-2303111131330110"></a>

## Direct properties — dst_prefix_list / 120012022001 / 3

<a id="canonical-1123201213103230-0212322122321330-2123100330003321-2322212312312213-1003202210102321-0231113301131021-2333113331203202-3123302003313213"></a>

<a id="canonical-0023231133320122-1210210113301232-3010200323310022-0301233000122323-0132320301112022-0223303310311101-0230333111031121-1330212212302333"></a>

## prefixes property — dst_prefix_list / 120012022001 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-3013211020301202-2322131102032202-1201232003322323-3110320312010030-3231203223121021-2111123111321100-0201321000031021-2313132302013123"></a>

## Next pages — dst_prefix_list / 120012022001 / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002320320313230-3110310102220102-3211132113331322-0111103330331100-3023131221232013-0322103101321322-2231023310131331-1323303013310031"></a>

## rule_list.rules.http_list — http_list / 011310123222 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.http_list

<a id="canonical-0303221302121213-3300222312030333-0201113302021220-0233320100320221-3132010201031003-0121131020030220-2113330031013200-2121113112200120"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112321223110210-0302331223021331-2003002300023211-3101222203023121-1330211101331031-2311331122120032-0032021022100303-1100221232203311"></a>

## Direct properties — http_list / 011310123222 / 3

- [http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2233313011311201-0112331132232311-1302312313130130-3213111133013310-0221230302221331-2221103121012302-0200000031133101-1313303121032132): complete subsection reference.

<a id="canonical-2122231131310231-1331312323213220-0211003302020012-3202320013133111-3020011021121232-0020110233031122-2332100233121203-0022301103112111"></a>

## Next pages — http_list / 011310123222 / 4

- [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2233313011311201-0112331132232311-1302312313130130-3213111133013310-0221230302221331-2221103121012302-0200000031133101-1313303121032132)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2233313011311201-0112331132232311-1302312313130130-3213111133013310-0221230302221331-2221103121012302-0200000031133101-1313303121032132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022110320321102-2133122032211102-1200201101033002-0111131201323121-3301333033013331-0322030101032013-1032112300312310-1032332002220122"></a>

## rule_list.rules.http_list.http_list — http_list / 110003221120 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220)
- rule_list.rules.http_list.http_list

<a id="canonical-3233101000301203-2221302110220122-3120023122101002-2103030031301331-2102112332233320-2132233021113322-3300223022101021-1110310001131020"></a>

Type: `"object"`. list nested block, Optional.

HTTP URLs. URLs for HTTP connections.

Upstream description:

URLs for HTTP connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_path",
    "path_exact_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_prefix_value"),
  validators.ConflictingListObjectAttributes("path_exact_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("path_prefix_value",
    "path_regex_value"),
  validators.ConflictingListObjectAttributes("regex_value",
    "suffix_value")}
```

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
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231202220002233-3232313002310112-0332310322331222-2012210203302310-0220213220120011-2333220002220332-2201011021003331-1221030013201213"></a>

## Direct properties — http_list / 110003221120 / 3

- [any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-2003033123330321-2120103022031003-0010013112313223-0333323210322231-0022213123230131-3321212132120223-3102310210211020-1003331133201333): complete subsection reference.

<a id="canonical-0110303312312100-3302332023333011-3013211102110311-1332321111223000-1211030231013022-0130233232003230-0302010002233123-1002000330312222"></a>

<a id="canonical-2230030333103300-0131220210001032-0321301120102300-2231322212101123-3110231323020200-0011111212302031-0003010301100202-0203221011021221"></a>

## exact_value property — http_list / 110003221120 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1030230101321100-1320202313303213-0103113132013002-0303123013001311-0230201220102033-0222131300323020-3001221033201333-2223323001002233"></a>

<a id="canonical-3321330203331300-1230110133301003-0220001133112331-0131231011223332-2011110210022212-0010301132020331-3310333210302130-2111320103132323"></a>

## path_exact_value property — http_list / 110003221120 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Upstream description:

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1023000131030012-0201003310033101-0033320003332302-3031201313032022-1001201121030000-2312330012311310-1013022130120212-0100002203013231"></a>

<a id="canonical-1211111010101011-3233130310020212-3313332120300222-1001332313230112-2220221211230002-2333102102001213-3212220032333301-1311213121313223"></a>

## path_prefix_value property — http_list / 110003221120 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2321103313110001-0220320211301003-1231110312331110-2222213123220123-3013113011312202-2133321323331223-1102220123121220-0331112123220121"></a>

<a id="canonical-1003101121200230-2112021301223332-3101320102300113-3332023133112030-3100323012303102-3200020031231112-2230222313331230-0031002103030033"></a>

## path_regex_value property — http_list / 110003221120 / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Upstream description:

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3103123221200021-3003232300100022-2111010130321012-0023121232312321-1312133230011200-0322300122033311-3020100022021201-3222312221301311"></a>

<a id="canonical-0312011322111022-0322023312310211-2232202221202210-2200321022223102-3310122001103131-2123111100333101-1320300222010313-0100310220133132"></a>

## regex_value property — http_list / 110003221120 / 8

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3312001323010103-2001203101301300-3011033023130120-0111033103113213-2003321023232010-1201123221220112-3103132112103011-0231230112101223"></a>

<a id="canonical-3031311013000023-0101032000011211-2233331220123023-1203001102012212-0322132103230131-1012321320221212-1011123131121332-2021311332200122"></a>

## suffix_value property — http_list / 110003221120 / 9

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1221112122220101-3213133132211002-1121213231103012-2203300212133000-2201211322033211-2301233220233111-1122302132321303-1220203123302203"></a>

## Next pages — http_list / 110003221120 / 10

- [rule_list.rules.http_list.http_list.any_path](resources--forward_proxy_policy--reference--group-001.md#canonical-2003033123330321-2120103022031003-0010013112313223-0333323210322231-0022213123230131-3321212132120223-3102310210211020-1003331133201333)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2003033123330321-2120103022031003-0010013112313223-0333323210322231-0022213123230131-3321212132120223-3102310210211020-1003331133201333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201122232021213-1233320230121333-3312103223001222-1112100121023132-0013130132002313-1011303100120003-1100301320230120-2123013312130203"></a>

## rule_list.rules.http_list.http_list.any_path — any_path / 111120021302 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220)
- [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2233313011311201-0112331132232311-1302312313130130-3213111133013310-0221230302221331-2221103121012302-0200000031133101-1313303121032132)
- rule_list.rules.http_list.http_list.any_path

<a id="canonical-3202321113213112-0133121003330220-1110021232002330-0003032332330231-2030303330111010-0032032002123322-3231021102212313-1330020232232221"></a>

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
any_path = {}
```

<a id="canonical-2320322233033202-3203332200133002-2121200313233313-1102113201032032-3301012312203221-0130313030002002-2202100102001021-2133022123311021"></a>

## Direct properties — any_path / 111120021302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213013100201020-2321002012032321-3310211313022131-0003132331131121-0020312210200321-3220320232201103-3022323031220300-2311103000102100"></a>

## Next pages — any_path / 111120021302 / 4

- [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2233313011311201-0112331132232311-1302312313130130-3213111133013310-0221230302221331-2221103121012302-0200000031133101-1313303121032132)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-0310221030310313-2132220021203023-2232111022112130-3332333211121113-3031033033311333-0320210103020113-3230210210002002-0321121231301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220221300000200-3312130203033302-0302231112100230-3100212132210211-1112130020310233-1213300213311122-0003320322220332-2332230213302333"></a>

## rule_list.rules.ip_prefix_set — ip_prefix_set / 212113300011 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.ip_prefix_set

<a id="canonical-3121023022033302-1031312030100132-0330222032021033-1032003012313021-1111132222010301-1201132021123313-3320303010120013-3311010121132023"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220131003222210-3132013323213110-1210321313030121-3321223011002221-1002011120103211-0021331322202203-2330010233333032-3012111013113331"></a>

## Direct properties — ip_prefix_set / 212113300011 / 3

<a id="canonical-2012010232113010-0022122301320321-2332020121012222-2012112123021011-0110010201230221-3100302001030100-3101220203020122-2322031231120023"></a>

<a id="canonical-0323102032112301-3201003002030131-2331323133011303-2100113101121331-2223120032332223-3003321013012013-1333132230122010-0232032002200203"></a>

## name property — ip_prefix_set / 212113300011 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0333111100230131-1132003010020103-2013122011020022-2210320330111033-3313310322023312-0233022323031230-2033211011121300-0303103022002103"></a>

<a id="canonical-2031100201011023-0330122211021020-1020301211100331-3331203100023033-0130002023203213-0310201310031102-3010223322312201-3321022031203311"></a>

## namespace property — ip_prefix_set / 212113300011 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3200333102332123-3301130310212333-3122102030101313-0300002032000321-2311210311332130-2232320100322102-1012333101323210-1302001331302021"></a>

<a id="canonical-0103102201130001-0011103202010023-3323033033131130-3100211003302101-3310031021112032-3223031301330302-0232333212200113-1020330313110003"></a>

## tenant property — ip_prefix_set / 212113300011 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-0120313222122202-2221331012300200-2103201133300221-3303010010311122-1032232032022230-3310230231103330-0110310313310320-2302101122203113"></a>

## Next pages — ip_prefix_set / 212113300011 / 7

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-2112013033001320-0321130123000130-0033123321022312-1123023302201232-3023123010010201-0202333213232300-1112213012003221-3122233032102202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033021101000110-2123233022020313-3100110013303321-1121210102011330-0232213221012022-2112120202133330-0013013130202133-1123211323102201"></a>

## rule_list.rules.label_selector — label_selector / 022200301021 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.label_selector

<a id="canonical-3121111211111121-1122302203201112-3102332311230230-3203303120100331-3021222232232112-3133330102113323-2112203100221130-3322102323110331"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132012113312113-3203012123200133-1301132221011123-2320200320313310-2310100102131120-2023301133300331-1223113032303002-2030302312313132"></a>

## Direct properties — label_selector / 022200301021 / 3

<a id="canonical-2301003312033221-2123222132203212-3121001210320023-2103033321110233-1230021311201000-2002231310010012-3003212231231022-2313022330311321"></a>

<a id="canonical-2122312123230321-1300200212010000-1230310221212321-3103003300301021-3220111110000100-1102123011131212-1112213333221321-3133332200333022"></a>

## expressions property — label_selector / 022200301021 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-3133330323011022-1101003112333010-0212231320131332-3002310311220212-3101220211230133-3331330122021220-0332103131011203-1301123111131113"></a>

## Next pages — label_selector / 022200301021 / 5

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-3030113001113311-1321300323302332-1203022110332230-1131331102123232-3132221000110231-0333211230303301-3323103223120032-0220113312111300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330021220320032-3313003311333003-0332000232321113-2221031322133022-0033230010130132-1032211333313120-1310122233023021-2012130203322213"></a>

## rule_list.rules.metadata — metadata / 131111223210 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.metadata

<a id="canonical-1332230121002020-3211213221030200-2100000310310030-2010212011132131-0020012313222302-2212122211112020-2130030012301333-1312320023312112"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321100210323302-0120301230333012-0203001132113133-2003011200211212-3020022202220001-3122212111221303-0122000223031321-1011211332200213"></a>

## Direct properties — metadata / 131111223210 / 3

<a id="canonical-0233103121111200-2102202012102232-1232130022221021-3011020121123011-0021131301002021-3100000210113020-2323223031000121-3320112203111311"></a>

<a id="canonical-3302021232210300-3332030201021002-3021032223131232-0001010212221330-1333010311133332-0311123331021222-1301111222020313-1220301020212303"></a>

## description_spec property — metadata / 131111223210 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1210211300000123-1010312231021100-1002312113332021-0211031021311120-0310310012311330-1220100012201131-1200010330123212-0223033020112330"></a>

<a id="canonical-2132011221220312-3213132203113033-0331111310131033-2223333032300231-2230120321100130-0302123303100310-2312302132011000-3311330110021221"></a>

## name property — metadata / 131111223210 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3332220000110321-1332233303102111-0023300012211021-1130132301312121-0313010223300322-0131330312131101-2111032202020230-3202320201301331"></a>

## Next pages — metadata / 131111223210 / 6

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-3310331010213110-3101213300003222-3021120021321003-0333023333133021-2032303012222011-1332203203222230-0321133031203200-3003323103232221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132110133233300-1023311300232033-0123102001311000-1212203113202132-2301303021300002-2210210000120133-2103331122310100-0233323212020231"></a>

## rule_list.rules.no_http_connect_port — no_http_connect_port / 023130322011 / 2

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-001.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.no_http_connect_port

<a id="canonical-0111023013123131-3100003330320311-1212021131321300-1233200332221332-3212323113123333-1123130023200200-0130023123113223-2103000121330203"></a>

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
no_http_connect_port = {}
```

<a id="canonical-3220220231320230-0231231323333100-1313323301023010-2223110232232131-2210323212330211-2203222332032231-0132003022332023-2330220120120300"></a>

## Direct properties — no_http_connect_port / 023130322011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223002130311203-2112003132000222-2012122013101121-3131311122111303-2023230322020010-0021021132012312-1221002033230022-1311031023311310"></a>

## Next pages — no_http_connect_port / 023130322011 / 4

- [rule_list.rules](resources--forward_proxy_policy--reference--group-001.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)

<a id="canonical-1310001202000220-0222212030212031-3110301233213311-0331231201211302-1222330112030113-0332221131200032-0000000110301303-1000000231020321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
