---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-3323032303112012-3022221132311311-2010023233121102-3130133130021112-2211023322231311-1131213033021222-3220020013010221-3010030100310012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_token` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- http_receiver.auth_token

<a id="canonical-2203000202112220-1133020201212101-0031233310023301-2031202223131123-2332210232322123-3320030020121211-2130301211222110-3102122202320212"></a>

Type: `"single"`. Computed.

Access Token. Authentication Token for access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3131332303323020-3112313303100033-1030032222102121-2213220201120013-1212313112231213-2302310022320010-2223122110310210-1033330300012213"></a>

### Direct properties for `http_receiver.auth_token`

- [token](data-sources--global_log_receiver--reference--group-003.md#canonical-2201133021212200-2303301223000231-3102300001331113-3023110323032000-3213313102021101-1310223123221022-2010313010210032-0323030222002033): complete subsection reference.

<a id="canonical-2201133021212200-2303301223000231-3102300001331113-3023110323032000-3213313102021101-1310223123221022-2010313010210032-0323030222002033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_token.token` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-003.md#canonical-3323032303112012-3022221132311311-2010023233121102-3130133130021112-2211023322231311-1131213033021222-3220020013010221-3010030100310012)
- http_receiver.auth_token.token

<a id="canonical-0201021301103231-3001330030020313-0200101121102202-2030112032102122-2300331232201200-0312111003103310-3213211323101101-3130111001031131"></a>

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

<a id="canonical-3133121231323002-1000200223133323-3203102033030301-2303033120211030-1033013022311201-0330030002020210-2101223313113212-2213213302211130"></a>

### Direct properties for `http_receiver.auth_token.token`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-3122131103202331-3003103122023132-0223133112022331-2232011333011223-0000203231220000-0031310023211023-1201213232013200-2132110210233021): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-0201023123111221-3303233302030022-1322020011222002-0313133211211212-0211110120023023-0211131312213200-1013113102111313-1111300112023212): complete subsection reference.

