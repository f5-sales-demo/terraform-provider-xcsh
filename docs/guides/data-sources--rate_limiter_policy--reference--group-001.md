---
page_title: "xcsh_rate_limiter_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy reference."
---

# xcsh_rate_limiter_policy reference

<a id="canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200113113120312-3111020010300321-1312100201202322-0320033312211023-1311123122112313-1313300232232233-0011300330310221-3301120023111111"></a>

## Property reference — Property reference / 311311313303 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- Property reference

<a id="canonical-0011103120313023-1303022102210001-3030121030321333-2200333310112121-1033112133022121-1303020132312320-2120133011211211-1211030102201012"></a>

## Direct properties — Property reference / 311311313303 / 3

<a id="canonical-3130012233022310-3122210320112130-1332033033220023-0000123300112001-0320333101101200-2012210133121130-2010313131221201-3323113203001200"></a>

<a id="canonical-3121001002021313-1100122003231230-2023331210123120-2111111322200002-3102320212130111-3210232012300320-2322302321333332-2013101331012333"></a>

## annotations property — Property reference / 311311313303 / 4

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

- [any_server](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2121000013301231-3221330132132110-0100321023133212-2223301013231102-2310221021120302-2132001033022231-1131031203333131-1111300212113321): complete subsection reference.

<a id="canonical-2121230222100122-0110200111000011-1200113101030001-2021212223011313-0100103303001203-2322030210112323-2133011232220333-3031033302101103"></a>

<a id="canonical-1222100302001310-3310301320032311-3111212122103030-2210011213223031-3320123101121311-2022231300110131-0121010210320211-3123020011123022"></a>

## description property — Property reference / 311311313303 / 5

Type: `"string"`. Computed.

Description of the RateLimiterPolicy.

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

<a id="canonical-0223220131100300-3100323321102330-3333213310022023-2203311301330122-0302131033030212-0310002210323011-2102322203230012-0233321000120301"></a>

<a id="canonical-0030112123300311-0100103100212233-0300320201201320-1130122102231312-3101131332313021-0231300012002230-2013313002112110-0313123231333321"></a>

## ID property — Property reference / 311311313303 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3211233211031203-0110212133133103-3130123011111312-2031012312331211-3102122213302321-2012230230031011-1120203211122223-1233132112211031"></a>

<a id="canonical-2231003311321321-1001121311030032-1213202013123033-0330213122123023-3103303232022322-1200112130310201-0201211033330000-1003032311133213"></a>

## labels property — Property reference / 311311313303 / 7

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

<a id="canonical-0023201023031113-2000023201311003-0131123123100113-2111313002332133-0210120102132222-0030000321023021-3321022003032023-2131332221131230"></a>

<a id="canonical-0202010131013111-1022102030321203-2203103322310231-3123033002320312-3132100020311203-0220200012132102-1230212233131100-3022232203311013"></a>

## name property — Property reference / 311311313303 / 8

Type: `"string"`. Required.

Name of the RateLimiterPolicy.

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

<a id="canonical-2011320210313301-0230021320333211-1231210020201132-1220213002122032-1103301030220312-0030020110122213-0223201211313333-0303023133200032"></a>

<a id="canonical-0320231221310001-1323023010122233-1203202123003033-1230220003130233-1020220032333010-1230102333331302-1323032010302230-2001222201102031"></a>

## namespace property — Property reference / 311311313303 / 9

Type: `"string"`. Required.

Namespace where the RateLimiterPolicy exists.

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

- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011): complete subsection reference.

<a id="canonical-1131133311301313-0022011121030013-0200023033130220-0021013132012103-3120211012110121-2220231113300301-2133001232203010-0313013310220122"></a>

<a id="canonical-0203313221031031-3311200333210033-2121003231033233-2323321123020121-1220000020103203-3320130302310210-3122121211320321-0011221113003012"></a>

## server_name property — Property reference / 311311313303 / 10

Type: `"string"`. Computed.

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server. The actual names for the server are extracted from the HTTP Host header and the name of the
virtual\_host for the request.

Upstream description:

Exclusive with \[any\_server server\_name\_matcher server\_selector\] The expected name of the
server. The actual names for the server are extracted from the HTTP Host header and the name of the
virtual\_host for the request.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [server_name_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3122202032011323-1213021132223201-3223120002010000-0333012013320120-1332310231213133-1210322022033020-3310220211132300-3233322110130012): complete subsection reference.

- [server_selector](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2300221102110132-2231203300303222-2331320002303300-1301313113022220-0110000303303010-0302231130010202-2312132002122201-3033221320120000): complete subsection reference.

<a id="canonical-2133010101220001-3203300303010030-0013323001120032-0232020001213033-2113032003210300-0230330010121032-2100220321333333-2123032210121011"></a>

