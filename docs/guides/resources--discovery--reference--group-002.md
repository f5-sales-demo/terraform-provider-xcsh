---
page_title: "xcsh_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery reference."
---

# xcsh_discovery reference

<a id="canonical-3332310112212022-3122012220002330-2121220211123313-1033230023102333-3330312121222111-2032023323130013-0003110121320123-2320322123230011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.isolated` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- discovery_k8s.access_info.isolated

<a id="canonical-3211300221101310-3111000233302102-0003303033023023-0013221003331123-1010310010222310-2323120000222312-2003211221021101-3103201311333023"></a>

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
isolated = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.kubeconfig_url` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- discovery_k8s.access_info.kubeconfig_url

<a id="canonical-1033220201210013-0311033123222002-3212333323131300-0213032312230030-0101322010020222-0302103230121203-2033313012320120-3023033010100033"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
kubeconfig_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103200102313011-2011113110132231-2032022300133100-3211020220310133-1110330313310210-3322130202103322-2002000323212203-3202211000100120"></a>

### Direct properties for `discovery_k8s.access_info.kubeconfig_url`

- [blindfold_secret_info](resources--discovery--reference--group-002.md#canonical-1121112003223220-2111211201312211-3313302223212113-0312231220131131-0222122203020203-0231232101302010-1303022030200300-3333302311300130): complete subsection reference.

- [clear_secret_info](resources--discovery--reference--group-002.md#canonical-0100213213212331-2003313330013022-3230232303212010-2201001131100021-2221103310311031-2131103302233031-0300232300023101-0000331310202031): complete subsection reference.

<a id="canonical-1121112003223220-2111211201312211-3313302223212113-0312231220131131-0222122203020203-0231232101302010-1303022030200300-3333302311300130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-002.md#canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130)
- discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info

<a id="canonical-3102230331112221-2100230112012011-1213331102102001-1033013222121133-1321300232220200-3103211011110120-3201221221313211-1102010012101102"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0223313122102332-2220222100300321-0310201332003033-3200001110233122-2212030000221112-0332331320333133-2333221102002023-0332031331133033"></a>

### Direct properties for `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info`

<a id="canonical-1103303221030211-2100031021202022-1103001330010000-1112000113220022-1131001210023200-3201010100303000-1012131100132131-2113231232320102"></a>

#### `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2032131212111000-2113001030322321-3103321103310131-3132300103123211-1000103021322300-3212300331122232-1110320333030020-3110313212203301"></a>

<a id="canonical-1130003121111100-3100332220131331-1231110033123013-2310333331101003-2020002022302003-3322321103123113-1130102233033030-2233012102100113"></a>

#### `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0300312312030123-0300210331211232-0102212022211232-0330031300100231-2331101300312113-1101112230112111-3332110333331023-1301032122010120"></a>

<a id="canonical-3313031010321223-2021200223121131-0002132232032101-1123213033012221-1233321033003133-0001000000322203-2120212111100011-1320321121313012"></a>

#### `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0100213213212331-2003313330013022-3230232303212010-2201001131100021-2221103310311031-2131103302233031-0300232300023101-0000331310202031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.kubeconfig_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- [discovery_k8s.access_info.kubeconfig_url](resources--discovery--reference--group-002.md#canonical-2103233030101201-2311203221133210-0220202230223032-3211312010223020-0101303231022121-2330110001021033-1320230101103202-0211001030103130)
- discovery_k8s.access_info.kubeconfig_url.clear_secret_info

<a id="canonical-0211013121120001-1030303311113122-0201210030321322-1123032022123210-0330030221311111-0321211123331033-1302021022031320-2133101012221301"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1331220233032312-3001310213200102-1010120300022331-2211331022321211-2122201022301111-1032311230302221-3132202122222011-0211000202002212"></a>

### Direct properties for `discovery_k8s.access_info.kubeconfig_url.clear_secret_info`

<a id="canonical-2321131111120312-3223313210213103-3202010223011233-3322230223003033-0213110012232220-0333301230031312-3101233002100232-2230232033203233"></a>

#### `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2301313313331230-3323110211231223-1221121122111023-2201111121121103-1002122012011313-2013200302123112-1302330001103103-1213111230320231"></a>

