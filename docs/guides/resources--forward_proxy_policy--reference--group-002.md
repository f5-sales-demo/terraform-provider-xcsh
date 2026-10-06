---
page_title: "xcsh_forward_proxy_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy reference."
---

# xcsh_forward_proxy_policy reference

<a id="canonical-0221230323223211-3010132100100233-3210122132012122-1223331132311301-3301002003203030-3311233232103222-1102003000313200-0033331023010010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_label_selector` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- proxy_label_selector

<a id="canonical-3111002303022031-3313302320131202-0002322100332212-1022112321301020-3203320320321333-0222031133023001-2312200122230321-1032320001202331"></a>

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
proxy_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311312113310003-1023132120210200-2013013303013130-2330301313010201-3213133020131103-1223232211010322-0201031302222310-3003232022213320"></a>

### Direct properties for `proxy_label_selector`

<a id="canonical-3201211130202332-2202302330013100-0300022120303313-1132021123100111-0200313000200102-2313001132212200-2112010202220210-1021223303003112"></a>

#### `proxy_label_selector.expressions` property

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

<a id="canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- rule_list

<a id="canonical-3120101031021000-0311222230303232-3110212032012330-3000223220331220-2233310120133312-0001101213022230-2123233112313210-2100201222030032"></a>

Type: `"object"`. single nested block, Optional.

Custom Rule List. List of custom rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2223002031000023-1313320302122313-0300332110220013-2122303312221303-0313310233101131-3020201203003021-2131302130002032-1021322230123032"></a>

### Direct properties for `rule_list`

- [rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330): complete subsection reference.

<a id="canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- rule_list.rules

<a id="canonical-3312120222023102-3011020210320310-3103002210201000-0003220102311011-3111011031221010-0103012112312202-3021233312200333-0333313230200112"></a>

Type: `"object"`. list nested block, Optional.

Custom Rule List. List of custom rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300322033120023-2302101301222021-0021102213010021-1302000201212331-2313321233011312-2222110001123022-2310010310131302-3301103002200022"></a>

### Direct properties for `rule_list.rules`

<a id="canonical-1210301032202321-3321023231033121-3012222222312121-2213000223102330-3221112323212121-2220323212113001-0301111201301003-1021313123100030"></a>

#### `rule_list.rules.action` property