## All schema paths — Property reference / 311311313303 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3130012233022310-3122210320112130-1332033033220023-0000123300112001-0320333101101200-2012210133121130-2010313131221201-3323113203001200) |
| `any_server` | [any_server](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2001321212100203-1000311320231102-0101223232313023-1300213222100123-3132303022131132-1310203101320230-1313123321021233-3131003232022003) |
| `description` | [description](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2121230222100122-0110200111000011-1200113101030001-2021212223011313-0100103303001203-2322030210112323-2133011232220333-3031033302101103) |
| `id` | [ID](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0223220131100300-3100323321102330-3333213310022023-2203311301330122-0302131033030212-0310002210323011-2102322203230012-0233321000120301) |
| `labels` | [labels](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3211233211031203-0110212133133103-3130123011111312-2031012312331211-3102122213302321-2012230230031011-1120203211122223-1233132112211031) |
| `name` | [name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0023201023031113-2000023201311003-0131123123100113-2111313002332133-0210120102132222-0030000321023021-3321022003032023-2131332221131230) |
| `namespace` | [namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2011320210313301-0230021320333211-1231210020201132-1220213002122032-1103301030220312-0030020110122213-0223201211313333-0303023133200032) |
| `rules` | [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0122302211123023-0212312301102332-2302112111230111-3020010302111321-3120102032312331-2120011000223030-3322331021003202-2300010232330322) |
| `rules.metadata` | [rules.metadata](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0101223301322301-1013211312022022-2112311211103101-1322302110001132-0110323200011021-0031322332232000-1010332030100330-3202212103213202) |
| `rules.metadata.description_spec` | [rules.metadata.description_spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2002321112131332-2330203300231313-1221002020000321-2112030132122032-1221233101133221-2331013010012200-2301213000101213-3220123321133312) |
| `rules.metadata.name` | [rules.metadata.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3200012232232020-0303100023203131-1020100112311130-2032312231101031-2112302301110331-0203001202322312-0310120302002112-0121001321212033) |
| `rules.spec` | [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0013200200321323-2130233311332302-0331311333211010-3123303102113132-0022311022003211-3121011303311103-1112102322311010-3321002102232021) |
| `rules.spec.any_asn` | [rules.spec.any_asn](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2211133200330030-0330010021213213-1331211100121012-1233200320330110-2111332003120001-2133000223003130-1300132000311231-2020331101033112) |
| `rules.spec.any_country` | [rules.spec.any_country](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1211031201233030-2121030111120101-0310302322002123-2233233133322232-3210020210112233-1020330133302232-2113303331120300-3110021322332021) |
| `rules.spec.any_ip` | [rules.spec.any_ip](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0311102232022310-2010200033233112-2223313033111103-0223022312012110-0213200310122033-2031223301132102-0313110001101130-1322103320033132) |
| `rules.spec.apply_rate_limiter` | [rules.spec.apply_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2320230212321103-0132032202031100-1323301002111003-2232223332221201-3122011122030023-1320033301111211-3023211020022032-0321210020330321) |
| `rules.spec.asn_list` | [rules.spec.asn_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3123203023101330-2101111333010212-2003302021230102-0321002023213221-3230013303023123-2030110131030121-2121211200321222-1012220131101223) |
| `rules.spec.asn_list.as_numbers` | [rules.spec.asn_list.as_numbers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2330020201233102-3002031322030310-2012330332333313-1231121231210332-0201113210023220-2220333113312023-0210201302011113-2012010220032221) |
| `rules.spec.asn_matcher` | [rules.spec.asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3311313022111002-0201002131220333-2310111202123213-2111312122332222-2032100033131023-1023011313130001-1100011110001320-3120012310231110) |
| `rules.spec.asn_matcher.asn_sets` | [rules.spec.asn_matcher.asn_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0302311102200001-2200232111320301-1101121223112232-0133302230223021-2110132102012020-1323232100322121-2030333021220332-3331333230331131) |
| `rules.spec.asn_matcher.asn_sets.kind` | [rules.spec.asn_matcher.asn_sets.kind](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0223230130330010-0132030331212332-3121222211131230-0011001310321030-2331220021003111-1213331000320000-0230001011221121-2111301012031201) |
| `rules.spec.asn_matcher.asn_sets.name` | [rules.spec.asn_matcher.asn_sets.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2300210302103310-1110331203003103-0022320200121002-1123211210122332-1313011112310220-2211110121132120-0021232213230302-2012331102310302) |
| `rules.spec.asn_matcher.asn_sets.namespace` | [rules.spec.asn_matcher.asn_sets.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2302113330130003-3001223221230001-2000000010010031-3230111310323311-3031012213323023-3211223023012312-0013033200333103-1120313311101333) |
| `rules.spec.asn_matcher.asn_sets.tenant` | [rules.spec.asn_matcher.asn_sets.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0300033023030201-1312233000210220-0220332013311202-2022313211012300-0002332013303032-2111223113010322-3121132223022133-2310212210210012) |
| `rules.spec.asn_matcher.asn_sets.uid` | [rules.spec.asn_matcher.asn_sets.uid](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3020011133300030-3102133222111120-2233110313002222-3211322010022133-2230302313030111-0301131132111210-3022012122300120-3333002312112123) |
| `rules.spec.bypass_rate_limiter` | [rules.spec.bypass_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0303310011001331-1301100121302202-0122100231122022-2210330313020110-3220032333332211-3000021221121111-1031230201112032-0230302011330211) |
| `rules.spec.country_list` | [rules.spec.country_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3023331101223023-2100320011100021-1200313201023231-1030021332313330-3000312010121030-2131131120113011-1120112331312332-0300302303233130) |
| `rules.spec.country_list.country_codes` | [rules.spec.country_list.country_codes](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3321323020202322-0130131311101332-2321010310311101-1133320110212112-2210203333320331-1222111211100333-3023011312313213-2302032310122010) |
| `rules.spec.country_list.invert_match` | [rules.spec.country_list.invert_match](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3323021130223302-2131233001230233-2213231120130010-2031301211123212-1232023222030212-1302201213120122-0100110202222101-0221133013231211) |
| `rules.spec.custom_rate_limiter` | [rules.spec.custom_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0113232333221103-2023200102021013-0002230203103123-2220203212001203-2210102101000332-0303110001200021-2120023130231200-3301121200032121) |
| `rules.spec.custom_rate_limiter.name` | [rules.spec.custom_rate_limiter.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1203030201133120-0211311202213220-1131100112013302-0001020221110032-3230221002000331-2221230203211230-3230212031002223-1310333010310121) |
| `rules.spec.custom_rate_limiter.namespace` | [rules.spec.custom_rate_limiter.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1211213201030020-0321332220320031-0222120101130120-0332001120122331-0332032221122202-0133201320322110-2021120032020020-0222211131332223) |
| `rules.spec.custom_rate_limiter.tenant` | [rules.spec.custom_rate_limiter.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3110100031001023-2110223132332311-1032130111330012-2333211320030030-2131300010122000-0103033130230233-3221023102121101-3002331110311230) |
| `rules.spec.domain_matcher` | [rules.spec.domain_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0320022310231323-0001223203333033-2003101211011210-3210031331202300-2201211213312121-1110002102003211-2021210211223330-1020213133200233) |
| `rules.spec.domain_matcher.exact_values` | [rules.spec.domain_matcher.exact_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3011323310220102-2010321332203013-2221113111220031-0133221212113002-0001321232023310-3110332231003021-0332133300121230-0000332313331000) |
| `rules.spec.domain_matcher.regex_values` | [rules.spec.domain_matcher.regex_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1202300203321232-0203302120033333-1110310130010130-3300131100221210-2001202102230311-0322210211010310-3201321332022310-1021332330230101) |
| `rules.spec.headers` | [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1110201011332120-1031102100303120-1003120332030112-3233303212320210-1031212032012231-1203132232120132-3232020132331131-3102133131330301) |
| `rules.spec.headers.check_not_present` | [rules.spec.headers.check_not_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2203033222000222-1232230021132131-3030323020032000-2001021112112331-0201113233120101-3112132320231220-1123202113021103-2212302111031132) |
| `rules.spec.headers.check_present` | [rules.spec.headers.check_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2011313332301002-0022132321020020-1331032001332011-3011313201111210-0103221132230301-1122311011012132-0210011031303223-1220111000000201) |
| `rules.spec.headers.invert_matcher` | [rules.spec.headers.invert_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2212201133001200-2010102320001210-3123221303121013-3122233230030211-3323323103121220-2032001310013202-1312101031030201-3300103100312221) |
| `rules.spec.headers.item` | [rules.spec.headers.item](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3121103113031321-1232211232123102-2033322233320003-1123003100030132-0010223320113302-3223323231300132-1002321031132310-2233032320012021) |
| `rules.spec.headers.item.exact_values` | [rules.spec.headers.item.exact_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1132320221221201-0133133132300033-3231332101111000-0010111101232011-0222100212101032-1300122010133310-0210333031033312-1303333022303013) |
| `rules.spec.headers.item.regex_values` | [rules.spec.headers.item.regex_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1303230320000330-3330131320110110-0211331203303133-3032023121300211-1102133221321102-3020133120122322-3032101333201203-1332003232030102) |
| `rules.spec.headers.item.transformers` | [rules.spec.headers.item.transformers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0312332112120312-1223010100221311-0220312113303033-1101212121213311-3101203322312302-1101322213130300-2232333312212023-3022210011303011) |
| `rules.spec.headers.name` | [rules.spec.headers.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3110020111321133-2332101021003222-0020121101202022-3021121011233330-0000121121001310-0002313001300303-1023302123113000-2020001030300212) |
| `rules.spec.http_method` | [rules.spec.http_method](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0200003020123320-2301212032101023-1300200031330023-1332200021133101-2210110232330103-2331331303110120-0120331113330000-3131002202330222) |
| `rules.spec.http_method.invert_matcher` | [rules.spec.http_method.invert_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0113012222313330-3031011203021233-1301303313011201-3022021333212203-1100313302221113-1000302013320101-1200013003210010-2211031301030331) |
| `rules.spec.http_method.methods` | [rules.spec.http_method.methods](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3132012012212220-0122000122202210-0232333212330323-0112301231001222-2002321010223100-0022333211000223-2000321121212002-1331313313221003) |
| `rules.spec.ip_matcher` | [rules.spec.ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2112020200312101-2213003110012023-3111311002023132-3121300113202311-0012111220031203-0013102203302031-2313320302003223-2033131222312032) |
| `rules.spec.ip_matcher.invert_matcher` | [rules.spec.ip_matcher.invert_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2003031211211033-0023203332310001-3033202300000123-1212333202231032-1110220302022130-1230101321323033-3110200120000231-3303230112331120) |
| `rules.spec.ip_matcher.prefix_sets` | [rules.spec.ip_matcher.prefix_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0330133233312112-3123020202030111-3101220221001021-2300111312210330-3022032020201030-1023013000003133-0220033202223111-0113121311132200) |
| `rules.spec.ip_matcher.prefix_sets.kind` | [rules.spec.ip_matcher.prefix_sets.kind](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1222220001131030-2230012211220010-0302010321313112-2032012330021033-1221130221300031-2103012301121011-2030020001122223-2233332101301103) |
| `rules.spec.ip_matcher.prefix_sets.name` | [rules.spec.ip_matcher.prefix_sets.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2121013003023223-0103101023312013-1213131310032032-2123103332230310-1321232220002233-0333212233232022-0003333123203220-1213110030031033) |
| `rules.spec.ip_matcher.prefix_sets.namespace` | [rules.spec.ip_matcher.prefix_sets.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3031312210332220-0201233110101012-0212120020201210-0312332221202022-3100111102210233-1133123213223132-1011231311322201-2031310312222122) |
| `rules.spec.ip_matcher.prefix_sets.tenant` | [rules.spec.ip_matcher.prefix_sets.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2230200113111321-0320101021001021-2320232302203300-3020110130101033-2101003231223130-0321220111230231-0211021312130312-1200131132310323) |
| `rules.spec.ip_matcher.prefix_sets.uid` | [rules.spec.ip_matcher.prefix_sets.uid](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0123112320323112-3010112322101003-2122310313023202-1223321123221331-3233001011323033-3332133030312223-1122110223201021-1332231031030000) |
| `rules.spec.ip_prefix_list` | [rules.spec.ip_prefix_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2213311021302313-3300311310330122-0212212033231113-3011313121230311-1330011130202110-0202111122213231-1021131013320003-1332131322201200) |
| `rules.spec.ip_prefix_list.invert_match` | [rules.spec.ip_prefix_list.invert_match](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0021331212232120-0322020101203033-3233210211132120-1013121212221110-1011330021210203-0201322023331013-2013102310132011-0312132311110132) |
| `rules.spec.ip_prefix_list.ip_prefixes` | [rules.spec.ip_prefix_list.ip_prefixes](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3112023220102330-1232321010231000-2221232000223012-3123022231212312-0031210033033022-3202301231231333-3232120133120311-0211301212313011) |
| `rules.spec.path` | [rules.spec.path](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2311012332211012-1312233330223231-0232101122113212-2010112103311311-0011232311220202-1202331100132312-0231011323102032-3221032221331310) |
| `rules.spec.path.encoded_path_matcher` | [rules.spec.path.encoded_path_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2133223230030131-0011310113201330-2103132021121100-0113102230031100-0203011000023120-0003122122112131-3311111202110222-2031231201020223) |
| `rules.spec.path.exact_values` | [rules.spec.path.exact_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3022021232010123-3200012000330201-3330213000201030-3023333123310203-1032022011133132-0121202321311331-3110332023003112-0222201030202200) |
| `rules.spec.path.invert_matcher` | [rules.spec.path.invert_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3000321201133310-0320003210201020-3321133233311130-0012112103203331-2021231100022303-0032013203023233-0323303201311112-3112330300200311) |
| `rules.spec.path.prefix_values` | [rules.spec.path.prefix_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2313231323223100-1132012222233330-2032230033330202-2320203230010213-0111133202331132-2002013302000302-0122003120120310-1331221013213232) |
| `rules.spec.path.regex_values` | [rules.spec.path.regex_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0101020333233202-1200223311122000-1200000000221202-0101303103311313-0300122013211112-1002223000300230-1210102223100222-2102022013223131) |
| `rules.spec.path.suffix_values` | [rules.spec.path.suffix_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0122320213010223-2111300223232000-2200330123013231-2323021230110233-3101210022012301-1210233333101113-3020113032000330-1111123111322033) |
| `rules.spec.path.transformers` | [rules.spec.path.transformers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1203001121011300-3120331121102320-0133033133132233-0210303100012233-3113221232231232-1230033010202333-3330101103320220-3102321122113232) |
| `rules.spec.segment_policy` | [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3011030132330212-1220223032003031-1102231212203213-2221300130330003-0201033133213333-0023303303331103-0120113012101203-0133031320131202) |
| `rules.spec.segment_policy.dst_any` | [rules.spec.segment_policy.dst_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2121313230303301-2100020332020300-2100213331003033-2010210003111021-2302300333110113-0023301030232033-0102133113322220-2013020201110230) |
| `rules.spec.segment_policy.dst_segments` | [rules.spec.segment_policy.dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1031220301022130-2022323332222120-2031203101000301-0330012223022203-3230211012112303-0203233323022323-0112300020311022-0321210201331232) |
| `rules.spec.segment_policy.dst_segments.segments` | [rules.spec.segment_policy.dst_segments.segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0032010113022333-3331312321133311-3303132213200003-3030032002330220-3320121300012103-0103023033330221-1212001102321333-2201131233320300) |
| `rules.spec.segment_policy.dst_segments.segments.name` | [rules.spec.segment_policy.dst_segments.segments.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1321231020010312-0321313211312300-3101020333322303-3113123323000133-3100233020230200-0130002213120332-2322103033032031-3122013110102233) |
| `rules.spec.segment_policy.dst_segments.segments.namespace` | [rules.spec.segment_policy.dst_segments.segments.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0133001102100213-3122202322200201-2122213203230003-1013323313201002-2220332023311011-2011230213002222-1301222323031313-2121112101112330) |
| `rules.spec.segment_policy.dst_segments.segments.tenant` | [rules.spec.segment_policy.dst_segments.segments.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3131111110203221-3002213003311212-2113111223220021-3030033113111200-2122320003110003-3013221200300033-1032101301323230-1331010020100230) |
| `rules.spec.segment_policy.intra_segment` | [rules.spec.segment_policy.intra_segment](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2101220011333001-2021220032012333-0302010222001310-1130323210030022-0011103033020222-3131112111221222-1223301202111130-3111332120321302) |
| `rules.spec.segment_policy.src_any` | [rules.spec.segment_policy.src_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2133103221101302-2010022203320323-1002100100311113-3020213212031122-2331200112223301-1313120132323330-2230310320112011-2102021221132102) |
| `rules.spec.segment_policy.src_segments` | [rules.spec.segment_policy.src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1120331301203123-2330032201311323-0022100213233110-3232013201113322-0221031323220200-1300233000300203-3103031112331133-3111213100013102) |
| `rules.spec.segment_policy.src_segments.segments` | [rules.spec.segment_policy.src_segments.segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3003011310222002-1303311302313111-1120312323032113-0101122123002230-1022123203220302-2313201233112312-0201210100013230-2020323032132032) |
| `rules.spec.segment_policy.src_segments.segments.name` | [rules.spec.segment_policy.src_segments.segments.name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0023031300101121-0230311331213133-2101212303330200-0212032102002313-3311130113101222-2203311311322021-2113232012013100-1002021111003213) |
| `rules.spec.segment_policy.src_segments.segments.namespace` | [rules.spec.segment_policy.src_segments.segments.namespace](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1031212202030120-0021221120323223-0123001012131123-3100230010030111-2103120012020112-3310321022133220-3020032031100222-3032123203122001) |
| `rules.spec.segment_policy.src_segments.segments.tenant` | [rules.spec.segment_policy.src_segments.segments.tenant](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3100221010032130-2110032031120230-3033211011321101-1320230120311232-0031303310310011-2223011221002123-0200110013330210-1210303132223023) |
| `server_name` | [server_name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1131133311301313-0022011121030013-0200023033130220-0021013132012103-3120211012110121-2220231113300301-2133001232203010-0313013310220122) |
| `server_name_matcher` | [server_name_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1120312101000322-3321303131321311-1322332210112222-3002333130320100-3332030020030203-2301012200300031-0122003311122200-3112032101011210) |
| `server_name_matcher.exact_values` | [server_name_matcher.exact_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2033020121101210-1220000011001213-3020012332031132-2330203313230322-1131331131013333-1003020202302300-1001230102021310-1030313033032201) |
| `server_name_matcher.regex_values` | [server_name_matcher.regex_values](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3233023101100012-0130301013223302-0321201303231322-0220132010112230-3010300023103330-2303103131313130-0101003203123312-1102303101013323) |
| `server_selector` | [server_selector](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1330331320203033-1300323033303021-0323110232023001-3312311332112311-0323100030122312-0320220320131120-1321303322330100-0033321012201122) |
| `server_selector.expressions` | [server_selector.expressions](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1223312111322011-2030213123122310-3031003211201300-3032030230220221-1221202120010300-0202333012300203-1113322313003013-2103133031300133) |

