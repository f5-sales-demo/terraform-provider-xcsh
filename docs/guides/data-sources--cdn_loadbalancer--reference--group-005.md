---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1013102102222130-0211200132123203-2020100332123111-1202103000300323-1132233120132112-1231302012223202-1003113210023001-0112100023202303"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_list — asn_list / 231020322333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-1122332223222331-0102103011002112-0113003012122131-0201200221110100-0332121322321321-0330213233130230-2103133010313121-0132302101200233"></a>

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

<a id="canonical-3231120033220130-3322201123110312-2213020000201000-3121100123323302-3220103131123200-0033301021302032-3010300311012210-2131112031111123"></a>

## Direct properties — asn_list / 231020322333 / 3

<a id="canonical-3201000120301201-1322203010031232-1332210022211110-2023023113302021-3101002101333220-0311322230012111-1331312021022332-2123323330223303"></a>

<a id="canonical-1221201313303131-2031330303012212-1313030202102022-1131300231321221-0210012122131113-0310323120311201-2023330201102102-2321302013321100"></a>

## as_numbers property — asn_list / 231020322333 / 4

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

<a id="canonical-3223223032021100-0231311110030200-1123103331310100-3323233012100103-3311022212012231-2031002031010110-1112303000212132-1100211312330311"></a>

## Next pages — asn_list / 231020322333 / 5

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0133022212232102-1112100110032100-0132032331300111-0012313212321133-2331322300300022-0030113112200233-2020233032030313-1100320203220131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313220031102222-0201021302022323-2213220330330023-0321213012321101-1032321102231003-3223330310233102-3012010203312002-0201102131121301"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher — asn_matcher / 203020132110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-3333133203131103-1031131013301012-0113202331000312-0323120023112230-0022122011201123-3330003202312212-3211311322132201-2011111213003332"></a>

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

<a id="canonical-1313230333230003-2313001102131220-1321333000131033-2210120320310212-0113310303222312-1230131310300211-0330333131212220-2132012130303313"></a>

## Direct properties — asn_matcher / 203020132110 / 3

- [asn_sets](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2220121222022033-3210122112113001-3031102003233303-1012310013103003-2211013302211113-3130313322223303-2200001230030220-2123011101033110): complete subsection reference.

<a id="canonical-3312321301022121-1130310312131332-0020302102123010-1310202112110223-0300230312311111-3232303123330202-0132323212100202-1302202131121330"></a>

## Next pages — asn_matcher / 203020132110 / 4

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2220121222022033-3210122112113001-3031102003233303-1012310013103003-2211013302211113-3130313322223303-2200001230030220-2123011101033110)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2220121222022033-3210122112113001-3031102003233303-1012310013103003-2211013302211113-3130313322223303-2200001230030220-2123011101033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033200330110130-2133130020230030-2023322102100203-2233110212031100-2203300321110201-1103212312210322-1010031020132101-2033032021131331"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 020100102101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0133022212232102-1112100110032100-0132032331300111-0012313212321133-2331322300300022-0030113112200233-2020233032030313-1100320203220131)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-1013030200122331-3022133322232030-1010103131020322-1131110213111200-3312221031022131-2001032331320331-2123233312121202-1102100222101222"></a>

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

<a id="canonical-3002121111303122-1220310331201122-1312111212322111-2020103022321321-3020022030313032-3110230222002311-1210330331300013-1011323330302323"></a>

## Direct properties — asn_sets / 020100102101 / 3

<a id="canonical-2132123033213203-0322213100312110-3031332330221003-0000222033301101-1031221223331333-2113311203022022-1023231123230003-1231323333102103"></a>

<a id="canonical-1201212310332313-1010331213232123-3202103100011311-1021233100110300-1101133331223322-3320331011010321-2212220212110220-1121011012032230"></a>

## kind property — asn_sets / 020100102101 / 4

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

<a id="canonical-1203120311113130-3212033001123200-2030201312231203-3001020013111123-3312221031001102-1110102122132232-1310031111000320-1000002110122032"></a>

<a id="canonical-2310323122132032-3320103333230211-0301313010120113-2000021230231310-0122131013113000-0132323011202313-0231011333232221-0131012132321201"></a>

## name property — asn_sets / 020100102101 / 5

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

<a id="canonical-3122100131032102-0022102003013312-0313023201313320-3131110301332203-2210211002011122-0111133302100113-1221300211132233-0222201033131203"></a>

<a id="canonical-1322200231320003-1030320000120223-1220100333233302-1022301220312321-2120120133233103-3222103301032121-1123230202210330-3323120200103201"></a>

## namespace property — asn_sets / 020100102101 / 6

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

<a id="canonical-2111121131302313-1323101312303303-0312211132300023-3103031220020130-0302310133011000-1332122013030210-3031011010032032-0221122321032012"></a>

<a id="canonical-2103310100123203-1110201100210121-2232311001313113-3122011302332211-1131012012203011-1002311303300003-0130032302310012-0322211221001311"></a>

## tenant property — asn_sets / 020100102101 / 7

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

<a id="canonical-2231332030330333-1012010322022001-3311223233111102-2002022033301313-1100112312032201-3223101001312301-0110300322312003-0022313230313032"></a>

<a id="canonical-0320102020333003-1013312102232102-2312220121102200-1231220200100212-3102212113130112-2221333031001222-2233122200311103-3033000033320022"></a>

## uid property — asn_sets / 020100102101 / 8

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

<a id="canonical-3112002322020020-1322220213103103-0302100112301132-0233231021003220-0332301103202132-0022323111333021-0220002130201021-3213110133100302"></a>

## Next pages — asn_sets / 020100102101 / 9

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0133022212232102-1112100110032100-0132032331300111-0012313212321133-2331322300300022-0030113112200233-2020233032030313-1100320203220131)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1102311331212020-3211200320200222-2332222313232022-1303200130323321-0201321202321212-0130211201023332-0132232331302303-2110212200321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131130111103033-0302103313230013-3031013033020021-2103002011011111-2210212021321022-1111003011201111-1022122013030222-1222030010031200"></a>

## api_rate_limit.server_url_rules.client_matcher.client_selector — client_selector / 002202113312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-1312031113123331-3221301101131122-2202003322201013-1321302202213210-2130302102322331-1331130310312030-1303032022221211-1222110201101212"></a>

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

<a id="canonical-0233220203100130-3100320123023022-1303110321111220-3002020022023322-0212132023122022-1000233322132033-3330332012122303-3322100200013201"></a>

## Direct properties — client_selector / 002202113312 / 3

<a id="canonical-3020321231322333-0123100220222000-1022231001301233-2100132111103101-2220313312303201-3210303321212110-0020002222230312-2300121012213033"></a>

<a id="canonical-2100002023232101-3323121320321122-1220113023032001-0313110122022211-2122020232031031-2202023233010102-1120121030331020-0211333311000101"></a>

## expressions property — client_selector / 002202113312 / 4

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

<a id="canonical-3311333002323223-0023133322233003-2131330110210113-0221310303313320-0331201120232330-2133332032013122-1030231132333221-2200100100222031"></a>

## Next pages — client_selector / 002202113312 / 5

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2232310130330032-0230113113110122-3311201323331201-1123310322203310-3122013132022210-0203031021101102-0303110220332330-1131012111133130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022000333100222-1002130121203333-2221320031120312-0032330033203320-1332012210010223-3102121221202231-1320101211232112-1103212002330003"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher — ip_matcher / 013103313021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-3000310002313233-3001331303011030-3011322130132023-0230102112302013-1011303122122003-2323100121003202-3223202012200123-0120132023121322"></a>

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

<a id="canonical-2330330103130022-1103322033220233-1132211020331212-3321111033301121-0331202123012230-1321120232321020-0003130323112223-0233010222203332"></a>

## Direct properties — ip_matcher / 013103313021 / 3

<a id="canonical-1220200011023323-0332210011023333-1301122312201222-3002210332311133-3233222332123330-3332112013000102-1220132110301013-3302311310011221"></a>

<a id="canonical-3313232133133123-0123212112131021-1000300233301001-2202021322330131-2031132220220223-2131303213121202-2223211211330130-3231312223122231"></a>

## invert_matcher property — ip_matcher / 013103313021 / 4

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

- [prefix_sets](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0332010221101322-0322100113002212-2103310010322103-0222021000202312-3013222033311032-0112203133013130-2010310120112022-0010113011221321): complete subsection reference.

<a id="canonical-0311203113032222-2300130131330131-0320331202221132-3201121003021321-2332332011023133-3013003223332032-3210132103331120-0122102302112311"></a>

## Next pages — ip_matcher / 013103313021 / 5

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0332010221101322-0322100113002212-2103310010322103-0222021000202312-3013222033311032-0112203133013130-2010310120112022-0010113011221321)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0332010221101322-0322100113002212-2103310010322103-0222021000202312-3013222033311032-0112203133013130-2010310120112022-0010113011221321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311322312212211-0330313300110111-1011233333312131-3011030212131310-1211103111000301-3212213323211213-0310302101201030-0310112133211232"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 232133021231 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2232310130330032-0230113113110122-3311201323331201-1123310322203310-3122013132022210-0203031021101102-0303110220332330-1131012111133130)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-1321232022111010-3100033302000012-1301031030003033-2003331011103313-2220132002220222-2330320310333003-2233112203332301-3120120122223321"></a>

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

