---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-1030001320322301-0023323210131012-3203012102020030-3023002133112201-2130012100302310-2131101333202310-0311000201121333-0111011222232312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.no_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- https_management.advertise_on_sli_vip.no_mtls

<a id="canonical-3032111232130310-0101311100010111-0203033311133313-3323030110112221-0112012301303303-1102032322232302-2222030221313220-3311011203100312"></a>

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

<a id="canonical-1320230022303131-0203122310311023-2233021033023131-1130023120121103-0320031031230321-0102221123321003-0132102211000013-2321200111102031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_certificates` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- https_management.advertise_on_sli_vip.tls_certificates

<a id="canonical-2002011332302033-2103010220231001-2131132223300011-1100010011003023-0012012023230311-3000323303302221-1000303100301321-0133233221301133"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3233010213022202-3013123220320312-1313333000013003-0012101122011021-0200311132220332-1303113200100032-2121200000322022-0013021320313201"></a>

### Direct properties for `https_management.advertise_on_sli_vip.tls_certificates`

<a id="canonical-2312312122020313-2200133300310220-0231213303203133-2110303222222003-1032300220221333-0333122130001031-3002102121223310-0131023010100202"></a>

#### `https_management.advertise_on_sli_vip.tls_certificates.certificate_url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-0233223013120313-2110121233321202-0300212230313310-2202102320321010-1133030322223131-0133133321102023-0231032023302231-2130330102023331): complete subsection reference.

<a id="canonical-1103133323103302-1120223211130233-2102120032312201-0200212300233100-2330302333101011-1210312133022013-1033031222232132-2220003313022311"></a>

<a id="canonical-2322211131012121-2120212002211132-1321313103230132-3323231023201113-2121223223112121-2012331210223231-0111112102013223-3222022202031000"></a>

#### `https_management.advertise_on_sli_vip.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-1231020032111213-0310230100213121-0201223123233212-0003322011101322-1122113122112311-3232132300223013-3100121111203320-2011322220022022): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-002.md#canonical-3202013233313032-1122321011023203-3110102003123020-2123222112301232-3113233132130213-0123211321312332-1213330030132020-2001231220212130): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-2022233010303110-0201030120210332-0132322130310222-2221330303032013-1212333100220200-2210223113213111-0103101130131212-0203000310212323): complete subsection reference.

<a id="canonical-0233223013120313-2110121233321202-0300212230313310-2202102320321010-1133030322223131-0133133321102023-0231032023302231-2130330102023331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-1320230022303131-0203122310311023-2233021033023131-1130023120121103-0320031031230321-0102221123321003-0132102211000013-2321200111102031)
- https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-2133333032023132-2001313021223012-1301232203301223-3030303331231333-3333332000223112-3311232121231200-3112333013211313-1113133210330330"></a>

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

<a id="canonical-3232120111303220-3211332313010023-3022031102203001-2133322213211230-3221333100321310-2112021121110000-1310202101313322-1232220022010320"></a>

### Direct properties for `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms`

<a id="canonical-2110220311230131-1023202021222203-0013333322022220-1032133002132112-2112322023200131-2031112313333233-0201231201231113-0332221133223231"></a>

#### `https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1231020032111213-0310230100213121-0201223123233212-0003322011101322-1122113122112311-3232132300223013-3100121111203320-2011322220022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-1320230022303131-0203122310311023-2233021033023131-1130023120121103-0320031031230321-0102221123321003-0132102211000013-2321200111102031)
- https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-0323202120103231-2101300122033322-2310221022102003-1013302220223313-1010213003023320-3123313001223132-2133202033032120-0113201010112122"></a>

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

<a id="canonical-3202013233313032-1122321011023203-3110102003123020-2123222112301232-3113233132130213-0123211321312332-1213330030132020-2001231220212130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-1320230022303131-0203122310311023-2233021033023131-1130023120121103-0320031031230321-0102221123321003-0132102211000013-2321200111102031)
- https_management.advertise_on_sli_vip.tls_certificates.private_key

<a id="canonical-2003023012103221-2303020132131020-0210210010030010-3033031330033322-1102233211203113-3020202112213031-3123112003323313-3101020032233313"></a>

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

<a id="canonical-1010032212032311-0320022212012110-2220333132233020-1232310101032303-3012023001300323-3023313313021030-1330232210000330-1331231132213113"></a>

### Direct properties for `https_management.advertise_on_sli_vip.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-0310011011322103-2313103131313133-2203101110103031-2321300230332012-2200232220133303-3001310300111020-0321232333110323-2111311333232312): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-3210200201032121-2232112000323331-1301110320031032-1011002021221300-1233223021020200-2231100320002302-3300000320223122-2331311301203123): complete subsection reference.

