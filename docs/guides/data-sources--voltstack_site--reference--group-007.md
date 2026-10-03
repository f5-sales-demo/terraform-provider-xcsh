---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-2022331022023221-2022020221301211-3301332210010121-2033202301100030-3301212031333100-2011202220321213-1303223323100113-2303120032032203"></a>

## Next pages — blindfold_secret_info / 130230311222 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0031221320012122-2002233213121101-3003200313031331-2212223302113100-2313302033333122-2032231231331013-0221000231001003-2110023331211220)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0223330300130022-0012123333120000-1011000003112310-3110211020323233-3111311202233031-3310030221212333-2212202123032122-2010322330313223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210121201203222-3003213313033131-0102123013200031-0211033323223022-1221301021313032-1312110113333231-3311211101100001-1103313312323213"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info — clear_secret_info / 033023203021 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0031221320012122-2002233213121101-3003200313031331-2212223302113100-2313302033333122-2032231231331013-0221000231001003-2110023331211220)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-3333131322330023-3231210120302333-3000321010211301-0201232120312203-1121021000122223-2230212220010133-1030201312003311-2213222322010001"></a>

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

<a id="canonical-0001110000103202-0002032323032233-0021323203033023-1221112312322033-3123011301301101-3113212320133123-0321020303321003-1031310032023111"></a>

## Direct properties — clear_secret_info / 033023203021 / 3

<a id="canonical-0320300110213310-0202010300303010-0013003110121201-3312122320013023-2132220323122030-2103123021332330-2003223133023000-1223131021130130"></a>

<a id="canonical-2213101302023333-2233303123210313-1332121112132232-3312300101022113-0302010233011210-2111123030222232-0113210231302022-3303111130000123"></a>

## provider_ref property — clear_secret_info / 033023203021 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0122312232323000-0100101100220310-2310001101020200-1012310220001233-2121222033231311-1011131332311003-1201222210300120-0203131110112031"></a>

<a id="canonical-2221103123300030-1130221230302222-0200222130203010-0331121222230313-1133130303332230-2321131011311212-3311003313023203-1021221310110322"></a>

## URL property — clear_secret_info / 033023203021 / 5

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

<a id="canonical-0102111323203213-1213013321223302-0212210311212023-3320323333312203-1031032101022121-2200320213131123-2110103132110000-2301200133322013"></a>

## Next pages — clear_secret_info / 033023203021 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](data-sources--voltstack_site--reference--group-006.md#canonical-0031221320012122-2002233213121101-3003200313031331-2212223302113100-2313302033333122-2032231231331013-0221000231001003-2110023331211220)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2123020023031200-2333220022121210-1120020130313012-0313132120312202-2233312231113231-2122223330311101-1312011233231233-1310013123333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231112213030013-3322001102301011-1133010212101312-2321312011103230-2103113302030210-0112032221311100-0132121120301112-1322010212310012"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap — no_chap / 110211303103 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-3100032222013331-3201203003101213-3012110020023333-3322311132030111-0220300012133202-2002002111122112-0000300003101322-0030211012332220"></a>

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

<a id="canonical-3222213022202210-0123010322102211-0102330013330021-3233102332021010-2211303223332030-3202001221013011-1123303013132132-0122023032102022"></a>

## Direct properties — no_chap / 110211303103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122302010002023-0011311300202333-2131231231001110-3313012212311103-0300111220021222-0332010131033101-3232223312133033-0203102112102200"></a>

## Next pages — no_chap / 110211303103 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2002230133022000-3011203023322232-2010110321231200-0013321303301320-2200010220303333-0123330233021231-1213023311113322-3001232310123311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113332313033131-3022231310131102-3133033030001211-3020002232333010-1032010013101122-1002323133020313-3211232301203030-2230203301222223"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password — password / 112003220211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-0322303320021002-3313220112203232-2001302300231303-2111131100222232-1213020011222230-2231332122213022-3300213231031212-1132000103002123"></a>

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

<a id="canonical-2133121112111003-3322001021023102-0103212330330022-1033011131213311-1131003223032131-1130230132003123-1030210300212233-2010222023132201"></a>

## Direct properties — password / 112003220211 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3331113232230020-0311320112113030-2001311010123022-1303230301311233-2130220210233330-1202313002302132-2123301122013132-3202302001030110): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3003132322113202-1102230022103032-2022121330232122-2220313101230313-3130301213221010-3130302133112130-2031120100113012-2000230203133233): complete subsection reference.

<a id="canonical-3002221011020102-2032200010233310-2112212300302021-1300333022013123-3301211130123103-0330103221031021-1220231010232020-3303012213122011"></a>

## Next pages — password / 112003220211 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3331113232230020-0311320112113030-2001311010123022-1303230301311233-2130220210233330-1202313002302132-2123301122013132-3202302001030110)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3003132322113202-1102230022103032-2022121330232122-2220313101230313-3130301213221010-3130302133112130-2031120100113012-2000230203133233)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3331113232230020-0311320112113030-2001311010123022-1303230301311233-2130220210233330-1202313002302132-2123301122013132-3202302001030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130220202233203-0320200212231220-1100120203110212-2210303313011001-2232112331120302-1121110211211300-3022112130200313-0133122030013122"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info — blindfold_secret_info / 133233023130 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-007.md#canonical-2002230133022000-3011203023322232-2010110321231200-0013321303301320-2200010220303333-0123330233021231-1213023311113322-3001232310123311)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-3122001110120232-1303303113133233-1222320200112301-3333230220323231-0220300000223000-0301322121203030-1211202001333322-1102100223103200"></a>

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

<a id="canonical-2231213032000333-0031031312300321-0321002222201212-0001131223011211-0012200202302213-0311331332213133-2303112011100320-2120333313320202"></a>

## Direct properties — blindfold_secret_info / 133233023130 / 3

<a id="canonical-2310030121100303-3133230230302321-1013012032121332-2312301303211131-3032131122133220-1110132223200023-3112021332110230-0132330020021321"></a>

<a id="canonical-3301230231102323-3323122321111010-2013003012331323-1112130203310330-3001220101102231-2323222311233321-3222033300302200-3322023220102121"></a>

## decryption_provider property — blindfold_secret_info / 133233023130 / 4

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

<a id="canonical-0122222113030221-3323211021033131-0111222011203232-3322020200012233-1003233100301203-2133332103121312-2210300323201002-3301221013020132"></a>

<a id="canonical-1022123233231231-3213133202120311-1131203002021332-0030203311101022-3002303000200001-0301212332100213-3133303311223230-2001201003030333"></a>

## location property — blindfold_secret_info / 133233023130 / 5

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

<a id="canonical-2120033123032200-1210120230103102-2203313010231123-3032211232302230-0121311312011031-3120103012333020-2331332021312330-1223121120101001"></a>

<a id="canonical-3101203201123101-2210212200002322-2032021122023131-3201203103311021-3011311200230121-1011232110013213-2103123210200120-1001020333101212"></a>

## store_provider property — blindfold_secret_info / 133233023130 / 6

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

<a id="canonical-2300223120310001-3333032032201320-3021022220330232-2010302111013021-3303303130012211-1221101022301022-0103000300002000-0030030113322300"></a>

## Next pages — blindfold_secret_info / 133233023130 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-007.md#canonical-2002230133022000-3011203023322232-2010110321231200-0013321303301320-2200010220303333-0123330233021231-1213023311113322-3001232310123311)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3003132322113202-1102230022103032-2022121330232122-2220313101230313-3130301213221010-3130302133112130-2031120100113012-2000230203133233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011211231123132-1133000221121100-1002130312131120-0312131202020332-0111322100211103-3023110232100030-1110032223230022-1303201231213110"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info — clear_secret_info / 213233111230 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-007.md#canonical-2002230133022000-3011203023322232-2010110321231200-0013321303301320-2200010220303333-0123330233021231-1213023311113322-3001232310123311)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-3120310303212131-3233322000213110-3010202033223113-1012321111222122-0302111213023021-0003320222103012-1330133200301132-0212130002202212"></a>

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

<a id="canonical-0110031302030031-0212112010132330-0201101010111333-1132312023331002-3031302230021303-0011320332001123-1131110310013231-2112000122330233"></a>

## Direct properties — clear_secret_info / 213233111230 / 3

<a id="canonical-0130323332132221-1311302330101020-3333301302030301-2322003100030223-2202030220222333-0321120121100200-2013021323331122-1102003231231210"></a>

<a id="canonical-0311310012222001-0333001030002332-0132313220110210-1023101311221201-2300121103111032-0233102001002110-0321312121232022-1201330112233001"></a>

## provider_ref property — clear_secret_info / 213233111230 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0301302302001330-2010112001001323-2011030330002212-2001132031202212-0322211020020031-2001212201300223-2231113221121023-2033300102102230"></a>

<a id="canonical-2203003210230011-2002220323213230-3003200002000121-1120101302012121-2323211220011303-3022301133300323-3223003303211323-2312001203002010"></a>

## URL property — clear_secret_info / 213233111230 / 5

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

<a id="canonical-3212131120232013-2210111220200101-0320223113313202-3013003120330121-2312320000123310-0031033201023123-3302031203030310-1321021303120100"></a>

## Next pages — clear_secret_info / 213233111230 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](data-sources--voltstack_site--reference--group-007.md#canonical-2002230133022000-3011203023322232-2010110321231200-0013321303301320-2200010220303333-0123330233021231-1213023311113322-3001232310123311)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1321303001321312-0013123013303321-1200013023232333-2120313211303332-0233020303202230-0101001210010021-1033011211222210-2222201031023232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223212300223310-3103330011332230-1031331201111233-1302101122102100-2030313200120300-2032311000113220-3021322231322212-0201020201111311"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage — storage / 233033302030 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-2120332002130322-1022130023333120-2132321030323113-0322121001203113-1223331213230121-0121021331120320-3223121120213110-1210310020320031"></a>

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

<a id="canonical-3121223101312133-3123222230202122-1303131101221301-3133132012320122-1003102102010203-3011222020031110-3003231023321120-0110303313101333"></a>

## Direct properties — storage / 233033302030 / 3

<a id="canonical-3320000002211023-1203322021312030-3330231121331121-3112120001120002-3133103101030211-2022130323013001-1200013201103112-3211012332132001"></a>

<a id="canonical-1032311201111033-2133210320303103-3033113012032031-1003003303312010-0220200030030122-3132321032211312-3130020101331000-3001331220120013"></a>

## labels property — storage / 233033302030 / 4

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

- [volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-2230321330333321-3200202322120330-3302033130001303-3232223012022201-2331222120212201-1203113002222231-2233320203133031-0211022032210131): complete subsection reference.

