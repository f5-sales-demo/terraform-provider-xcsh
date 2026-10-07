---
page_title: "xcsh_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery reference."
---

# xcsh_discovery reference

<a id="canonical-1103000123011011-3013311320020030-1100120212033010-3310000302030332-3003033102012113-2113313032223212-2301011302230023-1103020303110313"></a>

## Direct properties for `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info`

<a id="canonical-0330002231230302-2230033132220031-2110100212122002-1101311113000202-0320230000202203-0223121030331021-1113111232210001-1103002000022112"></a>

### `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3012322212212000-3333101232331313-3122110322101001-2322320130131232-1201302232220120-0313101031220020-2203020030123230-3312302302000100"></a>

<a id="canonical-2020222101202300-3302102013212110-3300000112210313-3220211302300302-0133103023033120-3231321220221112-0201122112230202-3013130313110311"></a>

### `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0000221330012012-2302022102003100-1020030310232110-0021302131123120-0223233223020101-3101330233200313-3233000300212213-2100230312210011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.isolated` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- discovery_k8s.access_info.isolated

<a id="canonical-3211013220013210-2330200001321322-1330221213221000-1221230133203311-1101010320302122-2111222033110231-0002130112322222-3101302221111111"></a>

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

<a id="canonical-1033022012013103-3002223202110122-0023123111320030-3303111222312023-3212132023330113-2310233213012123-2013212323130130-3100102223221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.kubeconfig_url` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- discovery_k8s.access_info.kubeconfig_url

<a id="canonical-1210333310133030-3021230123230212-3320112222212210-2300223123300123-1320101120020203-0323102033302210-2233301213121203-3132023020221323"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-2103332203313133-2113022103033321-2302222001303020-2220120033232232-1222213230301121-1132312022200201-3022213111121003-0323001223213012"></a>

### Direct properties for `discovery_k8s.access_info.kubeconfig_url`

