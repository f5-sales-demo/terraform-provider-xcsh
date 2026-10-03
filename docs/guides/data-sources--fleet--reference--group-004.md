---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-3103030303223011-2231202303033121-1200322230221120-3102232012130133-3323221301112333-2330101303030103-0033022000003233-0021223002303111"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info — clear_secret_info / 203101211121 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info

<a id="canonical-2223020212022331-3131033330011233-3232030101122132-1333030132122021-0031133120030033-2012201301110021-0012122110021011-3003030022302031"></a>

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

<a id="canonical-2022221033320203-0220111122203020-3113121000321301-3020012333031220-0021330302013311-2223301101100020-3102311110021221-2120223120131012"></a>

## Direct properties — clear_secret_info / 203101211121 / 3

<a id="canonical-3113003011211123-0330012032002031-3113301030003211-3230130333202022-0030123030102020-2123031311212233-3112300212311101-0131001120133333"></a>

<a id="canonical-1100113132122323-0223222210221011-3100120330030031-2212033002120333-2013010002310201-2233112221333232-3230223010231111-2110013101202212"></a>

## provider_ref property — clear_secret_info / 203101211121 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0221013221131232-1313102331122231-2322101201212133-0320331010130130-3002330132033011-3130022311212231-1200112123123321-0310212123102121"></a>

<a id="canonical-1332111003312110-3331021200231220-2032223123012000-1002003000010103-3033022022112131-1102112002113020-2111201232021220-0321033102320233"></a>

## URL property — clear_secret_info / 203101211121 / 5

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

<a id="canonical-1232032321231100-3202030113331212-1131321022312210-2111210210133122-0301131001030311-2033213303210130-2010233232201102-2001130230123110"></a>

## Next pages — clear_secret_info / 203101211121 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--fleet--reference--group-003.md#canonical-2222302212123213-1032300331301122-3210031123122321-2010312311333111-1323331323202131-0103320000013002-2310212230020120-1210311133011322)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322010223012100-1023301102303001-1300032202020102-0232012111233131-0232001010121223-0102210023322313-3321332013202121-3220321233030020"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret — chap_target_initiator_secret / 110320012131 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

<a id="canonical-2111032202300332-0202103323100031-0223020330201333-1011031112200332-0333131202301203-0133301002033001-3102120001132110-3122312230312033"></a>

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

<a id="canonical-1322303220231220-0032020002003223-2010021000012013-1012110021313331-0121200030111133-1013132132302203-1301113030013102-3012302201132322"></a>

## Direct properties — chap_target_initiator_secret / 110320012131 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-3322321230331020-0100100110023131-0111023101130232-3213312223203302-2201122113320010-2230002030222100-2312221020031100-1011330023231132): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-1310022120000032-1223310030013111-2232132100112231-2213123112022221-2012011031030031-3221030232210233-1030312121231330-0210300320102301): complete subsection reference.

<a id="canonical-3002331233001001-0212313330330110-3311002000210112-2320112313311332-2210321321231113-0202203133212111-1121132011012122-3300322131332130"></a>

## Next pages — chap_target_initiator_secret / 110320012131 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-3322321230331020-0100100110023131-0111023101130232-3213312223203302-2201122113320010-2230002030222100-2312221020031100-1011330023231132)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-1310022120000032-1223310030013111-2232132100112231-2213123112022221-2012011031030031-3221030232210233-1030312121231330-0210300320102301)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3322321230331020-0100100110023131-0111023101130232-3213312223203302-2201122113320010-2230002030222100-2312221020031100-1011330023231132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020302200331130-1210021023100303-3200133131302330-3003211320231121-2222321131101133-3010033033120113-3013121021223303-0121232003333101"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info — blindfold_secret_info / 313101333120 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-004.md#canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="canonical-2311100313330320-2301332312000222-0132012020012220-3312112032210300-0311211333202033-1123022020032013-0321320230230133-0133210020313031"></a>

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

<a id="canonical-3012130330023122-2331031020030003-3030223202021102-1023011022122001-1301220213332230-1302032131202331-3012100222231200-0133313022022213"></a>

## Direct properties — blindfold_secret_info / 313101333120 / 3

<a id="canonical-2211320130111310-1113210033212101-3030102032300111-0131130203300231-3111210302033221-0211213112213233-0123001102223320-3113203303300310"></a>

<a id="canonical-1321310232232311-1030230012323121-1330311323133103-3311021122210030-2122311001333203-1311013320320121-3330233102010003-3020200131300013"></a>

## decryption_provider property — blindfold_secret_info / 313101333120 / 4

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

<a id="canonical-1012211210321310-1321131033011022-0300322113231233-3210331033023011-2001221103330130-1213022100232211-1020321310331111-3032231012313210"></a>

<a id="canonical-2032120212103301-3212230300223111-0001011222021321-1031013021011120-3321323230121301-3223131333333230-3230321323222111-0000121311103200"></a>

## location property — blindfold_secret_info / 313101333120 / 5

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

<a id="canonical-3310131330110201-1330211300211302-0012022130321133-2011213013112032-2003333211310323-1122330203113230-3222200301023030-1013022113011332"></a>

<a id="canonical-0231132020322111-0132223103103321-0221232321202113-2032012001302233-1221220212201001-2010032301232232-1023033130331021-2000001210332023"></a>

## store_provider property — blindfold_secret_info / 313101333120 / 6

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

<a id="canonical-1303102300032132-0221320010303013-0111302130001103-0233121222001132-3220022023123331-3333300133220103-1301103211200200-1202222123031321"></a>

## Next pages — blindfold_secret_info / 313101333120 / 7

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-004.md#canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1310022120000032-1223310030013111-2232132100112231-2213123112022221-2012011031030031-3221030232210233-1030312121231330-0210300320102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321021030230002-2132200122333030-0211112101223203-3333100322021213-0102121310012210-2021313203203033-3111322212320000-0100120112011022"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info — clear_secret_info / 220321010011 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--fleet--reference--group-003.md#canonical-3032230231202013-3203021200032120-3001221132331021-0310010001010202-1202210101021232-1311113001032110-2231130122032322-1023203010122223)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-004.md#canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info

<a id="canonical-1011000022331322-2202333231231210-0330313312000132-3313012122010202-3233212323000010-3223133011120330-1322222332322100-1331230010220312"></a>

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

<a id="canonical-0322011120023133-3100221003003222-0103231230223301-1011303003300230-3303101201320123-0132233112020332-2202203100201203-2001200003111130"></a>

## Direct properties — clear_secret_info / 220321010011 / 3

<a id="canonical-3130210001220331-0233213130111131-0012120132323313-0311211332030320-1100220033012120-1020303113232212-3311233132001000-2320202111032102"></a>

<a id="canonical-0013333303223031-3122331011110003-0223121101302233-1023323032210113-3131102013010111-3102231013112210-1002232230020123-1113320312321301"></a>

## provider_ref property — clear_secret_info / 220321010011 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2013110222210210-1213300101120031-2233101311031203-0110310031232100-1311030201012233-0123330130010330-3311003321113101-3233110103031131"></a>

<a id="canonical-2123121030023013-2211030102102330-2330112303021030-0221110103022120-1013210311012001-0130213313201220-2131200230111001-2023013222032121"></a>

## URL property — clear_secret_info / 220321010011 / 5

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

<a id="canonical-0220130201211221-3011311311103231-0100103330131222-1032013312213323-1310231131313032-3211101232300131-1121222320230301-1023110331322013"></a>

## Next pages — clear_secret_info / 220321010011 / 6

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--fleet--reference--group-004.md#canonical-3310332213203121-2032233103313131-0111301002333130-3323212301221301-3320230030331300-2033210302113211-1123113110010333-1310330231220030)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0312231301000023-1223223123130323-1300021202220113-1200230003223230-0121123321012302-1111201303220322-3001223112113203-2120323230132130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023121211121031-1021021203300320-0203313231213020-0111002132220330-2213132100332323-0311030033003001-1300103111212303-2111210113103100"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults — volume_defaults / 003303100320 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="canonical-1332120030303100-1011030102300031-1320033110110031-2221101323321313-1001122120023202-1122303033122221-2011320210112110-3301302212233320"></a>

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

<a id="canonical-0320320311113030-3021103010222001-2233013311203320-3303332122102021-2131001332110330-0201012023311311-0022100321121113-3222321113332132"></a>

## Direct properties — volume_defaults / 003303100320 / 3

<a id="canonical-0223312230221031-2223213122223030-3201123231233102-1211211210013000-0000101231233033-1122112233103021-0211132103033023-3030330123123131"></a>

<a id="canonical-0033300200011232-0130301033331022-1000012112110113-3003033120103102-2133301023330333-0322102223000121-1100332101303120-3021333202133123"></a>

## adaptive_qos_policy property — volume_defaults / 003303100320 / 4

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

<a id="canonical-3011310020100202-3001033230102321-3010133231100130-2120311321122211-2001003321333132-0030130133301023-2023310113301230-1123311130320231"></a>

<a id="canonical-3132332302331232-2102303002101123-1211000021322200-3032212012021310-0113021110301113-2002222102232122-0332021322131130-3302213031223011"></a>

## encryption property — volume_defaults / 003303100320 / 5

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

<a id="canonical-1331313020000122-2332221302010221-0120000130322233-2112010230022122-0222233113312112-3331123001232100-1033222001000120-0131002032012223"></a>