<a id="canonical-2023322203010330-0033311322001223-2111112020001121-2231103032012132-3000201221130102-3021003032332010-0221331200031103-3311031110303203"></a>

<a id="canonical-0021210201010313-1123120101310332-0321013123333210-1030101311200100-0201011203013112-3032310300330010-0321232300233223-3020311200232332"></a>

## zone property — storage / 233033302030 / 5

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

<a id="canonical-2100033130201020-3102221322111221-3221313021310111-3331331011022012-2322001323032111-0301110021200223-2130231111222230-0130100130333123"></a>

## Next pages — storage / 233033302030 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-2230321330333321-3200202322120330-3302033130001303-3232223012022201-2331222120212201-1203113002222231-2233320203133031-0211022032210131)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2230321330333321-3200202322120330-3302033130001303-3232223012022201-2331222120212201-1203113002222231-2233320203133031-0211022032210131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030003120330231-2322302333001213-2020120200202220-2011130230032323-3013122212230210-0331301111321303-3122103330233031-1121032230111312"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults — volume_defaults / 032010101231 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--voltstack_site--reference--group-007.md#canonical-1321303001321312-0013123013303321-1200013023232333-2120313211303332-0233020303202230-0101001210010021-1033011211222210-2222201031023232)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-3100300110132302-3031120100303133-0032220131311302-2330201212023233-2322111031120133-1232211011333100-3131113031101010-0331301002200021"></a>

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

<a id="canonical-0223210300200023-1111022233311203-1032232013213000-1231022100313003-0132320232033312-1311123332002120-0310031113320030-1222223020013100"></a>

## Direct properties — volume_defaults / 032010101231 / 3

<a id="canonical-2312000021112133-0130101011132301-2010213013133220-2023310202312103-3101222132000001-1003302013013110-3303113000112122-1033320032201112"></a>

<a id="canonical-0301331310012121-0133002030021232-1200220101023233-2321230301130033-0201110100203202-2110312012330132-3223021011111323-2233223000101102"></a>

## adaptive_qos_policy property — volume_defaults / 032010101231 / 4

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

<a id="canonical-2000011331031210-1332212222001020-0232302103202312-1130303030311310-3320320230222333-0002322212332031-0323221110033212-2013012103133001"></a>

<a id="canonical-3221302231002331-1212011300203310-3312120120311321-3022013133031121-3131320300033333-0233130232100011-1123021001232002-3120010001313100"></a>

## encryption property — volume_defaults / 032010101231 / 5

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

<a id="canonical-0022201020033003-0023300230200310-3123202032212203-3033011111220020-0300211202312101-1003122032200133-2330212321321310-0232020213230110"></a>

<a id="canonical-1320330022111302-0313211301203030-3122111213100120-2222113011021232-1202032121102233-1301032001131332-2310032100023121-0022212212033031"></a>

## export_policy property — volume_defaults / 032010101231 / 6

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

