---
page_title: "xcsh_service_policy_rule reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule reference."
---

# xcsh_service_policy_rule reference

<a id="canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- Property reference

<a id="canonical-2130321223013330-2323002003032032-0033301310100333-3123101310300001-0213220200101223-1322213020120023-1121321213002300-1211121223201000"></a>

### Direct properties for `xcsh_service_policy_rule`

<a id="canonical-2312210211301110-2113032010223321-0013113121303023-2120013012120113-3202000012033000-3310313012032330-2030023113223301-0331330321022233"></a>

#### `action` property

Type: `"string"`. Required.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALLOW","DENY","NEXT_POLICY"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210022213223201-1333112033210201-1133103230023132-1103032300033102-2202020302021103-3310321021102021-1130130030311332-1013221302221310"></a>

<a id="canonical-1313322012002130-0001100320023221-3212100210023233-0230330312122121-3122130031013130-0110332013321331-3023313102023013-3201311103012210"></a>

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

- [any_asn](resources--service_policy_rule--reference--group-001.md#canonical-0221232230332013-1321203101233021-0003232132312013-0013203002120133-2300131221032201-0023201233030211-0232122000323001-1021303221010210): complete subsection reference.

- [any_client](resources--service_policy_rule--reference--group-001.md#canonical-0123012100122112-0010012320202103-1121103211320203-3102030000330101-0132100103211313-1111033030002021-2011133021332311-2330013222310310): complete subsection reference.

- [any_ip](resources--service_policy_rule--reference--group-001.md#canonical-2323033333212103-1131233331312011-2132333310220111-1000201311003232-3201032233230003-0111300321031231-0000200020031132-3212101202111300): complete subsection reference.

- [api_group_matcher](resources--service_policy_rule--reference--group-001.md#canonical-2001200123112031-0101111212320231-0323322202023310-1033112332303103-2323023202232130-1220331203300031-3220022122303100-3013111133233130): complete subsection reference.

- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-1032023103021302-0223330032311303-2203320210330011-1003220302332130-3220031001022020-2032013022121012-0303233103102223-2031323210322110): complete subsection reference.

- [asn_list](resources--service_policy_rule--reference--group-001.md#canonical-0021332221230033-0003220101123220-2020212222231011-0302311332321131-3023201300001102-0300232311002101-1131213221031303-1201030020323010): complete subsection reference.

- [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1303131313113301-2002011311201202-0113230121212323-0202022331303111-2010020122331301-0002311130331320-1111101133131202-0133331233122210): complete subsection reference.

- [body_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1001101320132112-3130213301123030-0312023110021102-1100300111322002-1210010333200210-2020102132223133-2012322132222311-2003212133113201): complete subsection reference.

- [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-2103123012221321-1103321133011201-1313232020222001-2012201222102131-0300222221031123-2133203022221201-3200132333302332-1223010122321102): complete subsection reference.

<a id="canonical-3210022202131202-2200000220321321-3121211013333232-2312333100032232-1221220300133213-1102211020130221-2200131102321312-1131313313300320"></a>

<a id="canonical-1100223233211110-0130112200032113-1121300231132133-0032000303122213-0031100111120223-1113331223301023-3112100212010210-0120003330023031"></a>

#### `client_name` property

Type: `"string"`. Optional, Computed.

Exclusive with \[any\_client client\_name\_matcher client\_selector ip\_threat\_category\_list\] The
expected name of the client invoking the request API. The predicate evaluates to true if any of the
actual names is the same as the expected client name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
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

- [client_name_matcher](resources--service_policy_rule--reference--group-001.md#canonical-0332123330301021-1100223121211022-3121010103210131-3033010222100300-1123002321301311-1112213102202113-1311100312010010-0101133331030013): complete subsection reference.

- [client_selector](resources--service_policy_rule--reference--group-001.md#canonical-1300312101033300-3320221301112302-3333112131101103-3101332010233112-1220113220003101-2212302222200022-2133323313312132-2031301301211130): complete subsection reference.

- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-0323223303303012-1323121213223000-3303312201212002-2112132210332220-0300220220230320-3122301023330023-1133223001001010-2203310132010333): complete subsection reference.

<a id="canonical-3332000012320201-1213103132201001-0002320203033002-2321331312222230-3300110123121130-2212322321122010-0033210320213132-1222033102012300"></a>

<a id="canonical-0031302202312032-1032033123003003-2133101131322333-0011322120331222-3123322122032033-0013220322131221-2203320311013322-3302300012120112"></a>

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

<a id="canonical-1321020300323111-2012222021220001-0302223323032023-2103101103110102-3121031101233330-2111132113010133-1231001101312000-0013003112003333"></a>

<a id="canonical-1003100232232100-2003030303332123-3131312232203131-0301330122131323-1331233113303202-0201330111011122-2321131020320121-3221302121320300"></a>

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

- [domain_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1332222130011133-0323233220013211-0330111222302013-3023320213212310-2333013213211313-0332132110001121-3220010230201023-2032103202233210): complete subsection reference.

<a id="canonical-0110021313112000-2113311332311002-3132332332310331-2020000112032120-1112021233211201-0301031133320310-0220100330232102-0223020212100332"></a>

<a id="canonical-3231131300233012-2021102013023232-3220233121333102-1103211323120310-2001332221232311-3000332023020333-3223202033111312-0000112122302011"></a>

#### `expiration_timestamp` property

Type: `"string"`. Optional, Computed.

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

- [headers](resources--service_policy_rule--reference--group-001.md#canonical-2020321201213200-3003101032002102-0003211310323331-0102232130323303-3322002211110311-3103210330303202-1130102333320230-1030212213102000): complete subsection reference.

- [http_method](resources--service_policy_rule--reference--group-001.md#canonical-3311223333320012-1301003001231331-1200011030303311-0111131102332203-0033001320311120-3102002212303313-3213012110031222-2313312311022232): complete subsection reference.

<a id="canonical-3201322011003103-0000220312130322-1102230323121100-2221122102123102-0021020033202203-1231102311330131-3220123000213002-0233313032031032"></a>

<a id="canonical-0030233003332313-3232023302320111-0120312333121301-2020302331212212-0132201010200311-0222013020102330-3200011210020301-1113020033211120"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1101030111001132-2130113101200012-1213033232233301-0331310212100221-2201110133311321-1321032003101012-0133122013101313-0213333230233312): complete subsection reference.

- [ip_prefix_list](resources--service_policy_rule--reference--group-001.md#canonical-1001031102013032-3012221302301103-1333012032022110-3313212031020033-3133033332330121-2311012133200203-3200331331023200-2110011322231321): complete subsection reference.

- [ip_threat_category_list](resources--service_policy_rule--reference--group-001.md#canonical-0012302311320221-1202233202103232-2011312303031222-3122303101021120-0303100330112201-0323021021121103-1213030202102213-1311133121210233): complete subsection reference.

- [ja4_tls_fingerprint](resources--service_policy_rule--reference--group-001.md#canonical-2232311312222031-0121212211131221-3223322103010300-2012020332313011-1130321013122111-3223223033331333-0020100121032303-3310222113002133): complete subsection reference.

- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113): complete subsection reference.

- [label_matcher](resources--service_policy_rule--reference--group-001.md#canonical-2020111211100111-2033231320322121-0121111211102333-3113322310323132-1300301223000233-0231032111303233-0333010313331210-0100231321203000): complete subsection reference.

<a id="canonical-2233000230231032-2210220331101212-3133132033220000-3110103111033000-0132232031010103-2310120000031231-1200232213120013-1002023222023222"></a>

<a id="canonical-3022020300311022-3102003121332313-0123112203023202-0302031011020312-1121233313300033-2022032021032222-1133232120013220-2103303322013002"></a>

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

<a id="canonical-3222311030321211-0233112310002323-2120331200212302-3330132211333103-2322201210032121-0133202002212101-0311113311300021-0130212112332011"></a>

<a id="canonical-3121020001110310-3220220312222111-2311121310121303-3231110210303220-0331103223303112-0202331203012130-2012312210131300-0030231222233002"></a>

#### `log_rule_evaluation` property

Type: `"bool"`. Optional, Computed.

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

- [mum_action](resources--service_policy_rule--reference--group-001.md#canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131): complete subsection reference.

<a id="canonical-3311330211110300-2322212221121203-2332323033302112-3100002022031232-0310333321013002-3102012212231212-0013201001131200-3121232213321221"></a>

<a id="canonical-0101131231103302-2200110302023222-3322121000322021-3231310001321102-3030003321033110-0202011001311033-2313311300313122-0330210300211330"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Service Policy Rule. Must be unique within the namespace.

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

<a id="canonical-1230330323221222-2131313222202020-1200012132022022-2303120013001111-3011322001112123-2132323323301301-0230012320210222-1022120231310003"></a>

<a id="canonical-1011013232011210-0301032013211130-1303132320133011-2321331130220223-2210032310011323-2032130123021013-2331210100332213-3102331211011213"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Service Policy Rule is created.

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

- [path](resources--service_policy_rule--reference--group-001.md#canonical-3010110322330023-1103101211122201-1133323113223332-1313232321223110-0101110303212213-1030200111110001-0302033113300012-0303132213131330): complete subsection reference.

- [port_matcher](resources--service_policy_rule--reference--group-001.md#canonical-2203222212003231-2202300100330200-1310211032320202-1313000110311333-0331313210023112-2112333302313030-1110023000021210-2033311231020133): complete subsection reference.

- [query_params](resources--service_policy_rule--reference--group-001.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101): complete subsection reference.

- [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012): complete subsection reference.

- [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0321323100221312-3112222323022232-0010121020212310-2330002103302030-1011120203032301-0030033112033132-2323303301011102-0131113333313221): complete subsection reference.

- [timeouts](resources--service_policy_rule--reference--group-002.md#canonical-1021310131303231-3111121233212022-1000112321100121-1023301303021202-3313011013101213-2203330203232131-3231103133122222-2233231203321311): complete subsection reference.

- [tls_fingerprint_matcher](resources--service_policy_rule--reference--group-002.md#canonical-2200000321123222-1121231032320123-3112320132222010-3101030202331210-2103131131000030-1333130312003211-1330202233211000-0302103210011223): complete subsection reference.

- [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-1201320322210230-3302110300023123-3200323003323032-1123021203110130-2133221133111213-0323113131313022-2121100020001222-3032001111321312): complete subsection reference.

<a id="canonical-3123103103300122-1023313133130103-1031332102031102-2312323201010313-0301331033213302-2313220300020120-2333120220012220-1303233110122123"></a>

### All schema paths for `xcsh_service_policy_rule`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `action` | [action](resources--service_policy_rule--reference--group-001.md#canonical-2312210211301110-2113032010223321-0013113121303023-2120013012120113-3202000012033000-3310313012032330-2030023113223301-0331330321022233) |
| `annotations` | [annotations](resources--service_policy_rule--reference--group-001.md#canonical-3210022213223201-1333112033210201-1133103230023132-1103032300033102-2202020302021103-3310321021102021-1130130030311332-1013221302221310) |
| `any_asn` | [any_asn](resources--service_policy_rule--reference--group-001.md#canonical-2021232023103020-1311221231322320-2222121231030210-3332011313321110-0002000113221131-1012300202111122-2321323031130323-2103001211202222) |
| `any_client` | [any_client](resources--service_policy_rule--reference--group-001.md#canonical-0332123200212031-2133003002311213-2132022222131132-3321011301221221-2201101101131212-3221312231000232-1021221221230320-3331123312230010) |
| `any_ip` | [any_ip](resources--service_policy_rule--reference--group-001.md#canonical-2332201302210331-2302231200033202-0311033331123333-1123320002301122-0102102012203001-3233102001102210-2233331031332022-3323231200100331) |
| `api_group_matcher` | [api_group_matcher](resources--service_policy_rule--reference--group-001.md#canonical-0020300113321302-2000001122310033-0200310032213011-0210023311311321-1010130332032110-2013022222333332-3233210310122300-3311101013003322) |
| `api_group_matcher.invert_matcher` | [api_group_matcher.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-0312320211233133-0313020131133210-0220331203302000-2201130201030200-2030103100110132-1130323221131331-3120331123302210-1212323231131203) |
| `api_group_matcher.match` | [api_group_matcher.match](resources--service_policy_rule--reference--group-001.md#canonical-0123011011030001-1131331003222003-3330122131002100-0010110220022302-3332300230332033-0112233210100221-1331223321213311-1223323233200120) |
| `arg_matchers` | [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-1222000023023301-0330002201002230-1222123321113320-1030320320222321-3022201123210032-3323011102313303-3031222002101112-3330212311132101) |
| `arg_matchers.check_not_present` | [arg_matchers.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-0201310201202121-0203012332231103-1132022032321312-3102123333330120-0013112011110203-0321332323223101-3311200220220100-0220113132131010) |
| `arg_matchers.check_present` | [arg_matchers.check_present](resources--service_policy_rule--reference--group-001.md#canonical-0101120133232000-3101033203320303-3222100012102001-1313133203302312-1332200321303302-3312302310323301-2020322312313233-1111312312221320) |
| `arg_matchers.invert_matcher` | [arg_matchers.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-2302221102322002-3020100222130332-0102033012320033-1311020322203203-1200321221131100-0003120303030202-3300233132121111-1210100112010003) |
| `arg_matchers.item` | [arg_matchers.item](resources--service_policy_rule--reference--group-001.md#canonical-2021321123232221-1030303120131220-2001200100013111-2233210321011133-3022023002213201-1023103103011321-3121131101032212-2301333312311002) |
| `arg_matchers.item.exact_values` | [arg_matchers.item.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-3321312331100031-1031120000202223-1101122322323321-3201111221122212-3230312231013011-3320123320133303-3301213101112030-2330210230322312) |
| `arg_matchers.item.regex_values` | [arg_matchers.item.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-3332131232013320-1220011323023332-2300030023000332-3000101321020231-1103301313011322-3010321322301130-1300022201120003-0111303121310223) |
| `arg_matchers.item.transformers` | [arg_matchers.item.transformers](resources--service_policy_rule--reference--group-001.md#canonical-3221220311013111-1102122110013000-3013023101202300-1311020221301333-1011211222120321-3000121300001200-2220221221131310-2210111223112323) |
| `arg_matchers.name` | [arg_matchers.name](resources--service_policy_rule--reference--group-001.md#canonical-2231202131023212-1003033232003033-3232033322210012-3230201322322023-1033001032021201-1201101213211200-0212202210101233-1003311301311202) |
| `asn_list` | [asn_list](resources--service_policy_rule--reference--group-001.md#canonical-2102013201202122-0110010331120030-0132033210110300-0310100302323203-0032200323333001-0333031101321330-3032130003323011-2032123322202222) |
| `asn_list.as_numbers` | [asn_list.as_numbers](resources--service_policy_rule--reference--group-001.md#canonical-1303000313130111-2210003313010021-3312021301100030-0003323212322133-0313112330211110-1111311032222121-0213020313211030-2121033003303120) |
| `asn_matcher` | [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1221102131001321-0211132300133300-3032122013030021-3330012330203320-1111302112032022-2010003103010103-3212002003231302-2032112213032110) |
| `asn_matcher.asn_sets` | [asn_matcher.asn_sets](resources--service_policy_rule--reference--group-001.md#canonical-0302030310311111-0213312232300232-1033212032111210-1333110021211213-1212033321100020-1012010103311002-3302331112220202-3130303222122223) |
| `asn_matcher.asn_sets.kind` | [asn_matcher.asn_sets.kind](resources--service_policy_rule--reference--group-001.md#canonical-2112123220221221-1230301211021312-0311111002031002-1232032123221323-2111120233020330-0132033031320332-1120320033123213-0001000130223222) |
| `asn_matcher.asn_sets.name` | [asn_matcher.asn_sets.name](resources--service_policy_rule--reference--group-001.md#canonical-2201313200031103-2312310101001013-2302231211113003-3330303113203011-3303010321013112-1232233303300122-3203103312201000-1032322122322131) |
| `asn_matcher.asn_sets.namespace` | [asn_matcher.asn_sets.namespace](resources--service_policy_rule--reference--group-001.md#canonical-1210131120210100-3321132310133131-1231113330332113-0223003303212133-3300110122012230-3102000232022032-1330013233113110-2222120332103222) |
| `asn_matcher.asn_sets.tenant` | [asn_matcher.asn_sets.tenant](resources--service_policy_rule--reference--group-001.md#canonical-2312222230230110-0013130002131321-0310302111111033-0312222010203113-2331223301132222-3112031313233003-0200132300031112-0132102311011120) |
| `asn_matcher.asn_sets.uid` | [asn_matcher.asn_sets.uid](resources--service_policy_rule--reference--group-001.md#canonical-2100130123023201-1001023220222322-0101202223322020-1321011301021303-2322020232213032-0233010003310232-2031223032122102-3200032311132032) |
| `body_matcher` | [body_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1122022323301132-1233100323013212-3122031331313102-0121301331200011-2133221311020330-2012233300032203-3123221112100101-3023322130333203) |
| `body_matcher.exact_values` | [body_matcher.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-2312120231210202-2121001330101303-3310130211123110-3021111113200323-0012211003120233-0032112333022311-2121002021211222-2000223302113123) |
| `body_matcher.regex_values` | [body_matcher.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-3233001100122201-3023001313201223-1103223310011230-0110213323203223-2200233111320111-3200221311211301-3002133302333303-2002303030131231) |
| `body_matcher.transformers` | [body_matcher.transformers](resources--service_policy_rule--reference--group-001.md#canonical-2113301112001032-1312310022101200-0112311313122310-1123000031313302-3231122331021301-3223020101131212-2012210323100202-0021332003021203) |
| `bot_action` | [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-0211333322223333-0012021033002113-3113330301122031-3133300333333103-3302312120011311-2133220231211123-2211120222101323-1233311200003102) |
| `bot_action.bot_skip_processing` | [bot_action.bot_skip_processing](resources--service_policy_rule--reference--group-001.md#canonical-2332113222010211-1120003301003132-1120011131222103-1221131231203032-1222203010220220-1001200310221033-2221333003103221-0320121310000013) |
| `bot_action.none` | [bot_action.none](resources--service_policy_rule--reference--group-001.md#canonical-2002301121233310-3003301303200212-2032210101120212-2033333010111110-3031310313021210-3220000020312130-3333101003311210-3112103312101102) |
| `client_name` | [client_name](resources--service_policy_rule--reference--group-001.md#canonical-3210022202131202-2200000220321321-3121211013333232-2312333100032232-1221220300133213-1102211020130221-2200131102321312-1131313313300320) |
| `client_name_matcher` | [client_name_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1112333331011102-2032331320312222-0130031223211101-3311210300301112-2323213120110120-3212330201033321-1113032330323313-1211222302020221) |
| `client_name_matcher.exact_values` | [client_name_matcher.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-1110002302110013-0211112020210133-2330013320302103-2333201210010220-0330030032311210-1222301210133232-2220203311112332-1330310222230330) |
| `client_name_matcher.regex_values` | [client_name_matcher.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-3010122320302101-1303110131221000-2032012203011012-1210233202233331-0223332303223200-1200212221321311-2102321232103232-1333000300210202) |
| `client_selector` | [client_selector](resources--service_policy_rule--reference--group-001.md#canonical-1100200013101011-2030121221031312-1013010130232011-0331330201133123-3232201232320220-1001101313033120-3032323121002232-0331120132130013) |
| `client_selector.expressions` | [client_selector.expressions](resources--service_policy_rule--reference--group-001.md#canonical-0201022300100121-2133322313133122-3312000321132200-3320031311133131-1003232121111211-2030122012103133-0131300111022233-1322000120220320) |
| `cookie_matchers` | [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-1303013313221111-3330333231120001-1133203100133112-0131212300130023-1202020011131101-2312130111203201-2320322122113212-0010302221203100) |
| `cookie_matchers.check_not_present` | [cookie_matchers.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-2020323310202031-1110323301130033-3101112230021230-2223322303011113-3300220123020101-1211300212203333-3310210312133013-0303221022121310) |
| `cookie_matchers.check_present` | [cookie_matchers.check_present](resources--service_policy_rule--reference--group-001.md#canonical-3100000301121103-3131133003130030-1030211212330230-3201021020321220-2111232012233001-0333110331302021-2131010333223332-0332211231022032) |
| `cookie_matchers.invert_matcher` | [cookie_matchers.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-0012303103120200-1111131102121321-1003222131132232-3330203031011113-3020332003102200-0031010023013231-0313230121232101-3313133300302323) |
| `cookie_matchers.item` | [cookie_matchers.item](resources--service_policy_rule--reference--group-001.md#canonical-3132212312222031-2330110023002212-1321033222303002-1211022200321232-3123211210312133-0121321232200202-3132232121213203-3320322331321303) |
| `cookie_matchers.item.exact_values` | [cookie_matchers.item.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-0110320321201122-0303000122100323-2132202201010121-0101111211310232-3310100331223202-2312332120002013-1111122210211312-3103211133022213) |
| `cookie_matchers.item.regex_values` | [cookie_matchers.item.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-0322003132210311-2132333213120032-3203313011303121-3323000123232023-1032002330103120-2122321320222130-2233210302133311-1002222331120011) |
| `cookie_matchers.item.transformers` | [cookie_matchers.item.transformers](resources--service_policy_rule--reference--group-001.md#canonical-3323213031202023-3321211300311001-1331013100111111-2120213001011231-0101103202203310-1320332133212220-0302200323113003-0103121213310232) |
| `cookie_matchers.name` | [cookie_matchers.name](resources--service_policy_rule--reference--group-001.md#canonical-0201213331312022-2313222020132223-0011303111302113-3003312311232033-3210122332001210-1120302122101103-1111330121121230-2302203100121223) |
| `description` | [description](resources--service_policy_rule--reference--group-001.md#canonical-3332000012320201-1213103132201001-0002320203033002-2321331312222230-3300110123121130-2212322321122010-0033210320213132-1222033102012300) |
| `disable` | [disable](resources--service_policy_rule--reference--group-001.md#canonical-1321020300323111-2012222021220001-0302223323032023-2103101103110102-3121031101233330-2111132113010133-1231001101312000-0013003112003333) |
| `domain_matcher` | [domain_matcher](resources--service_policy_rule--reference--group-001.md#canonical-3031311003030133-0022010310120120-2212123301010100-2110310032302121-1330130231313112-3220222112111310-0102133201222320-0311111131321321) |
| `domain_matcher.exact_values` | [domain_matcher.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-3011113220102231-1033210201321331-3103012112123212-2323233010212113-1211203210101022-0232303022122302-3000122322112033-1330201210220123) |
| `domain_matcher.regex_values` | [domain_matcher.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-2230211000033202-2300001210321311-2103313201113101-1120330200020033-1232310313103310-1333101222012020-1221021322222320-2322333012023111) |
| `expiration_timestamp` | [expiration_timestamp](resources--service_policy_rule--reference--group-001.md#canonical-0110021313112000-2113311332311002-3132332332310331-2020000112032120-1112021233211201-0301031133320310-0220100330232102-0223020212100332) |
| `headers` | [headers](resources--service_policy_rule--reference--group-001.md#canonical-3100231310321321-1112023301102132-2321103011213002-0221111301202230-0320230313103212-1012230100011130-1010213320020130-1331120322021311) |
| `headers.check_not_present` | [headers.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-0123223012103301-1212221122310102-1022100222001222-1311012231211230-0222020323023210-3230331110322210-1213220213021210-0202020212221221) |
| `headers.check_present` | [headers.check_present](resources--service_policy_rule--reference--group-001.md#canonical-3331101312120303-2320133322302031-3112310130213130-1120030312131023-0320011001030000-2110020311223122-2300302110330310-2212030031011321) |
| `headers.invert_matcher` | [headers.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-0201322233330101-0213311201330121-2210302013301321-2030230222130101-0311013001212102-2023100013013322-0021311022223030-0330210322202203) |
| `headers.item` | [headers.item](resources--service_policy_rule--reference--group-001.md#canonical-3002312200232200-3330330332211333-2030133223131002-2002223222032320-3002301212020313-2330203231323302-3322030031223333-0212011110201120) |
| `headers.item.exact_values` | [headers.item.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-0000011130223002-0000022233333002-0033301113031120-0323202221111320-3022132213130222-0012233021123033-0312201101033202-3201321331320223) |
| `headers.item.regex_values` | [headers.item.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-1123022002311322-0130001121222110-0123120333221302-1321201323322120-2023312203020203-0320311020023201-1022332030222322-3312120033211330) |
| `headers.item.transformers` | [headers.item.transformers](resources--service_policy_rule--reference--group-001.md#canonical-2210322210303121-2033320113321003-0231332213331211-0103013213121331-0310331310122302-0311010323200012-3011131303010220-0012033202110303) |
| `headers.name` | [headers.name](resources--service_policy_rule--reference--group-001.md#canonical-3330222210300200-2312111100120301-0120223332223033-0113233103323332-0113131133001103-3020002021301333-3211033321011300-1130010121211322) |
| `http_method` | [http_method](resources--service_policy_rule--reference--group-001.md#canonical-1312132130110201-0020131303101121-1001103103301103-2311301310022310-3310002321111100-0310131213202012-1302232332321213-3113321302232333) |
| `http_method.invert_matcher` | [http_method.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-0002002131101130-3022011120230131-2103001023212210-0022231310131000-1033200203311332-1003123210333030-2103222112333120-0210011331332302) |
| `http_method.methods` | [http_method.methods](resources--service_policy_rule--reference--group-001.md#canonical-1103132021332032-0302033230313120-1112302322210321-1011303211133230-2221102331130331-2012212123031333-3110301303031133-0120001301123322) |
| `id` | [ID](resources--service_policy_rule--reference--group-001.md#canonical-3201322011003103-0000220312130322-1102230323121100-2221122102123102-0021020033202203-1231102311330131-3220123000213002-0233313032031032) |
| `ip_matcher` | [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1202321232220311-2202112122103101-3030011112103123-0020033132221020-1311010010031120-1101202223111322-0322311103130230-0130112001011022) |
| `ip_matcher.invert_matcher` | [ip_matcher.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-3303301031033301-0101313213031213-3312003230101212-0121010123113210-0200321030120311-1133233332021111-3320223331002000-3201331101203130) |
| `ip_matcher.prefix_sets` | [ip_matcher.prefix_sets](resources--service_policy_rule--reference--group-001.md#canonical-3233013301230123-1131113210312133-0113101120022233-0201201323332003-1112221022101002-2302121130302222-2000230100232022-3001020313111201) |
| `ip_matcher.prefix_sets.kind` | [ip_matcher.prefix_sets.kind](resources--service_policy_rule--reference--group-001.md#canonical-3021012023330123-1212213123033200-0211103110122102-3103202302012231-0303101321033032-2100322301023201-2023202200311320-2303011333331021) |
| `ip_matcher.prefix_sets.name` | [ip_matcher.prefix_sets.name](resources--service_policy_rule--reference--group-001.md#canonical-3033013011102301-3220112000022223-3200032230222210-3021013231020330-2332103303333211-1020330210200313-3330310110320100-1132123311110303) |
| `ip_matcher.prefix_sets.namespace` | [ip_matcher.prefix_sets.namespace](resources--service_policy_rule--reference--group-001.md#canonical-0310000233130111-1320322323112300-2231312223231321-2221011220212232-1010031220213333-0210130332303123-3032111312110221-0201320321021102) |
| `ip_matcher.prefix_sets.tenant` | [ip_matcher.prefix_sets.tenant](resources--service_policy_rule--reference--group-001.md#canonical-1212123300023101-2332121200133103-2021203122012313-2000001322123321-1113230031211332-1101230232112231-2222332033110330-3223033213121023) |
| `ip_matcher.prefix_sets.uid` | [ip_matcher.prefix_sets.uid](resources--service_policy_rule--reference--group-001.md#canonical-2110212113221010-1213210023120200-2100101120133011-3333011230110203-0031203111331023-3201202031302102-1202222321211233-0213200031223202) |
| `ip_prefix_list` | [ip_prefix_list](resources--service_policy_rule--reference--group-001.md#canonical-1133010200322311-2203212333112121-2022320311033220-1222320331002313-2302333111223002-3112000131122011-0102211300332020-0320031012310030) |
| `ip_prefix_list.invert_match` | [ip_prefix_list.invert_match](resources--service_policy_rule--reference--group-001.md#canonical-1130333130002213-2122330233122100-2332030230323321-2200003210200111-0201000310211110-1133303300233302-1000312222030221-0232320121013020) |
| `ip_prefix_list.ip_prefixes` | [ip_prefix_list.ip_prefixes](resources--service_policy_rule--reference--group-001.md#canonical-0010320331303221-2231012110002312-1012001322030131-1130311133000123-1030121102000320-0001313132200321-2122003330311100-1232020033033011) |
| `ip_threat_category_list` | [ip_threat_category_list](resources--service_policy_rule--reference--group-001.md#canonical-0302200111131213-0020202112102320-2202002130110021-3103300012200032-1332110312113231-3320210001123321-3333203121031111-1100132203212331) |
| `ip_threat_category_list.ip_threat_categories` | [ip_threat_category_list.ip_threat_categories](resources--service_policy_rule--reference--group-001.md#canonical-2223300023201310-3013310111311301-1323233022033033-1000120312211111-1122120031101301-0000023210101132-2320130221002322-2311323332203202) |
| `ja4_tls_fingerprint` | [ja4_tls_fingerprint](resources--service_policy_rule--reference--group-001.md#canonical-2320112003011023-1000221201331002-0111220231302002-1211120101020002-1331131133033203-2220033003013011-2031203330133310-0311101013313003) |
| `ja4_tls_fingerprint.exact_values` | [ja4_tls_fingerprint.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-2320321310112103-2210311303010102-1310000112233211-3201211311332100-3011302022230030-1202131312310223-0213323210013133-3201131331010320) |
| `jwt_claims` | [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-0232111131010322-1023012332212110-0323323020302020-0011233321330000-1212323132301302-3033220331130113-0000321120220200-1122231031303301) |
| `jwt_claims.check_not_present` | [jwt_claims.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-3211003220132220-2111200033001312-0211232230111330-3000011132030100-2110300023103031-0203130031210320-3200222002201133-3221130312102132) |
| `jwt_claims.check_present` | [jwt_claims.check_present](resources--service_policy_rule--reference--group-001.md#canonical-3010233313020311-0333203110111132-3110113212022133-1020022110221232-3232011133320320-0003022122322222-0123212200202213-1032013222333231) |
| `jwt_claims.invert_matcher` | [jwt_claims.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1101030231301201-0101102022222020-2230211101302233-3011100312001211-1311032321002301-0230303211003202-0032101032133013-3312221012033300) |
| `jwt_claims.item` | [jwt_claims.item](resources--service_policy_rule--reference--group-001.md#canonical-3122203122311031-0320223010303311-2303023021112000-1200201201000221-2321323301132011-0002230313021110-1312102132223122-3221301301011011) |
| `jwt_claims.item.exact_values` | [jwt_claims.item.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-1131203020211002-1132013133221130-2330323330200033-3320120000301132-2021301102313331-2333232232133302-2001322133101011-1003030101331331) |
| `jwt_claims.item.regex_values` | [jwt_claims.item.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-0333211231120231-2221101333230013-2332220003030120-0320011313320131-3322220312003013-0201102031323122-3111010002110233-0121101120223200) |
| `jwt_claims.item.transformers` | [jwt_claims.item.transformers](resources--service_policy_rule--reference--group-001.md#canonical-2201100002003013-0013312020120131-3123210331020232-3233123332033123-3321302130020302-0013031102222131-3132331313323100-2203111311220022) |
| `jwt_claims.name` | [jwt_claims.name](resources--service_policy_rule--reference--group-001.md#canonical-1231130012032130-2301020010201101-2200222213010320-2333130103320103-1223033230213030-2311233211103102-3302030332233213-0331113201210101) |
| `label_matcher` | [label_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1312003200110013-1123113121200330-0303232330132123-3130321201330321-2333032023303210-3110021030303000-3113023202210211-2302312332233300) |
| `label_matcher.keys` | [label_matcher.keys](resources--service_policy_rule--reference--group-001.md#canonical-2303011232220331-2131220312311033-1100103323302001-1132322033300202-0223223333113312-1103211202000323-0011012300003323-0312120022330012) |
| `labels` | [labels](resources--service_policy_rule--reference--group-001.md#canonical-2233000230231032-2210220331101212-3133132033220000-3110103111033000-0132232031010103-2310120000031231-1200232213120013-1002023222023222) |
| `log_rule_evaluation` | [log_rule_evaluation](resources--service_policy_rule--reference--group-001.md#canonical-3222311030321211-0233112310002323-2120331200212302-3330132211333103-2322201210032121-0133202002212101-0311113311300021-0130212112332011) |
| `mum_action` | [mum_action](resources--service_policy_rule--reference--group-001.md#canonical-2221121310120320-0122232013020111-3111313002230103-0012100300022122-2233030220321313-1231102130223133-0323221002032002-3032000311011100) |
| `mum_action.default` | [mum_action.default](resources--service_policy_rule--reference--group-001.md#canonical-2102113111213301-3030032320301122-0101232002213212-0322211221310200-1101210110003000-2131012033221002-1022221121113013-1031100210203302) |
| `mum_action.skip_processing` | [mum_action.skip_processing](resources--service_policy_rule--reference--group-001.md#canonical-0312310211231330-1203021111031102-3103233002012120-1210130133030212-1322020300113131-3311120303211301-3330102301313021-0213020111213311) |
| `name` | [name](resources--service_policy_rule--reference--group-001.md#canonical-3311330211110300-2322212221121203-2332323033302112-3100002022031232-0310333321013002-3102012212231212-0013201001131200-3121232213321221) |
| `namespace` | [namespace](resources--service_policy_rule--reference--group-001.md#canonical-1230330323221222-2131313222202020-1200012132022022-2303120013001111-3011322001112123-2132323323301301-0230012320210222-1022120231310003) |
| `path` | [path](resources--service_policy_rule--reference--group-001.md#canonical-1222332103311013-2322202031102010-2310333211132301-2133303132123131-3101032220033221-2332323101120012-0023222202301233-0100221110233300) |
| `path.encoded_path_matcher` | [path.encoded_path_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1203321001120102-3230123100011102-3310122110021233-1010032220322303-1321202311111101-2101232203313223-3031113123221003-0123331113312113) |
| `path.exact_values` | [path.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-3300211023023011-0100330002020032-0221232220001003-2023130313203113-1110113023021102-0203110332303122-0301132030231013-0232211211002022) |
| `path.invert_matcher` | [path.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-0122202103323232-3332313000012033-0031003032201200-1303213100000221-1122210032020210-3221322012000233-0230203021001300-2003012113330110) |
| `path.prefix_values` | [path.prefix_values](resources--service_policy_rule--reference--group-001.md#canonical-0033131231323313-2133303031222313-3330221011331312-0132022203002230-3233300231311030-1232003330022330-1003213132011232-2033012301112130) |
| `path.regex_values` | [path.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-3231203123331011-1010021300103133-3100010310231333-2112002200321102-2213301011312013-2121110313020232-3301210032010221-2203323010311321) |
| `path.suffix_values` | [path.suffix_values](resources--service_policy_rule--reference--group-001.md#canonical-0012332022230031-0201113331012212-0221033003102023-2030201101331300-2322202311021232-1102202303102103-0022332222100002-3123213102000001) |
| `path.transformers` | [path.transformers](resources--service_policy_rule--reference--group-001.md#canonical-2332221310210123-1303210232003020-1011212330201122-0320023322033321-0332301020311103-1223013320302002-0213010132303030-1212211131110230) |
| `port_matcher` | [port_matcher](resources--service_policy_rule--reference--group-001.md#canonical-2322112120320222-3113003012110013-3110231221003322-0221120100321303-2132110222311133-2030211130312102-0231232102133322-2023023231200211) |
| `port_matcher.invert_matcher` | [port_matcher.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-0120221311003100-1000230031222100-1322323000312032-1031130332211231-1210111013010030-3210012031021302-3103210102000132-3032321110230021) |
| `port_matcher.ports` | [port_matcher.ports](resources--service_policy_rule--reference--group-001.md#canonical-2223320321103203-0023122212323001-3002223020320021-3213202020100230-1231130222132002-0132303232200030-3131023001013101-2010121223331303) |
| `query_params` | [query_params](resources--service_policy_rule--reference--group-001.md#canonical-1132121303300022-3222130331310311-3012221303320222-0300303322210021-2012131020223121-0203302231323312-3113331121030333-0011112032010322) |
| `query_params.check_not_present` | [query_params.check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-3202330001122331-1130330220001301-2201120212330221-2000333031220132-0331013333332133-0030011200120031-0232302232121212-1200221333202002) |
| `query_params.check_present` | [query_params.check_present](resources--service_policy_rule--reference--group-001.md#canonical-0203031220002032-2321320122323321-0330112110213330-0210112333032311-2300003111301101-1133223023031032-1103032000122312-3123133310113121) |
| `query_params.invert_matcher` | [query_params.invert_matcher](resources--service_policy_rule--reference--group-001.md#canonical-2033101023010311-0310112222301102-2100232312001011-1100313301133203-1131022123000031-0102213230111001-1303220330120300-0212300011311202) |
| `query_params.item` | [query_params.item](resources--service_policy_rule--reference--group-001.md#canonical-2233103232121213-2131110023313322-0313133320302120-1121000310110033-2111303130113330-0132132000302301-2101021032003213-3132031301203312) |
| `query_params.item.exact_values` | [query_params.item.exact_values](resources--service_policy_rule--reference--group-001.md#canonical-3323111210031331-1133123113121110-3131322121022313-1122302110012211-0011213013231301-2301000132230213-3332312100301303-2113311203232112) |
| `query_params.item.regex_values` | [query_params.item.regex_values](resources--service_policy_rule--reference--group-001.md#canonical-1003310032001103-3332232022122202-1031211322230031-3033231223011030-3201321033301331-3002013130130221-3000231122302010-2123131313211021) |
| `query_params.item.transformers` | [query_params.item.transformers](resources--service_policy_rule--reference--group-001.md#canonical-0201023212001002-2020130011313311-2230133011222110-2032132231300101-1303230311113220-3320322221201132-1332030201323221-2230300212220301) |
| `query_params.key` | [query_params.key](resources--service_policy_rule--reference--group-001.md#canonical-2100301210133310-2333223312301201-3131000303310131-1333223330212101-2121213331212121-2031222321033333-0202130301313312-3231000113313312) |
| `request_constraints` | [request_constraints](resources--service_policy_rule--reference--group-001.md#canonical-3321300223001100-2120321321003300-3311003010320003-2302220203322302-3333221002312102-0223300030120210-0203232233110201-3321030211133300) |
| `request_constraints.max_cookie_count_exceeds` | [request_constraints.max_cookie_count_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-2313200310211300-0211203202113123-0110012003032132-2111022330333210-2011232022312322-0302102323031210-2233021313131230-2331300203213300) |
| `request_constraints.max_cookie_count_none` | [request_constraints.max_cookie_count_none](resources--service_policy_rule--reference--group-002.md#canonical-3330003322002303-3112011103313310-2310133300001032-0301112113002100-2231301330310022-1113310301123111-1211300233030201-0231330331021111) |
| `request_constraints.max_cookie_key_size_exceeds` | [request_constraints.max_cookie_key_size_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-2020212101221102-2131103121230032-3211101120300013-1030032123131233-2001211302331313-0103012001113311-2113222133001310-3123100200033100) |
| `request_constraints.max_cookie_key_size_none` | [request_constraints.max_cookie_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0122132010120230-1020232110103112-3221012123023201-1301101202303202-0301310003310113-3210303313321301-0313331323231233-1323000201213320) |
| `request_constraints.max_cookie_value_size_exceeds` | [request_constraints.max_cookie_value_size_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-1021031032330313-1110311310220223-1130223231212220-1123010110233113-0020332200121200-1120100002222312-2321333002213310-0131331132002221) |
| `request_constraints.max_cookie_value_size_none` | [request_constraints.max_cookie_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0130202311032031-3021022220311213-2202111002300213-3112133122112130-3131100320201132-1230023003103202-2022301131222121-1130033202332102) |
| `request_constraints.max_header_count_exceeds` | [request_constraints.max_header_count_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-3223011123303132-0200013211133320-1333131213002323-3011012101112031-0231010330311111-1120113021100003-1212310200231232-0102032323220303) |
| `request_constraints.max_header_count_none` | [request_constraints.max_header_count_none](resources--service_policy_rule--reference--group-002.md#canonical-0231113303012011-0211220222302012-1200101001132312-3010010221132210-3233332300122222-0322021112300010-0300203121103023-3012210210021313) |
| `request_constraints.max_header_key_size_exceeds` | [request_constraints.max_header_key_size_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-0100220101123010-3112011323022110-1113131203013001-3322030123200030-2111302021101313-2133201313112312-0303122231321113-1113123000211112) |
| `request_constraints.max_header_key_size_none` | [request_constraints.max_header_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0203311333010103-1213002103233110-3332002301220100-0103103332020023-0001232020213231-0132101001323123-0021202133022020-3002032121223032) |
| `request_constraints.max_header_value_size_exceeds` | [request_constraints.max_header_value_size_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-2022303122012113-0223332122112030-2223312303331122-1120322031002123-0322301012001233-1331210232313113-1002202203022222-0300111231011013) |
| `request_constraints.max_header_value_size_none` | [request_constraints.max_header_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-3332010032003122-2002310031220112-0200323132113003-0232210311322332-1102020313301231-0100201222233130-1113113121320301-3123210320132232) |
| `request_constraints.max_parameter_count_exceeds` | [request_constraints.max_parameter_count_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-1233030203201113-1233220312102010-3013333203201322-1220312010103230-3222021021222031-2102103331312232-0212133332321003-0021011200132000) |
| `request_constraints.max_parameter_count_none` | [request_constraints.max_parameter_count_none](resources--service_policy_rule--reference--group-002.md#canonical-1110333230123311-0312322321312122-1330211131123201-0010320132111230-0111232322111312-0021010103323333-1112221012202133-1303112320333201) |
| `request_constraints.max_parameter_name_size_exceeds` | [request_constraints.max_parameter_name_size_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-0020332332023322-2123001321300233-0110333230222311-0032031322110002-1102101112103112-1033320301221033-1330032202122220-1302203122201011) |
| `request_constraints.max_parameter_name_size_none` | [request_constraints.max_parameter_name_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1223322223300233-0130113021331330-2323201312220132-1322022321310110-0202232103120301-1012123212220203-1011023221322021-3122033220301211) |
| `request_constraints.max_parameter_value_size_exceeds` | [request_constraints.max_parameter_value_size_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-0030320130001130-3333231301332021-2320213102111133-2332312131010011-2320030020100132-2333021032223221-3201132230222313-1302330101111230) |
| `request_constraints.max_parameter_value_size_none` | [request_constraints.max_parameter_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1323313230133220-0222102003010130-2130000311203101-2221202332011303-2303321031302101-3311022303220021-0120321203333320-1221213122130020) |
| `request_constraints.max_query_size_exceeds` | [request_constraints.max_query_size_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-3223021010002003-3011033230333121-3132233230301132-0133001310023221-3331100321113030-1233121111210133-2002313321310202-3330332331001332) |
| `request_constraints.max_query_size_none` | [request_constraints.max_query_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1203232220211220-0211100200300303-0022133022033311-1200031312211111-2322013231313110-0230232003011333-1100221300223112-3323222232010110) |
| `request_constraints.max_request_line_size_exceeds` | [request_constraints.max_request_line_size_exceeds](resources--service_policy_rule--reference--group-001.md#canonical-2111022021010231-3030203322312131-0133222011013033-2230101220213302-2323112200100313-2122120333133031-3020333331103222-3133012301010332) |
| `request_constraints.max_request_line_size_none` | [request_constraints.max_request_line_size_none](resources--service_policy_rule--reference--group-002.md#canonical-3033011001110323-3122020112332330-0230210321000010-1213233000333111-0210010303213123-2203220130031212-0021101202312223-1002021313203221) |
| `request_constraints.max_request_size_exceeds` | [request_constraints.max_request_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-3011003132011022-0130310200023122-2020210310200003-3111113130013010-1000222210303013-2302222331212011-1323123120232121-2303101010230003) |
| `request_constraints.max_request_size_none` | [request_constraints.max_request_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1110021232200100-3021311022230001-3331231112100033-1332110300222123-0222022201021032-2032300302303102-3200320111320133-3030022013230313) |
| `request_constraints.max_url_size_exceeds` | [request_constraints.max_url_size_exceeds](resources--service_policy_rule--reference--group-002.md#canonical-3002332102130132-1023102213020233-3301321313011303-3130311311302021-1211031221320201-3312013130003120-3001020230310222-0111101030121022) |
| `request_constraints.max_url_size_none` | [request_constraints.max_url_size_none](resources--service_policy_rule--reference--group-002.md#canonical-3212210122123013-0021120011333112-2333023110102022-1033202113011330-1332300213030323-1122132301332232-2333002000233000-2323333110322112) |
| `segment_policy` | [segment_policy](resources--service_policy_rule--reference--group-002.md#canonical-0000310300000020-3030312333212300-1033312300110000-2012001210311021-3333123312322130-3100112131120320-3322132313111200-1332233120101131) |
| `segment_policy.dst_any` | [segment_policy.dst_any](resources--service_policy_rule--reference--group-002.md#canonical-2332321232122233-1303323222111121-1220010021332231-2031211320212330-2001012101311130-0310102211102210-0322223312220030-2221230311322313) |
| `segment_policy.dst_segments` | [segment_policy.dst_segments](resources--service_policy_rule--reference--group-002.md#canonical-3113011300132303-2321300122002120-2020323213211011-1031022001212230-2122011332330022-3111201320100133-3002320222223022-0201022330113323) |
| `segment_policy.dst_segments.segments` | [segment_policy.dst_segments.segments](resources--service_policy_rule--reference--group-002.md#canonical-2202013210133013-0011231011230330-3111001023132023-1000201001113100-1231320031223030-1101010200302311-2021021111110302-0130312101001000) |
| `segment_policy.dst_segments.segments.name` | [segment_policy.dst_segments.segments.name](resources--service_policy_rule--reference--group-002.md#canonical-0031001033010112-3333100222122113-2333202111213332-1332030103113201-1113013303003120-0300211113000120-3223302103102220-0313301033231133) |
| `segment_policy.dst_segments.segments.namespace` | [segment_policy.dst_segments.segments.namespace](resources--service_policy_rule--reference--group-002.md#canonical-0213322130102120-3323021203320203-1220022200203100-1011323023311213-0332212223030110-0131231122113313-0110220312132221-3301321231031301) |
| `segment_policy.dst_segments.segments.tenant` | [segment_policy.dst_segments.segments.tenant](resources--service_policy_rule--reference--group-002.md#canonical-2221200313302300-2313130133003030-0001223103030203-0101131201221012-0011032323103103-1132112021123020-2231333133322031-2112333223333000) |
| `segment_policy.intra_segment` | [segment_policy.intra_segment](resources--service_policy_rule--reference--group-002.md#canonical-0230132313131123-0123032033333211-1203232110211110-3133123130013103-0333333132312203-3313310313311102-0133021021001210-3220000033031211) |
| `segment_policy.src_any` | [segment_policy.src_any](resources--service_policy_rule--reference--group-002.md#canonical-2133300232030230-3302313230113313-2230112133321120-1332211003133132-3132203311230011-0321232132203331-2303133222030121-3223131223223100) |
| `segment_policy.src_segments` | [segment_policy.src_segments](resources--service_policy_rule--reference--group-002.md#canonical-1010311123002311-3011223222303031-0010001002300303-1012200110130003-3312203003320212-0132220003222002-3010330210130202-2023310332233331) |
| `segment_policy.src_segments.segments` | [segment_policy.src_segments.segments](resources--service_policy_rule--reference--group-002.md#canonical-3033220131033301-0221220332120201-0230121221221110-1300033010331320-3231120233003320-1233310312232321-1213032210033323-0332010221113213) |
| `segment_policy.src_segments.segments.name` | [segment_policy.src_segments.segments.name](resources--service_policy_rule--reference--group-002.md#canonical-1021300030021222-1320103122013202-2022123003021031-0220021310110031-1312020031222001-0013203312301302-0332031320013132-0031333323032213) |
| `segment_policy.src_segments.segments.namespace` | [segment_policy.src_segments.segments.namespace](resources--service_policy_rule--reference--group-002.md#canonical-2313332313212032-1010011002013002-0003111333002113-0330032321023103-1321120322322300-0033020202201312-2331230000232031-1203010320231331) |
| `segment_policy.src_segments.segments.tenant` | [segment_policy.src_segments.segments.tenant](resources--service_policy_rule--reference--group-002.md#canonical-0301231033320131-0303133320301200-0021113210022300-0303123011102101-3022032301321332-1221021321211323-1330103010102103-1111212220302321) |
| `timeouts` | [timeouts](resources--service_policy_rule--reference--group-002.md#canonical-3220311022103123-2030031323233001-0321211311322230-1333222313202203-0233320303323213-3133330130021203-3233123132133020-0320211203311000) |
| `timeouts.create` | [timeouts.create](resources--service_policy_rule--reference--group-002.md#canonical-0222333130210210-2122032222313301-0321310123301310-3231302222101003-1030221121110102-2131020211020112-3112300131113212-2030203130111300) |
| `timeouts.delete` | [timeouts.delete](resources--service_policy_rule--reference--group-002.md#canonical-3310021130020030-3212212120032333-1100120210100100-3313113223313122-0222013210030010-0211003202300320-1211010013102103-1222023100022233) |
| `timeouts.read` | [timeouts.read](resources--service_policy_rule--reference--group-002.md#canonical-3013131233001012-2230120010033112-3331333322011332-3031030002030231-1222232132220233-0023310310001220-0212211202200112-2023303010113101) |
| `timeouts.update` | [timeouts.update](resources--service_policy_rule--reference--group-002.md#canonical-1120022113001331-2130320113031203-2010200322130313-2302311201332123-2312301110123333-1123322112031032-1122000320011102-2330131200103321) |
| `tls_fingerprint_matcher` | [tls_fingerprint_matcher](resources--service_policy_rule--reference--group-002.md#canonical-2300023332310032-3012332003021312-3311001110322002-3132320121320002-3212132123303111-2220301013203202-0001020200302013-3212122333321223) |
| `tls_fingerprint_matcher.classes` | [tls_fingerprint_matcher.classes](resources--service_policy_rule--reference--group-002.md#canonical-2213113001321332-2230010132132201-3131200121012322-3110112223120201-1022203210311023-1112312302231023-1032132201320101-1032321103021011) |
| `tls_fingerprint_matcher.exact_values` | [tls_fingerprint_matcher.exact_values](resources--service_policy_rule--reference--group-002.md#canonical-0311230130200231-1033100312020100-1322132333033120-0201003033310310-1023101322130221-0331223123301332-1332011031310202-1311100322001031) |
| `tls_fingerprint_matcher.excluded_values` | [tls_fingerprint_matcher.excluded_values](resources--service_policy_rule--reference--group-002.md#canonical-3120011321131203-0012123131322323-1101203311102121-3213221221210010-0302310333001031-1102100220102211-0000021211231302-1033300200332323) |
| `waf_action` | [waf_action](resources--service_policy_rule--reference--group-002.md#canonical-0112112222221031-1331310203113201-2000213313323030-0130211021321221-2010210010030031-3223302012121302-1312233021121013-0000331203202102) |
| `waf_action.app_firewall_detection_control` | [waf_action.app_firewall_detection_control](resources--service_policy_rule--reference--group-002.md#canonical-2302020223330013-0320000111031132-0032321313011101-1220221302020001-1023310022311120-3220322131302302-2112320233110232-1201003221033023) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts](resources--service_policy_rule--reference--group-002.md#canonical-0001100132222033-3233203111303132-1302101100302133-0001120212102001-0013332011013322-3332101122303200-2221332012000110-1332301022301122) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context](resources--service_policy_rule--reference--group-002.md#canonical-1112112313312322-3303313323212300-1000023203011031-2113002332231213-3101113000020001-2033311222112200-1300131112311013-3331021112301001) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.context_name](resources--service_policy_rule--reference--group-002.md#canonical-1002102131033021-1002323121230322-3221012131011021-2310003221323200-2122221201033021-0021231103230130-0330311130120232-2110300102211033) |
| `waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` | [waf_action.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type](resources--service_policy_rule--reference--group-002.md#canonical-1013333333330111-3301332333313001-0332010122001102-2333102303321222-2120201302231111-1302302312201221-2232022023200030-1001113202120131) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts](resources--service_policy_rule--reference--group-002.md#canonical-2231121312223131-1320132030312333-2213010200222123-3220300221010031-1332021211103133-3322232020002310-1330100201031010-3023130323130302) |
| `waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` | [waf_action.app_firewall_detection_control.exclude_bot_name_contexts.bot_name](resources--service_policy_rule--reference--group-002.md#canonical-1321301102011012-0310203102012333-2113223321021332-3113032022330013-3211001120101012-2333013030133200-1330031210330003-1010332331213121) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts` | [waf_action.app_firewall_detection_control.exclude_signature_contexts](resources--service_policy_rule--reference--group-002.md#canonical-2002232121300211-3232321313232120-2123203121030122-1110010100222330-3220123313310231-1203132200202311-2202013020330233-0330220012112210) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context](resources--service_policy_rule--reference--group-002.md#canonical-0302331300021002-2202300131013122-3312102321333212-1331330313212033-0203031212000012-3112230120120031-3303100213201232-1000032133303030) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.context_name](resources--service_policy_rule--reference--group-002.md#canonical-0211210033323210-1303120212112233-0220101030201301-1001122232200022-0330130023310100-2323332112012113-3021323233203322-3101001013023222) |
| `waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id` | [waf_action.app_firewall_detection_control.exclude_signature_contexts.signature_id](resources--service_policy_rule--reference--group-002.md#canonical-1303233330233001-2100200201011001-3133220013222203-0131100220332012-3113211013321200-2031322310321033-3212003110011311-0113123121223023) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts` | [waf_action.app_firewall_detection_control.exclude_violation_contexts](resources--service_policy_rule--reference--group-002.md#canonical-2002213110202130-2303030331103001-2313312323000001-1321231022313032-2011322002310132-2110121132211122-2230111031233322-3100013033001321) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context](resources--service_policy_rule--reference--group-002.md#canonical-2213132312321330-2111101320303132-2310111300031003-3311110023111200-2013133303030133-0000211013301221-1113132131211113-0121201011120111) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.context_name](resources--service_policy_rule--reference--group-002.md#canonical-0312113122213230-1031201020012332-2011002031322133-2322321323001012-3331200211033320-0011331021302131-0233301011003231-3213212101100231) |
| `waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` | [waf_action.app_firewall_detection_control.exclude_violation_contexts.exclude_violation](resources--service_policy_rule--reference--group-002.md#canonical-0333110202323322-1120213031231030-3332332221131011-0302102232020300-2012033000201110-1130100133333112-1210023010302103-3201000102230023) |
| `waf_action.none` | [waf_action.none](resources--service_policy_rule--reference--group-002.md#canonical-0203313012102132-1232220131100013-2102313200231111-2202023011200331-0321000311321331-2112212300301323-2303313202121100-0203313013320102) |
| `waf_action.waf_skip_processing` | [waf_action.waf_skip_processing](resources--service_policy_rule--reference--group-002.md#canonical-3313002212212011-1012223132001111-2033223121201032-0333003111331031-1323210230303010-3113320221320021-2300200010122121-0103203202012311) |

<a id="canonical-0221232230332013-1321203101233021-0003232132312013-0013203002120133-2300131221032201-0023201233030211-0232122000323001-1021303221010210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_asn` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- any_asn

<a id="canonical-2021232023103020-1311221231322320-2222121231030210-3332011313321110-0002000113221131-1012300202111122-2321323031130323-2103001211202222"></a>

Type: `["object", {}]`. Optional.

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

- [any_asn](resources--service_policy_rule--reference--group-001.md#canonical-2021232023103020-1311221231322320-2222121231030210-3332011313321110-0002000113221131-1012300202111122-2321323031130323-2103001211202222)
- [asn_list](resources--service_policy_rule--reference--group-001.md#canonical-2102013201202122-0110010331120030-0132033210110300-0310100302323203-0032200323333001-0333031101321330-3032130003323011-2032123322202222)
- [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1221102131001321-0211132300133300-3032122013030021-3330012330203320-1111302112032022-2010003103010103-3212002003231302-2032112213032110)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_asn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123012100122112-0010012320202103-1121103211320203-3102030000330101-0132100103211313-1111033030002021-2011133021332311-2330013222310310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_client` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- any_client

<a id="canonical-0332123200212031-2133003002311213-2132022222131132-3321011301221221-2201101101131212-3221312231000232-1021221221230320-3331123312230010"></a>

Type: `["object", {}]`. Optional.

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

- [any_client](resources--service_policy_rule--reference--group-001.md#canonical-0332123200212031-2133003002311213-2132022222131132-3321011301221221-2201101101131212-3221312231000232-1021221221230320-3331123312230010)
- [client_name](resources--service_policy_rule--reference--group-001.md#canonical-3210022202131202-2200000220321321-3121211013333232-2312333100032232-1221220300133213-1102211020130221-2200131102321312-1131313313300320)
- [client_name_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1112333331011102-2032331320312222-0130031223211101-3311210300301112-2323213120110120-3212330201033321-1113032330323313-1211222302020221)
- [client_selector](resources--service_policy_rule--reference--group-001.md#canonical-1100200013101011-2030121221031312-1013010130232011-0331330201133123-3232201232320220-1001101313033120-3032323121002232-0331120132130013)
- [ip_threat_category_list](resources--service_policy_rule--reference--group-001.md#canonical-0302200111131213-0020202112102320-2202002130110021-3103300012200032-1332110312113231-3320210001123321-3333203121031111-1100132203212331)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323033333212103-1131233331312011-2132333310220111-1000201311003232-3201032233230003-0111300321031231-0000200020031132-3212101202111300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `any_ip` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- any_ip

<a id="canonical-2332201302210331-2302231200033202-0311033331123333-1123320002301122-0102102012203001-3233102001102210-2233331031332022-3323231200100331"></a>

Type: `["object", {}]`. Optional.

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

- [any_ip](resources--service_policy_rule--reference--group-001.md#canonical-2332201302210331-2302231200033202-0311033331123333-1123320002301122-0102102012203001-3233102001102210-2233331031332022-3323231200100331)
- [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1202321232220311-2202112122103101-3030011112103123-0020033132221020-1311010010031120-1101202223111322-0322311103130230-0130112001011022)
- [ip_prefix_list](resources--service_policy_rule--reference--group-001.md#canonical-1133010200322311-2203212333112121-2022320311033220-1222320331002313-2302333111223002-3112000131122011-0102211300332020-0320031012310030)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001200123112031-0101111212320231-0323322202023310-1033112332303103-2323023202232130-1220331203300031-3220022122303100-3013111133233130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_group_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- api_group_matcher

<a id="canonical-0020300113321302-2000001122310033-0200310032213011-0210023311311321-1010130332032110-2013022222333332-3233210310122300-3311101013003322"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies a list of values for matching an input string. The match is considered
successful if the input value is present in the list. The result of the match is inverted if
invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("match")}
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
api_group_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322113003321113-1232213130100323-0320112112111120-1000021032331220-3321333213033300-0031123113221210-1223311021013022-0322110000012301"></a>

### Direct properties for `api_group_matcher`

<a id="canonical-0312320211233133-0313020131133210-0220331203302000-2201130201030200-2030103100110132-1130323221131331-3120331123302210-1212323231131203"></a>

#### `api_group_matcher.invert_matcher` property

Type: `"bool"`. Optional.

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

<a id="canonical-0123011011030001-1131331003222003-3330122131002100-0010110220022302-3332300230332033-0112233210100221-1331223321213311-1223323233200120"></a>

<a id="canonical-2033333131022213-0020320321310333-3301100230311310-2000101301201102-1113320013331032-0033011010220021-1122320231000331-0002031233013121"></a>

#### `api_group_matcher.match` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-1032023103021302-0223330032311303-2203320210330011-1003220302332130-3220031001022020-2032013022121012-0303233103102223-2031323210322110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `arg_matchers` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- arg_matchers

<a id="canonical-1222000023023301-0330002201002230-1222123321113320-1030320320222321-3022201123210032-3323011102313303-3031222002101112-3330212311132101"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

Terraform syntax:

```terraform
arg_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121010312022212-1022223333112023-3300212210211023-2031330003121301-3021023210323112-3112201130321231-3223001211121233-2230102323021231"></a>

### Direct properties for `arg_matchers`

- [check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-1221111202311312-0313312121110232-3032221101101202-1030222010000113-0223233211102003-0122112210311023-0033223223120233-2001311233123032): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-001.md#canonical-1001201231101011-0103320013111003-1011211222100311-3002301322320112-0103022302231001-2232301212333021-3300120302121333-3111033200222201): complete subsection reference.

<a id="canonical-2302221102322002-3020100222130332-0102033012320033-1311020322203203-1200321221131100-0003120303030202-3300233132121111-1210100112010003"></a>

<a id="canonical-2020102223232220-0033110202310210-1101003221032112-3333121112322220-3310000133200312-1303133110221231-0110213333001300-2221101020302310"></a>

#### `arg_matchers.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--service_policy_rule--reference--group-001.md#canonical-1223223233022123-1303233122202133-1311232203213010-1333321320320013-3100303311010202-2111101203322013-0303030003302222-3032101131001223): complete subsection reference.

<a id="canonical-2231202131023212-1003033232003033-3232033322210012-3230201322322023-1033001032021201-1201101213211200-0212202210101233-1003311301311202"></a>

<a id="canonical-2022110100212133-1320103311332130-0101101222020102-2112231022100333-2200201220021223-3210012021122322-2301212311031130-1010012203010001"></a>

#### `arg_matchers.name` property

Type: `"string"`. Optional.

A case-sensitive JSON path in the HTTP request body.

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

<a id="canonical-1221111202311312-0313312121110232-3032221101101202-1030222010000113-0223233211102003-0122112210311023-0033223223120233-2001311233123032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `arg_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-1032023103021302-0223330032311303-2203320210330011-1003220302332130-3220031001022020-2032013022121012-0303233103102223-2031323210322110)
- arg_matchers.check_not_present

<a id="canonical-0201310201202121-0203012332231103-1132022032321312-3102123333330120-0013112011110203-0321332323223101-3311200220220100-0220113132131010"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001201231101011-0103320013111003-1011211222100311-3002301322320112-0103022302231001-2232301212333021-3300120302121333-3111033200222201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `arg_matchers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-1032023103021302-0223330032311303-2203320210330011-1003220302332130-3220031001022020-2032013022121012-0303233103102223-2031323210322110)
- arg_matchers.check_present

<a id="canonical-0101120133232000-3101033203320303-3222100012102001-1313133203302312-1332200321303302-3312302310323301-2020322312313233-1111312312221320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223223233022123-1303233122202133-1311232203213010-1333321320320013-3100303311010202-2111101203322013-0303030003302222-3032101131001223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `arg_matchers.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [arg_matchers](resources--service_policy_rule--reference--group-001.md#canonical-1032023103021302-0223330032311303-2203320210330011-1003220302332130-3220031001022020-2032013022121012-0303233103102223-2031323210322110)
- arg_matchers.item

<a id="canonical-2021321123232221-1030303120131220-2001200100013111-2233210321011133-3022023002213201-1023103103011321-3121131101032212-2301333312311002"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011012131131001-3313323220223220-0313312100321032-0230331213233003-1302203211023311-0131100003323301-1332021010221330-1011000130220023"></a>

### Direct properties for `arg_matchers.item`

<a id="canonical-3321312331100031-1031120000202223-1101122322323321-3201111221122212-3230312231013011-3320123320133303-3301213101112030-2330210230322312"></a>

#### `arg_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-3332131232013320-1220011323023332-2300030023000332-3000101321020231-1103301313011322-3010321322301130-1300022201120003-0111303121310223"></a>

<a id="canonical-1230332220010002-0000101113003201-0313121302222201-3330122302110032-3123302313020030-3331011021102100-3321131033130331-2212232111032320"></a>

#### `arg_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-3221220311013111-1102122110013000-3013023101202300-1311020221301333-1011211222120321-3000121300001200-2220221221131310-2210111223112323"></a>

<a id="canonical-2232303002131131-2123033010230311-3212313330032002-2101211120130210-3031311021302332-0021002000212120-2122303203122321-0221223023320333"></a>

#### `arg_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-0021332221230033-0003220101123220-2020212222231011-0302311332321131-3023201300001102-0300232311002101-1131213221031303-1201030020323010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `asn_list` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- asn_list

<a id="canonical-2102013201202122-0110010331120030-0132033210110300-0310100302323203-0032200323333001-0333031101321330-3032130003323011-2032123322202222"></a>

Type: `"object"`. single nested block, Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011103123131320-0213313010113230-1230200002011202-1012311203023131-0132011113033111-0310002033102300-3003132123122010-2123213030331311"></a>

### Direct properties for `asn_list`

<a id="canonical-1303000313130111-2210003313010021-3312021301100030-0003323212322133-0313112330211110-1111311032222121-0213020313211030-2121033003303120"></a>

#### `asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1303131313113301-2002011311201202-0113230121212323-0202022331303111-2010020122331301-0002311130331320-1111101133131202-0133331233122210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `asn_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- asn_matcher

<a id="canonical-1221102131001321-0211132300133300-3032122013030021-3330012330203320-1111302112032022-2010003103010103-3212002003231302-2032112213032110"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100102021011321-1110320121022103-1323030331202133-3031123132223331-3303200320223023-3122310300230210-1032000221200330-2201100200321012"></a>

### Direct properties for `asn_matcher`

- [asn_sets](resources--service_policy_rule--reference--group-001.md#canonical-3310000112003100-1333100110033121-1133032303032212-1020213013010122-2323233101233301-1310330203010010-0033121210211033-1221223332300313): complete subsection reference.

<a id="canonical-3310000112003100-1333100110033121-1133032303032212-1020213013010122-2323233101233301-1310330203010010-0033121210211033-1221223332300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [asn_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1303131313113301-2002011311201202-0113230121212323-0202022331303111-2010020122331301-0002311130331320-1111101133131202-0133331233122210)
- asn_matcher.asn_sets

<a id="canonical-0302030310311111-0213312232300232-1033212032111210-1333110021211213-1212033321100020-1012010103311002-3302331112220202-3130303222122223"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111021022310033-2332311313333210-1121022211222233-2110331312321130-0030033232322102-2100011311031222-2313220222332011-1011320101121201"></a>

### Direct properties for `asn_matcher.asn_sets`

<a id="canonical-2112123220221221-1230301211021312-0311111002031002-1232032123221323-2111120233020330-0132033031320332-1120320033123213-0001000130223222"></a>

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

<a id="canonical-2201313200031103-2312310101001013-2302231211113003-3330303113203011-3303010321013112-1232233303300122-3203103312201000-1032322122322131"></a>

<a id="canonical-2121322312000121-1033101323313110-3200201223222211-3320103100003023-2000003001301202-1023211200301031-1012221221011011-2323312122233100"></a>

#### `asn_matcher.asn_sets.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1210131120210100-3321132310133131-1231113330332113-0223003303212133-3300110122012230-3102000232022032-1330013233113110-2222120332103222"></a>

<a id="canonical-0112133133233312-3120010011221130-0021002112132220-3310132030302310-1311213030333122-2203332303210203-2130220100021120-1031211113322133"></a>

#### `asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-2312222230230110-0013130002131321-0310302111111033-0312222010203113-2331223301132222-3112031313233003-0200132300031112-0132102311011120"></a>

<a id="canonical-0031113131313233-3002002113002012-3113013001310110-2232111133030313-0031232022023031-1222102003203011-3330122201000003-2202322121030310"></a>

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

<a id="canonical-2100130123023201-1001023220222322-0101202223322020-1321011301021303-2322020232213032-0233010003310232-2031223032122102-3200032311132032"></a>

<a id="canonical-1220111302222200-0323310101303013-0133131203032310-1121332200233001-1201121102021322-0030013131310311-2230031322112010-2213012212203301"></a>

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

<a id="canonical-1001101320132112-3130213301123030-0312023110021102-1100300111322002-1210010333200210-2020102132223133-2012322132222311-2003212133113201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `body_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- body_matcher

<a id="canonical-1122022323301132-1233100323013212-3122031331313102-0121301331200011-2133221311020330-2012233300032203-3123221112100101-3023322130333203"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
body_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010300322020012-0332303211021003-0222121123312303-0111021211220120-0333012112211030-1311203232000101-1233002302113123-1001232120112232"></a>

### Direct properties for `body_matcher`

<a id="canonical-2312120231210202-2121001330101303-3310130211123110-3021111113200323-0012211003120233-0032112333022311-2121002021211222-2000223302113123"></a>

#### `body_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-3233001100122201-3023001313201223-1103223310011230-0110213323203223-2200233111320111-3200221311211301-3002133302333303-2002303030131231"></a>

<a id="canonical-1310203220323323-0033122101312313-3302000001010033-1323100122031002-0000023303332303-2130322220013300-3012220330000110-1003302120113211"></a>

#### `body_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-2113301112001032-1312310022101200-0112311313122310-1123000031313302-3231122331021301-3223020101131212-2012210323100202-0021332003021203"></a>

<a id="canonical-0001102312003323-0323023100023322-0130303330321230-0001331312120233-2030230121302132-2001130231232222-2020203313012130-1322310032003111"></a>

#### `body_matcher.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-2103123012221321-1103321133011201-1313232020222001-2012201222102131-0300222221031123-2133203022221201-3200132333302332-1223010122321102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_action` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- bot_action

<a id="canonical-0211333322223333-0012021033002113-3113330301122031-3133300333333103-3302312120011311-2133220231211123-2211120222101323-1233311200003102"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("bot_skip_processing",
    "none")}
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
  "x-ves-oneof-field-action_type": "[\"bot_skip_processing\",\"none\"]"
}
```

Terraform syntax:

```terraform
bot_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012231113120032-2222201123331231-0101011011132313-1301133130220232-2122200211231013-1011022201211233-3012220313300112-3233212313312101"></a>

### Direct properties for `bot_action`

- [bot_skip_processing](resources--service_policy_rule--reference--group-001.md#canonical-0233301013220001-0030131113322111-0021333223232101-3120032120321323-3011203023132101-1331322332323111-0000100230312200-1011303013122202): complete subsection reference.

- [none](resources--service_policy_rule--reference--group-001.md#canonical-3122010311313110-2132131202133003-2311312313110302-2233020200211031-0213121313300102-1301033230330100-0013111110100012-3230001210223102): complete subsection reference.

<a id="canonical-0233301013220001-0030131113322111-0021333223232101-3120032120321323-3011203023132101-1331322332323111-0000100230312200-1011303013122202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_action.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-2103123012221321-1103321133011201-1313232020222001-2012201222102131-0300222221031123-2133203022221201-3200132333302332-1223010122321102)
- bot_action.bot_skip_processing

<a id="canonical-2332113222010211-1120003301003132-1120011131222103-1221131231203032-1222203010220220-1001200310221033-2221333003103221-0320121310000013"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
bot_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122010311313110-2132131202133003-2311312313110302-2233020200211031-0213121313300102-1301033230330100-0013111110100012-3230001210223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_action.none` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [bot_action](resources--service_policy_rule--reference--group-001.md#canonical-2103123012221321-1103321133011201-1313232020222001-2012201222102131-0300222221031123-2133203022221201-3200132333302332-1223010122321102)
- bot_action.none

<a id="canonical-2002301121233310-3003301303200212-2032210101120212-2033333010111110-3031310313021210-3220000020312130-3333101003311210-3112103312101102"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332123330301021-1100223121211022-3121010103210131-3033010222100300-1123002321301311-1112213102202113-1311100312010010-0101133331030013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_name_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- client_name_matcher

<a id="canonical-1112333331011102-2032331320312222-0130031223211101-3311210300301112-2323213120110120-3212330201033321-1113032330323313-1211222302020221"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
client_name_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111002233003301-0230102020230313-3130132033022101-0031312103231030-1232100011011202-2032311003212303-2013213231202100-1331130221233101"></a>

### Direct properties for `client_name_matcher`

<a id="canonical-1110002302110013-0211112020210133-2330013320302103-2333201210010220-0330030032311210-1222301210133232-2220203311112332-1330310222230330"></a>

#### `client_name_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-3010122320302101-1303110131221000-2032012203011012-1210233202233331-0223332303223200-1200212221321311-2102321232103232-1333000300210202"></a>

<a id="canonical-2002102003210011-3132110033213302-1210313002100023-3122010302312120-3100330311200230-2013332011003203-3102010130101110-3103030120030111"></a>

#### `client_name_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-1300312101033300-3320221301112302-3333112131101103-3101332010233112-1220113220003101-2212302222200022-2133323313312132-2031301301211130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_selector` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- client_selector

<a id="canonical-1100200013101011-2030121221031312-1013010130232011-0331330201133123-3232201232320220-1001101313033120-3032323121002232-0331120132130013"></a>

Type: `"object"`. single nested block, Optional.

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
EnumExtractionComplete: false
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032230212032320-1102022111100311-2302230313130322-3300313200132020-1020133012223212-1121323331112333-2302021311213003-3103102333120201"></a>

### Direct properties for `client_selector`

<a id="canonical-0201022300100121-2133322313133122-3312000321132200-3320031311133131-1003232121111211-2030122012103133-0131300111022233-1322000120220320"></a>

#### `client_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0323223303303012-1323121213223000-3303312201212002-2112132210332220-0300220220230320-3122301023330023-1133223001001010-2203310132010333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_matchers` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- cookie_matchers

<a id="canonical-1303013313221111-3330333231120001-1133203100133112-0131212300130023-1202020011131101-2312130111203201-2320322122113212-0010302221203100"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221220131332133-3021200302123112-0231323103231222-2120010003010130-1330323122210123-0013111220223113-2000222203130022-2003223012331010"></a>

### Direct properties for `cookie_matchers`

- [check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-2332100200033123-1121121103323301-1302112312031130-1030031133332310-1330303000013021-2231310101203301-2210302000133032-2232222303013203): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-001.md#canonical-0301100030213101-1111323301332122-1110231222230111-0301002120302310-2122012122030133-1220031231310322-3001102022032332-3122110120122303): complete subsection reference.

<a id="canonical-0012303103120200-1111131102121321-1003222131132232-3330203031011113-3020332003102200-0031010023013231-0313230121232101-3313133300302323"></a>

<a id="canonical-2100030232203222-1131200221012033-3133200311233103-2112211132331331-2020001213001210-0301111320112110-3000232021212231-0120333221301323"></a>

#### `cookie_matchers.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--service_policy_rule--reference--group-001.md#canonical-0030031330001303-3132122220201213-0322030033331323-3223212001023033-2111210301300310-0010310131021333-3131000032322232-3011201023111310): complete subsection reference.

<a id="canonical-0201213331312022-2313222020132223-0011303111302113-3003312311232033-3210122332001210-1120302122101103-1111330121121230-2302203100121223"></a>

<a id="canonical-2032330113103310-2222311223131131-3033111312131313-3020313310000113-2123312030023322-0113111122112322-0312322031301013-1120130022013110"></a>

#### `cookie_matchers.name` property

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

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

<a id="canonical-2332100200033123-1121121103323301-1302112312031130-1030031133332310-1330303000013021-2231310101203301-2210302000133032-2232222303013203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-0323223303303012-1323121213223000-3303312201212002-2112132210332220-0300220220230320-3122301023330023-1133223001001010-2203310132010333)
- cookie_matchers.check_not_present

<a id="canonical-2020323310202031-1110323301130033-3101112230021230-2223322303011113-3300220123020101-1211300212203333-3310210312133013-0303221022121310"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301100030213101-1111323301332122-1110231222230111-0301002120302310-2122012122030133-1220031231310322-3001102022032332-3122110120122303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-0323223303303012-1323121213223000-3303312201212002-2112132210332220-0300220220230320-3122301023330023-1133223001001010-2203310132010333)
- cookie_matchers.check_present

<a id="canonical-3100000301121103-3131133003130030-1030211212330230-3201021020321220-2111232012233001-0333110331302021-2131010333223332-0332211231022032"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030031330001303-3132122220201213-0322030033331323-3223212001023033-2111210301300310-0010310131021333-3131000032322232-3011201023111310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [cookie_matchers](resources--service_policy_rule--reference--group-001.md#canonical-0323223303303012-1323121213223000-3303312201212002-2112132210332220-0300220220230320-3122301023330023-1133223001001010-2203310132010333)
- cookie_matchers.item

<a id="canonical-3132212312222031-2330110023002212-1321033222303002-1211022200321232-3123211210312133-0121321232200202-3132232121213203-3320322331321303"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111313332033103-3302222212101231-0210210100010210-1103300100321020-1330030321103020-0122000003223120-2111121301023233-0231032012310303"></a>

### Direct properties for `cookie_matchers.item`

<a id="canonical-0110320321201122-0303000122100323-2132202201010121-0101111211310232-3310100331223202-2312332120002013-1111122210211312-3103211133022213"></a>

#### `cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-0322003132210311-2132333213120032-3203313011303121-3323000123232023-1032002330103120-2122321320222130-2233210302133311-1002222331120011"></a>

<a id="canonical-2022033311222202-2332200021131122-3102223000222221-2233230203220223-1012212033332012-3323111322321033-2100021232012012-0221111132030031"></a>

#### `cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-3323213031202023-3321211300311001-1331013100111111-2120213001011231-0101103202203310-1320332133212220-0302200323113003-0103121213310232"></a>

<a id="canonical-2312023312202023-0310203122323113-0321312102203313-1321011022310001-1303131132000323-0021300212033320-1201010223022130-3112333131320000"></a>

#### `cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-1332222130011133-0323233220013211-0330111222302013-3023320213212310-2333013213211313-0332132110001121-3220010230201023-2032103202233210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domain_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- domain_matcher

<a id="canonical-3031311003030133-0022010310120120-2212123301010100-2110310032302121-1330130231313112-3220222112111310-0102133201222320-0311111131321321"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300121322232223-2123312003310103-2231101302101113-3310233322012002-3021001111003132-3110232103333002-0311230101300311-3213121200323220"></a>

### Direct properties for `domain_matcher`

<a id="canonical-3011113220102231-1033210201321331-3103012112123212-2323233010212113-1211203210101022-0232303022122302-3000122322112033-1330201210220123"></a>

#### `domain_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-2230211000033202-2300001210321311-2103313201113101-1120330200020033-1232310313103310-1333101222012020-1221021322222320-2322333012023111"></a>

<a id="canonical-0101113001213122-3033311113313031-0213221011233202-3010120302230212-1301223302131002-3222220123103313-0213313001302033-1233000301212001"></a>

#### `domain_matcher.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-2020321201213200-3003101032002102-0003211310323331-0102232130323303-3322002211110311-3103210330303202-1130102333320230-1030212213102000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `headers` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- headers

<a id="canonical-3100231310321321-1112023301102132-2321103011213002-0221111301202230-0320230313103212-1012230100011130-1010213320020130-1331120322021311"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003020011131003-1322011203203002-3022012030130113-1121130122101210-2331001023321012-2200312032020030-2230122020110201-0303232001111300"></a>

### Direct properties for `headers`

- [check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-3301330211120321-0330221230100220-3302230311001230-3301221011010101-1210123012102200-0203201312303003-2213112321030120-3032120233122010): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-001.md#canonical-0120232011121103-3033223223013030-1320310310232213-3033003001301031-1202023313200122-3211311212001120-0330330202323031-1233223333200123): complete subsection reference.

<a id="canonical-0201322233330101-0213311201330121-2210302013301321-2030230222130101-0311013001212102-2023100013013322-0021311022223030-0330210322202203"></a>

<a id="canonical-1030023311332103-2320223202001101-2302320100123121-2102320213202102-3023311111223302-1002010001313202-0331233001132300-0321031101001101"></a>

#### `headers.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--service_policy_rule--reference--group-001.md#canonical-3003210030231111-2201132003202010-1201112211110122-0103021212200303-2210221210012010-0213120031010030-2320003302032010-1032212303033020): complete subsection reference.

<a id="canonical-3330222210300200-2312111100120301-0120223332223033-0113233103323332-0113131133001103-3020002021301333-3211033321011300-1130010121211322"></a>

<a id="canonical-3133222130003020-1022222302303210-1201331121212322-0310200200333232-0203332101030131-3003331002322102-0133301302022232-3321001310333310"></a>

#### `headers.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

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

<a id="canonical-3301330211120321-0330221230100220-3302230311001230-3301221011010101-1210123012102200-0203201312303003-2213112321030120-3032120233122010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `headers.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [headers](resources--service_policy_rule--reference--group-001.md#canonical-2020321201213200-3003101032002102-0003211310323331-0102232130323303-3322002211110311-3103210330303202-1130102333320230-1030212213102000)
- headers.check_not_present

<a id="canonical-0123223012103301-1212221122310102-1022100222001222-1311012231211230-0222020323023210-3230331110322210-1213220213021210-0202020212221221"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120232011121103-3033223223013030-1320310310232213-3033003001301031-1202023313200122-3211311212001120-0330330202323031-1233223333200123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `headers.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [headers](resources--service_policy_rule--reference--group-001.md#canonical-2020321201213200-3003101032002102-0003211310323331-0102232130323303-3322002211110311-3103210330303202-1130102333320230-1030212213102000)
- headers.check_present

<a id="canonical-3331101312120303-2320133322302031-3112310130213130-1120030312131023-0320011001030000-2110020311223122-2300302110330310-2212030031011321"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003210030231111-2201132003202010-1201112211110122-0103021212200303-2210221210012010-0213120031010030-2320003302032010-1032212303033020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `headers.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [headers](resources--service_policy_rule--reference--group-001.md#canonical-2020321201213200-3003101032002102-0003211310323331-0102232130323303-3322002211110311-3103210330303202-1130102333320230-1030212213102000)
- headers.item

<a id="canonical-3002312200232200-3330330332211333-2030133223131002-2002223222032320-3002301212020313-2330203231323302-3322030031223333-0212011110201120"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033323003321013-3023203222202011-0130212110323032-0123102210100221-2231131303212132-3310221223030212-0023002313133332-0333102302230210"></a>

### Direct properties for `headers.item`

<a id="canonical-0000011130223002-0000022233333002-0033301113031120-0323202221111320-3022132213130222-0012233021123033-0312201101033202-3201321331320223"></a>

#### `headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-1123022002311322-0130001121222110-0123120333221302-1321201323322120-2023312203020203-0320311020023201-1022332030222322-3312120033211330"></a>

<a id="canonical-3321323311222301-0112222103223103-1201222011200302-2211033113011021-3200323213102202-1002310133132010-3130012331300310-1031013203032231"></a>

#### `headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-2210322210303121-2033320113321003-0231332213331211-0103013213121331-0310331310122302-0311010323200012-3011131303010220-0012033202110303"></a>

<a id="canonical-1213112133010100-2223010310011102-1002332123310202-0101310123320323-3100120223111122-1202121232300113-0123103031223232-0131321222133122"></a>

#### `headers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-3311223333320012-1301003001231331-1200011030303311-0111131102332203-0033001320311120-3102002212303313-3213012110031222-2313312311022232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_method` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- http_method

<a id="canonical-1312132130110201-0020131303101121-1001103103301103-2311301310022310-3310002321111100-0310131213202012-1302232332321213-3113321302232333"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
http_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312001231113301-1301232300332213-3112302323233233-2123201210121221-2212313133212303-3110313131113200-0202223010220000-2032323120030122"></a>

### Direct properties for `http_method`

<a id="canonical-0002002131101130-3022011120230131-2103001023212210-0022231310131000-1033200203311332-1003123210333030-2103222112333120-0210011331332302"></a>

#### `http_method.invert_matcher` property

Type: `"bool"`. Optional.

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

<a id="canonical-1103132021332032-0302033230313120-1112302322210321-1011303211133230-2221102331130331-2012212123031333-3110301303031133-0120001301123322"></a>

<a id="canonical-0121313310012110-2112311220322232-3123123233221033-2223330132023211-0303232203311323-3103112102310310-2031131330032303-3221321323212101"></a>

#### `http_method.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-1101030111001132-2130113101200012-1213033232233301-0331310212100221-2201110133311321-1321032003101012-0133122013101313-0213333230233312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- ip_matcher

<a id="canonical-1202321232220311-2202112122103101-3030011112103123-0020033132221020-1311010010031120-1101202223111322-0322311103130230-0130112001011022"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211110122022200-2030003221001000-3013111100012321-3112311130300231-1002003323330220-1113013321010213-1110301113133012-1313201010320312"></a>

### Direct properties for `ip_matcher`

<a id="canonical-3303301031033301-0101313213031213-3312003230101212-0121010123113210-0200321030120311-1133233332021111-3320223331002000-3201331101203130"></a>

#### `ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--service_policy_rule--reference--group-001.md#canonical-0223113303331231-2132300210111111-1031012201033300-2022032222300011-3230123131322120-0001203203222221-3232302002233112-1132001223000312): complete subsection reference.

<a id="canonical-0223113303331231-2132300210111111-1031012201033300-2022032222300011-3230123131322120-0001203203222221-3232302002233112-1132001223000312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [ip_matcher](resources--service_policy_rule--reference--group-001.md#canonical-1101030111001132-2130113101200012-1213033232233301-0331310212100221-2201110133311321-1321032003101012-0133122013101313-0213333230233312)
- ip_matcher.prefix_sets

<a id="canonical-3233013301230123-1131113210312133-0113101120022233-0201201323332003-1112221022101002-2302121130302222-2000230100232022-3001020313111201"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320000110333130-0212110121002200-0221323023200131-3202200021132131-3113130120200301-3222230332311020-3123331020020031-0101031002320303"></a>

### Direct properties for `ip_matcher.prefix_sets`

<a id="canonical-3021012023330123-1212213123033200-0211103110122102-3103202302012231-0303101321033032-2100322301023201-2023202200311320-2303011333331021"></a>

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

<a id="canonical-3033013011102301-3220112000022223-3200032230222210-3021013231020330-2332103303333211-1020330210200313-3330310110320100-1132123311110303"></a>

<a id="canonical-1020032323032331-3220333333333220-0230111102122002-1022321310213211-2122113113032330-0333323033332000-0200030201221311-3321100231031313"></a>

#### `ip_matcher.prefix_sets.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0310000233130111-1320322323112300-2231312223231321-2221011220212232-1010031220213333-0210130332303123-3032111312110221-0201320321021102"></a>

<a id="canonical-0130032333232121-1001110311310212-2133220220221211-2101030310122121-1021230213010022-0002220201010021-1020210302033311-2233310103132220"></a>

#### `ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-1212123300023101-2332121200133103-2021203122012313-2000001322123321-1113230031211332-1101230232112231-2222332033110330-3223033213121023"></a>

<a id="canonical-1122023010112012-2111320011332020-0023320321203021-3123111023312200-0012332221101133-3130310303131232-0200130010301312-3120332202131233"></a>

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

<a id="canonical-2110212113221010-1213210023120200-2100101120133011-3333011230110203-0031203111331023-3201202031302102-1202222321211233-0213200031223202"></a>

<a id="canonical-0020032112000321-0211133023012332-3022223333133200-0333331000000311-3002013203001031-2212322301330220-2122322132222232-2100031013132123"></a>

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

<a id="canonical-1001031102013032-3012221302301103-1333012032022110-3313212031020033-3133033332330121-2311012133200203-3200331331023200-2110011322231321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_prefix_list` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- ip_prefix_list

<a id="canonical-1133010200322311-2203212333112121-2022320311033220-1222320331002313-2302333111223002-3112000131122011-0102211300332020-0320031012310030"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031213132202013-3221330310223200-0103101301133120-1130000232010302-1120101101012200-0320212100333130-0002211011000012-1311012003012130"></a>

### Direct properties for `ip_prefix_list`

<a id="canonical-1130333130002213-2122330233122100-2332030230323321-2200003210200111-0201000310211110-1133303300233302-1000312222030221-0232320121013020"></a>

#### `ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

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

<a id="canonical-0010320331303221-2231012110002312-1012001322030131-1130311133000123-1030121102000320-0001313132200321-2122003330311100-1232020033033011"></a>

<a id="canonical-3031011123232232-2131030312311320-1113321230311102-1202110020221202-1130022002231203-3331121312230133-1112332011232113-0123332010123102"></a>

#### `ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0012302311320221-1202233202103232-2011312303031222-3122303101021120-0303100330112201-0323021021121103-1213030202102213-1311133121210233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- ip_threat_category_list

<a id="canonical-0302200111131213-0020202112102320-2202002130110021-3103300012200032-1332110312113231-3320210001123321-3333203121031111-1100132203212331"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122330132100222-1132201203313322-2123322002211112-3233120002010222-3010010312302113-0322230011323001-1211210232123300-3303311303112221"></a>

### Direct properties for `ip_threat_category_list`

<a id="canonical-2223300023201310-3013310111311301-1323233022033033-1000120312211111-1122120031101301-0000023210101132-2320130221002322-2311323332203202"></a>

#### `ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2232311312222031-0121212211131221-3223322103010300-2012020332313011-1130321013122111-3223223033331333-0020100121032303-3310222113002133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ja4_tls_fingerprint` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- ja4_tls_fingerprint

<a id="canonical-2320112003011023-1000221201331002-0111220231302002-1211120101020002-1331131133033203-2220033003013011-2031203330133310-0311101013313003"></a>

Type: `"object"`. single nested block, Optional.

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

- [ja4_tls_fingerprint](resources--service_policy_rule--reference--group-001.md#canonical-2320112003011023-1000221201331002-0111220231302002-1211120101020002-1331131133033203-2220033003013011-2031203330133310-0311101013313003)
- [tls_fingerprint_matcher](resources--service_policy_rule--reference--group-002.md#canonical-2300023332310032-3012332003021312-3311001110322002-3132320121320002-3212132123303111-2220301013203202-0001020200302013-3212122333321223)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ja4_tls_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203312233301122-2123111210113030-0012000210333233-1312202222220312-3210330021332023-2012331300331032-3001001103302232-3103223000312122"></a>

### Direct properties for `ja4_tls_fingerprint`

<a id="canonical-2320321310112103-2210311303010102-1310000112233211-3201211311332100-3011302022230030-1202131312310223-0213323210013133-3201131331010320"></a>

#### `ja4_tls_fingerprint.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_claims` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- jwt_claims

<a id="canonical-0232111131010322-1023012332212110-0323323020302020-0011233321330000-1212323132301302-3033220331130113-0000321120220200-1122231031303301"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322222121113131-2121332023222001-2113310302001023-0311130032023032-3221230213322201-2232331011001233-0030221300013012-0233032030211221"></a>

### Direct properties for `jwt_claims`

- [check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-2303320103130132-2101102022031220-3233210122010200-3132323231322222-1202332312102313-0031113302231103-1221230211023220-1321231232313002): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-001.md#canonical-2031030212122331-1231122323313101-1302300202103032-0110222223333122-0001021313011023-2200310200000100-3330131313202003-3320120231032210): complete subsection reference.

<a id="canonical-1101030231301201-0101102022222020-2230211101302233-3011100312001211-1311032321002301-0230303211003202-0032101032133013-3312221012033300"></a>

<a id="canonical-2221011330122213-3332302032030231-1301320231221322-1302000222223210-2312212101113332-1032302301223232-0001222311300111-3211302203223211"></a>

#### `jwt_claims.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--service_policy_rule--reference--group-001.md#canonical-2311232010330201-3102323300101200-2132333132303330-0332202102312221-0220131212103330-0210002331320300-1023120010330210-2111203310303201): complete subsection reference.

<a id="canonical-1231130012032130-2301020010201101-2200222213010320-2333130103320103-1223033230213030-2311233211103102-3302030332233213-0331113201210101"></a>

<a id="canonical-2322011031311333-1001013201302123-0210302012303231-1313032300011200-1310032300010321-3202121232003121-2210103030322132-3132122101011112"></a>

#### `jwt_claims.name` property

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

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

<a id="canonical-2303320103130132-2101102022031220-3233210122010200-3132323231322222-1202332312102313-0031113302231103-1221230211023220-1321231232313002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113)
- jwt_claims.check_not_present

<a id="canonical-3211003220132220-2111200033001312-0211232230111330-3000011132030100-2110300023103031-0203130031210320-3200222002201133-3221130312102132"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031030212122331-1231122323313101-1302300202103032-0110222223333122-0001021313011023-2200310200000100-3330131313202003-3320120231032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113)
- jwt_claims.check_present

<a id="canonical-3010233313020311-0333203110111132-3110113212022133-1020022110221232-3232011133320320-0003022122322222-0123212200202213-1032013222333231"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311232010330201-3102323300101200-2132333132303330-0332202102312221-0220131212103330-0210002331320300-1023120010330210-2111203310303201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_claims.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [jwt_claims](resources--service_policy_rule--reference--group-001.md#canonical-3121210103300010-0232302121023210-3221102033333021-3313032210213000-2323222213320112-2302302332002332-3203103000231332-2012033133032113)
- jwt_claims.item

<a id="canonical-3122203122311031-0320223010303311-2303023021112000-1200201201000221-2321323301132011-0002230313021110-1312102132223122-3221301301011011"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021222330100113-3301003320111303-2311303300202321-2303032030200022-1223231301133011-1121012132100100-2322103031020022-3103220220303110"></a>

### Direct properties for `jwt_claims.item`

<a id="canonical-1131203020211002-1132013133221130-2330323330200033-3320120000301132-2021301102313331-2333232232133302-2001322133101011-1003030101331331"></a>

#### `jwt_claims.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-0333211231120231-2221101333230013-2332220003030120-0320011313320131-3322220312003013-0201102031323122-3111010002110233-0121101120223200"></a>

<a id="canonical-2201201221201131-1221222320003230-0112100301330321-2120111111213313-0023111300110130-3231302323313202-2222310312021203-3130123200131322"></a>

#### `jwt_claims.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-2201100002003013-0013312020120131-3123210331020232-3233123332033123-3321302130020302-0013031102222131-3132331313323100-2203111311220022"></a>

<a id="canonical-0311131213032130-3212233221030123-1312211210333311-3230230302130022-2320010333103301-1000121230223023-2200132231100310-3211000110030022"></a>

#### `jwt_claims.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-2020111211100111-2033231320322121-0121111211102333-3113322310323132-1300301223000233-0231032111303233-0333010313331210-0100231321203000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `label_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- label_matcher

<a id="canonical-1312003200110013-1123113121200330-0303232330132123-3130321201330321-2333032023303210-3110021030303000-3113023202210211-2302312332233300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123110323220300-0113311120120012-1110011111232311-3011302013222013-1311202220110311-1030303332223020-0313231332131313-0203132011330123"></a>

### Direct properties for `label_matcher`

<a id="canonical-2303011232220331-2131220312311033-1100103323302001-1132322033300202-0223223333113312-1103211202000323-0011012300003323-0312120022330012"></a>

#### `label_matcher.keys` property

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mum_action` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- mum_action

<a id="canonical-2221121310120320-0122232013020111-3111313002230103-0012100300022122-2233030220321313-1231102130223133-0323221002032002-3032000311011100"></a>

Type: `"object"`. single nested block, Optional.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default",
    "skip_processing")}
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
  "x-ves-oneof-field-action_type": "[\"default\",\"skip_processing\"]"
}
```

Terraform syntax:

```terraform
mum_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220201102001210-0021123312200103-1001201203032311-3023013020310220-2022203111122001-3221121000230013-3221223220132012-1311000103132021"></a>

### Direct properties for `mum_action`

- [default](resources--service_policy_rule--reference--group-001.md#canonical-3211102001223032-0311103320033133-3130122202310202-0331113011003213-3201120320303100-2321003200213302-1200120123321323-3112121031201102): complete subsection reference.

- [skip_processing](resources--service_policy_rule--reference--group-001.md#canonical-0122323332132112-1310201121311332-0000021312313130-2303132030300200-1103301333032313-0023213311211230-3300103103220222-3232213011100333): complete subsection reference.

<a id="canonical-3211102001223032-0311103320033133-3130122202310202-0331113011003213-3201120320303100-2321003200213302-1200120123321323-3112121031201102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mum_action.default` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [mum_action](resources--service_policy_rule--reference--group-001.md#canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131)
- mum_action.default

<a id="canonical-2102113111213301-3030032320301122-0101232002213212-0322211221310200-1101210110003000-2131012033221002-1022221121113013-1031100210203302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122323332132112-1310201121311332-0000021312313130-2303132030300200-1103301333032313-0023213311211230-3300103103220222-3232213011100333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mum_action.skip_processing` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [mum_action](resources--service_policy_rule--reference--group-001.md#canonical-2001130322132301-0223123123302223-3201113011220210-2021122302201100-0100103003130323-3131100121201120-2120001312003313-2010030020332131)
- mum_action.skip_processing

<a id="canonical-0312310211231330-1203021111031102-3103233002012120-1210130133030212-1322020300113131-3311120303211301-3330102301313021-0213020111213311"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010110322330023-1103101211122201-1133323113223332-1313232321223110-0101110303212213-1030200111110001-0302033113300012-0303132213131330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `path` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- path

<a id="canonical-1222332103311013-2322202031102010-2310333211132301-2133303132123131-3101032220033221-2332323101120012-0023222202301233-0100221110233300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030213002123113-1100011232010032-1333302032330312-0122202210023012-2302023020220130-3101002112120131-2201322011210322-1231012300310233"></a>

### Direct properties for `path`

<a id="canonical-1203321001120102-3230123100011102-3310122110021233-1010032220322303-1321202311111101-2101232203313223-3031113123221003-0123331113312113"></a>

#### `path.encoded_path_matcher` property

Type: `"bool"`. Optional.

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

<a id="canonical-3300211023023011-0100330002020032-0221232220001003-2023130313203113-1110113023021102-0203110332303122-0301132030231013-0232211211002022"></a>

<a id="canonical-3232113221313231-2201233211310220-3203302110222111-2000010103322220-2331100303102222-2123310030022131-3120211312102221-2102220032213021"></a>

#### `path.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact path values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-0122202103323232-3332313000012033-0031003032201200-1303213100000221-1122210032020210-3221322012000233-0230203021001300-2003012113330110"></a>

<a id="canonical-3102120230201203-0121323013002121-3133210112231203-3003210320020332-0031000130300303-2131121202132230-1321120132202013-1022321320211312"></a>

#### `path.invert_matcher` property

Type: `"bool"`. Optional.

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

<a id="canonical-0033131231323313-2133303031222313-3330221011331312-0132022203002230-3233300231311030-1232003330022330-1003213132011232-2033012301112130"></a>

<a id="canonical-3133110010320230-1322122213310131-2032031023031003-1102021120023220-3131330003233132-3310033033002113-3003220221020022-3111330302130323"></a>

#### `path.prefix_values` property

Type: `["list", "string"]`. Optional.

A list of path prefix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-3231203123331011-1010021300103133-3100010310231333-2112002200321102-2213301011312013-2121110313020232-3301210032010221-2203323010311321"></a>

<a id="canonical-3320212203220302-3002023322330112-2111301302122100-3121110100120013-1320122311220331-0100200101223101-2010302020302032-0221211001200222"></a>

#### `path.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-0012332022230031-0201113331012212-0221033003102023-2030201101331300-2322202311021232-1102202303102103-0022332222100002-3123213102000001"></a>

<a id="canonical-0331321321312231-2301111111330300-2113300312322310-2002112313202333-2030033311113010-1222200031223121-1211113313013020-0122312113032000"></a>

#### `path.suffix_values` property

Type: `["list", "string"]`. Optional.

A list of path suffix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-2332221310210123-1303210232003020-1011212330201122-0320023322033321-0332301020311103-1223013320302002-0213010132303030-1212211131110230"></a>

<a id="canonical-1121102022212021-0301212113202233-1020220211131331-1012113032201213-3310122112012130-1000233211201110-2212303220300322-0001312112200231"></a>

#### `path.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-2203222212003231-2202300100330200-1310211032320202-1313000110311333-0331313210023112-2112333302313030-1110023000021210-2033311231020133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `port_matcher` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- port_matcher

<a id="canonical-2322112120320222-3113003012110013-3110231221003322-0221120100321303-2132110222311133-2030211130312102-0231232102133322-2023023231200211"></a>

Type: `"object"`. single nested block, Optional.

Port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true. Server applies default when omitted.

Additional upstream details:

A port matcher specifies a list of port ranges as match criteria.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
port_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032130301303201-3310332220312231-2110203123201232-2330310112331221-2012131233111210-0011001330332132-0110220232022122-1113011011030131"></a>

### Direct properties for `port_matcher`

<a id="canonical-0120221311003100-1000230031222100-1322323000312032-1031130332211231-1210111013010030-3210012031021302-3103210102000132-3032321110230021"></a>

#### `port_matcher.invert_matcher` property

Type: `"bool"`. Optional.

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

<a id="canonical-2223320321103203-0023122212323001-3002223020320021-3213202020100230-1231130222132002-0132303232200030-3131023001013101-2010121223331303"></a>

<a id="canonical-0232303133221212-3332131223200113-2121331003110032-3131033211230031-1201202200123312-2131122022232322-1103300110112000-1131113331022031"></a>

#### `port_matcher.ports` property

Type: `["list", "string"]`. Optional.

List of strings, each of which is a single port value or a tuple of start and end port values
separated by '-'. The start and end values are considered to be part of the range.

Additional upstream details:

A list of strings, each of which is a single port value or a tuple of start and end port values
separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `query_params` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- query_params

<a id="canonical-1132121303300022-3222130331310311-3012221303320222-0300303322210021-2012131020223121-0203302231323312-3113331121030333-0011112032010322"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301200300003021-2303331200211221-2123130200012202-1201200111103323-1113302221122103-2112000032311330-1033133302222322-0301010333112101"></a>

### Direct properties for `query_params`

- [check_not_present](resources--service_policy_rule--reference--group-001.md#canonical-2233100010311321-0013300020103330-2232123000232130-2313333300323232-2021121122103012-3321120200110032-2222013322013111-3111323320011102): complete subsection reference.

- [check_present](resources--service_policy_rule--reference--group-001.md#canonical-2213203220132122-1123311113310120-0200132121322222-1320110200011231-3213213300100303-2100122210033231-1120321110122000-1001021132121000): complete subsection reference.

<a id="canonical-2033101023010311-0310112222301102-2100232312001011-1100313301133203-1131022123000031-0102213230111001-1303220330120300-0212300011311202"></a>

<a id="canonical-1211110211301031-0103111210102021-2201033020001011-3103130010320120-3301002313013200-0322010123020301-2202101203100230-2231303100212332"></a>

#### `query_params.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--service_policy_rule--reference--group-001.md#canonical-1311331210221220-1223101220322011-0301012322132023-0333111323012133-0011313130123200-3102231330313002-2322220322101021-1202323313133012): complete subsection reference.

<a id="canonical-2100301210133310-2333223312301201-3131000303310131-1333223330212101-2121213331212121-2031222321033333-0202130301313312-3231000113313312"></a>

<a id="canonical-0321302201211212-2332030000302032-0001133000113303-1120123103130103-1021010003321133-1310113312132210-1130111210132302-2102223001323221"></a>

#### `query_params.key` property

Type: `"string"`. Optional.

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
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

<a id="canonical-2233100010311321-0013300020103330-2232123000232130-2313333300323232-2021121122103012-3321120200110032-2222013322013111-3111323320011102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [query_params](resources--service_policy_rule--reference--group-001.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101)
- query_params.check_not_present

<a id="canonical-3202330001122331-1130330220001301-2201120212330221-2000333031220132-0331013333332133-0030011200120031-0232302232121212-1200221333202002"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213203220132122-1123311113310120-0200132121322222-1320110200011231-3213213300100303-2100122210033231-1120321110122000-1001021132121000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `query_params.check_present` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [query_params](resources--service_policy_rule--reference--group-001.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101)
- query_params.check_present

<a id="canonical-0203031220002032-2321320122323321-0330112110213330-0210112333032311-2300003111301101-1133223023031032-1103032000122312-3123133310113121"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311331210221220-1223101220322011-0301012322132023-0333111323012133-0011313130123200-3102231330313002-2322220322101021-1202323313133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `query_params.item` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [query_params](resources--service_policy_rule--reference--group-001.md#canonical-1012032130001102-2011001331322032-1103021200202322-0322031221221110-0133202010012100-1121312101103032-2022003321212212-0312123221010101)
- query_params.item

<a id="canonical-2233103232121213-2131110023313322-0313133320302120-1121000310110033-2111303130113330-0132132000302301-2101021032003213-3132031301203312"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012110112031103-0020130230300002-1103023101012131-0223013000331203-0113220000000012-0330103031000232-1313332132202000-2320310332221133"></a>

### Direct properties for `query_params.item`

<a id="canonical-3323111210031331-1133123113121110-3131322121022313-1122302110012211-0011213013231301-2301000132230213-3332312100301303-2113311203232112"></a>

#### `query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-1003310032001103-3332232022122202-1031211322230031-3033231223011030-3201321033301331-3002013130130221-3000231122302010-2123131313211021"></a>

<a id="canonical-3201022212200211-3222202301022310-3020023331110231-3012123201122001-3103121010033012-3120210102302122-1122000020300130-2201030112122010"></a>

#### `query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
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

<a id="canonical-0201023212001002-2020130011313311-2230133011222110-2032132231300101-1303230311113220-3320322221201132-1332030201323221-2230300212220301"></a>

<a id="canonical-2303101212301033-3103313231103000-0133330213200011-2112312002300320-3333221102010333-2300323013303120-1222002232123123-0320020012001133"></a>

#### `query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-3222030332220031-3233311201132021-2133203221221312-2032310111020301-0033321331133110-0320111032022023-2223313312230311-2022001130202012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_constraints` properties

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032)
- [Property reference](resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- request_constraints

<a id="canonical-3321300223001100-2120321321003300-3311003010320003-2302220203322302-3333221002312102-0223300030120210-0203232233110201-3321030211133300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request constraints.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_cookie_count_exceeds",
    "max_cookie_count_none"),
  validators.ConflictingObjectAttributes("max_cookie_key_size_exceeds",
    "max_cookie_key_size_none"),
  validators.ConflictingObjectAttributes("max_cookie_value_size_exceeds",
    "max_cookie_value_size_none"),
  validators.ConflictingObjectAttributes("max_header_count_exceeds",
    "max_header_count_none"),
  validators.ConflictingObjectAttributes("max_header_key_size_exceeds",
    "max_header_key_size_none"),
  validators.ConflictingObjectAttributes("max_header_value_size_exceeds",
    "max_header_value_size_none"),
  validators.ConflictingObjectAttributes("max_parameter_count_exceeds",
    "max_parameter_count_none"),
  validators.ConflictingObjectAttributes("max_parameter_name_size_exceeds",
    "max_parameter_name_size_none"),
  validators.ConflictingObjectAttributes("max_parameter_value_size_exceeds",
    "max_parameter_value_size_none"),
  validators.ConflictingObjectAttributes("max_query_size_exceeds",
    "max_query_size_none"),
  validators.ConflictingObjectAttributes("max_request_line_size_exceeds",
    "max_request_line_size_none"),
  validators.ConflictingObjectAttributes("max_request_size_exceeds",
    "max_request_size_none"),
  validators.ConflictingObjectAttributes("max_url_size_exceeds",
    "max_url_size_none")}
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

Terraform syntax:

```terraform
request_constraints {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010313321232010-1200202213331203-2130000310110112-2303030200003332-3110133011013202-1111201031323112-0033031012122320-3010030330311323"></a>

### Direct properties for `request_constraints`

<a id="canonical-2313200310211300-0211203202113123-0110012003032132-2111022330333210-2011232022312322-0302102323031210-2233021313131230-2331300203213300"></a>

#### `request_constraints.max_cookie_count_exceeds` property

Type: `"number"`. Optional.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_cookie_count_none](resources--service_policy_rule--reference--group-002.md#canonical-1122022011000132-2310232120103311-1302322200300230-2231133020012123-0101113203221130-3103331101010130-0200032210220122-2333200330022310): complete subsection reference.

<a id="canonical-2020212101221102-2131103121230032-3211101120300013-1030032123131233-2001211302331313-0103012001113311-2113222133001310-3123100200033100"></a>

<a id="canonical-2132312311232302-1133303013120203-2312321213222231-1123223111003300-1012221031131120-0033112132033121-3030303223110113-1300123130012223"></a>

#### `request_constraints.max_cookie_key_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_key\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_cookie_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-3210311330313210-1030032003013232-2310012332113123-0200213031011110-3032031210102331-2131213220101131-3321201211111213-3111123300123122): complete subsection reference.

<a id="canonical-1021031032330313-1110311310220223-1130223231212220-1123010110233113-0020332200121200-1120100002222312-2321333002213310-0131331132002221"></a>

<a id="canonical-3100102323203333-1221332333203330-1100023312210003-3033232022101202-0303211133222322-2303021002100332-3100111131302000-2130133300121013"></a>

#### `request_constraints.max_cookie_value_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_cookie\_value\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 32768),
}
```

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

- [max_cookie_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2210332330211102-3323302021220120-0333020011333130-3231233221330020-0311010101023031-1213120210302130-3032020321123303-3220013033000320): complete subsection reference.

<a id="canonical-3223011123303132-0200013211133320-1333131213002323-3011012101112031-0231010330311111-1120113021100003-1212310200231232-0102032323220303"></a>

<a id="canonical-2311303021203232-0132001003313010-2013333110323331-2000130331120023-2130010023132310-3020130113023310-0302301223031112-2222020022000033"></a>

#### `request_constraints.max_header_count_exceeds` property

Type: `"number"`. Optional.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 40),
}
```

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

- [max_header_count_none](resources--service_policy_rule--reference--group-002.md#canonical-2031333020012231-3100132320333220-1321002111211212-1121322112320113-1033112301122111-3221032333001212-3332000110133331-1031013122313000): complete subsection reference.

<a id="canonical-0100220101123010-3112011323022110-1113131203013001-3322030123200030-2111302021101313-2133201313112312-0303122231321113-1113123000211112"></a>

<a id="canonical-3331002233210112-3302321030230122-3322223201031102-0011212113332322-1220021231132332-1303213101320332-3330212001232000-3213232211313332"></a>

#### `request_constraints.max_header_key_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_header\_key\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_header_key_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1132101021313102-3332033311133003-2210133001000322-0010232331203032-2322133101011321-1202032320332312-3132033012032313-3332110002331110): complete subsection reference.

<a id="canonical-2022303122012113-0223332122112030-2223312303331122-1120322031002123-0322301012001233-1331210232313113-1002202203022222-0300111231011013"></a>

<a id="canonical-1132031131300201-1120302210233032-3002311311002011-1302200221013230-1100110333132002-0100133311322013-3131203030000022-2122110001012031"></a>

#### `request_constraints.max_header_value_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_header\_value\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 64000),
}
```

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

- [max_header_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2102120230210231-2100101103110120-1123203011201001-3031333211322330-1210231111011330-2310102033001102-1320213011013311-0021120132333103): complete subsection reference.

<a id="canonical-1233030203201113-1233220312102010-3013333203201322-1220312010103230-3222021021222031-2102103331312232-0212133332321003-0021011200132000"></a>

<a id="canonical-2222310223030030-3202100200200003-1021110310233331-1122212320100120-0233132131212220-3323013121020100-0131113121033021-3301112103333111"></a>

#### `request_constraints.max_parameter_count_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_count\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_parameter_count_none](resources--service_policy_rule--reference--group-002.md#canonical-3221300301123333-0233121001221021-1311302003000011-2010203103033011-3100131213301022-0011232323122220-0300231113313310-1313310111333232): complete subsection reference.

<a id="canonical-0020332332023322-2123001321300233-0110333230222311-0032031322110002-1102101112103112-1033320301221033-1330032202122220-1302203122201011"></a>

<a id="canonical-2110132002230123-2133210233202002-0310103021133012-1033200101002110-0120232233303030-1330312012011331-1230012103030112-3020001110133302"></a>

#### `request_constraints.max_parameter_name_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_name\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1024),
}
```

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

- [max_parameter_name_size_none](resources--service_policy_rule--reference--group-002.md#canonical-2230202133131132-2331002310102321-2203332230123312-3120102201131031-2000222113110033-2122102213110001-3312010333210032-0032213011111222): complete subsection reference.

<a id="canonical-0030320130001130-3333231301332021-2320213102111133-2332312131010011-2320030020100132-2333021032223221-3201132230222313-1302330101111230"></a>

<a id="canonical-1303120311332031-2112221311120100-2131203221102302-2212232021033332-1303033030220112-2033220010303000-0033102331021133-1133312121100201"></a>

#### `request_constraints.max_parameter_value_size_exceeds` property

Type: `"number"`. Optional.

Exclusive with \[max\_parameter\_value\_size\_none\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 1073741824),
}
```

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

- [max_parameter_value_size_none](resources--service_policy_rule--reference--group-002.md#canonical-1023121102232230-3110132110130012-2003011223012323-0330131103322022-1233311030100300-0113123310232012-3233100301000331-3200002330112200): complete subsection reference.

<a id="canonical-3223021010002003-3011033230333121-3132233230301132-0133001310023221-3331100321113030-1233121111210133-2002313321310202-3330332331001332"></a>

<a id="canonical-1132331031231100-0031203203302003-1011221311120002-3233330122311320-0003203031220202-3132201113030200-3202103232131203-1223031332120231"></a>

#### `request_constraints.max_query_size_exceeds` property

Type: `"number"`. Optional.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 60000),
}
```

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

- [max_query_size_none](resources--service_policy_rule--reference--group-002.md#canonical-0100103003131210-2220030203003130-1213013113101012-2123313132020002-3211122013110230-0202020232331211-3111300202321122-1301331102130200): complete subsection reference.

<a id="canonical-2111022021010231-3030203322312131-0133222011013033-2230101220213302-2323112200100313-2122120333133031-3020333331103222-3133012301010332"></a>
