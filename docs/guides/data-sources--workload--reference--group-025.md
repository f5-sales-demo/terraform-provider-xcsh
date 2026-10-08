---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1220202112323031-3111311011331020-1201000010030320-1311310133200321-2111102301320013-0013123012133223-1110133102012211-1130231012310231"></a>

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0213331323130200-2222333111310222-0011113020221013-1030103312301133-3010002220200223-3100001113233130-3033122310121101-2102023333201031"></a>

<a id="canonical-2120022220202312-3033332123232233-1223033320233331-3323122120210231-2223113103333003-0221310230311121-2233223203311232-0130312230030220"></a>

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0032320022012020-1333113003320330-2300102003230231-3000311133020020-1032310222312100-2013232033301300-2332001210210001-0300331022301100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-024.md#canonical-0231230331021312-3333101230213213-0310012020001323-2320030300101223-3001310131113102-0200031101113321-3330320112132230-2313300010123231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-3201320022123112-1222022111023021-1221332030322313-1130020100231313-1211301033311020-1230101030313310-3023031311032323-2012212102222102"></a>

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

<a id="canonical-1010221000110112-2101231031301222-3212121303112313-0300111201100101-0123113113211302-2121121331212012-0023032203123212-0131123000310002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-024.md#canonical-0231230331021312-3333101230213213-0310012020001323-2320030300101223-3001310131113102-0200031101113321-3330320112132230-2313300010123231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0303003022201303-3220013220300121-0222303232023133-0101203133311300-1122201212112003-2132122221012000-3331000102031313-3310020300002030"></a>

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

<a id="canonical-2010223110000300-3311233023031332-0312032010102230-2132233022313123-3302202131011230-0301113330213322-3020031030023202-3121113201103022"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-0321133233233101-0232030221301033-1201221100332311-2121223121010023-3020102323012123-3030322302003032-2131121001223223-1011033120330001"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2200320111212110-2213132330231320-1111023032333220-0030002111202320-0130113220001121-1203233130303332-2200032003200322-2003102132022133"></a>

<a id="canonical-2300233220320202-1020230312200310-2033033002012133-1130311203123021-0033122023013110-1210320322300130-1103311123210300-3010110322020112"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2213302213232031-0003111223002032-1231120313103133-2123303233100201-1301321031312232-3111300111320332-1311002030112322-2303023030022122"></a>

<a id="canonical-3332213001213323-1321021132103321-0023303231133131-3022012103332121-1312233003222110-3010110222011121-0001233111012311-2211213332321332"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0331200310110320-1302010120001200-1032220122300213-1303212000232222-3203101003221210-3212111332132200-1023120122331301-1203330313012132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-024.md#canonical-0231230331021312-3333101230213213-0310012020001323-2320030300101223-3001310131113102-0200031101113321-3330320112132230-2313300010123231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-1110311303113232-2023121310111310-2212003332110113-1121111321332011-3202213232121023-2000020110103330-0223201133102132-0201301103210011"></a>

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

<a id="canonical-1322220012202300-2332001221123030-2321023200031200-0000311110311323-0232032000010231-1022032232003030-1001112033221010-1303211233212100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](data-sources--workload--reference--group-023.md#canonical-3130202111132300-0301212030033211-3312303220203330-1002132121130131-3322330013202010-3013320012301202-1031201103211333-3201303313311311)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-024.md#canonical-1213321303121301-1102133032030121-0221111123231231-2130000023010203-0303210000223333-1100020323111330-0120221132032311-1020313231211133)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-024.md#canonical-0231230331021312-3333101230213213-0310012020001323-2320030300101223-3001310131113102-0200031101113321-3330320112132230-2313300010123231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2331101332002232-2012130220202102-1301322103110130-0103112100200030-0130331220300323-1111133111310131-3133333123301313-3320132100201210"></a>

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

<a id="canonical-1231201012010120-3000102310113303-0130113231333322-2010202120133310-3233313302013333-2100022200212003-3011212011133200-1021013000101332"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-3122233101333120-1001031213312033-1103333232110223-3103311013211203-3030022300031013-3233221131032012-3112333230212030-0032333000313132"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert

<a id="canonical-2103302320223330-3223031112203013-2032032023011100-3031332033103022-0301012330010100-2320202111013211-3321230020221033-1022223213213012"></a>

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

<a id="canonical-3210101302112030-3010002033330020-3230321002102022-1311332113202230-3113010000320002-0220300203223011-1233220010312213-1002211200320302"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert`

<a id="canonical-3333211112331210-3032131200133013-2033020233021232-2010120331001312-3200322010130321-1013123103000320-0103110232122023-1132323030113112"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.add_hsts` property

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

<a id="canonical-3111021012110133-0122021230000203-3200203031222332-2021311320023023-3230213031300321-1212201222133231-3322330312332300-1021121210322310"></a>