- [no_qos](data-sources--voltstack_site--reference--group-007.md#canonical-0212003001333232-3022323010003012-1120333312311213-3100331030213301-3133201002031222-2033112112220000-3310220331111110-2310101233032120): complete subsection reference.

<a id="canonical-1010130132233322-0331332100232321-0220213001020013-3130312000023112-1231201331013102-2113232001130313-1031302220223012-2311122330032322"></a>

<a id="canonical-3011011321000011-2002123032322010-1330122321200311-0301001211003112-1313211032231232-0221022222121101-2212203011311310-2333210002323330"></a>

## qos_policy property — volume_defaults / 032010101231 / 7

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

<a id="canonical-2103033012011302-0300302230030320-3210230030122311-1020110313330222-2103333302131311-1001330000010010-2300203133320211-1220013032102020"></a>

<a id="canonical-0213123113033211-3231213132221320-1021211023232320-1312022133233131-3123321102022103-1030021211331320-0313331211130000-1313210230031200"></a>

## security_style property — volume_defaults / 032010101231 / 8

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

<a id="canonical-0320122033233002-1122012033220203-3221303200001301-3110012330101200-1231100032222103-0213011013021122-2022131203010033-2332313230331131"></a>

<a id="canonical-2133303200230103-3131112023232122-1123300113321323-2221023203321112-0310232023220013-2200222223321130-2333331332300102-1312211003130101"></a>

## snapshot_dir property — volume_defaults / 032010101231 / 9

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

<a id="canonical-0202031302221133-3221310002103231-2313232330010300-3103312213122310-0122101231012120-2222302333001313-2002022031332310-1000233322313122"></a>

<a id="canonical-0112000333312133-1203232231021302-0100332211212310-1231023233100101-2231133033101303-3002100202213021-2130302002212001-2200120321320112"></a>

## snapshot_policy property — volume_defaults / 032010101231 / 10

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

<a id="canonical-1222233133003211-1330221101021300-0303130100033000-2103031300222322-1010223211111310-1310331022213300-1132303233111021-1012011113302110"></a>

<a id="canonical-2123110200213033-0133323301102202-0132221122013322-2023021012121321-2022003313233333-2102223321131233-3020021220010332-0232102320021233"></a>

## snapshot_reserve property — volume_defaults / 032010101231 / 11

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

<a id="canonical-1112222101230112-2111301030021132-0020202133122021-2313012223033310-1113333032003120-2300110020203033-1332021200003103-1011323300312210"></a>

<a id="canonical-2323023113202100-3103123031033111-3113203230132010-3220000130031211-1321202023303030-1102013110120112-3131113333111023-1112132103000210"></a>

## space_reserve property — volume_defaults / 032010101231 / 12

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

<a id="canonical-2100310012010210-0220200103111111-2303130213300303-1220201101003220-2333331030213231-1122110303301213-0021303230122200-0203222112013220"></a>

<a id="canonical-3231121223212210-0033032100112111-0200031131323103-3030323010331113-1102331333031200-3202330211200200-2131022022023323-3322231013232022"></a>

## split_on_clone property — volume_defaults / 032010101231 / 13

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

<a id="canonical-2333132002200020-3212000230221103-2203103202322303-0222123213003210-3233131121230330-0030332002322020-1320302322100021-1021030000112103"></a>

<a id="canonical-2230110112123000-3221213111210100-2101023323330310-2323022011303121-3221202113331012-1130023033000303-0332002232111001-1122033301221122"></a>

## tiering_policy property — volume_defaults / 032010101231 / 14

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

<a id="canonical-0130032221222232-1202233120030123-3230130230221332-0101130210232311-1110213023303122-1132101111021222-0030332332110121-3210112001101011"></a>

<a id="canonical-3112002201310002-3231202300120322-3200222111332332-0213013200003100-3002001131023320-2023103012032011-2131032302111121-2022303303001210"></a>

## unix_permissions property — volume_defaults / 032010101231 / 15

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

<a id="canonical-1313112031110002-1103230101311313-0310003110000122-2211110130333120-0210103000102200-3230321211201212-1312012310223210-3233223111110021"></a>

## Next pages — volume_defaults / 032010101231 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](data-sources--voltstack_site--reference--group-007.md#canonical-0212003001333232-3022323010003012-1120333312311213-3100331030213301-3133201002031222-2033112112220000-3310220331111110-2310101233032120)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--voltstack_site--reference--group-007.md#canonical-1321303001321312-0013123013303321-1200013023232333-2120313211303332-0233020303202230-0101001210010021-1033011211222210-2222201031023232)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0212003001333232-3022323010003012-1120333312311213-3100331030213301-3133201002031222-2033112112220000-3310220331111110-2310101233032120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021213230113111-2032112232233322-2200220330120131-3300233003123233-2311212132220311-3131200120232313-3332002110110031-3102310032121322"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos — no_qos / 010111301013 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](data-sources--voltstack_site--reference--group-007.md#canonical-1321303001321312-0013123013303321-1200013023232333-2120313211303332-0233020303202230-0101001210010021-1033011211222210-2222201031023232)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-2230321330333321-3200202322120330-3302033130001303-3232223012022201-2331222120212201-1203113002222231-2233320203133031-0211022032210131)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos

<a id="canonical-2113102110030111-1013202023031101-1013331231331010-0333113030212100-3312032303003011-0020303103002110-3211130003123212-0030221100100310"></a>

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

<a id="canonical-0333022221332210-3032313202211013-2322020111023210-3312031133012113-3102123320310221-1021110330032123-1322200212111000-2332221221302113"></a>

## Direct properties — no_qos / 010111301013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222022132210303-3313021110030023-1113010031103001-2000121200012120-1322330000332101-3102231232331033-1232032233213202-2033212133111220"></a>

## Next pages — no_qos / 010111301013 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-2230321330333321-3200202322120330-3302033130001303-3232223012022201-2331222120212201-1203113002222231-2233320203133031-0211022032210131)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012023203231121-2133212122301203-3231301301313232-0313120123302031-1221102032320223-3023012021001021-1333312200010213-0303020300001010"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap — use_chap / 013022201232 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="canonical-1031112030322312-0001303313032330-1310322101211123-1201200230322013-0101312333022220-2013220010033020-2233112012130233-2002010131020123"></a>

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

<a id="canonical-0023313001000123-2022220211232112-3013013021002312-3130223020132222-2320131020211020-3002130002222202-1320320301313323-3023330200220202"></a>

## Direct properties — use_chap / 013022201232 / 3

- [chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1000123101302213-2001121103222031-3202333322013201-3123001031022013-0213121300123220-2102002021132211-1000011020231320-0022202012132113): complete subsection reference.

- [chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1330023203122103-0013323102311222-0231011130103203-2021000011033130-1102311231300020-3030232113001112-0203331110101103-3213301302332020): complete subsection reference.

<a id="canonical-0312131002200323-2303013302332022-2003222301023001-3333303221003032-1230301323133020-1001123032133202-2222222003123210-1302221130021002"></a>

<a id="canonical-1221020233010030-0333100222323011-2110020231310212-3231101123320023-0232103311213120-1031011301222330-3112310032023333-2133100030231331"></a>

## chap_target_username property — use_chap / 013022201232 / 4

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

<a id="canonical-1312011020330103-2130333221310030-3122131002103003-3303212013300111-2211220313210201-0102012113310120-0330231012030003-0031012203101333"></a>

<a id="canonical-3302323110221310-3311213122233023-3001212010133312-2212000021103330-3113232230133133-2223200333230023-1311203213101003-0323300220032203"></a>

## chap_username property — use_chap / 013022201232 / 5

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

<a id="canonical-1031301220303302-1310012212102300-3002312200312230-0100322022111003-2102320002231133-3113201003130300-3030022303120331-2033030031000223"></a>

## Next pages — use_chap / 013022201232 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1000123101302213-2001121103222031-3202333322013201-3123001031022013-0213121300123220-2102002021132211-1000011020231320-0022202012132113)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1330023203122103-0013323102311222-0231011130103203-2021000011033130-1102311231300020-3030232113001112-0203331110101103-3213301302332020)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1000123101302213-2001121103222031-3202333322013201-3123001031022013-0213121300123220-2102002021132211-1000011020231320-0022202012132113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023032003322322-2033020013211301-1132113323220102-3310021230111332-0312123032201102-3121220332322223-1000233021033022-3032131303013333"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret — chap_initiator_secret / 333001103301 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="canonical-0001320010210220-3213312302330000-2321020110021223-0133002010232111-0110323212311303-2133130322233011-0013020003200302-2011001301131221"></a>

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

<a id="canonical-0211111310110011-2320303313002320-1232012021010132-3231131333212303-3001302310113202-1110331002030030-1222110332020210-0233022213201032"></a>

## Direct properties — chap_initiator_secret / 333001103301 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-0132221012200232-0003013101320313-2121003102313023-2022023312310212-2111213101012020-2320220030020120-3103131311321031-0120031020301102): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-1210032023201101-2102211311220311-1313133011023313-0332321110331010-1311320200233221-2033031003131030-2223032330030301-2321321121013003): complete subsection reference.

<a id="canonical-2123110230011132-1001012113213100-2303332113300103-1313002330210101-2103122030202033-1202022332322000-1210020323030130-0123301333312122"></a>

## Next pages — chap_initiator_secret / 333001103301 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-0132221012200232-0003013101320313-2121003102313023-2022023312310212-2111213101012020-2320220030020120-3103131311321031-0120031020301102)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-1210032023201101-2102211311220311-1313133011023313-0332321110331010-1311320200233221-2033031003131030-2223032330030301-2321321121013003)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0132221012200232-0003013101320313-2121003102313023-2022023312310212-2111213101012020-2320220030020120-3103131311321031-0120031020301102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210201303010222-3233020102221203-0201100220100021-3033213023000201-1132321011222120-1330310101101012-0321011202201233-1333233320113221"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info — blindfold_secret_info / 320121311203 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1000123101302213-2001121103222031-3202333322013201-3123001031022013-0213121300123220-2102002021132211-1000011020231320-0022202012132113)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info

<a id="canonical-1011121002001311-3203021232103121-0231311233233321-0102303213111101-0323031130023213-3231001022222332-2111131131100020-1101120321000302"></a>

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

<a id="canonical-2013103110103032-0201200311023323-2032321201302230-0023110201331333-3020202001111021-2231112113112213-0210031002230111-3112120312032232"></a>

## Direct properties — blindfold_secret_info / 320121311203 / 3

<a id="canonical-3223301020210230-3311323331020113-1132210112133031-3232030300033003-3003223110022211-0002202312232200-0300001212301332-1233032033311222"></a>

<a id="canonical-2033221230020233-0133310222023133-0131010000002321-2103202231113101-1133111303020323-1112001010333333-2111322121022322-1321111132200031"></a>

## decryption_provider property — blindfold_secret_info / 320121311203 / 4

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

<a id="canonical-1101201121233102-2203311300221210-3213023003022212-1203032303323210-0303303232102130-3231010013301212-2301202233120113-3312012101201030"></a>

<a id="canonical-2110012222302333-3121233113103231-2221002113102112-0300302022112030-2002330133303212-1022313020010330-2222312213013020-3013331311233002"></a>

## location property — blindfold_secret_info / 320121311203 / 5

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

<a id="canonical-1320302211320120-2320032222102322-2003322212122022-0302111102122322-2130121113001201-0222011113230103-3323010022202032-3221031313232032"></a>

<a id="canonical-0031221003221000-1112011122233210-2122230303303201-3013301332223303-3003332212312023-2230322031032303-0303300311123323-0211120112323300"></a>

## store_provider property — blindfold_secret_info / 320121311203 / 6

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

<a id="canonical-0001213030223101-0201231212310013-1133332312221122-3230130330312303-0031133301110333-2121012030323003-3130021203300230-0100121222311032"></a>

## Next pages — blindfold_secret_info / 320121311203 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1000123101302213-2001121103222031-3202333322013201-3123001031022013-0213121300123220-2102002021132211-1000011020231320-0022202012132113)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1210032023201101-2102211311220311-1313133011023313-0332321110331010-1311320200233221-2033031003131030-2223032330030301-2321321121013003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210223111321010-3021012022013002-0330132112031021-0002323023030032-0033303013311220-2221332322030110-1330232211323020-1223213210023100"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info — clear_secret_info / 133231333030 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1000123101302213-2001121103222031-3202333322013201-3123001031022013-0213121300123220-2102002021132211-1000011020231320-0022202012132113)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info

<a id="canonical-1323010021011233-2320130023221022-3010121103303322-1233230000311313-3232033103322332-2321023133130000-2202003230211011-2311313300031311"></a>

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

<a id="canonical-3310311130200300-0201110230133330-2311312221020321-0211300020303331-0130103013203222-2010330302030213-3212200012110011-3323231231213133"></a>

## Direct properties — clear_secret_info / 133231333030 / 3

<a id="canonical-2013123321312031-3222110313031233-1131220312331110-0002330013013011-0330310223302302-3322103233302102-3311110233221121-0010023011301130"></a>

<a id="canonical-3213233100303311-1210020201213102-2002332130310230-0221122000113121-1113000130113110-3331210202202131-0003033131232221-2310101112330111"></a>

## provider_ref property — clear_secret_info / 133231333030 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1331221230202200-1022133321002030-2233302101222221-3220110310311132-2012102130110200-3121322322033110-0323011323221231-2110132313003113"></a>

<a id="canonical-1222131313031011-0003212331033323-0102312303113302-2003003013013122-0132221223221122-3022003003120022-1020011020323233-0320133030332332"></a>

## URL property — clear_secret_info / 133231333030 / 5

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

<a id="canonical-0013201311202030-3113111322320131-0111122000020221-1121223301201313-1310002011331023-0311013320221030-2333003213033223-2330113201132221"></a>

## Next pages — clear_secret_info / 133231333030 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1000123101302213-2001121103222031-3202333322013201-3123001031022013-0213121300123220-2102002021132211-1000011020231320-0022202012132113)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1330023203122103-0013323102311222-0231011130103203-2021000011033130-1102311231300020-3030232113001112-0203331110101103-3213301302332020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312100000202031-0303211013133333-3332330123301101-2013111011233122-3133302121121223-2003020023122201-3102233233331121-2012002330110203"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret — chap_target_initiator_secret / 010022130320 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

<a id="canonical-2210331022131201-1011222120310033-0103200201033021-2322011131111120-1322302332331322-0313302021032213-0220012110331110-3330330013112330"></a>

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

<a id="canonical-0230123110022200-0013331013130020-3022000033222232-3030303310122222-1000230011200313-0010011130112023-2102210310313211-0120132121110101"></a>

## Direct properties — chap_target_initiator_secret / 010022130320 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3102132202222120-3330033002322011-3203021013223232-1003002011323330-3211203210103121-2102133023313320-0032011320301223-2003133312230030): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3121131002231123-2232003130213003-3321300120013323-3320111121131231-1313020202320103-3011133130301221-0201130103133102-3313013331010033): complete subsection reference.

<a id="canonical-2011301302330302-2311320102000203-0211003003102102-2010031022220032-0032122120130112-0203202222231210-0110132011210211-2003222302002022"></a>

## Next pages — chap_target_initiator_secret / 010022130320 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3102132202222120-3330033002322011-3203021013223232-1003002011323330-3211203210103121-2102133023313320-0032011320301223-2003133312230030)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3121131002231123-2232003130213003-3321300120013323-3320111121131231-1313020202320103-3011133130301221-0201130103133102-3313013331010033)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3102132202222120-3330033002322011-3203021013223232-1003002011323330-3211203210103121-2102133023313320-0032011320301223-2003133312230030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113111332313303-0222202023333132-1232230033132211-3011332022131302-0213120202101333-3222300032323011-3233211120101312-2032023133111011"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info — blindfold_secret_info / 230000222012 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1330023203122103-0013323102311222-0231011130103203-2021000011033130-1102311231300020-3030232113001112-0203331110101103-3213301302332020)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="canonical-1320011012122322-0211130001202312-0220122103132000-1301233310022330-1003030212102212-0013113032021100-0333233131220233-1011201222113012"></a>

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

<a id="canonical-2023123011331102-3130002201133023-1311002300213201-3031122110110123-3120001030023002-1032123121221322-0133010322013321-2000023323120330"></a>

## Direct properties — blindfold_secret_info / 230000222012 / 3

<a id="canonical-0230210210031301-2003132110121213-3012132321301010-3020232322310303-0221330222101313-0010131222020232-0002220222122030-2033301223300022"></a>

<a id="canonical-2311103211310133-0323030230212310-0320210022031012-2022331213120220-1222231013012101-2300221322230002-0012001013120102-2111323133330220"></a>

## decryption_provider property — blindfold_secret_info / 230000222012 / 4

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

<a id="canonical-0212301200323010-0301202230322031-1000101310223002-1133002220113221-1001130013231221-0133033033023220-0011100232122131-0112013220122031"></a>

<a id="canonical-0223223131130133-1123003022122202-2100231021100122-2233312210133130-3322223102303301-0202332122132222-1210103103302230-3321221231123332"></a>

## location property — blindfold_secret_info / 230000222012 / 5

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

<a id="canonical-2302221101211102-2310223213321313-1012031230003201-2213222221331333-0312301112232232-1032133331200332-2131301330133120-0223312023013031"></a>

<a id="canonical-0313222222031013-1312131123102321-1323202311201021-3333222011212332-1110322032112331-2021200201310221-3003330231202120-3303330310320313"></a>

## store_provider property — blindfold_secret_info / 230000222012 / 6

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

<a id="canonical-0210322301233203-3032221220210310-0203322002030113-1211120332331303-0232111322021032-2100333310130103-0301210031231312-2121031111213131"></a>

## Next pages — blindfold_secret_info / 230000222012 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1330023203122103-0013323102311222-0231011130103203-2021000011033130-1102311231300020-3030232113001112-0203331110101103-3213301302332020)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3121131002231123-2232003130213003-3321300120013323-3320111121131231-1313020202320103-3011133130301221-0201130103133102-3313013331010033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302120332001102-2130301001102321-3323320201332200-1230311121133220-2101133331122330-3301002203332133-1220321313312321-0102033110111121"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info — clear_secret_info / 203011231020 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](data-sources--voltstack_site--reference--group-007.md#canonical-3223320000310012-2211323020003202-2013321310202103-1200123111010301-2030113011213100-0100231010200231-1101330100223021-1133223303333221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1330023203122103-0013323102311222-0231011130103203-2021000011033130-1102311231300020-3030232113001112-0203331110101103-3213301302332020)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info

<a id="canonical-2301321123211013-1013332322000130-2203232120201012-2210110310122112-2323023203332113-2200231001123023-2020220133021010-2023331210000211"></a>

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

<a id="canonical-1020121333002103-3322101221032122-2030333121201131-3001310302311321-2030133113313311-1233103330233100-3311313020311123-1230030213210120"></a>

## Direct properties — clear_secret_info / 203011231020 / 3

<a id="canonical-0011021201320301-3231230203131332-0100203331303022-0101023200231120-0123221023212010-3230232312311313-2101222031210122-1120332223312210"></a>

<a id="canonical-1011231203211203-3202121212022331-3122320323111210-1203331103010030-3232322212120320-1110313212303300-1133021013133202-2333301201212212"></a>

## provider_ref property — clear_secret_info / 203011231020 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1021221321133210-0322000030333212-0202113132311203-0022322032201332-0231010113123131-1000210101011130-0113001100131322-2220230331003230"></a>

<a id="canonical-2033300031230213-2021030000103200-1002110012010023-2200101023232223-3012131100101003-1233030110321301-1122023010003321-3020033013110032"></a>

## URL property — clear_secret_info / 203011231020 / 5

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

<a id="canonical-1313313131033300-3103323331231033-2221132132033112-2233121310103021-2001012122203330-2221033210121031-2131210022033101-1302030102303312"></a>

## Next pages — clear_secret_info / 203011231020 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](data-sources--voltstack_site--reference--group-007.md#canonical-1330023203122103-0013323102311222-0231011130103203-2021000011033130-1102311231300020-3030232113001112-0203331110101103-3213301302332020)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0010213121112030-2023010302202201-0032322030031231-2203123331233303-3331013031313301-2111303021000302-2003330020101200-0100231311031011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313322122212002-1323211111311233-0211132110121002-0102113232100021-0033021211332213-3021023203303010-3322301222022111-2232102012332231"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults — volume_defaults / 121223211031 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="canonical-3031101323111132-3102103220311111-1222330312200222-2213013310330323-0111100123030022-2223203112310021-3301032311220030-0130012203313312"></a>

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

<a id="canonical-2021320331210131-3033311221030310-3211031223010021-0032222203300111-3330311032201020-2012331021301230-0013102311231123-0100312120313132"></a>

## Direct properties — volume_defaults / 121223211031 / 3

<a id="canonical-2313131110312330-1322121110110202-3103222022210211-0212022300322211-0131213313333113-1010032131110213-1102033212203000-3223131232130321"></a>

<a id="canonical-0031021212332112-3101330022031303-2003202210100223-1023011131131322-3220132130203203-1110101021133022-1301011311201221-2212013021112310"></a>

## adaptive_qos_policy property — volume_defaults / 121223211031 / 4

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

<a id="canonical-1220020112311003-0313003121220100-1020302032223102-0330003022003231-1020000310203123-0030300130113231-1133213220002321-3223100101132313"></a>

<a id="canonical-0020212123332302-2022330111111001-1013310231133302-3001101033220003-0332330303221200-3233132030331112-1102212002111233-1032020332321211"></a>

## encryption property — volume_defaults / 121223211031 / 5

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

<a id="canonical-3111103212213322-0122200002130132-2023003233302002-2302300010033331-2333202220213130-2000003011331231-2031312120201001-1301200222132211"></a>

<a id="canonical-3211221221322203-1212100112113231-0310010223223021-3021232032222103-2120331133321211-0301122102023021-1113102122120322-0101020010001203"></a>

## export_policy property — volume_defaults / 121223211031 / 6

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

- [no_qos](data-sources--voltstack_site--reference--group-007.md#canonical-0001013031313310-0323331301322111-3003023333000132-2222303302203300-3033313031303210-2113012231320002-0012022121330310-1230312220311000): complete subsection reference.

<a id="canonical-1003121300100012-1003120321202312-1102323032120212-2103212331013113-1013130232132002-1301320300231100-3222213130111323-3000120123330112"></a>

<a id="canonical-1202031110232332-2231301023131301-2232012013321023-2000133013002222-1333110202022321-1332223110121333-1303320313310201-2222231122003210"></a>

## qos_policy property — volume_defaults / 121223211031 / 7

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

<a id="canonical-0002200030210331-1220221131130323-0330323133211323-3031310032022121-3211313301003103-1303111122001300-2221202122201220-2101030301311003"></a>

<a id="canonical-0121332112032023-2033120211223232-0011132320003210-0313030212302303-3220002130132311-2233021023102003-2011323202032021-2331130013332322"></a>

## security_style property — volume_defaults / 121223211031 / 8

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

<a id="canonical-3030130131012102-1032301023232203-2211131001132100-0110310211000011-2310001130232102-1111111123122221-0222122232003000-1203113131200333"></a>

<a id="canonical-2031321330103300-2230333011121310-3213221113200331-1000202030011002-3033011210301232-0302032000000200-2310133002122300-3201023110133301"></a>

## snapshot_dir property — volume_defaults / 121223211031 / 9

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

<a id="canonical-0112033021030132-0223101323111011-1033033110213022-1012130330013022-2100330032002310-3213331100130200-3320230220320102-2332021133222003"></a>

<a id="canonical-0131111031321312-3301032233131010-2200021102013213-3123123323021101-3231312122211222-2133313001310330-3113201330021203-1000310202333011"></a>

## snapshot_policy property — volume_defaults / 121223211031 / 10

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

<a id="canonical-2013013211000113-2210032220032203-3331120221333213-2120213032332103-3201101222133300-2232000300330302-0031120013133020-1001320330021102"></a>

<a id="canonical-0332023211312300-0320221331133113-2223123132233201-0332100223210331-0120221201121223-2220223212200310-2022232311030011-2303020010300231"></a>

## snapshot_reserve property — volume_defaults / 121223211031 / 11

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

<a id="canonical-1131103001312110-0301011201322322-2320031203021000-1203122033023222-2220310331230011-0310010032222031-2301011323022330-1122301030311233"></a>

<a id="canonical-0220112100102101-1123021320202121-2202222030100311-0223003233333333-3203233213320230-2322221131300300-1300021120232211-3301330103220222"></a>

## space_reserve property — volume_defaults / 121223211031 / 12

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

<a id="canonical-0303032022313303-1130333210233300-3331211123331110-3013032330132210-1332020303200121-0311113102211311-3112032111000132-2300332203331222"></a>

<a id="canonical-2312313213310011-1013302003201310-0021210321213221-1032112222002101-0002200310212001-1313211230133210-2031100122031221-0221223013333102"></a>

## split_on_clone property — volume_defaults / 121223211031 / 13

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

<a id="canonical-2121010300213221-2102232122313322-2212031331010103-3003121210103311-3011332133212333-2033113230130232-2110030300020123-0320011300100330"></a>

<a id="canonical-0213232330220222-1111232021210021-1120001332210013-1002223203332031-2200031331121031-2130133121130323-3032031210232323-2222133333201032"></a>

## tiering_policy property — volume_defaults / 121223211031 / 14

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

<a id="canonical-1012231311231210-1123311103133012-2202210332233113-3233321133100121-1113302033103223-3302310330303012-0310300301311003-2223013300120232"></a>

<a id="canonical-1210010013320133-3021301230311022-1331031232223320-0201133013110123-0223322300021220-1232100001221033-1030230111001020-0010300223310200"></a>

## unix_permissions property — volume_defaults / 121223211031 / 15

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

<a id="canonical-2110002032121020-2001033023120000-3002001030111022-2203012323313011-2230331023330313-2203110122020113-3213233020213000-0111300233130023"></a>

## Next pages — volume_defaults / 121223211031 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](data-sources--voltstack_site--reference--group-007.md#canonical-0001013031313310-0323331301322111-3003023333000132-2222303302203300-3033313031303210-2113012231320002-0012022121330310-1230312220311000)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0001013031313310-0323331301322111-3003023333000132-2222303302203300-3033313031303210-2113012231320002-0012022121330310-1230312220311000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312302103133012-3022010223310312-2301213023033200-2201301212033133-2031320022000332-2200020010212233-2103330102022103-3132020310111133"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos — no_qos / 100130202313 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--reference--group-006.md#canonical-2030101222310112-2322023220122310-1122303121121233-1300110133032012-0222113333333132-3330010230303232-0310020023232113-0022232130300213)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--reference--group-006.md#canonical-1223002030203333-1121000122302200-2230010303202300-3121233302003303-2133030131021103-2010131113031110-0023022010021231-1033211212023103)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-0010213121112030-2023010302202201-0032322030031231-2203123331233303-3331013031313301-2111303021000302-2003330020101200-0100231311031011)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos

<a id="canonical-3130121123003333-2210102331013133-1020102212112233-3212022033232230-0023212313301302-2323223201310313-2200212211331013-0332213011011320"></a>

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

<a id="canonical-2333311123102001-2000322133122103-1332331023230013-3233002012332332-2321223311222200-1312302131030011-1020021013323030-0020120331110231"></a>

## Direct properties — no_qos / 100130202313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102130311003200-3113320021101013-3300031023101302-0210100330202223-2133202110201021-1130313011233223-3101303032131330-2300333310210130"></a>

## Next pages — no_qos / 100130202313 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](data-sources--voltstack_site--reference--group-007.md#canonical-0010213121112030-2023010302202201-0032322030031231-2203123331233303-3331013031313301-2111303021000302-2003330020101200-0100231311031011)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213101102031001-0232131131003020-0112022101121122-3321221003312103-2301213201031032-2212231302210021-0123233023102010-3312111010022200"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator — pure_service_orchestrator / 133120111132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator

<a id="canonical-1002002323222132-1101121322302101-0011213000010233-2303111213311013-2023003203021330-2332110210011211-1112130020002310-3323132031322322"></a>

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

<a id="canonical-2320000332010113-1022112110230301-2213310332331112-3133312222013212-0002312031100033-1022020201321313-1233213311000013-1230130013030133"></a>

## Direct properties — pure_service_orchestrator / 133120111132 / 3

- [arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320): complete subsection reference.

<a id="canonical-1010111210011322-1031223223232033-1032131302202133-3312202210100331-0311130110202122-0310001221013012-1122023303000000-3132210133200110"></a>

<a id="canonical-1103001313121302-3022002301211102-1331131211320100-0331132022111203-1003200320310311-0211023010230311-0201013202321200-0212030311021112"></a>

## cluster_id property — pure_service_orchestrator / 133120111132 / 4

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

<a id="canonical-0211111302232322-3310213230201123-0322133323012121-2032020112123133-3321300022321022-0033123232020233-3013213030233210-1000331102311333"></a>

<a id="canonical-3020023322031302-2230113101110313-1312120022001031-1132100103131123-0220221311123231-1111310330033230-3011312100033133-2010001021202231"></a>

## enable_storage_topology property — pure_service_orchestrator / 133120111132 / 5

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

<a id="canonical-1001313023202231-2113301133320010-1210103123321023-3332123222003110-1330330122101302-1033103333221213-1302120303002333-2313033201222101"></a>

<a id="canonical-0121033223313311-3213212112310233-2232233011300131-2112002212200012-1332010233031221-3233102101103212-2022111123122223-3311311303232321"></a>

## enable_strict_topology property — pure_service_orchestrator / 133120111132 / 6

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

<a id="canonical-1323213033010211-3132322201312330-3113312223100020-0210113101313023-3223032330002003-2210122021301033-2220323013222021-2022101203023212"></a>

## Next pages — pure_service_orchestrator / 133120111132 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132210213120303-2222312210030220-3121213212021003-3122311111020322-0130133332032312-1313121123123203-3032230323332012-2020103030321003"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays — arrays / 011030213101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="canonical-0330230330101002-1223131133013231-3312232022200131-2320201201003331-2331031212320230-3023020103312113-1201300323333231-1031031301100330"></a>

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

<a id="canonical-2232110301133213-0310330000022302-0002130223332333-2121001233002332-3230001233312223-2031311011111031-2130220121000332-2222132030332202"></a>

## Direct properties — arrays / 011030213101 / 3

- [flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-1100313000200101-0012223332010011-3020023012132003-0312200200133313-1032021123210212-0303200133033331-1110123212331311-3113021011001133): complete subsection reference.

- [flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0013120012030132-2200031302323020-2123201300021221-0202103100003211-2102331133013232-3213032000012323-2320220110001131-2213021022123330): complete subsection reference.

<a id="canonical-3331023103332311-2033100210031101-0300102222300122-1133213221132100-1032233330103102-1222131133200003-3101112311323223-2011311333203031"></a>

## Next pages — arrays / 011030213101 / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-1100313000200101-0012223332010011-3020023012132003-0312200200133313-1032021123210212-0303200133033331-1110123212331311-3113021011001133)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0013120012030132-2200031302323020-2123201300021221-0202103100003211-2102331133013232-3213032000012323-2320220110001131-2213021022123330)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1100313000200101-0012223332010011-3020023012132003-0312200200133313-1032021123210212-0303200133033331-1110123212331311-3113021011001133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303213123310123-2223223101210320-0133303010220012-3302001030311310-1110111131222220-1012310101311220-2100100201132333-2012022213321321"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array — flash_array / 110300333101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="canonical-2320322110123010-1312130330012100-0221330033200213-1102103023021001-3121130120202133-3121021123320331-2302010212332020-1210212323121102"></a>

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

<a id="canonical-2100002331030113-1001013120311101-0102013233133311-2331231310031132-3313102001000103-2113131323023310-2102321210220313-3233120110103122"></a>

## Direct properties — flash_array / 110300333101 / 3

<a id="canonical-3031130023023100-1321012021212020-3300213112223133-1202131330322320-0323113120000121-2031331313221000-2330112233022111-2100302232321010"></a>

<a id="canonical-3021220003331311-0213132130333332-3232212333303131-1310322300221011-3203003001121303-2323120202222211-3230330131322311-0322100321232103"></a>

## default_fs_opt property — flash_array / 110300333101 / 4

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

<a id="canonical-0313112321031200-0312020003033002-1312122213233022-3032102000233323-3330000023110223-2200323211310300-1000032301203031-3321032321120302"></a>

<a id="canonical-1001003222111133-2223310030003201-1331002023211021-2300323222221022-1103231303101221-1311010202320023-1031133223323133-1103020013032332"></a>

## default_fs_type property — flash_array / 110300333101 / 5

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

<a id="canonical-0332020303130202-3310330202000322-3102230123303300-1121020100032312-2112301321122022-0322323313322232-1313221131003102-2022301301103111"></a>

<a id="canonical-0223321331210113-0200100303112120-0031103032232022-0030021012323131-2310100110100323-3231122003200033-0102222313012320-1323100023331131"></a>

## default_mount_opts property — flash_array / 110300333101 / 6

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

<a id="canonical-2220023210221131-3032022103012011-2023223312002333-3213221122103122-2212311221323101-2221010130211330-0300022011310302-3131123323021333"></a>

<a id="canonical-1312021031112122-0220303321233032-2201201321213020-0302210110122332-1031310333321031-3022003211032300-0021200213022233-1311111012312222"></a>

## disable_preempt_attachments property — flash_array / 110300333101 / 7

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

- [flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-1012202022030311-0003003103002313-2302301121102031-3131022312012133-1330311130102202-0312303331122223-3011202121003013-3130333333230013): complete subsection reference.

<a id="canonical-3213203033213231-2100102212010231-1020100220120003-2103003101301203-3332213001001200-3121303030321201-0231102102211122-2100221000232200"></a>

<a id="canonical-3123011122313033-1211331101123001-2103021013310222-2011132103033232-2133133211223212-2320100303000300-1301030332101101-0322100001020022"></a>

## iscsi_login_timeout property — flash_array / 110300333101 / 8

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

<a id="canonical-0320002330010202-3002002311311033-2111221200111112-2303232113210022-1231211002203011-1223233310010023-3101132310121002-1300323021221011"></a>

<a id="canonical-1331323200330323-2102203232121302-0101002031310121-0233030333213312-0022003203010321-0303033133112332-2211032332323002-0220112011131332"></a>

## san_type property — flash_array / 110300333101 / 9

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

<a id="canonical-3001333200221303-2202333313102230-3113012023200233-2011101103330303-2013302321222001-2213122003103322-3213030003230332-2333303120011030"></a>

## Next pages — flash_array / 110300333101 / 10

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-1012202022030311-0003003103002313-2302301121102031-3131022312012133-1330311130102202-0312303331122223-3011202121003013-3130333333230013)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1012202022030311-0003003103002313-2302301121102031-3131022312012133-1330311130102202-0312303331122223-3011202121003013-3130333333230013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000323102303011-1320200111303132-0012323302211233-2031030200302033-0012021100110130-1221102320001012-0001200021011033-3030022212121223"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays — flash_arrays / 300102103203 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-1100313000200101-0012223332010011-3020023012132003-0312200200133313-1032021123210212-0303200133033331-1110123212331311-3113021011001133)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays

<a id="canonical-2101101333031300-2231330202232210-2303210332120121-1331311323332100-3230333313200223-0033002321023213-3032030331012133-2021222013300022"></a>

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

<a id="canonical-0303200132012122-1130303320303003-1033013111130102-1320111313200032-0001233312130310-3332013102000222-0021030201230332-1023030021201332"></a>

## Direct properties — flash_arrays / 300102103203 / 3

- [api_token](data-sources--voltstack_site--reference--group-007.md#canonical-1011100322023212-3021120232021313-1221023102222221-2030102101012231-2203133121232332-2311233001321021-2001023132011013-2303212123111101): complete subsection reference.

<a id="canonical-1032213222332110-1133231132110320-2332310322231313-1001221130210102-3331123000220130-1000210123311233-1220331300001031-1123103230011213"></a>

<a id="canonical-1103122310101013-3302012203332123-2320022220232133-3001130031130313-0213112322131220-3003200030101332-3312312230023020-2220200202333212"></a>

## labels property — flash_arrays / 300102103203 / 4

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

<a id="canonical-1121310121331003-0131003113301010-2310311103202322-0211221311111210-2331032032032223-3301211132120301-3102230223002020-2210213332312210"></a>

<a id="canonical-0032303002233010-1223310310122331-0120000031331223-2223202103301212-0123121123302212-0013130002302200-0030330200120002-0302313121320031"></a>

## mgmt_dns_name property — flash_arrays / 300102103203 / 5

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

<a id="canonical-0013301003202311-0310330132010331-2232101112312301-2210021221000032-3211303013010102-3230222023200333-2133223303001200-1001222131031213"></a>

<a id="canonical-1033233313133122-0112303210012103-2003120000213133-1300013122022111-1221331331123330-1300030331322010-2312231233031112-0320101303033123"></a>

## mgmt_ip property — flash_arrays / 300102103203 / 6

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

<a id="canonical-1323301110313021-3201203130320103-3302310223203101-3321100003133022-0133301220100313-1320323120321032-2003103011031102-3220011102002112"></a>

## Next pages — flash_arrays / 300102103203 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-1011100322023212-3021120232021313-1221023102222221-2030102101012231-2203133121232332-2311233001321021-2001023132011013-2303212123111101)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-1100313000200101-0012223332010011-3020023012132003-0312200200133313-1032021123210212-0303200133033331-1110123212331311-3113021011001133)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1011100322023212-3021120232021313-1221023102222221-2030102101012231-2203133121232332-2311233001321021-2001023132011013-2303212123111101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122202102300202-3311031002113101-3313312032300220-2013130111200133-2202101031231032-2013130020011210-1213301020200011-2010230233131313"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token — api_token / 230300120123 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-1100313000200101-0012223332010011-3020023012132003-0312200200133313-1032021123210212-0303200133033331-1110123212331311-3113021011001133)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-1012202022030311-0003003103002313-2302301121102031-3131022312012133-1330311130102202-0312303331122223-3011202121003013-3130333333230013)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token

<a id="canonical-1200212332312112-3010012330200312-2102101001223231-2023230012200332-1102311001101330-1313221020313201-0301302332222000-0331303020023200"></a>

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

<a id="canonical-0130013321000121-0321013121320221-2322323211330111-2023202321232232-0122112003321011-3120123331002202-1031201112100310-0213013120031222"></a>

## Direct properties — api_token / 230300120123 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-2302112012020122-2022130100103012-3001310213333113-2332012111333310-0332022031213303-0322323322100003-2311012133022330-2221222101223131): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-0202221333333011-1001323112132032-3103112301000223-3331112033303102-3023023132123322-3120321013000211-2113133212032001-2321201310102220): complete subsection reference.

<a id="canonical-1310012130033310-3232331210302110-2222233302210332-3233001221001332-0322010303223110-2302113010020003-3321211021300200-0033131100233100"></a>

## Next pages — api_token / 230300120123 / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-2302112012020122-2022130100103012-3001310213333113-2332012111333310-0332022031213303-0322323322100003-2311012133022330-2221222101223131)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-0202221333333011-1001323112132032-3103112301000223-3331112033303102-3023023132123322-3120321013000211-2113133212032001-2321201310102220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-1012202022030311-0003003103002313-2302301121102031-3131022312012133-1330311130102202-0312303331122223-3011202121003013-3130333333230013)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2302112012020122-2022130100103012-3001310213333113-2332012111333310-0332022031213303-0322323322100003-2311012133022330-2221222101223131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331002301133020-2333303013130211-0012030232110101-2013022331331220-0223202301221023-2320320101131233-2230210002230223-2102203203121331"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info — blindfold_secret_info / 222010032212 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-1100313000200101-0012223332010011-3020023012132003-0312200200133313-1032021123210212-0303200133033331-1110123212331311-3113021011001133)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-1012202022030311-0003003103002313-2302301121102031-3131022312012133-1330311130102202-0312303331122223-3011202121003013-3130333333230013)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-1011100322023212-3021120232021313-1221023102222221-2030102101012231-2203133121232332-2311233001321021-2001023132011013-2303212123111101)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="canonical-3302120200011112-0322302232131111-0132031103233000-1103320302312221-2320112232131032-3011230212212102-0200023101031320-3200033013132023"></a>

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

<a id="canonical-0323332112112122-1333111023221030-1321020121102203-1003300121330021-2132212102132032-3322023322220200-2320100230321101-1330000233102001"></a>

## Direct properties — blindfold_secret_info / 222010032212 / 3

<a id="canonical-1103303321002333-3130010202013213-1103133011030031-3323103112331200-1100100131013113-3321013221130030-1231320131022311-3000303202313132"></a>

<a id="canonical-3011013211311120-2332220333211322-3133201301110123-0310201100213033-0221331112032002-3133230102212301-0102201310003212-3113312023131320"></a>

## decryption_provider property — blindfold_secret_info / 222010032212 / 4

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

<a id="canonical-0323322101221210-2322320132111223-0022330300220033-1212330001302131-2102120011020031-1311123010332132-0221202332131133-1232013323332030"></a>

<a id="canonical-1313121033312031-1133322130203120-2320220203202130-0002311121230220-1131021320233221-3011103123300212-2213333121200010-0000233312100203"></a>

## location property — blindfold_secret_info / 222010032212 / 5

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

<a id="canonical-0000120003130002-0303311033013202-1012130210011313-3312313300023213-2331032110302020-0203010103211113-1020113232121113-3032103030311220"></a>

<a id="canonical-2032223123123321-1210132220212031-0230102311131233-1000202112122120-3300302132022031-0301331132330132-0131123002312120-3223132033321300"></a>

## store_provider property — blindfold_secret_info / 222010032212 / 6

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

<a id="canonical-1003231323312021-3030203113200331-1123310033031220-3032013132120000-1313303122103111-0122011210223310-0233331312302231-1333121322333223"></a>

## Next pages — blindfold_secret_info / 222010032212 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-1011100322023212-3021120232021313-1221023102222221-2030102101012231-2203133121232332-2311233001321021-2001023132011013-2303212123111101)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0202221333333011-1001323112132032-3103112301000223-3331112033303102-3023023132123322-3120321013000211-2113133212032001-2321201310102220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113032131013313-2301330022233223-2220200000300133-3101301100112101-2001323203023332-2311033220332202-0310203010201101-1130030231133113"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info — clear_secret_info / 121211013223 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](data-sources--voltstack_site--reference--group-007.md#canonical-1100313000200101-0012223332010011-3020023012132003-0312200200133313-1032021123210212-0303200133033331-1110123212331311-3113021011001133)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](data-sources--voltstack_site--reference--group-007.md#canonical-1012202022030311-0003003103002313-2302301121102031-3131022312012133-1330311130102202-0312303331122223-3011202121003013-3130333333230013)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-1011100322023212-3021120232021313-1221023102222221-2030102101012231-2203133121232332-2311233001321021-2001023132011013-2303212123111101)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info

<a id="canonical-1031203123213231-3130121210231301-0313320031301330-1212301030220021-3011020101112230-2302322130000023-2210301103033132-2202312023121023"></a>

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

<a id="canonical-0032330032111122-3322312122010330-1202210200330130-1220212003333133-1310121003313100-0213122221321030-2311013220313320-3020303021201130"></a>

## Direct properties — clear_secret_info / 121211013223 / 3

<a id="canonical-3331221311020021-1111030123313330-0022231113010303-2311033132000333-1120021303323021-0131112100013131-1110221320203032-3333300213122213"></a>

<a id="canonical-1033333212323320-2223110002010301-3233111233030133-2210231023131033-3031101121011223-2333212113332100-2230210220232223-2110222121312000"></a>

## provider_ref property — clear_secret_info / 121211013223 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1003313321231111-3311322021310110-0001100323332312-2130231202012300-1313121033211213-3133110011031201-3323311323012120-2010231003013021"></a>

<a id="canonical-1001302110001022-2133232303232330-0233203333232200-1220323023313133-1231211330203122-1231021110301012-3122230313300000-2000003303321311"></a>

## URL property — clear_secret_info / 121211013223 / 5

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

<a id="canonical-3213322102220330-0321012203301310-0220101203003001-0310300010312112-2200131120133012-1223112112111203-0033310032130102-2100222311210120"></a>

## Next pages — clear_secret_info / 121211013223 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-1011100322023212-3021120232021313-1221023102222221-2030102101012231-2203133121232332-2311233001321021-2001023132011013-2303212123111101)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0013120012030132-2200031302323020-2123201300021221-0202103100003211-2102331133013232-3213032000012323-2320220110001131-2213021022123330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113001012211221-1113000100221113-0333031111122000-2031231303212223-0131333002122033-2313321212100320-0031303331100313-3222222032222320"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade — flash_blade / 202121311132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="canonical-3222113322111020-1312002031021002-0230101301300013-1322302020132002-0110330320020102-2110021122012122-2322130312202131-1311020322101230"></a>

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

<a id="canonical-3033203130123000-2303312330232131-2331223220030030-2123212110221012-2131310021130302-3203120330012121-0011012300123021-2123010001100021"></a>

## Direct properties — flash_blade / 202121311132 / 3

<a id="canonical-2310232102031201-2333031013010132-0230122212021331-1202313002332120-1223001100022320-0030011201000010-0030322010130101-3303212322332103"></a>

<a id="canonical-1210022013022213-1013323233013112-1030202213112323-2312322133202033-0231121323133112-3220212013210332-3123301102302301-3132232333333303"></a>

## enable_snapshot_directory property — flash_blade / 202121311132 / 4

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

<a id="canonical-1120310300110332-2202230103301010-0012202130203003-1233131121200121-2201211101333332-0011103333330221-2330020013213002-0231010020101310"></a>

<a id="canonical-2233323322123212-0032111300320102-2000301023230200-1133202020022121-0330031230033131-0200220313231013-2021222123300000-2311211010210332"></a>

## export_rules property — flash_blade / 202121311132 / 5

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

- [flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-2223011023223210-3223121222301031-3103021131010200-3211213000230230-2200321220133333-1220230110200311-2331011211312102-3210012201110011): complete subsection reference.

<a id="canonical-0301232131231101-1210220313333232-1230020100202123-1120231211012030-3131211312333113-0220301010212121-0210333000303003-3223002103001120"></a>

## Next pages — flash_blade / 202121311132 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-2223011023223210-3223121222301031-3103021131010200-3211213000230230-2200321220133333-1220230110200311-2331011211312102-3210012201110011)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2223011023223210-3223121222301031-3103021131010200-3211213000230230-2200321220133333-1220230110200311-2331011211312102-3210012201110011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001313130133203-3310132212100000-2222110302231032-1331010103311202-1033320211122012-3233302001003130-0312221112121203-0312213000031233"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades — flash_blades / 122322002023 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0013120012030132-2200031302323020-2123201300021221-0202103100003211-2102331133013232-3213032000012323-2320220110001131-2213021022123330)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="canonical-3023331102312113-1332303312002031-2220210030012321-2320223110113222-3110323133222331-0201300330023023-3122331221131200-3033220223133102"></a>

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

<a id="canonical-3310230321120313-0233130312111121-1010202121223322-0023210213210321-1022001201300120-1012302320112120-2231201131120311-1010210203103120"></a>

## Direct properties — flash_blades / 122322002023 / 3

- [api_token](data-sources--voltstack_site--reference--group-007.md#canonical-2220021100130231-3331211120311320-2003133210202302-0132112301123102-0322333010313210-3021203032331331-3210030112213113-1312230000221031): complete subsection reference.

<a id="canonical-3213110332102313-1320210100011122-3330313131033320-3100211331032333-2232300213311323-3002133101320103-1212201032330102-0202332331013033"></a>

<a id="canonical-3320122021310102-2311221313303313-3121213202031200-3220112012100210-3020331003232112-2330102333001322-3330201330303230-2021223320122230"></a>

## labels property — flash_blades / 122322002023 / 4

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

<a id="canonical-2031131300211033-0201333103020320-2212121023211233-0212220220320011-2321131022311123-2030000201123321-3233101030300232-0103010031310031"></a>

<a id="canonical-2301121120230222-2332000130033230-0212133202301210-3103003100003111-1310311002210002-3203022203001232-0301033123112331-2132221131023233"></a>

## mgmt_dns_name property — flash_blades / 122322002023 / 5

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

<a id="canonical-0332020320333230-3313320232232200-3113103021203033-0112113310103002-2131233031001000-3230201201211132-1130230323013011-1231002121123023"></a>

<a id="canonical-1033231023212233-2201210212312201-2231030102030320-0210202202102202-0132310320023123-0120121013323300-3233023320121112-0133212222321103"></a>

## mgmt_ip property — flash_blades / 122322002023 / 6

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

<a id="canonical-1100121210121300-2130332112020013-2322111100002313-0320032111120003-3232032231112102-2302202203301100-1213300231300200-2031123132333112"></a>

<a id="canonical-0223211323212301-2122331131101201-1032113310201031-0022031123230202-1032220110120302-3122201311012212-0031332310033200-1011002311333213"></a>

## nfs_endpoint_dns_name property — flash_blades / 122322002023 / 7

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

<a id="canonical-0130221112330230-2203300012320101-1200300330030022-3103203110211013-2130321030131112-1310103131132233-3233122132110303-3202331200112112"></a>

<a id="canonical-1013223130002321-3121032120120321-3133320223210311-0010301322333330-0123320031103320-0321110122102033-0310023213010310-2003313203303010"></a>

## nfs_endpoint_ip property — flash_blades / 122322002023 / 8

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

<a id="canonical-0311231013320212-1202102130322210-3322203103311010-0122132032210031-3122133012033132-2120032221000113-0331322333201012-3111022312330320"></a>

## Next pages — flash_blades / 122322002023 / 9

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-2220021100130231-3331211120311320-2003133210202302-0132112301123102-0322333010313210-3021203032331331-3210030112213113-1312230000221031)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0013120012030132-2200031302323020-2123201300021221-0202103100003211-2102331133013232-3213032000012323-2320220110001131-2213021022123330)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2220021100130231-3331211120311320-2003133210202302-0132112301123102-0322333010313210-3021203032331331-3210030112213113-1312230000221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020221332123203-1130231000332230-1001012231100033-1110011002203020-1123230320120112-2311211322120020-2132111230200321-3321021021222232"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token — api_token / 303322301101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0013120012030132-2200031302323020-2123201300021221-0202103100003211-2102331133013232-3213032000012323-2320220110001131-2213021022123330)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-2223011023223210-3223121222301031-3103021131010200-3211213000230230-2200321220133333-1220230110200311-2331011211312102-3210012201110011)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="canonical-0230001333211202-3301122203033112-1221212113311001-2123111101031323-1233300102023303-2102232312121100-2122130123031112-3023012022202113"></a>

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

<a id="canonical-1213320311112133-0332212032110331-1222022323302211-2001232102003030-0123313120330023-0112220030100031-1312102200032232-1303031231010333"></a>

## Direct properties — api_token / 303322301101 / 3

- [blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3232322031331211-3032331210131031-1313010120030001-3101003213213201-2220030301333212-0311211013332303-1013301213002212-3231100323310123): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-2221212201330100-1302120332112020-3020122200012032-3221302030211002-0213223112302123-3011013032100010-2213302001013230-2111000213033031): complete subsection reference.

