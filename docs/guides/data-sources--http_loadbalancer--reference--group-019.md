---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0320110121232021-2013212130222103-3022303001320113-0020323232010100-2001123002033321-0030220023321111-3023330310300221-0133232113111210"></a>

## Next pages — xfcc_disabled / 130130330223 / 4

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3013120300333313-2132201100310030-1310220122232030-3302131210133100-0133030331022131-2300210103021313-2322311123331223-0123230011202223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222333212111012-3130112220010301-2302231001311132-3000023000212313-3021200223010322-2201121031012113-3103010200200100-3211133332233212"></a>

## https.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 321201210320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-2130001213013112-0103123210133333-3113021131112120-0212312202213003-3230133212032311-3331302202311213-0011222010101203-1302111030201321"></a>

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

<a id="canonical-1223322132231120-1301223300023333-2203321103212101-1033131302031120-1131312323113133-3232322201002022-0313030211000202-0222223230223012"></a>

## Direct properties — xfcc_options / 321201210320 / 3

<a id="canonical-3333310231313021-0311013111032003-2033201100230010-1010022102230101-2032032112122302-2213132201001211-1100303130313101-3212111013312112"></a>

<a id="canonical-3132023130113332-3011222110121110-0330202130012103-0020212203310202-1211103312102321-0311033331210021-1333033221121222-2220132110303013"></a>

## xfcc_header_elements property — xfcc_options / 321201210320 / 4

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

<a id="canonical-2101300111333231-3120031323122131-2230003333032103-3231001002111021-0221022210322100-2203132211311033-2332011302121003-2200230213232113"></a>

## Next pages — xfcc_options / 321201210320 / 5

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332212010321010-2212110033131020-1313310210002302-1222120102120133-0131311202303123-2320202303310120-0201112302302112-1022333300313102"></a>

## https.tls_parameters — tls_parameters / 233033102022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.tls_parameters

<a id="canonical-0020301130102023-0302311331221000-2131120320011101-3101303223201033-1102232300211302-3020101122010120-1203012131312021-1133211332021303"></a>

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

<a id="canonical-3203323303021120-1000001311001202-2230112031102301-3330233331210201-1321211212013031-2332203003112032-2002023230010330-0202023203231331"></a>

## Direct properties — tls_parameters / 233033102022 / 3

- [no_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3201121212312000-1320103233230333-2023311223003113-0102232030012220-3101110302231321-0222013011220221-3130301332332213-3101023121210333): complete subsection reference.

- [tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323): complete subsection reference.

- [tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232): complete subsection reference.

<a id="canonical-0202031210211211-2030221311330310-2030123332210031-2131310122232301-0113102132300011-1102023203003101-0330310102231223-1100122110023313"></a>

## Next pages — tls_parameters / 233033102022 / 4

- [https.tls_parameters.no_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3201121212312000-1320103233230333-2023311223003113-0102232030012220-3101110302231321-0222013011220221-3130301332332213-3101023121210333)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3201121212312000-1320103233230333-2023311223003113-0102232030012220-3101110302231321-0222013011220221-3130301332332213-3101023121210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231021200203203-3320222230133302-2313101231211002-2200122032331012-2132302321233020-1022100313320033-3113210311223333-3310112200222312"></a>

## https.tls_parameters.no_mtls — no_mtls / 220001010103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- https.tls_parameters.no_mtls

<a id="canonical-1133013132301033-0012221120320300-0131120212231022-0303100011302011-2033203022312021-3003011020332033-0000211221111010-1002332200000322"></a>

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

<a id="canonical-3213303310131213-0033310223113222-3311110320023320-3003110031233213-1030323301122110-2000120231031202-2320210020311223-0122021331321332"></a>

## Direct properties — no_mtls / 220001010103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033010333322231-2131133132120101-2301110211200131-0211111323031013-2203310123302102-3223121010313110-3211303222132232-0320332221333103"></a>

## Next pages — no_mtls / 220001010103 / 4

- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303102323022111-3100013000310320-2002301302232110-2320022130313020-2112113232311110-2213020000123231-1100211101023130-0032220222001020"></a>

## https.tls_parameters.tls_certificates — tls_certificates / 232232011212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- https.tls_parameters.tls_certificates

<a id="canonical-2331300233121013-2130300120030312-2333111120021031-3130330120200001-0233031132111220-3122203202120103-1123312001110310-2023111302123302"></a>

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

<a id="canonical-1123330113103223-0320300230100231-1210322130002302-2210111212301201-1330002121003332-1003213300122320-1220201231003212-2111321120122331"></a>

## Direct properties — tls_certificates / 232232011212 / 3

<a id="canonical-0321232132310301-0132122102303120-2011302233130030-0133310100123201-2211111020003030-0230030221210201-0132331031010201-2210203031301230"></a>

<a id="canonical-2112233223022301-1122232231223113-2122330231320001-1220210203202112-0021130022333132-0022122002013123-3203330101002222-3101113303112133"></a>

## certificate_url property — tls_certificates / 232232011212 / 4

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

