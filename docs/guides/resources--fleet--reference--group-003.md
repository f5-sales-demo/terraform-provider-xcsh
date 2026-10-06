---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-0313031320220001-1022121101320112-3101002202203311-0230200323122201-0100310223000021-2112000302223133-3231231321013323-1120320102230330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-002.md#canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info

<a id="canonical-1320123220002003-2203232200231032-3301031122212302-3122003300332103-3303331233110203-1302223022222111-0030311012311010-0120002332101312"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230212230222022-2312020221012010-1021101031330101-2013110211230213-0223110233302201-2332001101012030-3230021313302212-3010223201220333"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info`

<a id="canonical-0012222303202103-1030022203232022-1031300000330331-3322130231130130-1011212201110012-3311100003331202-1323301110333310-0112212110323021"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-2001103321320003-2233120032302001-0120133101032131-3032020323303122-3122021213013330-1202203022020033-3102100120030100-2211333311110001"></a>

<a id="canonical-2010230201013011-1210311131131210-2023021211132303-1002210023320000-2103323000133301-1032230000020230-3310320032223323-3013001013212011"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1321301321301330-0020122313331331-1032310130023223-0110301023103323-1133013313221230-3212331020012300-1311313200223303-3303202300011212"></a>

<a id="canonical-0032201133011022-2210203012030202-3101030313201211-1312231232132022-3023210310110200-3203210332210310-1130122002020013-0202221003330322"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-0103103020330320-2211332020322103-0111232101323212-2320200313030001-1201233210331002-3331223222012101-2133203032230302-1232113003313111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](resources--fleet--reference--group-002.md#canonical-1300320120013220-1111311020033332-2110201300332223-1330230012212101-3132130210000210-0310202102232330-2230320110022220-0211021223113033)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info

<a id="canonical-3233010122021111-2130130130011012-1321232312010131-3002100011220120-2133022111032011-3021203301111201-0103110110303311-0130302113321132"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230030001200103-1222333322233332-1332023020323020-0113233110130023-2022031220231033-1300012321233333-3133301211112200-1130101311121310"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info`

<a id="canonical-0023330102210212-2031110120200133-2030201223100222-0230332112020120-1013120322220231-1033220312320023-0103000200203101-1131230322332032"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0133202232200021-2122231020032100-3321302331221321-3211120221131301-3301322333210300-3332203210210032-0313102013122232-1332032021323233"></a>

<a id="canonical-3112213231320111-1120012320302111-2313212131022223-0233220312320303-1220131013002031-1101302023301023-0220113112330213-3012031100232110"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage

<a id="canonical-0323300303111030-2320210301123301-3031331000201323-2102132131230001-1220023300211123-0002211023213130-1220311130031020-3230300211333321"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122113201220021-1102201320331231-3012101112322200-3222030111000313-3001012030033013-1222320300220311-0310203033021102-3133033310203333"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage`

<a id="canonical-3113032131030330-3211313202123032-3130332011123121-3123100200302130-0300133332323112-1321320130121110-3022110213032121-3133322210232131"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.labels` property

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

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

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-2113320110201301-2101220231002013-3002022330010220-2021032231233120-3223323201302332-0302333022030323-1320122213233212-2012033130320012): complete subsection reference.

<a id="canonical-1111112113201000-1303200213203101-2032032301333011-1311133030021232-2323101133113011-3032022201201201-2111130221100323-3020103122303103"></a>

<a id="canonical-0221303032333333-3330012310232231-0122230310230022-0131112232333021-3132321200032010-1302011322310320-0330201011100313-2011010220300311"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.zone` property

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

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

<a id="canonical-2113320110201301-2101220231002013-3002022330010220-2021032231233120-3223323201302332-0302333022030323-1320122213233212-2012033130320012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults

<a id="canonical-0000132022102330-0231030023212231-2210011210103300-3232131333032210-3120303013311202-3023012311313331-2013301002010120-0032033232033201"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
```

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

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120000013111132-1320200020120213-2223123010002011-0103231010110211-0212210112130002-1010312103211131-3212002103311013-1232223211213120"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults`

<a id="canonical-1123313322201132-3201103312320230-2112321023323332-3010202112020231-2230023320323011-0200032231311022-1313231223302002-0201323330301231"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.adaptive_qos_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0330202233301311-1330201101233130-2220301010032200-1320002200310331-2313012333031101-0332031113033333-3321200312131200-2211311300002132"></a>

<a id="canonical-2020313200212122-0332302301110001-3012210320122021-1331210012022001-2000100231211102-1333303312021102-2230023003313213-0320022221112230"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.encryption` property

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2113330322231332-0002120111033202-0232012320313313-1320030220012233-0010110320120211-1033203000212023-3003031202332313-3112331100022131"></a>

<a id="canonical-3033113223312212-1330223201222120-2213310021301310-2210100211110021-1021102200323103-1101221301022102-0213312131100130-3330021013020322"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.export_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

- [no_qos](resources--fleet--reference--group-003.md#canonical-2313213003113030-0311202130033012-1012330231303212-2313221100211133-3311200223200332-3013012330201112-1030203113303131-2010233320301020): complete subsection reference.

<a id="canonical-3100311033130012-1330212011302313-0200201110132231-1220133130000032-3130332230222112-3310301121023133-1220300332232203-2231122220320232"></a>

<a id="canonical-1320330323010133-2231202111213120-1200231231310233-0033213232002120-2311220323211230-1223321301301003-0100012003101330-3231310332310101"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.qos_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3003310223321330-1332322213010222-1020201223033002-3233003103223001-0120020231233311-0202101321002001-1020020233221312-1302123333131101"></a>

<a id="canonical-1330311131302032-0233032213313202-0131123330121310-0100312302313203-0130221033122110-3223210021033300-3310002133033032-0110311021010011"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.security_style` property

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

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

<a id="canonical-1320011202312013-2113302320312302-2010221001130331-0033120100033000-1130231311321000-3102210112231002-2012032212022130-0233312020303213"></a>

<a id="canonical-1100203323102132-2021030310230001-3331033133130203-1122233032302013-2133321131220313-1033231132201013-2320230303102323-2330302100000212"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_dir` property

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311320330221233-0121331102122101-1333233233330300-2230132222032330-3320023113301320-2003201101103030-0122102131313332-1232222032330132"></a>

<a id="canonical-1132302112100013-3332311121300132-0000320100020323-3232200133222100-0213101022113312-0302011301010201-2303122031132132-3120013120010102"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

<a id="canonical-0321101321233220-2333320230121301-0113301002332022-2101121111033022-0203110011202231-3110330303303123-2022122131222231-0003022301022130"></a>

<a id="canonical-3021203230113233-3131122022112011-2000330232201101-2130313101111320-0021332231323230-3303013333311030-1021020223120033-2031011300320233"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.snapshot_reserve` property

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Additional upstream details:

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

<a id="canonical-1021021101001221-0200103022202113-3302100233210031-3302013022322301-3131231330323130-3112200233020000-2002202232320211-3210230303322003"></a>

<a id="canonical-3331000301132000-3133302313130303-1120100033230001-2022213312100333-1022033220312113-1123013302003321-2103202230111203-1200232232330213"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.space_reserve` property

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

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

<a id="canonical-2323101023330100-0201201012330331-1022232122101132-2133130323121332-2321330031110203-1330332310211000-3012102123232202-1213022303011320"></a>

<a id="canonical-1323311330012121-0322301131202223-1120111111203011-3301311133231333-2301122130032020-3312212102202330-0130323100113010-0011103220221010"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.split_on_clone` property

Type: `"bool"`. Optional.

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

<a id="canonical-3100021310223013-1023112233120302-2210203030020022-3030200301112000-1222010313313022-3333130133030001-3312101130231011-3332300300113000"></a>

<a id="canonical-1100000031020333-1303021130332110-3213132111002203-1101200020213223-1133200200001230-1122003033330133-3032300220032033-1002201110002120"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.tiering_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

<a id="canonical-2322020110010321-0033032323201203-2002323010021303-2330031222120100-1300330033230020-3102212133332221-0302310011312332-2331113333130203"></a>

<a id="canonical-2233330103211121-0220013212203220-0002310322310331-0303111001101321-0022220302023110-1322100010023320-0112323222133033-2001113202331313"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.unix_permissions` property

Type: `"number"`. Optional.

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

<a id="canonical-2313213003113030-0311202130033012-1012330231303212-2313221100211133-3311200223200332-3013012330201112-1030203113303131-2010233320301020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](resources--fleet--reference--group-003.md#canonical-0232211313301321-0212033213020113-3120021131102001-2231330003231332-2031221231130011-2302223013200320-3030323232233300-2202032103012300)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-2113320110201301-2101220231002013-3002022330010220-2021032231233120-3223323201302332-0302333022030323-1320122213233212-2012033130320012)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage.volume_defaults.no_qos

<a id="canonical-1111301331220233-3312003223022320-1323212032312332-0221033022323022-2221102121311121-3100031312312222-1233221313001301-0113103201310023"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_qos = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331201210221021-2100203101103023-3231010232113332-2331001322002230-3311121030003303-0210021021322130-2210111331110130-1031002103213331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults

<a id="canonical-0232310013002023-1113101303132333-1223201012331133-2300100203120113-2120003011000200-0120111033032303-3213002011031220-0003323310003012"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
```

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

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003110021311010-1323102113020201-2122122130300100-3002331311003232-2121323010002100-0033212003030222-3231222311023111-1130301023113120"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults`

<a id="canonical-1301310112321123-1103101102100223-3100202111102313-0202312230332100-2131202310321331-2222012023131302-1031000111121200-3300130223310233"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.adaptive_qos_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3221220112110001-3210323312300332-3302333312120201-2002213222212033-3212132112031003-2232002033333010-1312213200011310-2331002003110322"></a>

<a id="canonical-2301203210022330-2332233113313322-1323313333002003-2022123021201133-2313101100221331-3223111020311330-3000300111133201-2232120111210211"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.encryption` property

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3313032321030203-1012101202311213-3012030233011233-1122321112330110-1130311331233320-0030300011320130-0001333122033201-2312300201123302"></a>

<a id="canonical-3311211231212030-1301333332010011-0121132123331220-0310220320132010-2132100131330111-0202100233032230-0313322202200003-1313222122110222"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.export_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

- [no_qos](resources--fleet--reference--group-003.md#canonical-2233101330131022-2300110220203011-1230120022130033-2311302133323021-2033232112230323-3120331121123302-0131113211011002-0020113231123232): complete subsection reference.

<a id="canonical-1323211230012022-3210030032230001-0121213131022133-0331222030031312-1210123012231023-0220320201230323-1230232023110013-3231022112022120"></a>

<a id="canonical-3200202200202002-1013001210303012-1230311331222023-2211321233030133-3031000301223013-2131201310211233-1210103221121000-2110003122113013"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.qos_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3000210100233320-3030321021200212-2101320103302230-1330101123200012-2003021123001123-2211111001330002-0130300331332130-1223122010322122"></a>

<a id="canonical-1323312112310020-2031301111123120-0321300002213110-0233200231112300-3130233221300210-2211222202301213-1201303332230003-0003023232113033"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.security_style` property

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

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

<a id="canonical-1211301000102130-2231000200120211-2303313003120220-1120021301230112-0032110300130313-1031332120123030-3132010230120000-2332203302323122"></a>

<a id="canonical-0220001101333031-1102211101300032-1332102110312213-1333102331302122-2133131220003010-3321222013110103-1303011121122303-0230003320110303"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_dir` property

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1121311121212002-1200120212002203-2011102120031103-3003230220210022-3031132232221101-0330323010110301-0031333103330131-2123212111322232"></a>

<a id="canonical-1102111232110331-1002030321011003-0203002110012003-3220131021300112-2033320012022001-0313310211313011-0211201010233333-0321032232121312"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

<a id="canonical-1120023203110310-3021011102210320-2102213100322100-3330210232013310-3312112123230333-3332313203220000-3201202133233133-0203332021312010"></a>

<a id="canonical-1232323300030302-0112233223221131-3132301122322112-1331023011123300-1312223233022012-3022321101301222-1231102333121303-1312203000011101"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.snapshot_reserve` property

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Additional upstream details:

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

<a id="canonical-0322200131020031-2002313022011012-3303323230201123-2210023303232110-3000230003033033-2312233331230302-1032230031232220-3101130200313311"></a>

<a id="canonical-0022300111003121-1123203003120103-2030113131311333-1003001203032310-2330331122233200-3312111123231330-1322300313003200-0002002220102112"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.space_reserve` property

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

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

<a id="canonical-1132321301130121-2103311321210130-1032011211333013-1002211123312233-0302122032010210-2210313211023200-2000232111103331-0311122212123232"></a>

<a id="canonical-3000212000132301-1131211111001103-0102223031111002-1201303110303220-1202120213321111-1332230132222332-1302113200201011-1120010300002022"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.split_on_clone` property

Type: `"bool"`. Optional.

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

<a id="canonical-1301212333212322-2102120103213022-3201002132131123-2120133302011010-2132100320233101-2100202331112211-3123011111112122-2212213303320211"></a>

<a id="canonical-3101031312003313-3113130133320110-2112300232011221-3013220231323332-3313013122223003-1011203202132312-2210303133310001-2101333202001110"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.tiering_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

<a id="canonical-0223111201001002-0002122121223333-0020002021122221-0312300022013313-3301031030332211-3212010200321132-3200130031300213-3001322112303330"></a>

<a id="canonical-1333200221031132-1120200121010203-2211102231032300-3210110310320200-2231013001113320-1132030323311203-1321330302120302-2031312230320323"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.unix_permissions` property

Type: `"number"`. Optional.

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

<a id="canonical-2233101330131022-2300110220203011-1230120022130033-2311302133323021-2033232112230323-3120331121123302-0131113211011002-0020113231123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](resources--fleet--reference--group-002.md#canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](resources--fleet--reference--group-003.md#canonical-2331201210221021-2100203101103023-3231010232113332-2331001322002230-3311121030003303-0210021021322130-2210111331110130-1031002103213331)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults.no_qos

<a id="canonical-3002123320300132-0223331021300220-1210103031121331-2123031323213111-0011011022012122-0212310103013013-1233320313131233-0232112111213103"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_qos = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="canonical-3323210230230230-2010001013033130-1322132130110013-2001110133212220-1213311130133212-3130321202013113-0323121301331013-0123230022003303"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP SAN.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip"),
  validators.ConflictingObjectAttributes("no_chap",
    "use_chap")}
```

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

Terraform syntax:

```terraform
netapp_backend_ontap_san {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103132203202303-1232222020312202-3200020233210213-3122202011111122-1200223220300202-0220003122223111-3010120111323330-1230001020222133"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san`

<a id="canonical-1333220200332203-1102231001200321-2000221232122022-0220130302101023-2200231001302100-0232000312111132-2323131321132010-2132133033011210"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_certificate` property

Type: `"string"`. Optional.

Please Enter base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

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

- [client_private_key](resources--fleet--reference--group-003.md#canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321): complete subsection reference.

<a id="canonical-2002111031200132-3311123011032112-0110322030021311-0022332123323013-0011331321131221-2231102003123131-2122230212233012-0110011212102121"></a>

<a id="canonical-0123312221310201-3011302002112201-2122000102122112-1213320321233130-0311112220010320-0030013110030332-3010300301030230-2200021031120311"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_dns_name` property

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-2012111013231313-2100331100002022-1320213312331231-3222113322021232-2110300200223101-1333331120312331-3313231120310200-1322001011321300"></a>

<a id="canonical-0123233233021012-3223210300020220-2223333331320023-3223101130020132-2033303112232301-3301123111331131-3013333230231231-2332321020310312"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.data_lif_ip` property

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-2012011322130221-1123031013002011-2000133311221101-2123111023203221-1323332001233002-2223010010032112-3021000003310000-2111333002013232"></a>

<a id="canonical-2121130322010302-2111320232320233-0101103200220313-0303100001003033-2230110223132012-0130213120011131-2210030103330230-3133332312100100"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.igroup_name` property

Type: `"string"`. Optional.

Name of the igroup for SAN volumes to use.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1233203030100212-0313212330132230-2231021101303301-2310213112123011-0033133021101101-2300112212003123-0003120112302133-0101311231131222"></a>

<a id="canonical-1301033303102323-3001302021111221-1230323122113002-3020320222200011-0212211200001102-1200223222021030-0030133312130331-0013022300232122"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.labels` property

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

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

<a id="canonical-0131320010213210-1220000033003302-2312203332003003-0310033320002123-1000001013000200-3230032223120022-3033321322222013-1013300112103200"></a>

<a id="canonical-1200011333003332-1233313231100023-2031102123221131-0202311301032101-1023112312313001-2110132202001231-0132020032230310-1123331231022122"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_aggregate_usage` property

Type: `"number"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

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

<a id="canonical-3122211201223003-3021303220123210-0331301012310100-3103110210011103-2133220222313213-1300331312220101-2311131102010323-2330002322220121"></a>

<a id="canonical-0233320101213231-0132312101102323-3033003013122313-2332333232012201-0212102113322132-2002003113022111-2320132313023113-2100310013030122"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.limit_volume_size` property

Type: `"number"`. Optional.

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

<a id="canonical-0213102023320321-3313120133230223-0212213103112300-2103011130203320-2022033012100132-3122321200202302-2221130122322220-0323302313330002"></a>

<a id="canonical-1210213312120031-1302111112200131-2103000000211303-1210113231332221-1011020031021323-3110110022033100-3103312013220021-1102123220000332"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_dns_name` property

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-3321302332003231-2020020100012231-3102013012200321-3200033233110313-2202220203213011-0031320333333023-2202001110113300-2210223223310110"></a>

<a id="canonical-0020023032311103-1211022332023231-3000202010211311-2111002033102023-1123123000321303-0012132113301010-3003203330033021-0033000213023213"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.management_lif_ip` property

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

- [no_chap](resources--fleet--reference--group-003.md#canonical-2023212231132001-0233233303010331-1322332333230000-3022312201100210-0013202332122032-1001132113122332-0313002113102131-2232330213122212): complete subsection reference.

- [password](resources--fleet--reference--group-003.md#canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131): complete subsection reference.

<a id="canonical-3113021301202103-1120203201022013-3110131321022332-2321010111201202-0220123202230130-3133102021113222-2333233013003331-3321202212303110"></a>

<a id="canonical-3213012231310323-2301033323222101-0232211221210023-3001311021323201-3031231002200333-1303112120212123-3331030022201203-1312030331232121"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.region` property

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

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

- [storage](resources--fleet--reference--group-003.md#canonical-0132123032123033-1300031111020020-0122303313210220-0301213233313102-3000331003132200-0232211022231113-3313132300030110-0123333132302302): complete subsection reference.

<a id="canonical-1302220010102231-2313311230120312-0112100020331100-0031021323322033-2010110021211021-2302020303002012-3330323302221112-1300212220332022"></a>

<a id="canonical-3333132213110001-1211231131311231-1000213300103130-2001121323033023-3220330303203112-3210332133132120-2203331033232322-3121123210202233"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_driver_name` property

Type: `"string"`. Optional.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"),
}
```

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

<a id="canonical-2201002330330333-2002100130001202-1132110133003112-3313102201121221-1002332003103013-3111001332130233-3321230200111013-3100303203111031"></a>

<a id="canonical-2203001020003013-0233220031221212-1202110030311232-1031011102033321-0002221210121311-0322131032031111-1001002333110202-2103133103030321"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage_prefix` property

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 80),
}
```

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

<a id="canonical-1112213220332103-2331030233202121-3221310322030302-2203033311101133-2103220223012231-3222232222000313-2230203211003112-2022220020030332"></a>

<a id="canonical-0133001210331130-2013010330311230-1213213232210221-1322322212133123-0300212221221300-3102302312101223-0111013202231320-2302300011023303"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.svm` property

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1030122212211303-1033122101021312-1012213300312023-0001002122203202-3300123202301131-3100003011133013-3321313131100310-2022013213311032"></a>