<a id="canonical-3332132113301130-3102022123032000-2003220033310221-0200020112302321-1213030211211001-0122330332132211-1331111320311120-0033031110001323"></a>

## Next pages — api_token / 303322301101 / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-3232322031331211-3032331210131031-1313010120030001-3101003213213201-2220030301333212-0311211013332303-1013301213002212-3231100323310123)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](data-sources--voltstack_site--reference--group-007.md#canonical-2221212201330100-1302120332112020-3020122200012032-3221302030211002-0213223112302123-3011013032100010-2213302001013230-2111000213033031)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-2223011023223210-3223121222301031-3103021131010200-3211213000230230-2200321220133333-1220230110200311-2331011211312102-3210012201110011)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3232322031331211-3032331210131031-1313010120030001-3101003213213201-2220030301333212-0311211013332303-1013301213002212-3231100323310123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213302030031110-2020313032021111-1220103120000013-1112320211103000-0332232031220200-3210101012303332-2331120130111022-2331220120012001"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info — blindfold_secret_info / 100313202002 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0013120012030132-2200031302323020-2123201300021221-0202103100003211-2102331133013232-3213032000012323-2320220110001131-2213021022123330)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-2223011023223210-3223121222301031-3103021131010200-3211213000230230-2200321220133333-1220230110200311-2331011211312102-3210012201110011)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-2220021100130231-3331211120311320-2003133210202302-0132112301123102-0322333010313210-3021203032331331-3210030112213113-1312230000221031)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info

