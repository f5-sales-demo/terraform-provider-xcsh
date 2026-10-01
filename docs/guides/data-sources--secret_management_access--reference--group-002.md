---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220022102021121-1200031222213311-2122122001000121-1031332100323030-1232321301110331-1032332203112000-3203210300120223-3232101132333322"></a>

## access_info.vault_auth_info.app_role_auth — app_role_auth / 200313103311 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- access_info.vault_auth_info.app_role_auth

<a id="canonical-2003300311230032-3101330022111120-0301320313012000-3332130101131323-2013122012110023-2021212303131033-3101112300131203-0330303120003020"></a>

Type: `"single"`. Computed.

AppRoleAuthInfoType contains parameters for AppRole authentication in Hashicorp Vault.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3113233132302313-0203013331220012-0230112000020122-2023022301333001-2010113011001101-3331323120320033-3313130031203103-2123120022312201"></a>

## Direct properties — app_role_auth / 200313103311 / 3

<a id="canonical-3003333122011202-0301020320203102-1303032111133312-0322311230213230-2323121323310331-1030300103102231-1212321312021333-0022030321023212"></a>

<a id="canonical-0113002321112022-2201211232100120-0032303022222000-1323221032023202-0101202331001000-3320313122232010-2021030302321011-0211300020033332"></a>

## role_id property — app_role_auth / 200313103311 / 4

Type: `"string"`. Computed.

Role ID. Role-ID to be used for authentication.

Upstream description:

Role-ID to be used for authentication.

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

- [secret_id](data-sources--secret_management_access--reference--group-002.md#canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020): complete subsection reference.

<a id="canonical-3302011011031330-2102001112203131-3203203333223030-1303302003112323-0010203212000210-2202313013101213-0310322311122102-0130020003201301"></a>

## Next pages — app_role_auth / 200313103311 / 5

- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-002.md#canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222300321231010-0110301302231121-2012002000220002-2112311133313000-2231011000100102-1122001002230311-2320203021012030-2230130131002113"></a>

## access_info.vault_auth_info.app_role_auth.secret_id — secret_id / 002311222111 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-002.md#canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022)
- access_info.vault_auth_info.app_role_auth.secret_id

<a id="canonical-2301133211213110-3003223123030322-1223012110113120-3210202011132312-1301110333011231-0033210033323113-1211310200120311-2110123233032210"></a>

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

<a id="canonical-1032211033112122-0001330013012032-3031033100232022-2012013310200300-2103123003211203-2230233110132221-1221303001312321-3202123323003201"></a>

## Direct properties — secret_id / 002311222111 / 3

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-1313201332032131-2130101023010103-3323213031000023-0211023332102230-1100103002222230-2202201123302213-0303312223222203-1020132213130330): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-3210131331231023-2130202300320212-3120220201021330-1100313132122113-3000202123200023-3001310010123021-0110223232203320-0003110101320221): complete subsection reference.

<a id="canonical-3001330003312032-3213022231301223-1030002032111103-3023023232021130-3132101133120210-2323021010310123-1130232112011311-1011323223203302"></a>

## Next pages — secret_id / 002311222111 / 4

- [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-1313201332032131-2130101023010103-3323213031000023-0211023332102230-1100103002222230-2202201123302213-0303312223222203-1020132213130330)
- [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-3210131331231023-2130202300320212-3120220201021330-1100313132122113-3000202123200023-3001310010123021-0110223232203320-0003110101320221)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-002.md#canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1313201332032131-2130101023010103-3323213031000023-0211023332102230-1100103002222230-2202201123302213-0303312223222203-1020132213130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221320333303111-3312322133100232-0031223312003011-2221303203200110-2133313322222302-3102223101302023-1310311112211122-0230132031010012"></a>

## access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info — blindfold_secret_info / 212030000123 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-002.md#canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022)
- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-002.md#canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020)
- access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info

<a id="canonical-1123031301231032-3020202231313010-2332313230102321-2121322011001102-2301300313333300-2220111003331022-1133011301302331-2312210000121001"></a>

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

<a id="canonical-1002232002201200-2220012212103111-1003031231310200-1012300020111203-1002300022133121-2312221131202332-3020303300133220-2032033113202330"></a>

## Direct properties — blindfold_secret_info / 212030000123 / 3

<a id="canonical-3332221120110303-3131221100012212-0203120011321300-1003320322300312-2203113102020113-3021022100231131-0031322110102001-2123230021000022"></a>

