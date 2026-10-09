---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3133321021213011-1113001333203220-2231301222213211-1030201200200330-2132110221330000-0111200133021133-0111011222210101-0002033010023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-013.md#canonical-2010123112022313-2001011302203200-0301123031001302-1100132100302322-2211333103133032-0222112021112221-2320001100313200-2112222101210033)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-013.md#canonical-1023022002223010-1200201131111222-2100222110323303-0011013302311022-1030101202100223-0031230023213120-0233201102220332-1210022220321101)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path

<a id="canonical-1203120020111312-2023010332311301-2120301030002022-0210231012333230-0210302100101031-3213030332131132-2211212010203303-3030030213212233"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-1300012121123032-2201000110303023-3100111022331011-0302302303302312-3032222000003332-2303003100130113-2230110132100031-1120313310223330"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-2202033022310300-0330311112213231-0012221220110331-1300003211213031-3310232320122223-0132310302111131-1021121030000223-0131211130111000"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0030120322300001-2121322033213231-0133132032113232-0003331201330333-0133101111123220-1002120033211130-2312212202203122-0323002301000012"></a>

<a id="canonical-0231311300023000-1201002111333333-1010110110021230-3221210030001103-1112231231013222-3312332120003120-2103102020332102-0232031121112101"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0300201032333221-1231132022031100-1120230220202100-1112132310311031-2122001023323203-0002312223120211-1212110132110111-3021321320003122"></a>

<a id="canonical-0311002102202330-0011213301302123-1113321202102222-0013010031211313-3013210131000332-1120120332300233-3020200312330313-1300332132212112"></a>

#### `bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- bot_defense_advanced_protection.web_only.js_insertion_rules

<a id="canonical-2021213200202320-1330011221112232-3001313211303330-0023011001212133-0132001000202001-2101033221000030-2121133313131102-1001120021000311"></a>

Type: `"single"`. Computed.

This defines custom JavaScript insertion rules for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3002003231232222-3303031021223011-2313322230323012-1123303121222102-2232210022301022-3032032131313220-1001331300032011-3311210210311200"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules`

- [exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-0010302002010221-2230321221203322-1000112223123123-0011023301322002-1300001002002302-0302313120313131-3121300003031031-3320201121130112): complete subsection reference.

- [rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-2121020122222310-1300121031013302-3121333031332011-1030200121113022-2320232200223113-1032002023331330-1011303022200011-3200210301312122): complete subsection reference.

<a id="canonical-0010302002010221-2230321221203322-1000112223123123-0011023301322002-1300001002002302-0302313120313131-3121300003031031-3320201121130112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list

<a id="canonical-0033000033130001-0231301220221133-1302320333110200-2130021110130102-3302221222010120-0132233131132130-2223100012332222-1012010001122321"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0310002022200010-3030130220331302-0012102302022232-1131012121210221-3003022033012020-0112312101013011-2123011002303020-3321030102222303"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list`

- [any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-3300133033330331-2210030202330012-1100213021330022-1332123302113000-0223221121223123-2002320303100230-2312221101001112-1203000003222013): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-2211213333113213-1120230232213221-2102033313332323-3313012332122331-0001300233002310-0200322001312213-2021023223220213-3013002313230232): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-3011131320103023-3111331303223013-2322202333223321-2211003200022213-0312021022300133-1310332232131221-2113111121200321-2201002221033130): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-014.md#canonical-0221321212033331-1120112201321211-0212100200113111-0133113031302211-0302123132133302-3310201301003202-1300112001010230-1020022202120200): complete subsection reference.

<a id="canonical-3300133033330331-2210030202330012-1100213021330022-1332123302113000-0223221121223123-2002320303100230-2312221101001112-1203000003222013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-0010302002010221-2230321221203322-1000112223123123-0011023301322002-1300001002002302-0302313120313131-3121300003031031-3320201121130112)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain

<a id="canonical-1221123003122120-2122031123222322-1121022303012033-3002001301110301-1321103323123023-2310311313321100-1100303000330010-0223030120112322"></a>

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

<a id="canonical-2211213333113213-1120230232213221-2102033313332323-3313012332122331-0001300233002310-0200322001312213-2021023223220213-3013002313230232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-0010302002010221-2230321221203322-1000112223123123-0011023301322002-1300001002002302-0302313120313131-3121300003031031-3320201121130112)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain

<a id="canonical-1201000211010210-0133021333320102-2201003221013331-1102222120121122-1300013232023121-2201202313031331-3023221310221201-3131022333303002"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-1310020010010321-1002001223310123-1013311103310211-1013310012130003-2211013012013130-2312132030122213-0222223031100131-0331123220333231"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain`

<a id="canonical-2100111331323031-0001302310210003-3001100101322310-3030320300122221-2331323320210102-3002010123303211-3121011110000022-0232202330101330"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain.exact_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1023331332321002-2302233013032112-3332103332131012-2333111131023232-2323022020333323-3222212232201010-0010031103110330-3230331203013021"></a>

