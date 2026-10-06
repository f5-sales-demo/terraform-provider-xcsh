---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3011333222233131-0113333113011123-3012100120011302-1112101232033213-2032213333000311-1223123111101003-2131011023013300-1103130003100103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-009.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-009.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-009.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-009.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-009.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-2123113313220110-0310302301200313-3212113121210233-3200210232310020-0110010012200211-1300033200023320-0202132210212032-0010330003311023"></a>

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

<a id="canonical-3200221111202312-0013223323302122-0032030231030120-1230023110130233-0303311120220330-3211233013221030-1210003301033302-3231220123200101"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info`

<a id="canonical-3322003002213231-0230310333030102-1201313103300313-2312120312212110-3320222111221010-2231130331313002-3003011231302320-2333332011012303"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1331210200012223-3311013231010113-3302111203111003-1313113203320212-3203012011122210-2322112310002101-3212322121102220-3330313223320323"></a>

<a id="canonical-2320312111013312-3221230123132013-2103233102121111-0012122122130300-2211310022103323-1103013122200030-2033032113211323-2120323003203110"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.url` property

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

<a id="canonical-0223301322310022-2112210323323023-1010031200031021-3002003330300110-1311032202200232-0120121012010001-2331132132312022-1323130100110212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.disable_api_crawler` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-009.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-3133131002302221-3223111311103332-0120003223003223-0020102313300121-1303332123233333-0312011031011130-0021320031010010-3112031301233302"></a>

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
disable_api_crawler = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-1012300311230310-0001121003022332-0223030222213323-1030300202303311-1300332101121211-3123203113211231-2221201300122130-1311020030100120"></a>

Type: `"object"`. single nested block, Optional.

Select codebase and Repositories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302121320122110-0332000002303032-2001131200102122-3010221110120102-0302013002033123-0031111313103323-0330001323121113-2333111301322310"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan`

- [code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031): complete subsection reference.

<a id="canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-1133201031102003-3131001102211330-2110133203202212-2310302321232220-3030011203110303-2113121210013301-1120112103332321-3031023312012223"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for codebase integrations.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_repos",
    "selected_repos")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
