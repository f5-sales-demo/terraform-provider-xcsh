---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-1322111001132313-2003202232112000-0122202110231321-3202123133031220-3020310100023013-2321232212000003-2323322113220030-1322012201211322"></a>

## Direct properties — protected_endpoints / 001333211023 / 3

- [any_domain](data-sources--protected_application--reference--group-003.md#canonical-0212131231100333-2110011000010121-2031311012203203-1223132000302322-0333333123301001-2130310113101232-0130301333301102-1232032121111231): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-003.md#canonical-1231021222223112-0213002101313033-1233123023100112-2301210113023100-2132323003033223-3310130002230021-3302312332023010-2300022030133103): complete subsection reference.

- [flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300): complete subsection reference.

<a id="canonical-3303203312203031-0120303221223312-0331030213320330-1330310122213111-0232213031112330-2102000122301330-1321101100210332-3000011322133232"></a>

<a id="canonical-1123212220230333-3032300030333220-0033221033132210-1132022213133000-0031030333102000-3020321031123223-1312233223111301-2223303332311322"></a>

## http_methods property — protected_endpoints / 001333211023 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](data-sources--protected_application--reference--group-003.md#canonical-2021230303322000-2030220332003021-3120230122221302-3030222030333130-2213211332121312-1222122333301110-3220301200123102-0201311323123022): complete subsection reference.

- [mobile_client](data-sources--protected_application--reference--group-003.md#canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322): complete subsection reference.

<a id="canonical-2122012101213113-2221201200303102-2122011020233202-0221332133220331-0210231203321001-2123103300221131-3123322013331210-0213111111132332"></a>

<a id="canonical-1132012021001313-0012113120020310-2033121033130210-3212101220333111-1313123202022300-3213111013032112-0021133213112110-0302331232313123"></a>

## path property — protected_endpoints / 001333211023 / 5

Type: `"string"`. Computed.

Accepts wildcards \* to match multiple characters or ? To match a single character.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  }
}
```

<a id="canonical-3032233123200112-1222002132032020-0331330222131300-1331113200212013-3133301233321120-0030323330310210-1112032002302112-0231003232201101"></a>

<a id="canonical-2101203020200212-2232221031210213-3133133301331302-3110233200212232-0101122203130203-2030111323320001-3033112033203331-0023130232100300"></a>

## query property — protected_endpoints / 001333211023 / 6

Type: `"string"`. Computed.

Enter a regular expression to match your query parameters of interest.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
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

- [undefined_flow_label](data-sources--protected_application--reference--group-004.md#canonical-2130323322302022-0212320201322103-2202001122001133-3110232001330021-2300113223103221-0321313111323130-3330323312330233-3013133120303122): complete subsection reference.

- [web_client](data-sources--protected_application--reference--group-004.md#canonical-0021131221101020-3323020220213013-3221300121212033-3022302311332020-0101211313203023-0133110222010302-0221020001111012-0011131003021220): complete subsection reference.

- [web_mobile_client](data-sources--protected_application--reference--group-004.md#canonical-1221222210233203-0121321330221033-3221101000112032-3233233303213100-2313001001111110-3313022302233212-3003311322313220-1200102012332120): complete subsection reference.

<a id="canonical-3311020103010132-2131230021130203-2333310213300231-3013332200121212-0113330320021030-2320113130302031-2030131322102002-1002300223321131"></a>

## Next pages — protected_endpoints / 001333211023 / 7

- [cloudfront.protected_endpoints.any_domain](data-sources--protected_application--reference--group-003.md#canonical-0212131231100333-2110011000010121-2031311012203203-1223132000302322-0333333123301001-2130310113101232-0130301333301102-1232032121111231)
- [cloudfront.protected_endpoints.domain](data-sources--protected_application--reference--group-003.md#canonical-1231021222223112-0213002101313033-1233123023100112-2301210113023100-2132323003033223-3310130002230021-3302312332023010-2300022030133103)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.metadata](data-sources--protected_application--reference--group-003.md#canonical-2021230303322000-2030220332003021-3120230122221302-3030222030333130-2213211332121312-1222122333301110-3220301200123102-0201311323123022)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322)
- [cloudfront.protected_endpoints.undefined_flow_label](data-sources--protected_application--reference--group-004.md#canonical-2130323322302022-0212320201322103-2202001122001133-3110232001330021-2300113223103221-0321313111323130-3330323312330233-3013133120303122)
- [cloudfront.protected_endpoints.web_client](data-sources--protected_application--reference--group-004.md#canonical-0021131221101020-3323020220213013-3221300121212033-3022302311332020-0101211313203023-0133110222010302-0221020001111012-0011131003021220)
- [cloudfront.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-004.md#canonical-1221222210233203-0121321330221033-3221101000112032-3233233303213100-2313001001111110-3313022302233212-3003311322313220-1200102012332120)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0212131231100333-2110011000010121-2031311012203203-1223132000302322-0333333123301001-2130310113101232-0130301333301102-1232032121111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101231220121321-3032030211021130-2212021112122211-1020301313313313-2212031331123233-2003300200333102-3022003311120323-1001022200002310"></a>

## cloudfront.protected_endpoints.any_domain — any_domain / 232200301332 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- cloudfront.protected_endpoints.any_domain

<a id="canonical-0302122110010130-2031103103101000-0302002022210020-0022230202103223-3313033133303101-1313013020203010-1202213023111123-1323003011003211"></a>

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

<a id="canonical-0121001112010230-0221033203010301-1023131223100333-2232122103213033-3301101001222203-0132221201023300-2223103221322113-1322222231313023"></a>

## Direct properties — any_domain / 232200301332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003022230310203-1221202031301121-2022101232210101-0300323302333203-3031110101132011-3322010203330301-0022331011001022-2120010230011300"></a>

## Next pages — any_domain / 232200301332 / 4

- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1231021222223112-0213002101313033-1233123023100112-2301210113023100-2132323003033223-3310130002230021-3302312332023010-2300022030133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101122133013123-2212130121123032-0222222202311100-0010321330031123-1313002122310332-0032322312112120-1112233031102003-2032231012112233"></a>

## cloudfront.protected_endpoints.domain — domain / 002020210310 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- cloudfront.protected_endpoints.domain

<a id="canonical-2302231201023310-2313302100000222-3123132201202211-0313113312232311-3333220322013203-0011200020212100-0210213201332220-0202230123332103"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-3233103113312130-3102012223200113-0223213223203220-1132113311032332-1311111311020013-0222313003320302-0123030120031121-1323313201130212"></a>

## Direct properties — domain / 002020210310 / 3

<a id="canonical-2232133333212102-1110113103100322-3223321213132000-2110230230210100-1113302133010222-1223303130133001-2132011213223011-3011122111003330"></a>

<a id="canonical-2203313231112102-2313123223322332-3300303030123232-0201132302123221-0332302211012223-0111012003300223-1010320131022132-2332310101023203"></a>

## exact_value property — domain / 002020210310 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-2102020000130020-1230311222231003-2103223202013111-2021311330313121-3033021101322130-0230023011311030-2111113312202332-2000203312323320"></a>

<a id="canonical-3321002201313132-0103203310232333-2003010122330120-0031211133022303-0332311122312100-3032100202331232-0100122330300121-3023022211210300"></a>

## regex_value property — domain / 002020210310 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3100331203013213-2111233123322013-1323121230320310-3130113203100222-3032203021313231-2030330201023023-3120231113012002-0213033130221103"></a>

<a id="canonical-0100110111102010-1001021233132233-1323010030223310-0103103010011030-3023213320233133-3201021333122213-3033201331033012-0310331022303000"></a>

## suffix_value property — domain / 002020210310 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-2010231321231122-0021112211123012-2213030211100231-1010210022332033-3330232222000212-0322123002100232-2223201211323030-1301111033301231"></a>

## Next pages — domain / 002020210310 / 7

- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301233230313331-3230013102322213-0210011231220212-0220013203000230-3032230302201332-2000121221030220-3022023302130013-3303010221020013"></a>

## cloudfront.protected_endpoints.flow_label — flow_label / 013331013300 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- cloudfront.protected_endpoints.flow_label

<a id="canonical-2020033321111033-2202222133123020-2131133212213323-0222100102031133-2033032012121202-2313202322130212-3310021331311022-3013222210000203"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

<a id="canonical-2112023021330023-0003131002333331-0330323222023210-1032222232102102-1330223100322023-3211001220230301-0002110003212113-3312221102132211"></a>

## Direct properties — flow_label / 013331013300 / 3

- [account_management](data-sources--protected_application--reference--group-003.md#canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031): complete subsection reference.

- [authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201): complete subsection reference.

- [financial_services](data-sources--protected_application--reference--group-003.md#canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203): complete subsection reference.

- [flight](data-sources--protected_application--reference--group-003.md#canonical-3022031310300013-1303031212122003-1101220112133123-2031112213031201-2311331321103330-2230221333213230-0213031113013220-0222011332013103): complete subsection reference.

- [profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303): complete subsection reference.

- [search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003): complete subsection reference.

- [shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321): complete subsection reference.

<a id="canonical-1330331023222332-0132102321130003-2113332220322102-2233213000000221-0233033213033230-0001333213112123-1321233012000110-1320122001013030"></a>

## Next pages — flow_label / 013331013300 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203)
- [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--reference--group-003.md#canonical-3022031310300013-1303031212122003-1101220112133123-2031112213031201-2311331321103330-2230221333213230-0213031113013220-0222011332013103)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232002020310010-3133322030212111-3102123100323020-1200321301320211-0122230311210001-2003120120322221-3331201222230213-2123102311202120"></a>

## cloudfront.protected_endpoints.flow_label.account_management — account_management / 011311331101 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- cloudfront.protected_endpoints.flow_label.account_management

<a id="canonical-3200011001300121-1012011300020310-1220233013302021-3201222303301312-3211000322133331-3201312001031220-1222300101012023-2220323221201110"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Account Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

<a id="canonical-0123201112223200-1330100211121013-1023100311213230-0302012213333211-1110220033303033-0123302330321230-3100031000203013-2131321012232101"></a>

## Direct properties — account_management / 011311331101 / 3

- [create](data-sources--protected_application--reference--group-003.md#canonical-3210110300012033-1011300211230312-1131311032211333-0002200301302201-1300022300122112-1103112130210202-2032110033322133-3331233030023101): complete subsection reference.

- [password_reset](data-sources--protected_application--reference--group-003.md#canonical-0101310030221112-2222321012133122-1200303032322302-2212230000301113-0312010320102123-1130320133202333-2120310213213133-0122213121333010): complete subsection reference.

<a id="canonical-0010003133231112-2011211132232231-1322031131033213-2213213320233023-0202120002121332-0233313001132212-3110100000030202-3122112233320300"></a>

## Next pages — account_management / 011311331101 / 4

- [cloudfront.protected_endpoints.flow_label.account_management.create](data-sources--protected_application--reference--group-003.md#canonical-3210110300012033-1011300211230312-1131311032211333-0002200301302201-1300022300122112-1103112130210202-2032110033322133-3331233030023101)
- [cloudfront.protected_endpoints.flow_label.account_management.password_reset](data-sources--protected_application--reference--group-003.md#canonical-0101310030221112-2222321012133122-1200303032322302-2212230000301113-0312010320102123-1130320133202333-2120310213213133-0122213121333010)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3210110300012033-1011300211230312-1131311032211333-0002200301302201-1300022300122112-1103112130210202-2032110033322133-3331233030023101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003011310221223-1310322231221000-3312331211103011-3333111000210213-1113223103112303-3211213030203101-1320130330203210-3002230233303011"></a>

## cloudfront.protected_endpoints.flow_label.account_management.create — create / 022310330112 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031)
- cloudfront.protected_endpoints.flow_label.account_management.create

<a id="canonical-1201211312121211-1211310230220101-2220221303212022-2331111113122211-1220302231220100-0023210303121113-2212103013222033-3202101121310120"></a>

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

<a id="canonical-1301203331112232-1033013121222331-2020233023330003-3113301213213013-2131210330320312-2221123112001202-3320012312322030-1223111221112021"></a>

## Direct properties — create / 022310330112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120233311103230-0131302121032330-1130111333203102-1011103221013311-1000232032012321-2013233013231213-3120000310103333-2201033010121113"></a>

## Next pages — create / 022310330112 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0101310030221112-2222321012133122-1200303032322302-2212230000301113-0312010320102123-1130320133202333-2120310213213133-0122213121333010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223301233212030-0333323231001201-2100023230311030-0013033330200132-3020321302132103-3201231210223223-3030101330323001-1103332101323310"></a>

## cloudfront.protected_endpoints.flow_label.account_management.password_reset — password_reset / 312002103013 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031)
- cloudfront.protected_endpoints.flow_label.account_management.password_reset

<a id="canonical-3021002030312000-1311121111111201-2210211023013021-3133033102312302-3113131211312223-0323023003100120-2202320032213223-1000020231130232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for password reset.

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

<a id="canonical-0213212302312313-2232201000213130-3222301303101111-0222101212322330-0232231002021102-0322221203112021-1121021213021112-0023101000101223"></a>

## Direct properties — password_reset / 312002103013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013221030113312-2300230333131202-3211000303102320-1021032100011100-1022311012030113-3123100321013301-3332110221202331-2122023100121212"></a>

## Next pages — password_reset / 312002103013 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-1201000230102002-3032303001221220-1011311033111201-0122121320212232-1310132100233311-0231201102223133-3031233021232123-3001002122032031)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001111013103100-2232221332003303-1331232323222103-0302132200000013-0313300313122033-2331203320313022-1213131313320021-2101210322201131"></a>

## cloudfront.protected_endpoints.flow_label.authentication — authentication / 333110200122 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- cloudfront.protected_endpoints.flow_label.authentication

<a id="canonical-0133021313100221-0203311121003231-0123310101323110-3033333332123233-2021031323112332-0132232031002103-2103131331030100-0020021033001330"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Authentication Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

<a id="canonical-1231201023233023-0311221101123021-3022310031213012-0021130003301121-1002100320322010-1200022120321300-2021312211233312-0112220131023121"></a>

## Direct properties — authentication / 333110200122 / 3

- [login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222): complete subsection reference.

- [login_mfa](data-sources--protected_application--reference--group-003.md#canonical-0220321331123103-1002301300010023-2230032002223001-2221132032223202-2023100032222223-2301033221033123-3130101003320130-1003201201213100): complete subsection reference.

- [login_partner](data-sources--protected_application--reference--group-003.md#canonical-3031021001023310-1001303023303121-1333220222000312-3200311231032211-3023113112132201-3310303303233002-2002033331110212-1331321123030101): complete subsection reference.

- [logout](data-sources--protected_application--reference--group-003.md#canonical-1012212123203032-3332021312221320-2321023222231133-0101332313111032-2310012102321122-0310023322332100-1303310131212322-0131312113132230): complete subsection reference.

- [token_refresh](data-sources--protected_application--reference--group-003.md#canonical-2222011203231312-1332202010300230-2310220202122201-3231210200021000-1011313301220101-3100111223102101-0310312111022023-1330011320301212): complete subsection reference.

<a id="canonical-0030200033332121-2223033132331020-1312131113012033-0331122233312331-1220001001100032-2133031323302133-0322130220322211-0213212233313111"></a>

## Next pages — authentication / 333110200122 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](data-sources--protected_application--reference--group-003.md#canonical-0220321331123103-1002301300010023-2230032002223001-2221132032223202-2023100032222223-2301033221033123-3130101003320130-1003201201213100)
- [cloudfront.protected_endpoints.flow_label.authentication.login_partner](data-sources--protected_application--reference--group-003.md#canonical-3031021001023310-1001303023303121-1333220222000312-3200311231032211-3023113112132201-3310303303233002-2002033331110212-1331321123030101)
- [cloudfront.protected_endpoints.flow_label.authentication.logout](data-sources--protected_application--reference--group-003.md#canonical-1012212123203032-3332021312221320-2321023222231133-0101332313111032-2310012102321122-0310023322332100-1303310131212322-0131312113132230)
- [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](data-sources--protected_application--reference--group-003.md#canonical-2222011203231312-1332202010300230-2310220202122201-3231210200021000-1011313301220101-3100111223102101-0310312111022023-1330011320301212)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302003202331120-1222200302103023-1122221323221203-2101203313111211-0231322301321100-3323303212200331-3332022203233303-2120110323330000"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login — login / 222003022130 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- cloudfront.protected_endpoints.flow_label.authentication.login

<a id="canonical-1012103133031001-2312020033233231-0020320003313002-3020021322310220-0202223030120333-0201212223222101-0212211012133222-2022113330310122"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

<a id="canonical-3030221310132230-2031031333212133-1020023011222311-3223002223203020-3121230032120221-2221310203000022-2130323130003121-3123130310023120"></a>

## Direct properties — login / 222003022130 / 3

- [disable_transaction_result](data-sources--protected_application--reference--group-003.md#canonical-3122233301311321-0301320121032303-1300203222021111-1232303023101211-2132111311312200-3220121032232103-2221301133323323-0103022220203210): complete subsection reference.

- [transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022): complete subsection reference.

<a id="canonical-2302311132232210-2020130132022202-3122231202302100-2110212012003031-0330220011130221-2331320202022203-0020120123222031-2200130023112001"></a>

## Next pages — login / 222003022130 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](data-sources--protected_application--reference--group-003.md#canonical-3122233301311321-0301320121032303-1300203222021111-1232303023101211-2132111311312200-3220121032232103-2221301133323323-0103022220203210)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3122233301311321-0301320121032303-1300203222021111-1232303023101211-2132111311312200-3220121032232103-2221301133323323-0103022220203210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120230231323211-0200011012310030-1220110201130310-0331123231000303-2121310012023203-1223330331223300-1103000011032110-2020101320100123"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result — disable_transaction_result / 203120221022 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-0122200211003002-0130212322323301-1102122331111333-1022123311312002-3333100312023032-2121102301313122-2110203331120233-0200020230232231"></a>

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

<a id="canonical-3211302010310122-2210213331130001-2323302233232310-3012321300120332-3120312021201023-1000201213100312-1211312111020132-3230030302130200"></a>

## Direct properties — disable_transaction_result / 203120221022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130212300332311-2012201212002130-0220330232321203-3123103122023033-0010133320202222-2003310322000133-3032322212200322-2011332121112300"></a>

## Next pages — disable_transaction_result / 203120221022 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220000200220011-1302021300303333-0121021200301102-1320213220023333-3321133111111123-0130320213011200-1102011312001220-1031110023302201"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result — transaction_result / 333031033331 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-2212111022132332-2202111122203111-3101231323033320-1020002223321303-3111111021122133-3303310301331013-3013112101123012-2310000021322132"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2030210023023113-2213221012022210-0210000120322303-1210330300122132-0203030023031002-0323012332002023-1001201023302010-2221002033313232"></a>

## Direct properties — transaction_result / 333031033331 / 3

- [failure_conditions](data-sources--protected_application--reference--group-003.md#canonical-3130100221131213-2230321323333020-0320311132123322-2131333203032003-2303313031022313-1231112323232023-0030121210320003-2232000102332203): complete subsection reference.

- [success_conditions](data-sources--protected_application--reference--group-003.md#canonical-3230002002200023-1301221332312103-0112220022121110-1112223210312001-3220313313233103-3132033120232220-0203320100103002-3122321212032121): complete subsection reference.

<a id="canonical-0001012100233012-3100133310122122-2032003010103333-3121301003002003-1200111133232030-2201230302012221-0112331013101222-2103113232222311"></a>

## Next pages — transaction_result / 333031033331 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](data-sources--protected_application--reference--group-003.md#canonical-3130100221131213-2230321323333020-0320311132123322-2131333203032003-2303313031022313-1231112323232023-0030121210320003-2232000102332203)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](data-sources--protected_application--reference--group-003.md#canonical-3230002002200023-1301221332312103-0112220022121110-1112223210312001-3220313313233103-3132033120232220-0203320100103002-3122321212032121)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3130100221131213-2230321323333020-0320311132123322-2131333203032003-2303313031022313-1231112323232023-0030121210320003-2232000102332203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020310003001032-1311220301110213-1022100131133101-1133230220001310-3010323311003330-1133313222122122-3220122200300002-3121120031233032"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions — failure_conditions / 301110312033 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-3310220233022301-2120232333033111-2223213113122231-0133320130220230-3333333122031123-1222300203213213-2022013001223211-0213100330200111"></a>

Type: `"list"`. Computed.

Failure Conditions. Failure Conditions.

Upstream description:

Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1032233102312003-1022213222103321-1130231103001333-2331312330200213-0021130111000301-2122132100113001-2221003300122321-1102230031002020"></a>

## Direct properties — failure_conditions / 301110312033 / 3

<a id="canonical-0320230031002303-2222132230212120-2320030213131023-1320022222222023-2301202223300311-2220210231110221-3010023030120322-3112112210131202"></a>

<a id="canonical-0110123222110200-2213221101032331-1322013211321222-2223112301231310-0310303202330121-3331323311011220-2310231301201332-0110112002231012"></a>

## name property — failure_conditions / 301110312033 / 4

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1203203012103232-1030222322300120-3200233103020223-2120203231233323-3123202221003323-0232312200130221-2210201122230110-0313313001100030"></a>

<a id="canonical-0020230213133131-0210033021213322-3123232222121120-3002223311332302-0130221031230301-0033212100133131-0000313130211113-3330122320303001"></a>

## regex_values property — failure_conditions / 301110312033 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2223312011002310-1002321110003230-3313303301231320-1003102202313112-3313130310322121-3303223313303303-0331302230233012-2311011331222120"></a>

<a id="canonical-0300101033212001-3303100101232020-3031221211322112-2102131120212011-2003232133121331-2003120303100200-1031020200232022-2030202321310121"></a>

## status property — failure_conditions / 301110312033 / 6

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311101203002012-3221130010222203-0203302221102210-0221321011321103-1112303021212111-2330030120101120-0011012011113212-1331213031302232"></a>

## Next pages — failure_conditions / 301110312033 / 7

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3230002002200023-1301221332312103-0112220022121110-1112223210312001-3220313313233103-3132033120232220-0203320100103002-3122321212032121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133001331310312-2200223221010213-1232323011203110-1323320022121122-3011303222222330-3230111211300232-2332020011030211-3123033223001020"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions — success_conditions / 201221101021 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-3002133111222112-1231032002112122-3211131323230030-3013222033101021-0010013323132210-0231021130112001-2123321110123031-3323211011003222)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-3012221033231200-0121302301101011-1033122332123023-0011011020011121-0030030132332003-0330201132202312-2002330030231131-2313312232120302"></a>

Type: `"list"`. Computed.

Success Conditions. Success Conditions.

Upstream description:

Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1033310322213000-1121112300111212-3201230103223113-0002302332003032-2000131221330013-3123333133012311-2003033312103200-3313030320001213"></a>

## Direct properties — success_conditions / 201221101021 / 3

<a id="canonical-3003122003321110-2031233012202203-3021001023022032-0300220330000103-2022201123303110-1221232102123010-2331213333021013-3201122013303311"></a>

<a id="canonical-3303111303300310-0030032202220132-1001110122123210-1111311021032232-2101210322111032-0202003000232020-0333223321301010-1001303103023310"></a>

## name property — success_conditions / 201221101021 / 4

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0011213333313120-2230221211101311-0130111211213103-3331022332001203-2322021122111301-2131221013112032-1211210223132221-2321033333331320"></a>

<a id="canonical-1031313012110330-1321230220020103-0231001012311111-0210102032110223-1230230332203032-1023303130331220-1300230332011130-0030333123333332"></a>

## regex_values property — success_conditions / 201221101021 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2003212000301303-2113020321331212-1211311320301102-2332303030121022-2001302112233300-2121020303201000-2201033012102021-0323133011332232"></a>

<a id="canonical-1011102033120221-0322133022310302-1100013303132121-2222300000331011-2023033020233103-3301010213313301-3010323203300010-2112031310201022"></a>

## status property — success_conditions / 201221101021 / 6

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022023120000331-1112010023231002-1333322220013303-1121222122303223-3231332233110211-0332132301003000-3112102100221311-2031323210320220"></a>

## Next pages — success_conditions / 201221101021 / 7

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0013020233033310-3302013320031033-1303312000233302-1213302001321311-3221032102001213-3222212322133223-0002103132310330-0232001010120022)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0220321331123103-1002301300010023-2230032002223001-2221132032223202-2023100032222223-2301033221033123-3130101003320130-1003201201213100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002331121233312-1032210103300133-3033000332321121-2122312122202113-2012330210101221-1220013013031133-3001110130231233-3002311032120010"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login_mfa — login_mfa / 222232313000 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- cloudfront.protected_endpoints.flow_label.authentication.login_mfa

<a id="canonical-3220033113120103-2101233210303221-1002232301123312-0321023123232323-0121012110130333-0223130032323223-2230023112202232-3222033332211331"></a>

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

<a id="canonical-1120000221222112-2213020310201230-3133332320112012-1001201320013122-0203203301310211-2221212102132323-1001131133020120-0120221120222001"></a>

## Direct properties — login_mfa / 222232313000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012231302213111-3321213330210232-0000200122203030-3310013201112300-3123113201112303-1111312213300023-0101201210022001-1201033021321002"></a>

## Next pages — login_mfa / 222232313000 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3031021001023310-1001303023303121-1333220222000312-3200311231032211-3023113112132201-3310303303233002-2002033331110212-1331321123030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121023021313213-1331110232312121-1001202320013011-1132021023301133-1120010013203201-2010203023113033-0012020102121212-1001120013200102"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login_partner — login_partner / 223321123303 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- cloudfront.protected_endpoints.flow_label.authentication.login_partner

<a id="canonical-0302023012233013-3221000023200010-2031010330020232-1002320220210023-1003200130132023-1131132111021120-2113330032021110-0320331232321001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for login partner.

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

<a id="canonical-1113001101221320-1111113110210300-0120333103121320-0310100021202132-0032113123000133-0102113003103331-3122301003102101-3031330211323021"></a>

## Direct properties — login_partner / 223321123303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220112101010003-0121303113220120-3113331103322213-1232132021301322-3331121232213323-1313000320003022-0000002012230110-3011202030303330"></a>

## Next pages — login_partner / 223321123303 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1012212123203032-3332021312221320-2321023222231133-0101332313111032-2310012102321122-0310023322332100-1303310131212322-0131312113132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213333133313030-2203310332303001-3211303023102020-0020201223001102-3310101331321113-3222122300013213-3311021121112001-3312222330030000"></a>

## cloudfront.protected_endpoints.flow_label.authentication.logout — logout / 223331112032 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- cloudfront.protected_endpoints.flow_label.authentication.logout

<a id="canonical-1002313002311101-0331131111301323-0133232313003312-2222310222101123-0023101303203211-0133103013010102-1330011002323230-1102111211303301"></a>

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

<a id="canonical-2103122113010013-2030133022332221-0231303331222201-1221131232213303-2111110001233233-1230312122311331-0112032313021022-2313220003100203"></a>

## Direct properties — logout / 223331112032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022312133333322-1120120033122231-3010202303310112-3322003103300312-0221330301203202-1000230030312003-3320221032031102-2021310010201001"></a>

## Next pages — logout / 223331112032 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2222011203231312-1332202010300230-2310220202122201-3231210200021000-1011313301220101-3100111223102101-0310312111022023-1330011320301212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322111210212121-3300203032230113-0333201321101103-1203211021223211-3330300112111213-1323230113202322-0130322222031030-0103211302103231"></a>

## cloudfront.protected_endpoints.flow_label.authentication.token_refresh — token_refresh / 233000311131 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- cloudfront.protected_endpoints.flow_label.authentication.token_refresh

<a id="canonical-2301231112210332-3333321122233312-0221330323331200-3233330003013000-2132333300331222-3320331231013100-2121203323120331-0220220300031002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for token refresh.

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

<a id="canonical-1331102001122301-0033211233222311-3011313030110132-3332222100132331-1120201301321101-0120020332330200-2130032132120313-0323211102131210"></a>

## Direct properties — token_refresh / 233000311131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023012022202312-1011133031302201-2320031131312211-2101002113130211-0102233223200311-3121330100220231-1223031033311100-0120202101030120"></a>

## Next pages — token_refresh / 233000311131 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-1321102323021333-1211031102220000-2012323233002232-1321321030032120-3132201323302231-3110213320300012-2313030000030300-2311133222212201)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332021230101011-2333333103231200-2222120130322122-0301031332130310-0032312200132030-2221122132111031-3031110000303121-0332003133203010"></a>

## cloudfront.protected_endpoints.flow_label.financial_services — financial_services / 330031112213 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- cloudfront.protected_endpoints.flow_label.financial_services

<a id="canonical-2222023211012313-1012110111322331-1201302320020311-1020001302312203-0230132020032331-0231021211110203-2133200311112303-1100233031130313"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Financial Services Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

<a id="canonical-3231020022223030-3133002123113101-2233323030110000-2210002002300213-2030300000321330-3103303012333203-1031323033122303-2121033203211021"></a>

## Direct properties — financial_services / 330031112213 / 3

- [apply](data-sources--protected_application--reference--group-003.md#canonical-1332030300300100-2001130100010303-1321300101122113-0311303112010101-2232331302131321-1300211202320132-1101002131201122-0002232320123022): complete subsection reference.

- [money_transfer](data-sources--protected_application--reference--group-003.md#canonical-2232301131101210-0223321230301213-0112322230101032-2223033302333313-3221033230200213-0323102231333133-0301132200201000-0310332112300230): complete subsection reference.

<a id="canonical-2321230021312001-0223323213021020-1100213010330300-0112013212133020-0320303202303023-0202310233120001-2220012230221103-1212311111001133"></a>

## Next pages — financial_services / 330031112213 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services.apply](data-sources--protected_application--reference--group-003.md#canonical-1332030300300100-2001130100010303-1321300101122113-0311303112010101-2232331302131321-1300211202320132-1101002131201122-0002232320123022)
- [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](data-sources--protected_application--reference--group-003.md#canonical-2232301131101210-0223321230301213-0112322230101032-2223033302333313-3221033230200213-0323102231333133-0301132200201000-0310332112300230)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1332030300300100-2001130100010303-1321300101122113-0311303112010101-2232331302131321-1300211202320132-1101002131201122-0002232320123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310333120033330-3333100112222320-1101101102222100-1100020022231010-3200122013103301-2320101012320320-0231212133011331-2302320032130121"></a>

## cloudfront.protected_endpoints.flow_label.financial_services.apply — apply / 322101130300 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203)
- cloudfront.protected_endpoints.flow_label.financial_services.apply

<a id="canonical-1132102021302213-0122321030023122-1330021103020030-3302210322132330-0333200312300333-0012302301021030-3303312321330010-1123320001302102"></a>

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

<a id="canonical-3003212100013303-3001311012311020-1223321202130123-2200023222001212-3220311200200012-2222211100112213-1322301330010331-1001020021012120"></a>

## Direct properties — apply / 322101130300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022033100010122-0102313032300000-3210223030211333-3232020130113232-3120011222331011-2330000031120302-1200023001130123-2333133223332211"></a>

## Next pages — apply / 322101130300 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2232301131101210-0223321230301213-0112322230101032-2223033302333313-3221033230200213-0323102231333133-0301132200201000-0310332112300230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121102000112111-0002212210211323-2230211103013030-3301130232322112-1110021201110220-1231323330302123-3101033130000133-2113233131032230"></a>

## cloudfront.protected_endpoints.flow_label.financial_services.money_transfer — money_transfer / 113120123222 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203)
- cloudfront.protected_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-3302201311301221-1330002120222203-1313200313230013-3030023120220001-0002203001221011-0100100021200030-1211100222111002-2111232321310021"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for money transfer.

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

<a id="canonical-2000310312120001-3023212332123302-2010020011233230-1031210300302202-2220012123231212-1101201031121111-3120032320210233-1133201011211322"></a>

## Direct properties — money_transfer / 113120123222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211131020220101-0100301213320102-1321300112011033-3330231130311112-0212223320220101-1303110213023121-0002013233023110-0132101031220332"></a>

## Next pages — money_transfer / 113120123222 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-0011220233111020-1330012103322331-2202112020122103-0303112231301000-2010112002103002-1332331220303123-0303301320202331-2030123332030203)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3022031310300013-1303031212122003-1101220112133123-2031112213031201-2311331321103330-2230221333213230-0213031113013220-0222011332013103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220123220210031-1102331332213103-2323310222203330-3022103302132211-2303131010001013-0032100133120310-0303203002012323-1031302131212121"></a>

## cloudfront.protected_endpoints.flow_label.flight — flight / 211000012130 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- cloudfront.protected_endpoints.flow_label.flight

<a id="canonical-1232122302233330-0213332221113131-2320101130100000-1112203011031213-1222121212102130-0330113001231220-3132010210130132-1101310120300233"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Upstream description:

Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

<a id="canonical-1230313310320121-2011011202211030-3230200210013300-3232320210331233-0333032310201203-1131331020301021-0322223223312031-2113310312102002"></a>

## Direct properties — flight / 211000012130 / 3

- [checkin](data-sources--protected_application--reference--group-003.md#canonical-3023330122313330-2320320312330330-3312003310123233-2021330131000111-2302303131312313-3222131232323001-0220011213122221-0333101033212013): complete subsection reference.

<a id="canonical-2303303010000200-2311120312313323-2131110301310130-2303202101222123-1223220113232023-2223331113312023-3312110133310332-1301221003322232"></a>

## Next pages — flight / 211000012130 / 4

- [cloudfront.protected_endpoints.flow_label.flight.checkin](data-sources--protected_application--reference--group-003.md#canonical-3023330122313330-2320320312330330-3312003310123233-2021330131000111-2302303131312313-3222131232323001-0220011213122221-0333101033212013)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3023330122313330-2320320312330330-3312003310123233-2021330131000111-2302303131312313-3222131232323001-0220011213122221-0333101033212013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122021100000011-3033021112020010-0210011330310012-0103100232002113-1223332120300310-1000103102130220-3131203301013011-0303212310032001"></a>

## cloudfront.protected_endpoints.flow_label.flight.checkin — checkin / 112021200100 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--reference--group-003.md#canonical-3022031310300013-1303031212122003-1101220112133123-2031112213031201-2311331321103330-2230221333213230-0213031113013220-0222011332013103)
- cloudfront.protected_endpoints.flow_label.flight.checkin

<a id="canonical-3202003000010110-1333100201101313-0201201131310222-0033123002200113-2223002211020331-2010103312103302-0111312221103301-3102220011320322"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3002033210213001-1332232202223213-3010323333001120-3200132011202103-2122332120011233-1030033012010103-0122201031323112-3312230302302231"></a>

## Direct properties — checkin / 112021200100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201130220212101-3031022132011221-3310103003330111-1301101121321032-2023310110100223-3113201312223210-1303020023220113-1023221003132222"></a>

## Next pages — checkin / 112021200100 / 4

- [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--reference--group-003.md#canonical-3022031310300013-1303031212122003-1101220112133123-2031112213031201-2311331321103330-2230221333213230-0213031113013220-0222011332013103)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131220113230320-3120210030302131-2302232013132201-3332000133002220-2300012102210001-1310131000312210-2302220113130121-0020031230132012"></a>

## cloudfront.protected_endpoints.flow_label.profile_management — profile_management / 200202101200 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- cloudfront.protected_endpoints.flow_label.profile_management

<a id="canonical-2303110110213010-0232001133101302-3333101000112220-0120202302203331-2300031022332211-2333332101033002-0131203101001320-2320311002011232"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Profile Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

<a id="canonical-3232012023132211-1221121333323123-1123303212012001-1033120102303113-0123221113313101-0103312101100200-0133032300003333-0000112030130010"></a>

## Direct properties — profile_management / 200202101200 / 3

- [create](data-sources--protected_application--reference--group-003.md#canonical-2232302021010323-0301323323203311-2200330211003110-3221301223102230-3231113312102122-2031123321101113-1222102113210223-2213313230111321): complete subsection reference.

- [update](data-sources--protected_application--reference--group-003.md#canonical-1112121212322033-0313122000310130-1331020302112002-3011122213232100-1202213230302302-1020130102223133-2323113300121210-1301023320032233): complete subsection reference.

- [view](data-sources--protected_application--reference--group-003.md#canonical-0331031132000322-3212033022100310-2200111333110122-1213012103003213-1130233000213311-2211322131301013-3322303101002003-0213330200132221): complete subsection reference.

<a id="canonical-1332302211013213-3233111101010232-0223121120021310-0113313332032013-0033300103022332-2023121200002321-3300022111323010-2130201221121010"></a>

## Next pages — profile_management / 200202101200 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management.create](data-sources--protected_application--reference--group-003.md#canonical-2232302021010323-0301323323203311-2200330211003110-3221301223102230-3231113312102122-2031123321101113-1222102113210223-2213313230111321)
- [cloudfront.protected_endpoints.flow_label.profile_management.update](data-sources--protected_application--reference--group-003.md#canonical-1112121212322033-0313122000310130-1331020302112002-3011122213232100-1202213230302302-1020130102223133-2323113300121210-1301023320032233)
- [cloudfront.protected_endpoints.flow_label.profile_management.view](data-sources--protected_application--reference--group-003.md#canonical-0331031132000322-3212033022100310-2200111333110122-1213012103003213-1130233000213311-2211322131301013-3322303101002003-0213330200132221)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2232302021010323-0301323323203311-2200330211003110-3221301223102230-3231113312102122-2031123321101113-1222102113210223-2213313230111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312322221332201-2121102311020201-3001221233311301-3130202230012232-0011012111133333-1323330301223020-0202211103321120-0122221322332020"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.create — create / 102001030321 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- cloudfront.protected_endpoints.flow_label.profile_management.create

<a id="canonical-3033232101003130-0130121022102313-3003233210333122-3322323230200121-3022303311130021-3223022113223210-3332313020132301-1121121011203332"></a>

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

<a id="canonical-3000010221111111-1123122213231113-1011030310312120-1323323122031311-0311300313023312-0111223331230032-2212032320123230-3213101330001032"></a>

## Direct properties — create / 102001030321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221103012311130-1333311201303323-2002102332313331-2133130103302301-1300012332102121-1101000023123313-2132212021101123-2211020222100000"></a>

## Next pages — create / 102001030321 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1112121212322033-0313122000310130-1331020302112002-3011122213232100-1202213230302302-1020130102223133-2323113300121210-1301023320032233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300101201221011-1320032323011001-1110223310003023-3013013011220120-2230020330213212-2130131230223103-1203313301233311-3020122110200332"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.update — update / 023122231001 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- cloudfront.protected_endpoints.flow_label.profile_management.update

<a id="canonical-2023002023120122-0321323102011210-0021310330010132-1321001130212333-1323111030222233-0232213310111123-0320323213331010-2113210223003102"></a>

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

<a id="canonical-3212331303221312-1030203011323031-0011132001021320-1202211212231310-3300130122102001-3231013102322010-3110300221101102-1223201203123302"></a>

## Direct properties — update / 023122231001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201223120300321-1100200332202121-3022033332130200-2132330031112200-1201001321313122-2302213013311032-2220112101331231-0212111223013012"></a>

## Next pages — update / 023122231001 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0331031132000322-3212033022100310-2200111333110122-1213012103003213-1130233000213311-2211322131301013-3322303101002003-0213330200132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300122233220221-1000021013203000-2230220021333103-3330001233302331-1000200321130122-3011312231322322-2203010231300032-3002333300311221"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.view — view / 213322022131 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- cloudfront.protected_endpoints.flow_label.profile_management.view

<a id="canonical-0011211333313021-3300023230323311-1323231023201311-3101010312130033-1030231332211103-1212120320103111-1212033200232001-3003232213002103"></a>

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

<a id="canonical-0222203131232101-0120211210203033-3011333201112001-0133211331002112-2322211330103010-0233100002032331-2203021130332331-2231033013233033"></a>

## Direct properties — view / 213322022131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113033211221111-0000133303110030-0011301111303322-3213023210001311-2031133310121221-3220021013021333-3232202333032132-3030122001232210"></a>

## Next pages — view / 213322022131 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2101313131303013-1011212123321211-1000323123112313-0033102233312313-3202223332111233-2301310032010101-2021313232221011-2001032013103303)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213020203000121-2112030122110010-3313030110301000-1103002301223112-1020312212323131-1300223233211200-3110110333031210-3022302312331132"></a>

## cloudfront.protected_endpoints.flow_label.search — search / 333212011100 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- cloudfront.protected_endpoints.flow_label.search

<a id="canonical-2230122110213122-1110120131032001-3233210112103300-2013030213210220-2010002313212221-2311233132223310-2030032121230020-3001013202212023"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Upstream description:

Bot Defense Flow Label Search Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

<a id="canonical-2003032202321023-2303131123330320-0020221001200020-3332121013000011-3223111320312020-2320311323301120-2012132113222202-1320221022012321"></a>

## Direct properties — search / 333212011100 / 3

- [flight_search](data-sources--protected_application--reference--group-003.md#canonical-3133321302102133-1012302212210310-3312223232232012-2122313220103232-2130131321313132-0313111131120120-0301111110132221-3333120311231101): complete subsection reference.

- [product_search](data-sources--protected_application--reference--group-003.md#canonical-3021203312321323-1300320101103213-0221000203103100-1012100232112212-0320020302122331-0122312222220020-2221333313301021-0003113312313232): complete subsection reference.

- [reservation_search](data-sources--protected_application--reference--group-003.md#canonical-0021112012221110-0311122200132230-0021031131122110-2331303102123102-0301031133233112-1032231231221302-0002211111310001-1123110320230310): complete subsection reference.

- [room_search](data-sources--protected_application--reference--group-003.md#canonical-0002232210033023-2112013010311131-2233332110332220-0133033211000330-3313220210030120-2221010020322320-2103111223003113-0203121311321301): complete subsection reference.

<a id="canonical-3100002021102122-0120110233030230-0320023221102001-0023030031232323-1110201303021330-2323010012201003-1100100230013021-2033301232001101"></a>

## Next pages — search / 333212011100 / 4

- [cloudfront.protected_endpoints.flow_label.search.flight_search](data-sources--protected_application--reference--group-003.md#canonical-3133321302102133-1012302212210310-3312223232232012-2122313220103232-2130131321313132-0313111131120120-0301111110132221-3333120311231101)
- [cloudfront.protected_endpoints.flow_label.search.product_search](data-sources--protected_application--reference--group-003.md#canonical-3021203312321323-1300320101103213-0221000203103100-1012100232112212-0320020302122331-0122312222220020-2221333313301021-0003113312313232)
- [cloudfront.protected_endpoints.flow_label.search.reservation_search](data-sources--protected_application--reference--group-003.md#canonical-0021112012221110-0311122200132230-0021031131122110-2331303102123102-0301031133233112-1032231231221302-0002211111310001-1123110320230310)
- [cloudfront.protected_endpoints.flow_label.search.room_search](data-sources--protected_application--reference--group-003.md#canonical-0002232210033023-2112013010311131-2233332110332220-0133033211000330-3313220210030120-2221010020322320-2103111223003113-0203121311321301)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3133321302102133-1012302212210310-3312223232232012-2122313220103232-2130131321313132-0313111131120120-0301111110132221-3333120311231101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233110333012100-1233231302000130-3103232003230122-1203001120033100-0133320130002322-0312230331123022-3302211021110100-0012330302130023"></a>

## cloudfront.protected_endpoints.flow_label.search.flight_search — flight_search / 201032012110 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003)
- cloudfront.protected_endpoints.flow_label.search.flight_search

<a id="canonical-3320032311010300-1112022210110230-0231102330123022-0020200011121122-1221023121303013-3010213213002222-2002123103100203-0033302323101213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for flight search.

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

<a id="canonical-1203023111220111-2311200300101213-1131203210110102-3110221201013031-0221121323303100-1100033033011230-2202130333002122-0122021032323122"></a>

## Direct properties — flight_search / 201032012110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101230320320121-2200210131021110-2133123320222311-1131031201222200-2113331231210220-2211230202311302-3020120100003022-0203322002231003"></a>

## Next pages — flight_search / 201032012110 / 4

- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3021203312321323-1300320101103213-0221000203103100-1012100232112212-0320020302122331-0122312222220020-2221333313301021-0003113312313232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012000301123230-2133313033323230-1111023212121101-3230022201003020-0111312320020103-1212213001021220-1002322120103022-3323223223323201"></a>

## cloudfront.protected_endpoints.flow_label.search.product_search — product_search / 131033332330 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003)
- cloudfront.protected_endpoints.flow_label.search.product_search

<a id="canonical-2012312231223230-3221012223033223-1101331103003103-1022211110012232-1203233220331321-1131231202322112-3000232313303130-3223013313332001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for product search.

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

<a id="canonical-1321310213032122-2103210200331323-1102100031201002-3320120323002310-1233212221333210-3111202000312003-0220333000022320-3001110312202211"></a>

## Direct properties — product_search / 131033332330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013110020011010-1233022003210310-1333313300032223-1211223331002212-2112200122303122-0100021230120012-1331030223213312-2133213030203001"></a>

## Next pages — product_search / 131033332330 / 4

- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0021112012221110-0311122200132230-0021031131122110-2331303102123102-0301031133233112-1032231231221302-0002211111310001-1123110320230310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212120331102330-1101000321002010-0313001303201133-2321113001321100-1032011303003322-2102000131213130-2301322331303320-0301120021330030"></a>

## cloudfront.protected_endpoints.flow_label.search.reservation_search — reservation_search / 013103222333 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003)
- cloudfront.protected_endpoints.flow_label.search.reservation_search

<a id="canonical-3111021213123323-1331313023032322-3222322122210221-0320113013231302-3111323000211210-2230112110233120-0312030013331113-1100302232301312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reservation search.

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

<a id="canonical-2031223030131221-0223013132122220-1031203230323230-1113120132021223-3100231013313031-3001010213230201-1030020022331112-0300001121331310"></a>

## Direct properties — reservation_search / 013103222333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022201022200033-2332212303221200-2003030120131311-3032123233033330-0211001011021232-2012000313311301-0112033323233310-2203332310312211"></a>

## Next pages — reservation_search / 013103222333 / 4

- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0002232210033023-2112013010311131-2233332110332220-0133033211000330-3313220210030120-2221010020322320-2103111223003113-0203121311321301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331023113323023-3331312303333030-3313023021222102-0033301210130120-2022203212320131-1213202022132131-3032312200021233-3220121022111331"></a>

## cloudfront.protected_endpoints.flow_label.search.room_search — room_search / 321312222030 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003)
- cloudfront.protected_endpoints.flow_label.search.room_search

<a id="canonical-2022111222122031-1023322310111013-3202131313221030-3212000313112132-0300000320332311-0201122112123303-2033113121213203-1313110212131303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for room search.

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

<a id="canonical-2213110131000233-0000312323203120-2330202021203201-1320230210030321-3020002313200311-1213221331332103-2212200302231213-1301100122213033"></a>

## Direct properties — room_search / 321312222030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201032332301230-3311223000111300-0113200231220211-2021213110020121-0133101033302101-3330233031113321-3111323131110323-0021202103302321"></a>

## Next pages — room_search / 321312222030 / 4

- [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-1313223000232301-2211001332221312-1323112313113301-2113202013233212-1110323031322032-3101232010200320-3111103203333020-3311313011210003)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101101112110001-1300222031113213-3233111313313232-0232120023013301-2310210320013101-2212000230322311-2100213001302122-2021233021122301"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards — shopping_gift_cards / 321133321303 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards

<a id="canonical-1201220110210032-0222222220330220-2320333230100312-1221212001303310-1133121001233200-3233203133331003-3032131032201103-2101212313111010"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

<a id="canonical-2312001103101022-0000112123002003-1303133003221113-3122030212322111-3332032330103320-3030233302331133-0223020312210010-3301321030222033"></a>

## Direct properties — shopping_gift_cards / 321133321303 / 3

- [gift_card_make_purchase_with_gift_card](data-sources--protected_application--reference--group-003.md#canonical-3112010321101110-2123200221133012-2112301313000323-1100020231002131-3102020200210302-1211023210021023-1312102311113332-1132113113121022): complete subsection reference.

- [gift_card_validation](data-sources--protected_application--reference--group-003.md#canonical-1311130300033131-0321333221301331-3032232031101323-1132021021312302-3213120003333223-0122003131102131-3330331301212000-0013131012212332): complete subsection reference.

- [shop_add_to_cart](data-sources--protected_application--reference--group-003.md#canonical-1032132103303231-2021300131232011-3301321001031200-1301233031320121-2131130231223201-2212203113010032-3322200232112211-1031213230002203): complete subsection reference.

- [shop_checkout](data-sources--protected_application--reference--group-003.md#canonical-3201010213020232-0203123102222112-2333113002222302-2100003332221330-3221131301131201-3101231122213321-0202310033200233-3203133233303102): complete subsection reference.

- [shop_choose_seat](data-sources--protected_application--reference--group-003.md#canonical-0302112230300202-1310102123100022-0031302132132013-0212230202221120-0333321001102103-1012310210300210-1122130020220233-3311131203100002): complete subsection reference.

- [shop_enter_drawing_submission](data-sources--protected_application--reference--group-003.md#canonical-3200302210233310-1202113111201321-2123203030022113-0310313120303000-2030302200023132-1220322110102133-0330011221121002-0001302223233032): complete subsection reference.

- [shop_make_payment](data-sources--protected_application--reference--group-003.md#canonical-0312101130020121-0213313013321213-0010201311323120-1010323131002013-3012220232102113-0230223312002133-0200010222003000-1033001323022022): complete subsection reference.

- [shop_order](data-sources--protected_application--reference--group-003.md#canonical-1201121203100333-1123130020330311-3100212113101001-0020220002033122-3101001001333100-2313031111113311-1013123022102331-2033221102312000): complete subsection reference.

- [shop_price_inquiry](data-sources--protected_application--reference--group-003.md#canonical-2031113213300021-1101231331131112-3320201212332033-2110200311203313-3121213310223122-2112100301022000-0000002031021023-2211332033030311): complete subsection reference.

- [shop_promo_code_validation](data-sources--protected_application--reference--group-003.md#canonical-0013322120111000-1113110121032211-2031312111222212-1211200011102023-1100302211112133-3300033020321111-3133220031031100-2133231313112001): complete subsection reference.

- [shop_purchase_gift_card](data-sources--protected_application--reference--group-003.md#canonical-1033002030200301-1003111320302222-1132020010021010-0222301122303200-1000313222210111-2123100033031010-1223020013001202-1010300103132223): complete subsection reference.

- [shop_update_quantity](data-sources--protected_application--reference--group-003.md#canonical-2020322120330013-0120203002202100-0032032030022331-2022332121312233-0322201310210031-3223202313303211-3101202101221030-3013312330003023): complete subsection reference.

<a id="canonical-3131231210230201-3130021023000000-0201330122003322-3103231033003331-1322001003103232-1011031223033030-2130233233201110-2330112033023301"></a>

## Next pages — shopping_gift_cards / 321133321303 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](data-sources--protected_application--reference--group-003.md#canonical-3112010321101110-2123200221133012-2112301313000323-1100020231002131-3102020200210302-1211023210021023-1312102311113332-1132113113121022)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](data-sources--protected_application--reference--group-003.md#canonical-1311130300033131-0321333221301331-3032232031101323-1132021021312302-3213120003333223-0122003131102131-3330331301212000-0013131012212332)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](data-sources--protected_application--reference--group-003.md#canonical-1032132103303231-2021300131232011-3301321001031200-1301233031320121-2131130231223201-2212203113010032-3322200232112211-1031213230002203)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](data-sources--protected_application--reference--group-003.md#canonical-3201010213020232-0203123102222112-2333113002222302-2100003332221330-3221131301131201-3101231122213321-0202310033200233-3203133233303102)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](data-sources--protected_application--reference--group-003.md#canonical-0302112230300202-1310102123100022-0031302132132013-0212230202221120-0333321001102103-1012310210300210-1122130020220233-3311131203100002)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](data-sources--protected_application--reference--group-003.md#canonical-3200302210233310-1202113111201321-2123203030022113-0310313120303000-2030302200023132-1220322110102133-0330011221121002-0001302223233032)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](data-sources--protected_application--reference--group-003.md#canonical-0312101130020121-0213313013321213-0010201311323120-1010323131002013-3012220232102113-0230223312002133-0200010222003000-1033001323022022)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](data-sources--protected_application--reference--group-003.md#canonical-1201121203100333-1123130020330311-3100212113101001-0020220002033122-3101001001333100-2313031111113311-1013123022102331-2033221102312000)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](data-sources--protected_application--reference--group-003.md#canonical-2031113213300021-1101231331131112-3320201212332033-2110200311203313-3121213310223122-2112100301022000-0000002031021023-2211332033030311)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](data-sources--protected_application--reference--group-003.md#canonical-0013322120111000-1113110121032211-2031312111222212-1211200011102023-1100302211112133-3300033020321111-3133220031031100-2133231313112001)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](data-sources--protected_application--reference--group-003.md#canonical-1033002030200301-1003111320302222-1132020010021010-0222301122303200-1000313222210111-2123100033031010-1223020013001202-1010300103132223)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](data-sources--protected_application--reference--group-003.md#canonical-2020322120330013-0120203002202100-0032032030022331-2022332121312233-0322201310210031-3223202313303211-3101202101221030-3013312330003023)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3112010321101110-2123200221133012-2112301313000323-1100020231002131-3102020200210302-1211023210021023-1312102311113332-1132113113121022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331101323103110-2133331222210212-0020312231321131-0112223123212113-3101102311332010-3310231303133233-0313011213220103-3333011003122312"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card — gift_card_make_purchase_with_gift_card / 033320320231 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-3102312231001223-2231331203211103-3123132322130202-3111223003101031-3020020131203211-2133002122112123-3003202320000332-3301121130133012"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for gift card make purchase with gift card.

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

<a id="canonical-1023011031122011-0230130020113103-3212321013301013-1103232231021310-0013320231123231-0110200002130121-0003121030120332-0033203010122023"></a>

## Direct properties — gift_card_make_purchase_with_gift_card / 033320320231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002113313212333-2103022120230031-3131110020301230-2220200233223231-0031003232303310-1101130233021101-2213101333023213-2001022000321331"></a>

## Next pages — gift_card_make_purchase_with_gift_card / 033320320231 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1311130300033131-0321333221301331-3032232031101323-1132021021312302-3213120003333223-0122003131102131-3330331301212000-0013131012212332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013200332202123-1302330113001100-3011110121333213-3213322021221323-3310310223202320-2122110323013230-2011131000001133-0321202201021303"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation — gift_card_validation / 012111230110 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-2130201102102003-1312021230333312-0000121200230133-3301323131200211-0301101101220311-2300031213313001-3032330101012120-3110302222030300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for gift card validation.

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

<a id="canonical-3321113222112001-3133012123100120-1002020221010020-0022121010111201-0223211002202301-3112133230212030-2332021200202230-1201120332102220"></a>

## Direct properties — gift_card_validation / 012111230110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332020331010120-1103321330133220-3223103333221221-1212013212212013-1112230033230122-2331121123123312-0131101113112012-3331332110332223"></a>

## Next pages — gift_card_validation / 012111230110 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1032132103303231-2021300131232011-3301321001031200-1301233031320121-2131130231223201-2212203113010032-3322200232112211-1031213230002203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103122310013103-3001132110001303-1310300220022210-3130103020133301-2000221012230213-3103302011101233-2202233321300102-2312331303020011"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart — shop_add_to_cart / 233232332031 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-0103011120333033-0231032003232011-0233320210232300-1210032121223310-1210022233310113-2112130220202213-1122330302030201-1313100132301302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop add to cart.

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

<a id="canonical-1022133112213232-3220011031301003-3032000212002320-0220212030331312-2012022103032312-1101123102213302-3003030313331103-1202200132120003"></a>

## Direct properties — shop_add_to_cart / 233232332031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003220102311231-0230103112102001-3123130220103133-1031033302011313-1202120023201322-1131000311013032-3312220023003033-0221120303011320"></a>

## Next pages — shop_add_to_cart / 233232332031 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3201010213020232-0203123102222112-2333113002222302-2100003332221330-3221131301131201-3101231122213321-0202310033200233-3203133233303102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010011032201000-1230230112311212-1102130223302120-2333022121221131-0201031303121300-3112213332230221-0322233033100111-1120020222020111"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout — shop_checkout / 302113331302 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-0131112222223321-3220121211321201-2101011010012331-0113020213121030-1012320031200133-3110013002212021-3122030112102233-2001000032011221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop checkout.

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

<a id="canonical-3213201021023213-3132310033213113-1020311000012203-0021031323003112-1121220013012112-1023111213011010-3211003123300120-0100131212230230"></a>

## Direct properties — shop_checkout / 302113331302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030023300233231-2203000112010031-2110300220313111-0322123021013122-0131012221002121-2120311101211010-2311121333311322-1130301202200012"></a>

## Next pages — shop_checkout / 302113331302 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0302112230300202-1310102123100022-0031302132132013-0212230202221120-0333321001102103-1012310210300210-1122130020220233-3311131203100002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321230010320112-3223303331200211-3230133212322131-2113332002031330-2211013031313310-3312301203221031-1202101323323102-2100022121133322"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat — shop_choose_seat / 131313322123 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-1030022113131102-2020321122310200-0313133330201222-3330011100012132-0023122000313033-0013220331301123-1033012302111222-0130210213330111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop choose seat.

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

<a id="canonical-1130103110110313-2331220300021112-1033322302230110-0302211032031102-0310101230020310-2000121313213301-3220021110020330-0213303121033132"></a>

## Direct properties — shop_choose_seat / 131313322123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132132013211133-0210010331133330-0223211113321233-0302210133233232-3223032123003031-0023302112110023-1323221102120113-3033100103113232"></a>

## Next pages — shop_choose_seat / 131313322123 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3200302210233310-1202113111201321-2123203030022113-0310313120303000-2030302200023132-1220322110102133-0330011221121002-0001302223233032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120330000330003-1231223320332200-3010001232203211-1202231123320211-1310012330312312-2120300012011211-0330313010232000-1133210223002212"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission — shop_enter_drawing_submission / 013100333122 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-0230101322032220-1313022012323120-1300112112012300-1212020303123320-2330031231033031-2223113013110210-1102221233301221-2330130130322332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop enter drawing submission.

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

<a id="canonical-1320020121022331-1112132120130321-0132020121313231-2032231112001123-2122110231001110-3120111332133311-1302002113221222-3111201303133321"></a>

## Direct properties — shop_enter_drawing_submission / 013100333122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330323013230323-1132003031321010-1002132332132021-2330213221021120-2133331310110032-1032210332333311-1232303000211010-3002321002022222"></a>

## Next pages — shop_enter_drawing_submission / 013100333122 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0312101130020121-0213313013321213-0010201311323120-1010323131002013-3012220232102113-0230223312002133-0200010222003000-1033001323022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223312302212031-0000022320020323-3110033213003002-2213030031200221-1231212112223113-3330211310222120-0330302000103333-2103203012111232"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment — shop_make_payment / 210000021000 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-1212321202210332-1113220023331001-3323212222103212-2100231331302022-3001000310112001-0302012003333020-0110211232301122-1302220201002220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop make payment.

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

<a id="canonical-3210222033130100-1132001300331132-2023131210212012-0310323011012321-0022332101323022-0022301023231101-1022100032222000-2323300212003331"></a>

## Direct properties — shop_make_payment / 210000021000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211033022201103-3021322012020013-0210102231000112-2102132022032121-0122202211132233-1201221000201302-1111012210311331-3021132321020232"></a>

## Next pages — shop_make_payment / 210000021000 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1201121203100333-1123130020330311-3100212113101001-0020220002033122-3101001001333100-2313031111113311-1013123022102331-2033221102312000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131020013323321-0221323323300011-3230223223201332-2202231130313131-2332032102113131-2133323023332010-2201010201033121-0201033003011302"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order — shop_order / 031033301022 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-0322202003100020-0213321032201120-2333001311001320-0120213020330012-1032030212001210-0001020012311012-1111321230300222-1103123002022320"></a>

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

<a id="canonical-2131333100332331-3302212011111200-1302033232002322-0111221111030320-0322111232323323-1313131310331331-3133211220201123-2123112020201310"></a>

## Direct properties — shop_order / 031033301022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010112313320321-1003032230011021-0103002311020030-1121210233011321-1100312210221201-0201312110012002-1311301003222101-0323333333212120"></a>

## Next pages — shop_order / 031033301022 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2031113213300021-1101231331131112-3320201212332033-2110200311203313-3121213310223122-2112100301022000-0000002031021023-2211332033030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320223220131303-3300231130132033-3232303121011030-1220212003131211-1020231132101020-1101233201121233-3310313002121110-3110103222123212"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry — shop_price_inquiry / 303221020210 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-3132132203113113-2223113301110110-2120313233113330-0002031321133200-1101333131010320-3212311103011320-0222223330023220-1022331002233210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop price inquiry.

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

<a id="canonical-2003121230332130-2120111112032320-2211112312133230-0322020132230110-0133121211311201-3213102220020022-1101133312221202-0222200223331100"></a>

## Direct properties — shop_price_inquiry / 303221020210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222312100103011-2330211320232001-3213202003302100-2323320221311101-2332300233212000-0002203333003021-2313321303323211-3201111320023330"></a>

## Next pages — shop_price_inquiry / 303221020210 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0013322120111000-1113110121032211-2031312111222212-1211200011102023-1100302211112133-3300033020321111-3133220031031100-2133231313112001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101232300221111-1102210110231333-1000212201123133-3001322311030312-1202213031232100-3231312303002101-2333333022111203-1212003223010000"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation — shop_promo_code_validation / 111211112221 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation

<a id="canonical-1210312200102232-1233331312321120-0011131202100003-2122203303333111-2320003012003301-0003131010031022-0101213132300213-2310202321232302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop promo code validation.

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

<a id="canonical-0122100221331133-0313230011131323-1022111320302100-0102303313010010-0310122113100111-3200100001321322-3200112222323033-1013333303233303"></a>

## Direct properties — shop_promo_code_validation / 111211112221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331120011110103-1132201311321120-1000311311033032-2313301121302133-0322111203300330-3202031111022231-3002203122023311-0310130311223311"></a>

## Next pages — shop_promo_code_validation / 111211112221 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1033002030200301-1003111320302222-1132020010021010-0222301122303200-1000313222210111-2123100033031010-1223020013001202-1010300103132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110212302122100-1211101110102232-0112300200231211-3101310320321211-1312332322123131-2300021021330030-0132002331221312-0321132123333301"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card — shop_purchase_gift_card / 323003011210 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card

<a id="canonical-3022012322113332-2120133310310220-3023003111203110-2130202131230333-2002201112020122-0132013211033331-1310221213320301-2112023103201312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop purchase gift card.

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

<a id="canonical-0321011313233212-3332030021012021-0002320322101211-2333132030212233-0120332312212300-2302110202311100-0111122021303101-0222320302322102"></a>

## Direct properties — shop_purchase_gift_card / 323003011210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201133110011212-2301000233233333-0313111121132122-0213203231202132-0311002030330130-0123013320301001-0133322030221332-3130222123113030"></a>

## Next pages — shop_purchase_gift_card / 323003011210 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2020322120330013-0120203002202100-0032032030022331-2022332121312233-0322201310210031-3223202313303211-3101202101221030-3013312330003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310110132012212-1102122023332202-3111133311012023-1022200002113020-3223233333200203-2201001331023231-0113232023212113-1303233200032320"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity — shop_update_quantity / 310333001300 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-3032000120121202-2111103301013032-0021202200310132-0201022310021010-1113132210202231-1213002223120001-0203001111102022-1003210002011300)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity

<a id="canonical-0213020101000331-0113230001002102-0131200111013323-0230003223312211-2021310232121020-0022331330111110-1113110312202103-3003312220103331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for shop update quantity.

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

<a id="canonical-1112333330002220-2301120002130303-1002300130121111-1220230201232220-1120011102203233-2132212130032003-3021112013133203-3213301320132033"></a>

## Direct properties — shop_update_quantity / 310333001300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333203203232222-3321132133013102-0122033331310023-0101002120011313-1330102331300120-0220111223333230-2002322101133102-3231332113013331"></a>

## Next pages — shop_update_quantity / 310333001300 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-0223201300000202-3102010220222300-1300101113101310-2002223200102330-3333233210032301-1321110003200320-2203203131001100-0213112010210321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2021230303322000-2030220332003021-3120230122221302-3030222030333130-2213211332121312-1222122333301110-3220301200123102-0201311323123022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301233212222200-2120223221231130-0200033013210202-1011001303231111-3102033013002002-2113200002211231-1032233133002312-1133010220003130"></a>

## cloudfront.protected_endpoints.metadata — metadata / 211103012020 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- cloudfront.protected_endpoints.metadata

<a id="canonical-1110001100020130-3103033011210222-3100333313123103-3322333123021233-0321332013212031-0101103202330130-1121233111030123-2330001011002322"></a>

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

<a id="canonical-1010302101322031-0222333300212230-0213010121212013-3122010220133333-0010022130100002-2132130130112330-0123013101013103-1031003330113113"></a>

## Direct properties — metadata / 211103012020 / 3

<a id="canonical-0030231303213012-1312233012332211-0102320012120013-2133211022311301-3230233311032030-3013230220011202-3331000111332210-2210212233222322"></a>

<a id="canonical-2113222300020110-0102220022000121-2200111321000301-3131000310220103-3113222222110203-0332021113000012-3033233332203311-0022020103011213"></a>

## description_spec property — metadata / 211103012020 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3021313132000132-1121210301232100-1222221101003000-2132010321031130-1010321113321113-0321000121003011-3033120003001303-0321112013013222"></a>

<a id="canonical-3333020122111330-1330312303123322-1233101201022033-2002313221131133-3020311212320330-0302333010332223-2210033232113210-1200212323003301"></a>

## name property — metadata / 211103012020 / 5

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

<a id="canonical-0200012131023330-0200300013201110-0222003001122023-1101123131131313-1020330001221313-1023133232232202-1000312031332101-0231103203033201"></a>

## Next pages — metadata / 211103012020 / 6

- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332230000233120-2013133013123202-1232001300231233-0202221110303231-0302310123133313-3202033001220012-1311323033312320-0210312001002011"></a>

## cloudfront.protected_endpoints.mobile_client — mobile_client / 211112111000 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- cloudfront.protected_endpoints.mobile_client

<a id="canonical-0012112132311220-3113112100222332-2312111022211122-0100131000310023-1030012122303200-1201102321123003-3310000310201122-3222301102100232"></a>

Type: `"single"`. Computed.

Mobile Client. Mobile client configuration OPTIONS.

Upstream description:

Mobile client configuration OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

<a id="canonical-3202303320220013-2230103302232230-0131333302203222-3013310103310031-3010033211323333-0333302303003323-3323021013210233-1321200311001012"></a>

## Direct properties — mobile_client / 211112111000 / 3

- [block](data-sources--protected_application--reference--group-003.md#canonical-2031232202233300-3031001213312310-0023202021220202-0102222312013022-2222200002211321-1330001113322120-2123232120301223-3311331211223020): complete subsection reference.

- [continue](data-sources--protected_application--reference--group-003.md#canonical-0023212102233030-1310113233310133-1001111232030332-0031222031303032-1012003301313330-3122230011323003-1120310200220230-0010330121001203): complete subsection reference.

<a id="canonical-1321231133132230-0211213103132032-2122332021101022-3020322200210321-2013221320312132-0031332310103020-1220222003033200-1031130223130132"></a>

## Next pages — mobile_client / 211112111000 / 4

- [cloudfront.protected_endpoints.mobile_client.block](data-sources--protected_application--reference--group-003.md#canonical-2031232202233300-3031001213312310-0023202021220202-0102222312013022-2222200002211321-1330001113322120-2123232120301223-3311331211223020)
- [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-0023212102233030-1310113233310133-1001111232030332-0031222031303032-1012003301313330-3122230011323003-1120310200220230-0010330121001203)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2031232202233300-3031001213312310-0023202021220202-0102222312013022-2222200002211321-1330001113322120-2123232120301223-3311331211223020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331213130303132-1122222102310000-3101133303210130-0002003223013003-0221012332233013-0200202301120132-0030030022233122-2013101302203312"></a>

## cloudfront.protected_endpoints.mobile_client.block — block / 102323223110 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322)
- cloudfront.protected_endpoints.mobile_client.block

<a id="canonical-0310113123330330-1013111210332022-1031311202011212-1101221020222013-0212301000102101-1232323021132002-1322000231302002-3033201233120033"></a>

Type: `"single"`. Computed.

Block Response for Mobile. Block Response.

Upstream description:

Block Response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3031132012011120-0202111130232010-3223302213013003-2112012301332230-2031220103111303-0033313230322133-2110302130330033-2110332133122330"></a>

## Direct properties — block / 102323223110 / 3

<a id="canonical-0212111030330113-0001323233021023-2303223113021031-1232303222011333-3010202021012012-2302221210013300-1220212003103301-3113012212121301"></a>

<a id="canonical-1102223220222233-2001201021012112-2301021120311332-3123101030222003-0023210213030223-0201100013010201-3131001013100213-2020211103021103"></a>

## body property — block / 102323223110 / 4

Type: `"string"`. Computed.

Body. Custom body message.

Upstream description:

Custom body message.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-2101332323131011-3010000130112303-0022230311012213-0321233100300020-0222002120102112-0020333212330212-2303200312030102-3311133323223123"></a>

<a id="canonical-0103032220201112-0103200110221320-3332101203222021-3020302332010302-1133130102020100-1033022012003033-2102011232112023-1320212011312301"></a>

## content_type property — block / 102323223110 / 5

Type: `"string"`. Computed.

Content type to use in a block response.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3023210001121311-0321023300121333-2233320232333332-3321220333213110-0121032030033012-0333112010123201-0312001211101011-0003111332201032"></a>

<a id="canonical-2212213311012102-0130220212213232-0103323322310331-3232002130012330-0301011331030212-2120332102010320-3231320232000211-2233113030031110"></a>

## status property — block / 102323223110 / 6

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103233022132130-3110013211101333-1331120320212213-0301022221031102-1133132103331100-0003103210321220-1003111312333101-1000123131223020"></a>

## Next pages — block / 102323223110 / 7

- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0023212102233030-1310113233310133-1001111232030332-0031222031303032-1012003301313330-3122230011323003-1120310200220230-0010330121001203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330112131233103-2031201112221030-1011310322231022-3303121112201123-0200030021111031-1102130312110333-3322121112311203-0103310232023021"></a>

## cloudfront.protected_endpoints.mobile_client.continue — continue / 102333330201 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322)
- cloudfront.protected_endpoints.mobile_client.continue

<a id="canonical-2220132120301131-3031010223230013-0010201221201202-1013111230331312-1021013100210102-1023332233133023-3013210133133110-2023133121233213"></a>

Type: `"single"`. Computed.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

<a id="canonical-1102101231312110-3031210310221221-3232133232122212-2311312112010121-1131133310233021-0300323210232132-2313222322103103-1130012301113313"></a>

## Direct properties — continue / 102333330201 / 3

- [add_header](data-sources--protected_application--reference--group-003.md#canonical-2011020102310002-2131300002323130-2020332323212203-0121210233322021-1233212130003310-1312020201213212-3312002313303331-1031030330321212): complete subsection reference.

- [no_header](data-sources--protected_application--reference--group-003.md#canonical-0220130330201122-1021330210031120-0100201113100201-1333330120233113-3230132322322132-3312313210123032-3000212121232232-2321031232102202): complete subsection reference.

<a id="canonical-1110311031112101-3313100332133302-3000222222221201-0331320101003023-0301320202221302-2233221320312032-1231312113311212-3020302113003320"></a>

## Next pages — continue / 102333330201 / 4

- [cloudfront.protected_endpoints.mobile_client.continue.add_header](data-sources--protected_application--reference--group-003.md#canonical-2011020102310002-2131300002323130-2020332323212203-0121210233322021-1233212130003310-1312020201213212-3312002313303331-1031030330321212)
- [cloudfront.protected_endpoints.mobile_client.continue.no_header](data-sources--protected_application--reference--group-003.md#canonical-0220130330201122-1021330210031120-0100201113100201-1333330120233113-3230132322322132-3312313210123032-3000212121232232-2321031232102202)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2011020102310002-2131300002323130-2020332323212203-0121210233322021-1233212130003310-1312020201213212-3312002313303331-1031030330321212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311312310320130-2123010011103032-3202133022122103-0313333011112322-1012011232012201-0020222111321200-0022120101331203-1023211123331211"></a>

## cloudfront.protected_endpoints.mobile_client.continue.add_header — add_header / 321203121121 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [cloudfront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-002.md#canonical-1203303031112020-1302021321101330-1000021310000231-2203021210212133-3020313120132311-2213032031221000-3213301332331133-1310321302102200)
- [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-0300100133111022-1320133032023113-0200312122311310-2330312002011033-0123110113021233-1111122331121201-0322031011211223-1323112133031322)
- [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-0023212102233030-1310113233310133-1001111232030332-0031222031303032-1012003301313330-3122230011323003-1120310200220230-0010330121001203)
- cloudfront.protected_endpoints.mobile_client.continue.add_header

<a id="canonical-2310323321213111-1230121301201001-0323102301200112-3111003003013101-3121303212033003-0230212230332011-0100213230022330-3222032313003221"></a>

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

<a id="canonical-3201201101321121-0312200232132323-2032313001231031-1113311321002110-3200223020112023-0203203031031022-0103230110021010-0333030020212320"></a>

## Direct properties — add_header / 321203121121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210022112100000-0330121300210131-3313300200322112-2323021221300031-1322000322132322-3130020011333300-0233321222131101-1100312222022020"></a>

## Next pages — add_header / 321203121121 / 4

- [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-0023212102233030-1310113233310133-1001111232030332-0031222031303032-1012003301313330-3122230011323003-1120310200220230-0010330121001203)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0220130330201122-1021330210031120-0100201113100201-1333330120233113-3230132322322132-3312313210123032-3000212121232232-2321031232102202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
