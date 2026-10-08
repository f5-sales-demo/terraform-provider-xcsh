---
page_title: "xcsh_cdn_cache_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule reference."
---

# xcsh_cdn_cache_rule reference

<a id="canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- Property reference

<a id="canonical-2201023202011230-1323310022220200-1300131122002033-3131333110122332-1032211021012200-0301303230010000-2130032320020200-3212010102233221"></a>

### Direct properties for `xcsh_cdn_cache_rule`

<a id="canonical-0111122133213213-3021312210103203-0322023021022223-0100120002321031-2333220332132122-3132133121330201-0000210101201001-1301023031111220"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032): complete subsection reference.

<a id="canonical-0131303130023030-3312012020222020-0111000200302312-3301010210222123-1023303012113331-1332221230031133-1302212211330010-1001331300321213"></a>

<a id="canonical-3222032332232203-2212232133132322-1001012111012311-3012231331223012-1033312120123130-0032320332330133-0102321123110010-1031032230112220"></a>

#### `description` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1102100222302221-1003131020300132-1231201101301131-1303321221020330-3202322122023023-2111012320302003-3321223223211232-2001103030133001"></a>

<a id="canonical-3331123322032012-2010200013331300-2220111230311223-2232033123313213-0010100021032222-1031102233122320-1111113112120031-3131100201330311"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

<a id="canonical-3331311123222011-2113032232133333-1010211202103323-1000010333311003-1320203132133033-1113113231210223-2323220311033200-1233102230010232"></a>

<a id="canonical-0030020032013313-1123120333301132-1231033020330110-3130220220103330-3322013220332033-0032113031003021-3220103013213030-0021202322311313"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2302122000332223-2222303313203023-3200220122123030-2102021122331032-0320130310123210-3121031111330322-3121312210010020-0003023331330310"></a>

<a id="canonical-0012022003232022-2100332023300303-2232021213200013-2231101310032011-1122111233022033-0332013323002212-2221120123230122-1233213100331113"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-1222121031313220-0231003101031011-0303231303011012-1123113010030000-2111021300203210-0202101202033020-3322133030222220-2202120301120311"></a>

<a id="canonical-3132101021220311-2120311233020132-1321212233010012-0301121303311333-3211310002203210-3110122222011331-3220201001322310-3012112210333202"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CDN Cache Rule. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1203120010000232-3100013023201132-1202223122131303-0111200013013000-3121013202230002-1211211302101300-1002211230121012-3031023111002233"></a>