code_base_integrations {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333013212300302-2132321122200333-1100313110333221-3232313110102101-2330010333110330-3323101021203203-1002330132231123-3221310031313111"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`

- [all_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-1133222312233233-3230003202230132-0331020231301303-1013330021223021-2003332220203123-1333213033231321-2101011233031021-1301003213132232): complete subsection reference.

- [code_base_integration](resources--cdn_loadbalancer--reference--group-010.md#canonical-2121123213033030-0101132301023222-0210321103113032-3130313132023121-1023321313310131-1122011031001221-2122200023311221-1003322232210223): complete subsection reference.

- [selected_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-3332003330020233-1321203101311112-1032320222112212-0303123303022320-0130233233212212-2331122210022100-2012123323130120-2300212333133331): complete subsection reference.

<a id="canonical-1133222312233233-3230003202230132-0331020231301303-1013330021223021-2003332220203123-1333213033231321-2101011233031021-1301003213132232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-1000331013330222-3000113103113211-0001211010301300-1033013110223123-3221120202333012-2132002301113030-1001000130303322-0320300120233121"></a>

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
all_repos = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121123213033030-0101132301023222-0210321103113032-3130313132023121-1023321313310131-1122011031001221-2122200023311221-1003322232210223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-3200022321013010-2212130332311312-3132123120032023-3231000103301002-2121202022103002-3030130133310002-3310023202310011-0330112210321032"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111021330021332-0020032120013313-2310231313101220-1321010210201303-0000321301311332-3102122011332303-0023301101312010-0101311323021131"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration`

<a id="canonical-2220112313013320-1132100032313302-1000023013203022-2111020010001130-1013130133110121-3101130033123200-0020031032133123-1213112011220020"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2031312132122123-2121030213300113-1202213111200320-3230112321322101-2201031000211321-1020003321312133-3233303111003332-3210033322011222"></a>

<a id="canonical-1111021122133120-1300100202300011-1123300301131212-0002110221311103-3030203111331210-0021013101232333-2123330223102132-2113311222211213"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1300122230111202-2232322101032031-2012031232002231-3212101023303003-1231101131332320-2303301011231221-0311013221013232-2222110310033132"></a>

<a id="canonical-0032310122020112-2330023301133210-2121333231111111-0222231223113002-1100021003121202-0013130231301300-0233302332022113-0332212203310021"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3332003330020233-1321203101311112-1032320222112212-0303123303022320-0130233233212212-2331122210022100-2012123323130120-2300212333133331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-2023232133013100-2200211232213000-2201011220303123-1320133103031123-3102112121111312-2220023103310102-3100313313032210-1101013010020132"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("api_code_repo")}
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
selected_repos {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230001300133222-1222320301010301-3213330003233221-1111010030212233-1203312011200033-3231112333333121-0331210210121000-3210122110100203"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos`

<a id="canonical-0303003011020122-3333212220000322-2133133022133233-1221133210312310-0022110323330303-2231201103001302-2333311331321221-2022031032102221"></a>

#### `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos.api_code_repo` property

Type: `["list", "string"]`. Optional.

Code repository which contain API endpoints.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-2000000012223230-0211010000233313-1010013330320313-3003002303303013-3003012223220213-3332302231300133-2301220230113223-0113213212221220"></a>

Type: `"object"`. single nested block, Optional.

API Discovery Advanced Settings. API Discovery Advanced settings.

Receipt-pinned upstream constraints:

```json
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
custom_api_auth_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231221221302103-0211202002021321-2323021110133323-2303023132001020-2022202331221132-3322011233021212-1211030012303320-2102020231112210"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery`

- [api_discovery_ref](resources--cdn_loadbalancer--reference--group-010.md#canonical-1032323002010020-1022331203312301-0223203013030002-1003233130313211-0303313013223302-0222322103303123-0212103001230001-1121320112232330): complete subsection reference.

<a id="canonical-1032323002010020-1022331203312301-0223203013030002-1003233130313211-0303313013223302-0222322103303123-0212103001230001-1121320112232330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-2100223211110032-0022232233231102-3002001313210133-2120130303331110-3000320010313301-2113212302322012-0110123000000132-0013123202113003"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
api_discovery_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212101222033110-2113330102332322-3101332221113302-3020132121331313-1330201030121130-2221123003122231-0303023031102033-0213121133101020"></a>

### Direct properties for `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref`

<a id="canonical-1202013332032113-1311221103102202-0313222201100110-0220120112122133-1210103221112011-3121121213013331-3220132123302120-1002201232013312"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0211313023320013-0302112230002111-0003010130213213-1233001202103230-0323302221323222-3312132022011001-2003332110301021-2100211103311310"></a>

<a id="canonical-1321203231133220-1010000332303122-3122033002303320-2122132212222221-0113311001011012-2022220312110100-0200132002213202-1121032202300320"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1230233100131211-1010220312033333-2312032321233232-1303012123221001-0333021220203130-3333013233232321-1022100113322231-0102233302231111"></a>

<a id="canonical-3103110223120001-3220122002332200-3000230231202213-0311022300200313-2011231100133231-3110203233000211-3313100131331002-1313031212121213"></a>

#### `enable_api_discovery.custom_api_auth_discovery.api_discovery_ref.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1013221211221310-1031120002332131-0232001133000302-3000130332131232-1303121022332312-3012002312133111-1112121330233321-2233231122001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.default_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-3330312033101230-0332222200233332-0101032312130121-0031230222130131-0023200000201230-1103120303010303-1320313222002032-1020121123313123"></a>

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
default_api_auth_discovery = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331120230102020-0312102032323311-1020022132220113-1312101223301033-0301110002231123-3321223213133301-2033313330221330-1323102320001301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.disable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-1013231022323130-0133022002100312-2200002332003020-0032222010313033-3210300023103030-2312320031121232-0001223313301012-0033220122221221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable learn from redirect traffic.

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
disable_learn_from_redirect_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112030012103111-1212023301022210-3131303031130201-1232223310233020-1021210103230301-2110223201003332-1112033210222011-3100331330221113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.discovered_api_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.discovered_api_settings

<a id="canonical-3211302123213233-2213103023232312-1322230201223031-3113130221113323-3213022012202033-3103002213022020-0110212023330223-1131133203203212"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("purge_duration_for_inactive_discovered_apis")}
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
discovered_api_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221232302112131-0300300230320213-2213311201330311-2230032033031110-1112102323131210-1222333220010223-2032022332301002-2013112131112023"></a>

### Direct properties for `enable_api_discovery.discovered_api_settings`

<a id="canonical-0000013231022122-2030002133013321-3310220211312000-2200131132100102-3003021310122113-0002122013323233-1302113121220031-1110220311213120"></a>

#### `enable_api_discovery.discovered_api_settings.purge_duration_for_inactive_discovered_apis` property

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-0021233123301122-1013002210302002-1231011123110332-0101131021133112-1321303230003032-3032101231010222-0000133020332203-2201003101132010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.enable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-1330012322210220-2231101001022001-3233031123022210-0230202013013011-2230333011002231-3220333302031202-1002320112023320-1231321311011222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable learn from redirect traffic.

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
enable_learn_from_redirect_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_challenge

<a id="canonical-1201312312020123-3230111211203203-0323103120212310-2002022312313232-2200320302313111-0002122232302200-2233200202311231-0212233210001213"></a>

Type: `"object"`. single nested block, Optional.

Configure auto mitigation i.e risk based challenges for malicious users.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation")}
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
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

Terraform syntax:

```terraform
enable_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312320100311220-2121111023223010-2222123023231000-3231230201333232-2000321013003312-3122121312131333-2311101310003322-3032200103202113"></a>

### Direct properties for `enable_challenge`

- [captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012112303301110-3222223022111121-0322323321333132-0233230321301001-3011120230103011-1321123300033331-0013312223020331-0102200200231012): complete subsection reference.

- [default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-2212120120103301-3233031210133001-0010332200203330-3033001222120213-3112132200312113-1322100201132022-0200002302133322-0333302223330212): complete subsection reference.

- [default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121110313000012-3310321102200200-3113100211012323-3003132311310212-1230221132012011-1002211112133301-0031330011112203-0302213312133110): complete subsection reference.

- [default_mitigation_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-3003123010123300-3301100122332312-3002011132220020-2333111001110100-1212100130022031-2213210313032210-0121003122010333-3322202033313003): complete subsection reference.

- [js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-0312101012232110-0001011211003323-1030010110111113-0310302333212012-0111321322111020-3021022301102212-0232230030320121-0333232231023122): complete subsection reference.

- [malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-010.md#canonical-0323300102211113-0133120131332130-3221200230202231-0232201333102322-1131003112030231-0213300330113222-2103122033011023-1113122230202203): complete subsection reference.

<a id="canonical-1012112303301110-3222223022111121-0322323321333132-0233230321301001-3011120230103011-1321123300033331-0013312223020331-0102200200231012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-0211123311231020-0003112120310203-2220013131110110-0021333120132010-1232202233030213-1131200301211130-3001200210101112-0323130031020303"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
captcha_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022331321200331-2101333111332103-3211120120210111-3113020131111101-1020312332331130-1213113123003213-3122021112223312-3121002300110012"></a>

### Direct properties for `enable_challenge.captcha_challenge_parameters`

<a id="canonical-0113232331333331-3301201121111010-3210223121223023-3111113102333310-0313002211120202-0033130020101121-2320101211323120-0111223312232301"></a>

#### `enable_challenge.captcha_challenge_parameters.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3311131020313213-2132233320310202-2310102223101331-2332231120032002-1013210123311021-1133200320131330-1213200133232303-0011012332001002"></a>

<a id="canonical-2112000300021300-3033223223201021-1110032013133101-2223023020331232-1131322211233020-2221102003031030-3323132100031320-0232001211311013"></a>

#### `enable_challenge.captcha_challenge_parameters.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2212120120103301-3233031210133001-0010332200203330-3033001222120213-3112132200312113-1322100201132022-0200002302133322-0333302223330212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-3333113032103032-1001030133122233-0000132133113311-3120220111203201-0123211100233312-3311212113122220-2201021103002213-2021333032202220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default captcha challenge parameters.

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
default_captcha_challenge_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121110313000012-3310321102200200-3113100211012323-3003132311310212-1230221132012011-1002211112133301-0031330011112203-0302213312133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-1210331103002233-1230202212223322-0232231320321303-3010212202200130-3332331010221113-1011112221033100-3010101032032100-0011013032203001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default js challenge parameters.

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
default_js_challenge_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003123010123300-3301100122332312-3002011132220020-2333111001110100-1212100130022031-2213210313032210-0121003122010333-3322202033313003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_mitigation_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.default_mitigation_settings

<a id="canonical-0300221213130223-0123223121003201-0322310232332121-0202121230200012-2321332300132320-0223302031303011-2223230010221233-3011330222120330"></a>

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
default_mitigation_settings = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312101012232110-0001011211003323-1030010110111113-0310302333212012-0111321322111020-3021022301102212-0232230030320121-0333232231023122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.js_challenge_parameters

<a id="canonical-0323203300302101-1222220231002133-1303201020320013-0122301013313202-1231200110111332-2233103301103001-2123203121202012-3010322110322023"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
js_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202103312023101-2330031321110031-2331000030310023-0233233313320020-3102213330000302-2001200210100022-3001212222000230-0330031121220110"></a>

### Direct properties for `enable_challenge.js_challenge_parameters`

<a id="canonical-0033003031323131-2020101113101303-1303311130110220-2331313032022201-3310330101322221-1013302110103030-2220012103001131-3020213102210113"></a>

#### `enable_challenge.js_challenge_parameters.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2331310112023121-3121123302133023-2311110113133302-2100311211030323-0313310322203231-3220311320200112-3300131121012320-1122202120330202"></a>

<a id="canonical-2301000310332320-3313203212013322-3012332102102101-2330020211033202-2322230222330313-2303213003023010-0321323333203030-0220010133311100"></a>

#### `enable_challenge.js_challenge_parameters.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0202132112111103-3231121320123332-0013222333311010-2020101013220303-0230203203120232-3103212103230013-1213132122223213-2233212311210300"></a>

<a id="canonical-1310230120221300-0011231200200121-2100012232320232-3111020112312102-0312132220221310-0101122022313332-0132221102302200-2301033331022203"></a>

#### `enable_challenge.js_challenge_parameters.js_script_delay` property

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0323300102211113-0133120131332130-3221200230202231-0232201333102322-1131003112030231-0213300330113222-2103122033011023-1113122230202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.malicious_user_mitigation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.malicious_user_mitigation

<a id="canonical-3011302213322302-0112111011213203-2211103221013111-1102302231012201-2103101231221103-3211100203310131-1122133201203103-3020123030303130"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011030003301312-0323131112213330-1232013330203320-0223310012003231-2311010312330121-1213202330311231-1032001111220331-0212332130012303"></a>

### Direct properties for `enable_challenge.malicious_user_mitigation`

<a id="canonical-3230211000302300-2312302011233301-1233333123111232-0001122110132201-0231133230330302-0112222301013123-1102121201133032-3232220323103333"></a>

#### `enable_challenge.malicious_user_mitigation.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3122013203130123-3220023132222120-1101033220011321-0312333222120001-3201133220312033-1221213222320332-1221332331021321-0132312333121211"></a>

<a id="canonical-3202031212211010-1000310230001220-1222320033110102-1120112023332202-2310133002110323-1332120232200302-0213033100311013-2232331031012221"></a>

#### `enable_challenge.malicious_user_mitigation.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2322103131322131-0122312130133303-0300133222030010-1000211333211300-2121310322330221-1113023202220030-2202123212232232-1101202220020330"></a>

<a id="canonical-3000321313331123-3233131101110002-2022100221330130-2100333101131323-3023001223213330-0332320213100300-3122010313131301-1120213111323213"></a>

#### `enable_challenge.malicious_user_mitigation.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3132020110023330-0221123212220300-2303122132103021-2301032210323100-0112321332310001-2233030031121302-2100003011131031-0321200323303122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ip_reputation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_ip_reputation

<a id="canonical-1123223010010320-2010330211102303-3003132133301111-3013321110233132-3220003300022030-2113002030033003-3221112313032202-3320311130101320"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List. List of IP threat categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
enable_ip_reputation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1032321122230201-3300312212102222-2111201333313311-1200211223033002-0303130300310102-3012200321100321-3303211021120200-1023322223303132"></a>

### Direct properties for `enable_ip_reputation`

<a id="canonical-2222233221300223-2333220231013132-2233000201012203-2203002130020121-2110220333122330-3013100100110121-3121002111113221-3113320200211320"></a>

#### `enable_ip_reputation.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1303213031102030-3122302223322303-3002003203030013-0131022002010001-0002312200302222-3331123211021133-0230132330122210-3301210122110300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_malicious_user_detection

<a id="canonical-1123232203103010-2221320123213112-1303032131321013-2032233321303102-1212120310102022-0332230300221210-2020311101222333-2212332330213300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable malicious user detection.

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
enable_malicious_user_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301230133222123-3331320000010320-3300322323323202-1103213233103032-2103000031113100-2212030000211223-1213133100220200-0023212123220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_threat_mesh` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_threat_mesh

<a id="canonical-2123213233303030-1130113101312333-0212210331021330-3103233332320113-1210122212022213-2121031100211332-1031010102222213-2210032031223013"></a>

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
enable_threat_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- graphql_rules

<a id="canonical-3033113222332331-3132301221000332-2212230233231332-2330333210220301-0322131001130303-1131320103233222-1330301211123330-2232113032230103"></a>

Type: `"object"`. list nested block, Optional.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("exact_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("method_get",
    "method_post")}
```

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
graphql_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222132010310333-3320233301012221-2020133023031332-0102123222032003-2220313210230110-2102203012301223-3033102130122113-0311302333213312"></a>

### Direct properties for `graphql_rules`

- [any_domain](resources--cdn_loadbalancer--reference--group-010.md#canonical-0310312333200300-1001133112201301-2231012213003033-1211231123311203-0233000022332111-0012232012120121-1100033002132011-3333220110033032): complete subsection reference.

<a id="canonical-2112013200300110-3201131023132300-0102210030000203-0033322220111131-1213310331123121-0000301132321131-1232033122011130-3020003032010123"></a>

<a id="canonical-0032322322232311-1121101200302211-3120110320323313-3233031121122133-1102213020323321-1322333312032123-0021201210002113-3102230313032030"></a>

#### `graphql_rules.exact_path` property

Type: `"string"`. Optional.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Additional upstream details:

Default value is /GraphQL.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1100232233221102-3302033233011201-2230220112220103-2301130031231210-3322031300000320-0300022123221231-0333320021011211-3210112013322302"></a>

<a id="canonical-0102021211022121-2100310030312022-2302231322120111-1220112231223012-2111110103102221-1232212332121222-2301313032223300-1102221313203320"></a>

#### `graphql_rules.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [graphql_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-0303202121321220-1213302012121122-3010013331001221-0330202211132123-3330213012212101-3031001203002213-1113001320220013-0230201221030331): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-010.md#canonical-2111103032113221-2300330331032303-3021202023221012-1300020310021031-1001113031002312-0122031211203333-2111112020000120-2203030003322211): complete subsection reference.

- [method_get](resources--cdn_loadbalancer--reference--group-010.md#canonical-3112311302031310-2013010311020001-1012213233002331-3203312232102320-3332222023031121-2113023320012123-3003232203122311-3020001112232321): complete subsection reference.

- [method_post](resources--cdn_loadbalancer--reference--group-010.md#canonical-3101201020133222-2123332031201300-1233023232013013-3010301220131123-0000011132001021-1031001333131330-3303102132212030-1201311020130133): complete subsection reference.

<a id="canonical-3332221223013123-2033322221131331-3312213031321211-1013303313313001-2210013302100121-3132133122013122-2103110032030032-0321001312300003"></a>

<a id="canonical-0333313022001313-2023033012100312-3220102223121123-2001030323213122-0210330102231133-2021012012121320-2121210332123200-3330222312300211"></a>

#### `graphql_rules.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0310312333200300-1001133112201301-2231012213003033-1211231123311203-0233000022332111-0012232012120121-1100033002132011-3333220110033032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.any_domain

<a id="canonical-3001123003101213-2300323003102000-3130023130112210-2021233022011000-0110300003013222-3130322203032302-0231113220032230-2021130123001122"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303202121321220-1213302012121122-3010013331001221-0330202211132123-3330213012212101-3031001203002213-1113001320220013-0230201221030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.graphql_settings

<a id="canonical-0011113113220331-3333032022312102-1203133231233131-0132130130020232-2033011100101010-1313123013001231-3113103022230233-3130311100311331"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for GraphQL settings.

Additional upstream details:

GraphQL configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("max_batched_queries",
    "max_depth",
    "max_total_length"),
  validators.ConflictingObjectAttributes("disable_introspection",
    "enable_introspection")}
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
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

Terraform syntax:

```terraform
graphql_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202123010132002-3230320101122131-2021220211031013-3032201201231013-0231110210130131-0313222033300302-2102122032320221-1313331220210000"></a>

### Direct properties for `graphql_rules.graphql_settings`

- [disable_introspection](resources--cdn_loadbalancer--reference--group-010.md#canonical-3011202003003032-0001310103133331-0111311023220230-3202133313001031-2011303203032233-1303320103110121-1232112230112301-2301230212301233): complete subsection reference.

- [enable_introspection](resources--cdn_loadbalancer--reference--group-010.md#canonical-3323021120022333-1030001100012121-1030112033113022-0232222010022320-2122131332132002-3333121333003323-1202022302200200-1122220120032310): complete subsection reference.

<a id="canonical-3211220113030133-1200213123221232-0301220313213310-1310132222230210-0301101031200303-0102321203230220-0111320320121122-3312303312121320"></a>

<a id="canonical-3013101313013111-3332131033012111-3202322200212003-2211203230331033-3312131123113002-0122302031023113-1332330211111111-2003323021002323"></a>

#### `graphql_rules.graphql_settings.max_batched_queries` property

Type: `"number"`. Optional.

Specify maximum number of queries in a single batched request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-3213210311113222-2022312021003221-1300232213113101-2121303333001020-0320233232310022-0312101122203023-3131001222133123-1220211033100111"></a>

<a id="canonical-2002020312110111-0233123220103301-2022313022010013-1300101330033221-2022300223203200-0203302231103213-2103012222031110-2132123300003202"></a>

#### `graphql_rules.graphql_settings.max_depth` property

Type: `"number"`. Optional.

Specify maximum depth for the GraphQL query.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-0130230321231113-0331011323202233-2300320012123112-3101221221232012-0322003031303111-0330322010211212-0231010312200011-3122323313231011"></a>

<a id="canonical-1113310200103310-1202032301311101-0010032301022310-3020010322033132-0113121120011213-0003031203312223-1100101022123021-1123002211310211"></a>

#### `graphql_rules.graphql_settings.max_total_length` property

Type: `"number"`. Optional.

Specify maximum length in bytes for the GraphQL query.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 16386),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

<a id="canonical-3011202003003032-0001310103133331-0111311023220230-3202133313001031-2011303203032233-1303320103110121-1232112230112301-2301230212301233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.disable_introspection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- [graphql_rules.graphql_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-0303202121321220-1213302012121122-3010013331001221-0330202211132123-3330213012212101-3031001203002213-1113001320220013-0230201221030331)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-3101003223213131-0302023333100020-3330000313333233-1101213203002303-3202311323123010-3213133122001112-0332000000230023-2230222303323323"></a>

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
disable_introspection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323021120022333-1030001100012121-1030112033113022-0232222010022320-2122131332132002-3333121333003323-1202022302200200-1122220120032310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.enable_introspection` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- [graphql_rules.graphql_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-0303202121321220-1213302012121122-3010013331001221-0330202211132123-3330213012212101-3031001203002213-1113001320220013-0230201221030331)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-0112123030000133-3321231320123131-2303313110122003-3212213110203222-0101330120102132-0323102103232000-1113023320101333-1320031313001232"></a>

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
enable_introspection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111103032113221-2300330331032303-3021202023221012-1300020310021031-1001113031002312-0122031211203333-2111112020000120-2203030003322211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.metadata

<a id="canonical-2330312010203330-1012110013030332-1210331323020011-2231112203202213-0313123323013212-0130110311103332-2323332300321223-2022120113003213"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003123030230220-3012131020211133-0021000130002000-3300233302222322-3103221233131303-0030330023233123-2313322011011221-0103330023133010"></a>

### Direct properties for `graphql_rules.metadata`

<a id="canonical-2100133331032012-1032110132320011-3010120332330000-1121010332300111-0000103310321313-3111031131201030-1133213013323323-1012303231321321"></a>

#### `graphql_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2302112022103202-0000133103102303-3133221132130110-0100331212332000-3231332122020321-1322131111211312-1032133022010021-1102101120111303"></a>

<a id="canonical-1223222310133223-0121121213331201-0133332230330200-3100011002213010-0003110122013021-2221030133132132-1222111031322032-2002122132322133"></a>

#### `graphql_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3112311302031310-2013010311020001-1012213233002331-3203312232102320-3332222023031121-2113023320012123-3003232203122311-3020001112232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_get` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.method_get

<a id="canonical-3222233221101112-0202132120001201-1201222112323102-1202103122220113-3331222200313322-3003110310113022-2122010001233211-0321112111030303"></a>

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
method_get = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101201020133222-2123332031201300-1233023232013013-3010301220131123-0000011132001021-1031001333131330-3303102132212030-1201311020130133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_post` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223)
- graphql_rules.method_post

<a id="canonical-0013230202323121-1013332121001311-1321301023201110-1233301102312201-0202210313231333-2020101200110101-3203000113012012-3323102110321022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for method post.

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
method_post = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320012221332001-3023010200311133-0012202111211311-2312132120132210-3032223212022123-1212221020223222-2203232202220232-2203232310011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- http

<a id="canonical-1000121033311203-2210131322032220-3212323011131021-0001030112113101-0100303223210201-2212000012230011-3130012011101013-1133213221020323"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: http, https, https\_auto\_cert; Default: https\_auto\_cert\] HTTP Choice. Choice for
selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

OneOf alternatives in this subsection:

- [http](resources--cdn_loadbalancer--reference--group-010.md#canonical-1000121033311203-2210131322032220-3212323011131021-0001030112113101-0100303223210201-2212000012230011-3130012011101013-1133213221020323)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-0030210123302230-2132331033300110-1110303122211311-0023313300030000-1100221102102222-3331022302312113-0311203102112002-0232201201033012)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-011.md#canonical-1011001132000320-2032211231022312-1103022230201012-2202332311121101-0230203310001320-2210023110030310-2011102022223002-0133123021300313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110203032210122-1230103020310123-3322223110021200-3033232232110221-0312003020212010-1110130300322120-0230020222021330-2230331323000312"></a>

### Direct properties for `http`

<a id="canonical-0303010032021311-0113113123111312-1302003130011322-1230013312022021-1112001233233123-2210102012323213-3121011010311200-2212220131131323"></a>

#### `http.dns_volterra_managed` property

Type: `"bool"`. Optional.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311213121120331-0202110020301001-3231322230232210-0000100021323121-0031311021322332-0122010200112113-2311321303230103-2332030002221203"></a>

<a id="canonical-0200301131222331-2203030001230130-1113031030201120-2222112301002211-2333220033012001-1200211013133213-1233003011330030-0330223130010011"></a>

#### `http.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2010120000113220-1211130202013023-0201023113222223-3200110231021211-0210331202213222-1001011201001233-3202201031300000-0002023202212321"></a>

<a id="canonical-2211313223213231-0111232221031322-3003203212233110-0013203023002021-1002023222112203-3211200222332323-3031202100223022-3220023123120110"></a>

#### `http.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- https

<a id="canonical-0030210123302230-2132331033300110-1110303122211311-0023313300030000-1100221102102222-3331022302312113-0311203102112002-0232201201033012"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting CDN Distribution with bring your own certificates.

Receipt-pinned upstream constraints:

```json
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
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131222220220130-0003323101031201-1010302311020032-3320001010113301-1302311313223323-1313301103000221-0000101230220210-1303232120303122"></a>

### Direct properties for `https`

<a id="canonical-0133031133212130-1321032323001130-2233001313313310-2021301213211221-3010013102130311-0232123212130322-1200123303330023-2121223101130300"></a>

#### `https.add_hsts` property

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2331220202202202-3200110211012320-0122210103123013-1320121002030232-2130332230313130-0210020122313110-1323332113220011-1212031013210021"></a>

<a id="canonical-2003023323121320-3020311021020013-2110221002131212-3012232133013110-3023222210311113-3312022202202311-1212222211011002-0020100333202110"></a>

#### `https.http_redirect` property

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132): complete subsection reference.

<a id="canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- https.tls_cert_options

<a id="canonical-2002111312101220-1033231311000323-0002301330212330-1030113130123120-3220222031022023-1230233233211020-1022030031103100-0203200201210312"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert options.

Additional upstream details:

TLS Certificate OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_inline_params")}
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
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_inline_params\"]"
}
```

Terraform syntax:

```terraform
tls_cert_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002311210321321-1333231312332311-3210001222030203-0021213111231210-1232130013113323-3220013230223330-2222323310231330-3221311111121113"></a>

### Direct properties for `https.tls_cert_options`

- [tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101): complete subsection reference.

- [tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-0011130110103122-2323202003301013-2232002030333012-1023003212123231-1321201123320103-0032223033201200-2100133021322003-2300130211311100): complete subsection reference.

<a id="canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- https.tls_cert_options.tls_cert_params

<a id="canonical-3313331331032322-0103010213220221-3330121302020311-2110202321230310-3330110111331131-1030310222121120-2322013030000303-1111032312133023"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003113022313323-2210031312003010-3303000021003131-0013332230110220-0312113203133120-2322303101313301-0313222233132310-2111123302300333"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params`

- [certificates](resources--cdn_loadbalancer--reference--group-010.md#canonical-1020021333030211-3011200000212023-3210031201010321-0210213122231010-2022223303302133-3132210020202113-3223321030111210-3022032111122322): complete subsection reference.

- [no_mtls](resources--cdn_loadbalancer--reference--group-010.md#canonical-3021200213311112-0032113010312301-0030303030133012-1233012221331023-2133330311113211-2022022021310210-1112303023200211-2203030130332233): complete subsection reference.

- [tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120): complete subsection reference.

- [use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-2013031123320130-3023100003003001-0130010131231023-3100210112220023-0231230020210122-1000131313031002-0111001332022033-0102232331322320): complete subsection reference.

<a id="canonical-1020021333030211-3011200000212023-3210031201010321-0210213122231010-2022223303302133-3132210020202113-3223321030111210-3022032111122322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- https.tls_cert_options.tls_cert_params.certificates

<a id="canonical-2000100222122210-2202232030131201-1000203201201132-3003012121312033-1222211223203113-0031332231032230-3113323121212131-0001030300122310"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102032231110311-1212313012113130-2333121102013233-1300123330233021-0332002123102020-2321210132322002-3312011112222132-2123302023202203"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.certificates`

<a id="canonical-1012312010133120-2211131110100222-0013010023101020-3301211020023311-3033231002200100-0131123321312132-3123303221130223-3112012312023121"></a>

#### `https.tls_cert_options.tls_cert_params.certificates.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0300303022131203-1030100000300200-3220023200103311-0011123101303320-0012123200030330-0012302102020020-2310023012213131-1321023300123332"></a>

<a id="canonical-1223201000333233-1011023330012300-2302110212100302-3221101323121110-1003330203210332-3303012120213221-0132001221223230-2302110012232020"></a>

#### `https.tls_cert_options.tls_cert_params.certificates.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1111033010223022-3030333331231113-1013202231202022-2320201220323310-0220223121210233-2002302101313301-1011030111130210-0011132121022213"></a>

<a id="canonical-1020131101221220-1113012032022301-2122221020120001-3221323121113102-0330013110213111-1311212102103330-1303220313100323-2321012201033103"></a>

#### `https.tls_cert_options.tls_cert_params.certificates.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3021200213311112-0032113010312301-0030303030133012-1233012221331023-2133330311113211-2022022021310210-1112303023200211-2203030130332233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- https.tls_cert_options.tls_cert_params.no_mtls

<a id="canonical-3020222203232310-2012203131200230-1101313203221113-2012023111133011-1301113110223212-1223211001030323-1333310023003202-0331222302311012"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- https.tls_cert_options.tls_cert_params.tls_config

<a id="canonical-0100013130032010-0121010333130320-2133133111121102-3222011102303210-0000002031010032-0222303110302223-0121303010131223-3212231332331300"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202020102313333-1312111002302332-0133203323100123-0032100020201012-3000333130311113-0121311102000031-3331223332102300-2010322312201201"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.tls_config`

- [custom_security](resources--cdn_loadbalancer--reference--group-010.md#canonical-0323322323121333-1303311220212232-0223201133021022-1020011012312020-3312210303211202-3032220122202331-2312021023002111-1113232330001033): complete subsection reference.

- [default_security](resources--cdn_loadbalancer--reference--group-010.md#canonical-0120001133011022-3023333001330231-2022102023023132-0211132313213033-3031200311200322-0030122132223101-3202221322232221-3200120021201321): complete subsection reference.

- [low_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-3201130202012103-1300201123010231-0032201013002213-0220332212233302-3312112202322023-1002010112222303-1333323330101230-0111100031002332): complete subsection reference.

- [medium_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-2330012003201103-1033011111321030-3002023213331032-0223023310011123-1000213320323321-3203323130313132-2200211113332020-3200331311112110): complete subsection reference.

<a id="canonical-0323322323121333-1303311220212232-0223201133021022-1020011012312020-3312210303211202-3032220122202331-2312021023002111-1113232330001033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120)
- https.tls_cert_options.tls_cert_params.tls_config.custom_security

<a id="canonical-3112223131321101-1032002121311233-2310202330300110-3133303211203100-0211330033321103-1321333200012320-3132320321232001-1003223301120222"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303120023011001-2333320110111233-1212010313213023-2232023311121030-1032302100302130-3312231112110200-0230000221010031-3113102312013112"></a>

### Direct properties for `https.tls_cert_options.tls_cert_params.tls_config.custom_security`

<a id="canonical-0100332021321112-2330303211101120-1232302123021023-2212301301002122-2212111331331233-2312201300111333-0021233210200030-2002000131203221"></a>

#### `https.tls_cert_options.tls_cert_params.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1230300131320233-2311222221022223-2021320003010033-1331333003212101-0223232113002231-3111312332313212-0222322101111112-2211020210330132"></a>

<a id="canonical-3011000300130100-2212232001133310-0331203321221221-3213303033120011-0333330012202201-0020012000100230-2001202332210130-0003130132100103"></a>

#### `https.tls_cert_options.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0102032011131330-1013200231030311-2333302221222331-1000311020020320-0020032031131333-1210211322012020-1111322020320120-0300021110230201"></a>

<a id="canonical-1123113302302222-3103010111212002-0112112211221022-1320302213300311-0222031012112202-3112333320123233-0200033101323001-2230311011031030"></a>

#### `https.tls_cert_options.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120001133011022-3023333001330231-2022102023023132-0211132313213033-3031200311200322-0030122132223101-3202221322232221-3200120021201321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_options.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-0333210302232102-3132301301101323-1211333333001031-0100031110002301-1231020320311323-2013321101333113-2121323320223113-3103033031222132)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-3322313333000323-2110320030113113-1003100031023231-1120300133023120-2203322010232311-1313303213230301-3320003002011211-1312111301013101)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-1031331122100021-3131111330111203-2333333132031233-2130331021100120-1232131312311320-0120031130123220-1013301010302110-3333102021313120)
- https.tls_cert_options.tls_cert_params.tls_config.default_security

<a id="canonical-0312012022311011-3132222122030120-2230200111130023-2112023102112003-1022220122320020-2011003210122210-3010211202221120-2112020110101211"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.
