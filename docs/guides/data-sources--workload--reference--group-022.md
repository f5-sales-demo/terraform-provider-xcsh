---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0001310231113010-2300222132331331-1202322233202001-3132330301222132-0023333212323030-0221102011110130-0122220211010300-2331201130231301"></a>

## namespace property — trusted_ca / 012001310310 / 5

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

<a id="canonical-0320131211121030-0231301311123301-2321121201301031-2103332310203211-2002131233310110-0113233130313203-0322113013201111-3202321211133003"></a>

<a id="canonical-3321011233221212-1201002032013112-0302212112203230-0312021322103231-1002332110030022-3011012023200131-1012033320311322-1321300331200203"></a>

## tenant property — trusted_ca / 012001310310 / 6

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

<a id="canonical-2113003322330323-1332201003313220-1301021113001130-1320300101312012-3112222120223120-1003021001323123-1103310101030211-3030023032203222"></a>

## Next pages — trusted_ca / 012001310310 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0030111210231013-2222220121122130-2010310023010310-1130132313031002-1330222301322013-1210011310321303-2232121032200102-2310010213332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110210021122010-3131213031231131-1001222213221302-0132031330322332-2113321332032311-0002013313013030-3002010101103120-0001200020233302"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 310221202223 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-3203312331230212-0022231112012232-0201200221212231-0211111031100113-1020313300212121-3231110213333133-1003132331212301-0331122322200312"></a>

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

<a id="canonical-2132202011000201-2021010231100120-0030130132223200-2203133210101313-2200332033311003-0230023311221022-2103111002000313-1013330223230311"></a>

## Direct properties — xfcc_disabled / 310221202223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000221322022310-1113021132100210-1010201123321012-3003002331021310-3003130131000233-1020220232002333-3010113213111322-0303121021100102"></a>

## Next pages — xfcc_disabled / 310221202223 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3221100202303211-1330113112311332-1132323132323111-2020203132010123-0100312032302132-3010031232110212-2323021103032020-2102132331230010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022111030003201-2031132021213312-3021313331131313-2122100322023100-0220221303113113-2332032222123110-1000200102233320-3323100231212000"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 120103331322 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-021.md#canonical-0020221002300310-3320210022331231-2110020220012113-3303311310321122-1023231100033031-3133012031200203-1002323202122123-3323330220103313)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-3033003333232321-1322001101201210-0020223321032031-3302033013002231-1013112203123300-2302110002301322-0131032032110323-0213010222331110"></a>

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

<a id="canonical-3132333300221013-1300300323202122-1333330032031323-1032223333131331-1200331203213323-3201212231011313-3000001332120110-3110100202112202"></a>

## Direct properties — xfcc_options / 120103331322 / 3

<a id="canonical-0031322011101203-0312300020230301-3002131012031120-1201310331301030-1201100312100211-0121132233302330-0100110123111101-1020333201231202"></a>

<a id="canonical-3333000112110221-3010032202112213-1003120322113003-3300031333330001-0012203232221120-0213031323013213-2303000032331233-1011202101331002"></a>

## xfcc_header_elements property — xfcc_options / 120103331322 / 4

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

<a id="canonical-3130223122223012-2201031001301111-1330223222202313-3022312012313210-2113232100210211-3312220201110200-1132011100123320-0220303302203022"></a>

## Next pages — xfcc_options / 120103331322 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](data-sources--workload--reference--group-021.md#canonical-2120202120230123-1300020120013231-2032223101233223-1023310302320101-2320320032132221-0031023101132320-2230232211332312-0332131011020102)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101131201103221-3023001230203323-3022023022000033-1302032213132230-1011011133200320-3011022211312300-3210202311002303-0210130230022113"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters — tls_parameters / 121320313102 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-0103223033033133-2133130101020020-3033101100330100-3323031133022120-3113110221332132-3023032211123300-0012111230333020-3010133030032332"></a>

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

<a id="canonical-3001312212132201-1133200200332121-0331130333010022-1222230110012323-0033130032002212-0133113003202303-2231331012332000-3110101232111021"></a>

## Direct properties — tls_parameters / 121320313102 / 3