<a id="canonical-0310011011322103-2313103131313133-2203101110103031-2321300230332012-2200232220133303-3001310300111020-0321232333110323-2111311333232312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-1320230022303131-0203122310311023-2233021033023131-1130023120121103-0320031031230321-0102221123321003-0132102211000013-2321200111102031)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-3202013233313032-1122321011023203-3110102003123020-2123222112301232-3113233132130213-0123211321312332-1213330030132020-2001231220212130)
- https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3123121232110011-3313311313212012-1300330330333101-2101323003032312-1321001223210113-2112032011321230-2012101123033331-1222211302322332"></a>

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

<a id="canonical-0000232122111323-3201103132301130-0200331302210221-2032013310323112-2132203122331221-2112321111203323-0120331000000233-3021301300010223"></a>

### Direct properties for `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3203302012303220-3033030311100321-2230320330213323-0200031213332212-2211130033210112-1020003132330322-1100233323230003-0210013303221011"></a>

#### `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2102230020303312-1220011101001200-3301033332012313-3001120200230313-3233122011111103-1111122132213033-2031332231130303-3332303033210201"></a>

<a id="canonical-0032301112133100-1111102112133031-0301333133310200-3201312232130022-1120000103200020-0103002321003130-1311211223231100-3231123122120110"></a>

#### `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1323121223331321-2020200230211112-3003301123100233-0130303003231020-2110101013333120-3111312231213220-0023001230001010-2101222331310012"></a>

<a id="canonical-1201120122201303-0102311122133230-3110222121120202-1133113200231010-1122123112320302-1000220221312202-1301332300131121-2102231212200330"></a>

#### `https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3210200201032121-2232112000323331-1301110320031032-1011002021221300-1233223021020200-2231100320002302-3300000320223122-2331311301203123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-1320230022303131-0203122310311023-2233021033023131-1130023120121103-0320031031230321-0102221123321003-0132102211000013-2321200111102031)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-3202013233313032-1122321011023203-3110102003123020-2123222112301232-3113233132130213-0123211321312332-1213330030132020-2001231220212130)
- https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-3003133322202002-1213102113032213-2303212210103202-0010032222011122-1300221311222031-1133330111111210-2101230310032121-1200300120211103"></a>

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

<a id="canonical-0212031301332203-1020020133030003-1210331320310201-2321221110123103-3220213023210212-1313120031013212-2131011220231032-2003330322310330"></a>

### Direct properties for `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info`

<a id="canonical-2220300013322301-1031303333033030-3203012010023322-3320333300012310-1023211202132320-1031310202302212-1311030013021000-0113300033102301"></a>

#### `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3022100031020120-1211233320231230-2100323200110313-3033210313001013-1303213102322312-3102002020003301-0203102321001021-0131203123203001"></a>

<a id="canonical-1223133303320223-2300231332130021-2013022202033332-1001103130331013-2320010231032320-1121332332232013-0200000112011320-0113212212223132"></a>

#### `https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2022233010303110-0201030120210332-0132322130310222-2221330303032013-1212333100220200-2210223113213111-0103101130131212-0203000310212323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-1320230022303131-0203122310311023-2233021033023131-1130023120121103-0320031031230321-0102221123321003-0132102211000013-2321200111102031)
- https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults

<a id="canonical-0020213310000231-2233132011003122-2001133111200230-3033111323112101-0023321032022012-0220302303020021-0020010210310001-3233100221222123"></a>

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

<a id="canonical-0110322210323010-2020033212112220-2221001300203322-1221033012030103-3023103310321232-3310201013003121-3032200300302203-0233131002303201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_config` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- https_management.advertise_on_sli_vip.tls_config

<a id="canonical-0111120012132101-3232121210300120-3212320033032323-2212131200132202-2110121322023331-2030121310131123-1313310030210020-2320303320122310"></a>

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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-1122121220221331-3003121120011113-1010123302200203-3021001103230013-1031303332123231-1123210111002110-2210113002023120-0333321222322131"></a>

### Direct properties for `https_management.advertise_on_sli_vip.tls_config`