<a id="canonical-2221130300310021-2200323032112210-1023313230233303-2311322113122031-3012210020323233-2222233302113013-2111003223021021-3031301322231020"></a>

## export_policy property — volume_defaults / 003303100320 / 6

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

- [no_qos](data-sources--fleet--reference--group-004.md#canonical-0112010323101023-2331002330103032-1223131001001030-0221301023032130-0013032111330233-1231122211211221-0032203001033102-2302001301303210): complete subsection reference.

<a id="canonical-1001113023012001-3131000323101332-2012021220333012-0003122123313331-0303213333202120-2031203313001111-1320221032022110-0310231121103212"></a>

<a id="canonical-2001011132111210-1013200110220102-3120312331233131-3022132023111033-2320010321331211-2010002101120221-0131203332330203-3332121132212001"></a>

## qos_policy property — volume_defaults / 003303100320 / 7

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

<a id="canonical-2323313100102031-3223022212130221-3323123312200312-0100010032031300-1021202031130002-2033120300302300-1100133201212101-3011101000332110"></a>

<a id="canonical-0310212201233111-1100113120311310-2320000122333000-2010313102222222-2022203233220110-1023023302322323-1101321313032200-0311221301223023"></a>

## security_style property — volume_defaults / 003303100320 / 8

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

<a id="canonical-0303201022200003-3011211131102101-1332033220300030-3331223031121001-2211311003130023-3332233122023222-3030201033101011-2320130020313210"></a>

<a id="canonical-2302212123331202-3302312012301013-2100330110030223-1312010031010122-3210011210110233-1220121111211023-3010311223222133-2312210220333201"></a>

## snapshot_dir property — volume_defaults / 003303100320 / 9

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

<a id="canonical-2330002311031200-1323302032003100-3221003230330001-3231333232100011-0213330222321311-3030110000231210-0223011003330213-3321020023303103"></a>

<a id="canonical-2013321311223203-2123302212120200-3131113010133220-1132202023300033-1331223323333223-3201002132111033-0322120230000210-1302123232321203"></a>

## snapshot_policy property — volume_defaults / 003303100320 / 10

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

<a id="canonical-0010203023233122-0111302032300213-1300023003311020-0013222202311101-2312322231222022-1203011221233230-3120032300110000-1101230000031130"></a>

<a id="canonical-1112233133010120-0110200012300311-2131020130003102-3232022112012131-3212211233023013-0210231001023311-3102212203132100-0102300312212111"></a>

## snapshot_reserve property — volume_defaults / 003303100320 / 11

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

<a id="canonical-3032231002021200-2033111103132301-2331300103000022-3202013221031012-3300211032033002-2303300120022223-3131133023310111-0001221122132232"></a>

<a id="canonical-0232233221303000-1110030113221310-3223101012031120-3301001332333332-1233102212312010-0110233000301123-1132001030333030-0333321333321322"></a>

## space_reserve property — volume_defaults / 003303100320 / 12

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

<a id="canonical-0132231321211230-3002222301200020-1221220123322223-0113001131301220-2113312122022311-2213200320001023-0230013133022022-3211131033002002"></a>

<a id="canonical-2203223320331323-1212113010123321-2023200321323001-0022310010100032-3012123223322033-1323121031030302-0021110213132211-3322211102023323"></a>

## split_on_clone property — volume_defaults / 003303100320 / 13

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

<a id="canonical-2002101112132233-2132010020221320-1032230122321301-1231312010330200-1333220131012313-1310020323212332-3001333210011011-3123031102202001"></a>

<a id="canonical-0131030302020133-1222211030213212-2130301102101332-3000302311031331-0120230123000000-1031031013102313-0313202122111021-3323321231122330"></a>

## tiering_policy property — volume_defaults / 003303100320 / 14

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

<a id="canonical-3333032202313302-3232311230323102-0130002213122033-1002321012023232-3233102321222120-1230131003203310-0220111330012033-3033031132022130"></a>

<a id="canonical-1131231003032130-3333310123001123-3320001122323302-3021123232131302-1303331013302233-1222322003231003-1302113102301200-0311133312011213"></a>

## unix_permissions property — volume_defaults / 003303100320 / 15

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

<a id="canonical-2322303223021022-2030021330312323-3030201300321031-3112301112120003-2301012023132110-1200312001020220-1130212222102102-0113102302203101"></a>

## Next pages — volume_defaults / 003303100320 / 16

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](data-sources--fleet--reference--group-004.md#canonical-0112010323101023-2331002330103032-1223131001001030-0221301023032130-0013032111330233-1231122211211221-0032203001033102-2302001301303210)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0112010323101023-2331002330103032-1223131001001030-0221301023032130-0013032111330233-1231122211211221-0032203001033102-2302001301303210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003110003121012-1323200001301221-1323331022233223-2120310323301102-1303233212301222-3000222120001310-1233230221332330-3132221212311113"></a>

## storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos — no_qos / 113022313310 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.netapp_trident](data-sources--fleet--reference--group-003.md#canonical-3021312102012130-3321123001122010-2000002210023333-1120212201212033-3201013023101330-0211123300211123-0032120323300320-0031033232231321)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--fleet--reference--group-003.md#canonical-0223322303130221-1133222011231031-1230103011210112-0000011301133323-2032213310102002-1001222330231002-2331301223111213-2031331232310330)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--reference--group-004.md#canonical-0312231301000023-1223223123130323-1300021202220113-1200230003223230-0121123321012302-1111201303220322-3001223112113203-2120323230132130)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos

<a id="canonical-1103111130302023-0013203012111220-0020223120232110-0230113320122320-0310232321323221-3302022113031001-1123232211313310-3113322022133202"></a>

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

<a id="canonical-0221231211012312-3033301113011330-3222102321123012-0201203221202131-3003022002020132-0131122320000213-0203223011223320-2223021102311021"></a>

## Direct properties — no_qos / 113022313310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221130310102123-1230123310101230-3130112200230210-0131312111221130-3211201031331321-3213220331101112-2313001020011331-2113102222101032"></a>

## Next pages — no_qos / 113022313310 / 4

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--fleet--reference--group-004.md#canonical-0312231301000023-1223223123130323-1300021202220113-1200230003223230-0121123321012302-1111201303220322-3001223112113203-2120323230132130)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012322311113230-1323012232331202-1210232231323222-1013201033223223-0132010202201222-3322313332310320-2210312132102123-3111102001302222"></a>

## storage_device_list.storage_devices.pure_service_orchestrator — pure_service_orchestrator / 213230113103 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- storage_device_list.storage_devices.pure_service_orchestrator

<a id="canonical-3103200021113021-2113212020022021-3033203132002300-2030331010033303-1110331103303232-2221010220120101-2001201013001121-1002013130103233"></a>

Type: `"single"`. Computed.

Device configuration for Pure Storage Service Orchestrator.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223201301232000-3013300023110300-0311213211321330-2322101121120303-2213022011323022-0130211322030312-2310030213203100-3112130301131222"></a>

## Direct properties — pure_service_orchestrator / 213230113103 / 3

- [arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223): complete subsection reference.

<a id="canonical-2101031202010310-1031221221321122-0112033111020322-1313301032211103-1233131223333332-0202020103320112-0202332120012200-2220202203231332"></a>

<a id="canonical-0130230010312302-3000332230132312-2121120132023210-0321020210310221-0013331320301302-1230013100121321-0201233311320202-3132320032122122"></a>

## cluster_id property — pure_service_orchestrator / 213230113103 / 4

Type: `"string"`. Computed.

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays.

Upstream description:

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and
underscores.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 22,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 22,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9_]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  }
}
```

<a id="canonical-2100311003310111-1001222033122302-2211102021322333-1223100012322100-0233230012221110-3002332033331023-3321111002210030-3330021013101222"></a>

<a id="canonical-0003131202030103-0302021232133332-0031330231203333-0021001112212132-3310202201210313-0303210312131322-2011123201321220-2313111202201300"></a>

## enable_storage_topology property — pure_service_orchestrator / 213230113103 / 5

Type: `"bool"`. Computed.

Option is to enable/disable the csi topology feature for pso-csi.

Upstream description:

This option is to enable/disable the csi topology feature for pso-csi.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1323013031203322-3200122311011320-1310311133223331-2232312032311312-1122300022321220-3123113032020031-2122031111110210-1303021330101313"></a>

<a id="canonical-1033020102201110-3203310033032103-0030232232133300-0111023030331203-2221030101211012-2121111232100223-1013223122311101-2103003112032311"></a>

## enable_strict_topology property — pure_service_orchestrator / 213230113103 / 6

Type: `"bool"`. Computed.

Option is to enable/disable the strict csi topology feature for pso-csi.

Upstream description:

This option is to enable/disable the strict csi topology feature for pso-csi.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320333001300212-0012312013330333-0122202302220211-3332220200001302-0232221212031310-2122113103230201-2210211321300023-3020303130212231"></a>

## Next pages — pure_service_orchestrator / 213230113103 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130232020111113-2031301113101320-1300110202202303-3021111312131022-1010303012133011-0110211110101030-3102201133113122-2302312130122133"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays — arrays / 102320030111 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="canonical-1220133332301021-2220302313013310-2232231112111331-3331213003133231-0002230133132323-0301013231023201-3323320200020303-0002121312333132"></a>

Type: `"single"`. Computed.

Arrays Configuration. Device configuration for PSO Arrays.

Upstream description:

Device configuration for PSO Arrays.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1201331210331130-0320132103211223-3302023222121132-3320220303200233-1232301212120321-2330320203322300-0233212011010023-3001303000310323"></a>

## Direct properties — arrays / 102320030111 / 3

- [flash_array](data-sources--fleet--reference--group-004.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313): complete subsection reference.

- [flash_blade](data-sources--fleet--reference--group-004.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303): complete subsection reference.

<a id="canonical-2112132020220130-0231013102310011-1301212032123203-0203000132302302-1203303330021112-1012203111120313-3302100300102323-2003003012133003"></a>

## Next pages — arrays / 102320030111 / 4

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320302230001202-3322223233321100-1030330010311222-0220313201230322-3030121112132301-3211230001101211-0313322330230122-1020030000202223"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array — flash_array / 012233302103 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="canonical-0302201022230330-0012011321221330-0120310021103012-0300303211221032-0002221201013200-0302113310212312-0101323330202020-2223032211113330"></a>

Type: `"single"`. Computed.

Specify what storage flash arrays should be managed the plugin.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3101300303222121-0321203322002300-3311112012232211-3023023200202203-0213211201210023-1023032221303333-2113302120332033-3023223113003113"></a>

## Direct properties — flash_array / 012233302103 / 3

<a id="canonical-2131223011212022-0110323200102202-3102103122033213-1220231130221110-0000320233110000-0232013123213010-0031221003021212-0011211101013011"></a>

<a id="canonical-2000000131100020-3210201312320232-3230022212211110-1120123221030130-2001201111223232-3320202021002110-1001110232313331-3133200123010120"></a>

## default_fs_opt property — flash_array / 012233302103 / 4

Type: `"string"`. Computed.

Block volume default mkfs OPTIONS. Not recommended to change!

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

<a id="canonical-1130112130001332-3011220300332020-3031033330012131-3111221300123102-0111120122221201-3010230320223201-3310302032111100-3223203232330311"></a>

<a id="canonical-3321320113020113-3101200010333112-0220303200313013-3122133133101310-2102030123322223-2131312102323311-0002032201221321-0110112323311010"></a>

## default_fs_type property — flash_array / 012233302103 / 5

Type: `"string"`. Computed.

\[Enum: xfs|ext4\] Block volume default filesystem type. Not recommended to change!. Possible values
are \`xfs\`, \`ext4\`.

Upstream description:

Block volume default filesystem type. Not recommended to change!

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "xfs",
    "ext4"
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
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  }
}
```