<a id="canonical-3002110010110112-1313113330322103-0122323331213112-3012301001003332-3333102222101222-3112123001223302-3202032123211013-3123201003123201"></a>

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

<a id="canonical-0302121311120101-3130232223230131-3033302211001002-2210101222120003-2120011023211213-0313030233120020-1323031133020120-3333133001013301"></a>

## Direct properties — blindfold_secret_info / 100313202002 / 3

<a id="canonical-0313232122223120-1023333130131131-2232033313003331-0213212103230102-1301223231110301-1022113011301122-2121320003022010-1010313100230013"></a>

<a id="canonical-3202030203313000-1020102001321210-2230303221211322-2031303121231011-1010202011231302-3123211021330233-2310313202002301-1330231103330310"></a>

## decryption_provider property — blindfold_secret_info / 100313202002 / 4

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

<a id="canonical-2110322031231321-0111001220011311-1200021031010323-2220030101013233-0302032302211103-3321031112112022-0332001113130012-3133213101323133"></a>

<a id="canonical-2012023330123023-2323321002103312-2301123013030103-3333232130012203-2021211302212222-3103220301123021-0330133220123112-1320033332113231"></a>

## location property — blindfold_secret_info / 100313202002 / 5

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

<a id="canonical-0232011010302032-3133210321122010-0211321133233130-0122023101113032-2000010320303111-2011012032110011-0001302231112303-1130101212130201"></a>