<a id="canonical-2213032000303103-2320203211100211-3210023102100103-0100011103303033-2130111103100012-0230320022001121-1333232000311220-0101323303233331"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.trusted_ca_certificate` property

Type: `"string"`. Optional.

Please Enter base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

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

- [use_chap](resources--fleet--reference--group-003.md#canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013): complete subsection reference.

<a id="canonical-1022321112232130-0311002130232210-1330122220010222-1333111013033232-3120012303123320-3131322011310001-3021330000220020-3022323320101300"></a>

<a id="canonical-1132301303111332-3221313002031123-0020323023231201-0100020101103223-3330330112121331-1011101030200023-1002013001100202-0200111321133313"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.username` property

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-1023300110000030-3322320012222203-3312211121033330-0233123000012211-0313012003121021-3303103100101231-0022013323332110-0101320001330221): complete subsection reference.

<a id="canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-1333023203233113-2220031121323220-1231121203300023-0122013010232003-1220202332213131-1103300033133200-1102230102012002-2033131210130003"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311301232221000-1233223210212311-0112023001133311-2122330302122110-0112023332100223-1100223133033332-1320233302222203-3320012033330003"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key`

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-1220302221311032-2203010123221321-3213220210333022-3013321121313100-3230321211022333-2133120012002310-3012233113312320-3131330320030233): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-2133201323002230-0312012232123303-3112010321130320-1212102021103100-0300033231111230-0202002231132002-3332202023210003-3333313101202132): complete subsection reference.