<a id="canonical-3122131103202331-3003103122023132-0223133112022331-2232011333011223-0000203231220000-0031310023211023-1201213232013200-2132110210233021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_token.token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-003.md#canonical-3323032303112012-3022221132311311-2010023233121102-3130133130021112-2211023322231311-1131213033021222-3220020013010221-3010030100310012)
- [http_receiver.auth_token.token](data-sources--global_log_receiver--reference--group-003.md#canonical-2201133021212200-2303301223000231-3102300001331113-3023110323032000-3213313102021101-1310223123221022-2010313010210032-0323030222002033)
- http_receiver.auth_token.token.blindfold_secret_info

<a id="canonical-1012012301303002-1302322301202121-1000012132031113-1231221233301310-3222313323223112-1321233122103202-1203112213203100-3001321101121302"></a>

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

<a id="canonical-0110032332231030-0320310323131300-1330220213110200-0020003112222313-1221202101301111-0113032033003111-1202020311011202-1020120200310132"></a>

### Direct properties for `http_receiver.auth_token.token.blindfold_secret_info`

<a id="canonical-1301032122200101-0101220020222202-2232121022203212-3311133210010123-2120313220103032-1213302200032232-1210312301312232-1210012101211312"></a>

#### `http_receiver.auth_token.token.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0233123013332313-1102331232332101-3020200233000213-1101102002232032-2223212020310313-0320013201212121-0333021302210103-2012201003310322"></a>

<a id="canonical-0023311001233312-1201321013320111-0113113202111132-3103010223100211-3000233111200200-1211012023321312-0013111233202231-1310110210333311"></a>

#### `http_receiver.auth_token.token.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2020001013201213-1131120222230121-3312010002321331-1123020000321320-1312131332321203-1021213222232011-0230100110301013-0220131132200121"></a>

<a id="canonical-0312321232221233-3330333012102221-3010122102231120-1032313302012130-1001123111002333-1221320111123021-0230002102023032-2313110231120332"></a>

#### `http_receiver.auth_token.token.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0201023123111221-3303233302030022-1322020011222002-0313133211211212-0211110120023023-0211131312213200-1013113102111313-1111300112023212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_token.token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.auth_token](data-sources--global_log_receiver--reference--group-003.md#canonical-3323032303112012-3022221132311311-2010023233121102-3130133130021112-2211023322231311-1131213033021222-3220020013010221-3010030100310012)
- [http_receiver.auth_token.token](data-sources--global_log_receiver--reference--group-003.md#canonical-2201133021212200-2303301223000231-3102300001331113-3023110323032000-3213313102021101-1310223123221022-2010313010210032-0323030222002033)
- http_receiver.auth_token.token.clear_secret_info

<a id="canonical-1101010302032113-1203123130203303-3220122113321002-2113031301001331-0123112023310021-2331231311001010-2031002312030112-3203101331230122"></a>

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

<a id="canonical-1102222010312223-0112330233212300-3223021302303333-1211110110201103-0000120203221200-1102101000303232-1322022113012303-2112302223213331"></a>

### Direct properties for `http_receiver.auth_token.token.clear_secret_info`

<a id="canonical-3133000002133303-2103223113130302-0111121112010311-0123131331213103-1223222001302020-1301132031013121-1210200123103121-0233222021201221"></a>

#### `http_receiver.auth_token.token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3333011102322313-3302133232110023-3031310310133201-1212320332310130-2100131311010201-0013010331102123-2232212331112123-3233001303332101"></a>

<a id="canonical-3222303013320312-0212302313322322-1133213213000303-0130103333211202-3231031003303331-0011323301000013-0302303222130201-1302313222221201"></a>

#### `http_receiver.auth_token.token.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1323033310020132-0332013212020031-0212133332203131-1102212322332103-0133202002022001-1321122021112302-0003122230101102-1103000322223201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- http_receiver.batch

<a id="canonical-3033002330211213-2112021121332210-0032223330233100-0232112000221301-0101211123332200-2111212201311020-1130122003111211-1201230201202100"></a>

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

<a id="canonical-2202203301323201-1103213102030312-1202232002002122-0230323200222001-3301332313330123-0301301030223232-1220311033300112-3100011323023112"></a>

### Direct properties for `http_receiver.batch`

<a id="canonical-3120300312102212-2203000021131022-1120000010211023-1023012100011313-2230221201010213-0100320122111321-1010330211121020-0110200302030300"></a>

#### `http_receiver.batch.max_bytes` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-2133301222203022-1101003231200000-1203302021212111-0020202020312322-0220230003301132-2003321130003113-0002233030020211-1121121033222203): complete subsection reference.

<a id="canonical-2221101300030012-3321002023010312-1322213010011003-3200311300100302-0020200203301110-2101310002232013-3000102321031103-1130230131220102"></a>

<a id="canonical-2131213200031003-0203332011031123-0120013022220012-1333000023323002-0010311032011133-1021203013312203-0301011302221213-3223113111301123"></a>

#### `http_receiver.batch.max_events` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-1311211113012122-1110331310330102-3213202210330323-1333123300311320-1122300102202201-1332032132102320-2012302123011100-2022011020000230): complete subsection reference.

<a id="canonical-0102003133003300-1303311322102213-3212332232031303-0031013110013002-2213020320202111-1011213011001030-2100201113323131-0311032330122333"></a>

<a id="canonical-2221021123013321-3113133121020033-2133123120311231-0121011303301110-0321320110101031-1133302300312023-1020311231223102-1020300213231010"></a>

#### `http_receiver.batch.timeout_seconds` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-3301010000122032-3033331111021210-0102322010313311-1123200212213001-3331312102132101-1003112223213230-3033011002332220-1110301211003000): complete subsection reference.

<a id="canonical-2133301222203022-1101003231200000-1203302021212111-0020202020312322-0220230003301132-2003321130003113-0002233030020211-1121121033222203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-1323033310020132-0332013212020031-0212133332203131-1102212322332103-0133202002022001-1321122021112302-0003122230101102-1103000322223201)
- http_receiver.batch.max_bytes_disabled

<a id="canonical-0003000123021323-1203210121303101-1130112333130202-0212312121313220-0021001213033012-1131120323201311-2321202011321203-1323123301333111"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311211113012122-1110331310330102-3213202210330323-1333123300311320-1122300102202201-1332032132102320-2012302123011100-2022011020000230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-1323033310020132-0332013212020031-0212133332203131-1102212322332103-0133202002022001-1321122021112302-0003122230101102-1103000322223201)
- http_receiver.batch.max_events_disabled

<a id="canonical-1013132311021110-2211212013131133-0022002233233033-2303313010110012-1212313033112133-3010012102112030-0123311220033203-1310230101230000"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301010000122032-3033331111021210-0102322010313311-1123200212213001-3331312102132101-1003112223213230-3033011002332220-1110301211003000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-1323033310020132-0332013212020031-0212133332203131-1102212322332103-0133202002022001-1321122021112302-0003122230101102-1103000322223201)
- http_receiver.batch.timeout_seconds_default

<a id="canonical-0032213100301021-1200320130203012-1221102031021231-3223030013123122-2101232113013112-2232003223310313-2222031210301032-3100011331102013"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131223203223130-2000122203303022-1023002322213322-2120332113101311-3003321112120110-2022230313200233-1310131000132331-0100333201211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- http_receiver.compression

<a id="canonical-1230110031111023-2110110002013300-2331231310011020-0030213032233130-0121121133031320-3301012011020222-1101230213201022-2332132321133123"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Additional upstream details:

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

<a id="canonical-3103222132322032-1302203223130223-1020012200113123-2130332310120232-0033231122131103-2221322101330133-0103300101202003-2013201330110110"></a>

### Direct properties for `http_receiver.compression`

- [compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-0323033223122320-2023310003200132-3011103233122001-3230312221132131-3320311112121012-2013222310022000-3323010113020333-0011133331012012): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-2200031021200203-1013321021000212-2322313221131000-1132332101301223-1221020101021213-2203311130221111-0311020003210211-3011313110303321): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-3303023230322120-0212332020300101-2322122210232001-3212032303212200-1203103200230002-0133132032131013-3302010111001101-0000320111003002): complete subsection reference.

<a id="canonical-0323033223122320-2023310003200132-3011103233122001-3230312221132131-3320311112121012-2013222310022000-3323010113020333-0011133331012012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1131223203223130-2000122203303022-1023002322213322-2120332113101311-3003321112120110-2022230313200233-1310131000132331-0100333201211202)
- http_receiver.compression.compression_default

