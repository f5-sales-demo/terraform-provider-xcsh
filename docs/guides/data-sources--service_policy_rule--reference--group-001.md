---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010320022113122-2232111033123230-2101133001312012-2202023001331332-0000322310122002-3203213211133133-1323102211133313-2312022131231313"></a>

## Property reference — Property reference / 023333031212 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- Property reference

<a id="canonical-1331123322123232-0001023320322223-2302111121101101-1031230212123221-0200032302233322-1201232322122101-0202313033013201-0103230121011122"></a>

## Direct properties — Property reference / 023333031212 / 3

<a id="canonical-3300221333210321-1011012232032232-1122022031322312-3002331123211102-3221020032021203-0122201022100231-2202213021200132-1120111321102322"></a>

<a id="canonical-0202201010311130-3321012032331012-2310310100123300-1011323031001332-3101133110212202-0122202102130032-2011211000332202-2121223313213012"></a>

## action property — Property reference / 023333031212 / 4

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103210323313210-1332220213100230-2000023201300121-0201320331231213-2102111101023002-2033321211223322-0210100202232011-0212002002220301"></a>

<a id="canonical-0011100120123321-1120311230022001-3120003031322103-0231103032232101-1023031121331133-1103112023310220-0100323112031303-0222231120210022"></a>

## annotations property — Property reference / 023333031212 / 5

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

- [any_asn](data-sources--service_policy_rule--reference--group-001.md#canonical-3021332231300121-1013322201002112-2132213331221313-0333231230031023-0200113020331032-0333203121223033-3212000203002221-2232321123102030): complete subsection reference.

- [any_client](data-sources--service_policy_rule--reference--group-001.md#canonical-2201201330130100-1123001021023020-3303233322313203-3331102111111233-1313213321012112-3332023310232102-0333301231202200-2113111111230220): complete subsection reference.

- [any_ip](data-sources--service_policy_rule--reference--group-001.md#canonical-2010322331300203-1113233230123032-2023011223232330-0323222220231000-0113102103000012-1012103310023010-2111000200320203-1232223201332301): complete subsection reference.

- [api_group_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2323220103100030-2300221222133301-0312321332223112-2330231103123033-2032200011132020-3100000303330210-1222301113122130-2312000321110100): complete subsection reference.

- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112): complete subsection reference.

- [asn_list](data-sources--service_policy_rule--reference--group-001.md#canonical-2323110012210121-3300120010211011-3021130303320203-1302110103133022-2103210011332101-3311131100232303-0022131322313002-2203010100023112): complete subsection reference.

- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2123130101302113-2200222301333131-0110022202020010-1333032113332333-3123321211312310-0322220322033200-0100001233131121-1103233030221220): complete subsection reference.

- [body_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0023002331223020-0113212003033021-2231301311112131-1031101231203003-0013032323230200-0030122222320202-2000131202022210-1112020323222210): complete subsection reference.

- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320): complete subsection reference.

<a id="canonical-1221313132002113-0233102000032332-2030000020231122-2230220131010132-2101200200112002-3221330131232310-3033313133311110-2303132103310230"></a>

<a id="canonical-2130203320000031-1211212000323121-3110000200230230-2330300221210300-1100321333331201-0123013111222110-1302133033213203-0101123111332111"></a>

## client_name property — Property reference / 023333031212 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Upstream description:

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [client_name_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0333301122101131-3120331133111122-3030213133031021-1111322320200213-2203212111231210-1021112003222010-3011030011331310-3302320221031101): complete subsection reference.

- [client_selector](data-sources--service_policy_rule--reference--group-001.md#canonical-1123330002113022-1313333303031311-1222120120201120-1231033101313011-2011232002032001-3330222101101201-3032001310131123-1223131301212333): complete subsection reference.

- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232): complete subsection reference.

<a id="canonical-1001231000023020-3123032013102231-0121311123203321-2232130212313232-2123301112000200-1320201022211230-3213230123122221-2221212333002321"></a>

<a id="canonical-3213201023013330-2010211320222120-0121221232023322-0030231132203023-1303213130013100-2001133303203031-2030222100310122-3332021113002330"></a>

## description property — Property reference / 023333031212 / 7

Type: `"string"`. Computed.

Description of the ServicePolicyRule.

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

- [domain_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3322032130021031-3102103112002010-3312231133031210-2322012100203213-0213021201303211-2313203320103201-1031123311112323-2112122121122111): complete subsection reference.

<a id="canonical-3203330313130200-3202120103220101-1333122101113001-2133233330123133-0312110033231302-1003322332101230-2113131302122013-3012120021302223"></a>

<a id="canonical-2300320200232023-2300201230130231-2301310111021111-3322122210230032-2002302000013132-3021330203222023-0331110223020220-1113322002303023"></a>

## expiration_timestamp property — Property reference / 023333031212 / 8

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020): complete subsection reference.

- [http_method](data-sources--service_policy_rule--reference--group-001.md#canonical-0230303322301220-0201213302213132-3013033111230100-1212133331313130-2022033111330210-3011103212333123-3113001121120131-3123200320022331): complete subsection reference.

<a id="canonical-3113112111200320-0133311313223303-1302020322103102-3111302022123121-1111222120030012-0233210121012030-0323011330322223-2120010230112013"></a>

<a id="canonical-3130300310200012-3300211101101212-0121022111123231-1221122132131223-0130303221131302-3012102120122012-2203122123121111-0001310330331000"></a>

## ID property — Property reference / 023333031212 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1030121033313010-0032013003231230-2231312010033301-0022213330223022-2212312013310022-0011013023030122-2322233132210211-3302310303213130): complete subsection reference.

- [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-3311122011122130-1231221232013110-1321323211302210-3012120023003303-2303221232112331-3110301313311233-1110200101321222-0133333103002120): complete subsection reference.

- [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-3202102000312012-0320203211302133-2311121232212120-0010321012011000-1103221331232212-3320312321121320-0001000123312010-0122223300013313): complete subsection reference.

- [ja4_tls_fingerprint](data-sources--service_policy_rule--reference--group-001.md#canonical-3323210221002100-0120011222202210-3321322010321201-0231111312300102-1112122201231031-1313121202130301-0022111303332121-3233213330333003): complete subsection reference.

- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312): complete subsection reference.

- [label_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-3210311000320320-2030222121122100-0220110100321312-1121102131023230-3300133221311211-2200300323311103-1132001130323212-2010010230320023): complete subsection reference.

<a id="canonical-0232003211010103-3212200032012130-0033103220301110-1023021132213033-2010303221323033-3032202121111020-3122113102132303-3000031003003121"></a>

<a id="canonical-0023331201230003-3002323112132321-3032122323133111-1031101011112131-1123020200122213-0303232311100213-1001031332121112-0100001002001313"></a>

## labels property — Property reference / 023333031212 / 10

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

<a id="canonical-1101230321300123-1331021222121022-2201123223200312-3323333202100013-1222220113321013-1102310023211130-3231230301311112-2301221102220331"></a>

<a id="canonical-0110002230030001-0320232100331211-2123331101220032-0102312112032320-2200013313130222-0330110113220112-0220212110220123-1321303202132312"></a>

## log_rule_evaluation property — Property reference / 023333031212 / 11

Type: `"bool"`. Computed.

Log the rule match details along with the request and continue to evaluate rules in the sequence.

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

- [mum_action](data-sources--service_policy_rule--reference--group-002.md#canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303): complete subsection reference.

<a id="canonical-3312013302110003-1100231322120110-3233233123303221-3011301232312012-1101231120132120-2000111221201301-1333311102032203-0203223100201020"></a>

<a id="canonical-1203200033321233-3000232211322322-1320311100301032-3321010303012113-0101121322232302-1332232333212131-1332202202223310-2012202001303200"></a>

## name property — Property reference / 023333031212 / 12

Type: `"string"`. Required.

Name of the ServicePolicyRule.

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

<a id="canonical-3130223300023011-2022011331301002-2202120023233300-1330202223101222-2103122300213021-2030120103003232-0010121010200222-2001121132220331"></a>

<a id="canonical-0000231000332212-3013201321122031-1310133133020021-0032101030302112-3320001320310000-1021231333130100-2310212120303322-2102030323130003"></a>

## namespace property — Property reference / 023333031212 / 13

Type: `"string"`. Required.

Namespace where the ServicePolicyRule exists.

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

- [path](data-sources--service_policy_rule--reference--group-002.md#canonical-1202323031323101-3010322010211020-2201013123220300-0120101231322302-1330322232232230-0002133000323121-2021123103113212-2023022232033013): complete subsection reference.

- [port_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-2101210020323103-3303202202131321-1133232111123110-1302223220031001-1031030010123130-1010013200001221-1302232103001023-1320212133031120): complete subsection reference.

- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203): complete subsection reference.

- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010): complete subsection reference.

- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-0133301223123320-2232020220130121-1213202001010012-3222021111322322-0203313233311020-3313132002210111-1222220233313311-1202000232133022): complete subsection reference.

- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322): complete subsection reference.

<a id="canonical-2321332111022311-0011300200322131-1002013021021220-1013222201221022-1033022311210202-2220023203220133-0122030122332203-1212120031123001"></a>

