---
page_title: "xcsh_cdn_cache_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule reference."
---

# xcsh_cdn_cache_rule reference

<a id="canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121131333330303-0001030000311100-0013001202321102-2020233003011330-3032201110210102-2322032001220122-2323122000321103-1103112100231233"></a>

## Property reference — Property reference / 301103033121 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- Property reference

<a id="canonical-0130033313222220-0221103013223230-1010113311110131-1331031300100321-2212132123300322-0312011110220303-1020101031311233-3220033313023312"></a>

## Direct properties — Property reference / 301103033121 / 3

<a id="canonical-2123323121111232-0330232013330011-3001330103001203-3102312210303103-2011033103323310-0033302230111002-2310003210210213-1023103030101303"></a>

<a id="canonical-2011313332033112-3111300100323231-3312200300101303-2011220300123322-1013222213031323-0313200222213022-0331331211013030-1011011320211303"></a>

## annotations property — Property reference / 301103033121 / 4

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

- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313): complete subsection reference.

<a id="canonical-1311100201132012-1111032321230213-1113110123012033-1021030133231021-0213220023320123-2001032232231123-3312000011021123-0210221020211221"></a>

<a id="canonical-2022000222331003-0201010302000133-1020100123221312-0302322031320003-0031020111301303-3221210312211111-2220333210010210-0011202133121331"></a>

## description property — Property reference / 301103033121 / 5

Type: `"string"`. Computed.

Description of the CDNCacheRule.

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

<a id="canonical-3113201020231100-3113300101310320-1011201002200121-0300221313131113-3220202101310201-1200021011233323-1003032211113202-2210111202133103"></a>

<a id="canonical-1001010121310223-1003131110030031-2130021102303033-0112222230021210-3100200233331010-2032333031103110-2101001103022112-0100230202230000"></a>

## ID property — Property reference / 301103033121 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3332321231301311-0003101020033230-1211201302131001-3133021031323231-0300313332322231-0002010102321312-0322103021103302-2101220220221001"></a>

<a id="canonical-2111132010332210-0333023022300023-1113130123331200-3132222013301112-0002120213321021-2021200110132113-0222303300321202-0222222103331302"></a>

## labels property — Property reference / 301103033121 / 7

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

<a id="canonical-2130222311233021-0030303231231111-1001000111200133-1013130031300223-0101122322212011-2232201332111100-1100210303001023-3111210203222233"></a>

<a id="canonical-1033030321231321-1131200100212233-1210203000200321-2003233013323310-0001011010133213-0131023112013211-3031221012030301-3321331231122311"></a>

## name property — Property reference / 301103033121 / 8

Type: `"string"`. Required.

Name of the CDNCacheRule.

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

<a id="canonical-0101023202103201-2110133300110100-3301303121323202-0221120222010131-2200220200200003-3112310013132201-3233103130300100-2303113122233131"></a>

<a id="canonical-2330330112213231-2032030031231232-2012323101122331-1013122120031012-0102133310130002-3301200123330000-3220302030310303-3322313003331330"></a>

## namespace property — Property reference / 301103033121 / 9

Type: `"string"`. Required.

Namespace where the CDNCacheRule exists.

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

<a id="canonical-1011113032121212-2200001331101011-0322023230022123-2103023012231321-0201131030223032-3332013202110202-2131010311300203-0010302020002311"></a>

