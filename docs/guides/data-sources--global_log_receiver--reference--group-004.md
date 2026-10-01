---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-2230122120311331-1030031311332110-3231130303013330-2202013132020202-2201012203333310-3112223011103232-1231303030100131-2030110200203122"></a>

## trusted_ca_url property — use_tls / 302210111102 / 4

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-0210031033333100-3133001220230203-2132103201121222-2212031203000000-0323011113111310-0223101100211201-3130132031110000-2012211313201311"></a>

## Next pages — use_tls / 302210111102 / 5

- [qradar_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-3322301302112011-3033121332333012-0200331303130332-2222122221010101-0133233123032301-3122001023223333-3012013311301003-2311002323231113)
- [qradar_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-3033133203030300-0320022320111230-3121103002223122-2122001121220300-3311303212110211-3320101231000322-2200213311010132-1203010311323000)
- [qradar_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-0123123233012122-2113301330223202-2121020331223201-2212321333001230-1013330210322033-3012011223301023-1101221120323310-1121100232112320)
- [qradar_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-1020030030331112-0002020320132001-3321213311203211-2131010110102120-2030222032200022-1112312303103323-0211132122023330-2002320301320102)
- [qradar_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-3030223213022232-2013002132203010-1232301312001002-3012003330012130-1332310031333330-0313232120113211-0310122233131202-2332020002320113)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113)
- [qradar_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-1321113003023032-3133333102202200-3330133200020201-2210231001033103-1113301003303003-0302111220002111-0312310133310112-2033030022002322)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3322301302112011-3033121332333012-0200331303130332-2222122221010101-0133233123032301-3122001023223333-3012013311301003-2311002323231113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002213322211003-1201213103002103-2222121021330201-2120120103202211-3202302212111221-1021102232102332-3030310121030103-0323303231023100"></a>

## qradar_receiver.use_tls.disable_verify_certificate — disable_verify_certificate / 323303021313 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.disable_verify_certificate

<a id="canonical-1010110110313311-3022111300213103-2021321201103031-0010220303330001-3120330212103221-3311222332331301-2313001223210133-3102112211323120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

<a id="canonical-3002013011312032-1013120013133031-1101001102202120-1303300033320000-1213222003302033-1001020031312201-3312211031201122-3001300013203032"></a>

## Direct properties — disable_verify_certificate / 323303021313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302013123012331-0233231322301031-3012230233300012-2011011023130230-3310000122220221-2313030123011010-1213223102312120-2212132302210110"></a>

## Next pages — disable_verify_certificate / 323303021313 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3033133203030300-0320022320111230-3121103002223122-2122001121220300-3311303212110211-3320101231000322-2200213311010132-1203010311323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230331130203301-0331003130322130-0312220221231230-3002010100303231-2311220302101220-0311203211002231-2220213320111302-0122320230010113"></a>

## qradar_receiver.use_tls.disable_verify_hostname — disable_verify_hostname / 010110033323 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.disable_verify_hostname

<a id="canonical-0033032130120233-2223021003101311-3101310203222221-2100201013200230-2312010311131120-1121120223111102-1210313300121101-0222021301301201"></a>

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

<a id="canonical-0203230331130213-3032132122030031-1101201311203123-1023323331330203-3323103222120301-0330331211030301-0002112213011011-2003231200301323"></a>

## Direct properties — disable_verify_hostname / 010110033323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213302300123230-0021120100001013-2321310230022111-2131102233031103-1110323130110303-2003300302023100-3330001120012111-2112221210331311"></a>

## Next pages — disable_verify_hostname / 010110033323 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0123123233012122-2113301330223202-2121020331223201-2212321333001230-1013330210322033-3012011223301023-1101221120323310-1121100232112320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212102122101211-2312001212331301-2211103013033022-1232331301023023-0331032021323213-3302220321131220-3221032103002022-3113201101011233"></a>

## qradar_receiver.use_tls.enable_verify_certificate — enable_verify_certificate / 130111022011 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.enable_verify_certificate

<a id="canonical-3103312302111301-1231132032330333-3021000133010132-1332232332313323-1211033212323012-0210102013132320-3201021000322220-3212331133103003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

<a id="canonical-2230303210000100-1122302301220222-1221203202202213-3200023300111010-3332030313130323-2212210002132200-1220331121323320-1332200120220322"></a>

## Direct properties — enable_verify_certificate / 130111022011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202103133233320-3320101102233200-2132322212200223-3310131313133101-2100131012122331-1132002032332232-1100213022210120-2303313020032111"></a>

## Next pages — enable_verify_certificate / 130111022011 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1020030030331112-0002020320132001-3321213311203211-2131010110102120-2030222032200022-1112312303103323-0211132122023330-2002320301320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033100031312003-2200110211223012-1121011132312303-3222132022202302-0103223213303123-1133211101300030-0313101201101120-3303331011212113"></a>

## qradar_receiver.use_tls.enable_verify_hostname — enable_verify_hostname / 000302001122 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.enable_verify_hostname

<a id="canonical-2332003033202112-1303203011320320-2331011201020301-0322103001220033-0302010232213302-1223032102123333-0003313100001303-2303001121133003"></a>

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

<a id="canonical-3332021121011320-0112311222112001-3201313310020313-2130230331130130-3003201112211320-1210023010033311-1031232221000331-2300031101203320"></a>

## Direct properties — enable_verify_hostname / 000302001122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103020000223133-0030121122123000-3222030113113321-3022230011333112-0231100221211233-0122121212201311-2103000033022310-1112003202331112"></a>

## Next pages — enable_verify_hostname / 000302001122 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3030223213022232-2013002132203010-1232301312001002-3012003330012130-1332310031333330-0313232120113211-0310122233131202-2332020002320113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120212030002330-2200113330121333-0220300032203312-0121102311021321-2300032303231100-2331300031320300-2011200123031222-2203200020020213"></a>

## qradar_receiver.use_tls.mtls_disabled — mtls_disabled / 122303121020 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.mtls_disabled

<a id="canonical-1031002111110322-0020113230113332-0323211303022220-0203030213301032-0003032131113013-0300110032210130-2320310122133121-1223213301132030"></a>

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

<a id="canonical-3122303213112213-0102312333311033-0002303210332232-1132321210123131-3101121102332031-1223223032220232-3300322330122120-0010311332213303"></a>

## Direct properties — mtls_disabled / 122303121020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321301220011300-0132232220112333-1030333203012010-1232033001031030-0302002022222011-0321133310231300-1023312200333131-0112313031020301"></a>

## Next pages — mtls_disabled / 122303121020 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002333101133310-1222232103112033-2220032032130020-3020022223101332-3212210133112303-1211212311332033-1031021230232201-3310110012331002"></a>

## qradar_receiver.use_tls.mtls_enable — mtls_enable / 012113213231 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.mtls_enable

<a id="canonical-1102000323013231-3011332031221221-3223322211322130-3133331323223121-3232303103233103-0331221222131112-1223101302213111-3221301213001220"></a>

Type: `"single"`. Computed.

MTLS Client config allows configuration of mTLS client OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3330100222333030-3313203013121123-0022222120030120-3131213200133033-0322112130002112-1031111132232322-1212200210213032-3113120021333330"></a>

## Direct properties — mtls_enable / 012113213231 / 3

<a id="canonical-1023232113213120-0000322100312001-2123103300323311-1302120103233201-0201210131220111-2232202012112303-2122203032030310-1003310133231322"></a>

<a id="canonical-1303201212223303-2103310101211301-2223312203100100-1313320233133011-0122123301021123-0001013022133203-2200230103002123-2310123103232320"></a>

## certificate property — mtls_enable / 012113213231 / 4

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302): complete subsection reference.

<a id="canonical-2113213100202011-0312233320001233-2320121220220000-1010213103300231-0222010232202331-2232312211233221-1130033013110013-3122212312222033"></a>

## Next pages — mtls_enable / 012113213231 / 5

- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231133113310031-3213302313301300-3313110333020022-2332203100231332-1021301001130003-1113311233113002-3030013010020130-1002311133012020"></a>

## qradar_receiver.use_tls.mtls_enable.key_url — key_url / 322310303131 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113)
- qradar_receiver.use_tls.mtls_enable.key_url

<a id="canonical-3013222211230321-0300321101212300-0131201130103201-0110310111322300-1322020301322001-0002300130110132-1012112222220011-1002132110023013"></a>

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

<a id="canonical-3331132112322300-3133002111031222-0322312300230021-1011030030200102-3103233323213233-3313003010200100-3220030211303201-1331030333030031"></a>

## Direct properties — key_url / 322310303131 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3200100200210113-0202200100333101-0210100122110030-0203022302323032-1231112033230013-3023002002031003-0202102301122020-3021323223312111): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1332033023123302-2031003131031113-2323230312311202-0023312132113111-0000230321201033-3213103321131233-1032100022300220-2131301020321220): complete subsection reference.

<a id="canonical-2200232032322230-2023010322102011-3100133202320302-0110120321021021-2311012012333031-3121031322302022-2032333320211022-3222102030130331"></a>

## Next pages — key_url / 322310303131 / 4

- [qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3200100200210113-0202200100333101-0210100122110030-0203022302323032-1231112033230013-3023002002031003-0202102301122020-3021323223312111)
- [qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1332033023123302-2031003131031113-2323230312311202-0023312132113111-0000230321201033-3213103321131233-1032100022300220-2131301020321220)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3200100200210113-0202200100333101-0210100122110030-0203022302323032-1231112033230013-3023002002031003-0202102301122020-3021323223312111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211320312230020-2001013102310031-2212130222221123-1330212013021301-0101112320303310-0232121102223230-2223322032232210-3330120233203112"></a>

## qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — blindfold_secret_info / 031333202003 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113)
- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302)
- qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0033312112311302-2211223000112211-3020201001322220-0313222322010323-2310213203232110-0301332033031013-0211220200102303-0310133302001210"></a>

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

<a id="canonical-2000010310320312-0320110210021012-1012212210221302-1233003132201111-1110022103312302-1122010032123222-2013320311201122-0031111013230202"></a>

## Direct properties — blindfold_secret_info / 031333202003 / 3

