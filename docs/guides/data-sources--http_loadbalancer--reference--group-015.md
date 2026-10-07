---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3310301020231101-3102232211313030-0313003121033120-1131101232033221-3331001220130202-3021002222011110-0223010323310300-2013132220032321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-1221131131110100-1302232223010102-0032220310300312-3201101322331231-2300212321011313-2130013023010332-3231231330202323-2232213132013123)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-3030103201121213-2332211110111000-2320120312230331-0220221220331302-0301310003100201-2202112323012030-2030123030302202-3013121233123121"></a>

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

<a id="canonical-3312022311312330-3212022323203312-3013201321001031-1222032212323230-0110133020320221-1303312121210102-3221031110310002-2013311310023232"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.metadata`

<a id="canonical-0010303230111231-0221121301010233-3002322132023221-2031202122203003-2030131020112312-2201113001203212-0223311213333332-0301113122130012"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2232222111003110-0330031321302021-1302121122000003-2201121211131031-3212123000020021-1213222222230023-1011320101131310-0320202103300120"></a>

<a id="canonical-2031233023312232-3022003213032211-0303331213101301-3101302003330203-3310130332112310-1033101201122021-1032001332320023-0311200110312330"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1102212022221032-1132203202322130-3302222233202000-0212303110221001-1211220112112103-1200121210102230-2133132212310123-2313113330220302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--http_loadbalancer--reference--group-014.md#canonical-1221131131110100-1302232223010102-0032220310300312-3201101322331231-2300212321011313-2130013023010332-3231231330202323-2232213132013123)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-3200200032123033-3031101102231230-3100333311300331-3321220202333003-3301212031123232-0220130331100311-2333312202323022-0133231321313310"></a>

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

<a id="canonical-1321330100012010-2031232130112200-2031203233021123-1213233301312303-1303031333111133-2310121221012120-1003221030120213-0010221110232010"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.path`

<a id="canonical-1330021011303300-3101122100101112-0320331333100301-1120003013102021-2003023112102123-1021112113231311-2011232333112221-0213112212312202"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.path` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2213111022303031-0210120230313231-1003021003033231-2232103111331020-1222130133302112-0132232102002302-2120000020210032-3111113231232232"></a>

<a id="canonical-1231203120001333-3113220003123112-2213313033220303-3321023313310302-3210232112221002-1123211112211220-0203322233202033-3332202113110203"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.prefix` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2222231320110322-2221320130011321-1012020110321220-0332133322003201-3230313020002133-2030322030020203-3133123001211212-3302022220122011"></a>

<a id="canonical-0000233012301002-1031231003210101-1010302111122311-0000120301021122-0122103010230231-0303020312200321-1010311212011000-0122122310001030"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.regex` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0002130103311202-1210112220011212-3100302133302213-0101112031311302-0021330013310103-3002223310313331-1231013032312222-1101130033321301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-0320222122310121-3332220133321110-1212232223321232-3120313011113222-2200130232300021-2223313100113211-0100133332020301-3110302120122333"></a>

Type: `"list"`. Computed.

Required list of pages to insert Client-Side Defense client JavaScript.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0203031302200323-0100110130023100-2001311030010020-1302220232221012-2122222213301203-3303302111123202-1211131003220123-3310303311330121"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules`

- [any_domain](data-sources--http_loadbalancer--reference--group-015.md#canonical-3221113312332111-3222002020222013-2122312211232223-1030323013313233-3321120302023201-2311312132313333-3003230333211020-3321213223122102): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-015.md#canonical-3333220202202010-2111320001010110-1202210101222300-2002003212031333-1031010211213113-1213220013300102-3222020133333030-0233122231320102): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-015.md#canonical-1233112203313010-0012222033123300-3222022313321323-3132101023100332-0232021013101120-3300202030130332-1330230133133301-1102133221331113): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-015.md#canonical-2203320213003032-1030202233233330-3323210012103012-0220221000312111-2211210003121120-0023220111321323-2002221010202230-1313320201130312): complete subsection reference.

<a id="canonical-3221113312332111-3222002020222013-2122312211232223-1030323013313233-3321120302023201-2311312132313333-3003230333211020-3321213223122102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0002130103311202-1210112220011212-3100302133302213-0101112031311302-0021330013310103-3002223310313331-1231013032312222-1101130033321301)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-2230220302211231-1202200213010221-1002132223103200-1300302201010102-2130222103132312-3300213202233012-1213330123203110-2200231123021302"></a>

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

<a id="canonical-3333220202202010-2111320001010110-1202210101222300-2002003212031333-1031010211213113-1213220013300102-3222020133333030-0233122231320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0002130103311202-1210112220011212-3100302133302213-0101112031311302-0021330013310103-3002223310313331-1231013032312222-1101130033321301)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-2011203032213130-1312000302312113-2112023120300112-0103000211231302-0100213301001221-3223233133323020-3112032221120030-3022001303322121"></a>

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

<a id="canonical-0223032320203323-3133301221001222-0013010002001032-1133202122220001-1110213031120103-0123311132031111-0131321022132201-0001201331211102"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.domain`

<a id="canonical-0323310220130222-0103332331322311-0023211123020303-1022101101130331-1212232212302233-0021222322003123-0102123033130311-3000122220000110"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.exact_value` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2131103201012111-1002103100011220-2122223113220132-1100211111023303-1122033001221112-1201230001132210-2303332103333100-3211023222010302"></a>

<a id="canonical-2000211212103202-0112212000202122-3312332102320000-1320110012303133-2121213200101201-3030230321322022-1121023313132103-3321121010101231"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.regex_value` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0023322122203221-1013203113220110-2301321100102002-2221033322003110-2333000003111320-3113113301110330-0123120210123012-1001331223303210"></a>

<a id="canonical-0132022102223133-2332212201332210-3033120200110003-2313133133011311-2131312301222131-3031023002313112-0221201023012232-3012012002301021"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.suffix_value` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1233112203313010-0012222033123300-3222022313321323-3132101023100332-0232021013101120-3300202030130332-1330230133133301-1102133221331113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0002130103311202-1210112220011212-3100302133302213-0101112031311302-0021330013310103-3002223310313331-1231013032312222-1101130033321301)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-0233110300032130-1331110232003331-3321322100201210-3203313223331312-3033322120332223-3330123010112203-0013301130022302-1110131303120023"></a>

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