<a id="canonical-3002320321112130-0101332223120013-0311020211000301-3113321102220023-0001331012121233-3222322032322222-3033232230113130-0120110023323201"></a>

## decryption_provider property — blindfold_secret_info / 212030000123 / 4

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

<a id="canonical-2012012203201112-2213202132201311-3333332010202020-1330310201300220-1001332123001010-2322133103313313-0220302223232023-1330130021211123"></a>

<a id="canonical-1003211213000010-3333312321012202-0032132313012132-2021021230013201-2020012223030132-1012322122213102-2020313032012121-3033210110223010"></a>

## location property — blindfold_secret_info / 212030000123 / 5

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

<a id="canonical-1212122113102213-2312200130202113-0230330303320033-1022231322322123-0211110333130303-1321113133213112-1113111100012203-2121001021323130"></a>

<a id="canonical-0011200300320012-0133103032321332-0312012020222330-3202332312333013-2323023020030322-0212001210011001-1021300113113010-2330120033102020"></a>

## store_provider property — blindfold_secret_info / 212030000123 / 6

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

<a id="canonical-1101332310122231-1032231310320012-1210212033101231-3133112030023002-2230032102021002-0222223302030312-1231313332222230-0110131322123023"></a>

## Next pages — blindfold_secret_info / 212030000123 / 7

- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-002.md#canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3210131331231023-2130202300320212-3120220201021330-1100313132122113-3000202123200023-3001310010123021-0110223232203320-0003110101320221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003210301303323-3013233232112322-0232210012033213-0221033232010301-0220333110032121-2211031230102122-3232233123130130-1200322232223331"></a>

## access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info — clear_secret_info / 011112300333 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-002.md#canonical-0012301320210212-3110121123113000-2122321022210103-1322002120013000-2311021002123101-3121123121201133-3032133200103013-2133120000130022)
- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-002.md#canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020)
- access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info

<a id="canonical-3100233101303200-3021132120010212-1222112320313312-0321113032123231-0010020012201210-1300211221213200-3021001030132220-3210301222211221"></a>

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

<a id="canonical-0202133003031130-1103202200012101-3003033203223031-3301030132000000-1031321113030131-1130033111020112-3112100130023103-0121302113010232"></a>

## Direct properties — clear_secret_info / 011112300333 / 3

<a id="canonical-0111110323030220-1221133112102302-3020312111000201-0311122000330210-0330112102332332-3120022101211203-0303313312033320-0003121122320122"></a>

<a id="canonical-3310301131110322-1013000210030221-3203230023322332-1221302031233323-0211332300012311-1003201200123001-1331133313200101-1002123012210231"></a>

## provider_ref property — clear_secret_info / 011112300333 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0301122312022223-0320030333010321-3201331123120332-0111023123113100-0130203323013232-1132031333302111-0231322123002210-3113321120000333"></a>

<a id="canonical-1333321310012322-0222101021000310-3331131232202003-1212010010211303-3213210113112311-0330233012132203-0020302212120120-1123301331133120"></a>

## URL property — clear_secret_info / 011112300333 / 5

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

<a id="canonical-2212001223010001-3020032123002010-1323230133320133-3203131312201000-0312323002133000-1012211021202323-0300200032020231-1210202331310012"></a>

## Next pages — clear_secret_info / 011112300333 / 6

- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-002.md#canonical-2030321233233102-2312311211310012-1202222333323203-3112321211100133-1312211311100300-3011101111101112-0223223303202120-3300333012302020)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320302113303233-3300010302002100-2110000000023020-2332211221302101-0132030333012323-3312220302030011-2211103323332233-0100211020332322"></a>

## access_info.vault_auth_info.token — token / 300303232202 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- access_info.vault_auth_info.token

<a id="canonical-0022122320212220-3202033221010132-0222312101211131-2321320311103112-2202101021123031-2230111222301311-1222303121333330-2300210021233310"></a>

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

<a id="canonical-3321213021110302-0113011133232300-0332002102231023-1023202003122001-3123311202030210-3000023123021001-2111301330301111-1222223203220121"></a>

## Direct properties — token / 300303232202 / 3

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-3322012010221021-3220213301022210-3120103012201323-0210322010330132-3001131332202320-3212201213112200-1131211210133311-3303103222123010): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-1032012332233213-1220021321001313-2132023132202000-1000030120323133-3112100120311330-1100023213013300-1000230130203301-1020101322330113): complete subsection reference.