<a id="canonical-0322330321221331-1013132123001032-2101210332311030-0013320031333000-2032313230201312-0030310231233321-1331013030230103-3302233212110333"></a>

<a id="canonical-3212013012103203-2333203001332310-0123101222203221-3302213031221110-0223211123232321-2033100030023022-2013330101102102-2103311120013023"></a>

## decryption_provider property — blindfold_secret_info / 031333202003 / 4

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

<a id="canonical-3221232030221232-3013021202033023-2322132321313123-2200002001232131-0301103100332021-3022330111022110-0100131320022201-0002231122112022"></a>

<a id="canonical-1102220312222310-3203330013130010-2331003123200211-2022231312022210-1003030030101122-1311200131332032-1131100122310121-3001001321130300"></a>

## location property — blindfold_secret_info / 031333202003 / 5

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

<a id="canonical-1202122332200302-2222131032210313-1003230301231330-0123303132113322-2122113331120122-2311012020003120-2221001130132223-3032000321201012"></a>

<a id="canonical-2132332122301332-2032233020132230-3013030322321222-3221212321121220-3023223111003123-1122102312332211-2223010110023133-1201331201010212"></a>

## store_provider property — blindfold_secret_info / 031333202003 / 6

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

<a id="canonical-3311020220300120-1233231030101311-1200001232212302-0200010322201131-0220232011303100-3110302331010101-2112300032323012-0221300330131300"></a>

## Next pages — blindfold_secret_info / 031333202003 / 7

- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1332033023123302-2031003131031113-2323230312311202-0023312132113111-0000230321201033-3213103321131233-1032100022300220-2131301020321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210033330212023-2221030311030133-2033222220123110-1310300233221331-3113331211301211-0212112010312200-2130202001310123-1031313010203002"></a>

## qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info — clear_secret_info / 330110313121 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113)
- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302)
- qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-2201210312331202-0101132022131022-2222010023212102-2023013013203102-3312023322011122-3033102011133110-2031032213123303-3132011310320302"></a>

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

<a id="canonical-2112211110102312-0022130012220101-0302230003130223-0221301230031303-2030330132221302-0332321022112130-2023020032123122-0012110010310331"></a>

## Direct properties — clear_secret_info / 330110313121 / 3

<a id="canonical-0133230231030101-2220032310201011-3301300030132332-3121222122111013-2002203001121223-0100022131033210-1112000031321221-1230202330333311"></a>

<a id="canonical-2031203230310121-3313011231210322-3022013323300121-3310001323100131-1033022010310201-1210020112223201-2222230130223320-1200212033332311"></a>

## provider_ref property — clear_secret_info / 330110313121 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3133211121033330-0233013321011211-0020022001231222-3330023303233130-0322030301001130-2223303333312030-1110033220130211-2000002233032102"></a>

<a id="canonical-0221321022322301-2112021011000013-0023003223301010-1123033210202220-1201212321213111-1213311012220302-3310103013232001-3333020102021201"></a>

## URL property — clear_secret_info / 330110313121 / 5

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

<a id="canonical-3231012102132213-1201013122123112-3123330302130132-0322030212011030-0032102000012301-0233103321013212-1220011103233311-1231012133133330"></a>

## Next pages — clear_secret_info / 330110313121 / 6

- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1321113003023032-3133333102202200-3330133200020201-2210231001033103-1113301003303003-0302111220002111-0312310133310112-2033030022002322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332202301002003-0110221113202221-1312202031103221-0303020232223111-2232103233132233-0301001300130213-2323201011032312-0330300231322232"></a>

## qradar_receiver.use_tls.no_ca — no_ca / 121033303320 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.no_ca

<a id="canonical-2022312312103013-1230211221100032-2332102120220202-2023002223022230-3202111112303012-1120332111111212-0022023231311011-0001333322320123"></a>

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

<a id="canonical-1002332032002301-2201132303102333-1233323131311302-0011221200311111-3132133131200312-0220333203012210-0230232313130330-1131302212131133"></a>

## Direct properties — no_ca / 121033303320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313032103201302-1222212201220301-2130332100120303-3012323230223300-0003233320123012-2101003130232033-0333231030221102-0212023212001302"></a>

## Next pages — no_ca / 121033303320 / 4

- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110121300221033-0102231131001221-3123001021310210-3201032000132233-3310203021301202-1103030220232011-3013110311331032-3033332220310033"></a>

## request_logs — request_logs / 133132331130 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- request_logs

<a id="canonical-1032022001202310-0203211112112032-3223133020101320-1122331023331223-0011221222002323-0101233233111212-1002013321121320-1323300223110231"></a>

Type: `"single"`. Computed.

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Upstream description:

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-sampling_choice": "[\"sampled\",\"unsampled\"]"
}
```

<a id="canonical-1030313233000223-2000320022111203-1001212320303310-2000332232303011-2300130023003011-1130233220310120-2131110311222230-2130302212000000"></a>

## Direct properties — request_logs / 133132331130 / 3

- [sampled](data-sources--global_log_receiver--reference--group-004.md#canonical-3200100232122302-0211213120103021-3202321000131311-3212012110223110-1203323122230000-3020110233311101-1212230131233030-0233113222101002): complete subsection reference.

- [unsampled](data-sources--global_log_receiver--reference--group-004.md#canonical-1233211210202013-1333120132123330-0331113102232220-0202021333003133-1203103032313102-2012031101003210-2213123200023122-0233131311333203): complete subsection reference.

<a id="canonical-0211002033231213-2020020330123123-1323213013233221-3113310202031120-1122201220030012-3113211032001021-3312121001002002-0133333222322120"></a>

## Next pages — request_logs / 133132331130 / 4

- [request_logs.sampled](data-sources--global_log_receiver--reference--group-004.md#canonical-3200100232122302-0211213120103021-3202321000131311-3212012110223110-1203323122230000-3020110233311101-1212230131233030-0233113222101002)
- [request_logs.unsampled](data-sources--global_log_receiver--reference--group-004.md#canonical-1233211210202013-1333120132123330-0331113102232220-0202021333003133-1203103032313102-2012031101003210-2213123200023122-0233131311333203)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3200100232122302-0211213120103021-3202321000131311-3212012110223110-1203323122230000-3020110233311101-1212230131233030-0233113222101002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303102310102020-2003000022030022-0033321202111222-2020303111111300-1131100120010330-1232222212210012-1230321121230223-1111201102101110"></a>

## request_logs.sampled — sampled / 302112133231 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123)
- request_logs.sampled

<a id="canonical-0213000030003023-1302322002320330-1101232212233301-1333011030312222-0232332313312220-3002233121133303-1131323201303233-0201300101121120"></a>

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

<a id="canonical-1322023021221223-3021313032322133-0121011123312200-1102220333323120-3202200330011221-1212112332132001-3331033221003132-2232100002322010"></a>

## Direct properties — sampled / 302112133231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122003212213231-2133331313121322-0020203233123332-0211021113101220-3212130112013030-3211021021120100-3100011302232033-1031031330211010"></a>

## Next pages — sampled / 302112133231 / 4

- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1233211210202013-1333120132123330-0331113102232220-0202021333003133-1203103032313102-2012031101003210-2213123200023122-0233131311333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333220123132032-0212231223330222-0231303011102021-0023132133102233-1123323322132020-2222213221120330-3131132011323321-3210333132122123"></a>

## request_logs.unsampled — unsampled / 321231233121 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123)
- request_logs.unsampled

<a id="canonical-2302132121301121-0112133011310222-1221312213031332-3202011003220203-1221000010321133-3012123231311323-1213012233300221-1012211031332330"></a>

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

<a id="canonical-2033122330013222-3233331122203010-0111130131313332-0020202303323111-0313130130022322-0011300202323123-0113122000102301-2311212230013000"></a>

## Direct properties — unsampled / 321231233121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013020011313030-0213131300332221-3110213221120212-0312030313201131-0300302001212120-0313120313122332-0132111102330021-3311100132333333"></a>

## Next pages — unsampled / 321231233121 / 4

- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133231103322102-1213002210131131-0201302000222022-0320332113033310-3123333100301001-1020322133013213-1321032131221031-2003033321230120"></a>

## s3_receiver — s3_receiver / 113113102322 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- s3_receiver

<a id="canonical-3212111221313232-1220032220002121-1013323021001210-2102032312000332-1100102021232231-2013023201232020-2333203332302002-2022012110300003"></a>

Type: `"single"`. Computed.

S3 Configuration for Global Log Receiver.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2100000002021300-1022133020231232-3023121121110203-2110000123232133-0212030311011103-1022303320123012-1130103321123021-2112311100100102"></a>

## Direct properties — s3_receiver / 113113102322 / 3

- [aws_cred](data-sources--global_log_receiver--reference--group-004.md#canonical-0003321233311003-3301113211012122-0103013321010222-3300113211001131-0130113110200303-1032313231101020-3131112000301032-3333012230303003): complete subsection reference.

<a id="canonical-3130112022220123-0322010213101030-2030001220321011-3222113132131011-3021021310310021-2112332203022012-2013312231100122-0030030133210210"></a>

<a id="canonical-2200123331133032-2233110200330301-3310200020213203-1021223210233132-1222032312321330-0210000111101001-2313220332011103-2303121111221000"></a>

## aws_region property — s3_receiver / 113113102322 / 4

Type: `"string"`. Computed.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Upstream description:

AWS Region Name.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022): complete subsection reference.

<a id="canonical-0010213232000030-1120300322233310-2021103020303311-3220221331323312-0210232131011112-3122323121311010-1032300103222012-2032010201003301"></a>

<a id="canonical-0020302023130212-1222230012212121-2213201122030222-1013330103131031-0032022221132113-3002210121131033-0113212220010110-1312333231320013"></a>

## bucket property — s3_receiver / 113113102322 / 5

Type: `"string"`. Computed.

S3 Bucket Name. S3 Bucket Name.

Upstream description:

S3 Bucket Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  }
}
```

- [compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203): complete subsection reference.

- [filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330): complete subsection reference.

