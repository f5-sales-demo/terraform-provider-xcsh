---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1021123312221201-1312222301111302-0233031121333333-1112332031331023-2310013123003021-2321211002000112-1113110312032301-0122123101303130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-2230100233203213-1120300010112000-0033032222233023-1221003002002133-3333322302313312-3233023033001130-0331331300123133-1200013311023132"></a>

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

<a id="canonical-0033323113001332-3130211200030003-2313121202223102-3032311011020331-2032123031112202-0231010012213031-3203332330311022-3330003220023122"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_list`

<a id="canonical-2101323021223102-2303012012211000-2103231130211130-2102022110010130-1003022201233102-2132232032003010-2012233110323002-0330011313201310"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_list.as_numbers` property

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

<a id="canonical-0001122010010121-1201010123310202-2120312113113333-0023210203203012-1233023110101200-1222311231232122-0001033100311021-1031221311311303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-2101013120111112-3122211310311222-1103300330113333-1233323210220011-2323210023200112-3132030132033123-2023013203122013-1113113100212033"></a>

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

<a id="canonical-3101232110233100-3310212233331011-2000033000003120-1220223110321213-3213021221123222-2300100032002132-0323231122223102-2320121102213122"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_matcher`

- [asn_sets](data-sources--http_loadbalancer--reference--group-010.md#canonical-2133002000010013-2320002221020120-3111202331020021-0203113321123023-3230112033110322-0000001100020200-3230102030113203-1031233020323320): complete subsection reference.

<a id="canonical-2133002000010013-2320002221020120-3111202331020021-0203113321123023-3230112033110322-0000001100020200-3230102030113203-1031233020323320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001122010010121-1201010123310202-2120312113113333-0023210203203012-1233023110101200-1222311231232122-0001033100311021-1031221311311303)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-0203220112232122-0200221333212012-3121330310121232-0301222001311231-3303233112030303-1113201213020001-0003230032033331-2221313123212203"></a>

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

<a id="canonical-1322110230323131-1022322212221313-2203100311001021-0010212030321130-1332320133112013-3332313003111123-3330032102112011-2113301203320333"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-3013023003220313-0031031301023331-2220322221330021-3320003003130330-3212203200012210-0210331120230322-1203110212332121-2233123312300331"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.kind` property

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

<a id="canonical-0232112210211211-2230121221312012-2322133232222103-0222311133221312-1312322012103031-3122313130310133-2000110020222310-3130010020130112"></a>

<a id="canonical-0002032003201321-0111030022011221-2220130302130212-1100130332032131-2212201100202201-0333013202133012-2223320303131000-1231013220302103"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.name` property

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

<a id="canonical-2230131302021200-2100331220223312-3213033203222210-0113312123330003-1103131101022230-1121022013323230-1030333322201130-3202021311021133"></a>

<a id="canonical-1011002010323312-3331113003122011-0320331300121123-2313120301001331-2001001221313130-1312330322003101-0020101102300132-3311232021131212"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.namespace` property

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

<a id="canonical-2110113233303100-3130131213203230-0312223330133212-1022323201322210-1233210032032302-1331201331222302-0033122321333012-1303113200113301"></a>

<a id="canonical-2100230230103311-3210032121000200-0200332312322303-1232300303131020-1321223021330323-1313201311312112-2102303110022133-2101233030212033"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.tenant` property

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

<a id="canonical-2000201001110221-0022311333203321-0132230303110101-2211303112331013-3210301223301201-3300312231121333-1031100212100332-2031002212120012"></a>

<a id="canonical-3323331311331231-2203330200131113-0011101301112021-2200303321032032-3221202323103230-2333030130223311-1121203313301320-3203231102033230"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.uid` property

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

<a id="canonical-3111210223132323-0121012002312020-0220110001201123-3002220100131232-0100103213330022-1330010130033300-2313010211122202-1311133120322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-2221120130301002-1102203103110232-3100302301031032-1211000110222332-0232003102033031-0011322113331322-2300333003101223-0310233002021121"></a>

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

<a id="canonical-3312110121301322-3223011231121112-1212200122330023-1230313233333121-0120101131100203-2130330310212002-0032013300101002-0302211311213302"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.client_selector`

<a id="canonical-0102032011031333-0231003122211021-2003132312011311-3103130311112121-2313032021121213-0000321303011330-0100230030112230-0320303330000200"></a>

#### `api_rate_limit.server_url_rules.client_matcher.client_selector.expressions` property

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

<a id="canonical-3122322002000200-3120320003232323-0302010210201232-1233232103213311-3111211031112111-0220232001101101-2331023211232022-0213001120023010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-2003101222232223-0123222220222031-2332330331013332-0210313032132120-0003232200112101-0120101212002212-1001231310220322-0030303012012102"></a>

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

<a id="canonical-0133000112203320-2223221003111332-1123303023002233-2120022312301310-2123313210120022-1020112102010233-2220132321302310-1233301100133101"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher`

<a id="canonical-3233320113201312-0110200010133331-0313111323311222-2133030112001220-0221203020231233-0211222111101113-0202110110130133-2223003113211222"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-010.md#canonical-0313223320302213-2220023320001132-1232321221322311-1123233013103132-0302132301033333-2133232021132230-2303133000222011-1233002332010222): complete subsection reference.

<a id="canonical-0313223320302213-2220023320001132-1232321221322311-1123233013103132-0302132301033333-2133232021132230-2303133000222011-1233002332010222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-3122322002000200-3120320003232323-0302010210201232-1233232103213311-3111211031112111-0220232001101101-2331023211232022-0213001120023010)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-1233123132133311-1332102323120121-1212012223112221-3310101203000001-0311130313022002-0012332002330020-2021312110230230-0131113333132111"></a>

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

<a id="canonical-3113231301211303-2103321033203320-1021231302211211-3212001122012331-1300133222031312-0122012303302121-3203110202333222-0212032131302122"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-3102310110130110-0231232013220310-1312232213131023-1120010212100313-1121031132333010-0321103201331233-3120133113321131-0201313012020013"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.kind` property

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

<a id="canonical-3023111200123131-0011232233111212-2000133102003033-1301320020220313-3310212031332011-3232200220131201-1213323322130033-2302010132232232"></a>

<a id="canonical-3311333022013032-0013102322103100-1033110023213300-2223210021111233-2313202330310133-2030210322033310-2231203320321232-0020222032233221"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.name` property

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

<a id="canonical-3210002030110213-2123200203302201-0133233122330031-3033233313022301-3231133133332133-0032303210330222-2332312330131313-2103210133130110"></a>

<a id="canonical-2101320332230002-1331132123330000-1220103130021013-2131231022111121-2102032130301220-2323220011100302-3133333132023221-3211133320212232"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

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

<a id="canonical-0012021221231223-0013332123221130-0213101103022113-0232113030032133-3322120121103003-0220131123133331-0032100113312311-3310302013003120"></a>

<a id="canonical-3313113022300322-3102121111213333-1011321101002313-3000312111222303-2001001002031310-0031102223322020-1322321102311112-3113012233031310"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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

<a id="canonical-1220031100021111-0200211023230332-1302210301000011-0133031022331332-3223110013021331-1222113030103312-1300122202221102-3202122033333221"></a>

<a id="canonical-3221311003113113-1132003331030223-1032023302110010-1211230103002220-1312213033311100-1330333000032211-2223303111201020-1330313131121032"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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

<a id="canonical-0311222103112303-2230331030010021-0313331233320111-3132301213210203-1021012331121103-1101310110000222-3112130120323031-2300123230221020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-0321231203010132-1111222321113230-1112301030210202-3330300013022000-1213200213013131-0130232102231013-2021303222330013-1203232331211101"></a>

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

<a id="canonical-0301002002332200-1020323201221310-1233320022230031-1102132113211133-2100013210230111-2233000211103233-3222031202131322-2310312300312321"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list`