<a id="canonical-2200101303033110-1103203102110330-3102120322103311-3203021203002212-2103210111213001-2002213013123321-1033212313332110-1123203310023301"></a>

## Next pages — token / 300303232202 / 4

- [access_info.vault_auth_info.token.blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-3322012010221021-3220213301022210-3120103012201323-0210322010330132-3001131332202320-3212201213112200-1131211210133311-3303103222123010)
- [access_info.vault_auth_info.token.clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-1032012332233213-1220021321001313-2132023132202000-1000030120323133-3112100120311330-1100023213013300-1000230130203301-1020101322330113)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3322012010221021-3220213301022210-3120103012201323-0210322010330132-3001131332202320-3212201213112200-1131211210133311-3303103222123010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012103111030000-0322203133300023-3221112221133100-3000322220313003-1030200220030313-0231312202033333-0103110231103023-2233201220132132"></a>

## access_info.vault_auth_info.token.blindfold_secret_info — blindfold_secret_info / 331211130022 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203)
- access_info.vault_auth_info.token.blindfold_secret_info

<a id="canonical-2000101120331303-3230012210023113-2302112301002300-1312033301222103-2033032211313030-1102303102002311-0233021232002302-3132030301020322"></a>

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

<a id="canonical-1100220200030023-3333012332112021-1120122302330210-0102000322332103-0322021301000031-0222233233113210-3012232322203130-2031132112332211"></a>

## Direct properties — blindfold_secret_info / 331211130022 / 3

<a id="canonical-2123103202031232-3023230213333130-2111033331311223-0021031013201033-0012313001301213-0303213111300011-0233012332003320-0332021212303101"></a>

<a id="canonical-0310003222331202-0220003122202110-3202301132300001-2110031321313222-2001202222313232-1101132211123230-1012112012300012-1103333133030122"></a>

## decryption_provider property — blindfold_secret_info / 331211130022 / 4

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

<a id="canonical-2330203333102013-1201020331222013-3103132332000221-2222301231231220-1100023022033331-1031203020102231-3320313131030333-0112210330131331"></a>

<a id="canonical-2132100311203110-0312003132032133-2103322212021011-1303322202011133-0213113101212311-1210200013013323-3330031301230030-0022101131202212"></a>

## location property — blindfold_secret_info / 331211130022 / 5

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

<a id="canonical-1110201202123120-1123130112012000-1123231320031230-2232133100301222-3211303011333001-2210111230022023-1201303322113212-3213103321300030"></a>

<a id="canonical-3230032030200012-2201231200121210-3112023111032213-2301200311000032-2333200201130323-2200223130030301-0321213033221111-1302321311101102"></a>

## store_provider property — blindfold_secret_info / 331211130022 / 6

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

<a id="canonical-2330213003031010-1223033012002320-2220302201311123-2123013011222311-1020010212100232-2303311112120200-2100233200331013-3331212213300012"></a>

## Next pages — blindfold_secret_info / 331211130022 / 7

- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1032012332233213-1220021321001313-2132023132202000-1000030120323133-3112100120311330-1100023213013300-1000230130203301-1020101322330113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112122023211311-1321102010333133-0001120110133113-2330013030233110-0111132132000103-2210011111131022-0103202231123301-1001133031011233"></a>

## access_info.vault_auth_info.token.clear_secret_info — clear_secret_info / 002010012322 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-1322210000301221-3321003330330222-3030112221033300-3121302333020301-2222030302103023-2000100031132201-2000330210012102-1222232322020133)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-3032303213101100-1001300013332101-0301210111111320-2113001212300102-1031021233302032-1031020302223002-3013311200002010-2203121030101022)
- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203)
- access_info.vault_auth_info.token.clear_secret_info

<a id="canonical-1322100230302323-2211102110110000-0101322312020333-1300211312023200-3321301023133030-1330203320200110-0001032011213110-3213020033123313"></a>

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

<a id="canonical-1133300133020102-0213021003203023-0310230123221331-0330132031230211-0311322031113011-0121011222212331-1032133033200232-2033303232022113"></a>

## Direct properties — clear_secret_info / 002010012322 / 3

<a id="canonical-0010303232000132-0233132320002133-3220031313120133-1311103230331022-2010130101232303-0003031133222233-0300313333132302-1020113100331111"></a>

<a id="canonical-3300232000022000-0110302223020322-0031020131221120-0301201022323033-0310332132223003-3110211113120013-3203313232200223-1023112220032001"></a>