<a id="canonical-1231322030200102-0120122310210113-1313023031331331-0330213000133121-1031310030010101-1121202310033300-2012231320110032-0321302321110101"></a>

## Next pages — s3_receiver / 113113102322 / 6

- [s3_receiver.aws_cred](data-sources--global_log_receiver--reference--group-004.md#canonical-0003321233311003-3301113211012122-0103013321010222-3300113211001131-0130113110200303-1032313231101020-3131112000301032-3333012230303003)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0003321233311003-3301113211012122-0103013321010222-3300113211001131-0130113110200303-1032313231101020-3131112000301032-3333012230303003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120022313303002-3003003312003221-1212231211210200-0103202003102130-1212332322130012-3221233130232333-0000310233112302-2022010123210010"></a>

## s3_receiver.aws_cred — aws_cred / 010033101131 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- s3_receiver.aws_cred

<a id="canonical-3202012023230310-1232321333123111-2110301020021221-1222030102300112-1022302333122211-0002023120303122-3310223101003302-1232311223201111"></a>

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

<a id="canonical-1030310332111233-0331331331301321-2333133022131212-3312120102121021-0223020311020222-1011332002122000-0030302120121013-2101222130213212"></a>

## Direct properties — aws_cred / 010033101131 / 3

<a id="canonical-0322202223031112-3032032230302130-2330322222210210-2212330001231202-0213232332031201-0130003313021110-1121222111233103-0212002031023222"></a>

<a id="canonical-2130110301002013-0011333132020320-3321301031012011-0211300313313120-3133002111002233-2102030010122111-2022120033212122-2220322330123212"></a>

## name property — aws_cred / 010033101131 / 4

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

<a id="canonical-1300113000103321-2023233201122030-3113320102221202-2113021312022100-2011012330213333-3132130100213212-2010022121233111-2321133221202022"></a>

<a id="canonical-3200321200103112-3203333212310110-2122221300000222-1212211201102202-0131111332321232-0302310200030312-0320130022133211-2320221111031333"></a>

## namespace property — aws_cred / 010033101131 / 5

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

<a id="canonical-1313220200133133-0303012021322010-2322333212132001-0113210103003000-0102000301112103-3231322330303333-2231133012130130-0103231211302020"></a>

<a id="canonical-3023110223112023-3201032110021303-1001332113120330-2212232213313010-0323333023321132-0320213303113131-1330332002031131-3332111130012103"></a>

## tenant property — aws_cred / 010033101131 / 6

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

<a id="canonical-0132233313223113-1332000032332301-3013032312113003-2300011001212310-3231200331101331-3322012132233110-0121301320313003-2301131202200203"></a>

## Next pages — aws_cred / 010033101131 / 7

- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002300102312030-2032311020112032-3100100033121213-2022010310031211-1231133013233121-3112303200221112-0213121011322220-2300102032333303"></a>

## s3_receiver.batch — batch / 311010110300 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- s3_receiver.batch

<a id="canonical-3020230212103033-2133200232032131-3030102231220220-0212110022322010-1113213223103133-1201013033203131-3323110321302031-2023031103332131"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-1313030213203333-1120330033031333-0233302123122233-0232003202021122-1121230333331223-2033200123101120-0322130303003320-1312103030330201"></a>

## Direct properties — batch / 311010110300 / 3

<a id="canonical-1232003320120121-1030132103012101-3300312103030100-3012333211001311-1020133010333002-2030223013330000-3302300302123300-1231222033030221"></a>

<a id="canonical-3311000000132133-2130312330103033-0213123023123233-1130302323030122-0233003200332323-3333200131222030-2023102212002223-3113301302312123"></a>

## max_bytes property — batch / 311010110300 / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-0300002321302223-0023223112101120-0013133312201111-0221001213120312-2003102311100230-2300032303120102-0220101020333010-3102313200012130): complete subsection reference.

<a id="canonical-1221000313212123-3110311020011320-1131332220203311-0011331333033001-1102111323231121-0112300121002313-1121202210323000-1222032103221003"></a>

<a id="canonical-2222301333300233-3111102001212101-0300130111030111-2121022311023302-1232221122012101-3100201013133131-0120000203211213-1330301011232122"></a>

## max_events property — batch / 311010110300 / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-2311021311132023-1303002132300010-1100031113232002-2030101123030330-1232000102221200-1201130232132300-2001310112203113-0302210121200131): complete subsection reference.

<a id="canonical-1200121311230210-3122321333020103-0120222203010002-3232310321210101-0120232232013331-1121233123022031-1013302232323233-3231203010011013"></a>

<a id="canonical-0332323112111221-2320313102023010-1122021101232213-1032200133323121-2231111132211333-3211203310222210-1300211210032103-3233120130023133"></a>

## timeout_seconds property — batch / 311010110300 / 6

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-3102212213233231-2103011002112312-1332332211233203-3303222102232333-0311012131110002-0301232231020321-0032033130331320-1010223023231111): complete subsection reference.

<a id="canonical-0102101213121312-0023232002113301-3033223033113323-2122103120122131-2220131121003333-0011230322213101-1023211211200203-2211122210121312"></a>

## Next pages — batch / 311010110300 / 7

- [s3_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-0300002321302223-0023223112101120-0013133312201111-0221001213120312-2003102311100230-2300032303120102-0220101020333010-3102313200012130)
- [s3_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-2311021311132023-1303002132300010-1100031113232002-2030101123030330-1232000102221200-1201130232132300-2001310112203113-0302210121200131)
- [s3_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-3102212213233231-2103011002112312-1332332211233203-3303222102232333-0311012131110002-0301232231020321-0032033130331320-1010223023231111)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0300002321302223-0023223112101120-0013133312201111-0221001213120312-2003102311100230-2300032303120102-0220101020333010-3102313200012130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200312021102031-3101230023211000-3102201230122120-2310112203100133-2213122013022211-1010221300212313-0110103201000013-3311013333323122"></a>

## s3_receiver.batch.max_bytes_disabled — max_bytes_disabled / 223320300010 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- s3_receiver.batch.max_bytes_disabled

<a id="canonical-1031312331231133-2030212110221311-2021111101203223-2110103313222202-0013031132001221-1011313230222130-1223031002230130-2312311031102020"></a>

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

<a id="canonical-1212133210113211-0020033021110213-1301223100233333-2312221302022321-3223323033312103-0322212211311220-0233300033201101-2302301010303222"></a>

## Direct properties — max_bytes_disabled / 223320300010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322120301013112-2313211331223112-0221001210301321-1113023301322222-2332303300322132-1020031301120213-3021333231013002-2121103132120003"></a>

## Next pages — max_bytes_disabled / 223320300010 / 4

- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2311021311132023-1303002132300010-1100031113232002-2030101123030330-1232000102221200-1201130232132300-2001310112203113-0302210121200131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001323011133133-0133010001302002-0213023320301320-3310313021021121-1100031021103012-0323020100131001-0230222300232021-0031210102331212"></a>

## s3_receiver.batch.max_events_disabled — max_events_disabled / 312203332101 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- s3_receiver.batch.max_events_disabled

<a id="canonical-3123332003223121-2213310010110002-1023023102211100-2232123322210112-2133113023311001-2021021032101303-0001130231320313-1213021231002320"></a>

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

<a id="canonical-3113323021223010-0330330312130120-2230302333331331-2002211201032231-2003313213021111-1201123213203322-1130223103202330-2112110310111013"></a>

## Direct properties — max_events_disabled / 312203332101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211033133230131-2121200111020122-1231010002202100-1210113031100212-3302230103003221-2313330303020131-2121100100202223-1013331132121201"></a>

## Next pages — max_events_disabled / 312203332101 / 4

- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3102212213233231-2103011002112312-1332332211233203-3303222102232333-0311012131110002-0301232231020321-0032033130331320-1010223023231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301322121322302-1213012011111111-1303333122112312-0123013120030033-2131021232031021-0120003331021031-1123323230030310-2323331011122033"></a>

## s3_receiver.batch.timeout_seconds_default — timeout_seconds_default / 003012133223 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- s3_receiver.batch.timeout_seconds_default

<a id="canonical-1223311312003032-2320022133231021-1221010003222300-1223300001213203-3311231021211012-0232320301211112-0310102303102010-3220030131330313"></a>

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

<a id="canonical-2321223103203221-2133231031003323-0002130121013302-1332001230200323-3123102300033020-2333330103302310-2212122203001300-1123322300012220"></a>

## Direct properties — timeout_seconds_default / 003012133223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321310030230200-2103232130331303-0322121122330221-2122131220002301-0033032210200100-3222032020222023-3212301112202301-0332131001223322"></a>

## Next pages — timeout_seconds_default / 003012133223 / 4

- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212103303100211-1032200322232111-1103011211033122-0022331321122332-0000221012031323-3233213013320211-3131000331300020-1301031233221132"></a>

## s3_receiver.compression — compression / 030001102333 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- s3_receiver.compression

<a id="canonical-3001333213310201-2102302233230011-3032230023101000-3232110002022301-1102001231322133-2030103332321002-3002320222013201-0312013121300213"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-3120213012312321-1333331013233330-2233212100223312-3310323132312002-0303023310310110-2320211313303232-1020201130001333-3300000210111231"></a>

## Direct properties — compression / 030001102333 / 3

- [compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-2002302113301221-1031131313000313-2101220213231330-3113100123132011-3301030111033313-2221213230213223-3302132123230021-0321301201331112): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-3212132120322200-2302331133101021-2111211212223010-3022331020113303-2320201110202002-1320122310202113-0303332031320331-3200222123310302): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-3000302301201303-3102312231000030-2230122033111202-1100210120321131-3311002022330203-3223331330220131-2211300321013303-1221310330001320): complete subsection reference.

<a id="canonical-3330033021303233-1032003000323321-1103121233123111-3302210033312311-1020211200202221-1102300323233233-3030210300200003-3201133110312000"></a>

## Next pages — compression / 030001102333 / 4

