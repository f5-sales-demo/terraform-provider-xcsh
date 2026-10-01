---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2001300030332102-1310112222103330-2233331321222333-1101103201103233-3222330122102131-0002022203201100-3222300301221221-0113202311111102"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params — tls_cert_params / 222202103110 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-3000033031031211-1223100230330330-3022003003300111-2013230002130321-3003110311302332-1021322020330011-2030121311030032-0330003231012301"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

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

<a id="canonical-3120123011220332-1302201233223311-1222112302133332-2111122113121112-0213130202211203-0320222331212123-0120220312211312-0101201332202111"></a>

## Direct properties — tls_cert_params / 222202103110 / 3

- [certificates](data-sources--workload--reference--group-006.md#canonical-1233103322231013-3132102131120120-0023002223332110-0003300202230021-3330131032322121-3310320103022122-3331030033202010-1213132331303011): complete subsection reference.

- [no_mtls](data-sources--workload--reference--group-006.md#canonical-2130213330013201-2101021212210302-3111030201103132-2333333230211011-3012030013100002-2231110202033012-1002102220331311-3202200223133132): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303): complete subsection reference.

<a id="canonical-2101230231121003-2303133200333232-3101103311321031-3321121113113000-2221003100200210-3320301201322213-0101213112120322-2012313101012113"></a>

## Next pages — tls_cert_params / 222202103110 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates](data-sources--workload--reference--group-006.md#canonical-1233103322231013-3132102131120120-0023002223332110-0003300202230021-3330131032322121-3310320103022122-3331030033202010-1213132331303011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls](data-sources--workload--reference--group-006.md#canonical-2130213330013201-2101021212210302-3111030201103132-2333333230211011-3012030013100002-2231110202033012-1002102220331311-3202200223133132)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1233103322231013-3132102131120120-0023002223332110-0003300202230021-3330131032322121-3310320103022122-3331030033202010-1213132331303011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113032210213122-0232220332232122-2011002302333221-1100010232133313-0120331032312112-2003102223133113-0220112003001300-3120000322313111"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates — certificates / 010012023102 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-0033031102031232-2220100003220230-0233131122331023-2000230032332231-0321202022031203-0321220122233112-0223110100313323-0021332212123201"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0303220130122101-2202202333113330-2331202032031122-1010002320203300-0202021213133002-3020321211111132-1013133312203003-0011322213103332"></a>

## Direct properties — certificates / 010012023102 / 3

<a id="canonical-2111321023111133-0201323320312030-2210323023110032-3103320030330313-0223200020322120-2103332231301133-1230210013021001-1110200113213320"></a>

<a id="canonical-2221330220110022-2001013220112323-3031003230330223-3213220230132201-2112332132123000-0123133101233333-3211220001000103-1310020023132231"></a>

## name property — certificates / 010012023102 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0223022220300230-1103212122302000-3100231232320301-3332110211002101-1000322022111133-0230102313130023-1121000200302333-0320123332303121"></a>

<a id="canonical-3202232331321233-3312133323032312-2232320212331232-1302322232203022-2202113323133203-1032313000312321-0211212213130121-0130101001133010"></a>

## namespace property — certificates / 010012023102 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2122102103111113-2213310111202112-1311201212013123-2011312002232123-3301000010202023-2212111323031120-1300303211213102-3112021012230120"></a>

<a id="canonical-2010121012230033-2020222321311323-2013333013133212-3132320022231331-1112011303233132-3123333113132221-2020101311301220-1232233211133001"></a>

## tenant property — certificates / 010012023102 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1023322103111133-0332220201132213-0201121312301213-1112311013211321-3333211100230330-0033032210122122-3231122233311323-3223211312011110"></a>

## Next pages — certificates / 010012023102 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2130213330013201-2101021212210302-3111030201103132-2333333230211011-3012030013100002-2231110202033012-1002102220331311-3202200223133132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020331310313111-0312020203123220-3331203223130323-1110120311230300-2222120333001321-3011321233230030-2310123132010232-0332003200023101"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls — no_mtls / 212223223122 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-1320100213230332-1302331223212013-1020111301101310-0031323023111113-3010331110103333-3123021312130322-2132211221123131-2120201323322232"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2131332320033111-0331303032103022-2221200131221001-1130032111103031-2130301133201330-2233013111313200-2123032210330212-1203300121322002"></a>

## Direct properties — no_mtls / 212223223122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330010302000122-2021323123212221-3111031313331132-3210213322013222-2303100102122300-3120330011211311-1223011321301020-3022011110121310"></a>

## Next pages — no_mtls / 212223223122 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310333002320121-1301322300313201-1130001203223132-3100202222102102-2303100111330011-0303302010311321-2132333132102312-2332030103201321"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config — tls_config / 311223121203 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-3003023211012013-2020310121211100-3112000022100231-2130001332002322-2301213310303002-0003121021202310-2102322100332101-3023232131030020"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

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

<a id="canonical-0033231200310123-1021121210033212-2012031100222030-0111312232211333-2020203232100233-2311020100133232-1103223321023022-3030110210011202"></a>

## Direct properties — tls_config / 311223121203 / 3

- [custom_security](data-sources--workload--reference--group-006.md#canonical-0030030303321112-3223110120313032-0023101323032122-3003213122022111-3222030310211012-1002301210003013-3331030212013301-2003112330113131): complete subsection reference.

- [default_security](data-sources--workload--reference--group-006.md#canonical-3131322121310113-0112331022330003-0001120122201112-3031111202120013-1203032233020111-3333220113333033-2211023200230133-3120301300211101): complete subsection reference.

- [low_security](data-sources--workload--reference--group-006.md#canonical-2133032012213201-3121030023111313-3310021003223130-1301322320303233-0121133211203103-0220201030133003-0123300231032121-0313210322011201): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-006.md#canonical-1031323012313130-3302133223032022-0131232220111132-1132301232213332-2123021303231101-3023223301323011-0322310002031202-3112231133121213): complete subsection reference.

<a id="canonical-3310021012110000-3012320212223130-2023223322012320-0322300111311312-0131030302032201-0300332231311333-1001131110223000-3310113022020310"></a>

## Next pages — tls_config / 311223121203 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](data-sources--workload--reference--group-006.md#canonical-0030030303321112-3223110120313032-0023101323032122-3003213122022111-3222030310211012-1002301210003013-3331030212013301-2003112330113131)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](data-sources--workload--reference--group-006.md#canonical-3131322121310113-0112331022330003-0001120122201112-3031111202120013-1203032233020111-3333220113333033-2211023200230133-3120301300211101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](data-sources--workload--reference--group-006.md#canonical-2133032012213201-3121030023111313-3310021003223130-1301322320303233-0121133211203103-0220201030133003-0123300231032121-0313210322011201)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](data-sources--workload--reference--group-006.md#canonical-1031323012313130-3302133223032022-0131232220111132-1132301232213332-2123021303231101-3023223301323011-0322310002031202-3112231133121213)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0030030303321112-3223110120313032-0023101323032122-3003213122022111-3222030310211012-1002301210003013-3331030212013301-2003112330113131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312102023013332-2000212000003113-1100023223013103-3302331311313310-1232220232222331-2301012203112012-3312212232133222-1130330331100111"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — custom_security / 000303123131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-1300020121030101-1022323032032203-1020023213003200-3003022023332013-2111123111232000-2203201230122220-1031312101312300-0130000012002232"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

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

<a id="canonical-2033023203210122-0023000122002301-2203012101133022-3132311201132203-2223121133222012-3130033012003020-0110311111111232-0020332231112313"></a>

## Direct properties — custom_security / 000303123131 / 3

<a id="canonical-0203201222302023-2020032122320102-0133022212310123-0303300002112011-3211010103321111-1213301332123333-3013101021021122-2202320020230120"></a>

<a id="canonical-2003202002223331-2000302123011112-3332032321133120-0302021031330301-0010120101220211-2200312131220110-1231222102011131-2100333303030311"></a>

## cipher_suites property — custom_security / 000303123131 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0211212320021113-0133130133001101-2212232303202221-3033211330000100-2012203332002131-0030332222232111-3031023000232011-3003300010031032"></a>

<a id="canonical-3001033102211031-2313112123133302-0303031120022200-0230002100320123-2321323030233130-3103212223202323-0120122313132100-2003103233102212"></a>

## max_version property — custom_security / 000303123131 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-2032213100203033-3231031233231112-2213102120231200-3212020022302331-3233313003111102-2301301331130033-1202333201011210-1323210010332301"></a>

<a id="canonical-0010322332312331-3022013033031030-0302230132000233-3113221213320233-1110033213331110-1112231322222331-2330122022201201-1331221010233103"></a>

## min_version property — custom_security / 000303123131 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-2200230123300210-2100122003000321-1213022030013030-2032222313321322-1300202201133021-2221023212331232-3220013213302021-3213310111112013"></a>

## Next pages — custom_security / 000303123131 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3131322121310113-0112331022330003-0001120122201112-3031111202120013-1203032233020111-3333220113333033-2211023200230133-3120301300211101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102221111210222-3333231101210311-2130330020100021-1031021100311313-2212133011022101-0103221023302201-1312220312030100-2311201333010022"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security — default_security / 111003033301 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-1230032220230313-3320301103121210-2133003110302211-2321120101001022-0012020001123233-2010312022013321-2332001210030330-0231332130232203"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-0110022331113111-0013011133011203-0313233222232111-3010310100003132-1101113132330222-1111122011310232-0133033300212132-1121311110300321"></a>

## Direct properties — default_security / 111003033301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200011113311311-0321103001102313-2311221021211032-0133113001323121-0310200320013003-0220212101323001-1210313301013312-3220131333110211"></a>

## Next pages — default_security / 111003033301 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2133032012213201-3121030023111313-3310021003223130-1301322320303233-0121133211203103-0220201030133003-0123300231032121-0313210322011201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203332113321313-3311212310001322-2332011031211221-2322212203210331-1002201330121133-3302101230220130-0211032223103132-0130132113332022"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security — low_security / 312012132222 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-1103321002032101-0012011312013010-2300101310312302-0133010011132221-2101321102123021-3122122110230132-1220301112110301-3331210203031023"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2011213131111310-2110212221101332-1031122303210013-1311330221311313-0012321323111301-0011011121000321-2002032112111230-1011022121201311"></a>

## Direct properties — low_security / 312012132222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321301032011312-0332031210123031-3302201202320223-3233020013231313-3103223310111032-3200230001203013-0201333200200333-2021203222022132"></a>

## Next pages — low_security / 312012132222 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1031323012313130-3302133223032022-0131232220111132-1132301232213332-2123021303231101-3023223301323011-0322310002031202-3112231133121213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020322030202313-0220113303001133-2303333022002002-2133313033321213-3200032210303311-3022031112132003-2130123210122202-1232123023033323"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — medium_security / 112233320013 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-3322033110213033-1331332013323311-1123202331232201-0131130122312110-2010113220212010-1020300133222130-0202212202202202-3200323110200320"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-3101212331012110-1131030220301122-0100102121310312-1030032212013113-3113020123321212-2023130030313300-3200120222113121-2220033120312012"></a>

## Direct properties — medium_security / 112233320013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312003030130203-1331100200202131-3122013033233233-1113003220010030-0212331103113021-3303033011313320-2000103012200013-2213233001331311"></a>

## Next pages — medium_security / 112233320013 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](data-sources--workload--reference--group-006.md#canonical-2002302031101322-1232013122201203-3033322021021100-2302303111133202-1123111003311233-1301212232333133-2131333201222030-2333102102333001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333200003200300-3001100322232231-3013223121320232-3211030211213323-2123132111131201-2302320112130203-1120313331222200-2203131131220211"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls — use_mtls / 221012103001 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-2322311011012000-3100321001310010-1110133311132101-3103021033201322-3213323020033011-1222232211302010-1023002132122111-2221310003012301"></a>

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

<a id="canonical-3030130100330100-3010112321012313-3030233021021301-1021033111003232-0321323211001333-1130112113232202-2312103232300333-0122001033320221"></a>

## Direct properties — use_mtls / 221012103001 / 3

<a id="canonical-2231122030300233-0312120123023120-1132331010102011-1301323003223100-2331010330231311-2211130030133200-0102023133223001-1031223333321000"></a>

<a id="canonical-3311232232013130-3103013020003123-1313313022300211-3130321310030321-0100121221301003-3310201303023110-0200122331223023-0300101313101330"></a>

## client_certificate_optional property — use_mtls / 221012103001 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

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

- [crl](data-sources--workload--reference--group-006.md#canonical-2100112021010313-1101003332323120-2002003103202111-3213030123113023-2132000002031233-2200133303323202-3201320031033312-1311300100020023): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-006.md#canonical-0211103030021312-2130222322232112-0231333110332311-1103131311022021-0123233222320003-0221213333113310-0032112233300213-1311213012210002): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-006.md#canonical-0333232222000123-2202303321212203-3102321230233021-3003022333111032-1133210010030020-3023022232331023-3133110220213002-1030111232030201): complete subsection reference.

<a id="canonical-2031133111001003-3233133110103110-3213302022011001-3100322120031313-0320320020101111-1230023210201000-3311211013302002-0210112302030012"></a>

<a id="canonical-3031102312133101-2230113123113333-2310233120333131-1132230302030021-2013202221001101-2211002200323120-2031120221312003-2102322321021222"></a>

## trusted_ca_url property — use_mtls / 221012103001 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [xfcc_disabled](data-sources--workload--reference--group-006.md#canonical-1131301010313020-3212033012100000-1212110033202320-1322331311311021-2212330020102033-3203212210032133-2011233111121011-2112231333320130): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-006.md#canonical-1310032331213233-2001200023031311-1200203131000122-2113331310132110-3221300211221021-0230322330301312-3012310011022100-2320203013330031): complete subsection reference.

<a id="canonical-3103220100123331-3131100103122131-1102133232202033-2223210333003311-3322210002212110-1000232122021123-1132320110012312-0030133322132303"></a>

## Next pages — use_mtls / 221012103001 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](data-sources--workload--reference--group-006.md#canonical-2100112021010313-1101003332323120-2002003103202111-3213030123113023-2132000002031233-2200133303323202-3201320031033312-1311300100020023)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](data-sources--workload--reference--group-006.md#canonical-0211103030021312-2130222322232112-0231333110332311-1103131311022021-0123233222320003-0221213333113310-0032112233300213-1311213012210002)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](data-sources--workload--reference--group-006.md#canonical-0333232222000123-2202303321212203-3102321230233021-3003022333111032-1133210010030020-3023022232331023-3133110220213002-1030111232030201)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--workload--reference--group-006.md#canonical-1131301010313020-3212033012100000-1212110033202320-1322331311311021-2212330020102033-3203212210032133-2011233111121011-2112231333320130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](data-sources--workload--reference--group-006.md#canonical-1310032331213233-2001200023031311-1200203131000122-2113331310132110-3221300211221021-0230322330301312-3012310011022100-2320203013330031)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2100112021010313-1101003332323120-2002003103202111-3213030123113023-2132000002031233-2200133303323202-3201320031033312-1311300100020023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122102320320110-3330032223312120-2201102220013332-0022031320301212-1212333233131113-1133020301001130-1031011332320330-2230313322310332"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl — crl / 012202330231 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-0113020020232111-0323133310230100-1220303232031110-3221121203312212-1033300130110110-3003030322211032-2013012011123313-2221321020211031"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0202103111232333-3323211020221131-3330203222223323-2012201211223130-2233203102013301-0120000321111133-3100231232013011-1100010223032123"></a>

## Direct properties — crl / 012202330231 / 3

<a id="canonical-1200213300033123-3033111223231210-3032010233312323-0301201220112021-1302003003333103-1002332033332203-1233020221233233-3011210312221110"></a>

<a id="canonical-0313000000203101-2111022202231110-2003103213323033-2113322103123031-3302332101201031-3122220032002202-3102120103221002-1103122231030030"></a>

## name property — crl / 012202330231 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2312222023023220-0022131213012220-3232331111131221-3333231222311103-1302013312102120-1220111003133100-1201112020002030-3230230313220303"></a>

<a id="canonical-0130130232102101-2113330213303003-0202123330321030-2123100131103011-0320133112211011-1323011300321031-2113200030013201-2300222332233103"></a>

## namespace property — crl / 012202330231 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2111213223310130-2122322131230033-2122022011130000-1213001200122321-0232221111113032-2202123131013220-1231310222213301-3311301201210211"></a>

<a id="canonical-0302131111132233-2032331313311031-2300311201210220-3231132033120223-3003110302123320-1311202022310133-3030321232030001-0033012121011020"></a>

## tenant property — crl / 012202330231 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1301010321033300-0202211333320000-1332013300032100-2131012330303032-1132321301100201-1313001303110000-3102103221210113-2012003230320323"></a>

## Next pages — crl / 012202330231 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0211103030021312-2130222322232112-0231333110332311-1103131311022021-0123233222320003-0221213333113310-0032112233300213-1311213012210002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131121313010001-2012231333331200-0020303312012333-2031311230103310-1213010133021122-0232310002302323-1030121101031011-3131003220333233"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — no_crl / 223200220323 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-2230300012330303-1013010311013311-0322003221112132-2133210331020202-1331200211312330-2302010212331033-3200210121310113-3221131130020311"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-0200101123210223-1021001112021333-2003122130100012-0331122031023101-0202022000302331-1100313202033211-1013223112320123-0023230230032121"></a>

## Direct properties — no_crl / 223200220323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211023212300322-2220033210330123-3113313320221102-2303321112331322-1230230201313011-2111320210110102-0200202123022321-1031231113010310"></a>

## Next pages — no_crl / 223200220323 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0333232222000123-2202303321212203-3102321230233021-3003022333111032-1133210010030020-3023022232331023-3133110220213002-1030111232030201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301022230223230-2032311312201333-1311121113312300-0331003002202320-0130231012233132-2300032121133232-0310001320110323-3222213002312122"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 011131310222 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1223133010103203-0033201031220210-1110321331032212-3300131112322322-0303030232202130-0102021023211212-0113311132023113-2100231203010322"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2103120003330220-1133131320020030-3221213022233333-3212312131130101-1233110230110220-3113322223102322-1301233113022212-2020103210111221"></a>

## Direct properties — trusted_ca / 011131310222 / 3

<a id="canonical-3032012032233111-2021232100100122-2002000202230020-1310101133010021-2003122331313313-2213231230001013-0311202303332323-3130313211221211"></a>

<a id="canonical-3231312011230111-2130233312320111-2101023020120123-3210330222123202-2212333320211220-1111300133103133-3100230223333320-3211011002123033"></a>

## name property — trusted_ca / 011131310222 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1210131320300101-1032311131030230-2010120011001122-3323010101032123-0133033331321020-0110310203103213-2032110030230232-3021011211001213"></a>

<a id="canonical-1201122212131122-3202113013110123-2320212130233223-0210323121332333-3030230003213213-2213221021033123-1302012202133212-0121033010302133"></a>

## namespace property — trusted_ca / 011131310222 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2020303010200231-1313321013013011-1322303212130003-3021012202213102-2321100013112011-0113112132120332-3200321330111222-2303201113221202"></a>

<a id="canonical-0320211310001303-2012033011013120-0112033010102233-2310120302322131-3032201202300332-0103300121033333-3123010033033000-1003230103010121"></a>

## tenant property — trusted_ca / 011131310222 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0121133311001020-1231102100103030-1230123202120233-1212203303100110-2010100132033310-1033013211332301-3323013132001300-2221010110120001"></a>

## Next pages — trusted_ca / 011131310222 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1131301010313020-3212033012100000-1212110033202320-1322331311311021-2212330020102033-3203212210032133-2011233111121011-2112231333320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231032031231200-0122331203010230-2222332133133122-2301331110211032-1333312330102223-2231211311020113-0011223203003211-3330130212020301"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 202231332100 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-3110333331022300-2132001323301120-3033022021302002-1233101130332030-3032330131330221-2123133231201120-2031021231031210-3001321031112221"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-1220012321203330-1031011332111030-2103030030121122-3110023210232103-0120300013322223-0312010101232002-2233223212000333-3211330303112100"></a>

## Direct properties — xfcc_disabled / 202231332100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011333330031022-1130022012103032-3321200121333123-1301011233311033-1102021011021210-1012212313201102-0112101232131133-1231212110000312"></a>

## Next pages — xfcc_disabled / 202231332100 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1310032331213233-2001200023031311-1200203131000122-2113331310132110-3221300211221021-0230322330301312-3012310011022100-2320203013330031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332233130201320-1023031003031202-2001212313012302-2202032200123223-0003012020323123-1111012213222012-0001211022101000-0333011330101131"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 113131330221 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-005.md#canonical-0102232013120111-2110322312300021-3130310003001200-0200322320112020-1211033133001012-3223011030332110-3303332203233133-1313032331230033)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-1103000002111111-2112222223000311-2023322033002133-3211130003003111-0231230222222200-1223232003012012-2321222123231223-1011120110023010"></a>

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

<a id="canonical-1221200302012031-1213230111102202-2331103312201230-2021111230132221-3302213331230030-1300223211101123-3120022200323021-3110310100302302"></a>

## Direct properties — xfcc_options / 113131330221 / 3

<a id="canonical-0120232220331220-2012112101321011-0222121331322312-1220213013333222-2310320021301332-2232103313123301-1332331212202021-3113022000123222"></a>

<a id="canonical-3230011333332211-2133202131132213-3233032111313013-3232321330331233-2123010311313222-0022110113203301-1333111321323033-0120333311321320"></a>

## xfcc_header_elements property — xfcc_options / 113131330221 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-2023120310133123-1301223210330133-3130200032122230-1012032100201130-3202230021300100-3031222110112323-0223132113011000-1322300323300102"></a>

## Next pages — xfcc_options / 113131330221 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-006.md#canonical-2123132222300110-0312330003033133-1022331021222212-0013221002322220-2220120021013023-0111211022020213-0301301103203320-3003210130312303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023332021001330-1301113201202112-3022332102221231-2313303232121120-0321030201212320-1301220331103103-2133021033212231-2220331033003003"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters — tls_parameters / 132231311331 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-3333202213101201-2120030330220332-1002221033030011-3122110133311203-3130012030200011-0203310301202010-2201031103321311-3313201321002112"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

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

<a id="canonical-1203102320003132-3131103013031331-3111303302102320-3122131110113111-1102003311300322-3020311111120123-0322322323030212-2230333203133201"></a>

## Direct properties — tls_parameters / 132231311331 / 3

- [no_mtls](data-sources--workload--reference--group-006.md#canonical-3131112323220220-0220232022322203-1110321033100313-3120133012000301-0000030032202001-3033202313313222-3101021202233222-0000211002211131): complete subsection reference.

- [tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102): complete subsection reference.

<a id="canonical-3313100030233031-1331212020100231-0201021231010210-3003332112131021-3012212330331033-2203021030033001-2132132002012211-3230021033122020"></a>

## Next pages — tls_parameters / 132231311331 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls](data-sources--workload--reference--group-006.md#canonical-3131112323220220-0220232022322203-1110321033100313-3120133012000301-0000030032202001-3033202313313222-3101021202233222-0000211002211131)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3131112323220220-0220232022322203-1110321033100313-3120133012000301-0000030032202001-3033202313313222-3101021202233222-0000211002211131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210210212001231-1113131001331332-3323223031323201-1301321111122132-0203323310132013-0113210023101301-1321020022303032-0220310302221332"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls — no_mtls / 033213321103 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-3031000010013202-1310310021030222-1001011002001032-3203212303031130-3331202022113032-1102101222002023-2312033031003010-2001202100302031"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2322022320011130-0011000313011231-1330203222023123-0123232200230033-0130311133031102-2132322013300212-0210120302111230-0332021132233112"></a>

## Direct properties — no_mtls / 033213321103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101032303323033-0012332103333001-0311013331031122-1221312333203020-0012132100000313-3010333130021210-3033303131003233-2233200220202103"></a>

## Next pages — no_mtls / 033213321103 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232112331021110-2311121331022230-2112101012123030-1213322131212021-1000330011102231-2121103201013321-3123003333312301-0123033022231033"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates — tls_certificates / 212022321322 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-3223001323321221-1210302101210212-1201300121313233-2333302023233100-1333011113323200-1313033113233022-2222133023133022-3030111102213001"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3333011021313222-3330100320231033-3020231332303312-0001231031201221-0112132321031232-3110103003210030-2322320331133013-1212122200002020"></a>

## Direct properties — tls_certificates / 212022321322 / 3

<a id="canonical-2300012012303123-1212032331302001-1000220320100032-0303000030100120-3323111010322202-3131100331112102-2031022033221013-1013003123033201"></a>

<a id="canonical-3121003331023101-1121110303220210-3110303312131231-2032132320220312-3013313301220102-2311021011332013-1203320013123321-3223021201001120"></a>

## certificate_url property — tls_certificates / 212022321322 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [custom_hash_algorithms](data-sources--workload--reference--group-006.md#canonical-1201321122012001-2021001100221201-2100030330013102-0130223110223131-0311323013221132-2211020200310011-0100111232022220-3320022001000303): complete subsection reference.

<a id="canonical-0122203001022310-2122212230311112-2120032201110320-0213011120230302-1200332212102100-2302221232300221-0202103002301303-0021100130000211"></a>

<a id="canonical-0012321111102001-2312100023103311-1030001303001120-3322222101312320-3230232122320203-2230331011023120-0020323112122001-3023221101232201"></a>

## description_spec property — tls_certificates / 212022321322 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--workload--reference--group-006.md#canonical-0010200132201211-2311212331123201-1220322200210222-3212300122231311-2113331323103313-3320223002102003-1322233110212132-0010200021313311): complete subsection reference.

- [private_key](data-sources--workload--reference--group-006.md#canonical-0112002113222101-0333320311230112-2000002110222203-2133302133002132-0003120301112000-1023331323001120-0320000132230301-0200331320202113): complete subsection reference.

- [use_system_defaults](data-sources--workload--reference--group-006.md#canonical-1310323011103130-0300322132203323-1233131031023013-2002130323022100-2313013311123210-1113332311020102-0132203101013133-3121120333301302): complete subsection reference.

<a id="canonical-0133032320232312-1022122321121000-0200233200013311-2232201012220002-1332021222232201-1301221103020110-0220320033211303-1122330330010301"></a>

## Next pages — tls_certificates / 212022321322 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--workload--reference--group-006.md#canonical-1201321122012001-2021001100221201-2100030330013102-0130223110223131-0311323013221132-2211020200310011-0100111232022220-3320022001000303)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--workload--reference--group-006.md#canonical-0010200132201211-2311212331123201-1220322200210222-3212300122231311-2113331323103313-3320223002102003-1322233110212132-0010200021313311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-006.md#canonical-0112002113222101-0333320311230112-2000002110222203-2133302133002132-0003120301112000-1023331323001120-0320000132230301-0200331320202113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](data-sources--workload--reference--group-006.md#canonical-1310323011103130-0300322132203323-1233131031023013-2002130323022100-2313013311123210-1113332311020102-0132203101013133-3121120333301302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1201321122012001-2021001100221201-2100030330013102-0130223110223131-0311323013221132-2211020200310011-0100111232022220-3320022001000303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223322000321212-2102212202111301-3302110023322022-1030110330212110-1331132000200023-2211201000310131-3302223002222111-0300001300230001"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 133321211110 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-3220211122110311-3133220331111110-1301333031221101-0213013303030010-3023032010311022-3230300312311020-3023020131132030-2313200122013210"></a>

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

<a id="canonical-1320130122123111-0100232231322202-0302110200321011-2130002320211212-3212333212030221-3222233322303111-0312121022212311-0012133312121310"></a>

## Direct properties — custom_hash_algorithms / 133321211110 / 3

<a id="canonical-2012112123001332-2321200301122102-2233213322201330-2131222331321031-1303130233303321-2232303211030330-3221303103221113-0003302100233200"></a>

<a id="canonical-1001011032021320-0020211200120332-3232100101033111-0322331032313122-0200020301302133-1232132212021122-2302001321330211-0332011121101302"></a>

## hash_algorithms property — custom_hash_algorithms / 133321211110 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3220033310121333-2232122113020113-0203102213302331-1103100010232100-2333303011302001-3130032320102211-2000033002122331-3212132332301321"></a>

## Next pages — custom_hash_algorithms / 133321211110 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0010200132201211-2311212331123201-1220322200210222-3212300122231311-2113331323103313-3320223002102003-1322233110212132-0010200021313311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301320222102032-3203230030031122-0003011013002001-0023020300000201-2220201232002301-3301000223213213-1202101201313133-1130121021330332"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 131001120213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3000100200202212-2010120210200220-3001030132203321-0001020001123200-2020132213223201-1113201201233032-3312201032302101-3133010203122003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

Upstream description:

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

<a id="canonical-1133012021131320-0213301110102210-3010120010200121-1003001212132131-2302231322101330-0131321111102211-2131121210213200-1300300313031313"></a>

## Direct properties — disable_ocsp_stapling / 131001120213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222301032100011-0132001233210232-3010001022221022-1123330021033013-2030132331303022-2313310313023101-1333332302332112-3021011110311023"></a>

## Next pages — disable_ocsp_stapling / 131001120213 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0112002113222101-0333320311230112-2000002110222203-2133302133002132-0003120301112000-1023331323001120-0320000132230301-0200331320202113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221223230110013-0122313100103200-1020102031101022-1012232011023220-3020333011210312-0201201103313112-3330301021011200-2122100223033311"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — private_key / 122301223002 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-3312002210133100-2130201110121013-3302100200313302-3332301222323213-2233202121120230-2020100221130211-1210202023101212-2233122331311003"></a>

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

<a id="canonical-1313230030233032-1131031233122133-3210220122100231-3303111331013332-2131032311212301-0231132202012302-2323103102130123-2212103230133020"></a>

## Direct properties — private_key / 122301223002 / 3

- [blindfold_secret_info](data-sources--workload--reference--group-006.md#canonical-2121333231310100-2202101213230202-0123312212133121-0202333321030013-1323321033031033-2321133133031013-2111110101201030-2122110231013230): complete subsection reference.

- [clear_secret_info](data-sources--workload--reference--group-006.md#canonical-1212102301213113-3112232221323032-3101113332130313-2222012203202031-2022120010320100-1120120033332030-1213002211020331-3102311013123133): complete subsection reference.

<a id="canonical-2131131002210200-1302111031312233-1003022103102120-0330222231210312-3313002221322002-1220220030203011-1312100031010021-1110223302300202"></a>

## Next pages — private_key / 122301223002 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--workload--reference--group-006.md#canonical-2121333231310100-2202101213230202-0123312212133121-0202333321030013-1323321033031033-2321133133031013-2111110101201030-2122110231013230)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--workload--reference--group-006.md#canonical-1212102301213113-3112232221323032-3101113332130313-2222012203202031-2022120010320100-1120120033332030-1213002211020331-3102311013123133)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2121333231310100-2202101213230202-0123312212133121-0202333321030013-1323321033031033-2321133133031013-2111110101201030-2122110231013230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112012001203333-0023121323130021-0131133323201003-2323111020001021-3303201231301302-0303113032003120-3332223112213312-2113020312020320"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 131200012311 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-006.md#canonical-0112002113222101-0333320311230112-2000002110222203-2133302133002132-0003120301112000-1023331323001120-0320000132230301-0200331320202113)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3233012311031320-1301010321013302-0122323311033333-0133232012021132-2012210212323332-0211330201310210-1021321022030101-2103301022012130"></a>

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

<a id="canonical-0332202133122010-3312300200033113-1231012133013212-0321033032113210-1022320311222210-1132132131310030-0211313102200100-3202122033333110"></a>

## Direct properties — blindfold_secret_info / 131200012311 / 3

<a id="canonical-3121010021122210-0031131312121003-1232320120310211-3003322212213302-3303222210122322-2121112233132320-2102030000113033-1320232003332133"></a>

<a id="canonical-2101202230333102-2120330020033311-1202213102012000-3011312332303303-1301231132012232-1001112123203202-3010001010333010-1221031211200110"></a>

## decryption_provider property — blindfold_secret_info / 131200012311 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1322112030200101-1023120013000302-3230030222111010-3311310303323310-3130133231232123-3111303313301213-0112010003330300-1213330222203130"></a>

<a id="canonical-1002120023100020-2121203001301010-2302333131021200-0233311232033123-2330103033122101-0102120203312302-0333321203101002-2322213020131002"></a>

## location property — blindfold_secret_info / 131200012311 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2231321131331223-2310123030000323-3011231113233002-3202203320112313-2002331123213301-3111332203220133-2011220132211331-0212132130023212"></a>

<a id="canonical-2131121111110120-0031021310300211-2231101312312310-0112032003200110-3202311200222032-2022033131003311-0122312000203001-2311310013213130"></a>

## store_provider property — blindfold_secret_info / 131200012311 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1110100123231320-3333002013233032-0203330202030113-2103133131331113-0210230323031133-0001102010122220-0110033231132331-2012002022220121"></a>

## Next pages — blindfold_secret_info / 131200012311 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-006.md#canonical-0112002113222101-0333320311230112-2000002110222203-2133302133002132-0003120301112000-1023331323001120-0320000132230301-0200331320202113)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1212102301213113-3112232221323032-3101113332130313-2222012203202031-2022120010320100-1120120033332030-1213002211020331-3102311013123133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313233212322200-3103202221101121-0120213012000031-2112020010021221-2032030310030330-0211133330132211-0012331231112000-3320033102321331"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 222222313210 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-006.md#canonical-0112002113222101-0333320311230112-2000002110222203-2133302133002132-0003120301112000-1023331323001120-0320000132230301-0200331320202113)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-1120020022020120-2003012102121111-1330331111330022-0002302131222030-2002002220200112-2332333133013012-3203013121200032-1222131220230020"></a>

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

<a id="canonical-2002322033220022-3100110310300201-3303323132022202-3012121130312133-1023202222333332-3211111120100302-2211202133322113-3133100300023120"></a>

## Direct properties — clear_secret_info / 222222313210 / 3

<a id="canonical-2003222031310311-1001111002333322-3232013023230113-2130201233330302-1001131212030311-3221210001033310-1100132232330322-2230301031021000"></a>

<a id="canonical-3303303003111102-2123110200223201-2303111102000233-3021200130130232-0033212030102103-3201310111132213-0002000123031130-3300222021031333"></a>

## provider_ref property — clear_secret_info / 222222313210 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0121223202101332-3201021220013312-3221022112321333-1110231200033023-2202312022133220-2221202210201133-2223131213022013-1200320222123320"></a>

<a id="canonical-1130021212321003-2312033332202020-1131233221202320-2023101310003013-0301120011233133-2220030333000111-3233121300001313-2300301003202033"></a>

## URL property — clear_secret_info / 222222313210 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0002001313230321-1000130330323023-3222200032130010-1323033101301133-3100111103122212-3331110111222203-2111323123032023-3023302021001101"></a>

## Next pages — clear_secret_info / 222222313210 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-006.md#canonical-0112002113222101-0333320311230112-2000002110222203-2133302133002132-0003120301112000-1023331323001120-0320000132230301-0200331320202113)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1310323011103130-0300322132203323-1233131031023013-2002130323022100-2313013311123210-1113332311020102-0132203101013133-3121120333301302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133210113323333-3033321133132331-3311202021303312-2133230000310021-1022233022133231-1301123112122021-3130210311201200-2120210121210032"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 212331212032 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-0323111013033110-2031211121210030-2101311220010202-0310323211113101-1020203300200302-1003322210303021-0132312032230211-1332101023333033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

Upstream description:

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

<a id="canonical-3203021220201111-3102303033333020-1131133020332201-1300031013023133-2213300121221003-3010121021332213-0103031000331112-1213033031302222"></a>

## Direct properties — use_system_defaults / 212331212032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323221323013101-0132001122201313-1121220123222112-2011112100002011-0321021303122303-0221001200002203-3311233021333103-1001003210112123"></a>

## Next pages — use_system_defaults / 212331212032 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-006.md#canonical-2301221203132223-3031012223223001-2310313322201310-2332212103333322-3223201002103202-2320023103101132-2210221222013313-2322312223103130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032332133101301-0020320112213300-2112222003313222-1132211133311322-0200122210110202-3320120301013100-2203002210233132-3122300222320132"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config — tls_config / 322012122123 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-3331331200113210-2332312231210213-1100113302011121-1102000322301213-1220121123100311-2122223202100202-2233232233220110-3301322021202210"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

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

<a id="canonical-1103110211031120-0102330232202032-3020111030321300-0201033113131213-3310211300120003-2322133001102332-3233202230212020-1233112230032112"></a>

## Direct properties — tls_config / 322012122123 / 3

- [custom_security](data-sources--workload--reference--group-006.md#canonical-1233211231312010-2021020331101123-0122211212012011-1132032102220323-3010013210100112-1320112222033332-2320023121313322-0323233333311030): complete subsection reference.

- [default_security](data-sources--workload--reference--group-006.md#canonical-1021030310301020-0111113212130332-0211021220012000-3230102012023112-3111100233321202-3221233303032023-2232211220113012-0303103031231110): complete subsection reference.

- [low_security](data-sources--workload--reference--group-006.md#canonical-2213232010122122-2322113000101322-2110312223212010-3020211330323212-3313103201202210-2111031313322102-3110011031122003-0103333223301200): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-006.md#canonical-2023210311000132-0203302312323303-0033103213323013-1111123223331321-1201002220013321-1130102233302231-0231023313110301-0130331123031202): complete subsection reference.

<a id="canonical-2210212313220230-3011030333320100-0121132001103202-1133310311012103-1002030320003111-2201330313223121-0231212122222202-0332303111213110"></a>

## Next pages — tls_config / 322012122123 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security](data-sources--workload--reference--group-006.md#canonical-1233211231312010-2021020331101123-0122211212012011-1132032102220323-3010013210100112-1320112222033332-2320023121313322-0323233333311030)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security](data-sources--workload--reference--group-006.md#canonical-1021030310301020-0111113212130332-0211021220012000-3230102012023112-3111100233321202-3221233303032023-2232211220113012-0303103031231110)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security](data-sources--workload--reference--group-006.md#canonical-2213232010122122-2322113000101322-2110312223212010-3020211330323212-3313103201202210-2111031313322102-3110011031122003-0103333223301200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security](data-sources--workload--reference--group-006.md#canonical-2023210311000132-0203302312323303-0033103213323013-1111123223331321-1201002220013321-1130102233302231-0231023313110301-0130331123031202)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1233211231312010-2021020331101123-0122211212012011-1132032102220323-3010013210100112-1320112222033332-2320023121313322-0323233333311030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101122321233203-2000300303013101-3320013033300303-1100022211213300-3112231211102023-2111123231230310-0221130211112232-1301012200221012"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security — custom_security / 203220311200 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-0233332331101022-3032120031023133-2001030321110001-0231131001110121-0023120311013011-3100132100321133-1021203321321111-0230012033002300"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

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

<a id="canonical-3210213032213303-3221112322103332-3312303031103031-2330100230122330-2200202112213013-1133120111012303-2201010213111113-2031130210311033"></a>

## Direct properties — custom_security / 203220311200 / 3

<a id="canonical-0331031132331110-3102233020223132-1033120020130131-0313011321112110-2320231330200101-2311311123121012-2230202203101022-1323000230111100"></a>

<a id="canonical-0323122310112300-3111200300231123-3203213312133330-3022000233321011-1202333200100201-3232131231032102-0111100200101201-2231101203101021"></a>

## cipher_suites property — custom_security / 203220311200 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1321203031222102-2220333310011313-1111210211133313-0131322302131021-2031011020300130-3130030022033331-1220013133011113-3122133233301001"></a>

<a id="canonical-1022201312122312-2322220200303303-2110011132330301-2122321233302202-3131133123103222-3131103301002223-1203003320203020-2322020232133011"></a>

## max_version property — custom_security / 203220311200 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-0312023223223003-0331000213110003-1020301300231232-0030212311020301-3123130303221331-0132001312320011-3201322030033233-1222012202000321"></a>

<a id="canonical-1320320202230032-0010200300320013-0200013033302231-3022210102123313-2100231333103232-1222110101120022-1333230321003111-2212330113123102"></a>

## min_version property — custom_security / 203220311200 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-0212110312313300-0312323023032123-0313200000000133-0121113031202020-2313031120100220-2103331022130010-0020132000023322-3203221233203321"></a>

## Next pages — custom_security / 203220311200 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1021030310301020-0111113212130332-0211021220012000-3230102012023112-3111100233321202-3221233303032023-2232211220113012-0303103031231110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220003213001231-3330003203230001-0003133111321323-2031212100032112-2213320311320232-2132032101002301-3101332131123331-0200202010220121"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security — default_security / 300301013301 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-3110332212130123-1230020231320233-1133132032123331-0212331031030132-2123302223233303-1011333120222110-0110023212103102-2112022012313232"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-1200121201312310-0310001222321321-0030322210213301-2322032022112311-0300112032232310-1033231223322111-2223201201023202-1130322221002310"></a>

## Direct properties — default_security / 300301013301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323212102223333-3210213200002221-0320321103123313-3130032322313310-3223203221000102-3211333311232201-1113300033313332-2313120123302103"></a>

## Next pages — default_security / 300301013301 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2213232010122122-2322113000101322-2110312223212010-3020211330323212-3313103201202210-2111031313322102-3110011031122003-0103333223301200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001013313312332-0013231310300113-0110021022200302-1212310300300130-3211213113003330-1010213110300010-1231303221201223-3033110310312311"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security — low_security / 331000013202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-3223232123030130-1011203330131131-2311300332200031-0032333121133323-3331103131210013-3000110301203222-0110221112100000-2102131012013030"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2000021330333231-2233021223101323-3133220103332311-1302222323212202-0220323200311322-1012320323022102-0321313221301330-2012100122210021"></a>

## Direct properties — low_security / 331000013202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102000103002130-3113301310331221-1021132133330320-3111133101232322-3321221210013200-2222331211301133-0003320211331222-1012332201131322"></a>

## Next pages — low_security / 331000013202 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2023210311000132-0203302312323303-0033103213323013-1111123223331321-1201002220013321-1130102233302231-0231023313110301-0130331123031202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303221033312103-0320311330032231-2322313332002232-1110200112100323-3023120111323021-3222320131203001-3133323330103113-1230230133011021"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security — medium_security / 220131303210 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-1202301331131320-3311110030210230-3022332221323203-3231021122010321-1311312113122121-0313301001211111-3233103112122333-0312132010130231"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2031210032030220-3101030313200002-2230110220133001-0233321201312001-0333300203022011-0133313232210000-3222002300311121-2100332033320200"></a>

## Direct properties — medium_security / 220131303210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203001022023231-2301131120303120-0102221323231213-1110322100323012-3122323233230220-0321310110300302-2103103113012232-2230310233230020"></a>

## Next pages — medium_security / 220131303210 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-006.md#canonical-2213322032013202-3112133000020231-3020123023031101-3331202310121310-2203322130230101-3201331000031110-1120030000210213-0001303201121110)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323301221201030-0010303110200120-0022210320130010-1113221221010233-0231130332101123-3200020021001222-3002333020132220-3300222201320132"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls — use_mtls / 220301320030 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-1000232002133033-0032022332110212-3012121333101300-2231332121203233-2213310020010300-0000322020111321-2210110213222133-0102233012020030"></a>

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

<a id="canonical-2220022133122023-1311121101130301-2121103021313132-0003112012001223-1201320130313122-2321003210220331-3222131120003010-3303222132233223"></a>

## Direct properties — use_mtls / 220301320030 / 3

<a id="canonical-3330113200031031-3011210131233302-2202003032110002-3323330023022220-0102332111033321-2211203101232000-0131201211030033-1012332301133020"></a>

<a id="canonical-3000131020130123-0302300202002130-0322031212013222-3003233203012330-3331200120010100-0231212203311123-2003210231220232-0222120222023332"></a>

## client_certificate_optional property — use_mtls / 220301320030 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

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

- [crl](data-sources--workload--reference--group-006.md#canonical-1332100112202202-2011021330011322-3023133201101030-2323233132233002-3221213121123003-1120301313032121-0130131330220203-1231100000310101): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-006.md#canonical-2111030321113021-1003300000221331-3201230022120022-0131021210322123-0210333121221310-3023232312333102-2313131303210330-2333303201202222): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-006.md#canonical-2211100001003313-1000000131033111-2232310310032121-2301221130001132-1103011320133122-1310020231221213-2101020133103200-3312120010113331): complete subsection reference.

<a id="canonical-0102131113130113-3311111100212022-3002212133021311-2021100331233110-3303331222302332-2221102212103002-3321003221313123-2023202100210332"></a>

<a id="canonical-0130301030003331-2322203310011300-0031111220221311-3120310300100123-2300333312030221-3213121213023323-3330000103212313-1310120012333233"></a>

## trusted_ca_url property — use_mtls / 220301320030 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [xfcc_disabled](data-sources--workload--reference--group-006.md#canonical-0023021220022003-3131020323102102-3112121300133102-1023332113131301-0101102132032032-3112321203012311-3112312013132000-1221330032222312): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-006.md#canonical-2332032203103031-1210222110013121-1232330112000002-1213101303310103-1233133133010302-2131132110203032-2010311300113111-1113002221101333): complete subsection reference.

<a id="canonical-1323123003101100-0201331132030022-1003133213031030-2113021131221330-1323023021222101-0000020331320231-3213303023322120-2103333203013213"></a>

## Next pages — use_mtls / 220301320030 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](data-sources--workload--reference--group-006.md#canonical-1332100112202202-2011021330011322-3023133201101030-2323233132233002-3221213121123003-1120301313032121-0130131330220203-1231100000310101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](data-sources--workload--reference--group-006.md#canonical-2111030321113021-1003300000221331-3201230022120022-0131021210322123-0210333121221310-3023232312333102-2313131303210330-2333303201202222)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](data-sources--workload--reference--group-006.md#canonical-2211100001003313-1000000131033111-2232310310032121-2301221130001132-1103011320133122-1310020231221213-2101020133103200-3312120010113331)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--workload--reference--group-006.md#canonical-0023021220022003-3131020323102102-3112121300133102-1023332113131301-0101102132032032-3112321203012311-3112312013132000-1221330032222312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](data-sources--workload--reference--group-006.md#canonical-2332032203103031-1210222110013121-1232330112000002-1213101303310103-1233133133010302-2131132110203032-2010311300113111-1113002221101333)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1332100112202202-2011021330011322-3023133201101030-2323233132233002-3221213121123003-1120301313032121-0130131330220203-1231100000310101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222112033310022-1122030301230330-0003033103100000-1203212111012333-0101330003012212-3323202311320133-0011102023210212-2013300300201202"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — crl / 000021033320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-1312132000300310-1131312223322312-2111032233220232-2232103212322213-2230230121301012-2212210021112311-3101031033120221-1000311102131112"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0323020210310302-2232202332032212-0300220002222230-1030231211201233-2000012323001311-0332322003201123-2321222231111322-0001120222003020"></a>

## Direct properties — crl / 000021033320 / 3

<a id="canonical-3132032210111301-2023023322021110-1211331200323032-0100120233010330-0201131313201012-3201213011303200-1302003121303222-1033313111212333"></a>

<a id="canonical-1123302321103222-2202222232001001-3320302020302030-2010001233020231-1223213110301123-2223313123223230-2013002112232031-2123223130311300"></a>

## name property — crl / 000021033320 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0132111212110330-2332231320130013-0130111223120103-3213233020303223-0120100122313220-1022332230000333-3003332211321203-2201013230320223"></a>

<a id="canonical-3020011212000222-3033300320312110-3213101330223030-3122320311211020-1201231321030000-1212131123003022-0003110300111200-1321031323102321"></a>

## namespace property — crl / 000021033320 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1231032003322121-1223211003102111-1323011313222111-3233102320331121-0111331001112300-3303320001132012-0220200021233311-2130331200103110"></a>

<a id="canonical-2111302110220202-1023222112100121-1111123120031333-2213212203302021-2310000202112202-2130121231202121-0223003001222222-1000100032220220"></a>

## tenant property — crl / 000021033320 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1220310132021312-3132003232021201-3003103300231302-0310303231303132-2230322123303213-2103202032130120-0221010012331223-3121001010103123"></a>

## Next pages — crl / 000021033320 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2111030321113021-1003300000221331-3201230022120022-0131021210322123-0210333121221310-3023232312333102-2313131303210330-2333303201202222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112033121011212-2022201032231031-0023331022222301-2310031133223003-0222332010232121-3031323032212202-0123303202210301-3031220113301133"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — no_crl / 123300210100 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-1100323100133312-3220010023001113-3302033313021320-3332200220221030-3332321111201123-3212223022223010-2113313332001320-2330020031230312"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-1212121020120331-0221203330030123-0200122001332112-1103212300012102-3112112131111110-0100321312201122-1103123301313211-0231221002122333"></a>

## Direct properties — no_crl / 123300210100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010101321020313-1233230111220113-2003033331202331-2221133110312233-3022210031312031-0230012012330322-3222123310300200-2113211212203301"></a>

## Next pages — no_crl / 123300210100 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2211100001003313-1000000131033111-2232310310032121-2301221130001132-1103011320133122-1310020231221213-2101020133103200-3312120010113331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331330020331110-0133210021310200-1000130223123321-1221010033211202-3203130330202300-1021313120130322-0033211332230012-0012312330000322"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 030132012010 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1012122003023302-1203000110001031-3122301211110322-1020203010113311-0222201232001303-3230110131221011-1123222113110121-2023323313031210"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1302323232332001-2002130131222100-2103000212120322-1030001131231301-2112132132223313-2031023100321120-0202221211032232-0201133122001133"></a>

## Direct properties — trusted_ca / 030132012010 / 3

<a id="canonical-1200003133310013-1110220213301201-0101131120111232-0020302121222130-1331001300022131-3313302130133323-0032130101022001-2013112102101011"></a>

<a id="canonical-2222022331331023-0233003302033333-3300300012211321-0301202122022002-1331002131212211-3322021010211312-2233033111002100-0023001323301031"></a>

## name property — trusted_ca / 030132012010 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1130132311322301-1011222303231313-0211223300103201-1212310133121022-1310022232111003-1212030201131030-2333000312313003-2130301230020011"></a>

<a id="canonical-0030212110123122-1213023123333213-3100002002231020-3330012320330103-1000132212333321-0322220300022031-2031311020011031-3121122011330122"></a>

## namespace property — trusted_ca / 030132012010 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1123120133110112-1223322330203331-1002122022131010-3312311132110012-2322031233220010-0111322303102000-3120211221223331-0132311120023300"></a>

<a id="canonical-0330112330310231-2013001020000010-3233122003220131-2220032310101123-0003122333200111-2333201021201103-0313333203311213-2131003223213122"></a>

## tenant property — trusted_ca / 030132012010 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0323322003112212-1003000112031002-2013102321022230-3033113113213030-3002132030211211-0230020100322221-2112003320113030-1220321311123300"></a>

## Next pages — trusted_ca / 030132012010 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0023021220022003-3131020323102102-3112121300133102-1023332113131301-0101102132032032-3112321203012311-3112312013132000-1221330032222312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301310303111102-3123332022033232-3203221120103202-2330313121331221-1303200103212331-1212333010212033-1211021122020120-0023312210220331"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 001100330120 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-1212103200112221-3020132301202020-0102333111013311-2330221003112022-0202300130213010-3010103311001211-0011202312311111-0302301001223003"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-3003031300001101-0231302313133220-3130110101021100-0130020001121322-3331213231002123-1310013111002132-1302220012132213-0110320233322022"></a>

## Direct properties — xfcc_disabled / 001100330120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313323201313133-3333000212212001-0113030201230000-0212200103213233-1030030332120320-1303012303011000-1221032221210120-0011211322210132"></a>

## Next pages — xfcc_disabled / 001100330120 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2332032203103031-1210222110013121-1232330112000002-1213101303310103-1233133133010302-2131132110203032-2010311300113111-1113002221101333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010223332310133-2133021300132201-1301102233032213-3330013001113010-1013222020200121-3202003032221303-3111113223103300-1212201013002102"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 333211203111 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-005.md#canonical-3101230020212103-1000033022231131-3311300103311313-3131102020122310-2330011021322323-1001031320003132-2320130302320302-1103203103330101)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-006.md#canonical-3230312110002321-0231002302222111-1313023023222203-3020111230020103-3212002301303131-1013301332001321-2233020331123001-0102332102102232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-0212220210020302-1330202132101020-0331233011311200-1031000021333330-2220132130130100-1032303001213311-3133032320021101-1132111203212130"></a>

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

<a id="canonical-0101332222312123-2130113031320000-0021312021310200-3013323200023021-3132013310030312-2011231232122011-0320223301220211-2022313333013201"></a>

## Direct properties — xfcc_options / 333211203111 / 3

<a id="canonical-2021120030001003-3112010130121203-0312032133331103-2022331031330320-0003130330302111-0220102313230203-0121113123210322-0232213002101212"></a>

<a id="canonical-3110012021033323-1003323303330311-3211021211033001-0320302030000330-1022123112320200-1020102112303212-2211310132032233-3130212303212102"></a>

## xfcc_header_elements property — xfcc_options / 333211203111 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-3213021210222230-1011333201310233-1032011130313031-0323103122233012-1203321331230333-3000320131323322-0311301201120111-2213032102202121"></a>

## Next pages — xfcc_options / 333211203111 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-006.md#canonical-2322213330230332-0230331031131101-0321222230112320-2201310023230131-0213201223011000-1300311222103311-3210222221112123-0231232010320102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0212031122322200-2022231222300011-1333320232232011-0120223221003012-2010211322132130-1002332132031210-2031010103313020-3320013032133330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330011012022230-0230102102000123-2131101132000022-2313101101331020-2312210330020301-3012032203103202-1101330322312232-3021312120022301"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert — https_auto_cert / 023302212101 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [service](data-sources--workload--reference--group-005.md#canonical-1202331303220203-0201220210302110-3230233130203212-0211012200301123-2132232223102021-1100310220330100-0330302132103311-3331131002000211)
- [service.advertise_options](data-sources--workload--reference--group-005.md#canonical-1230023000110022-1010223003013010-1333020121230231-2221232101221211-2130003313320013-0210002201312132-3102123003022022-3202000301213333)
- [service.advertise_options.advertise_custom](data-sources--workload--reference--group-005.md#canonical-3332031113102310-3302223302333012-3120332000000030-2032331232202322-2211233120201013-3311001333011020-2221212213301221-1000320213301202)
- [service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-005.md#canonical-3102330212100003-1011312331332022-1230220200231323-0322023032231230-0301000013321010-2210330211111012-2002131311203010-2202033133210102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-005.md#canonical-1322323213310333-0001232011202131-1303231200230223-2123220213203321-3133010110001130-3103331221032223-1003000023032221-3033233301202130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert

<a id="canonical-0202020012201012-2002032101110031-2212023300112033-2330231221200330-3031321223223123-2203303032111012-3203032210333223-3133121102000333"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

<a id="canonical-1102101123321100-0213101222301032-0302310100031203-0002331231131011-3220232031211310-3023113322231002-3323202332320121-0031033230032333"></a>

## Direct properties — https_auto_cert / 023302212101 / 3

<a id="canonical-3121023230010300-0102012210121031-0232001221321133-3021222002020220-2321130222211232-1003012131333131-1120110211011020-2120232312011303"></a>
