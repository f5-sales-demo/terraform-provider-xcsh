---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1010100131212232-2001221030033123-2332233010121110-2300022021021323-3113021031010131-0323330112321223-3332333011330101-1011033003323202"></a>

## more_option.request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 012011312312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.request_headers_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3131323333013211-1120301133201010-1211330033032130-3121231320313312-1312311200203221-3112223310131110-3201331320102000-1013203220202110)
- [more_option.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-1131121000130320-1202130021200221-1000000220302321-0200030221313113-3002002321113003-3013132122301312-3130102313111120-1023001113122030)
- more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3220321331223321-1122030132112030-2102001031332301-2311233120300003-0031331321020010-0212221212230003-0000331310103221-1322300233123021"></a>

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

<a id="canonical-2123303133031001-1231111131211012-2231021233331101-2120102133332023-1321022113330121-2012023321310212-3130313330102320-0322022110021111"></a>

## Direct properties — blindfold_secret_info / 012011312312 / 3

<a id="canonical-3113211213011303-2331331030030013-1230232131313330-0030230012031310-0021301333030302-2321022003121100-3212323031110320-0123211232301023"></a>

<a id="canonical-0123011131230221-1232233010031030-2013232023302111-0312100312133311-2030022200021020-3123132301202330-3130300310002222-1200202200303121"></a>

## decryption_provider property — blindfold_secret_info / 012011312312 / 4

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

<a id="canonical-1011220203112322-2122122110322021-2201230313231222-0200232110203210-1000020220301101-0002121223331002-1222311233021130-1301322023330213"></a>

<a id="canonical-1002302132332111-2112133133313210-1011130110000311-3111101223020121-2113312130223022-1300230330022333-3001230013210322-3212103211013022"></a>

## location property — blindfold_secret_info / 012011312312 / 5

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

<a id="canonical-1200201102111100-3210010001130102-2022020023000101-3232011310123011-2032033321031021-1003011221303223-1320101222322123-3132001130000302"></a>

<a id="canonical-0222201202023331-0031010130210020-2000002012123323-2103230321002023-0002333012331032-0202213222113203-0120211031320103-2130120112213013"></a>

## store_provider property — blindfold_secret_info / 012011312312 / 6

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

<a id="canonical-2212322131120331-0013103101320100-3023301131122020-2303022311130221-2323331112130023-1213121103112123-1211010210230322-1302123012000113"></a>

## Next pages — blindfold_secret_info / 012011312312 / 7

- [more_option.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-1131121000130320-1202130021200221-1000000220302321-0200030221313113-3002002321113003-3013132122301312-3130102313111120-1023001113122030)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3333033032103310-0321020023231321-2331031212121033-0330213321121131-2123032230333122-1111202213202312-1211212221013132-3311312323321102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210200002031022-2321321331032332-2212120333002010-2123011311323120-2313122212123313-1312202001320020-1321133331003110-0132330311030221"></a>

## more_option.request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 120331100133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.request_headers_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3131323333013211-1120301133201010-1211330033032130-3121231320313312-1312311200203221-3112223310131110-3201331320102000-1013203220202110)
- [more_option.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-1131121000130320-1202130021200221-1000000220302321-0200030221313113-3002002321113003-3013132122301312-3130102313111120-1023001113122030)
- more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0032223111303303-3323103102122302-0213221210313210-0310220113112321-3332330102030202-3103313110113021-1212210312231213-3033112012322023"></a>

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

<a id="canonical-0222222122230213-0332222113320113-0302031132201230-3213031102332110-2013020020321110-2223213103223223-1022333123002200-2320113100112222"></a>

## Direct properties — clear_secret_info / 120331100133 / 3

<a id="canonical-0100032101231311-1211100122121223-3002311003332023-1212001121031111-3201012203133130-0030122130131333-0202023112321132-1230032032012300"></a>

<a id="canonical-0011123302022121-3211222322002303-3102211311330000-0231320121212133-2331231102220303-2033023333200020-1031311233100302-3113303032102001"></a>

## provider_ref property — clear_secret_info / 120331100133 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3110311303122221-2202310102233130-0112023222223122-0131213302312212-3303321010222330-1332211302213110-1223202322300303-0003303123310101"></a>

<a id="canonical-3131113113301032-3032313303021000-3333230222202010-3002123023032222-0220133102311320-3222020300031010-0233001220021013-3010100021211303"></a>

## URL property — clear_secret_info / 120331100133 / 5

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

<a id="canonical-0122300310123221-1331020220022011-3303013132121333-2201033310201331-1022030202200032-3111313233221113-0032233001301101-2033002313101031"></a>

## Next pages — clear_secret_info / 120331100133 / 6

- [more_option.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-1131121000130320-1202130021200221-1000000220302321-0200030221313113-3002002321113003-3013132122301312-3130102313111120-1023001113122030)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131023102123001-3330202032112311-2322223331121023-2121011233012320-3322222230121021-1032202203131201-2321112032223102-0323203311310211"></a>

## more_option.response_cookies_to_add — response_cookies_to_add / 200212330110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.response_cookies_to_add

<a id="canonical-2101101032203133-3212123103302333-1000313321122021-1111222312312303-1221213323120033-2211011210010322-1210003322020023-3311220302023212"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

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

<a id="canonical-0023303330003000-0023130033131122-2232231222020210-3102210123210302-1003133033030022-1132010202220030-2321003120230011-1133312020230231"></a>

## Direct properties — response_cookies_to_add / 200212330110 / 3

<a id="canonical-1301110313211001-0220322220320233-3112123012222112-2002310120110100-1300010323021320-3000233021233302-1201110122210313-2121002321201301"></a>

<a id="canonical-2221121311011331-2320022101001012-1001100021011200-3231013210320111-1012211311302031-1122330100320103-3301003002011333-2212322232301101"></a>

## add_domain property — response_cookies_to_add / 200212330110 / 4

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

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

<a id="canonical-1032120010211121-1313120021010212-2032210122310010-0020110213311231-3003113103011010-2030100311312003-1232313111331131-0321032031312320"></a>

<a id="canonical-0120231131230303-2233111220301332-0232011211110101-2021222130233232-2030211122323220-2030302123022331-1013200212022223-2330003101302300"></a>