<a id="canonical-1100121131310320-3232022323322030-2230333220200230-1311012303002011-2101201211100131-3112213313232330-3100103301211101-1222303030122322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200031021200203-1013321021000212-2322313221131000-1132332101301223-1221020101021213-2203311130221111-0311020003210211-3011313110303321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1131223203223130-2000122203303022-1023002322213322-2120332113101311-3003321112120110-2022230313200233-1310131000132331-0100333201211202)
- http_receiver.compression.compression_gzip

<a id="canonical-2301200110202101-2112323302023131-0323300123313122-2230102222020032-3232323112132210-0022330123000220-3331332103013230-3100030110131230"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303023230322120-0212332020300101-2322122210232001-3212032303212200-1203103200230002-0133132032131013-3302010111001101-0000320111003002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1131223203223130-2000122203303022-1023002322213322-2120332113101311-3003321112120110-2022230313200233-1310131000132331-0100333201211202)
- http_receiver.compression.compression_none

<a id="canonical-3321022233022230-1131021301100122-0322000321023003-2330000103121000-2203233233200210-1132221011232232-2113123112021132-2210033101323203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203222233032313-0213102110222011-2300132012120130-2123230023320010-0202222331102302-2230031332332231-2220302333122321-2010013030330222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- http_receiver.no_tls

<a id="canonical-1311130223100102-1002321010003212-3312331023231130-3133110201012032-0132132030023130-0222030110203332-2001333212132132-3203001321102320"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- http_receiver.use_tls

<a id="canonical-2222103220232031-2031032010121130-2211221233312033-0332301331300312-2120012113132301-1213303302130011-2202233100003010-2022303330332120"></a>

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

<a id="canonical-3103003023300212-2312322033300103-2032101203222102-3331121013101010-0223003230120133-0213313000111331-0030202120222022-2102021101221033"></a>

### Direct properties for `http_receiver.use_tls`

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-3121210011330223-0011111000110101-1030131212012132-2102023131112230-2220102020223101-2322231221313100-3232012030001031-1332131102220010): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-0023301120112000-3120122331121331-1122230222132231-2221031123010011-0312011213312333-0022213312120331-2211202002220002-2132230200032101): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-0120130211222022-3201212221100103-1101202210112001-1211110200113300-0103321001132022-2131230210313302-3300121031322300-1130330013133133): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-0111321111331201-1120111320131111-1121213021103220-3222102320100220-1021233221102302-3331200100022003-1030220103333200-3221310112302033): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-1221213212132032-2030312030121210-2101011333001223-2013110011120031-0212322100331312-0303103123212310-2311112312030320-2303223110233330): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-3013001012332100-1311123110003032-0230112313223203-2032232022012302-3123112231323032-2103101220012102-1323131022301203-0011200000223032): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-3113211300002122-3331111132231131-0202303003230033-0100302132213000-2302200133221203-0310001301322232-2010131303313100-1133201010330223): complete subsection reference.

<a id="canonical-0302313312222112-0030303133011212-3021220013013321-3123013131232303-3211233023211032-2030212111022112-0133112030200120-0003100032123032"></a>

<a id="canonical-1313031311221100-3221213022221233-2103110302223013-0223032020312330-1021212332222012-2031133310331111-2032230112113111-1201031212032033"></a>

#### `http_receiver.use_tls.trusted_ca_url` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3121210011330223-0011111000110101-1030131212012132-2102023131112230-2220102020223101-2322231221313100-3232012030001031-1332131102220010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- http_receiver.use_tls.disable_verify_certificate

<a id="canonical-1012311323221212-1102020011022033-2323111232000032-1111101011332113-1331301002323202-0301331220010002-1003330322203033-1022303203233220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023301120112000-3120122331121331-1122230222132231-2221031123010011-0312011213312333-0022213312120331-2211202002220002-2132230200032101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- http_receiver.use_tls.disable_verify_hostname

<a id="canonical-2300022131232303-0333023333023330-3301102313132320-3003310311323320-0320312230101203-2003111011031200-1012120022112223-2021122200120100"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120130211222022-3201212221100103-1101202210112001-1211110200113300-0103321001132022-2131230210313302-3300121031322300-1130330013133133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- http_receiver.use_tls.enable_verify_certificate

<a id="canonical-3232301210023212-0212212322110103-2200322003223232-1221013221010220-1121321120032311-2332122323030113-1331031302030120-2332110300022131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111321111331201-1120111320131111-1121213021103220-3222102320100220-1021233221102302-3331200100022003-1030220103333200-3221310112302033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- http_receiver.use_tls.enable_verify_hostname

<a id="canonical-3320320313031323-0100101101321230-0313122233131211-3313212130113120-3001211033122312-2100120310000110-3303320002133330-0011231032020230"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221213212132032-2030312030121210-2101011333001223-2013110011120031-0212322100331312-0303103123212310-2311112312030320-2303223110233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- http_receiver.use_tls.mtls_disabled