<a id="canonical-2123203033030022-3310021011321002-1321100120211111-0301220110023331-1031203010300010-2201231020301013-1220032303111330-3310210030132020"></a>

#### `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2233200231201012-1110023212011100-0220223121300122-2221312123303311-3102232213022203-2332332222223031-3301300320131201-2133101011010213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.reachable` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.access_info](resources--discovery--reference--group-001.md#canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320)
- discovery_k8s.access_info.reachable

<a id="canonical-2112131220231232-0301102320201020-1232302112130232-1133221121220011-2330130313333122-3230001102122021-1022033201022213-1332332002220132"></a>

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
reachable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112000203012301-3300021232321103-2332023112322323-2303202002010111-0021011211110313-0120122221203123-0121321110201012-0332133223123202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.default_all` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- discovery_k8s.default_all

<a id="canonical-2033200123113032-3001322221200211-2100033312102313-0312103132132013-1011101130212313-0001000000233301-2110230201220220-1320132000312113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default all.

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
default_all = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321010022030131-3311200301332101-3132022310131230-1233220021130131-3210333300201303-3321221200212203-2130012322120030-1013021110331232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.namespace_mapping` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- discovery_k8s.namespace_mapping

<a id="canonical-1130310213023212-1332013022102022-1111331100303002-2021113202231020-1230020021122231-3322220200003000-2322230130000120-0232111221000111"></a>

Type: `"object"`. single nested block, Optional.

