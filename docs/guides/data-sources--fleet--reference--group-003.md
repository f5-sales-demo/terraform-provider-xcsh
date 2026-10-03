---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-0331330221032001-3131001023112300-3201121021302030-0120013111022202-1320311100222321-0103122231321123-2230301210103332-1100122033203310"></a>

## storage_device_list.storage_devices.custom_storage — custom_storage / 322011033030 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- storage_device_list.storage_devices.custom_storage

<a id="canonical-2033333111100031-0322201230313331-2103013012013113-1333033103312210-1123011002330201-0221030302130102-3120322000132033-3300302230021112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for custom storage.

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

<a id="canonical-3223233032313322-3010312132232130-0000232010130013-3122010200103212-1120322312303201-0312321220022333-2001030001303022-2003313121121311"></a>

## Direct properties — custom_storage / 322011033030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022331102012200-1321003200321333-2210203001320032-1000100031012121-1003130322212312-2110123012312303-2312000001332123-1232102213101303"></a>

## Next pages — custom_storage / 322011033030 / 4

- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322132202132110-1010122231212331-1131103033010303-1000221031133030-3110010100020211-0100002232332013-1131120113003202-1030013312200212"></a>

## storage_device_list.storage_devices.hpe_storage — hpe_storage / 003200110111 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- storage_device_list.storage_devices.hpe_storage

<a id="canonical-1212112201320032-3131303222211210-1233131001130100-1110001310233102-2232031303300033-2321102313222101-0112011000331022-3321320102323221"></a>

Type: `"single"`. Computed.

Configuration parameter for hpe storage.

Upstream description:

Device configuration for HPE Storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103332220313023-2322132101133111-2221230321113232-1100233231232212-1201311320103013-0020211021201222-2032322232130020-3013031033131220"></a>

## Direct properties — hpe_storage / 003200110111 / 3

<a id="canonical-2113032022331123-3121132312013300-2100213212300022-1200013130113330-0031302223110322-2230213220201123-0303122230312220-1201101323111213"></a>

<a id="canonical-3201230333131332-1000311332111230-3212111131333033-1000021223011031-1312123011202210-3120303231031321-3002313230302331-3233322220032031"></a>

## api_server_port property — hpe_storage / 003200110111 / 4

Type: `"number"`. Computed.

Storage server Port. Enter Storage Server Port.

Upstream description:

Enter Storage Server Port.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [iscsi_chap_password](data-sources--fleet--reference--group-003.md#canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111): complete subsection reference.

<a id="canonical-0103200223000220-0331022103311121-3123010030033321-1020223100323331-3232132232112030-3301130310202332-0130200213032122-1200102022230233"></a>

<a id="canonical-3303001311033311-1303311321313130-0213313030123203-2131031011233012-0200212020120212-0120001113010333-1013230300321223-1323010133012103"></a>

## iscsi_chap_user property — hpe_storage / 003200110111 / 5

Type: `"string"`. Computed.

Chap Username to connect to the HPE storage.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [password](data-sources--fleet--reference--group-003.md#canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130): complete subsection reference.

<a id="canonical-2330312111310331-0223302321131011-1023231020313012-1133002323300110-1122101312331200-3223210202101120-2021030233003013-2132001013313112"></a>

<a id="canonical-3200032013232312-3011012113103221-1010010221333013-2021022211121002-3210031221312113-0303012312012113-3230310121312113-0311103212230200"></a>

## storage_server_ip_address property — hpe_storage / 003200110111 / 6

Type: `"string"`. Computed.

Storage Server IP address. Enter storage server IP address.

Upstream description:

Enter storage server IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1220002021111231-0031223333310000-1200032100200232-2300331102021131-3301131102001101-0111101202113003-0133222033113130-1101322020101033"></a>

<a id="canonical-0012213102001001-1100203300323223-2213233203312001-3311001200000323-3101102122012033-3201030031311003-3131110223202230-1122312202131212"></a>

## storage_server_name property — hpe_storage / 003200110111 / 7

Type: `"string"`. Computed.

Storage Server Name. Enter storage server Name.

Upstream description:

Enter storage server Name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0210031030002121-2031200113300102-2032121230132330-0112121111031321-1000322110223032-1333200011232300-2100113303200201-1120101233131121"></a>

<a id="canonical-3201202202101022-2023010120120211-3032330023023033-2322101123202321-0212032013211232-1311332200232323-2323323123303330-2231322333322333"></a>

## username property — hpe_storage / 003200110111 / 8

Type: `"string"`. Computed.

Username to connect to the HPE storage management IP.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0320312320013310-3331312212300200-3222000202333333-2000011132101111-3112131232311322-3111000013000102-0233020003112000-0333110031232012"></a>

## Next pages — hpe_storage / 003200110111 / 9

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-003.md#canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111)
- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033000033302103-3012231120313231-2311222130103333-3121210202220102-1333101033031223-0320230013011233-0221201132030230-0213000133010021"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password — iscsi_chap_password / 222321001122 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-003.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password

<a id="canonical-3002300233200011-1312213130100100-0232313112203300-2322011203133130-0211323013020113-0310103001321022-1112000003032201-2013223333322232"></a>

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

<a id="canonical-3322200331121021-1332133133022112-2313133212033333-1213003233333202-2210200310222120-3301311201200223-2202321112122202-1212203302213230"></a>

## Direct properties — iscsi_chap_password / 222321001122 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3300021311030112-0210120120101232-2333200121311000-2023123300221321-1110331011131132-1020130220212120-0120200321200012-2011022331022022): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0131300100122032-2130120010221130-3212110320101213-2000021321321023-0222111012211232-2031123113300133-2132113313131112-0232202213321332): complete subsection reference.

<a id="canonical-2211202023310201-0222230213312303-0323132300000123-1300030213133022-0011300111200110-1303301201212312-2121300230111223-0303110233231121"></a>

## Next pages — iscsi_chap_password / 222321001122 / 4

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3300021311030112-0210120120101232-2333200121311000-2023123300221321-1110331011131132-1020130220212120-0120200321200012-2011022331022022)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0131300100122032-2130120010221130-3212110320101213-2000021321321023-0222111012211232-2031123113300133-2132113313131112-0232202213321332)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-003.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3300021311030112-0210120120101232-2333200121311000-2023123300221321-1110331011131132-1020130220212120-0120200321200012-2011022331022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302103301013121-0233022032212223-3032330223033030-0020013301313331-3001122031022300-0132121221113021-2313223020302300-0111232120102301"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info — blindfold_secret_info / 222320322213 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-003.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-003.md#canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.blindfold_secret_info

<a id="canonical-2200102010012102-3010112120320011-2301020213233100-2001031132131231-0231033312202200-0301323321020001-0102031012133103-0112303331001112"></a>

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

<a id="canonical-1320311202130130-3000022022133003-0303313001322222-2023200113110033-3120021310123211-3003112231102323-0123311022130031-3101232330301002"></a>

## Direct properties — blindfold_secret_info / 222320322213 / 3

<a id="canonical-2230321312010112-1123113310200301-0321232033231221-3210000023212000-2200210103300302-3202303131301120-2012122202312322-0103031120220210"></a>

<a id="canonical-1332331310121102-0313032300233013-0011213202221013-0200023023200300-1201122232000323-0113100331300103-3011331233001231-2021002033033321"></a>

## decryption_provider property — blindfold_secret_info / 222320322213 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3311021013020013-0012130313112002-1303311121022223-1213021130111112-3012310201233333-2133330320202120-0132311300012033-2332310002001011"></a>

<a id="canonical-1012332002230120-0311102320101113-3222031123002122-0003322210301120-3332010302131021-0211233232100113-0111020232300210-0002302312210031"></a>

## location property — blindfold_secret_info / 222320322213 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0330322222223333-3003221111202133-1122230220231133-3002202031102032-0113310113001003-2300013021332132-1221033310311322-0300221203332030"></a>

<a id="canonical-3323222110003231-3333133102020310-0300031210121200-3200302001231220-0021231120122032-1221120120200232-0222012321010323-2101023313303200"></a>

## store_provider property — blindfold_secret_info / 222320322213 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3100010211122013-2110000223033022-2131003330012230-2020331322311201-1222233032301013-0200010333330110-1100221200120231-1232220022112330"></a>

## Next pages — blindfold_secret_info / 222320322213 / 7

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-003.md#canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0131300100122032-2130120010221130-3212110320101213-2000021321321023-0222111012211232-2031123113300133-2132113313131112-0232202213321332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203001101202111-2313232333331230-2320301101312232-2303320313120311-2211013211323320-2323001320021011-1013300203001221-2310223030221233"></a>

## storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info — clear_secret_info / 121000321011 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-003.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-003.md#canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111)
- storage_device_list.storage_devices.hpe_storage.iscsi_chap_password.clear_secret_info

<a id="canonical-3211331323113323-1330100002010213-0332211113233321-1223230302031003-1202033221022000-2313033120020202-2020110103330323-3031120321211021"></a>

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

<a id="canonical-2201331011013131-3123023113300012-0312231000112111-2021312101111003-3320202131100223-2122200320230112-0132323301020032-0002101022220122"></a>

## Direct properties — clear_secret_info / 121000321011 / 3

<a id="canonical-0210231030032132-0223203201003202-2233213313222302-2222023011331312-1023032100113211-0310233123311023-1232212111010010-0033303303321203"></a>

<a id="canonical-2323222121000213-3212330300023210-0333302203012222-1110123001120221-0220123303320211-0020133003312020-1332030021213210-3011320212222302"></a>

## provider_ref property — clear_secret_info / 121000321011 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2301121123110031-2113211011232000-3111330203211001-0123301300122330-0003310012112112-0110111010223303-3133200232132302-2030232200101323"></a>

<a id="canonical-2333001200102132-2331013211002122-0213130320320123-3111122101010310-1312233022102333-2130111311012220-1202330110311201-1221312130023303"></a>

## URL property — clear_secret_info / 121000321011 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2121221031000321-1003323223100003-2130131122111213-0332021113030320-0233012213132321-3032000003103133-0330100112103030-0002210111201222"></a>

## Next pages — clear_secret_info / 121000321011 / 6