<a id="canonical-2300310022022121-3111100311033321-3230202022101000-3000221220230210-2101121022003313-3310102130130313-0222303003321111-1023021000130120"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.append_server_name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [coalescing_options](data-sources--workload--reference--group-025.md#canonical-2133311121321113-0222103122223333-2122012130022232-0130003313312333-3220222303130030-1010102230033010-1303323101232113-3102020331132201): complete subsection reference.

<a id="canonical-1110223023331323-3013201230030011-3001310032230222-0310020332302322-2021301222112232-2113323113111032-2330120020030101-1133132232130020"></a>

<a id="canonical-3311330320011132-0311032331113030-3203132111300310-2210012331120323-0231101301300000-3131232230011303-3102201202120332-0200331210232110"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.connection_idle_timeout` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [default_header](data-sources--workload--reference--group-025.md#canonical-2111220010233112-2310212333012222-3223213320331000-0201300122101012-3010013120121011-2302203333303121-2123202220031112-2131112203212331): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-025.md#canonical-3200230200001001-1302132303311301-1310100100030123-2212022301103301-3103132330022123-2331221201122300-2320031212000232-1132011310301333): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-025.md#canonical-0111213230232223-2320133002012222-1130200102123121-3233303130133123-2202333020103223-1312103010330023-0320330033203212-1222011223212123): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-025.md#canonical-1333231203210122-1322302011031310-2323332002122331-2133223233002131-0031201111021000-2230131021222103-0000021020302133-2133333102120210): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-025.md#canonical-0130331103032203-1230202113332311-2330113032122321-2110030000322211-2133101322320003-2121233312103000-2001032231133033-0113111102113213): complete subsection reference.

<a id="canonical-1013332232110123-1210222011130230-1123002022001220-1311321320322131-3013233130030231-0010002220321221-0300330210213233-3031131003211323"></a>

<a id="canonical-1201101332231202-2323022221100012-1221331231302333-2123321013023311-2310233313211000-0031100130222021-1010023112333113-1231130221312113"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_redirect` property

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

- [no_mtls](data-sources--workload--reference--group-025.md#canonical-3031332232130313-1113300033303302-3202203221131322-1000012211312112-0231123201320023-0313133103122133-0201231100203221-0003132231113211): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-025.md#canonical-0332021101020222-3130230230022323-2203302010210022-3121321310010013-3113330231010132-1012022221101132-0120000202113301-0230310222221230): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-025.md#canonical-2333331233222233-0210001110231132-2333301023203222-2031131013031130-2313323221102333-1002210122013321-3213300313132322-0123233003032332): complete subsection reference.

<a id="canonical-3011012101221311-3120212013133023-1120033120122313-1131021222213221-3331320003213323-1000211111312101-2320032120010333-2201210222131012"></a>

<a id="canonical-3003213203133333-2330033123120312-2231312132313223-3012132333323001-3202010221121131-1010300132232121-0321103211211301-3101011231132022"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.port` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0332103012020021-0213123033132313-1323210331033232-2100323221202130-3303010001133032-3031200200202232-2111322033013001-0303033210330223"></a>

<a id="canonical-1312320321230121-2031301310011333-0233003032031333-3112012320203020-2032231001301032-0012322021320122-3012231213132333-3110112011012112"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0212322022233302-0222102021213200-0101021011201011-2303102221103333-0032003331020233-2300030133301030-2020213002222123-1020311031110200"></a>

<a id="canonical-1103302331133213-2302123300131332-0203202231022210-1023213322103032-1312101020022132-3200211211003102-0033002210312100-0303311111222033"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.server_name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [tls_config](data-sources--workload--reference--group-025.md#canonical-1020121131333112-2310331230322203-0002310223203320-3311021312310001-3002331002323331-0030223103110210-3113221020033311-3231123100110220): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-025.md#canonical-2220102110322333-3310111320012032-3312031312311011-3311231020200131-1231131001331220-2120133121131110-0032213010002000-0200332111303232): complete subsection reference.

<a id="canonical-2133311121321113-0222103122223333-2122012130022232-0130003313312333-3220222303130030-1010102230033010-1303323101232113-3102020331132201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-0210101332012200-0020003120220100-3301212020112031-3332112021211113-3231303221303100-2223230220122321-2102300220130022-1303232102222302"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

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

<a id="canonical-0103321023011023-1120203230310121-1230320013132202-2103003033212021-0332113322132210-3220333032030022-0232101123011000-1232322101121300"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options`

- [default_coalescing](data-sources--workload--reference--group-025.md#canonical-0023323031120210-2203100310321212-1221101333131101-2121213333211223-1100110000220031-1331123120220132-1332122032103021-1021111003100020): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-025.md#canonical-3132302102033001-0023130221232100-3300101002332102-2023121330012330-1201133313202333-1023123001332031-1302230021232032-2011000102230211): complete subsection reference.

<a id="canonical-0023323031120210-2203100310321212-1221101333131101-2121213333211223-1100110000220031-1331123120220132-1332122032103021-1021111003100020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-025.md#canonical-2133311121321113-0222103122223333-2122012130022232-0130003313312333-3220222303130030-1010102230033010-1303323101232113-3102020331132201)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-1110033213202113-2311022022120122-3220312202030300-3310212203123222-1003301122333203-2302230233301121-0231312113333120-2302211111002013"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-3132302102033001-0023130221232100-3300101002332102-2023121330012330-1201133313202333-1023123001332031-1302230021232032-2011000102230211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-025.md#canonical-2133311121321113-0222103122223333-2122012130022232-0130003313312333-3220222303130030-1010102230033010-1303323101232113-3102020331132201)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-0000323223232212-3130013110103102-2000021311331303-1301303200202032-0101102230030233-2023130003220001-0222201302030303-0033202322112301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-2111220010233112-2310212333012222-3223213320331000-0201300122101012-3010013120121011-2302203333303121-2123202220031112-2131112203212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-1100333322331232-1120332132102333-2301331332313032-1232021322231302-0110201000011122-1311101333121111-1212112032101333-1123312322112012"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-3200230200001001-1302132303311301-1310100100030123-2212022301103301-3103132330022123-2331221201122300-2320031212000232-1132011310301333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-3021003100113003-3222203231002213-0312033330132010-3133112101302121-1031230203212233-3311103220103011-3120211211222303-3311113320311332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-0111213230232223-2320133002012222-1130200102123121-3233303130133123-2202333020103223-1312103010330023-0320330033203212-1222011223212123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-1320031213300210-1223212012332002-3222020021331313-2001122323113003-1132223020302010-1323003222221023-3113113312200133-3303020001022212"></a>

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

<a id="canonical-1333231203210122-1322302011031310-2323332002122331-2133223233002131-0031201111021000-2230131021222103-0000021020302133-2133333102120210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-0322103120221220-1132111033103113-0110013121223002-1231233013103103-0220020332233011-0110001203230301-2110201310001330-1033103212203231"></a>

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

<a id="canonical-0130331103032203-1230202113332311-2330113032122321-2110030000322211-2133101322320003-2121233312103000-2001032231133033-0113111102113213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-1111113201111310-3310230113130013-0322111101100210-1133011012013212-2120301211112200-1013221311223221-1220300232110110-3323121231200011"></a>

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

<a id="canonical-3103100203001231-3100321210102011-1303211102032322-2110210031330113-2313233132300230-2011130202220023-3303011212022310-3200031323011310"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-2210332302012112-3010202233120313-0330212003320122-1210222201323030-3101013232202131-2230333032313321-3331100332301311-1232301302311120): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-025.md#canonical-3230230223103201-3330011021330203-1333131201210330-2223300332121011-2301313212333033-1033231123313212-3001023331001201-1033233113200032): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-025.md#canonical-2312110200301221-1203332303311322-0220001023212322-0012011112130101-0011101022330031-0130321103012211-0111223302110321-3021320322101110): complete subsection reference.

<a id="canonical-2210332302012112-3010202233120313-0330212003320122-1210222201323030-3101013232202131-2230333032313321-3331100332301311-1232301302311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-0130331103032203-1230202113332311-2330113032122321-2110030000322211-2133101322320003-2121233312103000-2001032231133033-0113111102113213)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3233002212200013-1011133302213201-1330012120313232-1300201212202021-2321121312222331-2100233220021113-3232021223200033-0322023200011323"></a>

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

<a id="canonical-0121030023213301-2020100001223131-0013213113333012-3201212103300230-0011230000132323-2011033312200212-0202120013320110-1031223210033000"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--workload--reference--group-025.md#canonical-3211323310302300-1020022001232231-0031022302231322-0303120213220310-3000203332012333-3332011231033311-1131020002131032-0321033202021330): complete subsection reference.

<a id="canonical-3211323310302300-1020022001232231-0031022302231322-0303120213220310-3000203332012333-3332011231033311-1131020002131032-0321033202021330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-0130331103032203-1230202113332311-2330113032122321-2110030000322211-2133101322320003-2121233312103000-2001032231133033-0113111102113213)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-2210332302012112-3010202233120313-0330212003320122-1210222201323030-3101013232202131-2230333032313321-3331100332301311-1232301302311120)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0233230131021003-0132212120123033-3321313303312011-0003020031323132-1011323222031030-0103030010111130-0231213101231022-0313233031022220"></a>

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

<a id="canonical-0302121103201301-0330213002013101-2231022211202221-0001332221030212-0221320032131010-1203103220013213-1122321111000331-3332321010100211"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--workload--reference--group-025.md#canonical-2111001132330310-0000011020131232-1113222333013323-0111300130221201-2301201232122122-3103233023320010-0213110203132222-3300020022101232): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-025.md#canonical-3332130033303321-1331133321323331-3300103332213010-0330302113000101-3112301221202323-2121011002211222-0001031301230010-0302220313012002): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-025.md#canonical-2110023311132303-2312123022033221-3230222231132230-2233210300331300-1100320131103000-0302213212323101-1103202100130313-2111212233320100): complete subsection reference.

<a id="canonical-2111001132330310-0000011020131232-1113222333013323-0111300130221201-2301201232122122-3103233023320010-0213110203132222-3300020022101232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-0130331103032203-1230202113332311-2330113032122321-2110030000322211-2133101322320003-2121233312103000-2001032231133033-0113111102113213)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-2210332302012112-3010202233120313-0330212003320122-1210222201323030-3101013232202131-2230333032313321-3331100332301311-1232301302311120)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-025.md#canonical-3211323310302300-1020022001232231-0031022302231322-0303120213220310-3000203332012333-3332011231033311-1131020002131032-0321033202021330)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0031200030023310-2231120122133222-1121131032311023-3220110112122002-2221323131020230-0031031110322221-0100200001213131-2203233131011100"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332130033303321-1331133321323331-3300103332213010-0330302113000101-3112301221202323-2121011002211222-0001031301230010-0302220313012002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-0130331103032203-1230202113332311-2330113032122321-2110030000322211-2133101322320003-2121233312103000-2001032231133033-0113111102113213)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-2210332302012112-3010202233120313-0330212003320122-1210222201323030-3101013232202131-2230333032313321-3331100332301311-1232301302311120)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-025.md#canonical-3211323310302300-1020022001232231-0031022302231322-0303120213220310-3000203332012333-3332011231033311-1131020002131032-0321033202021330)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-3002010300221213-2200021112322213-1001313001132313-3112303222331301-1002000312110121-3201313201000022-0332202120132012-3312130112110213"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110023311132303-2312123022033221-3230222231132230-2233210300331300-1100320131103000-0302213212323101-1103202100130313-2111212233320100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-0130331103032203-1230202113332311-2330113032122321-2110030000322211-2133101322320003-2121233312103000-2001032231133033-0113111102113213)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-025.md#canonical-2210332302012112-3010202233120313-0330212003320122-1210222201323030-3101013232202131-2230333032313321-3331100332301311-1232301302311120)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-025.md#canonical-3211323310302300-1020022001232231-0031022302231322-0303120213220310-3000203332012333-3332011231033311-1131020002131032-0321033202021330)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0331213130321200-3030121312033300-0221100310132210-2021310303023031-3030011222130332-2102302013003223-0100322130220332-0120100201123023"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230230223103201-3330011021330203-1333131201210330-2223300332121011-2301313212333033-1033231123313212-3001023331001201-1033233113200032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-0130331103032203-1230202113332311-2330113032122321-2110030000322211-2133101322320003-2121233312103000-2001032231133033-0113111102113213)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3003013220033332-0223132122110210-0111313301233032-0133033323021212-2331211310133303-3002100210133231-0103020033123111-3020222131012233"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-2312110200301221-1203332303311322-0220001023212322-0012011112130101-0011101022330031-0130321103012211-0111223302110321-3021320322101110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-025.md#canonical-0130331103032203-1230202113332311-2330113032122321-2110030000322211-2133101322320003-2121233312103000-2001032231133033-0113111102113213)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0021303332132131-3122031003110321-0201213012332131-3020332212312002-3231030113210013-3201333231032021-2323012022100133-3021331030201231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-3031332232130313-1113300033303302-3202203221131322-1000012211312112-0231123201320023-0313133103122133-0201231100203221-0003132231113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-3002020030122010-1203223012102032-2322233322002321-2113123010230011-1331232201311223-3203200210030023-0101021113301221-2330223330232232"></a>

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