Type: `"string"`. Optional.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [all_destinations](resources--forward_proxy_policy--reference--group-002.md#canonical-3303212022323032-0022002332220301-1322021013031312-2132120202033330-2121131310233020-0103000131331110-2012330122300033-1031202211311121): complete subsection reference.

- [all_sources](resources--forward_proxy_policy--reference--group-002.md#canonical-2222002001020101-1111232232301313-0120210021110313-1122331022321320-1033330212102033-2100023110012212-2222020302330033-1303301022233201): complete subsection reference.

- [dst_asn_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2011102313313211-2131223301223322-0233232012200022-3213221220130021-0300133023100031-0302010133211122-0000302210010220-0302330123310033): complete subsection reference.

- [dst_asn_set](resources--forward_proxy_policy--reference--group-002.md#canonical-0003132101030212-3123122112331210-3222323123002210-2001112030002030-0201331012012233-1132230211211100-0300021210200213-3101023102002203): complete subsection reference.

- [dst_ip_prefix_set](resources--forward_proxy_policy--reference--group-002.md#canonical-2212102330010013-2031123132223301-0131212022001000-2233000012210232-0111000131320100-3220031130123032-3123221133100232-2303223202311211): complete subsection reference.

- [dst_label_selector](resources--forward_proxy_policy--reference--group-002.md#canonical-2110220112000323-1230101213332333-3323222310110320-3030023320300210-3110111103122013-0113132230312113-0332013123321023-3023112213221112): complete subsection reference.

- [dst_prefix_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2220013131103110-0321123030130202-1211323031103321-3120101110022022-2312300220103132-3303110213302112-2320201002022211-3233200233130211): complete subsection reference.

- [http_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220): complete subsection reference.

- [ip_prefix_set](resources--forward_proxy_policy--reference--group-002.md#canonical-0310221030310313-2132220021203023-2232111022112130-3332333211121113-3031033033311333-0320210103020113-3230210210002002-0321121231301103): complete subsection reference.

- [label_selector](resources--forward_proxy_policy--reference--group-002.md#canonical-2112013033001320-0321130123000130-0033123321022312-1123023302201232-3023123010010201-0202333213232300-1112213012003221-3122233032102202): complete subsection reference.

- [metadata](resources--forward_proxy_policy--reference--group-002.md#canonical-3030113001113311-1321300323302332-1203022110332230-1131331102123232-3132221000110231-0333211230303301-3323103223120032-0220113312111300): complete subsection reference.

- [no_http_connect_port](resources--forward_proxy_policy--reference--group-002.md#canonical-3310331010213110-3101213300003222-3021120021321003-0333023333133021-2032303012222011-1332203203222230-0321133031203200-3003323103232221): complete subsection reference.

- [port_matcher](resources--forward_proxy_policy--reference--group-002.md#canonical-1310001202000220-0222212030212031-3110301233213311-0331231201211302-1222330112030113-0332221131200032-0000000110301303-1000000231020321): complete subsection reference.

- [prefix_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0312131012032312-0131110323211333-0123311001012011-2103130111123002-2020111100012001-3211231333323102-3111311122001102-1233223101130020): complete subsection reference.

- [tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0321220033202311-3123011210302013-3200210310012232-1230303120121032-2003133313330030-3321320211122231-2200231330033213-2211001300113003): complete subsection reference.

- [url_category_list](resources--forward_proxy_policy--reference--group-002.md#canonical-3133120013122023-1010000233103233-3100211201111331-1132333033310100-1001321120000211-0013313110101113-0312320322323313-0031022220302232): complete subsection reference.

<a id="canonical-3303212022323032-0022002332220301-1322021013031312-2132120202033330-2121131310233020-0103000131331110-2012330122300033-1031202211311121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_destinations` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.all_destinations

<a id="canonical-3132030112213100-2310103330022200-2131003300123302-1132133201031302-2103020101332133-1302013203011222-3012230100000133-1323023210222320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_destinations = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222002001020101-1111232232301313-0120210021110313-1122331022321320-1033330212102033-2100023110012212-2222020302330033-1303301022233201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_sources` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.all_sources

<a id="canonical-3212102110303220-1013322223223200-1012020002331200-0212001203213020-2323303320021303-3322102310313321-1021033112030103-3312210203030223"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_sources = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011102313313211-2131223301223322-0233232012200022-3213221220130021-0300133023100031-0302010133211122-0000302210010220-0302330123310033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_asn_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.dst_asn_list

<a id="canonical-0133131032312222-3031301123203003-2313303030213010-1313332220210321-3002120022010323-0202200023303100-1030300300322002-1200331011132330"></a>

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
dst_asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003132301232130-1200001233002332-1203013031020132-1123001131011213-1130103221233133-0131122230120000-3200003130130133-2220022330021232"></a>

### Direct properties for `rule_list.rules.dst_asn_list`

<a id="canonical-3122100001023130-1031310110202203-2320323301313322-2101212013320131-3303133021033233-0310230113000310-1322033110313122-0030033232102132"></a>

#### `rule_list.rules.dst_asn_list.as_numbers` property

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

<a id="canonical-0003132101030212-3123122112331210-3222323123002210-2001112030002030-0201331012012233-1132230211211100-0300021210200213-3101023102002203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_asn_set` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.dst_asn_set

<a id="canonical-2331112120122301-0133313301032221-0012303212331010-3021321001312212-2011002333330103-1011203112230202-1012223023111230-3122202012111221"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2110113321033112-2011001302312032-0200033122020011-3012321322230020-3002120202130020-0321103023111022-3322320132233101-3212310120313012"></a>

### Direct properties for `rule_list.rules.dst_asn_set`

<a id="canonical-2322000130221222-1321221031303133-0032001101310020-2130212203122212-3211100001133102-1000122221100303-1112313001031313-3032303020230001"></a>

#### `rule_list.rules.dst_asn_set.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2001131231213200-0003031233302220-0010301231230113-0233023021113333-0222111301013010-2321132202331030-0121201201113332-2332233323310122"></a>

<a id="canonical-2210303322031023-3113003303031033-1131233030232021-3030212030211003-1322123112200033-0033300100301302-0330211232310111-2103313300113212"></a>

#### `rule_list.rules.dst_asn_set.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3033301130331101-0201132001103210-3012030332110000-0031113331303210-3331302031021301-3023201223030311-0302000312031001-1010031303212302"></a>

<a id="canonical-3011112233132201-0032321323301231-1331203211300330-2120203233203323-3132202222132103-1220332012233110-1200313200003313-1122032101232120"></a>

#### `rule_list.rules.dst_asn_set.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2212102330010013-2031123132223301-0131212022001000-2233000012210232-0111000131320100-3220031130123032-3123221133100232-2303223202311211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_ip_prefix_set` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.dst_ip_prefix_set

<a id="canonical-0011232133000131-0011100200323101-3030033303023333-2022302203321221-0302332210113131-2213100122312100-0002032010322222-0203323300320333"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3020001111210133-2321100111000002-3320001132301210-3010203222002300-1011221023101320-0230032311223223-2001233221013221-1311020320303301"></a>

### Direct properties for `rule_list.rules.dst_ip_prefix_set`

<a id="canonical-2311223211201033-0000130323333133-1320113031333113-2303021333301121-0322023201012113-1321003022001233-2122312210302132-0212312201131203"></a>

#### `rule_list.rules.dst_ip_prefix_set.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3232300201110110-3313012332102222-1300330111333200-0012222112102123-3120123231220133-0033002310121010-2211122013030021-3233003303032033"></a>

<a id="canonical-3111010032002320-3302200020221100-0301203110331303-1003031102030000-0121030310311101-3323200203001331-2301030112201321-1331231102113303"></a>

#### `rule_list.rules.dst_ip_prefix_set.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0113010212002230-0211133322030033-2210022102230201-0330111011313212-0010131203211201-1201222120201001-2023221113111110-3110322102132021"></a>

<a id="canonical-3312102212101312-2002132002312321-1303310300103230-2031112133103110-1323031220213213-3232201310000102-0301122233010310-1231000000113000"></a>

#### `rule_list.rules.dst_ip_prefix_set.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2110220112000323-1230101213332333-3323222310110320-3030023320300210-3110111103122013-0113132230312113-0332013123321023-3023112213221112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_label_selector` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.dst_label_selector

<a id="canonical-3333031302001101-0030110031100012-3200112300222120-1330112310103100-0132330023312120-2222301223332202-1300021322022100-3323210112201323"></a>

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
dst_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130102032322232-2330003002311010-2213001033123112-0221033111220013-0123223111321013-3201222023133203-0233233032222202-3201023231312331"></a>

### Direct properties for `rule_list.rules.dst_label_selector`

<a id="canonical-0323111132202312-0031303221201210-3021012033322010-1310333200133110-2202123021111223-1323330200131233-0112302030000220-1003210113132111"></a>

#### `rule_list.rules.dst_label_selector.expressions` property

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

<a id="canonical-2220013131103110-0321123030130202-1211323031103321-3120101110022022-2312300220103132-3303110213302112-2320201002022211-3233200233130211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.dst_prefix_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
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

<a id="canonical-1100200021001013-3201110231320211-3011300101320221-1310301300130102-0333223003232101-0133032223200031-2111300111203230-2313330102112020"></a>

### Direct properties for `rule_list.rules.dst_prefix_list`

<a id="canonical-1123201213103230-0212322122321330-2123100330003321-2322212312312213-1003202210102321-0231113301131021-2333113331203202-3123302003313213"></a>

#### `rule_list.rules.dst_prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.http_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
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

<a id="canonical-0002320320313230-3110310102220102-3211132113331322-0111103330331100-3023131221232013-0322103101321322-2231023310131331-1323303013310031"></a>

### Direct properties for `rule_list.rules.http_list`

- [http_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2233313011311201-0112331132232311-1302312313130130-3213111133013310-0221230302221331-2221103121012302-0200000031133101-1313303121032132): complete subsection reference.

<a id="canonical-2233313011311201-0112331132232311-1302312313130130-3213111133013310-0221230302221331-2221103121012302-0200000031133101-1313303121032132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.http_list.http_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220)
- rule_list.rules.http_list.http_list

<a id="canonical-3233101000301203-2221302110220122-3120023122101002-2103030031301331-2102112332233320-2132233021113322-3300223022101021-1110310001131020"></a>

Type: `"object"`. list nested block, Optional.

HTTP URLs. URLs for HTTP connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

Terraform syntax:

```terraform
http_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022110320321102-2133122032211102-1200201101033002-0111131201323121-3301333033013331-0322030101032013-1032112300312310-1032332002220122"></a>

### Direct properties for `rule_list.rules.http_list.http_list`

- [any_path](resources--forward_proxy_policy--reference--group-002.md#canonical-2003033123330321-2120103022031003-0010013112313223-0333323210322231-0022213123230131-3321212132120223-3102310210211020-1003331133201333): complete subsection reference.

<a id="canonical-0110303312312100-3302332023333011-3013211102110311-1332321111223000-1211030231013022-0130233232003230-0302010002233123-1002000330312222"></a>

<a id="canonical-3231202220002233-3232313002310112-0332310322331222-2012210203302310-0220213220120011-2333220002220332-2201011021003331-1221030013201213"></a>

#### `rule_list.rules.http_list.http_list.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-1030230101321100-1320202313303213-0103113132013002-0303123013001311-0230201220102033-0222131300323020-3001221033201333-2223323001002233"></a>

<a id="canonical-2230030333103300-0131220210001032-0321301120102300-2231322212101123-3110231323020200-0011111212302031-0003010301100202-0203221011021221"></a>

#### `rule_list.rules.http_list.http_list.path_exact_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\_value path\_regex\_value\] Exact Path to match.

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

<a id="canonical-1023000131030012-0201003310033101-0033320003332302-3031201313032022-1001201121030000-2312330012311310-1013022130120212-0100002203013231"></a>

<a id="canonical-3321330203331300-1230110133301003-0220001133112331-0131231011223332-2011110210022212-0010301132020331-3310333210302130-2111320103132323"></a>

#### `rule_list.rules.http_list.http_list.path_prefix_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g '/abc/xyz'
will match '/abc/xyz/.\*'.

Additional upstream details:

Exclusive with \[any\_path path\_exact\_value path\_regex\_value\] Prefix of Path e.g "/abc/xyz"
will match "/abc/xyz/.\*"

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

<a id="canonical-2321103313110001-0220320211301003-1231110312331110-2222213123220123-3013113011312202-2133321323331223-1102220123121220-0331112123220121"></a>

<a id="canonical-1211111010101011-3233130310020212-3313332120300222-1001332313230112-2220221211230002-2333102102001213-3212220032333301-1311213121313223"></a>

#### `rule_list.rules.http_list.http_list.path_regex_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_exact\_value path\_prefix\_value\] Regular Expression value for the
Path to match.

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

<a id="canonical-3103123221200021-3003232300100022-2111010130321012-0023121232312321-1312133230011200-0322300122033311-3020100022021201-3222312221301311"></a>

<a id="canonical-1003101121200230-2112021301223332-3101320102300113-3332023133112030-3100323012303102-3200020031231112-2230222313331230-0031002103030033"></a>

#### `rule_list.rules.http_list.http_list.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-3312001323010103-2001203101301300-3011033023130120-0111033103113213-2003321023232010-1201123221220112-3103132112103011-0231230112101223"></a>

<a id="canonical-0312011322111022-0322023312310211-2232202221202210-2200321022223102-3310122001103131-2123111100333101-1320300222010313-0100310220133132"></a>

#### `rule_list.rules.http_list.http_list.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g 'xyz.com' will match
'\*.xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain names e.g "xyz.com" will match
"\*.xyz.com"

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

<a id="canonical-2003033123330321-2120103022031003-0010013112313223-0333323210322231-0022213123230131-3321212132120223-3102310210211020-1003331133201333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.http_list.http_list.any_path` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [rule_list.rules.http_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2221002323000201-0300021002302123-3110200203220001-1211121002132203-0201221220221001-1211313312213133-0310230023120021-1331032103220220)
- [rule_list.rules.http_list.http_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2233313011311201-0112331132232311-1302312313130130-3213111133013310-0221230302221331-2221103121012302-0200000031133101-1313303121032132)
- rule_list.rules.http_list.http_list.any_path

<a id="canonical-3202321113213112-0133121003330220-1110021232002330-0003032332330231-2030303330111010-0032032002123322-3231021102212313-1330020232232221"></a>

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
any_path = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310221030310313-2132220021203023-2232111022112130-3332333211121113-3031033033311333-0320210103020113-3230210210002002-0321121231301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.ip_prefix_set

<a id="canonical-3121023022033302-1031312030100132-0330222032021033-1032003012313021-1111132222010301-1201132021123313-3320303010120013-3311010121132023"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0220221300000200-3312130203033302-0302231112100230-3100212132210211-1112130020310233-1213300213311122-0003320322220332-2332230213302333"></a>

### Direct properties for `rule_list.rules.ip_prefix_set`

<a id="canonical-2012010232113010-0022122301320321-2332020121012222-2012112123021011-0110010201230221-3100302001030100-3101220203020122-2322031231120023"></a>

#### `rule_list.rules.ip_prefix_set.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0333111100230131-1132003010020103-2013122011020022-2210320330111033-3313310322023312-0233022323031230-2033211011121300-0303103022002103"></a>

<a id="canonical-1220131003222210-3132013323213110-1210321313030121-3321223011002221-1002011120103211-0021331322202203-2330010233333032-3012111013113331"></a>

#### `rule_list.rules.ip_prefix_set.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3200333102332123-3301130310212333-3122102030101313-0300002032000321-2311210311332130-2232320100322102-1012333101323210-1302001331302021"></a>

<a id="canonical-0323102032112301-3201003002030131-2331323133011303-2100113101121331-2223120032332223-3003321013012013-1333132230122010-0232032002200203"></a>

#### `rule_list.rules.ip_prefix_set.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2112013033001320-0321130123000130-0033123321022312-1123023302201232-3023123010010201-0202333213232300-1112213012003221-3122233032102202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.label_selector` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.label_selector

<a id="canonical-3121111211111121-1122302203201112-3102332311230230-3203303120100331-3021222232232112-3133330102113323-2112203100221130-3322102323110331"></a>

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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033021101000110-2123233022020313-3100110013303321-1121210102011330-0232213221012022-2112120202133330-0013013130202133-1123211323102201"></a>

### Direct properties for `rule_list.rules.label_selector`

<a id="canonical-2301003312033221-2123222132203212-3121001210320023-2103033321110233-1230021311201000-2002231310010012-3003212231231022-2313022330311321"></a>

#### `rule_list.rules.label_selector.expressions` property

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

<a id="canonical-3030113001113311-1321300323302332-1203022110332230-1131331102123232-3132221000110231-0333211230303301-3323103223120032-0220113312111300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.metadata` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.metadata

<a id="canonical-1332230121002020-3211213221030200-2100000310310030-2010212011132131-0020012313222302-2212122211112020-2130030012301333-1312320023312112"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2330021220320032-3313003311333003-0332000232321113-2221031322133022-0033230010130132-1032211333313120-1310122233023021-2012130203322213"></a>

### Direct properties for `rule_list.rules.metadata`

<a id="canonical-0233103121111200-2102202012102232-1232130022221021-3011020121123011-0021131301002021-3100000210113020-2323223031000121-3320112203111311"></a>

#### `rule_list.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1210211300000123-1010312231021100-1002312113332021-0211031021311120-0310310012311330-1220100012201131-1200010330123212-0223033020112330"></a>

<a id="canonical-3321100210323302-0120301230333012-0203001132113133-2003011200211212-3020022202220001-3122212111221303-0122000223031321-1011211332200213"></a>

#### `rule_list.rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-3310331010213110-3101213300003222-3021120021321003-0333023333133021-2032303012222011-1332203203222230-0321133031203200-3003323103232221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.no_http_connect_port` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.no_http_connect_port

<a id="canonical-0111023013123131-3100003330320311-1212021131321300-1233200332221332-3212323113123333-1123130023200200-0130023123113223-2103000121330203"></a>

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
no_http_connect_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310001202000220-0222212030212031-3110301233213311-0331231201211302-1222330112030113-0332221131200032-0000000110301303-1000000231020321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.port_matcher` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.port_matcher

<a id="canonical-2223210130031330-3322111110333000-0232120132010310-2310330231232132-3323020012121232-2033322201200330-0302302232022332-1031020232131121"></a>

Type: `"object"`. single nested block, Optional.

A port matcher specifies a list of port ranges as match criteria. The match is considered successful
if the input port falls within any of the port ranges. The result of the match is inverted if
invert\_matcher is true.

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

<a id="canonical-3201330013320233-3223133011203333-1331020110112323-3020203211003321-0001311333032023-0322321300132231-2322113222302212-0311032223032301"></a>

### Direct properties for `rule_list.rules.port_matcher`

<a id="canonical-2031230022033200-1220133000313003-3310210211031010-2312003221303330-0313231111130210-2032003311133131-0102223231332201-0322012233103232"></a>

#### `rule_list.rules.port_matcher.invert_matcher` property

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

<a id="canonical-3301203210020311-3231130311123200-1321231101330031-1101233123123203-1002322130031303-1010101013211122-0102111221101023-1322013330103312"></a>

<a id="canonical-1222202123030022-3222211120211132-2030000122211113-2030021300322300-2103331223331200-2021221222122312-0112212223330022-2002323122130101"></a>

#### `rule_list.rules.port_matcher.ports` property

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

<a id="canonical-0312131012032312-0131110323211333-0123311001012011-2103130111123002-2020111100012001-3211231333323102-3111311122001102-1233223101130020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.prefix_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.prefix_list

<a id="canonical-1211122323021202-0133301231103221-0220121222330021-2011313231220002-1302231303213021-2132320101120110-1120033302011123-0303330333031021"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010020032233031-3032330131231012-0310231132323010-2130200112221320-3032301110103222-3032200132322111-1110010122120231-0332311000113211"></a>

### Direct properties for `rule_list.rules.prefix_list`

<a id="canonical-3202103130210002-0303202322222201-3131323233020000-1032321012120302-3303222032300010-1310112310113321-1223132120323121-0021212203132213"></a>

#### `rule_list.rules.prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-0321220033202311-3123011210302013-3200210310012232-1230303120121032-2003133313330030-3321320211122231-2200231330033213-2211001300113003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.tls_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.tls_list

<a id="canonical-1210221123213132-0203202220231001-0221211321313010-3132220223103201-3330133110311231-3313102001230210-0012003131101132-3201023102012202"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101012300101023-2122312310212121-2022131123221103-1220020122201223-0002000232233012-1122210213130301-3033003002100233-3122232230322211"></a>

### Direct properties for `rule_list.rules.tls_list`

- [tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-1133323022102133-0031321013011302-2333300131111211-1232030103102131-3021311322113232-2331003013012020-3122100212320020-2220113100303010): complete subsection reference.

<a id="canonical-1133323022102133-0031321013011302-2333300131111211-1232030103102131-3021311322113232-2331003013012020-3122100212320020-2220113100303010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.tls_list.tls_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- [rule_list.rules.tls_list](resources--forward_proxy_policy--reference--group-002.md#canonical-0321220033202311-3123011210302013-3200210310012232-1230303120121032-2003133313330030-3321320211122231-2200231330033213-2211001300113003)
- rule_list.rules.tls_list.tls_list

<a id="canonical-0320220031133313-0130201231002103-1231201120112001-1331023212011310-2033322223202032-0301311201130303-1232213000333200-0320323303110202"></a>

Type: `"object"`. list nested block, Optional.

TLS Domains. Domains in SNI for TLS connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

Terraform syntax:

```terraform
tls_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020233102203123-1331131101100321-2302002203321213-2113223210121103-0301000310301332-3322020013033000-2133320021221212-1102033130302220"></a>

### Direct properties for `rule_list.rules.tls_list.tls_list`

<a id="canonical-2201020223300212-1232232212110133-3203021000333022-1301110310010322-1210022220123121-2222121002120100-3020223032103200-1303231320212113"></a>

#### `rule_list.rules.tls_list.tls_list.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-2333312232301233-2123321110303321-2300230212320103-1020203302000201-1000200322201031-1221131001330232-2221200101032223-1222132303101233"></a>

<a id="canonical-0030203232110213-1323010232321110-2232200000011032-0023311101212330-0110010120301013-1300211123131220-3110002003100122-2121001223113332"></a>

#### `rule_list.rules.tls_list.tls_list.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-2131322113012020-3133313202132232-2122222001332100-3103212103122010-1103322003212310-2212321111333220-2311013330300120-2002222202001222"></a>

<a id="canonical-0233100021313032-1303001221311132-2332210103121012-2120101011013322-2003110111303323-3031012223100003-0312100112200001-1120120010002122"></a>

#### `rule_list.rules.tls_list.tls_list.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-3133120013122023-1010000233103233-3100211201111331-1132333033310100-1001321120000211-0013313110101113-0312320322323313-0031022220302232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.url_category_list` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- [rule_list](resources--forward_proxy_policy--reference--group-002.md#canonical-2200222322222003-2320113212000003-2010311201313302-1220311011100232-2213331130301132-1011200321222020-2113222221233020-2330221022330120)
- [rule_list.rules](resources--forward_proxy_policy--reference--group-002.md#canonical-2233222310220010-1302001131122113-2331222302021230-3133211001321300-2110130032212032-2111320023110111-2102100032323333-1012303112322330)
- rule_list.rules.url_category_list

<a id="canonical-1101212112133130-0012123030323210-2001102332220101-0320131122131103-2002132301022320-1233000110003110-0210213223123000-0312130202013230"></a>

Type: `"object"`. single nested block, Optional.

URL Category List Type. List of URL categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url_categories")}
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
url_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122210130010012-1213210132311320-2201000133010213-0110223333233221-2333013211312303-0221130301030101-0013210110330102-1031231221331013"></a>

### Direct properties for `rule_list.rules.url_category_list`

<a id="canonical-2210032102001010-0111200021312310-2023030213021110-0033210311202033-2201330003033213-1100110230031303-3101132231122103-3330303323110030"></a>

#### `rule_list.rules.url_category_list.url_categories` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0022013123220103-0231302233112311-3331201312000023-3031022310310321-1300302233001110-2213110210212201-2331202121302302-0031230221002103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_forward_proxy_policy](../resources/forward_proxy_policy.md#canonical-3220033033101011-1003121322230233-1123133312201011-1101102202230100-3032211333331231-0001332310200021-3213230001300202-0033011011100123)
- [Property reference](resources--forward_proxy_policy--reference--group-001.md#canonical-1020112000212333-0003201022122210-3023133121003312-3230311112121002-1120320133012302-1112010113010113-1313022210231211-1033311300001112)
- timeouts

<a id="canonical-0312333203320031-2001323121301301-2113320111022333-1231122110212013-2112302103320203-2103010000220022-2200022121212202-0311001302331110"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112102332233122-3323131313012212-0220312310112012-2132313103010031-1300222202211313-0310301022312033-2212031312331332-0133111133012322"></a>

### Direct properties for `timeouts`

<a id="canonical-2213123131121111-2003023001310022-3233110121302133-3332021310311333-1223232223222031-2220332203230002-1030331232011310-2210102102320223"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0300322100323133-0201121113223010-2110130301231110-1203030033211001-1311113313112223-2330232130231302-3310211203102123-3031013202333322"></a>

<a id="canonical-3221230221333110-2002221133233211-3202021321132233-2202102323100110-3132203203103100-3230022222013110-3123030112332112-2032201311210322"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1303112313021102-3120111033121101-3203312232201202-3312120101300222-0022111202312021-2102001213232013-3022001023320321-3213031122111211"></a>

<a id="canonical-1320002123133032-3031320000333023-0300020113103120-2230300012210031-3031210113323112-0103013312233312-3332213212102023-3313222120323233"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1133330220013103-1010331003022300-2013122001033321-3033022332310323-0010301233300202-0112133323331020-1231212230120213-3301131233111100"></a>

<a id="canonical-2232002321230020-1320323313123130-1320312013202130-1223202122212233-3200302330230120-1321330201022311-1321203120002211-3122012103111032"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