<a id="canonical-2111310231211133-2031130013310011-1312313310013123-0012321012132132-0323223303011110-3022212220013210-3130310021102321-3312130121031030"></a>

## Direct properties — prefix_sets / 232133021231 / 3

<a id="canonical-3313202002001101-0311303213323301-2132323333323300-1310321322113333-3212010321020032-0021202032332133-2310201023220233-1112003232113201"></a>

<a id="canonical-2010231223302223-2203103031113212-1001132331212333-2200112222111103-1123213112331322-1012011112001233-0001200222012313-0213112123011130"></a>

## kind property — prefix_sets / 232133021231 / 4

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

<a id="canonical-2200333022333201-0233123023302100-0000130230223101-0012023322031333-1223202130220323-3123132201120222-1322132100232321-0201222032311011"></a>

<a id="canonical-0120133020230230-2011203203201003-2332133313311011-2332331023101212-1223333113230200-1320101031310233-0231301210200331-1200312323120321"></a>

## name property — prefix_sets / 232133021231 / 5

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

<a id="canonical-1013233212103223-1110132021220213-0001302111022313-2113203233111030-3202021321023012-1310002222131102-3101321303223301-0121031322233112"></a>

<a id="canonical-0123133013321132-2202231122211222-2212031122130022-1122022230012030-1032100230112203-2231321020121123-0200311301003231-2231011322133312"></a>

## namespace property — prefix_sets / 232133021231 / 6

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

<a id="canonical-3312331300003120-1013213231313033-2021130122100331-0332210001230133-3013011203201221-2132101111133312-2313211121202130-0132010131110030"></a>

<a id="canonical-0102110013001112-3130022111011311-1210201023003111-2230100312322212-0031021322320330-2130212103100010-0312110023112022-3212203210233011"></a>

## tenant property — prefix_sets / 232133021231 / 7

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

<a id="canonical-2012132320313213-0102133131020110-1020030120100030-1310112302033013-0000002130013133-0131311012311223-2210123322011213-0000031022101101"></a>

<a id="canonical-2300022212032121-3000010313103013-0311131103303312-2022311223231000-3221200013000212-1121332231332111-3113330031033333-3022202233103320"></a>

## uid property — prefix_sets / 232133021231 / 8

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

<a id="canonical-1133321332333332-0311300230233030-1232022121132131-0023303120321203-0110112312201201-3132123123133203-2322133113223221-2313232001031221"></a>

## Next pages — prefix_sets / 232133021231 / 9

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2232310130330032-0230113113110122-3311201323331201-1123310322203310-3122013132022210-0203031021101102-0303110220332330-1131012111133130)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1111210112232110-3133000020031222-3110321003330330-2102221330302033-1120133301333103-1120330011032222-1323100133211013-1203121330010323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103010200011033-2303013300303311-3301011332130110-0003030230213130-3203113203330303-3012012031120030-2102130230232103-3012000130112300"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_prefix_list — ip_prefix_list / 310311022003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-2121120311322110-2030300010233133-0031203330123032-2133330102112022-0000101121032132-1031231231222222-2223122113223302-1110322021103130"></a>

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

<a id="canonical-1332010120233221-0102200312133110-2021201030320001-2313311132222202-3132020011022102-1001232332133002-1222330120323200-0232232120201202"></a>

## Direct properties — ip_prefix_list / 310311022003 / 3

<a id="canonical-0010322323020310-0001302200322202-1233033130302121-3223132211300210-0123002333232311-3231003113203000-2310230321131231-1000021331001320"></a>

<a id="canonical-0000233011330321-1032220223101300-2211212223213021-1011222202303112-1323010233120200-1020230300001320-1022311120113321-0232220100311011"></a>

## invert_match property — ip_prefix_list / 310311022003 / 4

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

<a id="canonical-2232302110010022-3313103201102303-3213001131201132-2011200211123323-1012213133332211-0001030030301023-2010213033211030-3311023233210301"></a>

<a id="canonical-0003002120303132-3232113310013202-3221033300222323-3202120022030213-0230122221311030-2201011331213213-1322203113130100-2203232103213030"></a>

## ip_prefixes property — ip_prefix_list / 310311022003 / 5

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

<a id="canonical-3330321201011100-0012133102313000-1121102100230033-3203222133001321-3313200231323220-0220332211212221-3031220300330012-0212012322101221"></a>

## Next pages — ip_prefix_list / 310311022003 / 6

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3122022200133313-0102331311012322-3021212000002223-2000203333200202-1322331230332033-2221130003120212-3032002303233211-2303213201111032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010022221323220-0120103101010130-3013133301010330-0130320021030212-0121002032112332-0121230032211023-2131213111113230-0002132222001202"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 000221032203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-1012022123311200-2120033023303033-3011323210331001-2203231020330221-1022102320231312-0020013020103201-2301312131021323-3231222230000003"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="canonical-2311211321230230-3300213200320121-3202001121111300-2112001102213030-3221233031113202-3312122210211103-0130013221301013-0211303112223330"></a>

## Direct properties — ip_threat_category_list / 000221032203 / 3

<a id="canonical-0320202202323311-2002133230212320-2110333110032230-2130031003332221-1233032233200131-2232321112101211-1001333222103112-0203112032231100"></a>

<a id="canonical-2312103323333002-1010233311200320-0013211130232031-1031213111322200-0000333011012223-2231133313031003-0013303302100212-0100121212232023"></a>

## ip_threat_categories property — ip_threat_category_list / 000221032203 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2312003101013021-1120112332123011-0222101232013331-3210123102033321-1120300201002010-0001310001320020-3300303032332320-0313123113321021"></a>

## Next pages — ip_threat_category_list / 000221032203 / 5

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1331332223312102-1132112313212300-1023223213002222-2333213112000033-3131223010112201-3111213233133211-0202302112202003-2021131013002322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302303231110303-0211303113233320-3322310110201123-1210020132032023-3031032020112200-0320213223030322-1110012220320133-3322100210221010"></a>

## api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 221132230310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-3021132223233021-2333211013310010-0032232021113202-0222320100011201-0001233022332132-3021003332303121-0022033103133130-3121110100122233"></a>

Type: `"single"`. Computed.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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

<a id="canonical-2133030312322023-1130033313223232-2101313212111101-2233231103032320-1131132210111010-2020211332033123-3222023111022222-2132330030332103"></a>

## Direct properties — tls_fingerprint_matcher / 221132230310 / 3

<a id="canonical-0012102330111020-3310233201111112-2332121200331120-1001033001202111-2020303100332321-2333310233301102-1211210222203331-3012103310211103"></a>

<a id="canonical-1303200122112312-0100310300222003-0003120220111111-2010321332003030-3320131012113302-0310130101133221-2310201301211020-1320220323030212"></a>

## classes property — tls_fingerprint_matcher / 221132230310 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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

<a id="canonical-0111021302311313-2331022232331130-0213101223221101-1002133130302203-2201303033200111-0021320031021002-0011210303232221-3110013303301012"></a>

<a id="canonical-2123102012213112-3122021223332320-2331232333011211-1131232233011221-2110330120331320-1331220030200120-2010231011131112-3112220230133323"></a>

## exact_values property — tls_fingerprint_matcher / 221132230310 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2313030013121230-1322123313202333-1020223013211201-3322101001223332-0221212211303111-1203320021001200-1120133221300323-3020223021211100"></a>

<a id="canonical-2213210321010133-1103303230032123-2120033333021302-0012133010320332-3120020231220100-0301000101230233-0203233001313303-3320311233002002"></a>

## excluded_values property — tls_fingerprint_matcher / 221132230310 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1201031220123010-3001110032032121-0032103320323123-0221122311122303-2312231212323030-1303201323101102-2012211032312020-1302323002030323"></a>

## Next pages — tls_fingerprint_matcher / 221132230310 / 7

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011112321130220-1313221312212313-1000313010021322-2232032203321011-0210113103212000-0301332310203221-1313123333330323-2103231230301202"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter — inline_rate_limiter / 213332103003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-0110030131302220-0313212131333122-1233310303300113-3120132122303302-2332102313213003-2132001031302122-2002020220331122-3200312103000031"></a>

Type: `"single"`. Computed.

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

<a id="canonical-1123302101300210-1323021023201103-2133211321311112-3220222113322331-2123230200310111-2031033300011320-3320332330112011-0230103210230311"></a>

## Direct properties — inline_rate_limiter / 213332103003 / 3

- [ref_user_id](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3231133310000120-3211032302203321-1332331220130321-3313122302003213-3012233131133101-3321013331312301-2103321112331321-2310333123031130): complete subsection reference.

<a id="canonical-0231120203220331-1322303033132222-2213220101211312-0321132002223032-2103120320022030-1103223312003322-1013100023103322-1300230213333113"></a>

<a id="canonical-1121013211123320-1312302011320122-0301130110301311-0322333020330033-0233330131132003-2032113102220112-1002013300123003-1003322000221303"></a>

## threshold property — inline_rate_limiter / 213332103003 / 4