## All schema paths — Property reference / 301103033121 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2123323121111232-0330232013330011-3001330103001203-3102312210303103-2011033103323310-0033302230111002-2310003210210213-1023103030101303) |
| `cache_rules` | [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2121103330320332-0013230221022233-0031213113020023-1232102202103101-3313300030303333-2313023012110200-2303223022110310-3320113302032032) |
| `cache_rules.cache_bypass` | [cache_rules.cache_bypass](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3102233010102201-3230031031302301-1132020023103213-2330203102123133-0313112223021112-0030032312312100-3033113130101020-0323022211300110) |
| `cache_rules.eligible_for_cache` | [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1220212012220310-2301003111321231-2233301210100020-2333313212112330-0311012000103111-3031311010012102-1000323123321301-3211133303331332) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3033230031020310-2110201103021010-3101310102300123-2020123012012102-1223003232230303-1201122133020312-3322133312321113-1201020123131323) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2233122121312233-1323311211032223-3033222023311110-3322202303310113-3233111321002223-3320330001112320-1001313112313330-2323321211003232) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0103003302011320-1303020301033012-0311302322321203-1201023131223131-2013000112223133-3002200102331020-2233033220332211-0000100001101003) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0021020322023333-3110033111211330-0301323302211303-2012323002130033-3321003101111103-2023331032011223-3220012310313003-3332120320112021) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1323331223230103-0132113132011113-1333313313301010-0021000100222331-3120203120221301-0020313313312300-2301130200001201-2103222022301331) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3331320031000001-2230002231300120-3331212233311131-2132020010320012-0200031200333333-3232201113031201-3111022302330311-0211210020002232) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0101211220122223-1332012021000113-2310233020213110-1310212133320230-3230010131113110-2020003333002201-3112230021022200-3323220313030303) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0333303202223030-3112100322133013-2112020022011030-0230101213322213-3313113301312110-1123213112031332-2030211330022010-2132211001030322) |
| `cache_rules.rule_expression_list` | [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0021011212320013-2320223101012132-3103321020303111-3100212013033113-2312130132011300-3222322202103002-3313330331123310-2311231032101032) |
| `cache_rules.rule_expression_list.cache_rule_expression` | [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1320331221222300-1231000130211023-3321231010203233-2132113322113003-3212121002001030-3313100321123222-1233133112022030-2212003312123202) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3031002030030302-0232033000221213-1312121012122330-0213220302212103-1333200202022030-3023112001111202-1013233002122332-0312000001211123) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3100203001100020-1121130321311000-2200303123120221-2023110121231303-1313201222033023-3300333112100130-2022122201031200-3333103210101322) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1203113113031222-0202132031130212-0002210021321330-1330001311131111-2320110202102211-0310121322002101-0032120102111212-0223103011131101) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1002012022303201-2222122331231322-2203331112303113-3101113110232301-3333333113233311-3003213013102033-1113331302331333-2322221110300112) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1313011120210220-1200223023333201-3223023010120312-3001110210000012-2010203020101232-0021331300300211-2131123302230312-0313323211100303) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1020000200022122-0321033020010031-0002131303021210-2312101300312021-0203111321001130-0313332130121010-2233333111202331-1100333300213311) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3000130301113112-0131130321233201-2003333113120132-3222331020300033-3130232203103230-1033313123330203-2123302321231103-3313002212333031) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1300212030223032-0301021013120213-1130301112301212-1220333122122010-0032213132321203-0123213322101033-2103323322330210-0311030330010210) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3333212313320231-2011333213102323-2221302100231211-1231130133300133-2331300122033221-3200232323303011-0010132231131021-1112010030033031) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1222113113101223-1311301223222213-2301033100231223-2311320300210213-2002013332003000-1203031131312222-3011000303230200-2133021211033223) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0303203013023111-1012013222133032-3020100221023301-3203302212121023-0111222033121320-0233132332230101-3323223330133210-2103222301300300) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3030231221122020-3200301310203123-3013222000002223-0300121233101112-0333330123200010-0021301233103102-1101222112212303-1311320302123221) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0023220200110212-0100022320122331-1232011221232101-1001030000323231-3213202133300302-0221022323300322-2301121021010300-1001320203013121) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0330201212221021-3210112322310323-0213122001023112-0203233300011223-0222301210300112-3231023011101220-3203020001112311-2203303200220221) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0333211333130320-1111222120330323-0003121302301021-1003111220313000-3132010231003100-1331333202011333-0323113120032223-3001230132003021) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1031210311130300-3301113212113323-2103202301312322-0223013013020110-0101231001211202-2332302203122303-1213102132302130-0023232302222130) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1021321003230202-0221131113323121-1201222203023021-0001221231031331-0103130212302022-3032121323333230-2120012123332121-1222321313321003) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0012112110123331-3331320122331112-3120003011131120-1212130322123121-1133000202322133-1322300311310202-0001001203203300-3113233311122013) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3121221223211031-2131213001212221-2200300202122120-3110132002331232-0211020220010110-2031001211200120-0131333021203210-1000331302332133) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1112103220111321-0323210001120301-1012330013020212-2001133122113332-0231022312000230-0320231000333222-2223231102213113-3301320100121210) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0112201101120320-1211331333121312-3202223233032303-2022313133100312-0100203230132231-2200232032031202-3100311331030230-1012123101233220) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2321303303101023-1023113211001012-3210210303220201-0130002322330302-0332310300123001-0331110323121102-3211323313310003-2012013011213112) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2301030302132133-2221133031321020-0331003230011000-2111111311022111-1031101312310121-2030300201233111-1323323202320101-2022300121120203) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0202102301321112-3301001230011332-1301200023010200-1210132030221303-2321223310213323-2213100033311123-1201031120213101-3201123322310032) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match` | [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0121020033010023-3123013212113110-3322231010100131-1323331231031320-2021233211003213-0110100020203012-3021310111131100-3031012133231233) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2121032203312023-0330222111003030-3221312123111232-1111311313100130-3301021132012321-0213120202123303-1332200220331100-1103103223223322) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0213311310303033-0323300110120102-0021221102323030-2001311333012133-0123010212212113-1222001033133313-3000022311330103-3133131100233221) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3023320011222110-2110210233331133-2320121130012322-2221212302111132-3320223202022321-2103332303330121-3223301000101323-3013121310310130) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2010322020033011-1331112222033231-1133011311230000-1103310330302203-2321033302332130-2131313230000130-3223122322210020-1112332211203010) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0102033103211012-0111012211000230-2002301200302202-0313312010213322-1013130022130302-3200211201113303-3323101221303010-0111023333122200) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0302123222101320-2012213031223002-0101201120323332-3222333112120121-0320220232110302-1301202023033213-3130020121212230-2321101321321000) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3202203010002220-0112212203012220-2310113220003323-1021033002302230-2001300021232220-2312301103313311-3012102301222230-2132331231331303) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2221030001223331-0303120132121223-1133312321111202-0302011212311103-0020220133201321-2133013030333230-3202033300330122-2233023030111201) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3320112321221131-3032032221013002-2210333223010102-3303322010010030-1233330012021331-2112300012122300-1232333332033111-2111303222221202) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1022201302223133-2322333221231302-0213310121322300-1220012301112221-1233232300312002-2133022110301113-1133223121012122-3031303232033202) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2132100232123120-3323020021213231-1011300132133113-2121030333012011-0333132101201111-1313031031100033-3011113233201020-2130131013203111) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0132120021310232-0212013221331001-3132113203210323-3331121202111323-3211332320303130-2020002130000133-3111101322333111-3223022131222312) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0131002320323113-3220033222333011-2201033020101311-0323122033212232-2011220112313113-3010000130333313-3223102132000230-1123203131033010) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3322231301032212-1020130311021200-2011231000320212-2011202031031320-3220223313102110-0101211310121110-3303223131212000-2120032221211112) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0203132232230230-0003203330310203-3300011013101231-3030133301120021-3102010032023203-2202120102110003-0300310131303220-2133130133101201) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2301210112131010-0200311131232121-3203303121232032-3020000021033032-1003011301330031-1201211100102120-2211130001310112-2102300000010012) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3030230301103330-3013123101010300-0200120302103021-0220122011011010-1030022001031112-0112200113021000-0213102020231121-1132223230033210) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1330330330023130-2113202232202021-2121302210231103-1100113112310231-0212130230332212-0311112302203113-3200213011021323-2213211012312030) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1131232313130002-1110020302012231-3010210111303112-2233003000230002-3003122030202222-1113120022031122-0222332301030133-3102221103120033) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2311013132000011-2311300302012333-1000123301122332-3032331223022023-3313311223200230-3032312210000120-1222121122011321-0322312313333233) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2122210033100122-0203132033003300-1230133113223212-0131303023310001-0031331122120331-2000111233113133-3103002332201302-2331131212301322) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0101200102233321-2132210033230200-2020232210303000-1210032223212022-2331213313003100-1231321002110221-1030202321102031-2021122103020122) |
| `cache_rules.rule_expression_list.expression_name` | [cache_rules.rule_expression_list.expression_name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3310203013222210-3022112202200112-2331323312201110-2003302013120330-1200221312303333-2030233201013010-2001023203012100-2032031203021320) |
| `cache_rules.rule_name` | [cache_rules.rule_name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1120110030212113-1031032221012231-2312220330003122-0022001230221011-1311301333302323-0220120033032022-3022013211220113-1032233300211311) |
| `description` | [description](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1311100201132012-1111032321230213-1113110123012033-1021030133231021-0213220023320123-2001032232231123-3312000011021123-0210221020211221) |
| `id` | [ID](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3113201020231100-3113300101310320-1011201002200121-0300221313131113-3220202101310201-1200021011233323-1003032211113202-2210111202133103) |
| `labels` | [labels](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3332321231301311-0003101020033230-1211201302131001-3133021031323231-0300313332322231-0002010102321312-0322103021103302-2101220220221001) |
| `name` | [name](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2130222311233021-0030303231231111-1001000111200133-1013130031300223-0101122322212011-2232201332111100-1100210303001023-3111210203222233) |
| `namespace` | [namespace](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0101023202103201-2110133300110100-3301303121323202-0221120222010131-2200220200200003-3112310013132201-3233103130300100-2303113122233131) |

<a id="canonical-1123330132122221-0312113222221313-2001201121102133-0332231220120002-1020003003200330-3210130202311211-2100203030020000-3313123323030002"></a>

## Next pages — Property reference / 301103033121 / 11

- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021333300102130-1033000202120003-2032101300310131-2220003122322133-2020002233111100-2001301123131011-1301332013201330-0223013301221201"></a>

## cache_rules — cache_rules / 101212310033 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- cache_rules

<a id="canonical-2121103330320332-0013230221022233-0031213113020023-1232102202103101-3313300030303333-2313023012110200-2303223022110310-3320113302032032"></a>

Type: `"single"`. Computed.

Cache Rule. This defines a CDN Cache Rule.

Upstream description:

This defines a CDN Cache Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_actions": "[\"cache_bypass\",\"eligible_for_cache\"]"
}
```

<a id="canonical-0233332201121032-0033102100330313-0132213121000323-1113232100330003-0002001213021102-0311311300003011-3231103021130102-2122223100321211"></a>

## Direct properties — cache_rules / 101212310033 / 3

- [cache_bypass](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1200310201203301-0321332111330311-0121302101031310-0212113311033201-1320132031033312-0321303202132213-1021003233022111-0200121322130210): complete subsection reference.

