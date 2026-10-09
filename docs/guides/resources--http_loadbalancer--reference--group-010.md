---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3201313120222312-3232003221302110-0311130020311312-0301122020312122-3101232211013300-2031210131210003-3200321013030133-0011010102222102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-0113332230032013-2132020102112333-0000330033003122-1201230011312323-3230013130322012-1222202113231122-1311330130123322-1303302233213011"></a>

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
any_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022011202103021-0333123310221021-2132010333112323-3312330330102330-3013211101220331-1222231010222202-2200120320010011-0210121103221302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-3033220100120313-1021031312011213-3132233223122020-1220231332221022-1001301131003111-1203111020100123-2131230203133231-0001313213220130"></a>

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
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133033031101131-3302230123123022-1122323011022101-0320331123222100-0301010130230202-3013302033301123-3330213033221301-0203302213331121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-0320200130002221-3232133203102002-1012020302001133-1333331003232230-0110100103332022-2330320001203010-3020023220211123-3312001310221220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232321231030302-0300011322223102-2330033130000221-0101333330123131-0133303020223223-2210313203030020-3331321230312230-0330333232220230"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_list`

<a id="canonical-2122323111300010-1312122213330010-3300301111320010-1000100002313330-1201302113302133-0200312333303313-1020320330220310-2200220111222020"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2000002131130203-0313323223312023-2132220302130203-3322322103201031-1223203023121310-2021010311321311-0130130232033020-2200211201222213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-0133300223222310-3001330111103003-2202313203133321-1332000101230212-3101102232021031-3220121021201032-1112033233233122-2202132031131131"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000011323330333-1121210113221031-2222223223221101-1032212333300023-1333333032202000-0322130220212330-3022023110131303-2302222332012321"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_matcher`

- [asn_sets](resources--http_loadbalancer--reference--group-010.md#canonical-3222322102130103-0201110333302231-0120331323202112-3203033200323230-2003032231301100-3103302210231300-0200020130200120-0223211300331131): complete subsection reference.

<a id="canonical-3222322102130103-0201110333302231-0120331323202112-3203033200323230-2003032231301100-3103302210231300-0200020130200120-0223211300331131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-2000002131130203-0313323223312023-2132220302130203-3322322103201031-1223203023121310-2021010311321311-0130130232033020-2200211201222213)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-1230331022312210-0322033313310333-0310022113022113-3222023131210233-0022323332330200-1303320233330010-2300221321301212-3021020232013302"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3011320231333032-0322001012130200-3113121331213133-0022213310201132-2130320033012030-3103200013303131-3302303031131322-2020111321233221"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-1312212220211210-2331111212222132-0201201003111211-1322312000131232-0332230131013301-2331033231023012-2023121231012201-1031232001133332"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1131123223322103-0230331321222133-2101213231123120-3233001012320012-0221032020203101-2011323011300003-3032323121203321-2131330110002122"></a>

<a id="canonical-1321133330232221-0331131131021212-3033121200333221-2302322020131131-2020031203230210-3222201021331233-0123111220321033-3310233332220003"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2310232113002202-2120301332133300-0020303002231010-2130230020021000-1311012102110110-0220300333302331-3321121021320101-2033331020322103"></a>

<a id="canonical-2012320300023101-1021101100300122-2133220113203011-0203100311011201-2001020013010020-1321002020112101-3033023003210023-1123020010021301"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2333302222000021-0231032110002113-1102013023100000-0023021221031310-0323000033122333-0120132020013130-3103302231012013-2011113003112300"></a>

<a id="canonical-3331023322211023-0323322231222001-0301123113033333-1012233002221322-0201122002300310-0121033133331211-0020013331031010-3311103011001201"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3313113223130113-3121020310333021-0322110302112030-0323203123222233-1321202113222010-2222201231232123-1023211111301301-3220202221123013"></a>

<a id="canonical-3300202002212300-2123100002302102-2011031032031120-2022303022010120-2131311213101310-1220103300131133-2301110322111313-3023320212230303"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3221232222331013-2013333101232303-1211121213320301-0100113132313033-2331203203133003-2031021202220201-1321001103020332-3123212023212212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-3232230102200303-3123112230222022-1130013331131021-3210102111103232-0021031112300232-0100303303233310-1210132012003233-2011202330320132"></a>

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

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2322123231023210-1321011332213300-0002011320313012-0002110113003200-0000120132122012-1100212313113201-3121032131113220-2123302200002120"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.client_selector`

<a id="canonical-1030130202023322-0133220102202222-0220221223023321-2333223131220231-0332123100301103-2001322201032123-1201023111321001-2032201213123303"></a>

#### `api_rate_limit.server_url_rules.client_matcher.client_selector.expressions` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1201312112023220-3110102133111300-3200210011301312-1102221323231012-3201330200312020-0221113222213323-0212313312200010-1101310120100232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-3111212012331101-2220200112330133-2012303132011311-1222132200202101-1013223222220023-0031312030321003-0101320132032022-0111000212330212"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032322121131322-3302103111100123-3120030131100202-0103200232100232-3131313022111322-3330103203123000-2221013202030110-0212033203323033"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher`