<a id="canonical-2303103012103210-2213313332023310-0302013122302032-2302031232113202-0220002331231321-1312300023132122-1312313130310320-3003322003231120"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.metadata`

<a id="canonical-0112303231310210-2223330332310221-2022131111203212-2133123001201301-0212123312133303-3222331002223033-3033201011033132-1310101302303112"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3231113131322111-2231113102332101-3321012101120032-3030012021010301-1130233031133000-0031120333133000-1330010133310133-1020121001130123"></a>

<a id="canonical-3230111201001202-3302020331222330-1002332101310333-1033331321012102-1232312232001230-1302122000331102-1003223230111232-2002101201010023"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.metadata.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2203320213003032-1030202233233330-3323210012103012-0220221000312111-2211210003121120-0023220111321323-2002221010202230-1313320201130312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [client_side_defense](data-sources--http_loadbalancer--reference--group-014.md#canonical-0312133320130311-0010212012021202-3103200023313323-3101130322302231-1001031311132032-0131321220033101-0301230331000210-2023031011321233)
- [client_side_defense.policy](data-sources--http_loadbalancer--reference--group-014.md#canonical-0120132330203220-2203212210212301-2332111330000100-1223033222103223-1103223032333100-2212120112330122-0311130020303300-0132301332020223)
- [client_side_defense.policy.js_insertion_rules](data-sources--http_loadbalancer--reference--group-014.md#canonical-3031313030202312-0212302230013022-1101023031210201-3332223301001103-3200003200333310-1103010023201123-2032210221022310-0303322221001132)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0002130103311202-1210112220011212-3100302133302213-0101112031311302-0021330013310103-3002223310313331-1231013032312222-1101130033321301)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-2121033012220200-2001231301001132-3120213012021020-3110221210122223-0111323011201310-2020000120320223-1113313213012102-3101131313202130"></a>

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

<a id="canonical-2001110002230223-0103202100002321-0202120311330032-3122133100301021-3101010012002202-0131330311032121-1021321100030132-3222203213021313"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.path`

<a id="canonical-3021321100030113-2301112333022111-2312113313231203-2303023312322032-0111320113120110-1122002323131202-3232013010120100-1101011312121312"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.path` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1210211000210231-2202302021333231-0021030233031313-1330021331211332-3320030122230330-3202330112023031-3202323332203331-3122012030010202"></a>

<a id="canonical-3222131303000223-0112333021113102-1211023103122323-3101201131031031-0010002121310230-3312210122223033-0103222132101122-0212311112213331"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.prefix` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0123200113023010-0010002023131021-1122113233213011-0230112112130231-0321200031102233-1030102120113201-2112111323210023-1123112300011002"></a>