<a id="canonical-2121022311132032-3030013211332120-1112321332002023-3311000133313233-2301112031002003-3001023010220213-2333003120123013-2132231322202310"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain.regex_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1220303103210321-2103201031230221-0100300223100320-1311102223103002-3022033221320321-3303113220012303-3111200201033013-3122102121313132"></a>

<a id="canonical-0333112011312302-3212312021000113-1100032111123212-1302333332133022-3032211220333033-0032331112120312-0022100123231201-0002021012230010"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain.suffix_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3011131320103023-3111331303223013-2322202333223321-2211003200022213-0312021022300133-1310332232131221-2113111121200321-2201002221033130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-0010302002010221-2230321221203322-1000112223123123-0011023301322002-1300001002002302-0302313120313131-3121300003031031-3320201121130112)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata

<a id="canonical-1001301321120121-0233231331131320-0121130022313331-2012220303100111-3002220211301133-3103203303103311-0232102303222103-0112123212222012"></a>

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

<a id="canonical-1110122131223320-1201102113110203-2332321023333310-1011300221010001-2112313120020013-1222230211303033-1233323300320120-2002210232211312"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata`

<a id="canonical-1110202223311312-2031310210111311-3113320321123120-1020333312303223-3100313332122303-3222210100320203-2312222232113322-0032212130100120"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0331311202010031-3130120212220220-0123100202222102-0102211233301310-2103003120203220-1003210312110321-3320301112100200-0021132310102203"></a>

<a id="canonical-1023332210020213-0323223312320302-3023031201010321-1221203223100022-3213300023333022-1211201131011330-2100112221133331-1130221123121222"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata.name` property

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

<a id="canonical-0221321212033331-1120112201321211-0212100200113111-0133113031302211-0302123132133302-3310201301003202-1300112001010230-1020022202120200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-0010302002010221-2230321221203322-1000112223123123-0011023301322002-1300001002002302-0302313120313131-3121300003031031-3320201121130112)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path

<a id="canonical-2200133202333233-2311001023223201-3102223002003230-3330023331133302-3103112201110003-2021210322002112-3300130321122200-1001002010200110"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-2302010103303030-2121211300011321-0111333301000232-2100301321220311-1332203200112333-1220123220321331-0202101221233103-2012132330132303"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path`

<a id="canonical-1001113311120221-1201103111121310-1222110120020101-1323121033013221-3023321313100330-2233222012213310-3202033123320022-3321232233332002"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2222322003120303-3001011200221202-1313131313221323-1031302202021033-2311022103132333-0220102201133220-1301113220033132-3200100332223013"></a>

<a id="canonical-3033321031300130-0333102211330103-0301100302013210-2333120210231020-0310130102022103-2112021333031310-3030022120030302-3133220122310022"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3303123201023210-1010103002033012-3333211232200212-0213230103031230-1012022013321231-3112010322111213-0302133111110220-3331003210111320"></a>

<a id="canonical-1132121032232231-2222130011113320-3003020213012211-3001322301230310-2031123011211033-0212312310111301-3321321133223301-1331212313020100"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2121020122222310-1300121031013302-3121333031332011-1030200121113022-2320232200223113-1032002023331330-1011303022200011-3200210301312122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules

<a id="canonical-3131020121320001-2220130120122102-1200000323220101-2020130300221230-1332110000101131-1200321001023131-0031320020323022-3222012033001221"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0111131300022310-2321121120023233-0220013211120021-2001030031102101-0102030323120121-0001020022120232-0303113011120311-1102023310012021"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.rules`

- [any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-2330313113323102-1330212100001233-3112331330232121-3000103132220010-1033112311011130-1223032013122223-3022213322132323-0001212312310023): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-3023010200311010-3031013331313003-2323332321222311-2201211322323312-0132231331122331-2102200011313320-0113212312201020-2212023103332320): complete subsection reference.

<a id="canonical-2020210120201210-2303331033220303-1022022310311300-2033032112213332-2010210023213323-2210010232031011-2001223002130203-2013331012322122"></a>

<a id="canonical-1122020200021001-3321312212020203-2202002213101330-0310303012300033-1123211202030311-2203320120230200-0211233131011121-2330130002021230"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.javascript_location` property

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-0132101110002123-3233011323311002-0300231320331303-2021103320212302-2013333010100131-3202330020111230-3020203320321012-0330213003310133): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-014.md#canonical-3103032210221202-3203103300111103-0012201121133122-2332132223012333-2010212303210121-3101333110221032-3213300132021332-0213101030102210): complete subsection reference.

<a id="canonical-2330313113323102-1330212100001233-3112331330232121-3000103132220010-1033112311011130-1223032013122223-3022213322132323-0001212312310023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-2121020122222310-1300121031013302-3121333031332011-1030200121113022-2320232200223113-1032002023331330-1011303022200011-3200210301312122)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain

<a id="canonical-3331200120311022-3023111310130210-1321021013300300-1020032030310320-1102301122122311-2323211100020310-1312332330122210-1230102230112311"></a>

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

<a id="canonical-3023010200311010-3031013331313003-2323332321222311-2201211322323312-0132231331122331-2102200011313320-0113212312201020-2212023103332320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-2121020122222310-1300121031013302-3121333031332011-1030200121113022-2320232200223113-1032002023331330-1011303022200011-3200210301312122)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain

<a id="canonical-0010000011203221-1220203110300111-3010033112131213-1113131310222003-0233330010021201-2002320011030232-3203230301110203-2110100133031133"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-1003231112223300-0322123230010121-3233133122322201-0032112322231230-2311213201213032-3302130210020210-1021030232000112-1300200021012110"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain`