- [no_mtls](data-sources--workload--reference--group-022.md#canonical-3231200111230001-0232113102222133-3103130022103033-3000222111220333-2002100110103310-2313323033112231-3312011302333313-2030332210211321): complete subsection reference.

- [tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001): complete subsection reference.

- [tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322): complete subsection reference.

<a id="canonical-3232101223111101-2023013322312123-3221322120311132-2130013102012133-0331300001101312-0021332201032222-2123113302123212-0002020111131101"></a>

## Next pages — tls_parameters / 121320313102 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls](data-sources--workload--reference--group-022.md#canonical-3231200111230001-0232113102222133-3103130022103033-3000222111220333-2002100110103310-2313323033112231-3312011302333313-2030332210211321)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3231200111230001-0232113102222133-3103130022103033-3000222111220333-2002100110103310-2313323033112231-3312011302333313-2030332210211321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200023310101310-0301333211012120-1222200003321202-0333020133020120-0013113222030112-1101312312011011-0211312222312210-3112202201320013"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls — no_mtls / 002303330103 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-3203111113333010-2001103000233223-3212233233122122-0110033120133103-2132032232212131-3001110110303101-0232112322012101-0303230130120100"></a>

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

<a id="canonical-2321223322302210-0102231000000130-0221200310221133-0013000011101023-1220201100013200-3131203003300212-3033222023103103-2133212310122330"></a>

## Direct properties — no_mtls / 002303330103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110212003200031-1203210011210021-0333010030212231-1310322212203033-1211210122123221-1223022122232110-3202013331131203-0212231003333021"></a>

## Next pages — no_mtls / 002303330103 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220001031213312-3020200120013200-3310021201233320-0202320301203130-3323002211212120-2301133130031221-3201120331012113-0100131230021301"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates — tls_certificates / 300032322331 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-3003313211202113-2021100011112331-2013113002121110-1230132120122202-1001333111220123-0130103022032000-2012030320023012-0222011311030003"></a>

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

<a id="canonical-0330200022230010-3230133011212210-2032223103002203-3122232233212220-3313210111023031-2100300101101301-3111110232212310-1010131101031200"></a>

## Direct properties — tls_certificates / 300032322331 / 3

<a id="canonical-0021320233012323-1220012022030003-0331031132103220-3022022033230312-3111311003020110-2001221131131231-2320330301202330-3133233100032100"></a>

<a id="canonical-3113300131021001-0311323101230232-1011300113023221-2220220210233311-1220230000233330-2103321013201120-1212310332032023-1300021230232321"></a>

## certificate_url property — tls_certificates / 300032322331 / 4

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

- [custom_hash_algorithms](data-sources--workload--reference--group-022.md#canonical-0232323111210030-2222131202223002-1211111033031030-3032033203332223-0312232001123030-0222030200031310-1123300111221303-0312303221022212): complete subsection reference.

<a id="canonical-0230200013123230-2303103233120212-1013322212233010-0113331221120210-3301302011202101-3112311111312030-1133323212301023-1323022323210020"></a>

<a id="canonical-1223233202020330-0103200220203332-2210003312201233-3102010213001133-2311313031301311-2100301302313110-1212020033021210-3003312020222103"></a>

## description_spec property — tls_certificates / 300032322331 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--workload--reference--group-022.md#canonical-1101301113313023-1230011023121231-0220312222031231-1100013220330120-1003301213110123-2022202213001332-0313001302220030-2332211303000222): complete subsection reference.

- [private_key](data-sources--workload--reference--group-022.md#canonical-2000300020301313-0322332230111120-0301330331223200-3201132103312013-3200033102320123-1213320002213032-3123220302222201-3220313111200022): complete subsection reference.

- [use_system_defaults](data-sources--workload--reference--group-022.md#canonical-0233103130232303-0221321110230320-2111122030231332-1221212230221303-3221122020101003-2032122221233310-0132331130032110-2211011120331321): complete subsection reference.

<a id="canonical-2221000201013232-2122301211300102-0303231001122012-1312211121121223-1220303323102322-2321130133220231-2331201102312100-2311310323121032"></a>

## Next pages — tls_certificates / 300032322331 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--workload--reference--group-022.md#canonical-0232323111210030-2222131202223002-1211111033031030-3032033203332223-0312232001123030-0222030200031310-1123300111221303-0312303221022212)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--workload--reference--group-022.md#canonical-1101301113313023-1230011023121231-0220312222031231-1100013220330120-1003301213110123-2022202213001332-0313001302220030-2332211303000222)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-022.md#canonical-2000300020301313-0322332230111120-0301330331223200-3201132103312013-3200033102320123-1213320002213032-3123220302222201-3220313111200022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](data-sources--workload--reference--group-022.md#canonical-0233103130232303-0221321110230320-2111122030231332-1221212230221303-3221122020101003-2032122221233310-0132331130032110-2211011120331321)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0232323111210030-2222131202223002-1211111033031030-3032033203332223-0312232001123030-0222030200031310-1123300111221303-0312303221022212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120332003233031-2013111230013132-2321113132320231-3210313320323321-3101323122013330-2131302123113100-1101011123121031-3201113013221321"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 330123200002 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-0331100012101331-2211133103020003-0220130103213332-2023000311022321-3313201333200010-3010223111202102-2200100010312211-2200003301323300"></a>

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

<a id="canonical-3020212220320033-0220320010021130-0103221000033122-1203321130131300-2112322112213113-0012003300011331-2203121012322032-1232100122112101"></a>

## Direct properties — custom_hash_algorithms / 330123200002 / 3

<a id="canonical-2100112101203002-0303321001103002-1002103103121332-0012330100030332-0222002030221100-2100302332230312-1031133021112111-3122003010111230"></a>

<a id="canonical-3022133121121303-3331030033011200-1000110223302333-3120033131130013-1133312230321331-2013232031330212-0210102211103030-2201120332032122"></a>

## hash_algorithms property — custom_hash_algorithms / 330123200002 / 4

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

<a id="canonical-2223100122310000-3320300030321312-2302333030023132-0223231220210331-1020213001322221-3030011103120322-1313010112022003-1310230123111212"></a>

## Next pages — custom_hash_algorithms / 330123200002 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1101301113313023-1230011023121231-0220312222031231-1100013220330120-1003301213110123-2022202213001332-0313001302220030-2332211303000222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000312113103211-3112201000010330-1220322133101022-2221311222122231-1213133011320102-2122011013332121-3321332202023330-3331032233202012"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 223200300121 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-2223311031031320-2320102303003323-2232112022023123-3320222100233132-3111223322123311-1311032100111132-2102232032222023-1313202112223112"></a>

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

<a id="canonical-0021030203321121-1121233322012030-1110001002231212-2010100231033011-3311013200012211-1220331133300132-1020210221201112-3032222100232333"></a>

## Direct properties — disable_ocsp_stapling / 223200300121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020203030022331-2232111201232221-2033222103330113-3020330231222001-1231203121032111-1010230101130232-2110112110202210-0021220321011321"></a>

## Next pages — disable_ocsp_stapling / 223200300121 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2000300020301313-0322332230111120-0301330331223200-3201132103312013-3200033102320123-1213320002213032-3123220302222201-3220313111200022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021310211230011-3330223133131033-2213103212020333-2210122320002110-0320112033222330-1231022212132213-0013031002122101-0323211210222230"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — private_key / 131100131102 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-2210122223330221-3331012122201003-1221231113232132-2132130120131203-1213320112322333-3023101321313130-1331031210131101-1013333330312010"></a>

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

<a id="canonical-2311122303203032-3202211003220010-0210121220102220-0011330212311303-1031202132121123-0010300331111310-0201222103213102-2233133232302101"></a>

## Direct properties — private_key / 131100131102 / 3

- [blindfold_secret_info](data-sources--workload--reference--group-022.md#canonical-3030313223210003-0033301220330230-1003132301012121-0102312101313033-3211212133302000-1313001103132033-0010113112332203-2300321312203330): complete subsection reference.

- [clear_secret_info](data-sources--workload--reference--group-022.md#canonical-2201012012103103-2032210223330200-3023213332221130-1002110130103222-0013222322233031-3000020202220333-1301332020002223-3302322213000211): complete subsection reference.

<a id="canonical-0312101013120312-2303331223331320-2102031212022103-1133302203102210-2022223113022202-0030033101303030-3230002031001322-3030000231300003"></a>

## Next pages — private_key / 131100131102 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--workload--reference--group-022.md#canonical-3030313223210003-0033301220330230-1003132301012121-0102312101313033-3211212133302000-1313001103132033-0010113112332203-2300321312203330)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--workload--reference--group-022.md#canonical-2201012012103103-2032210223330200-3023213332221130-1002110130103222-0013222322233031-3000020202220333-1301332020002223-3302322213000211)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3030313223210003-0033301220330230-1003132301012121-0102312101313033-3211212133302000-1313001103132033-0010113112332203-2300321312203330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022131010020230-1031033003321302-1030223230200211-2313013212103131-2031023222201022-1101003203120202-1332100211031012-1203210213221200"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 113220032110 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-022.md#canonical-2000300020301313-0322332230111120-0301330331223200-3201132103312013-3200033102320123-1213320002213032-3123220302222201-3220313111200022)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0101331023123223-0303103333102132-1132031121121100-2013231130331231-3023110033303003-1322002032321220-1202303031031300-0203311201021022"></a>

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

<a id="canonical-3210333320131210-0102233303103321-2102332112033321-3321213221223021-1203131303310321-1031322010203032-3232213130231122-3012103311322200"></a>

## Direct properties — blindfold_secret_info / 113220032110 / 3

<a id="canonical-0221302202202311-2312021112030122-1033320110003031-1313323212201022-3132033121000203-3201321300012222-1010022032322121-0020130321231031"></a>

<a id="canonical-3003323111221211-3133120010102112-2330321033303210-2033103323221302-3033322230223302-3323302200200100-2132013212230320-0300231322102333"></a>

## decryption_provider property — blindfold_secret_info / 113220032110 / 4

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

<a id="canonical-3322012211223222-1210111022020221-1030230233020032-0232332131233333-3220221311133330-1030103300122030-0311133113010023-1322233031313332"></a>

<a id="canonical-2032121301110212-1100201030201020-2221032302022020-3311011122320203-0333000322110101-2101033201101011-3133021212032013-1212013222201013"></a>

## location property — blindfold_secret_info / 113220032110 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-0320303220330002-1312233031023110-2330032231000301-0002311112100332-1220332122000113-3311133023313112-2323103113132333-1113213223311012"></a>

<a id="canonical-0021310303232300-3322231333331003-3200312313010300-2322322233223132-3130122203213323-1001211112201030-0023003321013000-1132211122100002"></a>

## store_provider property — blindfold_secret_info / 113220032110 / 6

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

<a id="canonical-1330221021122233-2000232010100223-0233130212121003-1231300130112321-0013123303220120-2103113201023211-2223102120022212-1031321202012310"></a>

## Next pages — blindfold_secret_info / 113220032110 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-022.md#canonical-2000300020301313-0322332230111120-0301330331223200-3201132103312013-3200033102320123-1213320002213032-3123220302222201-3220313111200022)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2201012012103103-2032210223330200-3023213332221130-1002110130103222-0013222322233031-3000020202220333-1301332020002223-3302322213000211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332112100012302-2213211132032113-0122110213303020-2313012130000121-1000120200221202-1312333323030121-2302120102032332-3231103312123301"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 212231132202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-022.md#canonical-2000300020301313-0322332230111120-0301330331223200-3201132103312013-3200033102320123-1213320002213032-3123220302222201-3220313111200022)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-0102311301131101-3011313100302001-2110103021301311-1022311113010333-3222230310220130-1200221220023122-1332102120121221-1120120321231330"></a>

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

<a id="canonical-1202231130123221-2100231112313033-0110013130003301-2330101131213101-0300201023200131-1212133232333122-3210322303202231-2322100302320213"></a>

## Direct properties — clear_secret_info / 212231132202 / 3

<a id="canonical-3203021113110202-2120031113232232-3122011022031330-2031131302020323-0211201202201302-0201212311003230-0300311221000030-1001133211222233"></a>

<a id="canonical-3130333302212122-1132021122110310-2030111331031130-0232302300322312-0022002322202012-3201010201010132-2012232120013323-2323112113002303"></a>

## provider_ref property — clear_secret_info / 212231132202 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3023112100132300-1231001321333001-1011010003331002-1223313111100003-3220322200010332-3303002300322132-0300321011300101-0103131023103102"></a>

<a id="canonical-0131300211232200-3103020301333022-1011130011021333-3023123302103131-0320130023300112-2111003203302120-3100313111301013-1110230200313221"></a>

## URL property — clear_secret_info / 212231132202 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-2213010132202112-2201233130300023-1322130300332331-0033211020131000-1212301110331133-3013102002132132-0223123133002330-2333013331222303"></a>

## Next pages — clear_secret_info / 212231132202 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](data-sources--workload--reference--group-022.md#canonical-2000300020301313-0322332230111120-0301330331223200-3201132103312013-3200033102320123-1213320002213032-3123220302222201-3220313111200022)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0233103130232303-0221321110230320-2111122030231332-1221212230221303-3221122020101003-2032122221233310-0132331130032110-2211011120331321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101230232033033-3220020212321030-2301030123303131-1303103223111332-0022001001213211-0121021133031200-3302010210012332-1033301032111200"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 230011220113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-1121022301303133-1003001200203111-0133031202200122-3321013302131120-0131212333223130-2133321113023203-2123121332302031-3021313112223203"></a>

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

<a id="canonical-3112311102230201-3110300203302312-3000320120023103-2000311200023123-0323113323113010-2021001132021221-2031300330130011-0022322213102012"></a>

## Direct properties — use_system_defaults / 230011220113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123212011222220-0023222220100011-1110122000210032-2010131010002013-3102300230202000-0002313020332213-1231210223331132-2003012002103112"></a>

## Next pages — use_system_defaults / 230011220113 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](data-sources--workload--reference--group-022.md#canonical-2303102321200322-2223331031221031-0020113022300133-0320303032102020-2302213222211223-0230212333332330-2211002133222000-3201213130120001)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211232110200112-0032310013021111-0200113003112322-0302231330012033-0033112220020331-0202023333000122-1320322113200100-1202213223100201"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config — tls_config / 323111112311 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-0211200320101132-1223312110111222-2021323102000113-2000113212211301-0101232002003011-1212312132211330-0222023132200301-2211200220130013"></a>

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

<a id="canonical-3130311111220033-2123100310120000-0333300322330220-3012202133133300-0131113123211103-3123201331313011-2220333222302200-1011121203223212"></a>

## Direct properties — tls_config / 323111112311 / 3

- [custom_security](data-sources--workload--reference--group-022.md#canonical-2332213000011030-1321022302023322-2323312323211221-0332201321023300-3133020101230333-2331220111100201-0122301132201122-2133212130100220): complete subsection reference.

- [default_security](data-sources--workload--reference--group-022.md#canonical-1220230112020013-1033230003322103-0101003222112232-3220022221022230-0302020013133203-2310303101032220-2100201003333220-0331123003121221): complete subsection reference.

- [low_security](data-sources--workload--reference--group-022.md#canonical-1311013302002031-1111003303111231-0332030020020121-3132203323333033-1320031023100030-0212020303123103-3330301230332201-2022332032200232): complete subsection reference.

- [medium_security](data-sources--workload--reference--group-022.md#canonical-3030033000021223-3111020010232202-3201313030123120-0300333312120333-3301222030031012-2133333111322021-3001313131113123-0120332013132213): complete subsection reference.

<a id="canonical-2111200221320230-1213220112002121-1111000212230333-1021232111132222-3310120333022202-3010201231102312-2201232322033222-3232320112003301"></a>

## Next pages — tls_config / 323111112311 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security](data-sources--workload--reference--group-022.md#canonical-2332213000011030-1321022302023322-2323312323211221-0332201321023300-3133020101230333-2331220111100201-0122301132201122-2133212130100220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security](data-sources--workload--reference--group-022.md#canonical-1220230112020013-1033230003322103-0101003222112232-3220022221022230-0302020013133203-2310303101032220-2100201003333220-0331123003121221)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security](data-sources--workload--reference--group-022.md#canonical-1311013302002031-1111003303111231-0332030020020121-3132203323333033-1320031023100030-0212020303123103-3330301230332201-2022332032200232)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security](data-sources--workload--reference--group-022.md#canonical-3030033000021223-3111020010232202-3201313030123120-0300333312120333-3301222030031012-2133333111322021-3001313131113123-0120332013132213)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2332213000011030-1321022302023322-2323312323211221-0332201321023300-3133020101230333-2331220111100201-0122301132201122-2133212130100220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223031322221113-3012121321200300-3013122332300322-3102323231302102-0323230023233113-2020111132133023-2033213120113030-1231221101201033"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security — custom_security / 333332210331 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-1222302312111131-3313213010120322-0031022210233020-0013102233230100-3323222320023021-2101021011111020-0230331313111212-3312310231030130"></a>

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

<a id="canonical-2302112031233311-2121011322001311-2332301301311320-2332313112203331-0212200221203313-3332032321322320-1201303313331231-2311310330200312"></a>

## Direct properties — custom_security / 333332210331 / 3

<a id="canonical-2031232201000120-0313022201030211-3302133113201103-3331023111330222-2200312011233222-3021130233010212-1330112201001113-2003031222021011"></a>

<a id="canonical-0302120201030033-3100212001012031-1330132101131312-3230113203202230-2222021231221002-1330212110012102-1321333110202222-1212212210023010"></a>

## cipher_suites property — custom_security / 333332210331 / 4

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

<a id="canonical-0000210232110132-1123113122111322-1312022223002013-0022231203002030-2323113223122310-1110020332012232-2310102233130203-1223220310010113"></a>

<a id="canonical-3130130320012103-1213201223120010-0232220120230100-0112001220023101-2122301013323012-2230212131113031-2103132001230000-2202110303300310"></a>

## max_version property — custom_security / 333332210331 / 5

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

<a id="canonical-1003330020300021-1303132300231301-0312112232210001-2101002021332313-0200202103331202-3013112222130000-3102111122332231-1013030112122312"></a>

<a id="canonical-2002120330113131-0230223211221010-3210200303021310-1310101332112022-0323122103122322-3211013111013222-0001313001020211-0013012131322013"></a>

## min_version property — custom_security / 333332210331 / 6

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

<a id="canonical-1301331231301123-3210332011113211-1130231130113013-3302030022113201-2033023020211223-3233230330211330-0223102031300012-0033303001031200"></a>

## Next pages — custom_security / 333332210331 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1220230112020013-1033230003322103-0101003222112232-3220022221022230-0302020013133203-2310303101032220-2100201003333220-0331123003121221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011213330210113-2032300121013113-2030323121321003-2313020003232213-0210233032311321-2322210111021303-2123101303133031-1013100113133302"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security — default_security / 322320003330 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-1102203200131231-3223000331121002-3233023220012030-2101133311200101-2230113302100200-0113302331013300-2130010133110113-1032200221012121"></a>

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

<a id="canonical-2112030033210313-3033223121331200-2111003213221230-0223110331110103-2031302122100122-0331102232023303-3231202031023001-0213103312231311"></a>

## Direct properties — default_security / 322320003330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303330330032211-2002122121023003-1200113212213001-0223002002000311-2112003032213030-3202130112100003-1032331223212033-2102021312310110"></a>

## Next pages — default_security / 322320003330 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1311013302002031-1111003303111231-0332030020020121-3132203323333033-1320031023100030-0212020303123103-3330301230332201-2022332032200232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023210333002032-2112001302032300-0303311003001300-3311311213110131-0320101322113031-0120230333100222-1213221223010113-0230212032331011"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security — low_security / 022023303131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-1113201312032123-3310122021111020-2210230121300200-1331310100132202-2020130301310331-3123202021202120-1111220303013020-1313000012132200"></a>

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

<a id="canonical-1201003302023323-3203210320121300-3331113101000122-2123023033213330-2011202130231133-2320230220111122-1320233312223230-0132202311302002"></a>

## Direct properties — low_security / 022023303131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133033023233110-3310002021220212-2102322100221300-2130012011112223-0102110133321322-3210122021023020-0033003333223331-0312002032232322"></a>

## Next pages — low_security / 022023303131 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3030033000021223-3111020010232202-3201313030123120-0300333312120333-3301222030031012-2133333111322021-3001313131113123-0120332013132213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103112010200011-3203323123113213-0332331131022303-1333311022301101-1330223323223033-0011331223110301-2001103003200113-1111131012332132"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security — medium_security / 001230322203 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-1101221010332323-0001201012113323-1230232211020030-2303332213202032-1033233230131030-3211203320301221-1131103322211210-0022300002231130"></a>

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

<a id="canonical-2213110221322202-2232122223310222-0101322031101133-3020122012123201-2032033131133102-0031131120311100-3322312030212311-0022303211122310"></a>

## Direct properties — medium_security / 001230322203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232202213033120-3101133102123110-2323302130332122-0102001323130211-3012110203302113-0113031222131321-2302330131000112-0323230021310330"></a>

## Next pages — medium_security / 001230322203 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](data-sources--workload--reference--group-022.md#canonical-3112322103001202-3322033233311231-1102030211303311-3012000113330002-2313123200311203-2131113012330010-2213311202121332-1221333001020332)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233012113331131-0021332331331121-2201030302221110-0020320111013301-1033100033022231-2030112130110000-1322332011102112-1330030321300200"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls — use_mtls / 213330201203 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-3031220120320232-1101223321112101-3113103011322131-2230110013110330-3133112130133312-2322333201110020-1231032013012332-2011110131031330"></a>

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

<a id="canonical-3332121230212230-3232110000323010-3311102213233330-2213300223321011-1110320231303310-1002100300133230-1300120000313011-3003330230013230"></a>

## Direct properties — use_mtls / 213330201203 / 3

<a id="canonical-3023021022123222-3123020303302112-0012220010132003-1312013332112031-0322313313322222-2010023130021330-2323210113222012-3322111022230033"></a>

<a id="canonical-1121212330233113-3212130330231213-0222202231320033-0003202111330113-2032200302113023-0310330203020111-3221320123231110-2223101133132300"></a>

## client_certificate_optional property — use_mtls / 213330201203 / 4

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

- [crl](data-sources--workload--reference--group-022.md#canonical-1302332033021103-3230023212103203-0022210021101233-1011213131220332-0201032001102310-2203122210202111-0223220303201000-2202220220312132): complete subsection reference.

- [no_crl](data-sources--workload--reference--group-022.md#canonical-0303030313123000-1010330313030201-2210110301230010-3123032301212033-2113202121330010-3213011012102032-2023023323220213-2103101003003031): complete subsection reference.

- [trusted_ca](data-sources--workload--reference--group-022.md#canonical-0202320210323033-3300232032000321-1122010201201323-0133012012103120-3230201322021031-3300120121123312-1332000102212300-3021111202102023): complete subsection reference.

<a id="canonical-2130101312211222-1201001000313332-2303031022303121-1101230330120313-3110220233130221-2002312300021032-3021210032330300-3312130112102010"></a>

<a id="canonical-2100000231123131-1011212331223021-2220130100122330-2202323033331000-1011120032233103-1203030220103110-2210011033303123-2000221110030300"></a>

## trusted_ca_url property — use_mtls / 213330201203 / 5

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

- [xfcc_disabled](data-sources--workload--reference--group-022.md#canonical-0330211331233310-0002001003100020-2110000130201300-2233013100133030-0233321122222110-2130232233003202-0012023233303232-3302310232331003): complete subsection reference.

- [xfcc_options](data-sources--workload--reference--group-022.md#canonical-1031303310333110-0222030123111010-3030322013000110-2330311321110102-3101220021030011-1020103113331133-2130221113003221-2103321110120221): complete subsection reference.

<a id="canonical-2122312110223131-1231110303010132-3133030223010133-3013002322121031-0312231330322012-1330210200000320-2110222130221100-2200001211112300"></a>

## Next pages — use_mtls / 213330201203 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](data-sources--workload--reference--group-022.md#canonical-1302332033021103-3230023212103203-0022210021101233-1011213131220332-0201032001102310-2203122210202111-0223220303201000-2202220220312132)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](data-sources--workload--reference--group-022.md#canonical-0303030313123000-1010330313030201-2210110301230010-3123032301212033-2113202121330010-3213011012102032-2023023323220213-2103101003003031)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](data-sources--workload--reference--group-022.md#canonical-0202320210323033-3300232032000321-1122010201201323-0133012012103120-3230201322021031-3300120121123312-1332000102212300-3021111202102023)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--workload--reference--group-022.md#canonical-0330211331233310-0002001003100020-2110000130201300-2233013100133030-0233321122222110-2130232233003202-0012023233303232-3302310232331003)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](data-sources--workload--reference--group-022.md#canonical-1031303310333110-0222030123111010-3030322013000110-2330311321110102-3101220021030011-1020103113331133-2130221113003221-2103321110120221)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1302332033021103-3230023212103203-0022210021101233-1011213131220332-0201032001102310-2203122210202111-0223220303201000-2202220220312132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030103321320030-2133103110303321-3113013023221300-2201222222211003-2202312103131011-3213120313313102-3201131022333300-3122013013002211"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — crl / 323003133131 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-0011130030031132-0312023033320223-0101223113000230-0213023120011322-3021103031300302-3032132232101132-3332121200101032-1102030221211311"></a>

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

<a id="canonical-2200320220130301-1000033220032022-2230123013231122-2211021301320212-1001023223103130-1222102331230131-1221331311220033-2033032113210100"></a>

## Direct properties — crl / 323003133131 / 3

<a id="canonical-0000022300323222-0332321110003232-1332231302230332-2020200320230310-0312022111103330-1030012022332231-1120310011023001-2230020211221032"></a>

<a id="canonical-1331202231323122-3131032131332230-0032031233021102-1110233321233033-1333221001132201-2031200001213120-2231131210102202-1220301333130031"></a>

## name property — crl / 323003133131 / 4

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

<a id="canonical-1111102111130000-0301131103120230-2210011122221130-2212022321022032-2012303011211131-3330000222233320-3123212322322012-0010100223322022"></a>

<a id="canonical-1020102132233322-1311320013003202-3013213102322223-1231211130133301-1032113123020330-3123123313311332-3110330320221330-0021110303133200"></a>

## namespace property — crl / 323003133131 / 5

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

<a id="canonical-1122223330003110-3130311232313022-2211223322011330-1331020223023100-0011310300021022-1323110021331213-0021010001320032-1111023221001332"></a>

<a id="canonical-3120331121031302-0232333213011322-0313322333210323-1101031032200300-1230102331012032-1212302202203032-2001213011233332-2110300023212311"></a>

## tenant property — crl / 323003133131 / 6

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

<a id="canonical-1112112330200033-0220030320112022-1311100111222330-3030010330030221-3212112203220122-3000031223100310-1231120311131033-1122223210120203"></a>

## Next pages — crl / 323003133131 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0303030313123000-1010330313030201-2210110301230010-3123032301212033-2113202121330010-3213011012102032-2023023323220213-2103101003003031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000003203010003-3132322202330332-2233010323222223-3021321012023001-2222003023212021-3211223230032330-2210330330111002-0122213220020322"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — no_crl / 213200311330 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-1120012122332331-3232123101112130-1333111221003113-2000310111111220-1223321210020030-3310020321211202-3032201122113213-2001213133211122"></a>

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

<a id="canonical-2010122130231213-0220331220003121-3101201310023313-0103101232123121-2332132231210013-0333203201200333-3012311222030300-1021032023012130"></a>

## Direct properties — no_crl / 213200311330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010101112001022-3130321001332101-2311112302110100-1012220332121021-3010311131132022-3132221222312203-1000302200220030-1312011322302013"></a>

## Next pages — no_crl / 213200311330 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0202320210323033-3300232032000321-1122010201201323-0133012012103120-3230201322021031-3300120121123312-1332000102212300-3021111202102023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010300213302221-2030121001322232-1030012123002032-0122103231212332-1013013210300132-2120030122301330-1112101322202031-3203002002000122"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 330310300021 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1311223223111020-3111310203202232-0030111030310111-2122103021112202-0020322202101302-3122323112223230-2113321133133332-1002121022122200"></a>

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

<a id="canonical-1200231102013103-2221101030003022-1323311300311011-3212312033122100-0212220111220031-2212001332320331-2301332202122022-3322213202130022"></a>

## Direct properties — trusted_ca / 330310300021 / 3

<a id="canonical-1000312122322122-2002232111032102-1001322230302013-1200223111131312-3233202021130312-2212210210003331-2100111103122102-2331301001123103"></a>

<a id="canonical-3333220203013122-1031220323202011-3310300130010101-0012103121003200-0033020030212011-0121312010320221-3321213320021230-1220101203201223"></a>

## name property — trusted_ca / 330310300021 / 4

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

<a id="canonical-2123000122031033-1232132313032013-1300110033103132-0123030302132013-1300323202211303-3300232311112213-3303130203032303-1000311233321132"></a>

<a id="canonical-3012102322123103-3221010001120233-1023222102323102-3301120203022032-0321111203131310-2203103111030132-2002023203231000-3131033110323320"></a>

## namespace property — trusted_ca / 330310300021 / 5

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

<a id="canonical-0101321131230233-0001011103322101-2302210122100331-2110111110213100-2112201200231121-3111030221220113-3032311003021321-0131102330211011"></a>

<a id="canonical-1021020333332110-1021102122133222-0130313332100012-3120231132110222-3010002020223232-1003231212233111-3033130220202121-1113033232121131"></a>

## tenant property — trusted_ca / 330310300021 / 6

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

<a id="canonical-3111113321322103-1322122112221232-3112001121011123-2213323010001202-2120302123330033-1113302103112132-2303201002102030-3130110022213021"></a>

## Next pages — trusted_ca / 330310300021 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0330211331233310-0002001003100020-2110000130201300-2233013100133030-0233321122222110-2130232233003202-0012023233303232-3302310232331003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332232022211012-0002033003302330-2000003131023010-0000103313101332-2102313020310200-3220220330211020-1223130003122010-0131331233310012"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 303311103221 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-0303021300132123-1010130103100232-2030121021201000-0230333303101023-1032020310200123-3313323012022120-2203113203310201-1102201023223022"></a>

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

<a id="canonical-0130302131021011-3120301231221002-3030002313323103-2330123310202132-2220122013023110-0333302120230223-1210212310213020-0320213210311310"></a>

## Direct properties — xfcc_disabled / 303311103221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022202121010230-2033331003022113-1221100312120233-0120022320230002-2330222310111031-3013012323302020-0231133212311112-0231023031323120"></a>

## Next pages — xfcc_disabled / 303311103221 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1031303310333110-0222030123111010-3030322013000110-2330311321110102-3101220021030011-1020103113331133-2130221113003221-2103321110120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102002031203330-3233301223230111-2222033122330032-1312200203321311-2323030011032113-3331233022020330-2202320222313030-3212000030010200"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 112100200112 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](data-sources--workload--reference--group-021.md#canonical-2102010120332203-3300130310311113-2133023220220121-0222312012200312-2020311331222300-1021332202132301-0330212222120011-3302213111012231)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-022.md#canonical-0010123223300323-3122033232120131-1033211323302101-0202012323131002-0013101320130212-1220113010330021-1032021212320311-0300031310102303)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-1321020012213202-2233310220231122-0103331331212021-1032020222220013-0020011233101213-2211300123102230-3313011132312011-0100302323132130"></a>

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

<a id="canonical-0233203201213032-0100212202300331-2103213121202330-1211110210023100-0110133221003030-0103130023111221-0031312033011312-1132220122133300"></a>

## Direct properties — xfcc_options / 112100200112 / 3

<a id="canonical-1233033201220231-1022112223123101-0101220122300130-3020202013332032-0321110232202202-3202211230021320-2300202320021010-0232211101023132"></a>

<a id="canonical-0133213212213323-3332213331032330-2111230232130331-2213331233031220-1222012110013013-0333302030230220-1013030022220313-1011313030321110"></a>

## xfcc_header_elements property — xfcc_options / 112100200112 / 4

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

<a id="canonical-1020231320323112-1233023312202301-2331203021021123-1301132010010212-3223232102103020-2201322031120032-1331103010113130-1212233212302121"></a>

## Next pages — xfcc_options / 112100200112 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](data-sources--workload--reference--group-022.md#canonical-3300013223232120-0213310012303100-2201020132013200-0023231213231111-1000233001313111-1022001122331331-3200203330210203-2302331311311322)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320202213301103-3210202122330002-3131003122112211-0102122122220123-1310110231022101-0000021311102102-1213031231313011-0013223123321312"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert — https_auto_cert / 210200312031 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert

<a id="canonical-2211322212122310-0023103010113233-3101130302202121-0333032030220220-1333203322211333-2322103120331200-3230020300233300-1200022311311313"></a>

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

<a id="canonical-3202000033223302-0000131033031131-3211013323132121-1000321222020312-3300003233132132-0012021302000033-3103000002031002-0101330311032030"></a>

## Direct properties — https_auto_cert / 210200312031 / 3

<a id="canonical-1101111301100030-2020323103322313-3303311202233221-3020313002202222-2102012320032001-2231331113031321-1331131131103233-2221221332121223"></a>

<a id="canonical-0320120222030200-3202110211123221-2102230020212001-2002231200112023-0312030013111302-2131132212001322-3110333121223122-3103223011300112"></a>

## add_hsts property — https_auto_cert / 210200312031 / 4

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

<a id="canonical-1032200133031230-0002102132330000-1023223012101112-0323122030030101-1212130132233013-3112010312221311-3132331011103013-1233003203032322"></a>

<a id="canonical-0332021130330313-1113023012320032-0122023321220332-2131223112230112-3232110013012101-1331022211330201-2022032021103310-3132103112003312"></a>

## append_server_name property — https_auto_cert / 210200312031 / 5

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

- [coalescing_options](data-sources--workload--reference--group-022.md#canonical-3220221300123031-3112310310210012-0121102222200213-3320033313332130-3003323030222230-3320331101303021-0331130233030322-3022003012313103): complete subsection reference.

<a id="canonical-3000120023232033-3020321033002031-2000121033130113-1132103103120113-1030303312123012-3111231022122100-1001031003301222-0132220221211113"></a>

<a id="canonical-2312013202332112-2332121311110111-3133230302313211-2010220300301012-3301301300320030-1011331132133201-0213203002100222-0322323100230221"></a>

## connection_idle_timeout property — https_auto_cert / 210200312031 / 6

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

- [default_header](data-sources--workload--reference--group-022.md#canonical-3103231001222232-3310231220313303-2110233220123220-2111002123123310-1332203123031332-2220021122032012-3211103023302211-0022331213333322): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-022.md#canonical-1103302221331030-2332033032010130-2301133230131312-3203113112032121-3330203222031131-3212123001022000-3312002013222113-2333031333131233): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-022.md#canonical-3232031202133121-1102122001111122-2213323020321030-2313302031220030-2211030223212223-3032200231201003-2321313023201302-2332133223313012): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-022.md#canonical-3120002233022123-2001030022303022-3010033233022120-1010031020033023-1103222111301102-1310233112013220-3233222211320200-3000311023000232): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320): complete subsection reference.

<a id="canonical-2311213011022233-0021122111133321-2032000123001123-3230121110202210-0122313301211223-2233203331021310-0231222131111022-2202303121231202"></a>

<a id="canonical-2120310123133233-2303110330101203-0301132320111002-1021301130020302-3000201200132001-2310010021300302-0023110310211221-3112112313103223"></a>

## http_redirect property — https_auto_cert / 210200312031 / 7

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

- [no_mtls](data-sources--workload--reference--group-023.md#canonical-0102000212101103-1332202100200201-0020022323300210-0303223303003330-1231312230131200-0133331031320223-2133030210013201-3021133112100221): complete subsection reference.

- [non_default_loadbalancer](data-sources--workload--reference--group-023.md#canonical-3031212332220231-0321113012022312-0200221332210221-3132010220210300-3031230023332113-2003022322002010-1323013012111002-0230210102222131): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-023.md#canonical-3121300110323012-2311211111332231-1331033331133131-2013310211001020-0200001010132331-2300201012013200-0012030313110330-2201201212000113): complete subsection reference.

<a id="canonical-0210213100201211-2013200123033121-3302012123332023-1020002202213132-1130302133230131-1202131113103322-3301031333121302-0210300222020111"></a>

<a id="canonical-1301000211201323-0232221211332102-2112212033300021-1131232223301021-0202301233333031-2111012320110110-2132121023121223-2313100323333300"></a>

## port property — https_auto_cert / 210200312031 / 8

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

<a id="canonical-2321101213121032-0331001233322122-1233000322301101-3321101222131302-0103322030001003-2011231113133013-3033132022113301-0221303003131333"></a>

<a id="canonical-1320011221211320-1213110130223331-1322120210333301-0131001013103303-0222112130313321-0121303110332030-1121233301232022-2223312122202200"></a>

## port_ranges property — https_auto_cert / 210200312031 / 9

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

<a id="canonical-3000102221103203-2322002103231333-0233332211131331-3102133000122312-1312021232303101-2233003011130102-1230111301133313-3302331232001222"></a>

<a id="canonical-0030033313031003-3223213021130103-1230012032130020-3322200123202133-0322031032333112-2023232213301302-1100103330200300-0103231133023021"></a>

## server_name property — https_auto_cert / 210200312031 / 10

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

- [tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100): complete subsection reference.

- [use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031): complete subsection reference.

<a id="canonical-3301010302333020-3201001013000103-3230111101011011-0032301033203012-2023120201102210-0020101112113233-2300032012003132-1330222112100031"></a>

## Next pages — https_auto_cert / 210200312031 / 11

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-3220221300123031-3112310310210012-0121102222200213-3320033313332130-3003323030222230-3320331101303021-0331130233030322-3022003012313103)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header](data-sources--workload--reference--group-022.md#canonical-3103231001222232-3310231220313303-2110233220123220-2111002123123310-1332203123031332-2220021122032012-3211103023302211-0022331213333322)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](data-sources--workload--reference--group-022.md#canonical-1103302221331030-2332033032010130-2301133230131312-3203113112032121-3330203222031131-3212123001022000-3312002013222113-2333031333131233)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](data-sources--workload--reference--group-022.md#canonical-3232031202133121-1102122001111122-2213323020321030-2313302031220030-2211030223212223-3032200231201003-2321313023201302-2332133223313012)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](data-sources--workload--reference--group-022.md#canonical-3120002233022123-2001030022303022-3010033233022120-1010031020033023-1103222111301102-1310233112013220-3233222211320200-3000311023000232)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls](data-sources--workload--reference--group-023.md#canonical-0102000212101103-1332202100200201-0020022323300210-0303223303003330-1231312230131200-0133331031320223-2133030210013201-3021133112100221)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](data-sources--workload--reference--group-023.md#canonical-3031212332220231-0321113012022312-0200221332210221-3132010220210300-3031230023332113-2003022322002010-1323013012111002-0230210102222131)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through](data-sources--workload--reference--group-023.md#canonical-3121300110323012-2311211111332231-1331033331133131-2013310211001020-0200001010132331-2300201012013200-0012030313110330-2201201212000113)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](data-sources--workload--reference--group-023.md#canonical-2323123020222323-0201333323111230-1231221203230030-3330100031210233-0222100010111322-3313102010313113-3112021331102313-2133301202102100)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](data-sources--workload--reference--group-023.md#canonical-3332223322300330-0300332101012221-2000211222102303-3100312212012321-3332100132031130-1303221032333032-2111132330120012-0313030221023031)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3220221300123031-3112310310210012-0121102222200213-3320033313332130-3003323030222230-3320331101303021-0331130233030322-3022003012313103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203011231103021-1011103122230012-0332220231231001-2113231300332013-3232133102120310-0122000222033223-0023022120221301-2322120121302331"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options — coalescing_options / 212203110001 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-3322002313002220-0213202030030320-3211231331222213-3230103213331213-0321301301021120-0003231111311320-1310130311033311-3120111031122101"></a>

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

<a id="canonical-0003000021012130-2133210300201000-1133033112031222-2322321002002120-1013201330323113-2231122012312302-2310111031100230-1320121101331032"></a>

## Direct properties — coalescing_options / 212203110001 / 3

- [default_coalescing](data-sources--workload--reference--group-022.md#canonical-1102000103100232-2021100121133302-0303230223012323-3213111011031023-2231002120110112-3000100021201113-3001113320010133-3232031210012011): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-022.md#canonical-0302132310221120-3322313203202010-2321132130333331-3323312313223202-1212213030010302-1310233301100010-2011113102122310-2221133032101100): complete subsection reference.

<a id="canonical-1021331212303222-0331020101303202-2111123013032113-1022331332232020-0210120230001320-0220300210333233-2302233223332023-0230033201223333"></a>

## Next pages — coalescing_options / 212203110001 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](data-sources--workload--reference--group-022.md#canonical-1102000103100232-2021100121133302-0303230223012323-3213111011031023-2231002120110112-3000100021201113-3001113320010133-3232031210012011)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](data-sources--workload--reference--group-022.md#canonical-0302132310221120-3322313203202010-2321132130333331-3323312313223202-1212213030010302-1310233301100010-2011113102122310-2221133032101100)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1102000103100232-2021100121133302-0303230223012323-3213111011031023-2231002120110112-3000100021201113-3001113320010133-3232031210012011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312011323232212-3110032302121322-3113301203130012-1001103012233321-3030011101013011-2322202001321102-0112020122133010-2130312031232110"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 033010200103 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-3220221300123031-3112310310210012-0121102222200213-3320033313332130-3003323030222230-3320331101303021-0331130233030322-3022003012313103)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-3310022130112332-2230322321302032-0021112133030020-0132331100320310-1311101310020212-1031021210023213-0330010312330021-0031133020000332"></a>

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

<a id="canonical-2100003211213131-3311111122021312-2010330221221222-1221310102100013-1312333311211123-0120003323113322-1200331133230203-0030101311101312"></a>

## Direct properties — default_coalescing / 033010200103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010312033223123-1013002312323303-2100020212221120-1233323321111110-3201200302202220-1300020221301221-3120100301331222-3022330203030130"></a>

## Next pages — default_coalescing / 033010200103 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-3220221300123031-3112310310210012-0121102222200213-3320033313332130-3003323030222230-3320331101303021-0331130233030322-3022003012313103)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0302132310221120-3322313203202010-2321132130333331-3323312313223202-1212213030010302-1310233301100010-2011113102122310-2221133032101100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212113030123200-1113020212321002-3001010331030233-1002202313013310-3011120321111200-1232122212230021-2030031122031210-3231302310230233"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 022313103130 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-3220221300123031-3112310310210012-0121102222200213-3320033313332130-3003323030222230-3320331101303021-0331130233030322-3022003012313103)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-2120201033111001-1131002301020111-0320031023111233-2202130100223003-2323203023123233-1120010300223232-1011221023101212-2233211233210020"></a>

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

<a id="canonical-3233010313112322-0320331130121311-0022030233301103-2112210310130230-3321130022313220-0311331203130013-3320220322220103-1233320211233232"></a>

## Direct properties — strict_coalescing / 022313103130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233200303221313-2203010131133013-0101232122311002-2323011203311100-0213002000020230-1120211200013101-1230110111102021-2021223130321300"></a>

## Next pages — strict_coalescing / 022313103130 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](data-sources--workload--reference--group-022.md#canonical-3220221300123031-3112310310210012-0121102222200213-3320033313332130-3003323030222230-3320331101303021-0331130233030322-3022003012313103)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3103231001222232-3310231220313303-2110233220123220-2111002123123310-1332203123031332-2220021122032012-3211103023302211-0022331213333322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323332233303112-2101120012013232-2000201101032112-2101233203120233-3023133222210133-3232313300200111-2002230202111012-3231100101000320"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header — default_header / 123111112330 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-3012000213323302-3011200313332133-1131300012102120-3113223201332101-2121320333220112-3311021212111122-2302212130021003-3112223330100030"></a>

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

<a id="canonical-0011310321333301-3211011303230031-0321003212331020-0323000132332213-0213010222202320-2111332220130223-0220201323301333-0032103201111303"></a>

## Direct properties — default_header / 123111112330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111100132031321-3223102010200000-2010321133313003-0021323212123332-2030100121221032-2113032120310313-1301301002321320-0010001002300133"></a>

## Next pages — default_header / 123111112330 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1103302221331030-2332033032010130-2301133230131312-3203113112032121-3330203222031131-3212123001022000-3312002013222113-2333031333131233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201310202212302-2201311221023310-1020310222333330-3201331331010300-1103031202311021-2332303102210332-1321310203211012-3213211313022000"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — default_loadbalancer / 213232023033 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-3021210222222120-0222210120102101-0232023323030301-1020320102230310-0232331023012112-1101001201120120-1322311001211113-2012000302302011"></a>

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

<a id="canonical-1120323130322212-1022013232211230-0101230210033132-3213113032213221-1012300111313322-3200301100131023-2223130100003213-2001320202222002"></a>

## Direct properties — default_loadbalancer / 213232023033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132012233230210-3100313301012313-3033322233121103-2132102202103120-0221003201302333-2200200113301311-2302202220303110-1223021013311321"></a>

## Next pages — default_loadbalancer / 213232023033 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3232031202133121-1102122001111122-2213323020321030-2313302031220030-2211030223212223-3032200231201003-2321313023201302-2332133223313012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133200002010022-2023233232013330-0321111003231310-1032332210020200-2122323210020101-2313110122330311-1111030321323100-1330010022333132"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — disable_path_normalize / 310333033100 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-1322120123202111-0131123231200233-1131001123103023-0102323301013030-2100032300212230-3031103232223000-3222030121100120-3030203103222110"></a>

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

<a id="canonical-3001102221211212-3311223303121302-2312313002333230-0331131122223013-3212300233013312-3220312320211123-0123203001212100-1200233133020300"></a>

## Direct properties — disable_path_normalize / 310333033100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211122220202233-3302031100221031-2020001113201021-0021132001312302-0320200003300303-0133120013212302-2220112302020320-3001111221223222"></a>

## Next pages — disable_path_normalize / 310333033100 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3120002233022123-2001030022303022-3010033233022120-1010031020033023-1103222111301102-1310233112013220-3233222211320200-3000311023000232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120203002103300-0203003321213222-3321022101210021-0233100203201111-2102233113110220-2112310131000332-1212132311323021-3211030103210200"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — enable_path_normalize / 102113032033 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-2023120131313101-0002103002131111-2211023021310032-2313231111023111-0102001320311123-1220203330032213-0002001120301202-1231033011221211"></a>

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

<a id="canonical-0102300321023211-1212020211321122-3221033320311002-0222033312320323-2111012000332022-1222012223100033-2203010010321202-2222213311010203"></a>

## Direct properties — enable_path_normalize / 102113032033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320133103221002-0210221132313303-0133212233121033-2231331230132120-1000310031300301-1103303012313021-0132021223133223-0130102200230031"></a>

## Next pages — enable_path_normalize / 102113032033 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220021120110213-3331020321210011-1212311231332002-0010310220323030-2230332313201313-1122113322321132-2211020010111010-1213220103030000"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options — http_protocol_options / 002003021121 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-3230312021020011-1213032223101021-3010102001200111-3102030101201203-0101222133031113-1113033220233121-3122100013231031-3101112030303223"></a>

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

<a id="canonical-0103002031212233-2333301312230200-1200120131200012-0033131313130312-1200330313230322-0302131130023100-2323303122020221-1021003113121220"></a>

## Direct properties — http_protocol_options / 002003021121 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-3211312311101322-0011233123001133-2222321232121121-1130131311030122-1331023130321112-0202213032010323-2000222320222013-2310033313303222): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-023.md#canonical-1002130312012131-2130013003020321-3020012310312330-0003001123130301-3113012102022320-0031211011022012-0211201202322331-1102010222133300): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-023.md#canonical-3203112021002200-2003130103200020-3102221112202010-1001201203131130-1332310133132030-3132031202003220-2310330323112100-0220010200131331): complete subsection reference.

<a id="canonical-0220231102231223-3133130302211120-2331103223200110-2211330202321200-1100131012023132-2303012313200012-3213111202231333-1311201302103222"></a>

## Next pages — http_protocol_options / 002003021121 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-022.md#canonical-3211312311101322-0011233123001133-2222321232121121-1130131311030122-1331023130321112-0202213032010323-2000222320222013-2310033313303222)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-023.md#canonical-1002130312012131-2130013003020321-3020012310312330-0003001123130301-3113012102022320-0031211011022012-0211201202322331-1102010222133300)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-023.md#canonical-3203112021002200-2003130103200020-3102221112202010-1001201203131130-1332310133132030-3132031202003220-2310330323112100-0220010200131331)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3211312311101322-0011233123001133-2222321232121121-1130131311030122-1331023130321112-0202213032010323-2000222320222013-2310033313303222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230013233120102-0310002011221111-1230233321010030-2132220233023311-0013220111202231-0012001021122312-2210121310021223-3010203303011020"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 132010232322 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](data-sources--workload--reference--group-020.md#canonical-1220310313301222-1022021330323111-3000020103020220-0001323023100101-1212003022023123-2103220331130211-1221230310000302-1003233331103333)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](data-sources--workload--reference--group-020.md#canonical-0102110102031213-1121223312110232-2131020030001132-1300230131002230-2121033300310012-3210001131010120-0310232322332121-2033221232113110)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](data-sources--workload--reference--group-020.md#canonical-2312300331220231-0102030301212313-0030310003113303-2011030032133010-3321200100201222-0012310323122220-1122203212020211-1333030130110230)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-022.md#canonical-3120032122030013-2103202012301201-3331001221002131-1221031330101133-3201220221020011-0212213033013002-3102212233210301-0200220030321010)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1100323011112313-0121213122313301-1120011032130013-3000322123102220-1033312103232313-3033303200321133-1323322321131000-2233332223030311"></a>

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

<a id="canonical-3021322213021333-0120032331113103-0020003101023222-3023211311132023-3311031132012230-2230013310132000-0000110120103000-1033221230233202"></a>

## Direct properties — http_protocol_enable_v1_only / 132010232322 / 3

- [header_transformation](data-sources--workload--reference--group-022.md#canonical-0313000213023230-1103132333313023-2122323321233320-3113020132003202-0112333312000310-2012012221020031-1121331233130323-0313023300320013): complete subsection reference.

<a id="canonical-0221130130212220-2111330020321110-2322112212213011-1002210012323003-2212032221101113-1300220301102200-0033232120332030-2101132201110223"></a>

## Next pages — http_protocol_enable_v1_only / 132010232322 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-022.md#canonical-0313000213023230-1103132333313023-2122323321233320-3113020132003202-0112333312000310-2012012221020031-1121331233130323-0313023300320013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](data-sources--workload--reference--group-022.md#canonical-3103012111223102-2323200023212331-0010133311323102-1302201311002220-0233003131123331-1001102120301011-3332012213301122-1001000212330320)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0313000213023230-1103132333313023-2122323321233320-3113020132003202-0112333312000310-2012012221020031-1121331233130323-0313023300320013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
