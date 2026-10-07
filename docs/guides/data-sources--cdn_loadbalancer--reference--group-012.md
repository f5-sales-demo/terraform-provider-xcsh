---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1113003211013103-2102201110023000-0220203221113302-1200123023113023-2021000001322313-2123020102321113-2220013010330212-1203210121020300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1113100120020201-0123033213020223-2302310033123110-2201313231300232-0101133330302212-2012122210031301-1112131221311003-0212331320023121)
- https.tls_cert_options.tls_inline_params.tls_config.medium_security

<a id="canonical-0030211313323331-3100210121120213-3121002320111023-3101020113222002-1133320011123331-1233002211032131-3012202010031321-2003100002220000"></a>

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

<a id="canonical-1020121201233003-1212133021201001-2331021023212310-0111220113103030-3222211102322103-2331113010121022-0102011311113033-3111232102002113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- https.tls_cert_options.tls_inline_params.use_mtls

<a id="canonical-0103220012121010-1322200331303201-1310313021301233-1333121323033212-1230212200302023-2031130222120113-0210310010223030-0022130221001132"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-2033312121112200-3001120212130201-0101223333032210-0011111013332001-3100122231101302-1100112300013012-2313003132111231-0322333201330303"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.use_mtls`

<a id="canonical-2330121200221330-1201210000032001-1321232220321231-1013210020020012-2013102133311211-0303103000103301-1302130012300232-1030332023112231"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0120310201221301-0223121023223110-1210220001111131-3121203100010211-1032111131331222-3202011122233101-1122320332030332-0221321101000323): complete subsection reference.

- [no_crl](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0310332132132021-1112113001100331-3133111012201021-2110130113103231-1221121230130201-1201021001213221-0201101022122132-0010331123021330): complete subsection reference.

- [trusted_ca](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3132322112101321-1003112123012232-2103211321013131-1221123023023023-1101330210200330-0123110211232201-3132313110322002-3110211021222110): complete subsection reference.

<a id="canonical-2220002222221332-2130112112020101-3210010100113333-2003133212321221-3203000232200200-0231033013222132-2231320010313211-0233201212023300"></a>

<a id="canonical-0222310031231002-1111101012301303-0320032111213111-2000213312010201-1031000110023111-0331311120001121-3011102222001232-1312222213211200"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1131120333121322-2030012311223332-2300333301020123-0112001102020031-0033103310221302-1202332033103220-0012010212230321-2302013001332321): complete subsection reference.

- [xfcc_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2212120300203111-3013000322302121-3033212123323210-3131003001001012-1311103111221323-3011301102103212-2003123221230321-3132131123333220): complete subsection reference.

<a id="canonical-0120310201221301-0223121023223110-1210220001111131-3121203100010211-1032111131331222-3202011122233101-1122320332030332-0221321101000323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1020121201233003-1212133021201001-2331021023212310-0111220113103030-3222211102322103-2331113010121022-0102011311113033-3111232102002113)
- https.tls_cert_options.tls_inline_params.use_mtls.crl

<a id="canonical-0020332012310033-0201023211220211-2010203110031312-2321222233120220-0322211213313221-3232002102321303-0213201222133000-1113220330031312"></a>

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

<a id="canonical-0010200222021003-3200012232331213-1000230330131012-3302211002331101-2211220003003020-3131313113113021-2231122022031211-3233121202031322"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.use_mtls.crl`

<a id="canonical-1033331022132301-2201212012132022-3000302203023010-0123221333111223-3000132203320002-2133310020011313-2022101301302133-3310001103023323"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2102330223031322-3112122311221211-0003103320101013-0313103230322101-2020102033102321-2322301320131010-1121230201303130-2200222102120221"></a>

<a id="canonical-2001221300003302-0321320311130002-1331121032323210-3331100112312312-2130310212123330-1010202031111122-0001202212032212-3001200022030330"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.crl.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1123003130311210-2322101133011212-1213123210211200-0332221011130012-3123330330003323-0301103112313011-1211300233103221-3232100010233123"></a>

<a id="canonical-3030111223220213-3023011302310232-2323131011133310-2031221113021000-0031202312211002-0212130301033230-3202122102333332-3122111201211331"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.crl.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0310332132132021-1112113001100331-3133111012201021-2110130113103231-1221121230130201-1201021001213221-0201101022122132-0010331123021330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1020121201233003-1212133021201001-2331021023212310-0111220113103030-3222211102322103-2331113010121022-0102011311113033-3111232102002113)
- https.tls_cert_options.tls_inline_params.use_mtls.no_crl

<a id="canonical-0023101100121233-3112213033113221-1211120313032023-2200031202201221-2333001221231112-1302122021030200-2321011031101303-1002103301222213"></a>

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

<a id="canonical-3132322112101321-1003112123012232-2103211321013131-1221123023023023-1101330210200330-0123110211232201-3132313110322002-3110211021222110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1020121201233003-1212133021201001-2331021023212310-0111220113103030-3222211102322103-2331113010121022-0102011311113033-3111232102002113)
- https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca

<a id="canonical-1102033103202021-1201031210021312-3313132020023331-2111310302301002-0220112312120210-2232002023030233-1032030203120030-1031003212212003"></a>

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

<a id="canonical-1200333010113231-1220032021233321-2221032011103333-0202131301120210-1233313203230023-0303201003321120-2113131103100321-2302021222213311"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca`

<a id="canonical-3220222020321011-3013011001111333-3301013230230123-1020233201131303-2220302102322132-2302030121122102-2130333223110032-3321211013213033"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1332020220030010-1232032311023023-1213200231330300-3302110211310120-1112111120222322-1322232123301233-1203223202030201-2220120131021310"></a>

<a id="canonical-2022321030020233-1221311031010231-1130333310002122-0121210310123323-2122212132021013-3303232033233103-0233000133301232-2203332003022303"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2021123121330321-2030312112313203-2132120023021033-3311122222300133-2031332033313022-3332022001002323-0003003213222031-1100033033011000"></a>

<a id="canonical-1113310303200322-3202222101333221-0031103022200111-3021300002120232-2021130232103232-1101300332033131-3220101110212232-2200022200123310"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1131120333121322-2030012311223332-2300333301020123-0112001102020031-0033103310221302-1202332033103220-0012010212230321-2302013001332321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1020121201233003-1212133021201001-2331021023212310-0111220113103030-3222211102322103-2331113010121022-0102011311113033-3111232102002113)
- https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled

<a id="canonical-3322233102202311-2303311220133012-0132131213212000-0122300101221211-3320022122212103-3111022320023320-0220332023233232-1022020122203123"></a>

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

<a id="canonical-2212120300203111-3013000322302121-3033212123323210-3131003001001012-1311103111221323-3011301102103212-2003123221230321-3132131123333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3320031111223313-3223233303310013-3311313213013221-0331200101023022-3313123310020112-3232133301022211-2330003012130111-1012330031221211)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313032021211010-0011012123033300-3301113220311230-3323120333011222-2310123131130303-3113023231310330-2200321111112213-1211002321212010)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1020121201233003-1212133021201001-2331021023212310-0111220113103030-3222211102322103-2331113010121022-0102011311113033-3111232102002113)
- https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options

<a id="canonical-0011323300000111-0323102203222213-1110112102010031-1312032011230033-2100200020222132-1202013030020100-0121110203012013-0300031130110321"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3011021111030212-3112121002221122-1011301212131312-2203030302222100-3320030322120000-3002310203203222-1132311303123113-2030110223031211"></a>

### Direct properties for `https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options`

<a id="canonical-1131000103013311-3332311100100320-2022230000312130-2111130200322132-3200331120322201-1210001130210132-0311230233000121-0301323113101101"></a>

#### `https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-1022110223113311-1122120102013303-2121023120231002-0103032202213002-0201010020313211-3100130013113312-0132133320111300-3120303011211300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- https_auto_cert

<a id="canonical-1133023022021021-3123031201112323-2013213230333010-2121333120032213-2302232102103320-3300321310201302-0211002132112210-0203032111320311"></a>

Type: `"single"`. Computed.

Choice for selecting HTTPS CDN distribution with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0323213233232221-1301003233102303-3302120322033103-3310131023223313-0323110231201300-1031332223013202-0321210310132333-3102233001333333"></a>

### Direct properties for `https_auto_cert`

<a id="canonical-2021223011222111-1200113120032312-3030033023120000-1221012223131331-0120133122000111-0333220312022212-0333232130332130-1311202012210122"></a>

#### `https_auto_cert.add_hsts` property

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1011122302100311-0032322323210131-0112302221030021-2303321111220223-3213032211101211-2232320110203212-0223022201112220-3021331113200321"></a>

<a id="canonical-2132220310020101-3103020001200330-0222223111230303-0330110131123232-1332231030201023-3311031320022020-0320332312211312-3122210013111032"></a>

#### `https_auto_cert.http_redirect` property

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_config](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2200301313132132-1210221320321221-0030233012032220-1230220133023301-3311013122321210-0211333221233100-1300322331113102-1010120311010030): complete subsection reference.

<a id="canonical-2200301313132132-1210221320321221-0030233012032220-1230220133023301-3311013122321210-0211333221233100-1300322331113102-1010120311010030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1022110223113311-1122120102013303-2121023120231002-0103032202213002-0201010020313211-3100130013113312-0132133320111300-3120303011211300)
- https_auto_cert.tls_config

<a id="canonical-1313220100011113-1310133233113011-0300212123120121-3110033213012303-1132200332012111-2212012130002332-1302300112110003-3212300200223023"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"tls_11_plus\",\"tls_12_plus\"]"
}
```

<a id="canonical-1131033000232231-2301213330202332-2110202001121302-2021212130201132-2020231313123003-0113033201110202-3012213223231133-0031113203022212"></a>

### Direct properties for `https_auto_cert.tls_config`

- [tls_11_plus](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0032033121013330-2020220321131013-3112222001132103-0303322031033212-0021221232030222-3330032122333303-0313113001111113-1213301102233323): complete subsection reference.

- [tls_12_plus](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0103201312023121-2201330100100331-1202213303212320-0231123100100333-3102131211011133-1322000312211223-3223301000031212-3322303010233102): complete subsection reference.

<a id="canonical-0032033121013330-2020220321131013-3112222001132103-0303322031033212-0021221232030222-3330032122333303-0313113001111113-1213301102233323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.tls_11_plus` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1022110223113311-1122120102013303-2121023120231002-0103032202213002-0201010020313211-3100130013113312-0132133320111300-3120303011211300)
- [https_auto_cert.tls_config](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2200301313132132-1210221320321221-0030233012032220-1230220133023301-3311013122321210-0211333221233100-1300322331113102-1010120311010030)
- https_auto_cert.tls_config.tls_11_plus