Type: `"number"`. Computed.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-1301210211110023-2301300221030333-0213233131201333-1310021033032211-3332100030022300-1321001302121131-2323323200001132-3101011110312333"></a>

<a id="canonical-1121032112210132-0032011010033202-1131000323230110-1033331213022000-1123113033213130-3011211102212211-2221103012132332-1222321332203221"></a>

## unit property — inline_rate_limiter / 213332103003 / 5

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_http_lb_user_id](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3303003110311231-1300100313220132-1033031032233330-1213322120310203-3331231301333213-3121000011121200-2121221232332002-0010232122012112): complete subsection reference.

<a id="canonical-0303300203220001-0311311103003312-1302331321023001-2232021012202121-2333223121313312-2320202212223202-1123023311223202-0203131312001021"></a>

## Next pages — inline_rate_limiter / 213332103003 / 6

- [api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3231133310000120-3211032302203321-1332331220130321-3313122302003213-3012233131133101-3321013331312301-2103321112331321-2310333123031130)
- [api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3303003110311231-1300100313220132-1033031032233330-1213322120310203-3331231301333213-3121000011121200-2121221232332002-0010232122012112)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3231133310000120-3211032302203321-1332331220130321-3313122302003213-3012233131133101-3321013331312301-2103321112331321-2310333123031130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321231023311331-1002112220233010-2100320033223312-1233003021102232-3331333213131333-2001033230220200-2103001010013230-2121031230231132"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id — ref_user_id / 333212030012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-1331020331311213-2020310122021002-3312101201013121-2232130130023321-0321133210320320-3101303202013122-3233311113212121-3223011221021203"></a>

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

<a id="canonical-2323222223031111-3131333302230031-2031301123021321-1233101003332230-3230121101302302-0332022220032200-1123122132130213-1113211133201222"></a>

## Direct properties — ref_user_id / 333212030012 / 3

<a id="canonical-3003310123112230-0232203300133331-0232201131101233-3133022111200231-0203211012002210-3032101001130000-0331022001100323-2312102311002122"></a>

<a id="canonical-2301331111222011-1131133220130313-3021302100231033-3001302011003200-2001030033023302-0032211122300022-0330333011100213-3003201121222213"></a>

## name property — ref_user_id / 333212030012 / 4

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

<a id="canonical-3222122132111213-3000321232320301-3303300320010301-2321320103302211-1202030222303133-3100320230232023-1121320002223200-3131122110331001"></a>

<a id="canonical-3122110311130322-2310222022033323-2013233120032322-2032310021001122-0323012111322003-3130220122000101-2000231132230202-2031011221010310"></a>

## namespace property — ref_user_id / 333212030012 / 5

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

<a id="canonical-0120223123133001-0313030033202331-3110301222320212-1332302032131130-0111300110323322-3211333213323101-3013221121233010-1322131101000122"></a>

<a id="canonical-2123230131010331-0210001230023103-2212323320210331-0031323302101303-1132232020100202-0321220021022012-3133312023133210-3100020230031013"></a>

## tenant property — ref_user_id / 333212030012 / 6

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

<a id="canonical-2213011331231111-2003011303001002-2213103301132121-2130330011033003-1122103020211321-2302020101302020-1000132001333320-2212223212232221"></a>

## Next pages — ref_user_id / 333212030012 / 7

- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3303003110311231-1300100313220132-1033031032233330-1213322120310203-3331231301333213-3121000011121200-2121221232332002-0010232122012112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122233001213020-3312321033232303-3003300312310123-3132111330200211-2021200332213131-1201132121020133-2212100210102213-1231310312132020"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id — use_http_lb_user_id / 023012012133 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-0333322331300213-2011213312211130-2321023211112232-3012310323131123-0203112013321021-0101333023000103-1011113133313012-2133022112121121"></a>

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

<a id="canonical-2210011110200131-2203301320020300-3331302002313220-3202200232201001-1303130201001003-2311121032201232-1120213122113230-3101001223130002"></a>

## Direct properties — use_http_lb_user_id / 023012012133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323222112332233-0322113332230201-2211111113311121-3001122122201103-1320032011130311-3201002211013010-1203303130202233-0220322300202032"></a>

## Next pages — use_http_lb_user_id / 023012012133 / 4

- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1121003122321003-1030113023222001-1120323002113122-1103033302313332-0220011200302103-2120112030110322-1320212132033201-2023201113201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231123221100320-0202210120020220-2322023113313112-2030012202121223-2310132201201011-2122032111003312-0110202222132310-1032301322330133"></a>

## api_rate_limit.server_url_rules.ref_rate_limiter — ref_rate_limiter / 000213331213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-3012030002310130-2100121211230112-0100232010333101-0311103121231301-3202133021203023-0302221322021110-0011130013301213-0221303232010312"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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

<a id="canonical-0303120333023123-2031321033112233-3110320223323033-3330133122303113-1312022332321120-2122221233221222-0012023023003202-1201310020003112"></a>

## Direct properties — ref_rate_limiter / 000213331213 / 3

<a id="canonical-2333203222330323-2332021223301222-1023202003103332-1002023312201120-3112201212103001-3230013221020000-2130212020211003-2331100220020121"></a>

<a id="canonical-2032033102332130-0322023001333103-2330321212221133-0202132300112221-3120312031312122-2320210231132202-3020212131211132-0122102322110022"></a>

## name property — ref_rate_limiter / 000213331213 / 4

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

<a id="canonical-1013021132013231-3210212131112012-0101000013013101-2303103312231330-0222213012213001-1200203321312101-0122222230233120-0210201010301112"></a>

<a id="canonical-3310320020210201-3212200113122013-2331033101111232-0002302102210231-2001200300233100-3101222331110220-1220123321100332-0231002021211301"></a>

## namespace property — ref_rate_limiter / 000213331213 / 5

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

<a id="canonical-2300111310112011-1210102131202130-2013003102102221-3202333100201112-0200231113331212-2302113331130330-3312333331210132-2211022003130333"></a>

<a id="canonical-3210003201022003-2022311312303131-2310112223113103-0100011320130203-3232322321011121-1213113033022230-2332310302330323-2221232131023131"></a>

## tenant property — ref_rate_limiter / 000213331213 / 6

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

<a id="canonical-1003103332002032-1032031102001203-3303222231113003-3102030310112321-0231123310033332-1020221320011312-0010132111023211-3110130133121101"></a>

## Next pages — ref_rate_limiter / 000213331213 / 7

- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232021223100333-1322010032102030-1013112102010130-2101020223101001-2030321313333033-3311312320003100-3302021022100220-0002121210211033"></a>

## api_rate_limit.server_url_rules.request_matcher — request_matcher / 102230013003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-0003303123012330-2303130220210200-2311013220322133-3330120302100130-3133022220211221-1223011031223113-0123330232332230-3020203202310103"></a>

Type: `"single"`. Computed.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

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

<a id="canonical-3003233010210022-1201233312000023-2120021112233203-0121103212010233-3330010033303212-0230001211133321-2210102021001203-1220020102301323"></a>

## Direct properties — request_matcher / 102230013003 / 3

- [cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311): complete subsection reference.

- [headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102): complete subsection reference.

- [jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132): complete subsection reference.

- [query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010): complete subsection reference.

<a id="canonical-2031203113132020-2220123122033110-1300133121023323-3130002130330102-3002231110202123-2221021100113313-1223200212103232-3010201032003013"></a>

## Next pages — request_matcher / 102230013003 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100331302033103-2120123332121111-0002032132110333-1013021332011002-3012213211312311-1100310010033202-3202303112312121-2130123133103030"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers — cookie_matchers / 023212320033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-2321312320331132-3001003321112331-0010210132103201-0303133101112303-2002100311230302-2321112022310102-2000110021100120-3123213201210222"></a>

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

<a id="canonical-1101020302023022-0301120230221331-0210101123001303-2303212130122122-2012203111302220-1113332032331121-0121200322212323-0122022002232330"></a>

## Direct properties — cookie_matchers / 023212320033 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0320322103103020-2313313232322031-0000321023013022-0223111311311020-3122022310110320-3230111122313001-0120203130212333-1032002100332332): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2022310110121332-1001203013102033-3031021310111131-0203233032131331-3322211210121222-0023111130010123-1012012100333230-0201332112220201): complete subsection reference.

<a id="canonical-3130313330330220-2201300321203003-3002003201013121-2003310232311002-3201010001112331-1033211301203122-1210020031132122-3310031100033332"></a>

<a id="canonical-2110301110001020-3013223100332302-3102233322121030-0012203131132130-0121011123003303-2321031030322202-2222232331121033-0323211303330032"></a>

## invert_matcher property — cookie_matchers / 023212320033 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3010001213003122-0330233130333320-2101133003200200-2303102332223321-1033332013330213-3220230323322301-3323023021023003-3231030302330311): complete subsection reference.

<a id="canonical-3333302000133133-2231003101100211-3311200230223033-3110020031211002-0031000203101121-0131311011132130-1331013333233101-3231120002031213"></a>

<a id="canonical-3010300113212311-1100220030133220-1132300331033220-2323000111030321-1323211331112212-2231010130120010-0333311212131302-1132231003110300"></a>

