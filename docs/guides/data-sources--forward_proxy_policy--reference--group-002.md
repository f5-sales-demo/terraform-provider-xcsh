---
page_title: "xcsh_forward_proxy_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy reference."
---

# xcsh_forward_proxy_policy reference

<a id="canonical-0020332333203301-2021310003101213-3320300012002303-0310102300321113-3030210300310120-1330001133003332-2023331302203022-2123021322002233"></a>

## `network_connector.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2300322032101212-2102011102022232-1123320022202301-0232110132313232-2102300220002003-2112303223102033-2111222121323110-2233110200123221"></a>

<a id="canonical-2223110222303222-3023133122311123-0320330103021002-0133000100130330-0132230132330101-3101202321130020-0001232021333210-3322322301123013"></a>

## `network_connector.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1032301033033313-1101300232301332-0110322320310030-1032321212200203-0212320010031030-1100000123210022-0031113133102112-0301111301311103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_label_selector` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- proxy_label_selector

<a id="canonical-3320221102222111-2312332233123123-0022022013122310-0231213333221113-3301223112111020-0120030021112201-1210012120013200-1210110111112213"></a>

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

<a id="canonical-2102102213213222-0232300100102113-3211210123020002-3332032021233120-2133110103222122-1033000030212130-2233102103121310-0103311130322233"></a>

### Direct properties for `proxy_label_selector`

<a id="canonical-1300023203010112-2003122212123121-0320231300323302-0312320001120013-2122100132003011-2212000122231010-3133130330220312-0311103002312001"></a>

#### `proxy_label_selector.expressions` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- rule_list

<a id="canonical-3300003111010003-1022200013220321-2120312030000021-0322110310133010-2300202200330301-2102222012123212-3322012103010130-2332110210313311"></a>

Type: `"single"`. Computed.

Custom Rule List. List of custom rules.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1131103020121130-2231320103311321-3021310333212030-0210020132313111-0202211203211120-3333320233113230-3332113232021010-0230330310223032"></a>

### Direct properties for `rule_list`

- [rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003): complete subsection reference.

<a id="canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- rule_list.rules

<a id="canonical-3100302032212022-2021123222100320-3213331000002321-2003220103311020-0101020103332002-1313332302120213-0321120011110223-2110233133303132"></a>

Type: `"list"`. Computed.

Custom Rule List. List of custom rules.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0333202101113131-1322003211033310-1000011220000220-0103303023120310-1321213101302202-3301211203323012-2002203010000200-0313220121111032"></a>

### Direct properties for `rule_list.rules`

<a id="canonical-1131000023102020-1231010222000020-3001331200220123-0203303331330312-1031123113111032-1000301210102321-2102313200023203-1200130222231333"></a>

#### `rule_list.rules.action` property

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [all_destinations](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0132013321333203-2123012010320331-2230021013321013-1130111030022220-1132232330000022-2311021020311230-3133132103021322-3113330232312103): complete subsection reference.

- [all_sources](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1033231212331302-0311032231200320-1111232330003011-2021232302220313-1200220323300203-2323112301323222-3001111212012001-3300332012333313): complete subsection reference.

- [dst_asn_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1300121121133123-1331331113003123-2123002310303113-2213203323123203-2322233201100012-2333000202102110-2233223130210112-3013002000222300): complete subsection reference.

- [dst_asn_set](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0333101311200103-1201232200123231-1312310331322313-0330220231113033-1023301323021201-2233333012002321-0200302211110200-2303223023100212): complete subsection reference.

- [dst_ip_prefix_set](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3313033010201213-1303212323211203-3223031032231111-1000013301232331-2032030123132131-2222301002313201-2232022003332321-3221120203213211): complete subsection reference.

- [dst_label_selector](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0321111022030032-2322302230031301-2330333002000223-1223111132313113-0013121312101130-2001323322322232-3010000201300030-2311113000302123): complete subsection reference.

- [dst_prefix_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1302021011333221-1202321123122020-2132230322001012-2021003030013130-3003132133103002-0232330130210222-0130211121032322-1300232013202233): complete subsection reference.

- [http_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211): complete subsection reference.