<a id="canonical-3103120200110103-0320132002122112-0120133000000203-0330230102310113-3313233132103302-0131003133022223-3003232101230110-3010331200033002"></a>

## store_provider property — blindfold_secret_info / 100313202002 / 6

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

<a id="canonical-0200211122022100-2301221110103330-1322022223100300-3132101100112113-0132120323132303-2332210110321203-1022312310001332-1000100110131120"></a>

## Next pages — blindfold_secret_info / 100313202002 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-2220021100130231-3331211120311320-2003133210202302-0132112301123102-0322333010313210-3021203032331331-3210030112213113-1312230000221031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2221212201330100-1302120332112020-3020122200012032-3221302030211002-0213223112302123-3011013032100010-2213302001013230-2111000213033031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223333130133111-3001312311012021-2210212302031300-3002330300220300-1133201033321232-0130030322123031-3233132232101123-1200223122311011"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info — clear_secret_info / 101000231310 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--reference--group-006.md#canonical-3033003202233020-0021303031003033-0310101133213200-3220100120130010-3033202030323020-2002330332213112-3223210003002103-1332320221001320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](data-sources--voltstack_site--reference--group-007.md#canonical-1213202000013010-0030123312132321-3133223302121201-1033033330203210-1301220100012200-2120201000202330-3122231300132132-1103013121220220)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](data-sources--voltstack_site--reference--group-007.md#canonical-3133310312133321-0021000011010120-3332121223202101-3003033300323033-0300112310011322-1303221301023321-3311133132323002-3232020311032320)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](data-sources--voltstack_site--reference--group-007.md#canonical-0013120012030132-2200031302323020-2123201300021221-0202103100003211-2102331133013232-3213032000012323-2320220110001131-2213021022123330)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](data-sources--voltstack_site--reference--group-007.md#canonical-2223011023223210-3223121222301031-3103021131010200-3211213000230230-2200321220133333-1220230110200311-2331011211312102-3210012201110011)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-2220021100130231-3331211120311320-2003133210202302-0132112301123102-0322333010313210-3021203032331331-3210030112213113-1312230000221031)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info