<a id="canonical-3301003331030220-1012223130002322-3120222311120002-0212123212010322-0032303023330311-2111232030000203-1231200320032233-3323021023321003"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-3223100212012332-3223000321131022-3010203331300211-0121122212132232-1220030203223031-1202331311211333-3300130022022301-2021212013030131"></a>

<a id="canonical-2032201132213110-1222332012322233-2110021210133001-3321312021331200-2133100132313000-3013020000111032-0211012020303303-0202210121103320"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.ip_prefixes` property

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

<a id="canonical-3120113231122011-1303331120313130-0331123002200313-1103111321312013-3130331030123131-3022203021213101-0031030311322232-1321332310031012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-1111113122311232-0231320301111232-3212012222322311-1020212130032110-0230323201322222-3103212002330200-1310032031003321-0001112102301302"></a>

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

<a id="canonical-2230131320221010-2203331321012312-1213321103310023-3111203232300332-2030331101133303-0310100103003011-2313330011011122-0223001011123010"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list`

<a id="canonical-1130322220002122-2303311222021030-3013101220013012-3203102301232111-0232103011233210-3120131011331130-1013022023131111-3131323230202100"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

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

<a id="canonical-0210212023001300-2110302130121011-3111122332320113-1020310333033231-2233103310202212-0200131132000300-0101221231331202-2112211000022313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-3131032032221120-2201133003230122-1210200002321310-2100201003021313-1100001120203313-2333123013203020-3110333131223123-3121333312132303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2310213212311021-3310203011013033-3103213203202331-0100031113030023-3313110310320301-3102012213333000-0200000030031003-0012313333330021"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-2003223111000210-3331213300003311-2303120102300232-1001011022232212-3221113203310223-0021123100011302-0131122200032320-1020123033000201"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

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

<a id="canonical-1301021101230132-1332323013001201-2202201321101202-1100012311210232-0033333312010203-2210211030013033-3213220110030331-2032132301010110"></a>

<a id="canonical-3213213313221103-1132033020030223-3111213132103211-0131202113102301-0131311322101203-0000202031223123-2201230100202222-1111101203301120"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2230312021020331-3123110130010002-0331031321122023-2101003030000003-3202333013303000-2300020023111020-3210132100113001-1011333233022022"></a>

<a id="canonical-1022312100100213-0331301313032130-0302123230031020-1330000023212000-1300223013202130-3003020121113211-3021230110010002-2111032132330100"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0322033310030223-3022000203103231-3203020110202213-2221301013202210-3110312310203011-0023332220310100-1311220130202000-3201021231210030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-3221212222213002-1220022132221200-0231100220300323-3213023010213101-0000201302311203-0030323010002331-0331122233212220-0223023132132123"></a>

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

<a id="canonical-1010100333130102-2220302211130033-0220012213011022-1001201022203010-2103230213200200-3120122223231223-0202033223131031-1302122120301010"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter`

- [ref_user_id](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133002201002000-1311331001102102-0221132210022203-3110111113103223-0012003120201122-2201132233010231-1313232212320033-0311330322012301): complete subsection reference.

<a id="canonical-0220232332020210-2323230123332123-3331321310110130-0312001001231012-3310000300000102-2033010311230302-0233101321300210-1001321020200133"></a>

<a id="canonical-1301220302301123-2132211333111303-1022230131123313-2332111130013332-2111222321233202-2322022103010110-3020301113111230-1002303333120033"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.threshold` property

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

<a id="canonical-1031103020021220-0000100131121200-3022113311111112-3012200213010032-1010332033123333-1133300103002122-0333212022310210-1000321202230210"></a>

<a id="canonical-0301221333211220-1010201130320203-1001003333333321-2011320313303122-0303033221020010-2210303001301030-0301021233211302-3233303132231021"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.unit` property

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

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

- [use_http_lb_user_id](data-sources--http_loadbalancer--reference--group-010.md#canonical-2100233000123002-1312013200022320-3000111233302003-3233011130100121-0203023003231223-1023230323303033-3101100210312332-1133330332320003): complete subsection reference.

<a id="canonical-1133002201002000-1311331001102102-0221132210022203-3110111113103223-0012003120201122-2201132233010231-1313232212320033-0311330322012301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-010.md#canonical-0322033310030223-3022000203103231-3203020110202213-2221301013202210-3110312310203011-0023332220310100-1311220130202000-3201021231210030)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-2001103100010102-3221203311100013-2222131311311022-3003022232003021-0222010023200020-0200321010033131-2132001222210223-1212120021230132"></a>

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

<a id="canonical-2131333221211001-3200202103133231-2333303231321002-1312221223121302-1320213200130222-3132013330132030-3133112222123022-1313302023030130"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-3230310200103030-3300013200231001-0220330031023212-0330003321210133-0131021213000333-1031120111102022-1223332032011313-2233112231303110"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.name` property

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

<a id="canonical-3111000010331133-1103003113203133-3013232020302020-0003202221333212-0222231212121323-1312310232023020-2113331311312320-2233313221333000"></a>

<a id="canonical-1300001023121301-2300223313230302-2223000302332330-3222233312020321-3031023010123210-3223130121322020-0131033210023032-0121033023201311"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.namespace` property

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

<a id="canonical-1232200010001133-2112330323102232-0302312101133300-2131010321302300-0222313123110300-3210230211033331-0212230312121231-2021223200030220"></a>

<a id="canonical-3100101210122113-2010300230113330-0302203312211303-3221332202110003-1202322222202012-0131233120001101-2030120300021100-1222302311233332"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.tenant` property

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

<a id="canonical-2100233000123002-1312013200022320-3000111233302003-3233011130100121-0203023003231223-1023230323303033-3101100210312332-1133330332320003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-010.md#canonical-0322033310030223-3022000203103231-3203020110202213-2221301013202210-3110312310203011-0023332220310100-1311220130202000-3201021231210030)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-2102112301100122-0303021000312032-3111233320203130-1031331002211200-2011101221031213-0002032030210130-3103302120233113-1130030011020230"></a>

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

<a id="canonical-0230200231320022-0022211103313020-3313102213001013-2131023001201221-2123213222111222-2111231003133322-2133231102123132-3302111101201122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-1013121200000122-0313311233130023-3213232131320100-0012211113302122-1302111320112300-1332021003002030-2101012122321332-0233031131320003"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Additional upstream details:

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

<a id="canonical-3010000211003110-1010122223012130-2002100310121201-2100000010313122-1221223102122211-3131010011032132-1201330113320310-1212202023100123"></a>

### Direct properties for `api_rate_limit.server_url_rules.ref_rate_limiter`