<a id="canonical-1012011120001301-2332213130331001-1032300313132101-1002311203032112-2103313113011010-1211123010030223-0232101210313013-3323113301221232"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain.exact_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0102321132311121-1331321112113032-2033230321301112-2302120302333232-3131222332110001-1133232103220112-0033211133201313-0230322230311310"></a>

<a id="canonical-2212330223232233-3313001030101302-3131330010101332-2221121013121330-0133011110303310-3120300201010303-0122331301001223-1331120213002101"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain.regex_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0021311323231010-2331030000031103-2210220203231300-1323302121031132-0110022003101001-2122023221333011-1103320200330210-1330011121322013"></a>

<a id="canonical-2112233112002321-1331120221011012-0101300332012312-0022330001033101-1202221010232221-1210123221121231-1300002001023200-0110003330013112"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain.suffix_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0132101110002123-3233011323311002-0300231320331303-2021103320212302-2013333010100131-3202330020111230-3020203320321012-0330213003310133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-2121020122222310-1300121031013302-3121333031332011-1030200121113022-2320232200223113-1032002023331330-1011303022200011-3200210301312122)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata

<a id="canonical-2112113333132020-3232332221233222-0222010123301231-0323012113231120-3103111323213130-1032203132230101-1103001322203312-1331211122322013"></a>

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

<a id="canonical-0000211132213312-1200012110200020-2233320230331213-1310121021311232-0113100000133111-1003003203110332-1233022003010302-2330122112002130"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata`

<a id="canonical-2301330221332302-2233333231323332-3002131122312333-0022121203003021-3301211002332120-2133231101330311-2202200211232133-2011002223212202"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1313023223120012-1002001211011111-0323303311210333-0131313023303202-2121331330213320-0233211011130310-1310223332220313-1332212130020103"></a>

<a id="canonical-2311030333121221-3323310001120232-2112022323003013-0002021013103102-3330233111202220-0213301333302233-2133321321001222-0323112123113012"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata.name` property

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

<a id="canonical-3103032210221202-3203103300111103-0012201121133122-2332132223012333-2010212303210121-3101333110221032-3213300132021332-0213101030102210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-0012022320000133-1100030111130203-3221113200303100-0121022133001330-1330013300002100-1110203321132331-3103211113331200-0022030120222321)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-2121020122222310-1300121031013302-3121333031332011-1030200121113022-2320232200223113-1032002023331330-1011303022200011-3200210301312122)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path

<a id="canonical-2011033201323202-0121202310101313-1111323300003231-1213003103200023-3331221012201102-0011032303020100-0030021000133333-1111123113311230"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-2033210212130310-0113231133103130-3232332321132312-1313213222133122-2023332313030032-1231000112332102-0323220001123301-0000223223230010"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path`

<a id="canonical-1123123322012333-2231100231133010-0023301213213103-0223233223203101-3311303330322331-3010102032202011-3012023202301200-3332020103001212"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0330011300030230-1221022233320321-3210131010123303-2232332230211323-2101102331232130-2022012100001002-1202210302032033-2223202100231110"></a>

<a id="canonical-2221332122011202-3100000313320311-2321323231231230-1010021012202303-1333021333012332-3111120013103302-2033231113321020-3310312200202031"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2032003132203112-0210302221321012-3303202300231122-2002202333003101-0100001021222303-3021013210233213-1130030033001120-3031003012201232"></a>

<a id="canonical-1301000031002330-2222200210123210-2230102233131301-1313110310331010-0321101212311321-2123300332322203-3222201122112102-0223130211130232"></a>

#### `bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3331233301232310-3000013112111023-3201310202212232-0000121130110320-3113022211201020-0210222021021230-1220003332102023-2320212100030130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.web_only.web` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--reference--group-013.md#canonical-3312111000032123-2113232132201102-2211030313120002-2330103303233223-2011200103223002-1133120323103103-3122020003302002-0023203221322131)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--reference--group-013.md#canonical-3132003030001203-1302110211110223-3201131102022232-1002200113030130-2133320012301000-3013033120330212-1012201103300121-2311120322102003)
- bot_defense_advanced_protection.web_only.web

<a id="canonical-2130123323012011-2310203030312121-2300300332211120-0133003301233322-3100110330323122-0202310100102302-3321330313101003-0121201022022201"></a>

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

<a id="canonical-0322013111213132-1321313323330311-3232132110102322-2103113120330021-0300331131231322-2020133213121032-3231002102300311-3123130111100131"></a>

### Direct properties for `bot_defense_advanced_protection.web_only.web`