- [ip_prefix_set](data-sources--forward_proxy_policy--reference--group-002.md#canonical-2212103130223310-3330130302022032-0203030031113130-1003332132223122-1023111013210012-0030311012030221-2111023131032223-2232213101213000): complete subsection reference.

- [label_selector](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0121013023233311-1203120030030033-1332300301332300-0321210320000300-1112203010322331-1133332312311203-2212003212000211-0102321303223132): complete subsection reference.

- [metadata](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3223102201230333-0133323300113202-0020231011110213-1103132313011100-2303122321311003-3120230013203320-0203112302310020-0011101020213312): complete subsection reference.

- [no_http_connect_port](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3322110302231022-2313001321321311-0011112313123332-0331303200302032-1300033123310301-3131201202203220-0023301312310322-2323311123323123): complete subsection reference.

- [port_matcher](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3303223031202122-2130021033020311-3000130130333300-0233223101103033-0220002000332013-0100230323032120-1303100210111331-2311331133310002): complete subsection reference.

- [prefix_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3212130112120111-0303230130202111-3022100232223233-3002201320311201-2201332023302332-2211231010010220-3022300313111010-2321121102331231): complete subsection reference.

- [tls_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3130031113132122-1313310010332323-3012313301202331-1020032123210331-2033230333200121-0311201313100230-3313331323102333-1023320301301321): complete subsection reference.

- [url_category_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-0233202033222111-2122021223212100-2100002000120303-2133131211221023-2302211312302313-1031130223233300-0112000112200222-3200310020303232): complete subsection reference.

<a id="canonical-0132013321333203-2123012010320331-2230021013321013-1130111030022220-1132232330000022-2311021020311230-3133132103021322-3113330232312103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_destinations` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.all_destinations

<a id="canonical-3202113122323120-0131313120111220-3331020322313230-3012320022131230-3333301200321111-2210213110123333-1330333230301321-2002120302021203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all destinations.

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

<a id="canonical-1033231212331302-0311032231200320-1111232330003011-2021232302220313-1200220323300203-2323112301323222-3001111212012001-3300332012333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_sources` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.all_sources

<a id="canonical-0232322210231213-1331210222033021-1312001113030130-2111300303102332-2121333301322131-2302331100020102-1210101111321300-3032213222012221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all sources.

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

<a id="canonical-1300121121133123-1331331113003123-2123002310303113-2213203323123203-2322233201100012-2333000202102110-2233223130210112-3013002000222300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_asn_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_asn_list

<a id="canonical-0303030201113121-1103123211313320-2021122332033320-3110030310233020-3022311312010321-0002213022211101-2320012313111133-0132122022111120"></a>

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

<a id="canonical-3111333312011231-2222232133021202-2301302121213223-3323122010302021-0320303313201200-3002320022102302-1323131101322312-3100211101022211"></a>

### Direct properties for `rule_list.rules.dst_asn_list`

<a id="canonical-2003230122023011-2121213210123031-1113031023122221-2023230133300111-3123122013003332-2131330132000112-0031101033020121-1131233311303333"></a>

#### `rule_list.rules.dst_asn_list.as_numbers` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0333101311200103-1201232200123231-1312310331322313-0330220231113033-1023301323021201-2233333012002321-0200302211110200-2303223023100212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_asn_set` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_asn_set

<a id="canonical-3111311212203313-3311102310322000-2032122321103130-2302311032021331-1101223222020021-1221122022032002-0310111021133330-3110122110213322"></a>

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

<a id="canonical-2123311302223000-2103010031232022-3231011021220202-3030031310120231-0301232123102111-2123230332103221-3302133331322100-2021212021310113"></a>

### Direct properties for `rule_list.rules.dst_asn_set`

<a id="canonical-1101221330120033-1333020323112103-1323002010222002-2102032223122320-1132110312103222-2200012123201213-3112232111211030-1210132333120311"></a>

#### `rule_list.rules.dst_asn_set.name` property

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

<a id="canonical-2120023233120332-2313201203332121-3221303013000302-3313110123232132-3222003033113030-3112012021003122-3003202012203030-2221013220020202"></a>

<a id="canonical-3012302200103203-1121212122021230-1200032111231321-2001003310312110-2110300133012120-3330003111100202-2002022023202010-3023230103313021"></a>

#### `rule_list.rules.dst_asn_set.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1122030031232123-1012003211300232-1033023221232233-3000333032030121-1133000013010303-1323213322122323-3201110332302110-0211000113131201"></a>

<a id="canonical-1131333211321331-3330100303003312-0111011022023213-3302022312032101-3233320120103003-3220032321213210-1131022212231220-1310030012010012"></a>

#### `rule_list.rules.dst_asn_set.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3313033010201213-1303212323211203-3223031032231111-1000013301232331-2032030123132131-2222301002313201-2232022003332321-3221120203213211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_ip_prefix_set` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_ip_prefix_set

<a id="canonical-1103131002311232-1320220001232303-1312231231010331-3322012200212232-3223223122313132-2032031102120010-3322321001210302-2101003302223312"></a>

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

<a id="canonical-0032200220230332-3020330300211231-0211010322121110-0203222331313331-2011331333011012-2020100230201323-0331132131022121-0223332012030013"></a>

### Direct properties for `rule_list.rules.dst_ip_prefix_set`

<a id="canonical-2322102303003010-2022232312113301-1331313022032220-1131321010233223-1300303320133011-3231313011201333-0020111332302310-1132121021332323"></a>

#### `rule_list.rules.dst_ip_prefix_set.name` property

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

<a id="canonical-1330211321130102-3133131322230002-3231200232110233-1010310030122001-1312311213001131-2211112210021303-2333021103132022-0020101222301320"></a>

<a id="canonical-0123123110233000-0232230203123101-0002033101020121-3221101023220021-1032331312023231-1120311100123030-0032322023122332-0212103113203223"></a>

#### `rule_list.rules.dst_ip_prefix_set.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2232311300100111-3031232132100333-1033202333201013-3012320212003300-0311313133332212-0202323320231222-3120131210002123-0221113222323012"></a>

<a id="canonical-3221210000230031-1033010021200020-3022203003021233-3020012202001133-3130311021100022-3203310203323012-0101003330001311-1222031223013030"></a>

#### `rule_list.rules.dst_ip_prefix_set.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0321111022030032-2322302230031301-2330333002000223-1223111132313113-0013121312101130-2001323322322232-3010000201300030-2311113000302123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_label_selector` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_label_selector

<a id="canonical-0201201320303023-3100103321201122-1012310023331330-2231213211022200-1110312301211321-1002000012322203-1121022022022300-2000320212022123"></a>

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

<a id="canonical-0201220230001003-0003311330133011-2212332322202223-3302302121102110-0221223011213033-1302333222113022-3100103022321202-2313003131220033"></a>

### Direct properties for `rule_list.rules.dst_label_selector`

<a id="canonical-2312310323122132-1130022111033331-1322302131230201-2021323031003133-3010223222000120-1001330130221100-2210212100300220-0132211111233031"></a>

#### `rule_list.rules.dst_label_selector.expressions` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1302021011333221-1202321123122020-2132230322001012-2021003030013130-3003132133103002-0232330130210222-0130211121032322-1300232013202233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_prefix_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.dst_prefix_list

<a id="canonical-0223030022201003-2200200200212311-1002332201330230-0333233231033333-1022231312203330-3210013000132011-1110132233323131-2212022301232230"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0032023123012220-1031321303213132-1230121323331333-3010123110011210-2131231211232201-1203323213022203-2203233010333110-0111230013330231"></a>

### Direct properties for `rule_list.rules.dst_prefix_list`

<a id="canonical-3011333011112000-0301131132120001-2232312232221323-0233023333111321-3203230023021320-1300111201302213-1122210113213333-0012322011223032"></a>

#### `rule_list.rules.dst_prefix_list.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.http_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.http_list

<a id="canonical-0313322022003312-1011021132002233-2233113003012002-3000231031132201-2330222203212330-2232331110203021-0132310310120212-2113011020012122"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3121102111033322-1122313003321003-1103103211332003-2313002003100231-2330113213211311-0331021231232130-1130032132002132-0133133032223331"></a>

### Direct properties for `rule_list.rules.http_list`

- [http_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1003321321330001-0303131032030213-0232323233130012-2203312203321133-1321010000012031-2111113211232030-2120302312102011-3301121212202203): complete subsection reference.

<a id="canonical-1003321321330001-0303131032030213-0232323233130012-2203312203321133-1321010000012031-2111113211232030-2120302312102011-3301121212202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.http_list.http_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211)
- rule_list.rules.http_list.http_list

<a id="canonical-2332313220202000-3312100330020212-3112003002100213-3103003112132222-3313133221001310-3112333121203012-1230330210012213-1330013121333122"></a>

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

<a id="canonical-1212110001203133-2023222031112010-2111200201032113-2012322222330000-0103033211023030-1231032322332300-3212110010013321-1113130011303211"></a>

### Direct properties for `rule_list.rules.http_list.http_list`

- [any_path](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3333022102011120-3232120200301210-1001022300001130-0322221310302030-1333303212210200-0213321232313103-3223113020301220-2230030322330313): complete subsection reference.

<a id="canonical-3111112333130221-1303131331030032-1201221220120131-0313200330220210-2122220120012211-1012121203203220-2133211012103330-3222132032203203"></a>

<a id="canonical-0311021133333022-0030100223233131-2012112113302232-0321013213331101-3133231123111122-1121320032202311-2002131013300333-0200231220200132"></a>

#### `rule_list.rules.http_list.http_list.exact_value` property

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

<a id="canonical-3202020112112301-0130301032331013-3223031020012102-2221222111000222-3131102031122003-3110003202330003-0232201311303011-2001222202320000"></a>

<a id="canonical-1311132312002033-0221000302223313-3102103122220133-0302220010103200-0121000312210120-0123203100232113-2202121203131221-3221323011313103"></a>

#### `rule_list.rules.http_list.http_list.path_exact_value` property

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

<a id="canonical-0203302233233303-3101333323220301-3122101122123001-2131332212312103-2110111331131123-0230030030203131-1010302022300202-0312302103121033"></a>

<a id="canonical-3103320202020133-3130101201130012-1212111321321222-1020132320210133-3233012101120311-0002231030000321-3321001301100030-0112001312213120"></a>

#### `rule_list.rules.http_list.http_list.path_prefix_value` property

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

<a id="canonical-2101002331233101-2121123312030213-1212133022102300-1003221121320210-1313231012130201-1122023123030130-2330231333003100-2022021211212023"></a>

<a id="canonical-2212321210201002-2231032130303312-1011111002302210-0001122200003213-2203332023020030-0003221121330021-3133301321223201-0320022002332002"></a>

#### `rule_list.rules.http_list.http_list.path_regex_value` property

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

<a id="canonical-2231033013310301-1300010312031130-3103331320013211-0332101033001032-2213312312232303-2301232003320300-0210113011202003-2122311321121102"></a>

<a id="canonical-0010100313130122-0011213020133202-2100222011023230-2323101221230222-3231212010230110-1213201123323133-2302332320120311-1022120232100112"></a>

#### `rule_list.rules.http_list.http_list.regex_value` property

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

<a id="canonical-2222033202030201-0103312203123121-1130330031223122-3231321110323103-3201330223123021-1322032300300130-0000130001102111-1012212032120010"></a>

<a id="canonical-0112133112311310-2333020030312123-2303201023123210-2023310121101020-1202323222202322-1033313020032210-3022303122123031-1203131111313122"></a>

#### `rule_list.rules.http_list.http_list.suffix_value` property

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

<a id="canonical-3333022102011120-3232120200301210-1001022300001130-0322221310302030-1333303212210200-0213321232313103-3223113020301220-2230030322330313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.http_list.http_list.any_path` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [rule_list.rules.http_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211)
- [rule_list.rules.http_list.http_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-1003321321330001-0303131032030213-0232323233130012-2203312203321133-1321010000012031-2111113211232030-2120302312102011-3301121212202203)
- rule_list.rules.http_list.http_list.any_path

<a id="canonical-1000221323210102-2001232000102021-2022332131133200-3300211133331011-3312130111202030-1023022200210103-0011112210031123-3122032033202131"></a>

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

<a id="canonical-2212103130223310-3330130302022032-0203030031113130-1003332132223122-1023111013210012-0030311012030221-2111023131032223-2232213101213000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.ip_prefix_set

<a id="canonical-2221331210221003-2132123310203232-0200013110031023-0002101033031131-1221132113133023-3022000030320031-3023000332220103-2022031021310013"></a>

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

<a id="canonical-2133302302110032-1131133110032100-2202021133233311-0333223332013003-3333230203002130-3220003122111300-3303022220321333-2122113312102013"></a>

### Direct properties for `rule_list.rules.ip_prefix_set`

<a id="canonical-1230000211031033-1020001022123233-1012031230103223-2031203301201211-0020233121223212-3221202220022001-1230002323220221-2003330022033332"></a>

#### `rule_list.rules.ip_prefix_set.name` property

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

<a id="canonical-3231013322121303-3211023310222012-1212300102133111-0333313213331301-0121132232133202-2103321011331112-3123300233201132-1211030321020113"></a>

<a id="canonical-3133032233303201-0013210203233321-2331003133011001-2002210231312230-1210133110330022-2112202112100310-1320001203110213-1013013212020133"></a>

#### `rule_list.rules.ip_prefix_set.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0123201102220123-2130032032012111-3233022021022102-1312131023210020-3003132201332311-3012220202301203-1202330010320311-3320333321030031"></a>

<a id="canonical-2032220022221302-3023121333222332-0031210330100221-1321221221323311-0110030330130000-2030330131031110-3331232103312103-1202302323133320"></a>

#### `rule_list.rules.ip_prefix_set.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0121013023233311-1203120030030033-1332300301332300-0321210320000300-1112203010322331-1133332312311203-2212003212000211-0102321303223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.label_selector` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.label_selector

<a id="canonical-1232121312200123-3023221110001301-1102132213030122-0023311123012201-0231100200223121-2233212231230032-0320102101301123-1033320303203202"></a>

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

<a id="canonical-0332030011220132-2220301233333223-1113033303022122-1111131021333200-2122121331013020-3111122123213332-3211010301031101-3333003200203322"></a>

### Direct properties for `rule_list.rules.label_selector`

<a id="canonical-3330332011201323-1221230021010310-2332112102112030-1123222232210223-2121002033103331-3011232302122230-0013310022221131-1231021130120120"></a>

#### `rule_list.rules.label_selector.expressions` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3223102201230333-0133323300113202-0020231011110213-1103132313011100-2303122321311003-3120230013203320-0203112302310020-0011101020213312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.metadata` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.metadata

<a id="canonical-1210001310130000-3210012311233133-3203210131221230-2122331100231103-1232002322201110-3312010130102220-0001101001020211-0112012302211310"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0131221002102321-1003231000230030-2331302322320311-3121101011120300-2103012031231210-1102111012000211-0200220321311003-1301100030130030"></a>

### Direct properties for `rule_list.rules.metadata`

<a id="canonical-3022123321130201-0032231300012322-3003302202102003-3300010122103121-0311233211231032-2320102111100012-0032210002131133-3003213030131220"></a>

#### `rule_list.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2200322132332303-2323301312032032-1123321320220013-1031301233201230-2020131130311130-2013200130002311-1111312110323011-2312120023120021"></a>

<a id="canonical-0300022301021012-3222200131120112-1331103330102222-1132310331310221-3223310231012033-2200210313123330-0322300020111312-0122201003230311"></a>

#### `rule_list.rules.metadata.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3322110302231022-2313001321321311-0011112313123332-0331303200302032-1300033123310301-3131201202203220-0023301312310322-2323311123323123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.no_http_connect_port` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.no_http_connect_port

<a id="canonical-1100003212003111-2323300223301233-1031230030301310-3212211233102103-3011002303203300-0112210230132020-3010220233131302-0233121302100032"></a>

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

<a id="canonical-3303223031202122-2130021033020311-3000130130333300-0233223101103033-0220002000332013-0100230323032120-1303100210111331-2311331133310002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.port_matcher` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.port_matcher

<a id="canonical-0010111210103020-0201131031233012-3022203303022110-3320030321330122-1300200033301312-0013313002212230-0103213311310311-0302211222100201"></a>

Type: `"single"`. Computed.

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
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

<a id="canonical-0112232023130103-1201230121101032-3100131011331313-1113233211120221-2120232311113021-0030302201223320-3332332230311222-1132201033310231"></a>

### Direct properties for `rule_list.rules.port_matcher`

<a id="canonical-0201020021223012-2301010332213221-0102121212210210-2120330110231032-2003233113322032-0121223011011120-1012200323203231-3000132232033233"></a>

#### `rule_list.rules.port_matcher.invert_matcher` property

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

<a id="canonical-2310013313132102-1102331133322010-0223312113301331-1113002323122333-1222023100020223-1211001012200122-2300023031310131-1030222200333210"></a>

<a id="canonical-3303301220132130-1010000201003311-2112221033023011-1213231022013102-1013211012322133-0110202301212120-2112310211111021-0302203303112301"></a>

#### `rule_list.rules.port_matcher.ports` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3212130112120111-0303230130202111-3022100232223233-3002201320311201-2201332023302332-2211231010010220-3022300313111010-2321121102331231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.prefix_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.prefix_list

<a id="canonical-0011311212302311-1032132100233321-3030201131123332-0303230101110013-2301332123200111-2303313331112220-2231020112312331-1132213222233333"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3133133232023122-0333101020300313-3132100021132002-0310011111123132-1120021102302223-3223303313323112-2113331112301111-1012022332331302"></a>

### Direct properties for `rule_list.rules.prefix_list`

<a id="canonical-0003033122013012-1312321111102100-2111321330300023-0111213310330331-1233012132332232-1300222012221331-0200000311031133-3133120133121102"></a>

#### `rule_list.rules.prefix_list.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-3130031113132122-1313310010332323-3012313301202331-1020032123210331-2033230333200121-0311201313100230-3313331323102333-1023320301301321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.tls_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.tls_list

<a id="canonical-0310303133133202-3123020233103333-0330303332213232-3102210130001301-0203320323021232-1333013233022222-2023200032210202-2023003003301330"></a>

Type: `"single"`. Computed.

DomainListType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2011320012013222-2313333310203131-1313102232100323-1100032333200000-1023032110032030-2003223001333012-1033131302113202-0233010211323213"></a>

### Direct properties for `rule_list.rules.tls_list`

- [tls_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3010320131320220-1022001113232320-3102131330323313-0200102002102232-3133121213313032-1023122302010301-3211303001020120-1312123113010233): complete subsection reference.

<a id="canonical-3010320131320220-1022001113232320-3102131330323313-0200102002102232-3133121213313032-1023122302010301-3211303001020120-1312123113010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.tls_list.tls_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- [rule_list.rules.tls_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3130031113132122-1313310010332323-3012313301202331-1020032123210331-2033230333200121-0311201313100230-3313331323102333-1023320301301321)
- rule_list.rules.tls_list.tls_list

<a id="canonical-2221221222120112-0230021231220123-3301300121011210-0203100102003010-2231102212311222-2220120122210122-3001211332213332-3022002102200321"></a>

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

<a id="canonical-0203301102131023-1302012010213030-0001020131200110-2303332111323120-1001311133321012-3021122010033011-3300200320222001-0001101020113131"></a>

### Direct properties for `rule_list.rules.tls_list.tls_list`

<a id="canonical-1231232302320130-2023322210002011-2320030032332032-1011123111333213-0333310013131311-3210100331302133-3211133103003322-3331331100000321"></a>

#### `rule_list.rules.tls_list.tls_list.exact_value` property

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

<a id="canonical-3131121023120210-0013333111312111-2322322202212121-0022122110003230-1101133001020130-1131012220221333-1000012033303111-0100311333222213"></a>

<a id="canonical-2220221100312100-2122121230220012-2210320102211001-0312011221003113-1101313100322213-2123203011131032-2223002131102310-1331103230113132"></a>

#### `rule_list.rules.tls_list.tls_list.regex_value` property

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

<a id="canonical-0323013202223203-1131011033132200-0032201033002102-2011120320133203-1201121313211301-0000013202131003-0302033120113001-3210300002010233"></a>

<a id="canonical-0023022121300001-2020211302213222-1300213003311201-2100100133011333-2230012103132120-0301130112000023-1002331011201221-1311002220322200"></a>

#### `rule_list.rules.tls_list.tls_list.suffix_value` property

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

<a id="canonical-0233202033222111-2122021223212100-2100002000120303-2133131211221023-2302211312302313-1031130223233300-0112000112200222-3200310020303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.url_category_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md#canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203)
- [Property reference](data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [rule_list](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3210010030102222-1331103130112213-2033020020302222-2112133300213310-1131111132233200-3132121230310201-0100222323123311-0102303130230232)
- [rule_list.rules](data-sources--forward_proxy_policy--reference--group-002.md#canonical-3132203230220121-0130313322131100-3321120011030323-1103231211032031-3232310121212133-3323131133330001-1030230122003102-1021022230323003)
- rule_list.rules.url_category_list

<a id="canonical-3211012133002021-0031221310323002-1213030330010223-2320300332101200-2021130012032312-2120110323122333-3031003001021132-1121323213231210"></a>

Type: `"single"`. Computed.

URL Category List Type. List of URL categories.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1122222303203303-2211202101333112-2010000000300211-1121123130011312-0212033030020122-3312233022131303-1212130101011130-2302311210312301"></a>

### Direct properties for `rule_list.rules.url_category_list`

<a id="canonical-2131221100000212-2130132000023302-2131000020200320-1031112200021103-0002233003031312-0131003300330230-0132211220220100-1231300133212100"></a>

#### `rule_list.rules.url_category_list.url_categories` property

Type: `["list", "string"]`. Computed.

\[Enum:
UNCATEGORIZED|REAL\_ESTATE|COMPUTER\_AND\_INTERNET\_SECURITY|FINANCIAL\_SERVICES|BUSINESS\_AND\_ECONOMY|COMPUTER\_AND\_INTERNET\_INFO|AUCTIONS|SHOPPING|CULT\_AND\_OCCULT|TRAVEL|ABUSED\_DRUGS|ADULT\_AND\_PORNOGRAPHY|HOME\_AND\_GARDEN|MILITARY|SOCIAL\_NETWORKING|DEAD\_SITES|INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS|TRAINING\_AND\_TOOLS|DATING|SEX\_EDUCATION|RELIGION|ENTERTAINMENT\_AND\_ARTS|PERSONAL\_SITES\_AND\_BLOGS|LEGAL|LOCAL\_INFORMATION|STREAMING\_MEDIA|JOB\_SEARCH|GAMBLING|TRANSLATION|REFERENCE\_AND\_RESEARCH|SHAREWARE\_AND\_FREEWARE|PEER\_TO\_PEER|MARIJUANA|HACKING|GAMES|PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY|WEAPONS|PAY\_TO\_SURF|HUNTING\_AND\_FISHING|SOCIETY|EDUCATIONAL\_INSTITUTIONS|ONLINE\_GREETING\_CARDS|SPORTS|SWIMSUITS\_AND\_INTIMATE\_APPAREL|QUESTIONABLE|KIDS|HATE\_AND\_RACISM|PERSONAL\_STORAGE|VIOLENCE|KEYLOGGERS\_AND\_MONITORING|SEARCH\_ENGINES|INTERNET\_PORTALS|WEB\_ADVERTISEMENTS|CHEATING|GROSS|WEB\_BASED\_EMAIL|MALWARE\_SITES|PHISHING\_AND\_OTHER\_FRAUDS|PROXY\_AVOIDANCE\_AND\_ANONYMIZERS|SPYWARE\_AND\_ADWARE|MUSIC|GOVERNMENT|NUDITY|NEWS\_AND\_MEDIA|ILLEGAL|CONTENT\_DELIVERY\_NETWORKS|INTERNET\_COMMUNICATIONS|BOT\_NETS|ABORTION|HEALTH\_AND\_MEDICINE|CONFIRMED\_SPAM\_SOURCES|SPAM\_URLS|UNCONFIRMED\_SPAM\_SOURCES|OPEN\_HTTP\_PROXIES|DYNAMICALLY\_GENERATED\_CONTENT|PARKED\_DOMAINS|ALCOHOL\_AND\_TOBACCO|PRIVATE\_IP\_ADDRESSES|IMAGE\_AND\_VIDEO\_SEARCH|FASHION\_AND\_BEAUTY|RECREATION\_AND\_HOBBIES|MOTOR\_VEHICLES|WEB\_HOSTING\]
URL Categories. List of URL categories to be selected. Possible values are \`UNCATEGORIZED\`,
\`REAL\_ESTATE\`, \`COMPUTER\_AND\_INTERNET\_SECURITY\`, \`FINANCIAL\_SERVICES\`,
\`BUSINESS\_AND\_ECONOMY\`, \`COMPUTER\_AND\_INTERNET\_INFO\`, \`AUCTIONS\`, \`SHOPPING\`,
\`CULT\_AND\_OCCULT\`, \`TRAVEL\`, \`ABUSED\_DRUGS\`, \`ADULT\_AND\_PORNOGRAPHY\`,
\`HOME\_AND\_GARDEN\`, \`MILITARY\`, \`SOCIAL\_NETWORKING\`, \`DEAD\_SITES\`,
\`INDIVIDUAL\_STOCK\_ADVICE\_AND\_TOOLS\`, \`TRAINING\_AND\_TOOLS\`, \`DATING\`, \`SEX\_EDUCATION\`,
\`RELIGION\`, \`ENTERTAINMENT\_AND\_ARTS\`, \`PERSONAL\_SITES\_AND\_BLOGS\`, \`LEGAL\`,
\`LOCAL\_INFORMATION\`, \`STREAMING\_MEDIA\`, \`JOB\_SEARCH\`, \`GAMBLING\`, \`TRANSLATION\`,
\`REFERENCE\_AND\_RESEARCH\`, \`SHAREWARE\_AND\_FREEWARE\`, \`PEER\_TO\_PEER\`, \`MARIJUANA\`,
\`HACKING\`, \`GAMES\`, \`PHILOSOPHY\_AND\_POLITICAL\_ADVOCACY\`, \`WEAPONS\`, \`PAY\_TO\_SURF\`,
\`HUNTING\_AND\_FISHING\`, \`SOCIETY\`, \`EDUCATIONAL\_INSTITUTIONS\`, \`ONLINE\_GREETING\_CARDS\`,
\`SPORTS\`, \`SWIMSUITS\_AND\_INTIMATE\_APPAREL\`, \`QUESTIONABLE\`, \`KIDS\`,
\`HATE\_AND\_RACISM\`, \`PERSONAL\_STORAGE\`, \`VIOLENCE\`, \`KEYLOGGERS\_AND\_MONITORING\`,
\`SEARCH\_ENGINES\`, \`INTERNET\_PORTALS\`, \`WEB\_ADVERTISEMENTS\`, \`CHEATING\`, \`GROSS\`,
\`WEB\_BASED\_EMAIL\`, \`MALWARE\_SITES\`, \`PHISHING\_AND\_OTHER\_FRAUDS\`,
\`PROXY\_AVOIDANCE\_AND\_ANONYMIZERS\`, \`SPYWARE\_AND\_ADWARE\`, \`MUSIC\`, \`GOVERNMENT\`,
\`NUDITY\`, \`NEWS\_AND\_MEDIA\`, \`ILLEGAL\`, \`CONTENT\_DELIVERY\_NETWORKS\`,
\`INTERNET\_COMMUNICATIONS\`, \`BOT\_NETS\`, \`ABORTION\`, \`HEALTH\_AND\_MEDICINE\`,
\`CONFIRMED\_SPAM\_SOURCES\`, \`SPAM\_URLS\`, \`UNCONFIRMED\_SPAM\_SOURCES\`,
\`OPEN\_HTTP\_PROXIES\`, \`DYNAMICALLY\_GENERATED\_CONTENT\`, \`PARKED\_DOMAINS\`,
\`ALCOHOL\_AND\_TOBACCO\`, \`PRIVATE\_IP\_ADDRESSES\`, \`IMAGE\_AND\_VIDEO\_SEARCH\`,
\`FASHION\_AND\_BEAUTY\`, \`RECREATION\_AND\_HOBBIES\`, \`MOTOR\_VEHICLES\`, \`WEB\_HOSTING\`.
Defaults to \`UNCATEGORIZED\`.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