<a id="canonical-1023201103220230-2220101000133203-1232310201213323-3231013120233120-0320330232201002-2330330123230210-3320213323332322-0121111232331333"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.regex` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- cookie_stickiness

<a id="canonical-2103113203130312-1102020133122332-3020223011233233-0212010311022222-3332033111232322-2031132200102211-3100332001230023-0202303302313321"></a>

Type: `"single"`. Computed.

\[OneOf: cookie\_stickiness, least\_active, random, ring\_hash, round\_robin,
source\_ip\_stickiness; Default: round\_robin\] Two types of cookie affinity: 1. Passive. Takes a
cookie that's present in the cookies header and hashes on its value. 2. Generated. Generates and
sets a cookie with an expiration (TTL) on the first request from the client in its response to the
client, based on the endpoint the request gets..

Additional upstream details:

Two types of cookie affinity:

&#8203;1. Generates and sets a cookie with an expiration (TTL) on the first request from the client
in its response to the client, based on the endpoint the request gets sent to. The client then
presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

OneOf alternatives in this subsection:

- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-015.md#canonical-2103113203130312-1102020133122332-3020223011233233-0212010311022222-3332033111232322-2031132200102211-3100332001230023-0202303302313321)
- [least_active](data-sources--http_loadbalancer--reference--group-020.md#canonical-0320323021313031-3212211223022031-3213230131132232-3310022021123333-3021230311220030-1312120111233112-3301232133131302-2303211223003103)
- [random](data-sources--http_loadbalancer--reference--group-023.md#canonical-2123033220232332-2220220302011202-0000233120322210-2323011320220213-3231120212211230-3023111203101113-0102303132032332-3003021103131300)
- [ring_hash](data-sources--http_loadbalancer--reference--group-024.md#canonical-0323322103101000-0200310331030220-3311002233001310-1210203133131012-0201131201133211-2020032232103330-0213201232312002-3301321221002222)
- [round_robin](data-sources--http_loadbalancer--reference--group-024.md#canonical-0113223033112001-3033310213310132-3331011031030032-1122222000211010-1301100002000310-2102331021103223-3122121232331301-2211310020000133)
- [source_ip_stickiness](data-sources--http_loadbalancer--reference--group-026.md#canonical-2232300333201231-1322210302201100-2133311321132322-2102213131331030-1211231100230011-0131302130333003-1112000021020133-3102331223113200)

Select alternatives according to the provider validators above.

<a id="canonical-1212332003222311-2201032223121212-2222311202301132-2112310222111302-3333311323330310-1022310132033300-1123323221211002-1233211113300000"></a>

### Direct properties for `cookie_stickiness`

- [add_httponly](data-sources--http_loadbalancer--reference--group-015.md#canonical-1111101103313133-3120101212020231-2223023112332320-0311011332230002-3120002131012003-2133003223312022-3122011222330120-1201311311132223): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--reference--group-015.md#canonical-1112332111203033-1021012100323231-3003113112020321-1010020101111101-3332003010110220-0202112230201113-0012033233030203-0312102323002222): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-015.md#canonical-0201220223133021-3012321032331000-3302110122013110-3332003333021332-2303033213002200-2311200110121022-2321123130002103-0202002113323000): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-015.md#canonical-2203303122030113-0332330230031120-1130332331032220-0131223221210312-0200322101313023-2212313020110133-3321331123012031-2322012303123131): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-015.md#canonical-2133112113013023-1302233303233303-0131103002331302-3212101220012030-1121230232131331-0303120121122203-2133200102303211-2003313013132131): complete subsection reference.

<a id="canonical-1100323110310001-3210120112121231-2221001323320310-2122220333111113-2001323312020033-1221300232012102-1002233330122322-2201001122321333"></a>

<a id="canonical-3022210222223003-1030112312321122-2121213200301032-2320103321330201-0333320331002001-0010101121011011-1331023010030102-1130002003213113"></a>

#### `cookie_stickiness.name` property

Type: `"string"`. Computed.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2313320132032233-1201302101123031-0203202133322123-0233323232102031-1322311010103330-2032303110330302-3212020101111101-0122112332213323"></a>

<a id="canonical-1110301302023111-1323301321132033-1002230301230313-2033300220312121-1111330230101323-3023132321111000-2002002322132201-0202132233333211"></a>

#### `cookie_stickiness.path` property

Type: `"string"`. Computed.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-015.md#canonical-0223212213300332-2003330101023100-0130211331022011-1333212123112031-3200013011133013-3110311331332322-0301321130122031-2133120131312033): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-015.md#canonical-0030002313322203-2110032202222300-1220032300230012-2310023000022320-1131101122010213-1321303312002031-0133020031202110-3033110132111200): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-015.md#canonical-2230213103032230-0202013000210023-2333110231011201-0222212322130131-3020210321030330-2301212120210030-3011002310020331-3322322321222312): complete subsection reference.

<a id="canonical-1313221213312122-2231202211130221-1201322322122202-0130232333000302-0323123123323023-1213331322002010-3102331021301123-0323221120100333"></a>

<a id="canonical-1310121311011212-0111212133211323-2303322111221213-2331123231303310-3333200301231020-2333022031212300-0113213123320103-0220331211111111"></a>

#### `cookie_stickiness.ttl` property

Type: `"number"`. Computed.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1111101103313133-3120101212020231-2223023112332320-0311011332230002-3120002131012003-2133003223312022-3122011222330120-1201311311132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-015.md#canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323)
- cookie_stickiness.add_httponly

<a id="canonical-0320212231303221-2033110033003303-2231133213112020-1330211213320022-1031123002310112-2213213201010111-2023332323310313-3030112121202123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-1112332111203033-1021012100323231-3003113112020321-1010020101111101-3332003010110220-0202112230201113-0012033233030203-0312102323002222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-015.md#canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323)
- cookie_stickiness.add_secure

<a id="canonical-0210320021331203-0122313112013022-2312322320333321-1003012201101332-0012102013003302-3030322033313320-1110030123303310-2132333331010303"></a>

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

<a id="canonical-0201220223133021-3012321032331000-3302110122013110-3332003333021332-2303033213002200-2311200110121022-2321123130002103-0202002113323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-015.md#canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323)
- cookie_stickiness.ignore_httponly

<a id="canonical-2113311230113221-0222222000132321-3320212232123233-0101022032133302-1303201322220111-0001022003220303-1022100233133220-1000130220312030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-2203303122030113-0332330230031120-1130332331032220-0131223221210312-0200322101313023-2212313020110133-3321331123012031-2322012303123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-015.md#canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323)
- cookie_stickiness.ignore_samesite

<a id="canonical-2230122321201201-1122130333300321-1030102230032130-1101233112131021-3110220022201111-0203300320011133-2022122100010230-1011310002000120"></a>

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

<a id="canonical-2133112113013023-1302233303233303-0131103002331302-3212101220012030-1121230232131331-0303120121122203-2133200102303211-2003313013132131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-015.md#canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323)
- cookie_stickiness.ignore_secure

<a id="canonical-0032230033033032-0002322020010210-0202122200120200-1011212003032202-3023222210101321-1022233311103132-1323020201231020-0211232332331233"></a>

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

<a id="canonical-0223212213300332-2003330101023100-0130211331022011-1333212123112031-3200013011133013-3110311331332322-0301321130122031-2133120131312033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-015.md#canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323)
- cookie_stickiness.samesite_lax

<a id="canonical-3231302132231113-3313131032300230-0113102322131030-3001113300201131-1102212003110332-2111000023300120-2222101021200022-0023031120221203"></a>

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

<a id="canonical-0030002313322203-2110032202222300-1220032300230012-2310023000022320-1131101122010213-1321303312002031-0133020031202110-3033110132111200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-015.md#canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323)
- cookie_stickiness.samesite_none

<a id="canonical-0333130020023321-2332220300201123-0303002230212310-0333330111002321-3232221103112230-1211132122330331-1001101023211310-2100330102330102"></a>

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

<a id="canonical-2230213103032230-0202013000210023-2333110231011201-0222212322130131-3020210321030330-2301212120210030-3011002310020331-3322322321222312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cookie_stickiness.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [cookie_stickiness](data-sources--http_loadbalancer--reference--group-015.md#canonical-2230223021011213-0211213033300102-0023220233213021-1203031332112133-3003212222212310-1301211112321222-3230123212303111-0123111030322323)
- cookie_stickiness.samesite_strict

<a id="canonical-1201101000203311-1021210030111011-0211233002130112-2230210222230021-2331103323322002-3310301123023012-0312201203001123-0303021221211321"></a>

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

<a id="canonical-1033333121011121-2330112221131202-0332012230130101-2330221123112012-1131023310033132-0021220023112313-1111023210121300-3213221010122332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cors_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- cors_policy

<a id="canonical-1021230100330303-1012112213330232-3013022223203320-1222233102121020-2211103230133030-0133213002123033-2122102011110311-0030231133131302"></a>

Type: `"single"`. Computed.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.HTML Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2211102230110112-0013303221313301-0002112310012130-0222101030321212-0223123310323221-1231232122112013-3132011200113331-3100033232200323"></a>

### Direct properties for `cors_policy`

<a id="canonical-0121103200113310-3021112332303011-3223303103120100-0313023001120213-0113012321231110-1013330332110200-3323221100131001-2223200103022232"></a>

#### `cors_policy.allow_credentials` property

Type: `"bool"`. Computed.

Specifies whether the resource allows credentials.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2030001312023010-1332031120122122-1122133311210103-0031211131112002-1022132030010332-0223023320311313-1330300013002223-1221322301100100"></a>

<a id="canonical-2212301220333103-3330211311121213-2233233132101112-0200301302312232-3222220101010101-0321032133320103-2021232201210000-2320313322300021"></a>

#### `cors_policy.allow_headers` property

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-headers header.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0310223233211123-3020333301110031-3101020132202201-1023300321130132-2031110002202120-0113231320203202-0231012313130213-1021132323231302"></a>

<a id="canonical-3022213032130112-2113212232110211-1211322221320333-0313102210231211-2302003020210020-3011010322212321-3231130011333130-2101132132220203"></a>

#### `cors_policy.allow_methods` property

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-methods header.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-1210320132123120-0211131323023112-3023323320103233-2101032210003320-3121012130101330-3310010330131322-3011011230303300-2033313011311232"></a>

<a id="canonical-2013022010001230-1110321320001030-3311032123230100-1223211101112201-2230101321311323-2222031203031033-0203120031213031-1131132211002223"></a>

#### `cors_policy.allow_origin` property

Type: `["list", "string"]`. Computed.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0231303233110023-1332331303032300-2310130110330223-0121013300002003-1002011323001220-1100012122013021-0320322032221130-3020303003332320"></a>

<a id="canonical-0132123323111323-3220323003121303-1312130131023330-1131131103201233-2212003032301331-1021233232002233-0202201003120302-2300120320222032"></a>

#### `cors_policy.allow_origin_regex` property

Type: `["list", "string"]`. Computed.

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2031113032002023-1230210020230310-3210310302313303-0023113213331123-0211002131103013-2203310330021221-1123321013302003-2112002231002302"></a>

<a id="canonical-2313000103131200-3200211310222012-2021303212230020-2323212213210221-1103100121202330-0211031220201331-1013301231020212-0003111031223132"></a>

#### `cors_policy.disabled` property

Type: `"bool"`. Computed.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3130312123102310-3302210313011132-0321112331211321-0133122101203132-2012130110322013-0033020323112221-1210230323020022-2321200010212233"></a>

<a id="canonical-3201123221323221-1231032032120213-1230311111230303-3133332203100311-1112000331122312-1132022130112111-3111220323002310-0131200202100000"></a>

#### `cors_policy.expose_headers` property

Type: `"string"`. Computed.

Specifies the content for the access-control-expose-headers header.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3330301222222121-3000310331211201-3320301113013300-0003112130013230-0330101102310113-1311002101121211-2001023122132031-3031210321223230"></a>

<a id="canonical-3311122222303302-0300201210211033-3033020302133333-1033212210131231-1121231013320312-2330220003003330-2131200321222002-1021301113010310"></a>

#### `cors_policy.maximum_age` property

Type: `"number"`. Computed.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

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
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-3301321210130101-1003211310201200-2003003331000201-1010102230311020-3213103202230123-1233212032302110-1131123131102302-1031120113100113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- csrf_policy

<a id="canonical-2022003000222211-0212020320002311-0012310132330301-3303330232312023-2031012030120223-2020220112010200-0220203211001022-3312101322012201"></a>

Type: `"single"`. Computed.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

<a id="canonical-2133300302113133-3200312220030200-1233112330121212-2221011003013322-1213002033220211-0322203222211310-0213210221111101-0202122323100302"></a>

### Direct properties for `csrf_policy`

- [all_load_balancer_domains](data-sources--http_loadbalancer--reference--group-015.md#canonical-3322030002021231-0111321102322232-1101322212102202-2022311312130233-0032321310021030-3113300210113301-1233233001311332-0101110222130110): complete subsection reference.

- [custom_domain_list](data-sources--http_loadbalancer--reference--group-015.md#canonical-0232010010012222-2221211021013110-3003131302222003-1231133203301111-3333212011200323-1330111002210331-3020321320231201-3101212223012230): complete subsection reference.

- [disabled](data-sources--http_loadbalancer--reference--group-015.md#canonical-1312011032030321-2301233123320321-3121132123333003-1131333211130301-0103220210001330-1210033111332231-0001012023303020-3332210333230200): complete subsection reference.

<a id="canonical-3322030002021231-0111321102322232-1101322212102202-2022311312130233-0032321310021030-3113300210113301-1233233001311332-0101110222130110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.all_load_balancer_domains` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [csrf_policy](data-sources--http_loadbalancer--reference--group-015.md#canonical-3301321210130101-1003211310201200-2003003331000201-1010102230311020-3213103202230123-1233212032302110-1131123131102302-1031120113100113)
- csrf_policy.all_load_balancer_domains

<a id="canonical-3233233211000310-2303223122102112-3123133022130201-3203211200121211-0323301312333011-3023320300323111-0212213333023133-3020110213301203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all load balancer domains.

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

<a id="canonical-0232010010012222-2221211021013110-3003131302222003-1231133203301111-3333212011200323-1330111002210331-3020321320231201-3101212223012230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.custom_domain_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [csrf_policy](data-sources--http_loadbalancer--reference--group-015.md#canonical-3301321210130101-1003211310201200-2003003331000201-1010102230311020-3213103202230123-1233212032302110-1131123131102302-1031120113100113)
- csrf_policy.custom_domain_list

<a id="canonical-3213100113211000-0313121123033022-2210031103021203-3032003121120122-3100110102121100-3233113222310323-2222322020010003-1022120102232212"></a>

Type: `"single"`. Computed.

List of domain names used for Host header matching.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3212312020121001-1233020103203322-1030231232000302-0022020131232212-3300200213111203-3313201110301120-1131000220312331-1212201023212230"></a>

### Direct properties for `csrf_policy.custom_domain_list`

<a id="canonical-1301300133100023-2230223210123023-1031330302213301-2300231331232212-3132013211311103-1332211112003220-2202322132000131-2200003322320210"></a>

#### `csrf_policy.custom_domain_list.domains` property

Type: `["list", "string"]`. Computed.

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312011032030321-2301233123320321-3121132123333003-1131333211130301-0103220210001330-1210033111332231-0001012023303020-3332210333230200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `csrf_policy.disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [csrf_policy](data-sources--http_loadbalancer--reference--group-015.md#canonical-3301321210130101-1003211310201200-2003003331000201-1010102230311020-3213103202230123-1233212032302110-1131123131102302-1031120113100113)
- csrf_policy.disabled

<a id="canonical-2303001311323112-2301303110321233-2203321130211211-1103323031233122-2233012133211013-0011230030222210-1030021011301133-0002301123322102"></a>

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

<a id="canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- data_guard_rules

<a id="canonical-0220022120112300-0310212232200223-3331003021100112-1202321201032102-0300112323200222-3122032310131011-2232120012201100-2310211020130200"></a>

Type: `"list"`. Computed.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2320121222010313-3010120313113000-3121221212220132-3001331222233222-1002230122102113-0302200202220012-3032222130222333-1110132223320230"></a>

### Direct properties for `data_guard_rules`

- [any_domain](data-sources--http_loadbalancer--reference--group-015.md#canonical-0300302101330320-3211200123000012-1031002003101213-0002332101323012-0131223120200210-3322131331210203-0301132122032201-3100310230202111): complete subsection reference.

- [apply_data_guard](data-sources--http_loadbalancer--reference--group-015.md#canonical-0232321022032122-2332332021122100-2320200213020210-1332313001021133-1330002331331202-1311320032331003-0312203331213020-1120301103331331): complete subsection reference.

<a id="canonical-1201211013001020-2131123033311311-1013222130020333-3210313230030003-2223012203203313-3302212121121302-0332332222232333-0023133331233220"></a>

<a id="canonical-2301121320320310-1003300200220120-1202002010011012-1311021111102011-2203223000031202-3331302212201301-2200232300201302-2010013203013330"></a>

#### `data_guard_rules.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [metadata](data-sources--http_loadbalancer--reference--group-015.md#canonical-0133002021113121-3031213001032121-3100332220322133-3300030230013212-0331022012120010-0112030330012111-2022122103102221-0013213223301210): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-015.md#canonical-0201200113103120-1212023202111230-1022202310100130-2131303113032231-0000001012230232-1121331030333211-1103221220302011-0332230021210331): complete subsection reference.

- [skip_data_guard](data-sources--http_loadbalancer--reference--group-015.md#canonical-2111200221101132-1102030310103202-0123131133013021-2000001130233001-2321323033332102-1202013010233021-0002303120110202-1322310203021212): complete subsection reference.

<a id="canonical-1131002202233200-0320312200002330-0203031132321222-3131130012130230-1233213023213133-2132023123220033-1213223003222311-0221213023131200"></a>

<a id="canonical-1031322112132120-1103312033311320-0221212013002030-0200113312312322-1030102221113020-2113210130011233-0033231032212201-3320011311231231"></a>

#### `data_guard_rules.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0300302101330320-3211200123000012-1031002003101213-0002332101323012-0131223120200210-3322131331210203-0301132122032201-3100310230202111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003)
- data_guard_rules.any_domain

<a id="canonical-1033202322112331-3102123022111122-1121032220301323-2313230110002030-2222300013022202-3220230333011011-3323301022111232-0120211331103333"></a>

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

<a id="canonical-0232321022032122-2332332021122100-2320200213020210-1332313001021133-1330002331331202-1311320032331003-0312203331213020-1120301103331331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.apply_data_guard` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003)
- data_guard_rules.apply_data_guard

<a id="canonical-3123330130300112-2221101001312012-1123023023202320-1121332013023213-2333211332100300-3122322130033310-0203333031232301-0333212232030201"></a>

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

<a id="canonical-0133002021113121-3031213001032121-3100332220322133-3300030230013212-0331022012120010-0112030330012111-2022122103102221-0013213223301210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003)
- data_guard_rules.metadata

<a id="canonical-1011010103000001-3330023020132213-0112001022313300-0323230323320232-2020311232301211-2121023232311100-1323100313221211-2113003301311020"></a>

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

<a id="canonical-1312121210100200-2013321232131102-0032231020310101-0003321221202021-2133331202213011-2033131201221021-0202301131333232-1303020133112133"></a>

### Direct properties for `data_guard_rules.metadata`

<a id="canonical-2001311222121002-0233120032100333-1013210003222120-0322102100223321-3320231112111311-3111003122000033-2221022020202202-1133033121003013"></a>

#### `data_guard_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0130031032331101-3103132321322303-1200003003323132-1213221000222011-3213031001230132-1312213112311321-1321230132030110-1013322212000331"></a>

<a id="canonical-1032202222311233-3200011021101130-1201212111323333-3220000233222200-3110133201001231-3221120112111213-1231102210311211-1131222000133103"></a>

#### `data_guard_rules.metadata.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0201200113103120-1212023202111230-1022202310100130-2131303113032231-0000001012230232-1121331030333211-1103221220302011-0332230021210331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003)
- data_guard_rules.path

<a id="canonical-3000203000322121-0000031031101010-2112013002033303-0130232030133322-1202110011311010-2231231313013111-2012232233232210-3030113101210230"></a>

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

<a id="canonical-2002113022131110-2321231021211020-2332032331320331-3002032212123200-1222131233033331-3302101222200303-1011312200031331-1322203111233132"></a>

### Direct properties for `data_guard_rules.path`

<a id="canonical-3231333121123031-3232131010320022-0230103210323122-2103100332201310-2102321221001021-0121100023330123-2013300012330332-2300131230121201"></a>

#### `data_guard_rules.path.path` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2223020313313013-3031212302132220-0201101031002333-3010012232332113-0112302102302232-2320330021030122-0202330320010201-0301001320023002"></a>

<a id="canonical-2300303212331332-1321303033321233-1103221133113031-0023210100122032-0132003333012331-1222033122003101-3302310123220223-3021110101221210"></a>

#### `data_guard_rules.path.prefix` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1211230011131211-1210112202330331-0132133011300013-2301320001202123-0323313302010103-2033001103123222-3220230032230211-1120110211332003"></a>

<a id="canonical-2131302322222000-3021001210301211-0221212202203203-2211310030130301-2031201032132311-3300030023011221-0331311131310332-3323110121313131"></a>

#### `data_guard_rules.path.regex` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2111200221101132-1102030310103202-0123131133013021-2000001130233001-2321323033332102-1202013010233021-0002303120110202-1322310203021212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `data_guard_rules.skip_data_guard` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [data_guard_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-3102323311020130-1122103001002133-3000002211222110-3333021323121303-3331301010012011-3020222021303103-3231322003001201-3132013303031003)
- data_guard_rules.skip_data_guard

<a id="canonical-2112022130230301-0101311030002021-1032221133312010-2230032200210311-2020311310313113-1222131212020123-0202231202001333-2132323333323102"></a>

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

<a id="canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- ddos_mitigation_rules

<a id="canonical-0222330120211123-0011202332200013-2112301203110122-1110303322031332-1103131020323122-2323100131123103-3202211033221113-1010303020023112"></a>

Type: `"list"`. Computed.

Define manual mitigation rules to block L7 DDoS attacks.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
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

<a id="canonical-2122102303110000-1311220122212220-3010103102322121-2100200302031130-2202002202303000-0122210101121212-2301231011230030-0112123330330322"></a>

### Direct properties for `ddos_mitigation_rules`

- [block](data-sources--http_loadbalancer--reference--group-015.md#canonical-3123311102233033-2133220002012003-1200022032312311-2033132203131033-0302133003210102-0211133322321231-1023020222113331-2200302213010302): complete subsection reference.

- [ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203): complete subsection reference.

<a id="canonical-2203100123132033-1122223233202320-2032112011201002-3332121103001221-1330203133213012-3011022002221322-3233301023113233-0011123302003330"></a>

<a id="canonical-1323203003010022-1322230211320233-0123230010311313-2000220131233020-2011312012132031-2011333333331303-1000210223203122-1002320222233311"></a>

#### `ddos_mitigation_rules.expiration_timestamp` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-015.md#canonical-2211231022121230-1011323212121001-2133133122021303-3332030210202013-2210323321121202-3323133220200310-1123101311301033-1102111103021310): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-015.md#canonical-3211223222113323-1020030013300003-0322233310230102-1333301120222330-2330300300132111-2030313102210023-3301233101112230-0011300100122012): complete subsection reference.

<a id="canonical-3123311102233033-2133220002012003-1200022032312311-2033132203131033-0302133003210102-0211133322321231-1023020222113331-2200302213010302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- ddos_mitigation_rules.block

<a id="canonical-1210012321002200-3333111301020223-2212103132320332-0321130112320022-0312332321031123-1330132230122313-1333100201321203-2030000200321323"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-2032003231013220-1023210320302003-2231130230313201-3231322112330332-2021331312003021-1111331330303120-1023123233012013-1033231132332323"></a>

Type: `"single"`. Computed.

DDoS Client Source Choice. DDoS Mitigation sources to be blocked.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3013213121011330-3212301332131133-2001233313102221-3300221031003130-2320013102200330-0223103123102121-1301213012311112-3213113122003311"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source`

- [asn_list](data-sources--http_loadbalancer--reference--group-015.md#canonical-3330032330013000-1211133133122010-2130122100120230-0100010212032131-2332210311132233-0101302230121001-1213033123113100-2332311202002232): complete subsection reference.

<a id="canonical-0122330112132233-0210130302032203-2032023000213302-2321212112010300-3331132132022321-1103310013113213-0121111201020120-0103003310011321"></a>

<a id="canonical-0233213202130221-2112011311122202-1011031100122230-3320021022212123-2012202002131321-0223212120220212-3021131322221210-2233022333123032"></a>

#### `ddos_mitigation_rules.ddos_client_source.country_list` property

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Sources that are located in one of the countries in the given list. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ja4_tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-015.md#canonical-2022101130232011-2200203221023030-2232112210033120-2231121332031203-0320103110231212-0303033002020202-0213132313131332-0103023320030032): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-015.md#canonical-3221220000322331-1103032020303023-3303013302320311-3111100110122312-3021233331202213-0102203132223012-2320202201010211-0132302112322330): complete subsection reference.

<a id="canonical-3330032330013000-1211133133122010-2130122100120230-0100010212032131-2332210311132233-0101302230121001-1213033123113100-2332311202002232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-3030120101310201-3131323021220120-3002102113002201-0013130130220233-0130210121123012-1211312231321012-1301211121301000-2232130111211332"></a>

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

<a id="canonical-1123101310123232-3220010020330321-3201221220311312-1203032231221111-2122221030120212-1123112013223133-0333220112133112-1030102203313013"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source.asn_list`

<a id="canonical-1121022201203112-3300122201331132-1331003130111210-3211231332321222-0312202220333331-1133200331020111-2332332301303023-3333210120102220"></a>

#### `ddos_mitigation_rules.ddos_client_source.asn_list.as_numbers` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2022101130232011-2200203221023030-2232112210033120-2231121332031203-0320103110231212-0303033002020202-0213132313131332-0103023320030032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-1003033321212022-2320022013112220-3123033231332312-2010313110330210-0032313013233233-1102311232021301-3010000131303103-3110310312023311"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1233013322001311-0232020232330333-2003010212223223-3233231123013230-2212330130212321-3200222202031201-0322202122100110-1200302312100133"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher`

<a id="canonical-0110123013032300-0212020312223201-0311212211331220-1031103023220211-3011231222232213-3111301100133012-3121012222103202-1122000013033020"></a>

#### `ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3221220000322331-1103032020303023-3303013302320311-3111100110122312-3021233331202213-0102203132223012-2320202201010211-0132302112322330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- [ddos_mitigation_rules.ddos_client_source](data-sources--http_loadbalancer--reference--group-015.md#canonical-2010132332231223-1321201300322332-1003012200330023-2322321131102132-2330033003320010-2133300030311102-1211110222320210-1201200010002203)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-2201001123302320-2201213010210332-3220002202202121-2333202003320130-3001230102003121-2023231300103231-0213211333213020-2203303331200010"></a>

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

<a id="canonical-2020312333301333-0100303023222310-0001232320223221-0122010232312221-1333021313023230-3012303030131022-3223102332230030-0322013310213330"></a>

### Direct properties for `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher`

<a id="canonical-2123212230122013-0011230111323131-2102001210120213-0110230110231020-2313022021331012-0331013120110300-0212213323101131-0201310231131222"></a>

#### `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher.classes` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1113221201032101-3121131303312211-3132102112031311-3213030303300021-1030330332102323-0013333103311021-3313203333232120-2031130120212311"></a>

<a id="canonical-0302101110221203-1112002013201300-2212030113321311-0330222202203030-2000130301222312-1311333221330303-3221313300222303-2032201211232230"></a>

#### `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher.exact_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0201300110310020-0021120302031133-1323111331131122-2101203031102332-0311120221222212-2133020323200303-2013301113000310-3321120211112130"></a>

<a id="canonical-0213223011012331-0210012332311221-1032022322010301-1011300000301331-2212010032030213-2023133300222332-2012102302200221-3003311033213001"></a>

#### `ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher.excluded_values` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2211231022121230-1011323212121001-2133133122021303-3332030210202013-2210323321121202-3323133220200310-1123101311301033-1102111103021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-3201210333000031-1121021211121013-3302123100100212-3021231202223130-2103302202022131-2223203122303123-3011120202230210-0200221100302231"></a>

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

<a id="canonical-1103310032102032-0312302032320311-2221031133021120-0311212310100111-3111321313101030-3113210213131323-2122112323302103-2123222102032120"></a>

### Direct properties for `ddos_mitigation_rules.ip_prefix_list`

<a id="canonical-1030010202323302-2323000311201012-3032102022000302-3301321013011101-2030000030020010-2313211011231130-1100013202131223-1313222120220200"></a>

#### `ddos_mitigation_rules.ip_prefix_list.invert_match` property

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

<a id="canonical-2031231122103033-0322320332231110-3203003212000333-0300023302123212-1231210323200203-3211120210033123-0023010122020331-2312302222232000"></a>

<a id="canonical-1203313102130010-1332223211233023-2222201330111211-3012121332232110-1121223132001310-0002002322032101-3113003002110131-2133020320020202"></a>

#### `ddos_mitigation_rules.ip_prefix_list.ip_prefixes` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3211223222113323-1020030013300003-0322233310230102-1333301120222330-2330300300132111-2030313102210023-3301233101112230-0011300100122012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_mitigation_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [ddos_mitigation_rules](data-sources--http_loadbalancer--reference--group-015.md#canonical-0301200302321031-3023010021231312-0100023302102310-0312021220023102-2032230321111333-2300113021102000-1003123320300330-0230013100020001)
- ddos_mitigation_rules.metadata

<a id="canonical-3301003113312330-2102210322322323-0302133112133201-2303333111222011-1302111133110310-0320201212312311-3303313302233111-3222032210101222"></a>

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

<a id="canonical-0203110000200300-1323130211032113-2000023203101221-0310231110031102-0021021222323033-1112123030313321-1031230211231233-3331011323110300"></a>

### Direct properties for `ddos_mitigation_rules.metadata`

<a id="canonical-1011232300100310-3030023011313311-1332312030300031-1330110202303322-1000311203100011-2323310301222023-1113311121213321-2323202210321323"></a>

#### `ddos_mitigation_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0200102200220211-3121003332113211-0220212233230103-2121203103232231-1013113011123032-0002022112330330-2233113023231301-2031303020233213"></a>

<a id="canonical-1331003211002222-2320311011012330-3112201200332122-1232333013030131-1321023211003322-3120132320132233-1211333312213312-1220011223011321"></a>

#### `ddos_mitigation_rules.metadata.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1333210303310103-2202102333202012-2110112132031220-3101331311300122-0321301032320110-1231002300030202-0203120323320030-2333120320300331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- default_pool

<a id="canonical-2113122000131113-2331303012231120-1313102301020223-2110331202012112-1022100313100132-0300120312032331-3001102320000100-1333210220033320"></a>

Type: `"single"`. Computed.

\[OneOf: default\_pool, default\_pool\_list; Default: default\_pool\] Configuration parameter for
default pool.

Additional upstream details:

Shape of the origin pool specification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_port_choice": "[\"health_check_port\",\"same_as_endpoint_port\"]",
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

OneOf alternatives in this subsection:

- [default_pool](data-sources--http_loadbalancer--reference--group-015.md#canonical-2113122000131113-2331303012231120-1313102301020223-2110331202012112-1022100313100132-0300120312032331-3001102320000100-1333210220033320)
- [default_pool_list](data-sources--http_loadbalancer--reference--group-017.md#canonical-1130033132313333-3221230023302310-2121330233203033-0200002030033001-3110123103121011-1230303201033313-3011113212231322-1133312312033333)

Select alternatives according to the provider validators above.

<a id="canonical-3320102310020121-1003113131120100-0201333101311300-1233103331122112-3002320003302310-2123202330323222-0012131322232332-0001311020013322"></a>

### Direct properties for `default_pool`

- [advanced_options](data-sources--http_loadbalancer--reference--group-016.md#canonical-2032131223133102-0103033032013002-1132033101022201-3120101010321133-2000232323133112-2301313020232111-2302010213111311-3223110200321013): complete subsection reference.

- [automatic_port](data-sources--http_loadbalancer--reference--group-016.md#canonical-3333201113312220-0211103023020231-1231022023021132-3130330021322021-1011333013330100-1221301202121312-1123312313021112-0032303012213003): complete subsection reference.

<a id="canonical-0202020300030231-3010110031021002-1021333103322333-2031332121030231-1012222012330000-0133002321230300-3131001032302122-3332300302303012"></a>

<a id="canonical-0320310022202021-3201030130022211-3320102010231231-2103333312003320-3302202011333311-1202232300330012-0202010200311301-0302333330100112"></a>

#### `default_pool.endpoint_selection` property

Type: `"string"`. Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`. Server applies default when
omitted.

Additional upstream details:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

Receipt-pinned upstream constraints:

```json
{
  "default": "DISTRIBUTED",
  "enum": [
    "DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1103101202101131-2010321102031001-1203203131303113-1221132101310313-0110303222121002-2222301003103220-2200132132330123-2131100233300133"></a>