## name property — cookie_matchers / 023212320033 / 5

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1002010031201220-2202000231123303-0210322312230110-0112303211021030-0130231311213011-0103131100100021-3303012212321102-1221110223302102"></a>

## Next pages — cookie_matchers / 023212320033 / 6

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0320322103103020-2313313232322031-0000321023013022-0223111311311020-3122022310110320-3230111122313001-0120203130212333-1032002100332332)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2022310110121332-1001203013102033-3031021310111131-0203233032131331-3322211210121222-0023111130010123-1012012100333230-0201332112220201)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3010001213003122-0330233130333320-2101133003200200-2303102332223321-1033332013330213-3220230323322301-3323023021023003-3231030302330311)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0320322103103020-2313313232322031-0000321023013022-0223111311311020-3122022310110320-3230111122313001-0120203130212333-1032002100332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320130003322322-0113130321223132-0113102033302302-3011021103211320-1011200302000230-3301010333313320-1002321303021033-1123102323013223"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 003110123030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-3301233301232200-1033222231223322-3021332123020112-1313222012221332-0212300110120321-3323210111020230-0232000330012202-1311302232023000"></a>

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

<a id="canonical-2330221011210133-2123022203020222-0030103201002031-2202121023003122-1130220300332011-2232010103110130-0220123201213303-0211323133011332"></a>

## Direct properties — check_not_present / 003110123030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103333013000030-3331212313112021-3200323222110032-2332102331001220-0230223303003313-3123210132013312-2301012131210032-2222213110320233"></a>

## Next pages — check_not_present / 003110123030 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2022310110121332-1001203013102033-3031021310111131-0203233032131331-3322211210121222-0023111130010123-1012012100333230-0201332112220201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013232131231102-3311322101321331-1331031301010320-2220211013300323-1312321031212111-2102130330033221-1212103320312232-3231222110101131"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present — check_present / 101102301203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3311030303330321-2223222301222030-1201111223201101-0002031321231213-0021333132123120-2211310133010313-1122330101222211-1323002221100101"></a>

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

<a id="canonical-3232230300333110-1122311121112231-2001113203123222-3101212212111033-1301020013021130-0231221201030013-1112022033131101-2020110100020310"></a>

## Direct properties — check_present / 101102301203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010023221202321-1233001010132211-1331032003011321-0302002031212110-0113220210120212-1032321301311122-1231233333002000-3112012101302111"></a>

## Next pages — check_present / 101102301203 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3010001213003122-0330233130333320-2101133003200200-2303102332223321-1033332013330213-3220230323322301-3323023021023003-3231030302330311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101030120310000-2113301223203112-1200200031111122-1303310023013020-3030120003012123-3131031030112322-0111333103010332-0213100023133113"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item — item / 300322203103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-2222033032003020-2000032133110232-3330123303122131-2122101313120202-0130301121310031-3220333112012131-3010012300001102-0012122311203332"></a>

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

<a id="canonical-2023330133031203-3223222013202102-0132333033320322-0003032131121200-3330201303032010-2213232332200321-1121312332030332-0323232130200112"></a>

## Direct properties — item / 300322203103 / 3

<a id="canonical-1311220230222101-3013032320131100-3011133230013030-3032101311111332-3220131310233132-3213231110211301-3101311303133022-0300211102201310"></a>

<a id="canonical-3213303031020323-3002213111133031-1302100231312022-0122030301231301-1311130130101003-0200002133201032-1201220022222302-1103032022132203"></a>

## exact_values property — item / 300322203103 / 4

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

<a id="canonical-3020230330202131-1023131111201031-1123011021322003-1120122010032202-1332222213111112-2311132033233233-2112113030121322-3310303202313213"></a>

<a id="canonical-0221321321321102-3203211221121330-2123132312020021-0222333011111021-0021022013130001-3202010002311003-2011101210020232-3200013032112311"></a>

## regex_values property — item / 300322203103 / 5

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

<a id="canonical-3013122031202331-3033111301210112-2122300102102331-2321030230312010-3330012300012312-2200320303103030-0012000223123120-1220220201333333"></a>

<a id="canonical-1200100131311222-2300202023300111-1203313123331310-0312321123200210-0330232013002333-0032220332211000-3003332030311302-3222333112313310"></a>

## transformers property — item / 300322203103 / 6

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

<a id="canonical-2201003013311301-0232002123311000-2010023201020000-2010212102001132-1030310033000131-3203203033131132-3030122333132331-3230222231200300"></a>

## Next pages — item / 300322203103 / 7

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223332033033222-0102300210032022-2102032223033200-3003211302101211-3110022302221213-2312133001133102-0020022201322130-3000031300332311)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313333111123020-0012212310101111-3211321312120100-0120323230210311-1230311313101021-1011331023322000-0111130013311322-0311131101030201"></a>

## api_rate_limit.server_url_rules.request_matcher.headers — headers / 101012320013 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-0002320110110031-2130232230111222-1011030100322310-2031312120211013-2201031100213311-0223220232320133-3123132330102311-2103031020330032"></a>

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

<a id="canonical-3213011033213111-3032030331220132-0313301232220103-0331202202210020-0022303203111222-3011231301201013-1031321212123201-1110220103312313"></a>

## Direct properties — headers / 101012320013 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0010103220103033-0132101131333333-1311001203302210-1021232313303321-2303203110232331-1230322110222202-0032302113102102-3112010220031000): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2023231022000000-1131113332332132-2131233330010221-1013000120011020-2102220131033021-3320023310320102-0103311333021202-0232033103120221): complete subsection reference.

<a id="canonical-0202001211010110-2301231320311323-0311031230201023-1312213112320023-3300001001222011-3011231021010322-3320331133100123-2300000133011123"></a>

<a id="canonical-1302132013022321-3213310032013332-3322232003203102-0310330201303212-1323032111310213-2033323303302130-2020111000222303-0121022001033032"></a>

## invert_matcher property — headers / 101012320013 / 4

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3022310101300212-1322011313030311-0231333020333233-0011201200100310-2123320120310103-1213320233012230-3230112002012200-3123103010320233): complete subsection reference.

<a id="canonical-0001032012023003-1330103010000102-3223331332213202-3100110001321233-1302022120021023-3133203010012111-0102022301223203-0030212111323112"></a>

<a id="canonical-2210303120011120-3311022213322132-1301102013112000-3323102020033012-2333123323302132-1201121320131301-1120202212013211-1312322130112131"></a>

## name property — headers / 101012320013 / 5

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

<a id="canonical-3323020203021210-3213113031233333-3231112013102201-0200313301233211-0212113313310112-3000101111303313-3020200312223301-2200230123122330"></a>

## Next pages — headers / 101012320013 / 6

- [api_rate_limit.server_url_rules.request_matcher.headers.check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0010103220103033-0132101131333333-1311001203302210-1021232313303321-2303203110232331-1230322110222202-0032302113102102-3112010220031000)
- [api_rate_limit.server_url_rules.request_matcher.headers.check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2023231022000000-1131113332332132-2131233330010221-1013000120011020-2102220131033021-3320023310320102-0103311333021202-0232033103120221)
- [api_rate_limit.server_url_rules.request_matcher.headers.item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3022310101300212-1322011313030311-0231333020333233-0011201200100310-2123320120310103-1213320233012230-3230112002012200-3123103010320233)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0010103220103033-0132101131333333-1311001203302210-1021232313303321-2303203110232331-1230322110222202-0032302113102102-3112010220031000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131133210300131-2133220023212121-3330102213000202-0323130100101121-1223032130332111-0211121013012201-0233131220111032-2133022231200313"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_not_present — check_not_present / 021023133030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-3032100301212001-3011123110301203-2312201002311102-0111130011013331-2200031202230010-3233301222013211-3011120012231312-0232332203301132"></a>

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

<a id="canonical-1200302121331231-2021301120101211-1120113123110320-3330100102310231-2221121112310322-0333213230220101-1031101011320123-1031010312100011"></a>

## Direct properties — check_not_present / 021023133030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231222331302012-3012130132010312-0303231200122213-1001023123031331-0023310101213303-3010223331223132-3322223223123230-0211033200222213"></a>

## Next pages — check_not_present / 021023133030 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2023231022000000-1131113332332132-2131233330010221-1013000120011020-2102220131033021-3320023310320102-0103311333021202-0232033103120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113311303002230-0323003210312320-3003033322300021-0100100110312212-3210213220012220-2003333220021330-1332313012213321-3111131021010232"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_present — check_present / 301002221312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-2212232202201331-0330100202300311-2320013001230222-2101022021310321-0232031220231323-1220221332031323-0021022112002211-2002012212321323"></a>

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

<a id="canonical-3123113003031112-3330321330002212-0332121020302020-1102111203022001-2002203321323010-1333123133113333-2010310301332330-1132111031331112"></a>

## Direct properties — check_present / 301002221312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303311010012031-0113300022111201-0010013301331303-3333303301020221-2132312303230213-2322301201212332-0302310033231021-3220200032321312"></a>