- [eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0121222213031113-0202111113032130-3010030102132010-3111032331233113-3013002230321320-1032202211011320-2000321331321001-1222332312203123): complete subsection reference.

- [rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300): complete subsection reference.

<a id="canonical-1120110030212113-1031032221012231-2312220330003122-0022001230221011-1311301333302323-0220120033032022-3022013211220113-1032233300211311"></a>

<a id="canonical-2032233311103310-1012133203132121-1012003023112332-0312310310200302-0032013021212100-0321331332000012-0102112010100300-1011130330321213"></a>

## rule_name property — cache_rules / 101212310033 / 4

Type: `"string"`. Computed.

Rule Name. Name of the Cache Rule.

Upstream description:

Name of the Cache Rule.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1110010302230111-2022302320230133-0013330201300333-3211232100120012-2132321313313121-3102031012333123-2101203123333103-2233101121200101"></a>

## Next pages — cache_rules / 101212310033 / 5

- [cache_rules.cache_bypass](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1200310201203301-0321332111330311-0121302101031310-0212113311033201-1320132031033312-0321303202132213-1021003233022111-0200121322130210)
- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0121222213031113-0202111113032130-3010030102132010-3111032331233113-3013002230321320-1032202211011320-2000321331321001-1222332312203123)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-1200310201203301-0321332111330311-0121302101031310-0212113311033201-1320132031033312-0321303202132213-1021003233022111-0200121322130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011213101103011-3331023002222002-0223002033133102-2131213020233002-3300001021132023-1301031302202130-1000013212012112-3130203021100201"></a>

## cache_rules.cache_bypass — cache_bypass / 310002203133 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- cache_rules.cache_bypass

<a id="canonical-3102233010102201-3230031031302301-1132020023103213-2330203102123133-0313112223021112-0030032312312100-3033113130101020-0323022211300110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for cache bypass.

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

<a id="canonical-1300131301102113-1233030210313003-0320223200231321-1112312132123210-0222121030020101-1120031002123331-1303230220203330-2313122300023313"></a>

## Direct properties — cache_bypass / 310002203133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221303003231133-3122310122211220-1103232301112313-2003302301233100-2030213021323212-3222311310100012-3203300103123000-1321112311031013"></a>

## Next pages — cache_bypass / 310002203133 / 4

- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-0121222213031113-0202111113032130-3010030102132010-3111032331233113-3013002230321320-1032202211011320-2000321331321001-1222332312203123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132202113013002-0021011112011030-0110222011311012-0311221021201101-0113131302333012-0000231223212101-2212333223312213-2113330330021020"></a>

## cache_rules.eligible_for_cache — eligible_for_cache / 222210121333 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- cache_rules.eligible_for_cache

<a id="canonical-1220212012220310-2301003111321231-2233301210100020-2333313212112330-0311012000103111-3031311010012102-1000323123321301-3211133303331332"></a>

Type: `"single"`. Computed.

Configuration parameter for eligible for cache.

Upstream description:

List of OPTIONS for Cache Action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-eligible_for_cache": "[\"scheme_proxy_host_request_uri\",\"scheme_proxy_host_uri\"]"
}
```

<a id="canonical-1332203302332233-3303121130312032-0222032112001322-3113010210330030-1123301232330100-0010331320311201-3313000211203301-0103001203100001"></a>

## Direct properties — eligible_for_cache / 222210121333 / 3

- [scheme_proxy_host_request_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1133121110333221-2102301131222222-1101333101231332-1000333213102100-0000231312133030-1300101132033322-1010301231001112-3331032133211302): complete subsection reference.

- [scheme_proxy_host_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2322323230330032-1013101130203322-2002311100202112-3230333221330001-3011103200320000-1332113333221302-2013001123000010-2210131133101313): complete subsection reference.

<a id="canonical-2020200033031211-3231213310210121-0323111023231302-0303033123302322-1310310111133310-2021231322220312-3233201320302312-3312131022230312"></a>

## Next pages — eligible_for_cache / 222210121333 / 4

- [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1133121110333221-2102301131222222-1101333101231332-1000333213102100-0000231312133030-1300101132033322-1010301231001112-3331032133211302)
- [cache_rules.eligible_for_cache.scheme_proxy_host_uri](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2322323230330032-1013101130203322-2002311100202112-3230333221330001-3011103200320000-1332113333221302-2013001123000010-2210131133101313)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-1133121110333221-2102301131222222-1101333101231332-1000333213102100-0000231312133030-1300101132033322-1010301231001112-3331032133211302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221230323332220-1331201323121221-0321020000033202-3032130022210030-0221313033102133-2101201030033013-1023201130100300-0212200222333221"></a>

## cache_rules.eligible_for_cache.scheme_proxy_host_request_uri — scheme_proxy_host_request_uri / 323132002120 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0121222213031113-0202111113032130-3010030102132010-3111032331233113-3013002230321320-1032202211011320-2000321331321001-1222332312203123)
- cache_rules.eligible_for_cache.scheme_proxy_host_request_uri

<a id="canonical-3033230031020310-2110201103021010-3101310102300123-2020123012012102-1223003232230303-1201122133020312-3322133312321113-1201020123131323"></a>

Type: `"single"`. Computed.

Cache TTL Enable Props. Cache TTL Enable Values.

Upstream description:

Cache TTL Enable Values.

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

<a id="canonical-0112130233013200-0120313232301232-2210210201002003-1021000230021303-2121023312113303-2020221321033110-1100202123032112-2121321031103321"></a>

## Direct properties — scheme_proxy_host_request_uri / 323132002120 / 3

<a id="canonical-2233122121312233-1323311211032223-3033222023311110-3322202303310113-3233111321002223-3320330001112320-1001313112313330-2323321211003232"></a>

<a id="canonical-2221230101321023-3103033002021323-0012023122131031-2103001311221002-0320113131021201-2303232001001021-0221020302030233-1131211223000010"></a>

## cache_override property — scheme_proxy_host_request_uri / 323132002120 / 4

Type: `"bool"`. Computed.

Cache Override. Honour Cache Override.

Upstream description:

Honour Cache Override.

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

<a id="canonical-0103003302011320-1303020301033012-0311302322321203-1201023131223131-2013000112223133-3002200102331020-2233033220332211-0000100001101003"></a>

<a id="canonical-1013330322032030-1303233031122013-1300033211132232-1223000233121021-2001222333303122-3101120211122333-0301021030132233-3022001210123021"></a>

## cache_ttl property — scheme_proxy_host_request_uri / 323132002120 / 5

Type: `"string"`. Computed.

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

Upstream description:

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-0021020322023333-3110033111211330-0301323302211303-2012323002130033-3321003101111103-2023331032011223-3220012310313003-3332120320112021"></a>

<a id="canonical-3322323033113312-0222103212023122-2211233103020130-1230303231213212-2102001122301103-2311103123011112-3031333111103311-2033203000330213"></a>

## ignore_response_cookie property — scheme_proxy_host_request_uri / 323132002120 / 6

Type: `"bool"`. Computed.

By default, response will not be cached if set-cookie header is present. This option will override
the behavior and cache response even with set-cookie header present.

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

<a id="canonical-1232220112303120-3310021301322301-2323021300100220-1202112222312111-1131133210001301-2322210103303012-1202202301223233-2023223303123303"></a>

## Next pages — scheme_proxy_host_request_uri / 323132002120 / 7

- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0121222213031113-0202111113032130-3010030102132010-3111032331233113-3013002230321320-1032202211011320-2000321331321001-1222332312203123)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-2322323230330032-1013101130203322-2002311100202112-3230333221330001-3011103200320000-1332113333221302-2013001123000010-2210131133101313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130110023131203-0112012121311210-0230222313013003-2313011123311312-1031210223110101-0231300321030111-0132320130030323-1200213111121002"></a>

## cache_rules.eligible_for_cache.scheme_proxy_host_uri — scheme_proxy_host_uri / 001021330013 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0121222213031113-0202111113032130-3010030102132010-3111032331233113-3013002230321320-1032202211011320-2000321331321001-1222332312203123)
- cache_rules.eligible_for_cache.scheme_proxy_host_uri

<a id="canonical-1323331223230103-0132113132011113-1333313313301010-0021000100222331-3120203120221301-0020313313312300-2301130200001201-2103222022301331"></a>

Type: `"single"`. Computed.

Cache TTL Enable Props. Cache TTL Enable Values.

Upstream description:

Cache TTL Enable Values.

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

<a id="canonical-3332001130220100-2200102213330131-0320132012123001-1103122103202200-1320323120301213-2102111333121130-3330223121133330-0323012120120211"></a>

## Direct properties — scheme_proxy_host_uri / 001021330013 / 3

<a id="canonical-3331320031000001-2230002231300120-3331212233311131-2132020010320012-0200031200333333-3232201113031201-3111022302330311-0211210020002232"></a>

<a id="canonical-1210230111220323-0330030131100031-3323133013121021-1031222120033221-0030033112211313-1133222011320011-2122302331113113-3312000013120310"></a>

## cache_override property — scheme_proxy_host_uri / 001021330013 / 4

Type: `"bool"`. Computed.

Cache Override. Honour Cache Override.

Upstream description:

Honour Cache Override.

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

<a id="canonical-0101211220122223-1332012021000113-2310233020213110-1310212133320230-3230010131113110-2020003333002201-3112230021022200-3323220313030303"></a>

<a id="canonical-3311223232311230-2110002101133032-3010003313211202-2202132033203203-2333322330331012-0012323313211102-3111300011330001-1311231122310031"></a>

## cache_ttl property — scheme_proxy_host_uri / 001021330013 / 5

Type: `"string"`. Computed.

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

Upstream description:

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-0333303202223030-3112100322133013-2112020022011030-0230101213322213-3313113301312110-1123213112031332-2030211330022010-2132211001030322"></a>

<a id="canonical-1111333300132033-2200100321233212-2110100013020200-1133333013332330-1301210021002122-1013212022301200-0232220101033222-3232012010022112"></a>

## ignore_response_cookie property — scheme_proxy_host_uri / 001021330013 / 6

Type: `"bool"`. Computed.

By default, response will not be cached if set-cookie header is present. This option will override
the behavior and cache response even with set-cookie header present.

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

<a id="canonical-1110112112220030-0130002312313121-0023110223101021-1220110030101000-1332212232303032-0131031323223003-0132233220312120-1313230201103323"></a>

## Next pages — scheme_proxy_host_uri / 001021330013 / 7

- [cache_rules.eligible_for_cache](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0121222213031113-0202111113032130-3010030102132010-3111032331233113-3013002230321320-1032202211011320-2000321331321001-1222332312203123)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330221131200332-1023212110330122-1130223201033123-0331121220310210-0231102031103110-1200220033113300-2311300220230322-0033321213123203"></a>

## cache_rules.rule_expression_list — rule_expression_list / 032113323002 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- cache_rules.rule_expression_list

<a id="canonical-0021011212320013-2320223101012132-3103321020303111-3100212013033113-2312130132011300-3222322202103002-3313330331123310-2311231032101032"></a>

Type: `"list"`. Computed.

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs.

Upstream description:

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs..

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0220113112011030-0113023033020133-2121222313201220-1200333223231101-3002203003202200-1031021211311113-0330033000033233-2110130213220133"></a>

## Direct properties — rule_expression_list / 032113323002 / 3

- [cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320): complete subsection reference.

<a id="canonical-3310203013222210-3022112202200112-2331323312201110-2003302013120330-1200221312303333-2030233201013010-2001023203012100-2032031203021320"></a>

<a id="canonical-1023123331211300-3013330310330231-2232011001303101-2032213103023133-2203200012210122-0103332002202123-3302133033122321-3002120100000130"></a>

## expression_name property — rule_expression_list / 032113323002 / 4

Type: `"string"`. Computed.

Name of the Expressions items that are ANDed.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0102220032220102-1023331113303230-3220333121302002-2232001310122031-1031103020302322-3021322333221002-3210020320322132-0030321022221023"></a>

## Next pages — rule_expression_list / 032113323002 / 5

- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221101211122231-2322310300233332-0101103331233231-3202033322321123-1301303002012031-3012323203331311-2202000233022021-2011232232323300"></a>

## cache_rules.rule_expression_list.cache_rule_expression — cache_rule_expression / 103322311330 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- cache_rules.rule_expression_list.cache_rule_expression

<a id="canonical-1320331221222300-1231000130211023-3321231010203233-2132113322113003-3212121002001030-3313100321123222-1233133112022030-2212003312123202"></a>

Type: `"list"`. Computed.

The Cache Rule Expression Terms that are ANDed.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1120101123031322-3131233112303320-3032320113320132-3101121221101110-1113322100110020-0203032101222223-3021213001230213-1212311213213101"></a>

## Direct properties — cache_rule_expression / 103322311330 / 3

- [cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2310102321323033-1210212323323332-1202211203122230-2021133113112202-1003003121033123-3112112230201032-3011101011230231-0011311220230303): complete subsection reference.

- [cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1231220023220121-1002031303233033-0112323332132323-3002103230130102-3122330223000101-0300100210313312-1201132233330102-0313212323110010): complete subsection reference.

- [path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1222132303203102-1332002321001313-3101000221323333-3220013110303000-3013213221013111-3010322021132232-1322221213022000-1103300211302001): complete subsection reference.

- [query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3230302103122023-3030031120230010-3101223033311330-3033211100103332-2020123332320031-2222303302313230-3003010112200303-1012113223321112): complete subsection reference.

<a id="canonical-2003322103320011-3220100130103121-3231121032002131-2012313223103310-2011203001031301-3311213002333131-0211311221011102-3230001331310000"></a>

## Next pages — cache_rule_expression / 103322311330 / 4

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2310102321323033-1210212323323332-1202211203122230-2021133113112202-1003003121033123-3112112230201032-3011101011230231-0011311220230303)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1231220023220121-1002031303233033-0112323332132323-3002103230130102-3122330223000101-0300100210313312-1201132233330102-0313212323110010)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1222132303203102-1332002321001313-3101000221323333-3220013110303000-3013213221013111-3010322021132232-1322221213022000-1103300211302001)
- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3230302103122023-3030031120230010-3101223033311330-3033211100103332-2020123332320031-2222303302313230-3003010112200303-1012113223321112)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-2310102321323033-1210212323323332-1202211203122230-2021133113112202-1003003121033123-3112112230201032-3011101011230231-0011311220230303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201021033130323-3230033033003001-2133032130201122-2210212201003033-1023321210300110-3330310301201320-1232302032213001-0021310213103113"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cache_headers — cache_headers / 222302122122 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers

<a id="canonical-3031002030030302-0232033000221213-1312121012122330-0213220302212103-1333200202022030-3023112001111202-1013233002122332-0312000001211123"></a>

Type: `"list"`. Computed.

Configure cache rule headers to match the criteria.

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

<a id="canonical-0123011210102020-2201320122222110-1023302031223322-0131201002301101-3003003333221232-3111203330202223-2312320212003111-0210012010100220"></a>

## Direct properties — cache_headers / 222302122122 / 3

<a id="canonical-3100203001100020-1121130321311000-2200303123120221-2023110121231303-1313201222033023-3300333112100130-2022122201031200-3333103210101322"></a>

<a id="canonical-2131112221011211-3133331323000020-3311101200323312-3101123002221110-2200130303202301-3221321322132100-3333131323031030-2021201212131302"></a>

## name property — cache_headers / 222302122122 / 4

Type: `"string"`. Computed.

\[Enum: PROXY\_HOST|REFERER|SCHEME|USER\_AGENT\] - PROXY\_HOST: Proxy Host Name of the proxied
server - REFERER: Referer This is the address of the previous web page from which a link to the
currently requested page was followed - SCHEME: Scheme The HTTP scheme used: HTTP or HTTPS -
USER\_AGENT: User Agent The user agent string of the user agent. Possible values are
\`PROXY\_HOST\`, \`REFERER\`, \`SCHEME\`, \`USER\_AGENT\`. Defaults to \`PROXY\_HOST\`.

Upstream description:

&#8203;- PROXY\_HOST: Proxy Host

Name of the proxied server &#8203;- REFERER: Referer

This is the address of the previous web page from which a link to the currently requested page was
followed &#8203;- SCHEME: Scheme

The HTTP scheme used: HTTP or HTTPS &#8203;- USER\_AGENT: User Agent

The user agent string of the user agent.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROXY_HOST",
  "enum": [
    "PROXY_HOST",
    "REFERER",
    "SCHEME",
    "USER_AGENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0200213333313200-3101323231120230-3012100101231011-2310221130201002-1203001120030000-2232111301103211-1221113121212121-1210120201133322): complete subsection reference.

<a id="canonical-1300300011013320-2220222002211100-3122213010320113-1021000130300330-2330300230223110-2110213323233101-2230300203020323-0200323200310200"></a>

## Next pages — cache_headers / 222302122122 / 5

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0200213333313200-3101323231120230-3012100101231011-2310221130201002-1203001120030000-2232111301103211-1221113121212121-1210120201133322)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-0200213333313200-3101323231120230-3012100101231011-2310221130201002-1203001120030000-2232111301103211-1221113121212121-1210120201133322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323012132003203-0223320233032222-0030030032013230-2102023202000220-3022032312233300-3210233323212132-2101101120001033-2200230001021023"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator — operator / 332210023131 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2310102321323033-1210212323323332-1202211203122230-2021133113112202-1003003121033123-3112112230201032-3011101011230231-0011311220230303)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator

<a id="canonical-1203113113031222-0202132031130212-0002210021321330-1330001311131111-2320110202102211-0310121322002101-0032120102111212-0223103011131101"></a>

Type: `"single"`. Computed.

Operator

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

<a id="canonical-3303300231023220-2102210020021101-3000033021202333-1233230123201232-0020110233030012-0200313230231131-2322310232111101-1303000331003001"></a>

## Direct properties — operator / 332210023131 / 3

<a id="canonical-1002012022303201-2222122331231322-2203331112303113-3101113110232301-3333333113233311-3003213013102033-1113331302331333-2322221110300112"></a>

<a id="canonical-1331202021121032-2302120133111132-2200202301032012-0011322203032212-1122233013030032-3212031331110133-0032131013001011-0013202201312320"></a>

## contains property — operator / 332210023131 / 4

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The header value must include the specified value as a substring.

<a id="canonical-1313011120210220-1200223023333201-3223023010120312-3001110210000012-2010203020101232-0021331300300211-2131123302230312-0313323211100303"></a>

<a id="canonical-2033101233131113-0301111132133233-0301222022310030-3320113332102312-3231203111201313-1031333322321221-3333120013101233-1012130333311221"></a>

## does_not_contain property — operator / 332210023131 / 5

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not include the specified value as a substring.

<a id="canonical-1020000200022122-0321033020010031-0002131303021210-2312101300312021-0203111321001130-0313332130121010-2233333111202331-1100333300213311"></a>

<a id="canonical-3212031003133331-1211230113111200-3331131302303130-0333111321201023-2132033302120302-3113233032302121-0123213311232221-0022310311303100"></a>

## does_not_end_with property — operator / 332210023131 / 6

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not end with the specified value.

<a id="canonical-3000130301113112-0131130321233201-2003333113120132-3222331020300033-3130232203103230-1033313123330203-2123302321231103-3313002212333031"></a>

<a id="canonical-2302030001103103-1113221033303311-0033003132013200-2132201300112223-1302233313332322-1000330112223331-0131233310331222-1230212300323120"></a>

## does_not_equal property — operator / 332210023131 / 7

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not match the specified value.

<a id="canonical-1300212030223032-0301021013120213-1130301112301212-1220333122122010-0032213132321203-0123213322101033-2103323322330210-0311030330010210"></a>

<a id="canonical-0112330213031103-3132110020121220-3013120203203121-1012023330322013-0002203221131202-1130300000223310-1320323203012311-2030211110001311"></a>

## does_not_start_with property — operator / 332210023131 / 8

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The header value must not begin with the specified value.

<a id="canonical-3333212313320231-2011333213102323-2221302100231211-1231130133300133-2331300122033221-3200232323303011-0010132231131021-1112010030033031"></a>

<a id="canonical-2303323033321013-1013311300033220-2322321102330033-2221322111103231-2202013003131011-3230001310012122-0123203111213312-3100132001320011"></a>

## endswith property — operator / 332210023131 / 9

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The header value must end with the specified value.

<a id="canonical-1222113113101223-1311301223222213-2301033100231223-2311320300210213-2002013332003000-1203031131312222-3011000303230200-2133021211033223"></a>

<a id="canonical-3212201201320221-3220303310001202-0321002131213132-0300330310312313-3120021300001113-2221022200121013-2111122332121331-2332212311132330"></a>

## equals property — operator / 332210023131 / 10

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The header value must exactly match the specified value.

<a id="canonical-0303203013023111-1012013222133032-3020100221023301-3203302212121023-0111222033121320-0233132332230101-3323223330133210-2103222301300300"></a>

<a id="canonical-3312131213230331-2100130322120322-1100211210110013-3132123122313010-0210200100023011-0331313003232103-1212130032003011-2321102011211012"></a>

## match_regex property — operator / 332210023131 / 11

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The header value must match the specified regular expression pattern.

<a id="canonical-3030231221122020-3200301310203123-3013222000002223-0300121233101112-0333330123200010-0021301233103102-1101222112212303-1311320302123221"></a>

<a id="canonical-3122123030033110-3323012300313221-1200000212113333-1112013231010321-0112313323132003-2222023323012311-3303121333311321-1132220010311330"></a>

## startswith property — operator / 332210023131 / 12

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The header value must begin with the specified value.

<a id="canonical-2213002212032003-3020301303131033-0011212301203202-0130322022332333-3133321300323200-0113102010313021-2013121030303202-1222122023311131"></a>

## Next pages — operator / 332210023131 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2310102321323033-1210212323323332-1202211203122230-2021133113112202-1003003121033123-3112112230201032-3011101011230231-0011311220230303)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-1231220023220121-1002031303233033-0112323332132323-3002103230130102-3122330223000101-0300100210313312-1201132233330102-0313212323110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020011330211100-1032201313111031-3301033311023231-0320221321112132-1130131210303001-3230300022002332-3100013031020322-3123011103210013"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher — cookie_matcher / 302210230223 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher

<a id="canonical-0023220200110212-0100022320122331-1232011221232101-1001030000323231-3213202133300302-0221022323300322-2301121021010300-1001320203013121"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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

<a id="canonical-0233002013010232-2220013323033023-0221303030000221-1002330023102131-0321311110001023-0312010233100212-1220110331010211-0220311122003002"></a>

## Direct properties — cookie_matcher / 302210230223 / 3

<a id="canonical-0330201212221021-3210112322310323-0213122001023112-0203233300011223-0222301210300112-3231023011101220-3203020001112311-2203303200220221"></a>

<a id="canonical-1001311333122301-0030113230010000-3210232313020312-3101320213200012-1033110310023001-1200323310331123-3233020113100321-3311310313323121"></a>

## name property — cookie_matcher / 302210230223 / 4

Type: `"string"`. Computed.

Cookie Name. Enter the name of the cookie to match.

Upstream description:

Enter the name of the cookie to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1011103131033320-1110203132010000-2102212012302033-1231013330220023-2002203220021103-1232122203313330-2210103010112303-3323003231310000): complete subsection reference.

<a id="canonical-3231311310220031-1131011222323113-1221110002303332-3122203311231222-2231101232103120-1332003121203211-0032223112031030-3021210020011101"></a>

## Next pages — cookie_matcher / 302210230223 / 5

- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1011103131033320-1110203132010000-2102212012302033-1231013330220023-2002203220021103-1232122203313330-2210103010112303-3323003231310000)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-1011103131033320-1110203132010000-2102212012302033-1231013330220023-2002203220021103-1232122203313330-2210103010112303-3323003231310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212331333002211-0300300200301203-0311202100230310-0322000031321030-0202323302111021-1320221022212302-3131301022300000-1110233323012313"></a>

## cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator — operator / 002132323311 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1231220023220121-1002031303233033-0112323332132323-3002103230130102-3122330223000101-0300100210313312-1201132233330102-0313212323110010)
- cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator

<a id="canonical-0333211333130320-1111222120330323-0003121302301021-1003111220313000-3132010231003100-1331333202011333-0323113120032223-3001230132003021"></a>

Type: `"single"`. Computed.

Operator

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

<a id="canonical-0231101033332000-0323012010310213-2123033323101232-0023003301111002-3230301233102011-1033321022002222-1113230202222103-1000100310230103"></a>

## Direct properties — operator / 002132323311 / 3

<a id="canonical-1031210311130300-3301113212113323-2103202301312322-0223013013020110-0101231001211202-2332302203122303-1213102132302130-0023232302222130"></a>

<a id="canonical-1001113313333132-0221031301102210-0220020031023220-2232030022233120-1322033132133220-3110310210332232-1220311131323221-0113010002221010"></a>

## contains property — operator / 002132323311 / 4

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The cookie value must include the specified value as a substring.

<a id="canonical-1021321003230202-0221131113323121-1201222203023021-0001221231031331-0103130212302022-3032121323333230-2120012123332121-1222321313321003"></a>

<a id="canonical-2013311322003212-0013331001212003-0113133031322001-3221113200032333-2130220011231011-0110131013221013-2122101013103200-2030330003033331"></a>

## does_not_contain property — operator / 002132323311 / 5

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not include the specified value as a substring.

<a id="canonical-0012112110123331-3331320122331112-3120003011131120-1212130322123121-1133000202322133-1322300311310202-0001001203203300-3113233311122013"></a>

<a id="canonical-0010120220023130-1311302330130231-2131130130111321-2123211113330132-1020321020330202-3012320302231031-1120233313100111-1121002131312223"></a>

## does_not_end_with property — operator / 002132323311 / 6

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not end with the specified value.

<a id="canonical-3121221223211031-2131213001212221-2200300202122120-3110132002331232-0211020220010110-2031001211200120-0131333021203210-1000331302332133"></a>

<a id="canonical-0123123313320221-2103100030210200-3113121012232012-0213232020313300-0313200333221012-0102233331312311-0231302103122101-3030122221311002"></a>

## does_not_equal property — operator / 002132323311 / 7

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not match the specified value.

<a id="canonical-1112103220111321-0323210001120301-1012330013020212-2001133122113332-0231022312000230-0320231000333222-2223231102213113-3301320100121210"></a>

<a id="canonical-1102131333100333-0033220303321222-3212332100323322-0201213323212110-1331111021311122-3032023310320033-1302330211133220-0002220222112333"></a>

## does_not_start_with property — operator / 002132323311 / 8

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The cookie value must not begin with the specified value.

<a id="canonical-0112201101120320-1211331333121312-3202223233032303-2022313133100312-0100203230132231-2200232032031202-3100311331030230-1012123101233220"></a>

<a id="canonical-2023023101311233-2223332122233323-0011322111223021-2332132023010030-3103121003001103-2322221103020211-1212110333030212-3123321120212200"></a>

## endswith property — operator / 002132323311 / 9

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The cookie value must end with the specified value.

<a id="canonical-2321303303101023-1023113211001012-3210210303220201-0130002322330302-0332310300123001-0331110323121102-3211323313310003-2012013011213112"></a>

<a id="canonical-3301321103302333-3333310021212323-1033231232231012-3220033011033021-2210112122123133-2301000330111200-1013231233013331-0300021320100133"></a>

## equals property — operator / 002132323311 / 10

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The cookie value must exactly match the specified value.

<a id="canonical-2301030302132133-2221133031321020-0331003230011000-2111111311022111-1031101312310121-2030300201233111-1323323202320101-2022300121120203"></a>

<a id="canonical-1230132231220001-3033001031022201-0111020221122031-1333210020301001-0311113122110313-2321023213131320-0310020220032030-1212223001231013"></a>

## match_regex property — operator / 002132323311 / 11

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The cookie value must match the specified regular expression pattern in PCRE
format.

<a id="canonical-0202102301321112-3301001230011332-1301200023010200-1210132030221303-2321223310213323-2213100033311123-1201031120213101-3201123322310032"></a>

<a id="canonical-1332000033023022-3221000222211231-3302021202021123-1002303123022232-3112330022012311-0013031300122002-2310103122111320-3330232223202302"></a>

## startswith property — operator / 002132323311 / 12

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The cookie value must begin with the specified value.

<a id="canonical-3210032312131232-2101000030302303-3001221011011122-0211223232132130-0112012323211202-3322321331230313-0113011312300200-0030013032231010"></a>

## Next pages — operator / 002132323311 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1231220023220121-1002031303233033-0112323332132323-3002103230130102-3122330223000101-0300100210313312-1201132233330102-0313212323110010)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-1222132303203102-1332002321001313-3101000221323333-3220013110303000-3013213221013111-3010322021132232-1322221213022000-1103300211302001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013310111221323-0220032031021311-3001222121112010-1211033032210230-0133002321202123-1322021100330013-2131203330303302-3133103031232233"></a>

## cache_rules.rule_expression_list.cache_rule_expression.path_match — path_match / 312331111233 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- cache_rules.rule_expression_list.cache_rule_expression.path_match

<a id="canonical-0121020033010023-3123013212113110-3322231010100131-1323331231031320-2021233211003213-0110100020203012-3021310111131100-3031012133231233"></a>

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
  }
}
```