<a id="canonical-0011012110203303-3102330000302331-2213002013223320-1300310002202021-3220233310330223-0310312210312101-2321222103133012-3330130333311203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for tls 11 plus.

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

<a id="canonical-0103201312023121-2201330100100331-1202213303212320-0231123100100333-3102131211011133-1322000312211223-3223301000031212-3322303010233102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.tls_12_plus` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1022110223113311-1122120102013303-2121023120231002-0103032202213002-0201010020313211-3100130013113312-0132133320111300-3120303011211300)
- [https_auto_cert.tls_config](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2200301313132132-1210221320321221-0030233012032220-1230220133023301-3311013122321210-0211333221233100-1300322331113102-1010120311010030)
- https_auto_cert.tls_config.tls_12_plus

<a id="canonical-2330101220313011-1020200321333110-3110213223210230-1033033320031232-1333231322131311-1130231133002012-2323133031211032-2313011011331111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for tls 12 plus.

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

<a id="canonical-2302310323202313-1111033213121230-1303220321000110-2312133012002202-1332110200301033-3033101101203021-2210102112012310-2223033331003110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `js_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- js_challenge

<a id="canonical-3330221111033030-0012113213012002-2033023113003223-1233203023311203-1000312213200220-3232232211323130-0331012203001311-2211203302323021"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0110112223012003-1111033233100230-2100310031211010-0030101303320133-3231311321111320-0030101322212220-3110000030120312-0211032102230233"></a>

### Direct properties for `js_challenge`

<a id="canonical-0323120230101122-1201200123310213-3031120303021012-1122303222301011-1131312130200121-3321112111323123-2233203002033113-2132303120221001"></a>

#### `js_challenge.cookie_expiry` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1032121311331231-3112311022310000-3131311233300033-1031303032000012-1333321112310210-0011122200132301-0033231230321013-1121003200202112"></a>

<a id="canonical-2023333133113313-2030330001032333-1303201233210210-0322101321302013-2300330103103100-0333322202311222-2212113310033233-1001001013323220"></a>

#### `js_challenge.custom_page` property

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0001023023233102-2332322213231110-1311112122131000-1331312202203110-1210323220012001-2012312012022030-2013032202133303-0133310232213111"></a>

<a id="canonical-1130321200233113-2103211020101202-1310302303023020-2213312311033300-3220233111131231-0222312123011003-1330333110000123-2133323031321321"></a>

#### `js_challenge.js_script_delay` property

Type: `"number"`. Computed.

Delay introduced by JavaScript, in milliseconds.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- jwt_validation

<a id="canonical-2133012221102320-0130230332002103-0003112110130333-3332323110231032-0213213222113331-3322023222233002-0100021001020322-1230202303233322"></a>

Type: `"single"`. Computed.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-jwks_configuration": "[\"authorization_server\",\"jwks_config\"]"
}
```

<a id="canonical-1210021101033233-1210230120331202-0010001001313012-3123202100201222-2303033132233322-1301133033103213-0010310112021310-0333130210101130"></a>

### Direct properties for `jwt_validation`

- [action](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2121132033013110-0203011222103223-3211020302222303-0223111212332002-1211112230203201-2000211102021000-2110302303321020-0031111212212330): complete subsection reference.

- [authorization_server](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0310100010103312-1233220331111211-3302232321222020-1110011322003010-2103021012033230-3033203321133230-3110202223321102-1031201001311221): complete subsection reference.

- [jwks_config](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2233311333221001-0201023333322021-0001121233331231-2203102320310202-0203310201232311-0033100133232033-2311132220330122-3013022213113102): complete subsection reference.

- [mandatory_claims](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0013131103300111-3033111303113002-3012310103320120-3013323323201023-2222300130320021-1000122302322132-3333112211031230-1332333103310212): complete subsection reference.

- [reserved_claims](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2000130323022201-1203333322333203-2223333303103023-2132200101132302-3033023313001031-2201000201221213-1102001212002322-2130102102212330): complete subsection reference.

- [target](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2003033103100212-0221322000032122-1033301212022110-3311221200010103-3112220320320320-1031002021233000-3210122333012323-1113020133332202): complete subsection reference.

- [token_location](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1010220232131121-3133122220132302-2102331001310120-1103121302102313-0210031123201300-3211132222021113-1033033022012000-0210101200200032): complete subsection reference.

<a id="canonical-2121132033013110-0203011222103223-3211020302222303-0223111212332002-1211112230203201-2000211102021000-2110302303321020-0031111212212330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- jwt_validation.action

<a id="canonical-1023000110021232-0221301002220212-2032111110220203-0220230001321021-0211303230213332-3031002020003210-0302023021213302-2211113010003300"></a>

Type: `"single"`. Computed.

Action

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

<a id="canonical-1301202032021030-2002301230332332-2232103012011312-1320101033013312-0322110310222301-0201101203003102-0321320101002030-0320120211311212"></a>

### Direct properties for `jwt_validation.action`

- [block](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2301132022103110-0023313223111130-3303221301022222-2022300330312010-3203123333032313-1123002111100302-2330002110200131-3333313001312313): complete subsection reference.

- [report](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2012102131330300-0311233221131010-1232003120132232-1020222133321100-1130303102222221-0232230113112311-2021211301033103-3020022123303032): complete subsection reference.

<a id="canonical-2301132022103110-0023313223111130-3303221301022222-2022300330312010-3203123333032313-1123002111100302-2330002110200131-3333313001312313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.action](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2121132033013110-0203011222103223-3211020302222303-0223111212332002-1211112230203201-2000211102021000-2110302303321020-0031111212212330)
- jwt_validation.action.block

<a id="canonical-0213123211321130-0130020101320123-2030221323200313-3221110111000300-2310312210231332-1323101110013121-1300002031000313-2120101002121300"></a>

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

<a id="canonical-2012102131330300-0311233221131010-1232003120132232-1020222133321100-1130303102222221-0232230113112311-2021211301033103-3020022123303032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.action](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2121132033013110-0203011222103223-3211020302222303-0223111212332002-1211112230203201-2000211102021000-2110302303321020-0031111212212330)
- jwt_validation.action.report

<a id="canonical-2331113300213332-2132230023000003-3101111230031323-1000120130013201-3331300133100132-0011031122232310-2032101031330301-2203131210131002"></a>

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

<a id="canonical-0310100010103312-1233220331111211-3302232321222020-1110011322003010-2103021012033230-3033203321133230-3110202223321102-1031201001311221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- jwt_validation.authorization_server

<a id="canonical-2332301102123131-1021320021232010-1233320011101212-3033122233121003-3211331111232101-2200221003322100-0220103131102012-2013013112221010"></a>

Type: `"single"`. Computed.

Reference to Authorization Server object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3002101222103212-0332303030300333-3200101012100113-1221220231123102-2022322212130312-1000122023002022-1032133022201222-3320210132300322"></a>

### Direct properties for `jwt_validation.authorization_server`

- [authorization_servers](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2122211231213200-1001330323130103-1330330000300312-2012031011313132-3302120023223201-1313210133131230-3303132222033112-3103213100232300): complete subsection reference.

<a id="canonical-2122211231213200-1001330323130103-1330330000300312-2012031011313132-3302120023223201-1313210133131230-3303132222033112-3103213100232300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server.authorization_servers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.authorization_server](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0310100010103312-1233220331111211-3302232321222020-1110011322003010-2103021012033230-3033203321133230-3110202223321102-1031201001311221)
- jwt_validation.authorization_server.authorization_servers

<a id="canonical-1102301231121332-3221312300122122-1130110113332133-0331202201132100-2013323003012013-2131103203230020-3001112110111333-0003332012302212"></a>

Type: `"list"`. Computed.

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-3221223102322111-0332313111233211-1101101001132033-3330100321303123-0221220333103230-3322311121320013-0200210101012000-3223230233120102"></a>

### Direct properties for `jwt_validation.authorization_server.authorization_servers`

<a id="canonical-2222002331031210-2012011031011020-0303210231222010-3103322220133001-1120301120023221-0131230001321002-2212032103322011-3332331222300331"></a>

#### `jwt_validation.authorization_server.authorization_servers.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3323031223300111-2311023020312003-3200220103110131-3303330232221132-2101211112121031-0203210010010331-3100133033332333-2111211213323122"></a>

<a id="canonical-0013303311300223-0223330313302200-2001122320110103-1002103333233031-1230231313102230-1233110111331130-1223021033332003-1332301132011230"></a>

#### `jwt_validation.authorization_server.authorization_servers.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0002213131300022-3322211101033302-0033100001221230-0010312332201110-2032110121110203-0200033230003023-3031322021020330-2310231101232002"></a>

<a id="canonical-3213133200103030-0102320101021311-3210310311221322-3333013222332202-2113003122111013-3222211122032022-0023330132223201-0232212003313221"></a>

#### `jwt_validation.authorization_server.authorization_servers.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2233311333221001-0201023333322021-0001121233331231-2203102320310202-0203310201232311-0033100133232033-2311132220330122-3013022213113102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.jwks_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- jwt_validation.jwks_config

<a id="canonical-0301222032033311-2130022303003032-1133200233300213-1311320001233313-1111020022110301-0330111211102300-3302233122221032-0201230230322001"></a>

Type: `"single"`. Computed.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3010212110302032-0201301103221112-2332330131200003-3031022332132112-3112031112310100-2311312103232033-0222300131121203-3330023100232212"></a>

### Direct properties for `jwt_validation.jwks_config`

<a id="canonical-3022312030321201-3200222120302323-3201322132302221-2331022311112022-0020130231200021-0100312301003230-0233232101200122-1301312133131122"></a>

#### `jwt_validation.jwks_config.cleartext` property

Type: `"string"`. Computed.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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

<a id="canonical-0013131103300111-3033111303113002-3012310103320120-3013323323201023-2222300130320021-1000122302322132-3333112211031230-1332333103310212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.mandatory_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- jwt_validation.mandatory_claims