## Next pages — check_present / 301002221312 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3022310101300212-1322011313030311-0231333020333233-0011201200100310-2123320120310103-1213320233012230-3230112002012200-3123103010320233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130103303023130-2311332022121332-1100312132321212-2101122032102112-1033220331030303-3020003220023310-1133121122122230-2133230122121313"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.item — item / 231010302321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-0032312021212130-0221113013320233-3333312131223102-0101102231313120-2233322212100003-0112131023203012-3312321012132013-1030233332321031"></a>

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

<a id="canonical-3212011102332031-3120002023033213-2203032100321201-3302211203211120-1320333031021302-0322103233203032-2303023103111110-3012200212231210"></a>

## Direct properties — item / 231010302321 / 3

<a id="canonical-3020300303131202-0311222232121200-0100021333101012-1330312211331211-1122132133101230-2330200103302203-1123200111210200-2002222321322020"></a>

<a id="canonical-1033123332111311-3123321123011303-1030012131013110-0021111023112023-3111030120111203-2202323223010021-3133222301013031-1311330103201300"></a>

## exact_values property — item / 231010302321 / 4

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

<a id="canonical-2323312020111021-3210030131330332-3113322320330122-3332012331000200-2233223312310131-2111232032332311-1113132023230311-2200302103110132"></a>

<a id="canonical-0113011001221000-2121123220113333-1213102023223122-2313233323033301-0203233122222032-0120221230001111-2100202212332221-3001012333200333"></a>

## regex_values property — item / 231010302321 / 5

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

<a id="canonical-2231132202331201-3323013210321110-0132310321221032-3000310103000230-3132323020133022-2303120000210313-3231302203320030-2322123331231312"></a>

<a id="canonical-1011212212113101-1311100101332322-0213331220302230-0012022002330021-2002320230030101-0331100103313013-1233201113313001-1311320311220222"></a>

## transformers property — item / 231010302321 / 6

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

<a id="canonical-0301013323131030-1300202103022312-0322203231230022-0112101323332002-0323231200303022-2130101210111102-0012003233132212-0000130120131111"></a>

## Next pages — item / 231010302321 / 7

- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1223023130120230-3121222311203123-0202202221311321-3330223132220232-2222233332103120-1303223323122023-2112222123232133-3112121212303102)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202302120110111-0121133021220313-0332132103003021-1301221011203021-1321322202332102-2033101122102121-2032001012123031-1031123102132311"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims — jwt_claims / 311123103313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-0311232131121030-3211231300310002-1213211133131233-3302100221201122-1131323131221210-3200303231321210-3332123330103233-2113232232011022"></a>

Type: `"list"`. Computed.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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

<a id="canonical-1102121112101230-3330302333121311-2001223332330030-3132013103223312-2300211122003313-1031111310331133-1220200210010212-2003121130010131"></a>

## Direct properties — jwt_claims / 311123103313 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0333101330332230-2332030123201200-1303333113012133-2023231322030100-1201011231002223-0331323011102222-1212303110132330-3030321311110300): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0200313122001322-1313112000002023-2020202313101330-3223130102323210-1203332201311103-2031031020113210-2231211121301111-1213011330030223): complete subsection reference.

<a id="canonical-1213301222231210-3320022313322121-3230223321330320-2232022222032101-0203030021233112-3213101300020121-3333311030312202-3323333032222220"></a>

<a id="canonical-0030202121311030-3032331232131002-1211221201332311-1002201231330202-3022123021212130-0022110123001310-1102021013331102-1101112033011231"></a>

## invert_matcher property — jwt_claims / 311123103313 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2321232222300023-0132311022213301-2202132022301222-0322011310133320-3301313311210300-0202300020110100-2300303000223323-0101103132332333): complete subsection reference.

<a id="canonical-2321012033222321-1231311113000032-3132302210012233-0320003102102101-0103221032103320-3330213311221213-3010112001333231-3333012203100333"></a>

<a id="canonical-1033102210212211-1011211222301023-1130000013333231-1013210033010233-1202033321211203-3131233121031122-3203232013101210-0012023322113113"></a>

## name property — jwt_claims / 311123103313 / 5

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2210222312103231-2231100031321211-1202110120122130-1200323030230000-1320230130100103-2330203221333333-0011003200030200-3112220312011313"></a>

## Next pages — jwt_claims / 311123103313 / 6

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0333101330332230-2332030123201200-1303333113012133-2023231322030100-1201011231002223-0331323011102222-1212303110132330-3030321311110300)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0200313122001322-1313112000002023-2020202313101330-3223130102323210-1203332201311103-2031031020113210-2231211121301111-1213011330030223)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2321232222300023-0132311022213301-2202132022301222-0322011310133320-3301313311210300-0202300020110100-2300303000223323-0101103132332333)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0333101330332230-2332030123201200-1303333113012133-2023231322030100-1201011231002223-0331323011102222-1212303110132330-3030321311110300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123212331033331-1030011102023230-3322321022332100-2121133021210032-1203321301000322-3302302313031203-3201123311230213-0210022320221223"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 031302120111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3331010132013021-0333320112023002-1121331123203313-1220131011002200-1131331123132021-0020232203110200-1321201230121002-2132032132221000"></a>

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

<a id="canonical-0131110032312313-2313321010320122-3303330233221130-0311223130121133-3023021003133111-3110021123011301-0122011310032111-2332233013212313"></a>

## Direct properties — check_not_present / 031302120111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323123232000300-3023202232122303-2011212303223221-0030113233030103-2122332010133231-1001323310203322-3320102323211133-0023120000331323"></a>

## Next pages — check_not_present / 031302120111 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0200313122001322-1313112000002023-2020202313101330-3223130102323210-1203332201311103-2031031020113210-2231211121301111-1213011330030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120301030001130-0321101213020202-3303130122112022-1221120030303230-1212010002132321-2023010230302010-0300101323033210-0320020011203313"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present — check_present / 222032222323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2210120313000013-0333231223002211-0301121033200230-3120300122330023-1012231331211303-3123120121312120-3123222311120223-3302002120103111"></a>

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

<a id="canonical-1321122300301010-0213032021232103-3002320102100203-3131001222311011-0011113223130023-2131011312122212-3310111131132330-3020111213032311"></a>

## Direct properties — check_present / 222032222323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333003332331101-3301322002022202-2300001012121013-0111023213213301-3111101100313033-2220203211112113-0312130301100112-0022020313301031"></a>

## Next pages — check_present / 222032222323 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2321232222300023-0132311022213301-2202132022301222-0322011310133320-3301313311210300-0202300020110100-2300303000223323-0101103132332333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133233010233333-0112220200012111-0312323102103321-3322313133221023-3033111131320200-2130121332011132-3302321002103311-3200123100330312"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.item — item / 011211133230 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-0031121031022221-0110032202112033-0112331332132300-1311331323131221-2113120203012302-2302030321111102-1100201012013200-1230122002102232"></a>

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

<a id="canonical-0001301222220231-2100202333103033-3211011102030213-3030221203133212-3030100023002121-3232023312211302-1023033021322111-1102213023223001"></a>

## Direct properties — item / 011211133230 / 3

<a id="canonical-3302333201102030-2111322331302110-2121320100333132-0211301330312033-3332230302202322-1113110121301233-0321030320220012-1302220131132303"></a>

<a id="canonical-3201220313313121-0002220203220312-0312033032321223-2123300100221233-1032230032323023-3210031302111121-3100300110010012-3202323132121330"></a>

## exact_values property — item / 011211133230 / 4

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

<a id="canonical-1031323221223213-0302020331222222-3101012200112123-0312231232332121-0211022031122211-0031311011233320-2132223203020302-2333102023232231"></a>

<a id="canonical-3021321322223202-3103022003121312-0302002112113131-0033103202203013-3011322332231303-2102300300113111-2211013111310232-3212131130112232"></a>

## regex_values property — item / 011211133230 / 5

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

<a id="canonical-1212303320033110-0331100333120101-2023233113212030-2020103030233032-0112320301113001-3201331212030322-0330220032130331-3020231031013113"></a>

<a id="canonical-2112122013030122-0012323113300132-1211000131302230-2202211012231320-2211223001212310-0113102311112323-1301020013230301-3232231222321010"></a>

## transformers property — item / 011211133230 / 6

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

<a id="canonical-2120001011210023-2300222103230202-0300223103232002-1331030230110213-2103221330130021-0023103010302103-0222212033011103-1012330300223231"></a>

## Next pages — item / 011211133230 / 7

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3233123020000333-3312331013333322-3123301310211033-2202022112310020-1232022122113132-2212132301013313-2030130222010333-0220113122330132)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230100312120333-2323320222312131-0011200230121230-2121233232301120-2331220323320020-2122132121010002-1202110032101232-2010320300030003"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params — query_params / 112233000013 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-2120332013323121-0001231332031020-2113321023313332-2300013032231102-0232233011110030-3020311002303111-3303200213021131-3110203332132001"></a>

Type: `"list"`. Computed.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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

<a id="canonical-0110131220100301-1010102320023010-3023311313231330-0220301230300131-1111332101021110-3223030302221020-0200111313301011-0213332123330030"></a>