<a id="canonical-3212101332321101-3311003332322221-2213222233330200-0010302033301023-2030101121301331-1023231110200203-1130100210002201-1010201332322131"></a>

## Direct properties — path_match / 312331111233 / 3

- [operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1010230310032130-2010303023131230-0130113213231301-0231333131032033-2230300113001323-0123300113030322-3123010001230231-3332320100131020): complete subsection reference.

<a id="canonical-3331032010031020-1002012221322011-0031000323023300-3102231210222313-2120313003133120-0100111112123123-2120313200232022-3132202002122033"></a>

## Next pages — path_match / 312331111233 / 4

- [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1010230310032130-2010303023131230-0130113213231301-0231333131032033-2230300113001323-0123300113030322-3123010001230231-3332320100131020)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-1010230310032130-2010303023131230-0130113213231301-0231333131032033-2230300113001323-0123300113030322-3123010001230231-3332320100131020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230122111120302-2313211120302300-1231031102203102-1200312132302022-2101013220131212-2320322012110221-1120332213111333-3201212111031221"></a>

## cache_rules.rule_expression_list.cache_rule_expression.path_match.operator — operator / 322011110113 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1222132303203102-1332002321001313-3101000221323333-3220013110303000-3013213221013111-3010322021132232-1322221213022000-1103300211302001)
- cache_rules.rule_expression_list.cache_rule_expression.path_match.operator