## provider_ref property — clear_secret_info / 002010012322 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2123203212021322-2130123100211003-1133310310010113-0222011121122013-2113321003002302-3310303312032130-2010110013012303-1010030222321013"></a>

<a id="canonical-3010133111120321-3302323121322222-2312301011021221-3232323022310321-3320031331120333-2030300200111113-1110321201013312-3321310322313331"></a>

## URL property — clear_secret_info / 002010012322 / 5

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

<a id="canonical-0233132001130101-0303330130213031-0300230131302012-0230200023310121-2201021130312122-1212013322033222-2102312201312023-3220323300322202"></a>

## Next pages — clear_secret_info / 002010012322 / 6

- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-2200222113222202-3210312221321013-3230302101001032-0202022233223101-0311212122310232-3200300221230331-2113223010002011-2202110112012203)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122321222122310-2120120322130020-3013021302203333-3030223110001002-2303323211020311-0001121231012303-0010312213321012-0001231112130021"></a>

## where — where / 023021132100 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- where

<a id="canonical-3002002312300203-3333333000211233-1230031002223332-2233310003333231-0321230033230330-2132001131033032-1210230322002222-0123221101110303"></a>

Type: `"single"`. Computed.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

<a id="canonical-0000012111210333-3133322202212000-2330120313303311-0210232110132101-0212110030331201-0202123312222202-2003131233130113-1322112120133001"></a>

## Direct properties — where / 023021132100 / 3

- [site](data-sources--secret_management_access--reference--group-002.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321): complete subsection reference.

- [virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-1010100312222110-1101123221213010-1320012022131302-2121232311200201-0133002000213230-3132212101210021-0321212330221100-1322311110200013): complete subsection reference.

- [virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212): complete subsection reference.

<a id="canonical-0132203321220213-2012020223312232-0320133202310313-3300011233210131-0311331302321321-2122221010101003-1220012011011101-1332132013102103"></a>

## Next pages — where / 023021132100 / 4

- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- [where.virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-1010100312222110-1101123221213010-1320012022131302-2121232311200201-0133002000213230-3132212101210021-0321212330221100-1322311110200013)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022203332102202-1303321303031203-1013302032002312-0332122121012133-1210111310122120-0003210103001322-2300022033001001-2311102200113230"></a>

## where.site — site / 112330131201 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- where.site

<a id="canonical-3332301333021022-3213033023310121-1000023201332233-3131300113321222-2103321032312310-0030301012000231-3233102010210212-1120213112220330"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-1322330321000133-2021010333013111-1302312011121221-1321101211032013-0032221102212103-2231002101001221-1320130303023203-3111313212132301"></a>

## Direct properties — site / 112330131201 / 3

- [disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-0020112212231223-0132311323130321-1120222122103021-0120213121301200-3113000133332023-2332011232112303-3330133000113322-3001220111210031): complete subsection reference.

- [enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-2111130133133231-1331101233332023-0113032000223033-1012103203012032-3011031332333030-1132330302303130-0330112330331310-3021223003231021): complete subsection reference.

<a id="canonical-2210021123301000-3323031211313332-1301202310021222-3311220210320320-1222331201230200-3112100013321230-0021221113101111-2120133122103022"></a>

<a id="canonical-1022210001000032-2132011312213113-2030021332033332-0023130003130023-0311210000021102-2132110121132233-1301221333201310-0320103113321112"></a>

## network_type property — site / 112330131201 / 4

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--secret_management_access--reference--group-002.md#canonical-3322223332330102-2330232001102203-2213010010102323-3200103311331011-0100130131321202-0222213330321201-0100121301230212-1331121303010101): complete subsection reference.

<a id="canonical-0032011120011322-3232123323133313-1330120200333121-3000331210311102-3220130232213220-0310111223003122-1102321112310013-0013222032021033"></a>

## Next pages — site / 112330131201 / 5