## Direct properties — query_params / 112233000013 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3321231032133332-2332011032203202-1012330112120330-0122123202131132-0121022230003033-0312212130223213-0111002303001212-0310022131300332): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1311122023131330-2222100033212110-0313000012330101-0032333320201000-0023230122102111-3330130030323001-3023020010320130-3312121303331001): complete subsection reference.

<a id="canonical-0123031031032131-2123331122200230-0100031113201020-0323333120300001-0332133012111123-2120132213123200-2031211200132033-3211010123213110"></a>

<a id="canonical-2330102203011132-1222023211120110-3011331332201001-3211231021322033-0112022113012300-2001113223002321-2310323121312030-2213103222020310"></a>

## invert_matcher property — query_params / 112233000013 / 4

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2032112230003320-1101112311321321-0231130232220100-0310213022230312-0120022103302100-1032320313031302-0113333120322133-0311121311123213): complete subsection reference.

<a id="canonical-3031022111130221-3300100001330000-2132300121203021-0132230201301200-2112332311111320-2330312020031303-3110023123331310-1333023323030122"></a>

<a id="canonical-2020112022110100-1011200030003010-1013122133012233-0310303331011310-0102121230221123-1030031013021111-1122003021023132-0121001012202223"></a>

## key property — query_params / 112233000013 / 5

Type: `"string"`. Computed.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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

<a id="canonical-1133023103200320-0233121113121221-1323122310111331-3313033212031110-0121032212313331-0302101032300310-3220231030113330-3110222001330220"></a>

## Next pages — query_params / 112233000013 / 6

- [api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3321231032133332-2332011032203202-1012330112120330-0122123202131132-0121022230003033-0312212130223213-0111002303001212-0310022131300332)
- [api_rate_limit.server_url_rules.request_matcher.query_params.check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1311122023131330-2222100033212110-0313000012330101-0032333320201000-0023230122102111-3330130030323001-3023020010320130-3312121303331001)
- [api_rate_limit.server_url_rules.request_matcher.query_params.item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2032112230003320-1101112311321321-0231130232220100-0310213022230312-0120022103302100-1032320313031302-0113333120322133-0311121311123213)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3321231032133332-2332011032203202-1012330112120330-0122123202131132-0121022230003033-0312212130223213-0111002303001212-0310022131300332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331030221123131-0210113103310000-0012033220313101-0312131112223330-2123211133300332-1001213013002100-3313232030221110-3002033001231112"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present — check_not_present / 323200202120 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-2320000232203100-1223203022203202-1312221012333333-0023130233012101-0023321001111311-1122330021120132-0221222010112233-3231210000103302"></a>

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

<a id="canonical-3303120323103133-3112311121131333-1001130123320313-0130011213032021-0130230023221321-1032303103132332-3203111312011012-0032303120321132"></a>

## Direct properties — check_not_present / 323200202120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320012332020023-2110302322012220-0233211311110330-1202230332213022-2122012332132003-3001100311130003-0222311331112303-1031013222130311"></a>

## Next pages — check_not_present / 323200202120 / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1311122023131330-2222100033212110-0313000012330101-0032333320201000-0023230122102111-3330130030323001-3023020010320130-3312121303331001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210131101211033-1131023100333233-2313131020123313-3011010300122102-1303320321322231-1121321233013033-2002332121101103-1112210022110301"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.check_present — check_present / 030211121113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-2211000221001310-2101013132123312-2022302001032010-0210103133103031-2002231110332303-0031003011312011-0100220010331213-1131113132001101"></a>

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

<a id="canonical-1202013332101321-2001310031323233-0323230132100001-1130210233210101-0120111113001313-0000221213110022-0303322120203310-1332233213201033"></a>

## Direct properties — check_present / 030211121113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222312201212232-2102223020000023-2013121023332213-2232233031130131-2021221221130101-3122032101321312-3030110033102221-3031020120133222"></a>

## Next pages — check_present / 030211121113 / 4

- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2032112230003320-1101112311321321-0231130232220100-0310213022230312-0120022103302100-1032320313031302-0113333120322133-0311121311123213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021020302032110-2002002012323323-3020332202231200-1202210331122013-3301220320331032-0312020030102330-2002131322232002-1022031212112203"></a>

## api_rate_limit.server_url_rules.request_matcher.query_params.item — item / 003232323330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-2032131303311331-0101023203120001-3213001101113033-1320220003001220-0211233031111220-2001101231001110-1303103020110102-0310310131322230"></a>

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

<a id="canonical-1002323312200020-1323210031322013-2210130022221202-1221103210321003-2222111331012112-1310213122310020-3201202220120003-0201201112300322"></a>

## Direct properties — item / 003232323330 / 3

<a id="canonical-2201200112321223-0200222321303101-2311112023301103-0130033031032320-1233333030111031-2120132121231031-0030331312101133-3211213133013013"></a>

<a id="canonical-0021232212112323-2030023100233102-0312022120210201-3301222231100231-2132101323120030-1013233011123133-0332302203110333-3323111210001310"></a>

## exact_values property — item / 003232323330 / 4

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

<a id="canonical-0311121313001023-3220232303133032-1203023301100131-2110100211321021-2211120010213230-0021032112131013-3000110013302120-3303030320131101"></a>

<a id="canonical-3032120332322130-2212210313300001-2313302100212100-3213301101121203-3132320300121230-2111320231133303-1120221032030113-2201223322232333"></a>

## regex_values property — item / 003232323330 / 5

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

<a id="canonical-3233031002022320-0213001221202202-2011233313312201-1322020313100113-0033321031103102-3201103020011110-3023131222002131-0100003011133330"></a>

<a id="canonical-2223002222123332-0321110322301000-2030132222221022-1120031311201320-1112121022231202-1113331323021310-1223111331123213-0123313223212201"></a>

## transformers property — item / 003232323330 / 6

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

<a id="canonical-0102120303310231-0000110231113000-3213031113002331-3033311000311013-3210033030121000-1031222323223312-3113223303023130-0301100212012320"></a>

## Next pages — item / 003232323330 / 7

- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1123032003112112-3231033313121110-1313003001311222-1322232102232332-0003312203221233-2001323102303132-3332230100003002-1112001201232010)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022010321232222-2122020123131112-0202230323021112-3100033112033232-1102311110231003-1033123112212023-1100333010210130-0220132223111232"></a>

## api_specification — api_specification / 021232111302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- api_specification

<a id="canonical-0103332123202233-1321002331003032-0001313121330030-0110313123121212-2312131130012102-2302113002030233-2202330001303202-3131131301300132"></a>

Type: `"single"`. Computed.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Upstream description:

Settings for API specification (API definition, OpenAPI validation, etc.)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0103332123202233-1321002331003032-0001313121330030-0110313123121212-2312131130012102-2302113002030233-2202330001303202-3131131301300132)
- [disable_api_definition](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0123123210010032-2000110022103021-1311000120001310-2022020021332330-3303230013303001-1302033001330103-1133323112122130-2022201130112011)

Select alternatives according to the provider validators above.

<a id="canonical-3221321300103010-0020100211133332-0300031011300333-1211313002112332-2211122230211211-0202323330003023-1332320011223023-3003130033233110"></a>

## Direct properties — api_specification / 021232111302 / 3

- [api_definition](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0221212311312323-3110130002001100-0321312322230103-1130303111331232-2120030223003322-0011020222121321-0113022030133120-0320312233133320): complete subsection reference.

- [validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001): complete subsection reference.

- [validation_custom_list](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000): complete subsection reference.

- [validation_disabled](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-2200322100312233-1331332231131222-2201003130302321-0312322000110013-0131003103300102-0103311103220223-2320201221023111-2120100012312210): complete subsection reference.

<a id="canonical-1130031011002111-1330222013130131-3332311301322110-0310200101032130-0302202010223123-2313202102330310-0023011212012121-3110112322231221"></a>

## Next pages — api_specification / 021232111302 / 4

- [api_specification.api_definition](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0221212311312323-3110130002001100-0321312322230103-1130303111331232-2120030223003322-0011020222121321-0113022030133120-0320312233133320)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_disabled](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-2200322100312233-1331332231131222-2201003130302321-0312322000110013-0131003103300102-0103311103220223-2320201221023111-2120100012312210)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0221212311312323-3110130002001100-0321312322230103-1130303111331232-2120030223003322-0011020222121321-0113022030133120-0320312233133320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130302312233210-2222020023013332-0033310112200032-2200310220233300-0032303113120102-3230100112230321-2201330001112321-3120012330230023"></a>

## api_specification.api_definition — api_definition / 012311220100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- api_specification.api_definition

<a id="canonical-2210302210310001-1212112313200131-3012330201333320-0100310110301230-3033301332320232-3203330132020010-0111231210022311-2223332212320110"></a>

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

<a id="canonical-3213121330222113-2120023121233200-3223321320320033-0131313221200101-2201230303123231-3100113303203030-0031003211313232-0312103021323320"></a>

## Direct properties — api_definition / 012311220100 / 3

<a id="canonical-2330002002200110-1322022313322010-2121123101221102-3122020103212230-1033221021130223-2233133211123032-0010203320103012-3012102330013313"></a>

<a id="canonical-3330302002212212-2101311023332110-1001011311001131-3101012233320301-0010102200021100-2120213220102013-2032030233022323-3132123012312100"></a>