<a id="canonical-2121032203312023-0330222111003030-3221312123111232-1111311313100130-3301021132012321-0213120202123303-1332200220331100-1103103223223322"></a>

Type: `"single"`. Computed.

Operator

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

<a id="canonical-0311031311323101-1111203300331331-1033113322132101-0022033012011020-3101200310033123-2213131131020022-0120301002223111-3301221323332313"></a>

## Direct properties — operator / 322011110113 / 3

<a id="canonical-0213311310303033-0323300110120102-0021221102323030-2001311333012133-0123010212212113-1222001033133313-3000022311330103-3133131100233221"></a>

<a id="canonical-2033023311003023-0110003131211200-1303130001300301-3310122120121031-3210311000022222-1021013000310220-3110333222032331-1111210122321301"></a>

## contains property — operator / 322011110113 / 4

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The path must include the specified value as a substring, up to the
filename.

<a id="canonical-3023320011222110-2110210233331133-2320121130012322-2221212302111132-3320223202022321-2103332303330121-3223301000101323-3013121310310130"></a>

<a id="canonical-1330221132032000-2100103130101322-1020220001203010-1110213311213001-2002131323100223-0221021001223101-2313101023202211-1032302020001020"></a>

## does_not_contain property — operator / 322011110113 / 5

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not include the specified value as a substring, up to the filename.