Select the mapping between K8s namespaces from which services will be discovered and App Namespace
to which the discovered services will be shared.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("items")}
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
namespace_mapping {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002130212032223-1332132121001013-2320333022003012-0012103010313020-1200032232303001-2302031032022133-2122020030313211-3231310223230231"></a>

### Direct properties for `discovery_k8s.namespace_mapping`

- [items](resources--discovery--reference--group-002.md#canonical-3002302011033021-2032112031301321-3113010021213321-1211130323222133-0222322210032313-2033311001220112-2322102303032012-1211323103133123): complete subsection reference.

<a id="canonical-3002302011033021-2032112031301321-3113010021213321-1211130323222133-0222322210032313-2033311001220112-2322102303032012-1211323103133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.namespace_mapping.items` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.namespace_mapping](resources--discovery--reference--group-002.md#canonical-0321010022030131-3311200301332101-3132022310131230-1233220021130131-3210333300201303-3321221200212203-2130012322120030-1013021110331232)
- discovery_k8s.namespace_mapping.items

<a id="canonical-3112000132011301-1222010321120031-3203310201120101-2133032230221032-1210122030302003-2113022110300221-2203221203330020-3021321032100331"></a>

Type: `"object"`. list nested block, Optional.

Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be
mapped to a single App Namespace, which is determined by the first matched regular expression.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
items {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110013112112131-1032022130113303-3233202111113110-0123231002331032-0113133020211003-1121312132103323-2030012230300333-0003103230122312"></a>

### Direct properties for `discovery_k8s.namespace_mapping.items`

<a id="canonical-3201112330030203-1133302310200330-1102230113001203-3002032231013323-3222112223222010-1031010002130102-3301331110032301-2122202331201332"></a>

#### `discovery_k8s.namespace_mapping.items.namespace` property

Type: `"string"`. Optional, Computed.

F5XC Application Namespaces. Select a namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1130100022032112-2301113201110213-3223002220033003-0303132103013200-0001103312021020-0221103001202012-2132111201101330-3200113123101210"></a>

<a id="canonical-2131332121223110-1012322223313121-1210111110010013-2013330011302311-0210102133221333-3113032211002200-1110103020112331-2121332023113322"></a>

#### `discovery_k8s.namespace_mapping.items.namespace_regex` property

Type: `"string"`. Optional.

The regular expression here will be used to match K8s namespace(s).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- discovery_k8s.publish_info

<a id="canonical-3122222231121203-3012003210123330-0330311123322320-2213313323121323-0113123300203332-1113032311231020-0323112032313021-3223031122220121"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for publish info.

Additional upstream details:

K8s Configuration to publish VIPs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "dns_delegation"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("publish",
    "publish_fqdns")}
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
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"dns_delegation\",\"publish\",\"publish_fqdns\"]"
}
```

Terraform syntax:

```terraform
publish_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230120202310122-1100012123022121-3201101312103310-3230300302002012-1101102311031120-1310332211313022-2231322302301022-3030212132331112"></a>

### Direct properties for `discovery_k8s.publish_info`

- [disable_spec](resources--discovery--reference--group-002.md#canonical-3011020302001133-2202023030221333-1231203031313222-1323230201300200-0033331312113001-1023022000002023-1113301030132032-0031131021301222): complete subsection reference.

- [dns_delegation](resources--discovery--reference--group-002.md#canonical-0213033233320133-2313331011023030-3300221322303212-3012012303031332-2322003103212311-2232013223120231-0223002211302120-2121012230012320): complete subsection reference.

- [publish](resources--discovery--reference--group-002.md#canonical-2201132213231022-0230303111210031-0130220302231300-2323030120010232-2210223132002011-2002001223211111-3210202213010231-0112131333302012): complete subsection reference.

- [publish_fqdns](resources--discovery--reference--group-002.md#canonical-2320301002300300-0102133000123033-1212202333001332-0013302020101020-3320231121013113-0232002100202122-2121011332331332-2121202110113103): complete subsection reference.

<a id="canonical-3011020302001133-2202023030221333-1231203031313222-1323230201300200-0033331312113001-1023022000002023-1113301030132032-0031131021301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info.disable_spec` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.publish_info](resources--discovery--reference--group-002.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- discovery_k8s.publish_info.disable_spec

<a id="canonical-1131112011113030-1210133212003130-3311313320312101-1010030313001132-3013331233013300-1312020132001211-2322023011103123-0311120302310100"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213033233320133-2313331011023030-3300221322303212-3012012303031332-2322003103212311-2232013223120231-0223002211302120-2121012230012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info.dns_delegation` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.publish_info](resources--discovery--reference--group-002.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- discovery_k8s.publish_info.dns_delegation

<a id="canonical-0013022331002220-1313110232311321-3210232322133213-0021213003322120-1312002013122030-3330223222122301-3013202002112203-0303031222230100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for DNS delegation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("subdomain")}
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
dns_delegation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132203212213332-3233011123102333-3133001200031320-1223230111132131-3121002332213230-0111000031003200-0223102201300233-1112011220201333"></a>

### Direct properties for `discovery_k8s.publish_info.dns_delegation`

<a id="canonical-1332300010111012-3302320013202330-0111130212331013-0130321000210330-3310023321020232-0010333012131112-1002231112021123-2320002323102121"></a>

#### `discovery_k8s.publish_info.dns_delegation.dns_mode` property

Type: `"string"`. Optional.

\[Enum: CORE\_DNS|KUBE\_DNS\] Two modes are possible CoreDNS: Whether external K8s cluster is
running core-DNS KubeDNS: External K8s cluster is running kube-DNS. Possible values are
\`CORE\_DNS\`, \`KUBE\_DNS\`. Defaults to \`CORE\_DNS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CORE_DNS","KUBE_DNS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CORE_DNS",
    "KUBE_DNS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CORE_DNS",
  "enum": [
    "CORE_DNS",
    "KUBE_DNS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022211331330030-1332003123210002-3333231111322323-3323233032003011-2333223123033210-2202100211311321-3310101233321103-2333210311000310"></a>

<a id="canonical-3001232222022321-2011332132210212-3303111021103222-2331302212012001-3130223302311111-3021233130312132-0311203022032000-1010201300300132"></a>

#### `discovery_k8s.publish_info.dns_delegation.subdomain` property

Type: `"string"`. Optional.

The DNS subdomain for which F5XC will respond to DNS queries.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2201132213231022-0230303111210031-0130220302231300-2323030120010232-2210223132002011-2002001223211111-3210202213010231-0112131333302012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info.publish` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.publish_info](resources--discovery--reference--group-002.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- discovery_k8s.publish_info.publish

<a id="canonical-1313310112011001-2301122021130021-2233122213001303-3132203303203013-0132212011210030-1110203023102312-2032012331000200-0031120213233032"></a>

Type: `"object"`. single nested block, Optional.

K8SPublishType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("namespace")}
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
publish {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212002102030132-0031103231130313-0103302300010333-3313322032112201-0321010231032101-1230000100320222-2010132010122100-2122303222012111"></a>

### Direct properties for `discovery_k8s.publish_info.publish`

<a id="canonical-2301110233320203-0021313232121333-1220131003100110-3221033032110200-2302200212010212-1113301212023221-0210121133120020-1213010011022120"></a>

#### `discovery_k8s.publish_info.publish.namespace` property

Type: `"string"`. Optional, Computed.

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2320301002300300-0102133000123033-1212202333001332-0013302020101020-3320231121013113-0232002100202122-2121011332331332-2121202110113103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.publish_info.publish_fqdns` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [discovery_k8s](resources--discovery--reference--group-001.md#canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000)
- [discovery_k8s.publish_info](resources--discovery--reference--group-002.md#canonical-0203000021133233-3222203312230000-3212020101320013-0023011013001011-0302303200132012-2020233333033223-1131132020211023-1000310030022113)
- discovery_k8s.publish_info.publish_fqdns

<a id="canonical-1031022023310213-0232313203112320-1213302020313332-0311311033232210-2302323330201030-2113123312203013-2323321312021030-2222203323321201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for publish fqdns.

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
publish_fqdns = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332001313120022-1313323002112231-2021132023022331-0131012222332133-3231233201102020-0102032212311132-1210113121202130-1000203132133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_cluster_id` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- no_cluster_id

<a id="canonical-3002323030220322-0020200231302202-3111002300320220-3022101033030022-2021313300131030-0300131003010012-3302102103231001-3130001221002120"></a>

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
no_cluster_id = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133122212320132-3323032213033022-3031001120032311-3120311312312302-0221333302003213-0021323013310132-3013211230313033-1331213120121300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- timeouts

<a id="canonical-0012101032103333-1011032300121101-1320200130220320-3123221213101133-3033312102212232-3230131201211022-0330321023111232-2312201331323113"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332223100213303-1223222212222101-3123110331221210-1332210000203222-1201303001203212-3232130100013301-0310221201012121-1222323300131220"></a>

### Direct properties for `timeouts`

<a id="canonical-1112033131022101-2022013210222232-1032111211310023-3103132312130230-0332013311310200-1320023323331223-2123323113333110-2102223332202222"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2202202032033022-2300320013310321-3032002000003122-3311310323233233-1110211012322012-2223330110301111-0021011112302010-3303323300133200"></a>

<a id="canonical-2311232203213221-0223010331121013-1123002301311311-2203021113132110-3300202202112303-1301012220101320-2010210203111113-2022301020133330"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1011000201012302-1011300001301023-1101211332122121-0012132102332023-0132133332030322-2020132020301311-1330332300202211-1221120303031121"></a>

<a id="canonical-3222123220210310-2003133130222002-0103102112302132-1203311133301211-2020220322102312-3330022203030013-1311023020202101-3231030003000001"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3231100130200331-3321310232220101-2113300013331202-2123333311110333-0103113330100211-0002013301300101-3333023313130001-0202100110313023"></a>

<a id="canonical-1301100033110110-1202212100000300-3012120022002212-2322311033001203-1203031103101221-0330010023011133-1321301202312133-2203203130301132"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- where

<a id="canonical-1321103302233201-0332121301121100-3101133031230333-0133220331333201-0030002311302101-3220122332233201-0213131123032010-1102233222333213"></a>

Type: `"object"`. single nested block, Optional.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingObjectAttributes("virtual_network",
    "virtual_site")}
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032123011230031-2030312203122113-2302100022101232-1003222232120031-3103133201211302-0321013120201111-2302310022321010-0303111120201010"></a>

### Direct properties for `where`

- [site](resources--discovery--reference--group-002.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200): complete subsection reference.

- [virtual_network](resources--discovery--reference--group-002.md#canonical-1110120031123232-1103133033221300-2333022012020332-2112320331303032-0100001133321112-2231130320002221-2033320232220130-0031211123210313): complete subsection reference.

- [virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330): complete subsection reference.

<a id="canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- where.site

<a id="canonical-2012013100322022-1022120320103220-1213011202221212-0031301113211120-0121000101211123-3131302320112311-0011001103132122-2312321232311020"></a>

Type: `"object"`. single nested block, Optional.

This specifies a direct reference to a site configuration object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001332330111000-0210021310011302-0200010321220301-3231021112300313-1333300020102230-2321200023332123-1022101022002201-1213100103102223"></a>

### Direct properties for `where.site`

- [disable_internet_vip](resources--discovery--reference--group-002.md#canonical-1132310133301223-1201221101101031-0330233322032030-0023233223120110-0230220023201310-2333220211300312-0303123103120213-0001032331220103): complete subsection reference.

- [enable_internet_vip](resources--discovery--reference--group-002.md#canonical-2110231323102212-0020122101231302-3330133200323223-2201010003300321-2322213312332232-3102021022010321-1320030113212012-2303013132303030): complete subsection reference.

<a id="canonical-1223122302203100-3200331030230313-3033220330302303-2130131323322120-1023013321001130-2221000311303032-1010022102102223-2111232233001112"></a>

<a id="canonical-0001010210130201-0133120123311112-1231223132230233-1203201132201320-0232231031131232-2000320011310002-3011313311303211-1200223103230011"></a>

#### `where.site.network_type` property

Type: `"string"`. Optional.

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

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
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
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

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

- [ref](resources--discovery--reference--group-002.md#canonical-2032210203231300-1003003000132213-3101202330201122-0001100213311020-0131210232302123-3012213102000232-2331302321030002-1313330003132301): complete subsection reference.

<a id="canonical-1132310133301223-1201221101101031-0330233322032030-0023233223120110-0230220023201310-2333220211300312-0303123103120213-0001032331220103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.site](resources--discovery--reference--group-002.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- where.site.disable_internet_vip

<a id="canonical-2130130203303330-3032002112302222-3303311021103102-3220300023331330-2220110300021020-3130203123022331-2211123013020023-2212023112133033"></a>

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
disable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110231323102212-0020122101231302-3330133200323223-2201010003300321-2322213312332232-3102021022010321-1320030113212012-2303013132303030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.site](resources--discovery--reference--group-002.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- where.site.enable_internet_vip

<a id="canonical-2013232001120033-0302203321031212-1133020011332302-1333012221313120-2000331022223130-1101021100312211-0222032331230000-0230320102112321"></a>

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
enable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032210203231300-1003003000132213-3101202330201122-0001100213311020-0131210232302123-3012213102000232-2331302321030002-1313330003132301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.ref` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.site](resources--discovery--reference--group-002.md#canonical-2111032303022223-2300132122211213-0100203302230222-1011230233321231-1103310111023200-0013013033313032-1103302322310010-1210312011210200)
- where.site.ref

<a id="canonical-2033133220303202-3013313000222302-2301200020221123-1122313121332100-3130122201311311-0000112123213332-1303231310133313-1202031022130122"></a>

Type: `"object"`. list nested block, Optional.

Reference. A site direct reference.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002200330012003-3131013313111130-3231203222012133-0122002002320133-1210301220231323-3311120311001222-2100213101333210-3120021101121203"></a>

### Direct properties for `where.site.ref`

<a id="canonical-1113130301320303-0102201023122131-3030032223203001-1230013232303200-0011222101202102-0123132303213222-3202320010210311-3121021213212102"></a>

#### `where.site.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3031320213210302-3220122121323100-3210233133033321-1323022011022011-1200100211232201-0330332302310320-3033221000311312-1101131220220030"></a>

<a id="canonical-3032030203201003-3311221223223133-1022320331330313-1332212132002113-1131220021100010-1222112333111001-0303320021302313-0101033031030222"></a>

#### `where.site.ref.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2200112313310013-2202311201323300-3333100133332133-0002320320303203-2123202220032331-2330031322020203-2012123023330220-3100130010330302"></a>

<a id="canonical-3320210003221030-3331122330333020-1122321031001002-1031231011130103-0221301020223203-1331200022232010-3331200330210032-1032122111120101"></a>

#### `where.site.ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2330123230100021-3011022000123301-3010211231231122-2233001123111213-0111131030321332-2012031031200121-3220312312031010-3200110301112200"></a>

<a id="canonical-2232320303211301-0021221133330001-2111013030103113-1023313300221013-1122031322102221-2020302120320102-2313101300312023-1312110321233223"></a>

#### `where.site.ref.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3302231323230211-1230022212113010-2132212230313210-1112211220021102-2201112103222210-2121122200113332-1232212331001002-3133130123122231"></a>

<a id="canonical-1310333031022200-3300012323011311-0110312030013020-3310232301013210-3210233100020302-1220130132110110-3023021331023233-3022000330202020"></a>

#### `where.site.ref.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1110120031123232-1103133033221300-2333022012020332-2112320331303032-0100001133321112-2231130320002221-2033320232220130-0031211123210313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- where.virtual_network

<a id="canonical-0232303211103312-3331210310230232-2233302213331130-0103210120311211-0000201320122332-3000300231333221-2002100223333131-2311110230300022"></a>

Type: `"object"`. single nested block, Optional.

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002110222330132-2331301213110320-0200201123030012-3333111202102111-1112003211201023-2233113111232112-0023012111202121-2111331000120111"></a>

### Direct properties for `where.virtual_network`

- [ref](resources--discovery--reference--group-002.md#canonical-3312231132303023-0313321002210120-1100032032133210-2222122233111210-0001122010132331-0103313320302022-3012202230320312-2112101111311001): complete subsection reference.

<a id="canonical-3312231132303023-0313321002210120-1100032032133210-2222122233111210-0001122010132331-0103313320302022-3012202230320312-2112101111311001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network.ref` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.virtual_network](resources--discovery--reference--group-002.md#canonical-1110120031123232-1103133033221300-2333022012020332-2112320331303032-0100001133321112-2231130320002221-2033320232220130-0031211123210313)
- where.virtual_network.ref

<a id="canonical-0313030022322130-0201012102111211-1323203223021023-3221032231313023-2111033001103001-1132232322001222-1111311222123012-0332300301000303"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual network direct reference.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011201123211022-2122013223313033-0202232100330333-3331032131333021-2312313100010020-1232101211022020-3011021200300111-2332211223113102"></a>

### Direct properties for `where.virtual_network.ref`

<a id="canonical-2310303223212320-0321110030100012-0212032031110030-1130130311202121-1230312032102123-2131033233032120-1233201003302012-2213110133030220"></a>

#### `where.virtual_network.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1101213121031030-1010311321101310-2001033232001130-0101113320121100-1011000021311202-0020323232033133-3100213301032230-3332132120111020"></a>

<a id="canonical-3223023010133302-2003201332310312-0032112131120311-0303213021132223-2232220113133021-0300131203013302-0132010132121330-2312303330222013"></a>

#### `where.virtual_network.ref.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2230332102131303-2302330121300100-2031333221222002-3221200012111112-2203001021012013-2112222202300313-2100033221301220-3322312033320301"></a>

<a id="canonical-0130301121230310-2121220322300022-0113122111011233-1012123130330220-3202223031202230-1320333232112002-0220032112113103-1013132132020030"></a>

#### `where.virtual_network.ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2303110203312302-2123030210132103-3330002012000123-0313233002221332-2100021101111202-2113323111313320-0311032233113213-0133322332013222"></a>

<a id="canonical-3102220232223202-2210233012002001-1002332011132003-3103212313331230-1311313212231031-2223102000233032-2011001202102223-2233332231102302"></a>

#### `where.virtual_network.ref.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1330031302023030-1020030301031332-3120221321302000-2120021120311320-1033000332231101-1131022222011001-1133321202003102-1201031121313323"></a>

<a id="canonical-1210233302303003-2323112012331120-1213223233103232-3311022330031131-1032132321203331-0001301310022122-3120203201201220-0122233102013203"></a>

#### `where.virtual_network.ref.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- where.virtual_site

<a id="canonical-0322133000130021-2333312212222111-0020101213231023-2030310233221013-0020031031000102-3020113021303003-1120101000110032-0110002030131332"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311132013333132-3103023120120022-0331032332221013-1003332111032211-2102310102313011-3031002330113012-2212120110103302-0310133310333310"></a>

### Direct properties for `where.virtual_site`

- [disable_internet_vip](resources--discovery--reference--group-002.md#canonical-0012000002320212-1110003030333201-2200132323301111-0031003322022332-2131310131100122-1020102030213221-1322011213212030-0212233130001000): complete subsection reference.

- [enable_internet_vip](resources--discovery--reference--group-002.md#canonical-0300132310111200-0202012132102100-3100103102212123-0313102300221333-3102331310321103-2121102033002200-2111330113031323-1011211010002010): complete subsection reference.

<a id="canonical-2000131222330332-3332323003332032-2221323000032331-0203112013033230-3122310103120230-2102103122211223-3222321131122303-0133011000320313"></a>

<a id="canonical-2303030313021123-2330030223120001-0233122030132333-1210330223021333-1001201101003102-3323101212230033-2132133112120110-2022110120032230"></a>

#### `where.virtual_site.network_type` property

Type: `"string"`. Optional.

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

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
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
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

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

- [ref](resources--discovery--reference--group-002.md#canonical-2313102201211221-1123332010222132-3130002111113333-1310230310302021-1132233010120030-2130230020331133-3331301032013003-3310010212002322): complete subsection reference.

<a id="canonical-0012000002320212-1110003030333201-2200132323301111-0031003322022332-2131310131100122-1020102030213221-1322011213212030-0212233130001000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- where.virtual_site.disable_internet_vip

<a id="canonical-2312101123222101-2011011201100032-2221012332123301-3002130230322301-1332302020000333-1302213311003311-3230331320112310-2330320003131110"></a>

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
disable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300132310111200-0202012132102100-3100103102212123-0313102300221333-3102331310321103-2121102033002200-2111330113031323-1011211010002010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- where.virtual_site.enable_internet_vip

<a id="canonical-3021002022120123-1022010032122021-1110031333111120-2230322030033022-2320033033023333-2330031110211102-3331023012233013-1113232210103211"></a>

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
enable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313102201211221-1123332010222132-3130002111113333-1310230310302021-1132233010120030-2130230020331133-3331301032013003-3310010212002322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.ref` properties

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Property reference](resources--discovery--reference--group-001.md#canonical-1000122023131121-1311020012032030-2231333201001000-3132033312130132-3123110232132122-0111031230023103-1231303130211110-1311211000221222)
- [where](resources--discovery--reference--group-002.md#canonical-1023323010232302-1333131303112022-3300103023013000-3132230030212330-3332300120230032-3102010223021221-1000322012101233-3100000121300010)
- [where.virtual_site](resources--discovery--reference--group-002.md#canonical-2223322232211101-1111031212021233-2331110221020110-1113201232123120-1032003111222123-1221012223203322-0130200312133221-2222033300110330)
- where.virtual_site.ref

<a id="canonical-2223002323111232-1312332333333231-1100100132112221-0012333303020122-3102332001111121-1232303103331221-2230332023123001-3031110313130223"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual\_site direct reference.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231010303212211-0023120113230030-3333231321002131-2130101012130133-0111231331301120-0112302120111210-0202220021002122-1300201121111211"></a>

### Direct properties for `where.virtual_site.ref`

<a id="canonical-2310233223102032-3200231202102012-0332102020222202-2321213203231033-0131300210132221-2122310230200102-3213023230113112-2123333320120000"></a>

#### `where.virtual_site.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2321321030101011-1330320131023013-0333321210013022-1000011201113301-1300033212122000-3213323232321100-0210203301323011-0222000031001100"></a>

<a id="canonical-2011012132322233-3321001123303203-1100010203330303-1133103121123120-3121101211012321-2313323122023330-0320112212210223-0331002221322223"></a>

#### `where.virtual_site.ref.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1310133233122030-3311122223101211-1231133000123210-1233003002231310-0232031212202101-0000111202002232-0013132332033210-2002201110021133"></a>

<a id="canonical-3133131120333011-1320232223002011-3321131031133021-1311020320003011-2311130132030330-1302003231003220-3030112323033123-0230102220123013"></a>

#### `where.virtual_site.ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0130012332031021-3010032323230100-1022321222223111-3033330120201333-1313333323313330-0131100032320203-2032021301312030-0011231233211132"></a>

<a id="canonical-2231203131101323-0002113030323202-2100320303010001-2021222230230123-0003202313031020-3202310201021032-1010213113122212-3102232312122121"></a>

#### `where.virtual_site.ref.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3211223011011211-1221223322021300-1303303103320033-1103022230301222-3320010130100231-0101312233120020-2030010012112132-1120233103310133"></a>

<a id="canonical-3213333300111022-0130221231310203-3320202201230102-3213313203222230-1202023231333120-3121322303302000-2322322110231120-1203002022013033"></a>

#### `where.virtual_site.ref.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