- [s3_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-2002302113301221-1031131313000313-2101220213231330-3113100123132011-3301030111033313-2221213230213223-3302132123230021-0321301201331112)
- [s3_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-3212132120322200-2302331133101021-2111211212223010-3022331020113303-2320201110202002-1320122310202113-0303332031320331-3200222123310302)
- [s3_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-3000302301201303-3102312231000030-2230122033111202-1100210120321131-3311002022330203-3223331330220131-2211300321013303-1221310330001320)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2002302113301221-1031131313000313-2101220213231330-3113100123132011-3301030111033313-2221213230213223-3302132123230021-0321301201331112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132203303002113-2322010200303231-2023022102003002-1201322230320031-3200003001232021-1001121230102032-2033111330022312-1020301331332313"></a>

## s3_receiver.compression.compression_default — compression_default / 031212111001 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- s3_receiver.compression.compression_default

<a id="canonical-2101011132302010-1030032330223332-2331120003000330-0021322010212233-2103201031201120-1030331312213122-1321222100100313-2202013222120010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-3123001023110331-1030310230333013-3023133310112233-3232133210333203-0023222222311323-3203032230302203-2200120311313303-2001230310310120"></a>

## Direct properties — compression_default / 031212111001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113201013102313-2120222112120020-3000122030322100-2030212130001330-1011302303221023-0031313123013122-2131003031120101-0222031011130320"></a>

## Next pages — compression_default / 031212111001 / 4

- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3212132120322200-2302331133101021-2111211212223010-3022331020113303-2320201110202002-1320122310202113-0303332031320331-3200222123310302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330020230322123-1223230032322331-0112221113230303-2112011112000202-0001010130330330-2012131112033302-1121113122201032-0220312031203002"></a>

## s3_receiver.compression.compression_gzip — compression_gzip / 313023203102 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- s3_receiver.compression.compression_gzip

<a id="canonical-1220000312113133-2210122210013121-3002312323100031-2332022331021101-1020122102103021-0322231130303011-2012323132101103-0020110213121231"></a>

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

<a id="canonical-1002230200112311-3303130203323130-0000330223203030-1303201202233021-3321201003001231-0013010223010311-2200121323301032-1202013122033333"></a>

## Direct properties — compression_gzip / 313023203102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312133023110330-2230131021313101-2323331101130322-3102121220113133-3202330020330210-2230121302020222-3213202103233231-2201012320010021"></a>

## Next pages — compression_gzip / 313023203102 / 4

- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3000302301201303-3102312231000030-2230122033111202-1100210120321131-3311002022330203-3223331330220131-2211300321013303-1221310330001320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201303211013113-3312112201230122-0231022102002131-3330003113120202-0120221110303210-1120131323030113-0132011032013003-3202332030122031"></a>

## s3_receiver.compression.compression_none — compression_none / 012013112303 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- s3_receiver.compression.compression_none

<a id="canonical-0222021302001020-1002323122223230-2121022131132310-2000023221200211-1311330112010322-3011000012030021-3011202100113013-0200222130110312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-2202201221012120-0031000013000320-3212031211021031-2201123120103333-1111030013110022-0011101230013011-2031323310232231-2233231320313233"></a>

## Direct properties — compression_none / 012013112303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211323311211003-0202233130102102-3132301201212003-2200303010133032-1232212323110221-0111031110210021-0323330223122222-3220121032003012"></a>

## Next pages — compression_none / 012013112303 / 4

- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003000111200200-3210103200303003-3300223012311113-0331221222213221-0311013133301213-1223300023033331-0212220333200110-2311103130303300"></a>

## s3_receiver.filename_options — filename_options / 131003123012 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- s3_receiver.filename_options

<a id="canonical-1032230031213311-2300200122322002-0302230112310111-2120100032200020-2301110213210003-2021332303201111-1322333322030322-3113032031023012"></a>

Type: `"single"`. Computed.

Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint
bucket or file.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-folder": "[\"custom_folder\",\"log_type_folder\",\"no_folder\"]"
}
```

<a id="canonical-3002111132332323-3331322203320222-0022333320100321-3120321022321131-2121320110301033-3012301033103313-1203302331223012-3301200021211020"></a>

## Direct properties — filename_options / 131003123012 / 3

<a id="canonical-3230322132020033-0122030010132232-1220200103000203-3133103023310103-0333033321222301-3212011323001322-3230000102021202-0103220122222322"></a>

<a id="canonical-2202222032022102-0002000310211023-1032002331320332-2332322111330100-1331111210130320-2023331321302310-0113211220022132-3303331032303123"></a>

## custom_folder property — filename_options / 131003123012 / 4

Type: `"string"`. Computed.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match.

Upstream description:

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  }
}
```

- [log_type_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-0123203103002230-1021003313212300-1020132113232121-0123320120202112-3230033000000220-2331231122332210-1233332333022011-2023130223020331): complete subsection reference.

- [no_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-1310231121310111-1212121112320110-1322001223001031-2110333310003030-0213212011333213-2311022001120122-1010123230030300-1132113233200221): complete subsection reference.

<a id="canonical-0231220101111213-0111330023033103-0320202100211202-2232322022103313-3001313123123213-3000203230323320-1301112233312310-1210012313132333"></a>

## Next pages — filename_options / 131003123012 / 5

- [s3_receiver.filename_options.log_type_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-0123203103002230-1021003313212300-1020132113232121-0123320120202112-3230033000000220-2331231122332210-1233332333022011-2023130223020331)
- [s3_receiver.filename_options.no_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-1310231121310111-1212121112320110-1322001223001031-2110333310003030-0213212011333213-2311022001120122-1010123230030300-1132113233200221)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0123203103002230-1021003313212300-1020132113232121-0123320120202112-3230033000000220-2331231122332210-1233332333022011-2023130223020331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103202120101230-2222123101323020-2302010011030303-1312122111133332-0122122222132030-2030133321212120-2110013330211012-2102013131013103"></a>

## s3_receiver.filename_options.log_type_folder — log_type_folder / 000010020230 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330)
- s3_receiver.filename_options.log_type_folder

<a id="canonical-2322233312020231-2000120120233333-3113002301321110-0000132302223232-3132211033310322-3232132330332321-2202312231021022-3322112001213221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for log type folder.

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

<a id="canonical-3100323103020032-0010220131023313-0030313132213320-3120021322010011-3101101002233003-0013033210133303-1213110301223323-2010200003113133"></a>

## Direct properties — log_type_folder / 000010020230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102221301210110-2022212313330332-3031011331113103-1020122303202313-1221113221223330-3202321032120112-2232313103110310-2020300001322012"></a>

## Next pages — log_type_folder / 000010020230 / 4

- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1310231121310111-1212121112320110-1322001223001031-2110333310003030-0213212011333213-2311022001120122-1010123230030300-1132113233200221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322012013330023-2301010111121312-3313300211313000-2011312011330001-2023122223012323-3130333231010203-2320011022123012-3233303323301303"></a>

## s3_receiver.filename_options.no_folder — no_folder / 120113133111 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330)
- s3_receiver.filename_options.no_folder

<a id="canonical-3211122120332313-3121320303230030-0222131322213202-2002322003202101-3222031120323101-1030111030312000-2110032103222302-1132220333031010"></a>

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

<a id="canonical-3220010202202022-3011000022133203-2120200021210003-2031332033330310-1323130030021023-2332233032013033-2110213032012110-2113103020122303"></a>

## Direct properties — no_folder / 120113133111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322210321113312-1220330332122130-2203213322003121-1102000102010221-0001231211020330-0221310120231033-0203103130300230-0310321230112311"></a>

## Next pages — no_folder / 120113133111 / 4

- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1002031033031310-1233001100003213-2003102123102332-1202233033221123-1021111002311202-0121010220310110-1211132100130223-2001223323310023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322211002120030-0230300232120012-3322131300221001-0030033010301103-3111132000013133-1103303123222102-2323310203101210-1020033230300001"></a>

## security_events — security_events / 001200102020 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- security_events

<a id="canonical-2121201001133031-1031120231001131-1132311112122102-3321300232323210-0223013202122312-3113023102021320-3231333333203113-1212202033021010"></a>

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

<a id="canonical-0220120100010001-0131221333213312-3022220030120230-0333303332321313-3031000001021203-1131210320123020-1322032110021112-0310302103001131"></a>

## Direct properties — security_events / 001200102020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013301120030023-2023112310201203-1300302323113202-1000010231010323-0210130233030312-1213211313321233-3301230332201303-2030113233122100"></a>

## Next pages — security_events / 001200102020 / 4

- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211032111113110-3223123101201030-1212323121312111-3133022103220213-3000001110130110-2222303123001123-3011202310020303-3231231203211312"></a>

## splunk_receiver — splunk_receiver / 210010101311 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- splunk_receiver

<a id="canonical-2130322203222220-0110323213010302-1233212331033202-0013332333331213-2311200220221030-0313332130213320-3232333202032213-1331302023310301"></a>

Type: `"single"`. Computed.

Configuration for Splunk HEC Logs endpoint.

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

<a id="canonical-1331101103303233-0210012003033010-0322103102013101-0130102002301330-0311332312031330-2033232310300210-2010100312332332-3013033131133233"></a>

## Direct properties — splunk_receiver / 210010101311 / 3

- [batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012): complete subsection reference.

<a id="canonical-0122310212212110-3112223030311123-2232032230233300-3030032021332100-0032011231230113-2103000003021332-0101003333210312-0013311330230221"></a>

<a id="canonical-3313230320310131-0301033123031201-1121212233200301-3011120132223013-0030002011030201-3002101013030200-3023111320021103-0220001023103131"></a>

## endpoint property — splunk_receiver / 210010101311 / 4

Type: `"string"`. Computed.

Splunk HEC Logs Endpoint. Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

Upstream description:

Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^https?://[^\\s/$.?#].[^\\s]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
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

- [no_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-1233030330011113-1231121311123123-0322000010211013-3222020113322211-3022231031332102-3133330303211311-2020013231023221-3011331223022132): complete subsection reference.

- [splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022): complete subsection reference.

- [use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023): complete subsection reference.

<a id="canonical-3103030011232221-1213112100231020-0313230331303220-3323111001000013-2011130200303213-0330222031210021-2013301300110030-0000222012103232"></a>

## Next pages — splunk_receiver / 210010101311 / 5

- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- [splunk_receiver.no_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-1233030330011113-1231121311123123-0322000010211013-3222020113322211-3022231031332102-3133330303211311-2020013231023221-3011331223022132)
- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122311110321300-3231030002131320-0123203202321100-2303103102020122-2310201230303200-3322330133313313-3231332323230213-3330233230232202"></a>

## splunk_receiver.batch — batch / 022103020233 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.batch

<a id="canonical-0310220231331303-0331002120123031-0301111231212212-1023313103123122-0000330132030210-3020131210312123-1200013032123101-0001321123123300"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-1011103320300302-1220003033212003-0330030301013220-1330332010202203-3022120200103233-1330123223111321-2113023310112113-1010333201213332"></a>

## Direct properties — batch / 022103020233 / 3

<a id="canonical-1313123330123213-0003011032101233-3100023111111030-1132130103203023-0023311121132023-0211032310032013-2223222211222223-0103022301200332"></a>

<a id="canonical-0000303110021201-0102030121033201-1323223121100023-2301122013231130-1103121210130332-2323122002101220-2220010133112023-0002032220032101"></a>

## max_bytes property — batch / 022103020233 / 4

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-3111222321000321-0120131323312113-1033322130301101-0000302330020000-1113012130231021-1311210222120231-1120130030002310-1200111131321321): complete subsection reference.

<a id="canonical-0332001220230311-1010200201130011-2030132303020020-3211320112201213-1023222330200313-2110312100010333-2322322023310323-2132301201311211"></a>

<a id="canonical-3302212110012121-1200000021200313-1000111331310330-0101231012203302-2021031102320101-1220230213200113-2021100323101200-0312123122112200"></a>

## max_events property — batch / 022103020233 / 5

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1033000022022201-1103122311103101-1321000023212222-0311310033112000-2321233131230033-3030001232122222-1021112131301103-3310231220311003): complete subsection reference.

<a id="canonical-0210302133331323-2002310321031133-2310301132233113-2223020302300102-2000213122103310-1100222200330101-3210300131221100-3123021020202220"></a>

<a id="canonical-2301200313102021-1302332302132212-3230110102312200-3010212020133130-2132333021131100-2131323330020222-3310120003031002-0213232230000202"></a>

## timeout_seconds property — batch / 022103020233 / 6

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-1323132023122003-2323333011320033-1001313000033020-2232133013103301-2010201300322203-0012110111123301-3222312301131202-3213321233310032): complete subsection reference.

<a id="canonical-2133300010111202-2212232001231012-2330122301212003-3312322320203223-3332222320123023-1300201101313332-3212011103310203-1103232220133103"></a>

## Next pages — batch / 022103020233 / 7

- [splunk_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-3111222321000321-0120131323312113-1033322130301101-0000302330020000-1113012130231021-1311210222120231-1120130030002310-1200111131321321)
- [splunk_receiver.batch.max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1033000022022201-1103122311103101-1321000023212222-0311310033112000-2321233131230033-3030001232122222-1021112131301103-3310231220311003)
- [splunk_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-1323132023122003-2323333011320033-1001313000033020-2232133013103301-2010201300322203-0012110111123301-3222312301131202-3213321233310032)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3111222321000321-0120131323312113-1033322130301101-0000302330020000-1113012130231021-1311210222120231-1120130030002310-1200111131321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112111110312302-2023133301112320-0221312323211003-3020001230010033-0001110201300310-0023223023011213-3213203223032022-2001131131223311"></a>

## splunk_receiver.batch.max_bytes_disabled — max_bytes_disabled / 013111301311 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- splunk_receiver.batch.max_bytes_disabled

<a id="canonical-1232023303133101-3020111132101123-0102313223233313-2300012130123310-2122102302312132-2312110110030130-0023122000311322-2103232313120211"></a>

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

<a id="canonical-0210033302133213-1233101202122132-1310322220112311-0123110230131212-2130101212202301-3011100030221030-1223322310101210-2003220130020203"></a>

## Direct properties — max_bytes_disabled / 013111301311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121122331313130-1210211030313230-3102213111031103-3321133031002031-1023322320301013-2010310101130122-0113330101012100-3003201311210103"></a>

## Next pages — max_bytes_disabled / 013111301311 / 4

- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1033000022022201-1103122311103101-1321000023212222-0311310033112000-2321233131230033-3030001232122222-1021112131301103-3310231220311003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302230303112233-1122322032212202-2001300133321100-3103321333121332-2302020002210300-0032210032230302-1102200331030210-3300133313230223"></a>

## splunk_receiver.batch.max_events_disabled — max_events_disabled / 210311300213 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- splunk_receiver.batch.max_events_disabled

<a id="canonical-1311200213112200-3231012032111330-2021101211203231-1033333300110112-0230312323032022-1223132213220201-3133212203021103-0200330313121323"></a>

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

<a id="canonical-0122020020330202-2000200131212132-1023020230331232-3131300311103320-3033021132012112-2123213312222311-0300113220202132-0112102110131203"></a>

## Direct properties — max_events_disabled / 210311300213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201313320003113-3120232000213333-1322112131210332-2221131021310313-0203320312130112-1212111231021103-1322333222303033-2310131132221231"></a>

## Next pages — max_events_disabled / 210311300213 / 4

- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1323132023122003-2323333011320033-1001313000033020-2232133013103301-2010201300322203-0012110111123301-3222312301131202-3213321233310032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133301010213001-2102231110031321-3113123211000110-3122102312000332-0321222103122220-2110330011200321-0022112020122330-2101023111013232"></a>

## splunk_receiver.batch.timeout_seconds_default — timeout_seconds_default / 321203223111 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- splunk_receiver.batch.timeout_seconds_default

<a id="canonical-0102310221311100-3321300201103310-1120201321320102-2200031123021023-2313311331312130-2013200120021233-0020310121011101-3213322313303103"></a>

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

<a id="canonical-3133222211020333-0133311100120233-0133101333132100-3003020022310020-0232311002233000-3102112012322300-0201131100103211-1121213111302113"></a>

## Direct properties — timeout_seconds_default / 321203223111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011123110103213-1022003231223200-3100200221330132-0210330111111213-0220221003231110-1223012121101223-2021000001202032-2210313200031213"></a>

## Next pages — timeout_seconds_default / 321203223111 / 4

- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131001121202000-3131000000323211-2031100010002231-1313033332012111-3321301011321313-1212233032320102-0203311132001002-3320211103300330"></a>

## splunk_receiver.compression — compression / 001230202030 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.compression

<a id="canonical-1132323220133001-0201123332100322-1033310330132112-3213213013322122-0300012121231123-3020321013321232-3323213333203302-1030032221133032"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-2100013322012122-1003322333210032-1310300222222201-0131323323200202-1211202020122300-2113112233100102-0303030011213213-1323002322313303"></a>

## Direct properties — compression / 001230202030 / 3

- [compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-0220032330333000-1120010102301122-3020221220131131-3012330331312122-0213033222201331-3102203201332302-3101102300003212-0130332101101030): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-2212130212220231-3321010011123321-3102001102331131-0000211121213231-2131203130322331-2022003133031302-1012132212333122-3311202213132110): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-1132213011323301-0203201233202312-0103201103333001-1032323212333223-2331133103322031-2132321132233012-0223222222321301-0101121130022001): complete subsection reference.

<a id="canonical-3220210201230111-1120332210003012-3212221312110032-1112123301123312-2322013103000130-0220233321100003-2322321201033223-2202001103310331"></a>

## Next pages — compression / 001230202030 / 4

- [splunk_receiver.compression.compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-0220032330333000-1120010102301122-3020221220131131-3012330331312122-0213033222201331-3102203201332302-3101102300003212-0130332101101030)
- [splunk_receiver.compression.compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-2212130212220231-3321010011123321-3102001102331131-0000211121213231-2131203130322331-2022003133031302-1012132212333122-3311202213132110)
- [splunk_receiver.compression.compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-1132213011323301-0203201233202312-0103201103333001-1032323212333223-2331133103322031-2132321132233012-0223222222321301-0101121130022001)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0220032330333000-1120010102301122-3020221220131131-3012330331312122-0213033222201331-3102203201332302-3101102300003212-0130332101101030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121102232000213-0100121213302312-3013123312013133-0112110311332230-3021002223032310-2120123101002302-2110132122032210-2001121221302031"></a>

## splunk_receiver.compression.compression_default — compression_default / 212230110312 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- splunk_receiver.compression.compression_default

<a id="canonical-3311331031122302-0212312022022233-0312032020232210-2132203310211031-0103311001021322-1002222320000023-3211103312123213-0221312003323302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-2120213310213200-0010221233020233-3121212012221222-2103121022112103-1330130330032312-1212001322223301-2110223011200132-3131212123013232"></a>

## Direct properties — compression_default / 212230110312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302003223121110-2300231210032300-1122003200113110-3213311200320103-2032110220202032-1311023203330102-2000103200220000-0022301331012131"></a>

## Next pages — compression_default / 212230110312 / 4

- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2212130212220231-3321010011123321-3102001102331131-0000211121213231-2131203130322331-2022003133031302-1012132212333122-3311202213132110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201321132032302-3312000220322002-3300121032333032-0311331122331000-2001333213002102-3330310033231223-0010300320120301-1310320021022310"></a>

## splunk_receiver.compression.compression_gzip — compression_gzip / 311331211333 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- splunk_receiver.compression.compression_gzip

<a id="canonical-2132130303021302-2101200321303130-0221321111211022-2233003020031223-3323311320110311-1021220110220012-1303032031122203-3211100221123100"></a>

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

<a id="canonical-2113013031301231-3331003103020312-0003122321222022-0330211320331000-0303323013033022-1301030202112211-3323203330333121-1211222123112300"></a>

## Direct properties — compression_gzip / 311331211333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101221133311110-2312211000211010-0331130111023303-2132031230133010-2033012010002322-0033233011302031-1011303131120130-1203223103120300"></a>

## Next pages — compression_gzip / 311331211333 / 4

- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1132213011323301-0203201233202312-0103201103333001-1032323212333223-2331133103322031-2132321132233012-0223222222321301-0101121130022001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323022330011301-1120110022202033-0120313131002112-3210313300321322-0111211013333233-0313030331333131-3113221123220331-0210110100301013"></a>

## splunk_receiver.compression.compression_none — compression_none / 122001221010 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- splunk_receiver.compression.compression_none

<a id="canonical-2221100001033010-3102322311003213-2110202223011002-3112322101332001-0301201130313111-1322300310132002-2120010331112000-2012303132000132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-1013021210101321-2313333011000332-0202330221012211-1310013330000123-2200020300230333-3201212002100111-3032111021002203-1100013232100301"></a>

## Direct properties — compression_none / 122001221010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312121223021321-0210111110132022-2310220001231231-0311300101132202-1303023232332303-1012232032113333-3033333222232310-1023032023120303"></a>

## Next pages — compression_none / 122001221010 / 4

- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1233030330011113-1231121311123123-0322000010211013-3222020113322211-3022231031332102-3133330303211311-2020013231023221-3011331223022132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322310212330203-2200033023033323-1203111023311002-3012001130302332-0103000233322032-1300120111010313-3132133331322030-3312032030230323"></a>

## splunk_receiver.no_tls — no_tls / 232310332313 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.no_tls

<a id="canonical-1323320233332023-2333133131013212-0300231313001332-3002202330132300-0231200131032222-3223322211102210-2011000310220322-1312000232220302"></a>

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

<a id="canonical-1120132003131032-0323302333331221-0200000310112000-2131000122113200-1231300113133332-0112302101320012-3002133111123233-1111322121202000"></a>

## Direct properties — no_tls / 232310332313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121111133031013-2311300100323201-1132100110221031-1213312133130010-1320031131232023-0232112021332203-2230331122213110-1030303330233211"></a>

## Next pages — no_tls / 232310332313 / 4

- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300020222003203-0202021220032203-1220023102323221-3121231100033011-1103210131332120-0020012333122123-1132022011123233-2311332200310021"></a>

## splunk_receiver.splunk_hec_token — splunk_hec_token / 130122313102 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.splunk_hec_token

<a id="canonical-0010103202313023-3011132320331202-0102212312100212-3001200030210211-2311320200030020-3322033230203133-3023103200120212-2003331010321021"></a>

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

<a id="canonical-1303031212313311-1320201213130023-0002001121022010-3102223210102003-3123132320331121-3222203033112222-0131100313130303-2122302103000212"></a>

## Direct properties — splunk_hec_token / 130122313102 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1121031112302011-3332331311022312-3100203321103031-3310001103223212-3131020002020022-1312123023032000-1221330110000113-3222000201010332): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1021211032223203-1332333030312012-0210130231331232-3022003021333011-0000133333131223-2333213130002232-0311122212130021-0302011123212222): complete subsection reference.

<a id="canonical-3333101213221200-0331321220111312-2213201231113210-2002001123212210-3030021101333233-3220321031013310-0110322201111020-0011000221033223"></a>

## Next pages — splunk_hec_token / 130122313102 / 4

- [splunk_receiver.splunk_hec_token.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1121031112302011-3332331311022312-3100203321103031-3310001103223212-3131020002020022-1312123023032000-1221330110000113-3222000201010332)
- [splunk_receiver.splunk_hec_token.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1021211032223203-1332333030312012-0210130231331232-3022003021333011-0000133333131223-2333213130002232-0311122212130021-0302011123212222)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1121031112302011-3332331311022312-3100203321103031-3310001103223212-3131020002020022-1312123023032000-1221330110000113-3222000201010332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012001130102302-2212112312220002-1322301322103210-1101233010312102-3321322003030202-2232300330200030-0012300303013233-0321032233123120"></a>

## splunk_receiver.splunk_hec_token.blindfold_secret_info — blindfold_secret_info / 230013012232 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022)
- splunk_receiver.splunk_hec_token.blindfold_secret_info

<a id="canonical-3132210202011312-0201101331120113-1222232030200210-1230121333300323-1230101102232033-0020302320112213-2303000333023102-3032303211133120"></a>

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

<a id="canonical-1221212102133203-3311332310233121-3012023000003321-3032133322213120-1203223011311332-1312113011200232-2023302001211012-2000013113301130"></a>

## Direct properties — blindfold_secret_info / 230013012232 / 3

<a id="canonical-2303101132032203-0211112123202323-3211110030200203-0131223211012130-1110332210123011-2332310301332103-0210313012320213-1111132032323103"></a>

<a id="canonical-2333223332030002-2332021010201230-1323021202100312-0300110130220103-3000110333321310-3303232030211011-0120200013202102-1122311021220001"></a>

## decryption_provider property — blindfold_secret_info / 230013012232 / 4

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

<a id="canonical-0103031121311131-3323313022123223-3203111111300213-2300231302300202-0232120303332032-1101033203231301-1103202232102110-3013020000113203"></a>

<a id="canonical-2200333221103210-2020132001221112-1101111101302201-2312310232333200-2111322222200320-3032003323311133-0311110223213013-0102113312222031"></a>

## location property — blindfold_secret_info / 230013012232 / 5

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

<a id="canonical-1311131020313032-1133313011023101-3023031222233011-0100300020213101-1311132232113303-2121100212213012-1122100101202033-3133123220013330"></a>

<a id="canonical-1002012131333023-3212321021021230-1022010231203323-0002021031330013-2331303203101313-0312232312323231-1210133020022223-1212022331120030"></a>

## store_provider property — blindfold_secret_info / 230013012232 / 6

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

<a id="canonical-3122202010020231-2102232001111301-1320222310012031-1111213003231020-3223100111022012-2021030223002021-0201010332013103-0220302133221031"></a>

## Next pages — blindfold_secret_info / 230013012232 / 7

- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1021211032223203-1332333030312012-0210130231331232-3022003021333011-0000133333131223-2333213130002232-0311122212130021-0302011123212222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213122333302231-2020303233231130-0213033121001222-0222201332231332-0021000212010330-2001322031202333-2100133112132103-0021020220023002"></a>

## splunk_receiver.splunk_hec_token.clear_secret_info — clear_secret_info / 013022020011 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022)
- splunk_receiver.splunk_hec_token.clear_secret_info

<a id="canonical-0200023233103100-3011023230000130-0100301222313331-3121101032301320-0231302002301322-3101002221230110-1032013322310103-2122320232021113"></a>

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

<a id="canonical-3003121120202202-2022302230031302-0033220032303203-2202031010323322-0032222011103211-1201202032013323-2100023030300210-0331203333013103"></a>

## Direct properties — clear_secret_info / 013022020011 / 3

<a id="canonical-1333313132331203-1003130333221001-3222233022303012-3223123031320111-2321100003122123-1122303121202133-0033232202131203-2101121321213203"></a>

<a id="canonical-2300230010111131-0232222200111110-2212300001102032-1030011130232233-3111222301200123-2031013201030020-2111002112230322-2010022212312021"></a>

## provider_ref property — clear_secret_info / 013022020011 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0231130312323030-1130312230200320-2320111130021202-2131123210301103-1320111220031212-0201110010111301-2122203012333333-3321312303021333"></a>

<a id="canonical-1022102001320222-0323033202300030-1231111101330301-3113000303333101-1232032221310030-2312221321100102-0102121210013103-0031213331211022"></a>

## URL property — clear_secret_info / 013022020011 / 5

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

<a id="canonical-0302020313311323-3022212121131000-1313201312231202-2013322323033121-1223301333001021-2213002131022333-1032121310202320-3112111221012111"></a>

## Next pages — clear_secret_info / 013022020011 / 6

- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013233201323222-3233002200220302-3101002110021103-0021330131001102-0012200321330000-2231002100210210-2323011033023200-0031113033131311"></a>

## splunk_receiver.use_tls — use_tls / 102000012032 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.use_tls

<a id="canonical-1031212113011132-2011333230211332-2123311222333320-3331202123213021-1113210222203122-3231203212213130-1301320311201323-1313312332103102"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

<a id="canonical-2200211330232313-0333000100020220-0023231233200002-0212303030132223-0000333201011200-0021221132130100-3213221332210101-1302132221202122"></a>

## Direct properties — use_tls / 102000012032 / 3

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-0113132010013332-1122203011101010-0110033323000110-0101112021112202-2001131212021302-1002001101133213-1332223012130120-3301023201332102): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-2120013112032130-2201222001203321-0130203200120210-1312300211012122-3313331301223223-3233100331213300-3320320312310122-1032101010001213): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-1232330203130112-1311330323130000-1300003021122033-1222320003112203-0333031110331212-0200201112212310-3202313032312331-1112020123301121): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-3333222331323333-2030013030332323-0331321210300120-2212123201100313-0112313231221131-3023202031103100-1331233133033201-3300102121112321): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1131330031301212-3222002001020010-1223131122021130-3220131332023103-2120101201231023-2113023101202200-1113032233303001-1233100102003200): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-3222321113000113-2231233223133012-2032013200210220-0201021000323102-1331023333031133-0122211311330020-0200010321023212-3111003231323233): complete subsection reference.

<a id="canonical-1033200111011313-3230202023223222-2101030323123302-1003323303012031-2123100222300223-0111100333233002-2231311332002310-0220231120112202"></a>

<a id="canonical-0232033320233020-1120210213023032-3323213120110230-0302032123311030-0332102113120233-1210032232231331-0110330030220323-3120023100302331"></a>

## trusted_ca_url property — use_tls / 102000012032 / 4

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-1223001330120321-1122221212100102-2122333122103102-2320312321002231-3320113313030111-3012013010223022-2302210202020013-1223102301313110"></a>

## Next pages — use_tls / 102000012032 / 5