<a id="canonical-2010322020033011-1331112222033231-1133011311230000-1103310330302203-2321033302332130-2131313230000130-3223122322210020-1112332211203010"></a>

<a id="canonical-0330211010100103-2131332230011220-1002133303100330-0013130101233321-0011011222132330-1300123321100123-0302311101021231-3330133220113110"></a>

## does_not_end_with property — operator / 322011110113 / 6

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not end with the specified value, up to the filename.

<a id="canonical-0102033103211012-0111012211000230-2002301200302202-0313312010213322-1013130022130302-3200211201113303-3323101221303010-0111023333122200"></a>

<a id="canonical-2030001123031121-0101203130121133-1321023302212223-3100132203333312-2303033233331211-0311011013203312-2321112100301010-2200122301231113"></a>

## does_not_equal property — operator / 322011110113 / 7

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not match the specified value, up to the filename.

<a id="canonical-0302123222101320-2012213031223002-0101201120323332-3222333112120121-0320220232110302-1301202023033213-3130020121212230-2321101321321000"></a>

<a id="canonical-1202100112113011-2212331221300211-1323210303120310-0031133220100130-3121312013102231-1020213002121110-0313203221033130-3132110230133111"></a>

## does_not_start_with property — operator / 322011110113 / 8

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The path must not begin with the specified value, up to the filename.

<a id="canonical-3202203010002220-0112212203012220-2310113220003323-1021033002302230-2001300021232220-2312301103313311-3012102301222230-2132331231331303"></a>

<a id="canonical-2203203030020201-1100100120031312-3130110213120132-0220221223211013-1200000030301313-3221023112231233-3030013020232133-1013220012123333"></a>

## endswith property — operator / 322011110113 / 9

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The path must end with the specified value, up to the filename.

<a id="canonical-2221030001223331-0303120132121223-1133312321111202-0302011212311103-0020220133201321-2133013030333230-3202033300330122-2233023030111201"></a>

<a id="canonical-0010303013313010-2202230022123213-0023312333331302-3101103223300322-2111110310322200-3311121230231030-1310133022321310-2220032233020222"></a>

## equals property — operator / 322011110113 / 10

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The path must exactly match the specified value, up to the filename.

<a id="canonical-3320112321221131-3032032221013002-2210333223010102-3303322010010030-1233330012021331-2112300012122300-1232333332033111-2111303222221202"></a>

<a id="canonical-0031110131032013-2203021322303031-1320331332313012-1103000012222212-1311211012113132-0002222123233130-1310322000231203-2310110220213220"></a>

## match_regex property — operator / 322011110113 / 11

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The path must match the specified regular expression pattern in PCRE format.

<a id="canonical-1022201302223133-2322333221231302-0213310121322300-1220012301112221-1233232300312002-2133022110301113-1133223121012122-3031303232033202"></a>

<a id="canonical-3011230001321303-3012312130133032-1103200312323120-2301212303002300-2031203311132003-1002021301303311-3313320102313330-2303031220231102"></a>

## startswith property — operator / 322011110113 / 12

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The path must begin with the specified value, up to the filename.

<a id="canonical-1231013133201011-2031202231023122-0333333210111012-1003301300220013-0033102031013230-1230011003223002-3222332222110322-0331000320331332"></a>

## Next pages — operator / 322011110113 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.path_match](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1222132303203102-1332002321001313-3101000221323333-3220013110303000-3013213221013111-3010322021132232-1322221213022000-1103300211302001)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-3230302103122023-3030031120230010-3101223033311330-3033211100103332-2020123332320031-2222303302313230-3003010112200303-1012113223321112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223112323201323-3001000311233112-0201332330122313-1220112201110331-2320103002100323-2330002120230220-0202313132312130-1121000102213110"></a>

## cache_rules.rule_expression_list.cache_rule_expression.query_parameters — query_parameters / 230203003000 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters

<a id="canonical-2132100232123120-3323020021213231-1011300132133113-2121030333012011-0333132101201111-1313031031100033-3011113233201020-2130131013203111"></a>

Type: `"list"`. Computed.

Query Parameters. List of (key, value) query parameters.

Upstream description:

List of (key, value) query parameters.

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

<a id="canonical-1131131021010002-1320210301333220-2312301333113312-0001011203303301-0230113311033030-0301233003012113-1310312013031222-3011012333203310"></a>