<a id="canonical-1120230210122301-3310301120022111-0213202220122230-3133303130233221-2222023122231210-2223022200002322-1223131112332133-1233302120130312"></a>

## Next pages — Property reference / 311311313303 / 12

- [any_server](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2121000013301231-3221330132132110-0100321023133212-2223301013231102-2310221021120302-2132001033022231-1131031203333131-1111300212113321)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [server_name_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3122202032011323-1213021132223201-3223120002010000-0333012013320120-1332310231213133-1210322022033020-3310220211132300-3233322110130012)
- [server_selector](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2300221102110132-2231203300303222-2331320002303300-1301313113022220-0110000303303010-0302231130010202-2312132002122201-3033221320120000)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-2121000013301231-3221330132132110-0100321023133212-2223301013231102-2310221021120302-2132001033022231-1131031203333131-1111300212113321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323013312103023-2313202333300331-3102200330020000-1233103012201120-1000320012012031-3232032321020010-2223013322130101-1021033103120010"></a>

## any_server — any_server / 223030003103 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- any_server

<a id="canonical-2001321212100203-1000311320231102-0101223232313023-1300213222100123-3132303022131132-1310203101320230-1313123321021233-3131003232022003"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_server, server\_name, server\_name\_matcher, server\_selector\] Enable this option

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

- [any_server](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2001321212100203-1000311320231102-0101223232313023-1300213222100123-3132303022131132-1310203101320230-1313123321021233-3131003232022003)
- [server_name](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1131133311301313-0022011121030013-0200023033130220-0021013132012103-3120211012110121-2220231113300301-2133001232203010-0313013310220122)
- [server_name_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1120312101000322-3321303131321311-1322332210112222-3002333130320100-3332030020030203-2301012200300031-0122003311122200-3112032101011210)
- [server_selector](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1330331320203033-1300323033303021-0323110232023001-3312311332112311-0323100030122312-0320220320131120-1321303322330100-0033321012201122)

Select alternatives according to the provider validators above.

<a id="canonical-1000232123133001-2333111200210133-2121021320001122-1031231210001033-0201101222120210-3103221023213012-1030322312130123-0221112212212321"></a>

## Direct properties — any_server / 223030003103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330312100200031-2012133012001320-2133101202113020-0130010211222022-3213030132230231-1131201301322323-3003310013003020-0323310302203232"></a>

## Next pages — any_server / 223030003103 / 4

- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302201133002231-1003010203313010-2312321312103312-2330212022221112-3033212332210312-0213233112321120-3200000133031023-2000111120002332"></a>

## rules — rules / 212200131120 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- rules

<a id="canonical-0122302211123023-0212312301102332-2302112111230111-3020010302111321-3120102032312331-2120011000223030-3322331021003202-2300010232330322"></a>

Type: `"list"`. Computed.

List of RateLimiterRules that are evaluated sequentially till a matching rule is identified.
Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

A list of RateLimiterRules that are evaluated sequentially till a matching rule is identified.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2021223002020210-1300313212210222-1211232100113023-0120032331002103-3001133102203230-2111021212323203-2312212002320030-2233100112100220"></a>

## Direct properties — rules / 212200131120 / 3

- [metadata](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1301202323201131-1003133112300122-3023011323230300-2010010301002230-1331223311130332-2221223331323010-2301203311102132-1201002332311232): complete subsection reference.

- [spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323): complete subsection reference.

<a id="canonical-2320111322031202-1130330130221113-2001032323202302-2221300233000232-3223213113333002-1033121311333112-1233023202302121-3032312230301123"></a>

## Next pages — rules / 212200131120 / 4

- [rules.metadata](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1301202323201131-1003133112300122-3023011323230300-2010010301002230-1331223311130332-2221223331323010-2301203311102132-1201002332311232)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-1301202323201131-1003133112300122-3023011323230300-2010010301002230-1331223311130332-2221223331323010-2301203311102132-1201002332311232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010323322210013-2011200301331332-0101230323313333-3103230222323230-3020130000123031-1012222113022032-0323222222101333-2300013321201331"></a>

## rules.metadata — metadata / 001121221012 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- rules.metadata

<a id="canonical-0101223301322301-1013211312022022-2112311211103101-1322302110001132-0110323200011021-0031322332232000-1010332030100330-3202212103213202"></a>

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

<a id="canonical-2231100310000303-2213111133200023-2010310331011020-0310311301201300-1122110113303312-0203002301202301-1231221101203101-1230113113221011"></a>

## Direct properties — metadata / 001121221012 / 3

<a id="canonical-2002321112131332-2330203300231313-1221002020000321-2112030132122032-1221233101133221-2331013010012200-2301213000101213-3220123321133312"></a>

<a id="canonical-0323001033320332-2121001112001220-0021113121301211-3011033011130003-1033202202123303-2101121001323232-0223302120011210-3122132332101011"></a>

## description_spec property — metadata / 001121221012 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3200012232232020-0303100023203131-1020100112311130-2032312231101031-2112302301110331-0203001202322312-0310120302002112-0121001321212033"></a>

<a id="canonical-3223331230111020-3002000231022032-3023202033310323-3032331300030112-1110221000223303-3130233311213110-2202131030020100-1011023031330333"></a>

## name property — metadata / 001121221012 / 5

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

<a id="canonical-1302033001322100-3330002202233303-2232301113320303-3131300211313320-3322333332133232-2301000030103133-3302300011312112-0113211131310200"></a>

## Next pages — metadata / 001121221012 / 6

- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120222333332020-0213120320303300-2311103303021230-0212101103332202-1022011323332310-3122202003301230-1023331301213121-2221211031022301"></a>

## rules.spec — spec / 320320310230 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- rules.spec

<a id="canonical-0013200200321323-2130233311332302-0331311333211010-3123303102113132-0022311022003211-3121011303311103-1112102322311010-3321002102232021"></a>

Type: `"single"`. Computed.

Rate Limiter Rule Specification. Shape of Rate Limiter Rule.

Upstream description:

Shape of Rate Limiter Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"apply_rate_limiter\",\"bypass_rate_limiter\",\"custom_rate_limiter\"]",
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-country_choice": "[\"any_country\",\"country_list\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

<a id="canonical-1213303332212032-0033100111313313-1200210331003313-0312101113023130-2201322123102330-3213002100230320-1113010202303030-0321112200202002"></a>

## Direct properties — spec / 320320310230 / 3

- [any_asn](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1303030310132112-2132322300021121-1130020232311303-2012230022310201-3011001331132232-0222032032302123-2112030232120001-1102121321203333): complete subsection reference.

- [any_country](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2032132111200112-0212122320321102-2130210330330223-1000333220333020-0322323310212033-1333002032310300-0303132033112120-1113011331230220): complete subsection reference.

- [any_ip](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3030300000310132-3013013320203011-3131031020133123-2033121031233301-3222030031320212-1310210133001100-1303123330130002-3322123331231311): complete subsection reference.

- [apply_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0122133122213230-3020232121202333-3110012000201131-0323322111333113-0232131100200202-2120120302031031-1102001033310120-0210313223333230): complete subsection reference.

- [asn_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3012200013232231-0112112212122230-3221210023330312-2032320202230101-3323213213322333-0230123020101020-1310023203030122-3131322221333020): complete subsection reference.

- [asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3002221121111303-2331202203130210-1332331020200211-3010112211110012-0032023322323303-2113013110212101-2120023130220301-0322213133032000): complete subsection reference.

- [bypass_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2133012112012103-3221331232323300-0110230101211323-1203302221310110-3131123221212100-2203223010220212-2130311130310121-2301030133031122): complete subsection reference.

- [country_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3213003203031101-0230212312221233-0302020132221010-1321321122030321-3020101312332322-0031120212120232-3100021120013123-1321302011333122): complete subsection reference.

- [custom_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0211133231322303-0311133231212303-2113002322200333-0232112201021322-3320012111202101-1130031032312020-3112020010310301-2132011331023332): complete subsection reference.

- [domain_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2303121332011111-1232200221302002-3022032030001130-0230223111223022-3222030310332321-1010210020132012-0203313330323112-2222011312021032): complete subsection reference.

- [headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201): complete subsection reference.

- [http_method](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1003111120111010-1202121130322131-0303110233212223-0131132030100322-1120130333033103-1122133331132333-1312330333111123-3103233022322212): complete subsection reference.

- [ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2112002322133222-3333220133101103-0102002332002330-2211132331201332-1123022001030310-2011113311303031-1232200100012331-0313020121012011): complete subsection reference.

- [ip_prefix_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0303101313121210-0310111031310101-3220232113223221-0331032013032223-3231300001123323-3020301022000223-0010002203110030-2222030301212202): complete subsection reference.

- [path](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0121221232200020-2022030220023313-0020220130020131-2132003303231023-3100020030022230-2122103121031331-1112310202213023-1300031300222332): complete subsection reference.

- [segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132): complete subsection reference.

<a id="canonical-2021003022111211-2100330132211230-0030330302131221-1032332321332332-1231120330033331-3123030321130210-1021210301330132-2223333231302031"></a>

## Next pages — spec / 320320310230 / 4

- [rules.spec.any_asn](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1303030310132112-2132322300021121-1130020232311303-2012230022310201-3011001331132232-0222032032302123-2112030232120001-1102121321203333)
- [rules.spec.any_country](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2032132111200112-0212122320321102-2130210330330223-1000333220333020-0322323310212033-1333002032310300-0303132033112120-1113011331230220)
- [rules.spec.any_ip](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3030300000310132-3013013320203011-3131031020133123-2033121031233301-3222030031320212-1310210133001100-1303123330130002-3322123331231311)
- [rules.spec.apply_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0122133122213230-3020232121202333-3110012000201131-0323322111333113-0232131100200202-2120120302031031-1102001033310120-0210313223333230)
- [rules.spec.asn_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3012200013232231-0112112212122230-3221210023330312-2032320202230101-3323213213322333-0230123020101020-1310023203030122-3131322221333020)
- [rules.spec.asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3002221121111303-2331202203130210-1332331020200211-3010112211110012-0032023322323303-2113013110212101-2120023130220301-0322213133032000)
- [rules.spec.bypass_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2133012112012103-3221331232323300-0110230101211323-1203302221310110-3131123221212100-2203223010220212-2130311130310121-2301030133031122)
- [rules.spec.country_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3213003203031101-0230212312221233-0302020132221010-1321321122030321-3020101312332322-0031120212120232-3100021120013123-1321302011333122)
- [rules.spec.custom_rate_limiter](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0211133231322303-0311133231212303-2113002322200333-0232112201021322-3320012111202101-1130031032312020-3112020010310301-2132011331023332)
- [rules.spec.domain_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2303121332011111-1232200221302002-3022032030001130-0230223111223022-3222030310332321-1010210020132012-0203313330323112-2222011312021032)
- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201)
- [rules.spec.http_method](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1003111120111010-1202121130322131-0303110233212223-0131132030100322-1120130333033103-1122133331132333-1312330333111123-3103233022322212)
- [rules.spec.ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2112002322133222-3333220133101103-0102002332002330-2211132331201332-1123022001030310-2011113311303031-1232200100012331-0313020121012011)
- [rules.spec.ip_prefix_list](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0303101313121210-0310111031310101-3220232113223221-0331032013032223-3231300001123323-3020301022000223-0010002203110030-2222030301212202)
- [rules.spec.path](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0121221232200020-2022030220023313-0020220130020131-2132003303231023-3100020030022230-2122103121031331-1112310202213023-1300031300222332)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-1303030310132112-2132322300021121-1130020232311303-2012230022310201-3011001331132232-0222032032302123-2112030232120001-1102121321203333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201210121000332-1322021010213310-3310112032221212-2001202321230030-1021032221100132-0333220031332022-3033031201202213-1132122210321313"></a>