<a id="canonical-0321002223312233-0301230111211021-0322112010201112-3001032203203203-3303003232220102-2332100211100300-0111232232111330-1201130002222012"></a>

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

<a id="canonical-1110231132332111-2120302003330310-2131320200103013-3313321103323003-0002203032121100-1121331300201100-0322102113120313-1122011121232010"></a>

## Direct properties — clear_secret_info / 101000231310 / 3

<a id="canonical-3210200212330003-0202213023003032-0222211110230003-1210321302203110-0021110003311021-0222211333301323-3020013302223222-2003213112223320"></a>

<a id="canonical-1201032023320220-1311223222312011-2333312201122003-1302321013302333-0021031212110231-2130303323212101-2201220113213301-0220100103111221"></a>

## provider_ref property — clear_secret_info / 101000231310 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3312230020000201-0103022320302123-3303000112100032-3133211302303323-1221102001103122-1011202312323120-2112000101033302-2322130313222302"></a>

<a id="canonical-0020322031223123-2212221230200301-2123002301121332-3131100213001012-2023320000113100-1300113312223221-1113211303300232-3322220112001000"></a>

## URL property — clear_secret_info / 101000231310 / 5

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

<a id="canonical-0100210223302133-1111202201130302-3130000023313111-2333310230031220-1002132212023020-3233011103011001-2022330333130113-0312211101122001"></a>

## Next pages — clear_secret_info / 101000231310 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](data-sources--voltstack_site--reference--group-007.md#canonical-2220021100130231-3331211120311320-2003133210202302-0132112301123102-0322333010313210-3021203032331331-3210030112213113-1312230000221031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213231011020231-3000122310031022-2220012112012312-1031333130233231-3020212131213232-2103010203103120-1031231231231332-0100213301021331"></a>

## custom_storage_config.storage_interface_list — storage_interface_list / 233321333232 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- custom_storage_config.storage_interface_list

<a id="canonical-3133301121101220-2201301302020210-0321322303120330-2112311121300301-2300002213023302-0333032213201333-0313320030133100-0210002212013311"></a>

Type: `"single"`. Computed.

Configure storage interfaces for this App Stack site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1300103002200332-2003012331123320-1231000132101031-0033321103010321-3202331111102311-2132322002101112-3203022111332222-3301313111232010"></a>

## Direct properties — storage_interface_list / 233321333232 / 3

- [storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132): complete subsection reference.

<a id="canonical-0120211210303332-0313232200022032-3020322100100020-2332110220311220-3032321130320021-3003020313103030-3033013222232323-0133302133103133"></a>

## Next pages — storage_interface_list / 233321333232 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111310032002232-1231311312011111-3111223031102132-0120201302232101-1110212011121102-3002201211020211-2023333122300323-2122232123133210"></a>

## custom_storage_config.storage_interface_list.storage_interfaces — storage_interfaces / 320202232203 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- custom_storage_config.storage_interface_list.storage_interfaces

<a id="canonical-3333210301210311-0100023023031003-3311223330110122-2330202300310200-0312200133301211-2123312002022322-3310122002312221-2311313033022223"></a>

Type: `"list"`. Computed.

Configure storage interfaces for this App Stack site.

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

<a id="canonical-1333101131310113-2203330312200231-2330233203030231-3312132011133330-1323330231223331-2332131032313133-2202323112100211-2222001222100220"></a>

## Direct properties — storage_interfaces / 320202232203 / 3

