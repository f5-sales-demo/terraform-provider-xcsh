---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0303232100133222-2130330313201033-2201223103031001-0233203200301001-2312022223221013-1022131013303011-0023302033112233-2131022001020203"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 012303211313 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-3230312021230322-3310001132303313-1110012112202213-3330322321030323-1121221221133120-2321102112023333-2011212100221003-0220120111022302)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-3010030130130213-0233311203221211-0302301012133311-0223303023230132-3333021000312203-2031010110200121-0030233112301010-2010220232133001)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-3322300310001300-1103102103311122-2211132002003011-2311211201021303-1102033120011021-1010000301013030-3032231123011021-1132310111232203"></a>

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

<a id="canonical-0331321023313122-2302210020330221-2212300111210331-2310333232132023-3330121220231012-1002133012311313-0111013223300121-2200110000101322"></a>

## Direct properties — trusted_ca / 012303211313 / 3

<a id="canonical-0133000210001132-2033230203210323-0003011301010002-1130323330332223-0212020320331213-1121312033113200-1121213033132000-2101220312102213"></a>

<a id="canonical-0032021200333023-2223121310233130-2231032321000330-0112222112033232-2213132011011210-2130302120310310-0323121233233310-3111200130330222"></a>

## name property — trusted_ca / 012303211313 / 4

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

<a id="canonical-3303303123110322-0121121220033113-0231020101222113-3233230103011231-0111113311232203-0032331122030320-1012302133320331-0332033122333323"></a>

<a id="canonical-1200002123223313-3102023133020000-0201132200331301-1301002203031211-1000213210023331-0200021212013133-0203210011021202-0113102103130232"></a>

## namespace property — trusted_ca / 012303211313 / 5

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

<a id="canonical-0131333031111012-0022023102302100-2323101112233100-2000110133132013-2013001031112320-3002021032230233-1133210010213331-3100332210130122"></a>

<a id="canonical-2013213311303203-0123013112211022-1331101010001132-3213331033110022-0233211021301201-1230312022112223-3201033130210013-3202200113221331"></a>

## tenant property — trusted_ca / 012303211313 / 6

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

<a id="canonical-3023303223110130-1331233131103010-0132200133300312-2221033002312010-2300120102031211-0133323121011202-3300103011230332-2323213102032320"></a>

## Next pages — trusted_ca / 012303211313 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-3010030130130213-0233311203221211-0302301012133311-0223303023230132-3333021000312203-2031010110200121-0030233112301010-2010220232133001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1012001212200320-0310100102011003-3222132130103211-3211031230332103-3312022033113131-0013221102222302-0212302021221313-2331110022001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212012113313232-2133300221122122-2221230202210300-3220003100200213-1330111112131001-1002332132131332-2223131111231312-3012300023113130"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 012320313200 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-3230312021230322-3310001132303313-1110012112202213-3330322321030323-1121221221133120-2321102112023333-2011212100221003-0220120111022302)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-3010030130130213-0233311203221211-0302301012133311-0223303023230132-3333021000312203-2031010110200121-0030233112301010-2010220232133001)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-0100202313230010-0312222122311010-0133211113231113-1330222211230210-0232331001001032-2312332110210311-2112013321231223-0332230101011113"></a>

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

<a id="canonical-0010222201331131-2323002033203123-0000112231121003-2230211000312020-0200212031303330-1233203102321002-3020132231233102-2122030032222003"></a>

## Direct properties — xfcc_disabled / 012320313200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322213211100030-2311203332310321-2331132332101303-2113320130223033-3122032220302113-0132323320212131-2000332301032302-2213211111223031"></a>

## Next pages — xfcc_disabled / 012320313200 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-3010030130130213-0233311203221211-0302301012133311-0223303023230132-3333021000312203-2031010110200121-0030233112301010-2010220232133001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0232212102233202-2232200221001030-0101120013011032-1111211121231121-3313332102322123-1233301110312210-2131103100313333-1213203021113001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133032020002021-1213122300331232-2113333130001303-0323221111330030-0213120301331021-1102101020233122-3130211122223031-1003212300302201"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 021211322212 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-3230312021230322-3310001132303313-1110012112202213-3330322321030323-1121221221133120-2321102112023333-2011212100221003-0220120111022302)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-3010030130130213-0233311203221211-0302301012133311-0223303023230132-3333021000312203-2031010110200121-0030233112301010-2010220232133001)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2223223000231302-0133110103223000-3322022210110003-0233032133003322-2020011211121002-0201331030221120-1030112201210210-0122002312332003"></a>

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

<a id="canonical-2110322331333012-3201300122031301-0313301103113131-0120023021022012-1222201002301220-3311201233102313-2302033110331221-3130100311210120"></a>

## Direct properties — xfcc_options / 021211322212 / 3

<a id="canonical-0231211010213332-3033120220122311-3232111022211001-3211301302102222-1220312213330210-0123001102311210-0302301210203231-3011033112300033"></a>

<a id="canonical-0320300022210222-3111002333320023-3213323233102130-3310103302200211-1311210323201030-3220013031101013-1222100101202102-3121302233302333"></a>

## xfcc_header_elements property — xfcc_options / 021211322212 / 4

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

<a id="canonical-1232003132333000-1211310101313233-1002321223002210-1301201202331112-3023001121103202-0230001033000223-0233000103123022-2323113033330322"></a>

## Next pages — xfcc_options / 021211322212 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-018.md#canonical-3010030130130213-0233311203221211-0302301012133311-0223303023230132-3333021000312203-2031010110200121-0030233112301010-2010220232133001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233203200112031-3111032002111012-2100321311232101-2131332211313321-3112113201333011-1031233100233331-0221101013332000-2033111212032012"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert — https_auto_cert / 332012231131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert

<a id="canonical-1323011022132303-0113231111013132-2011301211203032-3101101002101320-0011223212131220-0310100030200201-0303202320101031-2030222220012302"></a>

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

<a id="canonical-1203200303211132-0213121211133022-0200020100013321-1111321032110000-1232111011330322-2013021203022323-3121032112002203-2232300110120010"></a>

## Direct properties — https_auto_cert / 332012231131 / 3

<a id="canonical-0002312331121002-0012122022323123-2213133032301331-3110031133222323-3131012031101100-0113113010021202-0302020230031200-3022323310231011"></a>

<a id="canonical-3131121111021010-2220020023120321-3221130012211333-2303003012101131-2002010223232213-3113202022102310-3031333333233303-0022121021213111"></a>

## add_hsts property — https_auto_cert / 332012231131 / 4

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

<a id="canonical-2003111003330202-3303201320122113-2323200301220010-2020031111122032-0002302311213323-3233212122201323-1312300222021132-3221313311330212"></a>

<a id="canonical-0100320030030222-2221103001303230-0111123233000320-3110231321320013-2010202220113201-2033011003101012-3132001023210010-1310211212311001"></a>

## append_server_name property — https_auto_cert / 332012231131 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-019.md#canonical-0322012333022122-0100132312200311-0311303222130121-3022223011010322-1103030132323301-2312013011122212-1111102310032131-0123132120313301): complete subsection reference.