<a id="canonical-1220302221311032-2203010123221321-3213220210333022-3013321121313100-3230321211022333-2133120012002310-3012233113312320-3131330320030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-1032101211332033-0323320121220133-0030003312332200-2302232033012231-1211320312213223-0122333220020121-1120010231313222-2211123111103123"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301302100310122-2023311332313120-0131022212013221-3121322332203032-1103203130331020-0230020210010210-0120023202310012-3103323232331312"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info`

<a id="canonical-0002330000000333-2100323201331121-3002321101203313-3110102123200333-0020000313303130-3202021233232233-1320003211001021-1200110222210303"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-2032213310223333-0013311123223100-3111132102303203-0003322311320032-2030321112030100-1320320130323200-2331310121031331-3330013320122121"></a>

<a id="canonical-1131211330002100-2301031103320113-0103100010220232-0202111123023201-0002023020031233-1202121001102303-0213313003231133-3320000202331200"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-3011122020112110-3020031131323020-0202000100122231-0233121000100330-0133310023323323-2133232223202010-0331030103132223-1231200221330320"></a>

<a id="canonical-1131212210312223-3032230032022232-0010122030121313-1202213033111132-1332111302123200-1222223123201033-2323223001233100-2232010211322312"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-2133201323002230-0312012232123303-3112010321130320-1212102021103100-0300033231111230-0202002231132002-3332202023210003-3333313101202132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--fleet--reference--group-003.md#canonical-2232002211332213-2313133232203222-2021113333022113-1101102033001221-0113320231122221-0233101120302221-1123133323322012-0101112100013321)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-2111210331213021-1030012213212013-0130210111021323-0011121310030321-0303010302333230-0100302222202033-3102121113230231-1100000232103210"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211101332201222-1213032302032012-0311012231331003-0123232003031000-1020302220332030-1230001302132322-0313133010110201-3121220323233010"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info`

<a id="canonical-3131103210331203-2220030230322322-1121010313023300-2012302011020301-3012223303213311-0020331021021002-0213202202010123-2102302102301012"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1332321023113022-2321210211203213-0313312231212203-2230333131200033-1303333013320213-3113020002212121-2313230130000301-1302213313331211"></a>

<a id="canonical-0303300130131312-0002123202113010-2202213223022211-1203021222110030-2032311032030100-2223113120322103-1103032131032032-0311321111313211"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-2023212231132001-0233233303010331-1322332333230000-3022312201100210-0013202332122032-1001132113122332-0313002113102131-2232330213122212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-2312231103301222-2303013202220233-2012103312203311-0302300223213112-1322111133033332-0310121011022233-1331102122201223-0120313021103321"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_chap = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-0311220320002233-2031131013333113-1230200132010013-0120222121311103-3212313300023203-0320022012312101-2320233010131023-2113321121332123"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110011311121122-3201300032331321-1210110001321121-2323003032032210-0201101010233111-2032203012101311-0023220202223120-1223213202031020"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password`

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0303321000302203-3000100313100310-2330132312013223-1102030110232131-0000101330011313-2231230031010101-0232312030313121-1102002233112220): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-2202220231210203-2000203110211122-1331231222330313-3123032313032231-3022303121212123-1121112330122201-0022010131023200-0032300013101100): complete subsection reference.

<a id="canonical-0303321000302203-3000100313100310-2330132312013223-1102030110232131-0000101330011313-2231230031010101-0232312030313121-1102002233112220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-0003322203023300-3012202302303003-1011322203212103-3310323103322303-0113121312302131-3113310101110330-1310000101320233-2103300022331221"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233101120232002-1120132101310132-2030322000021113-3213123012021101-3213011303333012-3122003030112113-1313203023213013-2312010221123302"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info`

<a id="canonical-3030322023203110-1012123230012320-1323010223002002-3311300301131233-0310313202312322-1221112030332123-0020311202111011-0020321231023301"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-3123003120331232-2033301223111112-0321030133113031-2231212012322030-3001101302101123-1103033210230203-2012230323123230-0310203333230223"></a>

<a id="canonical-1030012011221311-1233300023321310-3100310220303031-3113230312001100-3232013021333211-0133100103102031-1023110202231302-2030202012120332"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1300131113120232-0222132210030312-3023012320330301-3111101130300101-1022202121022013-3213032110331221-0011203330223113-0030301212123312"></a>

<a id="canonical-2313133222021012-0302330221012131-0002323203033021-1133203312302200-0111010002210223-2100111220332023-3133313200032003-3122320020133123"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-2202220231210203-2000203110211122-1331231222330313-3123032313032231-3022303121212123-1121112330122201-0022010131023200-0032300013101100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--fleet--reference--group-003.md#canonical-0332110131332013-0220121311102123-0210223010320022-2100123103202201-0131201101033031-0221312322121031-1323200232310111-3120331210320131)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-3231023312310010-0111102221002032-2121103001212000-2333130132120032-1223001032131223-2330230220233221-3021000010120311-2003030221203101"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302032112231113-3102310122221030-0223011121313221-3123032230232033-2102223311111121-2312121130230332-1031303001101322-0121033312021121"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info`

<a id="canonical-0130220311301111-3201001110102323-3311300222113302-2322132000223120-2330111012311320-3230003321112012-1301322313221213-2231333232022303"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0001230103223201-3103313212232313-3303202101333031-1332331332003120-2231101233330301-2332102301202003-3113223013332031-0323003011302313"></a>

<a id="canonical-3003100131030203-1313020102022202-0012111310200122-1002012223213221-0111102000330201-0303133211200030-3133031101002120-2310120122113131"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-0132123032123033-1300031111020020-0122303313210220-0301213233313102-3000331003132200-0232211022231113-3313132300030110-0123333132302302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-0211032110321131-3312003313011133-1233331030301322-3312031133133231-0213230312110110-0032012133303202-2321030123303101-2333231131032301"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232030311311221-3111320121321131-0301221111201203-3023330130013300-0121302323111321-0211333123010123-3103011120023112-2112222122311230"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage`

<a id="canonical-1031323230331001-1101112001113021-3003130101130201-3212102133311110-3211331013210001-3131231313220311-3011323000230200-2031233023110202"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.labels` property

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

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