<a id="canonical-0332021101020222-3130230230022323-2203302010210022-3121321310010013-3113330231010132-1012022221101132-0120000202113301-0230310222221230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-3221321012210233-2113130211312011-2023003223003112-0011001010201010-3312133222301122-0232100313300100-0113212221121220-0103333033233333"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-2333331233222233-0210001110231132-2333301023203222-2031131013031130-2313323221102333-1002210122013321-3213300313132322-0123233003032332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-3202202311003021-3103203203332122-0131010323100200-2131133321101023-1323300110221101-2210210020212331-1331101320122131-3200222112131203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-1020121131333112-2310331230322203-0002310223203320-3311021312310001-3002331002323331-0030223103110210-3113221020033311-3231123100110220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-3023201202323010-3033013233333011-0101312122020322-1213332322123211-0212030322331123-3211111222230001-3333200021112323-2031030330013212"></a>

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

<a id="canonical-3232211332223323-2030210220322101-2212302211010233-2332300202221300-0222111032201300-3233020300213112-3000132320313101-2011100210320122"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config`

- [custom_security](data-sources--workload--reference--group-025.md#canonical-0033302020200001-3321330233023133-2323101300023102-3023112121320003-1321310010032022-3020331133223123-3122221130313023-3232020110332233): complete subsection reference.

- [default_security](data-sources--workload--reference--group-025.md#canonical-3023213000021031-1231231022211303-2120022332310300-3320313331232120-1023003303112310-2300113201021321-0113111011130233-3112121323220013): complete subsection reference.

- [low_security](data-sources--workload--reference--group-025.md#canonical-1302013000120032-0331023113203201-1012133232212212-2103231022300300-0101213013213033-3031032213010000-3312133323303300-1312022023231113): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-025.md#canonical-1202201012303211-1031121032122300-2321333312000033-2022221301233120-1322011022013111-1320130221332110-2033102213312103-2331303123113230): complete subsection reference.

<a id="canonical-0033302020200001-3321330233023133-2323101300023102-3023112121320003-1321310010032022-3020331133223123-3122221130313023-3232020110332233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-025.md#canonical-1020121131333112-2310331230322203-0002310223203320-3311021312310001-3002331002323331-0030223103110210-3113221020033311-3231123100110220)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-0102333220000233-1130010220103303-3233230130023311-2203202323110011-1113300232331033-3202303231301101-3322323313221010-3131022100233311"></a>

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

<a id="canonical-0312122010023200-0122330310233123-0023200102102322-1221030233200213-1332132133330313-2010022103213033-3301300232300000-3301110033321211"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security`

