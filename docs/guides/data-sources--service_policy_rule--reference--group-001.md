---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- Property reference

<a id="canonical-2010320022113122-2232111033123230-2101133001312012-2202023001331332-0000322310122002-3203213211133133-1323102211133313-2312022131231313"></a>

### Direct properties for `xcsh_service_policy_rule`

<a id="canonical-3300221333210321-1011012232032232-1122022031322312-3002331123211102-3221020032021203-0122201022100231-2202213021200132-1120111321102322"></a>

#### `action` property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW|NEXT\_POLICY\] The rule action determines the disposition of the input request
API. If a policy matches a rule with an ALLOW action, the processing of the request proceeds
forward. If it matches a rule with a DENY action, the processing of the request is terminated and an
appropriate message/code returned to.. Possible values are \`DENY\`, \`ALLOW\`, \`NEXT\_POLICY\`.
Defaults to \`DENY\`.

Additional upstream details:

The rule action determines the disposition of the input request API. If it matches a rule with a
DENY action, the processing of the request is terminated and an appropriate message/code returned to
the originator. If it matches a rule with a NEXT\_POLICY\_SET action, evaluation of the current
policy set terminates and evaluation of the next policy set in the chain begins.

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

<a id="canonical-1331123322123232-0001023320322223-2302111121101101-1031230212123221-0200032302233322-1201232322122101-0202313033013201-0103230121011122"></a>

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

<a id="canonical-0202201010311130-3321012032331012-2310310100123300-1011323031001332-3101133110212202-0122202102130032-2011211000332202-2121223313213012"></a>

#### `client_name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0011100120123321-1120311230022001-3120003031322103-0231103032232101-1023031121331133-1103112023310220-0100323112031303-0222231120210022"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the ServicePolicyRule.

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

- [domain_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3322032130021031-3102103112002010-3312231133031210-2322012100203213-0213021201303211-2313203320103201-1031123311112323-2112122121122111): complete subsection reference.

<a id="canonical-3203330313130200-3202120103220101-1333122101113001-2133233330123133-0312110033231302-1003322332101230-2113131302122013-3012120021302223"></a>

<a id="canonical-2130203320000031-1211212000323121-3110000200230230-2330300221210300-1100321333331201-0123013111222110-1302133033213203-0101123111332111"></a>

#### `expiration_timestamp` property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3213201023013330-2010211320222120-0121221232023322-0030231132203023-1303213130013100-2001133303203031-2030222100310122-3332021113002330"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1030121033313010-0032013003231230-2231312010033301-0022213330223022-2212312013310022-0011013023030122-2322233132210211-3302310303213130): complete subsection reference.

- [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-3311122011122130-1231221232013110-1321323211302210-3012120023003303-2303221232112331-3110301313311233-1110200101321222-0133333103002120): complete subsection reference.

- [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-3202102000312012-0320203211302133-2311121232212120-0010321012011000-1103221331232212-3320312321121320-0001000123312010-0122223300013313): complete subsection reference.

- [ja4_tls_fingerprint](data-sources--service_policy_rule--reference--group-001.md#canonical-3323210221002100-0120011222202210-3321322010321201-0231111312300102-1112122201231031-1313121202130301-0022111303332121-3233213330333003): complete subsection reference.

- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312): complete subsection reference.

- [label_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3210311000320320-2030222121122100-0220110100321312-1121102131023230-3300133221311211-2200300323311103-1132001130323212-2010010230320023): complete subsection reference.

<a id="canonical-0232003211010103-3212200032012130-0033103220301110-1023021132213033-2010303221323033-3032202121111020-3122113102132303-3000031003003121"></a>

<a id="canonical-2300320200232023-2300201230130231-2301310111021111-3322122210230032-2002302000013132-3021330203222023-0331110223020220-1113322002303023"></a>

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

<a id="canonical-1101230321300123-1331021222121022-2201123223200312-3323333202100013-1222220113321013-1102310023211130-3231230301311112-2301221102220331"></a>

<a id="canonical-3130300310200012-3300211101101212-0121022111123231-1221122132131223-0130303221131302-3012102120122012-2203122123121111-0001310330331000"></a>

#### `log_rule_evaluation` property

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

- [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303): complete subsection reference.

<a id="canonical-3312013302110003-1100231322120110-3233233123303221-3011301232312012-1101231120132120-2000111221201301-1333311102032203-0203223100201020"></a>

<a id="canonical-0023331201230003-3002323112132321-3032122323133111-1031101011112131-1123020200122213-0303232311100213-1001031332121112-0100001002001313"></a>

#### `name` property

Type: `"string"`. Required.

Name of the ServicePolicyRule.

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

<a id="canonical-3130223300023011-2022011331301002-2202120023233300-1330202223101222-2103122300213021-2030120103003232-0010121010200222-2001121132220331"></a>

<a id="canonical-0110002230030001-0320232100331211-2123331101220032-0102312112032320-2200013313130222-0330110113220112-0220212110220123-1321303202132312"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the ServicePolicyRule exists.

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