<a id="canonical-0010222301110230-0311211203023000-3010201033013031-0331230300332203-1213100321300013-0020200332123210-2113021203110132-3302233303222012"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013001012332100-1311123110003032-0230112313223203-2032232022012302-3123112231323032-2103101220012102-1323131022301203-0011200000223032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- http_receiver.use_tls.mtls_enable

<a id="canonical-2113120302213213-1123011121122130-3212233323101200-3310302003320121-3223232330321201-0232202133330102-0021012322103320-2301311211223020"></a>

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

<a id="canonical-1210213221233321-3232202032100321-0013003200221033-3211200333030001-2032110011231002-0103120210312130-1110221323323102-0122000102223103"></a>

### Direct properties for `http_receiver.use_tls.mtls_enable`

<a id="canonical-3011232331102121-2330330322002033-2201333131022212-3123023222300013-3123221233010012-3301202013213232-3202322230223333-3213002333212222"></a>

#### `http_receiver.use_tls.mtls_enable.certificate` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-1133103010301301-0020232210020321-3030310100333102-3111211210210310-0231202311212021-0213010021023320-3010322033133313-2332103233321312): complete subsection reference.

<a id="canonical-1133103010301301-0020232210020321-3030310100333102-3111211210210310-0231202311212021-0213010021023320-3010322033133313-2332103233321312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-3013001012332100-1311123110003032-0230112313223203-2032232022012302-3123112231323032-2103101220012102-1323131022301203-0011200000223032)
- http_receiver.use_tls.mtls_enable.key_url

<a id="canonical-1113302222312031-0310021301203101-2021332030311232-1232130111100222-0303001200231112-0303012311213320-3202212012111113-3200112103330102"></a>

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

<a id="canonical-1033133203311120-0231120222012303-1203022001302311-1230333332312120-2212322000000010-0210032223100033-0233032232120201-3320133130330122"></a>

### Direct properties for `http_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-3311023020321022-3001233101130221-2312333032200203-2103220012320313-2232031301023023-3122033131233113-2201020311101323-0321013002212311): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-0102312301010323-0101310233032311-3203011200311011-1323233023200231-2201020323132312-3220133231222220-2123300110200312-1010333011223212): complete subsection reference.

<a id="canonical-3311023020321022-3001233101130221-2312333032200203-2103220012320313-2232031301023023-3122033131233113-2201020311101323-0321013002212311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-3013001012332100-1311123110003032-0230112313223203-2032232022012302-3123112231323032-2103101220012102-1323131022301203-0011200000223032)
- [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-1133103010301301-0020232210020321-3030310100333102-3111211210210310-0231202311212021-0213010021023320-3010322033133313-2332103233321312)
- http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-1313132100103300-2200020030302011-1223231310013233-0103203020020031-2003212223032200-3313032311120013-0321200000320303-1032002021021101"></a>

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

<a id="canonical-2130200123032301-0332303301100211-0102132113330212-3313001203132221-1002123322032121-3231333001121012-3122023032132010-0132211333101033"></a>

### Direct properties for `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-1321022223131123-1331320221102321-1122301032020222-3012132233001103-3132031312203023-1023303322300312-2300223303102230-0133331310112201"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3003300203011300-3330100031310112-3123212021122332-3321031300113210-3031111322030303-0003300203233330-0133103031011231-2132201003100100"></a>

<a id="canonical-0221302231203321-2032122001201230-0223230223001032-0212221310223330-1312333301302011-0232320221023203-1312231222013113-1220103133323003"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0021021202012031-2030011233031213-2020202332312033-2123101131202030-3130123233222133-2231320310030123-1132310301101033-2322001113221133"></a>

<a id="canonical-0220200322031021-1303201122121223-2210301020130220-3200330321032221-3313110211000303-1000010311033123-1003211333310123-0212210322003111"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0102312301010323-0101310233032311-3203011200311011-1323233023200231-2201020323132312-3220133231222220-2123300110200312-1010333011223212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-3013001012332100-1311123110003032-0230112313223203-2032232022012302-3123112231323032-2103101220012102-1323131022301203-0011200000223032)
- [http_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-1133103010301301-0020232210020321-3030310100333102-3111211210210310-0231202311212021-0213010021023320-3010322033133313-2332103233321312)
- http_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-3213122221111213-0130101301002232-1023211111130201-2011230120121000-0013203010303301-0313131121311333-3133112230023002-1221331111133202"></a>

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

<a id="canonical-2323212031101032-3302303010132230-1012020301002123-2011301332323001-2002033310233223-1011323331032322-3000021000133011-1112302102303021"></a>

### Direct properties for `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-1012123123310222-3001133331320300-1001331222101302-1120312301001233-2212011231120332-3231312102211222-3023003001031130-0222320020333331"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2021121112020000-3302012323231110-0002100122232233-1330322333103002-2323322132211202-2100123301003120-3012000313313032-3003000230002212"></a>

<a id="canonical-2020111213231010-3132203130220200-1211020210120030-1023301111332312-1130302032022203-2222203021230321-2222231020003302-1011223000030323"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3113211300002122-3331111132231131-0202303003230033-0100302132213000-2302200133221203-0310001301322232-2010131303313100-1133201010330223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330)
- http_receiver.use_tls.no_ca

<a id="canonical-1130022333100312-2220220200023033-2002231330211320-3113330221112013-3110303002101020-2210033032213230-3231301230010201-1311100130023201"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- kafka_receiver