<a id="canonical-0120112203301222-2310122121223203-1200210010301201-1201130121010233-1023220223313012-3302020133322313-3320223213332013-2012123001020111"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0111000300302301-1332031331321101-2030312111210111-0321012013123313-0312332223113033-1131301002010020-2303230203003121-1130032103312022"></a>

<a id="canonical-1331303023303101-1003230321311000-2033031120333110-1200133001210222-1122102312302012-1232201010001310-1121122110311031-1322100023300100"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security.max_version` property

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

<a id="canonical-2012323300102232-0131122313000222-0300233130302110-0032121331110303-0123112212233021-2302131023323033-0131031132130200-2001130221322331"></a>

<a id="canonical-2112100133213130-2312220030323020-1203022212110023-2020102121003022-3223333103031213-2032213201232011-1003001232133101-1233232330301221"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.custom_security.min_version` property

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

<a id="canonical-3023213000021031-1231231022211303-2120022332310300-3320313331232120-1023003303112310-2300113201021321-0113111011130233-3112121323220013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-025.md#canonical-1020121131333112-2310331230322203-0002310223203320-3311021312310001-3002331002323331-0030223103110210-3113221020033311-3231123100110220)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-2222313022222001-2233103321123110-2220212331300022-2030100112231323-1111211123112100-1031012100320002-1232121211123103-2133100331030201"></a>

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