<a id="canonical-2023022032110302-3212103323101011-3232222030203023-0113223003303133-2132100131301232-1113101212132220-3300123000201212-1022000323003120"></a>

#### `bot_defense_advanced_protection.web_only.web.name` property

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

<a id="canonical-3020200012322003-0130233021201002-0010310303110312-0100013212003103-2101110023002222-0223021030133131-2203311031300021-2122110232323030"></a>

<a id="canonical-1333233230100121-3023133333323232-2231230122030121-2002010032322133-1231210203220230-3331123001202113-1202331021222030-3210321111223003"></a>

#### `bot_defense_advanced_protection.web_only.web.namespace` property

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

<a id="canonical-0110003001313002-2102203013330133-2003122012013302-2230320032100323-1112130001212232-1230112302113200-2203021021211123-0023010311223200"></a>

<a id="canonical-1301320203103320-3121221123111220-3100113333310000-0232111023202231-1222320130311112-1311220012120230-3313120231133003-0200100023221320"></a>

#### `bot_defense_advanced_protection.web_only.web.tenant` property

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

<a id="canonical-2212332302031321-3221110201202123-1223303020332011-1212131213222121-1302331201111020-0113210331231011-0111302001212003-1022220130211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- caching_policy

<a id="canonical-3231032210101012-1032131000310210-3032330112212133-2200022320312100-0103300111022130-0010300332232232-0303120001030331-0100123022010021"></a>

Type: `"single"`. Computed.

\[OneOf: caching\_policy, disable\_caching; Default: disable\_caching\] Policy configuration for
this feature.

Additional upstream details:

Caching Policies for the CDN.

Receipt-pinned upstream constraints:

```json
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

- [caching_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-3231032210101012-1032131000310210-3032330112212133-2200022320312100-0103300111022130-0010300332232232-0303120001030331-0100123022010021)
- [disable_caching](data-sources--http_loadbalancer--reference--group-017.md#canonical-1312123312223031-3132131302203112-0321223312021102-2111003030332211-1303112101133122-3303033010123320-1323111123300010-0231330312020120)

Select alternatives according to the provider validators above.

<a id="canonical-1001100300313123-3131030120012303-3313212133222233-2012102223120332-3131231333231120-3211320233321230-1202011301022202-2012122032100222"></a>

### Direct properties for `caching_policy`

- [custom_cache_rule](data-sources--http_loadbalancer--reference--group-014.md#canonical-2103101010232231-0110012110211132-0322322122312321-3300211033231023-1310100312303203-2311312023231103-2232010032031201-0312132102010220): complete subsection reference.

- [default_cache_action](data-sources--http_loadbalancer--reference--group-014.md#canonical-2030011133231130-2202111323312110-1300223120331333-0200132202210011-1330111202102310-2332303231202013-1322220332013231-3031001313200331): complete subsection reference.

<a id="canonical-2103101010232231-0110012110211132-0322322122312321-3300211033231023-1310100312303203-2311312023231103-2232010032031201-0312132102010220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy.custom_cache_rule` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [caching_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-2212332302031321-3221110201202123-1223303020332011-1212131213222121-1302331201111020-0113210331231011-0111302001212003-1022220130211102)
- caching_policy.custom_cache_rule

<a id="canonical-3011031331112202-0203330320223122-1332201211103333-0223221123320112-0011311211210212-0333221233021303-3012212012101332-0230000320331323"></a>

Type: `"single"`. Computed.

Custom Cache Rules. Caching policies for CDN.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0301300211020122-3221302310210211-3012211130330302-0103312101211223-2111011030313310-0312212030202303-1020321212222313-3332332023013200"></a>

### Direct properties for `caching_policy.custom_cache_rule`

- [cdn_cache_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-2113020113101121-2020031022230031-0323020201120000-2203333033223101-0331102021031130-3120011033012301-2321222333001210-3333230211113011): complete subsection reference.