<a id="canonical-0111132332313022-1103313111230130-2120220332221012-1023112330330303-0023123012000223-0100310013110222-2102303332233311-1101220032010300"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.name` property

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

<a id="canonical-0212201131222203-2212203033320012-1203122132212121-3322230323200313-3012133221223030-3131201222222300-2133023202001302-1010311330013003"></a>

<a id="canonical-2222130202012132-0000011020201023-3220233033122032-2003133213002331-1213313330303132-1220312201030012-2321311322213033-0331323232202202"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.namespace` property

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

<a id="canonical-2212203000002221-0203031312222132-1123101232221132-2001031312003300-0121332201323223-0202110021100332-1201300020000323-0212313212322212"></a>

<a id="canonical-0121101011230201-0101222110001013-1000101102030120-2123330323101233-2302131021002221-2111003223113013-3120222313201132-1332023320310111"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.tenant` property

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

<a id="canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-1202233032331312-2310122121003203-3031122332130033-1102300202323313-1233201311022131-0222320102300333-1303012212233031-2103233301322321"></a>

Type: `"single"`. Computed.

Configuration parameter for request matcher.

Additional upstream details:

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

<a id="canonical-1023102013111200-1102120013201103-3200031010200033-1101002232203103-0013122032233301-3322233101331131-1320223220222313-1121123100320022"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher`

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-010.md#canonical-2302313313211200-2233133103311111-2001102203220223-1030103322333310-3122220033231220-2133220210332203-1302221003322011-1222232333103220): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202012122323002-2221011300323011-3002010321331011-0011132130110012-1002100123330121-1333131010222131-2110201113010212-2131031333112121): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-010.md#canonical-1003021201210032-1221023103130112-3122023300313011-3021330322231301-3310310002123332-1033032201020133-3231303220230110-3122221010023030): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-010.md#canonical-2200312211113322-0112121111320122-2320320321112330-3123202310120301-0202103223313212-2121113323301100-2031302101213301-3311210320012002): complete subsection reference.

<a id="canonical-2302313313211200-2233133103311111-2001102203220223-1030103322333310-3122220033231220-2133220210332203-1302221003322011-1222232333103220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-3133312003112133-1102100203132102-0131222320112030-3323323113303122-1131300330201011-3332321021301101-0101013112130120-0121103311100033"></a>

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

<a id="canonical-0022212011001323-0010021220031330-0110233330033221-1330210000211113-0101002221203023-0020131113301223-1132311132213222-1312122203101212"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-010.md#canonical-1220122103001202-0130013222320223-1123011323010222-1103012002211001-0221312130333010-0033222320323033-3230010302333202-2220233123101310): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-010.md#canonical-1220102002001333-1202301201210123-1320013102111102-1200033022323320-3322120131221112-2011002121333320-1200122133003011-2323221013222300): complete subsection reference.

<a id="canonical-3331021221330220-1121101103013333-3221131022113302-1231123112212021-1330313122103222-0321002000223222-3030130330132111-1102101230220102"></a>