- [volume_defaults](resources--fleet--reference--group-003.md#canonical-1331203201033220-3033100023220013-2010010023203222-2210330212120313-2310202223112112-2033230330001120-3220221121000221-1210120221222223): complete subsection reference.

<a id="canonical-2223300130001123-1020221130012330-2221203113311211-2110033211130301-1022321230301010-1032100301011002-0030213103210201-1220030110231322"></a>

<a id="canonical-2111322322231121-1310032223220110-3222110133211331-1332011010020220-1330002220303222-3032201220012213-2023120303001111-0133203023303120"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.zone` property

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

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

<a id="canonical-1331203201033220-3033100023220013-2010010023203222-2210330212120313-2310202223112112-2033230330001120-3220221121000221-1210120221222223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-0132123032123033-1300031111020020-0122303313210220-0301213233313102-3000331003132200-0232211022231113-3313132300030110-0123333132302302)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-2301022122323310-3321013323230120-2303310023223021-0120321200233130-2001120302001031-2010011333030111-2013212312311212-3000313000023301"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
```

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

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313001022132122-1113121033300002-1220110332122001-0033212033033132-1033303222113212-0220202210030212-1102000230110230-0033301221012210"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults`

<a id="canonical-1022212002223031-2221112020301210-1303300102113000-1020303122020010-3100003102012010-1031233301131313-1110033022022003-2130313101303201"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.adaptive_qos_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0330330022322210-3301101000302032-2132023210033101-0223111032030301-1030330213312101-2232021313233320-2330102021023110-3021223023131222"></a>

<a id="canonical-2013332323020033-3123121212110113-1213110310312221-2033203003112021-0023311301032212-2203123013112101-1303101033021112-1331123103001300"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.encryption` property

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2213021120201112-2003101113230331-0303213102200212-3002333313221013-1123212231202012-2122202023122230-2331112010023211-1210230113103131"></a>

<a id="canonical-3021313212030223-3121213213302312-3100030120221130-2120012223222101-1221313003133101-0102103023312023-3332212131122030-3010011032021300"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.export_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

- [no_qos](resources--fleet--reference--group-003.md#canonical-0230132003233011-1302200101120032-3302332032133012-3132303201232302-1223221123003013-2101021331312201-1332131002010030-2012322020112120): complete subsection reference.

<a id="canonical-3013310231230110-0213300133130022-2011011333202031-0123013132133012-1223333022121122-1210300330233101-0323113022030300-2312011313303321"></a>

<a id="canonical-0021322001333321-0022331022233211-1022221023012122-3213321133302121-0311320100322031-1322301211003212-1300203133330002-2232310200220032"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.qos_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-1310101232210213-1000222222131132-0133121323130303-0033233011121323-0210133022201023-2201322302232202-2303131002101211-2212100123203302"></a>

<a id="canonical-2033202030030212-1222221122121313-1312331233221113-2100032112120310-2123113111131323-2031323003023023-0131311002013020-1320330323100211"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.security_style` property

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

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

<a id="canonical-1303130223112010-2020311230033012-1010200302212311-1132131223113010-1000002311211003-0230333331023233-1010210321021131-3321031221021130"></a>

<a id="canonical-1112312221221002-1300222331201003-2312201310131110-3102003102112100-0023021101013310-3022121131201102-3132320101233132-1311322212310331"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_dir` property

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122312112000331-2013032020102312-3301120311111221-0323122213002000-0331132232122123-2003222330221131-2333313311213132-3110321122023323"></a>

<a id="canonical-0002002221013302-3302110231222211-1110212111332210-1200133002003300-0311202133002212-3301330220033103-1322302310122202-1030312313312022"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

<a id="canonical-1110102000233101-1223012221010300-1222313202211233-1031320301000003-3213331313022303-3122201313201001-1023112303020202-2221100201010332"></a>

<a id="canonical-0110222310113022-0303310003033121-0230102030003022-3120033312221122-0223302032031212-3202030023330222-2220121022000010-1231020333003103"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.snapshot_reserve` property

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Additional upstream details:

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

<a id="canonical-3112103233200230-0313213130201003-2202123221301131-3121223010212031-3201233333013003-2123223212023203-0111113312033301-1012201103313230"></a>

<a id="canonical-2223031211133302-3312321010021333-0322300202300303-1220013220022013-0101221201112111-3032222012112001-1100333130122131-1010220132223311"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.space_reserve` property

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

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

<a id="canonical-1103301002033212-2022222310112100-2131310122213301-1032320212232132-0220220222311223-0132112112313211-2132021310211011-1203222302313101"></a>

<a id="canonical-2112330222330300-1333210230313220-2333103311200000-3011230011200322-1201202012022031-2311300013031313-1330100201111203-0033213310122312"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.split_on_clone` property

Type: `"bool"`. Optional.

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

<a id="canonical-1320221323223103-1110012330110123-1023130132003333-2223102212222231-2100010222003201-1332213203222210-1202023032120332-1203033030030230"></a>

<a id="canonical-1101002032020002-0020312133201323-3213010120302233-2211113221201332-3003130312332003-2310131301001101-2303200321130111-1122321010010312"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.tiering_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

<a id="canonical-1012021000321233-1321020100222120-2001333133203313-3231332221020110-1300313202021020-0101210330030221-2323312100112121-2111332012010101"></a>

<a id="canonical-2022310121220110-1212022302210020-3201020013310312-3212230321001023-3111123230311211-0021022130223023-0231013011211103-0001203102211203"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.unix_permissions` property

Type: `"number"`. Optional.

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

<a id="canonical-0230132003233011-1302200101120032-3302332032133012-3132303201232302-1223221123003013-2101021331312201-1332131002010030-2012322020112120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--fleet--reference--group-003.md#canonical-0132123032123033-1300031111020020-0122303313210220-0301213233313102-3000331003132200-0232211022231113-3313132300030110-0123333132302302)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--fleet--reference--group-003.md#canonical-1331203201033220-3033100023220013-2010010023203222-2210330212120313-2310202223112112-2033230330001120-3220221121000221-1210120221222223)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos

<a id="canonical-1133003221223230-0303032211031120-3022301113132322-2030233300100020-3303121320031231-1323233311120011-0123023231130030-1111232031201013"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_qos = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="canonical-1133032223232231-2303221011010102-3232213321303033-0003320130211121-1213101030131133-3111310233100132-0202031331022310-1132230120303303"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_chap {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312323303130222-1131103213031121-1320032301322103-0201100130033103-0311123000000320-0013323212332102-0320202101210032-1303012303123333"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap`

- [chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-0033032203023221-3230223303120013-3120220030200030-0022302213110221-0011310133330111-1221021010221230-0021101120100321-0002031102332020): complete subsection reference.

- [chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-0302303232000211-2330001303022213-1220001222213313-0212232010021002-1331020111110122-2031200222303321-2311031120312300-1223011001120331): complete subsection reference.

<a id="canonical-2103013301232100-0221031213132010-3123220101022120-2203113032322203-1112313132030322-3112003331200031-0212031321113300-3102321031011013"></a>

<a id="canonical-0302231302202101-2131101322031112-0331322330113000-0222221303000123-3130303111010201-1331122202302012-2010131221121011-1011010033111222"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_username` property

Type: `"string"`. Optional.

Target username. Required if useCHAP=true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2100321023203102-2230022130112113-3022031232233322-3320201303330133-2001023212101322-0113202323301113-3121233032122022-3110330200112013"></a>

<a id="canonical-2132031222231030-2123110110202033-2313120233133023-2322213221310333-3213310113232112-0003132220210002-3201320301212330-1021212223100212"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_username` property

Type: `"string"`. Optional.

Inbound username. Required if useCHAP=true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0033032203023221-3230223303120013-3120220030200030-0022302213110221-0011310133330111-1221021010221230-0021101120100321-0002031102332020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="canonical-3003210202000223-3110123111132321-0210001220333232-1120002022022300-2031300223111011-0310211303323230-3333312113012331-3300210120023030"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
chap_initiator_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200030011201321-3101223021201010-0302320321301300-2322103011103332-0200321323311200-1330031020300231-2123223112023311-0001322322032123"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret`

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-1311022310213221-0333131231320002-0000033220210132-2103102300130033-3011330200113211-1133033231030033-3013210121001132-1201200310011320): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-1302003100223322-2113010313033312-0113123102230222-2222212203323032-3101311023013330-0213232223130200-1032311121000011-2102230310113110): complete subsection reference.

<a id="canonical-1311022310213221-0333131231320002-0000033220210132-2103102300130033-3011330200113211-1133033231030033-3013210121001132-1201200310011320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-0033032203023221-3230223303120013-3120220030200030-0022302213110221-0011310133330111-1221021010221230-0021101120100321-0002031102332020)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info

<a id="canonical-2113232130020120-0223210300032311-3132221222002303-0132322012300200-0131323110111002-1201130222111302-3203332332230022-3320010322120030"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232201003003130-2201021310112230-1333311033312222-1210200110133001-3210013022102321-1032303302330330-1311230223203001-0110020303100201"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info`

<a id="canonical-3032211003021020-1220203130030132-0231120122110112-3031012302030033-2231033103200320-1211110322300222-3320313321001130-2110130112230132"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-0333031232023020-2300221312011131-3311310322233330-3313123300202200-1220322120203322-3200321121333320-3212011231100220-0030022011002300"></a>

<a id="canonical-3113312011030032-0232001220031132-2020233300300022-1213330111323210-1101002302032100-2101023201210110-1033103331322200-1313011012320113"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1112101223312123-0311100031331300-3100300301320122-2321313213312202-3013123313001212-2221030222100221-3133232123233302-2200000233113123"></a>

<a id="canonical-0022233331313311-1123320121102012-3330112300102020-3333310022223232-0223330310110031-2331111321022332-0231312210220030-0101303233102123"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-1302003100223322-2113010313033312-0113123102230222-2222212203323032-3101311023013330-0213232223130200-1032311121000011-2102230310113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--fleet--reference--group-003.md#canonical-0033032203023221-3230223303120013-3120220030200030-0022302213110221-0011310133330111-1221021010221230-0021101120100321-0002031102332020)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info