<a id="canonical-0022100113020001-3121002310200223-3210222233120121-3222331210331221-1202113121012111-1330323223212110-2323310311010020-0020102113320103"></a>

Type: `"single"`. Computed.

Configurable Validation of mandatory Claims.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2102330120032011-0222101230203223-1032331323011210-1011232301101033-1201202012121213-2303103122133000-0112001101303012-2231013201213031"></a>

### Direct properties for `jwt_validation.mandatory_claims`

<a id="canonical-3113012231333322-2002333013231001-1123220030022122-3032031221212120-2221233200003222-0133021010333012-0121232010312310-2102010210322111"></a>

#### `jwt_validation.mandatory_claims.claim_names` property

Type: `["list", "string"]`. Computed.

Claim Names. Human-readable name for the resource

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

<a id="canonical-2000130323022201-1203333322333203-2223333303103023-2132200101132302-3033023313001031-2201000201221213-1102001212002322-2130102102212330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- jwt_validation.reserved_claims

<a id="canonical-1123330220330130-2123221113202103-2011002313013030-1031113303222023-1213123020100231-3313130221203033-0101122303100222-3302000221021313"></a>

Type: `"single"`. Computed.

Configurable Validation of reserved Claims.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-audience_validation": "[\"audience\",\"audience_disable\"]",
  "x-ves-oneof-field-issuer_validation": "[\"issuer\",\"issuer_disable\"]",
  "x-ves-oneof-field-validate_period": "[\"validate_period_disable\",\"validate_period_enable\"]"
}
```

<a id="canonical-1201100113122302-2231323331102132-1110021203210323-3123131130230103-1022102203333133-1211032200132320-0001123021233212-0200233113231312"></a>

### Direct properties for `jwt_validation.reserved_claims`

- [audience](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3202200320023003-0101010202132232-1233322023021013-0130002313222102-3032000203230110-2301232133002122-3031302201023200-0202222202331201): complete subsection reference.

- [audience_disable](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3020312033121000-2230001310322231-2330030201332023-3313100302320102-2303221002213323-1121233311212231-1001303320000121-2203031101233220): complete subsection reference.

<a id="canonical-3223020200021123-3210212031112022-2203000033012022-0023110101032213-0003231111132121-3022013121012301-2301023120031201-2002012132232031"></a>

<a id="canonical-2320320001002303-3000320301110133-2211301203102101-3300200302220213-0332023223332303-1112102030033332-0202311223221132-2131223122120323"></a>

#### `jwt_validation.reserved_claims.issuer` property

Type: `"string"`. Computed.

Exact Match. Exclusive with \[issuer\_disable\]

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

- [issuer_disable](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0003020220300223-0110122303133010-2202002103111321-1131223123332131-0002033031232201-2021031333012223-3033121203231320-2303023033303021): complete subsection reference.

- [validate_period_disable](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1220311131022133-0000112333312030-0213131313032020-0212110000310331-0013120030301233-1011021221021202-3011032303313222-1021212332120230): complete subsection reference.

- [validate_period_enable](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3313303032211312-1110230111100303-2110100033123232-0332321222310220-3110121213001313-3210011201110000-1201301211311102-0222110330103212): complete subsection reference.

<a id="canonical-3202200320023003-0101010202132232-1233322023021013-0130002313222102-3032000203230110-2301232133002122-3031302201023200-0202222202331201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.audience` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2000130323022201-1203333322333203-2223333303103023-2132200101132302-3033023313001031-2201000201221213-1102001212002322-2130102102212330)
- jwt_validation.reserved_claims.audience

<a id="canonical-1200212122120112-0100313031021101-2332210210022232-1220212033133102-2220330203223231-0233022212021122-1023022301230033-0011131232203101"></a>

Type: `"single"`. Computed.

Audiences

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2121001010033121-2320202223113313-1103230231121001-3003301112100212-2113131131210200-2332303123332123-2011023232110330-0023031112330033"></a>

### Direct properties for `jwt_validation.reserved_claims.audience`

<a id="canonical-0101020031123103-1330001332231221-2012203101301301-0132302210203323-0111101132120220-2200223303311100-0011312103301330-0331301121123111"></a>

#### `jwt_validation.reserved_claims.audience.audiences` property

Type: `["list", "string"]`. Computed.

Values. Configuration parameter for audiences

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

<a id="canonical-3020312033121000-2230001310322231-2330030201332023-3313100302320102-2303221002213323-1121233311212231-1001303320000121-2203031101233220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.audience_disable` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2000130323022201-1203333322333203-2223333303103023-2132200101132302-3033023313001031-2201000201221213-1102001212002322-2130102102212330)
- jwt_validation.reserved_claims.audience_disable

<a id="canonical-3303322232230123-2032120320321220-2032013111120320-2011230113012130-1111300323323202-2030131320122110-2100110200130121-1323312312310303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for audience disable.

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

<a id="canonical-0003020220300223-0110122303133010-2202002103111321-1131223123332131-0002033031232201-2021031333012223-3033121203231320-2303023033303021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.issuer_disable` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2000130323022201-1203333322333203-2223333303103023-2132200101132302-3033023313001031-2201000201221213-1102001212002322-2130102102212330)
- jwt_validation.reserved_claims.issuer_disable

<a id="canonical-3120232210320313-1001303202012333-3102221001100033-1202233220301330-0203103310202101-3112032233032133-3131030020302203-3123013331033201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for issuer disable.

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

<a id="canonical-1220311131022133-0000112333312030-0213131313032020-0212110000310331-0013120030301233-1011021221021202-3011032303313222-1021212332120230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.validate_period_disable` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2000130323022201-1203333322333203-2223333303103023-2132200101132302-3033023313001031-2201000201221213-1102001212002322-2130102102212330)
- jwt_validation.reserved_claims.validate_period_disable

<a id="canonical-2121200020131021-3202101032301103-2110100323200130-0031233010101230-3212212002010123-3013111223323012-0312101313123210-0100330010233213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for validate period disable.

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

<a id="canonical-3313303032211312-1110230111100303-2110100033123232-0332321222310220-3110121213001313-3210011201110000-1201301211311102-0222110330103212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.validate_period_enable` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2000130323022201-1203333322333203-2223333303103023-2132200101132302-3033023313001031-2201000201221213-1102001212002322-2130102102212330)
- jwt_validation.reserved_claims.validate_period_enable

<a id="canonical-0120203020120102-1111130033010030-2223211130222023-0131133122130333-1333023101230313-3332210232220232-1010001310300001-2022300212122031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for validate period enable.

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

<a id="canonical-2003033103100212-0221322000032122-1033301212022110-3311221200010103-3112220320320320-1031002021233000-3210122333012323-1113020133332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- jwt_validation.target

<a id="canonical-3113310232123102-2332222320120201-0131220203322011-1313200120311022-2122211213033233-2233322203222203-1230120102310103-0320203333223232"></a>

Type: `"single"`. Computed.

Define endpoints for which JWT token validation will be performed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target": "[\"all_endpoint\",\"api_groups\",\"base_paths\"]"
}
```

<a id="canonical-2023222330203222-2033212023030212-3101023103302003-0202313230130120-2312323321012303-3321323310313130-2112112121331322-0100011233110321"></a>

### Direct properties for `jwt_validation.target`

- [all_endpoint](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0312021102331203-3020203022033221-1331033201020023-2220012011002100-0210020222112212-0011021013000220-3000322201111100-2303231131222111): complete subsection reference.

- [api_groups](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2233121203122313-0223111032212010-3220210331311220-0332032011011020-1001012330311231-2022322020032112-3012210002021122-1332200333300020): complete subsection reference.

- [base_paths](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2231120101112320-2130230111301220-1112012300000110-1013330013333101-1123322331000012-1220221202200200-1032031112331100-3111001103000011): complete subsection reference.

<a id="canonical-0312021102331203-3020203022033221-1331033201020023-2220012011002100-0210020222112212-0011021013000220-3000322201111100-2303231131222111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.all_endpoint` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2003033103100212-0221322000032122-1033301212022110-3311221200010103-3112220320320320-1031002021233000-3210122333012323-1113020133332202)
- jwt_validation.target.all_endpoint

<a id="canonical-0013101011010210-1000223220213211-0021310303223213-0312202200000021-3310333230233302-3130332220321230-2032132113120013-0313132002330021"></a>

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

<a id="canonical-2233121203122313-0223111032212010-3220210331311220-0332032011011020-1001012330311231-2022322020032112-3012210002021122-1332200333300020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.api_groups` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2003033103100212-0221322000032122-1033301212022110-3311221200010103-3112220320320320-1031002021233000-3210122333012323-1113020133332202)
- jwt_validation.target.api_groups

<a id="canonical-0103233320321322-1203000231021210-0222303322123123-2303302020010201-1010020223221022-3001022220031130-0230031203201011-2233020002022300"></a>

Type: `"single"`. Computed.

API Groups.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3231331002023031-0213213230202003-2302121022230231-0311211020133021-0320121120213200-3301300102011323-0233331330022220-1120033310000300"></a>

### Direct properties for `jwt_validation.target.api_groups`

<a id="canonical-3231102233131032-1202301311222203-2311111133222233-3330120232303022-2022312021311331-3103202112320001-0033113100302021-0321210322302302"></a>

#### `jwt_validation.target.api_groups.api_groups` property

Type: `["list", "string"]`. Computed.

API Groups. Group or collection configuration

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2231120101112320-2130230111301220-1112012300000110-1013330013333101-1123322331000012-1220221202200200-1032031112331100-3111001103000011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.base_paths` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2003033103100212-0221322000032122-1033301212022110-3311221200010103-3112220320320320-1031002021233000-3210122333012323-1113020133332202)
- jwt_validation.target.base_paths

<a id="canonical-3201221130101132-1222101113211311-2323113320133211-3033201313030011-1132102032020210-0222033300111123-1311303010000321-0122030233222120"></a>

Type: `"single"`. Computed.

Base Paths.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3230321310302031-2030013220202312-0111012202010201-0301002132022020-3022202133101121-0032310231313201-3222002122312332-1121020333031303"></a>

### Direct properties for `jwt_validation.target.base_paths`

<a id="canonical-3130100200301011-3213100233230333-2133230022103112-2223102202230123-1021330012010011-3200120232122022-2030330012212001-3020002312322313"></a>

#### `jwt_validation.target.base_paths.base_paths` property

Type: `["list", "string"]`. Computed.

Prefix Values. File system or URL path

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1010220232131121-3133122220132302-2102331001310120-1103121302102313-0210031123201300-3211132222021113-1033033022012000-0210101200200032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.token_location` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- jwt_validation.token_location