- [path](data-sources--service_policy_rule--reference--group-001.md#canonical-1202323031323101-3010322010211020-2201013123220300-0120101231322302-1330322232232230-0002133000323121-2021123103113212-2023022232033013): complete subsection reference.

- [port_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2101210020323103-3303202202131321-1133232111123110-1302223220031001-1031030010123130-1010013200001221-1302232103001023-1320212133031120): complete subsection reference.

- [query_params](data-sources--service_policy_rule--reference--group-001.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203): complete subsection reference.

- [request_constraints](data-sources--service_policy_rule--reference--group-001.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010): complete subsection reference.

- [segment_policy](data-sources--service_policy_rule--reference--group-002.md#canonical-3110203330302013-0013231000222220-0332010312210230-3030022012013213-3313312233301211-2123100221121323-3220101200100332-3300123330120100): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--service_policy_rule--reference--group-002.md#canonical-0133301223123320-2232020220130121-1213202001010012-3222021111322322-0203313233311020-3313132002210111-1222220233313311-1202000232133022): complete subsection reference.

- [waf_action](data-sources--service_policy_rule--reference--group-002.md#canonical-1310200131200032-3220310022120032-2311002022323233-3020221313312231-2010023021321133-3101213020023220-2111322333220230-2120021231310322): complete subsection reference.

<a id="canonical-1203200033321233-3000232211322322-1320311100301032-3321010303012113-0101121322232302-1332232333212131-1332202202223310-2012202001303200"></a>

### All schema paths for `xcsh_service_policy_rule`

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
| `jwt_claims.item` | [jwt_claims.item](data-sources--service_policy_rule--reference--group-001.md#canonical-1211020202230012-3220003301020233-3203312001201133-3001032323312311-1110031110122001-3233023031130232-2231133233133311-3102023302121230) |
| `jwt_claims.item.exact_values` | [jwt_claims.item.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-2133301320131231-0020011102112213-2011122110332202-2333103013020020-3301311300322012-3103032201213131-1233011223002230-2122020100001123) |
| `jwt_claims.item.regex_values` | [jwt_claims.item.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-1332033123211330-0022220030300003-1121030333122222-0132221332332020-2102032112320103-3303031310013210-3323030010203022-3133103202032312) |
| `jwt_claims.item.transformers` | [jwt_claims.item.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-1101001201330102-0200132301102310-0123011303130023-2212101210003122-2321220202121011-1001031322130232-3123301302321232-0303313012123023) |
| `jwt_claims.name` | [jwt_claims.name](data-sources--service_policy_rule--reference--group-001.md#canonical-2211210223011032-2300331220200201-0022021130020203-0212201311012110-2322123302130323-2111033121303333-0101331113200302-2012000000333032) |
| `label_matcher` | [label_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0200330310302023-3300303133133011-2123020230313101-1331223303030333-1303021022210130-2310222013032010-2320211001011312-1130301211201100) |
| `label_matcher.keys` | [label_matcher.keys](data-sources--service_policy_rule--reference--group-001.md#canonical-1001130000233222-3122310003013300-1303221203220123-1323100333203231-2300311003330313-3220301303212000-0022220212232010-3022000123332132) |
| `labels` | [labels](data-sources--service_policy_rule--reference--group-001.md#canonical-0232003211010103-3212200032012130-0033103220301110-1023021132213033-2010303221323033-3032202121111020-3122113102132303-3000031003003121) |
| `log_rule_evaluation` | [log_rule_evaluation](data-sources--service_policy_rule--reference--group-001.md#canonical-1101230321300123-1331021222121022-2201123223200312-3323333202100013-1222220113321013-1102310023211130-3231230301311112-2301221102220331) |
| `mum_action` | [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-1232230311312032-2121331211020302-0213033312300333-0012232230033211-3200033003020130-2312132120000303-0211133100312120-1112302303011322) |
| `mum_action.default` | [mum_action.default](data-sources--service_policy_rule--reference--group-001.md#canonical-3033220123200322-2031112332001123-2020131332123131-2132100310012032-0330122211030322-1112132322010133-0123033030113010-1001102023302202) |
| `mum_action.skip_processing` | [mum_action.skip_processing](data-sources--service_policy_rule--reference--group-001.md#canonical-0100312101210231-0302220100200300-3002131133112312-0220212012323312-0100210000133313-0100021301232022-0211222230120232-0023333023103300) |
| `name` | [name](data-sources--service_policy_rule--reference--group-001.md#canonical-3312013302110003-1100231322120110-3233233123303221-3011301232312012-1101231120132120-2000111221201301-1333311102032203-0203223100201020) |
| `namespace` | [namespace](data-sources--service_policy_rule--reference--group-001.md#canonical-3130223300023011-2022011331301002-2202120023233300-1330202223101222-2103122300213021-2030120103003232-0010121010200222-2001121132220331) |
| `path` | [path](data-sources--service_policy_rule--reference--group-001.md#canonical-0031213201320113-0231310121110121-0121200331113332-0010202322232200-0011332112301101-0322223221223202-2103011210010201-0213023020212223) |
| `path.encoded_path_matcher` | [path.encoded_path_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2233111103220210-0022110311303002-1021121233222031-1032011321031322-3011332103003213-0010111122202122-0203220132302331-0132212011220102) |
| `path.exact_values` | [path.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-1021112133022313-1113200020120301-3112323111032321-1212111001101223-1223331200132322-2122031320002121-0031311213003312-3102303312330211) |
| `path.invert_matcher` | [path.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0122002222120103-2300311231001111-1311011013331230-2212113133301030-2200212210200301-3200133220230132-0032221313013111-2002222333212313) |
| `path.prefix_values` | [path.prefix_values](data-sources--service_policy_rule--reference--group-001.md#canonical-3323220100003020-2333331113121203-3310122032320133-0022011132123112-0122122331313132-2301220000322102-1220220211212231-3031330301200212) |
| `path.regex_values` | [path.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-2311313133011303-3332310211130130-0201203330202201-1300002230232211-0002001110301003-2323210233001032-3201002301122202-0121200000031030) |
| `path.suffix_values` | [path.suffix_values](data-sources--service_policy_rule--reference--group-001.md#canonical-1222230001213311-0322221330330021-2013011220303322-0002110130221311-0221001011101022-2230231000332100-1302022031220301-2022211030222111) |
| `path.transformers` | [path.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-0123102202300202-3231310213302101-3230011013211232-0303202020230313-1310020313203131-0130000132320312-3201220320000200-3231200022221202) |
| `port_matcher` | [port_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0131213012220100-3233021322333010-0232201212000122-1202123103003121-0133002011221000-3020313202313003-1303103303310102-0011031113301112) |
| `port_matcher.invert_matcher` | [port_matcher.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3330332300030201-2203330233023021-0002032132001020-2012201330321203-1020331230110211-1222103030122130-2321312211030033-3131103112310113) |
| `port_matcher.ports` | [port_matcher.ports](data-sources--service_policy_rule--reference--group-001.md#canonical-3101130102333312-2110110312330220-3020123011133333-0233110000213322-1222311330211123-0101303232310120-0122121013221113-3020002312112110) |
| `query_params` | [query_params](data-sources--service_policy_rule--reference--group-001.md#canonical-0110022223012231-1011022222022021-1231213210021012-2020023231312213-0120003203031103-1110212020200121-1320200132312313-1310013021311300) |
| `query_params.check_not_present` | [query_params.check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-3011101203123210-1300300210101020-3321130320311201-3310230230131332-2213302110000323-0122201221030113-2011323332311220-3000300112000112) |
| `query_params.check_present` | [query_params.check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-1002312133133323-0223313003230001-3331200333031002-1332121010001221-2012123200022211-2211201002322331-3223232333131320-1303131103120003) |
| `query_params.invert_matcher` | [query_params.invert_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1103033200210122-3201213211221303-2230013333312012-1222031203122022-2022101301223122-2330130211320000-2021230222233110-2301132123331121) |
| `query_params.item` | [query_params.item](data-sources--service_policy_rule--reference--group-001.md#canonical-0330101112120322-3333133122212133-2332122021103233-1322100002311122-3333313123111132-1133023013130202-0221213301303033-2102321000311213) |
| `query_params.item.exact_values` | [query_params.item.exact_values](data-sources--service_policy_rule--reference--group-001.md#canonical-0233120011232212-2331113232020021-3322333002132232-1133021001021230-3210112012102013-2212101320301333-1120222012313232-2232020233211120) |
| `query_params.item.regex_values` | [query_params.item.regex_values](data-sources--service_policy_rule--reference--group-001.md#canonical-1131323033011300-2123101133231321-3221130011313103-3132232220303320-3331131121200131-0002313032300221-2033320002333012-1331313001110000) |
| `query_params.item.transformers` | [query_params.item.transformers](data-sources--service_policy_rule--reference--group-001.md#canonical-0300320221022220-3000131220203013-1103213222132210-2120301022320023-2123003333033112-1000032112123101-0000111222300023-1001002002122313) |
| `query_params.key` | [query_params.key](data-sources--service_policy_rule--reference--group-001.md#canonical-0312121020113322-3203333312133313-1203103311320333-3121102221233032-3201130110310000-0000311232031333-3002323232211121-3023333131220320) |
| `request_constraints` | [request_constraints](data-sources--service_policy_rule--reference--group-001.md#canonical-0121303130321000-2010323021321020-0130112232322010-2222221123133001-3032121201323213-1210123132003033-2233210030211101-2031120112210020) |
| `request_constraints.max_cookie_count_exceeds` | [request_constraints.max_cookie_count_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-2133230320223220-1113123020232100-0233011100311311-2220223033232133-3111131323323231-3232322210332223-2332213312020201-3232132112311022) |
| `request_constraints.max_cookie_count_none` | [request_constraints.max_cookie_count_none](data-sources--service_policy_rule--reference--group-001.md#canonical-1030113020022100-3110132033201201-3012322102131031-2003100123311002-1013100132030132-0013103030302130-0022123120223321-2203020213100101) |
| `request_constraints.max_cookie_key_size_exceeds` | [request_constraints.max_cookie_key_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-1303231131001120-1001120031110232-1103010020023232-3331030223330223-1331132022322023-0322322123113202-0030300232030133-3003313021120013) |
| `request_constraints.max_cookie_key_size_none` | [request_constraints.max_cookie_key_size_none](data-sources--service_policy_rule--reference--group-001.md#canonical-1033333332111222-3003302032233333-0211010132033231-3310003311301003-3212033101301303-1032122303212112-2211301202111332-1031120033122313) |
| `request_constraints.max_cookie_value_size_exceeds` | [request_constraints.max_cookie_value_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-0103100211320112-3303303103023200-0221102113130302-2233122010100221-0132302032002032-0132232101102221-0211311320202031-3102102113001113) |
| `request_constraints.max_cookie_value_size_none` | [request_constraints.max_cookie_value_size_none](data-sources--service_policy_rule--reference--group-001.md#canonical-0210323111300301-1111012003030002-1322022032031323-3021122231020033-0221120021330323-3121231110103200-1201231031302223-2021232101312332) |
| `request_constraints.max_header_count_exceeds` | [request_constraints.max_header_count_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-1011320330133023-1012201121103201-0220020110013000-1103113123322021-0303331033111012-0222221110223300-0121300123013110-3012233211103301) |
| `request_constraints.max_header_count_none` | [request_constraints.max_header_count_none](data-sources--service_policy_rule--reference--group-001.md#canonical-1011200232203120-1120023222123130-3012332100333113-0203030232001031-2203132022312323-2223131332313013-2310122210301230-3211321030023300) |
| `request_constraints.max_header_key_size_exceeds` | [request_constraints.max_header_key_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-0021211210121302-2200300210020031-1011203031203000-1102310111013122-2210300122303131-1321322330121202-1101123303012123-1010010202232310) |
| `request_constraints.max_header_key_size_none` | [request_constraints.max_header_key_size_none](data-sources--service_policy_rule--reference--group-001.md#canonical-0101223120103230-1001211233301323-1112113320113113-1203303322022001-2011130311001020-3132003002210131-2331311332120212-1130111111331002) |
| `request_constraints.max_header_value_size_exceeds` | [request_constraints.max_header_value_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-3020010003200211-0221120102310101-0200113112132022-3220032210112210-1300312232100223-0312220230303033-3110111223313202-0233000323022122) |
| `request_constraints.max_header_value_size_none` | [request_constraints.max_header_value_size_none](data-sources--service_policy_rule--reference--group-001.md#canonical-3220222232320100-3012231320230101-0123211102311011-1301202121302323-0013311022300300-1220113302230332-0330210112021120-1301310000221100) |
| `request_constraints.max_parameter_count_exceeds` | [request_constraints.max_parameter_count_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-2230210011311131-3220103311102300-3112202003332123-2313200201113222-3011111012102202-2332010230231223-1312300231011212-2302110110222232) |
| `request_constraints.max_parameter_count_none` | [request_constraints.max_parameter_count_none](data-sources--service_policy_rule--reference--group-001.md#canonical-1300321310102013-2312333223021233-0221302313001122-0300013002001321-3023232221132022-0200210310231210-1323002132331321-2022020122222212) |
| `request_constraints.max_parameter_name_size_exceeds` | [request_constraints.max_parameter_name_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-2011121232102213-3111131112321201-0122200021020202-0232203201103211-1211301033333121-1201003213032123-1201033231132220-3230213123111223) |
| `request_constraints.max_parameter_name_size_none` | [request_constraints.max_parameter_name_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1312322021102121-2103201031100211-3230031321110122-3102130021333311-0100330121032110-2132301222032332-2213021332010103-2101111300003003) |
| `request_constraints.max_parameter_value_size_exceeds` | [request_constraints.max_parameter_value_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-1202300203103111-1022032020110102-0000101131000020-1323231311011110-0201130121012113-3001321201123102-3332012303122003-1231120333210300) |
| `request_constraints.max_parameter_value_size_none` | [request_constraints.max_parameter_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-2002332310321030-2000302113032022-3321111201101313-3332023001302200-0202033321102113-3011330102300031-0330322121001002-0303032201223003) |
| `request_constraints.max_query_size_exceeds` | [request_constraints.max_query_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-2111112211231200-3121212100122210-2203032122130232-1331022232111100-2303220210233302-1320120203011133-3332222213313110-2120333320201121) |
| `request_constraints.max_query_size_none` | [request_constraints.max_query_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1312001012111122-3011222231323112-0313303221200022-3123133210220220-2302201211112122-2021012313222222-3221332020130101-2231113210010302) |
| `request_constraints.max_request_line_size_exceeds` | [request_constraints.max_request_line_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-1102112313031123-1113223221133223-0021013202022231-0232200231230232-2121000203210332-0030122300330213-0313312231313112-1213332012132020) |
| `request_constraints.max_request_line_size_none` | [request_constraints.max_request_line_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0323333201031003-0221202303322213-1211000011202220-2103103033331130-0230033213322231-3022022322031102-2202101323201231-1310221322311022) |
| `request_constraints.max_request_size_exceeds` | [request_constraints.max_request_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-0223002231203300-3122020102321311-1002001001122023-1133210130323002-1103311212233233-3113200020230020-2233021302033331-3200001133030313) |
| `request_constraints.max_request_size_none` | [request_constraints.max_request_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3133031123331012-0233222031021311-2223123201320023-1333300122021021-3303210120101113-3223310112302002-0003022331123312-0323113122002112) |
| `request_constraints.max_url_size_exceeds` | [request_constraints.max_url_size_exceeds](data-sources--service_policy_rule--reference--group-001.md#canonical-2113220021310133-3312302133003221-2120222232020321-0201023231321213-2332101031203212-2330121103303231-2203230022033021-2222030023031102) |
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

<a id="canonical-3021332231300121-1013322201002112-2132213331221313-0333231230031023-0200113020331032-0333203121223033-3212000203002221-2232321123102030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_asn` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- any_asn

<a id="canonical-1103323332130021-1130010320022321-2023010233233023-0112310301203100-0302223100132002-2212120013210010-0233031113220223-3213322301131021"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_asn, asn\_list, asn\_matcher\] Enable this option

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

- [any_asn](data-sources--service_policy_rule--reference--group-001.md#canonical-1103323332130021-1130010320022321-2023010233233023-0112310301203100-0302223100132002-2212120013210010-0233031113220223-3213322301131021)
- [asn_list](data-sources--service_policy_rule--reference--group-001.md#canonical-1322323023222131-2322301032320000-2223112213032313-3033201110301033-2223100302210312-3211332033011002-1210033202322002-3303332112101301)
- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2121010320211310-2113000100221200-0120121023000110-0010200020000220-0231301010310211-3211123101233310-1200010331101032-2003100121101220)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201201330130100-1123001021023020-3303233322313203-3331102111111233-1313213321012112-3332023310232102-0333301231202200-2113111111230220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_client` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- any_client

<a id="canonical-2232333100231331-2001201301222032-1023312023330222-2113120310013130-2331123030332013-1210211003120133-0231011111012011-2101103222033221"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_client, client\_name, client\_name\_matcher, client\_selector,
ip\_threat\_category\_list\] Enable this option

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

- [any_client](data-sources--service_policy_rule--reference--group-001.md#canonical-2232333100231331-2001201301222032-1023312023330222-2113120310013130-2331123030332013-1210211003120133-0231011111012011-2101103222033221)
- [client_name](data-sources--service_policy_rule--reference--group-001.md#canonical-1221313132002113-0233102000032332-2030000020231122-2230220131010132-2101200200112002-3221330131232310-3033313133311110-2303132103310230)
- [client_name_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-0232200201211231-3003231012023103-1001131131111002-0223210230203122-3322021220320103-1300210332123320-3113122212120323-3300300232011012)
- [client_selector](data-sources--service_policy_rule--reference--group-001.md#canonical-3320212100230333-3113230330223211-2130302022230311-2022021132302231-3222023001010121-2203230303122230-0032302303130333-2100111332210200)
- [ip_threat_category_list](data-sources--service_policy_rule--reference--group-001.md#canonical-0232313301010221-2121221333323202-0022132312230200-1033310323333323-0202103311012132-3323213022030321-1001300131020020-0122112000033003)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010322331300203-1113233230123032-2023011223232330-0323222220231000-0113102103000012-1012103310023010-2111000200320203-1232223201332301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_ip` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- any_ip

<a id="canonical-0313320230011213-1302003001100002-0301310012032122-0221102332320102-1221122201331002-0333310301200222-3230230310132333-2103332320221303"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_ip, ip\_matcher, ip\_prefix\_list\] Enable this option

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

- [any_ip](data-sources--service_policy_rule--reference--group-001.md#canonical-0313320230011213-1302003001100002-0301310012032122-0221102332320102-1221122201331002-0333310301200222-3230230310132333-2103332320221303)
- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-3302333201312031-2022013030010023-3201022223002000-2220101331302101-1013322320220110-1011232213020221-0303123223222203-3122103000203321)
- [ip_prefix_list](data-sources--service_policy_rule--reference--group-001.md#canonical-1212211231233102-1100131031220010-0321101033131300-2032320131002110-3002323203102331-0002021013112023-2033203132320102-0123012333023100)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323220103100030-2300221222133301-0312321332223112-2330231103123033-2032200011132020-3100000303330210-1222301113122130-2312000321110100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_group_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- api_group_matcher

<a id="canonical-1320220012022000-1103133113330120-3330100332131202-1313123221132133-1000013231001122-2231030201200110-1012031012130032-1000221200100301"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0311210223312103-0133133313310033-1222310020223200-1200223002021000-3121132333331203-2231131331010010-2033232333200132-3213313120300031"></a>

### Direct properties for `api_group_matcher`

<a id="canonical-1121201002123021-2313002110321021-0021032111020011-1310102222322120-1100331102232330-2032212331121032-1221023112001333-2113131023001022"></a>

#### `api_group_matcher.invert_matcher` property

Type: `"bool"`. Computed.

Invert String Matcher. Invert the match result.

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

<a id="canonical-0311200232203032-3212320011113123-0321101003220321-3330003113122131-0113313231010023-1010133123202123-0021322021000221-1000032032300202"></a>

#### `api_group_matcher.match` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `arg_matchers` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- arg_matchers

<a id="canonical-2023212113302303-2202212212112323-0020110030001332-1231302201311002-1011021002323320-2303003201331232-0010200323010113-2102132333223122"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3321103222030332-1312223033200200-3023012302222230-0133100211121032-3100300233211011-0020030331212101-2202220100112010-2333312211210000"></a>

### Direct properties for `arg_matchers`

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-1123303301321301-1332102122030203-0103213021123112-1122022132212300-2000133122300001-1011300321201321-1200001201030111-1033220212221011): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2322323320313222-0233331221103131-1021102222110202-0010331131231122-1002310233303213-0020212331101022-3322233230331330-2323001000112201): complete subsection reference.

<a id="canonical-1330300031000022-1211111010033312-1002203330331233-1320233010320223-2202332230222233-1110323223211222-2010001003011330-1121011132230210"></a>

<a id="canonical-2301302000000211-3103200203012101-1211123003121131-1112001032211020-2301010320320333-3331323031110300-0121330020232331-0201320313302101"></a>

#### `arg_matchers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

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

<a id="canonical-1302332022203332-2233311121133323-0303220322201321-0013023123330202-3120020201203100-1002021131222223-2133020203110200-1322330211011333"></a>

#### `arg_matchers.name` property

Type: `"string"`. Computed.

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

<a id="canonical-1123303301321301-1332102122030203-0103213021123112-1122022132212300-2000133122300001-1011300321201321-1200001201030111-1033220212221011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `arg_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- arg_matchers.check_not_present

<a id="canonical-2022310320210323-3302312322030200-2022300112203032-3101232001201210-3021323131123130-1113003001210322-2022102001132111-0000332021311313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-2322323320313222-0233331221103131-1021102222110202-0010331131231122-1002310233303213-0020212331101022-3322233230331330-2323001000112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `arg_matchers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- arg_matchers.check_present

<a id="canonical-3333231321333211-2331112212120201-0220213022120303-3323221313310013-0212030203323020-1312332000133333-2222210230302223-3201013032131222"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-2332010133303230-0321112322213021-3021222010213233-2210110101312231-3321012133003211-2103322032020000-3210203232033232-2110301031331021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `arg_matchers.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [arg_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-0201003001102212-1330231130333231-3131312033232332-1303110320120133-3300322211211220-3130200000322013-3233103020220223-0012121000033112)
- arg_matchers.item

<a id="canonical-3000321222213202-1210333013030013-1022301313102211-1312332112212212-3023032002201020-2333123330013333-2201212321111313-0200122301202021"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3033311202031312-2131131032213033-1332102230011031-3200032100113013-0330102010320220-2101033323321130-1101111321002111-1131211100012103"></a>

### Direct properties for `arg_matchers.item`

<a id="canonical-0103000023210113-1323032230032211-3310113213303212-1331013203231220-0000113002133333-0100310230313330-3103033200121023-1233321211103221"></a>

#### `arg_matchers.item.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0011133023322123-3021321102033010-0103200130333133-1130030002002031-0001000033220230-2113013303321303-1133132032200233-0322330313212322"></a>

#### `arg_matchers.item.regex_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0112300013321332-2200022200001322-1332220013101021-2230213310001131-3323011130022131-2001130301222113-3300020310110031-1023120103321101"></a>

#### `arg_matchers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2323110012210121-3300120010211011-3021130303320203-1302110103133022-2103210011332101-3311131100232303-0022131322313002-2203010100023112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `asn_list` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- asn_list

<a id="canonical-1322323023222131-2322301032320000-2223112213032313-3033201110301033-2223100302210312-3211332033011002-1210033202322002-3303332112101301"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2002321300223121-2303321032200010-3330233320212201-2113132021030003-2221023102320030-2133113033020113-1033020331322002-0231322200301302"></a>

### Direct properties for `asn_list`

<a id="canonical-1022001300331020-2202331103312101-2313100103303220-0011132033111302-1002302021023200-1200131001012121-2303103033200030-1010210200311331"></a>

#### `asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

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

<a id="canonical-2123130101302113-2200222301333131-0110022202020010-1333032113332333-3123321211312310-0322220322033200-0100001233131121-1103233030221220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `asn_matcher` properties

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

<a id="canonical-0302103312101303-1230203300232202-3312232011332021-3211310102133203-2130332311030101-3131010103201010-1202310020003222-0323123010221023"></a>

### Direct properties for `asn_matcher`

- [asn_sets](data-sources--service_policy_rule--reference--group-001.md#canonical-3031222031221220-3103310132120011-3130222223310230-0202233201012300-0332220113201103-2031023123323212-1133222013221131-3300120301112032): complete subsection reference.

<a id="canonical-3031222031221220-3103310132120011-3130222223310230-0202233201012300-0332220113201103-2031023123323212-1133222013221131-3300120301112032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [asn_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-2123130101302113-2200222301333131-0110022202020010-1333032113332333-3123321211312310-0322220322033200-0100001233131121-1103233030221220)
- asn_matcher.asn_sets

<a id="canonical-2033120202232101-3022000030110322-1033022300320103-3111312033011103-3002032212313131-3002201210112300-1123220120233320-3101112332220223"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-0323022332131330-2301033112023202-2330012121030233-3113312321032001-2001000321312033-2030111020223122-1303231130313001-0332003200301101"></a>

### Direct properties for `asn_matcher.asn_sets`

<a id="canonical-2333310231033111-2010030030332020-0210003311313130-0302022110122101-0023232100332130-0101320310201110-0223021220121000-3322010011210220"></a>

#### `asn_matcher.asn_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1302212330223203-3232030233033131-1101200000012332-2321103101201313-0203113023310211-1130322023022331-2002323133330111-1212331133002311"></a>

<a id="canonical-3002113001312332-0120202103101123-0123323012010023-0202032012030121-1130123312021210-2320031111233013-0010323112321003-2203231011231223"></a>

#### `asn_matcher.asn_sets.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1210310130222022-0011121101002123-2233322001112023-1312322120033320-1130110002022222-2202200211220133-2321313231122332-2233111121012131"></a>

#### `asn_matcher.asn_sets.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-2312121211123100-3332221202221210-0003132231212111-2113113333130332-1021120312023123-3020001012003130-2030013021113321-0111101120323002"></a>

<a id="canonical-0203323212022132-0322012221331331-3030033323310030-1001333330131033-2130311321320002-3120103202031221-3210003320201100-0002301313033100"></a>

#### `asn_matcher.asn_sets.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1213212002301301-1303022130301333-1233233200002221-2332123311333303-1111302221010210-1003211031120010-3322021330231300-3103212221321333"></a>

#### `asn_matcher.asn_sets.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0023002331223020-0113212003033021-2231301311112131-1031101231203003-0013032323230200-0030122222320202-2000131202022210-1112020323222210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `body_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- body_matcher

<a id="canonical-0322031300230322-3320223211302110-0200330201013300-3331113030130121-2111032033313121-1300202310001211-1220222000133023-1102210221232130"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1031203031211110-3302331302122031-0130321133331033-3312303222321220-3033101001100231-2133013010300320-2230200023032031-3002020213332301"></a>

### Direct properties for `body_matcher`

<a id="canonical-3303022213021030-0110021020233322-3330101322202001-0310332121111113-1233011221312012-3312303233231022-0013323133332333-3100011020033310"></a>

#### `body_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1322320211220112-3130133102313211-2330133320333130-2220101003321211-2301013110320200-3012022100220201-3322011331333322-1111013011330300"></a>

#### `body_matcher.regex_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3233100023032302-3023102312330211-2013323131211003-3010021123301102-1303013123132003-2233232313032101-1030032211132321-3231100032200320"></a>

#### `body_matcher.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_action` properties

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

<a id="canonical-1310333120120132-1111003300001203-0321202313333333-3231232303212002-2101123313003310-3331303211021031-1230232012322001-3101121232100001"></a>

### Direct properties for `bot_action`

- [bot_skip_processing](data-sources--service_policy_rule--reference--group-001.md#canonical-3123330002023221-1203220000101333-3110103201331012-1110320320030010-2001322320030331-2013102122333203-0333110032110221-0033333030021210): complete subsection reference.

- [none](data-sources--service_policy_rule--reference--group-001.md#canonical-2213110333331103-2100331312312131-3212030302231233-2021213000333130-3100230230213320-1120023130222011-2333210022332221-0010211222000033): complete subsection reference.

<a id="canonical-3123330002023221-1203220000101333-3110103201331012-1110320320030010-2001322320030331-2013102122333203-0333110032110221-0033333030021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_action.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320)
- bot_action.bot_skip_processing

<a id="canonical-1131002201121333-1003201232021310-0321123333003223-1000202111223231-1001011313302301-1223332033122103-0201003032121122-3203220122021330"></a>

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

<a id="canonical-2213110333331103-2100331312312131-3212030302231233-2021213000333130-3100230230213320-1120023130222011-2333210022332221-0010211222000033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_action.none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [bot_action](data-sources--service_policy_rule--reference--group-001.md#canonical-1200211100231112-1113202321220331-0232133300012233-2333023320110203-1210230311121300-0320030320011301-2303123233001302-0200221101032320)
- bot_action.none

<a id="canonical-3131121022332130-2130113020011323-1330122113121132-0232121102112303-3133131000100320-1323113023011033-0231000113330103-3303033032211221"></a>

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

<a id="canonical-0333301122101131-3120331133111122-3030213133031021-1111322320200213-2203212111231210-1021112003222010-3011030011331310-3302320221031101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_name_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- client_name_matcher

<a id="canonical-0232200201211231-3003231012023103-1001131131111002-0223210230203122-3322021220320103-1300210332123320-3113122212120323-3300300232011012"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2322120320030222-3123213123330122-2011101202003101-3101333310012010-0322331131133000-1020131022330221-3133320100000100-0201101133331113"></a>

### Direct properties for `client_name_matcher`

<a id="canonical-0200231123130022-3230132131321111-1310300322220212-1210111101011003-1230013222323022-0300300101220331-2330201031311100-2231032331132223"></a>

#### `client_name_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2301303221100131-3123311133121221-3100302103332032-2022112332012202-3012033120321330-1130102123322000-1311002132123033-0011301000202112"></a>

#### `client_name_matcher.regex_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1123330002113022-1313333303031311-1222120120201120-1231033101313011-2011232002032001-3330222101101201-3032001310131123-1223131301212333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_selector` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- client_selector

<a id="canonical-3320212100230333-3113230330223211-2130302022230311-2022021132302231-3222023001010121-2203230303122230-0032302303130333-2100111332210200"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3010321020103330-3221003302212102-2311331113310220-3211223322131000-0100120001330302-1032320300201120-3112110103201000-2001322020023103"></a>

### Direct properties for `client_selector`

<a id="canonical-3032331103030210-1133203002110233-3210100131221322-3313121112120020-1333230331110320-2331123230333021-2211030333220220-0100011000023020"></a>

#### `client_selector.expressions` property

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

<a id="canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_matchers` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- cookie_matchers

<a id="canonical-2323300210310032-0232010312211032-3121033321210300-2022102221102220-0032203331223022-3211001221322131-1113123033312333-0011021112330132"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3222032222233323-2010300100032000-2101121031130113-2223300213133023-3133030123302200-2012213221132131-1113103203112220-2311233121301121"></a>

### Direct properties for `cookie_matchers`

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-0230120212020131-1033121330312033-3110201302120201-3310233231133213-1222230212012131-2322323002102231-1030100200333302-1102202203032132): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-3211312011203132-2300211012110001-2012221303302211-0132111220031320-3210131112013323-2112201022201003-1120221123331110-1230223102132010): complete subsection reference.

<a id="canonical-0013002301203230-0332012310002202-0300023223121031-2223012012303111-1332302003133021-2031230100003122-0023211012301212-3220212211131322"></a>

<a id="canonical-2110222022220113-0210301213321331-0100131220002030-1132212020023301-0100030111112332-2220202332222321-1111232001233011-0231211121230032"></a>

#### `cookie_matchers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

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

<a id="canonical-2301221122213020-1101130130210033-1011232332231210-3003311310213213-2130133033001332-2130323300203300-1013310233231212-2113100331203100"></a>

#### `cookie_matchers.name` property

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

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

<a id="canonical-0230120212020131-1033121330312033-3110201302120201-3310233231133213-1222230212012131-2322323002102231-1030100200333302-1102202203032132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- cookie_matchers.check_not_present

<a id="canonical-2013132320301232-3133111311010300-0131133300111303-0311010012021230-2302102020220011-3203022100030002-3002212022103100-3013322312120001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-3211312011203132-2300211012110001-2012221303302211-0132111220031320-3210131112013323-2112201022201003-1120221123331110-1230223102132010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- cookie_matchers.check_present

<a id="canonical-2013311203120230-2331110032222211-2211311300323012-0323303022103013-1301230302222010-1203111023020030-0202233322230111-2321122102132121"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-2022020011113232-1321230300321033-0022332103223100-0230313132332010-3212201313231232-2121022102113130-0230320201301312-2002300233232020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [cookie_matchers](data-sources--service_policy_rule--reference--group-001.md#canonical-2031203121000110-0123221133213200-1331032330313020-0123323130120121-1100231122132111-1132213020331101-0301132221212130-3033213103222232)
- cookie_matchers.item

<a id="canonical-2020030231312123-0023000321120302-0213232023301232-2220021010212001-2211203211133113-3331022220030311-3301332103320000-2032032333121012"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3130022023131212-2210103211113032-1131213022133013-2322121212122212-3201032133313233-1300001221112012-3020103312233210-3130313211000110"></a>

### Direct properties for `cookie_matchers.item`

<a id="canonical-1032302220230311-2022010222133133-3213111132232232-2233113122201231-3211321020222131-1313033300112310-0132301200330333-2220122110223022"></a>

#### `cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0200000223032103-2021010202212222-1221312223210110-0130000033212332-3211230322102200-0321111310130032-3211123123113112-1020310103212301"></a>

#### `cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1221202113123202-2200022222302103-1310112021120202-3203102023233212-2121332111332300-2230131332112110-1112311031123020-1020233013230302"></a>

#### `cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3322032130021031-3102103112002010-3312231133031210-2322012100203213-0213021201303211-2313203320103201-1031123311112323-2112122121122111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domain_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- domain_matcher

<a id="canonical-0312230300002332-0030301030103001-1131312203102213-0003101313212213-0231120303320021-2011323012110331-1302133110033011-3232301122011313"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0313133003020220-0211120102211020-0212210221121001-0112023313301302-2211322013222113-1122213121202002-2033310111312000-0202331100222132"></a>

### Direct properties for `domain_matcher`

<a id="canonical-1230222333132221-1112132201011112-2213013102102201-1312001011313202-1303301303003213-2010233212133231-2111003202231021-1000210032013213"></a>

#### `domain_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0003312031010232-1030210113312210-2111123020110312-0312110123032201-0031101302311000-3001302221022002-2013102200003211-0211110002100020"></a>

#### `domain_matcher.regex_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `headers` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- headers

<a id="canonical-1120302113020222-1102120012132113-3302321012332003-3120203331222113-0230323030021301-3103310320210203-3000013301131012-3321121202310100"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0233121213021301-0302101300312103-1113022100230201-2011100320023102-3003100003000022-0321323132213003-0331223121020313-3231312123313201"></a>

### Direct properties for `headers`

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2300201011320213-2033232121332232-1111030330032001-0020303322123320-2031120222332230-1010323130120303-1213320023310111-0201220122320212): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-0132310103132301-2333311002122003-2222131301231103-3033012203201001-3323231312121133-3103330320121331-3113200003103110-3030100103303203): complete subsection reference.

<a id="canonical-0012332132113303-0301233133332010-2213133232221002-0012132302313322-3132011113100132-0013231002310110-2222023101021222-2321011110300003"></a>

<a id="canonical-1231100203020001-0130212221210313-3032313330211210-0232301211120300-0212030132120231-0231012002032130-1203221130333000-0010333230330133"></a>

#### `headers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

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

<a id="canonical-3032000210131003-0201121310022302-0133300023223323-1323102322202130-2032112220100203-3320122220000323-1101223220132002-0002300302221031"></a>

#### `headers.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

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

<a id="canonical-2300201011320213-2033232121332232-1111030330032001-0020303322123320-2031120222332230-1010323130120303-1213320023310111-0201220122320212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `headers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- headers.check_not_present

<a id="canonical-2211121111132120-3200232303122130-0002221133130300-3013312211133111-1323233221121201-3122133023222021-2022213122122232-0301111312000032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-0132310103132301-2333311002122003-2222131301231103-3033012203201001-3323231312121133-3103330320121331-3113200003103110-3030100103303203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `headers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- headers.check_present

<a id="canonical-0003331212022223-3321211302032201-3300201323021101-0200110233313102-3232113101303102-3123112013002222-2013101120223220-1011000112321211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-1332032310102012-3230200311001310-1333011211032011-1112210033213300-3231303221301312-0231013333210022-0032131333320303-1201002312122002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `headers.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [headers](data-sources--service_policy_rule--reference--group-001.md#canonical-3021020212103033-0302112131332010-2112310002312102-0122011210200111-2220322223111311-1313202112311330-2312330110233111-0323320133333020)
- headers.item

<a id="canonical-1303303111021002-2032232212333310-2130123011031101-3011020031031313-0132233222210002-2003023231000321-2113111020203200-3023310001101202"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0023200103321232-3002031222123111-0123022211031321-2231011220312132-1122202021033322-3333121133213101-2013230201311120-0010221313011221"></a>

### Direct properties for `headers.item`

<a id="canonical-1033001323232301-1320001330220133-0221101023030311-1131001223011132-2000202033331333-2133110120302000-2113103201330021-1001230021001030"></a>

#### `headers.item.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2223112023023312-2100101003201031-3231200320201312-0030311023110220-2332300221310312-3113022103312012-3111220120220212-1120010230002133"></a>

#### `headers.item.regex_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3323232230232232-0012111123233111-3003320321322030-1310012121032012-3000220233101102-2220231102033300-0230311323123120-2102302232121331"></a>

#### `headers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0230303322301220-0201213302213132-3013033111230100-1212133331313130-2022033111330210-3011103212333123-3113001121120131-3123200320022331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_method` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- http_method

<a id="canonical-3301213020313120-0213220212231131-1012213020121123-0200300300300033-0021031321301311-2313202002332032-2223312031123311-2301031031122322"></a>

Type: `"single"`. Computed.

An HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
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

<a id="canonical-1023312133321203-2223202130123321-2120230101103010-3223310111123100-0012220210111121-0212022322320003-1331033322223200-2122013232301220"></a>

### Direct properties for `http_method`

<a id="canonical-2020310203000333-3013101230302023-0102020102330212-0223322032032120-1110132031023111-3213332333322010-0020132001113223-2122211130122210"></a>

#### `http_method.invert_matcher` property

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-1112210101012310-2211230211332232-1323312312130113-3302123223001000-1220323102231002-2300212113220133-0111203212111102-3113031010131102"></a>

#### `http_method.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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

<a id="canonical-1030121033313010-0032013003231230-2231312010033301-0022213330223022-2212312013310022-0011013023030122-2322233132210211-3302310303213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- ip_matcher

<a id="canonical-3302333201312031-2022013030010023-3201022223002000-2220101331302101-1013322320220110-1011232213020221-0303123223222203-3122103000203321"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3233031232132212-0223113100013012-2232312333231310-3131131202233032-2320232301013003-2030230022102112-1130003001200003-3323103210003230"></a>

### Direct properties for `ip_matcher`

<a id="canonical-1023131331122331-1313112022223201-1002232212231222-2003102020213023-3222130130222231-2311222012130010-2111023300320102-0000302121320032"></a>

#### `ip_matcher.invert_matcher` property

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

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

<a id="canonical-0212003310100320-2310130320000003-3310330100021131-0333300301130302-0231113230111322-1321211330032333-3311231333130020-0030032230203323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [ip_matcher](data-sources--service_policy_rule--reference--group-001.md#canonical-1030121033313010-0032013003231230-2231312010033301-0022213330223022-2212312013310022-0011013023030122-2322233132210211-3302310303213130)
- ip_matcher.prefix_sets

<a id="canonical-1023033132101223-1111310323230100-2323130021321203-1122010213023233-1301113112022030-3303103300103201-1321303113123033-3122312320110313"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-2130203003032332-0222011301030001-2022230301210303-3103221300331200-1202202033122132-0212320003032022-2321112132010022-3222300023133333"></a>

### Direct properties for `ip_matcher.prefix_sets`

<a id="canonical-2003303310031232-1230003000020013-2112013113331122-0033201111000100-2031212032223303-2131112331113323-1110203300321133-2211313112310213"></a>

#### `ip_matcher.prefix_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0223320332213011-1122210133111031-1311331131022110-0323211312312021-3321002322011312-2310320000020300-0231211311001003-2132023002310321"></a>

<a id="canonical-3320213213113221-0132023112023100-2130230100322321-1012300332211212-3111000010112113-1300223012202003-1213121101012213-3120022222122020"></a>

#### `ip_matcher.prefix_sets.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1223010113002120-3302231311021003-2323300010220130-3112333131211301-0213120133022221-0021311101101001-3020130203013202-3203230101112102"></a>

#### `ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-2002303223331002-0223031001132100-2331230333103332-0012311220321030-2301323100201131-2302033101121023-0122101003032201-0210020232021021"></a>

<a id="canonical-3010331311023022-0112211103220320-1033330201013022-2300321033031222-3130133022223222-1003220320012121-2200320210022023-2311200330233020"></a>

#### `ip_matcher.prefix_sets.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1122131220322312-1311013030113201-1233233112202211-3012101132213211-2212312211311232-3112113321022123-2332301301202323-1310333202210103"></a>

#### `ip_matcher.prefix_sets.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3311122011122130-1231221232013110-1321323211302210-3012120023003303-2303221232112331-3110301313311233-1110200101321222-0133333103002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_prefix_list` properties

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

<a id="canonical-2123003101223120-1211332330321111-2302021123031132-1301102332113123-1320301330232131-3033223223012011-1322112113233322-2020211332321301"></a>

### Direct properties for `ip_prefix_list`

<a id="canonical-2211113211211023-0031202313100232-3032131001322010-1123023100011020-1203001220210332-1220330030311032-1202313331203213-1332303222011023"></a>

#### `ip_prefix_list.invert_match` property

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="canonical-1013121113031003-2210030130311011-0320030002201021-2303002321233123-2301032032222011-2302102320213013-0102301000103133-3312133001130203"></a>

#### `ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

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

<a id="canonical-3202102000312012-0320203211302133-2311121232212120-0010321012011000-1103221331232212-3320312321121320-0001000123312010-0122223300013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- ip_threat_category_list

<a id="canonical-0232313301010221-2121221333323202-0022132312230200-1033310323333323-0202103311012132-3323213022030321-1001300131020020-0122112000033003"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

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

<a id="canonical-0033212023111011-0302232101120111-1022232131313123-2220032031021311-3223102022331022-3100022221101311-0310132001221032-3110021333033203"></a>

### Direct properties for `ip_threat_category_list`

<a id="canonical-0033223212103030-0323002230011333-1130010311310010-1030102300023201-3132312022120110-1001303023203221-0200123311230102-0120201102011323"></a>

#### `ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3323210221002100-0120011222202210-3321322010321201-0231111312300102-1112122201231031-1313121202130301-0022111303332121-3233213330333003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ja4_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- ja4_tls_fingerprint

<a id="canonical-3101102133303110-1023320010002013-1113133022313303-2312010203332111-3121223232213313-0310022312200300-0022000232200213-1333113233203220"></a>

Type: `"single"`. Computed.

\[OneOf: ja4\_tls\_fingerprint, tls\_fingerprint\_matcher\] Extended version of JA3 that includes
additional fields for more comprehensive fingerprinting of SSL/TLS clients and potentially has a
different structure and length.

Additional upstream details:

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

<a id="canonical-1111002112330023-2320002212012010-1110031022122121-2130022312322220-2230012331331230-2010212112121012-2022310230020311-3210201233133131"></a>

### Direct properties for `ja4_tls_fingerprint`

<a id="canonical-0312302121301122-2233012201001130-0310221300331333-0121203032310022-0213130102203212-2212131331310110-3101202103332232-3122323020323101"></a>

#### `ja4_tls_fingerprint.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_claims` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- jwt_claims

<a id="canonical-2100031211011310-0122001300303022-2231031011001203-1323002110133232-1013211312323311-0322302202321333-3031220203311031-2232011200202223"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2022022210333023-3133030022331122-3222100213300102-2321233012223130-2130212230021110-3320331203302012-1101111030003200-2232330311211300"></a>

### Direct properties for `jwt_claims`

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2003201333002211-0023223200322233-0010233312332121-3020313223303203-0133200232203000-2132301332332020-2203031223031230-1221200233030211): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-3210311130033300-0112032022230200-1033313023001203-3332333220012220-2120333013230302-2112013101303033-0311032123313002-0301332120111212): complete subsection reference.

<a id="canonical-1230331313133033-1201303031330010-1202121223210013-3112030200232200-1010130021200202-3033300320233112-1320123102130311-3303230121202011"></a>

<a id="canonical-1302303200132113-3033232332323220-1012201220130323-2320123110300330-3321023133333331-1113330101020201-1011223201123320-0021020321320311"></a>

#### `jwt_claims.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

<a id="canonical-0223111311001123-3331101220102101-2311332002022021-0212221100201300-0131233101313101-3020212011103221-2330033233033002-1312030301223022"></a>

#### `jwt_claims.name` property

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

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

<a id="canonical-2003201333002211-0023223200322233-0010233312332121-3020313223303203-0133200232203000-2132301332332020-2203031223031230-1221200233030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- jwt_claims.check_not_present

<a id="canonical-1120121133200333-2212031310311023-0002331000222231-0010330322231023-0020213210002121-3131330231121311-1230030020031312-3021301100200110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-3210311130033300-0112032022230200-1033313023001203-3332333220012220-2120333013230302-2112013101303033-0311032123313002-0301332120111212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- jwt_claims.check_present

<a id="canonical-1331221331111102-1102032022212000-0232000323111103-1000113012030123-3310133032313131-2100020123020213-3301021233310110-1032303101101021"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-1311030333210012-2303333100333100-2121112112001202-0001002112331211-0220301003203002-2102202032032130-3003023230311210-0203331033021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_claims.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [jwt_claims](data-sources--service_policy_rule--reference--group-001.md#canonical-3233303020320223-0201322230301223-1111013223211322-2311021222122102-0313111032112232-3202102031123001-0123113023331220-1103011102320312)
- jwt_claims.item

<a id="canonical-1211020202230012-3220003301020233-3203312001201133-3001032323312311-1110031110122001-3233023031130232-2231133233133311-3102023302121230"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3101203321210221-0110210200130002-0133300232120301-1132333333302213-0023222223113112-3122300010121033-3120102131012231-2221232213013212"></a>

### Direct properties for `jwt_claims.item`

<a id="canonical-2133301320131231-0020011102112213-2011122110332202-2333103013020020-3301311300322012-3103032201213131-1233011223002230-2122020100001123"></a>

#### `jwt_claims.item.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1332033123211330-0022220030300003-1121030333122222-0132221332332020-2102032112320103-3303031310013210-3323030010203022-3133103202032312"></a>

<a id="canonical-0301002301222012-3230101302310330-0012030321013132-2112333231313221-2210203131311320-2331130322201020-1301010013101223-1120022312030333"></a>

#### `jwt_claims.item.regex_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1101001201330102-0200132301102310-0123011303130023-2212101210003122-2321220202121011-1001031322130232-3123301302321232-0303313012123023"></a>

<a id="canonical-0002020211331233-2223330323101223-0322302330121032-0222030303233302-1332312022303031-2103212023003222-3332023010023123-1211313011001203"></a>

#### `jwt_claims.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3210311000320320-2030222121122100-0220110100321312-1121102131023230-3300133221311211-2200300323311103-1132001130323212-2010010230320023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `label_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- label_matcher

<a id="canonical-0200330310302023-3300303133133011-2123020230313101-1331223303030333-1303021022210130-2310222013032010-2320211001011312-1130301211201100"></a>

Type: `"single"`. Computed.

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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

<a id="canonical-1201100101223030-1222210022200130-3110103013323310-3231302113001013-0000120020213213-2110323322031023-2122323031023133-3132033321112122"></a>

### Direct properties for `label_matcher`

<a id="canonical-1001130000233222-3122310003013300-1303221203220123-1323100333203231-2300311003330313-3220301303212000-0022220212232010-3022000123332132"></a>

#### `label_matcher.keys` property

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mum_action` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- mum_action

<a id="canonical-1232230311312032-2121331211020302-0213033312300333-0012232230033211-3200033003020130-2312132120000303-0211133100312120-1112302303011322"></a>

Type: `"single"`. Computed.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"default\",\"skip_processing\"]"
}
```

<a id="canonical-0012020010000323-0302330223323331-1302001030301220-2320212130320120-2233103331302022-3010010223203020-0200311220010001-3323322311202331"></a>

### Direct properties for `mum_action`

- [default](data-sources--service_policy_rule--reference--group-001.md#canonical-2003020203321223-0123233121100233-1132113230022032-2232011221333200-1312120330233100-3131101132302300-1010003020000323-0231332001211221): complete subsection reference.

- [skip_processing](data-sources--service_policy_rule--reference--group-001.md#canonical-3212231102301113-1020002133121323-3321230232033030-2130133321112213-0102232131102032-1131133320123332-1322221320320130-1203133200212132): complete subsection reference.

<a id="canonical-2003020203321223-0123233121100233-1132113230022032-2232011221333200-1312120330233100-3131101132302300-1010003020000323-0231332001211221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mum_action.default` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303)
- mum_action.default

<a id="canonical-3033220123200322-2031112332001123-2020131332123131-2132100310012032-0330122211030322-1112132322010133-0123033030113010-1001102023302202"></a>

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

<a id="canonical-3212231102301113-1020002133121323-3321230232033030-2130133321112213-0102232131102032-1131133320123332-1322221320320130-1203133200212132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mum_action.skip_processing` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [mum_action](data-sources--service_policy_rule--reference--group-001.md#canonical-0103133211113112-2110333301000123-0132303232311023-0033020113003322-3132123120231322-1200002211210031-0222200323230132-3323113100322303)
- mum_action.skip_processing

<a id="canonical-0100312101210231-0302220100200300-3002131133112312-0220212012323312-0100210000133313-0100021301232022-0211222230120232-0023333023103300"></a>

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

<a id="canonical-1202323031323101-3010322010211020-2201013123220300-0120101231322302-1330322232232230-0002133000323121-2021123103113212-2023022232033013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `path` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- path

<a id="canonical-0031213201320113-0231310121110121-0121200331113332-0010202322232200-0011332112301101-0322223221223202-2103011210010201-0213023020212223"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0201230013200313-3302213210113000-0011102120121221-2030301100211110-0311000231012032-1010100111222302-0132320020023203-1032032332220223"></a>

### Direct properties for `path`

<a id="canonical-2233111103220210-0022110311303002-1021121233222031-1032011321031322-3011332103003213-0010111122202122-0203220132302331-0132212011220102"></a>

#### `path.encoded_path_matcher` property

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

<a id="canonical-1021112133022313-1113200020120301-3112323111032321-1212111001101223-1223331200132322-2122031320002121-0031311213003312-3102303312330211"></a>

<a id="canonical-1200201332313331-1003321122013311-3130023301210213-0120030200132321-1103303201002300-1222302101102121-2230110313211331-2123022100123111"></a>

#### `path.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0122002222120103-2300311231001111-1311011013331230-2212113133301030-2200212210200301-3200133220230132-0032221313013111-2002222333212313"></a>

<a id="canonical-2103013312320223-0313323320023202-2111123121003032-1011333112033023-1233030321323302-0203311233113323-1121333213022201-3010203032003311"></a>

#### `path.invert_matcher` property

Type: `"bool"`. Computed.

Invert Path Matcher. Invert the match result.

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

<a id="canonical-3323220100003020-2333331113121203-3310122032320133-0022011132123112-0122122331313132-2301220000322102-1220220211212231-3031330301200212"></a>

<a id="canonical-2130023232020331-3020020030333203-1112220333003302-3133300310031233-2112111331211111-1332011002300331-2231123033322301-3312101002002303"></a>

#### `path.prefix_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2311313133011303-3332310211130130-0201203330202201-1300002230232211-0002001110301003-2323210233001032-3201002301122202-0121200000031030"></a>

<a id="canonical-2113321122310130-0311002131233103-3222012111310111-1223321320133001-1001322112032012-0123310231302200-2333221032312012-1330001310002201"></a>

#### `path.regex_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1222230001213311-0322221330330021-2013011220303322-0002110130221311-0221001011101022-2230231000332100-1302022031220301-2022211030222111"></a>

<a id="canonical-0033000313300200-3313033232222211-2332101132233122-2033322120200003-3222320122012202-0101323132103222-3332210033012320-1200313130102211"></a>

#### `path.suffix_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0123102202300202-3231310213302101-3230011013211232-0303202020230313-1310020313203131-0130000132320312-3201220320000200-3231200022221202"></a>

<a id="canonical-2020022111320030-0010311220203003-3332203233302312-0320311132112201-1101021230012100-1210333110212232-2333232120211303-3133313301133001"></a>

#### `path.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2101210020323103-3303202202131321-1133232111123110-1302223220031001-1031030010123130-1010013200001221-1302232103001023-1320212133031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `port_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- port_matcher

<a id="canonical-0131213012220100-3233021322333010-0232201212000122-1202123103003121-0133002011221000-3020313202313003-1303103303310102-0011031113301112"></a>

Type: `"single"`. Computed.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true. Server applies default when omitted.

Additional upstream details:

A port matcher specifies a list of port ranges as match criteria.

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

<a id="canonical-0110312220032211-1120012011113000-3310133033333021-0130202310211223-3100333232103102-0102120310111113-1100232303310132-2121103200032112"></a>

### Direct properties for `port_matcher`

<a id="canonical-3330332300030201-2203330233023021-0002032132001020-2012201330321203-1020331230110211-1222103030122130-2321312211030033-3131103112310113"></a>

#### `port_matcher.invert_matcher` property

Type: `"bool"`. Computed.

Invert Port Matcher. Invert the match result.

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

<a id="canonical-3101130102333312-2110110312330220-3020123011133333-0233110000213322-1222311330211123-0101303232310120-0122121013221113-3020002312112110"></a>

<a id="canonical-1322320233123301-3203203331010312-2321110312200212-1130222300003321-0231201331210221-2331211302121333-3123333033303312-3333022000022131"></a>

#### `port_matcher.ports` property

Type: `["list", "string"]`. Computed.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Additional upstream details:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-".

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `query_params` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- query_params

<a id="canonical-0110022223012231-1011022222022021-1231213210021012-2020023231312213-0120003203031103-1110212020200121-1320200132312313-1310013021311300"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0212331311311303-0321322311230213-1201232202012013-2213132130222221-2132032230320210-1030013032233203-2202113100332202-0312123133111322"></a>

### Direct properties for `query_params`

- [check_not_present](data-sources--service_policy_rule--reference--group-001.md#canonical-3220033332213100-2033312120023122-2231132332010002-2033031303113023-2132212313323222-1001212132301000-1210033320000312-2101012002123022): complete subsection reference.

- [check_present](data-sources--service_policy_rule--reference--group-001.md#canonical-2330013133300202-1312212010221301-2102030201103121-3303231100230202-0000303311230222-0002320123000121-2011023000302222-1332232020022211): complete subsection reference.

<a id="canonical-1103033200210122-3201213211221303-2230013333312012-1222031203122022-2022101301223122-2330130211320000-2021230222233110-2301132123331121"></a>

<a id="canonical-2100333011300110-0021200200230121-0312201333333101-3120032300320331-0000320113322332-1202010221221010-2011233323031310-2030312001223132"></a>

#### `query_params.invert_matcher` property

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](data-sources--service_policy_rule--reference--group-001.md#canonical-3133001330323313-2012212130033223-1311101230031121-0231002310113211-3111102301030122-2300120031010000-0310221131132102-0302320330303002): complete subsection reference.

<a id="canonical-0312121020113322-3203333312133313-1203103311320333-3121102221233032-3201130110310000-0000311232031333-3002323232211121-3023333131220320"></a>

<a id="canonical-3313110021012103-2233212133313000-1031111133032031-1013203202012122-0032103112011203-0123120032101032-1320322300230333-3111301111323010"></a>

#### `query_params.key` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3220033332213100-2033312120023122-2231132332010002-2033031303113023-2132212313323222-1001212132301000-1210033320000312-2101012002123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [query_params](data-sources--service_policy_rule--reference--group-001.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- query_params.check_not_present

<a id="canonical-3011101203123210-1300300210101020-3321130320311201-3310230230131332-2213302110000323-0122201221030113-2011323332311220-3000300112000112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-2330013133300202-1312212010221301-2102030201103121-3303231100230202-0000303311230222-0002320123000121-2011023000302222-1332232020022211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `query_params.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [query_params](data-sources--service_policy_rule--reference--group-001.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- query_params.check_present

<a id="canonical-1002312133133323-0223313003230001-3331200333031002-1332121010001221-2012123200022211-2211201002322331-3223232333131320-1303131103120003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-3133001330323313-2012212130033223-1311101230031121-0231002310113211-3111102301030122-2300120031010000-0310221131132102-0302320330303002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `query_params.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [query_params](data-sources--service_policy_rule--reference--group-001.md#canonical-0010003202113012-3010221330311030-1033010003322312-0003011202030230-0310211010032310-1130310201331200-0130002311302322-3313032310023203)
- query_params.item

<a id="canonical-0330101112120322-3333133122212133-2332122021103233-1322100002311122-3333313123111132-1133023013130202-0221213301303033-2102321000311213"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1001313021121202-2103033310122110-1203221120001320-3010303002331202-3311022221310203-3232323112100213-2102032003220330-0002031123113230"></a>

### Direct properties for `query_params.item`

<a id="canonical-0233120011232212-2331113232020021-3322333002132232-1133021001021230-3210112012102013-2212101320301333-1120222012313232-2232020233211120"></a>

#### `query_params.item.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-1131323033011300-2123101133231321-3221130011313103-3132232220303320-3331131121200131-0002313032300221-2033320002333012-1331313001110000"></a>

<a id="canonical-1220201202012003-0030211021332300-0333020322220003-3133303231022221-2003323331321122-2210331201121132-1122111301311202-3313032033113202"></a>

#### `query_params.item.regex_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0300320221022220-3000131220203013-1103213222132210-2120301022320023-2123003333033112-1000032112123101-0000111222300023-1001002002122313"></a>

<a id="canonical-0323011033120131-0020203101231133-1102110000000130-3302110133223103-3122020113120011-1002013210221301-3032000021023132-3210031310210203"></a>

#### `query_params.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- request_constraints

<a id="canonical-0121303130321000-2010323021321020-0130112232322010-2222221123133001-3032121201323213-1210123132003033-2233210030211101-2031120112210020"></a>

Type: `"single"`. Computed.

Configuration parameter for request constraints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_cookie_count_choice": "[\"max_cookie_count_exceeds\",\"max_cookie_count_none\"]",
  "x-ves-oneof-field-max_cookie_key_size_choice": "[\"max_cookie_key_size_exceeds\",\"max_cookie_key_size_none\"]",
  "x-ves-oneof-field-max_cookie_value_size_choice": "[\"max_cookie_value_size_exceeds\",\"max_cookie_value_size_none\"]",
  "x-ves-oneof-field-max_header_count_choice": "[\"max_header_count_exceeds\",\"max_header_count_none\"]",
  "x-ves-oneof-field-max_header_key_size_choice": "[\"max_header_key_size_exceeds\",\"max_header_key_size_none\"]",
  "x-ves-oneof-field-max_header_value_size_choice": "[\"max_header_value_size_exceeds\",\"max_header_value_size_none\"]",
  "x-ves-oneof-field-max_parameter_count_choice": "[\"max_parameter_count_exceeds\",\"max_parameter_count_none\"]",
  "x-ves-oneof-field-max_parameter_name_size_choice": "[\"max_parameter_name_size_exceeds\",\"max_parameter_name_size_none\"]",
  "x-ves-oneof-field-max_parameter_value_size_choice": "[\"max_parameter_value_size_exceeds\",\"max_parameter_value_size_none\"]",
  "x-ves-oneof-field-max_query_size_choice": "[\"max_query_size_exceeds\",\"max_query_size_none\"]",
  "x-ves-oneof-field-max_request_line_size_choice": "[\"max_request_line_size_exceeds\",\"max_request_line_size_none\"]",
  "x-ves-oneof-field-max_request_size_choice": "[\"max_request_size_exceeds\",\"max_request_size_none\"]",
  "x-ves-oneof-field-max_url_size_choice": "[\"max_url_size_exceeds\",\"max_url_size_none\"]"
}
```

<a id="canonical-2111302313300233-0210121300112313-3001003322312011-3313123233023032-3311101002121303-2213330312121133-1233300012221322-0332122030220320"></a>

### Direct properties for `request_constraints`

<a id="canonical-2133230320223220-1113123020232100-0233011100311311-2220223033232133-3111131323323231-3232322210332223-2332213312020201-3232132112311022"></a>

#### `request_constraints.max_cookie_count_exceeds` property

Type: `"number"`. Computed.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_count_none](data-sources--service_policy_rule--reference--group-001.md#canonical-3322130202102222-2002033301313123-0322013111331003-3201133202022132-2301323122123113-2220101110303020-1322001332112212-0012003221013332): complete subsection reference.

<a id="canonical-1303231131001120-1001120031110232-1103010020023232-3331030223330223-1331132022322023-0322322123113202-0030300232030133-3003313021120013"></a>

<a id="canonical-1033010023301200-2220000033330102-2110120023232100-0111102303011111-2013321212223122-2210211012012132-2210333212220320-0211031023302021"></a>

#### `request_constraints.max_cookie_key_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_key\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_key_size_none](data-sources--service_policy_rule--reference--group-001.md#canonical-0000033220333232-3211001012013030-0300033323331123-3022331331122131-1222202312230032-2133033211233301-2210203013012211-1330313023103330): complete subsection reference.

<a id="canonical-0103100211320112-3303303103023200-0221102113130302-2233122010100221-0132302032002032-0132232101102221-0211311320202031-3102102113001113"></a>

<a id="canonical-3101030313130130-1232301313213212-0112210010120313-1222103203003133-1003202112002321-0301303032203122-3210002332211212-2112323011211211"></a>

#### `request_constraints.max_cookie_value_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_value\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

- [max_cookie_value_size_none](data-sources--service_policy_rule--reference--group-001.md#canonical-3021022221302101-1312003030231221-0002211013021200-2000330000130120-3310201001130112-0333223310022132-3321210010212212-2011313233033202): complete subsection reference.

<a id="canonical-1011320330133023-1012201121103201-0220020110013000-1103113123322021-0303331033111012-0222221110223300-0121300123013110-3012233211103301"></a>

<a id="canonical-1010322201331311-1003003003122212-1112311310032302-3032322120112301-2010301030311312-0001310020311230-2330000333101101-0131222121021111"></a>

#### `request_constraints.max_header_count_exceeds` property

Type: `"number"`. Computed.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  }
}
```

- [max_header_count_none](data-sources--service_policy_rule--reference--group-001.md#canonical-1123213300220232-2101212233311012-1023031233101111-1133011231331301-3023033020130231-2302310003221210-2031231031003312-2122012002210303): complete subsection reference.

<a id="canonical-0021211210121302-2200300210020031-1011203031203000-1102310111013122-2210300122303131-1321322330121202-1101123303012123-1010010202232310"></a>

<a id="canonical-0210100010303033-1321332101200102-1200332311202021-1102203221202013-3331123101023213-0212123203112013-0300011130131002-3300010231323332"></a>

#### `request_constraints.max_header_key_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_header\_key\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_header_key_size_none](data-sources--service_policy_rule--reference--group-001.md#canonical-3313303021111131-1302003331000213-3231032112323223-0300030232230302-2211032133132212-1330210212312310-3111332021210102-1021012223230013): complete subsection reference.

<a id="canonical-3020010003200211-0221120102310101-0200113112132022-3220032210112210-1300312232100223-0312220230303033-3110111223313202-0233000323022122"></a>

<a id="canonical-2201113231102010-3031010113122010-0211312112120223-3100212310002001-1333300012122002-0013020012213102-0113012002100110-0023121301320120"></a>

#### `request_constraints.max_header_value_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_header\_value\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [max_header_value_size_none](data-sources--service_policy_rule--reference--group-001.md#canonical-0020232011313002-2133122103132211-3000120033100032-3220320232021001-1032321100113000-3311323230103210-1311020111232100-3033222222030303): complete subsection reference.

<a id="canonical-2230210011311131-3220103311102300-3112202003332123-2313200201113222-3011111012102202-2332010230231223-1312300231011212-2302110110222232"></a>

<a id="canonical-1221323123220202-1221300102311323-3333020120021320-3301313302300100-0033120320213101-2221022122233312-1211130130320003-0231330112023201"></a>

#### `request_constraints.max_parameter_count_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_count\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_count_none](data-sources--service_policy_rule--reference--group-001.md#canonical-3033220222300300-0312132120213313-3011220332122230-1113333033131312-0302330300123213-0300300322310022-3310003202121003-0233013221122221): complete subsection reference.

<a id="canonical-2011121232102213-3111131112321201-0122200021020202-0232203201103211-1211301033333121-1201003213032123-1201033231132220-3230213123111223"></a>

<a id="canonical-1331310130201313-3213333221113222-1033321223011331-0323022020312321-1132231121300210-1133312203102200-0201012023012001-0220330030213320"></a>

#### `request_constraints.max_parameter_name_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_name\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_name_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-0011311103233113-0020332013313232-3021002032111022-0203011021103123-1221333330033203-0121320221312221-2311221221201123-1000231123111131): complete subsection reference.

<a id="canonical-1202300203103111-1022032020110102-0000101131000020-1323231311011110-0201130121012113-3001321201123102-3332012303122003-1231120333210300"></a>

<a id="canonical-3100031222020103-1222231223010013-1100032133001213-2131103312032300-1200223001030101-1322021013122130-0023301112103231-1011022220122001"></a>

#### `request_constraints.max_parameter_value_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_value\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1073741824,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  }
}
```

- [max_parameter_value_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1312213213333021-1112322013103230-2013321221110002-2031101322120113-2102200330010300-3230203301312001-1100102031023300-3310310200121322): complete subsection reference.

<a id="canonical-2111112211231200-3121212100122210-2203032122130232-1331022232111100-2303220210233302-1320120203011133-3332222213313110-2120333320201121"></a>

<a id="canonical-1033231013032123-1012203223010301-1030122211221313-1213103222200232-0113002122013312-2030003222000023-2213023130322312-0111002301201120"></a>

#### `request_constraints.max_query_size_exceeds` property

Type: `"number"`. Computed.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [max_query_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1221101332103232-3222233020213003-2223321331211011-0211013303022331-0032003131003020-3223331120000131-2313110313120021-2030012203112000): complete subsection reference.

<a id="canonical-1102112313031123-1113223221133223-0021013202022231-0232200231230232-2121000203210332-0030122300330213-0313312231313112-1213332012132020"></a>

<a id="canonical-1100222203121231-2033132120221211-1113320212312203-2311001032032033-0213202002303311-0131122332103002-2311021011003012-3012002223121030"></a>

#### `request_constraints.max_request_line_size_exceeds` property

Type: `"number"`. Computed.

Exclusive with \[max\_request\_line\_size\_none\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_line_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-1003333222232313-0033202110033020-1231100320211000-0303120101200113-1111233102320332-1322231222013131-1102322123013121-0022031230311130): complete subsection reference.

<a id="canonical-0223002231203300-3122020102321311-1002001001122023-1133210130323002-1103311212233233-3113200020230020-2233021302033331-3200001133030313"></a>

<a id="canonical-0020012301321333-2123102030103023-3303221013321023-2203101201213111-3012300001032201-1003321133200221-1203130233223211-3003131021202333"></a>

#### `request_constraints.max_request_size_exceeds` property

Type: `"number"`. Computed.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-3103320223032223-3030132100322323-3232212132012203-3001100112123333-3131111113110120-3330312202013022-1332020030033333-0122320311212323): complete subsection reference.

<a id="canonical-2113220021310133-3312302133003221-2120222232020321-0201023231321213-2332101031203212-2330121103303231-2203230022033021-2222030023031102"></a>

<a id="canonical-2200322120132320-1221331310131331-1233331102231132-1323222323111311-1202233123023101-3120033221322022-3231201201221332-1000313130332323"></a>

#### `request_constraints.max_url_size_exceeds` property

Type: `"number"`. Computed.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  }
}
```

- [max_url_size_none](data-sources--service_policy_rule--reference--group-002.md#canonical-2210112122120113-1023110233212102-2203103322200020-0322210012030020-3332223301032223-2020120330212011-1010021311233031-1011330313322303): complete subsection reference.

<a id="canonical-3322130202102222-2002033301313123-0322013111331003-3201133202022132-2301323122123113-2220101110303020-1322001332112212-0012003221013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_cookie_count_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-001.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_cookie_count_none

<a id="canonical-1030113020022100-3110132033201201-3012322102131031-2003100123311002-1013100132030132-0013103030302130-0022123120223321-2203020213100101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max cookie count none.

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

<a id="canonical-0000033220333232-3211001012013030-0300033323331123-3022331331122131-1222202312230032-2133033211233301-2210203013012211-1330313023103330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_cookie_key_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-001.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_cookie_key_size_none

<a id="canonical-1033333332111222-3003302032233333-0211010132033231-3310003311301003-3212033101301303-1032122303212112-2211301202111332-1031120033122313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max cookie key size none.

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

<a id="canonical-3021022221302101-1312003030231221-0002211013021200-2000330000130120-3310201001130112-0333223310022132-3321210010212212-2011313233033202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_cookie_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-001.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_cookie_value_size_none

<a id="canonical-0210323111300301-1111012003030002-1322022032031323-3021122231020033-0221120021330323-3121231110103200-1201231031302223-2021232101312332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max cookie value size none.

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

<a id="canonical-1123213300220232-2101212233311012-1023031233101111-1133011231331301-3023033020130231-2302310003221210-2031231031003312-2122012002210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_header_count_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-001.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_header_count_none

<a id="canonical-1011200232203120-1120023222123130-3012332100333113-0203030232001031-2203132022312323-2223131332313013-2310122210301230-3211321030023300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max header count none.

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

<a id="canonical-3313303021111131-1302003331000213-3231032112323223-0300030232230302-2211032133132212-1330210212312310-3111332021210102-1021012223230013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_header_key_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-001.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_header_key_size_none

<a id="canonical-0101223120103230-1001211233301323-1112113320113113-1203303322022001-2011130311001020-3132003002210131-2331311332120212-1130111111331002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max header key size none.

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

<a id="canonical-0020232011313002-2133122103132211-3000120033100032-3220320232021001-1032321100113000-3311323230103210-1311020111232100-3033222222030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_header_value_size_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-001.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_header_value_size_none

<a id="canonical-3220222232320100-3012231320230101-0123211102311011-1301202121302323-0013311022300300-1220113302230332-0330210112021120-1301310000221100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max header value size none.

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

<a id="canonical-3033220222300300-0312132120213313-3011220332122230-1113333033131312-0302330300123213-0300300322310022-3310003202121003-0233013221122221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints.max_parameter_count_none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133)
- [Property reference](data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [request_constraints](data-sources--service_policy_rule--reference--group-001.md#canonical-0223111210233232-2103113322330233-2210311312311232-3333113212213223-1333100211111321-1330302213030021-3023212303011331-2203110020313010)
- request_constraints.max_parameter_count_none

<a id="canonical-1300321310102013-2312333223021233-0221302313001122-0300013002001321-3023232221132022-0200210310231210-1323002132331321-2022020122222212"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max parameter count none.

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