<a id="canonical-3021111121102012-3212002311010020-3111321230311111-2200000103003303-0200201332213203-1223313220003122-3120011221011010-1203303003223332"></a>

<a id="canonical-2032321102033003-2003313221013221-3211222333310011-3223013203320212-0300010302111121-1010030033131130-0121302123031230-0102323233120220"></a>

## connection_idle_timeout property — https_auto_cert / 332012231131 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-019.md#canonical-2200130132022231-2002123300110310-2003330320002013-3233302321020131-3113330022110201-2300103331101032-1100003102020213-3003222221130110): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-019.md#canonical-1031201332122110-1012222320310122-1210232321233211-1100211003330002-1103312323010320-0212232302231100-3112200331231011-3111321331302321): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-019.md#canonical-0333232132032030-3133133333130221-3333130121003233-3313013310132111-1300301101032101-1222111230011103-0200313120010222-2000020103232000): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-019.md#canonical-3310031330333302-2300003212023313-3002323132300211-3202111001300101-2230233032203110-2313212031202102-1020223321123200-2303131333002003): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213): complete subsection reference.

<a id="canonical-3201230002020231-1311020331233103-2031123302013221-1321011002333221-3232330322111122-3300221213002202-2233330021232332-1202312301211220"></a>

<a id="canonical-2313032323121220-1222101000000313-0321123211020000-2230200322010333-2310131212122223-1031220010331213-3103232220201003-0220313210130331"></a>

## http_redirect property — https_auto_cert / 332012231131 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_mtls](data-sources--workload--reference--group-019.md#canonical-1301101230033112-3020330221011121-0310230213123321-3122232303313021-0313303222103221-2131330331331202-2220320110030033-3130123001220231): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-019.md#canonical-0123212111101211-3221323200311300-3220232332021201-2002310222233130-3321330132002101-0303203113020332-1203102002030031-1011021221210230): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-019.md#canonical-1101311113121322-3100323133230131-0321033001022313-1113030012312103-2110011131030201-3013233031033002-1300302030220331-1323302021203003): complete subsection reference.

<a id="canonical-2032321101133023-0103203130112302-1322032121012331-1312010122201030-2211321113210221-0201002313310221-2312012133210301-2230123031201232"></a>

<a id="canonical-1332200133303132-3300022220212031-1111111332030233-1321102101032022-1123203213331013-1332210323022102-1111203211101210-1313302231131033"></a>

## port property — https_auto_cert / 332012231131 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0312232023122312-2320010322102302-1211013202213031-0213113213220320-0023233010330203-2221010233331012-3003302210311300-2113123211110212"></a>

<a id="canonical-2210202021312010-2010220303032103-2110100020133013-1213011121332013-1202132012131020-0232203333112020-0012133330112331-2101210332013120"></a>

## port_ranges property — https_auto_cert / 332012231131 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-2311333130100300-2210113013300200-0011213001013310-2332100313032330-1210320103220122-0021032033330131-3032113021320320-3201230210123231"></a>

<a id="canonical-2203123031133320-2200003202212311-2302331110200302-1013200310212023-3022031123103030-2013133300300210-3013113130021303-3332223130213230"></a>

## server_name property — https_auto_cert / 332012231131 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310): complete subsection reference.

<a id="canonical-0202123330012013-1011333233102223-0033022010321033-2213312000330120-0000210111233101-3322122311232010-1000130232011121-3101231101130300"></a>

## Next pages — https_auto_cert / 332012231131 / 11

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-019.md#canonical-0322012333022122-0100132312200311-0311303222130121-3022223011010322-1103030132323301-2312013011122212-1111102310032131-0123132120313301)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header](data-sources--workload--reference--group-019.md#canonical-2200130132022231-2002123300110310-2003330320002013-3233302321020131-3113330022110201-2300103331101032-1100003102020213-3003222221130110)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](data-sources--workload--reference--group-019.md#canonical-1031201332122110-1012222320310122-1210232321233211-1100211003330002-1103312323010320-0212232302231100-3112200331231011-3111321331302321)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](data-sources--workload--reference--group-019.md#canonical-0333232132032030-3133133333130221-3333130121003233-3313013310132111-1300301101032101-1222111230011103-0200313120010222-2000020103232000)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](data-sources--workload--reference--group-019.md#canonical-3310031330333302-2300003212023313-3002323132300211-3202111001300101-2230233032203110-2313212031202102-1020223321123200-2303131333002003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls](data-sources--workload--reference--group-019.md#canonical-1301101230033112-3020330221011121-0310230213123321-3122232303313021-0313303222103221-2131330331331202-2220320110030033-3130123001220231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](data-sources--workload--reference--group-019.md#canonical-0123212111101211-3221323200311300-3220232332021201-2002310222233130-3321330132002101-0303203113020332-1203102002030031-1011021221210230)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through](data-sources--workload--reference--group-019.md#canonical-1101311113121322-3100323133230131-0321033001022313-1113030012312103-2110011131030201-3013233031033002-1300302030220331-1323302021203003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0322012333022122-0100132312200311-0311303222130121-3022223011010322-1103030132323301-2312013011122212-1111102310032131-0123132120313301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313113132231130-2103121323233121-2311322333220203-2202133333302313-3122221103333102-0122310023011310-0000210102133113-2333002310000013"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options — coalescing_options / 220333131201 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-2002310201230231-1120222212313231-2133321130131301-2231100102332101-3312330223232203-2302132100331012-1012110332013301-3113101122110323"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-0101013312230200-0123011203021011-1031002222100321-1131323003330320-3331302131112220-2102122123322011-2000200032121031-1311111320020211"></a>

## Direct properties — coalescing_options / 220333131201 / 3