<a id="canonical-1011112121210202-2033232331030021-1331311323303210-3130122001301310-0322123000231201-0132031230333122-0331231101212321-0333303103132130"></a>

Type: `"single"`. Computed.

Configuration parameter for token location.

Additional upstream details:

Location of JWT in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-token_location": "[\"bearer_token\"]"
}
```

<a id="canonical-0333332110210233-1231221133122120-3201211123311323-3021120012213230-2003212003300013-3212003012302333-2313100003302233-2120111222030210"></a>

### Direct properties for `jwt_validation.token_location`

- [bearer_token](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2323122310133033-0011322123033011-3232331222021000-3001331210002230-0220203131212011-2103223313021332-0212312101112222-3201220030011332): complete subsection reference.

<a id="canonical-2323122310133033-0011322123033011-3232331222021000-3001331210002230-0220203131212011-2103223313021332-0212312101112222-3201220030011332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.token_location.bearer_token` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310)
- [jwt_validation.token_location](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1010220232131121-3133122220132302-2102331001310120-1103121302102313-0210031123201300-3211132222021113-1033033022012000-0210101200200032)
- jwt_validation.token_location.bearer_token

<a id="canonical-2212221303222213-1221101121223113-3111022300323232-0231302203231103-3223111000132230-0000123030020010-0131232303302130-2120212201203232"></a>

Type: `"single"`. Computed.

Configuration parameter for bearer token.

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

<a id="canonical-1333333200120332-2013203303220313-1010312331020333-1200321100302113-0213120311020001-2122222320111122-1131233320112210-3120230011331311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- l7_ddos_action_block

<a id="canonical-0121332212030320-0213232321001000-3320203322101202-2300222332123020-1020133333031113-0032132211211312-2313313303221322-2310230103232103"></a>

Type: `["object", {}]`. Computed.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

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

- [l7_ddos_action_block](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0121332212030320-0213232321001000-3320203322101202-2300222332123020-1020133333031113-0032132211211312-2313313303221322-2310230103232103)
- [l7_ddos_action_default](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2320123303212320-1212123302322032-0323012003003120-0221123121231123-1030301103100322-2232012311013211-3223320103133231-0200303000112022)
- [l7_ddos_action_js_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1002031021002332-2223003333233333-3011103103202321-1233232033012322-3212032102022130-0133212311320220-1220011130213313-2002023032322211)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200300131023022-3211332031133233-1011202032330202-3130010032100021-3311130322311121-3313332200131021-3212002210303013-2323210011000003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_default` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- l7_ddos_action_default

<a id="canonical-2320123303212320-1212123302322032-0323012003003120-0221123121231123-1030301103100322-2232012311013211-3223320103133231-0200303000112022"></a>

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

<a id="canonical-1000121000031123-3101310311110223-2321303023200233-3323223212003123-0212021030302130-1321311333202303-0011001201301122-3311213103331020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_js_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- l7_ddos_action_js_challenge

<a id="canonical-1002031021002332-2223003333233333-3011103103202321-1233232033012322-3212032102022130-0133212311320220-1220011130213313-2002023032322211"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1323002213123321-1313130010311222-1221321221321102-1131221101320110-1233112123002312-3321302212013220-3333222313322132-2131333011023211"></a>

### Direct properties for `l7_ddos_action_js_challenge`

<a id="canonical-2102302110323031-0032021231100010-3011333012133231-0022323220130100-1230000122211113-3213031010203321-1122111330322002-3312113213121320"></a>

#### `l7_ddos_action_js_challenge.cookie_expiry` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0000203022231302-1011000311003121-3110121112330303-1332331101101112-2030021220003102-2313131110323033-3123133133111303-1110333030133210"></a>

<a id="canonical-3232211122100100-1323321133130002-3310123301010020-3133333312200130-2100330103301110-3022111021022002-0300222030323312-3011011202300323"></a>

#### `l7_ddos_action_js_challenge.custom_page` property

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1110131211312011-1202101231333320-2220211013131230-2220011011300002-3233113330032002-0110002011031120-1012300211212013-0211313213210233"></a>

<a id="canonical-0333011021232110-2133011022200231-0113301330202130-1332122201122022-0033112113112212-2021230320111013-0221112310301300-1021012023303110"></a>

#### `l7_ddos_action_js_challenge.js_script_delay` property

Type: `"number"`. Computed.

Delay introduced by JavaScript, in milliseconds.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1232110221313322-1131211111332023-1003031130302012-1213030301210330-0103133011322131-1112113122330123-0213220333222100-0202331332301210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- no_challenge

<a id="canonical-3310301110320021-0213111113212022-3132132101000131-3200220321201123-1332220102000021-3003110110011011-3300301312033311-3223313332033121"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-0100133220321213-2123201303210012-3220032001222313-3010303023322211-2133122331211121-0312122103223102-2233330202130123-2213221132032212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_service_policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- no_service_policies

<a id="canonical-2123133221100300-1332201322123031-0223020033302100-0012203203012302-3210301020331023-1032022110113221-1110220203031213-2123232211032300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no service policies.

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

<a id="canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- origin_pool

<a id="canonical-1002210021120022-1322322031020020-3120120233320133-0023022123313122-2113100222302300-0211323021212022-0032310332132230-2132123132203031"></a>

Type: `"single"`. Computed.

Configuration parameter for origin pool.

Additional upstream details:

Origin Pool for the CDN distribution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-1210212120103020-0331231033323011-3021223221201202-2310102200000222-1103211330021231-2130210211123023-1230123200110023-3222131003212312"></a>

### Direct properties for `origin_pool`

- [more_origin_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2100100112333133-0231133201332010-3330303123020131-2331222323022102-0213320032321313-1122011332212031-0323101302201211-0202110320102031): complete subsection reference.

- [no_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2031230312010033-2310122001111003-3300013203022310-1301312203232323-2121123131232111-0012210011021120-1300031223020201-1020132120200030): complete subsection reference.

<a id="canonical-1211310111003113-0102233210002031-3023120200103231-1012200213333312-0130023013200220-2003210322323300-3111031101020023-0100133221203200"></a>

<a id="canonical-0330211030112302-3033233101112301-3012311300001101-0312323020232212-2303012112221220-3231032101012330-2120103222110101-1001112323023332"></a>

#### `origin_pool.origin_request_timeout` property

Type: `"string"`. Computed.

Configures the time after which a request to the origin will time out waiting for a response.

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
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [origin_servers](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1011000010211201-1230313033201303-0203320122033323-0202110232110203-0003303302302131-3023122211323223-2231102320110130-3001132020010021): complete subsection reference.

- [public_name](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2301023123001322-2203122230111300-2322323332212113-1031310002121311-2133032320112203-2113122012330212-2313320111310111-1013330000001103): complete subsection reference.

- [use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100): complete subsection reference.

<a id="canonical-2100100112333133-0231133201332010-3330303123020131-2331222323022102-0213320032321313-1122011332212031-0323101302201211-0202110320102031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.more_origin_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- origin_pool.more_origin_options

<a id="canonical-1201031303313202-3031101303220202-3130030213011332-1300231001101330-3132211333302121-2102022322302322-0101211311213003-1012101320032013"></a>

Type: `"single"`. Computed.

Configuration parameter for more origin options.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2330111311322322-0310310300322000-0310110111211220-0313032303121220-3021110031122203-2121031121311321-0212211303331333-3113333322020100"></a>

### Direct properties for `origin_pool.more_origin_options`

<a id="canonical-3302103121021023-0323223023032030-1201031030213332-2001131302222322-0110232323233101-1230121331101231-0030303320310020-3110111232121313"></a>

#### `origin_pool.more_origin_options.enable_byte_range_request` property

Type: `"bool"`. Computed.

Choice to enable/disable byte range requests towards origin.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0230203211230020-0122303131100113-2221133133003212-3030012211023120-0210012032231000-2021110230123202-2013233321311231-0031012013020202"></a>

<a id="canonical-2012323202022212-0012302321021230-3210003230333123-0230232012102022-0232023313201002-3133233031312213-2302000310203101-1320302211300220"></a>

#### `origin_pool.more_origin_options.websocket_proxy` property

Type: `"bool"`. Computed.

Option to enable proxying of websocket connections to the origin server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2031230312010033-2310122001111003-3300013203022310-1301312203232323-2121123131232111-0012210011021120-1300031223020201-1020132120200030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.no_tls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- origin_pool.no_tls

<a id="canonical-0103230311210230-2213200230320123-2331023322211131-2011312121022322-2301030210202320-1012232000222100-2310102210021211-2010111301300031"></a>

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

<a id="canonical-1011000010211201-1230313033201303-0203320122033323-0202110232110203-0003303302302131-3023122211323223-2231102320110130-3001132020010021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.origin_servers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- origin_pool.origin_servers

<a id="canonical-3300202311123210-1223023123222222-1321000103011012-1313103320032011-0323003103321332-2310210332121202-2200122212123212-0333021113331223"></a>

Type: `"list"`. Computed.

List Of Origin Servers. List of original servers.

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

<a id="canonical-0121230200021322-3233132020021220-1130210313230220-3303211130010302-1010130320000001-1100130113112303-1320333201302313-0221123230311301"></a>

### Direct properties for `origin_pool.origin_servers`

<a id="canonical-0213313131221112-3102022022002311-1231213003102212-2211013330101110-3321202021302332-2020211032001133-1101201022033200-1100132333203313"></a>

#### `origin_pool.origin_servers.port` property

Type: `"number"`. Computed.

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [public_ip](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0132110013211010-2232021112020113-3313310211323030-1311303110123032-3300001100222120-3023112111022001-2123301122112220-2110220312233012): complete subsection reference.

- [public_name](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0202330321113332-2001002131310211-1200123332212303-2123222030011100-2030222321210331-0032313101101201-3010333101201132-2202101113102033): complete subsection reference.

<a id="canonical-0132110013211010-2232021112020113-3313310211323030-1311303110123032-3300001100222120-3023112111022001-2123301122112220-2110220312233012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.origin_servers](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1011000010211201-1230313033201303-0203320122033323-0202110232110203-0003303302302131-3023122211323223-2231102320110130-3001132020010021)
- origin_pool.origin_servers.public_ip

<a id="canonical-1200300021310231-2011311322321101-2231302232311211-2301220232220302-0102213323212212-3320321201233021-2012032231211222-0213033320122013"></a>

Type: `"single"`. Computed.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-3222212020100113-0203123230121213-1020221133332003-0010213223000103-1212321110002321-2220313201231110-2311100212201112-2121101100013333"></a>

### Direct properties for `origin_pool.origin_servers.public_ip`

<a id="canonical-3020030133010312-3313231023323231-1301303221232300-2113113121103001-0101321333002332-0020303020101211-0200010233210312-1113000333333030"></a>

#### `origin_pool.origin_servers.public_ip.ip` property

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0202330321113332-2001002131310211-1200123332212303-2123222030011100-2030222321210331-0032313101101201-3010333101201132-2202101113102033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.origin_servers](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1011000010211201-1230313033201303-0203320122033323-0202110232110203-0003303302302131-3023122211323223-2231102320110130-3001132020010021)
- origin_pool.origin_servers.public_name

<a id="canonical-3012320031330301-2011030221111213-0220013202223222-3310310111003331-2102030231231302-0133022223111332-2011122023233301-0322031132223322"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0332123102020002-2232030203000231-0333300333322322-1121311032000010-2133123020113122-0212112301023022-1211121203022121-3303232202101311"></a>

### Direct properties for `origin_pool.origin_servers.public_name`

<a id="canonical-0120323112331102-2332011210213201-3021003212203000-3032131131213322-2133222220323023-3112303030001130-1311211233113101-3333122123300312"></a>

#### `origin_pool.origin_servers.public_name.dns_name` property

Type: `"string"`. Computed.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2222022032032322-2213331100323322-2311001230112113-2310122133213113-0010023033130210-0123000320012201-2211001013333220-1322110100213233"></a>

<a id="canonical-2020033111013203-2223122111022112-1100212110313200-0100001213331331-3010230322321033-1300002203302320-1113310300321013-2010330002033321"></a>

#### `origin_pool.origin_servers.public_name.refresh_interval` property

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-2301023123001322-2203122230111300-2322323332212113-1031310002121311-2133032320112203-2113122012330212-2313320111310111-1013330000001103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.public_name` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- origin_pool.public_name