<a id="canonical-1211211332230210-3333201323331032-1113303300110210-1123133221120302-3220230220001202-0010212300113121-3111203030031132-0223013302231012"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](data-sources--http_loadbalancer--reference--group-010.md#canonical-0103000032223101-1312123211313033-2321121230110301-0300112312000013-0330221030002023-0031200103111203-2121133212002011-0201230032323322): complete subsection reference.

<a id="canonical-2300011113023121-0111102203200322-1320023301333223-3302110020130310-1220122120210120-3223023302013212-1312100101330233-0121310202232013"></a>

<a id="canonical-0123233102300222-0112102133101121-1311010231130020-3312332121001230-1320221120312201-3210323312130130-1012301030031122-1133321122011030"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.name` property

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

<a id="canonical-1220122103001202-0130013222320223-1123011323010222-1103012002211001-0221312130333010-0033222320323033-3230010302333202-2220233123101310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-010.md#canonical-2302313313211200-2233133103311111-2001102203220223-1030103322333310-3122220033231220-2133220210332203-1302221003322011-1222232333103220)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2131331302012012-2212320203333313-0231103313320131-1102322023213230-3203132200300102-0020032233201030-1320232223003023-0323231120330013"></a>

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

<a id="canonical-1220102002001333-1202301201210123-1320013102111102-1200033022323320-3322120131221112-2011002121333320-1200122133003011-2323221013222300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-010.md#canonical-2302313313211200-2233133103311111-2001102203220223-1030103322333310-3122220033231220-2133220210332203-1302221003322011-1222232333103220)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-0102132122021232-2033033102210122-1113030131230121-2103310010001300-3301123303012320-2102222113110313-0300331101101320-3121131201221203"></a>

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

<a id="canonical-0103000032223101-1312123211313033-2321121230110301-0300112312000013-0330221030002023-0031200103111203-2121133212002011-0201230032323322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-010.md#canonical-2302313313211200-2233133103311111-2001102203220223-1030103322333310-3122220033231220-2133220210332203-1302221003322011-1222232333103220)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-2010210113212230-2323321212331300-3012003012313013-3130220123133030-1000030120122111-3300120113133323-3323330230321322-1330001200203313"></a>

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

<a id="canonical-3022103033203321-0323232332113313-0020312321112013-2313131012230021-1212033103211200-1101222311311322-3012030103330111-1222033302130213"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item`

<a id="canonical-2222202332202213-2303102301233220-0300302012000300-3113232103010011-1300222001320213-1100031332122101-2230313231100231-2210313201112223"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.exact_values` property

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

<a id="canonical-1303302031201000-3303303102323000-1333003333301202-3111331230320133-2102311230000232-0113010213222211-2101303131232132-2232002320123303"></a>

<a id="canonical-3101233200132023-0212021103230212-2023332330213333-2321013220113302-2301023310300101-1102033210103123-1310202000221002-3011130012023112"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.regex_values` property

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

<a id="canonical-1000132203313113-0200303100212220-0023322221311103-0223131001333122-1322131323301011-0220210330200302-2102212331221330-1303113122311133"></a>

<a id="canonical-0011233101223121-2231303031002210-0101220003221222-1233113312300121-0031333131322213-0220131220001331-2201023113032303-0223222111201210"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.transformers` property

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

<a id="canonical-2202012122323002-2221011300323011-3002010321331011-0011132130110012-1002100123330121-1333131010222131-2110201113010212-2131031333112121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-1110210033100120-3300232221202331-3023301220302210-0323331233133313-3230203001023023-1312110323013211-3231232123131323-0000312131230333"></a>

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

<a id="canonical-3231112011230200-1100131300320020-3213313122033112-0133312021001110-1000012033221132-0330021213312232-3222112132133202-0022300310003111"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-010.md#canonical-2100203030122023-3001312021022202-3113132303000100-3311111222321303-0330021312001203-2111201101312020-3001200310200222-1110233011303001): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-010.md#canonical-2213013203113110-0213002132222131-3030313223333331-3301232332101322-1320131222111120-3222100313033000-0211323111311101-1310120030331301): complete subsection reference.

<a id="canonical-0312101311213102-1313303322333001-0101220013232022-0212100222233210-1120212213002202-1333300330021133-0033201011022010-1013122211032003"></a>

<a id="canonical-1012000322321121-1220203333010020-3102123332211222-3123101033220302-1000002223003101-2122023001203012-1030113210023102-0013213021212020"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.invert_matcher` property

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

- [item](data-sources--http_loadbalancer--reference--group-010.md#canonical-0331223210031012-0300211013322310-3230110231020002-3200333202301013-0001032033113003-2121200302220101-1031003223030100-3211122212220120): complete subsection reference.

<a id="canonical-3101122200103213-3013010333311030-0113202011110103-0301321001022013-1030301202122011-2200320202233110-1101033230230102-1030323003003110"></a>

<a id="canonical-1313003023111021-2232321023113332-3322022221312201-0201200203200323-0232130320000130-1313330103322313-2212021130231220-0033033321310113"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.name` property

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

<a id="canonical-2100203030122023-3001312021022202-3113132303000100-3311111222321303-0330021312001203-2111201101312020-3001200310200222-1110233011303001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202012122323002-2221011300323011-3002010321331011-0011132130110012-1002100123330121-1333131010222131-2110201113010212-2131031333112121)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-2020333130300130-0013100301122231-1311322132203310-3132323012131221-2120012232120222-3232311330211233-0222122023301101-1321113012110212"></a>

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

<a id="canonical-2213013203113110-0213002132222131-3030313223333331-3301232332101322-1320131222111120-3222100313033000-0211323111311101-1310120030331301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202012122323002-2221011300323011-3002010321331011-0011132130110012-1002100123330121-1333131010222131-2110201113010212-2131031333112121)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-2320122132011233-2100202011232232-3011313123233230-3220012120333030-2123231102132113-3302103321312132-3022202333212133-3210121032323032"></a>

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

<a id="canonical-0331223210031012-0300211013322310-3230110231020002-3200333202301013-0001032033113003-2121200302220101-1031003223030100-3211122212220120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202012122323002-2221011300323011-3002010321331011-0011132130110012-1002100123330121-1333131010222131-2110201113010212-2131031333112121)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-3003312333302023-0321110232230033-0122222130000333-1223003211201013-0222120220020220-2112133203231300-2103233311203021-1110122303112102"></a>

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

<a id="canonical-3310332132123320-2312013220201023-2133021200121233-0132201230001230-1223210310123021-0022200110130332-2333223233201232-2233303002231233"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers.item`

<a id="canonical-2202133010100110-1131122102320331-1022301212221110-3300213030031203-0102221330323022-0110203300100221-2032131333300033-1133332101101011"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.exact_values` property

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

<a id="canonical-0013033030332023-1033030032320313-1321201013100223-0033222032121201-1013231132231232-2102032033321230-3121031032332301-2030232111320123"></a>

<a id="canonical-2233112222130131-0012100030132320-3120323101303310-0312020101222223-1230311211102020-0203131023233220-2230101231011023-2310120031031003"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.regex_values` property

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

<a id="canonical-0033130310113303-1200222031001220-0322100013111010-0033201033223023-2133320212012212-2103222312032212-1032122332232112-0100122132200302"></a>

<a id="canonical-1131021233023031-0201102000330032-0311200132023300-1313010223020332-3011322222031311-3031203301331311-0020032201030010-1200122131320332"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.transformers` property

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

<a id="canonical-1003021201210032-1221023103130112-3122023300313011-3021330322231301-3310310002123332-1033032201020133-3231303220230110-3122221010023030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-1202200100103010-0220113322020010-2102230322321123-0111303201303120-1311233233322323-1211222023100223-2133120110013322-2300202211112320"></a>

Type: `"list"`. Computed.

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

<a id="canonical-2010102211323011-1323320122123101-2210102112323020-2013320123231100-3213321110121133-1113301211011300-2312223233332010-0011020210030100"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims`

- [check_not_present](data-sources--http_loadbalancer--reference--group-010.md#canonical-2300120101331223-3122230111223202-1230033302301000-0022310102311222-1122202023121221-3213023010102130-3232213003202333-1201110202212312): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-010.md#canonical-0311321122012003-1121010131230122-1011222130220310-3333222223011333-3130223311221111-2001213013022201-2101112111301002-0132030301030010): complete subsection reference.

<a id="canonical-3022313100031212-2101021313130022-3122133123010110-2311112113032202-3310320101131113-3301221112332133-2232003032012223-1000123113303301"></a>

<a id="canonical-3311333013103201-1313022202331103-2123230033211002-3210330032212210-1103020302213030-1330122223222023-1131301103210332-0011231231001113"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.invert_matcher` property

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

- [item](data-sources--http_loadbalancer--reference--group-010.md#canonical-0220132213032223-0011132031000321-1113221101121022-1103320223300232-3023331022103003-2123300312302231-0232111022010113-2002232001312300): complete subsection reference.

<a id="canonical-0210012212332112-1000311032200233-3012322030012202-1222000123303313-0110202212112021-1123223002103030-2333110103323020-0233003030120200"></a>

<a id="canonical-3103111233322202-1230231031303303-1323332002330221-3033301021101120-0000330331113013-0313322320301202-2333201212100023-1330313031233203"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.name` property

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

<a id="canonical-2300120101331223-3122230111223202-1230033302301000-0022310102311222-1122202023121221-3213023010102130-3232213003202333-1201110202212312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-010.md#canonical-1003021201210032-1221023103130112-3122023300313011-3021330322231301-3310310002123332-1033032201020133-3231303220230110-3122221010023030)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3222311213101313-1020200231013112-1310220113221010-0301003222332112-3322031033122300-3330103100002201-1201320032132310-3023231210032331"></a>

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

<a id="canonical-0311321122012003-1121010131230122-1011222130220310-3333222223011333-3130223311221111-2001213013022201-2101112111301002-0132030301030010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-010.md#canonical-1003021201210032-1221023103130112-3122023300313011-3021330322231301-3310310002123332-1033032201020133-3231303220230110-3122221010023030)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2120030300222312-3323212013103323-3111133331212200-1210013211330302-2100022310213321-2310200223210220-3210303303230302-0323033221203113"></a>

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

<a id="canonical-0220132213032223-0011132031000321-1113221101121022-1103320223300232-3023331022103003-2123300312302231-0232111022010113-2002232001312300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-010.md#canonical-1003021201210032-1221023103130112-3122023300313011-3021330322231301-3310310002123332-1033032201020133-3231303220230110-3122221010023030)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-3120011123330123-2132222232101113-0103202132102322-3013003333101010-3112302213013202-2233221211120320-2233213123133013-0102232233303211"></a>

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

<a id="canonical-2031202203133032-3313330300310231-3332333302221212-0022032001223130-2332132123132030-0333113213311211-0201201230223001-2110032110202312"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item`

<a id="canonical-1310122303102112-2120113222213232-2210320310010220-2010222330310001-3010131233330003-1223311112133023-0301212101131202-2323203103221130"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.exact_values` property

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

<a id="canonical-0332033300023333-2321310112013310-3213313223112021-2023121221013330-3000103230010102-2301111333320301-1011021213000213-3221122203200203"></a>

<a id="canonical-1300313003221101-2321333001000010-2313232000123100-3113321112201311-2223201000212013-3131102020102333-1323201013132310-1120023000201321"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.regex_values` property

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

<a id="canonical-1032203232222333-1300131310023310-3031003202332323-1210102133022110-3131122133321221-0220200120301111-3001101330223320-0121220111212303"></a>

<a id="canonical-2213031211323110-1000233333332333-0113103120133002-3223133112010102-2302132130130110-3233231320220023-3333220303312312-2111133231212033"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.transformers` property

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

<a id="canonical-2200312211113322-0112121111320122-2320320321112330-3123202310120301-0202103223313212-2121113323301100-2031302101213301-3311210320012002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-3010110132230131-2333310310111032-1031101100011022-1120020103330131-0223301230313020-2122022130103023-0010131110132031-0103020010320213"></a>

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

<a id="canonical-2230211001301201-1120132312123321-2212131110031330-2131323110301001-1311030303221013-3022101232301223-2102201111132230-3230213013322300"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params`

- [check_not_present](data-sources--http_loadbalancer--reference--group-010.md#canonical-3101000021230221-3101122113133320-0030012222032131-3030302320120011-3100111020002222-3033230011022323-2333303221330110-0033303103012000): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-010.md#canonical-1321111020333113-0221300303203131-2033210232121303-2133021310330303-0210203132312123-2212021332131133-3113130223131031-0230002221131222): complete subsection reference.

<a id="canonical-2300311321030100-1122102322031223-1010021303331121-0030031103021011-1110030000313012-3011203113301322-2033223321020122-1001301023203023"></a>

<a id="canonical-3323211202032203-0000101310330032-2030011310002212-0001300231323321-2033203033232113-3220210123213231-2220300133022231-3013120202213130"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.invert_matcher` property

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

- [item](data-sources--http_loadbalancer--reference--group-010.md#canonical-0320022211020230-1132121121232013-2022101300300202-2333103033313020-0032001330100020-3121101130003221-0213110123102030-2102022331000000): complete subsection reference.

<a id="canonical-3033003031030230-1313111031012321-1013212033110012-0322313233100212-3212223313330100-2320030033213010-2113123302122310-3312331321321030"></a>

<a id="canonical-3232113213000211-3102112211023030-2312002321310212-0233220120211232-3033122003222103-3113011113030120-3200120230010303-1010111013333200"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.key` property

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

<a id="canonical-3101000021230221-3101122113133320-0030012222032131-3030302320120011-3100111020002222-3033230011022323-2333303221330110-0033303103012000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-010.md#canonical-2200312211113322-0112121111320122-2320320321112330-3123202310120301-0202103223313212-2121113323301100-2031302101213301-3311210320012002)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-3220311002121023-2220013132202003-0323210233110232-1022233200013020-2230003300221121-1312132000312112-2031203113330101-3101011023302200"></a>

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

<a id="canonical-1321111020333113-0221300303203131-2033210232121303-2133021310330303-0210203132312123-2212021332131133-3113130223131031-0230002221131222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-010.md#canonical-2200312211113322-0112121111320122-2320320321112330-3123202310120301-0202103223313212-2121113323301100-2031302101213301-3311210320012002)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-1330132211101221-1312130323321202-2231312033333311-0202223301232302-1202202322301232-0110132032221023-3320320032332230-2110000332021102"></a>

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

<a id="canonical-0320022211020230-1132121121232013-2022101300300202-2333103033313020-0032001330100020-3121101130003221-0213110123102030-2102022331000000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313)
- [api_rate_limit.server_url_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-010.md#canonical-2200312211113322-0112121111320122-2320320321112330-3123202310120301-0202103223313212-2121113323301100-2031302101213301-3311210320012002)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-1132102113331102-2021100100202131-2133002131212312-3032102212332330-3220120312320310-2013100220030102-3113300221321033-0102213210102102"></a>

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

<a id="canonical-3202311211003130-2310332202111100-1230123332001223-1100121332202212-2020231233200132-2222133210121131-0110211310113103-0203031211210120"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params.item`

<a id="canonical-1221212310333003-0000101003211223-0302131110013011-0233032331133332-1230223031310202-3010112220110212-1030001302111213-0022300121330100"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.exact_values` property

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

<a id="canonical-0333013223232300-3020223110302210-0333003103332220-3120133332101123-2032130310122030-1130031203012310-3233301321201111-1031230201300212"></a>

<a id="canonical-2012011133020100-3031233112302313-3310101233201111-2211322312121011-3202331333332311-1033313321232322-0133113301303311-1221133103100333"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.regex_values` property

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

<a id="canonical-3021122110120103-2130312202201113-3313133003322311-0011200332122132-3301113330303012-2310002223323213-1103000331133302-0221200121020232"></a>

<a id="canonical-0322031213011020-1122121102032300-1002322220233002-0131332003313331-0003002201013120-3120001202033101-1233223203213112-1303320111100001"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.transformers` property

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

<a id="canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- api_specification

<a id="canonical-0233032022033311-3000012333333220-3023221323213013-0000232201012301-3323333020013020-2320312132000122-3131220121310300-1220331333210121"></a>

Type: `"single"`. Computed.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

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

- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-0233032022033311-3000012333333220-3023221323213013-0000232201012301-3323333020013020-2320312132000122-3131220121310300-1220331333210121)
- [disable_api_definition](data-sources--http_loadbalancer--reference--group-017.md#canonical-0310031002133033-3133323313302002-1023201212303220-0020212011021133-1133233311232010-3211233231213212-0132233203113201-3302223221232231)

Select alternatives according to the provider validators above.

<a id="canonical-3133121120130011-0330013300332230-3332120131223333-2211003101100331-2322303301320033-0330032111001331-0033312000221030-2003102322210332"></a>

### Direct properties for `api_specification`

- [api_definition](data-sources--http_loadbalancer--reference--group-010.md#canonical-1103232023002132-0123230310330020-0130010100010230-0310312113203032-2122111130333020-1300120230202002-2111220332131030-1122332210110331): complete subsection reference.

- [validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321): complete subsection reference.

- [validation_custom_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020): complete subsection reference.

- [validation_disabled](data-sources--http_loadbalancer--reference--group-011.md#canonical-2020131233310012-1202232333311311-0032120321023001-0232332130333112-0110131200031202-1001330331121013-3201323102010002-0200221330310320): complete subsection reference.

<a id="canonical-1103232023002132-0123230310330020-0130010100010230-0310312113203032-2122111130333020-1300120230202002-2111220332131030-1122332210110331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.api_definition` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- api_specification.api_definition

<a id="canonical-2213321323303212-0010130132311231-2123300301220302-2203002332023323-1320213213130000-2012100210002111-2022110300212131-0131312020302001"></a>

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

<a id="canonical-1230022021313013-3321110333230032-3222232003310322-2122002032321233-2113030113332133-1312300332133030-3211332322201012-0323000200221023"></a>

### Direct properties for `api_specification.api_definition`

<a id="canonical-0312012130301320-2002130201030010-3302012033321310-2221233022302133-1101103223020202-2120022213213111-2231302332021202-0321123121121011"></a>

#### `api_specification.api_definition.name` property

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

<a id="canonical-1231232020021030-0211310030033110-2212023321013322-3201331101020300-0202133321333121-0023211131031011-3213120122021231-2031010003032101"></a>

<a id="canonical-2103321302233123-1113030330310311-3320321133132011-0222011210211031-1132011111021123-1122323201110320-3303333011131100-0233001230002123"></a>

#### `api_specification.api_definition.namespace` property

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

<a id="canonical-0103331300021232-2003331033312233-0333122131033131-0033023320313312-0001302021131312-2331312323121323-0331230323213132-3133002331000313"></a>

<a id="canonical-0212122300200302-1333222200100012-3331100110000303-0303323320032120-1121022120322102-2231022002100212-2110222233113100-2231231103213112"></a>

#### `api_specification.api_definition.tenant` property

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

<a id="canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- api_specification.validation_all_spec_endpoints

<a id="canonical-0300130112101230-2300130023121222-0010331130201221-1302302202212001-1210100102000221-1223213302131230-2010000032301013-1333211033012130"></a>

Type: `"single"`. Computed.

API Inventory. Settings for API Inventory validation.

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

<a id="canonical-0223022003210212-2223221222230010-3333320103023020-2002221202011030-3312323332001111-0133201330331021-3111311022000020-3110020132233211"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints`

- [fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133): complete subsection reference.

- [settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033): complete subsection reference.

- [validation_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030): complete subsection reference.

<a id="canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-2020203003301001-1222321311112110-3222210120002030-0320221003230333-2320233212321113-2130312210101232-1100102212232010-0102201032010332"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

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

<a id="canonical-1303332103220211-0001221101013123-1223333133303030-2032212000101010-2130221012021123-2103013232033120-1132200130001002-2100131123330013"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode`

- [fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-010.md#canonical-0222323213331011-3331012010011322-3023222213303310-1232001033102120-3221113033113323-3331322111100332-1112331012300322-0102132302020030): complete subsection reference.

- [fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023): complete subsection reference.

<a id="canonical-0222323213331011-3331012010011322-3023222213303310-1232001033102120-3221113033113323-3331322111100332-1112331012300322-0102132302020030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-2203333311013003-0213110002311213-0031220323203321-2323321031122121-3210033201220303-0203312301320103-3111333320320111-2311012032011002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-1222013132233021-0211110132133301-2211202013130120-1012200102230330-0113201000032010-3012331303211203-2321113133030212-2200112001221012"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Additional upstream details:

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

<a id="canonical-1323313202320233-3213221121323100-3232330302333112-1210010302312102-1330201010103033-1303322221311231-3201310130101232-2222121032333003"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom`

- [open_api_validation_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222): complete subsection reference.

<a id="canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-0221321132110111-0233131232133021-3203013212312302-0302100021230212-1130001332313032-0001210332210223-1211000010021020-1303302231103210"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

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

<a id="canonical-1133200010130301-1112022123121000-2230100322313122-2323022213032211-0123002121120001-1131213313212210-2310131231231221-0201120201220011"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`

- [action_block](data-sources--http_loadbalancer--reference--group-010.md#canonical-1111001011322323-1010220020321130-2203212102130333-1211333023132120-0203003001300331-3123001222320311-3210103012113302-3122220200033110): complete subsection reference.

- [action_report](data-sources--http_loadbalancer--reference--group-010.md#canonical-2211103003120222-2122202132331112-3031300120010000-1302113112000323-1202201123031030-0131112223123102-0310122031103002-1303122220133032): complete subsection reference.

- [action_skip](data-sources--http_loadbalancer--reference--group-010.md#canonical-2200122231211020-0001023322111131-1011333211022033-3230202211313211-2003233200201223-2231111022311232-1202331222001010-0123331302201202): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--reference--group-010.md#canonical-0012223332023113-1301002030011201-3211230101223131-0303322232110312-2323130313210031-3123013220200023-0323220203200223-2333333023201213): complete subsection reference.

<a id="canonical-2002302031133210-0211100220211212-1033030032100322-0123322313203312-2132201323033011-0223311221332023-1013023001122310-1022233212312023"></a>

<a id="canonical-0010220103220202-1222001202312023-2133032203032021-2233020200310201-1203032033013203-3303313212033121-1012002132322112-1133221121230311"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` property

Type: `"string"`. Computed.

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

<a id="canonical-2323303311132223-2331001321333203-2311132012331000-2130121031121303-1330231020213331-1130110123030231-3000103230030231-3310002110212122"></a>

<a id="canonical-0110311210331002-0330230303201111-0322230101102320-3301301233203121-1203312303110031-0102200131100331-3113131111330022-0203103210101301"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` property

Type: `"string"`. Computed.

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

- [metadata](data-sources--http_loadbalancer--reference--group-010.md#canonical-3020032022020013-0023320022000021-2001032112013230-3013123000022110-0301133312020233-2211132210021022-0000030000302100-2113132002130101): complete subsection reference.

<a id="canonical-1111001011322323-1010220020321130-2203212102130333-1211333023132120-0203003001300331-3123001222320311-3210103012113302-3122220200033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-0202220202021311-3313312322000000-1132221113111002-2103210312011020-2022102302232020-3300311120220113-0332011031201332-2132300213303102"></a>

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

<a id="canonical-2211103003120222-2122202132331112-3031300120010000-1302113112000323-1202201123031030-0131112223123102-0310122031103002-1303122220133032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-1310112100131121-0321220210022000-1012330231100123-2211101101130233-0301031323221322-2313313111210222-2110002012211113-1311330000312312"></a>

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

<a id="canonical-2200122231211020-0001023322111131-1011333211022033-3230202211313211-2003233200201223-2231111022311232-1202331222001010-0123331302201202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-2031222113203133-0221013323103201-0033130300020022-1311301010131000-3323002123313302-2102131110331000-1012033123032102-3031013001201133"></a>

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

<a id="canonical-0012223332023113-1301002030011201-3211230101223131-0303322232110312-2323130313210031-3123013220200023-0323220203200223-2333333023201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-2000200310202003-2323022131000210-0112312312122320-1221221101313212-2301111203202011-2230130301222230-3020102120110220-2322130313333102"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2111321103311103-3310032103213310-2233133022002220-2013221113300310-2300021121122110-2001302032220332-3302332000333031-0221311131202223"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint`

<a id="canonical-0012102010033303-1300312320330112-1200231301230202-3121212130200103-1223031203131113-3021030323010013-1020030112030110-0311122333330332"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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

<a id="canonical-1102001330211100-2320102222003221-0232332110202133-2310230033303100-0130332300103100-2132022032202123-1220002330210232-1201332232323233"></a>

<a id="canonical-2330200212310300-1323223212303332-3132211110230110-2223211110313110-0002031310120323-2101312001230022-2013023020310213-1133300113233102"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Computed.

Path. Path to be matched.

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

<a id="canonical-3020032022020013-0023320022000021-2001032112013230-3013123000022110-0301133312020233-2211132210021022-0000030000302100-2113132002130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-0021302312030332-0303201132233001-3230312010323201-3110203320230233-3300110333031120-2132000002312131-2020313030320032-1313211230132131"></a>

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

<a id="canonical-2012023230010231-1333001032200221-3013121000001322-1121100130022221-3032300312100330-1100210202101003-2331302332232301-3133322131220103"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata`

<a id="canonical-2310110312203300-2330001223032212-0230212100032002-0120121220033030-0131130320223310-3122121100310000-2331132113303121-0231120021231321"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0212013303002210-3200020233333023-2330031031113023-2101131202001102-3231231122101031-0200200011330312-2223213321313031-1312111201110132"></a>

<a id="canonical-2031312120333230-0222002223322133-2311200320111311-3003200231210303-1333020001211002-1030133320223002-3132320013101211-2031122010031332"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name` property

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

<a id="canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-0333213303321311-0313303122100312-3320131121213203-1130020332003202-0111223011101032-1312321122020213-2230110300011032-2113022120000010"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Additional upstream details:

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

<a id="canonical-0330233000003022-0220200000212130-0013233001301022-2033123003212000-1323222000312223-1003313313322132-1331232012302002-3322231212220112"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings`

- [oversized_body_fail_validation](data-sources--http_loadbalancer--reference--group-010.md#canonical-3113232302202210-3022000123021322-1231130111032320-1011021230232011-3200213202302112-1323010132133111-0201032010333213-1120210212002131): complete subsection reference.

- [oversized_body_skip_validation](data-sources--http_loadbalancer--reference--group-010.md#canonical-1013232123013112-1211132010230200-1000112320320002-1100331013131201-0221220102322102-1002203213221302-3001313030000221-2301312000201220): complete subsection reference.

- [property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012): complete subsection reference.

- [property_validation_settings_default](data-sources--http_loadbalancer--reference--group-010.md#canonical-3220123201332000-0331112323303301-2101330223310203-0302311120123132-2333320103213120-2120010031023120-0103202000232213-1332331112110303): complete subsection reference.

<a id="canonical-3113232302202210-3022000123021322-1231130111032320-1011021230232011-3200213202302112-1323010132133111-0201032010333213-1120210212002131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-1110122223200132-0133023212132023-3202200303101103-2212103302001011-3133310122322132-2033110230323023-2132101312030320-0023320232102313"></a>

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

<a id="canonical-1013232123013112-1211132010230200-1000112320320002-1100331013131201-0221220102322102-1002203213221302-3001313030000221-2301312000201220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-0011210313111112-0202133011013010-0000102202112200-0303300310213203-0303332132232231-0131011000323303-1302101232222321-0101101223120023"></a>

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

<a id="canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-0001011221301303-2110113101100112-1330311223232331-3120233232020232-0100320001323300-2332313003103310-0231101112303002-2321320210202333"></a>

Type: `"single"`. Computed.

Configuration parameter for property validation settings custom.

Additional upstream details:

Custom property validation settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2023110133103232-2032031213032200-2012303121023310-0122232031231221-3002102303133132-0321020333101213-2131223112132301-2101210120223222"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom`

- [query_parameters](data-sources--http_loadbalancer--reference--group-010.md#canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330): complete subsection reference.

<a id="canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-2323332022220100-0331232002302100-1110122130222310-2002320100233221-0201003330213310-1123320133231112-0001222303301322-1332122222113220"></a>

Type: `"single"`. Computed.

Custom settings for query parameters validation.

<a id="canonical-3303010110001203-1131013221003320-0303101121001112-0201201303131213-0000322322232303-1101020012332030-0232310231031300-1300312220103202"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters`

- [allow_additional_parameters](data-sources--http_loadbalancer--reference--group-010.md#canonical-1333210300113120-0122203221120030-3311320212031000-2231122010213203-0323211101012030-0102110001203211-1333332122221230-1233323301120303): complete subsection reference.

- [disallow_additional_parameters](data-sources--http_loadbalancer--reference--group-010.md#canonical-0330321311010311-3133200223221032-1331111221033122-3102223010112303-2120123010133320-0121132231021010-0002230330113011-0323230113232020): complete subsection reference.

<a id="canonical-1333210300113120-0122203221120030-3311320212031000-2231122010213203-0323211101012030-0102110001203211-1333332122221230-1233323301120303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-010.md#canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-2231001101313222-3033100113331011-2232102101131120-2022013213112300-1003000311311302-1002000233312233-2202313223002013-3333202121011010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow additional parameters.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330321311010311-3133200223221032-1331111221033122-3102223010112303-2120123010133320-0121132231021010-0002230330113011-0323230113232020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-010.md#canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-3202032012013033-0001122011332023-1001332231300030-0031113232113133-3000033211323013-3031313220000213-3223323022330011-3223323023003000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disallow additional parameters.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220123201332000-0331112323303301-2101330223310203-0302311120123132-2333320103213120-2120010031023120-0103202000232213-1332331112110303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-2031230212122022-0101121101230231-3130132032332111-3212132000002021-2020101310023223-3132033201011202-0122213133130130-0222110120102102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for property validation settings default.

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

<a id="canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="canonical-2332132223223231-2223132301000320-3333121011113200-3231330300110130-1210013120122321-2310122202110312-2013211011032033-1332233133220032"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

<a id="canonical-1223232121213220-0130312121011221-3003202122033313-2122201002302133-3322202103110023-1111120033212012-3210321111222332-1011032100033021"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode`

- [response_validation_mode_active](data-sources--http_loadbalancer--reference--group-010.md#canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203): complete subsection reference.

- [skip_response_validation](data-sources--http_loadbalancer--reference--group-010.md#canonical-3132203122300201-2033033330013231-0011311333222001-3332223323003210-0223010111120230-0301121301032132-2233203322301220-2002213320322312): complete subsection reference.

- [skip_validation](data-sources--http_loadbalancer--reference--group-010.md#canonical-0301113120020203-0120232321323230-0300313133221220-2222202200111002-2030333130000033-3030031211203331-0003112022013001-0003302203031320): complete subsection reference.

- [validation_mode_active](data-sources--http_loadbalancer--reference--group-010.md#canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131): complete subsection reference.

<a id="canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active

<a id="canonical-2212232323103202-3313303212012333-0131311213231112-0012132220301231-3322322020210132-3121031023101201-0333120211302223-0210132322100310"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-1212123013123233-2231333002003112-0201112303111303-0310020213100223-0233231020123320-3202112212023302-3201211002312203-1323213000103003"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active`

- [enforcement_block](data-sources--http_loadbalancer--reference--group-010.md#canonical-3133312121132100-2220000100320020-2013322310103113-0230201130123330-1230221221013112-1312233012312312-1300132302112031-2131003021220133): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-010.md#canonical-0332012230010002-3213030320323032-1331200221132313-2332212311221301-0321010222031121-0002202320313002-0332131222103113-1320030200031221): complete subsection reference.

<a id="canonical-2300132123102120-0030321203031231-3011330332333223-0030032303221011-3231103320103322-3300302113233200-0032100101022322-3310010322330002"></a>

<a id="canonical-3321303231030113-1031230013030002-0213101032003310-2231301323021330-1023332201113103-2311123321000031-0310200303022302-0311003311320312"></a>

#### `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.response_validation_properties` property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3133312121132100-2220000100320020-2013322310103113-0230201130123330-1230221221013112-1312233012312312-1300132302112031-2131003021220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-010.md#canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-1213201023312212-2112132203123330-0313323321122123-2000330120030202-0013333222132331-3313133213212302-3102301010132331-2013313303311032"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0332012230010002-3213030320323032-1331200221132313-2332212311221301-0321010222031121-0002202320313002-0332131222103113-1320030200031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-010.md#canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-0221012213111000-3212031303322213-1203321233333201-3323200333203010-1313013001010221-2313033113322013-2320300302313131-2003230220020203"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3132203122300201-2033033330013231-0011311333222001-3332223323003210-0223010111120230-0301121301032132-2233203322301220-2002213320322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="canonical-1321022011123031-1033321031211332-3003121112223222-0333223212030313-0030022232331331-3020002323223202-2310013332313323-3011113030100222"></a>

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

<a id="canonical-0301113120020203-0120232321323230-0300313133221220-2222202200111002-2030333130000033-3030031211203331-0003112022013001-0003302203031320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.skip_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_validation

<a id="canonical-3322231322213301-1020321311302223-1022022113100221-3120220123222231-0103221031132233-0020112322213130-2020331123210001-1233012221300312"></a>

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

<a id="canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="canonical-1222230200100133-0230332211023322-0300213031011211-2123011232112112-3112302112100121-2110010300220102-1122031101033302-3210022300031021"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-1220011033321012-1331332230021002-2121023301111221-0222313122012012-3321112103032013-0122333123321112-1301021001323100-0132322121123021"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active`

- [enforcement_block](data-sources--http_loadbalancer--reference--group-010.md#canonical-3000232033312222-3220022322231321-2212102002220222-3211201111222322-0133233300232002-2133333311312331-1122013032102123-2210312021023001): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-010.md#canonical-3331103210000311-1011321023313011-0132201220120311-2010002300033130-2302121022123222-0202132331000202-2320201210220130-0213023012231213): complete subsection reference.

<a id="canonical-0013301023023212-3133232032102021-0100333011302200-3102032021133132-2211000130030330-0331131012121131-1032232213232032-2101233023310133"></a>

<a id="canonical-1221013302200333-0202330231200130-3213023312223112-0310312332023020-0101133212003330-0331021321120021-3302233132200200-0012321020003121"></a>

#### `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.request_validation_properties` property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3000232033312222-3220022322231321-2212102002220222-3211201111222322-0133233300232002-2133333311312331-1122013032102123-2210312021023001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-010.md#canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-0033033132310010-0033103331310012-0023232021330133-3223322133323211-3112220222111032-2022103202310301-0321030212330130-1130020002021213"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3331103210000311-1011321023313011-0132201220120311-2010002300033130-2302121022123222-0202132331000202-2320201210220130-0213023012231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-010.md#canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-1132310122233223-1300323121301100-2020102022330301-0311110112113010-0212220301303000-3232023202022012-1003130201120132-3300031210200330"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- api_specification.validation_custom_list

<a id="canonical-3101132312231321-3223323020302011-1221032111111101-2012333233023122-1333331130232323-3013122202333212-0102323010000300-0023000021312223"></a>

Type: `"single"`. Computed.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Additional upstream details:

Any other API-endpoint not listed will act according to "Fall Through Mode".

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

<a id="canonical-3033023323212213-0231200213001302-3020311303112011-0222213301103131-2120130310111310-1213013012021311-1310021322000313-0123101133000001"></a>

### Direct properties for `api_specification.validation_custom_list`

- [fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111): complete subsection reference.

- [open_api_validation_rules](data-sources--http_loadbalancer--reference--group-011.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231): complete subsection reference.

- [settings](data-sources--http_loadbalancer--reference--group-011.md#canonical-1112200302330330-0330323032110031-3333202023000202-3101020301022323-2022212332013123-0100001232113333-1301210303031331-2322220133323302): complete subsection reference.

<a id="canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- api_specification.validation_custom_list.fall_through_mode

<a id="canonical-0331300023303112-1113002331201200-2233031330023223-0022213102332232-1033320010020010-3121112100303020-1310023303300020-1123022212103322"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

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

<a id="canonical-2313220233202003-0010000112101233-2001122213301323-0103331133212131-3301331121223032-3200002022210122-1123121220321021-3220201110132233"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode`

- [fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-010.md#canonical-2021002022311320-3312333031031200-2230023210200122-3110221031110210-1132133211000120-2033323313231001-2022331211323130-2332030020002120): complete subsection reference.

- [fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233): complete subsection reference.

<a id="canonical-2021002022311320-3312333031031200-2230023210200122-3110221031110210-1132133211000120-2033323313231001-2022331211323130-2332030020002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow

<a id="canonical-2003031030233012-0313301022023321-3100000102330032-2013113223002303-0003033103200022-1020212201123100-1013201220002010-1323220133332011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

<a id="canonical-2320231033113033-0332333222102123-3222313000032232-1001213011022220-0010233233002123-0002013210133232-0011220113001322-2112200033122031"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Additional upstream details:

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

<a id="canonical-0302001331311032-1330222102123332-3100012221012013-0322311322122232-0310203132211333-0002132020213201-3330112233000111-3122101321100010"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom`

- [open_api_validation_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232): complete subsection reference.

<a id="canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-0330231301121223-0022222023331222-0001220202322223-1300232110332030-3131112302203100-2301001110213320-3300132023033202-3210122223000001"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

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

<a id="canonical-3303011210033032-0121020131221002-1322000233002212-3033012303033311-1020002320030222-1021120223212012-0102301131132223-3031301133210101"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`

- [action_block](data-sources--http_loadbalancer--reference--group-010.md#canonical-0111101003232101-3131100310020233-2010312231212110-2313002113010123-1333322033101130-3121111300312130-2322013132223311-3111133312323323): complete subsection reference.

- [action_report](data-sources--http_loadbalancer--reference--group-011.md#canonical-1110203031310300-1221132322033110-3232321323312231-0130233322323301-2121010221023323-3200012322000200-2213023201121021-1132202121220001): complete subsection reference.

- [action_skip](data-sources--http_loadbalancer--reference--group-011.md#canonical-0333203022231331-3123012100031001-3222301122310110-2201220333210122-0121121233130032-3203331100330231-2000101033132133-1100022230213010): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--reference--group-011.md#canonical-1221313211210230-1332320300101232-0201112130123232-1312013223130000-3102212012111222-1021232200203020-0131102013121221-3303113003212302): complete subsection reference.

<a id="canonical-3303201020220210-3020030211132101-1302232221133320-0221021300233011-1122002310123012-0111121213222121-1101133133211201-3131311033211100"></a>

<a id="canonical-0232110211302301-2332323002100332-0030022113130233-2201232023220033-2212331111101131-3113001131130112-2222020212300321-0031120230212020"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` property

Type: `"string"`. Computed.

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

<a id="canonical-1112120120013312-3332232321000213-0313120011233203-2303123320310222-3232303133100231-3101302333001232-0032200323103333-0310300321323212"></a>

<a id="canonical-3223232033321301-1321023023130111-1112200011322130-1331120001000222-0000001000221000-1000211101332223-1300111033101102-0000121013110123"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` property

Type: `"string"`. Computed.

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

- [metadata](data-sources--http_loadbalancer--reference--group-011.md#canonical-3230303110032112-1233220022032312-0230100031022321-3220321200100002-3120330332220313-3032320112122231-1220303221310301-2113322030112220): complete subsection reference.

<a id="canonical-0111101003232101-3131100310020233-2010312231212110-2313002113010123-1333322033101130-3121111300312130-2322013132223311-3111133312323323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-010.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-010.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-010.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-010.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-3000232303110113-2321011002333013-0213332021003332-3120011000333100-1132122213132202-0120001221103123-0121230121133233-3032233322312111"></a>

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