- [where.site.disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-0020112212231223-0132311323130321-1120222122103021-0120213121301200-3113000133332023-2332011232112303-3330133000113322-3001220111210031)
- [where.site.enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-2111130133133231-1331101233332023-0113032000223033-1012103203012032-3011031332333030-1132330302303130-0330112330331310-3021223003231021)
- [where.site.ref](data-sources--secret_management_access--reference--group-002.md#canonical-3322223332330102-2330232001102203-2213010010102323-3200103311331011-0100130131321202-0222213330321201-0100121301230212-1331121303010101)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0020112212231223-0132311323130321-1120222122103021-0120213121301200-3113000133332023-2332011232112303-3330133000113322-3001220111210031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312203212203032-1331203313232303-3203000112312030-1112110112021012-2030031212302100-0112012120100101-0032212013213201-0322203220031101"></a>

## where.site.disable_internet_vip — disable_internet_vip / 223200223330 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- where.site.disable_internet_vip

<a id="canonical-2012121010311023-0320322323221133-2231213331200223-0303303332211213-0300022003020012-2113213130312000-0131000211021201-0220110313303302"></a>

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

<a id="canonical-0020313321223020-2003302120132331-2011222213231003-3202323103020102-1123121111202101-0133003210203310-3322002322221233-3233113122003031"></a>

## Direct properties — disable_internet_vip / 223200223330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001023130223131-0310200100330201-0323200030020332-3320033003011023-2011011332031311-2222200323003131-1021323220101220-1312303210223012"></a>

## Next pages — disable_internet_vip / 223200223330 / 4

- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2111130133133231-1331101233332023-0113032000223033-1012103203012032-3011031332333030-1132330302303130-0330112330331310-3021223003231021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232220332303303-0121031113021131-0032223130330303-1100020001023103-1121131133312100-0111123203203033-2312031323223103-1222023301211022"></a>

## where.site.enable_internet_vip — enable_internet_vip / 312103103230 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- where.site.enable_internet_vip

<a id="canonical-1111122123111022-3033221332301213-0100023031220221-1221230002220021-0302100212310033-2001123033002000-2331010113112222-2120221221310123"></a>

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

<a id="canonical-1023011101333320-0112011301012123-1010001323003103-1030102000111032-0212302103321100-1300002220232323-3000201310232233-3223121122203222"></a>

## Direct properties — enable_internet_vip / 312103103230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120102303010213-0230120102130220-0212200210112232-3203300110200012-3201110330003312-0213330331100210-2202103321101101-1330022300000100"></a>

## Next pages — enable_internet_vip / 312103103230 / 4

- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-3322223332330102-2330232001102203-2213010010102323-3200103311331011-0100130131321202-0222213330321201-0100121301230212-1331121303010101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113313213100323-0233131322133001-2221321211003313-2201001332003123-0112100202112020-2220022320330013-1212102122032010-1133131211131103"></a>

## where.site.ref — ref / 030221003331 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- where.site.ref

<a id="canonical-2320023130132003-3222300321212333-0111020000010222-2033323012310311-2201013331210103-3033122023121112-1030120122210330-0031121111132021"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

Upstream description:

A site direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1323113113332311-1133303213320101-2032202111002211-2212023023010303-2232123013223232-2311021022310032-1321303310210033-2102121011010322"></a>

## Direct properties — ref / 030221003331 / 3

<a id="canonical-3223132211021303-3233223303220121-1101102203131010-3000103220102132-3133111113201222-0333301320020230-1100211301310112-0231011302000102"></a>

<a id="canonical-2320111310030010-3110220230332201-3031222021331220-3311202012030312-2301032323301331-3323210121001232-0312111301320301-2123302133031101"></a>

## kind property — ref / 030221003331 / 4

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

<a id="canonical-0030100011312200-3032013002233113-3313003313012123-1000310112313100-3223131023210120-2202311112123023-1030330031222233-3311212202002223"></a>

<a id="canonical-2303202103302231-0121112022020232-0232323001100111-1232202212310112-3031223131211101-0211213210211330-3122201312023330-1210213332000121"></a>

## name property — ref / 030221003331 / 5

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

<a id="canonical-0122223323201200-0210313020303120-0111132033302032-0300113022210113-3033033323121122-0033112310031321-2213033101303330-1311301302301023"></a>

<a id="canonical-3310112032113320-2221112302223011-2211323231330020-2203222323102301-0131211322113033-2223110132321022-0303001103032133-1201230201333302"></a>

## namespace property — ref / 030221003331 / 6

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

<a id="canonical-1133311323211200-3113102312303032-2210222003122001-3130233331203020-2103000302023203-2001012123133332-2202013331202303-3210301321330302"></a>

<a id="canonical-2113321202101132-0333233313211113-2313020113032223-3212032012131230-0012331033110001-3321032133220300-2331032320122021-0233033101313310"></a>

## tenant property — ref / 030221003331 / 7

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

<a id="canonical-3211200102122130-3011210131320233-3031132330313121-3032221130032300-1110221212131010-3110133100313223-1323102122201312-3332231020120212"></a>

<a id="canonical-2003033120231020-0223232002012210-2202020102333021-1002330100130130-3120022023230021-3002110122013300-2000310220003000-2303120122233020"></a>

## uid property — ref / 030221003331 / 8

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

<a id="canonical-3022103000121321-2101013222333032-2323321110120312-3031223200010120-3031133202202303-0321130222203031-3231233310112110-1133300320122213"></a>

## Next pages — ref / 030221003331 / 9

- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-2213021221013101-0232220310211310-3320010300110003-3011301313311110-2001301100021312-1233100133010012-1110203313003001-1331010131223321)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1010100312222110-1101123221213010-1320012022131302-2121232311200201-0133002000213230-3132212101210021-0321212330221100-1322311110200013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033210230130022-2121223303312102-1012001023332201-0331302120000311-0120130320133101-0021023222330123-2313301311131233-0223203220122213"></a>