<a id="canonical-2203013201303101-3030330033021301-1103130222112033-1122003330112030-1222132131131010-1103203132001032-2212100031032010-3020032111211321"></a>

Type: `"single"`. Computed.

Kafka Configuration for Global Log Receiver.

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

<a id="canonical-0213101131210113-2033333232222333-3001203222010011-2313311223120211-2311332012213113-1120031131200010-0120013230203203-1102000023231021"></a>

### Direct properties for `kafka_receiver`

- [batch](data-sources--global_log_receiver--reference--group-003.md#canonical-2101111322112303-1120032221330021-1031320201133133-3312123101031222-3111100003311222-2231313323211022-2032230032221231-3132302202013131): complete subsection reference.

<a id="canonical-1332023031001203-2013031021030231-2333220022313323-3231313313301031-1130012101013112-0211011313200030-2223223120313020-1032231122221331"></a>

<a id="canonical-3202131330133030-2230030222030103-0201113130112122-3100000033101102-0203030320101221-0110001323033221-2123032013321103-0301023031232010"></a>

#### `kafka_receiver.bootstrap_servers` property

Type: `["list", "string"]`. Computed.

List of host:port pairs of the Kafka brokers.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [compression](data-sources--global_log_receiver--reference--group-003.md#canonical-2322211302002023-0221320330100320-3311021103023112-0113301202202302-0023200222201123-1212001312313332-1221100211201320-2133103013012322): complete subsection reference.

<a id="canonical-2012333120313110-1131020310021033-0032333230112121-0300112112131013-0122211121323213-0033323201201223-2003132232121221-0222322101330123"></a>

<a id="canonical-1200112012300312-0021310012001302-0031221002121102-3311030133221031-3201101213302320-3022303231110332-2103302010111333-2213023003022332"></a>

#### `kafka_receiver.kafka_topic` property

Type: `"string"`. Computed.

The Kafka topic name to write events to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  }
}
```

- [no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3030211300033120-1323113022231332-1023203302233310-3233200230030213-1320011330333112-3032010023113013-0322311302022211-0013221012322101): complete subsection reference.

- [use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210): complete subsection reference.

<a id="canonical-2101111322112303-1120032221330021-1031320201133133-3312123101031222-3111100003311222-2231313323211022-2032230032221231-3132302202013131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- kafka_receiver.batch

<a id="canonical-1103322123332110-2312301222223013-3030323302031000-3033210312300312-0200330221132131-3012110022100232-2021001131233012-0002011100001301"></a>

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

<a id="canonical-3221301301122211-3313201312113120-0201030113313022-3113032332313021-2103112221322333-0200231332132013-0101121102113032-0002131001221030"></a>

### Direct properties for `kafka_receiver.batch`

<a id="canonical-2012330020221010-1211133223123213-1200110031230300-2021131210323110-1032021203313023-1031201211223211-1111300301111030-3023013110313030"></a>

#### `kafka_receiver.batch.max_bytes` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-1221013032033203-2220320223000221-1101111120033010-1111003220333023-0211322232233322-2322013313133113-2223200103320231-2131123210001012): complete subsection reference.

<a id="canonical-3303232321010230-0212022230010300-0000203323230112-1330122203020013-1132300202120332-3312122123322020-0131132302123210-3310102022322131"></a>

<a id="canonical-2230103321330121-1230002200033212-0303010321232031-2130022310323013-0312203130100022-1332013220333030-0232333300220122-1233302323311110"></a>

#### `kafka_receiver.batch.max_events` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [max_events_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-0230232211003013-0003322213330001-3130320332030120-1110231100023121-0133331110003103-1031312123332113-2220033111232022-0223020223331223): complete subsection reference.

<a id="canonical-2311311020333333-0300321223100302-0023121312213012-1203021023302230-0030231012010132-0221311033221333-3333032311201332-3030000231333232"></a>

<a id="canonical-2000230320010021-0012011021201201-2021210133200331-0000112220102323-1131311302323120-1313221232002231-0033121302103211-1220222000233203"></a>

#### `kafka_receiver.batch.timeout_seconds` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-003.md#canonical-3012302003103120-0333332201212300-3000202332101100-1122002101303221-3123031221303313-2310113123303010-1223123020212331-2030010223202002): complete subsection reference.

<a id="canonical-1221013032033203-2220320223000221-1101111120033010-1111003220333023-0211322232233322-2322013313133113-2223200103320231-2131123210001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-2101111322112303-1120032221330021-1031320201133133-3312123101031222-3111100003311222-2231313323211022-2032230032221231-3132302202013131)
- kafka_receiver.batch.max_bytes_disabled

<a id="canonical-3112302132301323-1021330023022202-0111232333331202-0300130322031002-1030022130133302-3033223132301201-2303113313220032-0203120131221221"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230232211003013-0003322213330001-3130320332030120-1110231100023121-0133331110003103-1031312123332113-2220033111232022-0223020223331223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-2101111322112303-1120032221330021-1031320201133133-3312123101031222-3111100003311222-2231313323211022-2032230032221231-3132302202013131)
- kafka_receiver.batch.max_events_disabled

<a id="canonical-3330200310000303-2103100211320103-1010331303313002-2202100201230130-3211200203200012-0031302130330332-1301010332101201-3220103200030020"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012302003103120-0333332201212300-3000202332101100-1122002101303221-3123031221303313-2310113123303010-1223123020212331-2030010223202002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.batch](data-sources--global_log_receiver--reference--group-003.md#canonical-2101111322112303-1120032221330021-1031320201133133-3312123101031222-3111100003311222-2231313323211022-2032230032221231-3132302202013131)
- kafka_receiver.batch.timeout_seconds_default

<a id="canonical-3132213100123201-3103312231110023-1221200233103133-1222331202010101-1103021133320201-1300010213133301-0000212033130022-0032210302112303"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322211302002023-0221320330100320-3311021103023112-0113301202202302-0023200222201123-1212001312313332-1221100211201320-2133103013012322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- kafka_receiver.compression

<a id="canonical-3122313310131001-0022220321013131-3310132233222032-2201132223011310-1233121232201212-1223222103112023-3302122222032323-1331330312133220"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Additional upstream details:

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

<a id="canonical-1102030012320301-0103020310320302-2000203310333020-2213131222023003-2212121121010210-3332011310012000-2300021321030131-2020131202203320"></a>

### Direct properties for `kafka_receiver.compression`

- [compression_default](data-sources--global_log_receiver--reference--group-003.md#canonical-3011020122320201-2321110010212201-3211110223201012-1123133020102123-3010300003111203-3020022021032311-2212030201231203-0321123302202302): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-003.md#canonical-3300213310211013-0330221230303311-0020113230330102-1020321303332003-1123001230200101-1120301011310301-3213133333120330-1200303211330210): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-003.md#canonical-0333332011320132-3330201313222031-2022301122201310-0222101221003311-3132032133313213-3210011211122221-1101222210331020-0331320020033232): complete subsection reference.

<a id="canonical-3011020122320201-2321110010212201-3211110223201012-1123133020102123-3010300003111203-3020022021032311-2212030201231203-0321123302202302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-2322211302002023-0221320330100320-3311021103023112-0113301202202302-0023200222201123-1212001312313332-1221100211201320-2133103013012322)
- kafka_receiver.compression.compression_default

<a id="canonical-3020100312300133-1000020221310322-0333010222100321-1320231212103302-0101021311211113-1203021120112312-2322202303212202-0213010003200103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300213310211013-0330221230303311-0020113230330102-1020321303332003-1123001230200101-1120301011310301-3213133333120330-1200303211330210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-2322211302002023-0221320330100320-3311021103023112-0113301202202302-0023200222201123-1212001312313332-1221100211201320-2133103013012322)
- kafka_receiver.compression.compression_gzip

<a id="canonical-2011112302112310-3303121002333320-2220113131113232-3013210311202323-2210302001230003-0000022221121002-3120210123110001-2213000333120011"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333332011320132-3330201313222031-2022301122201310-0222101221003311-3132032133313213-3210011211122221-1101222210331020-0331320020033232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.compression](data-sources--global_log_receiver--reference--group-003.md#canonical-2322211302002023-0221320330100320-3311021103023112-0113301202202302-0023200222201123-1212001312313332-1221100211201320-2133103013012322)
- kafka_receiver.compression.compression_none

<a id="canonical-3201322230122102-2200203201202022-2001133323103333-0313122001013132-1313020333221231-1032123012310031-0211103130030101-2332311212020003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030211300033120-1323113022231332-1023203302233310-3233200230030213-1320011330333112-3032010023113013-0322311302022211-0013221012322101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- kafka_receiver.no_tls

<a id="canonical-2331001030323113-1111330223311321-3200012112103010-3333010100231012-3233301323323300-3201212022021030-1030010322201101-0212123032221022"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- kafka_receiver.use_tls

<a id="canonical-3330132200100100-0020102213030203-2233233303310310-0302331011332221-2020111331103031-1012333222101213-0331200301100212-0232233220311331"></a>

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

<a id="canonical-0300313233310033-2011012320311101-0020132113213322-2103112111310122-2311001032330210-2013132020121131-0203323213202101-3131232102122113"></a>

### Direct properties for `kafka_receiver.use_tls`

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-2113022112332013-3022133121223332-3311203120123321-0112032101010223-2133001033033101-0303302123210212-3131311231123121-0210320123022310): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-1101221011132122-2211001012033112-3313111312132101-2133320033300111-3133013000210202-3033112120101310-3011000101012313-1332311321222131): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-003.md#canonical-2201112000030222-0223232021001202-3223130121200320-3131200321223323-1333023323031310-3312200122321031-0221233021112012-3221113131300013): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-003.md#canonical-1303210202210020-3021300233113012-2103311122232100-0132012121130002-1120030111331200-1220223300303020-3101012121231003-2100323333233023): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-003.md#canonical-0320122110232312-0212313331031130-2231012230110120-0011203330231122-2020231213301332-3021200303322002-0010212102011213-2013310213212031): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-1201231113022120-1200133201312002-3002301323003121-2033213120211230-0131220130030100-2131032203333322-3312211321220321-0223112000332010): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-003.md#canonical-1232311211211321-3322003303301101-2300121102210211-2233323313033033-1111322213122101-2212001312010312-1111312200311300-1201200111020023): complete subsection reference.

<a id="canonical-0231203313002022-1201122033032121-2323203200101030-0013002101233001-2233330320211102-2002300300301022-0230301231132210-1303300201021310"></a>

<a id="canonical-3021322122132121-3223322230213022-3133311011310020-1031022011032102-2223323010130230-0321020110332210-1011321203212233-1220311032001103"></a>

#### `kafka_receiver.use_tls.trusted_ca_url` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2113022112332013-3022133121223332-3311203120123321-0112032101010223-2133001033033101-0303302123210212-3131311231123121-0210320123022310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- kafka_receiver.use_tls.disable_verify_certificate

<a id="canonical-1223201221313231-0031123331012030-1030320120200222-2220312320121030-2333303213300310-3201211320023330-2231232033010012-2202302021202300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101221011132122-2211001012033112-3313111312132101-2133320033300111-3133013000210202-3033112120101310-3011000101012313-1332311321222131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- kafka_receiver.use_tls.disable_verify_hostname

<a id="canonical-1021002132200323-1222113013010021-0201231111300321-2000110302031311-2021130313333232-2220200301223021-0022310003211231-2320101022100033"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201112000030222-0223232021001202-3223130121200320-3131200321223323-1333023323031310-3312200122321031-0221233021112012-3221113131300013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- kafka_receiver.use_tls.enable_verify_certificate

<a id="canonical-1103211111121003-0132101102312313-3220310300312323-3012120102112300-2032301000302131-0230022213221213-3333130301313132-2103211130222302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303210202210020-3021300233113012-2103311122232100-0132012121130002-1120030111331200-1220223300303020-3101012121231003-2100323333233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- kafka_receiver.use_tls.enable_verify_hostname

<a id="canonical-3302011300132310-0001201322302032-0213032210102320-3212102222230210-3122233221232320-3033332223300020-2303331322332333-2201103203022210"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320122110232312-0212313331031130-2231012230110120-0011203330231122-2020231213301332-3021200303322002-0010212102011213-2013310213212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- kafka_receiver.use_tls.mtls_disabled

<a id="canonical-2012100003222003-1233022313301120-1202032301021033-2131330033301310-2300200000211030-3202023103132133-2003201201123232-1212300302130113"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201231113022120-1200133201312002-3002301323003121-2033213120211230-0131220130030100-2131032203333322-3312211321220321-0223112000332010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- kafka_receiver.use_tls.mtls_enable

<a id="canonical-1033200222323211-2320212012032013-3131002201001211-2013131003222232-1203022133203000-2021031231321210-2112122300322010-2202002121230202"></a>

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

<a id="canonical-2200310002203210-3210231032210012-0031131013302011-1022012223221211-3013010121333231-3202122123110121-3113011233231201-1110101221031033"></a>

### Direct properties for `kafka_receiver.use_tls.mtls_enable`

<a id="canonical-0122122313131013-1210211301200301-1211003300100120-0311211011232311-2203033120323220-0203201103000120-2132113123012311-0331233100002102"></a>

#### `kafka_receiver.use_tls.mtls_enable.certificate` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-2222122111310200-1310132022011200-3323332100011131-1213112103002021-1033303210022120-1010302010313212-2123222120011011-1230303122201211): complete subsection reference.

<a id="canonical-2222122111310200-1310132022011200-3323332100011131-1213112103002021-1033303210022120-1010302010313212-2123222120011011-1230303122201211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-1201231113022120-1200133201312002-3002301323003121-2033213120211230-0131220130030100-2131032203333322-3312211321220321-0223112000332010)
- kafka_receiver.use_tls.mtls_enable.key_url

<a id="canonical-3311122312201101-1321301021320310-3220323330121130-1201303312223123-3300312323213101-0103033331232022-1233001321313330-2231101120111233"></a>

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

<a id="canonical-2321113023020310-3202020002002030-2203322033033132-2202212113232211-0201121003231120-0113222310330222-3322210130200101-2102211323022233"></a>

### Direct properties for `kafka_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-0011220232102223-3120012301131310-0211131122200031-1202000111331113-3301112312321033-0121022001301003-1323123103031113-2322213033001210): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-003.md#canonical-1302122213102100-3120132021301020-2102031210132333-3200030001232311-1333002033031221-2330012220223232-2003222232332301-2231030312201113): complete subsection reference.