- [splunk_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-0113132010013332-1122203011101010-0110033323000110-0101112021112202-2001131212021302-1002001101133213-1332223012130120-3301023201332102)
- [splunk_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-2120013112032130-2201222001203321-0130203200120210-1312300211012122-3313331301223223-3233100331213300-3320320312310122-1032101010001213)
- [splunk_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-1232330203130112-1311330323130000-1300003021122033-1222320003112203-0333031110331212-0200201112212310-3202313032312331-1112020123301121)
- [splunk_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-3333222331323333-2030013030332323-0331321210300120-2212123201100313-0112313231221131-3023202031103100-1331233133033201-3300102121112321)
- [splunk_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1131330031301212-3222002001020010-1223131122021130-3220131332023103-2120101201231023-2113023101202200-1113032233303001-1233100102003200)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320)
- [splunk_receiver.use_tls.no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-3222321113000113-2231233223133012-2032013200210220-0201021000323102-1331023333031133-0122211311330020-0200010321023212-3111003231323233)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0113132010013332-1122203011101010-0110033323000110-0101112021112202-2001131212021302-1002001101133213-1332223012130120-3301023201332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320130322201203-3221231331010320-2023303313320010-0230001020201222-0233013002300220-1323210323022231-1001123031123313-2030011333213122"></a>

## splunk_receiver.use_tls.disable_verify_certificate — disable_verify_certificate / 032233121000 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.disable_verify_certificate

<a id="canonical-2223231332002313-0032022223310312-1231223133230100-2123121200320213-3033232222021103-2322032320322231-1123100133131031-3003321022111031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

<a id="canonical-3332000020012033-3311102021113232-2310210301331220-1313011332000301-1321231221011030-1131001333301302-1000333131200220-0302033133001232"></a>

## Direct properties — disable_verify_certificate / 032233121000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023000100110131-0030220231000012-1313332232123120-1233201223330131-1001031030132332-0021201030300223-3102010110131201-0010233321000231"></a>

## Next pages — disable_verify_certificate / 032233121000 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-2120013112032130-2201222001203321-0130203200120210-1312300211012122-3313331301223223-3233100331213300-3320320312310122-1032101010001213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032222231023320-1022310022122303-0330121113312233-0022010302332313-2011112311002301-1121222101110301-1030201333311211-0202023231322132"></a>

## splunk_receiver.use_tls.disable_verify_hostname — disable_verify_hostname / 110012131013 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.disable_verify_hostname

<a id="canonical-2220022321101202-3023033101013113-3113332220001112-2111312011023122-2301330200013003-3113232311121320-3130132320023333-2002100320323032"></a>

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

<a id="canonical-3233112333033320-2232321013131313-1331101130120320-2030211301203210-0221010332202232-0300023102330210-2130122303133313-1311001303033312"></a>

## Direct properties — disable_verify_hostname / 110012131013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012221301121001-1022030220122320-2222132112333201-2210310201003220-0133232132313201-3201000011023102-3310332102110123-0110013031032131"></a>

## Next pages — disable_verify_hostname / 110012131013 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1232330203130112-1311330323130000-1300003021122033-1222320003112203-0333031110331212-0200201112212310-3202313032312331-1112020123301121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202230111203322-3212303310130121-1203133322330023-1011302213012230-2222122222113033-0302123223020230-3122001221202232-0310210222212031"></a>

## splunk_receiver.use_tls.enable_verify_certificate — enable_verify_certificate / 232332130120 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.enable_verify_certificate

<a id="canonical-0230322230221202-2020112013000230-3123100122122031-1133022021313122-2031120331313201-1320223023301333-1021003222302010-2012313213103321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

<a id="canonical-0023000221232311-2310023132220110-0123330013030322-0302011320110322-2221213003212310-2321210113230123-2212130013103110-1202321101023130"></a>

## Direct properties — enable_verify_certificate / 232332130120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202313203312331-3010012122120202-2002212231010221-1012221222011031-3322323123002003-3100212321220320-0330211033232130-2131332301201203"></a>

## Next pages — enable_verify_certificate / 232332130120 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3333222331323333-2030013030332323-0331321210300120-2212123201100313-0112313231221131-3023202031103100-1331233133033201-3300102121112321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103230032301130-2110312021003112-2232002032331013-1032010210103103-3101033100203122-1210001003001201-0110133211111312-0322330300002000"></a>

## splunk_receiver.use_tls.enable_verify_hostname — enable_verify_hostname / 012310213102 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.enable_verify_hostname

<a id="canonical-2310123001312230-0021301000230220-1033231002320021-2222230121103303-0033022312001232-2131331000033020-3300210221033122-2323000113022022"></a>

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

<a id="canonical-1121000102133110-0031032221330322-2201213321310302-3313123013211210-1313011111230220-1202013032222220-3002131302222330-2211132221133212"></a>

## Direct properties — enable_verify_hostname / 012310213102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021032130312233-1012022300313111-1320221130221230-1321133220013212-0031302212201322-2222211110011322-3112203132130311-2122313022312302"></a>

## Next pages — enable_verify_hostname / 012310213102 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1131330031301212-3222002001020010-1223131122021130-3220131332023103-2120101201231023-2113023101202200-1113032233303001-1233100102003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111211133330322-3112022313113332-0201121112030310-0322212303110001-1002201333313300-1023122102311213-0020200023032233-0301321123122303"></a>

## splunk_receiver.use_tls.mtls_disabled — mtls_disabled / 102231130201 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.mtls_disabled

<a id="canonical-3012003232022201-3102013121131022-0232111011323021-2000213212120120-3133132033333333-0323100303331113-1110320011033330-1103000320110013"></a>

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

<a id="canonical-2001112210320123-3002323321311321-2200022311130322-3113123313010030-1102222201130323-0200030221010130-1300033110333211-2221303210002011"></a>

## Direct properties — mtls_disabled / 102231130201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002110322300310-1001032311101322-3003312232132201-3330222200311333-2021122032001210-2202102232233121-2132302321213221-2103210223113030"></a>

## Next pages — mtls_disabled / 102231130201 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133232103203220-3322133221320303-1302230111133201-0020231220030223-2130232130121200-3103210223301020-2222322301001213-3003321320302011"></a>

## splunk_receiver.use_tls.mtls_enable — mtls_enable / 102211311123 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.mtls_enable

<a id="canonical-2010302101030012-1210121202101231-1111311133202220-2021132103100332-3012001222103222-1023303233013021-2101111032233310-0101130022223233"></a>

Type: `"single"`. Computed.

MTLS Client config allows configuration of mTLS client OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3033333022133103-0221130102033223-0010101300311322-0123002331113001-0302133111010002-2300110303103111-1111111103311220-1332121113212203"></a>

## Direct properties — mtls_enable / 102211311123 / 3

<a id="canonical-0120010122112010-0002132022130120-2032300031113301-3100221130213203-2120213321103212-0323302233311113-3330022223033103-2032233232232220"></a>

<a id="canonical-2002320132201012-1311033210111211-0021102010020313-2223312133302001-1103312331202122-1030030211103033-2331001002202133-0323200101003222"></a>

## certificate property — mtls_enable / 102211311123 / 4

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200): complete subsection reference.

<a id="canonical-3221013312321103-3100122223213211-3223302003300030-3121233312022031-0131031310203030-3022030212301133-1312011210212331-1300220033010022"></a>

## Next pages — mtls_enable / 102211311123 / 5

- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120333113020023-0312023302102011-2300103220202222-0003122222332311-0122101311222130-2232032311032132-2231323322211033-2110022321331123"></a>

## splunk_receiver.use_tls.mtls_enable.key_url — key_url / 232331130303 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320)
- splunk_receiver.use_tls.mtls_enable.key_url

<a id="canonical-3113201231201020-0013002203113121-0333123202230231-3113112321232032-1132112323111003-1211202112312232-3121010132010333-3012103323301131"></a>

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

<a id="canonical-1202201331023131-3132303133212200-1323230203113223-3131323111110010-1320011030123202-1331203100111101-3133333321021112-0012021333200000"></a>

## Direct properties — key_url / 232331130303 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3202132321221132-0221233313210303-2120321211022013-3002031233031122-2113033031231200-2133201021312221-0102301201110301-2200130033132223): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1031033333021002-2013213002003321-1010311121211023-3132121203200222-1133012003000201-1103133202223011-3211002103301121-1221103121103021): complete subsection reference.

<a id="canonical-3211100001231023-0210002302002001-2122221033310213-1010113013301232-1200033032310000-1011101023222220-3011221331000311-3121132230312120"></a>

## Next pages — key_url / 232331130303 / 4

- [splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3202132321221132-0221233313210303-2120321211022013-3002031233031122-2113033031231200-2133201021312221-0102301201110301-2200130033132223)
- [splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1031033333021002-2013213002003321-1010311121211023-3132121203200222-1133012003000201-1103133202223011-3211002103301121-1221103121103021)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3202132321221132-0221233313210303-2120321211022013-3002031233031122-2113033031231200-2133201021312221-0102301201110301-2200130033132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023320133321331-3332211302213123-0100101103001210-1222332023210102-3101311201332212-3302303030001221-3002112113020122-2112211203210301"></a>

## splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — blindfold_secret_info / 013213011100 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320)
- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200)
- splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-2120321122323330-0221123321123322-0012200100202233-0213210301302221-0120131120112130-0133201332300111-2003321133012323-3233130130300122"></a>

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

<a id="canonical-1102103210113321-2300302010133113-1310300322033012-2110000102232300-2223123313130020-2122012302323313-1132212100003103-1121302203303233"></a>

## Direct properties — blindfold_secret_info / 013213011100 / 3

<a id="canonical-3102322113210211-3303301020102000-0201200131101202-1120120113220313-1032223012312300-3130033323031032-2321220201232212-0110200311032301"></a>

<a id="canonical-0001321232210320-1003323121303220-2323103232021003-1301121233121223-1012131320033313-1230000213231030-1312331023302131-1110130210333311"></a>

## decryption_provider property — blindfold_secret_info / 013213011100 / 4

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

<a id="canonical-0322110130022031-1030130230320232-2101300030000133-1323310321002003-2233010231323330-3003321030223122-1230303112011132-0003030330130130"></a>

<a id="canonical-3113132110120021-3232010302211200-0231020311023211-3103110330220131-0321020312203200-2101132033201131-0103333212030131-3302223110233200"></a>

## location property — blindfold_secret_info / 013213011100 / 5

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

<a id="canonical-1321131122030021-0010103220020013-0003002011330200-2320030102300313-3002321222230111-1333001301231323-1323130202223303-0011323200102022"></a>