- [storage_device_list.storage_devices.hpe_storage.iscsi_chap_password](data-sources--fleet--reference--group-003.md#canonical-2011133210003233-3131330213022333-1301302221301320-0331302020103111-1130013021301311-0223101011022203-2333310310111200-1313002103033111)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112002023113111-1333011001330321-3320020132200322-2120211203110100-2211233331113020-3230323033200301-0113110103001123-0201312301023213"></a>

## storage_device_list.storage_devices.hpe_storage.password — password / 333212121323 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-003.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- storage_device_list.storage_devices.hpe_storage.password

<a id="canonical-3022311202213013-0111211232111011-0020331030232013-2120323021100020-2021201033230011-3021113112111000-2021020310302203-0301123201132133"></a>

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

<a id="canonical-2113312203211011-2323310201331100-0303130132031120-1021113100311303-2102133121012132-3120020100112332-0030333213130210-3213302233303103"></a>

## Direct properties — password / 333212121323 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-0232130221121111-1223103201032122-3332221222133123-3012113112033200-0211033023311001-0211332330102000-0201210202331011-3220222021110002): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0130012033220110-0203211303002303-1313210113223103-1201210110022023-1002131232200332-2233130102010010-3132130111120002-0111021202233113): complete subsection reference.

<a id="canonical-0011332232312300-0332222303322333-0120100301032000-0330330100320231-3311100011303033-2303032322120212-2321032203231102-3302133130321321"></a>

## Next pages — password / 333212121323 / 4

- [storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-0232130221121111-1223103201032122-3332221222133123-3012113112033200-0211033023311001-0211332330102000-0201210202331011-3220222021110002)
- [storage_device_list.storage_devices.hpe_storage.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0130012033220110-0203211303002303-1313210113223103-1201210110022023-1002131232200332-2233130102010010-3132130111120002-0111021202233113)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-003.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0232130221121111-1223103201032122-3332221222133123-3012113112033200-0211033023311001-0211332330102000-0201210202331011-3220222021110002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002010210232032-1302201133011133-2131303110331221-0010301302213031-3022033330130232-1311012000011320-0302002023100010-1230330323210101"></a>

## storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info — blindfold_secret_info / 313130303212 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-003.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130)
- storage_device_list.storage_devices.hpe_storage.password.blindfold_secret_info

<a id="canonical-3320020032311103-3213233232313221-3330123202331211-1031232003012013-3032112103330110-1231002223302100-2323103131223302-1223023011021300"></a>

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

<a id="canonical-1020323120100123-2021300102130220-2233121303001230-0233320331233330-0103230102311010-0012210322021203-3003211210301031-2020121301133021"></a>

## Direct properties — blindfold_secret_info / 313130303212 / 3

<a id="canonical-1121212011333300-3332211131221132-2120320030021303-0230321023023200-3230101310131212-1331032231112012-1023022301123030-3200132100122322"></a>

<a id="canonical-2133011232211030-0000102133322010-2232302221113323-1123202300222323-2330110331130310-0013312333303302-3300303200133032-2303101213112102"></a>

## decryption_provider property — blindfold_secret_info / 313130303212 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2122101103100010-3011123211303121-2200132303232310-1032123211213331-3201303312002320-3330132112332320-0001320301103321-1213311022110200"></a>

<a id="canonical-0221030322311201-2202111330111322-0303032020331120-2031221300120212-3012300021001222-1111131333110213-1003110230303222-3233333231121103"></a>

## location property — blindfold_secret_info / 313130303212 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3302011003230232-2331023023303110-3221322311211002-0300130120011320-1231321330100130-1203203113210011-3133121012201012-1311332302303130"></a>

<a id="canonical-3121313111133121-0003302231013000-0222222032020020-3130202211030033-1023010213332223-3212010033301011-2121213113133201-2001301232100332"></a>

## store_provider property — blindfold_secret_info / 313130303212 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0031222133132310-2113313033300003-1033132130333131-3303302003023312-1113213322003020-1101312002100032-2222102311033201-3232231120011210"></a>

## Next pages — blindfold_secret_info / 313130303212 / 7

- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0130012033220110-0203211303002303-1313210113223103-1201210110022023-1002131232200332-2233130102010010-3132130111120002-0111021202233113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032131030231321-1101323201302203-2012132002120301-2322011312120210-2112130030032312-3322303233122202-1010221211211233-1000002310031310"></a>

## storage_device_list.storage_devices.hpe_storage.password.clear_secret_info — clear_secret_info / 112102310122 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.hpe_storage](data-sources--fleet--reference--group-003.md#canonical-3203031320022332-0223333222103200-2303232321210132-1210230200210302-1322120232020302-2110111133231300-0103102022030212-0203300323220111)
- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130)
- storage_device_list.storage_devices.hpe_storage.password.clear_secret_info

<a id="canonical-3000312121032023-3210130231331133-0033213103303301-0011323122303113-3003021012300000-2031213023013232-0320113132120331-0111213313003323"></a>

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

<a id="canonical-3033331002123323-2031010212122311-3332321313013301-1331021331323331-0320220031030230-1301210000002032-0120223233102000-2203131132123222"></a>

## Direct properties — clear_secret_info / 112102310122 / 3

<a id="canonical-1020302012322333-3121201332233132-0001323120030131-3002203233323322-2100110313121102-1100010213123330-3031011030230121-0231122001110120"></a>

<a id="canonical-1332022100000222-2102303012122213-0123030111100200-2230300313001111-0303332022003220-3002033132202123-0203211123111123-2323312113223231"></a>

## provider_ref property — clear_secret_info / 112102310122 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3132233123232231-2020322101010020-1103312120011203-3322330330220202-1331321322322012-1001232302322100-2212310211123330-1233232000101011"></a>

<a id="canonical-2330103001022000-1101112031303232-3111321231003321-3112021131313023-2312123001322200-1203103310112311-3132132132332010-2012120310321213"></a>

## URL property — clear_secret_info / 112102310122 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3211132331001332-3110233311332310-1022131030102303-1213122320101232-0210301133233222-0100232330232111-2010020033021020-3221111100301103"></a>

## Next pages — clear_secret_info / 112102310122 / 6