<a id="canonical-2121200323203330-3121023001232023-1210101100133020-1113301333110031-1330123120102102-1110100202001210-0221130213332121-1323022200130131"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0210202010002211-2020003013103222-1222100303232321-3100300322312121-2000301001232323-0202031001032311-1323123200222203-0233003102111232"></a>

### Direct properties for `origin_pool.public_name`

<a id="canonical-2312020021330231-3011320032102321-0331323331121312-0232022111030212-1021220303100012-0230010130021110-2122113130300213-3032310220133223"></a>

#### `origin_pool.public_name.dns_name` property

Type: `"string"`. Computed.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1102102003310233-2322230001012210-0130202302322320-1221231221030213-1331303231313012-1100032002121123-0320021310021112-3033130323130232"></a>

<a id="canonical-3320111000020003-2023133330110132-1100001100033110-2010211031213212-1133212330110103-2301300020032003-0220200000120020-2331320211233322"></a>

#### `origin_pool.public_name.refresh_interval` property

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- origin_pool.use_tls

<a id="canonical-0003200131302013-1200020211102123-1330312010211333-3101111130312011-3110123320101313-3001320131000312-3131113312323322-2203233312010210"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

<a id="canonical-3200021312113010-0020310121021003-2322031202202212-0113332022000313-0100022033231112-1132223233220333-0022022323210322-2212301333233112"></a>

### Direct properties for `origin_pool.use_tls`

- [default_session_key_caching](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0310120301010101-3131230121130233-1322302023102302-2002122310023320-2322302120001113-3203332201120201-2112233200130133-0023110312323020): complete subsection reference.

- [disable_session_key_caching](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2201331101003332-3031310321031033-1232133131310233-3020001012000112-0322223322031313-2022300331033122-3020203103122311-3213332311202002): complete subsection reference.

- [disable_sni](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1112232301011020-1301220131310321-3102321333232231-3320132321310223-3221222102332123-0333301222120123-0233200031031131-2303321013123113): complete subsection reference.

<a id="canonical-3220221323303223-2031333130023333-2213322232031032-1330100232230303-3012310213200320-2101123010233031-0102102210001302-0200110000332203"></a>

<a id="canonical-1132321212001001-1302110333303211-2201102330122230-2321322233132222-1103200221223021-1031120213112333-0010111333301031-3223310313213203"></a>

#### `origin_pool.use_tls.max_session_keys` property

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2020033002332013-1021323202131031-3110222302033120-1223333102121102-3323100321023302-2323112131221122-1033133332120012-3220010021201002): complete subsection reference.

- [skip_server_verification](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2113112032120133-1323010012112000-0202223231301013-3012221213132133-2332101033210031-1002030011123020-2132222010230020-2033311020131202): complete subsection reference.

<a id="canonical-1320312313113303-2011010031013223-0233221200333300-0311233110001221-2001012203022002-1100113221102110-3103313300202131-1230113003021201"></a>

<a id="canonical-0010300321100233-0320121323302323-1003132123132031-2012100333302013-2303332222001310-3312323330301311-0323213231312101-0022002310030020"></a>

#### `origin_pool.use_tls.sni` property

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2022333133133320-3033011011031311-3331212130302203-3121321002020011-1212032102233031-0201131111301100-1332222302313302-2202121331033310): complete subsection reference.

- [use_host_header_as_sni](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3131121130111212-2311123222211323-0331002210200320-2003032303221233-0000213012232130-1132330310310022-3201321302021203-1100331303020233): complete subsection reference.

- [use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103): complete subsection reference.

- [use_mtls_obj](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3232212022121320-1101213210330122-2130312210301222-2203232201212312-3302313120032112-3002112203020230-2133233230003033-2011101021023313): complete subsection reference.

- [use_server_verification](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3112310332001121-3003121211331210-0323212121333133-0320010213211123-2202202111320200-2331112301223210-3222130301021213-3331311112012010): complete subsection reference.

- [volterra_trusted_ca](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1321313303313232-3311030333223032-3001313302112102-3101102302200031-3102232021201232-2030300102023020-1121201330102011-3220200000303332): complete subsection reference.

<a id="canonical-0310120301010101-3131230121130233-1322302023102302-2002122310023320-2322302120001113-3203332201120201-2112233200130133-0023110312323020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.default_session_key_caching

<a id="canonical-1220322321200121-2012131033312202-3123222130311212-3211100313313130-2132110130212300-0002321103301121-1011131023332111-3130003310302122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-2201331101003332-3031310321031033-1232133131310233-3020001012000112-0322223322031313-2022300331033122-3020203103122311-3213332311202002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.disable_session_key_caching

<a id="canonical-0201033002303332-0203303111102201-0331210320200010-0313032230311133-2022210020011111-2303132233310232-2312221310113232-3221211011212031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

<a id="canonical-1112232301011020-1301220131310321-3102321333232231-3320132321310223-3221222102332123-0333301222120123-0233200031031131-2303321013123113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.disable_sni` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.disable_sni

<a id="canonical-0233322330131233-2212022020220203-3223233221201123-1033303310331200-1232033122201031-1232023203111112-3001002201102022-0303120031233032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

<a id="canonical-2020033002332013-1021323202131031-3110222302033120-1223333102121102-3323100321023302-2323112131221122-1033133332120012-3220010021201002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.no_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.no_mtls

<a id="canonical-1311231130323201-2303133230131011-3030022031321013-0030312230010331-0321311101020011-1030132120322120-1300330300301231-0011321331020023"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-2113112032120133-1323010012112000-0202223231301013-3012221213132133-2332101033210031-1002030011123020-2132222010230020-2033311020131202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.skip_server_verification` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.skip_server_verification

<a id="canonical-3200310112221312-1331331332310023-0033011131123110-0002333130230011-3123210001031030-2201023001323210-0203212032102001-0223103130201301"></a>

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

<a id="canonical-2022333133133320-3033011011031311-3331212130302203-3121321002020011-1212032102233031-0201131111301100-1332222302313302-2202121331033310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.tls_config

<a id="canonical-3322221321222102-2131123311002313-3211113311330330-0113220202020011-0301223011112002-0012201133113132-0213332023201020-0221221300211010"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-3010313330232233-0210313330330031-3011122232113210-3002320313332323-3202201221320111-3301230200132013-2310210221101012-1231011022221121"></a>

### Direct properties for `origin_pool.use_tls.tls_config`

- [custom_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3333202030220013-0331022323012312-2012001033223201-1123103033133002-0022232311002021-3333123030311030-3020023110100301-3002210012031333): complete subsection reference.

- [default_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1102033120223300-1202312202212211-3112221302331303-0202120131303001-0312310133313012-3332221002121330-2320223012030323-2121030032233201): complete subsection reference.

- [low_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2032231321331222-3002321322033210-3301303020220023-0323321221320020-1023330122302011-2222102300022031-2010310022202302-2022123213012011): complete subsection reference.

- [medium_security](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0233022303221102-3232323001322300-1303200300200301-1210311320222330-2111131101202203-3013313231230113-1111123223300331-3032213332101222): complete subsection reference.

<a id="canonical-3333202030220013-0331022323012312-2012001033223201-1123103033133002-0022232311002021-3333123030311030-3020023110100301-3002210012031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2022333133133320-3033011011031311-3331212130302203-3121321002020011-1212032102233031-0201131111301100-1332222302313302-2202121331033310)
- origin_pool.use_tls.tls_config.custom_security

<a id="canonical-2032302233113011-3123232332221320-0222113102123002-3010111210333013-2112310320113131-0333201122211132-2102310000332223-3311131321323111"></a>

Type: `"single"`. Computed.

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0132102221303212-3200323010030320-2003302202133203-0201110311222322-1102000332122323-1213223330112121-3110213123321121-0312301312130332"></a>

### Direct properties for `origin_pool.use_tls.tls_config.custom_security`

<a id="canonical-1000011121113032-3020221331202102-0133210003220200-0120020313321310-3003302103020132-3113013001110122-3331103232000321-2333232012320100"></a>

#### `origin_pool.use_tls.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103001210110330-2100112220010331-2333103011113323-1110302112230331-0122101331211133-0122230202332213-3011122212123001-2013001120302223"></a>