<a id="canonical-2113020113101121-2020031022230031-0323020201120000-2203333033223101-0331102021031130-3120011033012301-2321222333001210-3333230211113011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy.custom_cache_rule.cdn_cache_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [caching_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-2212332302031321-3221110201202123-1223303020332011-1212131213222121-1302331201111020-0113210331231011-0111302001212003-1022220130211102)
- [caching_policy.custom_cache_rule](data-sources--http_loadbalancer--reference--group-014.md#canonical-2103101010232231-0110012110211132-0322322122312321-3300211033231023-1310100312303203-2311312023231103-2232010032031201-0312132102010220)
- caching_policy.custom_cache_rule.cdn_cache_rules

<a id="canonical-3021032112222032-1021203012112300-3322131123221231-2310030321222132-1221110300300323-0203021103303223-2103333011023201-3232033000202132"></a>

Type: `"list"`. Computed.

Reference to CDN Cache Rule configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1122102332211102-1120220122332312-2310302130101030-3301230133301101-1112310232333110-2023303113103030-0211313201212201-1221031311130122"></a>

### Direct properties for `caching_policy.custom_cache_rule.cdn_cache_rules`

<a id="canonical-2110321003111031-1322220003223011-2311001022001232-0230120232200333-2212222302302301-0212330033201030-0232212231022322-1320030211103233"></a>

#### `caching_policy.custom_cache_rule.cdn_cache_rules.name` property

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

<a id="canonical-1030032120100311-0020012202121003-0311302313033303-1031303021103001-2100323121303002-2032311130103012-2003201211222120-2111303223123232"></a>

<a id="canonical-1000010321320102-2000030001003221-0302233332122131-1200002013102322-1010010130012130-0313311202103023-2313221212132121-1302211013112202"></a>

#### `caching_policy.custom_cache_rule.cdn_cache_rules.namespace` property

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

<a id="canonical-3031031001000310-3300220100232231-1003223212121012-2120131213101002-2103230022121102-2103222022011230-3211203312233010-2301300223112203"></a>

<a id="canonical-2112320013010302-0031233331330100-0303233030302311-2311322323131103-1022331003013112-1213030232223122-0100000232331113-0010221313230213"></a>

#### `caching_policy.custom_cache_rule.cdn_cache_rules.tenant` property

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

<a id="canonical-2030011133231130-2202111323312110-1300223120331333-0200132202210011-1330111202102310-2332303231202013-1322220332013231-3031001313200331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy.default_cache_action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [caching_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-2212332302031321-3221110201202123-1223303020332011-1212131213222121-1302331201111020-0113210331231011-0111302001212003-1022220130211102)
- caching_policy.default_cache_action

<a id="canonical-2232111220033313-2313312100222322-1100312330202000-0301213333310132-2211320322210131-0303312112321310-3031022133112322-0102331200012011"></a>

Type: `"single"`. Computed.

Default Cache Behaviour. This defines a Default Cache Action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

<a id="canonical-0230122112330231-2131300010100031-2132131101321230-3010202233031200-0030223022111213-0031123110000213-0213213311233333-1013322033000100"></a>

### Direct properties for `caching_policy.default_cache_action`

- [cache_disabled](data-sources--http_loadbalancer--reference--group-014.md#canonical-2222120133222110-0333102211011113-2313221321133031-3312012030311122-3222310012303032-2120001011212103-3100122021313001-2233212213123311): complete subsection reference.

<a id="canonical-1332012320110011-3200031222320111-2010221031010332-3312220001210322-2031033130103222-3001001032000032-3312232130120211-1230112123102031"></a>

<a id="canonical-3133233020030313-2312002101103212-3100033210003223-2022012220012103-2202332232011300-1021332123312210-0313021210222300-1113023311232322"></a>

#### `caching_policy.default_cache_action.cache_ttl_default` property

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-3111221121022131-2011203102312233-1300100002333131-2023302230002130-3021322131011132-3131203321103231-2330313132122000-3102011333001121"></a>

<a id="canonical-1312022031111231-3032130122330223-1231113210002022-2101233132030333-3313123201021132-0000203002111012-1312333031311302-1302333010031201"></a>

#### `caching_policy.default_cache_action.cache_ttl_override` property

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-2222120133222110-0333102211011113-2313221321133031-3312012030311122-3222310012303032-2120001011212103-3100122021313001-2233212213123311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `caching_policy.default_cache_action.cache_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [caching_policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-2212332302031321-3221110201202123-1223303020332011-1212131213222121-1302331201111020-0113210331231011-0111302001212003-1022220130211102)
- [caching_policy.default_cache_action](data-sources--http_loadbalancer--reference--group-014.md#canonical-2030011133231130-2202111323312110-1300223120331333-0200132202210011-1330111202102310-2332303231202013-1322220332013231-3031001313200331)
- caching_policy.default_cache_action.cache_disabled

<a id="canonical-3003302330033202-2202330332133133-2212231311231233-2013110132023033-2133002213203122-2301223320011312-3132110230232230-1000203302133303"></a>

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

<a id="canonical-2333201102102232-2011010100211101-3321232300202320-0130113200321012-3222011121333011-3213310310031212-2203212330330331-3220013313301331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- captcha_challenge

<a id="canonical-2110223332131312-2331013023232121-3103031223213312-1133233330313132-1213220302210232-1221232032113310-3202220230131120-3231303101333123"></a>

Type: `"single"`. Computed.

\[OneOf: captcha\_challenge, enable\_challenge, js\_challenge, no\_challenge,
policy\_based\_challenge; Default: no\_challenge\] Enables loadbalancer to perform captcha challenge
Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that
pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is
configured to do Captcha Challenge, it will redirect..

Additional upstream details:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha. When loadbalancer is configured to do Captcha
Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page
will have captcha challenge embedded in it. Client will be allowed to make the request only if the
captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid
Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only
human users can pass through. Generally, computers or bots are not capable of solving a captcha. You
can enable either JavaScript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
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

- [captcha_challenge](data-sources--http_loadbalancer--reference--group-014.md#canonical-2110223332131312-2331013023232121-3103031223213312-1133233330313132-1213220302210232-1221232032113310-3202220230131120-3231303101333123)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-2230132011120203-0011333021033130-1010233313001013-1100313220213320-3231313122231210-2033232202103112-3100002200122100-0000223330000031)
- [js_challenge](data-sources--http_loadbalancer--reference--group-019.md#canonical-1330313013020313-0000301213230310-2113112320001221-3102132110313013-3033321333122212-3213022133202312-1321331102121220-0332330321122003)
- [no_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-0121222330302201-1313000022222103-0000332201021130-0021331132102001-3010230020131302-3021233321102212-0331230031230303-0000030023131011)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-0000010110233333-3321322033031133-3333123333332012-3000223230222011-1100121210223232-2221310130103120-2032020232112022-0210222331102001)

Select alternatives according to the provider validators above.

<a id="canonical-0333131022011022-0310222012013131-3120200211210003-0211321332013120-3332012013121002-0201213012232223-3211132201002331-2211110231030133"></a>

### Direct properties for `captcha_challenge`

<a id="canonical-3300103233030221-1003133123012120-1132033220032103-1320133022011230-0332013012123033-0002332120010123-1112001001032313-2033133111203011"></a>

#### `captcha_challenge.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-1221210112030000-1300131302313102-3322211212313330-1110001012230230-1133122222003011-0011003322030323-0021223333203133-3223203330112202"></a>

<a id="canonical-0102203221213031-3310220211313101-2002122101211231-1123120131032202-1121221133032203-3033022132102210-2031003323003213-0213302321321002"></a>

#### `captcha_challenge.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- client_side_defense

<a id="canonical-3030331303311001-0012333100122100-2031232323331133-0121312200010333-1031220333310100-1133330100123030-0203233312012303-2323011200220013"></a>

Type: `"single"`. Computed.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Additional upstream details:

This defines various configuration OPTIONS for Client-Side Defense Policy.

Receipt-pinned upstream constraints:

```json
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

- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-3030331303311001-0012333100122100-2031232323331133-0121312200010333-1031220333310100-1133330100123030-0203233312012303-2323011200220013)
- [disable_client_side_defense](data-sources--http_loadbalancer--reference--group-017.md#canonical-3313223321312022-3132121202023311-2200012303220320-3120033023232211-0103033210021322-3121211210300330-1103022020313021-0322333013120023)

Select alternatives according to the provider validators above.

<a id="canonical-1300121011132121-1312323000013301-1030111230201010-0010332200011332-3113132032011313-1330131000030222-0233223211121022-2010303333220301"></a>

### Direct properties for `client_side_defense`

- [policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223): complete subsection reference.

<a id="canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- client_side_defense.policy

<a id="canonical-3303010333322331-3013201320311220-0012211021313123-2223323023222310-2330300103313201-0320231221003111-3303203202022213-0131133220211211"></a>

Type: `"single"`. Computed.

This defines various configuration OPTIONS for Client-Side Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

<a id="canonical-0113233200330312-2101003110022212-1111003212310023-2211101010222332-3022012222032231-2312222231203331-2121020203131332-3122331212010130"></a>

### Direct properties for `client_side_defense.policy`

- [disable_js_insert](data-sources--http_loadbalancer--reference--group-014.md#canonical-1111120331213133-3103110230332133-0333202102211301-3310010232020031-3103222021311100-3010320002012032-0112130100011002-2321022122233112): complete subsection reference.

- [js_insert_all_pages](data-sources--http_loadbalancer--reference--group-014.md#canonical-3302310012100320-3030020310133232-0202121212000322-2232110023321332-3323002333023313-1301201213220323-3022001212211013-3201001333311010): complete subsection reference.

- [js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-014.md#canonical-3332302132123321-1001101220011123-3023000220130311-3013001211311132-0000311302332201-2210102331010203-3012203311320232-1121020202012312): complete subsection reference.

- [js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132): complete subsection reference.

<a id="canonical-1111120331213133-3103110230332133-0333202102211301-3310010232020031-3103222021311100-3010320002012032-0112130100011002-2321022122233112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.disable_js_insert` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- client_side_defense.policy.disable_js_insert

<a id="canonical-0030103322300013-3123020210223010-0023122011300021-3200310331001311-2321130202110202-2120231031321121-2031132132011200-2312303222330211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable js insert.

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

<a id="canonical-3302310012100320-3030020310133232-0202121212000322-2232110023321332-3323002333023313-1301201213220323-3022001212211013-3201001333311010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-3230210100120200-1311323230122300-3320311020103030-3121002223000331-1311230033230210-3020333033310321-3213211320001021-3133202233113302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for js insert all pages.

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

<a id="canonical-3332302132123321-1001101220011123-3023000220130311-3013001211311132-0000311302332201-2210102331010203-3012203311320232-1121020202012312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-0021200202232031-2320230330133100-0101002320302000-3230112321331230-1101311213023303-3302133313002012-3330232321031013-0111201100120131"></a>

Type: `"single"`. Computed.

Insert Client-Side Defense JavaScript in all pages with the exceptions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100010000203033-2210030001131002-2013222001023122-0032202303330202-3001022230103221-3301023113232101-0313231231120132-1223313220301201"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except`

- [exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-3001131021322100-1033021100001131-0032031130201230-1013300210330323-0120301310321321-2031303231323121-3301322213110123-1223221202211030): complete subsection reference.

<a id="canonical-3001131021322100-1033021100001131-0032031130201230-1013300210330323-0120301310321321-2031303231323121-3301322213110123-1223221202211030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-014.md#canonical-3332302132123321-1001101220011123-3023000220130311-3013001211311132-0000311302332201-2210102331010203-3012203311320232-1121020202012312)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-0020222323222103-1300321020020111-1323000012220221-2130221122202032-0330102300112132-1210122000122023-1302330231110203-3221131012000130"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2032310302202122-0021310013131330-0111003110020332-0122022231212302-2103030310133033-2313000312013000-3331321322032231-3010130112030321"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list`

- [any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-1002223112233200-0111000333111122-2310012121030230-1310311002022220-0210330020033221-0120013032302202-1033330113012311-1122122300310102): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-2002131310221020-0011333113302213-2322031101031300-3332123032212102-2233113023233330-0210223010320233-2223020133210110-0223021030200011): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-014.md#canonical-1032021113102000-3032121303330033-0332013330121300-0320203112030331-1203001230111301-0303320330231311-2331111011132211-3233301322303221): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-014.md#canonical-0002311222031021-0320223112123323-1013202013133120-1301131312311320-2000320001022231-1230112101312323-1230212033233012-1112033000002202): complete subsection reference.

<a id="canonical-1002223112233200-0111000333111122-2310012121030230-1310311002022220-0210330020033221-0120013032302202-1033330113012311-1122122300310102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-014.md#canonical-3332302132123321-1001101220011123-3023000220130311-3013001211311132-0000311302332201-2210102331010203-3012203311320232-1121020202012312)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-3001131021322100-1033021100001131-0032031130201230-1013300210330323-0120301310321321-2031303231323121-3301322213110123-1223221202211030)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-3323323311000022-0221102212212132-1011101100303113-1222201212031320-1003120312233100-1033320233233223-3022322333322323-0212102010110223"></a>

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

<a id="canonical-2002131310221020-0011333113302213-2322031101031300-3332123032212102-2233113023233330-0210223010320233-2223020133210110-0223021030200011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-014.md#canonical-3332302132123321-1001101220011123-3023000220130311-3013001211311132-0000311302332201-2210102331010203-3012203311320232-1121020202012312)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-3001131021322100-1033021100001131-0032031130201230-1013300210330323-0120301310321321-2031303231323121-3301322213110123-1223221202211030)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-1313220300002222-0033320233012033-1033011133111012-3302201123121230-2311213200330131-2222001332302131-2032313312333110-2122000300313200"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-1102120121103200-3213313031323032-0120121112133100-2010101023321322-1321320100102002-1203210232202130-3220330031011110-1321320331313310"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain`

<a id="canonical-3101003013333233-2101101033312130-3103000122133112-2222112202122233-2330210023031303-1301322022130113-3030200322111002-1013201031211130"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain.exact_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3332110110302320-3303221011013112-2131001012031222-1222100212302223-2231232113320232-2130313303200133-0132012130231303-3220120232112310"></a>

<a id="canonical-2011101022232010-2120000132120000-1313230130112121-1123000322222012-3000210012122121-2303123210123201-0311202000200012-0013322230321231"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain.regex_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2133010313213101-0023000311021212-2011202302030312-0131123123023111-1322212332212102-3303113213210302-0130112010210131-1322030200333113"></a>

<a id="canonical-0300133003221000-2003120321203110-1122233030101033-3330202312101103-1322211301301002-2032332303203033-0133021201232021-3021213200131203"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain.suffix_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1032021113102000-3032121303330033-0332013330121300-0320203112030331-1203001230111301-0303320330231311-2331111011132211-3233301322303221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-014.md#canonical-3332302132123321-1001101220011123-3023000220130311-3013001211311132-0000311302332201-2210102331010203-3012203311320232-1121020202012312)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-3001131021322100-1033021100001131-0032031130201230-1013300210330323-0120301310321321-2031303231323121-3301322213110123-1223221202211030)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-1321113100003000-2003003101312132-2110102020112120-2300302210103331-2323101021101102-3102311321132313-1030030000012201-2231012332301022"></a>

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

<a id="canonical-3211202002002101-0110213302321301-0213113232320212-0121032321033133-1312010030222100-2021100021130200-0332223111332112-0131000121102210"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata`

<a id="canonical-1311030000102313-3232133121301111-3103131123302020-0332213131210211-3020003132231121-1030100300130133-3230033233022023-0232321121033300"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3223220232111333-2230113201322113-0301130321121133-1330230223231230-2121031000323320-3120030200202001-1300223232001210-2020100012322300"></a>

<a id="canonical-1031200211332202-1112300322220301-1102310212200032-1131333113031222-0313333211111102-2322212233311132-2132113033223113-0032312113121030"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata.name` property

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

<a id="canonical-0002311222031021-0320223112123323-1013202013133120-1301131312311320-2000320001022231-1230112101312323-1230212033233012-1112033000002202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--http_loadbalancer--reference--group-014.md#canonical-3332302132123321-1001101220011123-3023000220130311-3013001211311132-0000311302332201-2210102331010203-3012203311320232-1121020202012312)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-3001131021322100-1033021100001131-0032031130201230-1013300210330323-0120301310321321-2031303231323121-3301322213110123-1223221202211030)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-3030013023101202-3120301000321232-3310222311220113-1310300010103300-3000210212032100-3303000013012232-3303333022222232-1002032333103232"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-1310320221230100-3311101103112100-2111002202231313-0032033100122333-3232301332000212-3111123300321020-3012123132202232-1020000302311113"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-2233110121232210-1033122331322123-0132300122303031-2201221220110003-2303321000132011-3003333011023133-2030321003200311-0123221323330021"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3320123100213001-2121203203132220-3003011230120111-1301002102122121-0013312303111121-3130322301122212-0011132113213233-1203120301132023"></a>

<a id="canonical-3200123321013103-1302203020202121-2100130003021013-3122112223311123-2330120211302231-3002331202213230-1301221310233222-1221231332122132"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3013320310333002-2202121113311113-2312232330131010-2102230103213020-3001233133000102-1103010311100133-3231223322332030-3102020030021001"></a>

<a id="canonical-0120301033030210-0333323030211110-3102220002212133-2010332122311100-3330033133002113-1013021103112102-3100300030120111-2013133011202010"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-3032101311033121-3033012231331132-0331313300022323-1013123032323313-1313120121320110-2111222121131330-2101100012120123-3333123311131332"></a>

Type: `"single"`. Computed.

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1000323121132200-3322010213131123-0200033000000023-1132230032213013-1100121101103132-3131101100232002-1010132031212303-1210133200230023"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules`

- [exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-1221131131110100-1302232223010102-0032220310300312-3201101322331231-2300212321011313-2130013023010332-3231231330202323-2232213132013123): complete subsection reference.

- [rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0002130103311202-1210112220011212-3100302133302213-0101112031311302-0021330013310103-3002223310313331-1231013032312222-1101130033321301): complete subsection reference.

<a id="canonical-1221131131110100-1302232223010102-0032220310300312-3201101322331231-2300212321011313-2130013023010332-3231231330202323-2232213132013123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-1300313202320200-0102120323130221-0011331012203201-3020203230303113-1333111212131232-3323323021223323-2131303103102133-0012111110100022"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2010203333212003-2113300120230000-1202222030110031-1023211002301010-2311002232221112-0202011102002301-2300110311113303-1202011001113031"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list`

- [any_domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-3332022222113003-2331123321101210-2111222312302133-3333310032320023-3323023003320120-1030122033202221-0232011132212032-1012233122232321): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-014.md#canonical-3323110110310202-3111000200231032-2003020233110103-2131211022002032-3030311200000330-0321232013011111-2031221220112321-3013200213332011): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-015.md#canonical-3310301020231101-3102232211313030-0313003121033120-1131101232033221-3331001220130202-3021002222011110-0223010323310300-2013132220032321): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-015.md#canonical-1102212022221032-1132203202322130-3302222233202000-0212303110221001-1211220112112103-1200121210102230-2133132212310123-2313113330220302): complete subsection reference.

<a id="canonical-3332022222113003-2331123321101210-2111222312302133-3333310032320023-3323023003320120-1030122033202221-0232011132212032-1012233122232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-1221131131110100-1302232223010102-0032220310300312-3201101322331231-2300212321011313-2130013023010332-3231231330202323-2232213132013123)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-3313301003230211-1132133313201102-0201132133121030-0320010010131030-0132122223031013-0132200131222123-1110021022132213-2232300233303131"></a>

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

<a id="canonical-3323110110310202-3111000200231032-2003020233110103-2131211022002032-3030311200000330-0321232013011111-2031221220112321-3013200213332011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-1221131131110100-1302232223010102-0032220310300312-3201101322331231-2300212321011313-2130013023010332-3231231330202323-2232213132013123)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-0210031133300112-1020321020120213-0320230322323301-2311333232330023-0213202130020022-0333100032011221-0120112211201332-1222020211120013"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-1101332033111012-3011122233013323-3021203110312002-0320223223123330-1333023312201231-2232120032331123-0102132233010011-1302203323020220"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.domain`

<a id="canonical-1301112220030021-2032202130323223-2000322133133223-0301312212032301-1200223120101022-1021023120120221-3030300211103001-2131112313233313"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.exact_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3320311333020200-0121101032221010-1111202020301212-0321002332113332-0001003300112310-2331213113311232-3111310303220002-2003301232122313"></a>

<a id="canonical-2300333232320211-2111313130012113-3011133320020220-3232232130300130-3322303303100312-1000010310002030-1202322022312132-2112323103201211"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.regex_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2021330112132031-0211312113210031-2023121032130333-0013032233002021-3233231212223123-0010301302020320-0100033300113002-2120313221101131"></a>

<a id="canonical-1013222013011320-3232012010122232-1212312110010310-3000101321030113-2302013031013033-0322133131121102-3201103200121111-1121320131301331"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.suffix_value` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