## where.virtual_network — virtual_network / 023232222222 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- where.virtual_network

<a id="canonical-0330112231003221-3223230112200210-2101230313300030-0003230131212010-2032120121313002-1012323320223003-1310000313322201-1010102300232330"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3211032003000112-2033212030001310-0020200002211200-0003010000210033-0020022001301030-0001130332101023-0032101223100333-0331211222202300"></a>

## Direct properties — virtual_network / 023232222222 / 3

- [ref](data-sources--secret_management_access--reference--group-002.md#canonical-2321122003110310-1200023331323221-3113213010023032-0123112030221013-1132302221123030-2113312011101323-3130330322130320-1231001033133300): complete subsection reference.

<a id="canonical-3000333023312203-0332003300132022-3131233012231103-1320003313113323-1232221223212223-0023232333131332-3230113033223320-0101233012021311"></a>

## Next pages — virtual_network / 023232222222 / 4

- [where.virtual_network.ref](data-sources--secret_management_access--reference--group-002.md#canonical-2321122003110310-1200023331323221-3113213010023032-0123112030221013-1132302221123030-2113312011101323-3130330322130320-1231001033133300)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2321122003110310-1200023331323221-3113213010023032-0123112030221013-1132302221123030-2113312011101323-3130330322130320-1231001033133300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210132132020100-0102103001001223-1101222222113322-0303202011200222-1003100131333113-1211333120332223-1123020011011321-2233203323130301"></a>

## where.virtual_network.ref — ref / 320323233031 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-1010100312222110-1101123221213010-1320012022131302-2121232311200201-0133002000213230-3132212101210021-0321212330221100-1322311110200013)
- where.virtual_network.ref

<a id="canonical-1130233211123201-0120130313000322-0111210222311320-2022133233103301-3212131022333332-3131312323313102-2002132023230033-3100210201020200"></a>

Type: `"list"`. Computed.

Reference. A virtual network direct reference.

Upstream description:

A virtual network direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3301223213131003-3213101123001132-2320031023203000-3310001023222032-1320301100310033-1001102223312332-2003000112032303-1313303030003321"></a>

## Direct properties — ref / 320323233031 / 3

<a id="canonical-1213000012233011-3321212101033221-0310320123202300-0102231032000321-2233030002033002-3000201102111311-3123210201300121-0130233331122330"></a>

<a id="canonical-2010313333023310-0313220031211202-2211303310031203-2013010022101321-3110000223022113-1331212020310101-1110131121201012-1033001002110010"></a>

## kind property — ref / 320323233031 / 4

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

<a id="canonical-2310132331232320-3333330011323000-3123213032332310-3121321221202020-0120003012300221-2202133102201030-2012032323000331-0320233321222013"></a>

<a id="canonical-0233101033103200-2222033102201111-1032213322203033-3002221332123322-3121030213011032-3101133302212110-2030011220010010-0332001123203202"></a>

## name property — ref / 320323233031 / 5

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

<a id="canonical-1030131032213201-1113312220113313-3220321110301301-2001003223320010-2013333020133113-2103121320000313-0103010110323003-3033230301331000"></a>

<a id="canonical-1233233010300302-1232303311200223-3232332310220301-3031021102133302-2002110213212022-2103202200020020-0111330033122303-1303312323211123"></a>

## namespace property — ref / 320323233031 / 6

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

<a id="canonical-0010230300222313-0310010122111122-1003033123032201-2211303032310212-2102212132022321-0233321120212131-2213012203021101-3112021032331111"></a>

<a id="canonical-3121031232103022-0332301230220011-0232101022130332-0103200032312113-0222223321002130-2110033202010302-2231131312002101-1100130330232201"></a>

## tenant property — ref / 320323233031 / 7

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

<a id="canonical-1301030132220110-3202201203323322-2132223330211023-2030020002121122-3013220213231130-2222003131220322-2022102101023310-1203001102213032"></a>

<a id="canonical-0302231121220333-2213022231230312-2030231211231312-0221003220210112-3003123210212302-0230221223211233-1201000033222311-1323023312210113"></a>

## uid property — ref / 320323233031 / 8

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

<a id="canonical-2111220100332111-0032011100003123-1120321313030130-1210031011211003-3211122012103301-0133010133223111-2223302233210131-0312011123220310"></a>

## Next pages — ref / 320323233031 / 9

- [where.virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-1010100312222110-1101123221213010-1320012022131302-2121232311200201-0133002000213230-3132212101210021-0321212330221100-1322311110200013)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321001113133033-1112102222030233-1032202013301230-3320013220212001-3000312130312213-1210100012103002-1201221320220100-1332003201132222"></a>

## where.virtual_site — virtual_site / 123230202233 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- where.virtual_site

<a id="canonical-2001033311032022-3000322323110031-3221011302233023-1300031320330201-0112132010323310-2221301210010212-2301013033120303-2032321220111202"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-2233323121123332-3122111211203003-1112130033103302-3213030321120311-1311021033310001-3022230013311122-3133000111120131-3231332201330033"></a>

## Direct properties — virtual_site / 123230202233 / 3

- [disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-2202023202011210-2132033110202311-0001122031213032-3323000030010202-3213121322211221-0000213302332000-0021313210223312-2032331032311300): complete subsection reference.

- [enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-0231020332010130-0212221121331013-2113132202133330-3032233032130112-3322131230211113-0330102332002200-1230031220300210-3233103212332300): complete subsection reference.

<a id="canonical-1333032023122311-3110002120230000-0200021213113010-3121030313210303-2002002201001113-2301010232013211-2101322221203110-1323303303200021"></a>

<a id="canonical-1033321010202023-0333120032100203-0330110123332221-1120222232102211-0302201132022303-0301013113101123-3221121110300211-2122001302323020"></a>

## network_type property — virtual_site / 123230202233 / 4

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--secret_management_access--reference--group-002.md#canonical-1113100123233122-3321100023011022-3203000032000302-1320110201021112-2130013230100010-0213121312010030-0102200202033232-3333202000022121): complete subsection reference.

<a id="canonical-2311120101032311-0120020032313223-2101311330230300-1313213331333200-3020000133301333-0313313030311321-3303210312302320-1231302222220022"></a>

## Next pages — virtual_site / 123230202233 / 5

- [where.virtual_site.disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-2202023202011210-2132033110202311-0001122031213032-3323000030010202-3213121322211221-0000213302332000-0021313210223312-2032331032311300)
- [where.virtual_site.enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-0231020332010130-0212221121331013-2113132202133330-3032233032130112-3322131230211113-0330102332002200-1230031220300210-3233103212332300)
- [where.virtual_site.ref](data-sources--secret_management_access--reference--group-002.md#canonical-1113100123233122-3321100023011022-3203000032000302-1320110201021112-2130013230100010-0213121312010030-0102200202033232-3333202000022121)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-2202023202011210-2132033110202311-0001122031213032-3323000030010202-3213121322211221-0000213302332000-0021313210223312-2032331032311300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111132230332021-1212003203200011-3010130231223130-0012333001323302-1221220332122120-0101113321320031-2322223233221222-2311113332202131"></a>

## where.virtual_site.disable_internet_vip — disable_internet_vip / 113110012302 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- where.virtual_site.disable_internet_vip

<a id="canonical-1132313222220030-3302030131103023-0000322313002232-0221120003000322-2011313122022312-0333303011102330-0132121301002002-0223321212101203"></a>

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

<a id="canonical-1123203223230010-0123020223220221-0132002223021203-3002010110332031-1221322333230310-1213113331002221-1023213130101332-3122200021300203"></a>

## Direct properties — disable_internet_vip / 113110012302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002033120121100-0233130123222312-2310022323002110-0222210012013323-1311323010230002-2213103011123213-3130100201000202-3200333031132000"></a>

## Next pages — disable_internet_vip / 113110012302 / 4

- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-0231020332010130-0212221121331013-2113132202133330-3032233032130112-3322131230211113-0330102332002200-1230031220300210-3233103212332300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021220323131213-2220033010121021-3021101130203000-2002303031302232-0102320032131323-3123102021301232-1223001232003001-2323020022323121"></a>

## where.virtual_site.enable_internet_vip — enable_internet_vip / 333220231001 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- where.virtual_site.enable_internet_vip

<a id="canonical-1120002302323123-0012311313203013-0200032113300311-2030232003001321-3312010021210023-0201002020323002-0113303130112323-0301221133302001"></a>

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

<a id="canonical-0220330001033230-3031011320201110-0232213201003021-2303233000330233-1001012102101021-2302021232233002-3113211101232000-2001013031332133"></a>

## Direct properties — enable_internet_vip / 333220231001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020012322231102-1210201210033222-2200011130330321-1100113203231221-2121320001000011-1313203023030331-2121031022301112-1113003022333100"></a>

## Next pages — enable_internet_vip / 333220231001 / 4

- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)

<a id="canonical-1113100123233122-3321100023011022-3203000032000302-1320110201021112-2130013230100010-0213121312010030-0102200202033232-3333202000022121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203111110032231-2023333113003221-1112103123020222-2311301011010302-0013310110031231-3330010223320222-0011120100333320-0330130230131000"></a>

## where.virtual_site.ref — ref / 331110211302 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-0020132303030220-2213313121331030-0233220021033032-3310301211232111-0200012331323110-2102310201202212-0010320210113313-2031031233002101)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- where.virtual_site.ref

<a id="canonical-1232031000220131-0320210023132201-0233211020113323-3032000231123321-0020213122020321-2221002211321031-0031213203220013-3000022201231213"></a>

Type: `"list"`. Computed.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2212020133200212-3031200122100122-0313110311213000-3202123220310313-1123013033120211-2011213013102321-3002220022202130-0132033100213202"></a>

## Direct properties — ref / 331110211302 / 3

<a id="canonical-2202231031301311-3112211023032113-2302032230320220-0232123232133030-3313130222220100-3313301121023100-2113122112232221-0211030101330320"></a>

<a id="canonical-3020012133020322-0310210213021022-2213313133000123-3320022200231133-3010310010332322-3220332323333031-1321302211020101-3213131203210001"></a>

## kind property — ref / 331110211302 / 4

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

<a id="canonical-3232232030013323-0203121131202312-2322212303200312-3131010003100030-3213023132001000-1003301212231222-1211231122330020-2012012113233330"></a>

<a id="canonical-1022103310220132-1033021023110202-1020222220232112-0123333231210330-2003001231233200-0020221211211122-2320000132233001-1301102221321300"></a>

## name property — ref / 331110211302 / 5

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

<a id="canonical-1021230013230332-2210123030202320-0233211200001322-3301113201322201-0130231201123001-1211120132102222-2231003301111031-2313310131000020"></a>

<a id="canonical-2321033230330132-3221230213211331-2223213212010000-0120101122311322-3112102301100011-1110032313102311-3021001301102321-3010300231021220"></a>

## namespace property — ref / 331110211302 / 6

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

<a id="canonical-0220112000000022-1113232310030133-0333012222320322-3020331312311010-0123103320331032-0223231323123300-3322232031311202-1030330001001103"></a>

<a id="canonical-0320220210012132-3103033001020002-0002103312032211-2112130131123010-1321003222320102-3231100010232111-1332310012020221-0022330301312102"></a>

## tenant property — ref / 331110211302 / 7

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

<a id="canonical-3311022013110010-2133332013330030-2020231331130110-2233232200330123-3131303133133111-0033013113332232-0302321021000013-1023330330013103"></a>

<a id="canonical-2200333210132313-0310110232221211-2200331010312213-2132000313321332-1111130312322022-2020323332223033-1113022201213121-1103322010010310"></a>

## uid property — ref / 331110211302 / 8

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

<a id="canonical-2001302100210023-1310113002111232-1213233231101102-3001313202120101-1202312101233202-3230301310323032-0302310232333213-3001133312101131"></a>

## Next pages — ref / 331110211302 / 9

- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-1313312212212220-1131131232320221-3202033013012233-2322101003330200-0201230310303321-0023002110113211-1221120012012311-1220000103222212)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123)