<a id="canonical-1302013000120032-0331023113203201-1012133232212212-2103231022300300-0101213013213033-3031032213010000-3312133323303300-1312022023231113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-025.md#canonical-1020121131333112-2310331230322203-0002310223203320-3311021312310001-3002331002323331-0030223103110210-3113221020033311-3231123100110220)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-0100321121021121-0031301211231030-3103031200321133-3310021011132101-2322100212320300-1203023132210232-0133213001333300-0111231032100312"></a>

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

<a id="canonical-1202201012303211-1031121032122300-2321333312000033-2022221301233120-1322011022013111-1320130221332110-2033102213312103-2331303123113230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-025.md#canonical-1020121131333112-2310331230322203-0002310223203320-3311021312310001-3002331002323331-0030223103110210-3113221020033311-3231123100110220)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-2132230200301331-1121130023232033-1032002310133102-2312333011100302-0222012213020020-2323211211223200-0122131302230313-3120330111300101"></a>

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

<a id="canonical-2220102110322333-3310111320012032-3312031312311011-3311231020200131-1231131001331220-2120133121131110-0032213010002000-0200332111303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-1320021130222111-0101023003212102-0130012023300310-0230232302223130-3320032203211332-2010113302013021-1332112110022030-2323120013211322"></a>

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

<a id="canonical-0113331101211213-3332223012101322-2333332133021233-3001311201311220-3312003130022023-0333003111222220-0032213032322331-1323020221123111"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls`

<a id="canonical-1330102132013130-1000301133131001-3312010022323203-1313003130213233-3002101213012230-2132202223230100-0132032313122322-0322331020203010"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.client_certificate_optional` property

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