<a id="canonical-2121012011320310-3203302322321322-3320323000202022-3210323223000022-3303311333003122-0230013331113311-3323220003110213-3103032203021100"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032002130202120-3020031002113332-1222221101030031-1200200120121021-3303232111202100-3033311100313303-2210033023023130-3101311100222100"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info`

<a id="canonical-3103010221321100-0210021113231222-2303033220312212-0001233331300021-2302000112313100-3320323310012201-2113200122222110-0302100010010021"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1031330210211321-2002001031131302-2303323000203302-1332302210103030-2121210100010110-0202233220101212-1003221103212111-0120211320132313"></a>

<a id="canonical-2022222013210013-3303312211223113-3020133002002002-1210220202201132-1330230020320120-1103211221003230-3003101003312023-1000301212312010"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-0302303232000211-2330001303022213-1220001222213313-0212232010021002-1331020111110122-2031200222303321-2311031120312300-1223011001120331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

<a id="canonical-2203212113102333-0300233230302032-3020200303301320-1111120221332002-1121330102000322-3303213311130103-1132002311232013-0310303110320301"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
chap_target_initiator_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101223230220012-2113023010333030-2013010333202210-3012201332113232-0112202303230103-1211211320032113-3112020033321332-1323310121021113"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret`

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-0121302132030023-1013101303330133-2103323131013021-3211211231201332-2332022000230332-2010100122322313-2210110003033232-1333102111221121): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-1013220111023223-2002000311000303-0211210300200223-0012101221233121-2033110300233211-2100101033130202-3320222231332312-3132301132022211): complete subsection reference.

<a id="canonical-0121302132030023-1013101303330133-2103323131013021-3211211231201332-2332022000230332-2010100122322313-2210110003033232-1333102111221121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-0302303232000211-2330001303022213-1220001222213313-0212232010021002-1331020111110122-2031200222303321-2311031120312300-1223011001120331)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="canonical-0120222320012302-2020232212300003-3020323130332312-1131223130221321-1020133103220310-0110120013203122-1031020032222222-1100203201303233"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321130231233030-1321132003311312-2212220021313202-1230222210203021-1230031330131033-2011031210122012-2323312003122230-2230210321103031"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info`

<a id="canonical-2001300022010103-2301113102302022-1203013311032231-2102001232100033-3230000012233331-2100210003103301-2110011033001231-2223311311022012"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-0201020020101302-2121100130211133-2212110022221110-2032123110102331-2230212013002200-0313011101023322-1000332301320001-3323023002211022"></a>

<a id="canonical-0113321123332301-2320223231112220-0210322203321102-3020100021100211-1300200313220012-2033312001320020-0101003123201123-1322322120233001"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-2133300000223331-1030211311021131-0210003302231120-2012123220233233-2013303132221331-0030120223322112-2303333133021203-2130320122220013"></a>

<a id="canonical-1131012012213012-3101301003302330-1011001223100130-1030120211110332-0120232002310321-1112303013321022-3301322000022101-2030132311012213"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-1013220111023223-2002000311000303-0211210300200223-0012101221233121-2033110300233211-2100101033130202-3320222231332312-3132301132022211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--fleet--reference--group-003.md#canonical-0210223030221013-1133322303123113-0112201303102112-2213010000130321-2133111120133330-2233132011002112-3313333201032133-3311001320131013)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--fleet--reference--group-003.md#canonical-0302303232000211-2330001303022213-1220001222213313-0212232010021002-1331020111110122-2031200222303321-2311031120312300-1223011001120331)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info

<a id="canonical-0221032110302003-2203100123000322-0113102310010022-3032200301003232-1102201030302323-1312002231022000-2332320031113230-1211313112222221"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310001302203300-3330221112013102-1231303322122313-2101033011221223-1111121023103320-1030201330300232-3222022233301102-0001303223323031"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info`

<a id="canonical-1222023023332333-1133011301320210-3023213132101332-2030200230222032-2222323230322002-3113031001230323-3001320110000331-0100111311312113"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0231303220332321-3212002132011112-2103003330230313-0312203002111232-3030123332200132-1311232000300002-3112030220033031-2200031131311102"></a>

<a id="canonical-0302121022332201-3132233102113223-2221222031312000-0230113231012121-2220233211301332-2012300211313230-0110131023033322-1200223231210201"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-1023300110000030-3322320012222203-3312211121033330-0233123000012211-0313012003121021-3303103100101231-0022013323332110-0101320001330221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="canonical-2200223213222001-2013131101001201-0220221003001211-0131132313010112-2131033300030311-3210211013203310-0221033211333312-1031020230330220"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
```

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

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330231310203100-1212232030121300-0213331321302330-3102222303213031-0203121313111033-1230221131132321-0330120233331322-1301230303011311"></a>

### Direct properties for `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults`

<a id="canonical-1233033221100101-0303200212223302-2132120001021223-1231200232003112-1332110220030022-1111212332022202-3321223233033102-2002320232200211"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.adaptive_qos_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3320320012112032-0030112300230030-2112321013123101-2333321200231003-2223222311320020-2212033333020212-1301323232201011-3120200120301201"></a>

<a id="canonical-2113010221313331-1232023101011201-3122212332303323-0311323112223133-2213021203023121-3310021231100333-1003130221032102-3210010230202320"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.encryption` property

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1102122032033303-2133123121332312-0021213302202020-1322020301220200-1333122200020220-3212003302003012-1020012302131101-1013001100013200"></a>

<a id="canonical-2021222332213131-0030030210122100-2130233030112123-2112222231121213-3232212333211132-3201230001123310-2023321123121230-2111101211032003"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.export_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