<a id="canonical-2212013111102012-1213233212203102-3212233101123221-2202311212211122-3122022323231200-2223003110233020-2003011023231223-1210133323330301"></a>

## store_provider property — blindfold_secret_info / 013213011100 / 6

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

<a id="canonical-1202222231220103-2111000200123110-2311210233122220-3023311303101330-2111003331032003-0302221011222231-3123223132023021-0223030331223210"></a>

## Next pages — blindfold_secret_info / 013213011100 / 7

- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1031033333021002-2013213002003321-1010311121211023-3132121203200222-1133012003000201-1103133202223011-3211002103301121-1221103121103021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032232100001222-3203131310101220-3100020132132121-1321332033233023-2010311203212000-3223203320322002-3322022030322000-0323103321231132"></a>

## splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info — clear_secret_info / 103323333022 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320)
- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200)
- splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-3232103321112020-0023111312201010-0222130313012211-2213300313122012-0100033123321213-0233220100131002-3321211203112113-0003320232222233"></a>

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

<a id="canonical-3101200033232210-0113331302121033-0102103211323232-3011313033133102-0031232202022113-3130211220210332-0220120210223221-3012220312131220"></a>

## Direct properties — clear_secret_info / 103323333022 / 3

<a id="canonical-0301133101332311-1330020030200330-2300331102002000-1021120321000120-1001022003301003-1330300323101302-2121313320110321-3121220131213222"></a>

<a id="canonical-2320123201123000-2000023331121001-1300003303011012-3213131112213220-2223311100121122-2002310230110331-3100011101230011-3230232331323303"></a>

## provider_ref property — clear_secret_info / 103323333022 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3012023233110300-2100323311133121-2320013322321211-0321233302300023-3323203232303231-1103221001020023-2200121332021022-2233331333210333"></a>

<a id="canonical-1033003033233102-3102033230212112-3020230032313200-0220100001223321-2012130320320200-0330110121130333-0233130220030132-3013231022131122"></a>

## URL property — clear_secret_info / 103323333022 / 5

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

<a id="canonical-1203131102222222-2001100332332300-3323100221230322-0000130012300230-2113312110222310-2332110120333222-1303312021303300-0333300020132211"></a>

## Next pages — clear_secret_info / 103323333022 / 6

- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3222321113000113-2231233223133012-2032013200210220-0201021000323102-1331023333031133-0122211311330020-0200010321023212-3111003231323233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302102131213131-0211011003203303-2123113010313102-0102020122122300-3120322002010210-2023000211202113-2330202201311312-1010110220320120"></a>

## splunk_receiver.use_tls.no_ca — no_ca / 210302003132 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.no_ca

<a id="canonical-0121120030121223-3311221203133322-0331001300031213-2123200101302232-2330321212013011-1212222223213010-0333131113101330-2311122221133123"></a>

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

<a id="canonical-3032313222220002-3112001210303021-1220102222221003-2023022021310030-1032322300210121-3000203310203321-0100130333130030-1032000031300131"></a>

## Direct properties — no_ca / 210302003132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233032203331230-0103302002001133-1331200300202021-2303002110031321-1003313113211221-1331223323332203-0003010032212011-2321301102302003"></a>

## Next pages — no_ca / 210302003132 / 4

- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223002010101021-3213031103223213-0300103130100012-3021300122001011-3320120211110000-2110302322022102-1011210321233020-1112003100021000"></a>

## sumo_logic_receiver — sumo_logic_receiver / 200021210221 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- sumo_logic_receiver

<a id="canonical-2111201203021003-1010212310123022-2023221323331032-2212330321202303-1123222120131130-2113333103311120-3110301220010121-0200302122002201"></a>

Type: `"single"`. Computed.

Configuration parameter for sumo logic receiver.

Upstream description:

Configuration for SumoLogic endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0301103132010333-3300023333111002-3022011111303200-3202311300133233-2303130203013210-2021331313230012-2003301202212133-2002030213113002"></a>

## Direct properties — sumo_logic_receiver / 200021210221 / 3

- [url](data-sources--global_log_receiver--reference--group-004.md#canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221): complete subsection reference.

<a id="canonical-2331212112332113-3311110313322113-2232022002001211-1012233111230231-0201332102333220-2120102131022112-2222330031222023-2311030002031033"></a>

## Next pages — sumo_logic_receiver / 200021210221 / 4

- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031320133212111-2130212130031101-0313230021103012-1023012202102120-3233200200213333-1012302213003103-3122002322021101-3111203200122001"></a>

## sumo_logic_receiver.URL — URL / 233333330132 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001)
- sumo_logic_receiver.URL

<a id="canonical-3031113231133321-2013020123113033-3211311322001231-0323202201000203-0230303230001011-1131213213200200-0200233011132202-2110303323102013"></a>

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

<a id="canonical-2033100132322102-3312310333033122-2223222101321122-1131311103230111-1013111302330332-1223230203003020-1203313213313033-2302111023101301"></a>

## Direct properties — URL / 233333330132 / 3

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-0221110123003133-2102020331223201-3023110001211023-0202223323331111-0320021323323112-2331022323002200-3011132120211200-0120032000201002): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3022112033113130-3110121321311232-1122311012323133-0110011323101211-1021100232110232-0303123002223213-3012011201032230-0103120122031222): complete subsection reference.

<a id="canonical-3113232301033233-2020002221220113-0122131112212303-3300012202022020-2330002322001323-2330210201200121-2312132310013320-0110033031202313"></a>

## Next pages — URL / 233333330132 / 4

- [sumo_logic_receiver.url.blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-0221110123003133-2102020331223201-3023110001211023-0202223323331111-0320021323323112-2331022323002200-3011132120211200-0120032000201002)
- [sumo_logic_receiver.url.clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3022112033113130-3110121321311232-1122311012323133-0110011323101211-1021100232110232-0303123002223213-3012011201032230-0103120122031222)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-0221110123003133-2102020331223201-3023110001211023-0202223323331111-0320021323323112-2331022323002200-3011132120211200-0120032000201002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211100210333011-3330013302002101-0202312013011101-3322012312303132-3200100011101320-3102303102010300-3030002121100332-2202301211320130"></a>

## sumo_logic_receiver.URL.blindfold_secret_info — blindfold_secret_info / 011130300102 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001)
- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221)
- sumo_logic_receiver.URL.blindfold_secret_info

<a id="canonical-2323021002101033-1133310322210301-3120301223233331-3200112230302133-2031000301200123-0302300201133113-1211233300003302-2312300100313001"></a>

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

<a id="canonical-2330122321220021-3332230120021103-1300132322221032-1011211203200323-1121211212010200-3100201201201300-0311122022011132-2012212303113022"></a>

## Direct properties — blindfold_secret_info / 011130300102 / 3

<a id="canonical-1121232332123313-0120222303303113-2311222222302330-2122130020010333-1302110313023030-2122021101322213-3103010122102113-0110210101310032"></a>

<a id="canonical-0123032133200310-3231200302302212-1010213032120322-0110320230332001-1110132303131210-2301010022003103-0221323312032101-3220331311132302"></a>

## decryption_provider property — blindfold_secret_info / 011130300102 / 4

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

<a id="canonical-2300330231310021-1302012222330123-1310131131321132-2111021102122021-0020030123013003-1133032123120100-0102301331101023-1112121223002011"></a>

<a id="canonical-1232012220211201-3212111000021110-2231032100223123-1200131101320032-2002230020330333-2031233202231300-3202033133032032-0101220323132222"></a>

## location property — blindfold_secret_info / 011130300102 / 5

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

<a id="canonical-0301223313113202-1011312230223030-2022203332200310-1023102133313112-1120213221223033-2202331111013232-3322221021130132-3013323030231013"></a>

<a id="canonical-0012023010031233-3201023110011211-2210111231300331-3001333200200310-1030021012303211-1222113012232013-2121330301130121-3130030203101202"></a>

## store_provider property — blindfold_secret_info / 011130300102 / 6

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

<a id="canonical-1302023001112322-0022323023002113-2120123211021203-1333031110313131-1301330200323033-0200003302020123-1211320112121033-1321313103121320"></a>

## Next pages — blindfold_secret_info / 011130300102 / 7

- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)

<a id="canonical-3022112033113130-3110121321311232-1122311012323133-0110011323101211-1021100232110232-0303123002223213-3012011201032230-0103120122031222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001333232322200-1022300011103110-1230212301132020-2202232232022020-3012130130312020-0313120000321102-0112213033330002-0323220133002121"></a>

## sumo_logic_receiver.URL.clear_secret_info — clear_secret_info / 002201201202 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001)
- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221)
- sumo_logic_receiver.URL.clear_secret_info

<a id="canonical-0021203223313310-2211132223321003-3203301313223033-2031021001330323-0123202023020213-0021323022133130-1023113021001123-1313121020213002"></a>

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

<a id="canonical-2322332231103320-0112311202033322-1131111001331103-3113210312322200-1323230202032300-1313102323001003-0223220032003321-3222011333313002"></a>

## Direct properties — clear_secret_info / 002201201202 / 3

<a id="canonical-1303300110310232-2312131212230133-3211332002003123-1012111001131130-3201302200320122-3231113033320302-0320333210202031-3311112222222012"></a>

<a id="canonical-2033200000003102-3303223233330221-3232101132332112-2101110211022330-2012023121320313-1003220331100100-1310221332231111-3331331030300210"></a>

## provider_ref property — clear_secret_info / 002201201202 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1213131123010212-2301033211021332-0222321211113120-0202330202030231-0320200010301032-1102121133202132-2110231032122033-1320213100220203"></a>

<a id="canonical-3003103120213332-2230233311011230-3232202333222213-0321200211130032-2233322122100121-0002202212113023-0112220312013133-1010312001211120"></a>

## URL property — clear_secret_info / 002201201202 / 5

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

<a id="canonical-1013011120111332-2322003200033322-2020332130001001-3122233103330010-2330210100122002-0300320313032330-1313201023331312-2031211120022213"></a>

## Next pages — clear_secret_info / 002201201202 / 6

- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