## rules.spec.any_asn — any_asn / 010121012320 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.any_asn

<a id="canonical-2211133200330030-0330010021213213-1331211100121012-1233200320330110-2111332003120001-2133000223003130-1300132000311231-2020331101033112"></a>

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

<a id="canonical-1312023311231033-1331313212303121-1111113313120130-2023332311133302-0320030022211231-3320023003113023-1232031220121132-1112132121101022"></a>

## Direct properties — any_asn / 010121012320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000323312001012-0221211332122003-2130121302222321-0012132022223303-3220021300001101-2322020111312131-2222333023202012-2323231132212110"></a>

## Next pages — any_asn / 010121012320 / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-2032132111200112-0212122320321102-2130210330330223-1000333220333020-0322323310212033-1333002032310300-0303132033112120-1113011331230220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102003331011230-3113211221121130-3133312023233303-0311301233330301-3230012332311020-3121012132103202-3321023011311133-0330130021233013"></a>

## rules.spec.any_country — any_country / 133121330230 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.any_country

<a id="canonical-1211031201233030-2121030111120101-0310302322002123-2233233133322232-3210020210112233-1020330133302232-2113303331120300-3110021322332021"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for any country.

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

<a id="canonical-2002033103332223-0121311002120211-3121201223230312-0100013133021000-2300221331130121-0003023011311333-0032122003123210-0310211123201101"></a>

## Direct properties — any_country / 133121330230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303102012121133-1233221310000213-1111330220130001-0333232212133003-0200103031320002-3302133200333022-0301112123103012-0031231001210230"></a>

## Next pages — any_country / 133121330230 / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-3030300000310132-3013013320203011-3131031020133123-2033121031233301-3222030031320212-1310210133001100-1303123330130002-3322123331231311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021113110123121-0311020102332320-1301031211021223-2332132111231012-1113123030032020-1030123312222230-2100312122120022-0133100301021201"></a>

## rules.spec.any_ip — any_ip / 030222220222 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.any_ip

<a id="canonical-0311102232022310-2010200033233112-2223313033111103-0223022312012110-0213200310122033-2031223301132102-0313110001101130-1322103320033132"></a>

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

<a id="canonical-1313030230221023-1013202023011120-0032100230310311-2123230032220032-3002003321220301-3310110331212211-1201222111331302-2220323001301112"></a>

## Direct properties — any_ip / 030222220222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131011331200310-2001212213010011-0021023210032312-0232322002302313-2300302100321230-1102313001300333-2220303022120122-1301100302023122"></a>

## Next pages — any_ip / 030222220222 / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-0122133122213230-3020232121202333-3110012000201131-0323322111333113-0232131100200202-2120120302031031-1102001033310120-0210313223333230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202103011200013-1232002231332233-3032320313013031-3002310321031001-3202313002210302-1232331313123320-0021101333331231-2313231221332333"></a>

## rules.spec.apply_rate_limiter — apply_rate_limiter / 012221313320 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.apply_rate_limiter

<a id="canonical-2320230212321103-0132032202031100-1323301002111003-2232223332221201-3122011122030023-1320033301111211-3023211020022032-0321210020330321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for apply rate limiter.

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

<a id="canonical-3122221311133321-0113311110212221-3223000113013311-2301303222102220-0221222003112101-2212011311001200-2013330311213011-3001012221023222"></a>

## Direct properties — apply_rate_limiter / 012221313320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010333300021021-3111013033211033-2213312102000111-1333202113130232-3230111323130111-2333301312210212-2103131221012130-2023001320130233"></a>

## Next pages — apply_rate_limiter / 012221313320 / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-3012200013232231-0112112212122230-3221210023330312-2032320202230101-3323213213322333-0230123020101020-1310023203030122-3131322221333020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220132033322010-0332113102103221-3001033130030033-3312312002302301-0213002230223110-1231023133130310-2202303021202313-1113013313320322"></a>

## rules.spec.asn_list — asn_list / 021113312111 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.asn_list

<a id="canonical-3123203023101330-2101111333010212-2003302021230102-0321002023213221-3230013303023123-2030110131030121-2121211200321222-1012220131101223"></a>

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

<a id="canonical-0011223322210023-0122311313313203-1201002133201311-1330021330101322-1213112233222203-2130230230320212-0102202112220321-1110011020233303"></a>

## Direct properties — asn_list / 021113312111 / 3

<a id="canonical-2330020201233102-3002031322030310-2012330332333313-1231121231210332-0201113210023220-2220333113312023-0210201302011113-2012010220032221"></a>

<a id="canonical-2130101310000113-3030221013030010-1033301221200233-1033111212021011-3120112322222323-1333122310302022-1002301102211033-1333301303103012"></a>

## as_numbers property — asn_list / 021113312111 / 4

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

<a id="canonical-1203332020220312-3312000222010311-3231113121023122-0313211122221203-3313213311220202-3100231013302100-2132220032102331-2300123233331323"></a>

## Next pages — asn_list / 021113312111 / 5

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-3002221121111303-2331202203130210-1332331020200211-3010112211110012-0032023322323303-2113013110212101-2120023130220301-0322213133032000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322202203032313-2310322231322110-3213120023332101-3332200303223331-1131001332000122-0331030123023033-2203212120200011-1100110202330033"></a>

## rules.spec.asn_matcher — asn_matcher / 033011232013 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.asn_matcher

<a id="canonical-3311313022111002-0201002131220333-2310111202123213-2111312122332222-2032100033131023-1023011313130001-1100011110001320-3120012310231110"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-1130023002113300-0210102203222122-2221313302101223-1003300033332123-3202223200100322-3210112102102320-3030332233230010-2132320323322031"></a>

## Direct properties — asn_matcher / 033011232013 / 3

- [asn_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1302003110112131-1333311323211120-3221013032223322-2001302023030232-2033213001011123-2223012123201131-3030022133102213-2022010331231001): complete subsection reference.

<a id="canonical-0201232223022201-2131122220230323-2330103120203311-0200202332101210-3301220113201030-2003201222330013-0310330010322132-3213311210003230"></a>

## Next pages — asn_matcher / 033011232013 / 4

- [rules.spec.asn_matcher.asn_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1302003110112131-1333311323211120-3221013032223322-2001302023030232-2033213001011123-2223012123201131-3030022133102213-2022010331231001)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-1302003110112131-1333311323211120-3221013032223322-2001302023030232-2033213001011123-2223012123201131-3030022133102213-2022010331231001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001322323302001-3123233122011233-3311002211031132-0303333302212133-2303123320021110-1112031230023223-0033110131313312-2221032210322311"></a>

## rules.spec.asn_matcher.asn_sets — asn_sets / 210302102221 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3002221121111303-2331202203130210-1332331020200211-3010112211110012-0032023322323303-2113013110212101-2120023130220301-0322213133032000)
- rules.spec.asn_matcher.asn_sets

<a id="canonical-0302311102200001-2200232111320301-1101121223112232-0133302230223021-2110132102012020-1323232100322121-2030333021220332-3331333230331131"></a>

Type: `"list"`. Computed.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-2123333220201301-3010210201320302-0011131210232110-0132102031311012-1032222132003102-2330233322100321-2223301222133211-2130323112030333"></a>

## Direct properties — asn_sets / 210302102221 / 3

<a id="canonical-0223230130330010-0132030331212332-3121222211131230-0011001310321030-2331220021003111-1213331000320000-0230001011221121-2111301012031201"></a>

<a id="canonical-1211213113001113-3013312133031101-0211030201211330-3203232212031203-2010000311223301-2132023130311311-2330201322311121-3120213130002121"></a>

## kind property — asn_sets / 210302102221 / 4

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

<a id="canonical-2300210302103310-1110331203003103-0022320200121002-1123211210122332-1313011112310220-2211110121132120-0021232213230302-2012331102310302"></a>

<a id="canonical-3212302220113303-2111023120321223-0321120332233103-1312310320132123-2110100010031320-0113113301100220-1103100301123220-1212033032310310"></a>

## name property — asn_sets / 210302102221 / 5

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

<a id="canonical-2302113330130003-3001223221230001-2000000010010031-3230111310323311-3031012213323023-3211223023012312-0013033200333103-1120313311101333"></a>

<a id="canonical-1110002111220131-2200311300131232-1102212323333122-3023310122220021-1300030213212100-1002013113202210-0300033303001333-2330223303131321"></a>

## namespace property — asn_sets / 210302102221 / 6

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

<a id="canonical-0300033023030201-1312233000210220-0220332013311202-2022313211012300-0002332013303032-2111223113010322-3121132223022133-2310212210210012"></a>

<a id="canonical-0303213030332322-2111133110101000-2022131011221021-3013202230101301-2233030102331023-2333303210313300-0202333330321322-0302203323320333"></a>

## tenant property — asn_sets / 210302102221 / 7

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

<a id="canonical-3020011133300030-3102133222111120-2233110313002222-3211322010022133-2230302313030111-0301131132111210-3022012122300120-3333002312112123"></a>

<a id="canonical-3320032233100102-0233311233102112-3301232321301212-2233123311201321-0312230000202023-2023112213203230-0111233120012102-3212303123032310"></a>

## uid property — asn_sets / 210302102221 / 8

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

<a id="canonical-0122203101001311-2012131223030123-0001330302011200-0322303302023311-0133000022132113-3131021301321130-3300021030003230-0001113333212122"></a>

## Next pages — asn_sets / 210302102221 / 9

- [rules.spec.asn_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3002221121111303-2331202203130210-1332331020200211-3010112211110012-0032023322323303-2113013110212101-2120023130220301-0322213133032000)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-2133012112012103-3221331232323300-0110230101211323-1203302221310110-3131123221212100-2203223010220212-2130311130310121-2301030133031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032110112111332-3030321133322231-3312302101110332-2113021023133210-2301123023210320-1002013102302322-1103101123200322-3322122221030210"></a>

## rules.spec.bypass_rate_limiter — bypass_rate_limiter / 111020010133 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.bypass_rate_limiter

<a id="canonical-0303310011001331-1301100121302202-0122100231122022-2210330313020110-3220032333332211-3000021221121111-1031230201112032-0230302011330211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for bypass rate limiter.

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

<a id="canonical-2100100111233111-0120021000311030-3000311201203101-2223223120022223-0130110103031322-3231233200300012-0022012311300210-0130023022023312"></a>

## Direct properties — bypass_rate_limiter / 111020010133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120021223322303-3112201131211131-1221323030003010-0113322301130330-2101303213202323-2120121203011201-0123130303221232-1003323322303111"></a>

## Next pages — bypass_rate_limiter / 111020010133 / 4

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-3213003203031101-0230212312221233-0302020132221010-1321321122030321-3020101312332322-0031120212120232-3100021120013123-1321302011333122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011130230011331-2021210121212200-0000031203132103-0013003111312010-2132103021020031-1022103213123111-3010322320211332-2300120112203211"></a>

## rules.spec.country_list — country_list / 232203003232 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.country_list

<a id="canonical-3023331101223023-2100320011100021-1200313201023231-1030021332313330-3000312010121030-2131131120113011-1120112331312332-0300302303233130"></a>