<a id="canonical-0221222320013222-3110220123202203-1212023310030212-3133331231101111-3101013201222001-1120310200301123-2111000022311123-3222222132110312"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](resources--http_loadbalancer--reference--group-010.md#canonical-1111220221122301-1130033211133132-2103221030323230-1001320030112000-3032323021002323-0003122231221330-3000302230021030-3032112232201030): complete subsection reference.

<a id="canonical-1111220221122301-1130033211133132-2103221030323230-1001320030112000-3032323021002323-0003122231221330-3000302230021030-3032112232201030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-1201312112023220-3110102133111300-3200210011301312-1102221323231012-3201330200312020-0221113222213323-0212313312200010-1101310120100232)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-3011123330223303-0302302202320101-1023202031313100-0312132113311323-0330333030121113-2333130210002230-1101130122301220-1232233030202022"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0121000101331203-0131123312032203-2102201113313023-0022110333311010-0323233010031032-2031130103312123-1310300011002231-2023110202212012"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-0012121313331232-1300010123121220-2231120121223123-2302000213120323-0101112202221232-2332013310030320-2130110123102233-1020322233121220"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3331201001110013-3002200012002313-2213010030211232-1322310321303213-1100311233100121-3103130210210100-1121033300002123-2102123110311012"></a>

<a id="canonical-0101221021002103-0001013112301113-0010133120212101-2302111331020012-2012311313231221-2222301002020331-0020132303213213-1102220332030201"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3323313213312033-1002303132312030-0201231001111203-2311012333100110-1030223301113331-0333231101113133-1132210203121030-1003032133000231"></a>

<a id="canonical-2100220100301221-2011201321030010-3023232013222023-3011012203310032-1213332213103101-2231313220020332-1303233222102103-0210031113303003"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1032001322113122-3303111022131033-3221112002332213-1123213132211313-3222022323020032-2130321010121122-0130010201201230-0032001330113112"></a>

<a id="canonical-3303323222213301-1033122120133303-0320023002111032-0332101223320021-2102020011330120-2331103023200133-3131202313021121-0022023110120003"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0203123320301120-0103103103012013-3112131320213021-3221011133210131-3333323333222322-0333301101321332-2233331231322212-1130333311100112"></a>

<a id="canonical-1311122232303323-0303200001033020-3101300202110011-3102102200231220-1232131233200033-3212103331320100-0032210030103220-3012321311321231"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2203121131301233-2203201100322000-2220212301120232-2023030221110012-2001211200121321-1311112302313220-1333303302223032-3033032303332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-3023122011132231-2303000011023013-2311321220323022-3103122213212101-2021301031003213-0122333213301311-2012003121001202-1123210003123220"></a>

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

<a id="canonical-3332012031332212-3030233312221321-1323031000220002-3313320321220303-3130212303101132-2002311203020203-1002021311101032-0302223111022013"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list`

<a id="canonical-1103023001203200-0231123000031001-0323130230222031-3021323012201231-1221011203223221-1010020220203222-3321130022131313-0221332312123003"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-3331332102210030-1200100321010211-2210332202220101-2311011212332013-3330210112130232-3220203022110102-3033231210213312-3303031032122001"></a>

<a id="canonical-1302023211303330-0032300230321223-1121312232023221-0210022113001112-2331002223233302-1011030033012210-2223323113231130-1111000320010331"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1231330211322320-1122320030111011-1330301312012202-2310122321121110-0212323111010212-3313202200230031-0222223021333031-2130231112220022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-1300113312000132-0101322103330301-1102000231030003-1022001211030211-3332113032100303-1311330112310013-2010113001021300-3103022002020330"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310003031232313-3231303112323000-3300220122223001-1022310131313012-0332000020030320-0322331003131210-2113033310220320-2331030210012011"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list`

<a id="canonical-0022330011031210-2212330012212230-0322000330131301-1331201102133223-2233131200310020-0001323103333010-1000203222100212-1300330023011011"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2111102223312220-3103020330310103-3100313013131110-0322313121022222-0030312201003130-0023132031031331-1201122310213220-3210012002132112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1302113200032023-0310000120211211-3202002030113213-3323222032222031-2012001310021102-2210022030013010-3311133201210033-2133100121200210"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312213303122012-3031000210201033-1102100302011122-2022133330011131-3333333220033302-3123122131330112-2133021213122133-0302302203212311"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-3100012232313232-2211133211020333-1200020222202210-1033300020300213-0032230321201101-0310010131022221-0310122323103213-1110023132233002"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0222313131000022-2221120112303230-3230121022013332-2212220031101230-1221111210211221-1220020211211101-2212222033120021-3230002031011301"></a>

<a id="canonical-3021013001202121-0021013213301130-3302301021101200-3123230220321022-0132103131232323-0031011132200212-2011331303011020-2220220232033133"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0131112101211123-3130302033220120-2322002312003322-3133203201133322-1331102123212020-3201021301220113-0131200211120222-0310220022200310"></a>

<a id="canonical-3130130200013031-1212303033021112-0203313120011130-0220312100333030-0011120203031032-2020202320002122-2112120203123313-1021001203313320"></a>

#### `api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-1313000212311313-1331312031123311-3011103110321300-3001311203110301-0021012110133321-3032013102022222-3212102320003202-1131330031221133"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001033333231212-0022330112101212-1302110131212132-0102132200210120-0032211232313221-0203111003220000-2101202313233323-2333201010221300"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter`

- [ref_user_id](resources--http_loadbalancer--reference--group-010.md#canonical-3122110102021131-0210011132022301-2023111000333232-2031133321123202-0211131222132230-3013231033012330-0331030100133220-0330233330001203): complete subsection reference.

<a id="canonical-2122101101102120-3010332123211200-0332120012001022-2123300212313133-2013233021223022-0030300101001223-2023201313103003-0033200000211302"></a>

<a id="canonical-1301011110030100-0323320230201111-1122210311122133-2101230013221001-0133012223300211-0011013301210232-3332020200132310-2223322120121031"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.threshold` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2030211122303320-2101132213201232-0013311001223210-2013002320202231-1300311110122313-1120003210310113-0003232001232300-3121202322203103"></a>

<a id="canonical-3212212233202330-3020312121312300-3330211112033003-0121233220003230-3310303002320310-3103320131030210-0333031020102302-0001103202112103"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.unit` property

Type: `"string"`. Optional.

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

- [use_http_lb_user_id](resources--http_loadbalancer--reference--group-010.md#canonical-0111112213112131-0132010230220112-2313023320220012-2321101131120111-3313220011300212-1010223212133203-3012222122321311-2331010313001322): complete subsection reference.

<a id="canonical-3122110102021131-0210011132022301-2023111000333232-2031133321123202-0211131222132230-3013231033012330-0331030100133220-0330233330001203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-010.md#canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-0312321121332201-3113300202001302-3222322030131120-1213033113110320-3001121022122211-1032333330302202-3032130333220323-0003023001120332"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311333222131121-0302230321200102-2201233121311030-2233021120030020-2322012003232333-2002103131302002-2022303101002003-0213200032130211"></a>

### Direct properties for `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-3021002031022030-2031303233320112-1013032001120230-3131221232301001-2032310231322012-0031130332120332-2210020301120311-1200102032033332"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1202103030033230-3013100312012120-2120112330021110-0230000023021331-1120300333311200-0022010010331002-3031013312310001-2203011120032130"></a>

<a id="canonical-2322112331221333-0232012223102310-3132031030122111-3220213112321312-0110031220331330-1023333010303233-1010231320303030-2001102312033310"></a>

#### `api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2330331132022020-2122310211300012-1030000302220332-2011010301101020-0211312121123013-0011100303322303-2133320202011232-3032213021332120"></a>

<a id="canonical-2121212012101002-1303122213213002-2131333223221301-3303233022112322-2233223300030132-0232120322010102-2222223012023320-1310302133010010"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0111112213112131-0132010230220112-2313023320220012-2321101131120111-3313220011300212-1010223212133203-3012222122321311-2331010313001322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-010.md#canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-2321003232233303-0032003201123232-1223011301311110-2232231301312220-1133230212302020-1120312313312200-1112311333202303-1113021212223001"></a>

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
use_http_lb_user_id = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003023331222033-2221012320323333-0303300312100321-0001000033001020-3133220003001112-2131311032211230-1003011113023110-0101320100203112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-1330031203223123-3121121020100021-1113012332021303-0121331213121122-2320211100220000-0122130101021303-3300023020221223-0001002303233203"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300013011223122-0300320310133221-0111102233032322-3113202033131201-3312030133330232-3000220303223113-3303213202222202-3333331132013012"></a>

### Direct properties for `api_rate_limit.server_url_rules.ref_rate_limiter`

<a id="canonical-1310012201311023-0201332130332211-2211011303100323-2021212021311023-1133322120000201-3013023012100100-1123112131003310-2130223123011331"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0300103000212111-2001320101200213-0321013210111113-2312303311300012-1112212023100112-1110023321203031-3113111210013311-0021212212331001"></a>

<a id="canonical-2212210330200323-1231001321213201-0210123200031010-3133201331010120-3110112321313210-3300003103222221-1203220003212032-0333231030210332"></a>

#### `api_rate_limit.server_url_rules.ref_rate_limiter.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1211130120212002-2302203130300323-0103001320133332-2121031311132003-1201312100322313-1033312001211223-2322000021331000-3221013001122213"></a>

<a id="canonical-2221301230011110-1010131002312320-2011311201232321-1322120321211301-1223100001331122-3302310223230203-3013100002230202-0231121110223013"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-1113130322101231-2331120311211001-1030011233331110-1231131322213301-2013021322301312-1213232230202310-0122321003212010-0130330221223301"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312011021331002-0022203030222121-3010132210103331-3220322121311030-3111132200231031-0102132301013323-3211312102201301-2001112120013031"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher`

- [cookie_matchers](resources--http_loadbalancer--reference--group-010.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-010.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-010.md#canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-010.md#canonical-3132321313100031-2233313300131223-0233020023131212-1210213032330332-3320001101033110-1331311133030131-1220323302313300-0001103011302011): complete subsection reference.

<a id="canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-3133312200321003-0312332230100021-1223101200313030-2300010323313330-3030212021203021-0021033123112312-3112222303011000-3030210112313320"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1312131010113303-0132320133203210-0002031030103223-2311332120132203-0122320122030120-2122231101231230-1330223011112303-2300321302123210"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-010.md#canonical-2202121022331120-2021033031311003-2201002000201210-3113213321210313-2330120211331220-1302132220300001-3230021023012033-1220211313332010): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-010.md#canonical-1103012203322132-3332330132312221-3003020232323030-3010222032121131-2012301020000312-2230020212030331-2033011030222122-3330231013010230): complete subsection reference.

<a id="canonical-1303301023020022-3330102033230130-3203033103230002-2013010223333202-2133202202000210-2001231113213322-1001300002021210-0200300000023020"></a>

<a id="canonical-2333232233331121-0122023331000122-3320210222103003-0102022133123121-3313300121033232-0303330003221320-3222322311011222-3302233212001122"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-010.md#canonical-0233132330222300-0322203213003123-1110332230220311-2001222013301113-0113101031000130-2020223223231233-3121331032013112-3333332311132003): complete subsection reference.

<a id="canonical-1022103000322023-1231231223300022-2022023210002031-1310023320023333-1201232203211000-0131031012113030-0310121103202331-1021320113130302"></a>

<a id="canonical-1031021121302222-2023011001021023-0031230212132113-2320020323202333-2132001001222031-3321203301133320-1021122012313120-3310232010011030"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2202121022331120-2021033031311003-2201002000201210-3113213321210313-2330120211331220-1302132220300001-3230021023012033-1220211313332010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-010.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-3232131220012232-3032000200111030-2211230300030212-3123302200313202-3131003200202301-3303320300020232-3112123003202221-0321330103100020"></a>

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

<a id="canonical-1103012203322132-3332330132312221-3003020232323030-3010222032121131-2012301020000312-2230020212030331-2033011030222122-3330231013010230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-010.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3030031032203012-0202211030221232-0011100001330331-0131031332113201-1233332332111113-3110113103020231-0331231101011011-1003330023212011"></a>

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

<a id="canonical-0233132330222300-0322203213003123-1110332230220311-2001222013301113-0113101031000130-2020223223231233-3121331032013112-3333332311132003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-010.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-3311221321023323-3102310303101333-1323020331231111-0120233002132100-3130200031212220-0010202333130131-1203120112130320-3333223012212010"></a>

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

<a id="canonical-2132000212312223-3212232123202332-1313033130133033-2123331202323302-2100130233200103-2130013023132013-1233002112201103-1233033003313000"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item`

<a id="canonical-3101113213333002-1210002131212311-0130131211313233-1221110202231111-2200132111113332-0100322133121201-2232011130001133-1132002301031320"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0032100323230130-0030031020103202-0231333320220330-3122030103333210-2023010122032001-0332131230321001-3032032110110221-3111021200232220"></a>

<a id="canonical-1033113103033101-2222121011313113-0230131220231202-0030032001012031-3223122132212103-3303132301101200-0131300102303320-1012113010001311"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1302222113313131-3230012313112003-0101123100002011-2310322231032333-0013301321133310-3202001321233111-2320132113302032-3302101303211333"></a>

<a id="canonical-3312303202211310-0200132330211130-2321212000232021-3133030123200303-2330131213022001-1103231222101330-3003321330130020-1103331201303312"></a>

#### `api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-2331332131320001-3221232321332120-0302023121000121-2101001302102210-0313232113101231-3030111310231000-3101203101032120-3101300013023210"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2302111001100221-3231223313112012-3022203311111021-1203131232000320-0232232123321321-1221230001110031-3320301203010121-3310233320332231"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-010.md#canonical-1003032032032203-3312021311031000-2133021323020023-2222002033001332-2030331122213212-2030222230021231-3332011312110303-3230331132003103): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-010.md#canonical-1002220323233121-3113030212220310-0132310320312122-1021323123032112-1222202130201301-3001012131022113-0222331333230300-0303300110210303): complete subsection reference.

<a id="canonical-0121122020321032-2200200100101002-1002131033330033-0223021300132323-0211222331331322-1333000221333111-0221122202300013-3230100222122111"></a>

<a id="canonical-2103112302302132-2310111100230321-2032112302110201-2002002100311330-1231121303320232-3230120110032323-2231212110220301-2310030213102032"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-010.md#canonical-3101031232033010-3033203020010231-1313301131133032-0301312311312211-0311113321321321-3213333333130130-2311300320121103-0332011012130010): complete subsection reference.

<a id="canonical-0023023131112020-1120112011232333-2321132223221301-2023231001321310-3231213223313200-0103120133331101-2133331333033213-2300211323113002"></a>

<a id="canonical-0212000330311002-0112011023130131-0031223113001032-1311322312112211-2022101020201021-0310133230313132-0313302000102021-1212011302321002"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1003032032032203-3312021311031000-2133021323020023-2222002033001332-2030331122213212-2030222230021231-3332011312110303-3230331132003103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-010.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-0133332013231331-0213232131333002-3111113123031100-0323200120023013-1022320112312122-3333030233320002-0233120022310202-2011123201120112"></a>

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

<a id="canonical-1002220323233121-3113030212220310-0132310320312122-1021323123032112-1222202130201301-3001012131022113-0222331333230300-0303300110210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-010.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-2021210303201032-2223313010010001-3323202301000333-2321110112102111-2132011111133130-1011212321020022-1313003201100031-0033332221300100"></a>

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

<a id="canonical-3101031232033010-3033203020010231-1313301131133032-0301312311312211-0311113321321321-3213333333130130-2311300320121103-0332011012130010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-010.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-2230212202003002-1000331030232302-2113023323013203-1013010102313322-1223120322012310-0231232323030022-2232102001130222-1123230200331001"></a>

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

<a id="canonical-2220203212002311-1331020300122102-1010212230201022-0230113010232210-0103132303021122-2001130302131111-3030032231033111-2330121302301333"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.headers.item`

<a id="canonical-1210021013011130-2022020310322033-3033102111013002-2032122332010113-0232331111333030-0032021212010022-3013020021211320-3232103201311101"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1111213222012212-0132320211211101-0221230130311122-1233310123322301-2201031300213301-1132023221033310-3300210100120033-0110022001113032"></a>

<a id="canonical-0210300033001112-2103221031232230-3020312223232010-0130011111110113-3301211200031022-0022232131032031-1322221103330213-3312223233020233"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0203203032333323-0203120202012213-1232122032320100-3013231111110103-2202331310332301-1132123323232321-3203203132101201-3312112302302111"></a>

<a id="canonical-2022023202010233-0121301222303210-2212000100112031-1111303010333003-2203133233203033-0003012033112130-0012121103010313-3302231300032320"></a>

#### `api_rate_limit.server_url_rules.request_matcher.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-2122202123120102-0330310022311130-0010321120131103-1200312123023100-0112300133321012-0102113203002022-1313203213021003-1002200020100112"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2221211230332001-2132012232220123-2320213021210132-0230312010020300-1122200330300031-0122221110112320-3203113203132310-1002032301132001"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims`

- [check_not_present](resources--http_loadbalancer--reference--group-010.md#canonical-3110212003110301-0012213201332323-2231013003321122-3222231313211112-3123231322300031-3033313331110022-3011033203220223-1322230321101012): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-010.md#canonical-2111123000110232-0111130230112022-1302021303100330-0102030300233330-0230202021202330-1130333223323133-2112213212031212-1333112300103233): complete subsection reference.

<a id="canonical-1211011020203222-3102202230120200-0332210000002113-2313320310201211-1233022013020110-3311110301100202-1212120301111123-0201112213220100"></a>

<a id="canonical-3203022122011310-0323310032112023-0000021212221032-1100300020233301-1030003023222312-2300213301321021-1201302033202112-3313020023202112"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-010.md#canonical-2111202232213110-2023210223012321-3201222211232212-0122021020133110-2033032012020003-1320232321213322-1330131003233010-0000311123130312): complete subsection reference.

<a id="canonical-1310131200301010-3023303003301133-0000101000113021-1331020231123031-3313311111032233-1320123002201321-2112332202322132-1110312013120021"></a>

<a id="canonical-3130331031121130-1002331222302303-0303210032121001-1012020310112031-2111320300032020-0311303102300311-1233123310312033-3133223313000222"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3110212003110301-0012213201332323-2231013003321122-3222231313211112-3123231322300031-3033313331110022-3011033203220223-1322230321101012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-010.md#canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-1302130020323013-3311031003310020-0002202003112033-3311020103210013-1230310013233031-1133201120020232-2111123023103002-0302033200012301"></a>

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

<a id="canonical-2111123000110232-0111130230112022-1302021303100330-0102030300233330-0230202021202330-1130333223323133-2112213212031212-1333112300103233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-010.md#canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present

<a id="canonical-0130332033300132-0201111311123220-0220331203233321-3212120101222313-3323103113113233-2131010232101233-0011213022111321-2022303310030201"></a>

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

<a id="canonical-2111202232213110-2023210223012321-3201222211232212-0122021020133110-2033032012020003-1320232321213322-1330131003233010-0000311123130312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-010.md#canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.item

<a id="canonical-0302130032300333-2313000020022322-3031033123133330-3010301330333003-3223113123221200-2031100013203002-1320320122213212-2230131123321013"></a>

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

<a id="canonical-1311320201121113-3022332131330231-0231303003331101-1123000012001102-1020322223211010-3233230102220031-0103203312133333-0022022123201120"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item`

<a id="canonical-1321213121302203-0100312300130312-2122121012230210-2131233003200303-1230131003222101-3010130010310132-1202210223203130-3220021113131120"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0012300021202111-2013211011221222-1000130332303211-3123332110101033-1301023300333100-3030213100321030-0010000211310230-3110130312121323"></a>

<a id="canonical-2130312303320321-3223210031111221-0223221231032021-2132213231332101-0121310332123113-1123300113200232-0102323323310010-2000023212332201"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0003003232232300-1030322100212230-2032220110003320-0312313203230320-3331233333301321-1023111201201020-3022233032010330-2332333011101323"></a>

<a id="canonical-2300223233010132-1310210210111103-2230113203221232-1000311030203112-2123030131030322-0113222301332003-1320031123331212-1023332313011032"></a>

#### `api_rate_limit.server_url_rules.request_matcher.jwt_claims.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3132321313100031-2233313300131223-0233020023131212-1210213032330332-3320001101033110-1331311133030131-1220323302313300-0001103011302011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- api_rate_limit.server_url_rules.request_matcher.query_params

<a id="canonical-0013210100201310-3220313111102321-1310010332030013-0032213012221213-1133031210000121-2302222022023132-3231110102033120-1203123303101212"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1120020001301313-0011301031020322-1200000233012032-1203111121012113-3222322203312102-1103233302001121-2313333102031112-3313322303200131"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-010.md#canonical-3021120003000002-2122133132200033-2111100220122321-1130122312232121-1123322213333012-2222213213222122-0331233023132103-2030020303303221): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-010.md#canonical-2320113003330333-1023300333020000-3022102001232222-1312203332231230-3103213320102031-0122333320200333-1101001323103232-0033203302122112): complete subsection reference.

<a id="canonical-0332102132310122-0302031300223332-3030122222102101-0333220010312032-0213200000220312-0232021333203221-3100111100212222-3000013121331100"></a>

<a id="canonical-3002130023102221-0221003132322122-1122332230102130-1222120220220331-1223330311020133-2210230232100022-0310320332030201-3001223113211212"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-010.md#canonical-1301223110101123-3202011002021332-0211003100012100-3100330102311113-2223133333222122-0323210333101312-3103202001103333-1002211103220003): complete subsection reference.

<a id="canonical-0312102113031323-0203232310232010-0021130013123032-1331302113321320-3020023113102033-1223331013121121-0100011131112232-2310122211002123"></a>

<a id="canonical-3123123031313013-1101220203013123-2211303331301331-0001202032032331-3230121230000213-3132003331023220-2303302210113221-1133010311122331"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.key` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3021120003000002-2122133132200033-2111100220122321-1130122312232121-1123322213333012-2222213213222122-0331233023132103-2030020303303221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-010.md#canonical-3132321313100031-2233313300131223-0233020023131212-1210213032330332-3320001101033110-1331311133030131-1220323302313300-0001103011302011)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_not_present

<a id="canonical-3303123123322302-2013101223000120-3300302022211300-2323203031113010-3301212321310101-0030201121123221-1202013020032101-1021031010320112"></a>

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

<a id="canonical-2320113003330333-1023300333020000-3022102001232222-1312203332231230-3103213320102031-0122333320200333-1101001323103232-0033203302122112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-010.md#canonical-3132321313100031-2233313300131223-0233020023131212-1210213032330332-3320001101033110-1331311133030131-1220323302313300-0001103011302011)
- api_rate_limit.server_url_rules.request_matcher.query_params.check_present

<a id="canonical-2200320123023011-0300210232223132-2021020020313213-0000212303130233-0322022311231332-0222303303001013-3100132002013222-2222000232321033"></a>

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

<a id="canonical-1301223110101123-3202011002021332-0211003100012100-3100330102311113-2223133333222122-0323210333101312-3103202001103333-1002211103220003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-010.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-010.md#canonical-3132321313100031-2233313300131223-0233020023131212-1210213032330332-3320001101033110-1331311133030131-1220323302313300-0001103011302011)
- api_rate_limit.server_url_rules.request_matcher.query_params.item

<a id="canonical-1021303301232231-0133120023320231-2131233022011302-2113301000001232-0022122311120100-3023220002022333-2133220112223302-0003311303013203"></a>

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

<a id="canonical-1111202012022013-0010020032300221-3000202231200212-2333222203322211-2303321313233100-1322113321133330-1031213120212030-2021323230220120"></a>

### Direct properties for `api_rate_limit.server_url_rules.request_matcher.query_params.item`

<a id="canonical-3032012322211103-2330310312131330-3100000231122212-2012101101301211-1100210320331020-3213333323320300-0312112031130330-1121301110210300"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0213121313032033-2013333201331133-3213221021233230-0300101132003131-3323213110131311-1301102010310131-3203111323301022-3123221210000301"></a>

<a id="canonical-2013333313031203-3202102201101322-2103331011100010-1013011330032020-3000331133111223-1021101211122322-2030021321302230-2331312003210032"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2200222310313322-2103132331032303-3120000031100003-1011301112213002-1310201301120021-1003032011013233-2203333101102310-1300122202332232"></a>

<a id="canonical-1300030031101300-0310223232313012-0110221331123022-0333122031222302-0312303033203112-1032222103203132-0110112132123113-1122103223230332"></a>

#### `api_rate_limit.server_url_rules.request_matcher.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- api_specification

<a id="canonical-3012102112110133-1100013223322032-0331303111030231-1123212233101113-1333000321021110-3312102323020331-3120322310022201-1313013020100331"></a>

Type: `"object"`. single nested block, Optional.

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

- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3012102112110133-1100013223322032-0331303111030231-1123212233101113-1333000321021110-3312102323020331-3120322310022201-1313013020100331)
- [disable_api_definition](resources--http_loadbalancer--reference--group-017.md#canonical-0121212322222130-3103021101033321-1123000003031210-3103120000300132-1002112302102003-3020032021001211-1330203332020133-1300330021233313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132332003133333-3220031011313321-1220133230021020-2201231102300223-0323202322203211-0021232233012000-1330222331132133-1211233030333323"></a>

### Direct properties for `api_specification`

- [api_definition](resources--http_loadbalancer--reference--group-010.md#canonical-0003200003233123-2233111330123311-0220130030320331-0122321232003200-0211131112101211-2212201300231300-1101231020201330-3120302320032103): complete subsection reference.

- [validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130): complete subsection reference.

- [validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210): complete subsection reference.

- [validation_disabled](resources--http_loadbalancer--reference--group-011.md#canonical-0313322031101011-2030333022323320-1122212003130032-1133211032021000-3320233120012101-2233031012300132-0112202012322000-0031131021323032): complete subsection reference.

<a id="canonical-0003200003233123-2233111330123311-0220130030320331-0122321232003200-0211131112101211-2212201300231300-1101231020201330-3120302320032103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.api_definition` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- api_specification.api_definition

<a id="canonical-3001011331010122-2011203132332022-3123312213320200-0322203321301313-2121033021023031-3031200033002331-2000120230120010-3331230310033222"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
api_definition {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112102232312312-2000133332130300-3102201232020212-2220310223132010-2320330220201020-1230330023212301-1211011100321300-3310203133113013"></a>

### Direct properties for `api_specification.api_definition`

<a id="canonical-1031010200203201-0111300000321200-0131232030021031-0011320232131320-1032210321322123-2203002302002223-2322020130131001-2103020013220200"></a>

#### `api_specification.api_definition.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1310323102021021-3233332313011201-0103002021032320-2321201133330321-1123323320013103-1222232120023011-1123122121300202-2210332322113111"></a>

<a id="canonical-3223202002113220-3200000132300200-1203330231330233-3301230031021100-1211333313202133-3333112311012233-1210210233030302-1332223101301002"></a>

#### `api_specification.api_definition.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0133001123313112-3120012220322032-1030013100120010-0300303122113232-3222100011030023-1123010020220101-1121320213102133-2202132030222123"></a>

<a id="canonical-1301301202233222-1323232230212033-2311210031113223-0302230300312120-2101221012102200-3122202013303203-2032131011133120-3131020003223121"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- api_specification.validation_all_spec_endpoints

<a id="canonical-2133011132032101-3123211032100022-3211221120202100-3230103200023113-3133302232001223-3120301033120321-2321121010331031-2223323312001000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
validation_all_spec_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000302023131132-2231010301102203-1221001200010033-3233312030223012-3122112011323332-1301032123002210-2111133130001222-3332312121103013"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints`

- [fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032): complete subsection reference.

- [settings](resources--http_loadbalancer--reference--group-010.md#canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201): complete subsection reference.

- [validation_mode](resources--http_loadbalancer--reference--group-010.md#canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221): complete subsection reference.

<a id="canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-0322301020132010-3321032020323110-2210221200303111-2000010300121023-1003102013133310-3003031210231211-0200323312320130-3322200130320210"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010312303320323-2102123000223303-2011023221331133-0200312233320003-2010030023200130-2322333000133203-0020102120000031-1303031222020123"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode`

- [fall_through_mode_allow](resources--http_loadbalancer--reference--group-010.md#canonical-1001103200333100-1000033310223331-1103201132020320-1203011220013311-0100331220333111-0222320031233020-0121030110003232-2210022110302200): complete subsection reference.

- [fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-1122003301202012-3102002213101231-0321023122030331-2123231010001221-0213330100110323-3311230023031332-0300313322023002-0303230132321303): complete subsection reference.

<a id="canonical-1001103200333100-1000033310223331-1103201132020320-1203011220013311-0100331220333111-0222320031233020-0121030110003232-2210022110302200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-2033221300332311-0012320101302230-3131010101231203-2313231021230113-2233311001333310-1303323201020030-0213122203102103-0120132030130113"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
fall_through_mode_allow = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122003301202012-3102002213101231-0321023122030331-2123231010001221-0213330100110323-3311230023031332-0300313322023002-0303230132321303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-2101113022121033-3100020312123021-3312220210122332-1022102330302222-0322120031102133-0303121023312322-1022232200023030-2333102331000102"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202203010030310-1223002222123213-2311211303300302-3203033100302030-0132030101233333-0302230111012010-2033311230331112-1121013033100302"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom`

- [open_api_validation_rules](resources--http_loadbalancer--reference--group-010.md#canonical-2031101121132123-2231131132002010-2200021020303102-2030300323002103-2120310221133330-1330111012212103-2113000210332003-2210130300102300): complete subsection reference.

<a id="canonical-2031101121132123-2231131132002010-2200021020303102-2030300323002103-2120310221133330-1330111012212103-2113000210332003-2210130300102300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-1122003301202012-3102002213101231-0321023122030331-2123231010001221-0213330100110323-3311230023031332-0300313322023002-0303230132321303)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-0310132303031333-0133133101330033-0133222110000121-1022303223003122-1310123203021231-3020310232010112-3202230113022131-2220003310010321"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100201202201132-3120301230033023-2012231302200101-3033313100201212-2220230301232033-0302233321102331-2113033201123110-2133132121232200"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`

- [action_block](resources--http_loadbalancer--reference--group-010.md#canonical-3231322231130233-3222130232102123-1120123110112003-3321103323322213-3223102211122320-2010311221131011-0010212133113322-0103013130022110): complete subsection reference.

- [action_report](resources--http_loadbalancer--reference--group-010.md#canonical-0302333312110201-0110300010202113-3022111031130131-1202300103211311-1201032030032010-3011100111101301-0303231200010012-0232302031102301): complete subsection reference.

- [action_skip](resources--http_loadbalancer--reference--group-010.md#canonical-0131212003302220-2200201323110330-0023221301033121-1320020311320112-3010232220333303-0230210002032310-2231001220023232-1302021021230033): complete subsection reference.

- [api_endpoint](resources--http_loadbalancer--reference--group-010.md#canonical-2112011120332211-1233212201331030-1003312111300230-1202031201012121-1231023332313321-3123222110320023-3230323333102013-3111222022211111): complete subsection reference.

<a id="canonical-1111323210213200-3201303100231333-2012002323323223-1133323021111122-1301222112112010-1331122032132011-2123200230030331-0303133303120103"></a>

<a id="canonical-3023310232120231-0223231300131030-0312302103003332-2213312203013011-1023310302031011-3131120102232231-1223200002102133-1123322230321132"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0110103220210010-1212331330103002-2212310103203031-1330223222030132-0202013213312020-0321302003130230-2303011212003031-1212200202003210"></a>

<a id="canonical-2022113123321022-3230112010123123-3300003012223332-1213330311200112-0201321321213223-2221230010322321-2223210201012220-3322331331330312"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [metadata](resources--http_loadbalancer--reference--group-010.md#canonical-0112011313333020-0203222110213323-2201111003201311-1221003321233310-0231102310110332-1030123300021311-0321113231031311-3303010020020230): complete subsection reference.

<a id="canonical-3231322231130233-3222130232102123-1120123110112003-3321103323322213-3223102211122320-2010311221131011-0010212133113322-0103013130022110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-1122003301202012-3102002213101231-0321023122030331-2123231010001221-0213330100110323-3311230023031332-0300313322023002-0303230132321303)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-010.md#canonical-2031101121132123-2231131132002010-2200021020303102-2030300323002103-2120310221133330-1330111012212103-2113000210332003-2210130300102300)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-0003330303102133-0210121020110023-0310003330130013-1213020123301213-1101013331001000-2101303032120232-2310022032020003-3122300121123222"></a>

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
action_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302333312110201-0110300010202113-3022111031130131-1202300103211311-1201032030032010-3011100111101301-0303231200010012-0232302031102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-1122003301202012-3102002213101231-0321023122030331-2123231010001221-0213330100110323-3311230023031332-0300313322023002-0303230132321303)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-010.md#canonical-2031101121132123-2231131132002010-2200021020303102-2030300323002103-2120310221133330-1330111012212103-2113000210332003-2210130300102300)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-1330022230223201-0302111232101002-1302033332222020-0221122130103003-1131323121033130-1312313011120333-0322233111100232-3200333130121123"></a>

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
action_report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131212003302220-2200201323110330-0023221301033121-1320020311320112-3010232220333303-0230210002032310-2231001220023232-1302021021230033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-1122003301202012-3102002213101231-0321023122030331-2123231010001221-0213330100110323-3311230023031332-0300313322023002-0303230132321303)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-010.md#canonical-2031101121132123-2231131132002010-2200021020303102-2030300323002103-2120310221133330-1330111012212103-2113000210332003-2210130300102300)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-3002221130010210-2110003133203320-0101120112112122-2203202033231230-1310022020113112-2332000111233111-2322123010110322-2002212302313111"></a>

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
action_skip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112011120332211-1233212201331030-1003312111300230-1202031201012121-1231023332313321-3123222110320023-3230323333102013-3111222022211111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-1122003301202012-3102002213101231-0321023122030331-2123231010001221-0213330100110323-3311230023031332-0300313322023002-0303230132321303)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-010.md#canonical-2031101121132123-2231131132002010-2200021020303102-2030300323002103-2120310221133330-1330111012212103-2113000210332003-2210130300102300)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-1202220232220002-1033331010322310-0223110133122230-3111123132210021-0221011021020021-1232203121310000-0023111120022210-0020111200130302"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023323201131031-3301123201001313-2033031022112330-3133031010310311-3012002121332233-0023132000133031-2313023312201231-3031322021032121"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint`

<a id="canonical-1012310120113123-1313010113302311-1312323121030011-0103102313010233-1233220021102222-3023110020300022-1331233133312130-1200332010201112"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1232132133231200-3211113302333323-3021013223130100-1202222303213221-2120132102023213-2312231002312010-0003122002332201-3320012111232203"></a>

<a id="canonical-3222113320023123-2121112233103121-3131111320333311-2221220002203110-3312201130111321-3330311021011101-1100231030020210-1232001321122001"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0112011313333020-0203222110213323-2201111003201311-1221003321233310-0231102310110332-1030123300021311-0321113231031311-3303010020020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3133323000021231-1131033130121102-1113131122323101-2102200120003323-3320322110201323-2303330011203322-3231323030121320-3223313330232032)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-1122003301202012-3102002213101231-0321023122030331-2123231010001221-0213330100110323-3311230023031332-0300313322023002-0303230132321303)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-010.md#canonical-2031101121132123-2231131132002010-2200021020303102-2030300323002103-2120310221133330-1330111012212103-2113000210332003-2210130300102300)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-2130323311233000-1312201012223203-2232023213011120-3111110002331120-3332013212210311-2131123023010323-0100321322320202-0221001021201001"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332101203211311-1210011330222322-3301011131111310-3300222301222220-3311201311303312-1033003201233300-0031003031310322-3033120003133231"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata`

<a id="canonical-2330220203033130-0121003102312000-0110320100330231-2303031330232001-3131212022320011-2212303333130110-2201112302020323-2211132100302131"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-3113201202202012-1303003103333330-1120201033021311-0112113031111230-2100312122032020-3021030332021301-0121031301313313-3213031230333130"></a>

<a id="canonical-3021231123121111-0312230233010233-1203100130031212-3321110120211032-0011200112121332-2300333002312031-0213022202010300-0011330331010033"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-3130120312003022-1322132202333210-3111210223200122-1131230301021303-0310012032310200-3232310131120011-0011002303101020-3020012002220022"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112003131113220-2312203030221310-3120200222231133-3332331122111300-3032010130330101-3231300323030022-3113111331122102-2332102312212203"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings`

- [oversized_body_fail_validation](resources--http_loadbalancer--reference--group-010.md#canonical-3312330333211111-3000121221132132-2301210331233203-0223333303131130-3202321123113002-2230213011120130-1210221132132301-1333302112233220): complete subsection reference.

- [oversized_body_skip_validation](resources--http_loadbalancer--reference--group-010.md#canonical-1031212320313010-3123000102310331-0332121022023310-3320310021301101-0212320132021201-3103200122231220-1333103100320312-0321211120101002): complete subsection reference.

- [property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3312232232221322-1333321212031103-3231103331230300-1233333032313111-2113220211302002-1211312032321312-1113222000020233-3322023020230111): complete subsection reference.

- [property_validation_settings_default](resources--http_loadbalancer--reference--group-010.md#canonical-0220102302111311-1200332310001202-1022311232032302-3010330013211210-1322330110010232-0132010211233002-3211213002300121-1332310011020112): complete subsection reference.

<a id="canonical-3312330333211111-3000121221132132-2301210331233203-0223333303131130-3202321123113002-2230213011120130-1210221132132301-1333302112233220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-010.md#canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-0232003311003012-0120133210121330-1102010202300123-2100100112110133-1102112103121231-0022101020010313-3322320113210113-2310300101123222"></a>

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
oversized_body_fail_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031212320313010-3123000102310331-0332121022023310-3320310021301101-0212320132021201-3103200122231220-1333103100320312-0321211120101002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-010.md#canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-3321000021210022-1021321303110121-1132010221231023-2011332200123103-3323110010131212-1232010301030212-2032122321120123-1101302301012100"></a>

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
oversized_body_skip_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312232232221322-1333321212031103-3231103331230300-1233333032313111-2113220211302002-1211312032321312-1113222000020233-3322023020230111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-010.md#canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-3222103010112211-3111121223311021-0023312213331200-3300132032112031-3013311222331101-1012320000212302-3130211032233000-1333122230033002"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
property_validation_settings_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011331012113000-1011111233203120-0332223021322332-3312301323323232-0332233213323313-3333333121010022-2301303002203201-3330130031031302"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom`

- [query_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-1331332210231132-3211331311203231-3032323113010330-3023003330003301-0033232132220323-2110303320223002-3311020330311023-2030000013230320): complete subsection reference.

<a id="canonical-1331332210231132-3211331311203231-3032323113010330-3023003330003301-0033232132220323-2110303320223002-3311020330311023-2030000013230320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-010.md#canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3312232232221322-1333321212031103-3231103331230300-1233333032313111-2113220211302002-1211312032321312-1113222000020233-3322023020230111)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-3131000330000322-2312232012200201-2123123110313130-3210203103003130-0310100332311132-2201112222322211-2202002033321202-1003100122102121"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121302000021101-3213313203313211-0231011013122201-2230022031330203-1030013201021233-1320302030320010-1101231020220202-2232223321021321"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters`

- [allow_additional_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-2323332331010330-0012033102333123-3031231001230212-1031311002333011-2031233130230101-3023201112321202-2310220311022132-1122111032210333): complete subsection reference.

- [disallow_additional_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-2221112321030012-0222121013033103-3200032022132200-3323130320101021-2010301313223113-3231133321232232-0032100020033200-1222200033130031): complete subsection reference.

<a id="canonical-2323332331010330-0012033102333123-3031231001230212-1031311002333011-2031233130230101-3023201112321202-2310220311022132-1122111032210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-010.md#canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3312232232221322-1333321212031103-3231103331230300-1233333032313111-2113220211302002-1211312032321312-1113222000020233-3322023020230111)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-1331332210231132-3211331311203231-3032323113010330-3023003330003301-0033232132220323-2110303320223002-3311020330311023-2030000013230320)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-0020230231102331-1022311323102022-3103302322002123-3301332033113312-2220210112000232-3231030030031023-1331113233123311-3100311323130211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow additional parameters.

Terraform syntax:

```terraform
allow_additional_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221112321030012-0222121013033103-3200032022132200-3323130320101021-2010301313223113-3231133321232232-0032100020033200-1222200033130031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-010.md#canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3312232232221322-1333321212031103-3231103331230300-1233333032313111-2113220211302002-1211312032321312-1113222000020233-3322023020230111)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-010.md#canonical-1331332210231132-3211331311203231-3032323113010330-3023003330003301-0033232132220323-2110303320223002-3311020330311023-2030000013230320)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-1231231121300223-1102332212232111-2110231103230232-3021323332023321-0123002002310021-3200302302232233-3332020210021022-3101231033013312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disallow additional parameters.

Terraform syntax:

```terraform
disallow_additional_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220102302111311-1200332310001202-1022311232032302-3010330013211210-1322330110010232-0132010211233002-3211213002300121-1332310011020112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-010.md#canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-1322102100322233-2001333103302231-0001210123331333-1020233323230221-0202112301223131-2131322033331220-3222230101301131-1101000310132002"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
property_validation_settings_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="canonical-2023102321232331-3111120030032300-0113022233201331-2331203230222322-3201012112323032-1131100110212301-0010302223132211-0100031022022302"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120222233330131-0212033202211313-1100223023330123-3001302122130220-1220313010133120-0312203211023010-1221031121121213-1111020031331013"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode`

- [response_validation_mode_active](resources--http_loadbalancer--reference--group-010.md#canonical-3031311113220223-3302222123211030-0200331213031113-0013020312121310-2022321320211132-1302133022013022-3233023223333031-1220102113113133): complete subsection reference.

- [skip_response_validation](resources--http_loadbalancer--reference--group-010.md#canonical-0100033220202020-1131021231230323-2222013220201321-0211212331103312-3202020220130223-1220131120310001-1210121333303301-0003231322113031): complete subsection reference.

- [skip_validation](resources--http_loadbalancer--reference--group-010.md#canonical-2333132332321331-2120133301200233-1313323013013330-2300010023330103-2122213002121202-2000232332203321-2022301212110033-1302310200203000): complete subsection reference.

- [validation_mode_active](resources--http_loadbalancer--reference--group-010.md#canonical-2031323230023031-0320021103010303-2103210013020031-1312021133031331-3131111031220032-1202210132311022-2000323333331202-2130231312112322): complete subsection reference.

<a id="canonical-3031311113220223-3302222123211030-0200331213031113-0013020312121310-2022321320211132-1302133022013022-3233023223333031-1220102113113133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-010.md#canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active

<a id="canonical-0220110212102013-1330022112201122-3130303300033220-2031131301222132-0222113000101333-2000030213211123-0001210013313320-2231331113300310"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
response_validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322120001121131-0111323231320321-0210313033230100-3322312201202323-0022022123030320-1132222221313031-3103020211131020-2011103322023203"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active`

- [enforcement_block](resources--http_loadbalancer--reference--group-010.md#canonical-1131212031323201-2200200330122310-3000200010232011-3302031020323020-0310122103312313-1223100221213233-3210111131111303-1210320310130202): complete subsection reference.

- [enforcement_report](resources--http_loadbalancer--reference--group-010.md#canonical-1013012111113200-2213022201012030-2223033120201113-2311311012300130-3030103320001200-0313013100220103-2110200222311212-0220100302030222): complete subsection reference.

<a id="canonical-2312003110131333-2131312021301102-3012321103323112-3122131020211013-3201133120231222-1123310111110312-1202300100001220-3323111132131223"></a>

<a id="canonical-2302113110113022-3220212213103121-1123132112103021-2021100210111022-2330101311211110-0000211103003133-2101000303002023-1221221210211230"></a>

#### `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.response_validation_properties` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1131212031323201-2200200330122310-3000200010232011-3302031020323020-0310122103312313-1223100221213233-3210111131111303-1210320310130202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-010.md#canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-010.md#canonical-3031311113220223-3302222123211030-0200331213031113-0013020312121310-2022321320211132-1302133022013022-3233023223333031-1220102113113133)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-0021230002332302-0213302303011322-0202211320133130-1333222220010113-3232232123023301-0130231220100331-2330231302333120-3303223120101323"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013012111113200-2213022201012030-2223033120201113-2311311012300130-3030103320001200-0313013100220103-2110200222311212-0220100302030222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-010.md#canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-010.md#canonical-3031311113220223-3302222123211030-0200331213031113-0013020312121310-2022321320211132-1302133022013022-3233023223333031-1220102113113133)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-1303230001332011-3231313101333010-2013302103320203-2313113231000030-2233212120123122-2032011121031210-2023202001010223-0221111133320303"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100033220202020-1131021231230323-2222013220201321-0211212331103312-3202020220130223-1220131120310001-1210121333303301-0003231322113031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-010.md#canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="canonical-2230113131230212-1313100311212330-0301333103022101-2223000322303212-3123320133131332-3231100202320010-1103330013130110-2100322313100320"></a>

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
skip_response_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333132332321331-2120133301200233-1313323013013330-2300010023330103-2122213002121202-2000232332203321-2022301212110033-1302310200203000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.skip_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-010.md#canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_validation

<a id="canonical-3323021302001302-1322222101323203-1022321212202210-1000321322311311-0010130320101211-3321022013100013-2132010132201233-3222023100332213"></a>

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
skip_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031323230023031-0320021103010303-2103210013020031-1312021133031331-3131111031220032-1202210132311022-2000323333331202-2130231312112322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-010.md#canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="canonical-0212123020210200-2101033201113232-1021003101212031-0212002021000013-2232031331133231-0221230103220102-1130233221210132-3131100012133030"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230212000031223-3011113103012132-0101300321000221-3031112101102302-2302302200131202-1100320110012031-0332331100010223-3303303111130212"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active`

- [enforcement_block](resources--http_loadbalancer--reference--group-010.md#canonical-0012031010103320-0321311010111210-1321221023133332-1320033100023320-1322010130020120-3031302230210231-2223300133022102-0202023332321031): complete subsection reference.

- [enforcement_report](resources--http_loadbalancer--reference--group-010.md#canonical-2033202131331030-1332230103333233-0210320110320310-2133032012113221-3023300223201322-3111323313001003-3120223120133013-0011123212013300): complete subsection reference.

<a id="canonical-2012320231031201-1113123033001322-0013312012300323-0121210330100102-1213320132102233-3111001323311300-3303123113120321-0332130321020313"></a>

<a id="canonical-1303121112003220-3231010322321222-3233020020332232-2011101112312001-1131223012121323-3313312110230112-3012312010301022-1200020012010020"></a>

#### `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.request_validation_properties` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0012031010103320-0321311010111210-1321221023133332-1320033100023320-1322010130020120-3031302230210231-2223300133022102-0202023332321031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-010.md#canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-010.md#canonical-2031323230023031-0320021103010303-2103210013020031-1312021133031331-3131111031220032-1202210132311022-2000323333331202-2130231312112322)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-2010101020211132-0320121331220030-1001120011323023-1020330113330111-2301301203021132-3132121221310221-2223312101030003-2120320001212210"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033202131331030-1332230103333233-0210320110320310-2133032012113221-3023300223201322-3111323313001003-3120223120133013-0011123212013300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-010.md#canonical-0213222013003311-3002303331300130-2211021301221132-3112030212203001-1213013222321303-2313303103100220-3001020020113301-0020332221130221)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-010.md#canonical-2031323230023031-0320021103010303-2103210013020031-1312021133031331-3131111031220032-1202210132311022-2000323333331202-2130231312112322)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-0002012001033021-0123303122012220-3123113202102011-3221121011220000-0120323122110103-2112112310113012-0312001013302021-3201103331113212"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- api_specification.validation_custom_list

<a id="canonical-0012303032132100-2323311303132102-0010010010012110-3001133002301101-0012231012321212-1002310133202021-3030332320321133-2132110300200301"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
validation_custom_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201200303223000-1322002222212101-1310100223323103-1123003222003120-2231222122320002-1100100223003111-2010111022211130-2303201223000001"></a>

### Direct properties for `api_specification.validation_custom_list`

- [fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300): complete subsection reference.

- [open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2013010220033001-0013100030113300-0200322132301201-3020320120312010-1130321122020102-1023003033030302-3020100331302000-2333212011120001): complete subsection reference.

- [settings](resources--http_loadbalancer--reference--group-011.md#canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223): complete subsection reference.

<a id="canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- api_specification.validation_custom_list.fall_through_mode

<a id="canonical-0311320110301100-2100321110332103-1120232132132122-1230233111200100-3100213301010303-0122212121012301-3002322210233011-0231300031202031"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122212301000223-2210123321101200-3310120110313120-0010002230112030-0232120222211100-1133003001311203-2030231312120033-2331202011203132"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode`

- [fall_through_mode_allow](resources--http_loadbalancer--reference--group-010.md#canonical-2121203103001302-0011232200213011-0233012023312323-0233022221301222-0113010123003032-0132113322322210-1303331102100000-0032320012321020): complete subsection reference.

- [fall_through_mode_custom](resources--http_loadbalancer--reference--group-010.md#canonical-3321313311023320-1002202022131322-2210021110000000-1121001331331323-1321111011210331-3210031032312101-3003203002210210-2223110230111033): complete subsection reference.

<a id="canonical-2121203103001302-0011232200213011-0233012023312323-0233022221301222-0113010123003032-0132113322322210-1303331102100000-0032320012321020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow

<a id="canonical-2323231120110011-2002301212123031-2001033212103332-3012112230100300-1031113133230022-1000330202131030-3333203001212333-1101330101132021"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
fall_through_mode_allow = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321313311023320-1002202022131322-2210021110000000-1121001331331323-1321111011210331-3210031032312101-3003203002210210-2223110230111033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_specification](resources--http_loadbalancer--reference--group-010.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-010.md#canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-010.md#canonical-3231023210233220-1131013113311231-3110223001111211-3003322301130101-0022023310120132-3322003100031313-0323213202311223-1220032202220300)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

<a id="canonical-2022233211133330-3022331013202300-0002113203133331-0001220223201221-3112133203331220-1313112333021002-1312232010111300-3333001332211000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310200030233222-3210013103012231-2220312122120131-1023021010231303-0130102323201211-3222213320211003-2310110310020221-0201023222321120"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom`

- [open_api_validation_rules](resources--http_loadbalancer--reference--group-011.md#canonical-1302322013132300-0121110311233002-1110210221000212-3113113230211323-0100000013020112-0031120011301131-1310321331022103-2021003311133110): complete subsection reference.