<a id="canonical-0011220232102223-3120012301131310-0211131122200031-1202000111331113-3301112312321033-0121022001301003-1323123103031113-2322213033001210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-1201231113022120-1200133201312002-3002301323003121-2033213120211230-0131220130030100-2131032203333322-3312211321220321-0223112000332010)
- [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-2222122111310200-1310132022011200-3323332100011131-1213112103002021-1033303210022120-1010302010313212-2123222120011011-1230303122201211)
- kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0011011202120123-0021102131000112-3210332212001312-1031121022013120-0312332002221022-0023130221310331-0002300203011102-3022013223202011"></a>

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

<a id="canonical-0330123330120321-3010230100233021-2102321331100223-3001230020310330-1012302211331330-1131211103111032-3203001211213312-2120323132331301"></a>

### Direct properties for `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-1222112120321313-2313031003233123-2300131113101213-1212022030112223-0023221110313133-3102301311112221-1221110010113320-1332302131101211"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2210003301020133-0333331332133030-3031331300323302-1133230210022031-1121032202001020-1020123322020101-1302132031103333-2300321211133130"></a>

<a id="canonical-3310333013020100-2022002002003222-0033001301101003-1002131333030323-1310321322323233-3310323100200323-1103201002113010-1303223030210301"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2303221022012133-0310013031323222-2312023323111101-3203131313123231-1333332000302110-2232012130030033-1220123123323230-0232121022031010"></a>