- [custom_hash_algorithms](data-sources--http_loadbalancer--reference--group-019.md#canonical-2130310001100133-1220032202322233-3023101300330010-2022213021332233-0331030322031202-0303103112320000-0111232112030131-1200220111311232): complete subsection reference.

<a id="canonical-3333203201111120-2103122200230212-0032130122301331-1212031122331022-0102211233031120-1201310332120101-3322202220211013-0231303121103310"></a>

<a id="canonical-1201012331230232-2310100030103212-2200203223130322-2220112312231000-0223023202202030-0100022131013121-3330333222320300-1000300203132123"></a>

## description_spec property — tls_certificates / 232232011212 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--http_loadbalancer--reference--group-019.md#canonical-3121301022103202-0020210230210201-1233201302323220-3031133020213323-0212102312212311-1010311200212101-0113231212231020-3001233332313222): complete subsection reference.

- [private_key](data-sources--http_loadbalancer--reference--group-019.md#canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003): complete subsection reference.

- [use_system_defaults](data-sources--http_loadbalancer--reference--group-019.md#canonical-1001202002030212-2020213322121313-3222120131122030-2101300003120030-1023100100102232-2023123212311020-2132011210001120-2312203113300010): complete subsection reference.

<a id="canonical-2010132211020100-0033101103001112-3100311311310201-2303120132112213-2223101212233023-2022133131231231-3220200210123031-3311023112212300"></a>

## Next pages — tls_certificates / 232232011212 / 6

- [https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--http_loadbalancer--reference--group-019.md#canonical-2130310001100133-1220032202322233-3023101300330010-2022213021332233-0331030322031202-0303103112320000-0111232112030131-1200220111311232)
- [https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--http_loadbalancer--reference--group-019.md#canonical-3121301022103202-0020210230210201-1233201302323220-3031133020213323-0212102312212311-1010311200212101-0113231212231020-3001233332313222)
- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-019.md#canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003)
- [https.tls_parameters.tls_certificates.use_system_defaults](data-sources--http_loadbalancer--reference--group-019.md#canonical-1001202002030212-2020213322121313-3222120131122030-2101300003120030-1023100100102232-2023123212311020-2132011210001120-2312203113300010)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2130310001100133-1220032202322233-3023101300330010-2022213021332233-0331030322031202-0303103112320000-0111232112030131-1200220111311232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210122031100130-1231300211220031-1332333111331201-0131222201132212-3032321232001233-1112100101010230-2010211111302323-2013302231130323"></a>

## https.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 000021230011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-2301012321323123-1121333312221123-2201003201220203-3223121233031012-0102120113023023-3312101233322332-0100122131101132-2020011131320212"></a>

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

<a id="canonical-0301130031011130-2331213120320123-2012312102011221-2202202101012033-0202331200112222-1200031020003030-1111031100231033-3220213223121321"></a>

## Direct properties — custom_hash_algorithms / 000021230011 / 3

<a id="canonical-3021113100122302-0000130210020110-2001230311233233-2031303210122323-0330113132133121-2133302003012322-3200123132100203-1002100112120310"></a>

<a id="canonical-1320220130011001-2001012010331123-0313030132110020-2002300302032100-2322012311023323-3301100303300113-2330102330033321-2121103200333310"></a>

## hash_algorithms property — custom_hash_algorithms / 000021230011 / 4

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

<a id="canonical-0230102032233303-1230011322333030-2130201133133112-2233230321101110-1203212002100202-0203012032000112-0123110212311311-1031020121122202"></a>

## Next pages — custom_hash_algorithms / 000021230011 / 5

- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3121301022103202-0020210230210201-1233201302323220-3031133020213323-0212102312212311-1010311200212101-0113231212231020-3001233332313222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211332211120033-3233102002331201-2300032301220133-2202103201131122-0231230001100210-1130312203123120-2200323222203103-3300120012003323"></a>

## https.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 030221033012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-2303321200202001-1230303202333321-2113100211223203-3112101022220033-0022331201102130-1031012120330233-1011321112313213-1223300113201033"></a>

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

<a id="canonical-3201100222122011-0103103111313331-0033121020331302-2233211133031300-0331011301123211-1330310333102301-2300310230332313-2222301312222022"></a>

## Direct properties — disable_ocsp_stapling / 030221033012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111333002130122-0000030132132301-1133300013301322-0031113012320322-0121113031312131-2320013200031012-0131331300000133-3010330000223103"></a>

## Next pages — disable_ocsp_stapling / 030221033012 / 4

- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331132313230331-1102331311302330-0103232030033013-2001213201300130-3231302210220311-3103021120120122-0120200133032021-2223020113111320"></a>

## https.tls_parameters.tls_certificates.private_key — private_key / 310313213212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- https.tls_parameters.tls_certificates.private_key

<a id="canonical-2313001132021230-2120201201202320-0203233211013102-3233102301021310-2211303202223330-2320100122013221-1101230222221112-2303123102122020"></a>

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

<a id="canonical-3330230102123023-3021231302212113-3321221002103013-3032333132031033-0333003031201322-2003023033132331-1022000230021320-2032022102321201"></a>

## Direct properties — private_key / 310313213212 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-019.md#canonical-0101301211321131-1013101131023322-2201320033212201-3013133012210322-0103312012111132-2323212132131300-2011301221320220-1130300332223230): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-019.md#canonical-3200220310231302-0123030020201003-2102112011211023-1033132310020320-2300210300201101-0300312310233330-2220302223332121-2231230002232300): complete subsection reference.

<a id="canonical-1200013323032111-0230202220021020-3320123102212122-0112203101210000-2002020320130132-2021302031200103-0112200020033120-1033330023201021"></a>

## Next pages — private_key / 310313213212 / 4

- [https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-019.md#canonical-0101301211321131-1013101131023322-2201320033212201-3013133012210322-0103312012111132-2323212132131300-2011301221320220-1130300332223230)
- [https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--http_loadbalancer--reference--group-019.md#canonical-3200220310231302-0123030020201003-2102112011211023-1033132310020320-2300210300201101-0300312310233330-2220302223332121-2231230002232300)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0101301211321131-1013101131023322-2201320033212201-3013133012210322-0103312012111132-2323212132131300-2011301221320220-1130300332223230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302000113220201-2100311222220332-3313222033102020-1022132122212013-3100102231320310-2103023220312030-1123323030121031-3320230120302002"></a>

## https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 121301303233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-019.md#canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003)
- https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0202223010313203-1203021101200010-2021312021001231-0013320121302101-0012001112210030-2122001132231303-2123000103313302-3333320121001210"></a>

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

<a id="canonical-2323001001132123-3023211122202303-3013333232231130-2201130321030003-3313323203101233-1202210321203122-0212202121030200-2131120120211303"></a>

## Direct properties — blindfold_secret_info / 121301303233 / 3

<a id="canonical-1310301213302121-3001011021200303-0211031212222331-1001333133001001-0330231230322120-0021102000023001-2232321121033013-2010031222300333"></a>

<a id="canonical-2313010330222032-0030211031213111-3111120221301202-3211123011132312-3211230220021111-0322330211311123-2120331220122332-1333303131130302"></a>

## decryption_provider property — blindfold_secret_info / 121301303233 / 4

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

<a id="canonical-1010130230003110-1213302313020300-3130301233203002-1223110121333313-1020013331212011-0011002320223021-2331203322302033-1323301022003001"></a>

<a id="canonical-3220121003231112-0312113032201112-3003231303033112-2232331130021323-2331030031320100-1020211311113130-0310012113000001-3032121213013222"></a>

## location property — blindfold_secret_info / 121301303233 / 5

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

<a id="canonical-3122001032220323-2233312033213122-1132030232232332-2030321223030112-1133121100212221-0232221002021303-1131010310021312-0013022000312313"></a>

<a id="canonical-2012010130331211-2003031002221210-2021322122110230-1001011331012331-3131102031312011-0021102221110230-0311022213332112-3300012330003321"></a>

## store_provider property — blindfold_secret_info / 121301303233 / 6

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

<a id="canonical-1301303212321102-0011031130303011-3203200312311100-1302013230202021-1332101323122110-1211111122023010-0312321313101130-3002321201121312"></a>

## Next pages — blindfold_secret_info / 121301303233 / 7

- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-019.md#canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3200220310231302-0123030020201003-2102112011211023-1033132310020320-2300210300201101-0300312310233330-2220302223332121-2231230002232300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131202201213223-0312121122102020-0220112231022230-2213131232212232-2130011102010222-2320003202013011-3330221310132012-0331312211100020"></a>

## https.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 002131322322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-019.md#canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003)
- https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3003000001010212-0002002222301301-2211311001232023-1100313232000200-1100201101210121-2010133202031011-2211110202033131-0303011213002203"></a>

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

<a id="canonical-0200020102023212-1221222033220221-2312123100323023-1221100200030222-0213011013300320-1321121101302003-1310021100023213-2123300103100220"></a>

## Direct properties — clear_secret_info / 002131322322 / 3

<a id="canonical-0113313123013200-3010030103302103-3330323213022321-1323131032032201-2320021201221030-1321321111120311-1220133010101303-0031133103323112"></a>

<a id="canonical-3223131111301111-3211033103212231-3233011120011221-0320223232313321-2323120020210003-2011110030312133-1310213120113322-2112130032223123"></a>

## provider_ref property — clear_secret_info / 002131322322 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2111323220313311-2323121011113120-2000120101111332-2323121311230330-2331011013332032-2003011103210031-2230002223322130-0113000001002221"></a>

<a id="canonical-3001200220101121-3100322001320123-3133100121331133-2202031210231231-1022132011221330-2213230330121312-2232203331300031-1323221223311021"></a>

## URL property — clear_secret_info / 002131322322 / 5

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

<a id="canonical-3032021103330212-2030102211320020-3202102210322121-2312332012023022-0133330320230210-3012122230000112-0101211300201001-2220201110311021"></a>

## Next pages — clear_secret_info / 002131322322 / 6

- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-019.md#canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1001202002030212-2020213322121313-3222120131122030-2101300003120030-1023100100102232-2023123212311020-2132011210001120-2312203113300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312122012303121-0202212002132231-3100230033102122-3011212131123112-3021020030211023-1300033022203233-3111021203232331-3211022033301311"></a>

## https.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 112100333113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-2120222210031223-3021231123312333-0201330331321103-2213003223303003-3100120322310221-3232131301313012-1103010222231200-3233013132133101"></a>

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

<a id="canonical-1123031031202030-2232330302011001-2120133132231230-3203212301101021-2223022210300322-0212002211032022-3033223212233001-2130101001111220"></a>

## Direct properties — use_system_defaults / 112100333113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333303212103030-3303110110122132-3012023331303011-3033032212030233-3311111233311000-2223323221133111-3313222310130020-1211030003110302"></a>

## Next pages — use_system_defaults / 112100333113 / 4

- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-019.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200011301021321-2030123203130011-1132302030213220-0122120231123312-3303203010203000-2332010200202122-0020112203020230-2322300302112122"></a>

## https.tls_parameters.tls_config — tls_config / 212110320102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- https.tls_parameters.tls_config

<a id="canonical-2220011313112200-2112232331122113-1131221032002221-2113233011032321-3220133023133221-3213230311020232-0030101302320303-1333232302331113"></a>

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

<a id="canonical-2300220132331311-1212101203210033-2012223011213130-0020103120223021-0131303313323332-0100321211313310-2012110231331333-3022122012322221"></a>

## Direct properties — tls_config / 212110320102 / 3

- [custom_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-0120113003131211-1222320232220321-1103222033031223-1233013103213100-0121311111012012-1213331101001001-2101110031001332-0322012100203123): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-1301220222130220-0320332112003203-3001321233212300-0202203132012133-2333132211230130-2313200123031312-2211212302223221-1110103210022023): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-2100231223023213-2213001132101021-1332301103122232-3101232202021113-2223303200203321-2323102301200010-1200311003303100-1132213132011220): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-1231131120130120-0312320030302212-2303001023032013-3323330022210210-3012121333203322-0200200002123120-3220001332232003-2203301212212311): complete subsection reference.

<a id="canonical-1122333033121333-3020222332121122-0031301322021100-2312333213231311-3001323231320322-1110311123311011-3220330012022102-3103313000212332"></a>

## Next pages — tls_config / 212110320102 / 4

- [https.tls_parameters.tls_config.custom_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-0120113003131211-1222320232220321-1103222033031223-1233013103213100-0121311111012012-1213331101001001-2101110031001332-0322012100203123)
- [https.tls_parameters.tls_config.default_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-1301220222130220-0320332112003203-3001321233212300-0202203132012133-2333132211230130-2313200123031312-2211212302223221-1110103210022023)
- [https.tls_parameters.tls_config.low_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-2100231223023213-2213001132101021-1332301103122232-3101232202021113-2223303200203321-2323102301200010-1200311003303100-1132213132011220)
- [https.tls_parameters.tls_config.medium_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-1231131120130120-0312320030302212-2303001023032013-3323330022210210-3012121333203322-0200200002123120-3220001332232003-2203301212212311)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0120113003131211-1222320232220321-1103222033031223-1233013103213100-0121311111012012-1213331101001001-2101110031001332-0322012100203123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031131331102011-0311002313201012-3030100131020133-0321003222222101-0011001221310323-1000113123331233-1331011013101001-2311221022233031"></a>

## https.tls_parameters.tls_config.custom_security — custom_security / 133201122320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- https.tls_parameters.tls_config.custom_security

<a id="canonical-3123113312122123-2323302111113131-3023320311121203-2031000301102222-2121221221122331-0023010233201012-0130012231101113-2122321302303202"></a>

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

<a id="canonical-1031013010003002-3120020111002103-1333200321101321-1120311023221211-1331020010211112-3021222003200021-0121303212302130-1221032312010031"></a>

## Direct properties — custom_security / 133201122320 / 3

<a id="canonical-2022230313013222-0330033231121223-3233300001221212-3113312011301112-0322031213203113-0201101211032131-3322013123113301-1231202031210020"></a>

<a id="canonical-3020023103320011-2033333003233132-3020002310110012-3101022320101330-2222233222101011-3120203221020101-3010012233011021-2010013220312310"></a>

## cipher_suites property — custom_security / 133201122320 / 4

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

<a id="canonical-0010321301201120-2313111323301212-1113322212013132-2003111333123001-0302013303010211-1231100233220101-1312311113021011-1100332323212032"></a>

<a id="canonical-1123222230020220-1212303223031030-1203002023012230-3223201120132011-3311023122223201-0310112121120220-3012301231301022-2302110331013002"></a>

## max_version property — custom_security / 133201122320 / 5

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

<a id="canonical-1303213330302312-2231020120230100-0213210003012002-1203223022033000-2231303330012300-1130222311133330-1302033201122122-1022212130320221"></a>

<a id="canonical-1133212032100110-3011213233002030-0323031303123123-0113020221033233-2332001013121130-1031022022203022-3302302111213103-3131312130330223"></a>

## min_version property — custom_security / 133201122320 / 6

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

<a id="canonical-1212300030131103-0131000202013330-0001000333311020-3120311110033000-0120333230231310-3112103232212213-1030132310011312-3100311100023001"></a>

## Next pages — custom_security / 133201122320 / 7

- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1301220222130220-0320332112003203-3001321233212300-0202203132012133-2333132211230130-2313200123031312-2211212302223221-1110103210022023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131110202000113-2003320111010113-3331313300111122-3111033003212033-0120300002122111-2203010102210120-1312023020111121-3212323000202323"></a>

## https.tls_parameters.tls_config.default_security — default_security / 310333132321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- https.tls_parameters.tls_config.default_security

<a id="canonical-1033320200100101-2103311100000032-3122030322010212-3030201332221302-3123301323021102-0103320211012313-1221301131230211-3110111320221031"></a>

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

<a id="canonical-3301231121300022-0002301201010320-3221120221323233-2002120201213332-2333131220121222-0010232112012121-0311000233210321-3332330212023220"></a>

## Direct properties — default_security / 310333132321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122333302121121-2203002312011023-1332112111223100-2031311302000112-3100013220030321-3303021221033310-0021300303221003-2130301122012222"></a>

## Next pages — default_security / 310333132321 / 4

- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2100231223023213-2213001132101021-1332301103122232-3101232202021113-2223303200203321-2323102301200010-1200311003303100-1132213132011220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103221231202222-1220313111002331-1320212123111321-2332300131010200-1032220222133002-2213220221231000-2310120211120203-3102311130323301"></a>

## https.tls_parameters.tls_config.low_security — low_security / 230201023233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- https.tls_parameters.tls_config.low_security

<a id="canonical-2110311332123103-2222322021113002-0012112111021032-2133102131123302-2210002102123032-3321131232133131-0001130131121203-3001330230000032"></a>

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

<a id="canonical-3130230301110111-1223020222330303-2331313212113232-2033003013020011-0131101300023011-0101233301130303-0231333103103300-2221113310022000"></a>

## Direct properties — low_security / 230201023233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132210112231210-0302322121220202-2221023202213203-2332211212311131-1202231221132233-0001121231222021-0203323212133231-0001132331032112"></a>

## Next pages — low_security / 230201023233 / 4

- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1231131120130120-0312320030302212-2303001023032013-3323330022210210-3012121333203322-0200200002123120-3220001332232003-2203301212212311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300011122213010-0210113001010132-3121203001010111-1200102020311111-0233220131110311-2301112032021322-3221223210312003-0201321011110200"></a>

## https.tls_parameters.tls_config.medium_security — medium_security / 301121222231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- https.tls_parameters.tls_config.medium_security

<a id="canonical-2212312301311031-0103111011022020-3000200100230032-2220012002212012-1203131100221032-3121010030122022-0311011000222310-2203203211020133"></a>

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

<a id="canonical-0200001010301000-1203323133233100-2333322202032132-1302323222203232-2213321022100003-0020233321321132-1101333232022311-2100302022330213"></a>

## Direct properties — medium_security / 301121222231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113211011001003-3202233010200132-2030331000322203-0121002023232110-2312232030021123-0011010212200303-3221130300012312-2333211333020130"></a>

## Next pages — medium_security / 301121222231 / 4

- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132002212323321-1001133321012311-1202121120300330-0030311311131021-1231221110003303-2110113100022003-1011333333023120-2232213123032313"></a>

## https.tls_parameters.use_mtls — use_mtls / 110330220133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- https.tls_parameters.use_mtls

<a id="canonical-3322110332022200-3230111111213320-2220011031023301-1020312011321223-3321230122210123-0300333303233110-0330132033103022-3021132131312131"></a>

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

<a id="canonical-2032201021323333-1200132220113000-3112200100202012-0131020120012100-2031300033210320-2133302032202212-0133120321333022-2301302330220121"></a>

## Direct properties — use_mtls / 110330220133 / 3

<a id="canonical-2133102121003001-0011032220122220-1121132220011223-0330321311033002-1120023012200203-3001302121321103-0132003323010313-1122320220002332"></a>

<a id="canonical-1311213300020323-0310301323303213-3210223201302102-3111131112100231-0310311122000212-1002032130100013-1213132322331313-2313012111302320"></a>

## client_certificate_optional property — use_mtls / 110330220133 / 4

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

- [crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-2120200310111303-2123002332032223-1320022103031131-2131313231021033-3223303003130213-2332200022300000-3300301003112020-3121000323201322): complete subsection reference.

- [no_crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-3320113311333223-0331311201323111-3230310102122032-3110221022211232-3221133201031320-3221313011212200-0113023301232113-0213230313003003): complete subsection reference.

- [trusted_ca](data-sources--http_loadbalancer--reference--group-019.md#canonical-3130030222012122-3200001201311013-2320233200223213-2023020020303332-3100033030103312-3221133022132003-3311030330120211-3330220323230031): complete subsection reference.

<a id="canonical-0101003002023002-3233120221330120-2200002002013302-1103310210020312-0300321310112133-0211130302133323-1011233032302231-2131203122321003"></a>

<a id="canonical-2303000102310132-3010121100121003-0302310300220102-2000300123103323-1321330013112030-3031022113220201-1203120201202221-0031232303113133"></a>

## trusted_ca_url property — use_mtls / 110330220133 / 5

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

- [xfcc_disabled](data-sources--http_loadbalancer--reference--group-019.md#canonical-0211201201101012-2032030012112333-1023312213220032-3102332300203301-0211033020001201-3031133001311121-3033000210103130-2210310102021001): complete subsection reference.

- [xfcc_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-1322023323220103-0120111112033232-3301100122101000-1022023102030303-3333300212330203-2221222021232123-2032120200322113-2021012121033102): complete subsection reference.

<a id="canonical-3023330123131313-0203102112302333-3032023312130302-3012332333301302-2111123302131202-2030010031013212-2313231300303121-3202302310321132"></a>

## Next pages — use_mtls / 110330220133 / 6

- [https.tls_parameters.use_mtls.crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-2120200310111303-2123002332032223-1320022103031131-2131313231021033-3223303003130213-2332200022300000-3300301003112020-3121000323201322)
- [https.tls_parameters.use_mtls.no_crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-3320113311333223-0331311201323111-3230310102122032-3110221022211232-3221133201031320-3221313011212200-0113023301232113-0213230313003003)
- [https.tls_parameters.use_mtls.trusted_ca](data-sources--http_loadbalancer--reference--group-019.md#canonical-3130030222012122-3200001201311013-2320233200223213-2023020020303332-3100033030103312-3221133022132003-3311030330120211-3330220323230031)
- [https.tls_parameters.use_mtls.xfcc_disabled](data-sources--http_loadbalancer--reference--group-019.md#canonical-0211201201101012-2032030012112333-1023312213220032-3102332300203301-0211033020001201-3031133001311121-3033000210103130-2210310102021001)
- [https.tls_parameters.use_mtls.xfcc_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-1322023323220103-0120111112033232-3301100122101000-1022023102030303-3333300212330203-2221222021232123-2032120200322113-2021012121033102)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2120200310111303-2123002332032223-1320022103031131-2131313231021033-3223303003130213-2332200022300000-3300301003112020-3121000323201322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220120131101030-0211222310112301-3023233012022030-1230011211302332-2333211331032302-3300130231023210-1312123322002303-0131232320323330"></a>

## https.tls_parameters.use_mtls.crl — crl / 201033213320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.crl

<a id="canonical-1332020003312231-1221223330013003-2031200313212222-1113222203200011-1223122032331031-1321122102001233-2012120322120321-1011232323330132"></a>

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

<a id="canonical-3130002103312322-3131132130210303-3020212331201000-3301322122030210-1211023203333112-3003230212330122-1311102020133312-3111132221330100"></a>

## Direct properties — crl / 201033213320 / 3

<a id="canonical-3331002021302301-0330322310201211-1311202303312013-0133220300202113-3130220300131031-2331221032032311-0222132032020321-0010023301021130"></a>

<a id="canonical-0101103133320013-0233030320231122-2310011220210031-3300002332132302-1331321220223123-2030223131101332-2300301101003301-3010102322133310"></a>

## name property — crl / 201033213320 / 4

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

<a id="canonical-0303312231121300-0110220233110320-2330230213030231-0122101022331330-3121001011300300-1100331000002212-1212111020020011-3023023302203200"></a>

<a id="canonical-2301113131210103-2110111032302333-1332233111002002-1210333222113110-1300320211022023-1232013010101122-0313133121023300-0022020031333031"></a>

## namespace property — crl / 201033213320 / 5

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

<a id="canonical-2013331223121203-0132203020210001-0210122201122313-2322230223230030-1001301001310132-3100130030302013-1210312023331121-2102230131022133"></a>

<a id="canonical-0303131122312311-1323023100003222-3012131310332222-2210001232022303-1331013201011321-3033331002033002-0301022011221121-3213200112120302"></a>

## tenant property — crl / 201033213320 / 6

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

<a id="canonical-0121012223002321-0111210323331321-1023012020320330-3303220101000311-0103130311330333-0102020013022000-1201132122130012-2321301312030330"></a>

## Next pages — crl / 201033213320 / 7

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3320113311333223-0331311201323111-3230310102122032-3110221022211232-3221133201031320-3221313011212200-0113023301232113-0213230313003003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201300032203133-3000022122132011-2313333010100112-1201213020221232-0232310300203312-1132232111122321-3000101022202100-3223113032110012"></a>

## https.tls_parameters.use_mtls.no_crl — no_crl / 110002132112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.no_crl

<a id="canonical-3110223100130310-2032020003111330-3020130211313200-0312211213003311-1232311321302112-1222122032002010-1302300300133201-0233111230023213"></a>

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

<a id="canonical-2321301212312231-2211322232003111-3302323330220130-1300113303232330-1311103200123203-3120300131303111-3012131312102031-3213023030310322"></a>

## Direct properties — no_crl / 110002132112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210331112331111-1132020330000211-0330103010030332-2133222303003013-2232200320330002-0220310302013333-0132231311112301-1120301320123223"></a>

## Next pages — no_crl / 110002132112 / 4

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3130030222012122-3200001201311013-2320233200223213-2023020020303332-3100033030103312-3221133022132003-3311030330120211-3330220323230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210323333233023-1121133223033233-3001301113333332-0023101320000021-2333301131310320-3010120203322003-1200120130130023-2032030132202103"></a>

## https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 133310013123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0112200110031212-2210000101230001-2111323220223222-3301210332222213-1201113330201300-2031222231020231-0013220320030210-0032220303130233"></a>

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

<a id="canonical-0120211313312032-2121212002331030-0013202110213212-3130023232011031-1213120312133020-2221310233332121-1321211202323331-2123110113113101"></a>

## Direct properties — trusted_ca / 133310013123 / 3

<a id="canonical-0102213021210203-1213303232001121-3310020331312320-1011031023003012-0113102322131101-1320201123120003-1031110011210220-2313211302332112"></a>

<a id="canonical-3002131022003333-3331212122320233-3233223332033311-2202122203013103-2102323021203101-2310310331032013-3031301033211123-3313101302032111"></a>

## name property — trusted_ca / 133310013123 / 4

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

<a id="canonical-1131101131120302-2021211210103200-3013310033210232-1200133230222032-0003222220331123-1232122002103313-3023203331322132-0133302312311213"></a>

<a id="canonical-2233132131213223-0032120310130232-3323002010121122-3100120202331101-2211010022121112-0000223233212030-0222321233311201-0213232023330121"></a>

## namespace property — trusted_ca / 133310013123 / 5

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

<a id="canonical-3233011020300321-1332303333210002-2313023203122233-2123102120201100-2212011223310213-1011121121110031-2022322231301312-3000300003331111"></a>

<a id="canonical-3031220112212133-3223011032323211-1322222122313020-0031213000331112-1002310013031220-2202120330020233-0222323002010201-2231110301333021"></a>

## tenant property — trusted_ca / 133310013123 / 6

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

<a id="canonical-0310120030010203-2113131222230112-0133102011032301-0220211332312111-2021321221120321-1100233321221212-1113231323100000-0231313010023213"></a>

## Next pages — trusted_ca / 133310013123 / 7

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0211201201101012-2032030012112333-1023312213220032-3102332300203301-0211033020001201-3031133001311121-3033000210103130-2210310102021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230212200000302-1002203211301002-2033030333223321-0012020013132320-1021222210210200-0222311032020110-0200213301232103-0321110020213302"></a>

## https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 020312321120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-1331032301023001-3233221131202213-3013203003233100-3100310233000002-1133211110313230-2303132201302232-0321001233101233-3203303001210332"></a>

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

<a id="canonical-2321323201222102-1332213020101312-2203313320312033-1010233210120011-1300033212020320-2103302111311223-0032222312222120-0031000300101313"></a>

## Direct properties — xfcc_disabled / 020312321120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001133023030303-2110012132313220-1222301303331130-0321032333301312-3030223323030110-1231101331001331-1213032200120131-3211122311131001"></a>

## Next pages — xfcc_disabled / 020312321120 / 4

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1322023323220103-0120111112033232-3301100122101000-1022023102030303-3333300212330203-2221222021232123-2032120200322113-2021012121033102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202300213202300-0211233123302123-2311003020121030-1100101112100130-0131131122133011-3031312332220311-3301323333031233-0113023122232101"></a>

## https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 122000223302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-3123120232210113-1120201001032320-2111202300231311-1233011303113113-0033112111130132-2103102323301100-2121332011213003-0221103330122030"></a>

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

<a id="canonical-1220200001313221-1200232001233232-1320111323013202-2110320013002111-2100203013333012-1333020321102003-2310013232213100-3320202231033331"></a>

## Direct properties — xfcc_options / 122000223302 / 3

<a id="canonical-1201111313033303-2301101222133031-1131123132122133-1032332010230322-3330002213022112-0323121332320112-1311202313003020-1323232323011023"></a>

<a id="canonical-3113320012231311-1001211331201002-0303132032302001-3100131123111330-3011203321330301-0101002000200223-0131100022012233-0113011110300322"></a>

## xfcc_header_elements property — xfcc_options / 122000223302 / 4

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

<a id="canonical-2100033021233231-0022331010302333-0002131213003232-3023233120032223-1011002111201120-0330331312303202-0221331202230231-1321121310302312"></a>

## Next pages — xfcc_options / 122000223302 / 5

- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130023103110112-2011130000201230-0233132200200100-0131201212002100-1002133231310233-0330302200333230-3101332020123012-2010033311102203"></a>

## https_auto_cert — https_auto_cert / 201200220313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- https_auto_cert

<a id="canonical-2110233032332000-2230100230102301-1331113121321211-2111113312132303-0303213311122331-0023321211013320-1001321133320230-1213120312211202"></a>

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

<a id="canonical-2323002003113130-1213112321300033-1122030230200203-1220313320122123-3330022310212000-0121012101001022-2202121321121031-0330032313100300"></a>

## Direct properties — https_auto_cert / 201200220313 / 3

<a id="canonical-0311303033330303-3310230302321111-3310101311311301-0233300221203113-0203113302113003-0330110321330102-0222223303302301-2103113302231311"></a>

<a id="canonical-3101332222001001-3023201221303112-3301320331333030-1320310302220301-2313300103010113-2233200113223102-3130320120110200-0133222001020111"></a>

## add_hsts property — https_auto_cert / 201200220313 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header. Defaults to \`false\`. Server applies default
when omitted.

Upstream description:

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

<a id="canonical-0331322310020213-0323122013022122-2132011311010113-3333332311100003-0320221330101333-1110331123332312-3021233210201302-2230211233320122"></a>

<a id="canonical-3210213012000311-2333110300010302-2110001230231111-3122330011210003-1033203301220030-2233313210321123-1020121031223311-3310331221133123"></a>

## append_server_name property — https_auto_cert / 201200220313 / 5

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

- [coalescing_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130): complete subsection reference.

<a id="canonical-1022121000013012-2210122201230101-0323231032201000-3112220222201320-1021312213330001-1312122132113121-1120021201212110-3023013230002123"></a>

<a id="canonical-1113033023111211-2202321012032103-1123002301330322-0002320120211330-3222113000101211-0333130213020211-3232220103123133-3113002330121110"></a>

## connection_idle_timeout property — https_auto_cert / 201200220313 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Server
applies default when omitted.

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

- [default_header](data-sources--http_loadbalancer--reference--group-019.md#canonical-3313201031301311-2102131330032210-3313230003331111-0210121112312003-1313020011212333-0021220023223203-2231232032113133-1232222301222200): complete subsection reference.

- [default_loadbalancer](data-sources--http_loadbalancer--reference--group-019.md#canonical-1103130031312003-1131202320320111-3122032303332320-0133012311002333-0120323032300311-0210331101320210-0111202033012332-2321332303211221): complete subsection reference.

- [disable_path_normalize](data-sources--http_loadbalancer--reference--group-019.md#canonical-2230032312031133-3013111001221130-3332003130301132-0033310101300322-2330212002021033-2331122100220031-1313210230130110-3033330130320201): complete subsection reference.

- [enable_path_normalize](data-sources--http_loadbalancer--reference--group-019.md#canonical-0220213122201313-3032013121031020-3011110313200301-0022211020120231-3132020001103231-0223313003010232-3011313203012023-1001002312333313): complete subsection reference.

- [http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101): complete subsection reference.

<a id="canonical-2300013310330101-2231112312330110-2220123021033103-1121103012133100-0022010222230123-1010232133130323-1212203132330132-2103103120332300"></a>

<a id="canonical-1230101321231101-1230320102100023-2002312321012101-2232201210312312-3121000003321313-3123220221110333-1110211132233022-3323231133133023"></a>

## http_redirect property — https_auto_cert / 201200220313 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS. Defaults to \`false\`. Server applies
default when omitted.

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

- [no_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-2102301213102300-3130211303123110-1211031100221301-0101130110032303-3233220303001111-3333311120020322-3000202300233331-2213001132011033): complete subsection reference.

- [non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-019.md#canonical-3311320301322302-1303220013302033-0332111200020023-1101002220032022-2331300211113301-0130223132030323-1230211323100313-3210102020111331): complete subsection reference.

- [pass_through](data-sources--http_loadbalancer--reference--group-019.md#canonical-2221203212221022-2312321133113202-1222303012100333-2123103030313231-1012021303103113-3232300033113221-1000231113031103-1202131222211332): complete subsection reference.

<a id="canonical-1223132001313023-1132130130033013-2223113102320013-1233222230030003-3110033102033313-2313000012313310-3312300320110013-3301030113232203"></a>

<a id="canonical-3312132030303123-3132210000301113-3103331320313300-0330030022331303-3213003120313000-3003032002331023-1031303332023302-0012322023313300"></a>

## port property — https_auto_cert / 201200220313 / 8

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

<a id="canonical-2221033321122331-2320132130221131-2103331021120121-2232302330210213-3110120001033202-3302220101023133-2011000211211333-3231030031022003"></a>

<a id="canonical-1013310221302203-1203112230310213-0330102303330313-2303110102012031-1301320303002110-2313331313003121-2002332223301221-0022201232202220"></a>

## port_ranges property — https_auto_cert / 201200220313 / 9

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

<a id="canonical-3233311203010311-3020131011312031-1321323001322102-0030023112213323-3033223021201001-0100303110310320-2121310001110313-3013333121201101"></a>

<a id="canonical-3320122302223320-0311212103323000-2103303302022323-3223323201330002-3313230122133203-2232122123021210-3113331332303020-0022122233313030"></a>

## server_name property — https_auto_cert / 201200220313 / 10

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

- [tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330): complete subsection reference.

<a id="canonical-0112331012220222-3002302320222222-2020302012020123-1030023332001133-3202021213210312-2133111122132333-0200111232102023-2330212000231002"></a>

## Next pages — https_auto_cert / 201200220313 / 11

- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130)
- [https_auto_cert.default_header](data-sources--http_loadbalancer--reference--group-019.md#canonical-3313201031301311-2102131330032210-3313230003331111-0210121112312003-1313020011212333-0021220023223203-2231232032113133-1232222301222200)
- [https_auto_cert.default_loadbalancer](data-sources--http_loadbalancer--reference--group-019.md#canonical-1103130031312003-1131202320320111-3122032303332320-0133012311002333-0120323032300311-0210331101320210-0111202033012332-2321332303211221)
- [https_auto_cert.disable_path_normalize](data-sources--http_loadbalancer--reference--group-019.md#canonical-2230032312031133-3013111001221130-3332003130301132-0033310101300322-2330212002021033-2331122100220031-1313210230130110-3033330130320201)
- [https_auto_cert.enable_path_normalize](data-sources--http_loadbalancer--reference--group-019.md#canonical-0220213122201313-3032013121031020-3011110313200301-0022211020120231-3132020001103231-0223313003010232-3011313203012023-1001002312333313)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [https_auto_cert.no_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-2102301213102300-3130211303123110-1211031100221301-0101130110032303-3233220303001111-3333311120020322-3000202300233331-2213001132011033)
- [https_auto_cert.non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-019.md#canonical-3311320301322302-1303220013302033-0332111200020023-1101002220032022-2331300211113301-0130223132030323-1230211323100313-3210102020111331)
- [https_auto_cert.pass_through](data-sources--http_loadbalancer--reference--group-019.md#canonical-2221203212221022-2312321133113202-1222303012100333-2123103030313231-1012021303103113-3232300033113221-1000231113031103-1202131222211332)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231313032120122-0033212122033103-3013202002132130-0022202223233130-2022022121111200-0122203112020033-0103021202322132-3333303012222021"></a>

## https_auto_cert.coalescing_options — coalescing_options / 221331221023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.coalescing_options

<a id="canonical-3021311112321010-2123213221032100-1022221302021303-1233121321132130-3331311112020211-0230211310322133-3302321111130103-2202233300110022"></a>

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

<a id="canonical-0033003120301131-2133303200230022-1122111212031221-1323021111201022-3221102303210102-2112032003301331-2303022021323120-1300113130103320"></a>

## Direct properties — coalescing_options / 221331221023 / 3

- [default_coalescing](data-sources--http_loadbalancer--reference--group-019.md#canonical-0200233202323332-2013331221132310-0011132211023212-0130321121221120-1121220203202323-0213333300220112-1130200031032132-1233230203100003): complete subsection reference.

- [strict_coalescing](data-sources--http_loadbalancer--reference--group-019.md#canonical-3111003333102233-0313220330123210-2332032101011011-3213121121321113-1213132132201201-3102132322123112-3103222003222232-2031320322011023): complete subsection reference.

<a id="canonical-2312320200112320-2223231122113001-0300213121111332-1300320032221302-0100203200130233-2201122010111011-2212120023010032-2021333023133103"></a>

## Next pages — coalescing_options / 221331221023 / 4

- [https_auto_cert.coalescing_options.default_coalescing](data-sources--http_loadbalancer--reference--group-019.md#canonical-0200233202323332-2013331221132310-0011132211023212-0130321121221120-1121220203202323-0213333300220112-1130200031032132-1233230203100003)
- [https_auto_cert.coalescing_options.strict_coalescing](data-sources--http_loadbalancer--reference--group-019.md#canonical-3111003333102233-0313220330123210-2332032101011011-3213121121321113-1213132132201201-3102132322123112-3103222003222232-2031320322011023)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0200233202323332-2013331221132310-0011132211023212-0130321121221120-1121220203202323-0213333300220112-1130200031032132-1233230203100003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023203110210110-3330020313011232-2120311020021332-1131001330112220-3310221231210121-3123031032120221-0100023233013310-2113020332200100"></a>

## https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 221020312011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130)
- https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-3000211232211231-2211313012312213-3300122103111203-1012302301332313-1022012322010203-2101221333131013-2020100232013223-0133020302030203"></a>

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

<a id="canonical-3300010131320311-1010310201330123-2212221123112130-2023313230011312-0320301233112233-2122321231133010-1123313030111320-2131133033130223"></a>

## Direct properties — default_coalescing / 221020312011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231030203131313-2000022202022211-2103101223102023-1333221202000131-2102333300203000-1223201330323311-0222202310113032-2123323310301001"></a>

## Next pages — default_coalescing / 221020312011 / 4

- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3111003333102233-0313220330123210-2332032101011011-3213121121321113-1213132132201201-3102132322123112-3103222003222232-2031320322011023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003132010122332-1332301200211313-0302113120113303-0230203000010332-3202030232311113-1033332321033232-3112112033022100-1100003310222300"></a>

## https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 002310220012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130)
- https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-1232211103222321-1213020133213102-1223032112033203-0322233012300131-1233203033210032-2102300033313232-3330113321213222-3211213231320210"></a>

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

<a id="canonical-0010120103220020-2231312320201310-0311103302112331-0313000332030121-2210202233023322-1010210001312113-2303301011100231-3100232221312232"></a>

## Direct properties — strict_coalescing / 002310220012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222232230203221-0201313202333000-0201030210132333-0000222002011230-2330003100202100-3220011231021331-3201310111233030-0011200003310021"></a>

## Next pages — strict_coalescing / 002310220012 / 4

- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3313201031301311-2102131330032210-3313230003331111-0210121112312003-1313020011212333-0021220023223203-2231232032113133-1232222301222200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122331010221023-0333203002203200-3210313210033231-0203321301122022-2100202321322010-3103221201222121-2322130203030201-2233330011100222"></a>

## https_auto_cert.default_header — default_header / 303200321020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.default_header

<a id="canonical-1332131330320330-1321132103220233-2120221220011102-2103131200010313-3002220212013211-1330003002213112-1012332033331211-2320331201030303"></a>

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

<a id="canonical-1030111222222322-2300300201020233-0310332320112010-3121330202032232-2230333033211100-3010020332230320-1202312113332232-2300233102001022"></a>

## Direct properties — default_header / 303200321020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000213202333331-2023100032222333-3101300120033121-2021301312132000-0011321000103200-1112120203321001-0132232032110021-2110020123120110"></a>

## Next pages — default_header / 303200321020 / 4

- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1103130031312003-1131202320320111-3122032303332320-0133012311002333-0120323032300311-0210331101320210-0111202033012332-2321332303211221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103000032302222-2120332132010121-0002221011212102-0103021203230202-3120000303023331-0131132333100302-1212122312332020-3110321023303322"></a>

## https_auto_cert.default_loadbalancer — default_loadbalancer / 320032022300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.default_loadbalancer

<a id="canonical-1310002031103010-3003232121201212-0302302330012121-0203311003230212-3103021323233100-3033313333021321-0222012030132132-0003232323333032"></a>

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

<a id="canonical-0001311211323131-2011332301203202-3030103332023210-3331211210302011-1303211132033323-1302012101223302-2203103312101120-0133133320022120"></a>

## Direct properties — default_loadbalancer / 320032022300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103123111233111-1001203102301000-0012112320001031-1302211322301013-0303102013311332-2231311110132220-1112100120110211-0010020303032023"></a>

## Next pages — default_loadbalancer / 320032022300 / 4

- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2230032312031133-3013111001221130-3332003130301132-0033310101300322-2330212002021033-2331122100220031-1313210230130110-3033330130320201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103031202001012-1210231213300200-0133211320023200-1331331110122103-0010203203323333-0332003313030130-3033201120100002-1331001022220123"></a>

## https_auto_cert.disable_path_normalize — disable_path_normalize / 222200121000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.disable_path_normalize

<a id="canonical-2102321302210232-2332302131231010-0131312013131320-2012212301310233-0120332230332033-2112313030020221-3110202212112200-0332021003233212"></a>

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

<a id="canonical-2130103101032311-0002300322013231-0000122331202130-0032000130022203-1303321313300133-1232103002320302-0002221320132001-2320211100330002"></a>

## Direct properties — disable_path_normalize / 222200121000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233020303020231-1031020032032003-2101032210302020-0011110122101003-0103211303203111-2103120220113020-2232032312231012-0231323312031201"></a>

## Next pages — disable_path_normalize / 222200121000 / 4

- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0220213122201313-3032013121031020-3011110313200301-0022211020120231-3132020001103231-0223313003010232-3011313203012023-1001002312333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210010023100121-3103322003231013-0013303122310132-0121232312032101-3032220201213210-0100313310330321-2311232312311202-3120230201302203"></a>

## https_auto_cert.enable_path_normalize — enable_path_normalize / 131310102232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.enable_path_normalize

<a id="canonical-2022132030002222-1301101020102230-1130210022011301-0013220120003133-0213212031123332-0003110001233132-2012133123212302-0103012300312330"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-0231022320332120-1311310320002031-3331200311211021-0333123230011301-2000310130032002-0233011130111221-0030312100200033-2310302303101333"></a>

## Direct properties — enable_path_normalize / 131310102232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232101232003032-0101323321131000-2112131323310201-1331012321201100-1032020031022220-2003132031323232-1321022000103213-3011033002312200"></a>

## Next pages — enable_path_normalize / 131310102232 / 4

- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313331300031222-2013222033333222-1303033303311002-0100001131030221-2202122300121000-2110210022013001-0322300133330133-2101312223230113"></a>

## https_auto_cert.http_protocol_options — http_protocol_options / 012331233103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.http_protocol_options

<a id="canonical-1331022330003023-2221013213220111-2323330022030111-1122110020120211-0023131200032311-3103030130323100-0203233312031121-0021033122100321"></a>

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

<a id="canonical-1012203321212312-3330313302213222-0201120003331310-3003322221213012-3020011023011032-3232002321320322-0200222202313110-3211321323111023"></a>

## Direct properties — http_protocol_options / 012331233103 / 3

- [http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--http_loadbalancer--reference--group-019.md#canonical-1103132123031212-1312011311132313-1101103331230311-0211210100021021-0223211311301221-2201203120330131-0323223220202031-2113302232301212): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-0011132011130333-0210213032201022-0233112013130221-1333101332310100-3311123321233200-3302001330230303-1003013201132323-0021222312201002): complete subsection reference.

<a id="canonical-1100212222023131-1010313221333221-3022300101220020-3030332300313311-0311033021113210-3002321121202323-3203302131112020-1020012333100330"></a>

## Next pages — http_protocol_options / 012331233103 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--http_loadbalancer--reference--group-019.md#canonical-1103132123031212-1312011311132313-1101103331230311-0211210100021021-0223211311301221-2201203120330131-0323223220202031-2113302232301212)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-0011132011130333-0210213032201022-0233112013130221-1333101332310100-3311123321233200-3302001330230303-1003013201132323-0021222312201002)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120233130210211-3023110031222322-3213002310233102-3011212222003100-3100301003013202-3201132302302323-0133331020221021-2201121231200201"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 222212331112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3122230030213302-0110123203102111-1022330022220130-2110032233112112-2121030021330223-3013320302021203-2322202231033223-2000331212220021"></a>

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

<a id="canonical-3131003132110111-2100301112231103-2323223010220310-2013133103300111-3110212220100311-2103233101213132-0321310310233102-0202130103311213"></a>

## Direct properties — http_protocol_enable_v1_only / 222212331112 / 3

- [header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003): complete subsection reference.

<a id="canonical-1312203203233310-0012111113023303-0121311321123223-2112302330020023-3100021231212300-1123021210231023-2130323123022231-0012300030333202"></a>

## Next pages — http_protocol_enable_v1_only / 222212331112 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301121110102333-3003232202003310-0323313222301011-2132310311110310-0333233010320332-1303111322213100-3200210013113231-2031111001023021"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 101232110110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1121200303002003-1231323131233030-0223313122113033-0310223120311101-3013330213121112-3223101030231222-2213302113310223-0131211101133212"></a>

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

<a id="canonical-0221123331110320-3320122231123210-3131230212131101-0112331213320032-2203200322311022-3232213310302202-2311012013333300-2211113002201313"></a>

## Direct properties — header_transformation / 101232110110 / 3

- [default_header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-2202322213122002-0012033032233213-0030122021322013-3223102103300330-1323312223001023-0112231202203101-3133223023122033-1220233132213220): complete subsection reference.

- [preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1112032312121213-0223300322233013-1131303210001112-2002031310303110-2310232330302000-1132101213302010-2123323323313131-3001322312032112): complete subsection reference.

- [proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-2030103111230230-3220223102021020-2123201320222030-3002330311130323-0013101011122133-0321021212020230-1212012033213331-0333130322103031): complete subsection reference.

<a id="canonical-1033111233203000-1030221310213121-2133323010220133-2223310113320023-3211033000122001-0221132013010223-3100330023112310-3032010031200221"></a>

## Next pages — header_transformation / 101232110110 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-2202322213122002-0012033032233213-0030122021322013-3223102103300330-1323312223001023-0112231202203101-3133223023122033-1220233132213220)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1112032312121213-0223300322233013-1131303210001112-2002031310303110-2310232330302000-1132101213302010-2123323323313131-3001322312032112)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-2030103111230230-3220223102021020-2123201320222030-3002330311130323-0013101011122133-0321021212020230-1212012033213331-0333130322103031)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2202322213122002-0012033032233213-0030122021322013-3223102103300330-1323312223001023-0112231202203101-3133223023122033-1220233132213220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331132130313331-2100202322121201-0122331310303122-2230120030133030-3001031101000302-0302230230131200-2330220133323103-0313002003111023"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 231330223100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1203100033001003-2023033130302013-2020032021213230-1033003332123122-2203313230120030-3322310200311122-2212202211232313-1233031023330101"></a>

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

<a id="canonical-3101232133333000-0033220220113121-0122113312302022-2000003111111303-2333320122101220-3023333322120023-2022022002310233-3102101330300323"></a>

## Direct properties — default_header_transformation / 231330223100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233100333220331-1021021001331223-1302013002131111-1330020233020213-3122332013033002-1122130133232021-2101222120212311-1031301111032103"></a>

## Next pages — default_header_transformation / 231330223100 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1112032312121213-0223300322233013-1131303210001112-2002031310303110-2310232330302000-1132101213302010-2123323323313131-3001322312032112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300001103320220-1131302302033010-0020101210123203-2223020231323202-3332223211110032-2233201112222012-2310330103033202-1131231121220302"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 102130201103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-0100231133331333-1003133303001111-1002222012133210-0030100022123222-3033310202333212-3311203113313332-1321121303231003-2122022313103200"></a>

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

<a id="canonical-0021002201333013-2302113220021200-0332102103212031-1123032031203303-2222221321313133-3211331112100112-1011121102102311-3310210031132312"></a>

## Direct properties — preserve_case_header_transformation / 102130201103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122010123012322-0113221130213301-1323210303102312-3302021000020321-1232323003113132-2122221102111221-0210332333133111-1001310313333210"></a>

## Next pages — preserve_case_header_transformation / 102130201103 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2030103111230230-3220223102021020-2123201320222030-3002330311130323-0013101011122133-0321021212020230-1212012033213331-0333130322103031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231120031132022-2110201123020313-3331023320303231-3101232310311000-2223121112201131-3313313133033100-1032122211112023-2301103100100120"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 101300230211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1100112311121310-1111011320313012-1102212131122030-0022023012111331-1212312231000301-3000202102120203-1320132202013023-1210111022200332"></a>

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

<a id="canonical-1320132322013022-3030211330233221-2213022031132201-3012301000322320-0120032220012022-1000020221310323-2022202312031011-2021220320231101"></a>

## Direct properties — proper_case_header_transformation / 101300230211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311021021300201-2311111032101132-0123113321103213-3102322323310000-1000332310332213-2023202003113020-2210331232113233-2010121021301102"></a>

## Next pages — proper_case_header_transformation / 101300230211 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1103132123031212-1312011311132313-1101103331230311-0211210100021021-0223211311301221-2201203120330131-0323223220202031-2113302232301212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330132302001320-2130233120030102-1102123031233230-2310300121313023-3001313013123032-1300301110021312-0210232202230001-1333210332322333"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 212013332123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3032332011020133-0202230002131001-2120102202310001-1023233133030213-2003100020033220-3320202030203220-3002212120000331-3022002120130101"></a>

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

<a id="canonical-3110321200003132-1123122221202202-1221221333001100-0331332213323022-2303232210001223-3123111320301233-0203231120100120-1001212013212011"></a>

## Direct properties — http_protocol_enable_v1_v2 / 212013332123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011231130223132-0000032232003021-0002230231011022-1001300100200002-2213203011132221-3220211101223323-2131322231311302-2103120030302123"></a>

## Next pages — http_protocol_enable_v1_v2 / 212013332123 / 4

- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0011132011130333-0210213032201022-0233112013130221-1333101332310100-3311123321233200-3302001330230303-1003013201132323-0021222312201002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123002323000131-2232320331233232-1100113110002002-2202300301321031-0133221020103213-1020020120111012-1210011331133323-0101120233001220"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 023011303300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0110022001120101-2313001121311312-1030320012233132-3103202132121130-1000310302113232-1032111101020033-3322221023021120-1233231021100002"></a>

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

<a id="canonical-1230232103213032-1302232111222333-1223131331023130-0211013323101003-3103012320303031-1022211212111111-1031321323111223-0112220213102121"></a>

## Direct properties — http_protocol_enable_v2_only / 023011303300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030100331233003-0111000202321120-3033123220001320-2201002021133222-2103103102120030-0023311311010210-2302213201220323-2210003001233331"></a>

## Next pages — http_protocol_enable_v2_only / 023011303300 / 4

- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2102301213102300-3130211303123110-1211031100221301-0101130110032303-3233220303001111-3333311120020322-3000202300233331-2213001132011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210022113030321-0300332333120102-0231311020123223-0220332213120212-3211112210133230-3232101111232013-2201332333203033-1220120103122231"></a>

## https_auto_cert.no_mtls — no_mtls / 312202132323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.no_mtls

<a id="canonical-3102331103101200-1300200020133012-3212201023102022-1033211130112302-0111213322130103-0233311301011312-2202033231000203-3231303332020023"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-3310010013223221-1302212312110310-0232202033023213-2031010213312333-1002032333030322-3032010001330022-2111311210112331-1013132103213130"></a>

## Direct properties — no_mtls / 312202132323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113130321232111-3201202100111130-3013311020111201-3220000131300001-2000013211102220-2121313211102001-3312212330020303-2232130011301101"></a>

## Next pages — no_mtls / 312202132323 / 4

- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3311320301322302-1303220013302033-0332111200020023-1101002220032022-2331300211113301-0130223132030323-1230211323100313-3210102020111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013101232231202-2230322113333123-2030113211302201-3213101010321113-3032322022303312-3102221003332210-0301113131130133-0101201320331302"></a>

## https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 122220311110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.non_default_loadbalancer

<a id="canonical-1321112301132100-3311122130222111-1030000333030100-2021301202103013-2233220213212003-3313332221010022-0010031220303131-2301302131022111"></a>

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

<a id="canonical-3323230111301312-0103232012022103-2310330122123323-3211133132323132-3111200022323010-2312211110033101-3310013232122132-1301313203302023"></a>

## Direct properties — non_default_loadbalancer / 122220311110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010213022332202-3121331013001120-3032111233223212-0132120110000230-0020132100113203-1300001132113301-1210223123213223-3213321321033212"></a>

## Next pages — non_default_loadbalancer / 122220311110 / 4

- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2221203212221022-2312321133113202-1222303012100333-2123103030313231-1012021303103113-3232300033113221-1000231113031103-1202131222211332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211202122102302-0323312033200000-1031003322332233-3113023332332102-3101232121312200-3232003131010032-2122002032322101-2032132201112030"></a>

## https_auto_cert.pass_through — pass_through / 212221303113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.pass_through

<a id="canonical-0201223000223123-2113111120003131-1233012100023311-3210301320110223-1121212130032223-1212331021310210-1020112320000113-3020132023001321"></a>

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

<a id="canonical-0223132301332321-2311200312212130-1022122221102023-0211203310301011-1013013023111330-2030223302200221-2213020233332012-3231102230131302"></a>

## Direct properties — pass_through / 212221303113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013322201130230-0121320131001202-2033312100132100-3012022132331132-0011332302222231-0233033111303200-2210223203033220-1103213103221133"></a>

## Next pages — pass_through / 212221303113 / 4

- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200312213221312-3212221122333302-2222331300223230-2131010120200320-1232331202213211-3131010301021010-0013011212110021-2101122311023230"></a>

## https_auto_cert.tls_config — tls_config / 032302231213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.tls_config

<a id="canonical-1303230000123212-0021300330000120-1323323023203020-1103122301012131-0032202000033001-1210020010202233-0333213200223212-3023003032320310"></a>

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

<a id="canonical-2103331113223123-0000030121312221-2000013123213333-3021213031200213-1330001001030211-2233200230211023-0201122201222333-1330121001023033"></a>

## Direct properties — tls_config / 032302231213 / 3

- [custom_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-3321303330313313-1023102100223210-0012100231302202-0313102310302120-0232233111333113-0013230032202120-3011301313332221-2320312110030121): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-3311200230002203-1203320312313000-0133330003131232-1221012211132131-3002320003012210-3111203202112020-2012321231102301-0200133112313231): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-3133012211221031-2320213030030231-3002213122221231-3223321103200220-3333232002031112-0033232130010001-3010101230033111-2331300100130002): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-2320333312023320-3223013123032113-2113232310323110-2013111100113321-1000101020100330-1322323010122013-2120003020300200-0132323201132013): complete subsection reference.

<a id="canonical-3032222300130131-3300011020023013-1133123322302132-0132213311110020-0121311122010221-3001302221032033-3211302112230233-3101121333210012"></a>

## Next pages — tls_config / 032302231213 / 4

- [https_auto_cert.tls_config.custom_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-3321303330313313-1023102100223210-0012100231302202-0313102310302120-0232233111333113-0013230032202120-3011301313332221-2320312110030121)
- [https_auto_cert.tls_config.default_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-3311200230002203-1203320312313000-0133330003131232-1221012211132131-3002320003012210-3111203202112020-2012321231102301-0200133112313231)
- [https_auto_cert.tls_config.low_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-3133012211221031-2320213030030231-3002213122221231-3223321103200220-3333232002031112-0033232130010001-3010101230033111-2331300100130002)
- [https_auto_cert.tls_config.medium_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-2320333312023320-3223013123032113-2113232310323110-2013111100113321-1000101020100330-1322323010122013-2120003020300200-0132323201132013)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3321303330313313-1023102100223210-0012100231302202-0313102310302120-0232233111333113-0013230032202120-3011301313332221-2320312110030121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012300333031223-1312130013102103-1311221022232103-1122131212003333-1332000030322023-3113232111233210-2310002213031310-2023222010030320"></a>

## https_auto_cert.tls_config.custom_security — custom_security / 333031222313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- https_auto_cert.tls_config.custom_security

<a id="canonical-2121302111320221-2012230300301012-2311303023023203-1103002100023222-3020323101311201-2022130033101333-2302123212332100-2111033211131313"></a>

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

<a id="canonical-0010222111203131-1221021100232111-0301103102310323-1311103201123100-3233221331010112-1302323231311032-3331122110031212-3230101131120312"></a>

## Direct properties — custom_security / 333031222313 / 3

<a id="canonical-2312122002111323-1223301230013000-1302200210223223-1302222010310231-0223132032132222-3220322131023301-1011101121201123-1313303330210321"></a>

<a id="canonical-3311110332330122-3100033222013100-2332223310331220-2032123122001213-2121221012121132-0231202110331111-1121021230201023-1300202001120030"></a>

## cipher_suites property — custom_security / 333031222313 / 4

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

<a id="canonical-3003120113200102-3002111313221130-0133033001233200-3011111231303232-1213022101003011-2301203110110000-1123330323013310-2313110332000130"></a>

<a id="canonical-3301003223122123-3112223232320220-2230230012023012-3301032013200013-3320232233103223-3023132210123200-1001321220202322-0233322100322011"></a>

## max_version property — custom_security / 333031222313 / 5

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

<a id="canonical-1130201201321330-2311101220300322-0010320303322202-3020331330020103-0313320332031302-1332001113201232-0100322232130021-3033122221211221"></a>

<a id="canonical-2101203031111112-2300320102001320-2222002132321133-2300203313213111-0100302232100131-0113113002111130-3131311333113132-2211330333101213"></a>

## min_version property — custom_security / 333031222313 / 6

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

<a id="canonical-2320110030003110-1323333332303011-0000313331221211-2300021022310211-2232203302300130-1303023201332112-2221011131322013-0223002122333322"></a>

## Next pages — custom_security / 333031222313 / 7

- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3311200230002203-1203320312313000-0133330003131232-1221012211132131-3002320003012210-3111203202112020-2012321231102301-0200133112313231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110023003323023-3202123303310300-1113330323123300-2223032200030303-0001200121030223-3113210321220303-3323330300100322-1230112010100301"></a>

## https_auto_cert.tls_config.default_security — default_security / 313020302222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- https_auto_cert.tls_config.default_security

<a id="canonical-0110221032200233-0123030013302202-2222311311010320-2301021000313300-3102233231330121-3020302203200213-0310331222003021-2102200010332102"></a>

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

<a id="canonical-2303110331121002-2312002331001231-2222022201103001-3223210121321223-1223020233130232-1323102131233320-3310220331132023-1021221121030213"></a>

## Direct properties — default_security / 313020302222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001013230203323-1323130010210303-2203002311010111-2113321101122032-3023021011102100-3103330313010031-1331023122230122-0113321210333231"></a>

## Next pages — default_security / 313020302222 / 4

- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3133012211221031-2320213030030231-3002213122221231-3223321103200220-3333232002031112-0033232130010001-3010101230033111-2331300100130002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020002131233332-3231211330031203-2133132233113001-3330322032111223-1202120332302312-0233221313002220-0213112001210203-2012212013220303"></a>

## https_auto_cert.tls_config.low_security — low_security / 123222211312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- https_auto_cert.tls_config.low_security

<a id="canonical-3031000220033311-1311330330220321-1132202101203223-1021112003302031-2211203020123203-2111113021312212-0033201203132121-2311112023003112"></a>

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

<a id="canonical-0321012131211320-2233323320232121-2030302232321130-2302312111020103-1022302211320133-1332310311031011-2102213000212331-1003001132333210"></a>

## Direct properties — low_security / 123222211312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301303000102111-2232231230010113-0303332000102033-3122131002313123-2220333110230111-1300110101331231-1213023101001010-0210031003031113"></a>

## Next pages — low_security / 123222211312 / 4

- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2320333312023320-3223013123032113-2113232310323110-2013111100113321-1000101020100330-1322323010122013-2120003020300200-0132323201132013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201210000011232-2022033201100200-2121101130202220-2012220320032132-3022223223020022-3323221000202210-3230021131320302-0100300312130332"></a>

## https_auto_cert.tls_config.medium_security — medium_security / 002203210320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- https_auto_cert.tls_config.medium_security

<a id="canonical-3002200010300312-1320300313121032-3231232210323321-3013330020232213-2120002311330010-2202330323320321-3031312320231020-3221303220120130"></a>

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

<a id="canonical-3100012113323000-2101133101030020-1121223113321331-2302312303302120-2111121113230311-0002120022231202-0031311310213211-0301233202122003"></a>

## Direct properties — medium_security / 002203210320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131122123201231-3310103303223310-1311230031021213-0001023010133003-1211133000300320-3113223302010120-3203301002231231-0032331333212310"></a>

## Next pages — medium_security / 002203210320 / 4

- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213131323122132-1211202321233130-1110203012033300-3100211033202121-0232033231020033-2101213332332133-2312110200203232-0013032121103233"></a>

## https_auto_cert.use_mtls — use_mtls / 001133231303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.use_mtls

<a id="canonical-1332230210300023-0212122101223322-2213010231110222-3003011321323021-1120110102101332-1210100313330022-2022320011011210-3333230221131233"></a>

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

<a id="canonical-1301210033301112-3301312300100020-3310031031220023-3231333030311200-1211222322210323-2112103001003230-2201230012120003-3030233122032023"></a>

## Direct properties — use_mtls / 001133231303 / 3

<a id="canonical-2120103111000201-1101032320311321-2220213112202003-1231211112010220-1302232113233012-0203123333211201-0123310021003032-0102101123212302"></a>

<a id="canonical-1133231022303122-3011132023221233-2131331131231322-0100010023003132-3031221110332130-1302303301033000-2202332303221212-2032030012220031"></a>

## client_certificate_optional property — use_mtls / 001133231303 / 4

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

- [crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-3022211310311000-2113002323211002-3101033102103010-3332313030200030-2222011331300210-0333231203022020-1133233222102312-3022332112023033): complete subsection reference.

- [no_crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-3031122111221032-3301130300001112-3312231022102031-0132103220103110-3030131110101030-1032113130131001-0002230301223303-1013210033303222): complete subsection reference.

- [trusted_ca](data-sources--http_loadbalancer--reference--group-019.md#canonical-2110202012002322-1331000322220012-0322003302202201-1202301120221220-2203322233330010-3232030010201113-3002212331002320-1322000223330312): complete subsection reference.

<a id="canonical-1323300012312011-1213223312232322-2321231323120111-0232232021321030-2200313032022112-2332300113322011-0302000213333300-0123230131320220"></a>

<a id="canonical-1211110003221031-3333333022132100-3330323023013112-2033222320213302-3023120003333001-0122113101110000-2130023123031120-3330303232221101"></a>

## trusted_ca_url property — use_mtls / 001133231303 / 5

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

- [xfcc_disabled](data-sources--http_loadbalancer--reference--group-019.md#canonical-1323300021111131-0210300012100323-1030213301331200-2010322123033110-3012201130320312-3303210030320222-2213320012331332-1211233202213302): complete subsection reference.

- [xfcc_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-2320103323203311-2032012110110030-2033213110213001-3023133310030011-1310311100100221-3213202133022321-1103031113223332-1323011332023031): complete subsection reference.

<a id="canonical-1220222132111002-0200302321333031-0101323223301033-1322320330320222-1012320033101131-2131012001012021-1110011213323003-3203200230013230"></a>

## Next pages — use_mtls / 001133231303 / 6

- [https_auto_cert.use_mtls.crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-3022211310311000-2113002323211002-3101033102103010-3332313030200030-2222011331300210-0333231203022020-1133233222102312-3022332112023033)
- [https_auto_cert.use_mtls.no_crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-3031122111221032-3301130300001112-3312231022102031-0132103220103110-3030131110101030-1032113130131001-0002230301223303-1013210033303222)
- [https_auto_cert.use_mtls.trusted_ca](data-sources--http_loadbalancer--reference--group-019.md#canonical-2110202012002322-1331000322220012-0322003302202201-1202301120221220-2203322233330010-3232030010201113-3002212331002320-1322000223330312)
- [https_auto_cert.use_mtls.xfcc_disabled](data-sources--http_loadbalancer--reference--group-019.md#canonical-1323300021111131-0210300012100323-1030213301331200-2010322123033110-3012201130320312-3303210030320222-2213320012331332-1211233202213302)
- [https_auto_cert.use_mtls.xfcc_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-2320103323203311-2032012110110030-2033213110213001-3023133310030011-1310311100100221-3213202133022321-1103031113223332-1323011332023031)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3022211310311000-2113002323211002-3101033102103010-3332313030200030-2222011331300210-0333231203022020-1133233222102312-3022332112023033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032103131132202-0010003331100221-2332222332112022-1223133012202313-3031210200303233-1122300023201302-0010020211230101-0001020003301331"></a>

## https_auto_cert.use_mtls.crl — crl / 013313301123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.crl

<a id="canonical-3102203323230331-2332331121111301-1031021010032123-2300331230310203-1222022330101121-0323133203101230-0222130213103103-0301011102032323"></a>

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

<a id="canonical-0121103120302312-0312213130323222-0001321210102323-0203130223120303-2103113310220112-3300330202311130-3321200221302012-0300320233210111"></a>

## Direct properties — crl / 013313301123 / 3

<a id="canonical-1233201310332222-2221100001120323-1110223122023300-1133322230022011-1210211211023130-3012333233123133-1121102021133312-3121012000020333"></a>

<a id="canonical-3023302311010321-2101032123013111-2322233333003012-1332100002113222-2133111123330222-3120203221122232-3203232202112231-1300211113031002"></a>

## name property — crl / 013313301123 / 4

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

<a id="canonical-2103303330320031-0300230330132101-0302013113033022-1033110220302020-3312020221310222-1330233211203310-1220031212013001-0330232300111100"></a>

<a id="canonical-3133220001110310-0232211002201320-3133332111031332-0220223113113310-2133203231000033-1313310112002111-0311320213323011-1123320202221202"></a>

## namespace property — crl / 013313301123 / 5

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

<a id="canonical-0033223200330302-1110121331232322-1202010122133133-3310022323321310-2233331303213132-2031312313311312-0021323223030203-0222212332201322"></a>

<a id="canonical-2221312030322323-2031311013103011-1223213013130110-3122300221310033-2333300032330122-1033002101201223-3003121321320112-1322131311200011"></a>

## tenant property — crl / 013313301123 / 6

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

<a id="canonical-1331223222303333-2012112023002213-3332022333003203-1210101320112123-1203202310320331-2020110113111323-3220202221132110-0322232211133021"></a>

## Next pages — crl / 013313301123 / 7

- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3031122111221032-3301130300001112-3312231022102031-0132103220103110-3030131110101030-1032113130131001-0002230301223303-1013210033303222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302003102203212-1022213003303123-2333000021003133-0311202230111321-0010200103133200-2132301313130021-2123112222301331-2220003020330012"></a>

## https_auto_cert.use_mtls.no_crl — no_crl / 001111300313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.no_crl

<a id="canonical-0112313230133003-1022013212110230-3120323121133100-3202212220310113-0100231332310033-1323020012033203-1232113321333232-1233031323322132"></a>

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

<a id="canonical-3333323210311302-3110201303031020-1220003200101131-3332212311121212-2122330032031101-1130211011302333-3130131112112113-2103031301223301"></a>

## Direct properties — no_crl / 001111300313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200302031001211-2210212001300200-0102331230002000-3311001100322000-3321230103310212-0001331110323121-2032222330302310-0000101211213121"></a>

## Next pages — no_crl / 001111300313 / 4

- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2110202012002322-1331000322220012-0322003302202201-1202301120221220-2203322233330010-3232030010201113-3002212331002320-1322000223330312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311112330100300-2021213103233211-0211001220311002-2010021320132112-3010013303222212-0102030202312231-0331112330213311-3203321000123100"></a>

## https_auto_cert.use_mtls.trusted_ca — trusted_ca / 233110002103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.trusted_ca

<a id="canonical-1100231002200303-3331321100303210-0000032312113220-2003210220131330-3213130311333000-3323201330013313-3013230002233120-2323130010333130"></a>

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

<a id="canonical-1332121112200111-0130030033300131-3022313212101133-0313221312221012-0000101012200323-0221002321323110-2011130302132003-1123220221230000"></a>

## Direct properties — trusted_ca / 233110002103 / 3

<a id="canonical-0110311233213233-0100322222011201-0223322012022021-0022011031333233-1313202010331221-2201123000232313-2202102021301302-1322032103010002"></a>

<a id="canonical-1112333130031022-1222221000321120-3121122233031020-3002130322011021-0100200322322301-3010120102203321-3301230332223031-3021201021221223"></a>

## name property — trusted_ca / 233110002103 / 4

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

<a id="canonical-0123123101312031-3022300111202230-1230330311121213-3032110023123113-0301012213012032-3232302220332131-2302020133011210-1203013203313012"></a>

<a id="canonical-3200033132102330-2323002332033213-1203200200221030-0121203210011303-0033331021022112-3303031210000320-3131222021220310-2303302110200132"></a>

## namespace property — trusted_ca / 233110002103 / 5

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

<a id="canonical-3302300032003133-1110221312331012-3000020301331212-3203012123210202-3100331312021021-1120330023322232-0023232332201321-0022002203311012"></a>

<a id="canonical-2103222001322200-1102302221321113-2213302031103200-1313303023203233-0102311313133333-0021130021331203-1012031112222012-1101021323022311"></a>

## tenant property — trusted_ca / 233110002103 / 6

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

<a id="canonical-1032120111011231-0120111032010031-1001332101133310-0102303321313233-2303320111221111-0021201311133031-2012120313300200-2112012102330011"></a>

## Next pages — trusted_ca / 233110002103 / 7

- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1323300021111131-0210300012100323-1030213301331200-2010322123033110-3012201130320312-3303210030320222-2213320012331332-1211233202213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012101331013222-2222323313321230-2213102311202203-0301320023301312-0231112302220010-0211201232203211-1010313132222203-2002302222322120"></a>

## https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 001311112133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-0211231130230123-2232111000322000-2010130333011003-1120020001320111-3321011130302113-3022022003032023-0313003001331000-0333330313212130"></a>

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

<a id="canonical-3210013031112030-2131103021233230-2120331223002110-3312333003203203-1023333313321331-3023031103132223-0031021233023013-1221201111301200"></a>

## Direct properties — xfcc_disabled / 001311112133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100122332110310-1101202113223213-2323003130200020-3203302232220223-2220333100121011-1023122130202332-3123113331233223-1101223311310320"></a>

## Next pages — xfcc_disabled / 001311112133 / 4

- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2320103323203311-2032012110110030-2033213110213001-3023133310030011-1310311100100221-3213202133022321-1103031113223332-1323011332023031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233101301210011-2131022021232032-1121330213103121-2311021301332330-0000000123012313-3330233301323120-1202201133220212-0003333213310231"></a>

## https_auto_cert.use_mtls.xfcc_options — xfcc_options / 032110202111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.xfcc_options

<a id="canonical-3222033120312330-0220010302223110-3321013122203130-1222101301333322-1201332131110030-2011012200003302-0132223022330220-2020331330232323"></a>

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

<a id="canonical-1013233201031003-2232022030332010-0003001003012110-3023213230023031-2003123103223111-3000230013322020-0100211110001313-2112133320023333"></a>

## Direct properties — xfcc_options / 032110202111 / 3

<a id="canonical-3003200023111201-0000122322231303-3320122301312123-3213032301203302-3012031333332300-3202110103221133-0102001013320202-2131013231303010"></a>

<a id="canonical-2123003200020012-3031323333201230-1210101233032221-3332300020102331-3212001323321223-3021312130321120-0122123130211322-1312020111021030"></a>

## xfcc_header_elements property — xfcc_options / 032110202111 / 4

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

<a id="canonical-2112313100100230-1213102101201023-2020133002221003-1032303322301213-0021301001223103-1330233320231131-1301200311303222-2100300013132333"></a>

## Next pages — xfcc_options / 032110202111 / 5

- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3322121320310111-0021311031133102-3002213332023220-1023112102022302-3222311022023111-1222120222313303-1310032321220233-2003030001231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201123310322022-0120201111011021-0131311012133321-1303210002312131-1003221301313011-3223120033021300-1232033202113221-1201323120021031"></a>

## js_challenge — js_challenge / 213012032332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- js_challenge

<a id="canonical-1330313013020313-0000301213230310-2113112320001221-3102132110313013-3033321333122212-3213022133202312-1321331102121220-0332330321122003"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2121210011120231-1233022001120302-1210121121222200-3113213013100011-1211023212132201-0002202032313222-1201112033130211-1003023200113121"></a>

## Direct properties — js_challenge / 213012032332 / 3

<a id="canonical-3102210100303112-1020303021310313-2133302230122021-0030133322212223-2112033032132323-0231022220122101-3223033211310120-1010211100002121"></a>

<a id="canonical-0333000123111333-1013211202023121-2032111012203131-1022032120331321-0322332020020230-2211233022021120-0012110222010211-0221333302011322"></a>

## cookie_expiry property — js_challenge / 213012032332 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2322220013032133-3301003321120221-1102202030101101-1210303000213001-1323100103213320-2220110321031000-3031100201103013-2223023021000313"></a>

<a id="canonical-2113001001200120-2331010032000201-2313302320312003-0011003323313302-2223223030002021-0030313220132021-1333301323323022-2100002010011200"></a>

## custom_page property — js_challenge / 213012032332 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3032322122120211-2321313331012022-2313333033223322-0221212230003332-2122011202201200-2032300032312113-3110023321303313-0013121333103132"></a>

<a id="canonical-3330130131020323-3101131110121031-2210332123301321-3232202330121313-2320330023233231-3232211323213300-1001033020101101-1120131031101020"></a>

## js_script_delay property — js_challenge / 213012032332 / 6

Type: `"number"`. Computed.

Delay introduced by Javascript, in milliseconds.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3332203010123233-1313023113233110-1202221013001323-2131120223221332-1012300122310020-1101331133301103-3112103221123303-1000113011313002"></a>

## Next pages — js_challenge / 213012032332 / 7

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001103022130031-0101011122330022-2001222323012320-1303303003023022-3201023112313012-1321011131132302-1001111120021202-0030022332120110"></a>

## jwt_validation — jwt_validation / 213000213213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- jwt_validation

<a id="canonical-2333203013123213-3021320330302333-0312333311010330-3330210123032031-2133202033212130-1313333232130123-2011322310210232-0122201112222231"></a>

Type: `"single"`. Computed.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Upstream description:

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

<a id="canonical-1323001131021321-3022230311030222-0300010131000030-2312213132032101-3333222013030222-1323130211131102-1003230100210103-1222133302333003"></a>

## Direct properties — jwt_validation / 213000213213 / 3

- [action](data-sources--http_loadbalancer--reference--group-019.md#canonical-0330000101210233-3210012112023210-0100120200333133-1031102031030203-3332320022130222-1321213230203003-1303301123013331-1033300310312101): complete subsection reference.

- [authorization_server](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323202302020011-1013203000333112-1211331220023232-3232211311013021-0331110330133333-1013322200202110-2010313023231030-3213031121002022): complete subsection reference.

- [jwks_config](data-sources--http_loadbalancer--reference--group-020.md#canonical-1221233210212303-1303230013013133-2330013321313313-2331033212023133-0023202113210300-1000220111211222-1323110320030023-2001232301101300): complete subsection reference.

- [mandatory_claims](data-sources--http_loadbalancer--reference--group-020.md#canonical-1003113313333323-2312032022332312-2213233012232010-3132202210222310-1022032321220123-3210013112210001-2012002222033322-0232120300110111): complete subsection reference.

- [reserved_claims](data-sources--http_loadbalancer--reference--group-020.md#canonical-0303133023120031-0231031101210213-1112203310123213-2311122021023032-2031322322111201-0112323012111131-2112103030122332-0312113221123201): complete subsection reference.

- [target](data-sources--http_loadbalancer--reference--group-020.md#canonical-3320100111233100-2113302231232203-3303311321213110-0121302310123021-1231113212032011-3001021101020021-2023122021203001-2010102033133321): complete subsection reference.

- [token_location](data-sources--http_loadbalancer--reference--group-020.md#canonical-2101012003211203-1333223011213101-1311003002102303-2000301122201222-1111130223313030-0211031013100112-2032203303233133-3303103031322301): complete subsection reference.

<a id="canonical-1110111023311212-2011302211130221-1001022322033110-0121023113023320-3231231222023203-2112231022133332-3200102132331002-3133300123311011"></a>

## Next pages — jwt_validation / 213000213213 / 4

- [jwt_validation.action](data-sources--http_loadbalancer--reference--group-019.md#canonical-0330000101210233-3210012112023210-0100120200333133-1031102031030203-3332320022130222-1321213230203003-1303301123013331-1033300310312101)
- [jwt_validation.authorization_server](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323202302020011-1013203000333112-1211331220023232-3232211311013021-0331110330133333-1013322200202110-2010313023231030-3213031121002022)
- [jwt_validation.jwks_config](data-sources--http_loadbalancer--reference--group-020.md#canonical-1221233210212303-1303230013013133-2330013321313313-2331033212023133-0023202113210300-1000220111211222-1323110320030023-2001232301101300)
- [jwt_validation.mandatory_claims](data-sources--http_loadbalancer--reference--group-020.md#canonical-1003113313333323-2312032022332312-2213233012232010-3132202210222310-1022032321220123-3210013112210001-2012002222033322-0232120300110111)
- [jwt_validation.reserved_claims](data-sources--http_loadbalancer--reference--group-020.md#canonical-0303133023120031-0231031101210213-1112203310123213-2311122021023032-2031322322111201-0112323012111131-2112103030122332-0312113221123201)
- [jwt_validation.target](data-sources--http_loadbalancer--reference--group-020.md#canonical-3320100111233100-2113302231232203-3303311321213110-0121302310123021-1231113212032011-3001021101020021-2023122021203001-2010102033133321)
- [jwt_validation.token_location](data-sources--http_loadbalancer--reference--group-020.md#canonical-2101012003211203-1333223011213101-1311003002102303-2000301122201222-1111130223313030-0211031013100112-2032203303233133-3303103031322301)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0330000101210233-3210012112023210-0100120200333133-1031102031030203-3332320022130222-1321213230203003-1303301123013331-1033300310312101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100113233210203-1201203213113120-1100011112031102-3023030302323323-0203123200033111-0210301310023103-1002031100323321-1020000310112202"></a>

## jwt_validation.action — action / 332123313103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- jwt_validation.action

<a id="canonical-2322133200003120-0020322322031123-0112011202201012-2132213001020232-1002201123203303-2222321031110231-3003001213332310-0320110230013333"></a>

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

<a id="canonical-1000323131332302-3233320132002212-2002001123101323-1023103201023120-1213321323112013-0322201310331320-3032211002103130-3222021031121003"></a>

## Direct properties — action / 332123313103 / 3

- [block](data-sources--http_loadbalancer--reference--group-019.md#canonical-3310212213010012-1310313200231023-1002132133301020-2010130313022021-2211022313021303-3012312000013030-2123233331213130-0213302023010323): complete subsection reference.

- [report](data-sources--http_loadbalancer--reference--group-020.md#canonical-1232021203130011-1200023112011311-2022031033310210-0111312331122321-2101232202012233-2322330321033131-2003100213012100-3022220031123311): complete subsection reference.

<a id="canonical-1121123303331231-1000003131010320-0303301231320331-0213210010103031-1300322231030001-1113030131320131-2033011231013022-2300301022231331"></a>

## Next pages — action / 332123313103 / 4

- [jwt_validation.action.block](data-sources--http_loadbalancer--reference--group-019.md#canonical-3310212213010012-1310313200231023-1002132133301020-2010130313022021-2211022313021303-3012312000013030-2123233331213130-0213302023010323)
- [jwt_validation.action.report](data-sources--http_loadbalancer--reference--group-020.md#canonical-1232021203130011-1200023112011311-2022031033310210-0111312331122321-2101232202012233-2322330321033131-2003100213012100-3022220031123311)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3310212213010012-1310313200231023-1002132133301020-2010130313022021-2211022313021303-3012312000013030-2123233331213130-0213302023010323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