<a id="canonical-0122120312330032-0330333320122020-1323320100023132-2102131312031231-0321020130122312-1301213302302232-2220220213122210-0002233233101020"></a>

#### `origin_pool.use_tls.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2230303222030132-0232012233332101-0110323021013101-2212102030213320-2321100110013301-2012000001100111-0301332331213213-0223122100123232"></a>

<a id="canonical-2321211330110210-1003032220330112-1110021003211013-2223112223303102-2310011023030120-1123132320023101-2011102022311101-2200133022130220"></a>

#### `origin_pool.use_tls.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1102033120223300-1202312202212211-3112221302331303-0202120131303001-0312310133313012-3332221002121330-2320223012030323-2121030032233201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2022333133133320-3033011011031311-3331212130302203-3121321002020011-1212032102233031-0201131111301100-1332222302313302-2202121331033310)
- origin_pool.use_tls.tls_config.default_security

<a id="canonical-3303212233221333-0233200123021200-2230012332303131-3122100232100031-1322103132120212-2300123103031002-2201101312201221-0330103113012313"></a>

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

<a id="canonical-2032231321331222-3002321322033210-3301303020220023-0323321221320020-1023330122302011-2222102300022031-2010310022202302-2022123213012011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2022333133133320-3033011011031311-3331212130302203-3121321002020011-1212032102233031-0201131111301100-1332222302313302-2202121331033310)
- origin_pool.use_tls.tls_config.low_security

<a id="canonical-2303113301000013-1303000100020202-2203230031332302-0311112011021313-0222103003313131-0120100212130230-2010223131032113-2220010121133312"></a>

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

<a id="canonical-0233022303221102-3232323001322300-1303200300200301-1210311320222330-2111131101202203-3013313231230113-1111123223300331-3032213332101222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2022333133133320-3033011011031311-3331212130302203-3121321002020011-1212032102233031-0201131111301100-1332222302313302-2202121331033310)
- origin_pool.use_tls.tls_config.medium_security

<a id="canonical-0301302003102331-0321211223101013-2212303001022011-2001313201302330-3303121222201113-0120301201022232-3222200103312031-3301101032110110"></a>

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

<a id="canonical-3131121130111212-2311123222211323-0331002210200320-2003032303221233-0000213012232130-1132330310310022-3201321302021203-1100331303020233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.use_host_header_as_sni

<a id="canonical-1102023333130130-0302001203011203-2132223322111103-2200321002320020-2302122223313233-0322201303221220-2032213010331002-2102130323023212"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.use_mtls

<a id="canonical-1232030331102200-3233133210011213-3123011112012213-1113111113002111-0011330120312032-2002301123033010-1122010300303001-3202022113230303"></a>

Type: `"single"`. Computed.

MTLS Certificate. MTLS Client Certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3010323012010310-3323201312123132-1032101000032223-0133320203332002-1331000221013120-1230333301011112-2331101213120202-0200223230121033"></a>

### Direct properties for `origin_pool.use_tls.use_mtls`

- [tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0332330333021122-2010031101121322-1302112000322203-0032100232310203-1023101300103322-2331200012310001-2221102021203110-2200130323011033): complete subsection reference.

<a id="canonical-0332330333021122-2010031101121322-1302112000322203-0032100232310203-1023101300103322-2331200012310001-2221102021203110-2200130323011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103)
- origin_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-3313012130023131-2313102132303321-3003303232033101-1133020313323331-2121133332100102-2012112112100233-0222232313330213-3311101313123311"></a>

Type: `"list"`. Computed.

MTLS Client Certificate. MTLS Client Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2310302321230133-2233013300330011-2023103323230300-2221220301231323-3011131332022020-1133131323122033-1200213203332321-1313321303023012"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates`

<a id="canonical-2012112020301202-2003200122203031-1203323013302023-3222220313223310-2133220113032203-3121132320330211-3123301330232332-0222111332320031"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3313323013330310-2111332313220023-0022002313323303-0330231031130031-0221100132111001-1030002000023121-0222003232123031-2132320323003132): complete subsection reference.

<a id="canonical-1211001203022223-3310310223200320-0312111223332221-3322322323000131-2123200121133321-0132233010330011-3313222131013112-3001022002123210"></a>

<a id="canonical-1220213222200131-2211123113033123-2230213100122312-1111133310333221-0103023300132321-1312122103013303-1003332013022013-0132030122112221"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3111201220013323-1021020033100332-2322121121102120-1022101112202222-2322033232333131-3123023320013233-0120302020113113-1113013102133000): complete subsection reference.

- [private_key](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1131310320310132-1322011321200001-2220003132113021-3212311312132031-2310031330033111-0333130012300332-3231311120112201-0122232033233013): complete subsection reference.

- [use_system_defaults](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111121022101333-1333001323121031-0232230231332123-1330112130221332-0001012210313330-1212332100303110-3321212131321323-3300322231113223): complete subsection reference.

<a id="canonical-3313323013330310-2111332313220023-0022002313323303-0330231031130031-0221100132111001-1030002000023121-0222003232123031-2132320323003132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0332330333021122-2010031101121322-1302112000322203-0032100232310203-1023101300103322-2331200012310001-2221102021203110-2200130323011033)
- origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-1003013011230232-3303100222101301-3230331130333022-2100223112131300-0231223213300230-2330103132102213-0312333232211223-1220223312013321"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3231233320313122-1200030200100020-1222303002313202-3211202130302032-0203120012203011-1201013103000213-0020103300212031-0113222001220201"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms`

<a id="canonical-0303100011231233-3232032022202333-0033220302113023-2320330313331101-3111330011101103-3311111321021121-3303303231333302-2232302102023011"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3111201220013323-1021020033100332-2322121121102120-1022101112202222-2322033232333131-3123023320013233-0120302020113113-1113013102133000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0332330333021122-2010031101121322-1302112000322203-0032100232310203-1023101300103322-2331200012310001-2221102021203110-2200130323011033)
- origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-1111212022300320-2201313232203012-0101103122113031-0310321202230330-0301132311020000-3030001201211221-3022110203320303-1033230300133322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-1131310320310132-1322011321200001-2220003132113021-3212311312132031-2310031330033111-0333130012300332-3231311120112201-0122232033233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0332330333021122-2010031101121322-1302112000322203-0032100232310203-1023101300103322-2331200012310001-2221102021203110-2200130323011033)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-2032310002120302-1300302302020312-1111233003312012-3000202032312202-0021112221333000-2303212333111102-2221220202030320-0021203110123111"></a>

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

<a id="canonical-0221222202212313-2333220221213320-1020213102011212-0231111012220112-2222013301030201-2321131003032222-3302321002020121-1002132312221201"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2101313221330303-2100022112100332-2310002313310010-3200313022102131-0122000030331220-2101203301030220-0212020210220231-1231122033031000): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201303223120122-1321121302231203-0220220103131131-1003230030113032-0330202231231121-2331301001113123-0233003322120001-0013230333210330): complete subsection reference.

<a id="canonical-2101313221330303-2100022112100332-2310002313310010-3200313022102131-0122000030331220-2101203301030220-0212020210220231-1231122033031000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0332330333021122-2010031101121322-1302112000322203-0032100232310203-1023101300103322-2331200012310001-2221102021203110-2200130323011033)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1131310320310132-1322011321200001-2220003132113021-3212311312132031-2310031330033111-0333130012300332-3231311120112201-0122232033233013)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2310331011132231-3213103000100323-2232232211100031-1000332330310020-2200303103320331-0003312020213031-1010020331032033-2020230103332131"></a>

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

<a id="canonical-1213130023111313-2220131120101000-2013230320111110-0223303233111231-3221022002133121-1320001121212013-2133032332212120-2322220300120121"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-2130223303001320-1322030111002130-1102110122100300-3302223033333022-2312330011330023-0113311320013330-2031130003023103-2003003133100211"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0201111103121030-0123023200113303-2102132122212011-3230313020232012-2100100033021332-2300202100123202-2320022213311130-1112100002021222"></a>

<a id="canonical-2122321023131033-1100110010222111-0113233200323302-1232112202023311-1313310120113333-1221200210032310-0202212131202230-3311023022312302"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-3233230220200331-0001300303332010-1300312333333230-2313120202003302-2331223110231301-1222033011123301-1222013211001131-0211222200230330"></a>

<a id="canonical-1321233122323020-3230103103100133-2230100121221021-2121313332221310-3001032022200232-1320031222102223-2313203210301123-0322112133032112"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-0201303223120122-1321121302231203-0220220103131131-1003230030113032-0330202231231121-2331301001113123-0233003322120001-0013230333210330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0332330333021122-2010031101121322-1302112000322203-0032100232310203-1023101300103322-2331200012310001-2221102021203110-2200130323011033)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1131310320310132-1322011321200001-2220003132113021-3212311312132031-2310031330033111-0333130012300332-3231311120112201-0122232033233013)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-1130013220100320-0132011000212312-1023133330301313-1110002023133331-3303202112301313-0032300123001030-2322112231232203-1231330220301002"></a>

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

<a id="canonical-0123320102232213-3313002030112033-2313123312201321-0131123103321221-2321032331022002-3112000203130131-1200230202302001-2301201012322213"></a>

### Direct properties for `origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info`

<a id="canonical-2123222120221222-3132101330211201-0133013333122122-1221120300113212-3102233223111321-1332022320102033-0023021311222120-2131032311211231"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3030002200321013-1110232000101230-2110212221231001-2322021313021303-0113003300301021-0113123212323132-3211233003201320-2000320011230121"></a>

<a id="canonical-2323123230100322-0122013323321210-3031221233121102-1103002113033302-2203013333303201-0211330020222211-3130020221330331-0003032330131321"></a>

#### `origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-1111121022101333-1333001323121031-0232230231332123-1330112130221332-0001012210313330-1212332100303110-3321212131321323-3300322231113223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3013013030232013-0231120122010201-1022211022021122-1211201001221230-0010233133321012-0121033211103213-1310123020012320-0220112020020103)
- [origin_pool.use_tls.use_mtls.tls_certificates](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0332330333021122-2010031101121322-1302112000322203-0032100232310203-1023101300103322-2331200012310001-2221102021203110-2200130323011033)
- origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-1201230310222132-3022121130112003-0103330223332023-3212121110310202-2220210002130131-2301011023110110-3111001313130202-3210231113303100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-3232212022121320-1101213210330122-2130312210301222-2203232201212312-3302313120032112-3002112203020230-2133233230003033-2011101021023313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_mtls_obj` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.use_mtls_obj

<a id="canonical-2212222101121121-1001300212300033-3231123211013033-0000302131022210-0300330001011133-3023002303101030-3120032230333121-1233103300321100"></a>

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

<a id="canonical-2113211320031331-1001120220132310-3313210311232130-3220333123011310-1133220200230303-2223331233103122-2332231301012331-1001200203120103"></a>

### Direct properties for `origin_pool.use_tls.use_mtls_obj`

<a id="canonical-2102323333001313-0120130230301220-0320112002020311-2000132320303010-1303302120000301-1121121110131112-2333023203312001-1033032303033233"></a>

#### `origin_pool.use_tls.use_mtls_obj.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3000300210333222-2113021010211302-0323012202212331-2210013222231213-1000301121102001-0212023222002031-0121021211202202-3201132112111003"></a>