<a id="canonical-3322231033222011-1233010223303000-3321301203032013-2232121011213011-3111011330001110-0110003010232113-0213222332313002-2001103321330303"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1302122213102100-3120132021301020-2102031210132333-3200030001232311-1333002033031221-2330012220223232-2003222232332301-2231030312201113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- [kafka_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-003.md#canonical-1201231113022120-1200133201312002-3002301323003121-2033213120211230-0131220130030100-2131032203333322-3312211321220321-0223112000332010)
- [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-003.md#canonical-2222122111310200-1310132022011200-3323332100011131-1213112103002021-1033303210022120-1010302010313212-2123222120011011-1230303122201211)
- kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-2201011011210212-0111313312010303-3300222233031030-2130331000323300-3033202002200033-1303022313230333-1120330003022233-2001311103233300"></a>

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

<a id="canonical-1101322221112122-3113032001120213-0331302102201230-3203100122012003-0101331100032333-3032200013032113-1221303001022310-2213031020120321"></a>

### Direct properties for `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-2121112010001300-2222021231100003-0132200232330133-1111121122100011-3300303310332212-2200332121003012-1312120020310230-2332012133031121"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2030131233002212-2120212302120310-3323113020220031-2102233313023301-1120311233101131-3103123301010231-0323321323102022-3123211313311203"></a>

<a id="canonical-3212212232030113-0113322021133200-0020232212101332-2101002001021133-1020213211233210-1011210302211021-0301100300233001-0123330300033312"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1232311211211321-3322003303301101-2300121102210211-2233323313033033-1111322213122101-2212001312010312-1111312200311300-1201200111020023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [kafka_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-2302232221211301-1023330311131303-2131023213032202-2233001323102332-1100202333303230-1210103123023323-2203330332112122-0212332220012331)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-2202203223111223-3132210101332220-3320222310120203-3231120010211331-0231320233232231-1020203212300220-2022230100333130-3011123320220210)
- kafka_receiver.use_tls.no_ca

<a id="canonical-2203302101301230-1301203331100223-1032001221101212-1103110311322233-3032330010321003-3030113030302222-2203130222002333-3121011132121102"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222203022011122-1213111132203013-1201312120332200-3233301331202013-3012133313300230-3320033231321122-3201112101101230-3232031212020130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- new_relic_receiver

<a id="canonical-3132102203031223-0300022203231020-2033232013233211-3131300201011022-0210221130300123-2213111030003231-1002103002302030-1301031211011001"></a>

Type: `"single"`. Computed.

Configuration parameter for new relic receiver.

Additional upstream details:

Configuration for NewRelic endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-endpoint_choice": "[\"eu\",\"us\"]"
}
```

<a id="canonical-1321221003221000-3020021031003121-1133010123132233-3123122132212312-0303211012313032-1023213323021003-1222210022321210-1013202101032332"></a>

### Direct properties for `new_relic_receiver`

- [api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-2211310101013202-2223001023310323-3031310222032320-3233233120202231-0101322122330110-1132232033003313-0230330322030212-0213101112010303): complete subsection reference.

- [eu](data-sources--global_log_receiver--reference--group-004.md#canonical-3300010212112002-3213231331131321-2331110230220010-2012032031333333-0321230230113030-2111100323020013-3113212012233230-0110103010023113): complete subsection reference.

- [us](data-sources--global_log_receiver--reference--group-004.md#canonical-3012303210122113-2313212321310023-3132111332302132-0011323210220302-3103132011013301-3021100321221320-2203201323101102-2002031320002012): complete subsection reference.

<a id="canonical-2211310101013202-2223001023310323-3031310222032320-3233233120202231-0101322122330110-1132232033003313-0230330322030212-0213101112010303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.api_key` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-1222203022011122-1213111132203013-1201312120332200-3233301331202013-3012133313300230-3320033231321122-3201112101101230-3232031212020130)
- new_relic_receiver.api_key

<a id="canonical-1300211121201200-2001033220100202-1220231130130320-0012300201031003-0332101013010033-0213323210133310-2311223210302030-2231013130131301"></a>

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

<a id="canonical-1101003100212023-0132332121010000-0231203003301120-0311000301322211-3212122112120112-2312113122321131-2321332112130110-0233233221302112"></a>

### Direct properties for `new_relic_receiver.api_key`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-0123001111011103-1013010100301101-1212323323100111-3331023101321021-3022311012301132-0321303310102333-0001021212221210-2010110203012101): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1320221211133021-1033300103012120-1112030330032201-1031201332123101-1323200101300013-2222012210121020-2200130103311203-2032232322030331): complete subsection reference.
