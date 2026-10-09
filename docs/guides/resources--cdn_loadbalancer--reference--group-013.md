---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3223020032100323-0020200120223232-1323222300022131-0322111320311332-1131023131031221-1022311102320100-3210103133123110-2321012303123013"></a>

## `other_settings.header_options.request_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213): complete subsection reference.

<a id="canonical-3022011203323013-1301002301113101-1332330300221333-0222202010333310-1212222013112030-1000122303010103-2332321021103213-2010220131330323"></a>

<a id="canonical-1110113130130221-0230323120222211-3323211213030011-0312202131032122-1323320013200003-2123120011331010-3221111210122331-0030231122222023"></a>

## `other_settings.header_options.request_headers_to_add.value` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202)
- other_settings.header_options.request_headers_to_add.secret_value

<a id="canonical-3331010111222222-2100013220011320-2122101320302311-0110213332211311-1111322101222101-0132233212312033-3111312213220120-2110312032223000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323210230033302-2211221211332220-3110221112030300-0112221122332112-2012213130221121-2102120230002323-1130201310312320-2000021312102120"></a>

### Direct properties for `other_settings.header_options.request_headers_to_add.secret_value`

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-2331223000110010-0323233123202112-0231333012102230-1330312202120211-1230301032201200-1331220133030230-3103220021230133-0133003221223221): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-3313023202202302-3320212100333032-1100112130203220-1121213002122120-3322233331333121-3020330101330211-2101201221231103-0100103002332111): complete subsection reference.

<a id="canonical-2331223000110010-0323233123202112-0231333012102230-1330312202120211-1230301032201200-1331220133030230-3103220021230133-0133003221223221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202)
- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213)
- other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1113111123122232-1213133021302302-2133333120231312-0132301220321032-3013212233213313-3111322110011331-3121333002232110-3131313000300011"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332130131221001-2213132200303202-1221102230321321-1131100133322120-2232000102000312-1301001012211330-2131320231001302-1313021323131232"></a>

### Direct properties for `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-2110220223133230-0011202123303000-2311030122331121-0323222032222233-0010003031110333-3210102311210001-3111231233110302-3110323012100013"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2331232003201122-0021023020120103-0211123221313311-0102111312102301-2010101022130211-1213330333300103-3300333102221321-0020122233001112"></a>

<a id="canonical-2332120303110303-1002201231011222-2001003301113310-0311200222120100-3102032311223320-0011102302200032-2231202002003201-1000333112232232"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2101301211123232-0313013013012323-3003032331301032-2312332231002332-3230310222101313-2223203010001113-1300232210322212-3120011102110201"></a>

<a id="canonical-2030032311301100-0302321101303011-0103301113133210-0021130012121000-3332031111110200-3102002013030130-0120322012331022-0310110312233322"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3313023202202302-3320212100333032-1100112130203220-1121213002122120-3322233331333121-3020330101330211-2101201221231103-0100103002332111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-012.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202)
- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213)
- other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3011233032130321-2103211302133313-1121333010232200-0331000123300032-3032001203220201-2020031003223200-1121001222331130-0033011113220120"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222223021320011-3231131020012232-0023133120113210-0001103232122021-1333232112030003-0330323300022132-2232012011221013-2102213120303132"></a>

### Direct properties for `other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1212331233322232-1300113101321020-1023102103211302-0013031300031210-0230011131323323-1313331211313013-0013030031221121-2031202113320130"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3001032310102333-2323310302012001-0303032302012311-3313211303002100-2313020013233122-2011201211032211-2010332110312223-2001303103102103"></a>

<a id="canonical-2200333123210303-3012213231303121-1320201031111000-0021201212123101-3103330022030331-3212312113101122-0122023211333002-3202300122210202"></a>

#### `other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.response_headers_to_add` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- other_settings.header_options.response_headers_to_add

<a id="canonical-2100133013302121-2012033220023302-0222002231313313-2222131312032021-0221203003023003-0113333330322123-0032133230100331-3212231220030220"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203332003120211-2302131212033032-2023300133133321-3211331203121311-3011012100213112-3122212231030102-1311302011300003-3212230231202131"></a>