Type: `"single"`. Computed.

Country Codes List. List of Country Codes to match against.

Upstream description:

List of Country Codes to match against.

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

<a id="canonical-2310310023300031-1132013210201322-0230112200203222-3222113103010223-2101322321321330-0310210120301213-1000002032020102-3210302121122330"></a>

## Direct properties — country_list / 232203003232 / 3

<a id="canonical-3321323020202322-0130131311101332-2321010310311101-1133320110212112-2210203333320331-1222111211100333-3023011312313213-2302032310122010"></a>

<a id="canonical-2302231331121233-0201122330122220-3112312223313330-0322021010122031-1231012012312200-2003000200321120-3310303300230121-0013230032332311"></a>

## country_codes property — country_list / 232203003232 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Country Codes List. List of Country Codes. Possible values are \`COUNTRY\_NONE\`, \`COUNTRY\_AD\`,
\`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`, \`COUNTRY\_AI\`, \`COUNTRY\_AL\`,
\`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`, \`COUNTRY\_AQ\`, \`COUNTRY\_AR\`,
\`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`, \`COUNTRY\_AW\`, \`COUNTRY\_AX\`,
\`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`, \`COUNTRY\_BD\`, \`COUNTRY\_BE\`,
\`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`, \`COUNTRY\_BI\`, \`COUNTRY\_BJ\`,
\`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`, \`COUNTRY\_BO\`, \`COUNTRY\_BQ\`,
\`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`, \`COUNTRY\_BV\`, \`COUNTRY\_BW\`,
\`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`, \`COUNTRY\_CC\`, \`COUNTRY\_CD\`,
\`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`, \`COUNTRY\_CI\`, \`COUNTRY\_CK\`,
\`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`, \`COUNTRY\_CO\`, \`COUNTRY\_CR\`,
\`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`, \`COUNTRY\_CW\`, \`COUNTRY\_CX\`,
\`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`, \`COUNTRY\_DJ\`, \`COUNTRY\_DK\`,
\`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`, \`COUNTRY\_EC\`, \`COUNTRY\_EE\`,
\`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`, \`COUNTRY\_ES\`, \`COUNTRY\_ET\`,
\`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`, \`COUNTRY\_FM\`, \`COUNTRY\_FO\`,
\`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`, \`COUNTRY\_GD\`, \`COUNTRY\_GE\`,
\`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`, \`COUNTRY\_GI\`, \`COUNTRY\_GL\`,
\`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`, \`COUNTRY\_GQ\`, \`COUNTRY\_GR\`,
\`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`, \`COUNTRY\_GW\`, \`COUNTRY\_GY\`,
\`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`, \`COUNTRY\_HR\`, \`COUNTRY\_HT\`,
\`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`, \`COUNTRY\_IL\`, \`COUNTRY\_IM\`,
\`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`, \`COUNTRY\_IR\`, \`COUNTRY\_IS\`,
\`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`, \`COUNTRY\_JO\`, \`COUNTRY\_JP\`,
\`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`, \`COUNTRY\_KI\`, \`COUNTRY\_KM\`,
\`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`, \`COUNTRY\_KW\`, \`COUNTRY\_KY\`,
\`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`, \`COUNTRY\_LC\`, \`COUNTRY\_LI\`,
\`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`, \`COUNTRY\_LT\`, \`COUNTRY\_LU\`,
\`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`, \`COUNTRY\_MC\`, \`COUNTRY\_MD\`,
\`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`, \`COUNTRY\_MH\`, \`COUNTRY\_MK\`,
\`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`, \`COUNTRY\_MO\`, \`COUNTRY\_MP\`,
\`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`, \`COUNTRY\_MT\`, \`COUNTRY\_MU\`,
\`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`, \`COUNTRY\_MY\`, \`COUNTRY\_MZ\`,
\`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`, \`COUNTRY\_NF\`, \`COUNTRY\_NG\`,
\`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`, \`COUNTRY\_NP\`, \`COUNTRY\_NR\`,
\`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`, \`COUNTRY\_PA\`, \`COUNTRY\_PE\`,
\`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`, \`COUNTRY\_PK\`, \`COUNTRY\_PL\`,
\`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`, \`COUNTRY\_PS\`, \`COUNTRY\_PT\`,
\`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`, \`COUNTRY\_RE\`, \`COUNTRY\_RO\`,
\`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`, \`COUNTRY\_SA\`, \`COUNTRY\_SB\`,
\`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`, \`COUNTRY\_SG\`, \`COUNTRY\_SH\`,
\`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`, \`COUNTRY\_SL\`, \`COUNTRY\_SM\`,
\`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`, \`COUNTRY\_SS\`, \`COUNTRY\_ST\`,
\`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`, \`COUNTRY\_SZ\`, \`COUNTRY\_TC\`,
\`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`, \`COUNTRY\_TH\`, \`COUNTRY\_TJ\`,
\`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`, \`COUNTRY\_TN\`, \`COUNTRY\_TO\`,
\`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`, \`COUNTRY\_TW\`, \`COUNTRY\_TZ\`,
\`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`, \`COUNTRY\_US\`, \`COUNTRY\_UY\`,
\`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`, \`COUNTRY\_VE\`, \`COUNTRY\_VG\`,
\`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`, \`COUNTRY\_WF\`, \`COUNTRY\_WS\`,
\`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`, \`COUNTRY\_YT\`, \`COUNTRY\_ZA\`,
\`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

List of Country Codes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3323021130223302-2131233001230233-2213231120130010-2031301211123212-1232023222030212-1302201213120122-0100110202222101-0221133013231211"></a>

<a id="canonical-2123103112321311-1002333103302110-2303301012320222-2113323210000303-3112330300330000-2333112202012121-3020200302233101-3212323231101133"></a>

## invert_match property — country_list / 232203003232 / 5

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="canonical-3333310020111111-2031321230003033-1022021221110111-2000030123323211-1233212022132310-2132220102002130-3110211320233001-2002300123002312"></a>

## Next pages — country_list / 232203003232 / 6

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-0211133231322303-0311133231212303-2113002322200333-0232112201021322-3320012111202101-1130031032312020-3112020010310301-2132011331023332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130220233002131-3101131202133133-1110130033333031-2210310213222301-1321302231332022-0332133222302301-2203201031021023-2130002322222110"></a>

## rules.spec.custom_rate_limiter — custom_rate_limiter / 200111032132 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.custom_rate_limiter

<a id="canonical-0113232333221103-2023200102021013-0002230203103123-2220203212001203-2210102101000332-0303110001200021-2120023130231200-3301121200032121"></a>

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

<a id="canonical-3210002333322301-2100113110311220-2120123302103012-3212232222131012-1332230032021012-1300033201132220-0210032221332210-2300233230233031"></a>

## Direct properties — custom_rate_limiter / 200111032132 / 3

<a id="canonical-1203030201133120-0211311202213220-1131100112013302-0001020221110032-3230221002000331-2221230203211230-3230212031002223-1310333010310121"></a>

<a id="canonical-2133302011031310-0313130231121121-3110232012231031-3332211012211301-3121033003010010-3333023320311031-0302323231113100-2232032223301020"></a>

## name property — custom_rate_limiter / 200111032132 / 4

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

<a id="canonical-1211213201030020-0321332220320031-0222120101130120-0332001120122331-0332032221122202-0133201320322110-2021120032020020-0222211131332223"></a>

<a id="canonical-2332331100201101-1212300201111102-3322302023023022-0333222122013012-1012101323223121-1110201021200223-1301010300231313-2020222112011110"></a>

## namespace property — custom_rate_limiter / 200111032132 / 5

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

<a id="canonical-3110100031001023-2110223132332311-1032130111330012-2333211320030030-2131300010122000-0103033130230233-3221023102121101-3002331110311230"></a>

<a id="canonical-2323320333020320-1130233310230100-2030112001033321-1321220120030010-3202033331102000-3231031223010311-1021232002302103-1002312212333012"></a>

## tenant property — custom_rate_limiter / 200111032132 / 6

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

<a id="canonical-0021223032100123-0301211330302301-1232201311201211-0010203031302133-1032232000120222-1023202222020302-1210313210201312-1111010323331131"></a>

## Next pages — custom_rate_limiter / 200111032132 / 7

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-2303121332011111-1232200221302002-3022032030001130-0230223111223022-3222030310332321-1010210020132012-0203313330323112-2222011312021032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233113033022221-1100120110212002-3003030130320123-0130210032030120-0000313012020113-2331033012130311-1103030021212213-3201030201011102"></a>

## rules.spec.domain_matcher — domain_matcher / 221322302012 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.domain_matcher

<a id="canonical-0320022310231323-0001223203333033-2003101211011210-3210031331202300-2201211213312121-1110002102003211-2021210211223330-1020213133200233"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-3312010310231332-2113101332300120-3101012321131220-0001312012103312-0311120321020011-0220003213232033-3112233203023133-2122323231021110"></a>

## Direct properties — domain_matcher / 221322302012 / 3

<a id="canonical-3011323310220102-2010321332203013-2221113111220031-0133221212113002-0001321232023310-3110332231003021-0332133300121230-0000332313331000"></a>

<a id="canonical-0032031110221113-2302131222332131-3222132300031031-1302112113010201-2011230011232121-1320332031001202-1312023321222220-3310132013300210"></a>

## exact_values property — domain_matcher / 221322302012 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1202300203321232-0203302120033333-1110310130010130-3300131100221210-2001202102230311-0322210211010310-3201321332022310-1021332330230101"></a>

<a id="canonical-2302310323313031-2323321221133200-1032030003300330-3123022133221233-1300103120222020-0031101303300031-0322303311300313-0121233103102330"></a>

## regex_values property — domain_matcher / 221322302012 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3302012133332233-1021020011231323-2311203123223112-3122320112023101-3120222321103120-3213212303123011-2111130212221122-1233212011331301"></a>

## Next pages — domain_matcher / 221322302012 / 6

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321112122110230-1320200110212210-3102023122011103-1133023231323001-1122303332331022-1010031302310121-2323100302311001-1113133133232323"></a>

## rules.spec.headers — headers / 120003112030 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.headers

<a id="canonical-1110201011332120-1031102100303120-1003120332030112-3233303212320210-1031212032012231-1203132232120132-3232020132331131-3102133131330301"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
    "minItems": 0,
    "uniqueItems": false
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

<a id="canonical-1332322331102232-2232102032311212-3301012003210010-2303312113212103-3000111003222323-2231331131030110-3213320133331213-3101023130233032"></a>

## Direct properties — headers / 120003112030 / 3

- [check_not_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0330103033213213-2130232330122310-1331131313012223-3303222322131221-0210222202212223-0132200311223030-0203321030321313-1203001221220003): complete subsection reference.

- [check_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2322223113330131-1003313223221113-0112202010101022-0323131233032221-0211312332321231-3001220310310300-1310032231330100-3223300100121012): complete subsection reference.

<a id="canonical-2212201133001200-2010102320001210-3123221303121013-3122233230030211-3323323103121220-2032001310013202-1312101031030201-3300103100312221"></a>

<a id="canonical-1113220132233002-1120002302231303-1331021002111322-0223300210331122-1101022000030233-1233200011101130-0212213033232300-0232323302010000"></a>

## invert_matcher property — headers / 120003112030 / 4

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

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

- [item](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1312221303310012-3123131021330323-1323131123333322-0332313213231310-3100023023103003-2330033201130011-2002201321213233-3032211130231301): complete subsection reference.

<a id="canonical-3110020111321133-2332101021003222-0020121101202022-3021121011233330-0000121121001310-0002313001300303-1023302123113000-2020001030300212"></a>

<a id="canonical-3332031320220223-1010003001203100-3021021102232211-0020331301131002-1033310010302320-0010123232303201-2213312010323022-2233031203211322"></a>

## name property — headers / 120003112030 / 5

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3323202030321223-3330010030321110-2333321010233021-0102223131321131-0030332202021312-2302103010022202-3232222221333332-0203220013211331"></a>

## Next pages — headers / 120003112030 / 6

- [rules.spec.headers.check_not_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0330103033213213-2130232330122310-1331131313012223-3303222322131221-0210222202212223-0132200311223030-0203321030321313-1203001221220003)
- [rules.spec.headers.check_present](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2322223113330131-1003313223221113-0112202010101022-0323131233032221-0211312332321231-3001220310310300-1310032231330100-3223300100121012)
- [rules.spec.headers.item](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1312221303310012-3123131021330323-1323131123333322-0332313213231310-3100023023103003-2330033201130011-2002201321213233-3032211130231301)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-0330103033213213-2130232330122310-1331131313012223-3303222322131221-0210222202212223-0132200311223030-0203321030321313-1203001221220003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221020211303100-2311012102212120-0000302210131103-0110311113223101-3001211013222320-1132000032112130-2320131322132232-2031100321033011"></a>

## rules.spec.headers.check_not_present — check_not_present / 132033003333 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201)
- rules.spec.headers.check_not_present

<a id="canonical-2203033222000222-1232230021132131-3030323020032000-2001021112112331-0201113233120101-3112132320231220-1123202113021103-2212302111031132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-3021121211302101-1223121003101223-0030010102031120-2003033112233001-2201111231023220-1003211121202001-3032103330201130-2301003002312101"></a>

## Direct properties — check_not_present / 132033003333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222222103222301-1312110312010120-2322112023131221-2313232101133331-0320002101232123-2313001222332201-3001030103322022-3113200121112200"></a>

## Next pages — check_not_present / 132033003333 / 4

- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-2322223113330131-1003313223221113-0112202010101022-0323131233032221-0211312332321231-3001220310310300-1310032231330100-3223300100121012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231233303323132-2002032323030322-2302133133111022-0101221101103332-3312322120033233-0312212332112303-1111030321222000-0313303130322110"></a>

## rules.spec.headers.check_present — check_present / 110003230230 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201)
- rules.spec.headers.check_present

<a id="canonical-2011313332301002-0022132321020020-1331032001332011-3011313201111210-0103221132230301-1122311011012132-0210011031303223-1220111000000201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-3002002012300110-3030303100212023-3021322310123223-2301012013202201-2312012022131322-1111202210022311-1100203320300312-3211020220102231"></a>

## Direct properties — check_present / 110003230230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132013123032310-3322201300112311-1010101333310100-2133011111011113-3130022311202133-1103231311232113-2002332212321330-0031030112331101"></a>

## Next pages — check_present / 110003230230 / 4

- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-1312221303310012-3123131021330323-1323131123333322-0332313213231310-3100023023103003-2330033201130011-2002201321213233-3032211130231301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111310002211111-1030103033030133-0102003011131023-2331302310122020-1002132131133222-1221322220132102-2003221110221133-0120023312022011"></a>

## rules.spec.headers.item — item / 301002032333 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201)
- rules.spec.headers.item

<a id="canonical-3121103113031321-1232211232123102-2033322233320003-1123003100030132-0010223320113302-3223323231300132-1002321031132310-2233032320012021"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-0030320313321131-0100210120200120-3233121032122100-3311013013321031-1023033130301120-3230331332100310-3020310022133323-3322011121332320"></a>

## Direct properties — item / 301002032333 / 3

<a id="canonical-1132320221221201-0133133132300033-3231332101111000-0010111101232011-0222100212101032-1300122010133310-0210333031033312-1303333022303013"></a>

<a id="canonical-2100023023201302-3310003331113311-3032033010020332-2202130331010120-2100313111113301-3322312303010233-2112120112021201-2022323013221231"></a>

## exact_values property — item / 301002032333 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1303230320000330-3330131320110110-0211331203303133-3032023121300211-1102133221321102-3020133120122322-3032101333201203-1332003232030102"></a>

<a id="canonical-3310010312301022-0123013003123122-1331212021201222-1223311011131301-3301203031013030-2030002230120220-0012313212130130-0011202333023211"></a>

## regex_values property — item / 301002032333 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0312332112120312-1223010100221311-0220312113303033-1101212121213311-3101203322312302-1101322213130300-2232333312212023-3022210011303011"></a>

<a id="canonical-2000311021223001-3312233123113111-1022010130021212-3333323201320202-3211033211012323-1131102311333333-1223223221332003-3133202131212130"></a>

## transformers property — item / 301002032333 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2331030003133033-0212002333332232-3200212210231303-1311323110103103-3031110200101212-2301332012300232-0211022333200012-2011330032020302"></a>

## Next pages — item / 301002032333 / 7

- [rules.spec.headers](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3201330111112303-0211022230020313-1002132122212211-2223303310202112-2223112301032121-0203020131331203-3130000120102220-0030332211223201)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-1003111120111010-1202121130322131-0303110233212223-0131132030100322-1120130333033103-1122133331132333-1312330333111123-3103233022322212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133013113031030-0331302201232012-2210332102101232-2211112132300122-3210320110030020-0022322133221001-1310132120211001-0213022012202302"></a>

## rules.spec.http_method — http_method / 320113023100 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.http_method

<a id="canonical-0200003020123320-2301212032101023-1300200031330023-1332200021133101-2210110232330103-2331331303110120-0120331113330000-3131002202330222"></a>

Type: `"single"`. Computed.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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

<a id="canonical-0100321012313003-3313113030003011-3213133030011012-3330022300232331-3223201011333233-0121330201322010-2213101232311023-2131210211201022"></a>

## Direct properties — http_method / 320113023100 / 3

<a id="canonical-0113012222313330-3031011203021233-1301303313011201-3022021333212203-1100313302221113-1000302013320101-1200013003210010-2211031301030331"></a>

<a id="canonical-0031312221122322-0132033221101100-1210021003111233-1113113131320110-2112112003103203-1131221212332313-2332120220312130-1302333213332131"></a>

## invert_matcher property — http_method / 320113023100 / 4

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-3132012012212220-0122000122202210-0232333212330323-0112301231001222-2002321010223100-0022333211000223-2000321121212002-1331313313221003"></a>

<a id="canonical-3023300303233032-1212300012111221-0321220113223312-0103010313323011-2221323333012222-0220013122230230-0301133132133310-0133132100002330"></a>

## methods property — http_method / 320113023100 / 5

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3110310010103313-2003110113002300-1130122201331200-0133222132311231-2100233203331113-1122001102122131-3331102022003333-0022230000021012"></a>

## Next pages — http_method / 320113023100 / 6

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-2112002322133222-3333220133101103-0102002332002330-2211132331201332-1123022001030310-2011113311303031-1232200100012331-0313020121012011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330030123123121-2103032321223133-3102131000011331-2133330131101202-3112202223113200-0223210132223120-2312023032212022-1022130000110002"></a>

## rules.spec.ip_matcher — ip_matcher / 030030032323 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.ip_matcher

<a id="canonical-2112020200312101-2213003110012023-3111311002023132-3121300113202311-0012111220031203-0013102203302031-2313320302003223-2033131222312032"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-2030223331101000-2113211121211201-2100312100310033-2333112203323032-1321001020321231-3020023313203032-0301130312132233-3332201333200220"></a>

## Direct properties — ip_matcher / 030030032323 / 3

<a id="canonical-2003031211211033-0023203332310001-3033202300000123-1212333202231032-1110220302022130-1230101321323033-3110200120000231-3303230112331120"></a>

<a id="canonical-0313121123321023-2301200202211333-3233223303120200-0221030331130031-0222011131212330-0012000310133133-3110001320131103-3101231211202031"></a>

## invert_matcher property — ip_matcher / 030030032323 / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0113321110003202-3011310231121301-2103311031313121-3110323323022332-2000321020122212-0121310231202300-1132231101210111-1001212223010022): complete subsection reference.

<a id="canonical-2201222033302002-3333201123211223-1311112110033022-2030332213323202-2222033132212101-0112221323320132-3023202202113122-3102202003223011"></a>

## Next pages — ip_matcher / 030030032323 / 5

- [rules.spec.ip_matcher.prefix_sets](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0113321110003202-3011310231121301-2103311031313121-3110323323022332-2000321020122212-0121310231202300-1132231101210111-1001212223010022)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-0113321110003202-3011310231121301-2103311031313121-3110323323022332-2000321020122212-0121310231202300-1132231101210111-1001212223010022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323220103211131-2201310331123321-1003011233303202-0312323233203203-3110013212211202-1120130023103030-1010113011120230-1302023012303123"></a>

## rules.spec.ip_matcher.prefix_sets — prefix_sets / 222103032311 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2112002322133222-3333220133101103-0102002332002330-2211132331201332-1123022001030310-2011113311303031-1232200100012331-0313020121012011)
- rules.spec.ip_matcher.prefix_sets

<a id="canonical-0330133233312112-3123020202030111-3101220221001021-2300111312210330-3022032020201030-1023013000003133-0220033202223111-0113121311132200"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-0133021123232333-2032120112222120-2032022010212131-3101211020001000-3132010223032332-1232020332220200-3021030202021202-3303021212023112"></a>

## Direct properties — prefix_sets / 222103032311 / 3

<a id="canonical-1222220001131030-2230012211220010-0302010321313112-2032012330021033-1221130221300031-2103012301121011-2030020001122223-2233332101301103"></a>

<a id="canonical-3130031333101031-3020132321231300-0200203322203123-0231323120130330-2220102133212333-3230300302121110-1032230332020330-0332311230322033"></a>

## kind property — prefix_sets / 222103032311 / 4

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

<a id="canonical-2121013003023223-0103101023312013-1213131310032032-2123103332230310-1321232220002233-0333212233232022-0003333123203220-1213110030031033"></a>

<a id="canonical-3111011323312112-0120033111323302-2103113001333332-3033010132232123-1310022333032021-1113123013230213-3121300123032102-1002132213023221"></a>

## name property — prefix_sets / 222103032311 / 5

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

<a id="canonical-3031312210332220-0201233110101012-0212120020201210-0312332221202022-3100111102210233-1133123213223132-1011231311322201-2031310312222122"></a>

<a id="canonical-3331102212300133-1002001200113332-3122000022222230-2020102311220300-0232021322111213-3030211232311112-0321302032202030-2223300022331020"></a>

## namespace property — prefix_sets / 222103032311 / 6

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

<a id="canonical-2230200113111321-0320101021001021-2320232302203300-3020110130101033-2101003231223130-0321220111230231-0211021312130312-1200131132310323"></a>

<a id="canonical-2020001323223103-3022310003302032-1210313303321310-3310201103122102-3203323312033021-3301201002131010-3112012200221202-0001313131131221"></a>

## tenant property — prefix_sets / 222103032311 / 7

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

<a id="canonical-0123112320323112-3010112322101003-2122310313023202-1223321123221331-3233001011323033-3332133030312223-1122110223201021-1332231031030000"></a>

<a id="canonical-1101121111322300-0211322311032032-3233000030331303-0202133122300113-0120133003020121-2313000202023311-3101010021311322-1300303310230110"></a>

## uid property — prefix_sets / 222103032311 / 8

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

<a id="canonical-3213001123311120-0333232320233000-3311330113120202-0331130201331302-1021301013121002-2233120120320130-0033132012123131-0310301221301212"></a>

## Next pages — prefix_sets / 222103032311 / 9

- [rules.spec.ip_matcher](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2112002322133222-3333220133101103-0102002332002330-2211132331201332-1123022001030310-2011113311303031-1232200100012331-0313020121012011)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-0303101313121210-0310111031310101-3220232113223221-0331032013032223-3231300001123323-3020301022000223-0010002203110030-2222030301212202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220300103223303-0202110010100103-1321020310300012-0100020012002303-3212200033212201-3322020111332303-2311012311131201-3321312220201103"></a>

## rules.spec.ip_prefix_list — ip_prefix_list / 022001330223 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.ip_prefix_list

<a id="canonical-2213311021302313-3300311310330122-0212212033231113-3011313121230311-1330011130202110-0202111122213231-1021131013320003-1332131322201200"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-3220022223113111-1320110130021021-1103012001301321-2203011031230012-2100032130203330-3031320012220322-2120113333112310-1012130000200103"></a>

## Direct properties — ip_prefix_list / 022001330223 / 3

<a id="canonical-0021331212232120-0322020101203033-3233210211132120-1013121212221110-1011330021210203-0201322023331013-2013102310132011-0312132311110132"></a>

<a id="canonical-2230120010322021-1113222100031111-2233002232122130-3202201311211322-3003010300112021-3322203011230013-1330322321030323-0110220102310232"></a>

## invert_match property — ip_prefix_list / 022001330223 / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="canonical-3112023220102330-1232321010231000-2221232000223012-3123022231212312-0031210033033022-3202301231231333-3232120133120311-0211301212313011"></a>

<a id="canonical-2220310101230320-1321323330011212-0321122033233000-1213213103310100-3020312232121213-3120132000220203-0203032333312110-3303223332033131"></a>

## ip_prefixes property — ip_prefix_list / 022001330223 / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1133023310111202-0023322111213333-0220100322133112-0002033132231331-1210231330302221-3013230233123202-2330332131233120-3101120232003312"></a>

## Next pages — ip_prefix_list / 022001330223 / 6

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-0121221232200020-2022030220023313-0020220130020131-2132003303231023-3100020030022230-2122103121031331-1112310202213023-1300031300222332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313112213303323-0200031333003021-3031121320333202-3133001310010223-2131022320211112-3133310001310110-2132123303030012-2123320211021201"></a>

## rules.spec.path — path / 002320113232 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.path

<a id="canonical-2311012332211012-1312233330223231-0232101122113212-2010112103311311-0011232311220202-1202331100132312-0231011323102032-3221032221331310"></a>

Type: `"single"`. Computed.

Path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Upstream description:

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

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

<a id="canonical-0031302200303202-3331103311332011-1233003301002233-1310011320133113-0321012101203112-2222333112102200-3320132030001131-1001120102020230"></a>

## Direct properties — path / 002320113232 / 3

<a id="canonical-2133223230030131-0011310113201330-2103132021121100-0113102230031100-0203011000023120-0003122122112131-3311111202110222-2031231201020223"></a>

<a id="canonical-0330113203303231-0101220220300013-1312022310333013-1233120312131132-1312132000131131-1311330301032103-0003212222221010-0011113023300003"></a>

## encoded_path_matcher property — path / 002320113232 / 4

Type: `"bool"`. Computed.

Match against the encoded, escaped path.

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

<a id="canonical-3022021232010123-3200012000330201-3330213000201030-3023333123310203-1032022011133132-0121202321311331-3110332023003112-0222201030202200"></a>

<a id="canonical-1033301333102013-3123121113311233-1033201012022113-0333103230021130-0233002212020113-1030210023022230-0211330233200013-1112330122023313"></a>

## exact_values property — path / 002320113232 / 5

Type: `["list", "string"]`. Computed.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3000321201133310-0320003210201020-3321133233311130-0012112103203331-2021231100022303-0032013203023233-0323303201311112-3112330300200311"></a>

<a id="canonical-1032301013122231-1131132301113000-2303203331213001-2133312311222133-0212020332321320-2012033102020112-2130321323110120-3220023333333123"></a>

## invert_matcher property — path / 002320113232 / 6

Type: `"bool"`. Computed.

Invert Path Matcher. Invert the match result.

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

<a id="canonical-2313231323223100-1132012222233330-2032230033330202-2320203230010213-0111133202331132-2002013302000302-0122003120120310-1331221013213232"></a>

<a id="canonical-1123002010211102-0210030001302221-2111020301110020-1122032123021013-3232032002010313-3331032321100313-3031001112102331-2230210100100231"></a>

## prefix_values property — path / 002320113232 / 7

Type: `["list", "string"]`. Computed.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0101020333233202-1200223311122000-1200000000221202-0101303103311313-0300122013211112-1002223000300230-1210102223100222-2102022013223131"></a>

<a id="canonical-3021213123231133-0002103103310032-3012332033022321-0002022333311300-1212102313301320-2233323223311211-3022320020312100-1012033031211310"></a>

## regex_values property — path / 002320113232 / 8

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0122320213010223-2111300223232000-2200330123013231-2323021230110233-3101210022012301-1210233333101113-3020113032000330-1111123111322033"></a>

<a id="canonical-2110022001210110-1223122303032021-1130023310013012-3231301200313220-2212033020230123-3322132101323001-2302313101101312-3122033200123130"></a>

## suffix_values property — path / 002320113232 / 9

Type: `["list", "string"]`. Computed.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1203001121011300-3120331121102320-0133033133132233-0210303100012233-3113221232231232-1230033010202333-3330101103320220-3102321122113232"></a>

<a id="canonical-2033123122022222-2310020130103111-3300332102030010-3032212001311222-0321323132301302-0013123010333020-1202312123000000-2011322202011301"></a>

## transformers property — path / 002320113232 / 10

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1023201032220222-2010001302000313-2221212121031110-1233132113310100-3231001322103302-0300002203200122-2100113310103310-0211112231102030"></a>

## Next pages — path / 002320113232 / 11

- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021330020133230-3200232232222111-2002000013002213-2312203210133120-0203300210030201-0200010222010023-2200010332002110-2131101212313113"></a>

## rules.spec.segment_policy — segment_policy / 031123222013 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- rules.spec.segment_policy

<a id="canonical-3011030132330212-1220223032003031-1102231212203213-2221300130330003-0201033133213333-0023303303331103-0120113012101203-0133031320131202"></a>

Type: `"single"`. Computed.

Configure source and destination segment for policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

<a id="canonical-0031123102110303-1110122001001322-2012212002010032-3222122022223110-3121322021101213-3132030221010322-0321230113200000-1333331211211020"></a>

## Direct properties — segment_policy / 031123222013 / 3

- [dst_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0310202233111312-2012313311210321-3232330123330112-0231211100201233-3021030132300102-0123030312030221-3112123303220001-2300311222131010): complete subsection reference.

- [dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3001212202030123-2011032303010332-3011112032220233-3222313111213313-0223122011203130-1311231101202311-2200012000010031-1313320010021221): complete subsection reference.

- [intra_segment](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1200322321002333-1011021013103232-1121020122020222-3033213231211233-2211113111301231-0000321201101031-2102003223230000-1331133232112101): complete subsection reference.

- [src_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3022313321232102-3131320001031130-2103021011023233-3111303110111310-1330231100330302-0031320303233300-2321001133302103-0312103233332321): complete subsection reference.

- [src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2020302321132120-1213302111321230-1223112202013001-2011211311300300-2011032200000231-3212331030301211-0300222230200132-1320011130300203): complete subsection reference.

<a id="canonical-0020103300021233-0213213100002230-1033132203113203-1001202122201110-0003200002101333-1112010302033213-1200001211333132-1100032222212313"></a>

## Next pages — segment_policy / 031123222013 / 4

- [rules.spec.segment_policy.dst_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0310202233111312-2012313311210321-3232330123330112-0231211100201233-3021030132300102-0123030312030221-3112123303220001-2300311222131010)
- [rules.spec.segment_policy.dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3001212202030123-2011032303010332-3011112032220233-3222313111213313-0223122011203130-1311231101202311-2200012000010031-1313320010021221)
- [rules.spec.segment_policy.intra_segment](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1200322321002333-1011021013103232-1121020122020222-3033213231211233-2211113111301231-0000321201101031-2102003223230000-1331133232112101)
- [rules.spec.segment_policy.src_any](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3022313321232102-3131320001031130-2103021011023233-3111303110111310-1330231100330302-0031320303233300-2321001133302103-0312103233332321)
- [rules.spec.segment_policy.src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2020302321132120-1213302111321230-1223112202013001-2011211311300300-2011032200000231-3212331030301211-0300222230200132-1320011130300203)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-0310202233111312-2012313311210321-3232330123330112-0231211100201233-3021030132300102-0123030312030221-3112123303220001-2300311222131010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111102231310222-3223112032120311-2222202233012122-0021233322032312-1220132101020230-2031002221121121-1110001130202123-3213000101312203"></a>

## rules.spec.segment_policy.dst_any — dst_any / 113221300301 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- rules.spec.segment_policy.dst_any

<a id="canonical-2121313230303301-2100020332020300-2100213331003033-2010210003111021-2302300333110113-0023301030232033-0102133113322220-2013020201110230"></a>

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

<a id="canonical-0031333023321212-1211110031033312-3012130301222321-3131102302010030-3322202332111003-0311212302212103-1300333111002030-0323201010220212"></a>

## Direct properties — dst_any / 113221300301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012010000221110-3223021301231310-2013120000313222-3330222000230233-2022130300000031-0321001212123200-0303111021022023-3332332102033302"></a>

## Next pages — dst_any / 113221300301 / 4

- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-3001212202030123-2011032303010332-3011112032220233-3222313111213313-0223122011203130-1311231101202311-2200012000010031-1313320010021221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213032123122101-0103110120311103-3230013021311003-1223011000211123-3121122331321233-0022000113321102-1322031303231202-2301321000201030"></a>

## rules.spec.segment_policy.dst_segments — dst_segments / 201321111121 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- rules.spec.segment_policy.dst_segments

<a id="canonical-1031220301022130-2022323332222120-2031203101000301-0330012223022203-3230211012112303-0203233323022323-0112300020311022-0321210201331232"></a>

Type: `"single"`. Computed.

Configuration parameter for dst segments.

Upstream description:

List of references to Segments.

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

<a id="canonical-3032110313003031-1331123131231021-2111111230130112-1330033303133113-1232232322220331-3100300113111110-0321320010130323-0223303022121133"></a>

## Direct properties — dst_segments / 201321111121 / 3

- [segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1210132231003133-2132322321131230-3332111122200023-0213022122212112-2232333133223313-0010122012200102-2132000220012333-1331111132030010): complete subsection reference.

<a id="canonical-0103021320232310-3123031110303212-3332312212331332-3333112230310022-2212321323023221-2021331313321221-0022003332323232-0320222132202111"></a>

## Next pages — dst_segments / 201321111121 / 4

- [rules.spec.segment_policy.dst_segments.segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1210132231003133-2132322321131230-3332111122200023-0213022122212112-2232333133223313-0010122012200102-2132000220012333-1331111132030010)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-1210132231003133-2132322321131230-3332111122200023-0213022122212112-2232333133223313-0010122012200102-2132000220012333-1331111132030010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331232300303311-2310022301333323-3102301120111320-1202333103313100-0023201013000000-3330210230200221-1022023113021032-1320033321232333"></a>

## rules.spec.segment_policy.dst_segments.segments — segments / 211002102210 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- [rules.spec.segment_policy.dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3001212202030123-2011032303010332-3011112032220233-3222313111213313-0223122011203130-1311231101202311-2200012000010031-1313320010021221)
- rules.spec.segment_policy.dst_segments.segments

<a id="canonical-0032010113022333-3331312321133311-3303132213200003-3030032002330220-3320121300012103-0103023033330221-1212001102321333-2201131233320300"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Upstream description:

Select list of segments.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-3001101332020110-3021112303213231-1313113133101123-1031312320222310-3332011321023033-1020101302121012-0212222010230333-0331110332112212"></a>

## Direct properties — segments / 211002102210 / 3

<a id="canonical-1321231020010312-0321313211312300-3101020333322303-3113123323000133-3100233020230200-0130002213120332-2322103033032031-3122013110102233"></a>

<a id="canonical-3201103202021132-3031333221310311-3131010203201233-2010220023300332-3002301212101213-3222012112131020-0003201132230131-3211332022232010"></a>

## name property — segments / 211002102210 / 4

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

<a id="canonical-0133001102100213-3122202322200201-2122213203230003-1013323313201002-2220332023311011-2011230213002222-1301222323031313-2121112101112330"></a>

<a id="canonical-3023233103010221-3110102211311212-3230321312333313-2103220113323011-2322123100133130-2122030221332022-1322203122110010-1021032103121233"></a>

## namespace property — segments / 211002102210 / 5

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

<a id="canonical-3131111110203221-3002213003311212-2113111223220021-3030033113111200-2122320003110003-3013221200300033-1032101301323230-1331010020100230"></a>

<a id="canonical-3011313031311212-3110312232000000-1212120303211300-2122222000322010-2111101131313221-2013301031010233-3132113232312122-0320201120111223"></a>

## tenant property — segments / 211002102210 / 6

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

<a id="canonical-0122000212103013-1313030122301030-2310131000032303-3103331012303233-1032122133100003-3312112013113030-1102230132322212-2223223020030103"></a>

## Next pages — segments / 211002102210 / 7

- [rules.spec.segment_policy.dst_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-3001212202030123-2011032303010332-3011112032220233-3222313111213313-0223122011203130-1311231101202311-2200012000010031-1313320010021221)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-1200322321002333-1011021013103232-1121020122020222-3033213231211233-2211113111301231-0000321201101031-2102003223230000-1331133232112101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220232323012223-0131013013311003-0131310031110113-1300332312111200-2133313001301322-2131332012020122-1002332021210112-2000123232321031"></a>

## rules.spec.segment_policy.intra_segment — intra_segment / 211120000213 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- rules.spec.segment_policy.intra_segment

<a id="canonical-2101220011333001-2021220032012333-0302010222001310-1130323210030022-0011103033020222-3131112111221222-1223301202111130-3111332120321302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for intra segment.

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

<a id="canonical-0332012102231313-3001302333232030-2233310202102022-3332323010312213-0002100222002110-1101101232330102-3011121221203301-0201001320331132"></a>

## Direct properties — intra_segment / 211120000213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223233331301313-2123333320102213-2012322001203232-2300303202033313-1221030113331122-2131310102222020-1330102233122031-3321021301231111"></a>

## Next pages — intra_segment / 211120000213 / 4

- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-3022313321232102-3131320001031130-2103021011023233-3111303110111310-1330231100330302-0031320303233300-2321001133302103-0312103233332321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232301332130113-0032001210023330-1322232221100310-1323103012113012-3332132331332230-2102003130223131-1232101021003310-1031211102021033"></a>

## rules.spec.segment_policy.src_any — src_any / 332330320211 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- rules.spec.segment_policy.src_any

<a id="canonical-2133103221101302-2010022203320323-1002100100311113-3020213212031122-2331200112223301-1313120132323330-2230310320112011-2102021221132102"></a>

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

<a id="canonical-2122133031102312-0022220333312233-2211123130112313-2223310200103012-0212132123333310-2221002310131013-3310102222131311-3121123012210201"></a>

## Direct properties — src_any / 332330320211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210303030303210-3203033230100211-1020200202120300-3121131100331211-2212222022020301-3200211030032112-1203022220210312-0023030003233113"></a>

## Next pages — src_any / 332330320211 / 4

- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-2020302321132120-1213302111321230-1223112202013001-2011211311300300-2011032200000231-3212331030301211-0300222230200132-1320011130300203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110130313300333-3002200331111101-3313323020002011-1012312301111300-1102231313111210-1033130232012013-0022130301223102-2220323311002333"></a>

## rules.spec.segment_policy.src_segments — src_segments / 123120210230 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- rules.spec.segment_policy.src_segments

<a id="canonical-1120331301203123-2330032201311323-0022100213233110-3232013201113322-0221031323220200-1300233000300203-3103031112331133-3111213100013102"></a>

Type: `"single"`. Computed.

Configuration parameter for src segments.

Upstream description:

List of references to Segments.

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

<a id="canonical-0102012233300110-3010320032313100-2323132113222113-2212101020110310-0012322320331222-1202221101102032-2101113022022231-0223202022131323"></a>

## Direct properties — src_segments / 123120210230 / 3

- [segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1210111023320203-2132322302100130-3132022033031320-2223233013130102-1221101013013321-2100021030121331-1123020212131123-1201302012202110): complete subsection reference.

<a id="canonical-0003310031233322-3122033032311101-3132312223230221-3102101222023031-2210323333331032-3311031113101232-1031001300031233-1332131301001022"></a>

## Next pages — src_segments / 123120210230 / 4

- [rules.spec.segment_policy.src_segments.segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1210111023320203-2132322302100130-3132022033031320-2223233013130102-1221101013013321-2100021030121331-1123020212131123-1201302012202110)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-1210111023320203-2132322302100130-3132022033031320-2223233013130102-1221101013013321-2100021030121331-1123020212131123-1201302012202110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130212232121210-3022111013212033-0033010102330130-3333222131021103-0013201103103120-2030332101310332-0223201033000011-1321101311101220"></a>

## rules.spec.segment_policy.src_segments.segments — segments / 130100211132 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [rules](data-sources--rate_limiter_policy--reference--group-001.md#canonical-0230211323233232-3100333221102220-1012303310220132-1222111212001023-3011332112122021-1031133013002122-3300312030222013-3223210323131011)
- [rules.spec](data-sources--rate_limiter_policy--reference--group-001.md#canonical-1011120023321023-0220112121132013-1231303101030002-1321210221302301-3122322032022133-1120311332132033-0331320311321131-1021003120232323)
- [rules.spec.segment_policy](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2201003000230122-3113123312020223-0112001222320021-3330020112121032-2023120200333130-1132222322202230-1221020120201133-0012101223323132)
- [rules.spec.segment_policy.src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2020302321132120-1213302111321230-1223112202013001-2011211311300300-2011032200000231-3212331030301211-0300222230200132-1320011130300203)
- rules.spec.segment_policy.src_segments.segments

<a id="canonical-3003011310222002-1303311302313111-1120312323032113-0101122123002230-1022123203220302-2313201233112312-0201210100013230-2020323032132032"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Upstream description:

Select list of segments.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-1122200122112212-3322223103213332-3232102131322322-2120231001012302-0020201000023213-1120003212132320-1121022120101120-0330213013313201"></a>

## Direct properties — segments / 130100211132 / 3

<a id="canonical-0023031300101121-0230311331213133-2101212303330200-0212032102002313-3311130113101222-2203311311322021-2113232012013100-1002021111003213"></a>

<a id="canonical-2103003200220330-2303311320301221-1322321022113013-1132020310023232-3002020213123032-0111123012233232-3322312232020002-1331120220331032"></a>

## name property — segments / 130100211132 / 4

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

<a id="canonical-1031212202030120-0021221120323223-0123001012131123-3100230010030111-2103120012020112-3310321022133220-3020032031100222-3032123203122001"></a>

<a id="canonical-3313320022023212-2123302313112222-2000313210003120-2002101331103100-3220101113212222-0013233230112310-0100303321202110-0112010220223232"></a>

## namespace property — segments / 130100211132 / 5

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

<a id="canonical-3100221010032130-2110032031120230-3033211011321101-1320230120311232-0031303310310011-2223011221002123-0200110013330210-1210303132223023"></a>

<a id="canonical-3131310132211333-2332001201013000-3202333133213000-2300123123100132-1033200221020303-1220333221133110-0223221210120213-0211112023202110"></a>

## tenant property — segments / 130100211132 / 6

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

<a id="canonical-0130200211023203-0100302213033130-3112033331310300-2211113203202221-1020310301120220-2221023220323121-2231021203002213-3000033321032201"></a>

## Next pages — segments / 130100211132 / 7

- [rules.spec.segment_policy.src_segments](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2020302321132120-1213302111321230-1223112202013001-2011211311300300-2011032200000231-3212331030301211-0300222230200132-1320011130300203)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-3122202032011323-1213021132223201-3223120002010000-0333012013320120-1332310231213133-1210322022033020-3310220211132300-3233322110130012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123223010302310-1313220212103210-0320332120133102-0220311123001211-2332000221111011-0312323122320001-3302122031101222-3021121200210212"></a>

## server_name_matcher — server_name_matcher / 330320300033 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- server_name_matcher

<a id="canonical-1120312101000322-3321303131321311-1322332210112222-3002333130320100-3332030020030203-2301012200300031-0122003311122200-3112032101011210"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-1002032132320013-2113211220333102-0031213202223113-0223132120202312-2330003220021202-0122312003003001-0201200313133313-3101101220131201"></a>

## Direct properties — server_name_matcher / 330320300033 / 3

<a id="canonical-2033020121101210-1220000011001213-3020012332031132-2330203313230322-1131331131013333-1003020202302300-1001230102021310-1030313033032201"></a>

<a id="canonical-2323123301211322-2123203123032313-1331022101021213-0022233113032222-0211132233311300-2333312210232223-3212331311000121-3331222322223100"></a>

## exact_values property — server_name_matcher / 330320300033 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3233023101100012-0130301013223302-0321201303231322-0220132010112230-3010300023103330-2303103131313130-0101003203123312-1102303101013323"></a>

<a id="canonical-3021120013001123-1213310333221021-3220010010300203-2220112320331013-0101130112033022-1323121033313220-0002211013222022-1201110130323223"></a>

## regex_values property — server_name_matcher / 330320300033 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3321321020300011-2333022123321333-3021122232001321-2331032020230331-3310321010313133-1323112210230120-3322210102100323-1212100032302220"></a>

## Next pages — server_name_matcher / 330320300033 / 6

- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)

<a id="canonical-2300221102110132-2231203300303222-2331320002303300-1301313113022220-0110000303303010-0302231130010202-2312132002122201-3033221320120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201013032020313-3232333101333123-1132231330213211-3231222310321103-1211021301123021-3101233211030130-2000012122110210-3032032111220230"></a>

## server_selector — server_selector / 230113200130 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- server_selector

<a id="canonical-1330331320203033-1300323033303021-0323110232023001-3312311332112311-0323100030122312-0320220320131120-1321303322330100-0033321012201122"></a>

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

<a id="canonical-3012211020131220-1211200333020120-0230030113213013-2031022320320020-3203313103003312-3333132331312002-2032321221220131-1220013302110332"></a>

## Direct properties — server_selector / 230113200130 / 3

<a id="canonical-1223312111322011-2030213123122310-3031003211201300-3032030230220221-1221202120010300-0202333012300203-1113322313003013-2103133031300133"></a>

<a id="canonical-3023230332231220-3130022002010213-0222301020233302-0023230303310230-1321121011133311-2222200321131321-1103113330203313-3010030001110213"></a>

## expressions property — server_selector / 230113200130 / 4

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

<a id="canonical-2221211132013212-3223200333331131-1230100002003122-2002122333310032-1313001322320103-0321103031031100-0022013303323023-1033011321020122"></a>

## Next pages — server_selector / 230113200130 / 5

- [Property reference](data-sources--rate_limiter_policy--reference--group-001.md#canonical-2131122200130213-2131230332221110-0110331201121221-3023222311233223-3321323103111233-3333030022113203-1122030331001221-1000002322310330)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-3120113213323201-1201220031110110-1331102203311030-2201122323210330-2220101222123203-3313012030200221-2330311222323110-3113011013233110)