## name property — api_definition / 012311220100 / 4

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

<a id="canonical-2233222100310132-0221212023202002-0032031200121222-0210111120200201-2313000033233012-1302231000210320-2010300230221200-3210231300212023"></a>

<a id="canonical-1230001102032300-3011221110232110-1003300321130132-2320102222331102-1003310230332311-0100111112131033-3311112121000033-3231200013333212"></a>

## namespace property — api_definition / 012311220100 / 5

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

<a id="canonical-0032300300312233-0301021302333131-2220330102223131-3001132211212012-3113220030030213-0000211110133312-2332112321021122-0211332002231022"></a>

<a id="canonical-3322333302310011-0303013033120332-0013232332001211-0203303331313003-2120220032111212-0000023111211210-1223320033103303-3002213022323303"></a>

## tenant property — api_definition / 012311220100 / 6

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

<a id="canonical-1322331020332232-0311222011030121-1032220222102213-0201311220122323-1303313101012232-3201023112031200-1211211031232123-2000223120132121"></a>

## Next pages — api_definition / 012311220100 / 7

- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032011210011200-0022122113203030-3232322001231000-2010301323000222-1122223303303220-0232100102103312-1230201033213111-0012111300102032"></a>

## api_specification.validation_all_spec_endpoints — validation_all_spec_endpoints / 323101100312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- api_specification.validation_all_spec_endpoints

<a id="canonical-3001220010203320-1231332202120020-0133002130332030-2100121111302310-3210311313223123-3103313130233203-2132111013100133-0332121113013322"></a>

Type: `"single"`. Computed.

API Inventory. Settings for API Inventory validation.

Upstream description:

Settings for API Inventory validation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

<a id="canonical-2113201020112122-0321033122021233-3330202032313121-1331123311113323-0302230122202030-1310031203101312-2231210110000202-2031122110201123"></a>

## Direct properties — validation_all_spec_endpoints / 323101100312 / 3

- [fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120): complete subsection reference.

- [settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112): complete subsection reference.

- [validation_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221): complete subsection reference.

<a id="canonical-1101221003032330-3112132031221113-3110211221103230-1032321321103302-1303111100311201-3000330111313003-0103331121021210-3330211023321122"></a>

## Next pages — validation_all_spec_endpoints / 323101100312 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110010302033232-0301021221111112-3022012212032111-3330322113321100-3012102000132210-3133011131032331-3322200322011200-0112023232031211"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode — fall_through_mode / 303100300300 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-1301121333222130-3330221211230021-2211213103210000-2031031001130132-2301201331132023-3101133333301231-0020100220022031-2300122203033313"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

<a id="canonical-0012021301032202-1101203100233212-1302123210231010-2100022312311011-2131220013322321-3021001033231230-1020030232310133-2021010131232022"></a>

## Direct properties — fall_through_mode / 303100300300 / 3

- [fall_through_mode_allow](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1202011201033211-3121333123003020-3201332032202001-2321101021322012-3312111111323101-2002003022332203-1202202103321031-0301113233233222): complete subsection reference.

- [fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332): complete subsection reference.

<a id="canonical-2022111320323331-2311131031100101-0020331312330332-1201102002132202-1212323123031012-3011003021310030-0012332112103222-2332333103321223"></a>

## Next pages — fall_through_mode / 303100300300 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1202011201033211-3121333123003020-3201332032202001-2321101021322012-3312111111323101-2002003022332203-1202202103321031-0301113233233222)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1202011201033211-3121333123003020-3201332032202001-2321101021322012-3312111111323101-2002003022332203-1202202103321031-0301113233233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320021223323201-0031233001103001-3031122301033303-2011101220002200-3022021232203213-0122231133231122-3312130302303233-0313110110322032"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow — fall_through_mode_allow / 223013010121 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-0013132022213022-2222233103210031-2330133311111003-2233320212211013-3333020313133211-0033000033000330-1010030012101002-3221232200111130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-3213102230213001-1123102220012110-1123312101201120-2020023202211032-1121103321032202-2003031103312320-3311301221121333-0211003003112323"></a>

## Direct properties — fall_through_mode_allow / 223013010121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213000230013123-3112103013131031-1211211300001303-1002312020120332-0021212302122333-1013021203212222-3132232001110310-2332231221232313"></a>

## Next pages — fall_through_mode_allow / 223013010121 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031113032112321-0132220332021101-0322130001020313-2203223312302030-2300000220320030-2300203221201301-0113003123101123-1213030121233333"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom — fall_through_mode_custom / 131300211222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-1333202002310022-0033330331321002-3200100021133303-2300022300200001-1211333132212001-2332003120112103-2000011033302321-0123311023213103"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

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

<a id="canonical-2102103102203331-0333310233333032-1313000033123032-2002332200312301-3212230001311111-2320301033030232-3133320312121001-2303122321321231"></a>

## Direct properties — fall_through_mode_custom / 131300211222 / 3

- [open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210): complete subsection reference.

<a id="canonical-3323330102321220-2011120123212221-0321223003030113-0222111002102322-1202020223020211-1233013012332130-1303302201312110-3303023103312032"></a>

## Next pages — fall_through_mode_custom / 131300211222 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211012312122312-1333133112322302-0301320110110110-1223302122303033-2112230030100311-2320010212233132-1113002320100321-1300010211220330"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — open_api_validation_rules / 212231302012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-3000032313023320-3200302331002010-0201103232102211-3010311211020222-0213030310123200-1131031301013221-0212312202021223-2302022120011230"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2022223012111311-0303022201000001-1120033013020020-1322033323322021-3321001330001031-0211032230033010-0331300231222221-2102013322301003"></a>

## Direct properties — open_api_validation_rules / 212231302012 / 3

- [action_block](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2131223020301322-3003100021032330-1210101322232112-2113130012123200-2221213000203223-0303311110301331-3022022323312133-3132102301332331): complete subsection reference.

- [action_report](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0311030022230303-2110230221032132-0332031000200302-0201303011301232-1023000300322031-0310002101013313-3233013001313112-1122003202013102): complete subsection reference.

- [action_skip](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3203133033020130-2001121020123320-3332012003003030-2212112013111102-3121301133003313-2321233232221020-3012313221112133-1032013022202113): complete subsection reference.

- [api_endpoint](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3003231311130233-0223001212230013-0212113331312032-3013112131111230-2003131012233303-0202122333013000-1213303012102100-1023210112102301): complete subsection reference.

<a id="canonical-0333102001201213-3121323232033322-2321023220031320-0300312132000202-2211232303102210-3303333311213033-3131230230030331-2101330311023230"></a>

<a id="canonical-2032103033023311-3023013231331033-3221110212313033-2003331330000101-0130303202101103-1203020130221211-1122032130103333-3323313220203330"></a>

## api_group property — open_api_validation_rules / 212231302012 / 4

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1031231313332300-0021230131031003-2100312010111321-0113223113033301-0231023303003233-2010111332032231-2232222102003130-0020302312020100"></a>

<a id="canonical-0203030202003211-0130202221331302-1132130121313332-0230201233233011-3302303130310300-0333103330020123-2022110323333212-2002003113013033"></a>

## base_path property — open_api_validation_rules / 212231302012 / 5

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0012222112331100-0303231311123222-2221002000301331-3133032301313022-3303300021003313-0332320133322232-0311001301010103-3231133003323001): complete subsection reference.

<a id="canonical-1223210031210111-2102301211232213-1100210222210110-2100223100132331-2312103031232102-0200021003331310-2033130331113332-1313232302012202"></a>

## Next pages — open_api_validation_rules / 212231302012 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2131223020301322-3003100021032330-1210101322232112-2113130012123200-2221213000203223-0303311110301331-3022022323312133-3132102301332331)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0311030022230303-2110230221032132-0332031000200302-0201303011301232-1023000300322031-0310002101013313-3233013001313112-1122003202013102)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3203133033020130-2001121020123320-3332012003003030-2212112013111102-3121301133003313-2321233232221020-3012313221112133-1032013022202113)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3003231311130233-0223001212230013-0212113331312032-3013112131111230-2003131012233303-0202122333013000-1213303012102100-1023210112102301)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0012222112331100-0303231311123222-2221002000301331-3133032301313022-3303300021003313-0332320133322232-0311001301010103-3231133003323001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2131223020301322-3003100021032330-1210101322232112-2113130012123200-2221213000203223-0303311110301331-3022022323312133-3132102301332331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032310303102301-3131112022021200-2203331113111303-3211132220313133-2330102231332030-2130323330001220-2222230203200032-0333312220023021"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — action_block / 023133032300 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-1011301033030331-0030313301312313-1320211300303221-1121332223101020-0000200233133302-2210332303211322-0011300300312110-2021300010011012"></a>

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

<a id="canonical-3123203121213213-1033021130101033-3022122311330221-3000313020333123-2020333312022310-0230102303013333-2033132131312121-3220012323211102"></a>

## Direct properties — action_block / 023133032300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123022203302110-0002031311301032-3203211221020330-0301001132201001-2211203103000321-2121233200321310-2210301213310030-3020130121032302"></a>