### Direct properties for `other_settings.header_options.response_headers_to_add`

<a id="canonical-0120202310300101-0003121230201023-1221133000020113-1331000001103222-3330103103110313-1231122001223133-0123100331011333-0133133230303032"></a>

#### `other_settings.header_options.response_headers_to_add.append` property

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3123220211030303-1333013322323212-2230220310201130-2010123323030021-3121210111222022-2303220301221113-2311110121030210-1031120001330010"></a>

<a id="canonical-2131022111231021-1231130102131001-0310201321301001-3003033032112311-3123230133110331-3113201120213332-3313021003201332-1301021203213023"></a>

#### `other_settings.header_options.response_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122): complete subsection reference.

<a id="canonical-1300311232031203-2021130212202333-3201330123101231-1221002013231313-2203230031333113-2022000331120002-2233321231300131-2303130100303313"></a>

<a id="canonical-1013101221110112-2000110013022122-2332200203130300-3300130031330322-3202220233212331-1132100023213311-0101313010330033-1103133210331132"></a>

#### `other_settings.header_options.response_headers_to_add.value` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322)
- other_settings.header_options.response_headers_to_add.secret_value

<a id="canonical-2110103003010301-3213300103203013-0023222012012030-2221033001311233-3202021212133233-0000113232213302-0033300213010103-2002032301113121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021212121001000-1013332021101101-3322202003011102-3221212000331323-2122301232103102-1022333120111222-0300200010303101-0101323202130323"></a>

### Direct properties for `other_settings.header_options.response_headers_to_add.secret_value`

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-0110121132222012-0300301023333333-3030111211200313-3011123032013123-0313010311023311-1223212111130233-3210022201020223-2022201103322320): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-0103203311320333-0012221221223323-2121322202011033-0032311010003313-3200323120200113-0132231003130100-3100310000113203-2231110011203232): complete subsection reference.

<a id="canonical-0110121132222012-0300301023333333-3030111211200313-3011123032013123-0313010311023311-1223212111130233-3210022201020223-2022201103322320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322)
- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122)
- other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0112332002200010-2023311110302211-1023220011102331-1232022223321312-3000210322011320-3203102313330122-3022102113313323-1213203201112232"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322022130320100-0201320031013001-1210232301301213-0311200300111100-1102200320102300-2331012020232302-3002101111003131-1203011022220110"></a>

### Direct properties for `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1202312122323002-1130120131102031-3230200302102232-0122300131002303-2302012202231021-2010311001333332-1222332020232123-3031200031320011"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0221323133101313-1032012023303301-2201002302111022-1112012211213120-0201111000331003-0132122013112030-3200233311202213-2103300133021203"></a>

<a id="canonical-1123222203230110-1221331230022012-3320300002220332-0202330020201303-2222302020022032-2130013133312120-2322210011212122-3233231010212003"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1233030330020320-1111110211131311-2001220132001323-3123121001031131-0121101203133220-3112223033133333-1332321321102223-1300201112212000"></a>

<a id="canonical-3202213333310100-3211321013310302-1133231333231320-1202013021221101-3220010032133132-0010130010123321-0132212012230230-1100111300121321"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0103203311320333-0012221221223323-2121322202011033-0032311010003313-3200323120200113-0132231003130100-3100310000113203-2231110011203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322)
- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122)
- other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2311133021320023-2133001231310022-2322131301232223-3102331332221301-0033322213123233-3033110132230322-2113032323032233-1113212213331132"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210001323220021-1003003212223003-0131212100301320-3312221011020201-1220232332021232-2021330023122200-0000130023323322-2111223011223021"></a>

### Direct properties for `other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1301331330001002-0131110022123002-2113320200033002-3303022303202332-1312333011103131-2202210111221003-0303220112223203-0031110130022231"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1113132002110103-2330122303201312-2323313001131121-0123032213110122-2003013312012023-1110030021203033-0133233023012120-1112301002112133"></a>