<a id="canonical-1002033030301032-0331232230023021-2111213220233221-2310210110113312-0032233211011022-0223021200011110-0322010332303023-1320002030113112"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CDN Cache Rule is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [timeouts](resources--cdn_cache_rule--reference--group-001.md#canonical-1212033023211232-3031321101321003-2333222011200010-1230320100011230-0012013221003131-0331013111020303-1122320301310222-1203323113231230): complete subsection reference.

<a id="canonical-0301212300010102-0323131222310003-3201330303211310-2222110232300213-1332120333323232-3011332022102101-3001222231101321-2120111210102322"></a>

### All schema paths for `xcsh_cdn_cache_rule`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cdn_cache_rule--reference--group-001.md#canonical-0111122133213213-3021312210103203-0322023021022223-0100120002321031-2333220332132122-3132133121330201-0000210101201001-1301023031111220) |
| `cache_rules` | [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-1122333001102122-2321313023312201-2003012001230200-3203123031022233-0021232222212032-1311303112133020-2130102100203230-0003120301002331) |
| `cache_rules.cache_bypass` | [cache_rules.cache_bypass](resources--cdn_cache_rule--reference--group-001.md#canonical-0001321121022201-0201202313201011-0323303110311211-0231032230111003-3201330103133131-1312222210131201-2021301221110113-2103000102132332) |
| `cache_rules.eligible_for_cache` | [cache_rules.eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-3002130322212122-1200301323123012-0120010112323233-0132322212032022-1000321320200203-2331333102200301-0100101320133320-2320312201301232) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-3120113302013211-2213223031120302-3123322331330330-3233130223000310-1303230032130321-0011210211311300-2020303113231311-3021010233122133) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override](resources--cdn_cache_rule--reference--group-001.md#canonical-0220332030103300-3302102113302201-0111002001322222-1130320201311103-2131231033311000-1010000112113231-2201202323112333-1011213030011301) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl](resources--cdn_cache_rule--reference--group-001.md#canonical-0222301211122310-3003030220010121-3030232211110210-1112203122303230-3313312000303223-0203011232332133-1313112001200123-0320310131322121) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie](resources--cdn_cache_rule--reference--group-001.md#canonical-1333021222133303-1212332121012211-2301331002023131-0220123121312121-1020000033000223-2323130121230221-2021330312110303-1221123301311031) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-2000130033222132-3300122111031022-1220102330131002-3200003201303311-1123201102133332-2323032021331233-1021221320212220-0032230311013322) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override](resources--cdn_cache_rule--reference--group-001.md#canonical-3300323333112131-0331213103102322-1122122021021320-3010210133032330-0313031031220022-2210233201322220-2031212022103001-1233213033101231) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl](resources--cdn_cache_rule--reference--group-001.md#canonical-0201001220203111-1221021002030302-3022101111331100-1003200210300301-0201101111212122-2103121323103221-1202311232123230-1231230202101322) |
| `cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie` | [cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie](resources--cdn_cache_rule--reference--group-001.md#canonical-2001320103003002-3303230023300032-1203203320312210-3203033230301003-2330320323022022-2302032321020322-3323203002131022-3202101223111301) |
| `cache_rules.rule_expression_list` | [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-2011100211321213-0122103333303202-0100112210123213-0211020230100210-0000332133111003-3300102323103232-2201303300130101-2300300113120203) |
| `cache_rules.rule_expression_list.cache_rule_expression` | [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-2121110202321320-0311013113001213-1333012231133323-1331123030111011-1103313320330303-1202111230221213-3112023030311232-1233111230203331) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](resources--cdn_cache_rule--reference--group-001.md#canonical-3032323113232030-0230220031133000-3130123200030322-2202122120130111-0223200222100002-2332213313033111-0102112112223300-3313303101223221) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name](resources--cdn_cache_rule--reference--group-001.md#canonical-0002032133133102-2203003213211131-1300310223013123-1201330231113311-2300033223321213-2031013023200220-0232121021223231-3130323113000303) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-3100110202133212-1001323332131203-3100133301130210-2230111232122213-2331221133313133-2011211002130331-0223203330301301-2121110111211133) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains](resources--cdn_cache_rule--reference--group-001.md#canonical-3301202210301300-3301203122103133-0301323010123201-2302321333312200-3022103323333030-2033023212000020-1122032011031202-2012023011103102) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain](resources--cdn_cache_rule--reference--group-001.md#canonical-3330030202302130-0230300012313302-0031320111022032-3032120023031203-0130121302031112-1102323120302333-2231013010002123-3221113113212212) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with](resources--cdn_cache_rule--reference--group-001.md#canonical-2103123302232103-0123101231012130-0013103220222122-0200202022113322-3201133303113030-2230323312202100-1220211003200012-3313021210320321) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal](resources--cdn_cache_rule--reference--group-001.md#canonical-0122013221331030-2022331322221032-0010233223111233-0003110130201032-1202133121212030-3322313301011211-0101231213021202-0332200032130122) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with](resources--cdn_cache_rule--reference--group-001.md#canonical-3210002003303000-1311000003222312-1100123312131330-2220200213131311-1131022212011323-1011300021302320-2120120313103132-2002121100000200) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith](resources--cdn_cache_rule--reference--group-001.md#canonical-2012102133203210-2113021002102102-1220101232333012-2112313223302002-2120220301223023-0032332103012300-2323210211101111-1123133200212301) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals](resources--cdn_cache_rule--reference--group-001.md#canonical-3201231301323123-2111203001320011-3300213300003313-1101122201311230-1002221102130022-0322311113133101-2112203101302120-3333103110032132) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex](resources--cdn_cache_rule--reference--group-001.md#canonical-2202321000113311-1120130032110220-0122012120100200-3120030022111122-0212133331002223-3210130300213003-0021032100333312-3200101300300032) |
| `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith](resources--cdn_cache_rule--reference--group-001.md#canonical-0033011312033132-2231310003320332-3031223013111230-3003002110322032-2223313112222011-3320011011321330-1010323333011212-0322200003032232) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](resources--cdn_cache_rule--reference--group-001.md#canonical-3121002331002223-2221133032312221-2313120022332111-2102113221331301-3200110121103021-1010301110211321-2323202321012321-2322003311303202) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name](resources--cdn_cache_rule--reference--group-001.md#canonical-2030000212010232-2030112310201103-2322323131130212-2021133333322222-3010000302012230-0221013003103113-0020231302001210-2202011310001012) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-2213303211010032-0203332322103123-3202300002112002-2320023231123222-2121030233021102-0223003110310332-2302232032303331-3110331323011330) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains](resources--cdn_cache_rule--reference--group-001.md#canonical-2300010302231202-2013230322112322-3332303312101320-1232010221001221-1222100323011123-0322033130222321-1003310303312111-0221003023333221) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain](resources--cdn_cache_rule--reference--group-001.md#canonical-2121230321002301-3311211203211231-0021322210212223-3012131331023222-2302123220011003-1221003223231132-3202211212312333-0321301223133021) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with](resources--cdn_cache_rule--reference--group-001.md#canonical-0320110003010021-3110330200123110-3013211001312231-2330100231233010-3000022310213303-3130200231300013-2012311331223301-1201122110200333) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal](resources--cdn_cache_rule--reference--group-001.md#canonical-2000132120011101-0112033023103311-3333101111110301-1121210203323231-2230211010001100-0211103030313222-1032331231031022-1001031110012112) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with](resources--cdn_cache_rule--reference--group-001.md#canonical-3001213213112132-3231213131033001-2230220133022300-0212113100201322-2102323211201200-1321230122130302-0113111310121210-1331212023111211) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith](resources--cdn_cache_rule--reference--group-001.md#canonical-2113102132311133-0023100131021113-1300020231020331-2031103121023011-1132202310101221-0200103231100333-0302320000033231-1122013331111022) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals](resources--cdn_cache_rule--reference--group-001.md#canonical-0101100023121031-3100211101002302-2023233221112131-1002033110301031-1222032032232230-2202033121031010-1301312003013321-1030300210023230) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex](resources--cdn_cache_rule--reference--group-001.md#canonical-1201231213121301-0131130121120233-0332100210212020-1330313031133301-2002012330201112-1231133013010020-3003310103320302-2231212220113122) |
| `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith](resources--cdn_cache_rule--reference--group-001.md#canonical-3032332003211223-2132021233022222-3112120111120212-1211112031003023-0002013210000322-3032031323012023-2011212110310302-1222022022321333) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match` | [cache_rules.rule_expression_list.cache_rule_expression.path_match](resources--cdn_cache_rule--reference--group-001.md#canonical-2230021001103331-3022203323321121-2230012211001111-0232013200202331-2111103201031311-3002121100321122-3321301012223322-1213010331022000) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-2113002000330232-3310333211013110-2131322012001311-0103101332123312-1130211200300121-3220330213121130-2020212001330032-3002200203002221) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains](resources--cdn_cache_rule--reference--group-001.md#canonical-0213200012211013-1100322303201133-0102123132123202-3311211110100322-0303133010303012-1131122323203011-0320032231213020-3211011301203123) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain](resources--cdn_cache_rule--reference--group-001.md#canonical-3320022321220123-3013103213310130-2320320230212011-2311103103212331-1011320333011120-2112210320230323-1113121202120121-3132301103232303) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with](resources--cdn_cache_rule--reference--group-001.md#canonical-1131312101133211-0303013211130210-2223223012220300-1013201103032033-2321030213230030-0212120302320013-2020111212332220-2330200032213311) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal](resources--cdn_cache_rule--reference--group-001.md#canonical-0110133002331020-2033111101330333-0011123311230200-2011231330003113-3012011233122230-3202311030313222-2030212103110233-1333002202102132) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with](resources--cdn_cache_rule--reference--group-001.md#canonical-2313000330001203-3111210310213130-0332122321102120-0310120313231112-3310201312012113-2113321001023320-3311103110312120-2330210203113013) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith](resources--cdn_cache_rule--reference--group-001.md#canonical-3022101233122003-0023021000223021-3122031203220201-1321331030002231-0133312112012210-1211123221130012-0012213323103110-1032211303120132) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals](resources--cdn_cache_rule--reference--group-001.md#canonical-1122011310203203-1203202320103101-0302123103103223-3320310333222131-0222011222123313-0213320321323100-1203233010100232-0030301301221203) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex](resources--cdn_cache_rule--reference--group-001.md#canonical-3233112023321303-0121323330032223-2131223120020123-2031021103133002-2233033132233012-1321103010132002-2222313031120010-1130022331310013) |
| `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith](resources--cdn_cache_rule--reference--group-001.md#canonical-1133032102223222-2010122211101331-3123022200221100-1000221231302311-1131133233023213-2301120113022203-0010330123230303-0012212030002131) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](resources--cdn_cache_rule--reference--group-001.md#canonical-2311131300301233-2300103022032201-0122230231001312-3023000101130333-3011210201331112-0200120021113132-1303200123101220-3122203223102332) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key](resources--cdn_cache_rule--reference--group-001.md#canonical-2211330131032301-1133311302213321-3303203313201032-2001300230020113-3011003331003222-0110131011103312-3203221122002333-1221231110301331) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator](resources--cdn_cache_rule--reference--group-001.md#canonical-1311031211032010-3301220321300212-3123211200203102-1322223133312011-0211212222010333-1210312122121210-0313100221221221-0033203322013220) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains](resources--cdn_cache_rule--reference--group-001.md#canonical-1332202222133300-0033312013210122-2303111002203111-1303201211031130-3210122203302120-0100311212232113-2102120130201322-3133100332323323) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain](resources--cdn_cache_rule--reference--group-001.md#canonical-3011322332303023-3202201223130100-0030230330221202-0220121033202221-1001201213022003-3200301313332012-0223112132311120-1333111002212001) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with](resources--cdn_cache_rule--reference--group-001.md#canonical-0011001131123022-3303133101023102-0000033000022332-3101100033323302-2232012130122012-2131112303220200-2120102221230032-3231320323120031) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal](resources--cdn_cache_rule--reference--group-001.md#canonical-1013011232213233-0000230210030223-2223210002011301-2333023130001231-0103202033023022-0321220130101012-0113230010013101-2000103211210201) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with](resources--cdn_cache_rule--reference--group-001.md#canonical-2303233120122013-0302211131013232-0313300231331013-3333210012022030-0200300022030001-3212030211202133-0110112201011221-0022313123222311) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith](resources--cdn_cache_rule--reference--group-001.md#canonical-0311330313333020-0223002103103231-1012103010222310-0221223120222023-3221301102033013-3210131011010023-3131233121033201-2210021112003032) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals](resources--cdn_cache_rule--reference--group-001.md#canonical-3010132322321122-2212201201211010-3220022330322200-0223232122111003-3122200231222210-3113000031212201-3200330120303023-0003312300013021) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex](resources--cdn_cache_rule--reference--group-001.md#canonical-3311132011111121-3012302310323112-1121103201223232-3221123003331300-1133132232123001-1323223320322333-0330022032003302-0111013333302200) |
| `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith` | [cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith](resources--cdn_cache_rule--reference--group-001.md#canonical-2310331232223302-0113322303333210-0012130010031333-1102131301302322-3301113132102110-2331032200002000-3123001022120002-2031320102200022) |
| `cache_rules.rule_expression_list.expression_name` | [cache_rules.rule_expression_list.expression_name](resources--cdn_cache_rule--reference--group-001.md#canonical-1313000023203122-3013033222000132-1033323033321100-3110111112002200-0332002112302213-3022330212202230-1302001113133320-2202112232021231) |
| `cache_rules.rule_name` | [cache_rules.rule_name](resources--cdn_cache_rule--reference--group-001.md#canonical-3020100322333112-1201032221321222-3123033331300333-3133103223220011-3210322201313333-0130333200130222-0211300010302303-1030213132202213) |
| `description` | [description](resources--cdn_cache_rule--reference--group-001.md#canonical-0131303130023030-3312012020222020-0111000200302312-3301010210222123-1023303012113331-1332221230031133-1302212211330010-1001331300321213) |
| `disable` | [disable](resources--cdn_cache_rule--reference--group-001.md#canonical-1102100222302221-1003131020300132-1231201101301131-1303321221020330-3202322122023023-2111012320302003-3321223223211232-2001103030133001) |
| `id` | [ID](resources--cdn_cache_rule--reference--group-001.md#canonical-3331311123222011-2113032232133333-1010211202103323-1000010333311003-1320203132133033-1113113231210223-2323220311033200-1233102230010232) |
| `labels` | [labels](resources--cdn_cache_rule--reference--group-001.md#canonical-2302122000332223-2222303313203023-3200220122123030-2102021122331032-0320130310123210-3121031111330322-3121312210010020-0003023331330310) |
| `name` | [name](resources--cdn_cache_rule--reference--group-001.md#canonical-1222121031313220-0231003101031011-0303231303011012-1123113010030000-2111021300203210-0202101202033020-3322133030222220-2202120301120311) |
| `namespace` | [namespace](resources--cdn_cache_rule--reference--group-001.md#canonical-1203120010000232-3100013023201132-1202223122131303-0111200013013000-3121013202230002-1211211302101300-1002211230121012-3031023111002233) |
| `timeouts` | [timeouts](resources--cdn_cache_rule--reference--group-001.md#canonical-1003323221203111-2313300330023312-0013302333120001-2101012220223321-2313212100330322-2031321001013112-3020201022030230-1112113321000321) |
| `timeouts.create` | [timeouts.create](resources--cdn_cache_rule--reference--group-001.md#canonical-2210111332222310-0122012230032321-0202233121110332-3231202311331213-0120130013202230-0120323131321313-3232113202231320-1103021220331031) |
| `timeouts.delete` | [timeouts.delete](resources--cdn_cache_rule--reference--group-001.md#canonical-3212201313203022-0030033231023211-3120331011322112-3001322200310130-0201213220033222-3210032230320300-0012302230103131-3111021133001330) |
| `timeouts.read` | [timeouts.read](resources--cdn_cache_rule--reference--group-001.md#canonical-0132011103112100-2123213010010232-1113223322331200-2301101302232012-1122120133032222-1310030131011211-0213222103211000-1113133310212212) |
| `timeouts.update` | [timeouts.update](resources--cdn_cache_rule--reference--group-001.md#canonical-3101103011202003-2100201013223010-2021203232310131-2131010112230230-1222111131030013-0233123210220131-3002312032122021-3022302303200231) |

<a id="canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- cache_rules

<a id="canonical-1122333001102122-2321313023312201-2003012001230200-3203123031022233-0021232222212032-1311303112133020-2130102100203230-0003120301002331"></a>

Type: `"object"`. single nested block, Optional.

Cache Rule. This defines a CDN Cache Rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rule_expression_list",
    "rule_name"),
  validators.ConflictingObjectAttributes("cache_bypass",
    "eligible_for_cache")}
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
  "x-ves-oneof-field-cache_actions": "[\"cache_bypass\",\"eligible_for_cache\"]"
}
```

Terraform syntax:

```terraform
cache_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110313213211031-1021311022031211-0111023100333201-2210003101213332-1313111103222311-1203231333102120-0202013123213310-3010233013011012"></a>

### Direct properties for `cache_rules`

- [cache_bypass](resources--cdn_cache_rule--reference--group-001.md#canonical-2122013130232133-0111010230200202-2202321310010011-0203200021301121-0332132330000033-3332030221133201-0302003102311012-0110101110123212): complete subsection reference.

- [eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-3223232310102321-3003333113231003-0113233100223033-1002000121033202-2010331330331313-3230301030101011-2002313003222331-0320232110002100): complete subsection reference.

- [rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232): complete subsection reference.

<a id="canonical-3020100322333112-1201032221321222-3123033331300333-3133103223220011-3210322201313333-0130333200130222-0211300010302303-1030213132202213"></a>

<a id="canonical-3030000121120332-0130332030320233-2002100033131102-0313132003213010-2010030123112203-3032210121000123-0103233311033220-1312000203120123"></a>

#### `cache_rules.rule_name` property

Type: `"string"`. Optional.

Rule Name. Name of the Cache Rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2122013130232133-0111010230200202-2202321310010011-0203200021301121-0332132330000033-3332030221133201-0302003102311012-0110101110123212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.cache_bypass` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- cache_rules.cache_bypass

<a id="canonical-0001321121022201-0201202313201011-0323303110311211-0231032230111003-3201330103133131-1312222210131201-2021301221110113-2103000102132332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for cache bypass.

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

Terraform syntax:

```terraform
cache_bypass = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223232310102321-3003333113231003-0113233100223033-1002000121033202-2010331330331313-3230301030101011-2002313003222331-0320232110002100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.eligible_for_cache` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- cache_rules.eligible_for_cache

<a id="canonical-3002130322212122-1200301323123012-0120010112323233-0132322212032022-1000321320200203-2331333102200301-0100101320133320-2320312201301232"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eligible for cache.

Additional upstream details:

List of OPTIONS for Cache Action.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("scheme_proxy_host_request_uri",
    "scheme_proxy_host_uri")}
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
  "x-ves-oneof-field-eligible_for_cache": "[\"scheme_proxy_host_request_uri\",\"scheme_proxy_host_uri\"]"
}
```

Terraform syntax:

```terraform
eligible_for_cache {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212332002122202-2111233132121332-0202001131301222-3302313031311201-3120002122323301-1201100131313032-2200330003302222-3232001220321300"></a>

### Direct properties for `cache_rules.eligible_for_cache`

- [scheme_proxy_host_request_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-1323230121223200-0212212131033013-1021231333100002-1011020320033300-3131202303031112-3012110231231332-3222323310233031-2312123121103012): complete subsection reference.

- [scheme_proxy_host_uri](resources--cdn_cache_rule--reference--group-001.md#canonical-3131020013020013-3111002330212120-3331323132023223-2111313013102020-2031102010222030-2121100131011321-3100111011102223-0101003333131212): complete subsection reference.

<a id="canonical-1323230121223200-0212212131033013-1021231333100002-1011020320033300-3131202303031112-3012110231231332-3222323310233031-2312123121103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-3223232310102321-3003333113231003-0113233100223033-1002000121033202-2010331330331313-3230301030101011-2002313003222331-0320232110002100)
- cache_rules.eligible_for_cache.scheme_proxy_host_request_uri

<a id="canonical-3120113302013211-2213223031120302-3123322331330330-3233130223000310-1303230032130321-0011210211311300-2020303113231311-3021010233122133"></a>

Type: `"object"`. single nested block, Optional.

Cache TTL Enable Props. Cache TTL Enable Values.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cache_ttl")}
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
scheme_proxy_host_request_uri {
  # Configure direct properties listed below.
}
```

<a id="canonical-2331201033022020-0100320102033333-2231203102012220-2312002000002201-2002120233303201-1111223002311222-1213200003313112-1112010122230010"></a>

### Direct properties for `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri`

<a id="canonical-0220332030103300-3302102113302201-0111002001322222-1130320201311103-2131231033311000-1010000112113231-2201202323112333-1011213030011301"></a>

#### `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_override` property

Type: `"bool"`. Optional.

Cache Override. Honour Cache Override.

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

<a id="canonical-0222301211122310-3003030220010121-3030232211110210-1112203122303230-3313312000303223-0203011232332133-1313112001200123-0320310131322121"></a>

<a id="canonical-0120313320132022-1033203223111002-2120133223332332-2323132133320020-3001123211112003-1333303323023330-2131311112120210-0033313010002101"></a>

#### `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.cache_ttl` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1333021222133303-1212332121012211-2301331002023131-0220123121312121-1020000033000223-2323130121230221-2021330312110303-1221123301311031"></a>

<a id="canonical-1121300300223303-1132311013321031-3033333012330333-2110223130231221-1330101213013310-2113031331123321-1111013112200210-3001203103113230"></a>

#### `cache_rules.eligible_for_cache.scheme_proxy_host_request_uri.ignore_response_cookie` property

Type: `"bool"`. Optional.

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

<a id="canonical-3131020013020013-3111002330212120-3331323132023223-2111313013102020-2031102010222030-2121100131011321-3100111011102223-0101003333131212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.eligible_for_cache.scheme_proxy_host_uri` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.eligible_for_cache](resources--cdn_cache_rule--reference--group-001.md#canonical-3223232310102321-3003333113231003-0113233100223033-1002000121033202-2010331330331313-3230301030101011-2002313003222331-0320232110002100)
- cache_rules.eligible_for_cache.scheme_proxy_host_uri

<a id="canonical-2000130033222132-3300122111031022-1220102330131002-3200003201303311-1123201102133332-2323032021331233-1021221320212220-0032230311013322"></a>

Type: `"object"`. single nested block, Optional.

Cache TTL Enable Props. Cache TTL Enable Values.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cache_ttl")}
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
scheme_proxy_host_uri {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202322031133103-3331213000030021-3013100333101023-0121221330123120-3210031122200023-0102302303200121-0101131101331312-1313223010210211"></a>

### Direct properties for `cache_rules.eligible_for_cache.scheme_proxy_host_uri`

<a id="canonical-3300323333112131-0331213103102322-1122122021021320-3010210133032330-0313031031220022-2210233201322220-2031212022103001-1233213033101231"></a>

#### `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_override` property

Type: `"bool"`. Optional.

Cache Override. Honour Cache Override.

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

<a id="canonical-0201001220203111-1221021002030302-3022101111331100-1003200210300301-0201101111212122-2103121323103221-1202311232123230-1231230202101322"></a>

<a id="canonical-1223130223313313-3213020123001133-0120223033330102-2031203321231101-1031101002220110-2232131310222203-1302322330212330-0131230333301000"></a>

#### `cache_rules.eligible_for_cache.scheme_proxy_host_uri.cache_ttl` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2001320103003002-3303230023300032-1203203320312210-3203033230301003-2330320323022022-2302032321020322-3323203002131022-3202101223111301"></a>

<a id="canonical-2222002133331302-2231302311003201-0013312200131021-2021322330023111-2203032130301113-1301011010110230-2023331021010101-2021012121110301"></a>

#### `cache_rules.eligible_for_cache.scheme_proxy_host_uri.ignore_response_cookie` property

Type: `"bool"`. Optional.

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

<a id="canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- cache_rules.rule_expression_list

<a id="canonical-2011100211321213-0122103333303202-0100112210123213-0211020230100210-0000332133111003-3300102323103232-2201303300130101-2300300113120203"></a>

Type: `"object"`. list nested block, Optional.

Expressions are evaluated in the order in which they are specified. The evaluation stops when the
first rule match occurs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("cache_rule_expression",
    "expression_name")}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
rule_expression_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011103002333330-2001302211213100-2110100120132131-0001101030212013-0320000221320000-3022103322010210-1300221231031202-1121111301202012"></a>

### Direct properties for `cache_rules.rule_expression_list`

- [cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211): complete subsection reference.

<a id="canonical-1313000023203122-3013033222000132-1033323033321100-3110111112002200-0332002112302213-3022330212202230-1302001113133320-2202112232021231"></a>

<a id="canonical-1232221102300011-3033030323030110-2230201111100302-3230210333303330-1130230310031202-3013301001300320-2123330211313022-0003000010013032"></a>

#### `cache_rules.rule_expression_list.expression_name` property

Type: `"string"`. Optional.

Name of the Expressions items that are ANDed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list.cache_rule_expression` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232)
- cache_rules.rule_expression_list.cache_rule_expression

<a id="canonical-2121110202321320-0311013113001213-1333012231133323-1331123030111011-1103313320330303-1202111230221213-3112023030311232-1233111230203331"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
cache_rule_expression {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311012333201221-2123011023022112-0232220203301020-0311232223300312-1100130212200112-0211000022201223-2222123030210002-1330012302111322"></a>

### Direct properties for `cache_rules.rule_expression_list.cache_rule_expression`

- [cache_headers](resources--cdn_cache_rule--reference--group-001.md#canonical-0021201223133123-1200013121232330-0100330002102131-1323022222230023-3333012021332330-0333122303203033-2231113331132230-3201132323213220): complete subsection reference.

- [cookie_matcher](resources--cdn_cache_rule--reference--group-001.md#canonical-3331222303100022-2032300023213202-1132013320331232-2010223021302210-1102122233013032-3012330131230311-1122330220010203-1210211121131131): complete subsection reference.

- [path_match](resources--cdn_cache_rule--reference--group-001.md#canonical-1002100310130020-0310003113131311-2013020101322033-0120132133100322-0332310332020133-1111332232310300-0123303002223030-3330201313200213): complete subsection reference.

- [query_parameters](resources--cdn_cache_rule--reference--group-001.md#canonical-1132133310133103-1022330022230200-3201210301002123-0120022102320312-2100032230203203-1130330300300203-2022323102030133-0322103212202310): complete subsection reference.

<a id="canonical-0021201223133123-1200013121232330-0100330002102131-1323022222230023-3333012021332330-0333122303203033-2231113331132230-3201132323213220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list.cache_rule_expression.cache_headers` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers

<a id="canonical-3032323113232030-0230220031133000-3130123200030322-2202122120130111-0223200222100002-2332213313033111-0102112112223300-3313303101223221"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
cache_headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011002033201012-0233000321132303-0221302332102020-1211023002212311-2203011210222012-1033213130220303-0113101131231033-2332203011120032"></a>

### Direct properties for `cache_rules.rule_expression_list.cache_rule_expression.cache_headers`

<a id="canonical-0002032133133102-2203003213211131-1300310223013123-1201330231113311-2300033223321213-2031013023200220-0232121021223231-3130323113000303"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.name` property

Type: `"string"`. Optional.

\[Enum: PROXY\_HOST|REFERER|SCHEME|USER\_AGENT\] - PROXY\_HOST: Proxy hostname of the proxied
server - REFERER: Referer This is the address of the previous web page from which a link to the
currently requested page was followed - SCHEME: Scheme The HTTP scheme used: HTTP or HTTPS -
USER\_AGENT: User Agent The user agent string of the user agent. Possible values are
\`PROXY\_HOST\`, \`REFERER\`, \`SCHEME\`, \`USER\_AGENT\`. Defaults to \`PROXY\_HOST\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PROXY_HOST","REFERER","SCHEME","USER_AGENT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("PROXY_HOST",
    "REFERER",
    "SCHEME",
    "USER_AGENT"),
}
```

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

- [operator](resources--cdn_cache_rule--reference--group-001.md#canonical-1120031023302312-2211112032322312-2330311132102132-2232030231300231-0311110120003202-1201103113120222-2022011033212132-3133033010100013): complete subsection reference.

<a id="canonical-1120031023302312-2211112032322312-2330311132102132-2232030231300231-0311110120003202-1201103113120222-2022011033212132-3133033010100013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211)
- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers](resources--cdn_cache_rule--reference--group-001.md#canonical-0021201223133123-1200013121232330-0100330002102131-1323022222230023-3333012021332330-0333122303203033-2231113331132230-3201132323213220)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator

<a id="canonical-3100110202133212-1001323332131203-3100133301130210-2230111232122213-2331221133313133-2011211002130331-0223203330301301-2121110111211133"></a>

Type: `"object"`. single nested block, Optional.

Operator

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("contains",
    "does_not_contain"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("contains",
    "endswith"),
  validators.ConflictingObjectAttributes("contains",
    "equals"),
  validators.ConflictingObjectAttributes("contains",
    "match_regex"),
  validators.ConflictingObjectAttributes("contains",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "startswith"),
  validators.ConflictingObjectAttributes("endswith",
    "equals"),
  validators.ConflictingObjectAttributes("endswith",
    "match_regex"),
  validators.ConflictingObjectAttributes("endswith",
    "startswith"),
  validators.ConflictingObjectAttributes("equals",
    "match_regex"),
  validators.ConflictingObjectAttributes("equals",
    "startswith"),
  validators.ConflictingObjectAttributes("match_regex",
    "startswith")}
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
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

Terraform syntax:

```terraform
operator {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233031010232033-3102011120022032-1012131120132022-2330132202112220-1133230300022113-0331000212003013-1330110121201230-2300313233031213"></a>

### Direct properties for `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator`

<a id="canonical-3301202210301300-3301203122103133-0301323010123201-2302321333312200-3022103323333030-2033023212000020-1122032011031202-2012023011103102"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.contains` property

Type: `"string"`. Optional.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The header value must include the specified value as a substring.

<a id="canonical-3330030202302130-0230300012313302-0031320111022032-3032120023031203-0130121302031112-1102323120302333-2231013010002123-3221113113212212"></a>

<a id="canonical-0022003133000121-3112110223032211-1210131223112223-2120110231002302-0023310210312031-2233331023111003-2012110121032302-3103220031321332"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_contain` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not include the specified value as a substring.

<a id="canonical-2103123302232103-0123101231012130-0013103220222122-0200202022113322-3201133303113030-2230323312202100-1220211003200012-3313021210320321"></a>

<a id="canonical-3332303110123112-0332101000212212-1321023102032233-3223222133123030-2023310313210302-0312312202123302-3203301233223200-0131030102002121"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_end_with` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not end with the specified value.

<a id="canonical-0122013221331030-2022331322221032-0010233223111233-0003110130201032-1202133121212030-3322313301011211-0101231213021202-0332200032130122"></a>

<a id="canonical-1120111102122101-0223303003101300-3310210030231020-3322211323310320-1230032202100010-1210302211220123-3202033020311311-3332000000100001"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_equal` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The header value must not match the specified value.

<a id="canonical-3210002003303000-1311000003222312-1100123312131330-2220200213131311-1131022212011323-1011300021302320-2120120313103132-2002121100000200"></a>

<a id="canonical-2123011022122002-3233031312112313-0320300001013231-0022332033100233-1300103330211321-0310302102133333-3011122211000030-0210113331131200"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.does_not_start_with` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The header value must not begin with the specified value.

<a id="canonical-2012102133203210-2113021002102102-1220101232333012-2112313223302002-2120220301223023-0032332103012300-2323210211101111-1123133200212301"></a>

<a id="canonical-1120131323012310-0202020233031311-0112111211321101-0202120201002231-1111212012112133-2102320120202330-0121313002023022-2010131021122101"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.endswith` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The header value must end with the specified value.

<a id="canonical-3201231301323123-2111203001320011-3300213300003313-1101122201311230-1002221102130022-0322311113133101-2112203101302120-3333103110032132"></a>

<a id="canonical-1131002321331321-3213003110320200-1000321302003232-1111221010232000-2030020012212031-2112103201131102-1100332212320222-1103021322201330"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.equals` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The header value must exactly match the specified value.

<a id="canonical-2202321000113311-1120130032110220-0122012120100200-3120030022111122-0212133331002223-3210130300213003-0021032100333312-3200101300300032"></a>

<a id="canonical-0022221123310302-0103131112130321-3132012211022101-0121011010000010-3313332031321331-1121223033322010-0203101013330032-2201213220222031"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.match_regex` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The header value must match the specified regular expression pattern.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

<a id="canonical-0033011312033132-2231310003320332-3031223013111230-3003002110322032-2223313112222011-3320011011321330-1010323333011212-0322200003032232"></a>

<a id="canonical-2032001110322231-0120222332112331-0012023121331110-0121231312023011-1221313022303111-3032122100331022-3220112120101033-3022123332003331"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator.startswith` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The header value must begin with the specified value.

<a id="canonical-3331222303100022-2032300023213202-1132013320331232-2010223021302210-1102122233013032-3012330131230311-1122330220010203-1210211121131131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211)
- cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher

<a id="canonical-3121002331002223-2221133032312221-2313120022332111-2102113221331301-3200110121103021-1010301110211321-2323202321012321-2322003311303202"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
cookie_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200001122213223-3311011333202213-1033030003313313-3220000030310021-0211021101020103-3033321030110312-3103223123120023-1223220121331123"></a>

### Direct properties for `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher`

<a id="canonical-2030000212010232-2030112310201103-2322323131130212-2021133333322222-3010000302012230-0221013003103113-0020231302001210-2202011310001012"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.name` property

Type: `"string"`. Optional.

Cookie Name. Enter the name of the cookie to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [operator](resources--cdn_cache_rule--reference--group-001.md#canonical-3333321221103022-3231020113311333-2030101230122210-3332210112231310-2103212202000031-2131010223103312-2031200301202131-3010100033112333): complete subsection reference.

<a id="canonical-3333321221103022-3231020113311333-2030101230122210-3332210112231310-2103212202000031-2131010223103312-2031200301202131-3010100033112333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211)
- [cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher](resources--cdn_cache_rule--reference--group-001.md#canonical-3331222303100022-2032300023213202-1132013320331232-2010223021302210-1102122233013032-3012330131230311-1122330220010203-1210211121131131)
- cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator

<a id="canonical-2213303211010032-0203332322103123-3202300002112002-2320023231123222-2121030233021102-0223003110310332-2302232032303331-3110331323011330"></a>

Type: `"object"`. single nested block, Optional.

Operator

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("contains",
    "does_not_contain"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("contains",
    "endswith"),
  validators.ConflictingObjectAttributes("contains",
    "equals"),
  validators.ConflictingObjectAttributes("contains",
    "match_regex"),
  validators.ConflictingObjectAttributes("contains",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "startswith"),
  validators.ConflictingObjectAttributes("endswith",
    "equals"),
  validators.ConflictingObjectAttributes("endswith",
    "match_regex"),
  validators.ConflictingObjectAttributes("endswith",
    "startswith"),
  validators.ConflictingObjectAttributes("equals",
    "match_regex"),
  validators.ConflictingObjectAttributes("equals",
    "startswith"),
  validators.ConflictingObjectAttributes("match_regex",
    "startswith")}
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
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

Terraform syntax:

```terraform
operator {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120332012333102-1010101233320112-0123121332303002-1220011112120203-3100011301301302-1213111100310121-1030310123011230-1333022331203120"></a>

### Direct properties for `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator`

<a id="canonical-2300010302231202-2013230322112322-3332303312101320-1232010221001221-1222100323011123-0322033130222321-1003310303312111-0221003023333221"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.contains` property

Type: `"string"`. Optional.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The cookie value must include the specified value as a substring.

<a id="canonical-2121230321002301-3311211203211231-0021322210212223-3012131331023222-2302123220011003-1221003223231132-3202211212312333-0321301223133021"></a>

<a id="canonical-3313223000121322-1030202133123230-2020230232331112-2011132211330131-2220032123320121-3200221202302013-2131001330121000-2310302223010032"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_contain` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not include the specified value as a substring.

<a id="canonical-0320110003010021-3110330200123110-3013211001312231-2330100231233010-3000022310213303-3130200231300013-2012311331223301-1201122110200333"></a>

<a id="canonical-1212010300023100-1310301301021133-3323002021000332-0210020201113332-2210002320101232-0021310210331221-1300220322313132-2223230030200202"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_end_with` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not end with the specified value.

<a id="canonical-2000132120011101-0112033023103311-3333101111110301-1121210203323231-2230211010001100-0211103030313222-1032331231031022-1001031110012112"></a>

<a id="canonical-0231310222103232-3130330312332033-2322120301000211-0133332013111013-0231231033132123-3230231312203003-0111103013112010-3111231103103310"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_equal` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The cookie value must not match the specified value.

<a id="canonical-3001213213112132-3231213131033001-2230220133022300-0212113100201322-2102323211201200-1321230122130302-0113111310121210-1331212023111211"></a>

<a id="canonical-1213230222010130-3301310001123000-2001112203023023-3223303011010320-1231131212110003-0123011021332303-0030320012333103-1222321022300313"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.does_not_start_with` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The cookie value must not begin with the specified value.

<a id="canonical-2113102132311133-0023100131021113-1300020231020331-2031103121023011-1132202310101221-0200103231100333-0302320000033231-1122013331111022"></a>

<a id="canonical-3003021110133130-3222221230322203-2003313332313222-0202332213001113-2021331222313333-0201312332213332-3220323000202101-3211303011122110"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.endswith` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The cookie value must end with the specified value.

<a id="canonical-0101100023121031-3100211101002302-2023233221112131-1002033110301031-1222032032232230-2202033121031010-1301312003013321-1030300210023230"></a>

<a id="canonical-1113320200121210-2330123132020030-0101102120322211-3202213223003003-1102113112320020-2111233302300001-3111033232330010-0222331222030031"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.equals` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The cookie value must exactly match the specified value.

<a id="canonical-1201231213121301-0131130121120233-0332100210212020-1330313031133301-2002012330201112-1231133013010020-3003310103320302-2231212220113122"></a>

<a id="canonical-0331231230220211-3132221113221302-3020321032122113-1301212202032133-1331213300303031-3013032012312220-0012202002311130-1122222111110122"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.match_regex` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The cookie value must match the specified regular expression pattern in PCRE
format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

<a id="canonical-3032332003211223-2132021233022222-3112120111120212-1211112031003023-0002013210000322-3032031323012023-2011212110310302-1222022022321333"></a>

<a id="canonical-3033311022031030-3103301320332211-3121202132322013-1023010222022110-1113100121121110-0330022321223211-0131100100101301-3100122333011223"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.cookie_matcher.operator.startswith` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The cookie value must begin with the specified value.

<a id="canonical-1002100310130020-0310003113131311-2013020101322033-0120132133100322-0332310332020133-1111332232310300-0123303002223030-3330201313200213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list.cache_rule_expression.path_match` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211)
- cache_rules.rule_expression_list.cache_rule_expression.path_match

<a id="canonical-2230021001103331-3022203323321121-2230012211001111-0232013200202331-2111103201031311-3002121100321122-3321301012223322-1213010331022000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
path_match {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021020002202311-0311022002211303-0120211020130312-3321232122300323-0220003312133311-3011232001201321-0113031113102103-2313322231033330"></a>

### Direct properties for `cache_rules.rule_expression_list.cache_rule_expression.path_match`

- [operator](resources--cdn_cache_rule--reference--group-001.md#canonical-2302021031113222-2311330101313232-1203213322131312-2331013311010103-0132022321302301-0101133000133012-1021200111322103-1101023112023131): complete subsection reference.

<a id="canonical-2302021031113222-2311330101313232-1203213322131312-2331013311010103-0132022321302301-0101133000133012-1021200111322103-1101023112023131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211)
- [cache_rules.rule_expression_list.cache_rule_expression.path_match](resources--cdn_cache_rule--reference--group-001.md#canonical-1002100310130020-0310003113131311-2013020101322033-0120132133100322-0332310332020133-1111332232310300-0123303002223030-3330201313200213)
- cache_rules.rule_expression_list.cache_rule_expression.path_match.operator

<a id="canonical-2113002000330232-3310333211013110-2131322012001311-0103101332123312-1130211200300121-3220330213121130-2020212001330032-3002200203002221"></a>

Type: `"object"`. single nested block, Optional.

Operator

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("contains",
    "does_not_contain"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("contains",
    "endswith"),
  validators.ConflictingObjectAttributes("contains",
    "equals"),
  validators.ConflictingObjectAttributes("contains",
    "match_regex"),
  validators.ConflictingObjectAttributes("contains",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "startswith"),
  validators.ConflictingObjectAttributes("endswith",
    "equals"),
  validators.ConflictingObjectAttributes("endswith",
    "match_regex"),
  validators.ConflictingObjectAttributes("endswith",
    "startswith"),
  validators.ConflictingObjectAttributes("equals",
    "match_regex"),
  validators.ConflictingObjectAttributes("equals",
    "startswith"),
  validators.ConflictingObjectAttributes("match_regex",
    "startswith")}
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
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

Terraform syntax:

```terraform
operator {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122020012322200-2212320320220021-0203321301021133-1203131201131202-3301111011301020-1320002310131131-2031232033021103-1110122323012023"></a>

### Direct properties for `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator`

<a id="canonical-0213200012211013-1100322303201133-0102123132123202-3311211110100322-0303133010303012-1131122323203011-0320032231213020-3211011301203123"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.contains` property

Type: `"string"`. Optional.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The path must include the specified value as a substring, up to the
filename.

<a id="canonical-3320022321220123-3013103213310130-2320320230212011-2311103103212331-1011320333011120-2112210320230323-1113121202120121-3132301103232303"></a>

<a id="canonical-1303013001323313-2222302321123130-1032131212223200-3212322310031122-3302101223101131-2002003211000202-1112310003132120-3323302102020001"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_contain` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not include the specified value as a substring, up to the filename.

<a id="canonical-1131312101133211-0303013211130210-2223223012220300-1013201103032033-2321030213230030-0212120302320013-2020111212332220-2330200032213311"></a>

<a id="canonical-0233001322003032-2003122013231102-1010112110200003-2033231321303231-1120011011323230-2231122323113111-3303100012312203-2023133310302212"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_end_with` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not end with the specified value, up to the filename.

<a id="canonical-0110133002331020-2033111101330333-0011123311230200-2011231330003113-3012011233122230-3202311030313222-2030212103110233-1333002202102132"></a>

<a id="canonical-1103231012310130-3100103121201032-3313300302003311-2233121333112230-3222200123202232-3300011120131333-2211110233233021-1100210002002122"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_equal` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The path must not match the specified value, up to the filename.

<a id="canonical-2313000330001203-3111210310213130-0332122321102120-0310120313231112-3310201312012113-2113321001023320-3311103110312120-2330210203113013"></a>

<a id="canonical-1000233211233121-2232133123122302-0123211211101312-2020232311133202-2010320013021313-3001222300133131-1233001113200233-1220220003221331"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.does_not_start_with` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The path must not begin with the specified value, up to the filename.

<a id="canonical-3022101233122003-0023021000223021-3122031203220201-1321331030002231-0133312112012210-1211123221130012-0012213323103110-1032211303120132"></a>

<a id="canonical-1210211032130200-1021032233000101-3133221310331321-1022201301212023-1211002323303023-2033100130103230-3210232301132110-1223312011201322"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.endswith` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The path must end with the specified value, up to the filename.

<a id="canonical-1122011310203203-1203202320103101-0302123103103223-3320310333222131-0222011222123313-0213320321323100-1203233010100232-0030301301221203"></a>

<a id="canonical-1233312002111133-1011303231122130-3300031003122201-1130022033300211-3223313232312032-3003120003210033-0223231223013222-0102220013000302"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.equals` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The path must exactly match the specified value, up to the filename.

<a id="canonical-3233112023321303-0121323330032223-2131223120020123-2031021103133002-2233033132233012-1321103010132002-2222313031120010-1130022331310013"></a>

<a id="canonical-3112111330001301-2133010223130312-3301102313311002-2223202123001131-3032032222303031-3132112101023211-3002202312123210-0121220003233233"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.match_regex` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The path must match the specified regular expression pattern in PCRE format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

<a id="canonical-1133032102223222-2010122211101331-3123022200221100-1000221231302311-1131133233023213-2301120113022203-0010330123230303-0012212030002131"></a>

<a id="canonical-0330203133003130-3220131020213031-1202330000131003-2032300313133013-0133032023231302-0030320122002000-1123031121222231-2220231322303003"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.path_match.operator.startswith` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The path must begin with the specified value, up to the filename.

<a id="canonical-1132133310133103-1022330022230200-3201210301002123-0120022102320312-2100032230203203-1130330300300203-2022323102030133-0322103212202310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list.cache_rule_expression.query_parameters` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters

<a id="canonical-2311131300301233-2300103022032201-0122230231001312-3023000101130333-3011210201331112-0200120021113132-1303200123101220-3122203223102332"></a>

Type: `"object"`. list nested block, Optional.

Query Parameters. List of (key, value) query parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key")}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123303311001311-0211110300323021-2203302231312322-1111001011120201-3020321212012001-0100023321312032-0310113311233302-1203101301323020"></a>

### Direct properties for `cache_rules.rule_expression_list.cache_rule_expression.query_parameters`

<a id="canonical-2211330131032301-1133311302213321-3303203313201032-2001300230020113-3011003331003222-0110131011103312-3203221122002333-1221231110301331"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.key` property

Type: `"string"`. Optional.

The name of the query parameter to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [operator](resources--cdn_cache_rule--reference--group-001.md#canonical-2021003103122331-3221301001303300-3201223101100000-3230313230112122-1331131233033231-3332331020311132-0333222112203022-3322111131020231): complete subsection reference.

<a id="canonical-2021003103122331-3221301001303300-3201223101100000-3230313230112122-1331131233033231-3332331020311132-0333222112203022-3322111131020231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- [cache_rules](resources--cdn_cache_rule--reference--group-001.md#canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--reference--group-001.md#canonical-3110102003302131-2302312133210100-3031102231130310-1111030113012121-2333122333330303-0312333103333000-2013121121313131-3232213102311232)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--reference--group-001.md#canonical-3033332032033132-2030322222333310-3031222331013000-0212212120333312-0211133002303021-0110310230000322-3222230322223002-2321021232022211)
- [cache_rules.rule_expression_list.cache_rule_expression.query_parameters](resources--cdn_cache_rule--reference--group-001.md#canonical-1132133310133103-1022330022230200-3201210301002123-0120022102320312-2100032230203203-1130330300300203-2022323102030133-0322103212202310)
- cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator

<a id="canonical-1311031211032010-3301220321300212-3123211200203102-1322223133312011-0211212222010333-1210312122121210-0313100221221221-0033203322013220"></a>

Type: `"object"`. single nested block, Optional.

Operator

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("contains",
    "does_not_contain"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("contains",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("contains",
    "endswith"),
  validators.ConflictingObjectAttributes("contains",
    "equals"),
  validators.ConflictingObjectAttributes("contains",
    "match_regex"),
  validators.ConflictingObjectAttributes("contains",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_end_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_contain",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_equal"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_end_with",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "does_not_start_with"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_equal",
    "startswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "endswith"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "equals"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "match_regex"),
  validators.ConflictingObjectAttributes("does_not_start_with",
    "startswith"),
  validators.ConflictingObjectAttributes("endswith",
    "equals"),
  validators.ConflictingObjectAttributes("endswith",
    "match_regex"),
  validators.ConflictingObjectAttributes("endswith",
    "startswith"),
  validators.ConflictingObjectAttributes("equals",
    "match_regex"),
  validators.ConflictingObjectAttributes("equals",
    "startswith"),
  validators.ConflictingObjectAttributes("match_regex",
    "startswith")}
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
  "x-ves-oneof-field-cache_operator": "[\"Contains\",\"DoesNotContain\",\"DoesNotEndWith\",\"DoesNotEqual\",\"DoesNotStartWith\",\"Endswith\",\"Equals\",\"MatchRegex\",\"Startswith\"]"
}
```

Terraform syntax:

```terraform
operator {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231231311133201-0033210101111223-2121213312131231-1231012230333312-0132133203202213-1112123211003210-0130312301222030-1011313012012310"></a>

### Direct properties for `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator`

<a id="canonical-1332202222133300-0033312013210122-2303111002203111-1303201211031130-3210122203302120-0100311212232113-2102120130201322-3133100332323323"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.contains` property

Type: `"string"`. Optional.

Exclusive with \[DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals
MatchRegex Startswith\] The query parameter value must include the specified value as a substring.

<a id="canonical-3011322332303023-3202201223130100-0030230330221202-0220121033202221-1001201213022003-3200301313332012-0223112132311120-1333111002212001"></a>

<a id="canonical-2133323231103333-2122210022001230-2121010033120303-0111010020023313-0230010221113030-2111003313011103-2113232101230011-3003123120331120"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_contain` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not include the specified value as a substring.

<a id="canonical-0011001131123022-3303133101023102-0000033000022332-3101100033323302-2232012130122012-2131112303220200-2120102221230032-3231320323120031"></a>

<a id="canonical-2132211302021032-2312123322122001-0003203302113123-2212111301333000-0132011300223301-2310121301021110-0001020103033232-0021010232302213"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_end_with` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEqual DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not end with the specified value.

<a id="canonical-1013011232213233-0000230210030223-2223210002011301-2333023130001231-0103202033023022-0321220130101012-0113230010013101-2000103211210201"></a>

<a id="canonical-1000000310001321-1012231331033010-1013110300102320-1032131011313101-3302000031121001-0333013302203121-0031210200101321-0223323010331220"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_equal` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotStartWith Endswith Equals MatchRegex
Startswith\] The query parameter value must not match the specified value.

<a id="canonical-2303233120122013-0302211131013232-0313300231331013-3333210012022030-0200300022030001-3212030211202133-0110112201011221-0022313123222311"></a>

<a id="canonical-3221331112331100-2112012322133332-1303121013313121-0323130213001103-1330210230332010-1010300121020211-2021131312220211-0101200331002210"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.does_not_start_with` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual Endswith Equals MatchRegex
Startswith\] The query parameter value must not begin with the specified value.

<a id="canonical-0311330313333020-0223002103103231-1012103010222310-0221223120222023-3221301102033013-3210131011010023-3131233121033201-2210021112003032"></a>

<a id="canonical-2231310300012220-1010032213102123-3220001011010220-3131031221321112-3002322122000010-2110232012112213-0221331202033100-0200210300210030"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.endswith` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Equals
MatchRegex Startswith\] The query parameter value must end with the specified value.

<a id="canonical-3010132322321122-2212201201211010-3220022330322200-0223232122111003-3122200231222210-3113000031212201-3200330120303023-0003312300013021"></a>

<a id="canonical-0303332121321330-1313200133311231-3231231111303121-3333331130012103-2132202320022210-0223100322022331-2223300203203211-1220303321200310"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.equals` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
MatchRegex Startswith\] The query parameter value must exactly match the specified value.

<a id="canonical-3311132011111121-3012302310323112-1121103201223232-3221123003331300-1133132232123001-1323223320322333-0330022032003302-0111013333302200"></a>

<a id="canonical-0320211211000001-0121133133321222-2001302303032223-0222121110010022-2232222231032110-1330132203320320-3220221101032200-0212330331220031"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.match_regex` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals Startswith\] The query parameter value must match the specified regular expression pattern in
PCRE format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

<a id="canonical-2310331232223302-0113322303333210-0012130010031333-1102131301302322-3301113132102110-2331032200002000-3123001022120002-2031320102200022"></a>

<a id="canonical-0313301032110200-0203222313111121-2303221332102211-3310302021300101-0310130103000030-2221123321010103-1321103201030132-1001021313133133"></a>

#### `cache_rules.rule_expression_list.cache_rule_expression.query_parameters.operator.startswith` property

Type: `"string"`. Optional.

Exclusive with \[Contains DoesNotContain DoesNotEndWith DoesNotEqual DoesNotStartWith Endswith
Equals MatchRegex\] The query parameter value must begin with the specified value.

<a id="canonical-1212033023211232-3031321101321003-2333222011200010-1230320100011230-0012013221003131-0331013111020303-1122320301310222-1203323113231230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Property reference](resources--cdn_cache_rule--reference--group-001.md#canonical-0031322322320201-1123033010333010-3012313122023131-0002330011303300-2213110212022132-1021300113022011-2021133032110332-0012323030102001)
- timeouts

<a id="canonical-1003323221203111-2313300330023312-0013302333120001-2101012220223321-2313212100330322-2031321001013112-3020201022030230-1112113321000321"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232132221333133-2312201002211001-2122100030022001-3013030233231003-2311203202320231-1123211332002213-0023112333132313-0113003301323320"></a>

### Direct properties for `timeouts`

<a id="canonical-2210111332222310-0122012230032321-0202233121110332-3231202311331213-0120130013202230-0120323131321313-3232113202231320-1103021220331031"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3212201313203022-0030033231023211-3120331011322112-3001322200310130-0201213220033222-3210032230320300-0012302230103131-3111021133001330"></a>

<a id="canonical-2111001223300212-1033013003201123-0021210202230320-1100312132122102-1300321113220200-0102310100101102-1033203032302223-3020020202312102"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0132011103112100-2123213010010232-1113223322331200-2301101302232012-1122120133032222-1310030131011211-0213222103211000-1113133310212212"></a>

<a id="canonical-3212111320221101-3201232112013203-3121033330023201-0230213133332312-0002110110330112-2131113132031013-3302011030303223-3230301001233231"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3101103011202003-2100201013223010-2021203232310131-2131010112230230-1222111131030013-0233123210220131-3002312032122021-3022302303200231"></a>

<a id="canonical-2131320203212200-2212133022010301-3111310203021103-0111130331031222-2220121301310201-3313221311112032-0133111022123332-1312322210201301"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