## add_expiry property — response_cookies_to_add / 200212330110 / 5

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](data-sources--http_loadbalancer--reference--group-021.md#canonical-0300223001322130-2030032331032033-1003212011020133-2233320201331011-0030003131232030-2312222132220200-0001123123322032-2313020031031221): complete subsection reference.

- [add_partitioned](data-sources--http_loadbalancer--reference--group-021.md#canonical-2330103311303032-3130031021333331-1202200000210223-1010020332113233-3213302010103030-1032313211131110-1222033112230303-2200112031020203): complete subsection reference.

<a id="canonical-3001002123132133-3131310312330312-0130103221233000-3130210011032022-2023223111312112-1103023132013022-1001133202012223-1030000030323101"></a>

<a id="canonical-2020033200112223-3301102130232003-1313031010302003-1103331013222001-0100003231111203-2223110132111000-2202323031222231-2123133201101020"></a>

## add_path property — response_cookies_to_add / 200212330110 / 6

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](data-sources--http_loadbalancer--reference--group-021.md#canonical-1120032113222132-2213222211032003-1322223121013023-2312012011220113-2000301120202213-0212000001233232-0103230330300101-1111330203321121): complete subsection reference.

- [ignore_domain](data-sources--http_loadbalancer--reference--group-021.md#canonical-0312220213101021-3030212313010302-3202012101032100-2111031101023331-3120133123002032-2223302102000101-2330220300103220-0111233100003032): complete subsection reference.

- [ignore_expiry](data-sources--http_loadbalancer--reference--group-021.md#canonical-2022131312122120-0313230201310123-1101001012033031-3210112120013101-3230021320030132-3013030010002033-3312223102022221-3000132002223012): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-021.md#canonical-2330101032030000-0303210323310033-3012130033223300-0002321010133212-1002323223130123-3212131223020023-3310023220121103-3323022002320332): complete subsection reference.

- [ignore_max_age](data-sources--http_loadbalancer--reference--group-021.md#canonical-1020012011300002-0212210002102001-2332232100130203-1310112320331220-0321213103033111-3321203221331232-3233103020213010-2231210113023131): complete subsection reference.

- [ignore_partitioned](data-sources--http_loadbalancer--reference--group-021.md#canonical-1112032110322101-0323332202012031-1111232121233213-0100112000110132-1312301300330230-0203012013200310-2202212102303321-0201113021011112): complete subsection reference.

- [ignore_path](data-sources--http_loadbalancer--reference--group-021.md#canonical-1123231131001201-3301320030020311-0021020221303001-1213311000213321-1133011030330202-1122013213211230-2033103300113223-2230232203220113): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-021.md#canonical-3310323001033201-2132300030302321-0131212102122102-2311300130333121-1021223233022010-0110233321213022-3003311000212200-1100013303321211): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-021.md#canonical-1032301223020021-1112010033311012-1021301212331103-2133132101232310-2120231003021210-1022132222130213-1003021013201232-2031020120332230): complete subsection reference.

- [ignore_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3323210223311103-2002213230221021-1310310300003112-1331323212212201-2232002302031111-3003312203213313-0020013010022100-1310301302011320): complete subsection reference.

<a id="canonical-2023100221122231-3100023321123130-2232220200130011-3303030000313231-1201032133133301-3001212220020233-0101111331021320-0233313132103010"></a>

<a id="canonical-0232131133010032-1111032131020311-2202010202312002-1001220122122022-1001213232222202-1301012200133213-0012120310210002-3230220213022301"></a>

## max_age_value property — response_cookies_to_add / 200212330110 / 7

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-0112013303113300-1102301123010323-0020102330112213-1003313331031230-1131200231110012-0203032313110313-0212120133031311-1201210213323010"></a>

<a id="canonical-0110332131302103-2111012331122221-0111212100311222-0302303023220021-1200212332103132-0003233133123310-1101203232301232-3102231113232301"></a>

## name property — response_cookies_to_add / 200212330110 / 8

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1330320033311122-1203310312122300-2120132031222212-3023011232013222-3032220000322222-1221233102203133-2200302132302020-1213330323301123"></a>

<a id="canonical-3033302202303021-2330112203210001-3123322220331110-1120302001102133-2330332330203233-3310000303313100-1020222023133300-1211312021311332"></a>

## overwrite property — response_cookies_to_add / 200212330110 / 9

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-021.md#canonical-2213213033300102-1001312000303233-2322030203332322-2001210201311300-0230221021321222-1102121031331202-0203231312121202-3021220200223030): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-021.md#canonical-3033010232001333-3312013123310011-2122310320111100-2112121201231211-2223200010133310-2310133112233130-1113122300002022-0212302132121131): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-021.md#canonical-1013210113102030-0000002220023203-2120301333021130-3233130200233111-3000132120010233-3313013030021312-1003112120021100-2133312110113211): complete subsection reference.

- [secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122): complete subsection reference.

<a id="canonical-0321330303012123-1101331211313221-2200103112200211-3113322301101333-2211013222012200-1112201123333133-0310212333321032-3133101031120230"></a>

<a id="canonical-3022123002212023-0310011330113111-2223330232203102-1031102222131311-3103031121303222-1021122233000203-0013110030311201-3020220312022110"></a>

## value property — response_cookies_to_add / 200212330110 / 10

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

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

<a id="canonical-1112013020022133-1033130032222201-3322110003112031-3131321113301031-0330123102302312-2231010302310203-0320030032120112-1331232120111332"></a>

## Next pages — response_cookies_to_add / 200212330110 / 11

- [more_option.response_cookies_to_add.add_httponly](data-sources--http_loadbalancer--reference--group-021.md#canonical-0300223001322130-2030032331032033-1003212011020133-2233320201331011-0030003131232030-2312222132220200-0001123123322032-2313020031031221)
- [more_option.response_cookies_to_add.add_partitioned](data-sources--http_loadbalancer--reference--group-021.md#canonical-2330103311303032-3130031021333331-1202200000210223-1010020332113233-3213302010103030-1032313211131110-1222033112230303-2200112031020203)
- [more_option.response_cookies_to_add.add_secure](data-sources--http_loadbalancer--reference--group-021.md#canonical-1120032113222132-2213222211032003-1322223121013023-2312012011220113-2000301120202213-0212000001233232-0103230330300101-1111330203321121)
- [more_option.response_cookies_to_add.ignore_domain](data-sources--http_loadbalancer--reference--group-021.md#canonical-0312220213101021-3030212313010302-3202012101032100-2111031101023331-3120133123002032-2223302102000101-2330220300103220-0111233100003032)
- [more_option.response_cookies_to_add.ignore_expiry](data-sources--http_loadbalancer--reference--group-021.md#canonical-2022131312122120-0313230201310123-1101001012033031-3210112120013101-3230021320030132-3013030010002033-3312223102022221-3000132002223012)
- [more_option.response_cookies_to_add.ignore_httponly](data-sources--http_loadbalancer--reference--group-021.md#canonical-2330101032030000-0303210323310033-3012130033223300-0002321010133212-1002323223130123-3212131223020023-3310023220121103-3323022002320332)
- [more_option.response_cookies_to_add.ignore_max_age](data-sources--http_loadbalancer--reference--group-021.md#canonical-1020012011300002-0212210002102001-2332232100130203-1310112320331220-0321213103033111-3321203221331232-3233103020213010-2231210113023131)
- [more_option.response_cookies_to_add.ignore_partitioned](data-sources--http_loadbalancer--reference--group-021.md#canonical-1112032110322101-0323332202012031-1111232121233213-0100112000110132-1312301300330230-0203012013200310-2202212102303321-0201113021011112)
- [more_option.response_cookies_to_add.ignore_path](data-sources--http_loadbalancer--reference--group-021.md#canonical-1123231131001201-3301320030020311-0021020221303001-1213311000213321-1133011030330202-1122013213211230-2033103300113223-2230232203220113)
- [more_option.response_cookies_to_add.ignore_samesite](data-sources--http_loadbalancer--reference--group-021.md#canonical-3310323001033201-2132300030302321-0131212102122102-2311300130333121-1021223233022010-0110233321213022-3003311000212200-1100013303321211)
- [more_option.response_cookies_to_add.ignore_secure](data-sources--http_loadbalancer--reference--group-021.md#canonical-1032301223020021-1112010033311012-1021301212331103-2133132101232310-2120231003021210-1022132222130213-1003021013201232-2031020120332230)
- [more_option.response_cookies_to_add.ignore_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3323210223311103-2002213230221021-1310310300003112-1331323212212201-2232002302031111-3003312203213313-0020013010022100-1310301302011320)
- [more_option.response_cookies_to_add.samesite_lax](data-sources--http_loadbalancer--reference--group-021.md#canonical-2213213033300102-1001312000303233-2322030203332322-2001210201311300-0230221021321222-1102121031331202-0203231312121202-3021220200223030)
- [more_option.response_cookies_to_add.samesite_none](data-sources--http_loadbalancer--reference--group-021.md#canonical-3033010232001333-3312013123310011-2122310320111100-2112121201231211-2223200010133310-2310133112233130-1113122300002022-0212302132121131)
- [more_option.response_cookies_to_add.samesite_strict](data-sources--http_loadbalancer--reference--group-021.md#canonical-1013210113102030-0000002220023203-2120301333021130-3233130200233111-3000132120010233-3313013030021312-1003112120021100-2133312110113211)
- [more_option.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0300223001322130-2030032331032033-1003212011020133-2233320201331011-0030003131232030-2312222132220200-0001123123322032-2313020031031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331100010011220-3332122212103210-2231032102221331-3230330330211130-0311032103001303-0320001300200111-2331320011312202-0020201332313203"></a>

## more_option.response_cookies_to_add.add_httponly — add_httponly / 231020120211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.add_httponly

<a id="canonical-2210010000020033-0101102003330202-2013322110320320-1310021201333112-2311111331132102-0311322203120022-3320030101013302-1031023323102110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-1002003210302033-1120010220203233-3030202300310311-0331013100213203-2122033321113221-3121201301130232-0032200222013202-2103013032122033"></a>

## Direct properties — add_httponly / 231020120211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231022210022012-3201002303120213-3001100032001131-3321001120333323-1012132220331331-0320113012302233-3210312320322023-1303132312001002"></a>

## Next pages — add_httponly / 231020120211 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2330103311303032-3130031021333331-1202200000210223-1010020332113233-3213302010103030-1032313211131110-1222033112230303-2200112031020203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233331323132020-1012113200202313-3212322131302111-1032211010101221-2313030122103313-1011231130203300-2301112223312213-0322230012003023"></a>

## more_option.response_cookies_to_add.add_partitioned — add_partitioned / 203203012131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.add_partitioned

<a id="canonical-3212132312101300-1032332332010301-0031323103313131-0021302002002222-1003301203313223-0131222133020203-1201123330303002-0122201303031103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add partitioned.

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

<a id="canonical-0023022103203303-2101122021200200-0323112130000022-0033222112322120-1122332002322002-1022310322122020-3023232320223010-2001011213130133"></a>

## Direct properties — add_partitioned / 203203012131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003230110201220-1131231231301022-2220011002111012-0021112111302123-2012322330032133-3021230130010012-0102313022013020-2311232301233333"></a>

## Next pages — add_partitioned / 203203012131 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1120032113222132-2213222211032003-1322223121013023-2312012011220113-2000301120202213-0212000001233232-0103230330300101-1111330203321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322103203222300-3200223210103220-0312103021131231-2223020030022132-1302300213223313-0312110232210102-3331212131012212-2300220322333031"></a>

## more_option.response_cookies_to_add.add_secure — add_secure / 002231123130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.add_secure

<a id="canonical-0223113220332011-1101231133103130-1121030012013001-2330020323013132-0130110011223033-1232231012120011-2201321002330031-2122121121202231"></a>

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

<a id="canonical-3121323233032102-0020200133303231-0012310103031103-3021020130332202-3113332132300332-3111013302132021-3031232211322323-0113100201103200"></a>

## Direct properties — add_secure / 002231123130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232023020011031-2322333123320100-3001320102010101-1210312322130321-3123133003203221-0233310301111120-1032203200033232-0200320102210202"></a>

## Next pages — add_secure / 002231123130 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0312220213101021-3030212313010302-3202012101032100-2111031101023331-3120133123002032-2223302102000101-2330220300103220-0111233100003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112201010303202-2310330023000301-0211002032210333-3221011201033012-1011113121333331-1130230201011010-3213300131110231-2301021103120002"></a>

## more_option.response_cookies_to_add.ignore_domain — ignore_domain / 311303223321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_domain

<a id="canonical-0212331123303031-2311101123010131-2332322120301021-3122301122130112-3203021301201131-3122330312303331-2103220132231322-3030323201113100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore domain.

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

<a id="canonical-3220323213121310-3123323303201131-1212011302021033-2331312031131033-1132222110213222-2102002020000122-0113112320133000-1101223122111112"></a>

## Direct properties — ignore_domain / 311303223321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110213200120232-1230310333003303-3102001003113300-3320022011203032-0301130010222121-3023110333001103-0230012110101033-1000231302032121"></a>

## Next pages — ignore_domain / 311303223321 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2022131312122120-0313230201310123-1101001012033031-3210112120013101-3230021320030132-3013030010002033-3312223102022221-3000132002223012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320002121330031-1330102203121011-0313123221002001-3003123300010031-1013003332203230-2233203030133331-3023200203011222-1030103002303031"></a>

## more_option.response_cookies_to_add.ignore_expiry — ignore_expiry / 200002230210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-2010321213022011-2021003102023110-3112132321220201-3002300001112310-1210030210320033-0332112313232320-1232312101202222-2311231230212100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore expiry.

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

<a id="canonical-3213233330310102-2231233232323023-2022102301231132-2020331103301232-1122312320331231-2133112223110312-1022312301321131-3333101003023012"></a>

## Direct properties — ignore_expiry / 200002230210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010320031032101-2313333012222201-3323221033023032-1013203113113223-3020121103011030-3332010232313331-3212120000000033-3313121332121011"></a>

## Next pages — ignore_expiry / 200002230210 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2330101032030000-0303210323310033-3012130033223300-0002321010133212-1002323223130123-3212131223020023-3310023220121103-3323022002320332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013233012220013-3222010122111010-0232302332132300-0312323100000330-3120200221320213-1000031311230012-3111112210110000-2300323211113311"></a>

## more_option.response_cookies_to_add.ignore_httponly — ignore_httponly / 033232103122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-2132001130220023-2123200300303001-3302202213222001-1221102012002320-2320302123010323-2012102333333311-1110131003102320-2200003322123303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-2202230103101130-1323221120230120-1203301232112002-1323131303201012-3202311122002233-0323322312110202-1330302120113310-3200030023130230"></a>

## Direct properties — ignore_httponly / 033232103122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113222132231132-3033221300220302-2200112123313331-3210111202012120-1122200230223303-2030112201021301-1203310031233101-2031113120332033"></a>

## Next pages — ignore_httponly / 033232103122 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1020012011300002-0212210002102001-2332232100130203-1310112320331220-0321213103033111-3321203221331232-3233103020213010-2231210113023131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023302031322230-1122103221311312-2013213101112000-2000133113023230-3013011130202333-2132320210200022-0210001003230232-0113121231121103"></a>

## more_option.response_cookies_to_add.ignore_max_age — ignore_max_age / 212320101231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-0123110211333212-0223212133021233-2000031303313232-2300300033110303-2003022231221310-2033221232121130-1322302101011300-3022120101131213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore max age.

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

<a id="canonical-2223310212132132-0333120121110223-3010130210303323-3333020203002231-1213313322230102-3331200311200213-3333300020210022-2333310203221213"></a>

## Direct properties — ignore_max_age / 212320101231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212112301211110-3333301011102310-1023113021100023-1131121000021013-0031032220202221-3311202202220121-2322213201211113-3011031100021132"></a>

## Next pages — ignore_max_age / 212320101231 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1112032110322101-0323332202012031-1111232121233213-0100112000110132-1312301300330230-0203012013200310-2202212102303321-0201113021011112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020020231213003-3000221103303032-0230322130030121-1030310111333120-0023223122222320-1331203120223311-2021233113201211-0232210202222310"></a>

## more_option.response_cookies_to_add.ignore_partitioned — ignore_partitioned / 030111332313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-0021230111112231-3323322032231231-1222211031223232-1212310313101033-0112230022112322-2200212321011012-0122000102313002-3002332232203032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore partitioned.

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

<a id="canonical-1221131001323022-0010133011303322-1030030223110121-1220313001013111-0232201031122103-2203113020201110-2323032202130330-2312132210221121"></a>

## Direct properties — ignore_partitioned / 030111332313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103132120102030-2033332233111112-0112031032201020-3122021132030023-1222222100233003-0010302312112131-0231011103132230-2103333012100211"></a>

## Next pages — ignore_partitioned / 030111332313 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1123231131001201-3301320030020311-0021020221303001-1213311000213321-1133011030330202-1122013213211230-2033103300113223-2230232203220113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132331320213201-3333020011310113-0000010132111032-2322012003332312-0130321233232222-2310333022120300-2230221130012232-1321203110221301"></a>

## more_option.response_cookies_to_add.ignore_path — ignore_path / 131213130011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_path

<a id="canonical-3212210333221011-2033303231212210-3223302312222110-1001203020332203-2100212000300013-0122323220211203-1311002013202131-0132031121111101"></a>

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

<a id="canonical-0020130211233233-1000333012031220-3221121203221211-3320100000322221-1133331013322102-2200332103033023-3012230210102131-1032111331223331"></a>

## Direct properties — ignore_path / 131213130011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323321303123130-2201132000313023-2030131321233210-3122321031333301-1233030212303330-2333131230033010-2230313120032130-3332300110210022"></a>

## Next pages — ignore_path / 131213130011 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3310323001033201-2132300030302321-0131212102122102-2311300130333121-1021223233022010-0110233321213022-3003311000212200-1100013303321211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311213013113302-3310013010103333-3202103311112330-1210103323111021-1320011303122320-1123001230321230-1011310203020201-2200200003101301"></a>

## more_option.response_cookies_to_add.ignore_samesite — ignore_samesite / 032321001100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-2303301021020021-3021100310002011-0213120312212231-3221203310313020-0012200023001321-3100211301233233-3030313230332233-1211322222002131"></a>

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

<a id="canonical-3032222331032333-3332330002320301-1000233000030111-0200320113130302-0130020102111300-0110113030312010-0201112220132302-0303023120113021"></a>

## Direct properties — ignore_samesite / 032321001100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121201122313211-2012222320010021-0222121320001103-0300233330021223-2330002033121000-2123122003032012-0121221213321211-3123032333323211"></a>

## Next pages — ignore_samesite / 032321001100 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1032301223020021-1112010033311012-1021301212331103-2133132101232310-2120231003021210-1022132222130213-1003021013201232-2031020120332230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030230211302122-2020113002131011-2133212220210203-3021033030220003-0202013010132112-3033003312001203-2230302003312131-0011010131333223"></a>

## more_option.response_cookies_to_add.ignore_secure — ignore_secure / 123001023123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_secure

<a id="canonical-0111333201233001-1210201310023113-1313022230312320-2210002221032302-3131233002223302-3321031120200103-0021010321123203-1200003023331101"></a>

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

<a id="canonical-2313103221113311-2313302300112032-0201123213010120-2201221013230201-2121202203302313-2312013031201123-3123121320032012-3131111020302132"></a>

## Direct properties — ignore_secure / 123001023123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121321000322022-3022323210101031-1312212131211322-3323202233301231-1320212112301202-1012001211003002-2022230203100213-1230220003023133"></a>

## Next pages — ignore_secure / 123001023123 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3323210223311103-2002213230221021-1310310300003112-1331323212212201-2232002302031111-3003312203213313-0020013010022100-1310301302011320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113321130220330-3031032002201011-2323012000133202-0131113202231200-0200320300232221-0221013013203333-2223310111212111-3130102132232101"></a>

## more_option.response_cookies_to_add.ignore_value — ignore_value / 332331320030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.ignore_value

<a id="canonical-1213003230331300-1210022201122012-0223322131110033-1013332003303100-1221223232101031-2233232201311301-1030121222310022-0330131121123332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore value.

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

<a id="canonical-1010033331010233-3222121021310300-1120022033032113-0303312123300300-2211023033323200-3203032131023123-0210300213313330-2321021123310101"></a>

## Direct properties — ignore_value / 332331320030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233322313023032-3101302132210210-0300110022033023-0003303310312200-0301000111333311-1200333301100311-1132010231322120-0201333132021302"></a>

## Next pages — ignore_value / 332331320030 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2213213033300102-1001312000303233-2322030203332322-2001210201311300-0230221021321222-1102121031331202-0203231312121202-3021220200223030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201132102103023-3120020203302102-2131310332222320-0131222013011020-3311331111000020-0233103313001202-2323113001131300-1210000013112020"></a>

## more_option.response_cookies_to_add.samesite_lax — samesite_lax / 333320011333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.samesite_lax

<a id="canonical-3202200003013012-3232001111333212-1332212002202121-0312000000033102-1210133320023222-3002330212312132-1013102233333331-1313130320110120"></a>

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

<a id="canonical-0232013202203111-0030123202300230-1000011030111322-1332333030233321-3203012122122010-0033301313000112-3012220001321131-2321000330330002"></a>

## Direct properties — samesite_lax / 333320011333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210132103213303-1132113202303322-2001030210203121-3220021133333122-2332103313201022-3022220300230010-3231223232312301-1103330111332320"></a>

## Next pages — samesite_lax / 333320011333 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3033010232001333-3312013123310011-2122310320111100-2112121201231211-2223200010133310-2310133112233130-1113122300002022-0212302132121131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002133210332023-1102303033102010-1100132221001300-0233200020123333-1022223232013202-1221332032031333-2302320302001012-3223021132312011"></a>

## more_option.response_cookies_to_add.samesite_none — samesite_none / 002203210130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.samesite_none

<a id="canonical-2132321113321230-1122320201302110-0112133101012130-3222103001120103-2003312312133101-0030003303301200-0112222111322010-0300232013330321"></a>

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

<a id="canonical-2330121311120012-1101320312322023-1203122102210231-0201013001012123-0011102010020013-1003011021130012-0311112213022303-1030112120112102"></a>

## Direct properties — samesite_none / 002203210130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233233231001231-0201232032232010-0020312331131333-2303131111010122-2022103003220031-0010001222020020-3010012223022202-0031033132011330"></a>

## Next pages — samesite_none / 002203210130 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1013210113102030-0000002220023203-2120301333021130-3233130200233111-3000132120010233-3313013030021312-1003112120021100-2133312110113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303003330031110-1123330212011230-2302223232312011-1230030301023333-0331020022331201-0231203111301020-2320300230201223-3313203321020123"></a>

## more_option.response_cookies_to_add.samesite_strict — samesite_strict / 201023130131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.samesite_strict

<a id="canonical-1133020122101211-0002222210031222-0110222101113133-2130010111302303-0033021130311032-1003300313113323-1222131133331302-0000101111121033"></a>

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

<a id="canonical-1123211013323103-3032212112033230-3323011021123000-3120103100021101-3331013022010130-2000322223033012-2120030121200202-2130202333321012"></a>

## Direct properties — samesite_strict / 201023130131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003202030210223-2001001032022200-1223303121220201-2310001302010011-1012211023201120-2333302201303131-0102212132101303-2312111320322323"></a>

## Next pages — samesite_strict / 201023130131 / 4

- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100100033033332-0202002322133300-2212321103103311-1231023201113220-3321131031011122-0330132112123310-2211210202132223-0331122010001033"></a>

## more_option.response_cookies_to_add.secret_value — secret_value / 301112102133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- more_option.response_cookies_to_add.secret_value

<a id="canonical-3022202032232212-2130011311110303-2030201002012220-3200111200322300-1130120011210233-0231223220023312-2000203012102132-0002031313011133"></a>

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

<a id="canonical-3213231011232223-3223230300012302-3323231032100113-1333013113113322-3001301031001320-0230323203031210-2223002321113202-0101221010300130"></a>

## Direct properties — secret_value / 301112102133 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-3221030132320202-3330201112031213-2012330013120223-2021202223202021-0130212303210312-1232101010033313-0002202330201133-0131031300020113): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-2223031102111001-0010323200311031-0021330231133232-1220122101032213-0302030312303310-1200122123232132-2210220312130033-2131020131213301): complete subsection reference.

<a id="canonical-2210112030101001-3312303220322331-3322112212211221-2310330323221000-1122113032010331-3113003032021232-0112121210211200-2230312300223110"></a>

## Next pages — secret_value / 301112102133 / 4

- [more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-3221030132320202-3330201112031213-2012330013120223-2021202223202021-0130212303210312-1232101010033313-0002202330201133-0131031300020113)
- [more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-2223031102111001-0010323200311031-0021330231133232-1220122101032213-0302030312303310-1200122123232132-2210220312130033-2131020131213301)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3221030132320202-3330201112031213-2012330013120223-2021202223202021-0130212303210312-1232101010033313-0002202330201133-0131031300020113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012113300012231-0212313011000101-3231003120121113-1212020013320100-3132033100310203-3311201120021221-1311322223321101-1322021021010002"></a>

## more_option.response_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 012311002122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [more_option.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122)
- more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-2323322311100301-2023212100122202-1233211110221220-3332023210000132-1203230031303132-0220303002212302-1131233132101220-2030111132003323"></a>

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

<a id="canonical-0013313002320302-2002110301113200-1203103311330211-2232111303203301-0322030331131022-2210031023011003-3302123322000002-3320030212233223"></a>

## Direct properties — blindfold_secret_info / 012311002122 / 3

<a id="canonical-0211011303230133-3203001120322130-1202331322333221-0201023130331000-1120213103111323-3320302231010301-3021233010212313-0331100331303313"></a>

<a id="canonical-3322010321201321-2223211221003320-3223230210220011-3310022132002102-3130032212012030-2132320122303211-3101233020212021-0323232133231002"></a>

## decryption_provider property — blindfold_secret_info / 012311002122 / 4

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

<a id="canonical-1023312210333200-1213313101302330-2212313311302003-2332133213303200-2231333122302033-2203301120202003-3320112010302012-2031322031321022"></a>

<a id="canonical-3003313002013303-0133202210013322-0113310303012113-3100021230122020-3113101100032333-2012331331330220-2020211020222130-2321130203111020"></a>

## location property — blindfold_secret_info / 012311002122 / 5

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

<a id="canonical-0011110321232320-0100111113300113-3213113311322112-1311203132321203-2113203203313323-3011120023301032-3000312300332232-0113100232121010"></a>

<a id="canonical-3232102000013102-1023200311110032-0021212310220302-3021312323102112-1020021301003101-3031312220023203-0303213233333110-2130213033203230"></a>

## store_provider property — blindfold_secret_info / 012311002122 / 6

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

<a id="canonical-0333123013221022-0311000211212301-2130021210223003-1133010211313332-3121023000022310-3021001221231310-0202100211301030-3033122333100131"></a>

## Next pages — blindfold_secret_info / 012311002122 / 7

- [more_option.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2223031102111001-0010323200311031-0021330231133232-1220122101032213-0302030312303310-1200122123232132-2210220312130033-2131020131213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231112111223213-2300222111323221-0300203223112321-2100132031002012-1133230200210133-1313223230331313-1121300020212223-3223221201111321"></a>

## more_option.response_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 101121101003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132)
- [more_option.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122)
- more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-0131231201022222-1333320222323103-0010121310321311-2212113303322011-2133133222333103-3112223013113120-0112320102222133-1222213130101012"></a>

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

<a id="canonical-0111231020323321-1030333132121021-3020223313303321-3201200133003210-3221233321230132-3222123001332023-0000100231020120-0223221103310131"></a>

## Direct properties — clear_secret_info / 101121101003 / 3

<a id="canonical-3031012323202111-1033312233230033-3310230320022030-0311201100120221-0301001331111301-0020120133302312-0331302102122223-2233320200233323"></a>

<a id="canonical-1030101110003131-3221211231102131-2123022012210300-0301021122313033-1131232320121122-2023201220111230-2030120320330222-1011022202232202"></a>

## provider_ref property — clear_secret_info / 101121101003 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1233331323213123-1120332333300210-3212213310001101-2201321111011111-0311200002212012-3221030103033323-1110012322212112-2002030201122130"></a>

<a id="canonical-1201301312202023-0221033002123123-3232021023011312-0322101011203333-0113223001001320-2332223000201233-0203133233101132-0110311230020302"></a>

## URL property — clear_secret_info / 101121101003 / 5

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

<a id="canonical-3322230101202230-2133000322010321-2131000303120011-0302331001123321-3323013223223013-2021001001203121-1122120332023223-2232033233302212"></a>

## Next pages — clear_secret_info / 101121101003 / 6

- [more_option.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-0302230030012322-3321031112030212-2012330212323332-3012132212011311-0011120231202020-3203220032012223-3212032111310010-2023320221011122)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132130223220231-2000232002231310-2011311232001222-3212113133211113-2232222111011201-3031113223201103-1001222201100312-2021000120022031"></a>

## more_option.response_headers_to_add — response_headers_to_add / 013232311000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.response_headers_to_add

<a id="canonical-3111223013313030-2310133003221013-0030203200203220-3321132313210023-2302201203232323-2132200011303312-0131200132233310-0122322122332200"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3032210011023301-3003023330000320-3301010331220202-3221133030212300-3120010011232011-3223033300113132-2120132333310012-1121030233113032"></a>

## Direct properties — response_headers_to_add / 013232311000 / 3

<a id="canonical-1331131101121321-2100303111213210-1211030110110130-0023003230123112-2222132002011012-2001331130220011-1112122012003333-0220033220130011"></a>

<a id="canonical-0332301311123032-1322223210203130-3011213213320321-1100213033210121-3230202233110100-2112200330031302-0003200030333013-0302111302230020"></a>

## append property — response_headers_to_add / 013232311000 / 4

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2300221123233033-2212121033131301-3311312333300131-2300233213031302-1232001323203203-3031100311231230-3022331123003111-0003022031310113"></a>

<a id="canonical-1101201210211130-3230222321002322-2233321220320112-0033121303222300-1320302300210133-0332202330013012-3303303012232001-1013032032110010"></a>

## name property — response_headers_to_add / 013232311000 / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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

- [secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000): complete subsection reference.

<a id="canonical-1000202123102321-0310332211023122-1123011033213003-3223202222102023-3121311312013100-3002130213320120-0033031120002100-1301202310300132"></a>

<a id="canonical-2301021002031002-2310121300223213-2020300031313300-1021331133321123-3200322221010200-3023230102011222-0121122122133130-1232200111232300"></a>

## value property — response_headers_to_add / 013232311000 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

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

<a id="canonical-2233300030012331-1202022130200032-2332110311001123-0022322030222113-1101003033130303-2122101312011132-1221130033113333-0113210202031110"></a>

## Next pages — response_headers_to_add / 013232311000 / 7

- [more_option.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100311021010313-2232311013030101-2330322323310321-3111001132003000-1020120101211110-3310112311133212-3220123020201020-1133301102211302"></a>

## more_option.response_headers_to_add.secret_value — secret_value / 300002112122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_headers_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121)
- more_option.response_headers_to_add.secret_value

<a id="canonical-1231112121103033-1202232131311230-3203100001110310-0323220003313312-3130221102023320-2210222120123103-2331303230313333-2131222303332023"></a>

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

<a id="canonical-2033311331011312-3122332000223020-3231212332133221-1213123223100333-2011223003200323-0333130323120213-1302203220201120-2120202130032121"></a>

## Direct properties — secret_value / 300002112122 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-0210300123133110-2331100013321200-1031202323100310-1123013031321012-1021100031121211-2313132101122002-3110101102302301-2323102022030001): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-2303032030300103-0313202112202213-1331301030302200-0303313310021033-0313032233023013-0232112131331202-2110032133213001-0233112131102203): complete subsection reference.

<a id="canonical-3123320301000032-0132112323231321-0221023120032231-2101321033101001-1233023310032033-2222321111020223-2202023300123022-2200333213023332"></a>

## Next pages — secret_value / 300002112122 / 4

- [more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-0210300123133110-2331100013321200-1031202323100310-1123013031321012-1021100031121211-2313132101122002-3110101102302301-2323102022030001)
- [more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--http_loadbalancer--reference--group-021.md#canonical-2303032030300103-0313202112202213-1331301030302200-0303313310021033-0313032233023013-0232112131331202-2110032133213001-0233112131102203)
- [more_option.response_headers_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0210300123133110-2331100013321200-1031202323100310-1123013031321012-1021100031121211-2313132101122002-3110101102302301-2323102022030001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121101023331131-2022211020212212-3132320023311321-1203233003120210-3001003232303330-1110212203311131-2322313312233303-3232033231230133"></a>

## more_option.response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 032130111323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_headers_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121)
- [more_option.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000)
- more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3023312322233333-1300020231332312-1232113302011210-2013022230203010-1103333000103211-3101212131210021-3020031231213301-3200103332303223"></a>

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

<a id="canonical-3032011310232012-0030223112000131-3123201301322212-1012310223111002-1022121011312102-0003102212130232-2223030312233331-0300311111300321"></a>

## Direct properties — blindfold_secret_info / 032130111323 / 3

<a id="canonical-0130300122010313-0112020323030302-2233220033112221-2331003010121333-2213030022331332-0031002122130231-2220322232303112-1333010112001200"></a>

<a id="canonical-3203311312230011-1022320201003002-1023313210010030-2121031303013322-1122203013012230-1132031321131303-0002302012302203-0011220002230112"></a>

## decryption_provider property — blindfold_secret_info / 032130111323 / 4

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

<a id="canonical-2321202012112312-2121210021103303-3221202332221123-1332131031021010-0203301200110230-0033300211131302-2331331133013013-0332010322110211"></a>

<a id="canonical-1233102320301102-0000223100230000-3223330110310212-2211133101110021-0103023312021111-0321010011113032-1313202223301102-3222332211223103"></a>

## location property — blindfold_secret_info / 032130111323 / 5

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

<a id="canonical-0110201101200312-2023100222311020-2010003133303210-2302030003321031-0001201032131003-2221120320230332-1023120303222133-1231223120203331"></a>

<a id="canonical-0231201220323131-1223211120032010-1203020003100201-0022221323031000-0200132333121111-2332003032233132-3212031232232203-3012110013133020"></a>

## store_provider property — blindfold_secret_info / 032130111323 / 6

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

<a id="canonical-0013230030132303-2322021223213322-3122123212202203-1030120301213132-3303102013122311-1010301030212300-1132123220302012-1313200310032010"></a>

## Next pages — blindfold_secret_info / 032130111323 / 7

- [more_option.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2303032030300103-0313202112202213-1331301030302200-0303313310021033-0313032233023013-0232112131331202-2110032133213001-0233112131102203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012333113210231-2331100012013331-1311023212302123-0231311033102032-1201311001222003-3013030220320331-0130031021133133-0021111301123000"></a>

## more_option.response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 012303202321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.response_headers_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121)
- [more_option.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000)
- more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2022331010030303-0023102311110131-2123321313311110-2300302130002123-3110310201212123-2010022022032031-3030302311230112-2302312033131130"></a>

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

<a id="canonical-0103132202011123-0331222113302303-2110201101121313-3310200011333113-0212202230233020-2100302321211202-1020102300333311-3010002230213230"></a>

## Direct properties — clear_secret_info / 012303202321 / 3

<a id="canonical-0222321202020133-2101010112302122-2002331202212003-1221022113010222-0003030122213211-0213112011320220-3120011011130320-3111001120002221"></a>

<a id="canonical-1332200133320123-1111200131011001-0131220331211222-2011301200121312-3112303023213202-2331133330110030-2132113230130133-2211030000210011"></a>

## provider_ref property — clear_secret_info / 012303202321 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1103021303130222-2230210032223221-0313131103200031-0002310011200322-3111333312201020-0001002203331023-3333121110201003-3300010022133023"></a>

<a id="canonical-2212123303111302-1300313210120020-2131200300130232-0202121211012302-2120233311200010-0320333320212333-1301230230020021-2023023001310130"></a>

## URL property — clear_secret_info / 012303202321 / 5

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

<a id="canonical-3312132331101322-2220232210012110-3230013002000223-1221313312021023-2111302321121222-0003001230132231-1123102201130013-3202310201000221"></a>

## Next pages — clear_secret_info / 012303202321 / 6

- [more_option.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3212022231010330-0203331033101323-1001132002301230-2012330123131230-2110021112302030-2102103312212011-0310313023132302-0221201132030000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3312033113300312-1002132313113311-3233122303230311-2103222022222003-3000233023022122-2023233330322332-1333021223220210-1021232333203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011333302101123-0111021011120020-1220130131021320-3031102012333010-3231022332210323-3031311213103110-1321213022313212-2000003332223121"></a>

## multi_lb_app — multi_lb_app / 102330332310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- multi_lb_app

<a id="canonical-3303101211012301-0112301312033011-3101222111133102-2112323301002031-3333132011001210-1303100102132022-2302112333113013-3120233030123231"></a>

Type: `["object", {}]`. Computed.

\[OneOf: multi\_lb\_app, single\_lb\_app\] Configuration parameter for multi lb app.

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

OneOf alternatives in this subsection:

- [multi_lb_app](data-sources--http_loadbalancer--reference--group-021.md#canonical-3303101211012301-0112301312033011-3101222111133102-2112323301002031-3333132011001210-1303100102132022-2302112333113013-3120233030123231)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-0031200011131311-2233112323100033-0111001011133120-2221230201121322-2021123130030211-2222030012320123-0323231202112012-3300012012003232)

Select alternatives according to the provider validators above.

<a id="canonical-3312211231210323-0012221202033011-3322033103020032-1013031110231221-3013300203211311-0000122310130201-3102321112021103-2112323002323000"></a>

## Direct properties — multi_lb_app / 102330332310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213301021001333-1303330313020322-0311112201110330-0301321132021030-2332332021002213-0221310201332321-1012133302222332-0232021313032013"></a>

## Next pages — multi_lb_app / 102330332310 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0033322013033333-1012021230311113-3030032222123110-2130202012323311-2130011103202221-3200200101001031-2230202203012132-2012120233023302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020211110133230-3201232301331211-1221320031020332-0200002212313300-3313033311123131-1102310323330331-3302303033033102-3201220130210322"></a>

## no_challenge — no_challenge / 101211023022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- no_challenge

<a id="canonical-0121222330302201-1313000022222103-0000332201021130-0021331132102001-3010230020131302-3021233321102212-0331230031230303-0000030023131011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-1321121321300031-1301010130210121-1133023001201021-1132122133323020-3031030330330211-1321302030033112-2120123323220021-1011113132122313"></a>

## Direct properties — no_challenge / 101211023022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321221131312223-0121103200322122-1012013020213310-1133113002330011-1130320132012321-1211021203123310-2211013321133122-1200221031120223"></a>

## Next pages — no_challenge / 101211023022 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1012220231323000-0002101033211030-3111020332233130-0011331233023201-2102321003221301-1230013230021133-3230200102213220-0022221221200321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303013031232312-2010331232113012-1213222132330110-1002021233303100-1323121312321200-0212333200223310-3333101303311123-3132200300211332"></a>

## no_service_policies — no_service_policies / 113013333303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- no_service_policies

<a id="canonical-2210312233210311-0002123133331020-1220122121222001-0132233321112131-1121331112303320-2113021232203331-1202011022320022-1211301213031113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no service policies.

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

<a id="canonical-3023201032310322-0303110332000111-0213333211222303-3330132001222100-1032012332011230-0211311303310320-3122332322210123-3001122111003322"></a>

## Direct properties — no_service_policies / 113013333303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111131312312220-3320331113100003-2123103233333330-1200023133220303-3001012023132111-0232102310011220-1023202132211230-2130300212033011"></a>

## Next pages — no_service_policies / 113013333303 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213121323112302-0202100331103131-1101123211331323-3001020221001100-1002323322011021-3212300113210122-1311303211312300-0001022001100120"></a>

## origin_server_subset_rule_list — origin_server_subset_rule_list / 122120130331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- origin_server_subset_rule_list

<a id="canonical-0303030233300122-3332220002211232-0333301002233300-1211031003230112-0301311012023111-1131321021222333-3023203303301301-3202001221001112"></a>

Type: `"single"`. Computed.

Origin Server Subset Rule List Type. List of Origin Pools.

Upstream description:

List of Origin Pools.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1133122010130200-3132030311200213-2312103323231203-1003233212323322-1122020333300130-0002110120132202-3201323232031212-2003232232302103"></a>

## Direct properties — origin_server_subset_rule_list / 122120130331 / 3

- [origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113): complete subsection reference.

<a id="canonical-0203103021102321-2003321322123231-0123213023221332-3113032213000320-1310332301111003-2321203023220120-3333030002000001-3101332200111322"></a>

## Next pages — origin_server_subset_rule_list / 122120130331 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101122010233301-2100121131311320-0331012311012031-2100103013012201-3310320301013301-3331022030111033-0132332013032122-0321003131001223"></a>

## origin_server_subset_rule_list.origin_server_subset_rules — origin_server_subset_rules / 023103212032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- origin_server_subset_rule_list.origin_server_subset_rules

<a id="canonical-0002300201201300-0100200220100312-1111220332220113-1301331212323213-3312131311330033-0330313200012320-0310112101221300-3221112020123100"></a>

Type: `"list"`. Computed.

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to..

Upstream description:

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to define the correct order for Origin Server Subset to GET the intended result, rules are evaluated
from top to bottom in the list. When an Origin server subset rule is matched, then this selection
rule takes effect and no more rules are evaluated.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-0302023003211131-0001103033300022-0022323233313030-3100121221102312-0203032302320220-3023112022213012-2231011010120131-1211310002213033"></a>

## Direct properties — origin_server_subset_rules / 023103212032 / 3

- [any_asn](data-sources--http_loadbalancer--reference--group-021.md#canonical-2032000020020010-1123302032300201-2232220322231131-1012023101220201-0302130002001223-3333300112021112-2220300201332001-0211130212230322): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-021.md#canonical-3230013232200221-3212322310012222-2211001322100202-3333010132022233-1113223202303220-1030321113212203-2201011013213133-2132220200022021): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-2313113113021300-2203100123203023-0230212303310111-2310301301023231-1330302032312322-0121213323110311-3033331203021112-1331330030120322): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-2220012312230103-1201130010233021-3102111031311130-0010033203103323-0230113231130031-2111131202211032-2013212123010202-3220003023311113): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-021.md#canonical-2103123233211101-1110032203230202-2002021201211231-0022011033332321-3321332221312222-3333023200110202-0103113000313133-2232303110110310): complete subsection reference.

<a id="canonical-1031031132022321-1131002303133210-0020121223323233-0223031220032212-1232312123221133-1201130012222312-1323311221210300-2012131223201220"></a>

<a id="canonical-3122002002012221-3011212003022312-2323311021022003-0122032003320333-2010231221332323-2202012203311302-2103010303311000-1130232220332133"></a>

## country_codes property — origin_server_subset_rules / 023103212032 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Country Codes List. List of Country Codes. Possible values are \`COUNTRY\_NONE\`, \`COUNTRY\_AD\`,
\`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`, \`COUNTRY\_AI\`, \`COUNTRY\_AL\`,
\`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`, \`COUNTRY\_AQ\`, \`COUNTRY\_AR\`,
\`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`, \`COUNTRY\_AW\`, \`COUNTRY\_AX\`,
\`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`, \`COUNTRY\_BD\`, \`COUNTRY\_BE\`,
\`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`, \`COUNTRY\_BI\`, \`COUNTRY\_BJ\`,
\`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`, \`COUNTRY\_BO\`, \`COUNTRY\_BQ\`,
\`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`, \`COUNTRY\_BV\`, \`COUNTRY\_BW\`,
\`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`, \`COUNTRY\_CC\`, \`COUNTRY\_CD\`,
\`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`, \`COUNTRY\_CI\`, \`COUNTRY\_CK\`,
\`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`, \`COUNTRY\_CO\`, \`COUNTRY\_CR\`,
\`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`, \`COUNTRY\_CW\`, \`COUNTRY\_CX\`,
\`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`, \`COUNTRY\_DJ\`, \`COUNTRY\_DK\`,
\`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`, \`COUNTRY\_EC\`, \`COUNTRY\_EE\`,
\`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`, \`COUNTRY\_ES\`, \`COUNTRY\_ET\`,
\`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`, \`COUNTRY\_FM\`, \`COUNTRY\_FO\`,
\`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`, \`COUNTRY\_GD\`, \`COUNTRY\_GE\`,
\`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`, \`COUNTRY\_GI\`, \`COUNTRY\_GL\`,
\`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`, \`COUNTRY\_GQ\`, \`COUNTRY\_GR\`,
\`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`, \`COUNTRY\_GW\`, \`COUNTRY\_GY\`,
\`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`, \`COUNTRY\_HR\`, \`COUNTRY\_HT\`,
\`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`, \`COUNTRY\_IL\`, \`COUNTRY\_IM\`,
\`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`, \`COUNTRY\_IR\`, \`COUNTRY\_IS\`,
\`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`, \`COUNTRY\_JO\`, \`COUNTRY\_JP\`,
\`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`, \`COUNTRY\_KI\`, \`COUNTRY\_KM\`,
\`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`, \`COUNTRY\_KW\`, \`COUNTRY\_KY\`,
\`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`, \`COUNTRY\_LC\`, \`COUNTRY\_LI\`,
\`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`, \`COUNTRY\_LT\`, \`COUNTRY\_LU\`,
\`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`, \`COUNTRY\_MC\`, \`COUNTRY\_MD\`,
\`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`, \`COUNTRY\_MH\`, \`COUNTRY\_MK\`,
\`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`, \`COUNTRY\_MO\`, \`COUNTRY\_MP\`,
\`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`, \`COUNTRY\_MT\`, \`COUNTRY\_MU\`,
\`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`, \`COUNTRY\_MY\`, \`COUNTRY\_MZ\`,
\`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`, \`COUNTRY\_NF\`, \`COUNTRY\_NG\`,
\`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`, \`COUNTRY\_NP\`, \`COUNTRY\_NR\`,
\`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`, \`COUNTRY\_PA\`, \`COUNTRY\_PE\`,
\`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`, \`COUNTRY\_PK\`, \`COUNTRY\_PL\`,
\`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`, \`COUNTRY\_PS\`, \`COUNTRY\_PT\`,
\`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`, \`COUNTRY\_RE\`, \`COUNTRY\_RO\`,
\`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`, \`COUNTRY\_SA\`, \`COUNTRY\_SB\`,
\`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`, \`COUNTRY\_SG\`, \`COUNTRY\_SH\`,
\`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`, \`COUNTRY\_SL\`, \`COUNTRY\_SM\`,
\`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`, \`COUNTRY\_SS\`, \`COUNTRY\_ST\`,
\`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`, \`COUNTRY\_SZ\`, \`COUNTRY\_TC\`,
\`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`, \`COUNTRY\_TH\`, \`COUNTRY\_TJ\`,
\`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`, \`COUNTRY\_TN\`, \`COUNTRY\_TO\`,
\`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`, \`COUNTRY\_TW\`, \`COUNTRY\_TZ\`,
\`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`, \`COUNTRY\_US\`, \`COUNTRY\_UY\`,
\`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`, \`COUNTRY\_VE\`, \`COUNTRY\_VG\`,
\`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`, \`COUNTRY\_WF\`, \`COUNTRY\_WS\`,
\`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`, \`COUNTRY\_YT\`, \`COUNTRY\_ZA\`,
\`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

List of Country Codes.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-1303030032032131-1112131301012021-1211031323120231-1112220312122233-1330203201231213-0133101222312300-0200032222021123-2031211213323002): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3303312131012321-3302333210230231-2123131120233002-0011002032332003-3331302301210111-0311032201303101-0323212202312030-3230123212332022): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-021.md#canonical-3310311213320212-1221113303321301-3312133113302311-2203200311323223-1033032331103013-3312230022130121-0310110333322033-1132230223230331): complete subsection reference.

- [none](data-sources--http_loadbalancer--reference--group-021.md#canonical-0111233133120120-0323011111302022-2230023031221212-1230110330012322-2211321231001201-2002321333213021-1021202022222123-0122221101002230): complete subsection reference.

<a id="canonical-2112101032222020-2002233112211323-2000320210020201-0131120200012111-1202313023311220-2330312111213221-0010033321121112-3320011103201232"></a>

<a id="canonical-0111001230212122-3233230230000303-1232303031310300-3232123132113333-2102200111110331-2222130303303101-3322022123320130-0323132322223032"></a>

## origin_server_subsets_action property — origin_server_subset_rules / 023103212032 / 5

Type: `["map", "string"]`. Computed.

Add labels to select one or more origin servers.

Upstream description:

Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured
in the origin pool are: &#8203;1. Add labels to origin servers &#8203;2. Enable subset load
balancing in the Origin Server Subsets section and configure keys in origin server subsets classes.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1101221301132003-3003201321003130-2333302203011230-1311131331023220-0230300222023130-0102302332002323-1330311000220021-1000202130113113"></a>

<a id="canonical-3213103311123123-2010333210310131-2322320133120333-2333213033130013-2330203330312203-0100123231011220-2012013000220132-1223021130201202"></a>

## re_name_list property — origin_server_subset_rules / 023103212032 / 6

Type: `["list", "string"]`. Computed.

RE Names. List of RE names for match.

Upstream description:

List of RE names for match.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0110010303203102-0132310313213310-2021303313013332-2202020013113322-1303331130300300-3203230312300113-2131302323300131-3302122230222210"></a>

## Next pages — origin_server_subset_rules / 023103212032 / 7

- [origin_server_subset_rule_list.origin_server_subset_rules.any_asn](data-sources--http_loadbalancer--reference--group-021.md#canonical-2032000020020010-1123302032300201-2232220322231131-1012023101220201-0302130002001223-3333300112021112-2220300201332001-0211130212230322)
- [origin_server_subset_rule_list.origin_server_subset_rules.any_ip](data-sources--http_loadbalancer--reference--group-021.md#canonical-3230013232200221-3212322310012222-2211001322100202-3333010132022233-1113223202303220-1030321113212203-2201011013213133-2132220200022021)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-2313113113021300-2203100123203023-0230212303310111-2310301301023231-1330302032312322-0121213323110311-3033331203021112-1331330030120322)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-2220012312230103-1201130010233021-3102111031311130-0010033203103323-0230113231130031-2111131202211032-2013212123010202-3220003023311113)
- [origin_server_subset_rule_list.origin_server_subset_rules.client_selector](data-sources--http_loadbalancer--reference--group-021.md#canonical-2103123233211101-1110032203230202-2002021201211231-0022011033332321-3321332221312222-3333023200110202-0103113000313133-2232303110110310)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-1303030032032131-1112131301012021-1211031323120231-1112220312122233-1330203201231213-0133101222312300-0200032222021123-2031211213323002)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3303312131012321-3302333210230231-2123131120233002-0011002032332003-3331302301210111-0311032201303101-0323212202312030-3230123212332022)
- [origin_server_subset_rule_list.origin_server_subset_rules.metadata](data-sources--http_loadbalancer--reference--group-021.md#canonical-3310311213320212-1221113303321301-3312133113302311-2203200311323223-1033032331103013-3312230022130121-0310110333322033-1132230223230331)
- [origin_server_subset_rule_list.origin_server_subset_rules.none](data-sources--http_loadbalancer--reference--group-021.md#canonical-0111233133120120-0323011111302022-2230023031221212-1230110330012322-2211321231001201-2002321333213021-1021202022222123-0122221101002230)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2032000020020010-1123302032300201-2232220322231131-1012023101220201-0302130002001223-3333300112021112-2220300201332001-0211130212230322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330032111110011-1311101200031231-2130322123013023-2032013231221033-0313233100310032-3030010230013333-2321103232322122-1322002332122200"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.any_asn — any_asn / 230222120130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.any_asn

<a id="canonical-2323022213231322-1200212323010120-3121132003010100-3323023011323023-2313122223303121-2032230112031132-1203032322102001-2203231100121110"></a>

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

<a id="canonical-3032121030113131-3323123303130303-1000033311323031-1102230102120121-2300111210220021-2103110331122322-1220230120303023-0300100213033022"></a>

## Direct properties — any_asn / 230222120130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023231001022010-2230312131012310-2331103023332210-2132102223122333-1333023213011221-1212221321311330-1010321202300313-2002210333110321"></a>

## Next pages — any_asn / 230222120130 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3230013232200221-3212322310012222-2211001322100202-3333010132022233-1113223202303220-1030321113212203-2201011013213133-2132220200022021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032100013210312-2112110330003303-0123020323322220-1300110321301210-3132033022010211-1220122332100313-3230020201133020-3213222330222133"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.any_ip — any_ip / 011320211300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.any_ip

<a id="canonical-1230033312100101-3021031120123003-1332300130112103-1230202333220221-0210301032310103-0212002202203220-3313201302021302-1321131003001223"></a>

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

<a id="canonical-2313303331101011-3232310032331212-0021323022200130-3302323101230013-1000223231330331-3203030333302332-0213022322312221-0313222032212121"></a>

## Direct properties — any_ip / 011320211300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233112331302321-3200232322221130-1220031320122032-1002121323310201-0330223103012321-0321023123123021-1213110110112311-0032000313031223"></a>

## Next pages — any_ip / 011320211300 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2313113113021300-2203100123203023-0230212303310111-2310301301023231-1330302032312322-0121213323110311-3033331203021112-1331330030120322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122202023332102-3013322003002031-2320011032123000-1222212202030220-3233121333203011-1103211113302301-0003122130021333-0221132212311231"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.asn_list — asn_list / 102303002211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_list

<a id="canonical-3132333230033232-3032331310223000-3002012213023232-3022233020022233-2203211320312301-2330012201103111-1221132311003033-3222332022000130"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

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

<a id="canonical-1323102120012033-3101111012003020-2303301033220311-1301201101023031-0231221100122022-1211032323310030-3001033233012220-1323102210121023"></a>

## Direct properties — asn_list / 102303002211 / 3

<a id="canonical-2111022031310330-3223112120332001-1220303301120203-1033322331232102-2020311020320032-1302030013301323-0113121032302110-3132012233110331"></a>

<a id="canonical-2132311020100331-2233113321203222-1023200030330330-0321133001310120-2331121210022000-3223320313212231-1120001120103111-1220301101222032"></a>

## as_numbers property — asn_list / 102303002211 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

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

<a id="canonical-0231333113333010-2123213003320221-2130322202213200-2122231022333022-2211022022233300-3132011012203212-2113230300313322-1102020303301210"></a>

## Next pages — asn_list / 102303002211 / 5

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2220012312230103-1201130010233021-3102111031311130-0010033203103323-0230113231130031-2111131202211032-2013212123010202-3220003023311113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301211101222033-2301120011233131-2133323230201221-3033200313323102-2323301032202233-0210033211221033-0122321233110303-1313323023031110"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher — asn_matcher / 302311121123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher

<a id="canonical-3211120030102301-3013320101331223-3232121322113031-1301222002020320-2323031103130211-0311110201021013-3302000113001130-1202000010333100"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0103323321321010-0122112130200013-0300120102031101-0033220332130310-2132012121121103-3212323332101202-3203220313233003-2120030011221233"></a>

## Direct properties — asn_matcher / 302311121123 / 3

- [asn_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-1332131120323011-2223232132012303-2133332011330200-0230332230031002-1100111232122322-0112232222002300-1223203131110001-1100321320230323): complete subsection reference.

<a id="canonical-3200013221311030-3020123220001330-3333100103100200-3222121000000111-2030123002022210-1011101023313203-2212123202000321-1212002130213330"></a>

## Next pages — asn_matcher / 302311121123 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-1332131120323011-2223232132012303-2133332011330200-0230332230031002-1100111232122322-0112232222002300-1223203131110001-1100321320230323)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1332131120323011-2223232132012303-2133332011330200-0230332230031002-1100111232122322-0112232222002300-1223203131110001-1100321320230323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110022222021133-0323210123232111-3002313200331111-2010102013003100-0113300130121230-3123320322122230-3312111313220312-1013011130031212"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets — asn_sets / 122222110303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-2220012312230103-1201130010233021-3102111031311130-0010033203103323-0230113231130031-2111131202211032-2013212123010202-3220003023311113)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets

<a id="canonical-1002231233131302-2103201001220320-0033300110000312-0203123201311231-2133103103212012-1331311033303223-1111011332212001-2101020111223033"></a>

Type: `"list"`. Computed.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-1120320031201132-1212111122132132-3011331211303331-0211231200120133-1233321030213011-2131301011131230-0232130332100103-0002202031331231"></a>

## Direct properties — asn_sets / 122222110303 / 3

<a id="canonical-1123311103230003-3032132231223102-0100220221332103-3213031332032113-2000030032200200-0130322110200032-1333112302202332-0003200301122011"></a>

<a id="canonical-2210003222310321-0133202300121320-2133121110032121-0010200201302122-3021332103030111-2301230313330220-3210311223300330-1331231110321022"></a>

## kind property — asn_sets / 122222110303 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1330310200032312-1113031031023030-1233302020333301-0300031311220213-1210200131220010-1230201210021011-1232022103111332-3310031230100133"></a>

<a id="canonical-2213203211202112-0101331221120201-1300110333033222-2220023333200033-3000002202230223-2003120003110031-2313333100032112-2313101203031303"></a>

## name property — asn_sets / 122222110303 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3331021030223322-1122302112303210-0111332130232022-2110333231333003-2121012233231132-0130020323033120-1300031312032101-3112131222332210"></a>

<a id="canonical-0121231121103001-0212303200011101-2220003333031202-0232222002201120-0313330201300233-0301201101033032-3201220123033222-0033033030231030"></a>

## namespace property — asn_sets / 122222110303 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  }
}
```

<a id="canonical-2200032103002103-1013123232313102-0111300132312133-2100301132101202-3231312332110232-1110102220200202-1022121212203111-0330023030301331"></a>

<a id="canonical-3020000201002102-1031210310010322-2003302112122113-0233212112222332-2123021123322302-3222202013232311-2220120100021122-0111302223310130"></a>

## tenant property — asn_sets / 122222110303 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2211033002332221-3000213122311130-0200021103020320-3112101022222100-2223300333231202-1123122210022113-3003103003231212-0213020320130232"></a>

<a id="canonical-1223313031233213-3201032223231220-2122310003333332-3121002221021230-3310133323121111-0122201133103331-1003010131333311-0030332303011232"></a>

## uid property — asn_sets / 122222110303 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-1012013210003000-3102321202222021-2230222122201313-2212011210210213-1311211132231233-1311031120303223-1100310113200331-1323330100032333"></a>

## Next pages — asn_sets / 122222110303 / 9

- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-2220012312230103-1201130010233021-3102111031311130-0010033203103323-0230113231130031-2111131202211032-2013212123010202-3220003023311113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2103123233211101-1110032203230202-2002021201211231-0022011033332321-3321332221312222-3333023200110202-0103113000313133-2232303110110310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110021303320012-3100033032200212-1010322332233012-0300133310333100-3010113331110221-1310210321331323-3021210221222032-0212313122230333"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.client_selector — client_selector / 032320012123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.client_selector

<a id="canonical-0013310020221101-2000013030322220-2312313301110132-3303133002010113-2113210311203302-2200112011312013-2002323211023003-2301020201122203"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0033313300111111-2333030121310100-0203121002123023-1333222123310313-3120002110013113-3010030322311323-3003020100233111-3300031132220021"></a>

## Direct properties — client_selector / 032320012123 / 3

<a id="canonical-1203101203332301-0231113132313230-1320111130211102-2311002330103031-3001131133330032-2303111002310202-0222000013103020-0021012103320121"></a>

<a id="canonical-3222323001120102-2231330223221122-1101000203013202-1033233213112000-0231033020121232-0020300001301220-1213113101211312-0000323133011013"></a>

## expressions property — client_selector / 032320012123 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0131332021232111-3203232001031230-0331112223332232-0122132222232201-0301302130133130-3122032230012113-1230222323032032-2200220030332001"></a>

## Next pages — client_selector / 032320012123 / 5

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1303030032032131-1112131301012021-1211031323120231-1112220312122233-1330203201231213-0133101222312300-0200032222021123-2031211213323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230110223020300-1302013320232312-0322030333323011-2112211303332332-1100211332022000-1332022320323121-2332330120211033-3033130033223021"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher — ip_matcher / 313033221122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

<a id="canonical-2021203220112022-2301120121233022-0010132303113231-2131202131003110-1112112112021100-0132303122102301-3131332220231123-2132013030130231"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3212123011222201-1312033112130001-2011311010210331-1213201132020220-1031123120212222-1223033330223000-1033312133313201-1312203111102021"></a>

## Direct properties — ip_matcher / 313033221122 / 3

<a id="canonical-3322220000020103-2012222131020003-2122223033310221-3003211002130100-1120131003123100-3202122011311321-1223103122023012-2232220223030130"></a>

<a id="canonical-1030223100330222-2102102310022100-3012233021220330-0031030031332003-3323230131201212-1113212212223001-1121223113120310-3132302323231210"></a>

## invert_matcher property — ip_matcher / 313033221122 / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [prefix_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-0133331230123333-2203000212003303-3302122222033113-3131002322322033-2121033010211002-1221011313010111-3102230202231223-0131311131122131): complete subsection reference.

<a id="canonical-2103202022330000-0330330012331133-0101002330212323-3030102213211331-3030202322010331-0311223010220000-2301131021122031-0311221112003002"></a>

## Next pages — ip_matcher / 313033221122 / 5

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-021.md#canonical-0133331230123333-2203000212003303-3302122222033113-3131002322322033-2121033010211002-1221011313010111-3102230202231223-0131311131122131)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0133331230123333-2203000212003303-3302122222033113-3131002322322033-2121033010211002-1221011313010111-3102230202231223-0131311131122131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312320131100301-1020233101222131-1303303013211230-0330231013130330-0120223303221122-1123322213313302-1002302131030122-0100112201012110"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets — prefix_sets / 031231020220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-1303030032032131-1112131301012021-1211031323120231-1112220312122233-1330203201231213-0133101222312300-0200032222021123-2031211213323002)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets

<a id="canonical-2003132012113010-2213111022120321-0213102320332201-3321233032303000-1013322312103333-1033002130232210-3201230011310200-2022113200033213"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-1013312302032230-0221310023202122-1131121220310000-0302102333320023-2011111131020302-2230110030010023-1121021120303331-1331222221220031"></a>

## Direct properties — prefix_sets / 031231020220 / 3

<a id="canonical-2121301102320312-2132121311032011-3112021031120233-0123303312132101-3223011230003000-1030320232013121-2333320101223231-3211310102223332"></a>

<a id="canonical-2102032001230211-0331112131012322-0113202032323101-3021110333211310-3020301103120113-3332020013021201-3233313311113033-1333001113321133"></a>

## kind property — prefix_sets / 031231020220 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1233311020001130-2210003301131013-3323321121220022-0102000010012111-3320133101210321-0321213332203213-0121220011011112-1233002221233132"></a>

<a id="canonical-1231111013332133-0023030100102322-1202010310203123-1331210120010220-1233023300210132-0122320003311103-3032303102120303-2221211020100213"></a>

## name property — prefix_sets / 031231020220 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2100302321132222-2300002220222230-3121300113002133-1103113311121231-2312312101031011-2232123322101100-2013202313331011-3012033322203003"></a>

<a id="canonical-0103020121333032-3221101312013013-2133033303030100-2232303332322310-1002121122323220-2013121102102133-2303200202120103-0013211210033001"></a>

## namespace property — prefix_sets / 031231020220 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  }
}
```

<a id="canonical-1100011113321013-0032113230011302-0203122002300021-0101023203112233-0210010111210102-2332331131002111-1032211130323333-0321231012021021"></a>

<a id="canonical-0211013132032011-3002301313131301-2100003201232031-0122211333131031-2132300331022130-0232323213101133-2313220320021103-1133132313001023"></a>

## tenant property — prefix_sets / 031231020220 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2203233111030212-0201320112010200-1302223300233211-1312032110222223-1000113033200123-3110213001031101-1122312312122103-1103122302032323"></a>

<a id="canonical-1012101323211231-2323323333112011-3022112232133211-3321101013333020-1210213010233133-2200301122033103-1031103201201310-1230022223000121"></a>

## uid property — prefix_sets / 031231020220 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0211002132001322-3331102131022330-0221020013330112-3320232302033321-0123033303202110-2232123112233031-0033032211301213-3302001202303232"></a>

## Next pages — prefix_sets / 031231020220 / 9

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](data-sources--http_loadbalancer--reference--group-021.md#canonical-1303030032032131-1112131301012021-1211031323120231-1112220312122233-1330203201231213-0133101222312300-0200032222021123-2031211213323002)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3303312131012321-3302333210230231-2123131120233002-0011002032332003-3331302301210111-0311032201303101-0323212202312030-3230123212332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102002010002302-2320000302201231-2333231022111303-1131320231301301-2312102213121123-1101132002031020-1223022012222331-3223310120002311"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list — ip_prefix_list / 310323021011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list

<a id="canonical-1231120132111110-0302201332013110-3011313003213022-0101220021011203-0113001023210111-0313111021021100-0300232300220010-1111313132213303"></a>

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

<a id="canonical-2211331302322113-3202132031222113-2332231300213320-2210033210231110-0120313100003122-3222331012122230-3003303133003010-3101132103011322"></a>

## Direct properties — ip_prefix_list / 310323021011 / 3

<a id="canonical-1030213202333023-3023200133032220-2230300031222211-1121122130032013-2210013013011112-1220311003013211-1101310010000212-0202321203222000"></a>

<a id="canonical-1232310331312221-1231003033220330-3033000201313212-2200300120320222-1233002320300121-0322133023301003-1130003113210020-3303232132223121"></a>

## invert_match property — ip_prefix_list / 310323021011 / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1113232330210222-3320102230110200-1132232020031021-2130132110122222-0301011013321110-0103211310020230-1312032230012201-3131210312220100"></a>

<a id="canonical-1030200212032321-0013100332103032-3232100110212231-1132310220130021-2030202233200232-0320033123312121-2303232101020112-1333010121021212"></a>

## ip_prefixes property — ip_prefix_list / 310323021011 / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1103113213220311-1201130313132323-1310301301201021-3332121300020022-1213303131002023-1122012022312103-3333022333122133-0002302013020230"></a>

## Next pages — ip_prefix_list / 310323021011 / 6

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3310311213320212-1221113303321301-3312133113302311-2203200311323223-1033032331103013-3312230022130121-0310110333322033-1132230223230331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133101222230110-3330132132320300-0030332032133332-2201321222202100-0210021221302113-2003130022332300-0003200022121103-1220013203023303"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.metadata — metadata / 220103301232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.metadata

<a id="canonical-2300211303021310-2223312113033301-0022231122301233-2023300110020200-0122332131121231-0000211113301033-0010001121133230-0031100011130200"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-0313130002120332-0331213311221033-0301121333233030-3202320020031310-2230132323113010-0202111231332120-2330100331110232-3220103001013122"></a>

## Direct properties — metadata / 220103301232 / 3

<a id="canonical-3203310101131030-0112123223313210-0233122201320021-3011113310230213-1000102113222211-3112302223123013-3113110322012230-1303113123013013"></a>

<a id="canonical-0221020101322300-1023230133301111-3202331323011002-1201311333231222-0333033021032233-3331310320122321-2222002123000330-3233220111231200"></a>

## description_spec property — metadata / 220103301232 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0221011220020100-0010102100300023-2320320123032101-0011130120021001-3302003202133213-2323021212110231-0320310120323332-3200202302220113"></a>

<a id="canonical-1033002212212113-3213011230300011-2132311032122223-0122012002302233-3301303113220020-0010020033120332-1330023310133002-3212011123010213"></a>

## name property — metadata / 220103301232 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-3120033212203013-1220311103323122-2211300002032103-2003211011221113-2213122121312233-1022123212132322-2312201330012033-0132131210102000"></a>

## Next pages — metadata / 220103301232 / 6

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0111233133120120-0323011111302022-2230023031221212-1230110330012322-2211321231001201-2002321333213021-1021202022222123-0122221101002230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301200023313312-2333002201023001-3001212002303202-0333330222131320-2201322213001000-3102011031211031-1221330132033022-0210102003221310"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.none — none / 232102310333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.none

<a id="canonical-3220000002130022-1013102213221233-0020133331001223-2221211223130013-3103032131101232-2102232223203103-2301220023222223-2331231202200323"></a>

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

<a id="canonical-1033201331203130-0121332213013131-1230000222112231-3211011221113322-1220100111322232-0330210310310122-3203212020031333-1031303003321033"></a>

## Direct properties — none / 232102310333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321211331323130-3223013110330310-1331220300203200-2120232222000032-1231023332130130-0333013003322011-1232033312321101-0212203210030022"></a>

## Next pages — none / 232102310333 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122312030310313-1310321231321303-3023122000310030-3110231122300131-0321210321323122-2001222023113023-0123101113300210-1313221313100331"></a>

## policy_based_challenge — policy_based_challenge / 030201021020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- policy_based_challenge

<a id="canonical-0000010110233333-3321322033031133-3333123333332012-3000223230222011-1100121210223232-2221310130103120-2032020232112022-0210222331102001"></a>

Type: `"single"`. Computed.

Specifies the settings for policy rule based challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

<a id="canonical-3010322203022302-1310003302201303-0222103301111231-0322212331131133-2022321221213231-2233001130213010-1001000000102322-0122222203101213"></a>

## Direct properties — policy_based_challenge / 030201021020 / 3

- [always_enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-2200212312300130-0003330022130010-2311321111322302-2110000112113100-1213102032030312-1303002201233032-3003233221121221-1333312030100221): complete subsection reference.

- [always_enable_js_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133222000131230-0023032200003112-1111033230311311-3113100333330212-1131301231310001-0331222323011202-2222330220133121-3311322232103110): complete subsection reference.

- [captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-0031232123230311-0211102100302300-0122202222231323-1011002121100331-1103000211323133-0100003031322133-3132130021030323-0102120103131112): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-3033122122132230-3012313221123333-2130013103000000-2130320102210233-2202310122202323-0320010320033221-3010302311012312-1221131320330202): complete subsection reference.

- [default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-3210233223003011-1123023021123130-0201102132101200-0232230000032100-1303300132120223-2103303302313301-3111233202020030-1333221121021200): complete subsection reference.

- [default_mitigation_settings](data-sources--http_loadbalancer--reference--group-021.md#canonical-3001110002120201-1020301121231202-0010013333031201-0123200211120221-0211033021233231-0012210122330320-3220310110222103-2012203210302011): complete subsection reference.

- [default_temporary_blocking_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-2313131231020200-2210311011111232-0132022020301311-2203023202030330-1210133110122201-3131320010032300-0101231220213330-1321102212211301): complete subsection reference.

- [js_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-1312032111102320-0120213331303133-3332231110231301-3123320201021121-0100212113022232-2312321322331120-1131130322332031-0221220001021103): complete subsection reference.

- [malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-021.md#canonical-0301011213301031-1302102201002310-1003103333312111-2320002321312333-1213202200312033-3203123032001123-1313300002102203-0212131202331300): complete subsection reference.

- [no_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1120331230101123-3020020300310231-3110322023000131-2130202113002001-2202123002003002-0100220202023003-2230213230101000-2323120133200032): complete subsection reference.

- [rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212): complete subsection reference.

- [temporary_user_blocking](data-sources--http_loadbalancer--reference--group-022.md#canonical-3113023112302332-0333303302133233-0010212202113330-1313232030220213-3320210011012203-0101333020233000-0132210002222132-1003010103121113): complete subsection reference.

<a id="canonical-0212200133001300-0122032302213023-0311201101202001-2331013120100130-1012032203002321-2322330222130221-2010331330202223-1201301003112303"></a>

## Next pages — policy_based_challenge / 030201021020 / 4

- [policy_based_challenge.always_enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-2200212312300130-0003330022130010-2311321111322302-2110000112113100-1213102032030312-1303002201233032-3003233221121221-1333312030100221)
- [policy_based_challenge.always_enable_js_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133222000131230-0023032200003112-1111033230311311-3113100333330212-1131301231310001-0331222323011202-2222330220133121-3311322232103110)
- [policy_based_challenge.captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-0031232123230311-0211102100302300-0122202222231323-1011002121100331-1103000211323133-0100003031322133-3132130021030323-0102120103131112)
- [policy_based_challenge.default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-3033122122132230-3012313221123333-2130013103000000-2130320102210233-2202310122202323-0320010320033221-3010302311012312-1221131320330202)
- [policy_based_challenge.default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-3210233223003011-1123023021123130-0201102132101200-0232230000032100-1303300132120223-2103303302313301-3111233202020030-1333221121021200)
- [policy_based_challenge.default_mitigation_settings](data-sources--http_loadbalancer--reference--group-021.md#canonical-3001110002120201-1020301121231202-0010013333031201-0123200211120221-0211033021233231-0012210122330320-3220310110222103-2012203210302011)
- [policy_based_challenge.default_temporary_blocking_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-2313131231020200-2210311011111232-0132022020301311-2203023202030330-1210133110122201-3131320010032300-0101231220213330-1321102212211301)
- [policy_based_challenge.js_challenge_parameters](data-sources--http_loadbalancer--reference--group-021.md#canonical-1312032111102320-0120213331303133-3332231110231301-3123320201021121-0100212113022232-2312321322331120-1131130322332031-0221220001021103)
- [policy_based_challenge.malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-021.md#canonical-0301011213301031-1302102201002310-1003103333312111-2320002321312333-1213202200312033-3203123032001123-1313300002102203-0212131202331300)
- [policy_based_challenge.no_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1120331230101123-3020020300310231-3110322023000131-2130202113002001-2202123002003002-0100220202023003-2230213230101000-2323120133200032)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.temporary_user_blocking](data-sources--http_loadbalancer--reference--group-022.md#canonical-3113023112302332-0333303302133233-0010212202113330-1313232030220213-3320210011012203-0101333020233000-0132210002222132-1003010103121113)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2200212312300130-0003330022130010-2311321111322302-2110000112113100-1213102032030312-1303002201233032-3003233221121221-1333312030100221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012033110212000-0221132313112021-0230323031113210-1023030102300223-3023103111002302-0230211033203223-1211300020122331-3332203111002013"></a>

## policy_based_challenge.always_enable_captcha_challenge — always_enable_captcha_challenge / 121321010200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-3121012012013211-3323213133230010-0121322031120011-1103130301030223-3222203222102032-2112021300023120-2013210103032210-3301132303220102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for always enable captcha challenge.

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

<a id="canonical-2010132213300321-1001302133320322-0200130222302111-3031103322330320-1023132122010313-1123132302001303-1230110310230010-1221022232103002"></a>

## Direct properties — always_enable_captcha_challenge / 121321010200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303132031202121-1132002020211102-1213233302013331-1333031023021100-2231221132131020-2013223312103301-1023130012111011-3001112332022102"></a>

## Next pages — always_enable_captcha_challenge / 121321010200 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1133222000131230-0023032200003112-1111033230311311-3113100333330212-1131301231310001-0331222323011202-2222330220133121-3311322232103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101110203201130-1121331301232322-2020330310112333-2101330302030323-1133023322201232-3111311130220111-2201221133001232-3310100130211211"></a>

## policy_based_challenge.always_enable_js_challenge — always_enable_js_challenge / 210220202101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-3323022012313030-2233123221131010-3103212301010220-2311320330021102-3032213230100200-1300311310211001-0201222100022212-1312313303203102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for always enable js challenge.

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

<a id="canonical-0321121131220323-2002120312121203-3330301111220100-2230312012120000-3320030111332230-2201112232213012-0112232130330122-3110210332123203"></a>

## Direct properties — always_enable_js_challenge / 210220202101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201213001200120-3201202312000322-3012200001332220-2333120322211211-1202221003220013-3013210133132230-3120021223322021-0001001220211030"></a>

## Next pages — always_enable_js_challenge / 210220202101 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0031232123230311-0211102100302300-0122202222231323-1011002121100331-1103000211323133-0100003031322133-3132130021030323-0102120103131112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323030110003103-1132200231132213-2230302021230333-3310130221213203-3032010211100003-0221021002332211-2223102223112212-1111302322030003"></a>

## policy_based_challenge.captcha_challenge_parameters — captcha_challenge_parameters / 302130311132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-0132303121313313-3220013101030022-2202311122010311-3232101133202333-2223200233111313-2131200032013320-0012210131121122-0020123333020030"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

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

<a id="canonical-3010221100331301-3032023022032332-0001323331230133-3310231221221030-3331011111111130-2232320203102022-0320010311200321-2101330332010123"></a>

## Direct properties — captcha_challenge_parameters / 302130311132 / 3

<a id="canonical-3331203012200121-0120010122330033-1202213213203201-1300222033103301-3333233112100233-0123313301103302-1213311210312002-2110332011331003"></a>

<a id="canonical-2310000032233300-0311123030013213-1210000230300032-1001310230232101-3132102330131123-2111300231210112-1010212102332331-2032033133203213"></a>

## cookie_expiry property — captcha_challenge_parameters / 302130311132 / 4

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

<a id="canonical-0121122113120112-1310023211120300-3313002033321031-1210211310120221-2030302223201110-2201122323202202-1031221211221213-1330103120002300"></a>

<a id="canonical-2332220032213122-2301213001233132-2110012113123202-1011031212332330-2011100002330130-2110232201222223-3101231201210123-3232111002323320"></a>

## custom_page property — captcha_challenge_parameters / 302130311132 / 5

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

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

<a id="canonical-2220213223100322-3320002332023010-1012011210010032-1022132300010323-1030332233201021-3303101203311113-3031303123032010-3011310300110312"></a>

## Next pages — captcha_challenge_parameters / 302130311132 / 6

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3033122122132230-3012313221123333-2130013103000000-2130320102210233-2202310122202323-0320010320033221-3010302311012312-1221131320330202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113332202031213-1333130301331312-2231030000013002-3030300111000013-0021111311230020-0313132203201002-3233330013312002-0000100230033020"></a>

## policy_based_challenge.default_captcha_challenge_parameters — default_captcha_challenge_parameters / 232022333111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-3322101031112010-3003312211102330-3031110103022003-3303232011130333-1301313011313230-0221022133301212-0323333022322322-3213200032033013"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

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

<a id="canonical-2121013121200231-3110233223021301-2020320220301203-3212233311110321-3103003101330023-3123130000100203-0302021330012320-1202233120230323"></a>

## Direct properties — default_captcha_challenge_parameters / 232022333111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120123013010123-2003202222121023-0132012000021001-2132103121223100-1020203221013302-2212223133011213-3012220020332323-3010032210013001"></a>

## Next pages — default_captcha_challenge_parameters / 232022333111 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3210233223003011-1123023021123130-0201102132101200-0232230000032100-1303300132120223-2103303302313301-3111233202020030-1333221121021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130112022221212-0300100133222213-3303021230311210-2331300320220032-3333313003113300-2020220002033003-2202310302113021-1132133212212001"></a>

## policy_based_challenge.default_js_challenge_parameters — default_js_challenge_parameters / 232000333322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-2322330210031102-2200133111320321-2320012000121002-0200120130123211-0220321110202302-0213003230031330-0020320201232120-1010303313221201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default js challenge parameters.

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

<a id="canonical-3211322213101311-1021300002333323-1002013101300221-2130002321200310-0330310032121232-1110101321000211-0211013113330320-0330123202212133"></a>

## Direct properties — default_js_challenge_parameters / 232000333322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201010023122230-1133000132323123-0200323002332021-2102221003101102-3230330010233113-0021032233100032-2201210301302022-3010203202322312"></a>

## Next pages — default_js_challenge_parameters / 232000333322 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3001110002120201-1020301121231202-0010013333031201-0123200211120221-0211033021233231-0012210122330320-3220310110222103-2012203210302011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113133303100020-2111001122002121-3121311002323023-3233100130110112-3320300110203123-3310000113132100-2333310002023132-0130120333013121"></a>

## policy_based_challenge.default_mitigation_settings — default_mitigation_settings / 110201123233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-0220230311302122-1011113122131210-0222011121310121-3331233102121210-0322323201001323-0200133133213020-0131223021001230-3130121313330233"></a>

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

<a id="canonical-2113010112221102-3311301201211221-2230322022222021-3213130111033222-3203203231013000-2201212031322012-2103013002101220-0131212103200121"></a>

## Direct properties — default_mitigation_settings / 110201123233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310011231103103-3230213011312111-3221111011100302-1002313002011102-0323123212212023-2000133203322303-3211131121020323-3031222133311332"></a>

## Next pages — default_mitigation_settings / 110201123233 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2313131231020200-2210311011111232-0132022020301311-2203023202030330-1210133110122201-3131320010032300-0101231220213330-1321102212211301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313212010330021-0111131012122212-0031222130113322-1002301330002111-0132010110032223-3202213212003030-1223010322232330-1231130320033232"></a>

## policy_based_challenge.default_temporary_blocking_parameters — default_temporary_blocking_parameters / 323103110233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-0200121023310230-1111301200331020-3210230111033103-0330102303213221-3233123210101111-1330011310032202-1131330201332023-2010232133033311"></a>

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

<a id="canonical-3203102000202222-1302010212301331-0300110201131123-0121310032203200-2002202012313331-3331301303101022-0320232333020311-1012211302211201"></a>

## Direct properties — default_temporary_blocking_parameters / 323103110233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232320103022113-1022011011002113-2013230001320023-1201112002120000-0020011200333011-2320032320130000-3122201232301020-3222331331330221"></a>

## Next pages — default_temporary_blocking_parameters / 323103110233 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1312032111102320-0120213331303133-3332231110231301-3123320201021121-0100212113022232-2312321322331120-1131130322332031-0221220001021103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301303203313331-2132111103102132-1012123313131101-2331131223000320-1221113201133002-2121110300131322-3113100313220201-2030110311210332"></a>

## policy_based_challenge.js_challenge_parameters — js_challenge_parameters / 031223312330 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-3203002322011330-3113332212120003-3222001200220312-0023331112003013-3222203310231300-3101021132311100-2120031023013102-0023222003122331"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

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

<a id="canonical-1302333213221022-3212233321133023-3013002123103131-3320333310010111-0030120022111330-3230122310331120-2022123233220301-2322112232231213"></a>

## Direct properties — js_challenge_parameters / 031223312330 / 3

<a id="canonical-0003133331312000-0011121031123221-0300213132113213-3303132102233313-1003302032233220-0222233301123102-1100123320313320-1320332312310201"></a>

<a id="canonical-3011132210331333-3211132131303021-3300310120233212-1001202122113231-3201001133011112-3030011131122001-3112203201132023-3321123230223122"></a>

## cookie_expiry property — js_challenge_parameters / 031223312330 / 4

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

<a id="canonical-1203231033310133-2331231220130123-3310212231013323-2030122130212002-2311120313300000-2122002323123320-0000333103231300-1231333102013321"></a>

<a id="canonical-0003203011013321-2313231232220113-2000030130320231-0330002301201310-2311023002312101-2311013230022331-2003200111233232-2231012233031230"></a>

## custom_page property — js_challenge_parameters / 031223312330 / 5

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

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

<a id="canonical-3210002030013222-3132222333023121-3311233312321021-0203132021330121-0213221101111323-3222033001100232-0212232023133312-0321010202120201"></a>

<a id="canonical-3120101002001020-2221113221111021-2032010123201122-2013110231201021-1221302310230333-2121030203010013-2100031301011013-2303210011222302"></a>

## js_script_delay property — js_challenge_parameters / 031223312330 / 6

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

<a id="canonical-2220200123002211-2310332131112103-0331311333023233-3033310110033111-0303031230131232-3122121132203202-0022212030111232-0022000331221123"></a>

## Next pages — js_challenge_parameters / 031223312330 / 7

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0301011213301031-1302102201002310-1003103333312111-2320002321312333-1213202200312033-3203123032001123-1313300002102203-0212131202331300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122332230331111-1311302011111131-0121223330301011-2213120100222133-2132020113213201-3302203012121211-3020330021200220-3133133221331120"></a>

## policy_based_challenge.malicious_user_mitigation — malicious_user_mitigation / 311233112021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-3302122032032323-0332311103021220-0132023113133220-3003130332001212-3111220003211231-1110003110321010-0003221332013300-1001111321300033"></a>

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

<a id="canonical-0020212133010222-1203012132113013-2123231121303312-3020130032203302-1303311130310222-0233103130012032-2312000132222023-2223013221321101"></a>

## Direct properties — malicious_user_mitigation / 311233112021 / 3

<a id="canonical-3330201221322233-2133222010300220-1100200311221330-0313222022323022-1201310113103303-2200313211221012-3231222222203011-1220213100113032"></a>

<a id="canonical-3311200300333221-3123332220000310-0210130032311012-2303111032021331-3102230333221301-2212313332103013-3113021100023003-0332322113131011"></a>

## name property — malicious_user_mitigation / 311233112021 / 4

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

<a id="canonical-3332302130211120-1121101230201213-0320201321001023-0211121111021023-3231233123303320-2331021100212230-3323112312033131-0130222313003100"></a>

<a id="canonical-1210312322330333-0002223300130030-0132120323022110-3202122002303102-0110311002312112-0223000201330112-1103311333211232-0022030102222300"></a>

## namespace property — malicious_user_mitigation / 311233112021 / 5

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

<a id="canonical-0002030331333323-2000210011222100-2120332221021100-3032121320203300-2303322220203310-3312021112001220-0312121230132122-2023102032132110"></a>

<a id="canonical-3322032123102013-3210233002120221-0200203033322001-2123323021131111-2132120101131121-1132323213323023-1011032132302231-3032213132011103"></a>

## tenant property — malicious_user_mitigation / 311233112021 / 6

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

<a id="canonical-1122320110233302-3221132113303012-2212213331000300-3322321021202301-0103000212321323-0101223133300030-3000021232000101-3101200032313100"></a>

## Next pages — malicious_user_mitigation / 311233112021 / 7

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1120331230101123-3020020300310231-3110322023000131-2130202113002001-2202123002003002-0100220202023003-2230213230101000-2323120133200032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123110233312020-1022233011233002-0021300001023011-0202010030310301-1211312311121003-1223112133332013-3233101222232110-3032311323210000"></a>

## policy_based_challenge.no_challenge — no_challenge / 112120222111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.no_challenge

<a id="canonical-2320130130323121-1010221101023232-0101312303131301-3231300231100303-0112001120310333-1120122121200301-3300103322312213-0200222212323132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-3102120200133112-3300031121033321-0113110333201320-3203131300101102-0223011023321000-3010320010331030-0112213311233211-1320311002230231"></a>

## Direct properties — no_challenge / 112120222111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021223333003210-2213102301002032-1211223333013302-0331000111013322-0130031121132010-1301313322002021-1023331013103213-1221201213323212"></a>

## Next pages — no_challenge / 112120222111 / 4

- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102223020023111-3023202121320022-1003213121120210-1331011330330202-3002203132120110-1302031321231103-2221121232022010-3331323211111013"></a>

## policy_based_challenge.rule_list — rule_list / 002101111302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.rule_list

<a id="canonical-0112033020000303-3130130031301111-1031310202013211-1310223133303330-2212102323023212-1220110221030312-0132121103133023-2131123033033021"></a>

Type: `"single"`. Computed.

List of challenge rules to be used in policy based challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0210330303023200-0003012021000203-2221133332230011-1003132323102000-1133202110033101-2031010030221203-0202203010130311-3022032032303302"></a>

## Direct properties — rule_list / 002101111302 / 3

- [rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003): complete subsection reference.

<a id="canonical-3130013320232131-1111233313023011-2111211020020132-0312230320301233-0000332031230003-3321131111033202-2130012211320331-2203010112201313"></a>

## Next pages — rule_list / 002101111302 / 4

- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102303221202220-1111220031122222-1223330230233210-2020101033330201-0212330203210331-2311223030130111-2203332212321103-0201010321231231"></a>

## policy_based_challenge.rule_list.rules — rules / 212323130002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-021.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- policy_based_challenge.rule_list.rules

<a id="canonical-0002011213112120-2203023133222132-2112033212121333-0313333330000000-1002333110220330-2302313333221300-2223121032312303-0222110022133313"></a>

Type: `"list"`. Computed.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
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

<a id="canonical-0303301231203231-3201211210301303-2102230313332332-0021301030222201-2032121222100012-2030322010201202-0320003333100202-2020133121321332"></a>

## Direct properties — rules / 212323130002 / 3

- [metadata](data-sources--http_loadbalancer--reference--group-021.md#canonical-2220202321033213-0130010312331133-2002322003111230-2100213323113320-3321130302213322-1013132322011023-2110203123300232-2230122032313120): complete subsection reference.

- [spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020): complete subsection reference.

<a id="canonical-1120130303003131-3212300002301020-0330210122230120-1231110011100131-1103232233101200-2313203303312110-2101003232003311-1112232113303102"></a>

## Next pages — rules / 212323130002 / 4

- [policy_based_challenge.rule_list.rules.metadata](data-sources--http_loadbalancer--reference--group-021.md#canonical-2220202321033213-0130010312331133-2002322003111230-2100213323113320-3321130302213322-1013132322011023-2110203123300232-2230122032313120)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2220202321033213-0130010312331133-2002322003111230-2100213323113320-3321130302213322-1013132322011023-2110203123300232-2230122032313120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