- [custom_security](data-sources--nfv_service--reference--group-002.md#canonical-3001303023030313-0210100222230330-0000332301001333-3021232202033302-1232301033012130-1101020331031200-3021033210212220-1030133211200110): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-002.md#canonical-3003300102130220-1212221033213232-3221030202031102-2120331120303133-2111113021220203-1113220222222123-1123223010103312-3011130032323111): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-002.md#canonical-0020113322313331-2122133303022111-0230210100212313-2321030102231213-2323113110230111-1211102031203302-0030132001122321-2230033121332330): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-002.md#canonical-3322222221231111-2023310311232200-1132131211333000-2120330033102320-2002013303210101-3032310210202110-3001012320313122-3230301010003113): complete subsection reference.

<a id="canonical-3001303023030313-0210100222230330-0000332301001333-3021232202033302-1232301033012130-1101020331031200-3021033210212220-1030133211200110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-0110322210323010-2020033212112220-2221001300203322-1221033012030103-3023103310321232-3310201013003121-3032200300302203-0233131002303201)
- https_management.advertise_on_sli_vip.tls_config.custom_security

<a id="canonical-1323203100003013-2120031013223203-2121212313211112-0101132213112302-0200302233311101-2130111231033012-1033310010023110-1031303113012210"></a>

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

<a id="canonical-3320123001210230-3121122211012311-1131323002130321-0022300033302112-3022002301001223-2212202230333013-1103122210012130-3302122031232321"></a>

### Direct properties for `https_management.advertise_on_sli_vip.tls_config.custom_security`

<a id="canonical-3013102311030301-3020301121310222-3200112312300101-0021221311200022-3213112203031322-3011001220302232-1100013231120310-0332332302131331"></a>

#### `https_management.advertise_on_sli_vip.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1201100111002231-3121312012123012-2102230231333013-2102320210002330-1301210201332211-0031310013101311-1111111301001112-2211030232233230"></a>

<a id="canonical-0301301320301213-0122001230320213-1323200312213202-1101330321322002-1330330222200033-1011122310112322-0021210123221021-0011023332201122"></a>

#### `https_management.advertise_on_sli_vip.tls_config.custom_security.max_version` property

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

<a id="canonical-2323123121233223-1121100202301211-2223103223301111-0111010102010001-2021212211020302-1201232002323032-2000033110221100-1103001030100312"></a>

<a id="canonical-0130232313021010-2200320220030211-2333011333330210-1013031232103033-1331032231222010-2302031201331103-2021303121033022-2300132023201031"></a>

#### `https_management.advertise_on_sli_vip.tls_config.custom_security.min_version` property

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

<a id="canonical-3003300102130220-1212221033213232-3221030202031102-2120331120303133-2111113021220203-1113220222222123-1123223010103312-3011130032323111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-0110322210323010-2020033212112220-2221001300203322-1221033012030103-3023103310321232-3310201013003121-3032200300302203-0233131002303201)
- https_management.advertise_on_sli_vip.tls_config.default_security

<a id="canonical-0011111012213111-0231313131222030-0200233203223202-0131212311233033-1032121021312232-1203103102201110-0130131123103211-1123133323330221"></a>

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

<a id="canonical-0020113322313331-2122133303022111-0230210100212313-2321030102231213-2323113110230111-1211102031203302-0030132001122321-2230033121332330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-0110322210323010-2020033212112220-2221001300203322-1221033012030103-3023103310321232-3310201013003121-3032200300302203-0233131002303201)
- https_management.advertise_on_sli_vip.tls_config.low_security

<a id="canonical-0021201100033021-2022323301302233-2001011031322022-3000300013110112-3303310322313230-0332313022123133-0203013230111103-0031223331233113"></a>

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

<a id="canonical-3322222221231111-2023310311232200-1132131211333000-2120330033102320-2002013303210101-3032310210202110-3001012320313122-3230301010003113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-0110322210323010-2020033212112220-2221001300203322-1221033012030103-3023103310321232-3310201013003121-3032200300302203-0233131002303201)
- https_management.advertise_on_sli_vip.tls_config.medium_security

<a id="canonical-1002321133101300-1322312322332303-3032232333223223-2010220113200231-2220012022220132-3113101121012023-2002003123201231-1123100032212010"></a>

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

<a id="canonical-2300213023312122-3130032213033103-1123003220011120-0012033021323320-0202332113303322-0300001320031111-0113021103211133-3332310322320302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.use_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- https_management.advertise_on_sli_vip.use_mtls

<a id="canonical-1001022002323202-1003220103203221-3102321122222223-1220123331133033-0000000203210302-2311232223020231-3111330110201032-3021021110011332"></a>

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

<a id="canonical-2323332303300020-1000231300233130-2212132211311130-0030112330123111-1033121330203010-1222100002213322-0302223330033200-3320203130311000"></a>

### Direct properties for `https_management.advertise_on_sli_vip.use_mtls`

<a id="canonical-1321000201030302-0302101210302121-3202123002320020-0031030032231130-3322011302221230-3010321322213001-0003331310310023-0303030303032112"></a>

#### `https_management.advertise_on_sli_vip.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--nfv_service--reference--group-002.md#canonical-1122113023010121-0120120212121322-3122002222002100-1300212102313113-1232021322200032-3231120201022110-0010211320010320-3030120031213000): complete subsection reference.

- [no_crl](data-sources--nfv_service--reference--group-002.md#canonical-0130012211321301-2122021032213303-2101102030211101-1022212213102112-0311220211322112-0131203311333100-1330000030010333-0310003301223211): complete subsection reference.

- [trusted_ca](data-sources--nfv_service--reference--group-002.md#canonical-1302001323323303-2233012202022331-3200331122303011-0033313111213321-0023001013220011-2012210111023012-2033000302203311-1120011313122000): complete subsection reference.

<a id="canonical-1113312021332032-0322131313313120-1022030122113200-2311130121020202-0012301332211230-1120203221010310-0222203130131131-2122113131232031"></a>

<a id="canonical-2011312101131010-2223100022000320-0212233322210132-0310323132331111-3132210303302202-2323020301130003-0120122211223231-0311022111101203"></a>

#### `https_management.advertise_on_sli_vip.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [xfcc_disabled](data-sources--nfv_service--reference--group-002.md#canonical-0233321331200221-2131320020330212-0021102310202313-3303132301001023-3230122330330010-3222321111213312-2120313231211130-1110111023200322): complete subsection reference.

- [xfcc_options](data-sources--nfv_service--reference--group-002.md#canonical-2300320103003133-1221031333331121-0033223103101333-1112323023122231-1211213201311113-2221131221320031-3131313133300132-0001200202302110): complete subsection reference.

<a id="canonical-1122113023010121-0120120212121322-3122002222002100-1300212102313113-1232021322200032-3231120201022110-0010211320010320-3030120031213000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2300213023312122-3130032213033103-1123003220011120-0012033021323320-0202332113303322-0300001320031111-0113021103211133-3332310322320302)
- https_management.advertise_on_sli_vip.use_mtls.crl

<a id="canonical-2111311003313112-0212233103031023-3031003311030302-2101000110320333-3220033330211321-1303013332213120-1211021222001331-0210032203232210"></a>

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

<a id="canonical-2212220003003230-2202102203130200-2322130122022032-2323202322002003-0121303322010131-0121201031010332-2111121032333210-3130203121320000"></a>

### Direct properties for `https_management.advertise_on_sli_vip.use_mtls.crl`

<a id="canonical-0311310310131200-3003323300011133-3211003311302221-3000103220220102-1021130120012011-3300311102201331-0303122321213101-3123123022103121"></a>

#### `https_management.advertise_on_sli_vip.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3223312123303120-2102223013111003-3002311210103002-1000331300002211-0300022121110232-3022032131121120-1231320013000330-0300303030322101"></a>

<a id="canonical-1332211300010010-2021323301231020-0310203010303101-3023201323013012-3310302320320133-2100033113200212-1200101000231012-1032312210312330"></a>

#### `https_management.advertise_on_sli_vip.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1121131201333310-3021103220131320-2220213000111313-2332230122303000-1320310123302003-2103021001120101-1110123300011230-3003021330310132"></a>

<a id="canonical-2230112133133322-1303311023003022-0313330311011132-1132211013323332-3301002231101203-3113030023230111-2100202322221120-2011232021222302"></a>

#### `https_management.advertise_on_sli_vip.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0130012211321301-2122021032213303-2101102030211101-1022212213102112-0311220211322112-0131203311333100-1330000030010333-0310003301223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2300213023312122-3130032213033103-1123003220011120-0012033021323320-0202332113303322-0300001320031111-0113021103211133-3332310322320302)
- https_management.advertise_on_sli_vip.use_mtls.no_crl

<a id="canonical-1203102112003003-2222321112021301-2231202123113330-2201101232323102-0212220013233320-1121021123203311-3220332111033222-3302110132312200"></a>

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

<a id="canonical-1302001323323303-2233012202022331-3200331122303011-0033313111213321-0023001013220011-2012210111023012-2033000302203311-1120011313122000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2300213023312122-3130032213033103-1123003220011120-0012033021323320-0202332113303322-0300001320031111-0113021103211133-3332310322320302)
- https_management.advertise_on_sli_vip.use_mtls.trusted_ca

<a id="canonical-1113102033001013-1230300323112202-2303102330011011-2130120110011203-0300033333023231-3201031330100333-3023202031232110-3322121020033103"></a>

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

<a id="canonical-2330212103203300-1310311010221321-3102132120010202-1030031300312323-0313311002333321-2023003032200113-2010231331311123-0331000201330023"></a>

### Direct properties for `https_management.advertise_on_sli_vip.use_mtls.trusted_ca`

<a id="canonical-1122032020201031-3031203120130222-3310201212101200-0302003332203231-0021001110101021-1223233210022331-1303002132011333-0122120223133213"></a>

#### `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0120223310210023-2023133302323012-1302221123102120-3131120031222320-0322322110130130-0133223000002303-2011200102301022-2033020210013211"></a>

<a id="canonical-3113022232302332-3122311233220022-0001310233331032-2120010131320230-3032330231123012-0212011120120232-1033012130220022-1033030002120222"></a>

#### `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2213122103030213-0013221003001121-1023032033231202-3233020111321300-3101332112323322-0331123011202210-2323012210322022-0230320020212330"></a>

<a id="canonical-2020330101130321-0211212232033221-2331331312331202-2223112031120313-2033021003320311-2123000003212310-3031022121101313-0130023010220121"></a>

#### `https_management.advertise_on_sli_vip.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0233321331200221-2131320020330212-0021102310202313-3303132301001023-3230122330330010-3222321111213312-2120313231211130-1110111023200322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2300213023312122-3130032213033103-1123003220011120-0012033021323320-0202332113303322-0300001320031111-0113021103211133-3332310322320302)
- https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled

<a id="canonical-0231200133003302-2200110003112212-0232101020203002-3333030301310333-1032130213023313-1310132100232200-2010131100031221-3020201033101333"></a>

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

<a id="canonical-2300320103003133-1221031333331121-0033223103101333-1112323023122231-1211213201311113-2221131221320031-3131313133300132-0001200202302110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_sli_vip.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--reference--group-001.md#canonical-2103233100100200-1030211313230233-2302302311300133-0321323232230132-2211220230113113-0320013303102211-3030032113312011-2011121301323112)
- [https_management.advertise_on_sli_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2300213023312122-3130032213033103-1123003220011120-0012033021323320-0202332113303322-0300001320031111-0113021103211133-3332310322320302)
- https_management.advertise_on_sli_vip.use_mtls.xfcc_options

<a id="canonical-1203011120223321-1122002020121213-0232211113032302-3002320021302233-2010033032012130-0002101221211220-0003003032332212-2123112233021230"></a>

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

<a id="canonical-1233102012310101-2121220132321330-1002231130302120-0110333213313012-2003023210321022-2301002021021303-3130020103320033-0032333110221011"></a>

### Direct properties for `https_management.advertise_on_sli_vip.use_mtls.xfcc_options`

<a id="canonical-1200202322022330-3131313001210130-0131110313131220-3121213132113322-1201132121032011-0300031110022303-3133120210103231-1011010112102021"></a>

#### `https_management.advertise_on_sli_vip.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- https_management.advertise_on_slo_internet_vip

<a id="canonical-1203311020012322-1033102332022113-3010133021301203-3032233002303212-1000302333022202-2130111233333110-1111002132103113-0233112103101201"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-0320203231132031-3200001103231013-0101213301102201-2223220302221002-1103120000021001-0113223022100012-1022003102001012-2302330313102313"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip`

- [no_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2222311220220120-1033133303300102-0302311332322221-0033101220103010-3312300331132010-1301233220203010-2131223131100100-3211221320202331): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-0301331331032233-1133002331031101-3330213221020131-0232023323233211-0132321030100232-1131120210101333-0001003102003212-0131312223100312): complete subsection reference.

- [tls_config](data-sources--nfv_service--reference--group-002.md#canonical-1320330030222102-1230112010022200-2233010333113112-3202330221202212-1103020221102021-2001331331001221-3213223132020201-0331212313133013): complete subsection reference.

- [use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2133311313230323-2312031102220113-0231032111301301-2101330223210021-1220120303003222-0301102232300332-3233200231301001-0002100012111322): complete subsection reference.

<a id="canonical-2222311220220120-1033133303300102-0302311332322221-0033101220103010-3312300331132010-1301233220203010-2131223131100100-3211221320202331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.no_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- https_management.advertise_on_slo_internet_vip.no_mtls

<a id="canonical-3232111202132332-1232222320111223-1022120320323000-0232303031023312-1122022222320111-1012200212003222-3111023101120302-3210313323101303"></a>

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

<a id="canonical-0301331331032233-1133002331031101-3330213221020131-0232023323233211-0132321030100232-1131120210101333-0001003102003212-0131312223100312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_certificates` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- https_management.advertise_on_slo_internet_vip.tls_certificates

<a id="canonical-1130033232300013-0312233123122301-3300210001110330-3322203312320000-0013310112310303-2132031322320023-2331102130030313-3133201223130301"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1122320302311220-3302230313111003-3120031022010113-0311133110321302-0113211232331032-3131321001001322-2332202332200123-1230220031221311"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.tls_certificates`

<a id="canonical-1010202301331020-3122320111311201-2122232120002320-2213011123111012-3110201232021231-1030321021111022-2301302001101021-2122121232312122"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_certificates.certificate_url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [custom_hash_algorithms](data-sources--nfv_service--reference--group-002.md#canonical-1303013122120222-0303021213213032-1110021233120012-0102322301301330-1010300112130213-3203303133213321-1120232121223033-1011213130032130): complete subsection reference.

<a id="canonical-3003233333100213-3013231033112131-1032121222102330-1213122122232301-3201330122030113-0123312021333302-3320020010002232-2100122222133231"></a>

<a id="canonical-1103222232112003-1110231101221031-1221323332101021-2113110301020302-3202320000131102-1200210002332001-1321133320003132-1123303123301003"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--reference--group-002.md#canonical-3203331030110321-1331311010210130-2232311322320213-0202100022223113-0321302133100220-2331002112022222-1300220232221323-0313110130222233): complete subsection reference.

- [private_key](data-sources--nfv_service--reference--group-002.md#canonical-1303301201200331-2302003222013303-1310233222130300-3302331303220213-1313131121322233-2000122000310220-2031313012030211-1122211230122030): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--reference--group-002.md#canonical-2021131033222010-0103230000223333-3200032000022211-1112001102132330-1022231321323312-0303112001112233-3201003110020330-1220011312130213): complete subsection reference.

<a id="canonical-1303013122120222-0303021213213032-1110021233120012-0102322301301330-1010300112130213-3203303133213321-1120232121223033-1011213130032130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-0301331331032233-1133002331031101-3330213221020131-0232023323233211-0132321030100232-1131120210101333-0001003102003212-0131312223100312)
- https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-3111312223003312-2332031022132100-0121333010211321-1300133203003213-3321221123030123-1003230312301233-1300123331023112-2231112232011300"></a>

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

<a id="canonical-1102123210103120-1033011220130221-3002112103120021-2333120212102133-1100331312221001-2233312111032013-0021120211332013-0132123002321232"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms`

<a id="canonical-3031202111312232-3233323230022211-2031321222202333-3231200310203321-3110321002230332-2132000213312020-0201103301012013-1230230001333202"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3203331030110321-1331311010210130-2232311322320213-0202100022223113-0321302133100220-2331002112022222-1300220232221323-0313110130222233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-0301331331032233-1133002331031101-3330213221020131-0232023323233211-0132321030100232-1131120210101333-0001003102003212-0131312223100312)
- https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-2113310221322311-3000133132033220-2000130022203102-3121301220223331-0000301101033122-0121323121201131-2222222332022311-2320311110202332"></a>

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

<a id="canonical-1303301201200331-2302003222013303-1310233222130300-3302331303220213-1313131121322233-2000122000310220-2031313012030211-1122211230122030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-0301331331032233-1133002331031101-3330213221020131-0232023323233211-0132321030100232-1131120210101333-0001003102003212-0131312223100312)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key

<a id="canonical-3011002031002112-3303303102112202-3332200211030323-3100231213033022-2211321321302123-0020232011313020-1311113110331110-0213202101330031"></a>

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

<a id="canonical-1130132232212220-1311312031232223-0331223230232113-0020321020211122-1110202332331110-0200120302231111-0021222032123233-3200212002102221"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-3222120133330231-3321231302010230-0220202213110303-2311020330230311-2211023220213223-1303311212202003-3313132311022011-1232133000033302): complete subsection reference.

- [clear_secret_info](data-sources--nfv_service--reference--group-002.md#canonical-1312103300011133-3331111020103020-2201012311012111-1333020212323220-2220010133230331-1320130000123121-1012112101311011-1311222203332323): complete subsection reference.

<a id="canonical-3222120133330231-3321231302010230-0220202213110303-2311020330230311-2211023220213223-1303311212202003-3313132311022011-1232133000033302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-0301331331032233-1133002331031101-3330213221020131-0232023323233211-0132321030100232-1131120210101333-0001003102003212-0131312223100312)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-1303301201200331-2302003222013303-1310233222130300-3302331303220213-1313131121322233-2000122000310220-2031313012030211-1122211230122030)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1213212033222310-2323301103130310-2331013331321013-2100030220102033-0332320231201200-0220120131031000-0201332120110320-3330222303333111"></a>

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

<a id="canonical-1020221103330222-1322222121320130-3311311112100021-2202302011113211-2110300012222031-1210213110020230-0233321231111132-0010002221112203"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0023023033103323-0011210333113002-3313110002022112-0313210010203221-0222201333100030-3312023130131003-3100212213030030-2131200223312111"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3011121102330011-0132232233033110-1003332122000231-1023301012020020-2033301213201320-2231211010013200-1001321310310121-2303333232313020"></a>

<a id="canonical-0112312031320202-3302111200333313-0032313200321112-2330303321233020-2133031122221001-3122103002310132-3220033100111012-2032221122132211"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2013221312203301-2111202220223331-3133201103210313-1333100102103331-0122000120010030-3120113121112333-2311123110222001-3102320311300303"></a>

<a id="canonical-0311332121133312-0032231121002120-0201001330311311-3001111332303020-0020133200120331-2210113200321201-0112132013110323-0033031233020031"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1312103300011133-3331111020103020-2201012311012111-1333020212323220-2220010133230331-1320130000123121-1012112101311011-1311222203332323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-0301331331032233-1133002331031101-3330213221020131-0232023323233211-0132321030100232-1131120210101333-0001003102003212-0131312223100312)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](data-sources--nfv_service--reference--group-002.md#canonical-1303301201200331-2302003222013303-1310233222130300-3302331303220213-1313131121322233-2000122000310220-2031313012030211-1122211230122030)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-1303210011000222-1110120000133001-3322101000312030-1100032221002100-3133022222210332-2023231201312131-2201331012013102-3332113321121310"></a>

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

<a id="canonical-0013132313013232-0201333021030101-2000201211020012-0232202321201131-2202122030022002-0202313202111031-2103030202000232-2123132022031133"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info`

<a id="canonical-1000002003101031-3200003330120130-0300311311301331-3020210232111202-3120221003101313-3020323013130031-0201203332331002-2131013011113301"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0021321111030313-1321320233123130-3203012230300301-1303113301122101-3331300312030130-0323231031131022-0233002213201112-0031321203302230"></a>

<a id="canonical-3220313103301212-1333111320232201-0112210201111201-1002101010011122-1311133211113301-3133123302101301-1102020332013032-3213022323231333"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2021131033222010-0103230000223333-3200032000022211-1112001102132330-1022231321323312-0303112001112233-3201003110020330-1220011312130213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--reference--group-002.md#canonical-0301331331032233-1133002331031101-3330213221020131-0232023323233211-0132321030100232-1131120210101333-0001003102003212-0131312223100312)
- https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults

<a id="canonical-2013231231231111-2002232233030233-3011211000112013-3130221030120032-1233013220130230-2311132113232223-1330132111111130-1233030210311301"></a>

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

<a id="canonical-1320330030222102-1230112010022200-2233010333113112-3202330221202212-1103020221102021-2001331331001221-3213223132020201-0331212313133013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_config` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- https_management.advertise_on_slo_internet_vip.tls_config

<a id="canonical-2133101321102213-3301133322312003-0321023330322310-1231221102032100-0111222103330301-0121021033013033-3310020021031010-0311012233012233"></a>

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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-3300102200220333-2123202110232101-2332131202211120-1133210023013332-2311233122111223-2300132233100113-3133332330031121-3202000222130022"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.tls_config`

- [custom_security](data-sources--nfv_service--reference--group-002.md#canonical-2331100332032222-1012112131010300-2230301101201112-0121010202312123-3131132311331330-0320230011333323-0331322020330013-0113313122132201): complete subsection reference.

- [default_security](data-sources--nfv_service--reference--group-002.md#canonical-0021023202311333-2031333013032131-0102333102213103-2120223201023122-2210312003022311-1100002113233133-3230123012231101-2333321231300111): complete subsection reference.

- [low_security](data-sources--nfv_service--reference--group-002.md#canonical-1021102033121331-2023003111101010-2320323103031303-2221123321013222-1223102103022010-2201331303201230-3101113211001230-2330021103200313): complete subsection reference.

- [medium_security](data-sources--nfv_service--reference--group-002.md#canonical-1122131120301323-0111030202322123-1213320003111133-3300020011212012-3010312200221120-3211311112121221-3332013231121211-0331123223002210): complete subsection reference.

<a id="canonical-2331100332032222-1012112131010300-2230301101201112-0121010202312123-3131132311331330-0320230011333323-0331322020330013-0113313122132201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-1320330030222102-1230112010022200-2233010333113112-3202330221202212-1103020221102021-2001331331001221-3213223132020201-0331212313133013)
- https_management.advertise_on_slo_internet_vip.tls_config.custom_security

<a id="canonical-1100210220010022-1333331213031102-2111010103133000-3331302233000033-3211032012102222-1022111322032302-0332303310331322-3032323202231223"></a>

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

<a id="canonical-0012000030003030-3320211113311301-3120301220003011-0220221320231003-0312001203030203-0121331201310023-3023312321300332-3221021103030022"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.tls_config.custom_security`

<a id="canonical-2012313333233111-0301320331031001-0101120230312102-0223022022222111-2223001010301300-1133232313201313-0101011000212103-3301102330212110"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0203313312110220-2030130022202300-0013111002113122-0301330101110311-0220112123012100-3313211221233120-3312121001320012-0220302222001202"></a>

<a id="canonical-2111333311213313-1033312111200131-2022321120121223-0023203313203323-1032002103011313-1223203220110021-0213111202111312-0102031213231211"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version` property

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

<a id="canonical-0230122022302300-0032323100102022-3210231213213201-2112323132111131-1000011133300221-2000100322200230-0221313223022012-0213022122303013"></a>

<a id="canonical-2013030020100100-3113101011303111-0100311223302333-2231033111223103-1132131001332012-0003202302112333-3211010322102212-1300303111213112"></a>

#### `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version` property

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

<a id="canonical-0021023202311333-2031333013032131-0102333102213103-2120223201023122-2210312003022311-1100002113233133-3230123012231101-2333321231300111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-1320330030222102-1230112010022200-2233010333113112-3202330221202212-1103020221102021-2001331331001221-3213223132020201-0331212313133013)
- https_management.advertise_on_slo_internet_vip.tls_config.default_security

<a id="canonical-3010213023313123-1232320013323033-0233033000211101-3302201302130303-1022103032320213-3123020133110310-2103112020120203-1331332312210200"></a>

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

<a id="canonical-1021102033121331-2023003111101010-2320323103031303-2221123321013222-1223102103022010-2201331303201230-3101113211001230-2330021103200313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-1320330030222102-1230112010022200-2233010333113112-3202330221202212-1103020221102021-2001331331001221-3213223132020201-0331212313133013)
- https_management.advertise_on_slo_internet_vip.tls_config.low_security

<a id="canonical-3212111233100100-2312313110310021-0121302321322213-2123011031322322-1303222303123211-1001211032130330-3101101121230131-3022323133113323"></a>

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

<a id="canonical-1122131120301323-0111030202322123-1213320003111133-3300020011212012-3010312200221120-3211311112121221-3332013231121211-0331123223002210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--reference--group-002.md#canonical-1320330030222102-1230112010022200-2233010333113112-3202330221202212-1103020221102021-2001331331001221-3213223132020201-0331212313133013)
- https_management.advertise_on_slo_internet_vip.tls_config.medium_security

<a id="canonical-3300030310231313-3300001210102222-3103200200003330-1332330123310210-0331021121303210-3231133003211312-2020300323113210-0001110133323331"></a>

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

<a id="canonical-2133311313230323-2312031102220113-0231032111301301-2101330223210021-1220120303003222-0301102232300332-3233200231301001-0002100012111322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- https_management.advertise_on_slo_internet_vip.use_mtls

<a id="canonical-3313100011123200-0111211110233233-3222320301030333-2130023333331312-1110122032320312-0023202002320032-1202031120002202-0100313303112131"></a>

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

<a id="canonical-0133221212103110-1203212101021013-0332320232000032-1323333321122133-3201210003000122-3231311111212300-1033223303222111-0112312001103003"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.use_mtls`

<a id="canonical-1223302010310002-3032230321001023-0301213100122303-1033023312330103-0331220330130123-0303130003300022-2321120013331032-0231130202002012"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--nfv_service--reference--group-002.md#canonical-1212313011012221-3302103300332020-3211311222302202-0310303230202012-2100000120132002-0221101111113031-2232322120112002-2310100131001313): complete subsection reference.

- [no_crl](data-sources--nfv_service--reference--group-003.md#canonical-2023003310212210-0333201133031110-1002202332200232-0222222010323012-0030320111222123-3332300202131113-0102011031023033-0203211211323021): complete subsection reference.

- [trusted_ca](data-sources--nfv_service--reference--group-003.md#canonical-2031221030012123-1321032110300131-3211202312233202-1302222030301332-1201222310331232-1221312202031221-0012002022330313-2320211120102002): complete subsection reference.

<a id="canonical-0203300203210101-3012300320213023-1310103332331102-3022011013230211-1112320211121220-0303121021322102-0220012112221322-0021000211131202"></a>

<a id="canonical-3200030211102020-2330131210203200-2111300231332001-1130303333120320-3300030032201213-2021121301222133-1200120031233010-1000220132233033"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [xfcc_disabled](data-sources--nfv_service--reference--group-003.md#canonical-1122022200012120-2120301112202011-3300222101201102-2213231323302120-1323121000233121-0331231101131322-0302312323113031-2231030212302013): complete subsection reference.

- [xfcc_options](data-sources--nfv_service--reference--group-003.md#canonical-3133220123300002-0032012031333210-1302212011031201-2301012310203002-0120210222030320-2013132212330100-3032333231330332-1313001023320203): complete subsection reference.

<a id="canonical-1212313011012221-3302103300332020-3211311222302202-0310303230202012-2100000120132002-0221101111113031-2232322120112002-2310100131001313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300)
- [Property reference](data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [https_management](data-sources--nfv_service--reference--group-001.md#canonical-0322223112333211-3323030031122111-0131002222131221-2002123332012320-0013112013120020-2030102011102030-0020230010011012-0231301033231112)
- [https_management.advertise_on_slo_internet_vip](data-sources--nfv_service--reference--group-002.md#canonical-0110230011001003-1231323320021322-1011022232000223-2320332312110222-2231203031232231-3110223033333112-2310310031003031-3311231331100133)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--reference--group-002.md#canonical-2133311313230323-2312031102220113-0231032111301301-2101330223210021-1220120303003222-0301102232300332-3233200231301001-0002100012111322)
- https_management.advertise_on_slo_internet_vip.use_mtls.crl

<a id="canonical-1100102110000330-1313323030210113-0230232331201102-0220120030221021-2110121323303300-1213002011030110-2322030122222033-1331013123101020"></a>

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

<a id="canonical-3320132102022322-1212121320332322-0223201001311213-3212203023102013-1111331330112333-3113110232021202-0320332332310221-0102011210331202"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.use_mtls.crl`

<a id="canonical-1212132101303310-3213021103310002-2020030023020311-0222112213310303-1212033103033013-1202311022220121-1320000330103230-2000331011310212"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2110121001033031-0113223121121100-2331130010202132-2313233202030100-0003222221331320-1201331211222101-0032330003121230-2311033330132203"></a>

<a id="canonical-2021130302220130-2000130013130221-0333321131020100-0310231001302023-2021201302112123-1031313100233000-1233221110322001-0112002000233101"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2302001202321232-1332032300010312-3222221303311323-3233210013102213-1103332023133312-0203230101012331-0031320131331202-1111202030103113"></a>

<a id="canonical-3222100201211212-2132213101123203-1101302030102102-0020332311132233-1131103312230212-1132111203223000-1102320210102033-1301111012201001"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