<a id="canonical-0203330122023013-3033231103312003-3221213300313123-1022023012223011-2323101330012133-0300030321203123-1013020330101220-3331221030302100"></a>

<a id="canonical-2031010233223300-0010030102300311-1233312103213233-0100323112012321-2001220221223110-3301222000113022-0003100110003311-3133301203331103"></a>

## description_spec property — storage_interfaces / 320202232203 / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [labels](data-sources--voltstack_site--reference--group-007.md#canonical-1022011320301013-0223203220011323-2033112301003011-0201002202101013-1211301012112103-0220002300033201-2210001111323203-1221130212133320): complete subsection reference.

- [storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112): complete subsection reference.

<a id="canonical-2301111133220211-0103022013010232-3132230002211333-1133132123220123-2013300312233111-0331133221110001-2223300232012300-2213323303322300"></a>

## Next pages — storage_interfaces / 320202232203 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.labels](data-sources--voltstack_site--reference--group-007.md#canonical-1022011320301013-0223203220011323-2033112301003011-0201002202101013-1211301012112103-0220002300033201-2210001111323203-1221130212133320)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1022011320301013-0223203220011323-2033112301003011-0201002202101013-1211301012112103-0220002300033201-2210001111323203-1221130212133320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233023203031302-1001113122331212-2311300130131012-3302333323132010-2223102202030331-0020022332020212-2030013032221120-0000013112132203"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.labels — labels / 002311211101 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- custom_storage_config.storage_interface_list.storage_interfaces.labels

<a id="canonical-1110230201131103-2333213221001213-0210333100112223-0111121312233133-3020201023213211-1100103331303002-3130112101110211-3321101311111330"></a>

Type: `"single"`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1321110322011323-0200320000331321-0133303220320313-1230000010030003-3112121202023102-2121323312303320-3012133123133202-3223212001202032"></a>

## Direct properties — labels / 002311211101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213020103030320-1130203233233313-1221001301331023-3310112032222112-2310013022130220-3322211320133213-2302113120110010-0100002111110300"></a>

## Next pages — labels / 002311211101 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312323211113230-3230233222303110-0003322313322310-0011200000222013-3303032331233023-0102232302333031-2102211320200113-0312323203330111"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface — storage_interface / 231222300112 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface

<a id="canonical-2030131330111030-3132012023311102-1333032333030022-0322100103132003-0202001031121010-3132323312113131-1033112021212003-3102113231301022"></a>

Type: `"single"`. Computed.

Configuration parameter for storage interface.

Upstream description:

Ethernet Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

<a id="canonical-0310002130032022-2322130021310220-1203312000110222-0030311303122220-3120313103110003-0220100302330310-3102302220222233-3010221110300312"></a>

## Direct properties — storage_interface / 231222300112 / 3

- [cluster](data-sources--voltstack_site--reference--group-007.md#canonical-1001021030031020-0031312220331332-0303130102002300-0330311001320100-0301232313320333-0102220003001230-1311312200312213-3122313213213212): complete subsection reference.

<a id="canonical-1013322002131302-2202031102322232-1022223312011112-2132300232221110-2231003131003023-2301132221202110-1003211101110131-2121033022301320"></a>

<a id="canonical-0110121231200103-3230323013321002-2111211210130112-0223000003133222-3212230112323210-0113003201100212-3200323221011131-0011120200331211"></a>

## device property — storage_interface / 231222300112 / 4

Type: `"string"`. Computed.

Interface configuration for the ethernet device.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [dhcp_client](data-sources--voltstack_site--reference--group-008.md#canonical-3120112120113101-1213121321233012-0100130112310002-1223222201003330-3330020113332313-3010333233121130-2113202213030110-0033300233231013): complete subsection reference.

- [dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203): complete subsection reference.

- [ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122): complete subsection reference.

- [is_primary](data-sources--voltstack_site--reference--group-008.md#canonical-1220202132132123-2001001023131212-2022033031313001-1021013121202233-3222221323021101-2030033312333122-2031021320000120-0111303130010220): complete subsection reference.

- [monitor](data-sources--voltstack_site--reference--group-008.md#canonical-2110112312110023-0030332011121312-2133230302333032-3202002201323033-2223113200303030-2013321031230213-3313031022003300-3321210001020313): complete subsection reference.

- [monitor_disabled](data-sources--voltstack_site--reference--group-008.md#canonical-2030023310001011-2320112113220133-1112333003301203-1030201120132303-3213031312011323-3231203201022102-0122100332013323-0031211120021101): complete subsection reference.

<a id="canonical-1231232312330032-2320201120331322-0310223313111121-0320002332023101-3320303303220033-0032331223103000-0001133113210133-1132002031122302"></a>

<a id="canonical-3003201233332011-3132033321323120-2101203232011211-3223103301322100-1300212103111220-3133011032213312-0332311302032110-1113031202230230"></a>

## mtu property — storage_interface / 231222300112 / 5

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-0320032312103201-1103000230110221-2120111022313313-3231123111311033-2101321132232012-1111322130303230-3222102031321103-0123232111033011): complete subsection reference.

<a id="canonical-3123132123011112-1203131121333231-1130011000003200-3013123232021011-1101210230002001-0323211203222220-2210220220010201-1003211030103330"></a>

<a id="canonical-1102322313213313-1301311001333212-0212001100332010-3112120013021130-1211000032331001-2030100232011202-2110121103302000-0011132300332101"></a>

## node property — storage_interface / 231222300112 / 6

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](data-sources--voltstack_site--reference--group-008.md#canonical-3202123223312121-2201331233210200-2003113020033322-1122203000011230-3002212211321012-2120120300202323-0000303222101123-2121113222112011): complete subsection reference.

<a id="canonical-3110031020121323-1020311332301122-0310233210103100-1201113301312000-2322310003310121-1213302321111222-1100133203323002-1321131230023110"></a>

<a id="canonical-2121211001333232-2010303202300001-3102113100030322-0332032223001011-0132223102103030-0101100213301002-0200230030111230-1033012032003001"></a>

## priority property — storage_interface / 231222300112 / 7

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](data-sources--voltstack_site--reference--group-008.md#canonical-1102121322230302-0333023220300213-1230120212220211-2302312033010013-0002301002132122-3233222020101032-1023111203132021-1021122230301213): complete subsection reference.

- [site_local_network](data-sources--voltstack_site--reference--group-008.md#canonical-3010303102201030-0212132332013133-2020120320211111-0000011003030021-1212320132231103-0021223113211322-0210131310312221-3333032022023003): complete subsection reference.

- [static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2001102210110011-2301000210121201-1001131031120112-2200331212101002-1111010023012221-2020300312031230-3002311223111202-2113013123122122): complete subsection reference.

- [static_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-0201303232323320-3233203102011101-0232000020301320-2330223130030202-1211111211002013-3123331333003232-1310300331030201-2231230000000131): complete subsection reference.

- [storage_network](data-sources--voltstack_site--reference--group-008.md#canonical-3201101213330023-0130032023121030-2300130300313111-1303202211122101-0230110103133123-2120112313121323-2200033023122233-0200313213123323): complete subsection reference.

- [untagged](data-sources--voltstack_site--reference--group-008.md#canonical-0102000330232030-2013030113312230-1220203131223031-1112231133220313-0321203100221212-1322001103231112-0222311123022203-1300310312313132): complete subsection reference.

<a id="canonical-2200330202032120-1022011232013003-3232031010312322-1203111100023213-1322003122300210-3011333233333220-3133312222111312-1022303300031110"></a>

<a id="canonical-1033230203321332-3003333303130231-1022302231002023-3210131322122222-3220300030211011-3203133201033011-2233011003223032-1010230012123321"></a>

## vlan_id property — storage_interface / 231222300112 / 8

Type: `"number"`. Computed.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-2211132212333303-0210321220010323-3210103003110322-0110230110312131-0321330301133202-2002033311133332-0200001201323220-2201122123320130"></a>

## Next pages — storage_interface / 231222300112 / 9

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster](data-sources--voltstack_site--reference--group-007.md#canonical-1001021030031020-0031312220331332-0303130102002300-0330311001320100-0301232313320333-0102220003001230-1311312200312213-3122313213213212)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client](data-sources--voltstack_site--reference--group-008.md#canonical-3120112120113101-1213121321233012-0100130112310002-1223222201003330-3330020113332313-3010333233121130-2113202213030110-0033300233231013)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary](data-sources--voltstack_site--reference--group-008.md#canonical-1220202132132123-2001001023131212-2022033031313001-1021013121202233-3222221323021101-2030033312333122-2031021320000120-0111303130010220)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor](data-sources--voltstack_site--reference--group-008.md#canonical-2110112312110023-0030332011121312-2133230302333032-3202002201323033-2223113200303030-2013321031230213-3313031022003300-3321210001020313)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled](data-sources--voltstack_site--reference--group-008.md#canonical-2030023310001011-2320112113220133-1112333003301203-1030201120132303-3213031312011323-3231203201022102-0122100332013323-0031211120021101)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-0320032312103201-1103000230110221-2120111022313313-3231123111311033-2101321132232012-1111322130303230-3222102031321103-0123232111033011)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary](data-sources--voltstack_site--reference--group-008.md#canonical-3202123223312121-2201331233210200-2003113020033322-1122203000011230-3002212211321012-2120120300202323-0000303222101123-2121113222112011)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network](data-sources--voltstack_site--reference--group-008.md#canonical-1102121322230302-0333023220300213-1230120212220211-2302312033010013-0002301002132122-3233222020101032-1023111203132021-1021122230301213)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network](data-sources--voltstack_site--reference--group-008.md#canonical-3010303102201030-0212132332013133-2020120320211111-0000011003030021-1212320132231103-0021223113211322-0210131310312221-3333032022023003)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2001102210110011-2301000210121201-1001131031120112-2200331212101002-1111010023012221-2020300312031230-3002311223111202-2113013123122122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-0201303232323320-3233203102011101-0232000020301320-2330223130030202-1211111211002013-3123331333003232-1310300331030201-2231230000000131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.storage_network](data-sources--voltstack_site--reference--group-008.md#canonical-3201101213330023-0130032023121030-2300130300313111-1303202211122101-0230110103133123-2120112313121323-2200033023122233-0200313213123323)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.untagged](data-sources--voltstack_site--reference--group-008.md#canonical-0102000330232030-2013030113312230-1220203131223031-1112231133220313-0321203100221212-1322001103231112-0222311123022203-1300310312313132)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1001021030031020-0031312220331332-0303130102002300-0330311001320100-0301232313320333-0102220003001230-1311312200312213-3122313213213212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