- [blindfold_secret_info](data-sources--discovery--reference--group-002.md#canonical-3212211310223113-2210302111113110-1201312323203133-0023331301303132-1210301202320031-1100223012033000-2200102022223033-1312103323301301): complete subsection reference.

- [clear_secret_info](data-sources--discovery--reference--group-002.md#canonical-3202030220212203-1321100010331233-1112030230010110-0113133200302301-3023002223121301-1001330221132223-0211133222132211-2210002100013122): complete subsection reference.

<a id="canonical-3212211310223113-2210302111113110-1201312323203133-0023331301303132-1210301202320031-1100223012033000-2200102022223033-1312103323301301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--reference--group-002.md#canonical-1033022012013103-3002223202110122-0023123111320030-3303111222312023-3212132023330113-2310233213012123-2013212323130130-3100102223221311)
- discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info

<a id="canonical-2103322020011101-3313203020002120-3201221120300203-3122112302300230-2200211003322202-0233030112010113-1201103110030311-0202310123033202"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2213032303001032-2133103021013301-1201001120012231-0120023120230113-0021111002032210-3131312002132333-2001301221021133-1133323122111030"></a>

### Direct properties for `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info`

<a id="canonical-2203221012022212-1202130331031000-0021222100332331-2003231033230210-3213212330322012-0123020231101130-3013221010003102-1033130102121221"></a>

#### `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-0010123321311301-0323231023212031-1302202131310310-1332303000122013-1033111003102001-1222201210323112-2010301033001200-1203010232113121"></a>

<a id="canonical-2000202023131331-3122313111332223-3112102233310131-0101320203110133-2001221111123323-0102232321302122-3332221113102033-0033200103330201"></a>

#### `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0112123011032013-2132213100130212-1311203132233320-3313201030022113-2220110323221123-1102203120301210-0111121311203023-1210123101123211"></a>

<a id="canonical-1113132003323002-0021030233323113-3132301312110021-1313100101313112-1313010021110113-2330001302202132-1201320010123312-1131000221102032"></a>

#### `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-3202030220212203-1321100010331233-1112030230010110-0113133200302301-3023002223121301-1001330221132223-0211133222132211-2210002100013122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.kubeconfig_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--reference--group-002.md#canonical-1033022012013103-3002223202110122-0023123111320030-3303111222312023-3212132023330113-2310233213012123-2013212323130130-3100102223221311)
- discovery_k8s.access_info.kubeconfig_url.clear_secret_info

<a id="canonical-0231333002203201-2101132012300331-2010311023131333-3120021323222211-2103101131202033-3222012113302022-1012323222310031-1200133200131213"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0032123031302310-0133220032333212-2012212011322013-1012033331021230-3011320111031132-2223123310311233-1230310230311231-1210202313310332"></a>

### Direct properties for `discovery_k8s.access_info.kubeconfig_url.clear_secret_info`

<a id="canonical-0032001230230323-3112321020310322-1022010310002021-3323121022001021-3211321100300010-2210201010311030-3012213003132222-3300300120202121"></a>

#### `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1300231300130022-3233030222011120-2331321013211122-1022322023002030-0131020330030111-1022132310222111-2311003233001132-0300103032200323"></a>

<a id="canonical-1320122201123013-0232022233200001-3331202212033120-2203022111022320-0313230003312231-3233232112013311-1131133122030121-2120101130212212"></a>

#### `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1312013231130203-2031112120032320-1112122010012202-1211033002212132-2303102311110103-0201313021012021-0000203103133233-1201230330203320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.reachable` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- discovery_k8s.access_info.reachable

<a id="canonical-3130223331201311-2313022032233003-1121212101330122-0033110121001230-2122133332111102-0102323321310103-0010133223330300-2121103123311021"></a>

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

<a id="canonical-0111030212233202-0132113112330210-3102303023102000-2311213102011302-2231201111012332-2233001123200131-3010310220333110-2310213013031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.default_all` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- discovery_k8s.default_all

<a id="canonical-3210321213011110-2311302010321030-2122310222321021-0033023130232021-3001203030020011-3203122223013103-2233212222031122-3023033123221030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default all.

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

<a id="canonical-3303013301213222-3332332202321221-2311112031001231-0323332330033132-2230203122010120-1131213122133310-1033023232121000-2320222012200021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.namespace_mapping` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- discovery_k8s.namespace_mapping

<a id="canonical-2001322221023102-2130230202311020-3233201020110333-3311030322102203-3133111323220001-0000023002021113-3120203211322012-3230020231112012"></a>

Type: `"single"`. Computed.

Select the mapping between K8s namespaces from which services will be discovered and App Namespace
to which the discovered services will be shared.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1310301132011303-0002123333100020-2313121001012212-0311130233033332-0300121212202000-2210111130320212-3111033001100123-1301133020031221"></a>

### Direct properties for `discovery_k8s.namespace_mapping`

- [items](data-sources--discovery--reference--group-002.md#canonical-2230122013023000-3200231101230132-0122123232020331-2200331033033122-1203030111321231-0132012113103322-0010011312102120-2112231103110320): complete subsection reference.

<a id="canonical-2230122013023000-3200231101230132-0122123232020331-2200331033033122-1203030111321231-0132012113103322-0010011312102120-2112231103110320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.namespace_mapping.items` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.namespace_mapping](data-sources--discovery--reference--group-002.md#canonical-3303013301213222-3332332202321221-2311112031001231-0323332330033132-2230203122010120-1131213122133310-1033023232121000-2320222012200021)
- discovery_k8s.namespace_mapping.items

<a id="canonical-2132211121203312-1232000332123322-2321011010003113-0223211010221100-1132322000322210-3023223021101123-0021010123333033-0202131111213100"></a>

Type: `"list"`. Computed.

Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be
mapped to a single App Namespace, which is determined by the first matched regular expression.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0232322201103121-3310330031332120-3223133012313031-0033032221132222-3301232210011110-1120121332200230-0132102302300230-1303131130301201"></a>

### Direct properties for `discovery_k8s.namespace_mapping.items`

<a id="canonical-0320032001030203-2303232300302021-3210203012111012-2303230001210303-2110230230213102-3333120022312133-2012312103100202-3223222020012232"></a>

#### `discovery_k8s.namespace_mapping.items.namespace` property

Type: `"string"`. Computed.

F5XC Application Namespaces. Select a namespace.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0222322220213002-0132130023220100-3202202323203131-3200201001102201-1302012200331001-2122301300332220-3130230231201313-2031032130122311"></a>

<a id="canonical-1313231011312102-0030131212210210-0110221103332322-0131132020303010-1323301030333311-0120020321300201-3103230312132013-0003033002330022"></a>

#### `discovery_k8s.namespace_mapping.items.namespace_regex` property

Type: `"string"`. Computed.

The regular expression here will be used to match K8s namespace(s).

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2000131012030001-1310223011223322-3222122000322101-3130330332001002-2203022320301110-1121220212302102-1202210220223323-1101112222203112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- discovery_k8s.publish_info

<a id="canonical-0121022203221333-2232012212222101-0221331211133021-2301301002013333-0211010013100101-2022020131010031-2220113021102211-0330022332012122"></a>

Type: `"single"`. Computed.

Configuration parameter for publish info.

Additional upstream details:

K8s Configuration to publish VIPs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"dns_delegation\",\"publish\",\"publish_fqdns\"]"
}
```

<a id="canonical-0212213030303322-1332201232313032-0013303112133133-0012023111013301-0222201010301302-0011213130330003-2330132013021322-0010132322103213"></a>

### Direct properties for `discovery_k8s.publish_info`

- [disable_spec](data-sources--discovery--reference--group-002.md#canonical-0322020333110210-3031113323033111-1221103002213113-0233330130321131-1310120301120022-1330213130332110-2312101313113230-3013332111320101): complete subsection reference.

- [dns_delegation](data-sources--discovery--reference--group-002.md#canonical-0021110011333022-3031323120010221-3331021222110233-2032301330012323-0021101112313013-0221323000310213-0010322112333203-2312212121122233): complete subsection reference.

- [publish](data-sources--discovery--reference--group-002.md#canonical-0231330210033210-2333231230323101-2113020031321112-2300002130220010-0221021003322202-2013133032323202-3031333202303110-3302312021010333): complete subsection reference.

- [publish_fqdns](data-sources--discovery--reference--group-002.md#canonical-2232120000331013-1132210320233210-3313203301013320-1032222021113020-2322321122021331-1111013230001100-2230102132101002-1002322022031012): complete subsection reference.

<a id="canonical-0322020333110210-3031113323033111-1221103002213113-0233330130321131-1310120301120022-1330213130332110-2312101313113230-3013332111320101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info.disable_spec` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.publish_info](data-sources--discovery--reference--group-002.md#canonical-2000131012030001-1310223011223322-3222122000322101-3130330332001002-2203022320301110-1121220212302102-1202210220223323-1101112222203112)
- discovery_k8s.publish_info.disable_spec

<a id="canonical-1102232033130113-1120111201131032-3320202232120321-3213020130101131-0203300103110113-1222013203312221-3132212303033223-1033323111230300"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021110011333022-3031323120010221-3331021222110233-2032301330012323-0021101112313013-0221323000310213-0010322112333203-2312212121122233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info.dns_delegation` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.publish_info](data-sources--discovery--reference--group-002.md#canonical-2000131012030001-1310223011223322-3222122000322101-3130330332001002-2203022320301110-1121220212302102-1202210220223323-1101112222203112)
- discovery_k8s.publish_info.dns_delegation

<a id="canonical-1021331020323302-2333002120113112-2002133202131312-0201020103322210-2120122033320121-2002102100300302-2222012103230032-1303203032022232"></a>

Type: `"single"`. Computed.

Configuration parameter for DNS delegation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1322301112023211-0221100202102033-0310310332130212-2121302121333233-3222332032231001-0220212021003133-2321021321302222-1302010330221113"></a>

### Direct properties for `discovery_k8s.publish_info.dns_delegation`

<a id="canonical-0231112021101120-0033030231033102-1232203020121220-3223313222011223-2202200003021312-2023322011030202-0001100311122320-0220021103230121"></a>

#### `discovery_k8s.publish_info.dns_delegation.dns_mode` property

Type: `"string"`. Computed.

\[Enum: CORE\_DNS|KUBE\_DNS\] Two modes are possible CoreDNS: Whether external K8s cluster is
running core-DNS KubeDNS: External K8s cluster is running kube-DNS. Possible values are
\`CORE\_DNS\`, \`KUBE\_DNS\`. Defaults to \`CORE\_DNS\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "CORE_DNS",
  "enum": [
    "CORE_DNS",
    "KUBE_DNS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0130121213233301-3133321120322023-0223202200133303-1003031123333213-0031122030232322-2120021012122120-2111230031013132-3230310213131003"></a>

<a id="canonical-2012300010203112-0112200102103322-2100120311032030-1303012021101010-3222103110210111-0330132003331232-3310123231133100-2301121203103232"></a>

#### `discovery_k8s.publish_info.dns_delegation.subdomain` property

Type: `"string"`. Computed.

The DNS subdomain for which F5XC will respond to DNS queries.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0231330210033210-2333231230323101-2113020031321112-2300002130220010-0221021003322202-2013133032323202-3031333202303110-3302312021010333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info.publish` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.publish_info](data-sources--discovery--reference--group-002.md#canonical-2000131012030001-1310223011223322-3222122000322101-3130330332001002-2203022320301110-1121220212302102-1202210220223323-1101112222203112)
- discovery_k8s.publish_info.publish

<a id="canonical-0202130231222233-1321100212331103-3010120031323110-1230321133033033-3130320232310221-1312332213131320-3212303032020003-3212331001323000"></a>

Type: `"single"`. Computed.

K8SPublishType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2123102330102102-0331331003323121-3110000013301102-3332001330123102-3231312012021032-3120101032212022-3010133323012133-3323032032003321"></a>

### Direct properties for `discovery_k8s.publish_info.publish`

<a id="canonical-3223312102132300-3103022022023301-2020030101033011-1122332231222112-3022201011213010-0323023033222211-3020020103130003-2103321320313113"></a>

#### `discovery_k8s.publish_info.publish.namespace` property

Type: `"string"`. Computed.

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2232120000331013-1132210320233210-3313203301013320-1032222021113020-2322321122021331-1111013230001100-2230102132101002-1002322022031012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info.publish_fqdns` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.publish_info](data-sources--discovery--reference--group-002.md#canonical-2000131012030001-1310223011223322-3222122000322101-3130330332001002-2203022320301110-1121220212302102-1202210220223323-1101112222203112)
- discovery_k8s.publish_info.publish_fqdns

<a id="canonical-0330202321302013-2222011222223213-1230333311230000-2303221022122201-1301000322000302-3200112332332011-1221103232210031-0123011232332202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for publish fqdns.

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

<a id="canonical-3022232112013312-1120023000331333-0002001130200231-2023010022021302-2321320033321023-1313213310100331-3000313221232320-1132122300112130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_cluster_id` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- no_cluster_id

<a id="canonical-2330131222130101-1223121232130000-3310203111002300-3111030300122323-2032022200033233-3120020220301310-3232021301122000-1013013012212302"></a>

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

<a id="canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- where

<a id="canonical-1001023320211113-3130223330020112-0313320233030103-1120010333033202-2102010130130203-2210230232101121-0112310133203302-0100332000302133"></a>

Type: `"single"`. Computed.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

<a id="canonical-0030112103021301-0230030003303303-0012303011033000-1301201233202021-0303120221000320-3203200201130221-2211131133213331-3113123131300211"></a>

### Direct properties for `where`

- [site](data-sources--discovery--reference--group-002.md#canonical-1131111212212130-2300200213120302-1010130223111031-2031223120011323-3032013013110101-2231203231031223-1132131231333231-3003122031021212): complete subsection reference.

- [virtual_network](data-sources--discovery--reference--group-002.md#canonical-2000100203122133-0301232002323310-3233310132013200-2022221330110313-1210122133103010-3121123023112123-0033000321111322-1210332230303323): complete subsection reference.

- [virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122): complete subsection reference.

<a id="canonical-1131111212212130-2300200213120302-1010130223111031-2031223120011323-3032013013110101-2231203231031223-1132131231333231-3003122031021212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- where.site

<a id="canonical-2332030123000123-1332002322102101-0303233200130021-2101101032121320-3201023103031301-0202301323132301-0310021233310111-0331123230200211"></a>

Type: `"single"`. Computed.

This specifies a direct reference to a site configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-0310100001332221-0302330010113112-3101100322300332-1103330013313023-1101333222203100-3032000033303133-1023211313322330-1133211103032302"></a>

### Direct properties for `where.site`

- [disable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-2232230002030311-0013122112232130-2301203002020322-2023331313233001-1001220222231000-0213220233130211-0101220300203032-1121213233200231): complete subsection reference.

- [enable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-2220130012323232-0022222012102312-3122111302020020-1321013332322001-2201322133001002-0222201010223233-0230122022113231-3032012320302332): complete subsection reference.

<a id="canonical-3011001110022122-1021013002103021-3111220011103030-3133323223130323-0001321000033021-1211300123300233-1323313130122100-0020221132001121"></a>

<a id="canonical-3103030211230213-2302333213233100-0230022033032030-3001301013110131-0102032030211312-2310101003310132-1102230032303030-2332300220131221"></a>

#### `where.site.network_type` property

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--discovery--reference--group-002.md#canonical-0131223101303030-2100222021133330-2100033110332033-3201131000100200-3010112003213001-1211212130030102-2111211031100100-3132121230021312): complete subsection reference.

<a id="canonical-2232230002030311-0013122112232130-2301203002020322-2023331313233001-1001220222231000-0213220233130211-0101220300203032-1121213233200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.site](data-sources--discovery--reference--group-002.md#canonical-1131111212212130-2300200213120302-1010130223111031-2031223120011323-3032013013110101-2231203231031223-1132131231333231-3003122031021212)
- where.site.disable_internet_vip

<a id="canonical-2030031300113323-0010110223221003-2023031302000332-0202232312331130-3200003300003210-2213312023200100-3102003132012132-0210323030030333"></a>

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

<a id="canonical-2220130012323232-0022222012102312-3122111302020020-1321013332322001-2201322133001002-0222201010223233-0230122022113231-3032012320302332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.site](data-sources--discovery--reference--group-002.md#canonical-1131111212212130-2300200213120302-1010130223111031-2031223120011323-3032013013110101-2231203231031223-1132131231333231-3003122031021212)
- where.site.enable_internet_vip

<a id="canonical-1002020120312333-3331230301213201-0231232130130323-0323320202000020-2231202331021120-3300103120102220-3201223022100020-3113100130213023"></a>

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

<a id="canonical-0131223101303030-2100222021133330-2100033110332033-3201131000100200-3010112003213001-1211212130030102-2111211031100100-3132121230021312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.ref` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.site](data-sources--discovery--reference--group-002.md#canonical-1131111212212130-2300200213120302-1010130223111031-2031223120011323-3032013013110101-2231203231031223-1132131231333231-3003122031021212)
- where.site.ref

<a id="canonical-0013011012032002-2112121110233203-1122200300130232-2012333230310231-1320333102200021-0322033323202123-2103333231012031-0330103020110233"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1103202203330123-3213201022120231-0102202332130300-0200111132031122-0110002313310332-0122000202031300-1313110230100032-3123130200313321"></a>

### Direct properties for `where.site.ref`

<a id="canonical-0130011331102230-3220113112221202-2333300113030221-1020201321030101-3102312302100330-0300330110222101-3302220111031132-0203213013203210"></a>

#### `where.site.ref.kind` property

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

<a id="canonical-2232020130130030-0030222113021212-3231232032023230-0303213221011011-2301001333130122-0323313302132213-3201001112203333-3131110003103033"></a>

<a id="canonical-2030101200013011-1300333301000112-1203120033101302-3332320011232010-3303133112033020-0013321203233310-1123220232030102-0031130111011320"></a>

#### `where.site.ref.name` property

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

<a id="canonical-2031213121232012-3222123110002322-0211200101231210-1232110301112322-1103320011211231-2331210032122001-1022202030223313-0031320030332202"></a>

<a id="canonical-2030310322223320-3121122301011022-0110022210210001-0302123323132223-0203120303221032-0023331133330023-0021032101110011-2233122130103023"></a>

#### `where.site.ref.namespace` property

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0110012123223223-1023210333000311-0210323110232032-1311130202213332-1310320022122221-3112213323230121-0111233233000102-3022221031030330"></a>

<a id="canonical-2102312323112310-2023032112020201-3112233130231102-1330230231012331-1302223333333133-0323310331302201-2113003102031131-1001332322113203"></a>

#### `where.site.ref.tenant` property

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

<a id="canonical-3312030213330313-1332031321031101-2323013000300013-3323200111023202-2323012322230033-0012222230233033-3202202220000211-2202120113211123"></a>

<a id="canonical-0321013112103321-3113013221011111-2230301213301000-0332220232102330-2030003012002220-2310231222210033-1230233110112000-0223030033321200"></a>

#### `where.site.ref.uid` property

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

<a id="canonical-2000100203122133-0301232002323310-3233310132013200-2022221330110313-1210122133103010-3121123023112123-0033000321111322-1210332230303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- where.virtual_network

<a id="canonical-0223030102121002-2123230021330200-3103200211323201-0233313121200102-3030330120233330-2303002133322202-1203330011023332-3221332231311202"></a>

Type: `"single"`. Computed.

This specifies a direct reference to a network configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1201200132211110-3113322023031203-3230302013021210-2102102103133200-1303310331201231-3201123032102012-3002013310331233-0321003101212001"></a>

### Direct properties for `where.virtual_network`

- [ref](data-sources--discovery--reference--group-002.md#canonical-0221301333220223-3213322202103331-1212321320032223-0112231031101100-0101112320131210-3323311320113022-2131130300221222-0232102211101321): complete subsection reference.

<a id="canonical-0221301333220223-3213322202103331-1212321320032223-0112231031101100-0101112320131210-3323311320113022-2131130300221222-0232102211101321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network.ref` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.virtual_network](data-sources--discovery--reference--group-002.md#canonical-2000100203122133-0301232002323310-3233310132013200-2022221330110313-1210122133103010-3121123023112123-0033000321111322-1210332230303323)
- where.virtual_network.ref

<a id="canonical-1311001320310131-0100302300002102-2320032213230102-2100003002213030-0323131120301301-1031233201321211-2122331331032011-1210210212110213"></a>

Type: `"list"`. Computed.

Reference. A virtual network direct reference.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2221013000030012-3022321102312102-1302131332233133-1012203133313021-1221030130013221-3311002011301130-2010121013000311-0201132202112022"></a>

### Direct properties for `where.virtual_network.ref`

<a id="canonical-2112002132212001-2003233120203311-3032321100120121-3220031312213212-0121202212111221-2313211032311021-2012201330312103-2211232201111111"></a>

#### `where.virtual_network.ref.kind` property

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

<a id="canonical-0102103211220203-0301333001112111-1032012320200032-1202222312310012-1311322120002311-2123222120321221-3021303023302110-3210123131013230"></a>

<a id="canonical-3210130020132330-2303310201023011-3011330001223211-2320112000031123-2200102331310100-1033230331313103-1013320130100112-2111202332330330"></a>

#### `where.virtual_network.ref.name` property

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

<a id="canonical-2121022210310020-2010032232022003-3003333101133323-1002203312000231-3232312022221021-1020333301101231-3110112210120212-0003210012112232"></a>

<a id="canonical-3313100012023203-3202021230301021-2102121032030022-2210032300211002-3033122300300122-0300331202000132-3323003312323223-2022222332302132"></a>

#### `where.virtual_network.ref.namespace` property

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0122032130001322-1231202123102221-3121021003121013-1010113003330231-0231211122130013-1223130121002202-2201101230210001-3201023333311010"></a>

<a id="canonical-0330103230210002-2023131122013200-3000313131202303-3332020130030231-3212123120231201-3022003312202221-2202130203022010-3120213322311130"></a>

#### `where.virtual_network.ref.tenant` property

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

<a id="canonical-2303322321221301-3110012132223100-1333011120031230-1031222221000231-3212331333012313-0331033003122111-1200313121011202-0232212310312122"></a>

<a id="canonical-3001021000223021-1121113220222231-1300113323123110-1321211202120312-3120213012130021-2111002001131322-2102333203333221-0112131030130300"></a>

#### `where.virtual_network.ref.uid` property

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

<a id="canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- where.virtual_site

<a id="canonical-3121122332022100-2322022213221212-0030221202020033-0013010231331211-3232021320122012-3112321021330201-0311211221311133-0333021120321313"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-1233032212031121-2312310231023322-3133221322310012-3021132001032230-0021030133022310-3230012002121203-0213012222100223-0112103303011213"></a>

### Direct properties for `where.virtual_site`

- [disable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-2310001332100202-0111030221230112-1322023213221100-2012333111000330-1232321320031301-0131300201321203-2321211122012232-1333223332300122): complete subsection reference.

- [enable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-2020131222103111-2223331232232312-3212101302212223-1200010333003320-3310010232102320-3303103331031332-1000011103202330-2011022321023222): complete subsection reference.

<a id="canonical-3320020330211301-1102122020300322-2013331030233300-3113002012313023-0313002001233231-3100203330121111-1331013312002022-0000232011033021"></a>

<a id="canonical-2122313100301200-0130300323013330-0102203110012001-1231012023200023-0120003210331312-1201210111002321-0012121230231333-3103020123000310"></a>

#### `where.virtual_site.network_type` property

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--discovery--reference--group-002.md#canonical-2112101231113022-3113100010022100-2233130012103200-0011131222221220-1320202012310013-3232220012222010-1222300020330213-3022121113131011): complete subsection reference.

<a id="canonical-2310001332100202-0111030221230112-1322023213221100-2012333111000330-1232321320031301-0131300201321203-2321211122012232-1333223332300122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122)
- where.virtual_site.disable_internet_vip

<a id="canonical-3121010010312310-2012012330223202-0133303311112213-1300211123033101-3112301131231222-1200110102021221-3123330030130012-0323213030231012"></a>

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

<a id="canonical-2020131222103111-2223331232232312-3212101302212223-1200010333003320-3310010232102320-3303103331031332-1000011103202330-2011022321023222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122)
- where.virtual_site.enable_internet_vip

<a id="canonical-3011122123133320-1303211013132313-1112123011003222-0003232113101030-3110122102333102-1322203212322021-1021003202331110-3013333001212003"></a>

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

<a id="canonical-2112101231113022-3113100010022100-2233130012103200-0011131222221220-1320202012310013-3232220012222010-1222300020330213-3022121113131011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.ref` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110)
- [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-2133223112022020-1203131112101120-1123221132012123-1012101130010321-2213300002121320-2233022032323113-3111330010231222-0322222111031122)
- where.virtual_site.ref

<a id="canonical-0003002230031332-3213000322300231-0312212202233210-1002020331313320-2002103231103312-2201201210233320-0032021202013310-3001311232330013"></a>

Type: `"list"`. Computed.

Reference. A virtual\_site direct reference.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2001330203223203-3001022021322213-1013232321311303-1200220100010130-1221233323111001-3002001231020222-2102200211202010-3313320221202301"></a>

### Direct properties for `where.virtual_site.ref`

<a id="canonical-1230000133212313-1302201130321013-0300211302313300-3103213201302210-2101020230310013-0223031133213113-3320202002211122-3131330010330301"></a>

#### `where.virtual_site.ref.kind` property

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

<a id="canonical-2121102110133313-3103303223012200-3012220100312133-1311300020001322-3113312213223103-1301113333120331-3303113330100031-1311031232233200"></a>

<a id="canonical-2003131210022320-0103031120202002-3113112123202003-1312320021331003-3213102311102033-2021133223323032-0203333022101113-3202233332002211"></a>

#### `where.virtual_site.ref.name` property

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

<a id="canonical-0312223301202212-2200221213013001-0230121133112010-0313300000303022-1332232122113310-1102020233010220-2111030231100022-2211212130213320"></a>

<a id="canonical-0013031212112310-0030133023131203-1232002333222021-3013213031123022-0101131310221331-3220311100020202-1200130133201213-0313133120020200"></a>

#### `where.virtual_site.ref.namespace` property

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320113201012322-3011103012123211-1121033302223312-0032222302222130-3012132013310330-0020323320333332-2010230311000323-2303301212020011"></a>

<a id="canonical-3131103003010030-2211322213103301-0233322231213102-0001013112101321-3130213211012203-2031100303303322-0001130230202200-3120231331320133"></a>

#### `where.virtual_site.ref.tenant` property

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

<a id="canonical-3212130122302310-0120202023312212-3301200101013110-3102303122203200-3003002201111120-2001020013020220-3032223301302123-3101011122011122"></a>

<a id="canonical-1333113210230112-0130212231331303-0032330220321132-0200102013321232-3113200311303213-1212113110023200-3133022321331100-2211332302012211"></a>

#### `where.virtual_site.ref.uid` property

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