<a id="canonical-3100213001301130-1202202332231231-2033220222000333-2211011103011112-0120113301100223-2313311213301030-2202312131000101-2333211113023330"></a>

#### `other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.logging_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- other_settings.logging_options

<a id="canonical-2033033311310033-3112233220000212-1322020022231103-2131003221022103-0311112120332210-2131300130332320-2130303332121013-0233030022123001"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS related to logging.

Receipt-pinned upstream constraints:

```json
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
logging_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233311130111113-2212233223031111-3031032211223033-0222103101200222-3201101310313303-0213010203013302-2120023222101301-2102211321123202"></a>

### Direct properties for `other_settings.logging_options`

- [client_log_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-3100220303110020-0030111031001233-0010210000322313-1211310232011312-0233010320311301-2000102113132210-1320131102010223-3123110102211210): complete subsection reference.

- [origin_log_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-1101322013323301-2223110001000120-2120123233101121-3222003111233313-0333012210002223-1032223320210211-1130233103030323-1200113031130213): complete subsection reference.

<a id="canonical-3100220303110020-0030111031001233-0010210000322313-1211310232011312-0233010320311301-2000102113132210-1320131102010223-3123110102211210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.logging_options.client_log_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130)
- other_settings.logging_options.client_log_options

<a id="canonical-3020233031032001-3201110011132302-2201300010020331-1023202122110123-3000102123102133-3120020003022223-0120030233213111-0031033322312100"></a>

Type: `"object"`. single nested block, Optional.

Headers to Log. List of headers to Log.

Receipt-pinned upstream constraints:

```json
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
client_log_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300122323023333-2013222212333023-3221113110211210-1223213022200200-0231302323001212-1210010203212211-2003221101001321-0301132330030232"></a>

### Direct properties for `other_settings.logging_options.client_log_options`

<a id="canonical-1233012030130202-1023102102231022-0222222001110020-1030012310332100-2233223020203103-2322331320213013-3223013301010210-1012321232033213"></a>

#### `other_settings.logging_options.client_log_options.header_list` property

Type: `["list", "string"]`. Optional.

Headers. List of headers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1101322013323301-2223110001000120-2120123233101121-3222003111233313-0333012210002223-1032223320210211-1130233103030323-1200113031130213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.logging_options.origin_log_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130)
- other_settings.logging_options.origin_log_options

<a id="canonical-3112203131202231-2220133232103133-0221031232322333-0321132211002021-1221101331011213-1210003033123122-1213333031301302-2323211021031001"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin log options.

Additional upstream details:

List of headers to Log.

Receipt-pinned upstream constraints:

```json
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
origin_log_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000101000312332-1121010332022021-0233323223230020-1310111333220131-0100031202103213-0130320021223001-1320220130002020-2302312020130020"></a>

### Direct properties for `other_settings.logging_options.origin_log_options`

<a id="canonical-1313310301020133-1130021212132030-0203232122021213-1212013300001012-3323120222003320-3223011301032120-2201120322302300-0001003300322303"></a>

#### `other_settings.logging_options.origin_log_options.header_list` property

Type: `["list", "string"]`. Optional.

Headers. List of headers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- policy_based_challenge

<a id="canonical-3110233220333200-3212201333323201-2020113221230221-3322221202220102-2132220310202212-2231113002233113-3013031220202301-3203221133222020"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
policy_based_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133221103112111-0111132110110330-0331313021013031-0301110310300000-2023223112010321-1201233010121020-2123311002133211-2122301222131112"></a>

### Direct properties for `policy_based_challenge`

- [always_enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3231100132202302-1113312322120131-1133222211000302-2032310321101021-3323031100331200-0021213230032033-0021220113101101-0303301220221031): complete subsection reference.

- [always_enable_js_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3220313211233110-3203030311210203-1321222221301101-1311012231310212-1002121131030213-1332100110321123-1313111220310010-3112031113010121): complete subsection reference.

- [captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-1323000010110021-0032203011323300-0232330032321021-3212103130130000-0310233331320031-1320030320020332-2123111310102012-2130301311021333): complete subsection reference.

- [default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-3121331210122303-0333010020032030-2311300133222010-3103013023220220-0032113202212121-1033021203220133-1020013323120333-3130220030122211): complete subsection reference.

- [default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-1221023230023032-0022213230332233-1213130311032013-3101122202311233-3211011101333210-3031100032112110-0303320212301221-2011302321120101): complete subsection reference.

- [default_mitigation_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-0120231001300001-3322103130120122-3301130030331000-1001202220203231-2211200202233330-0302211021201312-0303302302033012-1012010021111100): complete subsection reference.

- [default_temporary_blocking_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-2022330021132232-2102221101220110-1010002220220113-3323023011020013-0103130102102110-1122300013010123-0103121301130020-2113313130200320): complete subsection reference.

- [js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-2013120300230010-1203120210000023-2003001321022311-1201003231201221-1322132022103111-3103302022311213-3030323332012121-0321033313031002): complete subsection reference.

- [malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-013.md#canonical-2213030000002202-0023323001313300-0100232113302220-2320331033032120-2132100010013330-2111101121100221-1132222032102331-1203323123332212): complete subsection reference.

- [no_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213200121002130-3323213132112330-2113023033200123-1131211303231121-3200332321302331-1110000333012132-2102221010223301-2021200003100300): complete subsection reference.

- [rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301): complete subsection reference.

- [temporary_user_blocking](resources--cdn_loadbalancer--reference--group-014.md#canonical-2020113203001122-0131003133101010-3112033112103122-1230132210230022-1110101121230003-3011322320122320-2001301222311323-1030230313310012): complete subsection reference.

<a id="canonical-3231100132202302-1113312322120131-1133222211000302-2032310321101021-3323031100331200-0021213230032033-0021220113101101-0303301220221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.always_enable_captcha_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-1131120110030330-2121022011120212-3121010110212203-0301010000032110-0032320221300310-0030301302231031-3012202210112333-3030302201303101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable captcha challenge.

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
always_enable_captcha_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220313211233110-3203030311210203-1321222221301101-1311012231310212-1002121131030213-1332100110321123-1313111220310010-3112031113010121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.always_enable_js_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-0121203231131213-3120133332233320-3013211300022321-2330013301232131-1302320221010310-2003330202213123-3300031122023110-0210020201320300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable js challenge.

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
always_enable_js_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323000010110021-0032203011323300-0232330032321021-3212103130130000-0310233331320031-1320030320020332-2123111310102012-2130301311021333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-3210133222033210-0303032110030301-1220310011130010-2101212131112203-1031002120002012-2323301311022000-3330012232302323-1331233323321011"></a>

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

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3133320031321200-1112110310332011-2200320321113313-1312310021331303-1002311033313203-0320222213313331-2100102030233013-2233032233221022"></a>

### Direct properties for `policy_based_challenge.captcha_challenge_parameters`

<a id="canonical-2231213221002121-2010200023222031-3031012201121023-1120320321200333-0012331332330012-3033120230023331-0223222231210323-0011313222003122"></a>

#### `policy_based_challenge.captcha_challenge_parameters.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0013000233133012-0031121131213211-1110121003130030-0333123233233201-1203330222213111-1100100222211032-3021003330003200-2123102220031302"></a>

<a id="canonical-0010020001201113-1002130323111223-1313133032201233-3033113213230012-3000133301332133-0323331320121232-2013212130123321-1101300100312130"></a>

#### `policy_based_challenge.captcha_challenge_parameters.custom_page` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3121331210122303-0333010020032030-2311300133222010-3103013023220220-0032113202212121-1033021203220133-1020013323120333-3130220030122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-2333021013130001-3303131231231130-2201323131321121-0200202130213031-0310013030130321-2021201021100100-0133222302331131-0313322100303302"></a>

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

<a id="canonical-1221023230023032-0022213230332233-1213130311032013-3101122202311233-3211011101333210-3031100032112110-0303320212301221-2011302321120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-0131110301113233-0221003002212302-2220223023001221-1231013031200030-1232013322113113-2321033312311323-0210232111000321-3202101101221312"></a>

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