<a id="canonical-0132220000033201-2300203220332002-0130202332313112-3210112021210002-0222003022130221-2230102323330311-1221120123120111-3220023321112333"></a>

#### `origin_pool.use_tls.use_mtls_obj.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3123202331231103-1233313201033002-3310131321232213-0200110330122233-2023322120202203-1312231223023220-3003333102021101-3132030203212202"></a>

<a id="canonical-3032313101103102-2023113200311310-1210222030113300-0333130100330310-3122222300110202-1323102221313222-2022310123032202-2020203321223033"></a>

#### `origin_pool.use_tls.use_mtls_obj.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3112310332001121-3003121211331210-0323212121333133-0320010213211123-2202202111320200-2331112301223210-3222130301021213-3331311112012010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_server_verification` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.use_server_verification

<a id="canonical-0300001301100022-0332213123032322-1302100020220221-0030201332302023-2001123020003210-3000103032000213-2032101002321330-0131230321013300"></a>

Type: `"single"`. Computed.

Configuration parameter for use server verification.

Additional upstream details:

Upstream TLS Validation Context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-2322313320120020-0012302001202033-2220231013033021-0010311131132011-0201300222203213-0220212121210313-0303113200212321-1121301313312031"></a>

### Direct properties for `origin_pool.use_tls.use_server_verification`

- [trusted_ca](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2331002121333112-3322210331123203-3303210112000232-1201112020203310-3332322132320031-1303002032212311-1102303201202212-1013333212020310): complete subsection reference.

<a id="canonical-1313233300031302-2211310020230013-1011302030020023-0333130221001003-0023302311030121-1302112221301002-2012103110301201-2233111300312320"></a>

<a id="canonical-3301012100301233-0212221200132132-2323230210300122-2232010100101033-2233002321112310-3202123321013333-0331201123330222-3011331131021323"></a>

#### `origin_pool.use_tls.use_server_verification.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2331002121333112-3322210331123203-3303210112000232-1201112020203310-3332322132320031-1303002032212311-1102303201202212-1013333212020310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.use_server_verification.trusted_ca` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- [origin_pool.use_tls.use_server_verification](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3112310332001121-3003121211331210-0323212121333133-0320010213211123-2202202111320200-2331112301223210-3222130301021213-3331311112012010)
- origin_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-2201012222010133-0003030133010220-0303310121312110-2023032313301321-2010100201113132-0123023330120310-1302322021113022-1200120320301000"></a>

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

<a id="canonical-3112123101123003-3200302332130230-0012233303110230-0031230001321000-3122303301210323-2020202303100310-0021302032011020-1212031223323212"></a>

### Direct properties for `origin_pool.use_tls.use_server_verification.trusted_ca`

<a id="canonical-0033321231111321-0021122100002033-0002313012010032-2033001032012121-0101022311110313-1300000023213001-0333001332123000-0011032303331300"></a>

#### `origin_pool.use_tls.use_server_verification.trusted_ca.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0120031303102010-3202033233202311-2100031110100130-2333122310300311-0133220022220333-2112012003320311-3022131020031122-1333303333000111"></a>

<a id="canonical-1103212210002101-0020321313132110-0132022212121020-0112132021300211-1300312302231010-3201031231213322-1111221002033110-2310122031112310"></a>

#### `origin_pool.use_tls.use_server_verification.trusted_ca.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1213323002030231-2213233300012332-3200310211110321-2220111312301032-0220000213111312-0111200113001202-0112010122102210-0120221303113323"></a>

<a id="canonical-3210021231302110-1321230032030230-0221312211031013-0333102020030112-3220313113022223-3232103303020202-1303210300122020-1320330310323012"></a>

#### `origin_pool.use_tls.use_server_verification.trusted_ca.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1321313303313232-3311030333223032-3001313302112102-3101102302200031-3102232021201232-2030300102023020-1121201330102011-3220200000303332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pool.use_tls.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0201211120230232-3020210323211012-3211212003311111-1101131212030001-0201030200023011-1211030022131011-0001120121022211-0132303232331100)
- origin_pool.use_tls.volterra_trusted_ca

<a id="canonical-3332331120312231-2022232330210111-3001213123022211-2210330212332232-3111100232211101-3321130211300000-2110132132301030-2310320121131211"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

<a id="canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- other_settings

<a id="canonical-1321303200011012-3100221032322110-0011320103002120-3033302110300233-0220132132031230-3312120123201231-2210310211120122-1332100310112230"></a>

Type: `"single"`. Computed.

Configuration parameter for other settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3113101003001200-0301013210321000-3311122231112031-2111231222103331-1211020233000122-3030112011003221-3202003123211312-2320010233133030"></a>

### Direct properties for `other_settings`

<a id="canonical-1111113223320002-0133122330331233-2133330322103310-3210122012032333-1221032100023321-2310132223333301-2032022001213011-1100311013221022"></a>

#### `other_settings.add_location` property

Type: `"bool"`. Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231): complete subsection reference.

- [logging_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2223103101102133-1321201122230001-0323100123120022-2010212213323231-0212002021131321-2330330122301331-0201233222123130-0011222322220210): complete subsection reference.

<a id="canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- other_settings.header_options

<a id="canonical-2122303333301233-1210311310210202-0101331012301200-2010103121331120-2312011111233321-0002210031121001-0110003011011200-0221012311011001"></a>

Type: `"single"`. Computed.

This defines various OPTIONS related to request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3323123300203200-1310333022220203-2030003110223002-3332301210013132-0102101002100212-0123130012213033-3332201322222313-2211012330123020"></a>

### Direct properties for `other_settings.header_options`

- [request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3001213110300203-3233102123311030-2011212033310213-3232321133112132-2102223022021022-2021132230010012-3103233330210011-3111310110311331): complete subsection reference.

<a id="canonical-1221202022030103-2211032031200223-2131220320012210-0203302211133032-0122320203200331-2313231321302133-1123030113222303-1213203132101221"></a>

<a id="canonical-2010010231010000-0221203220333332-3232311302203021-2003302100333300-1300311121201013-3030032001112203-0230102321023200-1123203030102303"></a>

#### `other_settings.header_options.request_headers_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1121012221312303-3200232310320122-0023103101313232-3333202101100212-1323331213323001-3101131011230223-0123021323200101-3332232323332201): complete subsection reference.

<a id="canonical-0330210033202202-3312200022331010-0330123001012311-2113121203123210-0123131301000113-3130032121013021-0111120031133123-0210220301323213"></a>

<a id="canonical-3310333012113321-3023232012300222-2213302100210330-0012101213331033-1131331130201131-0123120011211332-3202131331001311-3310331333033021"></a>

#### `other_settings.header_options.response_headers_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3001213110300203-3233102123311030-2011212033310213-3232321133112132-2102223022021022-2021132230010012-3103233330210011-3111310110311331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231)
- other_settings.header_options.request_headers_to_add

<a id="canonical-1210313011331123-3211302113230112-1031022031021211-3322123310231021-1200101033130000-3123010001000202-1012112100333111-0001123310003012"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3322301231312012-0332231113213301-1222133003211131-3330102101331210-0112313013222230-1113323000100120-3313210200120133-1310020200323021"></a>

### Direct properties for `other_settings.header_options.request_headers_to_add`

<a id="canonical-3033333303233031-2323311321321213-1233100113300020-1213110131220233-3013210331333111-0121302111013020-0303023221132123-0302320220302213"></a>

#### `other_settings.header_options.request_headers_to_add.append` property

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2021301030032031-0211102113130131-2312222133333111-1021132312003100-3221111113113021-1033313221132222-0032120230102111-0110313213322011"></a>

<a id="canonical-0111022110122000-1103310232000011-3312222102321323-1030302032033100-3003321011100123-3323030020230203-2303112113200131-2323203213301230"></a>

#### `other_settings.header_options.request_headers_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0330133111300023-1232203101020001-3233213322303122-3101323300332213-3222013020110000-3001231011202210-2220022013313012-1012032221231023): complete subsection reference.

<a id="canonical-1313313213132110-1103123313221213-0323120301002002-1120221321100311-3111321120313100-0023002313300302-0303302010233333-3200123032111311"></a>

<a id="canonical-2320312301020202-2210021323331020-0332111111001312-0021100320010200-2213103303101120-2020000003030201-1112122013101331-2121120202320130"></a>

#### `other_settings.header_options.request_headers_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0330133111300023-1232203101020001-3233213322303122-3101323300332213-3222013020110000-3001231011202210-2220022013313012-1012032221231023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231)
- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3001213110300203-3233102123311030-2011212033310213-3232321133112132-2102223022021022-2021132230010012-3103233330210011-3111310110311331)
- other_settings.header_options.request_headers_to_add.secret_value

<a id="canonical-2331302311020210-1323312122133310-2111033113201121-1000101332110113-1231001011010100-1122032200220300-1102200210132213-0322232312003330"></a>

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

<a id="canonical-1312300303201012-2232121323321131-1202301303100220-3013121203110203-2001120130022212-0331221023201301-0013022201020222-3323001233021111"></a>

### Direct properties for `other_settings.header_options.request_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1213210321223201-2100011113202212-1313321110032121-1210121001230220-2203332121112030-2122223230011023-1003110121220331-2023111101120320): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2123220110321212-1010103230021222-3012001003301222-1220302123122022-1123301212210101-2200212011223100-2222012231203233-0323310123021231): complete subsection reference.