<a id="canonical-2111332201122310-1111021121110001-1020230131023332-0111303001130013-1321332210011010-3011021231333313-2330002203310222-2233032111212102"></a>

<a id="canonical-2033223211101023-1132200120022110-2201012322133332-2130132232130331-3333221102302333-1311231123313003-3233112112231010-0121102110301112"></a>

## default_mount_opts property — flash_array / 012233302103 / 6

Type: `["list", "string"]`. Computed.

Block volume default filesystem mount OPTIONS. Not recommended to change!

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312021032030033-0330322302300111-1231130221323323-1203030311310311-1310320211220302-0201200231101221-3301133210231120-3012120331002102"></a>

<a id="canonical-1002000233231232-2111231111020223-1022120333201221-3020112003331031-1002021300011222-0023031003010023-1030120212011113-2130023003010221"></a>

## disable_preempt_attachments property — flash_array / 012233302103 / 7

Type: `"bool"`. Computed.

Disable Preempt Attachments. Enable/Disable attachment preemption!

Upstream description:

Enable/Disable attachment preemption!

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [flash_arrays](data-sources--fleet--reference--group-004.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102): complete subsection reference.

<a id="canonical-3112101121232013-1311320131013102-1213323211210312-3003330003321133-1023101001021133-1032312121322322-2220212123121302-2303030132133312"></a>

<a id="canonical-0202112003311320-2201233332322230-3103300212220113-2112032110103022-3203023210211032-1201333222003030-1110003311330113-0123033330102230"></a>

## iscsi_login_timeout property — flash_array / 012233302103 / 8

Type: `"number"`. Computed.

ISCSI login timeout in seconds. Not recommended to change!

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0302122123033100-0300322222122110-0131213232020222-0100030303311300-2100323000202203-0301101023301203-3112320301113213-1000301301033213"></a>

<a id="canonical-3102312011332120-2213323101221021-0111232030201231-0212300001002022-0102311230130012-2120301232231323-3230211133000022-2203231022312222"></a>

## san_type property — flash_array / 012233302103 / 9

Type: `"string"`. Computed.

\[Enum: ISCSI|FC\] Block volume access protocol, either ISCSI or FC. Possible values are \`ISCSI\`,
\`FC\`.

Upstream description:

Block volume access protocol, either ISCSI or FC.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ISCSI",
    "FC"
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
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  }
}
```

<a id="canonical-2132123012202303-0201202312112200-1010120202021021-3112031222030201-2020232300110300-2211112000232132-0120133333322133-0233303213133111"></a>

## Next pages — flash_array / 012233302103 / 10

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103120103112332-3110333122331132-2223310021122202-0001111212302010-0020123323103000-2113020133303323-2133301002001233-1321213001101020"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays — flash_arrays / 221100321201 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays

<a id="canonical-3103113323013021-2013123110001301-0032021233130100-0021032201110132-2311200320133212-3200330121131233-0010210220220301-1222312111012121"></a>

Type: `"list"`. Computed.

For FlashArrays you must set the 'mgmt\_endpoint' and 'api\_token'.

Upstream description:

For FlashArrays you must set the "mgmt\_endpoint" and "api\_token"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333123121311212-1021331101110000-2312232221321010-1021023233121322-0331302002213320-1331310010011011-2311300021303123-3133132310321223"></a>

## Direct properties — flash_arrays / 221100321201 / 3

- [api_token](data-sources--fleet--reference--group-004.md#canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321): complete subsection reference.

<a id="canonical-1112222310111031-2023132301131211-2021102110323103-2132021132303230-0011321132112203-0332222000221312-0313301331033311-3031333302322022"></a>

<a id="canonical-1233131233301300-0220133131022230-2223011301212210-1013102320100101-1033023320332021-1120132322031230-3312331213021210-0003013301131103"></a>

## labels property — flash_arrays / 221100321201 / 4

Type: `["map", "string"]`. Computed.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Upstream description:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

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

<a id="canonical-0200001223210200-3012103103032021-1001302330322103-2310112332122211-0213121022001203-0131313323330230-1023131312323120-1130002320320110"></a>

<a id="canonical-3123332322003322-0302301201022221-0211310303013332-1320020203132220-1020010212220202-1203101100132220-3132000122031202-0302302013320102"></a>

## mgmt_dns_name property — flash_arrays / 221100321201 / 5

Type: `"string"`. Computed.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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

<a id="canonical-1002222202122210-1022200331213131-0330331202330031-2101113133230232-3002120313020201-1022032203033121-3210010302310020-2012130033302321"></a>

<a id="canonical-3013202220101001-1013130223101212-3320201302211111-3221332232100222-0310003131030330-0123310322020122-0102221233022202-0303020323331100"></a>

## mgmt_ip property — flash_arrays / 221100321201 / 6

Type: `"string"`. Computed.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

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

<a id="canonical-1123032122013121-0031030323302230-1033021300122310-3332112313120303-3332010032101233-0122222231233113-2320223300001123-0311331323030031"></a>

## Next pages — flash_arrays / 221100321201 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311301022320301-1130023202332110-2102120111033001-2200202322311303-1221231213200021-3000231110303321-0120103111322313-1330333222311201"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token — api_token / 232021213220 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token

<a id="canonical-0312220003010031-3202200103111110-2313231221013322-0223033121020023-2122311002211321-1000000122320110-1331102130132220-3302030133210122"></a>

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

<a id="canonical-1210022210302211-0130311303233222-2203012123111010-3322203000233113-1232311233332323-0002300301211210-1001311012000231-0030120003000103"></a>

## Direct properties — api_token / 232021213220 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-0101333221021310-3111321213132322-1332102111332101-2000231131021001-3230121032302101-0131122211220312-2021023333201121-2310302303030320): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-3330220222201321-1021033010023002-1000120201020120-2012312101103123-0200202233223112-1322120311312223-2332003311111013-0113001001311222): complete subsection reference.

<a id="canonical-1300211203030212-0001003332301100-0331203312023030-0221103223021233-0001330220020111-1222212010330210-3101301010233032-1022012310210011"></a>

## Next pages — api_token / 232021213220 / 4

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-0101333221021310-3111321213132322-1332102111332101-2000231131021001-3230121032302101-0131122211220312-2021023333201121-2310302303030320)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-3330220222201321-1021033010023002-1000120201020120-2012312101103123-0200202233223112-1322120311312223-2332003311111013-0113001001311222)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0101333221021310-3111321213132322-1332102111332101-2000231131021001-3230121032302101-0131122211220312-2021023333201121-2310302303030320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233320310002301-2010300131203322-0300003330222203-2120100132301002-0121223331233031-3330302020230311-1103010310222201-2231301301323023"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info — blindfold_secret_info / 223220031232 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="canonical-0303332331030102-3213230100301230-2102130323020100-1223212223021120-3213212310022210-0133322103031123-1230212210233232-1102330200231000"></a>

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