## Next pages — action_block / 023133032300 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0311030022230303-2110230221032132-0332031000200302-0201303011301232-1023000300322031-0310002101013313-3233013001313112-1122003202013102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012021311123022-0112010201313122-0003320021022103-1001113233000333-3013111032032130-0332230210032010-2111301303320333-3033320102200130"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — action_report / 330003312232 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-2320332220300133-2022112102101312-3321101133133123-2102313230033313-0203132222010201-1202200100122031-3003022121002010-3110300320200032"></a>

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

<a id="canonical-3311023201202233-0233310032101211-2322210203003121-3303331230012102-3033220301130311-0222120131230222-1100011123333022-2330213133332301"></a>

## Direct properties — action_report / 330003312232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021011310322322-1102132121110111-2022111330332303-2333233331231002-3131301221131123-0233310333133320-3213321110213230-3233223202111231"></a>

## Next pages — action_report / 330003312232 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3203133033020130-2001121020123320-3332012003003030-2212112013111102-3121301133003313-2321233232221020-3012313221112133-1032013022202113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101020211022201-3231113323320000-3130220223221011-1102331321200230-1312222132110010-1010132213331113-0101030010311330-2113213301121012"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — action_skip / 213120302323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-0112022101130002-2003210300233310-2231202030201101-1103102332000002-1132113330132033-3231330033123323-0121111330032030-0223313211020223"></a>

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

<a id="canonical-2203100220001103-3302132232200200-0202131011011220-2001121210310033-3313212030131112-0021130131132101-3311122222120000-0122021001002223"></a>

## Direct properties — action_skip / 213120302323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131100021021311-0300310223030210-0030100310120122-0320312220232200-2033303103022210-0113302030312021-3022331130213020-2120101012010123"></a>

## Next pages — action_skip / 213120302323 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3003231311130233-0223001212230013-0212113331312032-3013112131111230-2003131012233303-0202122333013000-1213303012102100-1023210112102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010131213122012-3233011011200323-1232112130132103-1111020211332022-1102213120210321-3020323230333300-0200030322012011-1101110333303202"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_endpoint / 200211111333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-2231010133121310-1213202332330220-3131313012012322-2102330131213232-3033103032012002-2021120000022022-3220122031001301-1010323010030002"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

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

<a id="canonical-0230221133131230-1101310033101100-2312000233123303-1203211201223103-2220103310011332-2323321232300202-1100133221111210-0233003312022333"></a>

## Direct properties — api_endpoint / 200211111333 / 3

<a id="canonical-0320303303011133-3223022321011132-0220030233310103-3323100302200013-3011031202222223-1330102022221120-1031302220202122-3011013320020212"></a>

<a id="canonical-3322113100033201-3102020232110230-0110322203110201-0200231313303003-0000310100311122-2013223302330312-0011222213012210-2000021031203132"></a>

## methods property — api_endpoint / 200211111333 / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

<a id="canonical-0100012310323110-2033032033332021-2110201022021010-2001320022330222-3212223023030013-0313122201210320-2321331202000031-0300220031100032"></a>

<a id="canonical-3300113131031000-1003230310131203-0002312101012213-2110303311233220-1222100012301321-3013202133010011-0021023222021300-3300111303100320"></a>

## path property — api_endpoint / 200211111333 / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-0121022211033012-1021123013100311-3032202323111311-1213332220311210-3231301333211200-2212203133301000-3003301313320132-0322310001112103"></a>

## Next pages — api_endpoint / 200211111333 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0012222112331100-0303231311123222-2221002000301331-3133032301313022-3303300021003313-0332320133322232-0311001301010103-3231133003323001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202230030332003-2322203210132301-0320320133011122-1210031302121102-3323213302212122-2031313031032302-3002223210303101-2133301213133121"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — metadata / 103112230133 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-1210120100332311-1102322230101003-0323332011000321-3211100001212023-2202023223321033-2013333123111231-1220303010020203-1313211133132111"></a>

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

<a id="canonical-3032223012123233-3123230220333223-2113200001023202-0201121211033220-0213132112011231-1020001231131312-2000221002322311-0323312313312330"></a>

## Direct properties — metadata / 103112230133 / 3

<a id="canonical-2323033323300131-3211213022333330-3120313031103010-1223022011312211-2332310123322322-1220031333110132-0302031012002131-3102310311032211"></a>

<a id="canonical-1020121012000011-0131210102000032-2333313122013100-2133111130231201-3300211303100331-3222330303330302-1112121101310233-3213112331020110"></a>

## description_spec property — metadata / 103112230133 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3012111110001310-0320332122300121-1123003131000000-3023312012020212-2011011233323223-2301031221302122-3333321131001112-1103011221300101"></a>

<a id="canonical-2221001212231131-0203131321003313-1013231221131323-1221303110210221-0222221333332132-0310102211333101-3230123323323331-0323003122031301"></a>

## name property — metadata / 103112230133 / 5

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

<a id="canonical-3223132323031011-0312130132202033-1131300110211323-3032202210223110-1222211212230330-0301110101012112-2113221320202131-2210132231122301"></a>

## Next pages — metadata / 103112230133 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001011011100223-1320111012132220-3320323111211000-1331123232221020-0012231223011032-2101000323203223-3122220302222131-1322301113020012"></a>

## api_specification.validation_all_spec_endpoints.settings — settings / 130323310323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-2132231120232323-3102332213300103-0332011323311313-0131201112131010-2032123333330201-2111330101221222-3231130223201002-2110232312220113"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

<a id="canonical-3221223300303303-1000211132020210-1010322223222333-2021120331020312-3021200332312120-0112013033332012-1311301303321131-2332303100102022"></a>

## Direct properties — settings / 130323310323 / 3

- [oversized_body_fail_validation](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2032123213203300-3012321101203231-3031322012132330-2312112230303101-1031030110332030-3131132312022323-2022002033121210-2032113331022230): complete subsection reference.

- [oversized_body_skip_validation](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1322123231323112-0302010010111030-3222232133100230-0302132111101110-2322103311313233-2012201321013331-0021022112030220-0030013322321220): complete subsection reference.

- [property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3122223131200201-1103022231023032-3220222100222030-2032030233023233-1233300210331120-2030212211003130-2003213120012111-0313311330013111): complete subsection reference.

- [property_validation_settings_default](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1231323101123001-1202002133333233-2213212021020013-3220133302210223-0011013301202230-1002310131033101-2030310111110301-1111211200312123): complete subsection reference.

<a id="canonical-1330110210103000-1221312311203113-0131010000023230-3122001300020213-3002001000323303-2223103231121020-3101221020033031-2220102331010020"></a>

## Next pages — settings / 130323310323 / 4

- [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2032123213203300-3012321101203231-3031322012132330-2312112230303101-1031030110332030-3131132312022323-2022002033121210-2032113331022230)
- [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1322123231323112-0302010010111030-3222232133100230-0302132111101110-2322103311313233-2012201321013331-0021022112030220-0030013322321220)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3122223131200201-1103022231023032-3220222100222030-2032030233023233-1233300210331120-2030212211003130-2003213120012111-0313311330013111)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1231323101123001-1202002133333233-2213212021020013-3220133302210223-0011013301202230-1002310131033101-2030310111110301-1111211200312123)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2032123213203300-3012321101203231-3031322012132330-2312112230303101-1031030110332030-3131132312022323-2022002033121210-2032113331022230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213100030301203-2313312113031110-1103211220130101-0303132323320231-1020031233123332-1002202300023320-3310201321313322-2110221230023031"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation — oversized_body_fail_validation / 121123313133 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-1122110330103303-2131131331232203-1100122021233121-2321131321103122-2013213013022220-3211231201321130-2302001110333230-3210330220220223"></a>

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

<a id="canonical-2231030103010003-2232030333130223-2321111103212202-0311210311000102-1101233133123032-0031113312312210-2110023230011031-1033031303301223"></a>

## Direct properties — oversized_body_fail_validation / 121123313133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213111202323032-3200333020332132-2121002112310310-3020131102001222-0130023233021121-2022111123310322-2022123223023323-0023232103333303"></a>

## Next pages — oversized_body_fail_validation / 121123313133 / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1322123231323112-0302010010111030-3222232133100230-0302132111101110-2322103311313233-2012201321013331-0021022112030220-0030013322321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302112311111111-2330333221310131-2003101310030122-1122332221023133-3022130220010233-1112220112213232-3232112121030202-1110331203003013"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation — oversized_body_skip_validation / 333110310302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-2130301011103132-1333231212023113-2120212020233303-0112102120333003-2110221122310133-0101023231031330-0321302320021213-3321011202300002"></a>

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

<a id="canonical-2201001112112101-1031200011301103-0211032210121132-3222121123132200-0011110302320221-1003221012032100-3100101321033000-3122213133212321"></a>

## Direct properties — oversized_body_skip_validation / 333110310302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122002113220321-0310113122313200-1013202031230310-0113331223102203-0200300303122000-1101221201032122-0311031323301030-2121221020200213"></a>

## Next pages — oversized_body_skip_validation / 333110310302 / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3122223131200201-1103022231023032-3220222100222030-2032030233023233-1233300210331120-2030212211003130-2003213120012111-0313311330013111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