- [no_qos](resources--fleet--reference--group-003.md#canonical-0010013200030223-3310333100232211-2233333121120233-0323111122200013-0213032020121213-1313323313010311-3031303002020332-3001033202332322): complete subsection reference.

<a id="canonical-0322011223022022-2122231232231022-2113122011130013-0003130232302232-3131012012212333-3112130102112310-1000001032101111-1211012110032121"></a>

<a id="canonical-0012322213013210-1130322110231002-2231003011033002-2111111112013210-0200231010110202-2101310013230210-3201110230011002-1100330221213302"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.qos_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-2310333231000300-2332130023313201-1210201002133133-3333312101312230-1313323200001012-3001003230021112-0312113122301220-2312100020303030"></a>

<a id="canonical-2123112221011210-1301301301223322-2012212222002110-0313001033311221-0010311130132002-1311223323021212-1011311222030130-0333113221123212"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.security_style` property

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

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

<a id="canonical-1321313120000012-2213101203110002-0010333033021221-2221120232032202-0212123013233001-0133033113221203-0231213032112201-2100230312121130"></a>

<a id="canonical-3212131230003330-2103032312013011-1321022330011122-0133203303133131-1210013332300321-2031330021012331-2232111021012311-0212323230322032"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_dir` property

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100000203103031-0331131022003323-3030213023001223-1211023230133200-0233220233201212-2231031303201012-0000131323210222-3023233330302303"></a>

<a id="canonical-3232211021311320-1211222200012123-0001000022003001-0113223200102132-1300333233110232-0121333033130012-3132223023300310-1300012211212123"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

<a id="canonical-1013120300222302-3132231001320120-3212210101231302-0120213131230323-2300120102012302-1203203302212302-1322021302032223-3301312131013202"></a>

<a id="canonical-1101123023022011-3023002200230000-0000310010010012-1030202313201331-1331102300230133-2331132122320201-3312210332002132-0210310330000133"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.snapshot_reserve` property

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Additional upstream details:

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

<a id="canonical-0333120331121113-0123202130122202-2111302020132003-0031210101222303-2133130113322221-0000002012313012-2111201201112223-0301022200012220"></a>

<a id="canonical-3032131212030001-3300330221213132-1100201221111321-0113230130033211-1212331301130010-3323310102302220-0013102312213102-1203021003013022"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.space_reserve` property

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

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

<a id="canonical-1231010021221211-3031113203120211-0222202200311101-3210332030213131-1011031300223231-1022133220003102-1011123302321032-0301312000001233"></a>

<a id="canonical-0013111321120132-3001312100212102-2210322120133320-1210212233211122-3123232111312232-0211313122201002-1221203132103031-2222213032112032"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.split_on_clone` property

Type: `"bool"`. Optional.

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

<a id="canonical-3220300000313321-3223310230331233-3213010031032332-2331332203203303-3113302132120231-2200112101330032-3120221301220020-1203030000112031"></a>

<a id="canonical-3121131321103201-3322313033302213-3111320212010130-3332221121013301-2211110113223221-1130111213212002-0002023003233000-3221133302322302"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.tiering_policy` property

Type: `"string"`. Optional.

Policy configuration for this feature.

Additional upstream details:

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

<a id="canonical-0210210300313230-1221233221203300-0213323321030210-1133121200213230-0213231212020112-3123122002133301-0222323031331110-1122102332200212"></a>

<a id="canonical-3233100021120022-3223201130323200-0223021210233002-2331200213131303-0210122332110302-0311222111201212-0102230300221111-1012103233331132"></a>

#### `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.unix_permissions` property

Type: `"number"`. Optional.

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

<a id="canonical-0010013200030223-3310333100232211-2233333121120233-0323111122200013-0213032020121213-1313323313010311-3031303002020332-3001033202332322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.netapp_trident](resources--fleet--reference--group-002.md#canonical-2221120100320330-2103110113123100-1120330002302000-1201000323033013-3233130203103031-0231013330113102-1203110130230313-0232230300113023)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--fleet--reference--group-003.md#canonical-2122222203033102-2023313021121311-3131333222303102-2201111021211102-3123103323033132-3201131003113032-1000113100211333-1000323123113101)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--fleet--reference--group-003.md#canonical-1023300110000030-3322320012222203-3312211121033330-0233123000012211-0313012003121021-3303103100101231-0022013323332110-0101320001330221)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos

<a id="canonical-1003221313203203-0131000233301323-1320110301120322-2003020321330111-0302303133000000-1323012221101002-2031232212331331-2101022120221112"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_qos = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- storage_device_list.storage_devices.pure_service_orchestrator

<a id="canonical-0110313201102010-2001230300022320-3333102331213003-3202322213123220-0330311102032233-2011123003100323-3123312222132201-0102023231311000"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for Pure Storage Service Orchestrator.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120020022132210-0222323003323203-2122103312012021-1221022211322033-2220220002313101-3212230032120021-0211101022103112-3210222330321010"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator`

- [arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023): complete subsection reference.

<a id="canonical-1331110033133223-2300211312201210-1332001021233302-0122323110132320-3101013322002012-1001120102301233-1120123331333013-3302111032101313"></a>

<a id="canonical-1223303120222033-0202210111100013-3031201013033220-3123233331233210-3023302310031211-2320231010100200-2130312023321213-2301210020113022"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.cluster_id` property

Type: `"string"`. Optional.

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and
underscores.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 22),
}
```

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

<a id="canonical-1111210320012020-1033020233321323-1131111131221231-3131112033331312-1321322221333323-2033021232211320-1113223021230322-3233011032033031"></a>

<a id="canonical-0213320212330103-2232313300112313-1023220131122332-3013201020011233-0201031311011301-3033031102010121-2021303323012110-1103110001000102"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.enable_storage_topology` property

Type: `"bool"`. Optional.

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

<a id="canonical-1113301010311301-2202113322220212-1311203220223202-0212122320223120-0202033100323200-2111011232311113-0312022312033303-3333311032203332"></a>

<a id="canonical-0110101330221203-1120221022203030-0032101200322232-3213113210200222-3221000332312211-0110312120100013-3322100303102322-3130201212101120"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.enable_strict_topology` property

Type: `"bool"`. Optional.

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

<a id="canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="canonical-1222201323200223-0321321103232112-2101102012212110-2230211100212302-3120201131012133-0122123320122131-1120031330230223-2222131333102002"></a>

Type: `"object"`. single nested block, Optional.

Arrays Configuration. Device configuration for PSO Arrays.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
arrays {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003221210123322-0121121000303123-3011031031323032-1320011313001030-2201332003203320-2011011301011132-2201130232010012-1112220100221120"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays`

- [flash_array](resources--fleet--reference--group-003.md#canonical-3310212110120233-0203102010300103-0010123120220030-2230322323333131-1220032330101310-0310031202323012-0303210003020210-2012233111320021): complete subsection reference.

- [flash_blade](resources--fleet--reference--group-003.md#canonical-0310302300113020-1101003202302211-3132311023131201-1202333212121221-2103200023032132-0302011312131112-3231020221301103-1202323212022301): complete subsection reference.

<a id="canonical-3310212110120233-0203102010300103-0010123120220030-2230322323333131-1220032330101310-0310031202323012-0303210003020210-2012233111320021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="canonical-1021111023131203-1303203112233301-2002220300323203-0300230103101022-2012023323222203-0330222323321201-3101203223223121-0010130302200230"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash arrays should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("default_fs_type",
    "flash_arrays",
    "iscsi_login_timeout",
    "san_type")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
flash_array {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302210311130210-3231023113031110-0013331210201010-3002231302210333-1300120210113221-0201031110223231-0202122320202101-3321132132112320"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array`

<a id="canonical-1021210010230023-3110332022313103-0330002032030330-1201301332120303-2101103122320113-1211120320100121-3322012200230220-3103203231312113"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_opt` property

Type: `"string"`. Optional.

Block volume default mkfs OPTIONS. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3021101210023032-0030112230331303-0101031312121010-0200120300211101-2102212101323000-1102303213301030-3333013033010303-0031200001013301"></a>

<a id="canonical-0132232123122311-1010120202112200-0110300102313330-1221232100210221-2032211103200003-1011231212331202-3100021101103111-1322000033101112"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_fs_type` property

Type: `"string"`. Optional.

\[Enum: xfs|ext4\] Block volume default filesystem type. Not recommended to change!. Possible values
are \`xfs\`, \`ext4\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("xfs",
    "ext4"),
}
```

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

<a id="canonical-3031023332200232-0110013010310002-0002113023012112-3021213302323001-1303203102112121-1110331310112023-1011101303313303-1020211113210020"></a>

<a id="canonical-0301033300030303-0333132321322032-1332010131301222-2321022303111021-2010121002211310-3120010332202230-0232202302033333-0321012313301110"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.default_mount_opts` property

Type: `["list", "string"]`. Optional.

Block volume default filesystem mount OPTIONS. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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

<a id="canonical-1132301300033232-0001123222301013-0311313132231011-3210112211231032-2211122112231121-3022231201033220-0111133200112021-2120130103132323"></a>

<a id="canonical-3300122003100202-1123112123030331-1120000103303213-0022011221230003-3333313233002312-2130102233011100-1010120033120331-1322331223223030"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.disable_preempt_attachments` property

Type: `"bool"`. Optional.

Disable Preempt Attachments. Enable/Disable attachment preemption!

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [flash_arrays](resources--fleet--reference--group-003.md#canonical-3122132333220122-1300031212021013-0110322313232102-0002113131220310-0223220110220000-1200021213121031-3203202320330330-1121332311122133): complete subsection reference.

<a id="canonical-2321312103123003-2121222313033232-3021132110212020-2213112330120232-2112120020010112-0102130303323012-0110312333121031-3110033103012103"></a>

<a id="canonical-0131120200332320-0311321303001323-0310310131113111-3031013030222000-0102000121221213-2011223120120003-1223101230002102-1012213323213010"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.iscsi_login_timeout` property

Type: `"number"`. Optional.

ISCSI login timeout in seconds. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

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

<a id="canonical-3113010312330010-1223310023332132-3301111002031231-0222212331222222-2231013120010021-0230020020333013-3220330123203220-3232022203031133"></a>

<a id="canonical-0220003212020012-3230101330121030-2323313031102203-2231210233222001-1002111120311000-2010301233100303-0002121031032101-1310303323030313"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.san_type` property

Type: `"string"`. Optional.

\[Enum: ISCSI|FC\] Block volume access protocol, either ISCSI or FC. Possible values are \`ISCSI\`,
\`FC\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ISCSI",
    "FC"),
}
```

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

<a id="canonical-3122132333220122-1300031212021013-0110322313232102-0002113131220310-0223220110220000-1200021213121031-3203202320330330-1121332311122133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-003.md#canonical-3310212110120233-0203102010300103-0010123120220030-2230322323333131-1220032330101310-0310031202323012-0303210003020210-2012233111320021)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays

<a id="canonical-0222320010220211-2123002013303203-3333301122231310-0230300003321001-0202302201110212-2133200332011323-2133311022303221-3103021322023333"></a>

Type: `"object"`. list nested block, Optional.

For FlashArrays you must set the 'mgmt\_endpoint' and 'api\_token'.

Additional upstream details:

For FlashArrays you must set the "mgmt\_endpoint" and "api\_token"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mgmt_dns_name",
    "mgmt_ip")}
```

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

Terraform syntax:

```terraform
flash_arrays {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222332203230321-1200233310111203-0221333111031000-0103230230110030-3320113221020102-3131322323201003-2321230013022101-1122201222112231"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays`

- [api_token](resources--fleet--reference--group-003.md#canonical-0210002132032322-3322330030230331-3102011013110000-1233123000022301-3003003333130031-3321311033333321-1021230330211331-0020013012210203): complete subsection reference.

<a id="canonical-0122012211131120-3020111130200102-1213203032300133-0331122103010303-1223122320130133-0312202213031013-2320132320012002-3031211203120012"></a>

<a id="canonical-1022123122201011-0311312103003133-2221112322212220-3330212031330110-3332103313232021-0313212001130011-2100300121311120-3200313203330311"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.labels` property

Type: `["map", "string"]`. Optional.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Additional upstream details:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

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

<a id="canonical-1003210222212020-0102012303200303-2202312332300003-3331202221111232-3012003023001333-0112310332221303-3213300200003213-3220323013210030"></a>

<a id="canonical-3010333101303311-2122302032013310-0101101112311030-1033323033010210-0330120113212133-0330130002301010-1310032223213232-0133030032322221"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_dns_name` property

Type: `"string"`. Optional.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-2331110210223122-2033232302321203-0123130303121203-3033201113013303-1330103011131011-1232000023212220-2222122313303312-0321220301303003"></a>

<a id="canonical-3303220303130130-2201131320033023-0100133331222301-3021113032310021-0230333122031012-2113311013022030-1000003013002121-0110230001212121"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.mgmt_ip` property

Type: `"string"`. Optional.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-0210002132032322-3322330030230331-3102011013110000-1233123000022301-3003003333130031-3321311033333321-1021230330211331-0020013012210203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-003.md#canonical-3310212110120233-0203102010300103-0010123120220030-2230322323333131-1220032330101310-0310031202323012-0303210003020210-2012233111320021)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-003.md#canonical-3122132333220122-1300031212021013-0110322313232102-0002113131220310-0223220110220000-1200021213121031-3203202320330330-1121332311122133)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token

<a id="canonical-2020310001212022-3022231101313223-0021030012120323-0302231322202311-0200011201322302-1030200300011003-2301112303023011-2222020120131031"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000233131221221-1121230223203021-0312211021033202-2103213221202110-2020012202331100-3311110300101332-3230031011321323-3133133032233011"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token`

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-3322310000203022-3222000303221130-3312130212322301-2010232221331200-0123032312123002-1201201201101120-1302101123332032-0222332301231223): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-0222101001121303-1010300003111121-1223113103201010-1232323311132110-2121213330013321-2332020331213101-2110032333121110-0000310300321323): complete subsection reference.

<a id="canonical-3322310000203022-3222000303221130-3312130212322301-2010232221331200-0123032312123002-1201201201101120-1302101123332032-0222332301231223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-003.md#canonical-3310212110120233-0203102010300103-0010123120220030-2230322323333131-1220032330101310-0310031202323012-0303210003020210-2012233111320021)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-003.md#canonical-3122132333220122-1300031212021013-0110322313232102-0002113131220310-0223220110220000-1200021213121031-3203202320330330-1121332311122133)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--fleet--reference--group-003.md#canonical-0210002132032322-3322330030230331-3102011013110000-1233123000022301-3003003333130031-3321311033333321-1021230330211331-0020013012210203)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="canonical-1213302211323001-0300033000030021-2313321101022102-3023100122201031-0213010332101123-0013320013222202-1021030031223233-3201100033233203"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010231120310233-0023310332020230-0120212003020000-1122022311032111-1122130101031132-1010201010231330-2333010202322300-2130001220100323"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info`

<a id="canonical-3010123102230001-2311312122221132-3330323312202313-1213333122103111-0303020031330312-0013221011310301-3232302103033233-2103011110222012"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-0003303203223300-0111032331313023-1111311322010212-2110032033213231-0030222333022321-3210222101102221-0010231323013001-0333130001121203"></a>

<a id="canonical-0312000302102320-1112203120322123-3022113332130111-2030231211130301-0333112110322312-2321121013331203-3023210020233211-1132330330213022"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-2022102110211320-1003013032202230-2103231311102331-3302331210132020-3331332321000330-0222331130020221-0203230303332331-1113123003013023"></a>

<a id="canonical-1000003220313320-0230321011100022-3131330122222000-2121120122003033-0101321301223103-0212110233131001-3202023203312331-1102021210323112"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-0222101001121303-1010300003111121-1223113103201010-1232323311132110-2121213330013321-2332020331213101-2110032333121110-0000310300321323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--fleet--reference--group-003.md#canonical-3310212110120233-0203102010300103-0010123120220030-2230322323333131-1220032330101310-0310031202323012-0303210003020210-2012233111320021)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--fleet--reference--group-003.md#canonical-3122132333220122-1300031212021013-0110322313232102-0002113131220310-0223220110220000-1200021213121031-3203202320330330-1121332311122133)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--fleet--reference--group-003.md#canonical-0210002132032322-3322330030230331-3102011013110000-1233123000022301-3003003333130031-3321311033333321-1021230330211331-0020013012210203)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info

<a id="canonical-0302301301220231-2002130230331230-0332230330203230-1311211111121333-1202101300202213-0210012130101120-3012202320123011-1322033003302212"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123330321300003-1000023233100311-3110213112313023-2312120303121112-0310220232201313-1313302321013313-1103220010120002-1323223332101123"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info`

<a id="canonical-1200311110032231-3332023103310203-2220030130233111-2032301203230010-0312331031112031-2030100133230210-2331032020110012-2003031210021020"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2132302021230031-3103130211103231-2033111200331322-2200100221130211-2333320013023011-1221311033220011-2332010310311111-3122130201001100"></a>

<a id="canonical-1210012210012112-1201301103002030-0331031100123220-2020112130210302-3323012003323223-0331302202133202-3102303101233231-0113131313311012"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-0310302300113020-1101003202302211-3132311023131201-1202333212121221-2103200023032132-0302011312131112-3231020221301103-1202323212022301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="canonical-0023122320012230-0330131303220123-3103002103331211-2133231322302301-3110121323320011-2332132202123020-3033300223100202-1333322310000301"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash blades should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("flash_blades")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
flash_blade {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203032232311322-0201132220211130-3223333011321022-1200310112233021-0010011012302031-0021321332130330-0020203232320021-0010223221002111"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade`

<a id="canonical-0013320101021123-0320312211220000-1132302220011111-2121032130100310-1203010322101233-2231010202303131-1013113300001232-3101212002120211"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.enable_snapshot_directory` property

Type: `"bool"`. Optional.

Enable Snapshot Directory. Enable/Disable FlashBlade snapshots.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2020221110230312-3020310023221101-3121230311100110-0221130201222213-1202312132331320-2222321123101130-1032332123023330-2212200101130303"></a>

<a id="canonical-1332132100223021-1012123121011123-3210012020231233-3310320210012123-3321120120201300-3023213110003302-0023013103133133-3322012201123232"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.export_rules` property

Type: `"string"`. Optional.

NFS Export Rules. NFS Export rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 250),
}
```

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

- [flash_blades](resources--fleet--reference--group-003.md#canonical-2223232302311022-1021012202203331-2003121103333303-2102321103222121-1101000130232101-3232123232311220-0233312220020222-0221302333033323): complete subsection reference.

<a id="canonical-2223232302311022-1021012202203331-2003121103333303-2102321103222121-1101000130232101-3232123232311220-0233312220020222-0221302333033323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-003.md#canonical-0310302300113020-1101003202302211-3132311023131201-1202333212121221-2103200023032132-0302011312131112-3231020221301103-1202323212022301)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="canonical-3130013232110200-0123201330102301-1123020122131022-1223233130012333-3001323030121323-0333210213131330-2220112201323123-2331111121123001"></a>

Type: `"object"`. list nested block, Optional.

For FlashBlades you must set the 'mgmt\_endpoint', 'api\_token' and nfs\_endpoint.

Additional upstream details:

For FlashBlades you must set the "mgmt\_endpoint", "api\_token" and nfs\_endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mgmt_dns_name",
    "mgmt_ip"),
  validators.ConflictingListObjectAttributes("nfs_endpoint_dns_name",
    "nfs_endpoint_ip")}
```

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

Terraform syntax:

```terraform
flash_blades {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033022023333003-2330311232011313-0213123233013101-2120111120202012-3020203300121322-2030310203012303-3233200010013023-1033231213230003"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades`

- [api_token](resources--fleet--reference--group-003.md#canonical-3003301101020100-0233130231113013-2213021211303012-0220010031202032-3021210300123032-3310332022203131-1210033032123322-2210032321320100): complete subsection reference.

<a id="canonical-2320012023001332-3021021123023103-2202201031100230-0302332101211000-3023213022300120-3003320133022221-3122300000203323-0113322213023003"></a>

<a id="canonical-2023000010211200-1230203102113300-0330103320120231-0133330303211030-2033130130311331-3020011001022223-3313001232012023-2203330213321321"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.labels` property

Type: `["map", "string"]`. Optional.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Additional upstream details:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

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

<a id="canonical-1301333311311000-0120022021123202-0002231233013310-1232332011312023-2203203202202203-1203313032323111-0231322303231133-2302322230023102"></a>

<a id="canonical-1000131132321201-2301211111203023-2130000113022013-2112213011313203-2011212112103202-2203122003333032-1100311111221020-3010003111103031"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_dns_name` property

Type: `"string"`. Optional.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-1130222232301223-1010210003211211-3230203130201333-2331022332033211-0313113313332110-2111113322213012-3013110103112320-3320131310322112"></a>

<a id="canonical-0333110210032332-3133201112033332-2322212213221223-1013133212101330-3131231133313133-1033202213013002-1123222323302221-1323312010333213"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.mgmt_ip` property

Type: `"string"`. Optional.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-2110312131030113-1301300012011101-3023122011133012-3020133220032212-2011303112132012-1003103322230323-1202303303032311-1031232013023130"></a>

<a id="canonical-0110001022003002-3112031113103112-0213210310310033-0031121103010233-1302012230032033-3131223123130212-3220323110211020-0233221113120333"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_dns_name` property

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-3220011001133032-0323121013222213-2003312300331232-3312003120132001-3103133330001221-2002202312311222-0120001323213232-3233100030311032"></a>

<a id="canonical-1322212102123110-0133320312000202-2103110122232312-2303122100002210-1332221333331203-3230022331330222-0212212310121131-3101111202013313"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.nfs_endpoint_ip` property

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-3003301101020100-0233130231113013-2213021211303012-0220010031202032-3021210300123032-3310332022203131-1210033032123322-2210032321320100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-003.md#canonical-0310302300113020-1101003202302211-3132311023131201-1202333212121221-2103200023032132-0302011312131112-3231020221301103-1202323212022301)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-003.md#canonical-2223232302311022-1021012202203331-2003121103333303-2102321103222121-1101000130232101-3232123232311220-0233312220020222-0221302333033323)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="canonical-1003112101311133-0022112320333102-0111000201022221-0011303200023232-1033203203300013-2112233032322202-1311100303101011-2322111330133322"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201212102313001-0103003003300220-0123120223231323-3220201112021010-0101011133201221-3302022120001021-2020123333131222-0310032001231120"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token`

- [blindfold_secret_info](resources--fleet--reference--group-003.md#canonical-2302333121320233-3000130020023011-2300113131201320-1322231231001331-3011022301032010-0012332233203013-1030320121002130-1101002203003303): complete subsection reference.

- [clear_secret_info](resources--fleet--reference--group-003.md#canonical-2010011112000302-1232221132002002-3312110330233223-3121331203120003-3011332101312002-0102013302132223-3011003000111313-3112033221012023): complete subsection reference.

<a id="canonical-2302333121320233-3000130020023011-2300113131201320-1322231231001331-3011022301032010-0012332233203013-1030320121002130-1101002203003303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-003.md#canonical-0310302300113020-1101003202302211-3132311023131201-1202333212121221-2103200023032132-0302011312131112-3231020221301103-1202323212022301)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-003.md#canonical-2223232302311022-1021012202203331-2003121103333303-2102321103222121-1101000130232101-3232123232311220-0233312220020222-0221302333033323)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--fleet--reference--group-003.md#canonical-3003301101020100-0233130231113013-2213021211303012-0220010031202032-3021210300123032-3310332022203131-1210033032123322-2210032321320100)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info

<a id="canonical-0211300102302120-0302002001203300-3210221310310133-2003003030310223-1031330020233101-1210130203313200-3102110100222210-2133300230011003"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010211310111110-3001310322322223-2011323312222222-0101122333203303-0100222301231233-0212021102302121-0303102122230023-1020121002233313"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info`

<a id="canonical-3333112032131022-3022313202022220-1133033323221303-2233332232233032-3101030233023012-2101113013001222-0021011223312320-0011002223210323"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-3320101031133313-1111110203102032-2103302011111302-1003300211310030-1100211222003311-2223303110220212-0013322330200010-3221110212310031"></a>

<a id="canonical-3202122111121033-2303132122101122-3021222321301133-2101302233213132-0301200231312223-0322112333033310-3010101012020032-2301330131010003"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1110303102000301-3023033332010012-2001302111122221-2130202301011230-0103213322221211-2313333121331212-3113223033300100-2111313030333332"></a>

<a id="canonical-3313022112232301-2310223232310203-3233030221001111-3201210112211223-1312010323232012-0010301212030003-3003130103320032-1020301233223221"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-2010011112000302-1232221132002002-3312110330233223-3121331203120003-3011332101312002-0102013302132223-3011003000111313-3112033221012023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_device_list](resources--fleet--reference--group-002.md#canonical-2120213230221113-1333030012101112-2301222000203003-0123012120201213-1323232311211033-3022303032122220-1133001002321021-1101312223021101)
- [storage_device_list.storage_devices](resources--fleet--reference--group-002.md#canonical-1121310030023210-2313133122303233-3230332100212103-3010021230232322-3223210233032123-3200023200220121-3112220232033121-3020030131031221)
- [storage_device_list.storage_devices.pure_service_orchestrator](resources--fleet--reference--group-003.md#canonical-0002230120203323-3100203133101203-1033110303010030-1013320110112131-2330233023330013-1213113110311020-3003031003233122-3102232221001010)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--fleet--reference--group-003.md#canonical-1111203111312030-1320103010112000-2022221321200323-0223011123121012-3121030110021223-2133102323103132-2023030013223201-3320023100132023)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--fleet--reference--group-003.md#canonical-0310302300113020-1101003202302211-3132311023131201-1202333212121221-2103200023032132-0302011312131112-3231020221301103-1202323212022301)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--fleet--reference--group-003.md#canonical-2223232302311022-1021012202203331-2003121103333303-2102321103222121-1101000130232101-3232123232311220-0233312220020222-0221302333033323)
- [storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--fleet--reference--group-003.md#canonical-3003301101020100-0233130231113013-2213021211303012-0220010031202032-3021210300123032-3310332022203131-1210033032123322-2210032321320100)
- storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info

<a id="canonical-2013312123332222-1003010032313113-0303111133112000-3023211131121323-0211122130000011-3000021213211133-3320121202202112-2010302121321132"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221302303113303-3121023013110023-3330233303232032-2312202002112013-2200103331300132-0001032020023220-0110011033013211-3322113211232001"></a>

### Direct properties for `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info`

<a id="canonical-3012123101310132-0300310112132131-0130312102211220-0323101333032102-3133000201132230-0101010212320010-1131202333332312-1000332303030022"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1030312111123001-1310310332331001-0322101122103310-2003003121332300-2331210231003013-1231110312233031-1310212312212110-0111001303032122"></a>

<a id="canonical-2302020212332231-1131220300001002-3131313323222122-3013201301033011-2332100231232301-2203111231213223-2222311311303311-3030333311133332"></a>

#### `storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-2200223121321313-3202113301130230-3113322333033203-2122021133112132-3302320001211330-3311111102111033-2201322232210330-0102122100010132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_interface_list` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- storage_interface_list

<a id="canonical-2223313213113003-2010202220123320-3322302112013032-3001322211322310-2322330013330331-1010013302213030-2221310300112203-1322231330200313"></a>

Type: `"object"`. single nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
storage_interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103203032002231-2323200320300100-2120201221101230-3200311013311301-3032113003320212-0123030202321021-1012212301301112-0233221022020113"></a>

### Direct properties for `storage_interface_list`

- [interfaces](resources--fleet--reference--group-003.md#canonical-0210310131121313-1301013332313200-0102312022133012-3211101302122202-1132201232000013-2200303202131312-1233103030001133-0212033313312133): complete subsection reference.

<a id="canonical-0210310131121313-1301013332313200-0102312022133012-3211101302122202-1132201232000013-2200303202131312-1233103030001133-0212033313312133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_interface_list.interfaces` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_interface_list](resources--fleet--reference--group-003.md#canonical-2200223121321313-3202113301130230-3113322333033203-2122021133112132-3302320001211330-3311111102111033-2201322232210330-0102122100010132)
- storage_interface_list.interfaces

<a id="canonical-3222302313101310-0000213310101302-1333001013221232-1000321010203200-2230333031200013-3320013322312231-0031030021031210-3131233203313131"></a>

Type: `"object"`. list nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111013131232010-2200032303120332-1222212110113102-2102103003033221-1302112312110302-0213030310332211-3033002301323001-0010200132222003"></a>

### Direct properties for `storage_interface_list.interfaces`

<a id="canonical-0001010310110001-1200131311112010-2231313101122232-0322212100003322-3123320023203123-3301022220130213-3033322311320202-3102232200131213"></a>

#### `storage_interface_list.interfaces.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3331312331220012-3303310000313201-0003322220301233-1213002222333210-0023023023021011-2130220012020132-3323310110113311-3012111312100313"></a>

<a id="canonical-2003313302103032-3331310320203202-1032120130223130-1211233301023111-3132202101002212-3112230010220022-1203311323011102-3231000013200133"></a>

#### `storage_interface_list.interfaces.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1112313210200113-0201231203031213-0311323110130011-1311330212123223-2023110213000013-3222023102113131-3022312230222133-0003220101330101"></a>

<a id="canonical-0221000031103010-3032112120322023-3120030302300312-2023121101201111-3020202013200021-3332333311220230-1013131021330022-1110320000113200"></a>

#### `storage_interface_list.interfaces.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- storage_static_routes

<a id="canonical-3300220221311210-2223123102101211-3220330103110130-2013011230213113-0302003112110032-0211100233331131-0000213033313103-2010020102112211"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for storage static routes.

Additional upstream details:

List of storage static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_routes")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
storage_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313032300003323-1200212133312312-0212122021211232-1303300232332111-2112103212022002-0103033330303011-2310121212100231-3130322003330220"></a>

### Direct properties for `storage_static_routes`

- [storage_routes](resources--fleet--reference--group-003.md#canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231): complete subsection reference.

<a id="canonical-2103110110001123-1000122230223022-2202332021231220-0033011003313113-1322332012002001-0120330220321111-3021132100100332-2230130122031231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `storage_static_routes.storage_routes` properties

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_static_routes](resources--fleet--reference--group-003.md#canonical-3310103001011013-0133232333013322-2230330101201311-3021013203230121-1032321030203020-3003121131331312-2101312003321212-2012211010032330)
- storage_static_routes.storage_routes

<a id="canonical-1113103203120230-0103213211302232-2101121302323232-1122111212313231-2311322311200031-1121301130131233-1130133302322133-3311302231122210"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of storage static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("subnets")}
```

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

Terraform syntax:

```terraform
storage_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220311232321311-2101323223330230-1103013003202300-2110122210120013-0010021330313201-3023010223031310-3102321301330120-0123212122123231"></a>

### Direct properties for `storage_static_routes.storage_routes`

<a id="canonical-2033130301313321-1232220332313121-3311112212031300-2211323031330223-1200021003201131-3310133300223310-2120201103231320-3022031112310003"></a>

#### `storage_static_routes.storage_routes.attrs` property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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

- [labels](resources--fleet--reference--group-004.md#canonical-3122130211122322-1301231021112223-1012121131110203-0212120302301103-0321330301322030-1322303112321312-2232110100220131-3212330121123110): complete subsection reference.

- [nexthop](resources--fleet--reference--group-004.md#canonical-2323102010131331-0222022123303131-0001320120102230-3002231130201221-0320012123110310-2111211303213332-2101222302031312-1321301311201213): complete subsection reference.

- [subnets](resources--fleet--reference--group-004.md#canonical-0011031323102333-2133013032021320-3023030012331100-2101022003022112-3302133213002111-2211232221112303-0333322213031131-1131101123022222): complete subsection reference.