<a id="canonical-1310220002012122-2033102103332030-1231112231231011-1332122013320310-2321001333220212-3103102131022201-1221122211202303-0102131211131030"></a>

## Direct properties — blindfold_secret_info / 223220031232 / 3

<a id="canonical-1301213221022011-2231221221103202-3000112200213012-2231000223301013-1021212033020000-0233132122102103-3112003130032300-1031203031001113"></a>

<a id="canonical-0313303012323030-0312103333220020-0202132032013021-2232333030211323-2331120112212122-2133202222221332-1132332012012121-0023233032102001"></a>

## decryption_provider property — blindfold_secret_info / 223220031232 / 4

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

<a id="canonical-1120002232121121-0020133330312102-1233002123022023-2222100012031210-3333012000012201-2312103322322322-2001001013303301-0320031321323322"></a>

<a id="canonical-1321230011111321-1120223130311230-1100120310333322-1003112031110312-1012011233123100-0121121230213212-3302333013323123-2312300210300220"></a>

## location property — blindfold_secret_info / 223220031232 / 5

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

<a id="canonical-1112231322311101-0130133230303320-1202231133221001-0131233031230031-3112310032112301-0322201122320222-3112130212301223-2021001220103100"></a>

<a id="canonical-3131233230322312-3002203232012022-3000000312200301-3111300101231011-2110221013201220-2302013202202032-0310111220101103-3321213223310330"></a>

## store_provider property — blindfold_secret_info / 223220031232 / 6

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

<a id="canonical-1211210321002233-0230121323113310-1313333211311132-2330130200333113-0321202121310100-1003323013002032-1010320023020121-1213012001212111"></a>

## Next pages — blindfold_secret_info / 223220031232 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3330220222201321-1021033010023002-1000120201020120-2012312101103123-0200202233223112-1322120311312223-2332003311111013-0113001001311222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213303001212122-2033013233300030-1130221020213230-0210213011001331-2333112110001133-1212131323331131-1122113011011231-1020321333103003"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info — clear_secret_info / 311122012123 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--fleet--reference--group-004.md#canonical-3100333230102320-2201210300121102-1332331302002131-1233110231203031-3013323101032112-2023313222111010-1131233313003323-0221211230200313)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--fleet--reference--group-004.md#canonical-3302231123231130-1200133220000322-1120123000331211-2120021120300333-0312210101021321-3311211303232112-3012230333112203-3302201320332102)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info

<a id="canonical-0021013322031303-1233222001012201-1212301210302101-2132101033133303-2030233032122331-1202122300122130-3001002022322231-1131000320223323"></a>

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

<a id="canonical-3303222333233310-3232032030331111-0013011013202233-2222330332100213-2110201222010123-1313020113231100-2322132221102311-3230211211310110"></a>

## Direct properties — clear_secret_info / 311122012123 / 3

<a id="canonical-0233021131130012-1003022133313212-2323213030011331-0103220221112132-2101301100023102-0300300203023310-0012332223131302-2031201303321010"></a>

<a id="canonical-1102122311131112-3022201000103020-0113121231123202-3321311331221000-0231030202000212-1112202101003013-0013322331033331-0302023033123313"></a>

## provider_ref property — clear_secret_info / 311122012123 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3231022313133212-2310231300232001-3103202220310300-3110013003310323-1122322232300132-2000030202322112-0312013133130130-1022013302120000"></a>

<a id="canonical-2232220003200311-1000330002221103-0230103133311113-3100301131221232-2332121100230101-1113013112201303-2001222210320213-0112132021232300"></a>

## URL property — clear_secret_info / 311122012123 / 5

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

<a id="canonical-3012210313301021-0110203233130213-1110211131313130-3010032320102323-2021321300320003-3020022001021333-0102321112031100-3031123321122202"></a>

## Next pages — clear_secret_info / 311122012123 / 6

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--fleet--reference--group-004.md#canonical-3321023201313323-3303221121321200-2200021032301011-3013011202132111-3220321023021033-2022001332013210-2231012012223321-3222320330103321)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100332111313031-3230033321232130-3302020322210023-3310030032330012-2120332212002302-1331223222010210-0230120010101011-1021022013013203"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade — flash_blade / 201102201201 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="canonical-3301212300021000-3022111311233332-3222331022202022-3011032333032301-1301213320333212-3331122330001310-3311232031032221-1313100121032322"></a>

Type: `"single"`. Computed.

Specify what storage flash blades should be managed the plugin.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3323013131030030-1201331313300312-0100332100203333-0013010100121002-1032031121113312-2112333030310302-0301210300110301-0010001230231313"></a>

## Direct properties — flash_blade / 201102201201 / 3

<a id="canonical-0023111222020120-2130321133020000-1031123032323023-2302302220113121-2102102203132201-1202302320212330-1020122100230011-0030123122231300"></a>

<a id="canonical-3331222021012211-2001213312310001-2222232311320303-2103011020101113-1223010323222013-0322000112003012-3333033202210230-3302302023310331"></a>

## enable_snapshot_directory property — flash_blade / 201102201201 / 4

Type: `"bool"`. Computed.

Enable Snapshot Directory. Enable/Disable FlashBlade snapshots.

Upstream description:

Enable/Disable FlashBlade snapshots.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233020311321021-3333233212201203-0110201322310133-1232233133231221-1013131132302333-0333213300033002-0101200203201323-0231123121011010"></a>

<a id="canonical-1010031102330300-3231023203010130-1100112012011303-1033031213223212-3330221110020130-0312220201131000-1210201312322011-3233333210031012"></a>

## export_rules property — flash_blade / 201102201201 / 5

Type: `"string"`. Computed.

NFS Export Rules. NFS Export rules.

Upstream description:

NFS Export rules.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 250,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 250,
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
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [flash_blades](data-sources--fleet--reference--group-004.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330): complete subsection reference.

<a id="canonical-0000002331130333-0113323121111210-0033232203121120-0002103300030002-2301031223020211-3220323002312201-1003313023321033-2200210033331002"></a>

## Next pages — flash_blade / 201102201201 / 6

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021013331320010-3111333332121013-1112320131223031-2123012320030331-3032100123130031-3120011012130230-2122001103320222-1102101033102032"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades — flash_blades / 321203303122 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="canonical-3113322203312112-2320103102230022-3101033023320312-2123121313032122-2320123021333000-2221033111110100-2303310202201022-2321230333121021"></a>

Type: `"list"`. Computed.

For FlashBlades you must set the 'mgmt\_endpoint', 'api\_token' and nfs\_endpoint.

Upstream description:

For FlashBlades you must set the "mgmt\_endpoint", "api\_token" and nfs\_endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3221223202021003-3330311031310312-1001110311103230-0023221210311310-2133020202333110-1301300102301203-0121230321011302-2020202133030021"></a>

## Direct properties — flash_blades / 321203303122 / 3