- [storage_device_list.storage_devices.hpe_storage.password](data-sources--fleet--reference--group-003.md#canonical-1121120011223221-2200323302012231-1222012300231133-3120111033031022-0210303130321210-2320103330323021-3103002202230203-3131013302202130)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213100112130102-2203203113213102-1332112333112032-0331322222010132-3311112113203011-3111022322103013-0131120033310010-0202211103132132"></a>

## storage_device_list.storage_devices.netapp_trident — netapp_trident / 311022130312 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- storage_device_list.storage_devices.netapp_trident

<a id="canonical-3212322031313002-0000103303323330-0201110200213003-0303311233302032-2033023220010002-1102111001222103-1310010210112302-1100231333201300"></a>

Type: `"single"`. Computed.

Device configuration for NetApp Trident Storage.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-backend_choice": "[\"netapp_backend_ontap_nas\",\"netapp_backend_ontap_san\"]"
}
```

<a id="canonical-1312323121011031-3322113301010200-0023232010103130-1022103132200013-0221301300101100-2033011002113121-2301310311323013-1133302112122313"></a>

## Direct properties — netapp_trident / 311022130312 / 3

- [netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133): complete subsection reference.

- [netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330): complete subsection reference.

<a id="canonical-2213112111322220-1022013213202330-2233320033313100-3232112210213101-1121220333133122-3013003331321002-0203033122132232-1313130200301232"></a>

## Next pages — netapp_trident / 311022130312 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332030332211201-3023203000310230-3300101323020222-2200022311000033-3021312220123121-0302303031223210-2222013212021310-0002302333131330"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas — netapp_backend_ontap_nas / 100100300033 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="canonical-3233032001002013-3021100023130033-2100003322311201-2203310113201001-3033030102202132-3321311003001223-3130331230000323-2312111022010011"></a>

Type: `"single"`. Computed.

Configuration of storage backend for NetApp ONTAP NAS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

<a id="canonical-3123121220110301-3033031023033211-0330012112333003-3121023313113302-2110222210333000-1013322330301112-0020032100320002-2020323323313232"></a>

## Direct properties — netapp_backend_ontap_nas / 100100300033 / 3

- [auto_export_cidrs](data-sources--fleet--reference--group-003.md#canonical-3113120030020313-1132100122323130-1130312112201311-3233031212003201-2311312103330310-1001203133331302-1223332101303331-3202012131011033): complete subsection reference.

<a id="canonical-3311021333011023-1313302203121110-0231211011221321-2202320120020030-3322121330110331-3000203122312210-1312112302000000-1301333003313022"></a>

<a id="canonical-1013111133012210-3213013130100120-2221213320303020-2200122322032100-3120023322022230-1311132320212112-0330200013311002-0201223221221213"></a>

## auto_export_policy property — netapp_backend_ontap_nas / 100100300033 / 4

Type: `"bool"`. Computed.

Policy configuration for this feature.

Upstream description:

Enable automatic export policy creation and updating.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2101013322032001-2111233131211122-2122232203003020-3310111113112320-1323100332111232-2300112231213101-2211103221110302-1232313020120213"></a>

<a id="canonical-3330321302223221-2230313023110131-1111011002032201-0333112213332033-3133131322331201-1301020111012110-3103201202011110-2223213031023220"></a>

## backend_name property — netapp_backend_ontap_nas / 100100300033 / 5

Type: `"string"`. Computed.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Upstream description:

Configuration of Backend Name. Driver is name + "\_" + dataLIF.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 50,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2103232230102223-3323030322231200-1221320222113322-2023003022323121-0123030120020203-2233331121302110-2310102011310211-0213021311311103"></a>

<a id="canonical-0133312202033221-1222210100311322-1210201131313300-2023122111021220-2223222012013210-3303220301001223-3231331123213210-0332133333102012"></a>

## client_certificate property — netapp_backend_ontap_nas / 100100300033 / 6

Type: `"string"`. Computed.

Please Enter base64-encoded value of client certificate. Used for certificate-based auth.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](data-sources--fleet--reference--group-003.md#canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213): complete subsection reference.

<a id="canonical-3201310203302233-1310321113133113-3220230230131001-1323133101132221-1101223120032303-3232223132320012-1133311011203120-2021021131200121"></a>

<a id="canonical-0123000220103210-0310130331312310-3313110221210232-0202100032022220-1333213211331333-2301101203132120-3113322332322330-3111111203033103"></a>

## data_lif_dns_name property — netapp_backend_ontap_nas / 100100300033 / 7

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1300031232220123-2121120221113122-0100131301001201-3233203200203222-2132300030031322-1110101212200211-2112001303112210-3201111100211210"></a>

<a id="canonical-0010022002131102-2202303033103031-1203102122311130-2212001220332223-3212030131231031-1030033221033011-3120210003012213-3012123211301012"></a>

## data_lif_ip property — netapp_backend_ontap_nas / 100100300033 / 8

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3112233010300131-3311231210010230-3321020232322120-3002100133233330-0010130122031233-2132333033322020-0203333103333101-0110330022112000"></a>

<a id="canonical-1021302333001120-3120123210213121-1001320101232031-0113130213223323-1003210100221021-1333022320030312-2220203102103311-3302033221010102"></a>

## labels property — netapp_backend_ontap_nas / 100100300033 / 9

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-1200312103123212-2013113231211003-1032211132231000-2223232100023021-1310200200111021-3011230032231100-1223313202000131-1301113100110131"></a>

<a id="canonical-2200011031022030-0311203312020122-3332100010333112-3020311121330012-3111121100210123-1003321011310000-2022200002200321-3330010331103210"></a>

## limit_aggregate_usage property — netapp_backend_ontap_nas / 100100300033 / 10

Type: `"string"`. Computed.

Fail provisioning if usage is above this percentage. Not enforced by default.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2002322111030132-0312100202112100-0333302021010111-3332103310123103-2320201302201233-0021032131100032-2212331323112120-2121010030133123"></a>

<a id="canonical-3201000230213230-2110302101310132-0213020202122122-1221032020032020-0333110000321203-0010010213133000-2223211231012323-3131103110113020"></a>

## limit_volume_size property — netapp_backend_ontap_nas / 100100300033 / 11

Type: `"string"`. Computed.

Fail provisioning if requested volume size is above this value. Not enforced by default.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1331031333022103-2130301132322311-0022011222311023-0013203312223030-1133302231010320-2212001202321221-1200032222220021-2012223323332023"></a>

<a id="canonical-3212100330132133-3231303123032130-1111231133022331-2323011123011131-1203320212301112-2220333312211332-2111121103223010-2130113312301330"></a>

## management_lif_dns_name property — netapp_backend_ontap_nas / 100100300033 / 12

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1011201032102120-3333111222212223-2113030202133313-3310311310212011-3323201333203103-0103320133001002-3002031203003321-1030333131012233"></a>

<a id="canonical-0122023203110012-0101112220212322-1203033131232033-2013330110312330-2030012233001113-0323201312221003-1331110201313333-3132100113121323"></a>

## management_lif_ip property — netapp_backend_ontap_nas / 100100300033 / 13

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0233011322112312-1200231111020212-3313213022310210-1203203122002003-3302213232333303-3330132312330322-0031310221312321-3223201033130311"></a>

<a id="canonical-2130300321330323-2310312301331101-3231230230323213-1233233110020332-3100320112320031-0321323223233113-3332230102310110-0333301012103132"></a>

## nfs_mount_options property — netapp_backend_ontap_nas / 100100300033 / 14

Type: `"string"`. Computed.

Comma-separated list of NFS mount OPTIONS. Not enforced by default.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [password](data-sources--fleet--reference--group-003.md#canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200): complete subsection reference.

<a id="canonical-0303213303210031-0212122212133211-1201332213210201-0201203231300110-3203231311012210-2102330313322000-1200130200322313-1130111223323011"></a>

<a id="canonical-3121322322131231-3131230111002223-2113332121122203-0020023213013121-2022222002032013-3201103123103011-1112311133110003-2232013222221230"></a>

## region property — netapp_backend_ontap_nas / 100100300033 / 15

Type: `"string"`. Computed.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [storage](data-sources--fleet--reference--group-003.md#canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133): complete subsection reference.

<a id="canonical-3033301213032113-2330323113102132-1000203032311032-2323202020233120-1010203222202022-2021103123120212-2100323331303323-0021201011210132"></a>

<a id="canonical-3030313202110132-2211320002322123-1203112133331032-2032022012303321-3011002201010103-2101130120333330-3130320033000201-2330231121222110"></a>

## storage_driver_name property — netapp_backend_ontap_nas / 100100300033 / 16

Type: `"string"`. Computed.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-3020113222210123-0110200213132011-2223211130012322-2012122101330310-3113333132213000-2010230202230113-2132002220032311-3311122111210300"></a>

<a id="canonical-3020221233303032-3330001003121230-1132233211333022-1332122200121220-3121230122120222-1113332212033322-2213003020312212-1330111320023200"></a>

## storage_prefix property — netapp_backend_ontap_nas / 100100300033 / 17

Type: `"string"`. Computed.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0013110023120213-1130221032313331-2120200112300133-0133222300011102-0102233332313311-3011303212201110-2133230222313002-1101000330221321"></a>

<a id="canonical-2123232111202330-2002130311220003-0023000212320211-2310310220021131-1312032020001101-2330110010200322-3002331312200330-3231102030033012"></a>

## svm property — netapp_backend_ontap_nas / 100100300033 / 18

Type: `"string"`. Computed.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2023303332033200-3230202031222033-3211200213300120-2133011200121013-2003130003031133-1213323202210112-2111330020213302-2000031223311021"></a>

<a id="canonical-2311233100002112-3110012330130103-1130310113100232-3013120123231032-3112303130030001-0223231331323032-0223322211012001-1303320023222231"></a>

## trusted_ca_certificate property — netapp_backend_ontap_nas / 100100300033 / 19

Type: `"string"`. Computed.

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

<a id="canonical-3300131302300022-3302103033201332-2222303322201123-3010020033220222-1300332321122311-3030112010203203-2031222223013202-2132032023303120"></a>

<a id="canonical-2211100200213201-0100110131120331-1322132103323222-1130233122112120-3112111311211213-1012332100211000-1311221220011111-3032000031302222"></a>

## username property — netapp_backend_ontap_nas / 100100300033 / 20

Type: `"string"`. Computed.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1031002301120210-1211113332013001-0221212030021230-1213323311322101-0103111130010301-1132312332210131-1123103220103002-1121031232101023): complete subsection reference.

<a id="canonical-0333122333300221-0323132030230121-0130000211230320-2210100113332222-1132113323212323-3033121322011303-3300330120010031-0212031230213223"></a>

## Next pages — netapp_backend_ontap_nas / 100100300033 / 21

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](data-sources--fleet--reference--group-003.md#canonical-3113120030020313-1132100122323130-1130312112201311-3233031212003201-2311312103330310-1001203133331302-1223332101303331-3202012131011033)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-003.md#canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1031002301120210-1211113332013001-0221212030021230-1213323311322101-0103111130010301-1132312332210131-1123103220103002-1121031232101023)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3113120030020313-1132100122323130-1130312112201311-3233031212003201-2311312103330310-1001203133331302-1223332101303331-3202012131011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313102023221201-2122210233131100-2222202311012331-2321123100022330-3203011101303233-0223202220110230-0301001330133301-0110233121113332"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs — auto_export_cidrs / 012010311201 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs

<a id="canonical-0022021131213113-0121021030013132-0221122233332210-2001200320013331-0212121130300121-3332023312033233-2220321021010203-3013013210223102"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3320103332012211-2222031032332233-1232320103101302-2030320131301320-2112003213213232-3210133023231313-1003022030133133-2031301013030230"></a>

## Direct properties — auto_export_cidrs / 012010311201 / 3

<a id="canonical-0201223101123222-1202333220121203-3232303023011121-2122013033113003-2021032010030002-0133012131013210-1312022003110232-0133313110100110"></a>

<a id="canonical-0120102231020023-3120032203200231-2220213312121013-3313002131211022-2103100131002330-1032111010213010-0001130230230223-0330031012210100"></a>

## prefixes property — auto_export_cidrs / 012010311201 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3013230003020203-3312201301310332-0232200201303103-0003222212032300-2100310030000330-3032122321130100-2110233112312111-1021203103321020"></a>

## Next pages — auto_export_cidrs / 012010311201 / 5

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300021103101101-0020131211032110-0221220321311331-2312130031311312-0311031321211323-1230101210112331-3203021102232303-0112132033230011"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key — client_private_key / 313000131120 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

<a id="canonical-3101033301333120-2110302232130213-0201321203213322-1123313312020232-2312120120131001-1213131231233120-3222221121100113-0203332022300002"></a>

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

<a id="canonical-3300012133132223-2212012233303322-1132111231312300-1113231231231112-0023110012030110-0112133320000120-2113122233020203-2120211113233131"></a>

## Direct properties — client_private_key / 313000131120 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3301232302303033-1230110300113201-1202000003222232-1202332230320113-2203001331333023-1233213001321013-2213302102102121-1231313100210302): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-2223120110212013-1333023221011333-2320332030300203-0020211013323313-1332233000122311-1212112012322013-1331203213320121-2312002122022220): complete subsection reference.

<a id="canonical-1032101102310011-1221231223131230-3211122312002012-3111120333313322-0202302020122322-1112003230320311-2233121221233111-1003220103132123"></a>

## Next pages — client_private_key / 313000131120 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3301232302303033-1230110300113201-1202000003222232-1202332230320113-2203001331333023-1233213001321013-2213302102102121-1231313100210302)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-2223120110212013-1333023221011333-2320332030300203-0020211013323313-1332233000122311-1212112012322013-1331203213320121-2312002122022220)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3301232302303033-1230110300113201-1202000003222232-1202332230320113-2203001331333023-1233213001321013-2213302102102121-1231313100210302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331330021110120-3210220323332323-0003212223222300-0313301332301330-3031212331220132-1000323022012300-3213032130030223-0321030102200121"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info — blindfold_secret_info / 230200311321 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info

<a id="canonical-3232222332110310-0311213202021101-0301110302122331-3033101031021033-1331101132212330-1122010322001030-2332123223120310-2121001123100112"></a>

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

<a id="canonical-3103303310113203-2302000312001311-1021210123110132-0233203012003333-3230331011311312-2220133300230310-2212203300021300-1321030201323001"></a>

## Direct properties — blindfold_secret_info / 230200311321 / 3

<a id="canonical-1113011332230003-0101021320122313-0332011133010102-0000332100002300-3013311023122312-1011131323001120-3021100233201120-1320003100322021"></a>

<a id="canonical-1211030121302233-3311111301021133-1020032222120320-1321300212003303-2310231123210103-0221331322123220-1011131033233331-2201231003033003"></a>

## decryption_provider property — blindfold_secret_info / 230200311321 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0301002221330110-3302011030032031-2132313110130202-2103320123121211-0012332012220302-0221111222123200-2222131303213131-3032030020030022"></a>

<a id="canonical-2032111231320121-2023221201232211-1112221330232121-2202302321131332-3200010201333110-3000030002302223-3231111002011312-2200230232221000"></a>

## location property — blindfold_secret_info / 230200311321 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3100012231103323-3320312112013330-0130221220221231-3333103233121103-2013300123030021-2231012011201111-2203023333232320-2331321233023020"></a>

<a id="canonical-1033311131032310-3100103310213220-2131320002120031-2033232030302302-0113223103102323-2030132112022103-3312302102123132-3112020333013122"></a>

## store_provider property — blindfold_secret_info / 230200311321 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0022321122133233-0303133123200203-1120223322223301-3130012112321102-0330320132320210-0323300110030102-1203300001323121-2032003313332103"></a>

## Next pages — blindfold_secret_info / 230200311321 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2223120110212013-1333023221011333-2320332030300203-0020211013323313-1332233000122311-1212112012322013-1331203213320121-2312002122022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101201011310322-3230122120331212-3122210113012311-3300321121033213-1223131303210003-1013331210013010-0231031032322033-1102232310212003"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info — clear_secret_info / 022311011132 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info

<a id="canonical-0110122013122001-3130103312010112-2310311132223230-1220030100313323-0223300111121000-1121113022220112-1300110003113121-2021133111233333"></a>

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

<a id="canonical-2103120333033113-2001330133310231-3331321032000202-1121221101030021-1000133303322323-1131011303012021-0323121110222032-3320213120133231"></a>

## Direct properties — clear_secret_info / 022311011132 / 3

<a id="canonical-2120012311222112-0020313220110200-2122130123030221-2132200101303003-2301212202013202-2203333310130312-1031102110033212-0010023203032311"></a>

<a id="canonical-3000233020132010-2210312023200030-3323233110020013-3320312231031303-0231211201221023-3012133232001122-2312200302123120-3203131013103002"></a>

## provider_ref property — clear_secret_info / 022311011132 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1202012022133132-3331023301021132-1130332333233021-3231331213303023-2022201310100312-1312000231322302-1133211110320222-2033103231323111"></a>

<a id="canonical-3123001020012203-1201311200230133-2000012113311310-2320120031313113-3220132200213321-1010312102013311-3313330112311210-1311133301203311"></a>

## URL property — clear_secret_info / 022311011132 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0311312030031332-1221123210233202-3030101233112300-3321011123212230-0011310302201212-2201021322022102-1101333002120030-0113203322301230"></a>

## Next pages — clear_secret_info / 022311011132 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](data-sources--fleet--reference--group-003.md#canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330311332130303-1323013323033221-1313013203030000-0030201003110203-1131233112220010-3302011001000003-1313130003301320-1200123333100222"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password — password / 133320233110 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password

<a id="canonical-0132202332301213-1211220100201111-3230100100111111-2020130003211121-2210003212200321-1032201202230332-0222201302310333-1022201101302333"></a>

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

<a id="canonical-3301032232011302-2202321001033101-0010110101130200-1013230300211001-2013103002030320-2100301311031210-1203233332131303-3312202113211211"></a>

## Direct properties — password / 133320233110 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-0133033201021033-3323220103331213-2133010000231320-2321212112323001-1222222102230011-3003012022312013-2000310022310102-2122321232321121): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0032221332011333-1012012111123231-0332010233101131-1200212030322313-2122331111212210-2033013012133301-3003333000211101-1032300121212131): complete subsection reference.

<a id="canonical-1323123302321203-1111330310213002-1313033320213110-1123201102211231-1232220302130023-3112103121321232-3102311213010120-2033120222102213"></a>

## Next pages — password / 133320233110 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-0133033201021033-3323220103331213-2133010000231320-2321212112323001-1222222102230011-3003012022312013-2000310022310102-2122321232321121)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0032221332011333-1012012111123231-0332010233101131-1200212030322313-2122331111212210-2033013012133301-3003333000211101-1032300121212131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0133033201021033-3323220103331213-2133010000231320-2321212112323001-1222222102230011-3003012022312013-2000310022310102-2122321232321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010220031132233-0331033303321223-1030312112213213-3003003111303132-3130323131200000-0021220211321211-3123002313102311-0122231021333330"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info — blindfold_secret_info / 320322322233 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-1200002210003303-2130111011102030-0212312132101323-2011331010232132-0101300121102223-1331322222301211-1000223300021112-3013202110201232"></a>

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

<a id="canonical-2213220003323233-3210311010000022-0230212102331212-3221300223222231-1321310100321021-1302031302201232-3033031230101022-1111010121312312"></a>

## Direct properties — blindfold_secret_info / 320322322233 / 3

<a id="canonical-0100113300020123-0100110013011301-3120213311111022-3133113232100213-1011311002121301-3200101130032130-3323113000131310-3012010110203310"></a>

<a id="canonical-0322031332111112-3213121122022311-0013312012100320-1312133000021130-0321231123222022-0232302231320131-2233311211131120-1131101223210220"></a>

## decryption_provider property — blindfold_secret_info / 320322322233 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0210030313301320-2002111031101330-0323200101121213-3000203212210213-1010113023221202-0301321303123132-0021200201222002-3302303032212303"></a>

<a id="canonical-2033332303023103-1213223120030212-2032021131103323-3131022212312313-2132010301032230-3220333001333332-1320130333233212-1322231320332213"></a>

## location property — blindfold_secret_info / 320322322233 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2223110010330200-1100202321121202-3130012333133333-2223121303123032-2032300010310231-1022300121110210-3322301101312001-2201321013220302"></a>

<a id="canonical-3112120213311310-2121030133021102-0203122021223203-1233233122320001-2330211233213321-3101012320010010-3001310323221101-3312033100313213"></a>

## store_provider property — blindfold_secret_info / 320322322233 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1022033233310201-1221012232000201-0212010121303113-2330111132003003-1223123303223212-0311212213232303-1031020110020003-3202113200001212"></a>

## Next pages — blindfold_secret_info / 320322322233 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0032221332011333-1012012111123231-0332010233101131-1200212030322313-2122331111212210-2033013012133301-3003333000211101-1032300121212131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213022333033222-1221121133201110-1120021230201120-0030311302332022-0222020233313303-3112201030132012-2220110212031003-3303322210230230"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info — clear_secret_info / 231022010102 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-0111321322003023-2003101110132012-1020311222322113-0023233131322123-1033003033033120-2222312032031333-0020132101133010-1030112320321033"></a>

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

<a id="canonical-0300223323322211-2220301212023101-0120021031231103-2332103301013120-2011023030330312-2232231011312310-3011210022110211-2030313230332231"></a>

## Direct properties — clear_secret_info / 231022010102 / 3

<a id="canonical-3032211231103322-1122130200212330-1000102211010010-2223001221300020-1101122220201322-0331330203202022-3331312230002120-3300311023302221"></a>

<a id="canonical-0033011133110220-2012223221112033-2221001011232312-0331020233021223-0303030310132131-3122033103230102-0111112111023021-0033031113313130"></a>

## provider_ref property — clear_secret_info / 231022010102 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2203002033230233-2132033313220301-2121202301113222-2132233111133222-2232111130120130-2030321310011122-3321030120201101-0321331310132022"></a>

<a id="canonical-2330120102310222-0031220231103103-0101123101212311-0221113101322111-0201002331211321-2210210322302201-1233132310310232-2131100222302301"></a>

## URL property — clear_secret_info / 231022010102 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1221202311112231-2230011111112332-0223111020002300-2003022021021332-1232112233113203-1203320313210233-2212300311233100-1330112213222333"></a>

## Next pages — clear_secret_info / 231022010102 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](data-sources--fleet--reference--group-003.md#canonical-2111022122333231-1020223333303033-3133230013103111-0100210323332213-1002022333132012-2031131203001201-3222003120130130-2112332322021200)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202122020233303-2322302002233202-1101332233211011-1010122211000112-1132321113100023-0313111123200213-3103321200210331-2121021033201232"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage — storage / 222333130020 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-1022312031011321-1112330310300202-1221302332301303-1322221323332033-3120203020001121-3120320311023210-2031222121231101-3003213333303313"></a>

Type: `"list"`. Computed.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2233003103332321-0013112110302332-1121331303323031-1301210330202130-2121212310132120-2011302221300010-0013112213021330-2310230213331130"></a>

## Direct properties — storage / 222333130020 / 3

<a id="canonical-0101032211001230-0130131011011021-2111232232033300-2212322033102233-0321323331032310-2013021133221113-2321102210130021-0021033123133030"></a>

<a id="canonical-1323012310222121-0300132210310103-0212230202321212-2331233121113313-2132013212323131-3030000021230023-0331222123020020-0030300100003113"></a>

## labels property — storage / 222333130020 / 4

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1211033321203322-2132121023112301-1231232232311122-2110303320220002-0122310010221301-3323203132223102-1121213222001002-1103112012220111): complete subsection reference.

<a id="canonical-2201031100230100-0021012102011032-3231200133001101-0333231311131001-0013002132101011-1020300301313120-0120000203321310-0021320030100002"></a>

<a id="canonical-3232222031331120-2123010301302111-3233022110020311-2301133131010303-2133333300220311-2233012121223233-2020112132121213-3113131201013320"></a>

## zone property — storage / 222333130020 / 5

Type: `"string"`. Computed.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2030123021112130-2120110030201032-1322232211302320-0111332110321331-3332322102200120-2133201203103320-3300000303100032-0112110120333110"></a>

## Next pages — storage / 222333130020 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1211033321203322-2132121023112301-1231232232311122-2110303320220002-0122310010221301-3323203132223102-1121213222001002-1103112012220111)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1211033321203322-2132121023112301-1231232232311122-2110303320220002-0122310010221301-3323203132223102-1121213222001002-1103112012220111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103020130300030-0302100303112122-0331223022222303-1232112330133103-0113303102011231-1100203022113001-0332203122033200-1231011200232122"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults — volume_defaults / 221133311303 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-003.md#canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-2222002033211122-2322213312112112-1212313101200013-1303033333133103-1103301233211120-3100031131002033-3100313211110333-2300111231100331"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

<a id="canonical-0230203120002021-2223101203301121-1203033130200033-0311333210211011-0011031131030110-1200013100021013-3120202320113121-1312013223033312"></a>

## Direct properties — volume_defaults / 221133311303 / 3

<a id="canonical-0121302320001102-2021300120212030-3232232211113020-3023120010331133-2122112321210032-0210110013230330-3323212202132323-2230010310120022"></a>

<a id="canonical-2100111231130103-0223330312220213-0011310113000120-0021132022223213-3110333132202203-2302221023223031-2120200120120131-1022331120001020"></a>

## adaptive_qos_policy property — volume_defaults / 221133311303 / 4

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1132133313322022-3001320213301011-3332122233212030-3110000010213223-2133030313031231-3313110123103231-3121020203001223-3220122111002033"></a>

<a id="canonical-2212310112131130-3320011102220202-2211301232320110-0220130312200221-1110120102120021-0302022230103212-1322000232002131-2003332312223022"></a>

## encryption property — volume_defaults / 221133311303 / 5

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100211330213101-0233201122303301-0130022000113311-3303012320122001-0001211130120130-0123311223001302-2112113322101212-1000301133002033"></a>

<a id="canonical-3332203333233132-1030130000012312-3123212312110121-2322211011010333-1211031101203022-0020230332221321-2010331033230222-0033232303032220"></a>

## export_policy property — volume_defaults / 221133311303 / 6

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-3113121113310310-2213300013020110-3121130113032000-0201301203122212-1202103110001302-2031103010102301-0200203322033132-3030200300000132): complete subsection reference.

<a id="canonical-1132012000333113-3321133321101222-1202123300330020-0310012332001202-1000310003122023-0211100030311113-1010221003131201-0013020212000023"></a>

<a id="canonical-3123122003212321-0330210112000133-2130333002330330-1213202012212031-3030132221233112-3202020022033333-0313112102302121-1213310130331223"></a>

## qos_policy property — volume_defaults / 221133311303 / 7

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2020131220322220-1322011001002223-2313100103112230-1122130101120220-3132130003200012-3201232322020110-1110231111333032-0132021223020332"></a>

<a id="canonical-1301322202210022-2020113123320022-2311301202233022-1331011121110311-3121131200320323-0121130323033132-2023231122010130-0320032010000002"></a>

## security_style property — volume_defaults / 221133311303 / 8

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2323110203001112-2012323102131010-0301302021133110-0112111313120201-2233223101232301-3321203000113031-0101112233231302-1033023330300301"></a>

<a id="canonical-0100220221013003-2303203320011333-0000113331121302-1020320022313231-0100102112213313-3123033000303101-3221021030202020-0320300223302111"></a>

## snapshot_dir property — volume_defaults / 221133311303 / 9

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1303122323302033-3131333030232321-2313002231333131-3311131001220301-1210110231203220-1301012230123221-2231000202122122-0303032002230331"></a>

<a id="canonical-3030211030210130-2212330032232012-2220211331023321-0220221223111301-0002112013023123-0211320131313102-0333231222003031-3132011310312121"></a>

## snapshot_policy property — volume_defaults / 221133311303 / 10

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1312131131123032-2020120012332023-3102333320300113-0320221030020130-2012111112110103-2020313121010311-1031310022310111-0302223221110023"></a>

<a id="canonical-0003021301333111-1110031013013322-2101212300310122-3311130012123021-1221211020230200-0233000132311002-3010310102112030-3300201011320112"></a>

## snapshot_reserve property — volume_defaults / 221133311303 / 11

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3131013312012312-1221200322320120-2103233022302222-0330201200313313-0322313102201012-3331202100112033-1202232021301002-1010221033233211"></a>

<a id="canonical-3213223321011332-3230320002313323-0301022302000331-0313202012113122-0033122320300313-2120312111222201-1320132213133030-0133032202102213"></a>

## space_reserve property — volume_defaults / 221133311303 / 12

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-2222231031120220-3323023000012211-2310110332110121-3110130203033022-3113023330310310-0230121012100103-1223312211323302-3131220302202021"></a>

<a id="canonical-0121332311201113-3022021223311201-2100202132102303-1232222222033133-3333000131203130-3312300323031010-2310120233012013-1103102203203002"></a>

## split_on_clone property — volume_defaults / 221133311303 / 13

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2003323003031210-3331110123011303-0333223000000320-3123133313300202-3102212001210132-2330102220330023-0210032300113001-1011213222202033"></a>

<a id="canonical-1000301313131032-0230210321022213-1312100102211022-2100023032310323-1012220303301121-2333221330323010-3100123102203313-2333110301123113"></a>

## tiering_policy property — volume_defaults / 221133311303 / 14

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2331313110233112-1333312133031210-2332203033231031-2120132011321132-0130231222231221-2010210021031212-1333212213203222-3210222231310203"></a>

<a id="canonical-1002111102130113-0203203012131032-2110110121330311-0333331003000232-0030121011013101-0320112123311313-0132133101321203-3320230331013233"></a>

## unix_permissions property — volume_defaults / 221133311303 / 15

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3302022113331101-3230111031300311-1233331302113311-2222311001032000-0230222222133303-3323202023100312-3222113022102300-1123311033123221"></a>

## Next pages — volume_defaults / 221133311303 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-3113121113310310-2213300013020110-3121130113032000-0201301203122212-1202103110001302-2031103010102301-0200203322033132-3030200300000132)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-003.md#canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3113121113310310-2213300013020110-3121130113032000-0201301203122212-1202103110001302-2031103010102301-0200203322033132-3030200300000132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020333223130131-0121003100130232-2111311210212012-1033112320222201-0132111012220332-0323130311200201-3313313132022100-2330121012032222"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos — no_qos / 222002313121 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](data-sources--fleet--reference--group-003.md#canonical-2110120020013122-0302220310310103-1003013222131120-1301111323321323-2112112320103320-0131001012111321-0012200230303231-1111121213230133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1211033321203322-2132121023112301-1231232232311122-2110303320220002-0122310010221301-3323203132223102-1121213222001002-1103112012220111)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-0332332220231030-0323322002200213-1231110312111111-0121233021032223-1110300032110103-3303230002322032-0330230022121001-1012200231021333"></a>

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

<a id="canonical-3113123033323130-3132121233021311-2022001012121302-1013300111002132-2021123223300233-0101011111202103-3120011331103230-3330300313322232"></a>

## Direct properties — no_qos / 222002313121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213202030011220-1103300201130133-0302220231020221-2303232232000332-0121033231111232-0330002233233003-1131321010203213-2112221120121222"></a>

## Next pages — no_qos / 222002313121 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1211033321203322-2132121023112301-1231232232311122-2110303320220002-0122310010221301-3323203132223102-1121213222001002-1103112012220111)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1031002301120210-1211113332013001-0221212030021230-1213323311322101-0103111130010301-1132312332210131-1123103220103002-1121031232101023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122230031200031-2210230232011312-3223111200312003-3013300201030011-3200321133012323-2212001213111101-1101313031222222-2223113121102100"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults — volume_defaults / 001312233003 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-3120220122313122-3321001102020332-2132201023202013-1001003321001101-2333303320313313-3102211332313313-3031321010311233-2122233011223210"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

<a id="canonical-1101101312120103-2103123113103233-3213010110030021-3123000200331110-1201202222332333-3132113001132223-2331111323332013-1130102311231003"></a>

## Direct properties — volume_defaults / 001312233003 / 3

<a id="canonical-2223031010002302-3101233322331303-3301000120232131-2110231230033001-1132130011332202-0320001313233201-2110300210201133-3200122100132003"></a>

<a id="canonical-3032100331101020-2231101020000010-1211223200303031-3110300133311113-0123202110313210-3121030233231013-1020111020001312-1303020002220313"></a>

## adaptive_qos_policy property — volume_defaults / 001312233003 / 4

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3110212201321300-0333001203312101-2013211131100310-0311233112002320-2012011200131223-1123001112100212-1003201112320003-1033123020110010"></a>

<a id="canonical-3322312022201122-1021333120122202-3203323121011011-0303031312330211-3232003330300123-0233301202101131-0113222111312210-3220012010130001"></a>

## encryption property — volume_defaults / 001312233003 / 5

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3000221123203222-3300232012131211-0233002213021312-2333111003222213-2230223123232122-1202100011113032-0012301301021111-1120131100021221"></a>

<a id="canonical-2230311310222303-1122300133023111-0330031102203000-3331302210332222-1313212230123031-1013000311030310-2233130211112102-1131100303103120"></a>

## export_policy property — volume_defaults / 001312233003 / 6

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-2201230221322113-3100212201331313-2122002222033101-0232333201101302-2223232110333230-1323231320121001-0133012121022011-3011221133303100): complete subsection reference.

<a id="canonical-2213300111320222-2003130111233203-0200002221303123-2101201300102232-2022311333321233-2101331132101320-3330022310311023-0233122032330303"></a>

<a id="canonical-0303313110103001-1123022302203020-3030020313200213-1113132111302202-0020200122121101-2222331323020131-3322321223321131-3202200202232000"></a>

## qos_policy property — volume_defaults / 001312233003 / 7

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1332011221032211-3112212301230301-0032113121023101-3102202022302112-2313110212013233-0122213221033310-2222220231013232-1101123332013023"></a>

<a id="canonical-0012112220213220-3311033212020310-2202230201202110-0220202131201010-0110223130313303-2132133101200302-0101122201111122-3210033312122011"></a>

## security_style property — volume_defaults / 001312233003 / 8

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1333313030112300-0123000330132300-3313022130311233-2111333223003130-1120033330021323-2301212323320322-2202303321301112-3201213003200033"></a>

<a id="canonical-0201300212301203-0220013223332222-3002022221223102-0113331121023323-2331022230021130-0002221323332031-1110302111110120-1123031202012001"></a>

## snapshot_dir property — volume_defaults / 001312233003 / 9

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2012221032212021-3111203233312123-3230033223333031-2001212222230220-1131233321223300-1120012202320210-0000001023221022-1233313020110303"></a>

<a id="canonical-2121212101203222-2230232301011223-2011222111222100-3102011003321233-1223330131323300-1032212301101032-3233010130301012-1211230331122011"></a>

## snapshot_policy property — volume_defaults / 001312233003 / 10

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3210113120113130-1302003203220102-1301322313023200-1203210302320230-1000331222221113-0202222231021122-0121110203013100-2030230202000222"></a>

<a id="canonical-3113232001202021-2021301103121322-3222032302202101-2021131210321320-2021121120231110-1121323203113010-0210113030130221-0010023210012002"></a>

## snapshot_reserve property — volume_defaults / 001312233003 / 11

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0030331122233001-3012301030303032-0202131032313222-3313322111323201-1031132322131310-3012132302323103-2021231030232021-2113102323323322"></a>

<a id="canonical-2212231231023010-2333013203201320-3021213120222331-0311113120312331-3312121000201123-3023223222200310-1221031213122012-2002223331100100"></a>

## space_reserve property — volume_defaults / 001312233003 / 12

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-1322331331203332-3023110133211300-3131212012101011-2223322332121320-3112030321111222-3321300001011101-1102130011203213-3031120013201031"></a>

<a id="canonical-2333303331330200-3111033120013112-3111332212202101-0311032113021310-0323011322213022-0120200112320233-2121322300302033-1313033203330030"></a>

## split_on_clone property — volume_defaults / 001312233003 / 13

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0323011110310122-1331301002221123-3020113030121221-0003330131212322-0003103102133031-0333020212203002-2333201120311323-3220230211113002"></a>

<a id="canonical-3011033221331010-1302231033032120-0233012010131312-3311121232333212-0110303012103210-0303112130223131-2212020011312010-3031332212120331"></a>

## tiering_policy property — volume_defaults / 001312233003 / 14

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2203210122321211-2332220031010113-3222322310010033-1021322210033003-0133121332112012-3030111213321120-3212103032012001-3312231312101323"></a>

<a id="canonical-0020120221211231-2302221121330002-0101201233302233-2103123311322131-3223303010100120-3232031112202211-1031202131110201-3033011232133111"></a>

## unix_permissions property — volume_defaults / 001312233003 / 15

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2231221131120022-0303231123133210-0020332330033212-0100120123031331-2203210323210213-3233221001321212-0232110000323311-2030203103221103"></a>

## Next pages — volume_defaults / 001312233003 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-2201230221322113-3100212201331313-2122002222033101-0232333201101302-2223232110333230-1323231320121001-0133012121022011-3011221133303100)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2201230221322113-3100212201331313-2122002222033101-0232333201101302-2223232110333230-1323231320121001-0133012121022011-3011221133303100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222220220130313-0202300233231231-0222031122103211-0120230002020002-2030222010321022-1310320132231130-3333210120102220-2111013120203320"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos — no_qos / 311102321003 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](data-sources--fleet--reference--group-003.md#canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1031002301120210-1211113332013001-0221212030021230-1213323311322101-0103111130010301-1132312332210131-1123103220103002-1121031232101023)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos

<a id="canonical-2021001022300031-3202100022231310-0300011112223000-3013303023101202-3120133011203220-1232312001320012-1101010021010320-1223000203133202"></a>

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

<a id="canonical-2102330000101132-2121020033130030-2331312113010130-0130233221020201-3332013313230120-0230311233303323-2102110301131321-1301222313212330"></a>

## Direct properties — no_qos / 311102321003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311313030112030-3102011201013122-1330012000300233-0100210321230112-2131023313002110-3213011233323330-2313030003333022-3031133213030032"></a>

## Next pages — no_qos / 311102321003 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-1031002301120210-1211113332013001-0221212030021230-1213323311322101-0103111130010301-1132312332210131-1123103220103002-1121031232101023)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102222112122100-3331311013202113-1330213022312020-3221232033232021-2021123132301110-2302233211102030-3110122002101030-3111201230210122"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san — netapp_backend_ontap_san / 313333212301 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="canonical-0312112002310313-0011230320033000-0221223220022300-0123133033010010-0331311100212213-2001331113201230-3132333113002002-0311003130300301"></a>

Type: `"single"`. Computed.

Configuration of storage backend for NetApp ONTAP SAN.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-chap_choice": "[\"no_chap\",\"use_chap\"]",
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

<a id="canonical-0133303300013300-2231332001331100-2210130212121233-2312130202030030-2312121211112210-2300323031301213-3030032302301032-2031331103211213"></a>

## Direct properties — netapp_backend_ontap_san / 313333212301 / 3

<a id="canonical-1103102323312322-1132320211201101-0100201201230033-3003012001303301-1121023102010221-2013113032121030-2122102102120222-1311111311322002"></a>

<a id="canonical-0113332323033231-3033331112001321-1101213201321312-2010021300211110-2300230102331112-2033021121232033-2023110300230230-2013313120022231"></a>

## client_certificate property — netapp_backend_ontap_san / 313333212301 / 4

Type: `"string"`. Computed.

Please Enter base64-encoded value of client certificate. Used for certificate-based auth.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](data-sources--fleet--reference--group-003.md#canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302): complete subsection reference.

<a id="canonical-1223011231032031-3320201301010320-1212323222010312-3100102111303203-0111130310121120-1213013121213201-3132310111210203-2022213301201310"></a>

<a id="canonical-1322201132321320-1322300322210313-3132011000101230-3201321231012102-0011133233331130-0013201210102112-2033230311111121-2331202201022303"></a>

## data_lif_dns_name property — netapp_backend_ontap_san / 313333212301 / 5

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1213302220302210-2213013233012032-0033220221130121-0330121333331331-0302201300211201-1112011013213202-3030311030323131-0112212103200302"></a>

<a id="canonical-2111000311131323-2121030132120300-1202320121122231-3100301223111002-1201033013332013-1203332031331103-0111231210001333-0223320323220131"></a>

## data_lif_ip property — netapp_backend_ontap_san / 313333212301 / 6

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2010323130231333-1023302320210111-2123330310131313-3333011123033023-1133133031201000-2321203301301021-2103022231110210-1321030331211003"></a>

<a id="canonical-0222313231021313-2332212310031131-1113002201300200-3122020120001213-2333030231220012-3313101312333121-1112011020130122-3000012202013202"></a>

## igroup_name property — netapp_backend_ontap_san / 313333212301 / 7

Type: `"string"`. Computed.

Name of the igroup for SAN volumes to use.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0222003121002303-1132113230233322-2101231313301301-2013213121330333-0231003300112203-0302210120022030-3020103211321322-1110032332121003"></a>

<a id="canonical-1021131013003123-2202221321211310-0013203233312313-1220111201313333-1302313220111301-1210103021011103-2001103031133112-1200331322013013"></a>

## labels property — netapp_backend_ontap_san / 313333212301 / 8

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3222131122000132-2121230300303330-1122203111112130-2331232021333331-3122100301113323-0000302131130221-2222002222211300-1022303212302231"></a>

<a id="canonical-3233121211033022-2110210222011312-1120212231030120-0132301103102310-1132112222223131-1211013001110000-1023030011110113-0332003310113302"></a>

## limit_aggregate_usage property — netapp_backend_ontap_san / 313333212301 / 9

Type: `"number"`. Computed.

Fail provisioning if usage is above this percentage. Not enforced by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0133123200301012-0120230111120202-0030200213020320-2132121010212232-1230210211113322-0311332323031013-1313320102131101-0312111031102122"></a>

<a id="canonical-2300220021133031-3133322301133130-1022120312300020-3103310002321031-0031233333220330-3220123023222132-0220211222210221-0333003001230132"></a>

## limit_volume_size property — netapp_backend_ontap_san / 313333212301 / 10

Type: `"number"`. Computed.

Fail provisioning if requested volume size in GBi is above this value. Not enforced by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3310233010320233-2100333322313023-3220103023333232-2112103123313330-1023022301221002-2233022002021122-3200101010322033-3020201001323031"></a>

<a id="canonical-3011301103222203-0102022103102332-2203330203233131-1301121303311111-0203221111231231-0323031320221132-1001101110230103-1001233001223323"></a>

## management_lif_dns_name property — netapp_backend_ontap_san / 313333212301 / 11

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1000220103221121-0321233031310222-3211232303013303-1020032031132101-1011330302010231-3032103023031321-3202301033110210-0103301113310032"></a>

<a id="canonical-1102300020231203-0302212121331220-0121303113001201-0222111002221322-3122313201222301-2202331203123113-0331020222331013-1302002032311203"></a>

## management_lif_ip property — netapp_backend_ontap_san / 313333212301 / 12

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [no_chap](data-sources--fleet--reference--group-003.md#canonical-3313201320201133-2313002020200110-1323011320323131-2311012102232011-1201313313233301-0323203012311221-2032102101000000-1203131310313310): complete subsection reference.

- [password](data-sources--fleet--reference--group-003.md#canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213): complete subsection reference.

<a id="canonical-2020022010103223-2021130102210223-3133323031212321-3013332032332021-0013230203210011-1010302103032132-2220313301332002-3000002101331313"></a>

<a id="canonical-1002213230121131-2031120210121312-1010212333111333-1020320003233110-0033332203221303-3222321101333123-1002132121122210-3001103022230130"></a>

## region property — netapp_backend_ontap_san / 313333212301 / 13

Type: `"string"`. Computed.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [storage](data-sources--fleet--reference--group-003.md#canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000): complete subsection reference.

<a id="canonical-0211322203302201-2300121302203201-3133221110121030-0033122031303030-2031031031223231-1002210021211121-3032200221120123-3001310123301101"></a>

<a id="canonical-2223022032001132-0313310210111323-0023122212010223-2013102330131302-1313131313331322-0001312202322331-3101322310123112-1032321100122320"></a>

## storage_driver_name property — netapp_backend_ontap_san / 313333212301 / 14

Type: `"string"`. Computed.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-0322122121233033-3001031213030200-3002030113102013-3110230203230110-0231000221210030-1321222321220313-3201003112211322-1221133301322123"></a>

<a id="canonical-3320111133012103-1120131210330222-3131303001123321-1323132132232001-3220311203120010-1310212103100100-0001120321200021-3331312020111023"></a>

## storage_prefix property — netapp_backend_ontap_san / 313333212301 / 15

Type: `"string"`. Computed.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 80,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 80,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0133200121312230-3113011211332221-1102303012011010-2233221332110000-2313021102100330-2213210012012211-2313021120322331-2111132110313330"></a>

<a id="canonical-2032112031213301-3130211303303201-0132221022323131-1233133202000121-2310133112233012-0233020011321300-1113331213311331-2300131331203222"></a>

## svm property — netapp_backend_ontap_san / 313333212301 / 16

Type: `"string"`. Computed.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3312110300012211-1112031120233122-1103331132210021-1133011111131232-1212110011221330-3311233220113312-0121232102130320-3200131222013010"></a>

<a id="canonical-0203220212131011-2030213232221211-2032031110330201-3010323131133003-3013220013120030-0333123030332020-2222202200231023-0022020322020012"></a>

## trusted_ca_certificate property — netapp_backend_ontap_san / 313333212301 / 17

Type: `"string"`. Computed.

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223): complete subsection reference.

<a id="canonical-2033003311022231-2100213321002033-3013221033210111-3310100003012033-2001012002010201-1010200313133320-0112303303223033-1302012333123013"></a>

<a id="canonical-1110022011202011-1011310331330320-1030012002333000-2313221013311233-1211112012110001-0302120100011030-3011012313211211-0212310102131132"></a>

## username property — netapp_backend_ontap_san / 313333212301 / 18

Type: `"string"`. Computed.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--fleet--reference--group-004.md#canonical-0312231301000023-1223223123130323-1300021202220113-1200230003223230-0121123321012302-1111201303220322-3001223112113203-2120323230132130): complete subsection reference.

<a id="canonical-3121103200111133-1113231131332213-0131122003232331-3002233102102120-2021302300000013-1012133012330032-0101130333100200-2200013222121311"></a>

## Next pages — netapp_backend_ontap_san / 313333212301 / 19

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](data-sources--fleet--reference--group-003.md#canonical-3313201320201133-2313002020200110-1323011320323131-2311012102232011-1201313313233301-0323203012311221-2032102101000000-1203131310313310)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--reference--group-004.md#canonical-0312231301000023-1223223123130323-1300021202220113-1200230003223230-0121123321012302-1111201303220322-3001223112113203-2120323230132130)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003331212010330-3332003032232030-3002202002123321-3200300302111102-3303302231202000-1230322322201300-3111302020123000-0112311103032231"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key — client_private_key / 230130230333 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-3031003302120210-1112103033332012-2322330201031200-2232132110032223-0033200211101111-3100223210012331-0233223020222332-1103232300111231"></a>

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

<a id="canonical-0222331111203312-3112213032033001-1302002321321201-0001333131002222-2123102110011300-2213210321222111-1120031120322011-0111321123133212"></a>

## Direct properties — client_private_key / 230130230333 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3222030100002311-0011211111301020-1200320221123312-1020200020121232-0322000012302113-1203030000200313-1021010202032220-2222030133030010): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-3221203331201322-2001320000331111-2011101102210002-1121330001202201-3002200013300231-3111332102131332-0022013300100031-2012302021332301): complete subsection reference.

<a id="canonical-3121013000010310-3120202201133120-3100023103103220-1211131333303003-0013122311120030-3320003321222011-2013020131131123-2310313303230010"></a>

## Next pages — client_private_key / 230130230333 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3222030100002311-0011211111301020-1200320221123312-1020200020121232-0322000012302113-1203030000200313-1021010202032220-2222030133030010)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-3221203331201322-2001320000331111-2011101102210002-1121330001202201-3002200013300231-3111332102131332-0022013300100031-2012302021332301)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3222030100002311-0011211111301020-1200320221123312-1020200020121232-0322000012302113-1203030000200313-1021010202032220-2222030133030010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201210220322122-2301333022112202-3210122122213303-1003031320220232-2200221330011123-2123102110033221-1221311321122223-1000322200221323"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info — blindfold_secret_info / 111123031121 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-3200233233132022-2111132032132010-2303203113011311-3101221233200013-2100121322133213-0210030311010220-2002222103030121-1113323122100031"></a>

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

<a id="canonical-3032122001203032-1113203302301100-2033333213100311-0300001022001030-3333131101023103-1010312022021330-0013313113322312-1113021023011310"></a>

## Direct properties — blindfold_secret_info / 111123031121 / 3

<a id="canonical-1110023111312112-1231021123032003-1110211030300021-1230232130122322-1331132003000320-2210131221201221-3012133233002220-3203203121010223"></a>

<a id="canonical-0223220100230001-3001211312030220-2121022110322311-3002013123221122-2202230030023031-0221032133033203-0333131210121122-0020122010322310"></a>

## decryption_provider property — blindfold_secret_info / 111123031121 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0003321100231310-2010123033102232-3132333113201000-1312223331300120-1033311313322221-1303120021223003-1300102132220210-3101001201310121"></a>

<a id="canonical-1111022323333321-0032022131022122-2312101001012133-1133230310211003-0133312111321312-0300020223023021-1301320031301001-1132020020312121"></a>

## location property — blindfold_secret_info / 111123031121 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1201000000021200-0331323302302310-1003021011230300-2110012333013200-2300111332112113-1132032111001232-0232231221321130-0123333033300031"></a>

<a id="canonical-2122133210223031-0320311100233022-2121020310131203-2023111102310312-3020220010323112-3112132123322111-0300120222022012-1310123031033021"></a>

## store_provider property — blindfold_secret_info / 111123031121 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2000033200322323-1312012202121321-0320132311201221-3333120021001030-1021030302321033-3310003222331231-2300301013213010-3020221030213231"></a>

## Next pages — blindfold_secret_info / 111123031121 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3221203331201322-2001320000331111-2011101102210002-1121330001202201-3002200013300231-3111332102131332-0022013300100031-2012302021332301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200013201012031-2332323010221023-2031122121102010-3203210010031331-2120011130310202-2102033120213321-0302020321122011-2300111022232310"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info — clear_secret_info / 222011323200 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-1323002212231221-3031200330111102-1201032232101100-0132302300021133-3201311113123001-2331222030313221-0201031310312331-1131230222322201"></a>

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

<a id="canonical-3101322133020212-3100200103300120-1112012031310202-0332322333030010-0233200211321333-0120320120023122-2303212312320002-1021313020201312"></a>

## Direct properties — clear_secret_info / 222011323200 / 3

<a id="canonical-2103313223212303-1112220332121103-3003133033030230-3030302132003320-3121100112312031-1033230013310333-0123301002022112-3133132211303100"></a>

<a id="canonical-1230022132011212-3033001010101123-2310101330313032-1132332010032322-3012001220021311-0301033022123102-3102020313333131-0102312110302010"></a>

## provider_ref property — clear_secret_info / 222011323200 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0232202332113011-1113232303133010-1201101302202131-3113031120131213-3133313220323232-1320133020022231-2201231310322021-1121021222230021"></a>

<a id="canonical-1112100103323003-0310330003201322-3133120122202030-0200130033231021-3133331302022010-2123213220031001-1213020300223300-2101030302120301"></a>

## URL property — clear_secret_info / 222011323200 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3133303020212000-2201101123130310-0012100011212132-0223210322210333-1203001211101232-3201021010231313-3030120300332132-2203031132301312"></a>

## Next pages — clear_secret_info / 222011323200 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--fleet--reference--group-003.md#canonical-1302022310332322-2210300011312230-0302201131212100-2011210313122032-0300202010231330-3313111122220121-0333031121210010-3322300012301302)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3313201320201133-2313002020200110-1323011320323131-2311012102232011-1201313313233301-0323203012311221-2032102101000000-1203131310313310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032021013322313-3101311212013210-3313031020202322-0022003213223323-3222130112030113-0012131320130013-1031300013212111-2310131021030102"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap — no_chap / 300231110330 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-0310031112220200-0123213310312233-3212003112311332-2030013013221131-1302002123030320-3021012122132210-3301133303230322-1013013303310122"></a>

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

<a id="canonical-1112311132112322-2221111002100321-2320323102323212-3111322120320203-1112100000113210-2032022230203222-0222212130120132-3301132200221103"></a>

## Direct properties — no_chap / 300231110330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322120111132122-3212030300322103-3122311200210223-2012032012313110-1220010320312222-1122223131321001-3111312012123301-2011322030322222"></a>

## Next pages — no_chap / 300231110330 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333001330023233-3322230003120030-3300002032133020-2102312111102122-1010132220111232-3221301303032013-1033013113230133-0003223021100123"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password — password / 310101200220 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-0310232031000322-2020333220332233-3102000221213313-3330031001223123-0013203031031021-1002202203222203-0221000322201112-2223022211013200"></a>

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

<a id="canonical-2313122323111213-2111010203332113-3310013001100311-2033002033021300-2220203121112133-2111321212210300-2131110332020230-2133012012003000"></a>

## Direct properties — password / 310101200220 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3133112300100001-2200211002101323-2112032312301132-1032203130010222-3100032002101311-0212023330122312-3012221231200212-0103213130302323): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-3102200112033112-2132102311221003-0031231023020300-2333323202030130-0103122031133233-2223001132020233-0121102213220120-3000003302100031): complete subsection reference.

<a id="canonical-1001103132210002-3101130221031000-0013133031223331-2122101011131003-2312200223213233-3230012002020130-2213330202110001-1223311103333312"></a>

## Next pages — password / 310101200220 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-3133112300100001-2200211002101323-2112032312301132-1032203130010222-3100032002101311-0212023330122312-3012221231200212-0103213130302323)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-3102200112033112-2132102311221003-0031231023020300-2333323202030130-0103122031133233-2223001132020233-0121102213220120-3000003302100031)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3133112300100001-2200211002101323-2112032312301132-1032203130010222-3100032002101311-0212023330122312-3012221231200212-0103213130302323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023111003310133-1313013220222203-0201233232222222-2321213221323012-3312033013110300-0020013231020131-2103330022110203-1131233333133332"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info — blindfold_secret_info / 030230312232 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-1011032010132111-1210033133110121-1211320232203110-3101201121202321-0102231300121111-0120310322130210-2210000002030013-1023222120222301"></a>

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

<a id="canonical-1221200301120023-1113310023231013-2131111212100020-0103012232320312-0230031031203203-1201213231201103-3011303113112021-1312111301331031"></a>

## Direct properties — blindfold_secret_info / 030230312232 / 3

<a id="canonical-3021103221322131-1213221210233200-0030213231113000-2311013000332121-2131311002003011-2103211010220220-2222311310110100-1111311302301102"></a>

<a id="canonical-0301320300030020-1130310302223322-3321013210223103-0013112113131121-1012013131030012-1032012011333311-0232301100233103-3202021223321323"></a>

## decryption_provider property — blindfold_secret_info / 030230312232 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0222222312313033-3133322100001012-1231132131002202-1013021310210311-3021112311102221-2033333010311201-1120013323231001-2213312332021123"></a>

<a id="canonical-1010330221023100-3333022313213111-3112223201111133-0223301210020322-0012331020032300-2332003121120323-2222013031130130-2021001032010020"></a>

## location property — blindfold_secret_info / 030230312232 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1311121320101323-2023130201222003-1231113013033131-2200120011331231-2331331110122302-0100022032123312-0020003332313002-0200213003031333"></a>

<a id="canonical-2331301122203021-1131202203130122-2032320032121300-1212133110321212-2313112133331213-0131233223102013-0001212220302232-1203131332203320"></a>

## store_provider property — blindfold_secret_info / 030230312232 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0310300120121220-1101222213032130-0210203300313001-3300230223333133-0010022112332210-0321313111113131-1110021000222221-0002113301302102"></a>

## Next pages — blindfold_secret_info / 030230312232 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3102200112033112-2132102311221003-0031231023020300-2333323202030130-0103122031133233-2223001132020233-0121102213220120-3000003302100031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233301323121313-3120022312030212-3213231203312232-0111211110013022-0303213221232313-3023202213330002-2311021010111102-0022323131212313"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info — clear_secret_info / 113031313101 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-3231010021101212-2330203013213232-0333103330222012-1312023210230211-2221020033302012-0121121131222112-1013322011201003-2301211312132103"></a>

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

<a id="canonical-2112102110332132-0031123011130100-1333013031101120-3333302311331130-1133033031011011-0023032230332100-1102111013111033-0100321102330133"></a>

## Direct properties — clear_secret_info / 113031313101 / 3

<a id="canonical-3302311001112122-0023202123033212-3231102203230031-1232032102132122-3210030220020013-0323103131303101-1131320333133010-1231121320313321"></a>

<a id="canonical-3200001121003010-0000110022221020-1131201110103000-0213132110100012-2023220021120301-2332001211132022-3020110213033133-2223013102320233"></a>

## provider_ref property — clear_secret_info / 113031313101 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1002323110330003-1303310332322020-1312023031110032-1020230110223320-0212312120311031-2200132223213012-2120220033102100-0023101031101302"></a>

<a id="canonical-3113133011010212-0332003033323131-0203021202230112-1312012101121011-0301010101110300-2131012120233102-0012131313203111-3110232033211130"></a>

## URL property — clear_secret_info / 113031313101 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1212320323110113-0320121101320003-2221003331232323-3320110010313311-3113221012101323-3013000012221203-0203330203212321-0321113200133002"></a>

## Next pages — clear_secret_info / 113031313101 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--fleet--reference--group-003.md#canonical-1120102000102001-2032321311131031-1110201233301132-2222110002303111-1003233110303022-0232203211001321-3022023132001121-0023311012331213)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103311021021320-0113033131313300-1230120323220111-3120301303103302-3323323302011202-3310320221311230-1330023013331313-2232121320111333"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage — storage / 301322131311 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-0310312313110012-0202213323202030-3012112021002203-2301311202031310-2103221132110001-0023130211002111-0210010331321201-1200210032113110"></a>

Type: `"list"`. Computed.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0002012330133012-2301311100001211-3211222131301022-0030231311030112-2111022231012123-3122200210330102-1203100021032022-0031121220000203"></a>

## Direct properties — storage / 301322131311 / 3

<a id="canonical-0132220222322321-2032001131231003-1212230231132231-3302231220001233-3313323222000001-1211131031132002-2303313112002232-2123023323200031"></a>

<a id="canonical-1200220132211220-1013301330310132-3332100312123201-1311320120111221-3012122312121211-2302313333002323-1010102121222211-0302101212222230"></a>

## labels property — storage / 301322131311 / 4

Type: `["map", "string"]`. Computed.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](data-sources--fleet--reference--group-003.md#canonical-0100022222011211-1020313330202211-0130320331202011-1032331232231310-1221120203231123-2000021330110310-2302032001301201-2122232132023331): complete subsection reference.

<a id="canonical-3010122312301113-3020310312300202-3033213301101001-0200202230333111-1303333323302230-1332000233331021-3222213231220222-1300022312133010"></a>

<a id="canonical-0323323100320133-2332012112111230-3210221333330110-0032122301210331-3032112222203132-1313302231221131-3202323102212331-1301231311301201"></a>

## zone property — storage / 301322131311 / 5

Type: `"string"`. Computed.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1220221031123230-3023010101322321-2300012302322033-3311020321202213-2201330211322312-0133130031233322-1120010131221100-0232312322321220"></a>

## Next pages — storage / 301322131311 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-0100022222011211-1020313330202211-0130320331202011-1032331232231310-1221120203231123-2000021330110310-2302032001301201-2122232132023331)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0100022222011211-1020313330202211-0130320331202011-1032331232231310-1221120203231123-2000021330110310-2302032001301201-2122232132023331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322010122123212-1022312300320000-0330132233233231-1130131210323313-0032322023132330-1323223111232213-3320111312233003-3123220121032201"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults — volume_defaults / 032322322113 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-2203103131332010-2222303201333330-3323333211012333-1320201310300112-3120133211123031-2122033311121220-2001122113131130-3311303210030310"></a>

Type: `"single"`. Computed.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

<a id="canonical-1320210102303213-1323210112313021-0111123133321222-2030233001302003-2232233223032122-1200323322121330-0212121000012332-2222211222333123"></a>

## Direct properties — volume_defaults / 032322322113 / 3

<a id="canonical-2013320102233201-3300111130201133-0011223320001200-0010210032310022-2321000111002213-2232022010133321-0312331301210202-1132023010030333"></a>

<a id="canonical-0301200213013223-3112211102101130-2201101000300021-2300220331010212-2032102112013231-3022322012202310-0311121111230303-0321013331200032"></a>

## adaptive_qos_policy property — volume_defaults / 032322322113 / 4

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3330000302201033-1120323012301320-3302033320011131-3220113032320022-2310123013203023-2102010002303032-3113031231122213-2032122022110203"></a>

<a id="canonical-3130212201033033-2013101032111211-2010133322111011-2010203133133001-0102130331221220-1213110021033223-1232221121110002-0103322102023103"></a>

## encryption property — volume_defaults / 032322322113 / 5

Type: `"bool"`. Computed.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2101020230122302-3200133102212201-2011220101220020-2300330233212323-2203132311010020-0203002210202303-2002222113310032-1330123110313022"></a>

<a id="canonical-0311310132131101-2210123122101102-3223133131130200-3012102331100320-1023031030220233-2032022211131033-0220122323030232-2100132221031113"></a>

## export_policy property — volume_defaults / 032322322113 / 6

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [no_qos](data-sources--fleet--reference--group-003.md#canonical-2033121012023220-2332020112110101-0321031031111313-3200123113133023-2100333003111122-2133231032311332-3013313230331031-2230330130112112): complete subsection reference.

<a id="canonical-3010230100021132-0200112020110321-0002330300103020-2313223300322312-2011320003031203-0332012320210031-3032222303311300-3301112211003130"></a>

<a id="canonical-1031021120201322-1202310303132201-3000010100031302-2121320011320003-2003100300211111-3230000212313201-2201020320220211-3002312101033233"></a>

## qos_policy property — volume_defaults / 032322322113 / 7

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0221000231020020-1332131130032312-2330311131020220-1231102320021312-1112310200203312-0320200232131301-0321033201202230-3221201021220012"></a>

<a id="canonical-1010302110211122-0102032112201011-0212211321331320-1200222011121130-1012312122020200-2012230022133320-3211121100021231-3312220031102031"></a>

## security_style property — volume_defaults / 032322322113 / 8

Type: `"string"`. Computed.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1123232011332130-0011201032221000-0211223223223103-1120202213011013-0331203301010112-2223011021023321-2213023232202223-1330211323001300"></a>

<a id="canonical-3221123211012120-1230320223033120-1011210312331211-2130103011212202-0203033121023001-0200113311022332-0300013001100202-3120000320101110"></a>

## snapshot_dir property — volume_defaults / 032322322113 / 9

Type: `"bool"`. Computed.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1102211102123332-1112230032312300-1313233212211231-1212133131200020-0333113120301323-1122122222132031-2210020331323023-0011312223223201"></a>

<a id="canonical-2301303101301130-0301112133132031-0322323103211023-3230021333000011-3003011232210120-2133200200022123-3313323310212230-1330320001010010"></a>

## snapshot_policy property — volume_defaults / 032322322113 / 10

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2002300030332321-3033310313110003-2002101330113031-2311001310032112-2003003201021131-2100203203012100-0303210220011123-2210211003101101"></a>

<a id="canonical-1131012132021033-1322200221212033-3213222232122123-2103321012201313-2312222033012332-0223110002223132-1323103332103123-0013031022122323"></a>

## snapshot_reserve property — volume_defaults / 032322322113 / 11

Type: `"string"`. Computed.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3002023203302231-2221001032023310-2211222202231023-3020032100330233-3030003201330023-3101130302023023-2202210130311310-3030000332031022"></a>

<a id="canonical-0120103221121031-1112100001130301-2220121212312103-1002010022121133-1210003132110232-2023323002322322-2130303222033302-2033000322010232"></a>

## space_reserve property — volume_defaults / 032322322113 / 12

Type: `"string"`. Computed.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-2320313331131302-1112030310210210-0230101233021223-2211101021030133-3101033331211320-2011033131232112-3302221120330201-3001031212120113"></a>

<a id="canonical-2122221010230111-0001022230330210-3100200221102323-0102302210311112-1123103300021210-1311211320131332-1111222333032023-0033201010010100"></a>

## split_on_clone property — volume_defaults / 032322322113 / 13

Type: `"bool"`. Computed.

Split a clone from its parent upon creation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1030303301321303-0302111101300323-1021221100110002-2030111031101201-0322331311323320-2031313011331000-3302203021211212-2301132032231002"></a>

<a id="canonical-1133321122001310-2322022001300020-3103310231032133-0323012230302231-3133120300220311-0101220232001130-1212012222011120-1022320012301212"></a>

## tiering_policy property — volume_defaults / 032322322113 / 14

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2012300231311000-3021200322013233-2313312033023211-2130003133132132-2012112312011210-0202102231223003-1123213130300332-1202111130113001"></a>

<a id="canonical-2032111303022133-0300112100223320-1030123231202321-3133132333232212-2021000022203113-0221202010211313-2323122221303301-1130302330200033"></a>

## unix_permissions property — volume_defaults / 032322322113 / 15

Type: `"number"`. Computed.

Unix permission mode for new volumes. All allowed 777.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3012230130203232-2020321121121200-3201032011303333-0010222103313201-3230313022313122-3212103323013021-2203031323220120-2301330121102331"></a>

## Next pages — volume_defaults / 032322322113 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](data-sources--fleet--reference--group-003.md#canonical-2033121012023220-2332020112110101-0321031031111313-3200123113133023-2100333003111122-2133231032311332-3013313230331031-2230330130112112)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2033121012023220-2332020112110101-0321031031111313-3200123113133023-2100333003111122-2133231032311332-3013313230331031-2230330130112112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222022231100030-2320120003010310-1321001100231031-0003001310010233-2020333133123233-1203210300313300-0221333333023133-2020332210233102"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos — no_qos / 003312002212 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--fleet--reference--group-003.md#canonical-0203301222213122-0311200133322320-2330230003102220-0010021312213121-1003010303003013-1100202112013010-1112331202022330-1012101223313000)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-0100022222011211-1020313330202211-0130320331202011-1032331232231310-1221120203231123-2000021330110310-2302032001301201-2122232132023331)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos

<a id="canonical-2132313213220021-3232233203203002-1111201021233232-3200200321100203-3333021312131130-3132023011210130-0201023132223020-0213202033021210"></a>

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

<a id="canonical-3312000222230323-0023312110323122-3013021332023212-0221221332131121-0100303311231201-0010003100112312-1130212000021122-2301103030010321"></a>

## Direct properties — no_qos / 003312002212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330303212320001-0111103111312301-3213121321102301-3202022131323330-3233113123202211-2333333012231312-0111100101213121-2220202001010333"></a>

## Next pages — no_qos / 003312002212 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--fleet--reference--group-003.md#canonical-0100022222011211-1020313330202211-0130320331202011-1032331232231310-1221120203231123-2000021330110310-2302032001301201-2122232132023331)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320330312330121-3022322211030221-3101031020221202-3120000120211003-0302303222023233-3222201231131122-1323002202322332-2232112312202331"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap — use_chap / 111201121303 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="canonical-3022032202232031-2330032132010002-2220122221011220-2323310200322320-2323222110012203-3302020322231021-3023011223311111-2203231010130110"></a>

Type: `"single"`. Computed.

Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0313003001312201-0132113320302333-2333003311221000-0312321310132333-2111031030230210-0100101131131211-1033031222023000-3132233230032110"></a>

## Direct properties — use_chap / 111201121303 / 3

- [chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322): complete subsection reference.

- [chap_target_initiator_secret](data-sources--fleet--reference--group-004.md#canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030): complete subsection reference.

<a id="canonical-2131322311223003-1221000333013200-3013031110012322-3213230321322110-3231130002311102-0000331300131010-3331222203103330-1100210003121301"></a>

<a id="canonical-2022323222201131-3011200220212201-1010212300202313-3310213111222021-2211320221001330-0302220202200203-0000312001201010-2331231330322320"></a>

## chap_target_username property — use_chap / 111201121303 / 4

Type: `"string"`. Computed.

Target username. Required if useCHAP=true.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3133233032230213-1222110300230203-0123211033313023-1120202133101201-3021103202131313-0313201022132200-2001122200012300-0311003232302201"></a>

<a id="canonical-3003130322110111-3210110020330123-3233110320002220-2001220302211222-3003123132221101-0130123023111303-2021210032303103-2330123330330333"></a>

## chap_username property — use_chap / 111201121303 / 5

Type: `"string"`. Computed.

Inbound username. Required if useCHAP=true.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2303232131031133-1221010233313131-3202000112221022-3320101313000213-0120123133313330-1102033302001313-0222122011330332-1001332230011223"></a>

## Next pages — use_chap / 111201121303 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-004.md#canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013013321032021-0131010233001101-2230000013321012-3323103023330233-1001301002311131-0111203111223331-1331230031120120-3112031122323021"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret — chap_initiator_secret / 020330000302 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="canonical-1121122311111333-2010100230213332-2302011132320211-3220012023210112-1010310010132010-0131110031200101-1020110123220331-2021130301303103"></a>

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

<a id="canonical-2031110330120101-2322021323011122-2231220221210023-3223332133210223-1222020223212320-2321013233113032-3132002210313031-1310222323300313"></a>

## Direct properties — chap_initiator_secret / 020330000302 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-2020211030320032-0322010313303130-0100222030313123-2032033001123011-1213103330003331-1320230100311212-0131130132220300-1021102203013301): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0112301002112002-2212002033302312-3101021132222000-0102320320013121-2303232230000130-3001211103013021-2032311020113303-0002202113300332): complete subsection reference.

<a id="canonical-0100323222303222-2032110233020020-3222021321221301-0331313211133110-2132310123333102-2332333121122112-2123233000133113-0311310233001332"></a>

## Next pages — chap_initiator_secret / 020330000302 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](data-sources--fleet--reference--group-003.md#canonical-2020211030320032-0322010313303130-0100222030313123-2032033001123011-1213103330003331-1320230100311212-0131130132220300-1021102203013301)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](data-sources--fleet--reference--group-003.md#canonical-0112301002112002-2212002033302312-3101021132222000-0102320320013121-2303232230000130-3001211103013021-2032311020113303-0002202113300332)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2020211030320032-0322010313303130-0100222030313123-2032033001123011-1213103330003331-1320230100311212-0131130132220300-1021102203013301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203210301203310-2301201213000230-2232111121100001-0130000100102201-0313111011001001-2103030321213112-1023332201320333-3211001302122011"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info — blindfold_secret_info / 212330103312 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info

<a id="canonical-1210312020222122-0311102012211233-0021003031210230-0131220221032203-1113101200321321-2112312121030212-0112300123322021-0130212223322123"></a>

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

<a id="canonical-2330213200212001-0232312033232323-2113211010112220-3111232231312101-2002230001001130-3031333323102211-3111310202011323-2201102010213023"></a>

## Direct properties — blindfold_secret_info / 212330103312 / 3

<a id="canonical-1120220122032302-3032231022031201-2013133023021022-0020021013132013-2012013320112300-3310232103032333-2321010330013200-2131123232332222"></a>

<a id="canonical-1021302301203220-0011320010120121-3323112121121212-1032200203021131-2033222111130102-3233102123033112-0022010131321211-1301201121312023"></a>

## decryption_provider property — blindfold_secret_info / 212330103312 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1301222302320302-1033103201301010-0013032000032113-2333132021002022-2110311221301021-2123130322322300-3333132132331232-2010121002220231"></a>

<a id="canonical-3212201230111100-2030220303311023-2110322013302231-3231013001311312-1133123002220303-2022103223000231-1302103100321323-1130121311210102"></a>

## location property — blindfold_secret_info / 212330103312 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3133211210222110-1303113032323320-3021121023221203-2212300223123323-0323311023313213-1220312210223000-0320210331013030-2232200130220230"></a>

<a id="canonical-0031022101132012-2033302101113312-2233321022031311-2323013211010132-1333002200130130-3312101003013122-2132023132023012-1320101212321112"></a>

## store_provider property — blindfold_secret_info / 212330103312 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0210330131101303-2122112210122220-2132001300200203-0332310022213020-0333032132332300-1330021110020220-1213312122130110-0110100120012023"></a>

## Next pages — blindfold_secret_info / 212330103312 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0112301002112002-2212002033302312-3101021132222000-0102320320013121-2303232230000130-3001211103013021-2032311020113303-0002202113300332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