<a id="canonical-0120231001300001-3322103130120122-3301130030331000-1001202220203231-2211200202233330-0302211021201312-0303302302033012-1012010021111100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_mitigation_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-0032223011111213-1231211321302322-3022212312221033-2233032331033023-3230330232133203-3012230330131300-2231300122130213-0111223010331232"></a>

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

<a id="canonical-2022330021132232-2102221101220110-1010002220220113-3323023011020013-0103130102102110-1122300013010123-0103121301130020-2113313130200320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_temporary_blocking_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-3022111031011030-0313211220122230-3301102013121223-1311301021112031-2020222331330120-0001201313032101-1302230302201201-1012231113323200"></a>

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
default_temporary_blocking_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013120300230010-1203120210000023-2003001321022311-1201003231201221-1322132022103111-3103302022311213-3030323332012121-0321033313031002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-0033101233000012-0000301013201301-1230230132113303-1112320010000300-0331003323332021-2010023132211210-1322301300232330-0311000133201020"></a>

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

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2311310031033022-0220212220330010-1222013330303303-2301100000200003-2031233200031322-3220321331132000-1303232121323333-0033330003110230"></a>

### Direct properties for `policy_based_challenge.js_challenge_parameters`

<a id="canonical-2010021012230323-3330222231032332-1312110330211311-3231303020302113-1311000320133020-2110013300233200-2010322012313321-1000103031320311"></a>

#### `policy_based_challenge.js_challenge_parameters.cookie_expiry` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0100233011111322-2232013101131321-1331010202121101-0212333300122232-2311231221002032-3321123200323300-0312022112332301-0300002103122030"></a>

<a id="canonical-0301010313202210-0300201323200232-0321110223131132-3322132032230333-0013021233320101-0122100201001032-0302223310212112-3212030221103201"></a>

#### `policy_based_challenge.js_challenge_parameters.custom_page` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3233313101111033-2221120123131300-2032103233310322-1032130302030201-3023203302003111-3133223210012213-3321232230330221-1212323012220011"></a>

<a id="canonical-3131300301101111-3300002311202220-1321320021122002-1123211201321022-3003302322010002-3232313123020111-2110030302122230-3313323230000301"></a>

#### `policy_based_challenge.js_challenge_parameters.js_script_delay` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2213030000002202-0023323001313300-0100232113302220-2320331033032120-2132100010013330-2111101121100221-1132222032102331-1203323123332212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.malicious_user_mitigation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-2311303230123023-3132330022020112-1333202130031322-1022120030031201-3203003330100113-3111332102303323-0212333003200202-2032301211213002"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232223331020030-3033130013210010-0013123312221203-0123303330333210-3023023131232130-2201322200120221-1311112101320131-3213312000210300"></a>

### Direct properties for `policy_based_challenge.malicious_user_mitigation`

<a id="canonical-0210203220323123-0311121102223133-1100131231133212-3222010110021133-3313230010032303-2231322202230021-1010031212201220-3221121223310023"></a>

#### `policy_based_challenge.malicious_user_mitigation.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3013323212221331-2210322102103101-2330322322323021-0223101323321031-2231032210220213-2200332121131113-1033202003233022-2120031321233131"></a>

<a id="canonical-2023230330021010-3212100223131220-3021200232302130-1323200300011120-0112133103303022-0202313330130321-1102233012021120-2212100000023022"></a>

#### `policy_based_challenge.malicious_user_mitigation.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2103321000331011-2131302121010202-1303220331021330-2213100301223013-0220323120123332-2101200111023312-0020232132200123-0213133303030330"></a>

<a id="canonical-3331121222330330-0230211012133223-3013121230221102-2121310001002112-1121010323100101-1301311202030021-3101131111002332-2201231230101111"></a>

#### `policy_based_challenge.malicious_user_mitigation.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3213200121002130-3323213132112330-2113023033200123-1131211303231121-3200332321302331-1110000333012132-2102221010223301-2021200003100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.no_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.no_challenge