## All schema paths — Property reference / 023333031212 / 14

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](data-sources--service_policy_rule--reference--group-001.md#canonical-3300221333210321-1011012232032232-1122022031322312-3002331123211102-3221020032021203-0122201022100231-2202213021200132-1120111321102322) |
| `annotations` | [annotations](data-sources--service_policy_rule--reference--group-001.md#canonical-2103210323313210-1332220213100230-2000023201300121-0201320331231213-2102111101023002-2033321211223322-0210100202232011-0212002002220301) |
| `any_asn` | [any_asn](data-sources--service_policy_rule--reference--group-001.md#canonical-1103323332130021-1130010320022321-2023010233233023-0112310301203100-0302223100132002-2212120013210010-0233031113220223-3213322301131021) |
| `any_client` | [any_client](data-sources--service_policy_rule--reference--group-001.md#canonical-2232333100231331-2001201301222032-1023312023330222-2113120310013130-2331123030332013-1210211003120133-0231011111012011-2101103222033221) |
| `any_ip` | [any_ip](data-sources--service_policy_rule--reference--group-001.md#canonical-0313320230011213-1302003001100002-0301310012032122-0221102332320102-1221122201331002-0333310301200222-3230230310132333-2103332320221303) |
| `api_group_matcher` | [api_group_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1320220012022000-1103133113330120-3330100332131202-1313123221132133-1000013231001122-2231030201200110-1012031012130032-1000221200100301) |
| `api_group_matcher.invert_matcher` | [api_group_matcher.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1121201002123021-2313002110321021-0021032111020011-1310102222322120-1100331102232330-2032212331121032-1221023112001333-2113131023001022) |
| `api_group_matcher.match` | [api_group_matcher.match](data-sources--service_policy_rule--reference--group-001.md#canonical-1002111223312002-0110313130211000-2102121312000332-1101113001011010-2311121020000030-0020103003332122-1030103231023313-1010131210112002) |
| `arg_matchers` | [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2023212113302303-2202212212112323-0020110030001332-1231302201311002-1011021002323320-2303003201331232-0010200323010113-2102132333223122) |
| `arg_matchers.check_not_present` | [arg_matchers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2022310320210323-3302312322030200-2022300112203032-3101232001201210-3021323131123130-1113003001210322-2022102001132111-0000332021311313) |
| `arg_matchers.check_present` | [arg_matchers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-3333231321333211-2331112212120201-0220213022120303-3323221313310013-0212030203323020-1312332000133333-2222210230302223-3201013032131222) |
| `arg_matchers.invert_matcher` | [arg_matchers.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1330300031000022-1211111010033312-1002203330331233-1320233010320223-2202332230222233-1110323223211222-2010001003011330-1121011132230210) |
| `arg_matchers.item` | [arg_matchers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-3000321222213202-1210333013030013-1022301313102211-1312332112212212-3023032002201020-2333123330013333-2201212321111313-0200122301202021) |
| `arg_matchers.item.exact_values` | [arg_matchers.item.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-0103000023210113-1323032230032211-3310113213303212-1331013203231220-0000113002133333-0100310230313330-3103033200121023-1233321211103221) |
| `arg_matchers.item.regex_values` | [arg_matchers.item.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-3202212002113011-0223332322032230-0032230121311033-1333313220010112-0302130322212213-1011022313112031-2220203122202121-0203011201003032) |
| `arg_matchers.item.transformers` | [arg_matchers.item.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-3211120131220303-1121331002231131-2030323133203321-2200222111221312-2213333133301000-1333122201300000-2322210331001323-0100002332202223) |
| `arg_matchers.name` | [arg_matchers.name](data-sources--service_policy_rule--reference--group-001.md#canonical-3203003223201332-2201333110210221-1212203203231100-0030300031323122-2312203031332323-0303132302101212-2201212120220302-0121020202131333) |
| `asn_list` | [asn_list](data-sources--service_policy_rule--reference--group-001.md#canonical-1322323023222131-2322301032320000-2223112213032313-3033201110301033-2223100302210312-3211332033011002-1210033202322002-3303332112101301) |
| `asn_list.as_numbers` | [asn_list.as_numbers](data-sources--service_policy_rule--reference--group-001.md#canonical-1022001300331020-2202331103312101-2313100103303220-0011132033111302-1002302021023200-1200131001012121-2303103033200030-1010210200311331) |
| `asn_matcher` | [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2121010320211310-2113000100221200-0120121023000110-0010200020000220-0231301010310211-3211123101233310-1200010331101032-2003100121101220) |
| `asn_matcher.asn_sets` | [asn_matcher.asn_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-2033120202232101-3022000030110322-1033022300320103-3111312033011103-3002032212313131-3002201210112300-1123220120233320-3101112332220223) |
| `asn_matcher.asn_sets.kind` | [asn_matcher.asn_sets.kind](data-sources--service_policy_rule--reference--group-001.md#canonical-2333310231033111-2010030030332020-0210003311313130-0302022110122101-0023232100332130-0101320310201110-0223021220121000-3322010011210220) |
| `asn_matcher.asn_sets.name` | [asn_matcher.asn_sets.name](data-sources--service_policy_rule--reference--group-001.md#canonical-1302212330223203-3232030233033131-1101200000012332-2321103101201313-0203113023310211-1130322023022331-2002323133330111-1212331133002311) |
| `asn_matcher.asn_sets.namespace` | [asn_matcher.asn_sets.namespace](data-sources--service_policy_rule--reference--group-001.md#canonical-3122303003312323-3112113313321021-0322320022303010-0023012003212233-2301101113011132-0303121031133123-0203032020112111-0000000212033213) |
| `asn_matcher.asn_sets.tenant` | [asn_matcher.asn_sets.tenant](data-sources--service_policy_rule--reference--group-001.md#canonical-2312121211123100-3332221202221210-0003132231212111-2113113333130332-1021120312023123-3020001012003130-2030013021113321-0111101120323002) |
| `asn_matcher.asn_sets.uid` | [asn_matcher.asn_sets.uid](data-sources--service_policy_rule--reference--group-001.md#canonical-0032023312223013-0231111303232001-1210331123102230-1033111030332322-3133101101332133-2200303131032320-3130230123332220-0003103030002010) |
| `body_matcher` | [body_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0322031300230322-3320223211302110-0200330201013300-3331113030130121-2111032033313121-1300202310001211-1220222000133023-1102210221232130) |
| `body_matcher.exact_values` | [body_matcher.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-3303022213021030-0110021020233322-3330101322202001-0310332121111113-1233011221312012-3312303233231022-0013323133332333-3100011020033310) |
| `body_matcher.regex_values` | [body_matcher.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-0221021000002133-1000233002232022-0001123301310110-3011203123221323-3213200330312101-2310223322100112-3300212111213022-1331330321130110) |
| `body_matcher.transformers` | [body_matcher.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-3222132232321323-3320210030331000-3333100032021211-0022002030300033-0013001210020111-1133031223023112-2023310001200222-3312331300011211) |
| `bot_action` | [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-0012123222233210-0203221210313202-0102211121020221-1200321221120113-2123021121210332-0020001102203112-3331332033333213-1112233020212130) |
| `bot_action.bot_skip_processing` | [bot_action.bot_skip_processing](data-sources--service_policy_rule--reference--group-001.md#canonical-1131002201121333-1003201232021310-0321123333003223-1000202111223231-1001011313302301-1223332033122103-0201003032121122-3203220122021330) |
| `bot_action.none` | [bot_action.none](data-sources--service_policy_rule--reference--group-001.md#canonical-3131121022332130-2130113020011323-1330122113121132-0232121102112303-3133131000100320-1323113023011033-0231000113330103-3303033032211221) |
| `client_name` | [client_name](data-sources--service_policy_rule--reference--group-001.md#canonical-1221313132002113-0233102000032332-2030000020231122-2230220131010132-2101200200112002-3221330131232310-3033313133311110-2303132103310230) |
| `client_name_matcher` | [client_name_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0232200201211231-3003231012023103-1001131131111002-0223210230203122-3322021220320103-1300210332123320-3113122212120323-3300300232011012) |
| `client_name_matcher.exact_values` | [client_name_matcher.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-0200231123130022-3230132131321111-1310300322220212-1210111101011003-1230013222323022-0300300101220331-2330201031311100-2231032331132223) |
| `client_name_matcher.regex_values` | [client_name_matcher.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-0301331113130321-2132312221223123-0212121003003102-1223300001312223-2031210211220223-0210122102033220-1132312010321121-3301200000213210) |
| `client_selector` | [client_selector](data-sources--service_policy_rule--reference--group-001.md#canonical-3320212100230333-3113230330223211-2130302022230311-2022021132302231-3222023001010121-2203230303122230-0032302303130333-2100111332210200) |
| `client_selector.expressions` | [client_selector.expressions](data-sources--service_policy_rule--reference--group-001.md#canonical-3032331103030210-1133203002110233-3210100131221322-3313121112120020-1333230331110320-2331123230333021-2211030333220220-0100011000023020) |
| `cookie_matchers` | [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2323300210310032-0232010312211032-3121033321210300-2022102221102220-0032203331223022-3211001221322131-1113123033312333-0011021112330132) |
| `cookie_matchers.check_not_present` | [cookie_matchers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2013132320301232-3133111311010300-0131133300111303-0311010012021230-2302102020220011-3203022100030002-3002212022103100-3013322312120001) |
| `cookie_matchers.check_present` | [cookie_matchers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2013311203120230-2331110032222211-2211311300323012-0323303022103013-1301230302222010-1203111023020030-0202233322230111-2321122102132121) |
| `cookie_matchers.invert_matcher` | [cookie_matchers.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0013002301203230-0332012310002202-0300023223121031-2223012012303111-1332302003133021-2031230100003122-0023211012301212-3220212211131322) |
| `cookie_matchers.item` | [cookie_matchers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-2020030231312123-0023000321120302-0213232023301232-2220021010212001-2211203211133113-3331022220030311-3301332103320000-2032032333121012) |
| `cookie_matchers.item.exact_values` | [cookie_matchers.item.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-1032302220230311-2022010222133133-3213111132232232-2233113122201231-3211321020222131-1313033300112310-0132301200330333-2220122110223022) |
| `cookie_matchers.item.regex_values` | [cookie_matchers.item.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-2321312001123322-3201000321120301-1031113213030320-1120012302130210-2203133201322031-3232130202100110-3013211200132330-2212120103310032) |
| `cookie_matchers.item.transformers` | [cookie_matchers.item.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-1113201331310030-1210331201233200-2121313113231310-1213131221212123-3312300012202031-2330012022333300-0100321112202220-1013313310110312) |
| `cookie_matchers.name` | [cookie_matchers.name](data-sources--service_policy_rule--reference--group-001.md#canonical-2023132330032330-1023022213000230-2111212012321103-0001123203321323-0012321303302113-2300031001210303-0313130033120300-1012133113003210) |
| `description` | [description](data-sources--service_policy_rule--reference--group-001.md#canonical-1001231000023020-3123032013102231-0121311123203321-2232130212313232-2123301112000200-1320201022211230-3213230123122221-2221212333002321) |
| `domain_matcher` | [domain_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0312230300002332-0030301030103001-1131312203102213-0003101313212213-0231120303320021-2011323012110331-1302133110033011-3232301122011313) |
| `domain_matcher.exact_values` | [domain_matcher.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-1230222333132221-1112132201011112-2213013102102201-1312001011313202-1303301303003213-2010233212133231-2111003202231021-1000210032013213) |
| `domain_matcher.regex_values` | [domain_matcher.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-2201220032223120-1312231320011123-2301120323133003-3332231201202231-2133300030032210-2031301110132003-0133100101100123-0103121031011111) |
| `expiration_timestamp` | [expiration_timestamp](data-sources--service_policy_rule--reference--group-001.md#canonical-3203330313130200-3202120103220101-1333122101113001-2133233330123133-0312110033231302-1003322332101230-2113131302122013-3012120021302223) |
| `headers` | [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-1120302113020222-1102120012132113-3302321012332003-3120203331222113-0230323030021301-3103310320210203-3000013301131012-3321121202310100) |
| `headers.check_not_present` | [headers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2211121111132120-3200232303122130-0002221133130300-3013312211133111-1323233221121201-3122133023222021-2022213122122232-0301111312000032) |
| `headers.check_present` | [headers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-0003331212022223-3321211302032201-3300201323021101-0200110233313102-3232113101303102-3123112013002222-2013101120223220-1011000112321211) |
| `headers.invert_matcher` | [headers.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0012332132113303-0301233133332010-2213133232221002-0012132302313322-3132011113100132-0013231002310110-2222023101021222-2321011110300003) |
| `headers.item` | [headers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-1303303111021002-2032232212333310-2130123011031101-3011020031031313-0132233222210002-2003023231000321-2113111020203200-3023310001101202) |
| `headers.item.exact_values` | [headers.item.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-1033001323232301-1320001330220133-0221101023030311-1131001223011132-2000202033331333-2133110120302000-2113103201330021-1001230021001030) |
| `headers.item.regex_values` | [headers.item.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-0033230213000030-2103223223221230-1030132002131133-3123002323132021-2022301211031311-2312023130322110-2233012320112330-3113130331320011) |
| `headers.item.transformers` | [headers.item.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-0332223022221133-0311113233133330-3211212113222030-2131002021203202-0111011131023132-1221110103000313-3101010302330020-1100213200330010) |
| `headers.name` | [headers.name](data-sources--service_policy_rule--reference--group-001.md#canonical-3101220313113133-0100212120010231-1311331103213012-3010211033021202-1310031233120301-0011033333331311-0123203210203210-2021103133210120) |
| `http_method` | [http_method](data-sources--service_policy_rule--reference--group-001.md#canonical-3301213020313120-0213220212231131-1012213020121123-0200300300300033-0021031321301311-2313202002332032-2223312031123311-2301031031122322) |
| `http_method.invert_matcher` | [http_method.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2020310203000333-3013101230302023-0102020102330212-0223322032032120-1110132031023111-3213332333322010-0020132001113223-2122211130122210) |
| `http_method.methods` | [http_method.methods](data-sources--service_policy_rule--reference--group-001.md#canonical-0220201002321202-0320310022002331-3302301333200231-2212102033302031-2220301001002031-0332012031120023-3002133220022232-2232110330331131) |
| `id` | [ID](data-sources--service_policy_rule--reference--group-001.md#canonical-3113112111200320-0133311313223303-1302020322103102-3111302022123121-1111222120030012-0233210121012030-0323011330322223-2120010230112013) |
| `ip_matcher` | [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3302333201312031-2022013030010023-3201022223002000-2220101331302101-1013322320220110-1011232213020221-0303123223222203-3122103000203321) |
| `ip_matcher.invert_matcher` | [ip_matcher.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1023131331122331-1313112022223201-1002232212231222-2003102020213023-3222130130222231-2311222012130010-2111023300320102-0000302121320032) |
| `ip_matcher.prefix_sets` | [ip_matcher.prefix_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-1023033132101223-1111310323230100-2323130021321203-1122010213023233-1301113112022030-3303103300103201-1321303113123033-3122312320110313) |
| `ip_matcher.prefix_sets.kind` | [ip_matcher.prefix_sets.kind](data-sources--service_policy_rule--reference--group-001.md#canonical-2003303310031232-1230003000020013-2112013113331122-0033201111000100-2031212032223303-2131112331113323-1110203300321133-2211313112310213) |
| `ip_matcher.prefix_sets.name` | [ip_matcher.prefix_sets.name](data-sources--service_policy_rule--reference--group-001.md#canonical-0223320332213011-1122210133111031-1311331131022110-0323211312312021-3321002322011312-2310320000020300-0231211311001003-2132023002310321) |
| `ip_matcher.prefix_sets.namespace` | [ip_matcher.prefix_sets.namespace](data-sources--service_policy_rule--reference--group-001.md#canonical-3102003200033120-1220322122232023-3010330210113033-0011213013131233-1111203003222221-2021131012223033-2121312231133223-0303313031032113) |
| `ip_matcher.prefix_sets.tenant` | [ip_matcher.prefix_sets.tenant](data-sources--service_policy_rule--reference--group-001.md#canonical-2002303223331002-0223031001132100-2331230333103332-0012311220321030-2301323100201131-2302033101121023-0122101003032201-0210020232021021) |
| `ip_matcher.prefix_sets.uid` | [ip_matcher.prefix_sets.uid](data-sources--service_policy_rule--reference--group-001.md#canonical-3123031133003010-1111300000011120-0103331011211131-0101031322330102-2002332201223301-1332313311100012-3112132310122330-3011013231122231) |
| `ip_prefix_list` | [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-1212211231233102-1100131031220010-0321101033131300-2032320131002110-3002323203102331-0002021013112023-2033203132320102-0123012333023100) |
| `ip_prefix_list.invert_match` | [ip_prefix_list.invert_match](data-sources--service_policy_rule--reference--group-001.md#canonical-2211113211211023-0031202313100232-3032131001322010-1123023100011020-1203001220210332-1220330030311032-1202313331203213-1332303222011023) |
| `ip_prefix_list.ip_prefixes` | [ip_prefix_list.ip_prefixes](data-sources--service_policy_rule--reference--group-001.md#canonical-2032031130000130-0313110213110013-0130003132310212-2131023301201233-3033101033300232-1323310030301130-0300213213011022-2100230113031230) |
| `ip_threat_category_list` | [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-0232313301010221-2121221333323202-0022132312230200-1033310323333323-0202103311012132-3323213022030321-1001300131020020-0122112000033003) |
| `ip_threat_category_list.ip_threat_categories` | [ip_threat_category_list.ip_threat_categories](data-sources--service_policy_rule--reference--group-001.md#canonical-0033223212103030-0323002230011333-1130010311310010-1030102300023201-3132312022120110-1001303023203221-0200123311230102-0120201102011323) |
| `ja4_tls_fingerprint` | [ja4_tls_fingerprint](data-sources--service_policy_rule--reference--group-001.md#canonical-3101102133303110-1023320010002013-1113133022313303-2312010203332111-3121223232213313-0310022312200300-0022000232200213-1333113233203220) |
| `ja4_tls_fingerprint.exact_values` | [ja4_tls_fingerprint.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-0312302121301122-2233012201001130-0310221300331333-0121203032310022-0213130102203212-2212131331310110-3101202103332232-3122323020323101) |
| `jwt_claims` | [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-2100031211011310-0122001300303022-2231031011001203-1323002110133232-1013211312323311-0322302202321333-3031220203311031-2232011200202223) |
| `jwt_claims.check_not_present` | [jwt_claims.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-1120121133200333-2212031310311023-0002331000222231-0010330322231023-0020213210002121-3131330231121311-1230030020031312-3021301100200110) |
| `jwt_claims.check_present` | [jwt_claims.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-1331221331111102-1102032022212000-0232000323111103-1000113012030123-3310133032313131-2100020123020213-3301021233310110-1032303101101021) |
| `jwt_claims.invert_matcher` | [jwt_claims.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1230331313133033-1201303031330010-1202121223210013-3112030200232200-1010130021200202-3033300320233112-1320123102130311-3303230121202011) |
| `jwt_claims.item` | [jwt_claims.item](data-sources--service_policy_rule--reference--group-002.md#canonical-1211020202230012-3220003301020233-3203312001201133-3001032323312311-1110031110122001-3233023031130232-2231133233133311-3102023302121230) |
| `jwt_claims.item.exact_values` | [jwt_claims.item.exact_values](data-sources--service_policy_rule--reference--group-002.md#canonical-2133301320131231-0020011102112213-2011122110332202-2333103013020020-3301311300322012-3103032201213131-1233011223002230-2122020100001123) |
| `jwt_claims.item.regex_values` | [jwt_claims.item.regex_values](data-sources--service_policy_rule--reference--group-002.md#canonical-1332033123211330-0022220030300003-1121030333122222-0132221332332020-2102032112320103-3303031310013210-3323030010203022-3133103202032312) |
| `jwt_claims.item.transformers` | [jwt_claims.item.transformers](data-sources--service_policy_rule--reference--group-002.md#canonical-1101001201330102-0200132301102310-0123011303130023-2212101210003122-2321220202121011-1001031322130232-3123301302321232-0303313012123023) |
| `jwt_claims.name` | [jwt_claims.name](data-sources--service_policy_rule--reference--group-001.md#canonical-2211210223011032-2300331220200201-0022021130020203-0212201311012110-2322123302130323-2111033121303333-0101331113200302-2012000000333032) |
| `label_matcher` | [label_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-0200330310302023-3300303133133011-2123020230313101-1331223303030333-1303021022210130-2310222013032010-2320211001011312-1130301211201100) |
| `label_matcher.keys` | [label_matcher.keys](data-sources--service_policy_rule--reference--group-002.md#canonical-1001130000233222-3122310003013300-1303221203220123-1323100333203231-2300311003330313-3220301303212000-0022220212232010-3022000123332132) |
| `labels` | [labels](data-sources--service_policy_rule--reference--group-001.md#canonical-0232003211010103-3212200032012130-0033103220301110-1023021132213033-2010303221323033-3032202121111020-3122113102132303-3000031003003121) |
| `log_rule_evaluation` | [log_rule_evaluation](data-sources--service_policy_rule--reference--group-001.md#canonical-1101230321300123-1331021222121022-2201123223200312-3323333202100013-1222220113321013-1102310023211130-3231230301311112-2301221102220331) |
| `mum_action` | [mum_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1232230311312032-2121331211020302-0213033312300333-0012232230033211-3200033003020130-2312132120000303-0211133100312120-1112302303011322) |
| `mum_action.default` | [mum_action.default](data-sources--service_policy_rule--reference--group-002.md#canonical-3033220123200322-2031112332001123-2020131332123131-2132100310012032-0330122211030322-1112132322010133-0123033030113010-1001102023302202) |
| `mum_action.skip_processing` | [mum_action.skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-0100312101210231-0302220100200300-3002131133112312-0220212012323312-0100210000133313-0100021301232022-0211222230120232-0023333023103300) |
| `name` | [name](data-sources--service_policy_rule--reference--group-001.md#canonical-3312013302110003-1100231322120110-3233233123303221-3011301232312012-1101231120132120-2000111221201301-1333311102032203-0203223100201020) |
| `namespace` | [namespace](data-sources--service_policy_rule--reference--group-001.md#canonical-3130223300023011-2022011331301002-2202120023233300-1330202223101222-2103122300213021-2030120103003232-0010121010200222-2001121132220331) |
| `path` | [path](data-sources--service_policy_rule--reference--group-002.md#canonical-0031213201320113-0231310121110121-0121200331113332-0010202322232200-0011332112301101-0322223221223202-2103011210010201-0213023020212223) |
| `path.encoded_path_matcher` | [path.encoded_path_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-2233111103220210-0022110311303002-1021121233222031-1032011321031322-3011332103003213-0010111122202122-0203220132302331-0132212011220102) |
| `path.exact_values` | [path.exact_values](data-sources--service_policy_rule--reference--group-002.md#canonical-1021112133022313-1113200020120301-3112323111032321-1212111001101223-1223331200132322-2122031320002121-0031311213003312-3102303312330211) |
| `path.invert_matcher` | [path.invert_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-0122002222120103-2300311231001111-1311011013331230-2212113133301030-2200212210200301-3200133220230132-0032221313013111-2002222333212313) |
| `path.prefix_values` | [path.prefix_values](data-sources--service_policy_rule--reference--group-002.md#canonical-3323220100003020-2333331113121203-3310122032320133-0022011132123112-0122122331313132-2301220000322102-1220220211212231-3031330301200212) |
| `path.regex_values` | [path.regex_values](data-sources--service_policy_rule--reference--group-002.md#canonical-2311313133011303-3332310211130130-0201203330202201-1300002230232211-0002001110301003-2323210233001032-3201002301122202-0121200000031030) |
| `path.suffix_values` | [path.suffix_values](data-sources--service_policy_rule--reference--group-002.md#canonical-1222230001213311-0322221330330021-2013011220303322-0002110130221311-0221001011101022-2230231000332100-1302022031220301-2022211030222111) |
| `path.transformers` | [path.transformers](data-sources--service_policy_rule--reference--group-002.md#canonical-0123102202300202-3231310213302101-3230011013211232-0303202020230313-1310020313203131-0130000132320312-3201220320000200-3231200022221202) |
| `port_matcher` | [port_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-0131213012220100-3233021322333010-0232201212000122-1202123103003121-0133002011221000-3020313202313003-1303103303310102-0011031113301112) |
| `port_matcher.invert_matcher` | [port_matcher.invert_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-3330332300030201-2203330233023021-0002032132001020-2012201330321203-1020331230110211-1222103030122130-2321312211030033-3131103112310113) |
| `port_matcher.ports` | [port_matcher.ports](data-sources--service_policy_rule--reference--group-002.md#canonical-3101130102333312-2110110312330220-3020123011133333-0233110000213322-1222311330211123-0101303232310120-0122121013221113-3020002312112110) |
| `query_params` | [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-0110022223012231-1011022222022021-1231213210021012-2020023231312213-0120003203031103-1110212020200121-1320200132312313-1310013021311300) |
| `query_params.check_not_present` | [query_params.check_not_present](data-sources--service_policy_rule--reference--group-002.md#canonical-3011101203123210-1300300210101020-3321130320311201-3310230230131332-2213302110000323-0122201221030113-2011323332311220-3000300112000112) |
| `query_params.check_present` | [query_params.check_present](data-sources--service_policy_rule--reference--group-002.md#canonical-1002312133133323-0223313003230001-3331200333031002-1332121010001221-2012123200022211-2211201002322331-3223232333131320-1303131103120003) |
| `query_params.invert_matcher` | [query_params.invert_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-1103033200210122-3201213211221303-2230013333312012-1222031203122022-2022101301223122-2330130211320000-2021230222233110-2301132123331121) |
| `query_params.item` | [query_params.item](data-sources--service_policy_rule--reference--group-002.md#canonical-0330101112120322-3333133122212133-2332122021103233-1322100002311122-3333313123111132-1133023013130202-0221213301303033-2102321000311213) |
| `query_params.item.exact_values` | [query_params.item.exact_values](data-sources--service_policy_rule--reference--group-002.md#canonical-0233120011232212-2331113232020021-3322333002132232-1133021001021230-3210112012102013-2212101320301333-1120222012313232-2232020233211120) |
| `query_params.item.regex_values` | [query_params.item.regex_values](data-sources--service_policy_rule--reference--group-002.md#canonical-1131323033011300-2123101133231321-3221130011313103-3132232220303320-3331131121200131-0002313032300221-2033320002333012-1331313001110000) |
| `query_params.item.transformers` | [query_params.item.transformers](data-sources--service_policy_rule--reference--group-002.md#canonical-0300320221022220-3000131220203013-1103213222132210-2120301022320023-2123003333033112-1000032112123101-0000111222300023-1001002002122313) |
| `query_params.key` | [query_params.key](data-sources--service_policy_rule--reference--group-002.md#canonical-0312121020113322-3203333312133313-1203103311320333-3121102221233032-3201130110310000-0000311232031333-3002323232211121-3023333131220320) |
| `request_constraints` | [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0121303130321000-2010323021321020-0130112232322010-2222221123133001-3032121201323213-1210123132003033-2233210030211101-2031120112210020) |
| `request_constraints.max_cookie_count_exceeds` | [request_constraints.max_cookie_count_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-2133230320223220-1113123020232100-0233011100311311-2220223033232133-3111131323323231-3232322210332223-2332213312020201-3232132112311022) |
| `request_constraints.max_cookie_count_none` | [request_constraints.max_cookie_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1030113020022100-3110132033201201-3012322102131031-2003100123311002-1013100132030132-0013103030302130-0022123120223321-2203020213100101) |
| `request_constraints.max_cookie_key_size_exceeds` | [request_constraints.max_cookie_key_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-1303231131001120-1001120031110232-1103010020023232-3331030223330223-1331132022322023-0322322123113202-0030300232030133-3003313021120013) |
| `request_constraints.max_cookie_key_size_none` | [request_constraints.max_cookie_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1033333332111222-3003302032233333-0211010132033231-3310003311301003-3212033101301303-1032122303212112-2211301202111332-1031120033122313) |
| `request_constraints.max_cookie_value_size_exceeds` | [request_constraints.max_cookie_value_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-0103100211320112-3303303103023200-0221102113130302-2233122010100221-0132302032002032-0132232101102221-0211311320202031-3102102113001113) |
| `request_constraints.max_cookie_value_size_none` | [request_constraints.max_cookie_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0210323111300301-1111012003030002-1322022032031323-3021122231020033-0221120021330323-3121231110103200-1201231031302223-2021232101312332) |
| `request_constraints.max_header_count_exceeds` | [request_constraints.max_header_count_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-1011320330133023-1012201121103201-0220020110013000-1103113123322021-0303331033111012-0222221110223300-0121300123013110-3012233211103301) |
| `request_constraints.max_header_count_none` | [request_constraints.max_header_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1011200232203120-1120023222123130-3012332100333113-0203030232001031-2203132022312323-2223131332313013-2310122210301230-3211321030023300) |
| `request_constraints.max_header_key_size_exceeds` | [request_constraints.max_header_key_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-0021211210121302-2200300210020031-1011203031203000-1102310111013122-2210300122303131-1321322330121202-1101123303012123-1010010202232310) |
| `request_constraints.max_header_key_size_none` | [request_constraints.max_header_key_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0101223120103230-1001211233301323-1112113320113113-1203303322022001-2011130311001020-3132003002210131-2331311332120212-1130111111331002) |
| `request_constraints.max_header_value_size_exceeds` | [request_constraints.max_header_value_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-3020010003200211-0221120102310101-0200113112132022-3220032210112210-1300312232100223-0312220230303033-3110111223313202-0233000323022122) |
| `request_constraints.max_header_value_size_none` | [request_constraints.max_header_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3220222232320100-3012231320230101-0123211102311011-1301202121302323-0013311022300300-1220113302230332-0330210112021120-1301310000221100) |
| `request_constraints.max_parameter_count_exceeds` | [request_constraints.max_parameter_count_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-2230210011311131-3220103311102300-3112202003332123-2313200201113222-3011111012102202-2332010230231223-1312300231011212-2302110110222232) |
| `request_constraints.max_parameter_count_none` | [request_constraints.max_parameter_count_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1300321310102013-2312333223021233-0221302313001122-0300013002001321-3023232221132022-0200210310231210-1323002132331321-2022020122222212) |
| `request_constraints.max_parameter_name_size_exceeds` | [request_constraints.max_parameter_name_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-2011121232102213-3111131112321201-0122200021020202-0232203201103211-1211301033333121-1201003213032123-1201033231132220-3230213123111223) |
| `request_constraints.max_parameter_name_size_none` | [request_constraints.max_parameter_name_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1312322021102121-2103201031100211-3230031321110122-3102130021333311-0100330121032110-2132301222032332-2213021332010103-2101111300003003) |
| `request_constraints.max_parameter_value_size_exceeds` | [request_constraints.max_parameter_value_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-1202300203103111-1022032020110102-0000101131000020-1323231311011110-0201130121012113-3001321201123102-3332012303122003-1231120333210300) |
| `request_constraints.max_parameter_value_size_none` | [request_constraints.max_parameter_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-2002332310321030-2000302113032022-3321111201101313-3332023001302200-0202033321102113-3011330102300031-0330322121001002-0303032201223003) |
| `request_constraints.max_query_size_exceeds` | [request_constraints.max_query_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-2111112211231200-3121212100122210-2203032122130232-1331022232111100-2303220210233302-1320120203011133-3332222213313110-2120333320201121) |
| `request_constraints.max_query_size_none` | [request_constraints.max_query_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1312001012111122-3011222231323112-0313303221200022-3123133210220220-2302201211112122-2021012313222222-3221332020130101-2231113210010302) |
| `request_constraints.max_request_line_size_exceeds` | [request_constraints.max_request_line_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-1102112313031123-1113223221133223-0021013202022231-0232200231230232-2121000203210332-0030122300330213-0313312231313112-1213332012132020) |
| `request_constraints.max_request_line_size_none` | [request_constraints.max_request_line_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0323333201031003-0221202303322213-1211000011202220-2103103033331130-0230033213322231-3022022322031102-2202101323201231-1310221322311022) |
| `request_constraints.max_request_size_exceeds` | [request_constraints.max_request_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-0223002231203300-3122020102321311-1002001001122023-1133210130323002-1103311212233233-3113200020230020-2233021302033331-3200001133030313) |
| `request_constraints.max_request_size_none` | [request_constraints.max_request_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3133031123331012-0233222031021311-2223123201320023-1333300122021021-3303210120101113-3223310112302002-0003022331123312-0323113122002112) |
| `request_constraints.max_url_size_exceeds` | [request_constraints.max_url_size_exceeds](data-sources--service_policy_rule--reference--group-002.md#canonical-2113220021310133-3312302133003221-2120222232020321-0201023231321213-2332101031203212-2330121103303231-2203230022033021-2222030023031102) |
| `request_constraints.max_url_size_none` | [request_constraints.max_url_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-2231332203033130-2130030012322122-0213330212320022-0232123320130031-0333022121213313-2130200132132020-3022211301330302-3230321033321012) |
| `segment_policy` | [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-2031020012133013-2110330033230313-0220002102100223-3000000223311301-2113321321312110-0312203203302331-1220302301011132-2020123231020323) |
| `segment_policy.dst_any` | [segment_policy.dst_any](data-sources--service_policy_rule--reference--group-002.md#canonical-3313032031020120-1013330202221033-2110130331221320-2111001310103112-1122012202100103-0233120001322132-3213000321300112-3223302212323022) |
| `segment_policy.dst_segments` | [segment_policy.dst_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-0302121223330221-2323103000132013-1012113101320232-2330302231221131-1321322203212230-1220201211000102-2112131033200213-0230231031221233) |
| `segment_policy.dst_segments.segments` | [segment_policy.dst_segments.segments](data-sources--service_policy_rule--reference--group-002.md#canonical-3232232211120100-3120230303110001-1102300322102301-3022023303310101-2020200322022331-2313200300212333-3103002113013110-0132133233132312) |
| `segment_policy.dst_segments.segments.name` | [segment_policy.dst_segments.segments.name](data-sources--service_policy_rule--reference--group-002.md#canonical-3000332302212201-0223022223100003-3221203230123302-0211132331113323-3121331020011130-2302132300122321-2112322002213313-2302120011033133) |
| `segment_policy.dst_segments.segments.namespace` | [segment_policy.dst_segments.segments.namespace](data-sources--service_policy_rule--reference--group-002.md#canonical-0200330213223012-1002321113132022-2112311013013021-2213220213133301-2313231010012113-2021133230311122-1202320030200222-1303123223133302) |
| `segment_policy.dst_segments.segments.tenant` | [segment_policy.dst_segments.segments.tenant](data-sources--service_policy_rule--reference--group-002.md#canonical-2111113012033321-2301213202120313-0301120211112311-0202222112212212-3031332031110122-1133311113133011-2112032012230300-0111123122332120) |
| `segment_policy.intra_segment` | [segment_policy.intra_segment](data-sources--service_policy_rule--reference--group-002.md#canonical-3112210131131120-1002021133031301-2030012123301030-0322332021312201-2020230200203122-2103312230122302-2111301122321032-2031310133310120) |
| `segment_policy.src_any` | [segment_policy.src_any](data-sources--service_policy_rule--reference--group-002.md#canonical-0122132221001123-3333313131030020-0322321233112230-2221301320210123-1012231022310201-3112130223001331-2101111320212020-1320131003333212) |
| `segment_policy.src_segments` | [segment_policy.src_segments](data-sources--service_policy_rule--reference--group-002.md#canonical-1331220222003103-0321323122232233-1220010130200003-0211232323121110-3113211321121002-0203320300311103-3022011211201023-2120312032120000) |
| `segment_policy.src_segments.segments` | [segment_policy.src_segments.segments](data-sources--service_policy_rule--reference--group-002.md#canonical-2122202332330112-1003103211130200-3211002200113032-3332311003220322-3111333032313300-1212002102110302-2101112101020131-2021100111101023) |
| `segment_policy.src_segments.segments.name` | [segment_policy.src_segments.segments.name](data-sources--service_policy_rule--reference--group-002.md#canonical-0303203332220210-2100313121200330-0333010311222322-3130223110002112-3302120000111133-1033010200322013-0010330213312301-0333110132102000) |
| `segment_policy.src_segments.segments.namespace` | [segment_policy.src_segments.segments.namespace](data-sources--service_policy_rule--reference--group-002.md#canonical-3220122331000311-1012132222132112-3100333013233133-1130222033231310-3021002332003020-0031331121202332-0013131101302211-0120200120221231) |
| `segment_policy.src_segments.segments.tenant` | [segment_policy.src_segments.segments.tenant](data-sources--service_policy_rule--reference--group-002.md#canonical-2100203211202333-0002303211210220-3303302120230023-2122221102032310-3322320030122033-3031300020132213-1031310210312030-0113223203220302) |
| `tls_fingerprint_matcher` | [tls_fingerprint_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-3210102313131132-1011210000200333-2001031333102312-0013313211310201-0323202000203330-1101132130313031-1020332221301212-0331111311133010) |
| `tls_fingerprint_matcher.classes` | [tls_fingerprint_matcher.classes](data-sources--service_policy_rule--reference--group-002.md#canonical-1201310202132333-2011313030300331-1111021022231213-0300320033123200-1001313130111100-0202010121032222-0302120232113203-0211132203223320) |
| `tls_fingerprint_matcher.exact_values` | [tls_fingerprint_matcher.exact_values](data-sources--service_policy_rule--reference--group-002.md#canonical-0111300310230322-0313031200211310-2203021030102020-2220212211033201-2121323001320022-1313300002012111-2202223323213133-2323023032321023) |
| `tls_fingerprint_matcher.excluded_values` | [tls_fingerprint_matcher.excluded_values](data-sources--service_policy_rule--reference--group-002.md#canonical-2311311013303022-2012303321220002-1112212003323102-3023102302300221-1311031022303202-0211211312323221-2323132003000201-0132232332130003) |
| `waf_action` | [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1220313201012031-1031100221113212-2331033213200133-0012122003101020-1232003033001223-0233201223112301-1301311031211221-1212332121012202) |
| `waf_action.app_firewall_detection_control` | [waf_action.app_firewall_detection_control](data-sources--service_policy_rule--reference--group-002.md#canonical-2331311321100102-3032303201210011-1330122001130320-0220011133303031-2032020232100221-2212011011122111-1201010021202110-1112003013022322) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-2110333001322033-3212330132312023-2330020120202231-0121221113313303-0101321111020230-1022323002031220-3022010131310031-1331113203131100) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](data-sources--service_policy_rule--reference--group-002.md#canonical-2113310221211320-2213003332212012-1010311030023122-3131030133333303-2023021211310101-1333201322312013-0123003330103331-0123001101320300) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](data-sources--service_policy_rule--reference--group-002.md#canonical-3300002011323003-0100210313232001-1110331121321031-2231223130233123-3231311103331202-0102113320311201-0121023102100021-1213311223213221) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](data-sources--service_policy_rule--reference--group-002.md#canonical-1211033121233210-2001330110022110-1212012101323131-1231002323300101-1332301302330003-1030013122233100-2001002000032232-1212302103320212) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-1323012021120132-1031132233323102-3211013020103321-2112000300232212-2023312121323103-0231322012312103-3211313001323113-0220330100211112) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](data-sources--service_policy_rule--reference--group-002.md#canonical-1232210032121031-2102112302203332-0311331022231000-1203203211311212-0003220223012213-2222223113122222-0000203231330200-0133122123020101) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts` | [waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-0322300030322013-2013113332113001-3200111021300020-1302012212311003-1202002210220120-1010211321110331-2002123332003200-3003012101302230) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context](data-sources--service_policy_rule--reference--group-002.md#canonical-1231302122311322-0001222220331223-1102032101332033-0330312010310000-1033203001012233-2313310113220301-0221312031013023-0121203013312203) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](data-sources--service_policy_rule--reference--group-002.md#canonical-0203021032133203-3300110231231312-0022101322201303-2231112010212010-2032033320333023-0133332111302112-2103033013300202-3210003332301013) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](data-sources--service_policy_rule--reference--group-002.md#canonical-0232033020320133-3020132120123303-2103000223030333-2312313233022222-2023031302120310-3231213311020303-0232301220322330-2000012010230320) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts` | [waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy_rule--reference--group-002.md#canonical-3003011100221213-3333131100032013-2321120031100020-3001030001013121-3003223202131031-1323203230200311-1233330110300100-2130031301233321) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context](data-sources--service_policy_rule--reference--group-002.md#canonical-2123203313000031-1200103131002010-1030132003131312-2133021111313213-1011233331103331-0131003113102120-0320203331032220-3202131210111012) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](data-sources--service_policy_rule--reference--group-002.md#canonical-1123323020223102-0222030001322223-2000011333213332-2223032301212022-2300311102133032-1100121133013232-1331020112200312-1001310101011233) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](data-sources--service_policy_rule--reference--group-002.md#canonical-2321313031323311-1212331103003003-2010021010303023-2211310001132312-0303210332002232-3011223201030122-0223103203003320-0010220212112320) |
| `waf_action.none` | [waf_action.none](data-sources--service_policy_rule--reference--group-002.md#canonical-1312012132001212-2311312302222231-1013103310113003-2220023033202101-2211301333032312-3302222131002000-2321012332232203-0303132310333322) |
| `waf_action.waf_skip_processing` | [waf_action.waf_skip_processing](data-sources--service_policy_rule--reference--group-002.md#canonical-1222213123001231-0232113320132011-3130011302332122-3301210313321330-3323320310132212-3122133011230113-2020312100111231-1022031213132011) |

<a id="canonical-3313002133133133-3012322311200101-1213122213102212-1333211210111100-1103100233030310-0213222020002123-0333030213120130-2002313021232302"></a>

## Next pages — Property reference / 023333031212 / 15

- [any_asn](data-sources--service_policy_rule--reference--group-001.md#canonical-3021332231300121-1013322201002112-2132213331221313-0333231230031023-0200113020331032-0333203121223033-3212000203002221-2232321123102030)
- [any_client](data-sources--service_policy_rule--reference--group-001.md#canonical-2201201330130100-1123001021023020-3303233322313203-3331102111111233-1313213321012112-3332023310232102-0333301231202200-2113111111230220)
- [any_ip](data-sources--service_policy_rule--reference--group-001.md#canonical-2010322331300203-1113233230123032-2023011223232330-0323222220231000-0113102103000012-1012103310023010-2111000200320203-1232223201332301)
- [api_group_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2323220103100030-2300221222133301-0312321332223112-2330231103123033-2032200011132020-3100000303330210-1222301113122130-2312000321110100)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- [asn_list](data-sources--service_policy_rule--reference--group-001.md#canonical-2323110012210121-3300120010211011-3021130303320203-1302110103133022-2103210011332101-3311131100232303-0022131322313002-2203010100023112)
- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2123130101302113-2200222301333131-0110022202020010-1333032113332333-3123321211312310-0322220322033200-0100001233131121-1103233030221220)
- [body_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0023002331223020-0113212003033021-2231301311112131-1031101231203003-0013032323230200-0030122222320202-2000131202022210-1112020323222210)
- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320)
- [client_name_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0333301122101131-3120331133111122-3030213133031021-1111322320200213-2203212111231210-1021112003222010-3011030011331310-3302320221031101)
- [client_selector](data-sources--service_policy_rule--reference--group-001.md#canonical-1123330002113022-1313333303031311-1222120120201120-1231033101313011-2011232002032001-3330222101101201-3032001310131123-1223131301212333)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- [domain_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3322032130021031-3102103112002010-3312231133031210-2322012100203213-0213021201303211-2313203320103201-1031123311112323-2112122121122111)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- [http_method](data-sources--service_policy_rule--reference--group-001.md#canonical-0230303322301220-0201213302213132-3013033111230100-1212133331313130-2022033111330210-3011103212333123-3113001121120131-3123200320022331)
- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1030121033313010-0032013003231230-2231312010033301-0022213330223022-2212312013310022-0011013023030122-2322233132210211-3302310303213130)
- [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-3311122011122130-1231221232013110-1321323211302210-3012120023003303-2303221232112331-3110301313311233-1110200101321222-0133333103002120)
- [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-3202102000312012-0320203211302133-2311121232212120-0010321012011000-1103221331232212-3320312321121320-0001000123312010-0122223300013313)
- [ja4_tls_fingerprint](data-sources--service_policy_rule--reference--group-001.md#canonical-3323210221002100-0120011222202210-3321322010321201-0231111312300102-1112122201231031-1313121202130301-0022111303332121-3233213330333003)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- [label_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-3210311000320320-2030222121122100-0220110100321312-1121102131023230-3300133221311211-2200300323311103-1132001130323212-2010010230320023)
- [mum_action](data-sources--service_policy_rule--reference--group-002.md#canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303)
- [path](data-sources--service_policy_rule--reference--group-002.md#canonical-1202323031323101-3010322010211020-2201013123220300-0120101231322302-1330322232232230-0002133000323121-2021123103113212-2023022232033013)
- [port_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-2101210020323103-3303202202131321-1133232111123110-1302223220031001-1031030010123130-1010013200001221-1302232103001023-1320212133031120)
- [query_params](data-sources--service_policy_rule--reference--group-002.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- [request_constraints](data-sources--service_policy_rule--reference--group-002.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100)
- [tls_fingerprint_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-0133301223123320-2232020220130121-1213202001010012-3222021111322322-0203313233311020-3313132002210111-1222220233313311-1202000232133022)
- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3021332231300121-1013322201002112-2132213331221313-0333231230031023-0200113020331032-0333203121223033-3212000203002221-2232321123102030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110222232132320-3301301221202011-0133300113120130-3313101233020312-3021022132113310-2233310000032003-3313000300101002-0132101331003100"></a>

## any_asn — any_asn / 222021222010 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- any_asn

<a id="canonical-1103323332130021-1130010320022321-2023010233233023-0112310301203100-0302223100132002-2212120013210010-0233031113220223-3213322301131021"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_asn, asn\_list, asn\_matcher\] Enable this option

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

- [any_asn](data-sources--service_policy_rule--reference--group-001.md#canonical-1103323332130021-1130010320022321-2023010233233023-0112310301203100-0302223100132002-2212120013210010-0233031113220223-3213322301131021)
- [asn_list](data-sources--service_policy_rule--reference--group-001.md#canonical-1322323023222131-2322301032320000-2223112213032313-3033201110301033-2223100302210312-3211332033011002-1210033202322002-3303332112101301)
- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2121010320211310-2113000100221200-0120121023000110-0010200020000220-0231301010310211-3211123101233310-1200010331101032-2003100121101220)

Select alternatives according to the provider validators above.

<a id="canonical-0300203323310112-1332331320132020-3030211310220033-2311202210201120-3032221301001301-3310120022220311-3012211012201330-3133113102021220"></a>

## Direct properties — any_asn / 222021222010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213220200311213-3312021212030033-2300001301033001-2210202230101302-0132223231100113-1300030300323311-2102201021212301-0010121312200233"></a>

## Next pages — any_asn / 222021222010 / 4

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2201201330130100-1123001021023020-3303233322313203-3331102111111233-1313213321012112-3332023310232102-0333301231202200-2113111111230220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000331303203021-3212030002300321-0302222203323233-3333020303301032-3321110311333230-3323323302111313-2030113300332203-1300103102013231"></a>

## any_client — any_client / 320122322222 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- any_client

<a id="canonical-2232333100231331-2001201301222032-1023312023330222-2113120310013130-2331123030332013-1210211003120133-0231011111012011-2101103222033221"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_client, client\_name, client\_name\_matcher, client\_selector,
ip\_threat\_category\_list\] Enable this option

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

- [any_client](data-sources--service_policy_rule--reference--group-001.md#canonical-2232333100231331-2001201301222032-1023312023330222-2113120310013130-2331123030332013-1210211003120133-0231011111012011-2101103222033221)
- [client_name](data-sources--service_policy_rule--reference--group-001.md#canonical-1221313132002113-0233102000032332-2030000020231122-2230220131010132-2101200200112002-3221330131232310-3033313133311110-2303132103310230)
- [client_name_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0232200201211231-3003231012023103-1001131131111002-0223210230203122-3322021220320103-1300210332123320-3113122212120323-3300300232011012)
- [client_selector](data-sources--service_policy_rule--reference--group-001.md#canonical-3320212100230333-3113230330223211-2130302022230311-2022021132302231-3222023001010121-2203230303122230-0032302303130333-2100111332210200)
- [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-0232313301010221-2121221333323202-0022132312230200-1033310323333323-0202103311012132-3323213022030321-1001300131020020-0122112000033003)

Select alternatives according to the provider validators above.

<a id="canonical-1011212202200321-1120323210203300-3110210011020331-3321232232223300-0333002001103120-1213302202030322-3020312310132022-2220123101031031"></a>

## Direct properties — any_client / 320122322222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323132230223202-3030220201231323-3113323223230321-3321322232211103-3033202201012233-2303030012330112-2220020231303302-3111301132000111"></a>

## Next pages — any_client / 320122322222 / 4

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2010322331300203-1113233230123032-2023011223232330-0323222220231000-0113102103000012-1012103310023010-2111000200320203-1232223201332301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130021312102111-1103222003020323-0302100110222102-3333133310032133-2003302323021122-1212022031311033-3313332300320310-3321000031001023"></a>

## any_ip — any_ip / 033203011203 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- any_ip

<a id="canonical-0313320230011213-1302003001100002-0301310012032122-0221102332320102-1221122201331002-0333310301200222-3230230310132333-2103332320221303"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_ip, ip\_matcher, ip\_prefix\_list\] Enable this option

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

- [any_ip](data-sources--service_policy_rule--reference--group-001.md#canonical-0313320230011213-1302003001100002-0301310012032122-0221102332320102-1221122201331002-0333310301200222-3230230310132333-2103332320221303)
- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3302333201312031-2022013030010023-3201022223002000-2220101331302101-1013322320220110-1011232213020221-0303123223222203-3122103000203321)
- [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-1212211231233102-1100131031220010-0321101033131300-2032320131002110-3002323203102331-0002021013112023-2033203132320102-0123012333023100)

Select alternatives according to the provider validators above.

<a id="canonical-1113333231131200-1333332303122103-2322201123223321-2123313001012211-2001010330332333-1033212011031030-2232021023123032-0033311220333112"></a>

## Direct properties — any_ip / 033203011203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121300211200122-1120331223110201-2111320003222302-0120022132313202-2013133312011322-1321120011031122-3310102113322233-2203022310003201"></a>

## Next pages — any_ip / 033203011203 / 4

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2323220103100030-2300221222133301-0312321332223112-2330231103123033-2032200011132020-3100000303330210-1222301113122130-2312000321110100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311210223312103-0133133313310033-1222310020223200-1200223002021000-3121132333331203-2231131331010010-2033232333200132-3213313120300031"></a>

## api_group_matcher — api_group_matcher / 320010231102 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- api_group_matcher

<a id="canonical-1320220012022000-1103133113330120-3330100332131202-1313123221132133-1000013231001122-2231030201200110-1012031012130032-1000221200100301"></a>

Type: `"single"`. Computed.

Matcher specifies a list of values for matching an input string. The match is considered successful
if the input value is present in the list. The result of the match is inverted if invert\_matcher is
true.

Upstream description:

A matcher specifies a list of values for matching an input string. The match is considered
successful if the input value is present in the list. The result of the match is inverted if
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

<a id="canonical-0311200232203032-3212320011113123-0321101003220321-3330003113122131-0113313231010023-1010133123202123-0021322021000221-1000032032300202"></a>

## Direct properties — api_group_matcher / 320010231102 / 3

<a id="canonical-1121201002123021-2313002110321021-0021032111020011-1310102222322120-1100331102232330-2032212331121032-1221023112001333-2113131023001022"></a>

<a id="canonical-0212030132331110-0100200012200121-3021102233331312-3323331033021122-2001113003021231-0101013311102002-3200310012000003-0301210200221112"></a>

## invert_matcher property — api_group_matcher / 320010231102 / 4

Type: `"bool"`. Computed.

Invert String Matcher. Invert the match result.

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

<a id="canonical-1002111223312002-0110313130211000-2102121312000332-1101113001011010-2311121020000030-0020103003332122-1030103231023313-1010131210112002"></a>

<a id="canonical-0212231211120010-2220333122000011-1232013221013000-2323330221102210-2101122111200003-3122332100011332-3102231212320220-2232112100111322"></a>

## match property — api_group_matcher / 320010231102 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0021220331211212-2103302231021012-3003202120222112-2120032220201301-1120031202333032-3122313221021203-0212301200213320-1321130310220323"></a>

## Next pages — api_group_matcher / 320010231102 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321103222030332-1312223033200200-3023012302222230-0133100211121032-3100300233211011-0020030331212101-2202220100112010-2333312211210000"></a>

## arg_matchers — arg_matchers / 113020011000 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- arg_matchers

<a id="canonical-2023212113302303-2202212212112323-0020110030001332-1231302201311002-1011021002323320-2303003201331232-0010200323010113-2102132333223122"></a>

Type: `"list"`. Computed.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

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

<a id="canonical-2301302000000211-3103200203012101-1211123003121131-1112001032211020-2301010320320333-3331323031110300-0121330020232331-0201320313302101"></a>

## Direct properties — arg_matchers / 113020011000 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-1123303301321301-1332102122030203-0103213021123112-1122022132212300-2000133122300001-1011300321201321-1200001201030111-1033220212221011): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2322323320313222-0233331221103131-1021102222110202-0010331131231122-1002310233303213-0020212331101022-3322233230331330-2323001000112201): complete subsection reference.

<a id="canonical-1330300031000022-1211111010033312-1002203330331233-1320233010320223-2202332230222233-1110323223211222-2010001003011330-1121011132230210"></a>

<a id="canonical-1302332022203332-2233311121133323-0303220322201321-0013023123330202-3120020201203100-1002021131222223-2133020203110200-1322330211011333"></a>

## invert_matcher property — arg_matchers / 113020011000 / 4

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

- [item](data-sources--service_policy_rule--reference--group-001.md#canonical-2332010133303230-0321112322213021-3021222010213233-2210110101312231-3321012133003211-2103322032020000-3210203232033232-2110301031331021): complete subsection reference.

<a id="canonical-3203003223201332-2201333110210221-1212203203231100-0030300031323122-2312203031332323-0303132302101212-2201212120220302-0121020202131333"></a>

<a id="canonical-3302110023131213-1333213031033212-0313102322123202-0323332223203312-0201020221010233-1001223021003023-1323020222022213-2123033133013000"></a>

## name property — arg_matchers / 113020011000 / 5

Type: `"string"`. Computed.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

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
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1133101221132132-1110133032300011-0031211133002030-1233011001310001-0100110332332111-1100110210333033-2231201232213311-3111013131121023"></a>

## Next pages — arg_matchers / 113020011000 / 6

- [arg_matchers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-1123303301321301-1332102122030203-0103213021123112-1122022132212300-2000133122300001-1011300321201321-1200001201030111-1033220212221011)
- [arg_matchers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2322323320313222-0233331221103131-1021102222110202-0010331131231122-1002310233303213-0020212331101022-3322233230331330-2323001000112201)
- [arg_matchers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-2332010133303230-0321112322213021-3021222010213233-2210110101312231-3321012133003211-2103322032020000-3210203232033232-2110301031331021)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1123303301321301-1332102122030203-0103213021123112-1122022132212300-2000133122300001-1011300321201321-1200001201030111-1033220212221011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203230333020301-2312213322112211-0033223300213203-3223011022313131-1303312003101320-3032131230000012-0312312222131200-0131110332331023"></a>

## arg_matchers.check_not_present — check_not_present / 331122230002 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- arg_matchers.check_not_present

<a id="canonical-2022310320210323-3302312322030200-2022300112203032-3101232001201210-3021323131123130-1113003001210322-2022102001132111-0000332021311313"></a>

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

<a id="canonical-2130000321030332-1003123222303022-0333232022130121-2233020222310301-2122133013030013-0000321302110121-3203331202200003-1001303022101312"></a>

## Direct properties — check_not_present / 331122230002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013102011303032-0321232202121331-2121310303232330-3320303021301032-0103113211103122-0003311231213320-2110233121210303-0212303110233122"></a>

## Next pages — check_not_present / 331122230002 / 4

- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2322323320313222-0233331221103131-1021102222110202-0010331131231122-1002310233303213-0020212331101022-3322233230331330-2323001000112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323313030012303-1330303200210003-2300303222220232-0200201231212032-2311121122222020-1103302011211321-2031312030322103-1303211130022213"></a>

## arg_matchers.check_present — check_present / 220132220131 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- arg_matchers.check_present

<a id="canonical-3333231321333211-2331112212120201-0220213022120303-3323221313310013-0212030203323020-1312332000133333-2222210230302223-3201013032131222"></a>

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

<a id="canonical-0311130230122332-0301312213122010-1330032011103313-3320111102313230-1112300030130330-3330020132111202-3200100001013202-1120211232212022"></a>

## Direct properties — check_present / 220132220131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223222120101020-3120000121103132-3223101230123020-0032321031233000-1321200002123031-0133332203120312-3113012220123211-2213311301300010"></a>

## Next pages — check_present / 220132220131 / 4

- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2332010133303230-0321112322213021-3021222010213233-2210110101312231-3321012133003211-2103322032020000-3210203232033232-2110301031331021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033311202031312-2131131032213033-1332102230011031-3200032100113013-0330102010320220-2101033323321130-1101111321002111-1131211100012103"></a>

## arg_matchers.item — item / 230023200203 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- arg_matchers.item

<a id="canonical-3000321222213202-1210333013030013-1022301313102211-1312332112212212-3023032002201020-2333123330013333-2201212321111313-0200122301202021"></a>

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

<a id="canonical-0011133023322123-3021321102033010-0103200130333133-1130030002002031-0001000033220230-2113013303321303-1133132032200233-0322330313212322"></a>

## Direct properties — item / 230023200203 / 3

<a id="canonical-0103000023210113-1323032230032211-3310113213303212-1331013203231220-0000113002133333-0100310230313330-3103033200121023-1233321211103221"></a>

<a id="canonical-0112300013321332-2200022200001322-1332220013101021-2230213310001131-3323011130022131-2001130301222113-3300020310110031-1023120103321101"></a>

## exact_values property — item / 230023200203 / 4

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

<a id="canonical-3202212002113011-0223332322032230-0032230121311033-1333313220010112-0302130322212213-1011022313112031-2220203122202121-0203011201003032"></a>

<a id="canonical-0230022123320322-3023010120222232-3030111201331221-0020312133210110-2330231133003001-1220111101133312-1000010321233331-1013022002302023"></a>

## regex_values property — item / 230023200203 / 5

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

<a id="canonical-3211120131220303-1121331002231131-2030323133203321-2200222111221312-2213333133301000-1333122201300000-2322210331001323-0100002332202223"></a>

<a id="canonical-0330123101213003-1123312022331220-3020120311222213-2301320023230231-1030021210021022-1120120232132220-0112002013010201-1121132321022331"></a>

## transformers property — item / 230023200203 / 6

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0210003020011101-2332023120200033-0232200132123232-3101100032211323-1211203231202210-1331112330030030-1222323131221302-2301100323203232"></a>

## Next pages — item / 230023200203 / 7

- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2323110012210121-3300120010211011-3021130303320203-1302110103133022-2103210011332101-3311131100232303-0022131322313002-2203010100023112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002321300223121-2303321032200010-3330233320212201-2113132021030003-2221023102320030-2133113033020113-1033020331322002-0231322200301302"></a>

## asn_list — asn_list / 030320011032 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- asn_list

<a id="canonical-1322323023222131-2322301032320000-2223112213032313-3033201110301033-2223100302210312-3211332033011002-1210033202322002-3303332112101301"></a>

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

<a id="canonical-1201030033131233-2102322211010123-1133100332200201-0323303000301102-2201310221113111-1331230220200100-2212122232130322-1002222333333030"></a>

## Direct properties — asn_list / 030320011032 / 3

<a id="canonical-1022001300331020-2202331103312101-2313100103303220-0011132033111302-1002302021023200-1200131001012121-2303103033200030-1010210200311331"></a>

<a id="canonical-0332110121110200-3233232323320201-3030323011103223-1223121110321123-3201320131232122-3203102000131130-1320103030203232-0000130003001020"></a>

## as_numbers property — asn_list / 030320011032 / 4

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

<a id="canonical-3130302130201110-3011030333200300-0001000121011331-3121020130032120-2000101311201331-3003201110103201-3102330321123111-2120000323132030"></a>

## Next pages — asn_list / 030320011032 / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2123130101302113-2200222301333131-0110022202020010-1333032113332333-3123321211312310-0322220322033200-0100001233131121-1103233030221220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302103312101303-1230203300232202-3312232011332021-3211310102133203-2130332311030101-3131010103201010-1202310020003222-0323123010221023"></a>

## asn_matcher — asn_matcher / 113120321023 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- asn_matcher

<a id="canonical-2121010320211310-2113000100221200-0120121023000110-0010200020000220-0231301010310211-3211123101233310-1200010331101032-2003100121101220"></a>

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

<a id="canonical-2312003022023132-0021332220331112-1333123113211332-0200212131002032-3101220003203101-3213131231233230-3101221321230200-0133310331320322"></a>

## Direct properties — asn_matcher / 113120321023 / 3

- [asn_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-3031222031221220-3103310132120011-3130222223310230-0202233201012300-0332220113201103-2031023123323212-1133222013221131-3300120301112032): complete subsection reference.

<a id="canonical-3202020232030010-0232231331301300-0233032032211103-1232233102300210-0212211111211300-2232313201033110-0310032233221132-1301201222320010"></a>

## Next pages — asn_matcher / 113120321023 / 4

- [asn_matcher.asn_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-3031222031221220-3103310132120011-3130222223310230-0202233201012300-0332220113201103-2031023123323212-1133222013221131-3300120301112032)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3031222031221220-3103310132120011-3130222223310230-0202233201012300-0332220113201103-2031023123323212-1133222013221131-3300120301112032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323022332131330-2301033112023202-2330012121030233-3113312321032001-2001000321312033-2030111020223122-1303231130313001-0332003200301101"></a>

## asn_matcher.asn_sets — asn_sets / 031302133210 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2123130101302113-2200222301333131-0110022202020010-1333032113332333-3123321211312310-0322220322033200-0100001233131121-1103233030221220)
- asn_matcher.asn_sets

<a id="canonical-2033120202232101-3022000030110322-1033022300320103-3111312033011103-3002032212313131-3002201210112300-1123220120233320-3101112332220223"></a>

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-3002113001312332-0120202103101123-0123323012010023-0202032012030121-1130123312021210-2320031111233013-0010323112321003-2203231011231223"></a>

## Direct properties — asn_sets / 031302133210 / 3

<a id="canonical-2333310231033111-2010030030332020-0210003311313130-0302022110122101-0023232100332130-0101320310201110-0223021220121000-3322010011210220"></a>

<a id="canonical-1210310130222022-0011121101002123-2233322001112023-1312322120033320-1130110002022222-2202200211220133-2321313231122332-2233111121012131"></a>

## kind property — asn_sets / 031302133210 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1302212330223203-3232030233033131-1101200000012332-2321103101201313-0203113023310211-1130322023022331-2002323133330111-1212331133002311"></a>

<a id="canonical-0203323212022132-0322012221331331-3030033323310030-1001333330131033-2130311321320002-3120103202031221-3210003320201100-0002301313033100"></a>

## name property — asn_sets / 031302133210 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3122303003312323-3112113313321021-0322320022303010-0023012003212233-2301101113011132-0303121031133123-0203032020112111-0000000212033213"></a>

<a id="canonical-1213212002301301-1303022130301333-1233233200002221-2332123311333303-1111302221010210-1003211031120010-3322021330231300-3103212221321333"></a>

## namespace property — asn_sets / 031302133210 / 6

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

<a id="canonical-2312121211123100-3332221202221210-0003132231212111-2113113333130332-1021120312023123-3020001012003130-2030013021113321-0111101120323002"></a>

<a id="canonical-1210031020133103-3010331311233021-3310130103123101-1101132322113132-2331321331301033-0301333232220101-2200202021021233-3212020223130101"></a>

## tenant property — asn_sets / 031302133210 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0032023312223013-0231111303232001-1210331123102230-1033111030332322-3133101101332133-2200303131032320-3130230123332220-0003103030002010"></a>

<a id="canonical-1321103112020112-1122131131120100-1221131023312112-0313121003333131-1020323101311231-0033300111103120-0031102021331131-2013130233222221"></a>

## uid property — asn_sets / 031302133210 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2132002221111120-1311222300220333-2001212331011110-3132223101001033-1313103002322333-1232300032122322-0020311010322133-2013000101310000"></a>

## Next pages — asn_sets / 031302133210 / 9

- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2123130101302113-2200222301333131-0110022202020010-1333032113332333-3123321211312310-0322220322033200-0100001233131121-1103233030221220)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0023002331223020-0113212003033021-2231301311112131-1031101231203003-0013032323230200-0030122222320202-2000131202022210-1112020323222210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031203031211110-3302331302122031-0130321133331033-3312303222321220-3033101001100231-2133013010300320-2230200023032031-3002020213332301"></a>

## body_matcher — body_matcher / 332101320320 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- body_matcher

<a id="canonical-0322031300230322-3320223211302110-0200330201013300-3331113030130121-2111032033313121-1300202310001211-1220222000133023-1102210221232130"></a>

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

<a id="canonical-1322320211220112-3130133102313211-2330133320333130-2220101003321211-2301013110320200-3012022100220201-3322011331333322-1111013011330300"></a>

## Direct properties — body_matcher / 332101320320 / 3

<a id="canonical-3303022213021030-0110021020233322-3330101322202001-0310332121111113-1233011221312012-3312303233231022-0013323133332333-3100011020033310"></a>

<a id="canonical-3233100023032302-3023102312330211-2013323131211003-3010021123301102-1303013123132003-2233232313032101-1030032211132321-3231100032200320"></a>

## exact_values property — body_matcher / 332101320320 / 4

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

<a id="canonical-0221021000002133-1000233002232022-0001123301310110-3011203123221323-3213200330312101-2310223322100112-3300212111213022-1331330321130110"></a>

<a id="canonical-3202010322033000-3210333203130232-0220112311133220-0321130110210321-2012022201013322-1320330321022120-1331011300312322-3002333100013021"></a>

## regex_values property — body_matcher / 332101320320 / 5

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

<a id="canonical-3222132232321323-3320210030331000-3333100032021211-0022002030300033-0013001210020111-1133031223023112-2023310001200222-3312331300011211"></a>

<a id="canonical-0203120323002123-1130311231022010-3220012112132102-2312030311301320-3212123203313331-1000301110021311-3311031011131022-1202222312103031"></a>

## transformers property — body_matcher / 332101320320 / 6

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1002323212030203-3233202220221122-0120001102113221-2122312231130221-3133112122133320-3233102220232002-3011032130332001-0131030321313231"></a>

## Next pages — body_matcher / 332101320320 / 7

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310333120120132-1111003300001203-0321202313333333-3231232303212002-2101123313003310-3331303211021031-1230232012322001-3101121232100001"></a>

## bot_action — bot_action / 323111110113 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- bot_action

<a id="canonical-0012123222233210-0203221210313202-0102211121020221-1200321221120113-2123021121210332-0020001102203112-3331332033333213-1112233020212130"></a>

Type: `"single"`. Computed.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"bot_skip_processing\",\"none\"]"
}
```

<a id="canonical-2012302031110220-1000233222200010-3120101331232200-2033302123211020-3110211220121322-0222223130213000-1030101120231223-2223133331022011"></a>

## Direct properties — bot_action / 323111110113 / 3

- [bot_skip_processing](data-sources--service_policy_rule--reference--group-001.md#canonical-3123330002023221-1203220000101333-3110103201331012-1110320320030010-2001322320030331-2013102122333203-0333110032110221-0033333030021210): complete subsection reference.

- [none](data-sources--service_policy_rule--reference--group-001.md#canonical-2213110333331103-2100331312312131-3212030302231233-2021213000333130-3100230230213320-1120023130222011-2333210022332221-0010211222000033): complete subsection reference.

<a id="canonical-0223130222212121-3131021202313222-3120313021010130-3132030020103311-2001220030133200-2313113231113130-1213233312003111-0313131332113033"></a>

## Next pages — bot_action / 323111110113 / 4

- [bot_action.bot_skip_processing](data-sources--service_policy_rule--reference--group-001.md#canonical-3123330002023221-1203220000101333-3110103201331012-1110320320030010-2001322320030331-2013102122333203-0333110032110221-0033333030021210)
- [bot_action.none](data-sources--service_policy_rule--reference--group-001.md#canonical-2213110333331103-2100331312312131-3212030302231233-2021213000333130-3100230230213320-1120023130222011-2333210022332221-0010211222000033)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3123330002023221-1203220000101333-3110103201331012-1110320320030010-2001322320030331-2013102122333203-0333110032110221-0033333030021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000012232002232-2212023132320230-3323330232113231-0201222010132033-1333310232321233-3003233000021131-3130230323122211-0012231221221033"></a>

## bot_action.bot_skip_processing — bot_skip_processing / 222011103112 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320)
- bot_action.bot_skip_processing

<a id="canonical-1131002201121333-1003201232021310-0321123333003223-1000202111223231-1001011313302301-1223332033122103-0201003032121122-3203220122021330"></a>

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

<a id="canonical-1012132021133300-1313102123122300-3303120022100031-1102222202221233-1111113200300110-0320120012101211-2310103021233311-1122322020232221"></a>

## Direct properties — bot_skip_processing / 222011103112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120222332201222-3110322013013130-2222100301320301-0011230332110113-1031330203113330-2203112101203213-1231231022333331-0103312220112001"></a>

## Next pages — bot_skip_processing / 222011103112 / 4

- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2213110333331103-2100331312312131-3212030302231233-2021213000333130-3100230230213320-1120023130222011-2333210022332221-0010211222000033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222222200002211-0121131220131002-3010021223012130-1031030322111132-2030020202303131-2033132201320310-0213233331001200-3303213122003213"></a>

## bot_action.none — none / 223031000003 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320)
- bot_action.none

<a id="canonical-3131121022332130-2130113020011323-1330122113121132-0232121102112303-3133131000100320-1323113023011033-0231000113330103-3303033032211221"></a>

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

<a id="canonical-1230303113330102-2031030322120122-2300100202201121-2113203112113310-2223101223310003-3302231211023211-0121331311100121-3123311222130103"></a>

## Direct properties — none / 223031000003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311001011100320-1200021130210311-3130100211323030-2322311232113130-0331211201031331-3313233223233132-3333130023310000-1120233223002201"></a>

## Next pages — none / 223031000003 / 4

- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0333301122101131-3120331133111122-3030213133031021-1111322320200213-2203212111231210-1021112003222010-3011030011331310-3302320221031101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322120320030222-3123213123330122-2011101202003101-3101333310012010-0322331131133000-1020131022330221-3133320100000100-0201101133331113"></a>

## client_name_matcher — client_name_matcher / 132313333223 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- client_name_matcher

<a id="canonical-0232200201211231-3003231012023103-1001131131111002-0223210230203122-3322021220320103-1300210332123320-3113122212120323-3300300232011012"></a>

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

<a id="canonical-2301303221100131-3123311133121221-3100302103332032-2022112332012202-3012033120321330-1130102123322000-1311002132123033-0011301000202112"></a>

## Direct properties — client_name_matcher / 132313333223 / 3

<a id="canonical-0200231123130022-3230132131321111-1310300322220212-1210111101011003-1230013222323022-0300300101220331-2330201031311100-2231032331132223"></a>

<a id="canonical-3011130232311112-3333301021003023-1310331102102120-0130332212303030-3333201202033012-3211313200230322-1113122231033111-3130200321200223"></a>

## exact_values property — client_name_matcher / 132313333223 / 4

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

<a id="canonical-0301331113130321-2132312221223123-0212121003003102-1223300001312223-2031210211220223-0210122102033220-1132312010321121-3301200000213210"></a>

<a id="canonical-2023221110101223-2103213020121123-0333313220233100-1122230202112020-2321120033103123-3031323023110211-1220312332331010-3121332231003232"></a>

## regex_values property — client_name_matcher / 132313333223 / 5

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

<a id="canonical-0312320310323133-3021022212101031-2213203133221223-0233212233020223-1013202322130301-1301212300112301-1331012322312231-3203201300222200"></a>

## Next pages — client_name_matcher / 132313333223 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1123330002113022-1313333303031311-1222120120201120-1231033101313011-2011232002032001-3330222101101201-3032001310131123-1223131301212333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010321020103330-3221003302212102-2311331113310220-3211223322131000-0100120001330302-1032320300201120-3112110103201000-2001322020023103"></a>

## client_selector — client_selector / 201320022303 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- client_selector

<a id="canonical-3320212100230333-3113230330223211-2130302022230311-2022021132302231-3222023001010121-2203230303122230-0032302303130333-2100111332210200"></a>

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

<a id="canonical-2012031110130003-3331323132201302-0130220300111303-3032313132023230-1110200012133133-0331301211213230-3311003011232320-2111011212230300"></a>

## Direct properties — client_selector / 201320022303 / 3

<a id="canonical-3032331103030210-1133203002110233-3210100131221322-3313121112120020-1333230331110320-2331123230333021-2211030333220220-0100011000023020"></a>

<a id="canonical-0333302103100132-0011132200000131-2103231333321213-0323031331131120-2211220100311001-1012203311113233-0303113112110323-3301122002302110"></a>

## expressions property — client_selector / 201320022303 / 4

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

<a id="canonical-3310132322333102-1233000000003201-1312002301132023-3013123021121031-0032313231113231-1110231333023301-2131110322303313-0100211002321323"></a>

## Next pages — client_selector / 201320022303 / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222032222233323-2010300100032000-2101121031130113-2223300213133023-3133030123302200-2012213221132131-1113103203112220-2311233121301121"></a>

## cookie_matchers — cookie_matchers / 101111103320 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- cookie_matchers

<a id="canonical-2323300210310032-0232010312211032-3121033321210300-2022102221102220-0032203331223022-3211001221322131-1113123033312333-0011021112330132"></a>

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2110222022220113-0210301213321331-0100131220002030-1132212020023301-0100030111112332-2220202332222321-1111232001233011-0231211121230032"></a>

## Direct properties — cookie_matchers / 101111103320 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-0230120212020131-1033121330312033-3110201302120201-3310233231133213-1222230212012131-2322323002102231-1030100200333302-1102202203032132): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-3211312011203132-2300211012110001-2012221303302211-0132111220031320-3210131112013323-2112201022201003-1120221123331110-1230223102132010): complete subsection reference.

<a id="canonical-0013002301203230-0332012310002202-0300023223121031-2223012012303111-1332302003133021-2031230100003122-0023211012301212-3220212211131322"></a>

<a id="canonical-2301221122213020-1101130130210033-1011232332231210-3003311310213213-2130133033001332-2130323300203300-1013310233231212-2113100331203100"></a>

## invert_matcher property — cookie_matchers / 101111103320 / 4

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

- [item](data-sources--service_policy_rule--reference--group-001.md#canonical-2022020011113232-1321230300321033-0022332103223100-0230313132332010-3212201313231232-2121022102113130-0230320201301312-2002300233232020): complete subsection reference.

<a id="canonical-2023132330032330-1023022213000230-2111212012321103-0001123203321323-0012321303302113-2300031001210303-0313130033120300-1012133113003210"></a>

<a id="canonical-1023201033002230-0200300233312323-2112220321212331-0101231312201003-0031000100310132-3022012021303202-1103302322230130-3010213111133311"></a>

## name property — cookie_matchers / 101111103320 / 5

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

<a id="canonical-1210031003320312-1111230022000322-3200120212001103-3210033112313100-2322302031011133-2000312133031221-3132221331223320-3130132210301311"></a>

## Next pages — cookie_matchers / 101111103320 / 6

- [cookie_matchers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-0230120212020131-1033121330312033-3110201302120201-3310233231133213-1222230212012131-2322323002102231-1030100200333302-1102202203032132)
- [cookie_matchers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-3211312011203132-2300211012110001-2012221303302211-0132111220031320-3210131112013323-2112201022201003-1120221123331110-1230223102132010)
- [cookie_matchers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-2022020011113232-1321230300321033-0022332103223100-0230313132332010-3212201313231232-2121022102113130-0230320201301312-2002300233232020)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0230120212020131-1033121330312033-3110201302120201-3310233231133213-1222230212012131-2322323002102231-1030100200333302-1102202203032132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133000310121022-2032231023220032-0130131333312211-1111312100201103-2033210011213032-2012331322100222-1030230132123310-1203120211313322"></a>

## cookie_matchers.check_not_present — check_not_present / 103101201130 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- cookie_matchers.check_not_present

<a id="canonical-2013132320301232-3133111311010300-0131133300111303-0311010012021230-2302102020220011-3203022100030002-3002212022103100-3013322312120001"></a>

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

<a id="canonical-0313211003121103-1231331102131323-0100231102102002-3211223322010031-1012300110320033-2131012331330120-2131313111232313-1310030313021330"></a>

## Direct properties — check_not_present / 103101201130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032123201210211-3130023301221322-3120203223133222-2012323020212131-3232202323330331-2331303303330003-3310210201120012-2002333013111132"></a>

## Next pages — check_not_present / 103101201130 / 4

- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3211312011203132-2300211012110001-2012221303302211-0132111220031320-3210131112013323-2112201022201003-1120221123331110-1230223102132010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310120100031322-1121012323120131-2300021331021312-2201312222313233-1121130323322332-1112022032020133-1100322011001100-0031123102213102"></a>

## cookie_matchers.check_present — check_present / 122103030133 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- cookie_matchers.check_present

<a id="canonical-2013311203120230-2331110032222211-2211311300323012-0323303022103013-1301230302222010-1203111023020030-0202233322230111-2321122102132121"></a>

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

<a id="canonical-3111212022301130-1121012323000002-1120131032011210-1021103320010320-3212213022123322-3113233213312222-1103003011212232-1312031102201322"></a>

## Direct properties — check_present / 122103030133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210232031202200-3331213021122322-3001003130221212-1300003032031101-3011202012220132-3313212320111311-1222202100010023-3221320323000012"></a>

## Next pages — check_present / 122103030133 / 4

- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2022020011113232-1321230300321033-0022332103223100-0230313132332010-3212201313231232-2121022102113130-0230320201301312-2002300233232020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130022023131212-2210103211113032-1131213022133013-2322121212122212-3201032133313233-1300001221112012-3020103312233210-3130313211000110"></a>

## cookie_matchers.item — item / 013103200002 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- cookie_matchers.item

<a id="canonical-2020030231312123-0023000321120302-0213232023301232-2220021010212001-2211203211133113-3331022220030311-3301332103320000-2032032333121012"></a>

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

<a id="canonical-0200000223032103-2021010202212222-1221312223210110-0130000033212332-3211230322102200-0321111310130032-3211123123113112-1020310103212301"></a>

## Direct properties — item / 013103200002 / 3

<a id="canonical-1032302220230311-2022010222133133-3213111132232232-2233113122201231-3211321020222131-1313033300112310-0132301200330333-2220122110223022"></a>

<a id="canonical-1221202113123202-2200022222302103-1310112021120202-3203102023233212-2121332111332300-2230131332112110-1112311031123020-1020233013230302"></a>

## exact_values property — item / 013103200002 / 4

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

<a id="canonical-2321312001123322-3201000321120301-1031113213030320-1120012302130210-2203133201322031-3232130202100110-3013211200132330-2212120103310032"></a>

<a id="canonical-1030132023320130-3013321333220003-3200123201201211-2022223112131133-0013101033010222-2110022230211010-2223221110213201-1333021210020032"></a>

## regex_values property — item / 013103200002 / 5

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

<a id="canonical-1113201331310030-1210331201233200-2121313113231310-1213131221212123-3312300012202031-2330012022333300-0100321112202220-1013313310110312"></a>

<a id="canonical-1101303212030102-0101312310013330-1023021122123211-3113032002112321-1023210102331333-0100111333003301-0020102221312022-2230210003003313"></a>

## transformers property — item / 013103200002 / 6

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1202202200132013-3310130232010000-1223211132331222-2312311300203220-2100300221100111-2232032000003312-0011331100220010-2202001020132223"></a>

## Next pages — item / 013103200002 / 7

- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3322032130021031-3102103112002010-3312231133031210-2322012100203213-0213021201303211-2313203320103201-1031123311112323-2112122121122111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313133003020220-0211120102211020-0212210221121001-0112023313301302-2211322013222113-1122213121202002-2033310111312000-0202331100222132"></a>

## domain_matcher — domain_matcher / 221210031223 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- domain_matcher

<a id="canonical-0312230300002332-0030301030103001-1131312203102213-0003101313212213-0231120303320021-2011323012110331-1302133110033011-3232301122011313"></a>

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

<a id="canonical-0003312031010232-1030210113312210-2111123020110312-0312110123032201-0031101302311000-3001302221022002-2013102200003211-0211110002100020"></a>

## Direct properties — domain_matcher / 221210031223 / 3

<a id="canonical-1230222333132221-1112132201011112-2213013102102201-1312001011313202-1303301303003213-2010233212133231-2111003202231021-1000210032013213"></a>

<a id="canonical-1012311303232022-2030301001031122-2030121120020122-2333130121123303-3130300102120132-0132121300320003-3010211331003321-1323210312300101"></a>

## exact_values property — domain_matcher / 221210031223 / 4

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

<a id="canonical-2201220032223120-1312231320011123-2301120323133003-3332231201202231-2133300030032210-2031301110132003-0133100101100123-0103121031011111"></a>

<a id="canonical-3132031322322220-0200130221030321-1331023113303123-0223310222112002-2110000202133010-1010201202021210-1223120323301300-2112131211021011"></a>

## regex_values property — domain_matcher / 221210031223 / 5

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

<a id="canonical-3210200330102022-2323121101200332-2320223013010322-2122133201001121-3112011201120121-3311012323201231-3032100010122103-3003322201302231"></a>

## Next pages — domain_matcher / 221210031223 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233121213021301-0302101300312103-1113022100230201-2011100320023102-3003100003000022-0321323132213003-0331223121020313-3231312123313201"></a>

## headers — headers / 031211123121 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- headers

<a id="canonical-1120302113020222-1102120012132113-3302321012332003-3120203331222113-0230323030021301-3103310320210203-3000013301131012-3321121202310100"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1231100203020001-0130212221210313-3032313330211210-0232301211120300-0212030132120231-0231012002032130-1203221130333000-0010333230330133"></a>

## Direct properties — headers / 031211123121 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2300201011320213-2033232121332232-1111030330032001-0020303322123320-2031120222332230-1010323130120303-1213320023310111-0201220122320212): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-0132310103132301-2333311002122003-2222131301231103-3033012203201001-3323231312121133-3103330320121331-3113200003103110-3030100103303203): complete subsection reference.

<a id="canonical-0012332132113303-0301233133332010-2213133232221002-0012132302313322-3132011113100132-0013231002310110-2222023101021222-2321011110300003"></a>

<a id="canonical-3032000210131003-0201121310022302-0133300023223323-1323102322202130-2032112220100203-3320122220000323-1101223220132002-0002300302221031"></a>

## invert_matcher property — headers / 031211123121 / 4

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

- [item](data-sources--service_policy_rule--reference--group-001.md#canonical-1332032310102012-3230200311001310-1333011211032011-1112210033213300-3231303221301312-0231013333210022-0032131333320303-1201002312122002): complete subsection reference.

<a id="canonical-3101220313113133-0100212120010231-1311331103213012-3010211033021202-1310031233120301-0011033333331311-0123203210203210-2021103133210120"></a>

<a id="canonical-1011202103330132-1212313023222131-0303310212113113-0320133011023000-0110310230201221-0310211301300021-2331200133020003-3103123021123210"></a>

## name property — headers / 031211123121 / 5

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

<a id="canonical-0020320132011120-0322331220231131-1322002303133011-0222130230220020-0300122100101330-1323220331220120-1310100331233003-3320010013031003"></a>

## Next pages — headers / 031211123121 / 6

- [headers.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2300201011320213-2033232121332232-1111030330032001-0020303322123320-2031120222332230-1010323130120303-1213320023310111-0201220122320212)
- [headers.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-0132310103132301-2333311002122003-2222131301231103-3033012203201001-3323231312121133-3103330320121331-3113200003103110-3030100103303203)
- [headers.item](data-sources--service_policy_rule--reference--group-001.md#canonical-1332032310102012-3230200311001310-1333011211032011-1112210033213300-3231303221301312-0231013333210022-0032131333320303-1201002312122002)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2300201011320213-2033232121332232-1111030330032001-0020303322123320-2031120222332230-1010323130120303-1213320023310111-0201220122320212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011001113033331-2122311230233332-0112302111311023-2223000203200230-3130211033203013-0202222323111303-1022232023310310-0111301231333303"></a>

## headers.check_not_present — check_not_present / 311030033223 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- headers.check_not_present

<a id="canonical-2211121111132120-3200232303122130-0002221133130300-3013312211133111-1323233221121201-3122133023222021-2022213122122232-0301111312000032"></a>

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

<a id="canonical-0030210233300200-0331302312031322-2131131112130213-2100230313221231-1331030332001320-2230033320032001-3232331100102332-3031013133221203"></a>

## Direct properties — check_not_present / 311030033223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302131301032112-0033121310001321-0311001323321030-3020220222002220-0201030122201021-3211201310002332-3201033221202221-2321310203233230"></a>

## Next pages — check_not_present / 311030033223 / 4

- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0132310103132301-2333311002122003-2222131301231103-3033012203201001-3323231312121133-3103330320121331-3113200003103110-3030100103303203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102313312230321-2233223023232233-2212113310210102-2132012010230223-1313300200212000-1122320320210212-0012212332322212-3311223033210030"></a>

## headers.check_present — check_present / 110200310003 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- headers.check_present

<a id="canonical-0003331212022223-3321211302032201-3300201323021101-0200110233313102-3232113101303102-3123112013002222-2013101120223220-1011000112321211"></a>

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

<a id="canonical-2120331300002010-1021113320023000-3223221201203101-1331212301230133-2333131120131112-2310330112333233-1020200323222312-0013112323101311"></a>

## Direct properties — check_present / 110200310003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320332012131332-0133102033133030-2010103033123311-0113103133101110-2130002202331011-3203223002222211-2101203201231121-1022021221212302"></a>

## Next pages — check_present / 110200310003 / 4

- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1332032310102012-3230200311001310-1333011211032011-1112210033213300-3231303221301312-0231013333210022-0032131333320303-1201002312122002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023200103321232-3002031222123111-0123022211031321-2231011220312132-1122202021033322-3333121133213101-2013230201311120-0010221313011221"></a>

## headers.item — item / 211220103301 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- headers.item

<a id="canonical-1303303111021002-2032232212333310-2130123011031101-3011020031031313-0132233222210002-2003023231000321-2113111020203200-3023310001101202"></a>

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

<a id="canonical-2223112023023312-2100101003201031-3231200320201312-0030311023110220-2332300221310312-3113022103312012-3111220120220212-1120010230002133"></a>

## Direct properties — item / 211220103301 / 3

<a id="canonical-1033001323232301-1320001330220133-0221101023030311-1131001223011132-2000202033331333-2133110120302000-2113103201330021-1001230021001030"></a>

<a id="canonical-3323232230232232-0012111123233111-3003320321322030-1310012121032012-3000220233101102-2220231102033300-0230311323123120-2102302232121331"></a>

## exact_values property — item / 211220103301 / 4

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

<a id="canonical-0033230213000030-2103223223221230-1030132002131133-3123002323132021-2022301211031311-2312023130322110-2233012320112330-3113130331320011"></a>

<a id="canonical-2030312132223203-0230312302033210-1211200110012022-2223031110311302-3102000113013300-2211002020013312-3221133032300013-0230331313310120"></a>

## regex_values property — item / 211220103301 / 5

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

<a id="canonical-0332223022221133-0311113233133330-3211212113222030-2131002021203202-0111011131023132-1221110103000313-3101010302330020-1100213200330010"></a>

<a id="canonical-3333210133003002-3000223010132002-1110313002022030-2210030011120010-2112310000333030-3120123001300322-1133011033203201-2313331100310103"></a>

## transformers property — item / 211220103301 / 6

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2223023300310322-1212201113112203-3000132031210133-3132203232231101-3101030310013132-2032132201322131-0111122013013311-0312003211023312"></a>

## Next pages — item / 211220103301 / 7

- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0230303322301220-0201213302213132-3013033111230100-1212133331313130-2022033111330210-3011103212333123-3113001121120131-3123200320022331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023312133321203-2223202130123321-2120230101103010-3223310111123100-0012220210111121-0212022322320003-1331033322223200-2122013232301220"></a>

## http_method — http_method / 332331221213 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- http_method

<a id="canonical-3301213020313120-0213220212231131-1012213020121123-0200300300300033-0021031321301311-2313202002332032-2223312031123311-2301031031122322"></a>

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

<a id="canonical-1112210101012310-2211230211332232-1323312312130113-3302123223001000-1220323102231002-2300212113220133-0111203212111102-3113031010131102"></a>

## Direct properties — http_method / 332331221213 / 3

<a id="canonical-2020310203000333-3013101230302023-0102020102330212-0223322032032120-1110132031023111-3213332333322010-0020132001113223-2122211130122210"></a>

<a id="canonical-0122103123012003-1210121320130122-2130020201320231-2012023110230200-0312100021213102-3221311010302222-0332020312321200-1103121202223122"></a>

## invert_matcher property — http_method / 332331221213 / 4

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

<a id="canonical-0220201002321202-0320310022002331-3302301333200231-2212102033302031-2220301001002031-0332012031120023-3002133220022232-2232110330331131"></a>

<a id="canonical-2302212130030130-1313200120331321-1030230301012303-2103221202111010-3001022123022012-3110133312113313-1122330220021010-0313200001202322"></a>

## methods property — http_method / 332331221213 / 5

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

<a id="canonical-1220301132330032-3133213332303330-1321113323111213-0000202323013103-3331202033003103-1132232010132230-0211212032333230-0212330313122200"></a>

## Next pages — http_method / 332331221213 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1030121033313010-0032013003231230-2231312010033301-0022213330223022-2212312013310022-0011013023030122-2322233132210211-3302310303213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233031232132212-0223113100013012-2232312333231310-3131131202233032-2320232301013003-2030230022102112-1130003001200003-3323103210003230"></a>

## ip_matcher — ip_matcher / 001121330001 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- ip_matcher

<a id="canonical-3302333201312031-2022013030010023-3201022223002000-2220101331302101-1013322320220110-1011232213020221-0303123223222203-3122103000203321"></a>

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

<a id="canonical-2220121300132023-2233003030200203-2100000023013321-0203323221333112-0112110321122203-3302201003121001-2312020120213120-1132020003213223"></a>

## Direct properties — ip_matcher / 001121330001 / 3

<a id="canonical-1023131331122331-1313112022223201-1002232212231222-2003102020213023-3222130130222231-2311222012130010-2111023300320102-0000302121320032"></a>

<a id="canonical-1322101321113102-2132330221112300-0011300011111333-0123231133103023-0121022102233032-3012300322302033-0032203222232011-2232000021113023"></a>

## invert_matcher property — ip_matcher / 001121330001 / 4

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

- [prefix_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-0212003310100320-2310130320000003-3310330100021131-0333300301130302-0231113230111322-1321211330032333-3311231333130020-0030032230203323): complete subsection reference.

<a id="canonical-0023001003033033-3112333012231230-2312312312001111-2203211113230022-0103111032333001-0332301332203031-0322231302312000-3130232232030231"></a>

## Next pages — ip_matcher / 001121330001 / 5

- [ip_matcher.prefix_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-0212003310100320-2310130320000003-3310330100021131-0333300301130302-0231113230111322-1321211330032333-3311231333130020-0030032230203323)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-0212003310100320-2310130320000003-3310330100021131-0333300301130302-0231113230111322-1321211330032333-3311231333130020-0030032230203323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130203003032332-0222011301030001-2022230301210303-3103221300331200-1202202033122132-0212320003032022-2321112132010022-3222300023133333"></a>

## ip_matcher.prefix_sets — prefix_sets / 032022110221 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1030121033313010-0032013003231230-2231312010033301-0022213330223022-2212312013310022-0011013023030122-2322233132210211-3302310303213130)
- ip_matcher.prefix_sets

<a id="canonical-1023033132101223-1111310323230100-2323130021321203-1122010213023233-1301113112022030-3303103300103201-1321303113123033-3122312320110313"></a>

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-3320213213113221-0132023112023100-2130230100322321-1012300332211212-3111000010112113-1300223012202003-1213121101012213-3120022222122020"></a>

## Direct properties — prefix_sets / 032022110221 / 3

<a id="canonical-2003303310031232-1230003000020013-2112013113331122-0033201111000100-2031212032223303-2131112331113323-1110203300321133-2211313112310213"></a>

<a id="canonical-1223010113002120-3302231311021003-2323300010220130-3112333131211301-0213120133022221-0021311101101001-3020130203013202-3203230101112102"></a>

## kind property — prefix_sets / 032022110221 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0223320332213011-1122210133111031-1311331131022110-0323211312312021-3321002322011312-2310320000020300-0231211311001003-2132023002310321"></a>

<a id="canonical-3010331311023022-0112211103220320-1033330201013022-2300321033031222-3130133022223222-1003220320012121-2200320210022023-2311200330233020"></a>

## name property — prefix_sets / 032022110221 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3102003200033120-1220322122232023-3010330210113033-0011213013131233-1111203003222221-2021131012223033-2121312231133223-0303313031032113"></a>

<a id="canonical-1122131220322312-1311013030113201-1233233112202211-3012101132213211-2212312211311232-3112113321022123-2332301301202323-1310333202210103"></a>

## namespace property — prefix_sets / 032022110221 / 6

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

<a id="canonical-2002303223331002-0223031001132100-2331230333103332-0012311220321030-2301323100201131-2302033101121023-0122101003032201-0210020232021021"></a>

<a id="canonical-0022211121132221-0132312303020100-1300003111100230-0022022211113220-3332003232112312-1210200223031120-1203201001030030-3223232122011100"></a>

## tenant property — prefix_sets / 032022110221 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3123031133003010-1111300000011120-0103331011211131-0101031322330102-2002332201223301-1332313311100012-3112132310122330-3011013231122231"></a>

<a id="canonical-0222002002212013-3003002212121013-3133203131212221-1303203201023232-0113112331200032-0201001300330311-3301303212120332-3321323010320321"></a>

## uid property — prefix_sets / 032022110221 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0102111000203321-3131312101133312-1032031300022000-0021332303331232-2222230320223031-0313132211102232-0312213332332020-2311030332000102"></a>

## Next pages — prefix_sets / 032022110221 / 9

- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1030121033313010-0032013003231230-2231312010033301-0022213330223022-2212312013310022-0011013023030122-2322233132210211-3302310303213130)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3311122011122130-1231221232013110-1321323211302210-3012120023003303-2303221232112331-3110301313311233-1110200101321222-0133333103002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123003101223120-1211332330321111-2302021123031132-1301102332113123-1320301330232131-3033223223012011-1322112113233322-2020211332321301"></a>

## ip_prefix_list — ip_prefix_list / 101013103001 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- ip_prefix_list

<a id="canonical-1212211231233102-1100131031220010-0321101033131300-2032320131002110-3002323203102331-0002021013112023-2033203132320102-0123012333023100"></a>

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

<a id="canonical-1013121113031003-2210030130311011-0320030002201021-2303002321233123-2301032032222011-2302102320213013-0102301000103133-3312133001130203"></a>

## Direct properties — ip_prefix_list / 101013103001 / 3

<a id="canonical-2211113211211023-0031202313100232-3032131001322010-1123023100011020-1203001220210332-1220330030311032-1202313331203213-1332303222011023"></a>

<a id="canonical-0220013233020123-3101332230032333-2220000100021203-1002220031012302-2002323333223230-3232213303123131-1301111112033011-1003013021203310"></a>

## invert_match property — ip_prefix_list / 101013103001 / 4

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

<a id="canonical-2032031130000130-0313110213110013-0130003132310212-2131023301201233-3033101033300232-1323310030301130-0300213213011022-2100230113031230"></a>

<a id="canonical-2200231301203330-3010123322233322-1201313311022200-0031103122322003-2011321131030132-2331010233012010-1300222320031012-1132102301312102"></a>

## ip_prefixes property — ip_prefix_list / 101013103001 / 5

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

<a id="canonical-3213113331110223-0220322012023302-0100000100023131-2223121221300333-2312003200220333-2311321031323300-0001110333030233-0101332120112200"></a>

## Next pages — ip_prefix_list / 101013103001 / 6

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3202102000312012-0320203211302133-2311121232212120-0010321012011000-1103221331232212-3320312321121320-0001000123312010-0122223300013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033212023111011-0302232101120111-1022232131313123-2220032031021311-3223102022331022-3100022221101311-0310132001221032-3110021333033203"></a>

## ip_threat_category_list — ip_threat_category_list / 112232123311 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- ip_threat_category_list

<a id="canonical-0232313301010221-2121221333323202-0022132312230200-1033310323333323-0202103311012132-3323213022030321-1001300131020020-0122112000033003"></a>

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

<a id="canonical-3200001122123323-0331011330233100-2230122120221211-1210120333322120-2011231031122302-1320310102201121-3013120210232003-2122310320322102"></a>

## Direct properties — ip_threat_category_list / 112232123311 / 3

<a id="canonical-0033223212103030-0323002230011333-1130010311310010-1030102300023201-3132312022120110-1001303023203221-0200123311230102-0120201102011323"></a>

<a id="canonical-3132333331201313-0001223210311023-2300033313220230-2210023101102111-1322321233210033-0331012302222113-3013020232200203-2121303310030103"></a>

## ip_threat_categories property — ip_threat_category_list / 112232123311 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1202122210320132-0013101021103312-0030220033003032-1222203001302130-1010301300231102-3031033032011222-2223002113200220-0303002120031011"></a>

## Next pages — ip_threat_category_list / 112232123311 / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3323210221002100-0120011222202210-3321322010321201-0231111312300102-1112122201231031-1313121202130301-0022111303332121-3233213330333003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111002112330023-2320002212012010-1110031022122121-2130022312322220-2230012331331230-2010212112121012-2022310230020311-3210201233133131"></a>

## ja4_tls_fingerprint — ja4_tls_fingerprint / 330330113202 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- ja4_tls_fingerprint

<a id="canonical-3101102133303110-1023320010002013-1113133022313303-2312010203332111-3121223232213313-0310022312200300-0022000232200213-1333113233203220"></a>

Type: `"single"`. Computed.

\[OneOf: ja4\_tls\_fingerprint, tls\_fingerprint\_matcher\] Extended version of JA3 that includes
additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a
different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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

- [ja4_tls_fingerprint](data-sources--service_policy_rule--reference--group-001.md#canonical-3101102133303110-1023320010002013-1113133022313303-2312010203332111-3121223232213313-0310022312200300-0022000232200213-1333113233203220)
- [tls_fingerprint_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-3210102313131132-1011210000200333-2001031333102312-0013313211310201-0323202000203330-1101132130313031-1020332221301212-0331111311133010)

Select alternatives according to the provider validators above.

<a id="canonical-0202122013322013-1303303022210213-1330103102002332-1113100333012023-2323233020200012-0322022330003033-2201002222221010-1233220013302002"></a>

## Direct properties — ja4_tls_fingerprint / 330330113202 / 3

<a id="canonical-0312302121301122-2233012201001130-0310221300331333-0121203032310022-0213130102203212-2212131331310110-3101202103332232-3122323020323101"></a>

<a id="canonical-3012232010013130-0033332011133221-3213012032010333-0132113303100230-2222330122330200-0000233123110022-0221200322200220-2330202221231011"></a>

## exact_values property — ja4_tls_fingerprint / 330330113202 / 4

Type: `["list", "string"]`. Computed.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3232331122332301-3210202132202223-2111210332222123-0113102213321331-3020102230013200-3120333301211103-1020020223130333-0333001203200201"></a>

## Next pages — ja4_tls_fingerprint / 330330113202 / 5

- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022022210333023-3133030022331122-3222100213300102-2321233012223130-2130212230021110-3320331203302012-1101111030003200-2232330311211300"></a>

## jwt_claims — jwt_claims / 123123211300 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- jwt_claims

<a id="canonical-2100031211011310-0122001300303022-2231031011001203-1323002110133232-1013211312323311-0322302202321333-3031220203311031-2232011200202223"></a>

Type: `"list"`. Computed.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true.

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

<a id="canonical-1302303200132113-3033232332323220-1012201220130323-2320123110300330-3321023133333331-1113330101020201-1011223201123320-0021020321320311"></a>

## Direct properties — jwt_claims / 123123211300 / 3

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2003201333002211-0023223200322233-0010233312332121-3020313223303203-0133200232203000-2132301332332020-2203031223031230-1221200233030211): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-3210311130033300-0112032022230200-1033313023001203-3332333220012220-2120333013230302-2112013101303033-0311032123313002-0301332120111212): complete subsection reference.

<a id="canonical-1230331313133033-1201303031330010-1202121223210013-3112030200232200-1010130021200202-3033300320233112-1320123102130311-3303230121202011"></a>

<a id="canonical-0223111311001123-3331101220102101-2311332002022021-0212221100201300-0131233101313101-3020212011103221-2330033233033002-1312030301223022"></a>

## invert_matcher property — jwt_claims / 123123211300 / 4

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

- [item](data-sources--service_policy_rule--reference--group-001.md#canonical-1311030333210012-2303333100333100-2121112112001202-0001002112331211-0220301003203002-2102202032032130-3003023230311210-0203331033021210): complete subsection reference.

<a id="canonical-2211210223011032-2300331220200201-0022021130020203-0212201311012110-2322123302130323-2111033121303333-0101331113200302-2012000000333032"></a>

<a id="canonical-2011211212031320-3303232230301232-1123121321033301-3201301011131022-1110021302113202-0303110303103210-1003101221030231-2021100210232120"></a>

## name property — jwt_claims / 123123211300 / 5

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

<a id="canonical-0231032210023311-0303203320120321-2001020132131330-1301133332300303-2002121032202321-2300122100133230-3030333001110123-3001322200032322"></a>

## Next pages — jwt_claims / 123123211300 / 6

- [jwt_claims.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2003201333002211-0023223200322233-0010233312332121-3020313223303203-0133200232203000-2132301332332020-2203031223031230-1221200233030211)
- [jwt_claims.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-3210311130033300-0112032022230200-1033313023001203-3332333220012220-2120333013230302-2112013101303033-0311032123313002-0301332120111212)
- [jwt_claims.item](data-sources--service_policy_rule--reference--group-001.md#canonical-1311030333210012-2303333100333100-2121112112001202-0001002112331211-0220301003203002-2102202032032130-3003023230311210-0203331033021210)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-2003201333002211-0023223200322233-0010233312332121-3020313223303203-0133200232203000-2132301332332020-2203031223031230-1221200233030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110303031232310-1000131020013330-1210203133020032-1300300030100002-2311113131001131-3331113011323000-1103001031123323-2331230103232323"></a>

## jwt_claims.check_not_present — check_not_present / 331333223022 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- jwt_claims.check_not_present

<a id="canonical-1120121133200333-2212031310311023-0002331000222231-0010330322231023-0020213210002121-3131330231121311-1230030020031312-3021301100200110"></a>

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

<a id="canonical-2231130110120201-3023020200123232-0110221122021133-1201023121203310-2130130311312220-1202231031130100-3000203330303103-0322010020312231"></a>

## Direct properties — check_not_present / 331333223022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330001303220200-1003113011003333-0301332103021231-1300000110133213-3021002213310332-2000211301121110-2223003330303211-3033103333303111"></a>

## Next pages — check_not_present / 331333223022 / 4

- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-3210311130033300-0112032022230200-1033313023001203-3332333220012220-2120333013230302-2112013101303033-0311032123313002-0301332120111212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301120122222302-3000210112220312-3111301110010313-1232000113220113-0220311022313301-0101310302230023-2232013002120311-1232203112011201"></a>

## jwt_claims.check_present — check_present / 030112103222 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- jwt_claims.check_present

<a id="canonical-1331221331111102-1102032022212000-0232000323111103-1000113012030123-3310133032313131-2100020123020213-3301021233310110-1032303101101021"></a>

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

<a id="canonical-0022100220201301-2033322212300113-1021223211232232-0213130311231030-2123333033011022-2232022021303030-2322322323310102-1123300331313231"></a>

## Direct properties — check_present / 030112103222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002001101303302-2002222031200101-2000222110210321-0123230322030102-2322012113332033-3320312030013002-1010232021120230-0001313023320123"></a>

## Next pages — check_present / 030112103222 / 4

- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)

<a id="canonical-1311030333210012-2303333100333100-2121112112001202-0001002112331211-0220301003203002-2102202032032130-3003023230311210-0203331033021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