- [crl](data-sources--workload--reference--group-025.md#canonical-1013033200233230-1201220131032333-1121013110202111-3310101300310103-3002302330221200-2333003332112020-0232312123003003-1001010031032231): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-025.md#canonical-0122033132313021-2112210020320110-0133230103123230-0202310221000300-0303030020233330-2303131020232210-3231022113131120-2320000002203112): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-025.md#canonical-2022121213320132-2202302100330233-3013003002331120-3022132221300003-0202122012331313-0100221320310320-0313133100210020-0130200302121233): complete subsection reference.

<a id="canonical-2310123301331000-2131330101113133-0010311130133023-2310202000322332-3103130211020332-1233320311302012-0333031120121100-0333213200203021"></a>

<a id="canonical-2233230001120013-3213122212222133-1022010320212013-3132303122033102-1023133230303123-0321302310202231-2310331032030232-2022011102311112"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [xfcc_disabled](data-sources--workload--reference--group-025.md#canonical-3101332213003202-3112122010322232-2231312220233232-2211003222122222-3302030120012201-0202321003320333-2010023312230013-0211001322023120): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-025.md#canonical-1220210010013003-0022222302310312-1012013031030103-3101331310333233-1022310302102031-1311201111000001-3331122103313013-3021302312132312): complete subsection reference.

<a id="canonical-1013033200233230-1201220131032333-1121013110202111-3310101300310103-3002302330221200-2333003332112020-0232312123003003-1001010031032231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-025.md#canonical-2220102110322333-3310111320012032-3312031312311011-3311231020200131-1231131001331220-2120133121131110-0032213010002000-0200332111303232)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-2120011313311133-3311013202100000-3103001232331030-1301003230130000-1333322003002333-2212223103322221-0133113102020232-2320012221201001"></a>

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

<a id="canonical-1310231012123012-2312123032022103-0220002020333320-2120210203110322-2211002311130012-2211222023123331-3233202303131012-3023021012122333"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl`

<a id="canonical-2130031102330000-2123021002110311-0213032301133232-3302011213231233-1213122131323101-0001111132110300-3023233102320300-2211022221221031"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2021111033332211-2002203201303002-1212300110330213-2232003121110202-0013311101122220-0022021111132122-2303313210023001-1222011023123312"></a>

<a id="canonical-3122233030233201-1323321121111121-1013102322323310-2000102132301302-2112300200002223-3301032201201203-3220311330202230-3031222001021023"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0112012100111010-0330323222132323-3222112132120001-1131113310132013-2113232202302300-2311303223021303-0012222231232022-2003220201302303"></a>

<a id="canonical-0300312120130301-3021002201111013-3303112130103202-0023203203133033-0133201022302132-1210333220321133-1113221123113223-2222010333112000"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0122033132313021-2112210020320110-0133230103123230-0202310221000300-0303030020233330-2303131020232210-3231022113131120-2320000002203112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-025.md#canonical-2220102110322333-3310111320012032-3312031312311011-3311231020200131-1231131001331220-2120133121131110-0032213010002000-0200332111303232)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-2100321321023120-2000323320022000-1033011223020321-3123022130023020-1020110223011330-3110130110121013-2312322210013332-1122313212022222"></a>

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

<a id="canonical-2022121213320132-2202302100330233-3013003002331120-3022132221300003-0202122012331313-0100221320310320-0313133100210020-0130200302121233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-025.md#canonical-2220102110322333-3310111320012032-3312031312311011-3311231020200131-1231131001331220-2120133121131110-0032213010002000-0200332111303232)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-0200300230302122-3331102220201301-0011110331322330-2322232013103333-0021200122221001-0200321010210122-3021311332010011-1222220210031131"></a>

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

<a id="canonical-0121230123023033-0113020323100320-1113313020012111-3101133200332322-2212123302222331-0231012232122132-1210231301111220-0212100312232100"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-1122100212100030-1231031311211202-0321121330213023-0330021210000320-0313000130013202-1332033200133122-0120122031310100-3321302131313132"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3322210321233130-0023133210123131-2022122211110212-1011103003200233-2223111230113011-3302002112113101-2001032320113303-1003130200222303"></a>

<a id="canonical-0121103203120010-1230213111111310-3021301113200011-0101311031201200-0203321111132003-2322212033130120-3000122012202231-1233223020121023"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3031132200120010-3231230002212120-3331312021210231-1310212100300101-0021231012330113-0333121110003032-2313123033013123-0131310233311200"></a>

<a id="canonical-1310310222202020-2012132031222222-1020333203202202-1302210121033003-1203131032303321-3300321032131020-2331322230201311-0000230120001222"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3101332213003202-3112122010322232-2231312220233232-2211003222122222-3302030120012201-0202321003320333-2010023312230013-0211001322023120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-025.md#canonical-2220102110322333-3310111320012032-3312031312311011-3311231020200131-1231131001331220-2120133121131110-0032213010002000-0200332111303232)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-0302001331112320-1303301133000321-2031021030311331-0033212013023012-2202201223220200-2000012330311203-2111210001301121-3001303100033220"></a>

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

<a id="canonical-1220210010013003-0022222302310312-1012013031030103-3101331310333233-1022310302102031-1311201111000001-3331122103313013-3021302312132312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-025.md#canonical-2202113120330223-0103121120023301-3332222222333111-2112113203120113-0331031000100122-2110333300021322-3333230230101301-0211032212120200)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-025.md#canonical-2220102110322333-3310111320012032-3312031312311011-3311231020200131-1231131001331220-2120133121131110-0032213010002000-0200332111303232)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-3332102030032023-3303103311123012-2220111232132100-2110223122213231-3202232011203100-3011113221002122-2230122230101030-1301130202130230"></a>

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

<a id="canonical-0032112030110203-0103032023123300-1333102001033232-3032103310133122-1033213123001131-1313120333132121-1030011300130002-3223011130233221"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-2011121232303221-3312311110010320-3103112133331333-1311123202313030-3132130331030130-0220200032021133-3031332332202223-0021201303310101"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes

<a id="canonical-3020123211003101-2221231012112010-2212121021312331-1023012011323220-0331320301203220-1312021303210130-2300313223103301-3122101220102220"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1210120010301032-2310110002211231-2223033013131213-3100301102231200-0232322203220110-0003120120302303-0102101232331230-1322222202221110"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes`

- [routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301): complete subsection reference.

<a id="canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes

<a id="canonical-2312001001301021-0022131332112222-1130021020213020-3302321331032323-3300220310312123-3302103330102233-3031320223031202-3020001203102022"></a>

Type: `"list"`. Computed.

Routes. Routes for this loadbalancer.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3031111321303031-1110200200130233-1110122212121323-3323323130222023-1022132113321100-0333103000012002-0301130011212022-1210010110202213"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes`

- [custom_route_object](data-sources--workload--reference--group-025.md#canonical-3103102322002232-0021210201020220-3303203222210332-0001030133010303-1300032001131010-3022211131330033-2121321020331220-2313331100202110): complete subsection reference.

- [direct_response_route](data-sources--workload--reference--group-025.md#canonical-1322010312231223-1120130203032332-3022130323003102-3100122202032130-0130310231031322-0201133021131021-2213120113000200-0133131333000032): complete subsection reference.

- [redirect_route](data-sources--workload--reference--group-026.md#canonical-1200110130102320-0121231112030201-3130002232103131-0312211110213031-1221230122133313-3012231113202013-3232233232221320-2103301123023102): complete subsection reference.

- [simple_route](data-sources--workload--reference--group-026.md#canonical-2023011001130113-1300333200123000-3030302321122132-3232120332212030-2310223313003023-3213131323211332-1213202213220213-1203103011331322): complete subsection reference.

<a id="canonical-3103102322002232-0021210201020220-3303203222210332-0001030133010303-1300032001131010-3022211131330033-2121321020331220-2313331100202110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-3130203133130121-3333311202332032-0323123032211331-0220201202120131-3101230222312101-1201300312223032-2201200313200020-0223232130032022"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2312330023303332-3033032000212333-1113120233311103-1012002023323100-3013313012211221-1032232211211103-2221121311011023-1030133130201233"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object`

- [caching_disable](data-sources--workload--reference--group-025.md#canonical-0011132113102123-1202011322123221-0321123333112310-0313131312112100-3312102201303030-1302023211333031-1200132013231023-2021131202122322): complete subsection reference.

- [caching_inherit](data-sources--workload--reference--group-025.md#canonical-2120031231320222-0313123011331023-0013221010122033-3221033103110112-1120003022011333-0301333201310022-1030222120100231-2332123302232121): complete subsection reference.

- [route_ref](data-sources--workload--reference--group-025.md#canonical-0311132032311331-3003210112123211-3100220222223212-3021331301000213-0002020330103233-0310010212100131-0211313103311131-1200120011303033): complete subsection reference.

<a id="canonical-0011132113102123-1202011322123221-0321123333112310-0313131312112100-3312102201303030-1302023211333031-1200132013231023-2021131202122322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-025.md#canonical-3103102322002232-0021210201020220-3303203222210332-0001030133010303-1300032001131010-3022211131330033-2121321020331220-2313331100202110)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-2201301311121020-3300213121221131-0013233202013310-1232223302102311-2222112313033301-3022211223233021-1011120210001331-0200310113131020"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching disable.

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

<a id="canonical-2120031231320222-0313123011331023-0013221010122033-3221033103110112-1120003022011333-0301333201310022-1030222120100231-2332123302232121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-025.md#canonical-3103102322002232-0021210201020220-3303203222210332-0001030133010303-1300032001131010-3022211131330033-2121321020331220-2313331100202110)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-3130332331310112-3323332210003112-1232312100010012-0021100020232120-1123332333112023-2032333133120121-1102131301021210-0223001331300330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching inherit.

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

<a id="canonical-0311132032311331-3003210112123211-3100220222223212-3021331301000213-0002020330103233-0310010212100131-0211313103311131-1200120011303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object](data-sources--workload--reference--group-025.md#canonical-3103102322002232-0021210201020220-3303203222210332-0001030133010303-1300032001131010-3022211131330033-2121321020331220-2313331100202110)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-2033033310201111-2330012231203221-0012103323232210-1200112311010302-1320003231132133-3332102032132103-1222300120223332-2332023031330221"></a>

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

<a id="canonical-0310031330000013-2121032200310122-0000300111221203-1222302301110210-3310212321221213-0333032101002132-0320133200332320-2312220232301022"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref`

<a id="canonical-1030121200301301-0132310012312011-3230010032000102-3010302210003303-1231111331320220-1232022221033132-2120022330232232-1001100000002303"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0102221002302110-0012023000302112-2312331022113033-3330002130221113-0221021231112211-1230303122123020-3302003320310323-2103100301013330"></a>

<a id="canonical-2200102003320221-3320330101102130-1131112211132331-0123021311130101-1010021011300301-2220332123223131-3213030223033302-2013100322230332"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0022020021100202-2100213212223001-1220103222123202-1001102101333101-1023022103112103-0230020222202020-2111122303222312-0233301232121312"></a>

<a id="canonical-2300322130230101-0212022021032332-3130203110301313-0131012311310300-2033322010030233-2023020301110023-1201221330213102-1220210011203302"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1322010312231223-1120130203032332-3022130323003102-3100122202032130-0130310231031322-0201133021131021-2213120113000200-0133131333000032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-3231122232223223-3030333212133102-3311020030310231-1110331232100323-1333201030100132-3210032302200310-1222102201300023-2320132220232120"></a>

Type: `"single"`. Computed.

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1030103002022323-1210113331302103-2233201232233203-0023231230033111-0010112232232002-2133112331113202-3122300220021323-3000313332023311"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route`

- [headers](data-sources--workload--reference--group-025.md#canonical-1011303103102300-1123211203333223-2221123122221012-1220332201010222-1113030201202033-3103133020221303-3001331010220200-2330332132203012): complete subsection reference.

<a id="canonical-1132332301320210-3230032012002032-3322032300113020-1323002000313112-0030022302111000-1112330210320221-1013030101011210-1332201133010133"></a>

<a id="canonical-0303333102102232-1100232010212112-0221023213023233-1331122002030010-1201121002103122-1100120031310001-0210222302301030-3233332322103313"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.http_method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--workload--reference--group-025.md#canonical-1212013302213213-2100300000330222-0301113210100301-1032333123100311-2223212010102000-3300212000103201-0231010300220311-1000120011313232): complete subsection reference.

- [path](data-sources--workload--reference--group-025.md#canonical-2103330332123133-1011130300111023-1312223021200120-0120012100310303-2000023000201231-0231132333113000-3000313021201102-1032101010013331): complete subsection reference.

- [route_direct_response](data-sources--workload--reference--group-025.md#canonical-0102120200303102-3220121212113313-0311311203311303-2211321331021213-2131322113100102-3213200223202023-1233120303013222-0000013120002022): complete subsection reference.

<a id="canonical-1011303103102300-1123211203333223-2221123122221012-1220332201010222-1113030201202033-3103133020221303-3001331010220200-2330332132203012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-025.md#canonical-1322010312231223-1120130203032332-3022130323003102-3100122202032130-0130310231031322-0201133021131021-2213120113000200-0133131333000032)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-0222032202331213-3300323123200100-1002131300332303-2330223131012123-0110303320001312-2000133322223313-2221301330211031-1230301210310230"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 0,
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

<a id="canonical-0100112233212222-3033203111032233-1213303323103311-2221033320012213-3302332203302222-2322123320300103-3230210220212220-2303211220221030"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers`

<a id="canonical-2210130231200232-1333111031102221-3011310323131131-1020202031011132-0000212003200322-3202310010003102-3221023313110200-2213012331120111"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.exact` property

Type: `"string"`. Computed.

Exclusive with \[presence regular expression\] Header value to match exactly.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-0022311212300130-3333021303011230-0300200003203202-0211102200210013-1210200210311031-1122111201202120-3011032011331000-1033332012320230"></a>

<a id="canonical-3323131022232122-1211302301202312-0122122312022013-0322300021102113-0133302313210110-0323232131211320-3100331230232231-0233131010030131"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.invert_match` property

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3211232302232332-2220132220020333-2330133311113230-3123223201013331-2012311011022211-2330222333332101-3120221133200130-3321113303230132"></a>

<a id="canonical-0212121001300023-3312022030110030-3312331031211301-0112002120032313-3110110131001230-2030122101102302-2112301313222220-0120012212023211"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2332212001111320-2031023113213232-3331030113021102-2002001032121312-2330012200101033-2033133020321222-1032103333022321-2313130002003130"></a>

<a id="canonical-1002113101302103-0000213001133113-0330133201201310-0312100103121012-2111012023002112-1233320133303031-3210202223323013-3310103200202022"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.presence` property

Type: `"bool"`. Computed.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200133313022130-2233220233122000-2030033033302101-0100332310022322-2033021213333032-1303332003200333-0321033013012310-0233213221021231"></a>

<a id="canonical-0103310332132201-3023320310322101-0001310000131322-1002230221123102-2301000222112100-0122300101103220-3012331102331313-3300323310230112"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1212013302213213-2100300000330222-0301113210100301-1032333123100311-2223212010102000-3300212000103201-0231010300220311-1000120011313232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-025.md#canonical-1322010312231223-1120130203032332-3022130323003102-3100122202032130-0130310231031322-0201133021131021-2213120113000200-0133131333000032)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-2202312012101321-2313133032013112-1011020233313203-0310020111001333-0303013222311313-3021333212221010-2131222133201212-1310000132330221"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-1131001302111330-2303113303321222-0313222332200303-0212102012321221-0122200000131121-2013110330310322-3203202300013211-0333312302200301"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port`

- [no_port_match](data-sources--workload--reference--group-025.md#canonical-0230330310223010-2203000202333001-1203010221022132-1200022323011331-2312312200001002-0110222112131333-1002322020002102-3203103103131313): complete subsection reference.

<a id="canonical-3310302322203111-1012300303303030-3123130102213201-2231120032321310-0221102030131102-1223232202312322-0011303302311213-1203003112130131"></a>

<a id="canonical-0102033313301011-1303233201303332-0101130311201020-1323102212011022-2201001101112003-3123133033131321-2320330011203001-3001103313120213"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0102033200232130-1222103011213020-2120230210022010-1313303100333020-3211301022122221-2233123122002312-3311211023131202-3112033123232311"></a>

<a id="canonical-3223112121001311-0112233312203031-0302121230132022-2302311202013200-3223330232003333-3013000000010313-2133223110331030-1300122031311203"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-0230330310223010-2203000202333001-1203010221022132-1200022323011331-2312312200001002-0110222112131333-1002322020002102-3203103103131313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-025.md#canonical-1322010312231223-1120130203032332-3022130323003102-3100122202032130-0130310231031322-0201133021131021-2213120113000200-0133131333000032)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](data-sources--workload--reference--group-025.md#canonical-1212013302213213-2100300000330222-0301113210100301-1032333123100311-2223212010102000-3300212000103201-0231010300220311-1000120011313232)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-0320313311113311-1200221112132312-1033011113302110-3331122030103121-2212330232313111-1101030333330203-2001300322220220-3021130230022112"></a>

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

<a id="canonical-2103330332123133-1011130300111023-1312223021200120-0120012100310303-2000023000201231-0231132333113000-3000313021201102-1032101010013331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-025.md#canonical-1322010312231223-1120130203032332-3022130323003102-3100122202032130-0130310231031322-0201133021131021-2213120113000200-0133131333000032)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-2110201133212113-2322032231112320-3323232113001310-3003300001221222-3022132233022020-3213232032121211-0120230210103301-0212133332113312"></a>

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

<a id="canonical-0322001330023220-1300320230122203-1022030311011211-2023312023301321-2101213102122110-2123003210113131-1212300321301122-1320211132332103"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path`

<a id="canonical-0233012021211003-3103213103332111-1230110232032122-0312333220023030-3002223100221112-1200021001312321-3203303121233303-0110012230103220"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path.path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0003321012013331-1111301021120113-3121130020111313-2011020011330002-3011210212322111-0013112210200122-2313330110120322-3033102032213211"></a>

<a id="canonical-3033333031322332-0330202212032330-1200013303302022-3003300121022310-0200313123120220-1211133132212022-0211031123133013-1222220320002003"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path.prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0011222020213112-1102122230120023-0201131012132213-2310100010230103-1302210001103213-0233202230102232-0302021103202012-1231222030113210"></a>

<a id="canonical-1123313332113222-0331131001022211-0101321302200012-1123331123132103-3303331201020202-1023100211032321-0030200311100110-3200131001221301"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path.regex` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0102120200303102-3220121212113313-0311311203311303-2211321331021213-2131322113100102-3213200223202023-1233120303013222-0000013120002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response` properties

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-016.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-016.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.port](data-sources--workload--reference--group-023.md#canonical-2011132303102222-3023301203121330-2233013320233300-0112111313030012-1103313203030222-1002221110321002-3110022320101011-2333330131330303)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--reference--group-023.md#canonical-0122221112123120-0023010121003330-3033210010113330-0100100333002100-0211222302303110-3202021130000322-3023123330231032-2033332332032221)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--reference--group-025.md#canonical-1121003100212032-3332312231032301-0223112200303302-0330012100210020-0203121122320321-2230332022321022-0312211301020331-1321232022223123)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--reference--group-025.md#canonical-1333032133311133-3312033230203310-0200222212221223-0001211003102321-3000032210232113-1213013112112301-1122220230112320-1131011333303301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](data-sources--workload--reference--group-025.md#canonical-1322010312231223-1120130203032332-3022130323003102-3100122202032130-0130310231031322-0201133021131021-2213120113000200-0133131333000032)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-1330100221310102-2112121311002121-2113330311302332-0310101133310232-1212132323230030-1301303210123003-1011222230202302-0320031331310230"></a>

Type: `"single"`. Computed.

Send this direct response in case of route match action is direct response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1332130210013002-3031322000211023-1021100102233130-3313333333323320-2110101111301121-3311000311002311-3330221132213322-2303002232121222"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response`

<a id="canonical-1201312300330110-2232001323313031-3323223011300123-1020213000112020-1012100200011011-3333200232002200-3320023220003022-3313331020223201"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_body_encoded` property

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2110213012220033-3120233122000303-0003121330133201-2223230103021113-2212302220311132-1233133021302132-2113010212112121-1100131231103020"></a>

<a id="canonical-3022303012020131-0110201031131231-0233232111103012-1002311321312003-3322112113000230-1321003113312300-3030221201100033-0300030302222330"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response.response_code` property

Type: `"number"`. Computed.

Response Code. Response code to send.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```