<a id="canonical-0203202320131112-2302320221210211-3212131320032120-0320111122010132-0330111000211332-1210110133233323-2201113130103030-2031022102122131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no challenge.

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
no_challenge = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.rule_list

<a id="canonical-2123222330010033-0203022221121212-2113002200232332-3211212012222233-0111101003133132-2031003031110011-3233303210233021-0211120333103132"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030230222011133-2133003030332211-2310120020002300-2300000000213102-1300031331322321-0023032320020123-1302121331321113-0303222030001011"></a>

### Direct properties for `policy_based_challenge.rule_list`

- [rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111): complete subsection reference.

<a id="canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- policy_based_challenge.rule_list.rules

<a id="canonical-1101021211020011-0030321102302103-3003000331310110-0123033123211030-0221003232023011-2230023210001323-2303323322303303-2012020013101331"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301131101132230-0110121133030203-1323313331022113-1031233130303033-2313021213331331-1120332312321121-3111121202101300-1133033313011300"></a>

### Direct properties for `policy_based_challenge.rule_list.rules`

- [metadata](resources--cdn_loadbalancer--reference--group-013.md#canonical-3312022112120312-2330122132101001-0023311121210013-0120002103000332-2202100021232030-3133210320123312-0311302202001302-0322131111030012): complete subsection reference.

- [spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201): complete subsection reference.

<a id="canonical-3312022112120312-2330122132101001-0023311121210013-0120002103000332-2202100021232030-3133210320123312-0311302202001302-0322131111030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-0101332312003220-0213301222131132-1301030213013320-1311201032302202-2113001233131031-3121021333332102-3301313213032331-3222220200332313"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332130100022203-0031311203102011-3220220000322220-1213013011013310-0323220112113312-2303100323003130-3130021231311121-3020223033120021"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.metadata`

<a id="canonical-1333021033320221-3120102123212323-3003122103122212-1220003233013002-1101032110013032-3103032131103002-0030303202211102-3232211113120130"></a>

#### `policy_based_challenge.rule_list.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-2333300211031022-1222300310101122-3000113000323213-3212220302131122-0131030311021333-0033032230113120-0131132131022303-3313132030103111"></a>

<a id="canonical-1013232112323301-1033310000200133-3121120301210333-3202310230113030-2323120312111022-0333101101333000-2333110003103330-1013122001012101"></a>

#### `policy_based_challenge.rule_list.rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-2100022312032302-2032323222000310-3320020323323030-0311311230302003-0203013303301133-1303003301003030-0233222330313000-1012020030201303"></a>

Type: `"object"`. single nested block, Optional.

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010323020212330-3200001013012332-1320023301013123-1331300302020213-3331000301222013-2211103230302112-0023011120231032-1321100003100020"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec`

- [any_asn](resources--cdn_loadbalancer--reference--group-013.md#canonical-3121222113032223-1300220113211122-3102303011211231-0001101101121000-1212013123030310-1110302210322232-2112200222032232-3210130202023000): complete subsection reference.

- [any_client](resources--cdn_loadbalancer--reference--group-013.md#canonical-1302001220201330-1331232223202101-2220300033110233-1330233330100223-1231113112012330-0032302122023001-1202003201120313-0123231113010300): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110221022133232-2300032333010202-2101013323001311-3132202333321222-0301330230222000-2122121322011301-1212020100230302-1233203323313101): complete subsection reference.

- [arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003233223300201-0312002012103020-1303002203022023-2111233231333112-1000323100130012-1320213001201033-1202333023000302-1013322233111312): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003102132023010-3231100012022023-0021011213211001-3321012032332023-1221003122112110-0003331033100121-1032002020102300-1103330212100101): complete subsection reference.

- [body_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-0020310012101102-3000230032112032-0132110110013312-1020032331321312-1203133212331130-0123320022003020-2120221312120320-3221212002033010): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-014.md#canonical-3230133003213133-3011200302131310-1110312033300333-3311222103011221-1231200100131103-1131223213322203-3302022113011000-2001203101232132): complete subsection reference.

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3313200200332201-1200210121121211-3220003300020012-2002110001123113-3233122133221310-0212323322101310-1333012331212122-2002300120333130): complete subsection reference.

- [disable_challenge](resources--cdn_loadbalancer--reference--group-014.md#canonical-0032332002211123-1202012303111223-1300130212323212-3131233302210131-2031312311130100-1303223302131323-2102102310301121-2010003203213221): complete subsection reference.

- [domain_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-1031330331223002-1133112231311120-0010033010221011-3200232312203013-1321321213331110-3211120301031200-0232002103001122-1233011032121333): complete subsection reference.

- [enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-014.md#canonical-3103311303303301-3120122203132131-2113220010310230-3210210112101313-1321321213233223-2130133230221223-3223033330123030-1202010323003320): complete subsection reference.

- [enable_javascript_challenge](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101321300022020-2012121100121213-1300112332120133-3301000323233111-0133300213322003-1130203021233202-1031230021202010-0002003123331010): complete subsection reference.

<a id="canonical-2122200121320131-3201120023221012-1211131301012120-0013032102322322-2123011123130213-0111223233003021-1323002221223313-1222020032311002"></a>

<a id="canonical-0030230131311301-2020221301310303-3212333322310202-0300223002331230-2032203122001110-3102022311312130-1210021321020111-0000001000313232"></a>

#### `policy_based_challenge.rule_list.rules.spec.expiration_timestamp` property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020): complete subsection reference.

- [http_method](resources--cdn_loadbalancer--reference--group-014.md#canonical-2312131312221300-0230011100003021-0021131212211331-0123333322230033-1210310032303322-1320033320200030-2221332222001320-2330310311021003): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-1221203300113022-2132333200200013-1003100021200202-1120013111223113-1232003120333102-0311332302111030-2000323030332330-3022032101132020): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-1012033320010022-2222231020323103-2321000021201310-3031023311331022-0012002023131321-2230011300323100-3130312013222320-2213121223200210): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-014.md#canonical-1323031133100220-3003131300131310-1203112022000311-1230021132323020-2132121211131031-2202311300331212-3111230320020110-3011011131332111): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-3232022003302213-0032131233110110-3231221110233032-1010300113202110-1012221322120230-1322201202221221-2233333330333200-3030322000303331): complete subsection reference.

<a id="canonical-3121222113032223-1300220113211122-3102303011211231-0001101101121000-1212013123030310-1110302210322232-2112200222032232-3210130202023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_asn` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-3112231331323121-3320022221133013-3131221010013300-2211320300230111-2013000321333320-1111132033001000-3201301203311200-2233300001002020"></a>

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
any_asn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302001220201330-1331232223202101-2220300033110233-1330233330100223-1231113112012330-0032302122023001-1202003201120313-0123231113010300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_client` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-0112212111231331-2131010303013302-2030110211100123-3213013230330300-0133311300333000-3013202033201200-3212013031102103-2321213223202212"></a>

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
any_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110221022133232-2300032333010202-2101013323001311-3132202333321222-0301330230222000-2122121322011301-1212020100230302-1233203323313101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-3110230313303220-3032103230312330-2231013003013003-2033332203002113-0310200302220123-0200123232211321-3120331230122210-3000232103113021"></a>

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
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-1202331333202023-0331330213201312-2032323202013121-0330311313200330-1201333013012020-0333112213110211-1321223333001020-3212001201310000"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
arg_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301202331323033-0232132122202030-2113200300002120-3301220012303200-0221212202133013-2231032101323000-0233131302233000-3011023123121300"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.arg_matchers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-1013323232133001-1333023212002100-1022133020223331-2113302220020330-3233023001021000-0022321120211103-1311320012110312-2021332021031131): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-1210121133032101-2033110320303231-0021123031202121-1020223313123313-3013232013322003-1121031320130203-0110222312210333-1021213130201333): complete subsection reference.

<a id="canonical-1020030031110113-0012330311023011-3011133220032122-2013323112220222-2213122332113131-1010222333012211-0323300020000013-1030120232331011"></a>

<a id="canonical-0132312122011001-0321011322330233-1101123122101000-1233203213133300-0323023231023101-1010210330033301-3231023231321013-0012111032221322"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--cdn_loadbalancer--reference--group-013.md#canonical-0130223113313203-1133021030001223-3111100303130003-3130100310030322-1203103132102121-2101230102123033-1022122313200300-2113313231323113): complete subsection reference.

<a id="canonical-2110331012213113-3212213132303313-0201133213000113-2012122203332113-3303130311300232-3111230022130313-2323131212020332-0300222131022312"></a>

<a id="canonical-3331311120311211-0000231302023223-0002113212211330-0310000203031213-2322131102222130-0232011033133121-3001233203021321-0220231033330023"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.name` property

Type: `"string"`. Optional.

A case-sensitive JSON path in the HTTP request body.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1013323232133001-1333023212002100-1022133020223331-2113302220020330-3233023001021000-0022321120211103-1311320012110312-2021332021031131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-2013303011030110-2120110113222231-1010332303112313-2222102230111220-0320010011011303-3103100212221020-1001002311020001-3021030333000111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210121133032101-2033110320303231-0021123031202121-1020223313123313-3013232013322003-1121031320130203-0110222312210333-1021213130201333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-1032021110132311-1323302102033111-0022101301130301-0233312020213013-3323332212002102-1321213200321013-0100110333001202-3101303110120303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130223113313203-1133021030001223-3111100303130003-3130100310030322-1203103132102121-2101230102123033-1022122313200300-2113313231323113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-2321013130100133-3100211211012323-1330213213310232-3211201323310010-2132311001132132-0103313033102300-1033120102222332-0220330032231000"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102033130031202-0010123202112213-0223233303111232-0332322102333201-1133030330330233-0130100022311200-3101332220101311-1011102030230201"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.arg_matchers.item`

<a id="canonical-2213031201121112-1030213023202122-0333301300030100-2003230131113303-2213122021121211-0222330023110232-1000330122122111-1111213030102033"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1322130112200211-2323032300330112-3033013312221023-3013300210003320-1010122112020221-0101101220022100-0102221001132020-2012321323002112"></a>

<a id="canonical-1113010202011103-3232230013011030-0233123320003101-1200332111321100-0121232020333120-0123020100020213-2131231232103211-2113301122021321"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2111032001223332-1331033102323001-2022331201321132-1331333320310203-1101303103200021-3011030013201012-1023031003331031-1321231132013213"></a>

<a id="canonical-3023311201033203-2222303221010013-3300023222320323-3031020133203130-3312231113321303-2230022221210301-1212101300203221-0000331120010100"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0003233223300201-0312002012103020-1303002203022023-2111233231333112-1000323100130012-1320213001201033-1202333023000302-1013322233111312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-3203002021210222-0203103302001100-0313033230220021-0232012022120333-3103310130200201-0032102301021332-1030222233032220-2133231210111220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202332121121312-3111001003333230-0322333012103222-2122032300011232-1000100313103111-2130033223002331-1322121103321333-0322312231202122"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_list`

<a id="canonical-1323232023302012-2300320120020232-2103100331230112-2332223223210323-0032031010232102-2302222333203203-1213303211001121-3120201210113223"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0003102132023010-3231100012022023-0021011213211001-3321012032332023-1221003122112110-0003331033100121-1032002020102300-1103330212100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-3231331210032231-2123312121132312-0323213001332133-0011312120300023-2211111233320302-1302103230010132-2313131301331222-0300323203312230"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001020020113120-3003023102231320-0103330023133001-3220333200223211-0310002023120120-1030023112203322-1211013331313330-2021333321121320"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_matcher`

- [asn_sets](resources--cdn_loadbalancer--reference--group-014.md#canonical-0003030013130330-2331113132100113-3133013032111101-2330111202210322-2212222032222202-2223222120111202-1220132122303010-0212112303222130): complete subsection reference.