## Direct properties — query_parameters / 230203003000 / 3

<a id="canonical-0132120021310232-0212013221331001-3132113203210323-3331121202111323-3211332320303130-2020002130000133-3111101322333111-3223022131222312"></a>

<a id="canonical-3022123332321021-1310303323120333-0101212133013121-1322331103221033-2312102023132300-0303333020223131-0233111110023320-2311222313212001"></a>

## key property — query_parameters / 230203003000 / 4

Type: `"string"`. Computed.

The name of the query parameter to match.

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

- [operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0321133023101223-3333223330220120-0031130023011103-1210101333233333-3013013311003310-0300000230211002-3213201331211300-1130102020320113): complete subsection reference.

<a id="canonical-0111213110012101-0312212120231103-1031012230001013-3331322123033031-0112311003233003-2301013033330322-0310103233100223-2220100110203211"></a>

## Next pages — query_parameters / 230203003000 / 5

- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](data-sources--cdn_cache_rule--reference--group-001.md#canonical-0321133023101223-3333223330220120-0031130023011103-1210101333233333-3013013311003310-0300000230211002-3213201331211300-1130102020320113)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)

<a id="canonical-0321133023101223-3333223330220120-0031130023011103-1210101333233333-3013013311003310-0300000230211002-3213201331211300-1130102020320113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320021312133102-1331003200013201-3132133110020210-0221132111123331-2233122013331113-3002022010110011-2311130213211013-0202201330220022"></a>

## cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator — operator / 102122221131 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
- [Property reference](data-sources--cdn_cache_rule--reference--group-001.md#canonical-2320313102000230-1230033101112212-1013201230012001-2233112321110123-2313202103031130-1022000000100313-3212211012111121-2302303321033121)
- [cache_rules](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313)
- [cache_rules.rule_expression_list](data-sources--cdn_cache_rule--reference--group-001.md#canonical-1123331013102310-2233232031003023-2332200331323011-2322120001321310-0201033310311231-1132311301211121-3330122102201003-1122321330032300)
- [cache_rules.rule_expression_list.cache_rule_expression](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3213102220223202-2032111323011021-1113120113131332-1131011120323320-0232001103202221-2310132300300222-2200213030022123-3003322013201320)
- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3230302103122023-3030031120230010-3101223033311330-3033211100103332-2020123332320031-2222303302313230-3003010112200303-1012113223321112)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator

<a id="canonical-0131002320323113-3220033222333011-2201033020101311-0323122033212232-2011220112313113-3010000130333313-3223102132000230-1123203131033010"></a>

Type: `"single"`. Computed.

Operator

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

<a id="canonical-0323231331203232-1321100300232033-1303010011232201-3212131330121201-0100223232202313-2210112223300332-3303321320223021-2131231223233032"></a>

## Direct properties — operator / 102122221131 / 3

<a id="canonical-3322231301032212-1020130311021200-2011231000320212-2011202031031320-3220223313102110-0101211310121110-3303223131212000-2120032221211112"></a>

<a id="canonical-0210120122002032-0200210102320000-3102231211320133-2012302110103333-2133123012130320-1113311303303122-3103213331122331-1211130030332000"></a>

## contains property — operator / 102122221131 / 4

Type: `"string"`. Computed.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The query parameter value must include the specified value as a substring.

<a id="canonical-0203132232230230-0003203330310203-3300011013101231-3030133301120021-3102010032023203-2202120102110003-0300310131303220-2133130133101201"></a>

<a id="canonical-2213222332222201-2133210030202111-3333333022130131-1123220322331312-0310202031001023-3031323000121310-3310102000032022-2002132013012010"></a>

## does_not_contain property — operator / 102122221131 / 5

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not include the specified value as a substring.

<a id="canonical-2301210112131010-0200311131232121-3203303121232032-3020000021033032-1003011301330031-1201211100102120-2211130001310112-2102300000010012"></a>

<a id="canonical-3121220101102320-1003302231131101-2100010222131313-1311130030031022-1110322003033003-0103122322020013-1330332331333221-3232322302331110"></a>

## does_not_end_with property — operator / 102122221131 / 6

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not end with the specified value.

<a id="canonical-3030230301103330-3013123101010300-0200120302103021-0220122011011010-1030022001031112-0112200113021000-0213102020231121-1132223230033210"></a>

<a id="canonical-1203331131333111-1320330203132032-2323302101233202-1003030121002303-2333000301221320-3323111031220201-3010210230332013-3021232133203022"></a>

## does_not_equal property — operator / 102122221131 / 7

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not match the specified value.

<a id="canonical-1330330330023130-2113202232202021-2121302210231103-1100113112310231-0212130230332212-0311112302203113-3200213011021323-2213211012312030"></a>

<a id="canonical-1200012110031011-2302313101231211-0012302002330302-0232323021011012-2032210202122222-2200301033010131-2313303332133200-1133323001223203"></a>

## does_not_start_with property — operator / 102122221131 / 8

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The query parameter value must not begin with the specified value.

<a id="canonical-1131232313130002-1110020302012231-3010210111303112-2233003000230002-3003122030202222-1113120022031122-0222332301030133-3102221103120033"></a>

<a id="canonical-3122132231031300-3013210301112020-3021311203113203-1121222303331012-0123003033202101-3301033131302321-1000033220230223-3131201202320032"></a>

## endswith property — operator / 102122221131 / 9

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The query parameter value must end with the specified value.

<a id="canonical-2311013132000011-2311300302012333-1000123301122332-3032331223022023-3313311223200230-3032312210000120-1222121122011321-0322312313333233"></a>

<a id="canonical-2301001201203020-0320113023021301-0033112133001232-1212100010120330-2303213203201003-2320120101223220-3010103332030321-2303031122122232"></a>

## equals property — operator / 102122221131 / 10

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The query parameter value must exactly match the specified value.

<a id="canonical-2122210033100122-0203132033003300-1230133113223212-0131303023310001-0031331122120331-2000111233113133-3103002332201302-2331131212301322"></a>

<a id="canonical-3011323121232130-3133133223233320-1120130131020320-3023201310103000-0010021320113032-0100033330103023-1003311202101311-1330303110223232"></a>

## match_regex property — operator / 102122221131 / 11

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The query parameter value must match the specified regular expression pattern in
PCRE format.

<a id="canonical-0101200102233321-2132210033230200-2020232210303000-1210032223212022-2331213313003100-1231321002110221-1030202321102031-2021122103020122"></a>

<a id="canonical-2032332110210033-3011221010221123-3120313013110103-3122200301302232-1232311213020200-1301023211123232-2033221122201302-1302132330222332"></a>

## startswith property — operator / 102122221131 / 12

Type: `"string"`. Computed.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The query parameter value must begin with the specified value.

<a id="canonical-3213133133130321-0113002322322012-1321320330230202-1102013302211132-2230332222322123-1101133200302031-3233333001130303-0003331033101002"></a>

## Next pages — operator / 102122221131 / 13

- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](data-sources--cdn_cache_rule--reference--group-001.md#canonical-3230302103122023-3030031120230010-3101223033311330-3033211100103332-2020123332320031-2222303302313230-3003010112200303-1012113223321112)
- [xcsh_cdn_cache_rule](../data-sources/cdn_cache_rule.md#canonical-1302132311222223-2013120100100131-2322023003310212-2230310003033201-2101302310132220-3120031003232130-1130213222020202-2012321300112021)