- [api_token](data-sources--fleet--reference--group-004.md#canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223): complete subsection reference.

<a id="canonical-2333020233221223-1203231313233301-0011030002121232-1321030012230003-0212131322032202-3203121132213300-3333002320200322-0111332122332113"></a>

<a id="canonical-3223123232323132-2222020220301001-0300111221022221-1230200103022132-2330032222333220-0123223203000323-1003230101131111-1133331021132333"></a>

## labels property — flash_blades / 321203303122 / 4

Type: `["map", "string"]`. Computed.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Upstream description:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

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

<a id="canonical-0122232110112012-3300332203201123-3010320001220022-2132312332032022-3203111132202033-0113021022002101-2230213123302100-3110131023000101"></a>

<a id="canonical-0320030311021031-2021133113323231-2232310103220033-2210301102312131-1233002212322033-0023300211212010-3021101213022001-1222000230131133"></a>

## mgmt_dns_name property — flash_blades / 321203303122 / 5

Type: `"string"`. Computed.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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

<a id="canonical-3101133123322320-3112101220012312-0100301313203222-0313121231233223-1132123032233133-2130031233203210-0120301230333002-3013022301113032"></a>

<a id="canonical-1302000022123003-2120302230020230-0322103123301112-1330200200211103-3120033213031222-3131211003320321-2201312132222030-1133232101122010"></a>

## mgmt_ip property — flash_blades / 321203303122 / 6

Type: `"string"`. Computed.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

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

<a id="canonical-3131003020012111-0320223201311102-2313012303001211-2322001033221203-3311332031133023-3010221210303013-2020220111103120-2221022020202032"></a>

<a id="canonical-2133033220023022-1123222110113130-1013103233122211-2223302113011302-3332101003220001-1102232333131002-1113123332032231-1220101300110021"></a>

## nfs_endpoint_dns_name property — flash_blades / 321203303122 / 7

Type: `"string"`. Computed.

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

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

<a id="canonical-1232001313211322-3320331330033020-2121120223221232-0133312101320332-2033032303030333-0023311101013021-0300012310331231-1023033220223132"></a>

<a id="canonical-1002010022211111-3023013310100333-0312011333102111-0230320332032302-3331213110120112-2222013011220121-1310220122132023-2130012123010001"></a>

## nfs_endpoint_ip property — flash_blades / 321203303122 / 8

Type: `"string"`. Computed.

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

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

<a id="canonical-3330003232013013-3300100003120221-2132233102213301-1030001031133120-1223033000312301-3311301130020031-1132010232222111-0223310120023123"></a>

## Next pages — flash_blades / 321203303122 / 9

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101310111211122-2232121101112312-0330130110130003-0202131130023112-0033323123033012-0331201330310003-2010120303213103-0012311313221010"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token — api_token / 121010023333 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="canonical-0021033303010333-2212110030332110-1021101310013202-3130011001030221-1230101032033020-2010231023223021-1331300022303313-2002202112102001"></a>

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

<a id="canonical-0003111013001111-2230330120312022-2111122103021103-0220022100123002-1110032323331012-0313011322032211-0222321232310220-3222223221101320"></a>

## Direct properties — api_token / 121010023333 / 3

- [blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-2333102332103103-1110001013122020-1102332123133211-2133121132100030-1100202030011331-3323300101312320-3013011312302131-1020010311121322): complete subsection reference.

- [clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-0322013121102330-2231012313300203-1112311002132201-3321010010320213-1221030023112200-0021012100023330-0130023031203211-2322330031223302): complete subsection reference.

<a id="canonical-0332322300330203-3320102311133222-1320221110102211-0020301002113210-2111121311013301-1232223310012133-1130223100310102-2320101113233333"></a>

## Next pages — api_token / 121010023333 / 4

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](data-sources--fleet--reference--group-004.md#canonical-2333102332103103-1110001013122020-1102332123133211-2133121132100030-1100202030011331-3323300101312320-3013011312302131-1020010311121322)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](data-sources--fleet--reference--group-004.md#canonical-0322013121102330-2231012313300203-1112311002132201-3321010010320213-1221030023112200-0021012100023330-0130023031203211-2322330031223302)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2333102332103103-1110001013122020-1102332123133211-2133121132100030-1100202030011331-3323300101312320-3013011312302131-1020010311121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131120121313122-3111101012331230-0030112000323212-0230103222312323-1222223211010323-2320132032120202-2312033021013230-3132100021002313"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info — blindfold_secret_info / 320133121022 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info

<a id="canonical-1112133021203233-0120323301013121-2310320323113212-1213322002201232-1201210133330322-3031022123330112-3200321232300102-1322202130011321"></a>

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

<a id="canonical-1200122200202001-2022312301023310-0112310021103312-3012302310200112-1032021001200022-3001312023003321-1300302223132022-3111231022013312"></a>

## Direct properties — blindfold_secret_info / 320133121022 / 3

<a id="canonical-3200121032302100-1223033021102330-0212023000201132-2003102322122100-0121002101021123-0121131122210223-2000001100131110-3020130000303330"></a>

<a id="canonical-1312213332331300-0030322313021211-2131331313131332-3312231311211211-1132200103311032-1020231023003023-3002223031120013-3212033300213012"></a>

## decryption_provider property — blindfold_secret_info / 320133121022 / 4

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

<a id="canonical-2033022230311112-3332001132303112-0011201221103031-1212103232021201-1031313223211322-0231311213203213-1010110221330213-2213131120022331"></a>

<a id="canonical-3323020213110210-1003303011303121-2223111133300132-1113000120330200-1001233303222122-1131012301123212-0331222332231010-3321223113211111"></a>

## location property — blindfold_secret_info / 320133121022 / 5

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

<a id="canonical-3333203001133002-1213102111310013-1302100000233223-2103221320222321-1201322011130331-1132032220331213-2012211200100232-0031323033000032"></a>

<a id="canonical-2221023002001100-1222122011113202-3223100312111012-2131023231102332-1101302011212210-0202223311030021-2300130233033313-0100231230200312"></a>

## store_provider property — blindfold_secret_info / 320133121022 / 6

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

<a id="canonical-2122203030303203-2300103120110131-2223100110111132-0102130120032301-3221100001213110-1333222122201202-3101100003132110-1331110201212000"></a>

## Next pages — blindfold_secret_info / 320133121022 / 7

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0322013121102330-2231012313300203-1112311002132201-3321010010320213-1221030023112200-0021012100023330-0130023031203211-2322330031223302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200332203033220-2220302322013012-3320223011310231-2130200120120012-1232020001120200-0121231332123213-0020320023311212-2001113311213020"></a>

## storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info — clear_secret_info / 022002300101 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_device_list](data-sources--fleet--reference--group-002.md#canonical-2203120030033203-2032323000023021-1211113020233233-3010321330312301-3231230033331312-3132222101232013-1033011222003031-1331202320031021)
- [storage_device_list.storage_devices](data-sources--fleet--reference--group-002.md#canonical-2032223221033232-1222013030202102-0212323232123023-3231330112103313-1020231233021030-0133312302201010-2330111032210001-2111001101130211)
- [storage_device_list.storage_devices.pure_service_orchestrator](data-sources--fleet--reference--group-004.md#canonical-0232331033131323-1012000032222131-3103231213311210-3202202123003123-3003333321120112-1102302232131101-2330131202301013-1332210310100022)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--fleet--reference--group-004.md#canonical-2022021200123133-1012311022310131-2011001032320220-1011211211021301-3032232301200012-2111121022022131-3102103001000323-2322001003132223)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--fleet--reference--group-004.md#canonical-2032231100232113-1121332300210032-2300122032301201-1022223222230332-0320022001320313-0322210021221233-1130002200302111-0110333133311303)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--fleet--reference--group-004.md#canonical-3330123332213122-0311333303020203-2100301133303321-2131232010222233-1100023113000033-3113311212320100-1102102332012100-0222103102231330)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info

<a id="canonical-3003220022210311-0331322102200122-2321001333331321-2301120010110122-0021231013302303-3230323001313012-3210201221000130-2201031121321213"></a>

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

<a id="canonical-0132002002331111-0322322031112312-2332001022013100-2311023023312033-0310000020223112-3133320312030133-3330232303133311-0001221012321221"></a>

## Direct properties — clear_secret_info / 022002300101 / 3

<a id="canonical-0121230202312130-2223021112300133-1102032031212112-0322322113100211-1001030211122201-2011321022333332-2003312022331102-3330021111013212"></a>

<a id="canonical-3021233320321102-1120012132333321-1312032132200021-0033031031120002-2033132102122233-3223110211112011-1301220000331013-0111203332222020"></a>

## provider_ref property — clear_secret_info / 022002300101 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0230221020030020-3311113313312211-0101000331332133-3223033303212102-1311303002131123-2111313230332302-0132312032231113-2203102321221102"></a>

<a id="canonical-2222132222130031-2233131210012200-2222222123311232-2020002022223131-1023310313200211-2000100003230013-0100002313001233-0331112221012033"></a>

## URL property — clear_secret_info / 022002300101 / 5

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

<a id="canonical-3000111323020000-1121112130013023-1132022212032100-0021202201020202-1010010310102330-3202021110221120-0013103223200321-2301303211010133"></a>

## Next pages — clear_secret_info / 022002300101 / 6

- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--fleet--reference--group-004.md#canonical-0120031312202210-1213230232032302-3221130111010021-1123100131110111-2100101020033100-0310021321221020-0012013030020321-2020100103212223)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1332201313101100-2300203232300003-3203012102133211-1111001020322331-3011023323013302-0102200332002030-0003133323233202-2131113031033200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020030020103323-2021113002321103-0230130232123121-2130203122102320-0203211221113100-2102301301121223-1202002023303131-2032032310002233"></a>

## storage_interface_list — storage_interface_list / 123210111322 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- storage_interface_list

<a id="canonical-2121301312201313-2003303021323200-1113203311322120-1012103002103230-2013220110320003-0220331131210223-0321010211132303-2320301002023322"></a>

Type: `"single"`. Computed.

Add all interfaces belonging to this fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1202011322223122-3023200000131103-1132313201332202-2031023310332101-0033001223120131-3121201331110323-3122123310331300-0232111331020003"></a>

## Direct properties — storage_interface_list / 123210111322 / 3

- [interfaces](data-sources--fleet--reference--group-004.md#canonical-2033322200112100-2130213320101310-2230111211133032-2103331223111023-0200333312020201-1020302212030332-0231123132030110-2203122202333330): complete subsection reference.

<a id="canonical-3212011003101031-1233201311130120-1010211003232232-3230011113131200-2100023333201020-2211231310000113-3032112231022013-1121012202103231"></a>

## Next pages — storage_interface_list / 123210111322 / 4

- [storage_interface_list.interfaces](data-sources--fleet--reference--group-004.md#canonical-2033322200112100-2130213320101310-2230111211133032-2103331223111023-0200333312020201-1020302212030332-0231123132030110-2203122202333330)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2033322200112100-2130213320101310-2230111211133032-2103331223111023-0200333312020201-1020302212030332-0231123132030110-2203122202333330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203111313210223-2213332322103331-2200323020111133-0320332111212212-3210032223232301-2302213223023223-0200130131330111-1221121012220020"></a>

## storage_interface_list.interfaces — interfaces / 300100321023 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_interface_list](data-sources--fleet--reference--group-004.md#canonical-1332201313101100-2300203232300003-3203012102133211-1111001020322331-3011023323013302-0102200332002030-0003133323233202-2131113031033200)
- storage_interface_list.interfaces

<a id="canonical-1122201013213312-1213321213321112-2333130011001003-0003030300101031-2312332033001313-3003313032230033-1202330033301311-2111302313321313"></a>

Type: `"list"`. Computed.

Add all interfaces belonging to this fleet.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3010031132122100-0130310130000010-3011022231321322-1330122032012312-1103203100312223-0133321220322023-0331303303321321-1230122231211111"></a>

## Direct properties — interfaces / 300100321023 / 3

<a id="canonical-0212332102111200-3320032000321203-1202332311211102-0201332310320001-2123322132033202-0311313020113102-2231302303222121-1211123330131031"></a>

<a id="canonical-3133321030211233-1012112302110312-1123023332202321-3310222101223010-3310123203313301-3103121022213103-1003230203101202-0002103021133002"></a>

## name property — interfaces / 300100321023 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3301020201020100-2010220120303110-0100133332112211-2322233313332201-2310030111130300-2030312131130013-2202002101112113-0302100130332202"></a>

<a id="canonical-1013230233112001-0203312020302031-1122000102220110-3300332003211123-1031123123213031-1010200213112201-1322200323330102-1310330032310231"></a>

## namespace property — interfaces / 300100321023 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3101313223200101-0132311203200313-1322123133110022-3230123003300203-2312213320223020-0123013210323210-3233330203023100-3222222230220121"></a>

<a id="canonical-3231202122230330-1022003130210310-0200200130121030-3112233101030320-1002001211103332-0323201111111113-1122330002301221-3213330102132303"></a>

## tenant property — interfaces / 300100321023 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1101321122021001-2311103231303013-3323000030313013-0123120123101100-3312000322120103-0301211120131323-1130002330121221-3121032000022321"></a>

## Next pages — interfaces / 300100321023 / 7

- [storage_interface_list](data-sources--fleet--reference--group-004.md#canonical-1332201313101100-2300203232300003-3203012102133211-1111001020322331-3011023323013302-0102200332002030-0003133323233202-2131113031033200)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010301101233302-3011020210223321-3013132323001000-3313321130303220-3213131020232232-2132301223210102-1202020111101132-3221101121000131"></a>

## storage_static_routes — storage_static_routes / 112131133001 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- storage_static_routes

<a id="canonical-3012331233311111-0123111321131231-3213032311022211-3000231010022111-1001031021200012-0323022200021110-3333332210321120-3330100233233210"></a>

Type: `"single"`. Computed.

Configuration parameter for storage static routes.

Upstream description:

List of storage static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2131122111331020-0222202301120301-3010300310323031-3200330333003112-1303211221121021-3232202102211210-2000000001222023-1322131111233200"></a>

## Direct properties — storage_static_routes / 112131133001 / 3

- [storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110): complete subsection reference.

<a id="canonical-2300313200203222-1200312333130133-0121120230023331-1011120310101031-3330330030110020-3023231021321000-2232003021221002-1013330230202220"></a>

## Next pages — storage_static_routes / 112131133001 / 4

- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330222132012233-0313302200032001-1321332010223330-2222100013001112-2020301100213020-0133103120012020-1020231212120302-0330200323330211"></a>

## storage_static_routes.storage_routes — storage_routes / 320301210302 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- storage_static_routes.storage_routes

<a id="canonical-0002121312211212-2131010023231311-2303033332000203-0203113331203322-3113121222031102-0211010112002032-0330023132100311-1113010230012303"></a>

Type: `"list"`. Computed.

List of Static Routes. List of storage static routes.

Upstream description:

List of storage static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1101033032321020-1022102133133233-1003313222032011-2003230131323211-2201210010211102-1223020022101011-2233320000002332-0033111100203302"></a>

## Direct properties — storage_routes / 320301210302 / 3

<a id="canonical-3310113203102321-0220331313021223-1210113113222021-0310011211133303-0111110313211010-0032313110102321-3010010021213202-2333111201101201"></a>

<a id="canonical-0013201232200032-2211221233312030-2312003000000220-1111120233102213-0223103212101002-0022131012222321-1111213231210323-0202032031323310"></a>

## attrs property — storage_routes / 320301210302 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--fleet--reference--group-004.md#canonical-2333322330302012-3111121113301032-1013103232322112-0231022000222301-2313120302310130-2020010130032321-0112300013211020-1302031132120021): complete subsection reference.

- [nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000): complete subsection reference.

- [subnets](data-sources--fleet--reference--group-004.md#canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303): complete subsection reference.

<a id="canonical-2330001321202310-2320230321320310-0032021131120111-1031102100021012-0232000122123101-3113130210133012-1320203222212111-0003113133213021"></a>

## Next pages — storage_routes / 320301210302 / 5

- [storage_static_routes.storage_routes.labels](data-sources--fleet--reference--group-004.md#canonical-2333322330302012-3111121113301032-1013103232322112-0231022000222301-2313120302310130-2020010130032321-0112300013211020-1302031132120021)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2333322330302012-3111121113301032-1013103232322112-0231022000222301-2313120302310130-2020010130032321-0112300013211020-1302031132120021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211131331333333-3100002221220333-3210330032012100-3002133110200023-2120332020220310-3032022021211112-0020310113332202-0013320102123320"></a>

## storage_static_routes.storage_routes.labels — labels / 020333322233 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- storage_static_routes.storage_routes.labels

<a id="canonical-3103023220123112-1331103110032301-1100200232101013-3201303231110003-3013312332030102-3302312322303003-1232210221131112-0212323010101333"></a>

Type: `"single"`. Computed.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0011200200333120-3203010002312101-3113000010320101-2221102302331331-3101113200230000-2010321211023300-3001222100233000-1013332230333230"></a>

## Direct properties — labels / 020333322233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032322301000113-3302010002130130-1110103330112221-1110303330210133-1300322000120301-3003321010221303-1323120131022100-1002020211030103"></a>

## Next pages — labels / 020333322233 / 4

- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222203102100133-2303233202321200-3010121131100212-3111021310320120-2130121302112021-2332123302110000-0011013332222130-3113011201301131"></a>

## storage_static_routes.storage_routes.nexthop — nexthop / 102332113230 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- storage_static_routes.storage_routes.nexthop

<a id="canonical-1332000133212021-1023230000201303-0003013230221303-0220302202132323-3232023313031201-2332132010133323-3301122210110313-3101120330300132"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1103213022311333-1312200331202303-0321220020300030-0001200331110012-3120003320103113-3111130010133333-0033313131212211-3313010111033222"></a>

## Direct properties — nexthop / 102332113230 / 3

- [interface](data-sources--fleet--reference--group-004.md#canonical-3000023121020110-0012320102123202-2331230130030331-1223220011332102-2133301101233130-3332220022023312-2032222003102112-3123121021332302): complete subsection reference.

- [nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312): complete subsection reference.

<a id="canonical-1130201332320032-2211130120201322-1313001230111021-1213111203222323-3030233222333221-3213301320002103-1002302121200331-2111221210121322"></a>

<a id="canonical-1113221122333100-3233200120203301-2113103231132003-1033021103011002-1100032202220212-2023202121133333-0000212132000032-2302001131000303"></a>

## type property — nexthop / 102332113230 / 4

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3003011320221132-2321232213030233-2230111100330323-0021310331121020-3212230322311120-3023101223013031-0121332133233123-1010233123112112"></a>

## Next pages — nexthop / 102332113230 / 5

- [storage_static_routes.storage_routes.nexthop.interface](data-sources--fleet--reference--group-004.md#canonical-3000023121020110-0012320102123202-2331230130030331-1223220011332102-2133301101233130-3332220022023312-2032222003102112-3123121021332302)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3000023121020110-0012320102123202-2331230130030331-1223220011332102-2133301101233130-3332220022023312-2032222003102112-3123121021332302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303231313011203-1132102101010223-1032131132110232-3320223031131020-1331321033202121-3300130102233303-1033013011123331-2013010102132332"></a>

## storage_static_routes.storage_routes.nexthop.interface — interface / 332310113323 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- storage_static_routes.storage_routes.nexthop.interface

<a id="canonical-1101320030003212-2103123221332333-0010002001331313-1232300103011123-1223231022210033-2010030223001133-3203133330300112-1112003030322011"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0202210033323031-3301033030103102-3122103321100033-1220132111223031-0021331200320330-3121221020201002-1021331112320100-0201320122012112"></a>

## Direct properties — interface / 332310113323 / 3

<a id="canonical-0221330321302132-3103132202133002-2123302200102120-1201032132113300-2230333012002331-1033033010200322-0021210303103330-2011101200020230"></a>

<a id="canonical-1103210302220220-2211111013230121-0311132000130120-1322301232000013-2013301121032201-0330002133331211-2310303001110302-1130122113233301"></a>

## kind property — interface / 332310113323 / 4

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

<a id="canonical-2212023012122320-3321321103022033-1201010103312111-0331321313230123-1012322300122323-0302112321100020-0002130110323230-2012022121320113"></a>

<a id="canonical-1212113210300031-1001321132311132-3313210333133020-0211122000122211-0003122311101311-3300321312301000-2223110201120010-3203020232102010"></a>

## name property — interface / 332310113323 / 5

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

<a id="canonical-1030310330232201-2233001002222303-2103102022310111-1311113102221330-0031330210101032-3223300030102220-2133220131110133-2211032323122200"></a>

<a id="canonical-1012301001021223-2203001122312112-3330113212113231-2013133201312323-3213002032200232-1231300012021020-2223112222232001-0022003111001011"></a>

## namespace property — interface / 332310113323 / 6

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1021030213203231-2301210022200121-3320002231001133-2101212100303013-0201113232130320-0023220100223220-1133002213000322-2022132303302202"></a>

<a id="canonical-1101130001223331-1203030002113210-1313311100000212-3121132023031032-1133210232001330-2202321033021323-3132300130031131-0032022213101213"></a>

## tenant property — interface / 332310113323 / 7

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

<a id="canonical-1001311231020131-1013001312320310-0013210013303131-0001300123233331-3023023102332011-1223101103222230-1010102020011321-2101323011120121"></a>

<a id="canonical-3130212212311022-0300003211330132-3233110330323000-3101131112010103-1311213220033232-0311332301121201-2323113320133303-1303211213123032"></a>

## uid property — interface / 332310113323 / 8

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

<a id="canonical-3033303001110323-1203133110132101-3120013200020213-3302232120223200-2232012032033113-1203011222023330-0011330310230101-2312213031311231"></a>

## Next pages — interface / 332310113323 / 9

- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322100033331133-3121031220002323-2101131332222332-3203213031010312-1103012133030310-3300331201303033-2302011332230031-3223111223021113"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address — nexthop_address / 020312031012 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- storage_static_routes.storage_routes.nexthop.nexthop_address

<a id="canonical-0323212123132300-1112331232012132-3220213323231021-3003031110230301-0232111220311022-3123030133110010-2112312012122202-3121103331200021"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-1323031301001223-1110000313221010-1003023332133211-2023121313221131-1113111201131200-0130120333321233-2100013032100110-0123133023002320"></a>

## Direct properties — nexthop_address / 020312031012 / 3

- [dual_stack](data-sources--fleet--reference--group-004.md#canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033): complete subsection reference.

- [IPv4](data-sources--fleet--reference--group-004.md#canonical-1203230330233330-0320331023002231-0012122030311100-3112313322120232-1203113223020201-3033112121033312-1031200202331201-2321200330101200): complete subsection reference.

- [IPv6](data-sources--fleet--reference--group-004.md#canonical-0300332322220010-0132002021012213-1323101311003021-1232221222232332-3223212333122100-3312321023311121-3323201230133223-2201330333332321): complete subsection reference.

<a id="canonical-2301102220221223-0210213303132223-0300123111313002-0020202101012331-2120013330113202-0213103230011033-1031223013002132-3333221101120133"></a>

## Next pages — nexthop_address / 020312031012 / 4

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv4](data-sources--fleet--reference--group-004.md#canonical-1203230330233330-0320331023002231-0012122030311100-3112313322120232-1203113223020201-3033112121033312-1031200202331201-2321200330101200)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.ipv6](data-sources--fleet--reference--group-004.md#canonical-0300332322220010-0132002021012213-1323101311003021-1232221222232332-3223212333122100-3312321023311121-3323201230133223-2201330333332321)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032131302021221-2200021221313020-3302100312000100-2021332100112023-0021222010011011-2332313021213302-2011232310311202-3133103123210012"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack — dual_stack / 311112311232 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack

<a id="canonical-1103200233003002-1221133113200013-2203032010230202-1301330010003110-1130233112112303-3311033311031021-1223022331201021-2210133303232312"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0122211021322321-1010302332331031-1210203001332313-0123321011031120-0200233110112131-3112100013322000-1100300201232211-3010312103123332"></a>

## Direct properties — dual_stack / 311112311232 / 3

- [IPv4](data-sources--fleet--reference--group-004.md#canonical-1211023231012201-1301320231100000-1321021223020323-3332012322121331-1302122322213022-2212230133103320-0320110133121011-2223212112002302): complete subsection reference.

- [IPv6](data-sources--fleet--reference--group-004.md#canonical-2011021111102013-3023110120202033-0331312021311220-1123332323023030-1220322032013232-2210203012032320-1210230133222310-0120020131233230): complete subsection reference.

<a id="canonical-3212111303131300-0111211132220102-3322301303311232-0123321123010330-2012311001032210-3212300013021212-1211112030033333-1112203012301332"></a>

## Next pages — dual_stack / 311112311232 / 4

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv4](data-sources--fleet--reference--group-004.md#canonical-1211023231012201-1301320231100000-1321021223020323-3332012322121331-1302122322213022-2212230133103320-0320110133121011-2223212112002302)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.ipv6](data-sources--fleet--reference--group-004.md#canonical-2011021111102013-3023110120202033-0331312021311220-1123332323023030-1220322032013232-2210203012032320-1210230133222310-0120020131233230)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1211023231012201-1301320231100000-1321021223020323-3332012322121331-1302122322213022-2212230133103320-0320110133121011-2223212112002302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030031111033213-1332312212200001-0200210201313230-1321200112322222-0132102211000102-1021003011321313-0102230100010310-3020312332123310"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 000020333211 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-3011310122320321-2331020130230233-2322331013201032-2302311201113132-3030132023113010-1032003010212323-3210313113330220-0302321003223023"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0012203312312032-2303113330312311-2212203110332112-1020020113300321-0103302110320331-0220100111330011-3233213002011020-2212310102212001"></a>

## Direct properties — IPv4 / 000020333211 / 3

<a id="canonical-3030312222002332-2211112333101022-0322122013122221-3102032202132012-0120010001132030-0200313331123031-0200023233311101-2302132032023210"></a>

<a id="canonical-2132310121130330-2032203031003132-0133201010122020-0223323100003213-0020303020322030-2322100120123323-1103111012001100-3302333133112130"></a>

## addr property — IPv4 / 000020333211 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-3032233203130211-3121303130102233-3230030121001000-0112212303331113-2223020021100210-2302233111123231-0001101000101020-2000321011213202"></a>

## Next pages — IPv4 / 000020333211 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2011021111102013-3023110120202033-0331312021311220-1123332323023030-1220322032013232-2210203012032320-1210230133222310-0120020131233230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000211000332223-0211021130310012-1203002021202332-3321013110210002-0221023231002103-1202233010100303-3211000012210101-2100232110202003"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 331323201000 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033)
- storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-2110212021130032-1202110122000102-2022332332133233-1021113013101033-0003310011220110-0300000112312130-1131012212223231-2010220331210113"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122022012212111-2013333032312213-0312122322011332-2001302312322302-3313113223120023-1230222030222012-3001013020021113-1321023001012103"></a>

## Direct properties — IPv6 / 331323201000 / 3

<a id="canonical-0031332122230123-2320102032122203-2003313022030331-1120323013330312-2023023211113331-1121100332012200-3000230100211031-1011302210221012"></a>

<a id="canonical-1313331311000330-3131130323102210-2010331002301201-0010313212233220-1202102200112102-2212322323012202-3103210313211202-2221111110103101"></a>

## addr property — IPv6 / 331323201000 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0222032101202301-1322010300232300-1210302232313321-1202030312213230-0122213223220330-2300321111023020-2201330101110122-3112202320221332"></a>

## Next pages — IPv6 / 331323201000 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address.dual_stack](data-sources--fleet--reference--group-004.md#canonical-2301223133112311-2000202212230202-1202322323023311-0210223000310120-3222233322231210-1311332211322211-2310111023212312-3102211221223033)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1203230330233330-0320331023002231-0012122030311100-3112313322120232-1203113223020201-3033112121033312-1031200202331201-2321200330101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003200210300123-0332023300112331-0333310310011323-0000103232111132-3012131331321223-2002013000233320-3013333231133123-3310112002011331"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.IPv4 — IPv4 / 202323303301 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- storage_static_routes.storage_routes.nexthop.nexthop_address.IPv4

<a id="canonical-0110331321002230-2010332312131131-2130132323111022-0211020101231312-1310131222211333-2131211122133231-1030010113211303-1111302000302131"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0231333202232003-2232033012231121-3033011323122111-1133033123112211-1110021021022231-1031122123122022-3013123311220222-2233322022112320"></a>

## Direct properties — IPv4 / 202323303301 / 3

<a id="canonical-3013233232200232-3131103301210110-3200030331021322-0210323303213331-0201202013202110-1212001331011012-3300010111111203-3131223130120313"></a>

<a id="canonical-2031021131231011-2333101210313323-2223003101312022-0320002102302220-0112132230331310-3122112333302010-0312310210222033-2303122332321311"></a>

## addr property — IPv4 / 202323303301 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-3102200332032113-0120321122102333-2120023231210221-2202012030000131-3200320121310303-2102301011032032-2203221020131223-0000111220222123"></a>

## Next pages — IPv4 / 202323303301 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-0300332322220010-0132002021012213-1323101311003021-1232221222232332-3223212333122100-3312321023311121-3323201230133223-2201330333332321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220202332232211-2332201301301023-2012123011123221-3322101331303013-1323003320303323-3001023231113112-2023013130322110-3032210111312320"></a>

## storage_static_routes.storage_routes.nexthop.nexthop_address.IPv6 — IPv6 / 012310102311 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.nexthop](data-sources--fleet--reference--group-004.md#canonical-2300310022230112-0312031201003113-0031001300232212-2032120302303013-0011333313112022-0300131011132000-0313323203013231-2020030103023000)
- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- storage_static_routes.storage_routes.nexthop.nexthop_address.IPv6

<a id="canonical-2220130211322320-0302002013002220-1322021313011120-0103321002032111-1021101300231010-1302203202301020-3222203303312210-2031333123230212"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0211130212222211-2200131201202001-2021123300003130-3232111002332311-2232223311133211-1230211323023321-1021100203301223-2112210202032310"></a>

## Direct properties — IPv6 / 012310102311 / 3

<a id="canonical-2300133101021302-1202013020231010-0113210321032102-2003313301230211-2320211110230220-3202012010223001-2220033111010303-3233131032032232"></a>

<a id="canonical-3000212120303022-2011222131201222-2000330203303003-2301321010232213-1311131312320003-2101231302113333-1320211112202022-3300030000321023"></a>

## addr property — IPv6 / 012310102311 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2110300020011231-3001330131301223-0310111102000123-1231021022313130-1123333101112110-3111011030103223-2031332120100113-3113220032321033"></a>

## Next pages — IPv6 / 012310102311 / 5

- [storage_static_routes.storage_routes.nexthop.nexthop_address](data-sources--fleet--reference--group-004.md#canonical-0133113223233112-2133220231303223-1313211122221213-3222103010003321-3212220031020101-3332102130123013-0120121330221102-1323133103022312)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232232012202220-2021312303232001-1331312121030033-0303233312222010-3023110013210120-0113130213212031-0312202100202202-1020231030301301"></a>

## storage_static_routes.storage_routes.subnets — subnets / 103122130131 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- storage_static_routes.storage_routes.subnets

<a id="canonical-0121220231232033-2201120120322120-2020231221003131-1112312103331211-1033023301303323-0003123032310111-0003201333213233-1113303033221230"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-2023010300322100-0020302020021030-0013323122012313-3303100211202131-3303200032011130-2311330222211303-2210113323232310-3113230122320230"></a>

## Direct properties — subnets / 103122130131 / 3

- [IPv4](data-sources--fleet--reference--group-004.md#canonical-1132211300313231-3232222201233203-0221010020302300-3323323100303012-2223321212111321-1130121030332303-1000311123133212-3333130120132012): complete subsection reference.

- [IPv6](data-sources--fleet--reference--group-004.md#canonical-1311003212203132-3330033022230231-0100323000203331-0023132223330130-2131030223221112-1223012010100231-2203010313311022-2010130333101313): complete subsection reference.

<a id="canonical-1313301231202012-2330021100121312-3333002130002000-2330013320133300-2231003210221103-1221200220301203-1120233221232010-2133202322323313"></a>

## Next pages — subnets / 103122130131 / 4

- [storage_static_routes.storage_routes.subnets.ipv4](data-sources--fleet--reference--group-004.md#canonical-1132211300313231-3232222201233203-0221010020302300-3323323100303012-2223321212111321-1130121030332303-1000311123133212-3333130120132012)
- [storage_static_routes.storage_routes.subnets.ipv6](data-sources--fleet--reference--group-004.md#canonical-1311003212203132-3330033022230231-0100323000203331-0023132223330130-2131030223221112-1223012010100231-2203010313311022-2010130333101313)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1132211300313231-3232222201233203-0221010020302300-3323323100303012-2223321212111321-1130121030332303-1000311123133212-3333130120132012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301122333000303-1221320230203233-1032121123333123-0301213332102210-1232213101133101-1223021302212100-0112120203023231-3133103310022203"></a>

## storage_static_routes.storage_routes.subnets.IPv4 — IPv4 / 021223010321 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303)
- storage_static_routes.storage_routes.subnets.IPv4

<a id="canonical-3111021123013120-1113020100122333-0331321321332011-1130023020312203-3021313031121130-0120232323121310-2323100221032020-3213122112301233"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0322322131303101-2213020203312311-0120031020200100-0312330221012220-3301312212211301-3221203231313221-2310311313320232-0013210301033303"></a>

## Direct properties — IPv4 / 021223010321 / 3

<a id="canonical-1010201232103233-3212022220013201-1002123030101122-0302001220102323-0102322310321333-1021012131133233-2330301101201121-1301020122301222"></a>

<a id="canonical-1312131311132320-2310300030220010-2330001320101201-0013103312032231-3331233022212103-0031302113003211-0233222312323013-1221333121131333"></a>

## plen property — IPv4 / 021223010321 / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3330122023022133-3021202132103230-0202020213021100-1330111313110213-1222233312330011-2021023332121002-0232120230103223-3222203120111201"></a>

<a id="canonical-2332120220001303-0100311000000203-2320312100022022-0313303102200302-3232122203233032-1200211231030010-3302001332221112-0002221231113211"></a>

## prefix property — IPv4 / 021223010321 / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-2133322330122100-1031103203231023-2023011113020021-2000221110000230-0300103002213211-3322200102311013-0122013133032331-1122130112321133"></a>

## Next pages — IPv4 / 021223010321 / 6

- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-1311003212203132-3330033022230231-0100323000203331-0023132223330130-2131030223221112-1223012010100231-2203010313311022-2010130333101313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132312323222133-3233303311302212-1220111221023310-0233300231002012-3221210303030210-3020331011330133-0130232013210022-0100220222203330"></a>

## storage_static_routes.storage_routes.subnets.IPv6 — IPv6 / 222103003013 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [storage_static_routes](data-sources--fleet--reference--group-004.md#canonical-2122200211031101-0310202002203113-1003213330032121-3020122111123030-1110020002232312-1321211020300222-1230230232002021-3132001113201300)
- [storage_static_routes.storage_routes](data-sources--fleet--reference--group-004.md#canonical-3330321223102110-2223112202210110-1100312321211232-2212230012130201-1020231212020133-1302332032122201-2030131220013312-0310330030133110)
- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303)
- storage_static_routes.storage_routes.subnets.IPv6

<a id="canonical-1330332231122313-3220022201313212-3332030102322033-3021113101203313-2003130320102221-0320322202003210-3203013113103010-3333033111020013"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3322302111302111-1120301010233233-3311300100030030-3310330310301213-1220031320112120-0211101310202201-3331330131020110-2221313221320222"></a>

## Direct properties — IPv6 / 222103003013 / 3

<a id="canonical-2232103210312231-1221103313021222-0123100131001300-1200212030320313-1001231022002031-0320022201220110-2131100210201001-1313203333010000"></a>

<a id="canonical-3100121200303020-0020120103033331-1300212133020211-1101330311102310-0011323231233022-2211202023122120-3321120010033000-0303010210123310"></a>

## plen property — IPv6 / 222103003013 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1112320312302203-0021021011033332-0032012322211232-2012212033000321-3202031300201103-0032002230032012-1022200231333022-2021310011201221"></a>

<a id="canonical-0321221210033133-2120131303203232-3133022130111030-2120203111231310-1102102002010330-0330032101013201-0123313223201131-2131122200100122"></a>

## prefix property — IPv6 / 222103003013 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2221103311210220-0113322000121120-3301100301132202-0301011223000303-3211212320211033-2012211310320022-1210130102223320-3330000101132301"></a>

## Next pages — IPv6 / 222103003013 / 6

- [storage_static_routes.storage_routes.subnets](data-sources--fleet--reference--group-004.md#canonical-3221013101022303-0031020003333031-1030012032322102-2001322333132322-0323001032121312-3112313131203323-0220202103032320-2230222202110303)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)

<a id="canonical-2132300120030302-2100113132002202-2230322111111011-2301031030022032-0303102003030130-2111000030223010-1011003201131323-2230112300300113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000313021232330-1323201102333031-1221303131021210-3301331031301211-0232332022122311-0023110111033333-0212033301333321-1002100222121102"></a>

## usb_policy — usb_policy / 021223222301 / 2

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- usb_policy

<a id="canonical-0031113210110330-1320123122203030-3102000002211020-0300002012223023-3100323131030300-1303202220232201-2302030110112130-2113203111101330"></a>

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

<a id="canonical-3210132013011232-1312130111102132-0133113321233211-2320032220310202-2312122301133120-3201012020121112-1302121231001022-2101131213220002"></a>

## Direct properties — usb_policy / 021223222301 / 3

<a id="canonical-0301221010202211-0132101021111021-1312222323232010-2212232001211122-3211332100232100-1323120210013011-3020222032301300-3132112021201001"></a>

<a id="canonical-2123210022331012-0320211112323311-0113030201312333-0121112302200203-2203322203231022-3013210321203222-1113320102311313-3002132002311121"></a>

## name property — usb_policy / 021223222301 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0003220322203031-2113233001023022-2013102230031303-0031013310022321-0122032333010211-1010010131111121-0333333203300100-2320010023301322"></a>

<a id="canonical-0132301211330201-0203001033133323-3201302233221200-3011110233101002-2321333010200221-0211021031201113-2210301013020320-3130201110102200"></a>

## namespace property — usb_policy / 021223222301 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2011103303103012-0131103303222023-1331321033102010-0030230212112233-0301102001302101-3113023320012123-0322223321300332-1232203200023021"></a>

<a id="canonical-2210203222220110-0233103001021212-2201332103312010-0020130102331223-2013110311300220-1320101000112223-0212302202133310-2023120210210110"></a>

## tenant property — usb_policy / 021223222301 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1233230000112113-2113012311333012-0120022320332233-0000300203032210-3213302003112001-1020213031111010-1001132312030111-2310012102203320"></a>

## Next pages — usb_policy / 021223222301 / 7

- [Property reference](data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