- [default_coalescing](data-sources--workload--reference--group-019.md#canonical-3203202101030032-1201120220211003-2103223120030003-3201110131133113-1210233002033101-2013122112111331-0002103130202113-1210010023200130): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-019.md#canonical-2300030120121300-3211123232301212-3310031313331032-0210001203202122-2231000320110011-0103003230002002-0233123333032321-2320121022103310): complete subsection reference.

<a id="canonical-3023112323021313-0222031302333230-0102201112212001-2232121311103313-0200203300110133-1020301223322203-2022232312302300-0203023021311321"></a>

## Next pages — coalescing_options / 220333131201 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](data-sources--workload--reference--group-019.md#canonical-3203202101030032-1201120220211003-2103223120030003-3201110131133113-1210233002033101-2013122112111331-0002103130202113-1210010023200130)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](data-sources--workload--reference--group-019.md#canonical-2300030120121300-3211123232301212-3310031313331032-0210001203202122-2231000320110011-0103003230002002-0233123333032321-2320121022103310)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3203202101030032-1201120220211003-2103223120030003-3201110131133113-1210233002033101-2013122112111331-0002103130202113-1210010023200130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213020310300300-3223103321101301-1031313031022203-3102030323301321-3311103013100013-3120111020223111-2120233022203330-3133113203021133"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 121033012322 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-019.md#canonical-0322012333022122-0100132312200311-0311303222130121-3022223011010322-1103030132323301-2312013011122212-1111102310032131-0123132120313301)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-3003303300012013-3000010131311321-2213023032101122-2220001331323203-2331102221201112-0003110200200313-3100330212002122-3231223201322101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-2113100230123300-3333023013303231-1301220313011302-0213023221101233-1321201303322320-0011210211013120-2001100313213001-2022110000021011"></a>

## Direct properties — default_coalescing / 121033012322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001130323011002-0220020200230232-1203011302333210-1011321232110302-3001031221103313-0320323233222112-0203011302021322-2133023131323101"></a>

## Next pages — default_coalescing / 121033012322 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-019.md#canonical-0322012333022122-0100132312200311-0311303222130121-3022223011010322-1103030132323301-2312013011122212-1111102310032131-0123132120313301)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2300030120121300-3211123232301212-3310031313331032-0210001203202122-2231000320110011-0103003230002002-0233123333032321-2320121022103310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212211100321031-1123322212033011-2023200313010323-3220103211001012-0021323300302112-2321222300111030-0100202031000333-0030220222211130"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 310033013313 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-019.md#canonical-0322012333022122-0100132312200311-0311303222130121-3022223011010322-1103030132323301-2312013011122212-1111102310032131-0123132120313301)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3231032300030132-1312013312002231-3120111323303113-0032023112202323-3122220011303023-2130121101103102-3303222023101131-0303330222031321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-3201213311112003-2332301331330010-3313203300102331-0200112220300211-0200012011000023-3023203132311331-3003110320333030-2303223221302032"></a>

## Direct properties — strict_coalescing / 310033013313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200020300323130-0323221212012103-2311100001210110-2133111322132331-0320223320311322-3011120312202201-1122032333333101-0103221022012030"></a>

## Next pages — strict_coalescing / 310033013313 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-019.md#canonical-0322012333022122-0100132312200311-0311303222130121-3022223011010322-1103030132323301-2312013011122212-1111102310032131-0123132120313301)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2200130132022231-2002123300110310-2003330320002013-3233302321020131-3113330022110201-2300103331101032-1100003102020213-3003222221130110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332120332123033-1131010212312311-0201213111110130-3321101011321133-0230132232031013-0210021112033213-1002100330102203-2323200212310203"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header — default_header / 323231323312 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-1032111321213211-0203020121202230-1122202030101230-3131300322321022-0100203311011212-1300220001213221-2003231020130301-2020323012212333"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-3020122103002102-1011210211021100-1121003310000002-2320132010200002-1001231113122213-2233301033220201-1322200030120302-2332101123203033"></a>

## Direct properties — default_header / 323231323312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102233232000121-1012220131021311-1302330310130221-3202010231113022-2100303230210222-0101300313111021-3210033112002213-1322121022002232"></a>

## Next pages — default_header / 323231323312 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1031201332122110-1012222320310122-1210232321233211-1100211003330002-1103312323010320-0212232302231100-3112200331231011-3111321331302321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313021112023130-1233003101103233-0301033130023003-0323210333213310-2101332221023301-2331102200033100-2322320033021222-1210230130332220"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — default_loadbalancer / 330102311230 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-2323120012020002-3001030132132021-3222022312113001-3323202032031121-3002320201200213-0220213020200320-1103220032010332-1203230032210301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-1133003300211021-1132010113100020-0321013031330022-2232133300333131-0313002231313310-3202311012030001-0233130102102210-2112002232331132"></a>

## Direct properties — default_loadbalancer / 330102311230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220312120333031-3002312222030003-3021011323120013-1303113120002132-2002100232012021-2123032332033233-0020012023010203-0033231213001321"></a>

## Next pages — default_loadbalancer / 330102311230 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0333232132032030-3133133333130221-3333130121003233-3313013310132111-1300301101032101-1222111230011103-0200313120010222-2000020103232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122132102203222-2030032221233022-3030102021031311-1111303123320110-1002101100022103-3013121133003233-1112202233013012-3111303031110230"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — disable_path_normalize / 113022301202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-0133112121312302-1003333130133001-3213333333320221-1232001221033113-2223310030000120-1033233322023011-1303120120203012-0010022021300030"></a>

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

<a id="canonical-3322301112312030-0320312013023010-0012032031330303-2031130210130130-3000331322211000-2230332122012001-2113022301002122-2103311303110121"></a>

## Direct properties — disable_path_normalize / 113022301202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020320301333010-1010311201102030-3020301010011322-1233003130010222-0031233323033302-3012131113312211-3112110313030210-1103232202300002"></a>

## Next pages — disable_path_normalize / 113022301202 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3310031330333302-2300003212023313-3002323132300211-3202111001300101-2230233032203110-2313212031202102-1020223321123200-2303131333002003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300003210012332-3201330321013211-2333120332023203-1221113113002013-1011313022110103-1212300132102022-2331031113301020-1101310210321023"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — enable_path_normalize / 002031303032 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-0132113013221030-1111310303003330-3131201220232031-0321322033122100-1323212000103030-2300311022102331-0100033231012022-2012302232011312"></a>

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

<a id="canonical-3130302101130320-1211012312011121-2112023312331223-2003003222022123-3112222313023313-0131121023302321-2221332221122022-3003010322231303"></a>

## Direct properties — enable_path_normalize / 002031303032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102220003101231-1303100010000232-3201311330313210-2332223210103311-2022030113200313-2220330321020300-2001220112300032-2121033333203231"></a>

## Next pages — enable_path_normalize / 002031303032 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201220023322303-2132101301321010-2201201013112023-1230302033332233-0211202030211213-2210233201230103-2320210110211300-1321103313300300"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options — http_protocol_options / 213111033030 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-1022013210323031-0233031113303320-2310312230002321-0330233312221301-3023120313111031-3103303103131210-2332113223303022-0201301312313023"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-0200210131030123-0232302202030233-2132102133330031-0032303113321130-3322121212101330-1030331103303321-0303301221203333-3111121132102132"></a>

## Direct properties — http_protocol_options / 213111033030 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-019.md#canonical-1331001111003321-3100333210320120-3211000111030300-1030331020032322-2201110200020300-1133030011013022-0213323130331212-0132021030001003): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-019.md#canonical-3210220221112133-1001033100000123-2011102101310230-3200003233233100-3131132210120203-3332101011332031-0221303211003303-1130302321201301): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-019.md#canonical-2013300131201030-1132023011100221-3001232111223130-1130302210013231-1112303012131331-0122130030223102-2300021000331130-1013020032030033): complete subsection reference.

<a id="canonical-0301230021003302-0133120233022130-0311000321213310-3200220222122200-2313131000221212-2202032210321323-2000200030002123-2110222212030031"></a>

## Next pages — http_protocol_options / 213111033030 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-019.md#canonical-1331001111003321-3100333210320120-3211000111030300-1030331020032322-2201110200020300-1133030011013022-0213323130331212-0132021030001003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-019.md#canonical-3210220221112133-1001033100000123-2011102101310230-3200003233233100-3131132210120203-3332101011332031-0221303211003303-1130302321201301)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-019.md#canonical-2013300131201030-1132023011100221-3001232111223130-1130302210013231-1112303012131331-0122130030223102-2300021000331130-1013020032030033)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1331001111003321-3100333210320120-3211000111030300-1030331020032322-2201110200020300-1133030011013022-0213323130331212-0132021030001003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133022013120200-0000301002213120-1102112121333233-3330220031223313-0110301331223111-0000221103103121-1210001320002222-3030210102303010"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 313123232212 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0102223023311320-1110301212110030-1122323011202113-2323211312311010-0021020212221112-1321120230131300-3230023202012110-2203023021123010"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2012322100312011-1031020331112103-3331103332302122-0222310021022233-0203212322122013-0333133200331031-1102212131233312-1300131320032102"></a>

## Direct properties — http_protocol_enable_v1_only / 313123232212 / 3

- [header_transformation](data-sources--workload--reference--group-019.md#canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212): complete subsection reference.

<a id="canonical-2300012122321302-2000013232322332-1322231002002231-0032220133121312-3010223223111230-3012112101310010-1313113303210001-0222300213022000"></a>

## Next pages — http_protocol_enable_v1_only / 313123232212 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-019.md#canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012300222030230-1332033021031333-0203233303013330-0032000123112101-0011231211123213-2200103000123112-1230213210213202-0020011303011121"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 313130320221 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-019.md#canonical-1331001111003321-3100333210320120-3211000111030300-1030331020032322-2201110200020300-1133030011013022-0213323130331212-0132021030001003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0201332032011302-1331132203003120-3001001211210121-1001130303012333-0223131013320030-1112331131231323-3131012300212212-3323222003211010"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-0121032331102122-3102302213330332-2011233330232103-0211111001301132-0323333103113223-0012000031002322-3302020312212123-2101321330200011"></a>

## Direct properties — header_transformation / 313130320221 / 3

- [default_header_transformation](data-sources--workload--reference--group-019.md#canonical-2323013002103203-1132203301232010-1131012203310223-1121001021111100-3210300211320300-1301332113312113-1021332220201212-0122120022233123): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-019.md#canonical-0213311311313023-1233100103133200-2121312222120233-0310131113010322-1012120230310013-2301321323211002-3203211133130121-3011313311113202): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-019.md#canonical-1232301100021113-0101201110133211-3323201300201311-1132111120322111-3331321110103310-2013003312222232-1200332313010110-0320032000320313): complete subsection reference.

<a id="canonical-2021331301231211-0200313303022311-3232310231013201-2313320233302202-0113323001312301-3023011131301032-3103001200033003-1213331332230221"></a>

## Next pages — header_transformation / 313130320221 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-019.md#canonical-2323013002103203-1132203301232010-1131012203310223-1121001021111100-3210300211320300-1301332113312113-1021332220201212-0122120022233123)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-019.md#canonical-0213311311313023-1233100103133200-2121312222120233-0310131113010322-1012120230310013-2301321323211002-3203211133130121-3011313311113202)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-019.md#canonical-1232301100021113-0101201110133211-3323201300201311-1132111120322111-3331321110103310-2013003312222232-1200332313010110-0320032000320313)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-019.md#canonical-1331001111003321-3100333210320120-3211000111030300-1030331020032322-2201110200020300-1133030011013022-0213323130331212-0132021030001003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2323013002103203-1132203301232010-1131012203310223-1121001021111100-3210300211320300-1301332113312113-1021332220201212-0122120022233123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332032031020112-1110103003033233-3120212102311102-0001202303332222-0021001230310230-1311301201013220-1221312101132233-1102100120023203"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 021113212131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-019.md#canonical-1331001111003321-3100333210320120-3211000111030300-1030331020032322-2201110200020300-1133030011013022-0213323130331212-0132021030001003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-019.md#canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1203200333313121-1113122301122111-2012211132202000-1213122311233332-2032212032123311-2113023021003131-3300311232210311-1221102330202320"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0030033313131131-2120333021222002-0301113321331230-1021003310111212-0300223111112033-3331000132030001-3221333022002231-2323121010223301"></a>

## Direct properties — default_header_transformation / 021113212131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300201020013111-1220301213132202-1020233021300211-3222103211132123-2202212332220121-0232020123123000-1130032122031322-1210221023110312"></a>

## Next pages — default_header_transformation / 021113212131 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-019.md#canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0213311311313023-1233100103133200-2121312222120233-0310131113010322-1012120230310013-2301321323211002-3203211133130121-3011313311113202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132303021010031-2130200011230101-3233012231332223-2310030021232223-1332211321020103-2010132302231031-1300001100312121-0032111113211323"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 100031101031 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-019.md#canonical-1331001111003321-3100333210320120-3211000111030300-1030331020032322-2201110200020300-1133030011013022-0213323130331212-0132021030001003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-019.md#canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-3123233121331201-2102312303233300-3231321012312231-2100221000030021-0311103231000003-3221230202323311-2113333323303210-0122333010300212"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0201223313222221-1303313011131001-1033322203013233-3211021210023313-0301013003332112-2211101113122330-2201111332212210-1002302101233001"></a>

## Direct properties — preserve_case_header_transformation / 100031101031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133321201101301-1201303210131321-2330030110210213-1211003121311112-0111211132312230-1101021031303133-0122223300303133-0330002330002002"></a>

## Next pages — preserve_case_header_transformation / 100031101031 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-019.md#canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1232301100021113-0101201110133211-3323201300201311-1132111120322111-3331321110103310-2013003312222232-1200332313010110-0320032000320313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021323130323133-3222110111231233-3302023301323303-0221011312233122-0021322232223003-1110213003222102-3130322330032123-3031132110230102"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 031001033320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-019.md#canonical-1331001111003321-3100333210320120-3211000111030300-1030331020032322-2201110200020300-1133030011013022-0213323130331212-0132021030001003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-019.md#canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1300210033111300-2001211023033233-1132322131212031-2323220031200113-1210232313320311-0023230130100021-3020320100301113-3212200010230211"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203202001332113-3302202331121110-0332130300012302-3212030302330212-0302012203130130-0031121210000112-2013000112231210-2300301110123211"></a>

## Direct properties — proper_case_header_transformation / 031001033320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130121221202012-3311321000200030-2211013322221312-1022330022230010-3133011330223202-0103133022013112-2010313301022023-0000020002102133"></a>

## Next pages — proper_case_header_transformation / 031001033320 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-019.md#canonical-2311313320332200-0301211301330231-3320333312220303-1101102132303331-0132003311210012-0000110032203012-2311100231221010-0130200121001212)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3210220221112133-1001033100000123-2011102101310230-3200003233233100-3131132210120203-3332101011332031-0221303211003303-1130302321201301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300312311012302-3002133011131303-2220223300000211-2033011123120113-0032121212132200-3020123302013111-3202121012010133-2112231221032121"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 011012210012 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3332333210311202-1232300233133211-1201010323320212-2231331213103311-0233131000310010-0300101203003123-2233133020022131-3132111331130233"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-1113322232322012-1002320123031202-1233021210011230-0230103311122032-0101102200010320-3032203200213013-1312122013231233-3303203333223003"></a>

## Direct properties — http_protocol_enable_v1_v2 / 011012210012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203320110301111-1210311103010200-1111030301230313-0313201110210310-0200202123200313-3213300121232323-1010021031230020-0200011330102110"></a>

## Next pages — http_protocol_enable_v1_v2 / 011012210012 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2013300131201030-1132023011100221-3001232111223130-1130302210013231-1112303012131331-0122130030223102-2300021000331130-1013020032030033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231111323200220-3323213103111131-0033200231321223-3032322311210230-3022133200121001-1023302203030033-2311010233133111-2321020030003123"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 012311011101 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3020010330013332-0200320111012222-1113332133313230-3102312030202330-3210203332200220-0210323122223201-0202133030331030-1131203312200300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-3010310121302202-2032121133130331-2101210023002232-3010023033032231-2101231013100333-1022111030112202-0111301220000111-2330031323011012"></a>

## Direct properties — http_protocol_enable_v2_only / 012311011101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022221000021023-1030302100133030-1322312310031311-3232021032320230-1103120021132122-1000031021033020-0212213122322131-1133222002103213"></a>

## Next pages — http_protocol_enable_v2_only / 012311011101 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-019.md#canonical-0022310130212032-0323122331011101-3332212302320201-3111120200022203-1132023320300101-1221213311120321-1023201200131322-3213331323300213)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1301101230033112-3020330221011121-0310230213123321-3122232303313021-0313303222103221-2131330331331202-2220320110030033-3130123001220231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322311101233023-0323322003133121-1002200302321310-1212003122233100-2011330311113201-2321010133223203-3311110121020002-1333001033221001"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls — no_mtls / 300121030010 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-0203320200003333-0110002133301210-1100210233130323-0313110223133321-3233203310000131-2330300212023101-3200223121111322-0331020012300202"></a>

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

<a id="canonical-3320320300333023-1331212310310201-0332032330203321-2230301021031332-2330233301333101-3111201220001312-1232321220323011-1133210021112210"></a>

## Direct properties — no_mtls / 300121030010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323010111203301-2022110233112130-2033322113022111-0323130311220030-1232011231012110-2331132123211311-0133122320000101-1133332302331333"></a>

## Next pages — no_mtls / 300121030010 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0123212111101211-3221323200311300-3220232332021201-2002310222233130-3321330132002101-0303203113020332-1203102002030031-1011021221210230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322102220203113-1030130112000112-2302001120122013-2233231212130202-3102020303020210-3011033212231301-0323032002031110-2010333010113210"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 332232320131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-0133312201133132-0303122211012323-3132100133122332-3330321313011321-0101222313212203-0000200221301231-0302310110003332-0321100133101223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-2031222313102223-1321323300132203-1322121312303013-0300023303232320-3202002010313231-0300120103122022-3300120022123133-1222320321201223"></a>

## Direct properties — non_default_loadbalancer / 332232320131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213213322321213-0222331203112113-2313113113230223-1101200323023101-1031133132211212-0220331120123100-1122130110003130-3030210302223203"></a>

## Next pages — non_default_loadbalancer / 332232320131 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1101311113121322-3100323133230131-0321033001022313-1113030012312103-2110011131030201-3013233031033002-1300302030220331-1323302021203003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313021232200323-2023233320120303-2221012021120303-3011110110000013-3202213321311102-1301223130023102-2033312020302321-1021300121331333"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through — pass_through / 323021313111 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-1022311312133201-0021311321100122-1002121323112233-0030220022100000-0233102100313132-3312011203220030-1110021231200002-0330302333203020"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-3322333101022200-0103000113223032-0032311302022213-3023223103112203-3310032211331001-2012323230331120-2331002010020023-1132220202012022"></a>

## Direct properties — pass_through / 323021313111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312023211320212-0210332231201132-1122302332033030-1331112013331111-2100200003011303-3003312332020203-1210103210123300-3200332303131201"></a>

## Next pages — pass_through / 323021313111 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201210201120121-0133021323100211-3230211200113111-3230030011322301-2031213100031002-1311202112303321-2003001101300201-2013101022012002"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config — tls_config / 211033001312 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-2303001302113120-0033131212123132-1020133320301310-3022303221032210-1020210022101202-1121321321132000-3223010102122010-2100211021303300"></a>

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

<a id="canonical-0130212100302120-2212132303030222-1212010100001122-1113011032233121-0301200013003303-3000331201301020-1220012311321203-3111012302310133"></a>

## Direct properties — tls_config / 211033001312 / 3

- [custom_security](data-sources--workload--reference--group-019.md#canonical-2130031201022220-0133221130201233-2102011123210221-2202110130121332-1220113211222100-1230021313011331-2022310200200003-0210210131100320): complete subsection reference.

- [default_security](data-sources--workload--reference--group-019.md#canonical-3123213332103212-3130303033213300-1110022100031103-1213231303323131-0322023002013011-3320301030322233-3101011003103123-2301000010323030): complete subsection reference.

- [low_security](data-sources--workload--reference--group-019.md#canonical-3022033303210332-3200132323122021-1033222301312132-3231011132121212-0033002031033030-0003010122311011-2310332103101321-3301300130002310): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-019.md#canonical-1012220221211101-2202223232222013-1313013232111320-2203112323012220-0210231312322300-3212313222311322-2213103031220123-0320011201221301): complete subsection reference.

<a id="canonical-1221013212210212-2201313311132210-3032102032112110-2201113103331201-3012302323113230-3221012131102022-0031210031131002-1023312102010220"></a>

## Next pages — tls_config / 211033001312 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security](data-sources--workload--reference--group-019.md#canonical-2130031201022220-0133221130201233-2102011123210221-2202110130121332-1220113211222100-1230021313011331-2022310200200003-0210210131100320)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security](data-sources--workload--reference--group-019.md#canonical-3123213332103212-3130303033213300-1110022100031103-1213231303323131-0322023002013011-3320301030322233-3101011003103123-2301000010323030)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security](data-sources--workload--reference--group-019.md#canonical-3022033303210332-3200132323122021-1033222301312132-3231011132121212-0033002031033030-0003010122311011-2310332103101321-3301300130002310)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security](data-sources--workload--reference--group-019.md#canonical-1012220221211101-2202223232222013-1313013232111320-2203112323012220-0210231312322300-3212313222311322-2213103031220123-0320011201221301)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2130031201022220-0133221130201233-2102011123210221-2202110130121332-1220113211222100-1230021313011331-2022310200200003-0210210131100320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002320121000233-3000020103102133-1000213220113020-2312120300033220-3233200200231031-3003120113023112-3020223321333320-0001333231320201"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security — custom_security / 302101220120 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-2213303212322020-1033020031020310-0012212100123333-0131220230130033-3000302100100312-3233003032030232-3031012011001002-1032003223303022"></a>

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

<a id="canonical-1103112210321321-2013303313111210-0031332120132032-2102100010133330-1031003210131031-0011120223213123-3322103210131221-1030031113133011"></a>

## Direct properties — custom_security / 302101220120 / 3

<a id="canonical-1020020021133101-2211033033300013-0001132230323032-3111021230001031-2023101121112323-3131333330011310-3332333002311101-3201012312012200"></a>

<a id="canonical-3210320200233321-2320122020303123-0103131323320100-2131233122003132-3202333233303103-0311021011313320-0331212132302233-0320121102303210"></a>

## cipher_suites property — custom_security / 302101220120 / 4

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

<a id="canonical-0110303333111330-2330122323201020-0123003213110323-0131313123322112-0013011022011030-2313220033113031-0110320203133300-2210110332300133"></a>

<a id="canonical-2201321102310212-0113212212110220-2130330321011323-0222002113123221-1313302213220012-3111313332302322-0023123222132311-1001321313300321"></a>

## max_version property — custom_security / 302101220120 / 5

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

<a id="canonical-3121121201013210-0100313320333113-2132211110133002-2313231123032131-0131303123001233-1221032202131231-3133133321221231-0003122121113331"></a>

<a id="canonical-2300013033132022-3013133111122132-0302112323010233-1032032231110100-3021111231322302-1130022102011020-0003022101332223-2002221023301123"></a>

## min_version property — custom_security / 302101220120 / 6

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

<a id="canonical-0023230003221220-0221220300011303-3220002032121122-2120121222121011-1001323002223311-0200033323310103-3131330313200111-1331210323203301"></a>

## Next pages — custom_security / 302101220120 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3123213332103212-3130303033213300-1110022100031103-1213231303323131-0322023002013011-3320301030322233-3101011003103123-2301000010323030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021230000320211-1212021022033113-0232322223003203-1303233222101030-2232012130233332-0023101230203102-0323332211023220-1221000230231123"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security — default_security / 012223023310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-1202132113333001-3331110011203111-1001203212131313-3033200222222121-2032100301121220-3221333111302313-2020312310300130-0032221311020003"></a>

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

<a id="canonical-1001323103103133-2202123303113111-2200311332311222-0121313013013220-0312031110023120-1030011332203023-0130233101310211-0032213200023210"></a>

## Direct properties — default_security / 012223023310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320200001020033-0232333022002230-3132021120132212-2003023032131332-3103232113210332-0132011203001002-0001001000311230-0213310120203002"></a>

## Next pages — default_security / 012223023310 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3022033303210332-3200132323122021-1033222301312132-3231011132121212-0033002031033030-0003010122311011-2310332103101321-3301300130002310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321212122313213-0302111011001213-0232202203112020-1101311013133213-0201002302220013-3103011123113311-3312310101201222-3033331123223003"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security — low_security / 223112031331 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-1200302303301020-0031210101130122-0310303330101321-0223323331301322-2231033330320322-2031133220021011-0212010212200132-2310033030032020"></a>

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

<a id="canonical-0013200201112003-3122133132301223-1302121300001201-3210322100121332-1333301003022333-3211102201132221-0213123203302130-3213302102230232"></a>

## Direct properties — low_security / 223112031331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011302223012313-1111230020221101-2312130211312232-3132100110033010-1112021301331212-2212020002231022-3001232001123320-2330022332032131"></a>

## Next pages — low_security / 223112031331 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1012220221211101-2202223232222013-1313013232111320-2203112323012220-0210231312322300-3212313222311322-2213103031220123-0320011201221301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320010102313221-2110301100132210-2310000213202322-3100021200010120-1131222221133002-0132201320101002-3323130202012322-3011210221011112"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security — medium_security / 123322311011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-0232123300032021-2013201002213311-2120113213322121-3211333223121333-0012100323002033-1232120020233013-1203031200330222-0211120301020100"></a>

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

<a id="canonical-0322010102023200-2021000311330230-0221210232232101-1023010021313211-1031233212230302-3221132220221333-0101333320201121-2320230110001013"></a>

## Direct properties — medium_security / 123322311011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101213133301030-0321231121301003-3221100211203122-2201212221201121-1021312312121202-2220101221003002-2332200320030221-3330020231203200"></a>

## Next pages — medium_security / 123322311011 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-019.md#canonical-1030011121311021-2200321031212131-1100003220323232-3210300133021023-1111032322220103-1211222223110003-3130132122311121-1001333020311130)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220311133000222-0121220202313110-3333200321323010-3000013210201002-3121310203102102-2102221330100011-0001203010303112-1010333213020011"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls — use_mtls / 222210310303 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-0322332201323022-1122120232033203-1301032100113303-2022321320233020-3222311100101232-0333121220100013-1231201023021310-0120220103130101"></a>

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

<a id="canonical-2012332122220321-0110110322320300-1213011032100211-3120011302221300-0222211333320221-3131310311110210-2333020312010331-3322302303230301"></a>

## Direct properties — use_mtls / 222210310303 / 3

<a id="canonical-3110311221311233-3022011032032203-3021120332322012-2121113031221022-0111333003010211-0321031023321113-2221132213133110-1332032111310330"></a>

<a id="canonical-1232320313101023-2311233033010312-0211203301201102-3222203130323330-2220230022123131-3122023332232221-3031301330331122-0012133113321312"></a>

## client_certificate_optional property — use_mtls / 222210310303 / 4

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

- [crl](data-sources--workload--reference--group-019.md#canonical-3331013300223021-2021212101022320-0232331132021013-2201032323231321-3301211130331122-0031020301030023-0202000012031300-1201131000331000): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-019.md#canonical-3300203031320333-2202000101011002-2311012301013130-1223110233110110-2231020123010023-2132033123123233-0013011302010313-0103111122001032): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-019.md#canonical-3222203102011301-1122112132000001-3303011031321001-3021322030033131-3311101211302021-3322211310122003-0202222133220230-1221303023230112): complete subsection reference.

<a id="canonical-2220131330323220-3211003123233123-1133002033203110-1001012220310313-2302233202230231-2000201210303020-3021023001003111-2111022330323323"></a>

<a id="canonical-0300310033123212-1321322003302212-0310023233300223-1100311012003011-2003311200312031-2110022333211331-0033303232320301-0230201103313221"></a>

## trusted_ca_url property — use_mtls / 222210310303 / 5

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

- [xfcc_disabled](data-sources--workload--reference--group-019.md#canonical-1032212120033232-0302320103321232-0212301010212301-3002200131312311-0300103001211311-1021322302233201-1210000122302100-3331221200122301): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-019.md#canonical-0233310023013011-2212030213101032-3320003003000333-0303120231113103-1120310323222012-0213031012132300-3303001333220021-0031330222131333): complete subsection reference.

<a id="canonical-0001213333303001-0231212101333033-0201021112320203-2012032122301332-1330232033311110-2101030233023302-2231232032212111-2010200031212202"></a>

## Next pages — use_mtls / 222210310303 / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl](data-sources--workload--reference--group-019.md#canonical-3331013300223021-2021212101022320-0232331132021013-2201032323231321-3301211130331122-0031020301030023-0202000012031300-1201131000331000)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl](data-sources--workload--reference--group-019.md#canonical-3300203031320333-2202000101011002-2311012301013130-1223110233110110-2231020123010023-2132033123123233-0013011302010313-0103111122001032)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](data-sources--workload--reference--group-019.md#canonical-3222203102011301-1122112132000001-3303011031321001-3021322030033131-3311101211302021-3322211310122003-0202222133220230-1221303023230112)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](data-sources--workload--reference--group-019.md#canonical-1032212120033232-0302320103321232-0212301010212301-3002200131312311-0300103001211311-1021322302233201-1210000122302100-3331221200122301)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](data-sources--workload--reference--group-019.md#canonical-0233310023013011-2212030213101032-3320003003000333-0303120231113103-1120310323222012-0213031012132300-3303001333220021-0031330222131333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3331013300223021-2021212101022320-0232331132021013-2201032323231321-3301211130331122-0031020301030023-0202000012031300-1201131000331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223302330222120-2021323032313233-1133313123232313-2201003222133023-1200032200013112-0202302200122112-0002013303123320-1221203102222202"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl — crl / 103202122313 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-3223031202302333-0032130122133211-2233230331100122-2330313321101013-1132321130333033-3000230031213320-3330032123300111-0003233031133012"></a>

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

<a id="canonical-0230310233021131-3313221000221120-1332002131031300-2121221221030310-3201130132030011-2320311231002101-3210123121323032-2220312201030223"></a>

## Direct properties — crl / 103202122313 / 3

<a id="canonical-0202012223101120-1321132210021023-2111303003220202-1012110110023230-0020113010332100-0300102203021132-1323320000220332-1213103121102203"></a>

<a id="canonical-1303231122210310-1002212020112022-0300231123122113-2300322232022012-3212313300322131-0131002131300302-0323102220223200-2221333130212230"></a>

## name property — crl / 103202122313 / 4

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

<a id="canonical-3122213323122000-0111021110032233-0222212313322123-0333211000313122-0030220032212311-2311303320130111-0013322132110020-1222010320321103"></a>

<a id="canonical-3013220122123122-1021123232130311-0231113210203201-0030211322223333-3100100120332102-2320023002320010-0021310100102011-1221003002311033"></a>

## namespace property — crl / 103202122313 / 5

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

<a id="canonical-0133000300010110-0103013022112212-3131200003211031-0013230311312012-3132023122230203-3120122013012202-2000220231100210-0000333220021112"></a>

<a id="canonical-0321332323221030-0202330033023223-2220313122133310-1020002023313003-1132000120321312-2231012023200011-3302211003323133-3012122210132132"></a>

## tenant property — crl / 103202122313 / 6

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

<a id="canonical-2110212021302323-1112322213223301-2312031130220033-3210021111212322-3231210310311233-2310011111131200-2322231223121110-3003130213231023"></a>

## Next pages — crl / 103202122313 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3300203031320333-2202000101011002-2311012301013130-1223110233110110-2231020123010023-2132033123123233-0013011302010313-0103111122001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332212201312113-3311210223110032-3131120210122010-1220202030330233-2010102030122320-2232301321323111-1303012311223020-3221003313203100"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl — no_crl / 300001313312 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-3320132212121120-1113302022120121-1021110030200010-2002123232132003-3322010000311202-0333022213123323-0030023031012101-0130030123000100"></a>

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

<a id="canonical-1210301323203100-2233322113222333-2013101302121200-0021313310313103-1220312012003120-1023213223111210-2011010212221302-3312310123232132"></a>

## Direct properties — no_crl / 300001313312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222213331300023-3123331323322102-1033131212213302-0112200312221221-3321100131301233-1201232012200220-3313133102310010-1331101220012100"></a>

## Next pages — no_crl / 300001313312 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3222203102011301-1122112132000001-3303011031321001-3021322030033131-3311101211302021-3322211310122003-0202222133220230-1221303023230112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031303321303311-0110123231023220-3330200332211003-1012130222120111-1001220321123311-0203011313200231-2000111221203101-0302330012033222"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca — trusted_ca / 132321102113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-3010300223202003-3311231102112000-2323001021002330-3100200112021113-0310332321130332-2210101122110230-0332212011323222-2131003331221213"></a>

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

<a id="canonical-0320131123001332-2303222200003100-0232033200100321-0111303023220322-2132230113002233-1000300113333113-2033222121113120-3311333322231311"></a>

## Direct properties — trusted_ca / 132321102113 / 3

<a id="canonical-3201311012310121-0302113122032100-2331231121303031-0331013032213312-1030101330102331-3003001100200130-3103333113320022-0120210022030000"></a>

<a id="canonical-3322123020121312-3302312022120010-3032120211202113-0331300220230310-3022033301321320-1223120221010020-2021230212031331-1320022231020232"></a>

## name property — trusted_ca / 132321102113 / 4

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

<a id="canonical-2230102101203221-0321322323221202-3333100333323101-3302021120022022-1300323121302110-3232212022123123-3210332010202212-0122022200330112"></a>

<a id="canonical-0111222201100130-1023201322130023-1030303301322213-1200131123102220-3301313132233331-1201021122311103-1131122233011323-1200213101033000"></a>

## namespace property — trusted_ca / 132321102113 / 5

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

<a id="canonical-1020301202011130-1320300130103101-0232312233030322-3102223200313303-1302333002312011-1132221102210233-1002031011320201-3110031333001311"></a>

<a id="canonical-2023331212220101-3133122322030000-2221220023001210-1132231110023003-3123132311312022-0310330320211311-3000112320002301-1013222032222012"></a>

## tenant property — trusted_ca / 132321102113 / 6

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

<a id="canonical-0122103120021320-0000133212302010-0203010002112220-3203212001301323-0110131203110110-1313101122112210-3300202020331202-1030003332222332"></a>

## Next pages — trusted_ca / 132321102113 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1032212120033232-0302320103321232-0212301010212301-3002200131312311-0300103001211311-1021322302233201-1210000122302100-3331221200122301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131131320301323-3300210003312101-2121132332123020-0121011313220131-1103131322121013-2003011323310323-2231302012120032-0311310103310301"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 323130020202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-2313200121300223-3200111113003103-1322021303033010-0000123212220131-2202133022223221-3030322313100131-0102120311112210-2011112013020033"></a>

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

<a id="canonical-1022230031100033-3103112221113130-0230310301123110-1121232033233331-3333213023121221-1132220213200031-1323110333313333-1303310110031013"></a>

## Direct properties — xfcc_disabled / 323130020202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113303133322032-0000003011120303-1101120132120123-0102203300301000-0311322313023121-2300133122020330-2313011231223220-3010023111232301"></a>

## Next pages — xfcc_disabled / 323130020202 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0233310023013011-2212030213101032-3320003003000333-0303120231113103-1120310323222012-0213031012132300-3303001333220021-0031330222131333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023032233321322-0303132332033321-0333103121210021-0230201021202101-1000133032000131-3311331231002021-0131221102212303-1132033303312103"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options — xfcc_options / 012331012100 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-0203021111130232-2201123322313020-1112222330233233-0103032333012220-3201002303110030-0310103220310223-0131212203310223-0322002130031333"></a>

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

<a id="canonical-1111102031031130-2231121312220111-2212302101100033-1333202311012322-1311102222001333-0323201120220210-1210011211310302-2012302021231103"></a>

## Direct properties — xfcc_options / 012331012100 / 3

<a id="canonical-0013021210122023-1300012233221231-0023213113221012-2330120130110223-2032133100330132-2201213131100030-2123321303223032-0331123323323122"></a>

<a id="canonical-3032220033001322-3123220110002021-2200010123230113-3210200303021120-3003121310122110-3110001033032123-3132133311111320-3023010103023013"></a>

## xfcc_header_elements property — xfcc_options / 012331012100 / 4

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

<a id="canonical-3030332133022130-1031112310222301-3331022300013231-3221132031303102-2221323231210231-3312231333000011-1002322322103223-3230211330032003"></a>

## Next pages — xfcc_options / 012331012100 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-019.md#canonical-0133002120303232-1213230123211103-1003121320120302-3332113120122012-3003331102211032-3102100210323311-3213022210020102-3302133103000310)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300300232020012-1112333332103111-0312101113113013-3111023102322322-2111112320331133-2000132022001002-1011202312100220-0020111123130320"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes — specific_routes / 122100221123 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes

<a id="canonical-0012003120033331-3110121011013232-3100003112101033-0023300212031133-2120210210132213-1113102311112321-2103310331000210-1033221333331330"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1231113333131303-0233000100212312-1232222220313111-3021230002100330-3003212020012330-0123000200333321-2010332203201011-0003102200303201"></a>

## Direct properties — specific_routes / 122100221123 / 3

- [routes](data-sources--workload--reference--group-019.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221): complete subsection reference.

<a id="canonical-1103221131221020-3302133001022012-3121310000331032-0203331032203122-1222100000100312-3322221230211021-0020213310122132-1022000331131031"></a>

## Next pages — specific_routes / 122100221123 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-019.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033322303211323-0120000003302210-2313311112131223-1133132301131303-1201230230303230-0133302110330201-2332023230311330-0320030321011312"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes — routes / 101213213003 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-019.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-3133321110130321-0010132323232201-2220033300100302-3003312330013032-0110300021300013-3210213112022021-1203301133022020-0301021203032010"></a>

Type: `"list"`. Computed.

Routes. Routes for this loadbalancer.

Upstream description:

Routes for this loadbalancer.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-0032021023030323-3131013213323022-0312120221332131-3211202333002212-3000033211013213-0103023231302200-1132321302212220-1213022120200010"></a>

## Direct properties — routes / 101213213003 / 3

- [custom_route_object](data-sources--workload--reference--group-019.md#canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021): complete subsection reference.

- [direct_response_route](data-sources--workload--reference--group-020.md#canonical-2002311320330000-0130001113100201-1112221113021002-1032300311301213-1120312301131132-3120112000312202-2323022323313233-3303332120203323): complete subsection reference.

- [redirect_route](data-sources--workload--reference--group-020.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003): complete subsection reference.

- [simple_route](data-sources--workload--reference--group-020.md#canonical-3212320223022021-3213110113203220-3102231111023223-3311013332312230-0223232223121323-2120320230310210-1222202102333233-0031000030332110): complete subsection reference.

<a id="canonical-0231221223211103-3311300331300132-2020123130322300-2012203002221231-1113300232111212-2102021223033210-2211111202310331-0200303323203103"></a>

## Next pages — routes / 101213213003 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-019.md#canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-020.md#canonical-2002311320330000-0130001113100201-1112221113021002-1032300311301213-1120312301131132-3120112000312202-2323022323313233-3303332120203323)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](data-sources--workload--reference--group-020.md#canonical-1321230220033300-2331122222200003-3002001213201022-1231111322210322-2122202133002220-2020331313130011-0023111323322033-3013101310010003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](data-sources--workload--reference--group-020.md#canonical-3212320223022021-3213110113203220-3102231111023223-3311013332312230-0223232223121323-2120320230310210-1222202102333233-0031000030332110)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-019.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001302120013112-3222210110100333-3033201322202103-3221321012213100-1110112222100202-2213211220213122-3331312002312213-3121101310333122"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object — custom_route_object / 102112312232 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-019.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-019.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-2330132123222013-0210003011200332-1321102010020303-1123312033030232-3023033112303110-0303312213321332-0012332313002011-1233332021111233"></a>

Type: `"single"`. Computed.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

<a id="canonical-0023230303121333-2033322133022210-1202013200232030-3033133122030113-2201212022020022-1123300122011330-2323321332310103-1220120121212212"></a>

## Direct properties — custom_route_object / 102112312232 / 3

- [caching_disable](data-sources--workload--reference--group-019.md#canonical-0231321100133302-3222022021103000-3322113023221133-0112321031003210-1003000003133113-1320031223313021-1303000003313320-0103201223233223): complete subsection reference.

- [caching_inherit](data-sources--workload--reference--group-019.md#canonical-3323201223312101-3102322021300131-0000221012003302-3002120133332002-3302032130033002-1203002022001011-0301203232330003-3132013023222201): complete subsection reference.

- [route_ref](data-sources--workload--reference--group-020.md#canonical-3333110013011302-2000302332230032-2012313330021211-3211023012223321-1233010212213312-3133203011021211-3002000302202132-0313312131211201): complete subsection reference.

<a id="canonical-3223331113013211-1230003011302111-1120312003221012-0301031000222233-2030031333321121-2000203021010011-2023222321323323-0133101201223213"></a>

## Next pages — custom_route_object / 102112312232 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](data-sources--workload--reference--group-019.md#canonical-0231321100133302-3222022021103000-3322113023221133-0112321031003210-1003000003133113-1320031223313021-1303000003313320-0103201223233223)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](data-sources--workload--reference--group-019.md#canonical-3323201223312101-3102322021300131-0000221012003302-3002120133332002-3302032130033002-1203002022001011-0301203232330003-3132013023222201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](data-sources--workload--reference--group-020.md#canonical-3333110013011302-2000302332230032-2012313330021211-3211023012223321-1233010212213312-3133203011021211-3002000302202132-0313312131211201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-019.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0231321100133302-3222022021103000-3322113023221133-0112321031003210-1003000003133113-1320031223313021-1303000003313320-0103201223233223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202330100101202-0102201112033130-3132331322233010-0232220312311120-0231223333223122-0123302222233032-2320132202233032-1303301233323322"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable — caching_disable / 313301220022 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-019.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-019.md#canonical-3301103000210323-1301210301023200-3022211310211003-1212303120002000-2233121112310223-3213001031031302-1030112321313111-0022233133331221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-019.md#canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-0332323313303233-2321101330101103-2302113110322122-1100331100133132-3030133331313010-2322011133003302-1022100231222313-0330333012312031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching disable.

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

<a id="canonical-3003312013203100-3003302120021220-2123301122232331-3103221202321020-3221220113333231-0030122202221220-1213012233023332-0323213300130023"></a>

## Direct properties — caching_disable / 313301220022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101022232110322-0022320101130313-3113332303312233-0011101013123132-0011203030303221-0303012103130011-3203310230103000-3110020321102121"></a>

## Next pages — caching_disable / 313301220022 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-019.md#canonical-1020111301330012-2123311021030130-1222110320323130-3231113100311032-1312221202113231-2323202110010213-1222313032223302-1222312330303021)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3323201223312101-3102322021300131-0000221012003302-3002120133332002-3302032130033002-1203002022001011-0301203232330003-3132013023222201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