<a id="canonical-1213210321223201-2100011113202212-1313321110032121-1210121001230220-2203332121112030-2122223230011023-1003110121220331-2023111101120320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231)
- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3001213110300203-3233102123311030-2011212033310213-3232321133112132-2102223022021022-2021132230010012-3103233330210011-3111310110311331)
- [other_settings.header_options.request_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0330133111300023-1232203101020001-3233213322303122-3101323300332213-3222013020110000-3001231011202210-2220022013313012-1012032221231023)
- other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0011022222303222-0332102213003011-0123113021313122-0211233101212221-1331312033032102-2030222030001103-0201223311322220-3332210123130332"></a>

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

<a id="canonical-3232122112002320-0210330220010332-2121203030322331-0320200220202111-0102230213222301-0210200102302021-1211203013211013-1103221233330320"></a>

### Direct properties for `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-2230031001110130-1110310311023023-1231100303030202-0211233211211011-3233301313310020-3302012200123211-1133112102133132-0112103332310300"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2112121013200231-1311200002231322-1103210232322103-2221022032231213-0333210122202103-2111330022200331-1112303300301023-1121113310213233"></a>

<a id="canonical-0313203132330003-1003133330331203-2022131020001222-3211211230212233-1301100113221222-0120123133203321-0001310223023321-3231233311313230"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-2221002100110001-0230001010233333-1303232331213210-2133013023323302-0033211010300003-3200013221310221-3231222211211002-1122111330230322"></a>

<a id="canonical-1011113031323222-1211030032131102-3231032301202021-3023300203321232-2103203101122131-0200021232211030-0210021231130130-0102111203100301"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-2123220110321212-1010103230021222-3012001003301222-1220302123122022-1123301212210101-2200212011223100-2222012231203233-0323310123021231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231)
- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3001213110300203-3233102123311030-2011212033310213-3232321133112132-2102223022021022-2021132230010012-3103233330210011-3111310110311331)
- [other_settings.header_options.request_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0330133111300023-1232203101020001-3233213322303122-3101323300332213-3222013020110000-3001231011202210-2220022013313012-1012032221231023)
- other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3030320100323312-3301131001301211-1202323003232220-0132213103112202-3210223222131203-0102123110332000-2103001121202210-0221013223212112"></a>

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

<a id="canonical-0033103213203322-0033221323033003-2201313301310223-2120110110212323-0130200311212221-1103121120303123-2201131133130120-0222230312213102"></a>

### Direct properties for `other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-2323020021020102-1320202322133303-2022101130110303-0211211200313231-3101132210200222-1030022301130111-2101222230233210-2103021201222303"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2313211313122203-2320320322222031-1021220230102320-1322020200112231-1120331313012200-0230130211223123-1112202010020100-2000132003302012"></a>

<a id="canonical-3213211010322121-2112102000302021-3311102010231200-2332101131010200-3301133020033302-2122111032202001-0310113130233123-1222111320223212"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-1121012221312303-3200232310320122-0023103101313232-3333202101100212-1323331213323001-3101131011230223-0123021323200101-3332232323332201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.response_headers_to_add` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231)
- other_settings.header_options.response_headers_to_add

<a id="canonical-1312220232000010-1213332110131300-2320331002213023-0120022021120203-3011321222320213-2102031131003132-1020212200302302-3331012333233123"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1001012010032200-2020303112100022-1130213233100213-0223202112221320-3202130133232103-1120202233103301-2121002020332323-1032101333023110"></a>

### Direct properties for `other_settings.header_options.response_headers_to_add`

<a id="canonical-2013313003231301-1011011220101333-2121223003111003-1233003123322032-1310233023132011-0231000010131012-1103003232022022-1000022221321212"></a>

#### `other_settings.header_options.response_headers_to_add.append` property

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0123111021230033-1030030203132013-3001220332001303-0201101230120001-2201321312311331-1123001210212002-1230303201100221-1223231232111123"></a>

<a id="canonical-2130321000122330-1300302300220001-2233321200303311-1210130322233030-1123012010323302-2012001212122310-3132011122330332-3032100322313211"></a>

#### `other_settings.header_options.response_headers_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3113030023011322-2102333220022101-1021022200110333-1213020300210232-2231022221320323-2302200032120332-3100231223231112-2221010210022312): complete subsection reference.

<a id="canonical-0222221210013032-2200002032132000-1333031111322333-1320000323130100-0031202231112000-1021103202110033-0200321030330211-3311320322122111"></a>

<a id="canonical-0320323302002032-2311321021223203-3012111302200300-2323112102310223-0232003022310222-2320003012032113-0231112311120311-3103302200232103"></a>

#### `other_settings.header_options.response_headers_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3113030023011322-2102333220022101-1021022200110333-1213020300210232-2231022221320323-2302200032120332-3100231223231112-2221010210022312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1121012221312303-3200232310320122-0023103101313232-3333202101100212-1323331213323001-3101131011230223-0123021323200101-3332232323332201)
- other_settings.header_options.response_headers_to_add.secret_value

<a id="canonical-3323021330302113-0231002202112011-0033121011011001-2012031302202311-3122002122212211-2023103021030120-1022223003131132-1320010110011112"></a>

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

<a id="canonical-0213100123111130-2200313331100011-3112233101203321-2123212310300103-2101123030031130-2210221113130311-2331213021320322-0013333201301300"></a>

### Direct properties for `other_settings.header_options.response_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2302100302212031-0322333201311102-1230122101233113-2211111233331031-0012303320113022-1133123231131311-3102231231003322-3331202321122033): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2123222131331002-3221231222213332-2022313221113011-3012022211022103-3131013311210213-2210211100033303-2103221010112233-3220002212210301): complete subsection reference.

<a id="canonical-2302100302212031-0322333201311102-1230122101233113-2211111233331031-0012303320113022-1133123231131311-3102231231003322-3331202321122033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1121012221312303-3200232310320122-0023103101313232-3333202101100212-1323331213323001-3101131011230223-0123021323200101-3332232323332201)
- [other_settings.header_options.response_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3113030023011322-2102333220022101-1021022200110333-1213020300210232-2231022221320323-2302200032120332-3100231223231112-2221010210022312)
- other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0302312310313032-2100013231231231-2100300132203211-2001312100321021-3010003320330012-1303333201330110-2012001113233200-3001111002113120"></a>

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

<a id="canonical-2111030033022210-0123032333230131-0203013202103100-0211231133200331-3200321111012303-2101132131103231-3230330013131133-1123211300121213"></a>

### Direct properties for `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1002111033231202-3222232300101230-0110201111013210-2033031022200002-1012031132011303-0312312333201222-3201312022013322-2001032313303001"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0103032313010102-2030231133021011-3001333233112100-1232312222212130-0200011212122012-0230113100102033-1101323301302203-1101220203330203"></a>

<a id="canonical-0122001032021332-0103033300221220-2222313211321120-2221011023220000-1021132300002320-3230033321132111-0203323012130201-3103120022130122"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3230101211301013-1310313021332202-0330331013303111-2113131232332331-1331102203203230-0232011230320110-2321320310130311-1102200322301220"></a>

<a id="canonical-3220201031131312-0032010030303121-3332033332202202-0132020103110232-0312213221030010-2131133212331202-0203030123201020-1201233300332331"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-2123222131331002-3221231222213332-2022313221113011-3012022211022103-3131013311210213-2210211100033303-2103221010112233-3220002212210301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.header_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1111103301003200-3130132203322010-3022023313321022-3230221000330020-3003012133331121-2031010120210231-0320203010131123-0110303302211231)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1121012221312303-3200232310320122-0023103101313232-3333202101100212-1323331213323001-3101131011230223-0123021323200101-3332232323332201)
- [other_settings.header_options.response_headers_to_add.secret_value](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3113030023011322-2102333220022101-1021022200110333-1213020300210232-2231022221320323-2302200032120332-3100231223231112-2221010210022312)
- other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0113130010011030-3022323311031331-1210013030301121-1201231312302122-1322030232033131-3332211322211332-0002200212113000-3110122123031320"></a>

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

<a id="canonical-3210320301023112-2323123321122120-1112210002031011-0310222330120230-2110001113220331-2132301320113121-0322300101003333-0210030212202311"></a>

### Direct properties for `other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-3302010200123330-3003201321331021-1100221003231303-2220103033210302-2133021023123002-2330202023100133-0102320300012131-2333301312113232"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3111321031322121-0232013303222313-1221031110003032-3233302312111330-0131231211203100-0112120002203113-3221330211203030-2332121003010122"></a>

<a id="canonical-3231231303031321-3123222333032223-1322103331031303-1103221000021301-3011223031022013-1033303302333210-2010000333323333-3322222001012321"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-2223103101102133-1321201122230001-0323100123120022-2010212213323231-0212002021131321-2330330122301331-0201233222123130-0011222322220210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.logging_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- other_settings.logging_options

<a id="canonical-1200032211021213-1312332122330211-3000123023131321-3222022222003223-1013122003223032-2011232313010320-2020030331202232-0020010312233211"></a>

Type: `"single"`. Computed.

This defines various OPTIONS related to logging.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3012330303223202-3323101232012032-0301013330212223-1200001323200012-0101213102222332-1321130211020030-2233301200313310-0203210030323131"></a>

### Direct properties for `other_settings.logging_options`

- [client_log_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3032101221312010-3330322012323121-3323031000021002-2032133212332312-3003210021131221-3312230312230000-0102323130200333-1320210332321022): complete subsection reference.

- [origin_log_options](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1310022203322020-1113212000111233-1003332001203230-3101200021031032-1320202113312011-1202222032211131-2200332101002333-0131103023312300): complete subsection reference.

<a id="canonical-3032101221312010-3330322012323121-3323031000021002-2032133212332312-3003210021131221-3312230312230000-0102323130200333-1320210332321022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.logging_options.client_log_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.logging_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2223103101102133-1321201122230001-0323100123120022-2010212213323231-0212002021131321-2330330122301331-0201233222123130-0011222322220210)
- other_settings.logging_options.client_log_options

<a id="canonical-0103232313110322-1311031200311313-0320111103030230-2123010112113003-0233232300002211-2011100021320010-0103121121101102-3311122303310212"></a>

Type: `"single"`. Computed.

Headers to Log. List of headers to Log.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
