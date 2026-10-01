---
page_title: "xcsh_user_identification reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification reference."
---

# xcsh_user_identification reference

<a id="canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c12667076bda08ee7d09cdb58ee247f29be24cd2545d13b3abc8be2d013ac604"></a>

## Property reference — Property reference / e17b1686cc9c / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- Property reference

<a id="canonical-0d104a13b90ee6bac46e291d03563e789489501bd595cad435b112d439b76b53"></a>

## Direct properties — Property reference / e17b1686cc9c / 3

<a id="canonical-a7d9e1ad63297db713eea6ace2e67ca4cd851254e716155e30f58e8cead8c93c"></a>

<a id="canonical-c3eadf6eaed00f15d8391222cdb215d4949f4b44a859e7a187c60bccee1714be"></a>

## annotations property — Property reference / e17b1686cc9c / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-72e3fe29765ba2dd2aa7d03137f515cda37e83b72b919b9e3c0055fb05a7d01c"></a>

<a id="canonical-558bd1fbd2ac8e87d10f0e249a4a4a6f11cbc98606348a831ca8a0d618f9d3f0"></a>

## description property — Property reference / e17b1686cc9c / 5

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
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
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-6bf148338eb4500d005752fd2f340c42216a77a684ada1cf36920d8c7f5153cc"></a>

<a id="canonical-3edb6115e894ba73145f86c98628403c6e86f4229f0119e3d4453f68ef95d5d3"></a>

## disable property — Property reference / e17b1686cc9c / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5e107cb3ad650a5b8e9b02eeac3e99968970c98572ddf9ee7ecf5831406091e3"></a>

<a id="canonical-037ebfd2c8576baa4842398e6219ec263e16766047ae21938006f4ea17b06d94"></a>

## id property — Property reference / e17b1686cc9c / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-ca75edb4267dabf93bf070709490e951645a0b6dec92b3d60e9d44eac74002a3"></a>

<a id="canonical-1bb3871cb27db5bd4b04927aefeaf2bb261cbf459a411298134dbe233ce717f0"></a>

## labels property — Property reference / e17b1686cc9c / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fddc1bd69b87afbaf58b7fd27ff042c0f672b2008c594664adc88214f5005247"></a>

<a id="canonical-aa26767119f3643bd46eedb117fbf3205c29aeb02ffc5438aeec2ca8efc39a58"></a>

## name property — Property reference / e17b1686cc9c / 9

Type: `"string"`. Required.

Name of the User Identification. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-bc75740b53c2bd49b4ae020aec893440557063d9259cc4f225aefcac5aa4c90c"></a>

<a id="canonical-df8f373dd0aac1681f084103340b06948fdf3e631270f277bd85f4da3f425ecf"></a>

## namespace property — Property reference / e17b1686cc9c / 10

Type: `"string"`. Required.

Namespace where the User Identification is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023): complete subsection reference.

- [timeouts](resources--user_identification--reference--group-001.md#canonical-2214772ec6b124fd5f11279c689f3717409c0a48fdd1a5372b3ee45bdacaad3b): complete subsection reference.

<a id="canonical-038754948a66d3f2de40d6f6da789500b35afcda9efe91c752c5cce5834ad746"></a>

## All schema paths — Property reference / e17b1686cc9c / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--user_identification--reference--group-001.md#canonical-a7d9e1ad63297db713eea6ace2e67ca4cd851254e716155e30f58e8cead8c93c) |
| `description` | [description](resources--user_identification--reference--group-001.md#canonical-72e3fe29765ba2dd2aa7d03137f515cda37e83b72b919b9e3c0055fb05a7d01c) |
| `disable` | [disable](resources--user_identification--reference--group-001.md#canonical-6bf148338eb4500d005752fd2f340c42216a77a684ada1cf36920d8c7f5153cc) |
| `id` | [id](resources--user_identification--reference--group-001.md#canonical-5e107cb3ad650a5b8e9b02eeac3e99968970c98572ddf9ee7ecf5831406091e3) |
| `labels` | [labels](resources--user_identification--reference--group-001.md#canonical-ca75edb4267dabf93bf070709490e951645a0b6dec92b3d60e9d44eac74002a3) |
| `name` | [name](resources--user_identification--reference--group-001.md#canonical-fddc1bd69b87afbaf58b7fd27ff042c0f672b2008c594664adc88214f5005247) |
| `namespace` | [namespace](resources--user_identification--reference--group-001.md#canonical-bc75740b53c2bd49b4ae020aec893440557063d9259cc4f225aefcac5aa4c90c) |
| `rules` | [rules](resources--user_identification--reference--group-001.md#canonical-38ea4b540e90f9e83d3a943732e29991344115ee2209f7e55cb021e25c8f6f77) |
| `rules.client_asn` | [rules.client_asn](resources--user_identification--reference--group-001.md#canonical-f3d6fc5199cc15be8226cfe35f0a16fef140884e281f6dcd6c21ac21ba330eab) |
| `rules.client_city` | [rules.client_city](resources--user_identification--reference--group-001.md#canonical-5ed136c2497b3bbfd6cdcd15d143d1293ecaf463f4a9058e81ae5103f36df562) |
| `rules.client_country` | [rules.client_country](resources--user_identification--reference--group-001.md#canonical-8722eb505ab5316285a2a7fa89c1899ff6a00ec918f676571b73bc1124dc0eec) |
| `rules.client_ip` | [rules.client_ip](resources--user_identification--reference--group-001.md#canonical-32c5cdfc22baba2cde3d4b1a8127887eb0716219bc1d9c131c6433975aa36530) |
| `rules.client_region` | [rules.client_region](resources--user_identification--reference--group-001.md#canonical-1f817f6382a267d4d8447209093c2cfa1a4de7636eb5498a6eac6484f7fbd161) |
| `rules.cookie_name` | [rules.cookie_name](resources--user_identification--reference--group-001.md#canonical-106b7c5abb603e1c5c718b17da71b86448837ca3861dddafe99fee474e352750) |
| `rules.http_header_name` | [rules.http_header_name](resources--user_identification--reference--group-001.md#canonical-e84973517463a1a28465e65b51bc6443ee023de2d726c816892386e8c8b6a3a3) |
| `rules.ip_and_http_header_name` | [rules.ip_and_http_header_name](resources--user_identification--reference--group-001.md#canonical-0c7a034702f77599a713f3893d7a7a6826a11b766ef12440e6a5cf9d98124b80) |
| `rules.ip_and_ja4_tls_fingerprint` | [rules.ip_and_ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-f0f78707dc8380d6ad755b6ebc9c01a4524c78415677cfc792c99edcddc74915) |
| `rules.ip_and_tls_fingerprint` | [rules.ip_and_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-54f58357d656ce5796a58c10afcbd26da203c56355d165e5e2b221c29d302eff) |
| `rules.ja4_tls_fingerprint` | [rules.ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-16bab3491c82d1fe5929bb9de8bbf91bff2417664d36c87868931bd9574d9131) |
| `rules.jwt_claim_name` | [rules.jwt_claim_name](resources--user_identification--reference--group-001.md#canonical-d8cc052e489df71df720eb79daa0ed8df463d5ca07a3ca45d69afeb0c53ce4ac) |
| `rules.none` | [rules.none](resources--user_identification--reference--group-001.md#canonical-aec08e1ec393a992cb9f88ed91c49c847ffaf0e8396f74645d4e97df2f7d94a3) |
| `rules.query_param_key` | [rules.query_param_key](resources--user_identification--reference--group-001.md#canonical-1b851bedc60c730507dbdeb7ac5efc3223ec4a14ea70e69e213da2fb2496f44f) |
| `rules.tls_fingerprint` | [rules.tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-4623b727cb8a5d5d55fa46c51fe2f7bedca61fa93e8c64d70f98931f7c304fda) |
| `timeouts` | [timeouts](resources--user_identification--reference--group-001.md#canonical-a6db2fe5aeeed2eeb6a5744f604c3c4b26a10836b566a6405c2f76f1d5121884) |
| `timeouts.create` | [timeouts.create](resources--user_identification--reference--group-001.md#canonical-6c16fed40df1cdc01e2e9a67bc783502bd9080e558284413339b240dd1b1f078) |
| `timeouts.delete` | [timeouts.delete](resources--user_identification--reference--group-001.md#canonical-d7f4590eb4c6ac9f8bd754d867b7c18dcf293b9eb8765892d19d50bc467b1542) |
| `timeouts.read` | [timeouts.read](resources--user_identification--reference--group-001.md#canonical-d59149c4678ee58e9a36eb9d61f563bc2fdc92af49de7a60c06b62b9e7fde06b) |
| `timeouts.update` | [timeouts.update](resources--user_identification--reference--group-001.md#canonical-b289ede587ce59365af89eef7ac8209863478b84219630e9c948e1f40a113a44) |

<a id="canonical-7b6d1b0573dcc417c73eb748b997c06fceb5ff929591ecbbd4ba60dcc32138c2"></a>

## Next pages — Property reference / e17b1686cc9c / 12

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [timeouts](resources--user_identification--reference--group-001.md#canonical-2214772ec6b124fd5f11279c689f3717409c0a48fdd1a5372b3ee45bdacaad3b)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc99cf0e6ea8365b3e4b5e3566de2536a851f1dfde61fc5641d4792c38fc4087"></a>

## rules — rules / 20c6f5483e38 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- rules

<a id="canonical-38ea4b540e90f9e83d3a943732e29991344115ee2209f7e55cb021e25c8f6f77"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules that are evaluated sequentially against the input fields extracted from an API
request in order to determine a user identifier. Evaluation of the rules is terminated once a user
identifier has been extracted.

Upstream description:

An ordered list of rules that are evaluated sequentially against the input fields extracted from an
API request in order to determine a user identifier. Evaluation of the rules is terminated once a
user identifier has been extracted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("client_asn",
    "client_city"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_country"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_asn",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_asn",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_asn",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_asn",
    "none"),
  validators.ConflictingListObjectAttributes("client_asn",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_asn",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_country"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_city",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_city",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_city",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_city",
    "none"),
  validators.ConflictingListObjectAttributes("client_city",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_city",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "client_ip"),
  validators.ConflictingListObjectAttributes("client_country",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_country",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_country",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_country",
    "none"),
  validators.ConflictingListObjectAttributes("client_country",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_country",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "client_region"),
  validators.ConflictingListObjectAttributes("client_ip",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_ip",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_ip",
    "none"),
  validators.ConflictingListObjectAttributes("client_ip",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_ip",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "cookie_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("client_region",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("client_region",
    "none"),
  validators.ConflictingListObjectAttributes("client_region",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("client_region",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "http_header_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "none"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("cookie_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_http_header_name"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "none"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("http_header_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ip_and_ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_http_header_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "ip_and_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_ja4_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "ja4_tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ip_and_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "jwt_claim_name"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "none"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("ja4_tls_fingerprint",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "none"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("jwt_claim_name",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("none",
    "query_param_key"),
  validators.ConflictingListObjectAttributes("none",
    "tls_fingerprint"),
  validators.ConflictingListObjectAttributes("query_param_key",
    "tls_fingerprint")}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-e61573f62de80da440755229258b3bbfffddad8ebf07dc882a98f7f3779376fb"></a>

## Direct properties — rules / 20c6f5483e38 / 3

- [client_asn](resources--user_identification--reference--group-001.md#canonical-2367b986b505e91088fa8ed14f76beced6ac7faa2e25c9dad48685c39cf446ca): complete subsection reference.

- [client_city](resources--user_identification--reference--group-001.md#canonical-f573913dfb266cb88c6ccd0f59873ed30eb3f780e9505cd801a2d3dd14d4daf4): complete subsection reference.

- [client_country](resources--user_identification--reference--group-001.md#canonical-6685d7eb826077b67583250ab575f25ef170a5b2a93a6b3132b6d9f68616e43d): complete subsection reference.

- [client_ip](resources--user_identification--reference--group-001.md#canonical-07a03f0b25d9e74d4cb196367e3dd5a08f220d5d355c65a3ee6527078323ddf2): complete subsection reference.

- [client_region](resources--user_identification--reference--group-001.md#canonical-644eeb31b74424e7b5259a5d94470a91ad3af53caf9534d43365fd6e201ec6b1): complete subsection reference.

<a id="canonical-106b7c5abb603e1c5c718b17da71b86448837ca3861dddafe99fee474e352750"></a>

<a id="canonical-0e15d0b8dc296eb6c491421ade876332a40a77833569850a0c2a8e6cf861ca18"></a>

## cookie_name property — rules / 20c6f5483e38 / 4

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user identifier.

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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-e84973517463a1a28465e65b51bc6443ee023de2d726c816892386e8c8b6a3a3"></a>

<a id="canonical-7f34b07ae703bebc99a8b1095917062ceb92f1799598b31f540268d04d18aa23"></a>

## http_header_name property — rules / 20c6f5483e38 / 5

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user identifier.

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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0c7a034702f77599a713f3893d7a7a6826a11b766ef12440e6a5cf9d98124b80"></a>

<a id="canonical-dfe22e6adcdc02d52644e95e1f6736eb355db01ae82fc0ea360f7e0534028d12"></a>

## ip_and_http_header_name property — rules / 20c6f5483e38 / 6

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [ip_and_ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-668129970e274757516d407e9183455bb494fd696d83d5298f482909c3472567): complete subsection reference.

- [ip_and_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-b563986be14bb8118b211cfa6e16ab4959af1ae56d319cbddb94cb833b07eb40): complete subsection reference.

- [ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-48c71f798a7a049156dd1bcd55dc1c8b826a35f3bac003a8fc582f4a6feb78e0): complete subsection reference.

<a id="canonical-d8cc052e489df71df720eb79daa0ed8df463d5ca07a3ca45d69afeb0c53ce4ac"></a>

<a id="canonical-92d955c58a62adb5abf298bb17d50bdf4dc4f415cbe0dec3077a3e6c2cf28f09"></a>

## jwt_claim_name property — rules / 20c6f5483e38 / 7

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [none](resources--user_identification--reference--group-001.md#canonical-ed5a06323ca9efb4cdd87159494ec67b9fd6d0fdf7b0dfec3f0eb2d94ab3a022): complete subsection reference.

<a id="canonical-1b851bedc60c730507dbdeb7ac5efc3223ec4a14ea70e69e213da2fb2496f44f"></a>

<a id="canonical-32d7d2fefffbec27da77bde1e87a63e8718db997ca457a61bcb8aa1549c831ec"></a>

## query_param_key property — rules / 20c6f5483e38 / 8

Type: `"string"`. Optional.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user identifier.

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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-6621d980a09598d15e02b5c2d09dcc011c7439e22d4a4b10db7e099b56c954a0): complete subsection reference.

<a id="canonical-c5a8de7d8b76139b0a6f6abc1d20eab219681cf17941d5c74c148b28847041ef"></a>

## Next pages — rules / 20c6f5483e38 / 9

- [rules.client_asn](resources--user_identification--reference--group-001.md#canonical-2367b986b505e91088fa8ed14f76beced6ac7faa2e25c9dad48685c39cf446ca)
- [rules.client_city](resources--user_identification--reference--group-001.md#canonical-f573913dfb266cb88c6ccd0f59873ed30eb3f780e9505cd801a2d3dd14d4daf4)
- [rules.client_country](resources--user_identification--reference--group-001.md#canonical-6685d7eb826077b67583250ab575f25ef170a5b2a93a6b3132b6d9f68616e43d)
- [rules.client_ip](resources--user_identification--reference--group-001.md#canonical-07a03f0b25d9e74d4cb196367e3dd5a08f220d5d355c65a3ee6527078323ddf2)
- [rules.client_region](resources--user_identification--reference--group-001.md#canonical-644eeb31b74424e7b5259a5d94470a91ad3af53caf9534d43365fd6e201ec6b1)
- [rules.ip_and_ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-668129970e274757516d407e9183455bb494fd696d83d5298f482909c3472567)
- [rules.ip_and_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-b563986be14bb8118b211cfa6e16ab4959af1ae56d319cbddb94cb833b07eb40)
- [rules.ja4_tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-48c71f798a7a049156dd1bcd55dc1c8b826a35f3bac003a8fc582f4a6feb78e0)
- [rules.none](resources--user_identification--reference--group-001.md#canonical-ed5a06323ca9efb4cdd87159494ec67b9fd6d0fdf7b0dfec3f0eb2d94ab3a022)
- [rules.tls_fingerprint](resources--user_identification--reference--group-001.md#canonical-6621d980a09598d15e02b5c2d09dcc011c7439e22d4a4b10db7e099b56c954a0)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-2367b986b505e91088fa8ed14f76beced6ac7faa2e25c9dad48685c39cf446ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7485d08065bb27fbe8370d487b3f51fca0bae9e9e7a8c5c257407f79b8acf8ff"></a>

## rules.client_asn — rules.client_asn / c7590115c665 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.client_asn

<a id="canonical-f3d6fc5199cc15be8226cfe35f0a16fef140884e281f6dcd6c21ac21ba330eab"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
client_asn = {}
```

<a id="canonical-cc4cac9b82e1012dae648063be8f024d01aefa361e77555f48317c6812bb3db3"></a>

## Direct properties — rules.client_asn / c7590115c665 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa7adfdc4ed29b4d75fc6a0581d11e4fa468d39b60e8d05b3287b073c4009271"></a>

## Next pages — rules.client_asn / c7590115c665 / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-f573913dfb266cb88c6ccd0f59873ed30eb3f780e9505cd801a2d3dd14d4daf4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fe9615ba135b854df3f819f8bfd214f8aef6ede9a4cd7242b33edb73adfbc9a"></a>

## rules.client_city — rules.client_city / 4a2ce3a01161 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.client_city

<a id="canonical-5ed136c2497b3bbfd6cdcd15d143d1293ecaf463f4a9058e81ae5103f36df562"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
client_city = {}
```

<a id="canonical-16b030660298ab83d579649d71785d5146a779edd32d1db975842d8cd072f148"></a>

## Direct properties — rules.client_city / 4a2ce3a01161 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9e30fc7a4d6650d8461a4325e2cc94432af794473cafd70513a6ce48a1c2836"></a>

## Next pages — rules.client_city / 4a2ce3a01161 / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-6685d7eb826077b67583250ab575f25ef170a5b2a93a6b3132b6d9f68616e43d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7de32b19af18cc4ef2cdfaba0590a844298cb6b2be6ff5ebd55cc4ee13debe02"></a>

## rules.client_country — rules.client_country / f06f34a61285 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.client_country

<a id="canonical-8722eb505ab5316285a2a7fa89c1899ff6a00ec918f676571b73bc1124dc0eec"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
client_country = {}
```

<a id="canonical-719595c3e301c52e7b256c934c391e43b7b17235c1dfb82e28b7da7952cc9c51"></a>

## Direct properties — rules.client_country / f06f34a61285 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0efa29305beaef8240ea86c9a145c250c5405b782cb0aaae9efaaadd7fccd7a2"></a>

## Next pages — rules.client_country / f06f34a61285 / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-07a03f0b25d9e74d4cb196367e3dd5a08f220d5d355c65a3ee6527078323ddf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fe1edcce5c1970fac7321820f18d05a3648793d54121a3a970f8229d2fb06ab"></a>

## rules.client_ip — rules.client_ip / 27f9818dfa7d / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.client_ip

<a id="canonical-32c5cdfc22baba2cde3d4b1a8127887eb0716219bc1d9c131c6433975aa36530"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
client_ip = {}
```

<a id="canonical-4909134723c1352a908fa525a77a8e2171835a096e94de1076088960c8de19da"></a>

## Direct properties — rules.client_ip / 27f9818dfa7d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c14c3b4890c3d865a2ee767940100b12a6c7970dcfd89166de266bfac212209"></a>

## Next pages — rules.client_ip / 27f9818dfa7d / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-644eeb31b74424e7b5259a5d94470a91ad3af53caf9534d43365fd6e201ec6b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59a99b30a18a0783e6b58e648688a072b5348b5a7d87388f43968f14b7fbff61"></a>

## rules.client_region — rules.client_region / de4296b19644 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.client_region

<a id="canonical-1f817f6382a267d4d8447209093c2cfa1a4de7636eb5498a6eac6484f7fbd161"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
client_region = {}
```

<a id="canonical-9e1f11eb0a51b1b5d11f0520f378df070f53c9cfa0777c6004908c6a833f4d0e"></a>

## Direct properties — rules.client_region / de4296b19644 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-70d8084cec000b4fb8d6283227eb6a0fb7774bae5900dc2b19b36f719130138d"></a>

## Next pages — rules.client_region / de4296b19644 / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-668129970e274757516d407e9183455bb494fd696d83d5298f482909c3472567"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-302a9aa1c3a50b8fd9e35ee5e02eb824053a825786e0429fd92f6f22c1776c2a"></a>

## rules.ip_and_ja4_tls_fingerprint — rules.ip_and_ja4_tls_fingerprint / ea65ad099567 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.ip_and_ja4_tls_fingerprint

<a id="canonical-f0f78707dc8380d6ad755b6ebc9c01a4524c78415677cfc792c99edcddc74915"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ip_and_ja4_tls_fingerprint = {}
```

<a id="canonical-1a539c85b44869fb53d87fe04b63ad27cfd858eaf745b346b0bef2fc01bab037"></a>

## Direct properties — rules.ip_and_ja4_tls_fingerprint / ea65ad099567 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-054872e9f75eefec0b7c90f0986f0d1f033e60459829b548ffc7acb9023f1aa9"></a>

## Next pages — rules.ip_and_ja4_tls_fingerprint / ea65ad099567 / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-b563986be14bb8118b211cfa6e16ab4959af1ae56d319cbddb94cb833b07eb40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9110e886b96a7cd3ac1a93cccbc342dba76ad9a5dc25f71b287950226240a599"></a>

## rules.ip_and_tls_fingerprint — rules.ip_and_tls_fingerprint / 5dedd5d137cf / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.ip_and_tls_fingerprint

<a id="canonical-54f58357d656ce5796a58c10afcbd26da203c56355d165e5e2b221c29d302eff"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ip_and_tls_fingerprint = {}
```

<a id="canonical-95952f18ae262cc7fe71bd0523d72b3b1909839244279317afb0c2ab7b6d7ddf"></a>

## Direct properties — rules.ip_and_tls_fingerprint / 5dedd5d137cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89d56435cd16e413a8b7b74e2230a85099f0230a97a3502ff24e42cc41641f56"></a>

## Next pages — rules.ip_and_tls_fingerprint / 5dedd5d137cf / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-48c71f798a7a049156dd1bcd55dc1c8b826a35f3bac003a8fc582f4a6feb78e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8c9a97be3e120a0cd178134f1db749040e41dffa4505d37fb721fc7e01d69cd"></a>

## rules.ja4_tls_fingerprint — rules.ja4_tls_fingerprint / a155a1acc1c1 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.ja4_tls_fingerprint

<a id="canonical-16bab3491c82d1fe5929bb9de8bbf91bff2417664d36c87868931bd9574d9131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ja4 tls fingerprint.

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

Terraform syntax:

```terraform
ja4_tls_fingerprint = {}
```

<a id="canonical-4cfc3b97012a9c88b4910a6dc0b28ed1dcc29e298f61364d0afebd847a659590"></a>

## Direct properties — rules.ja4_tls_fingerprint / a155a1acc1c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39ac187745fd3359475fa677016553aac7d408d55f4bd097962f7bb42f5b383c"></a>

## Next pages — rules.ja4_tls_fingerprint / a155a1acc1c1 / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-ed5a06323ca9efb4cdd87159494ec67b9fd6d0fdf7b0dfec3f0eb2d94ab3a022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9128735dc7f0957dec938886606ec78553a15becc583fd73b80349e8e72a2287"></a>

## rules.none — rules.none / 469a324710e6 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.none

<a id="canonical-aec08e1ec393a992cb9f88ed91c49c847ffaf0e8396f74645d4e97df2f7d94a3"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
none = {}
```

<a id="canonical-406a70727444c0b8ae11c9238fa2008c43685cbde83e3ab9606ce9e59e816ba4"></a>

## Direct properties — rules.none / 469a324710e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-576def3066fdbc36bf9fad4b099de880b5b2d5c6425848b29277e8482754f869"></a>

## Next pages — rules.none / 469a324710e6 / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-6621d980a09598d15e02b5c2d09dcc011c7439e22d4a4b10db7e099b56c954a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d87d1b436d84b2ac914e9149475d3d4819d00a9d3c62c362699f6170000c1e4"></a>

## rules.tls_fingerprint — rules.tls_fingerprint / dce1bed56618 / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- rules.tls_fingerprint

<a id="canonical-4623b727cb8a5d5d55fa46c51fe2f7bedca61fa93e8c64d70f98931f7c304fda"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for tls fingerprint.

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

Terraform syntax:

```terraform
tls_fingerprint = {}
```

<a id="canonical-83637c38db15a2df45fc562b59446bb2e3d8e4200282d0a17e1637480b67a747"></a>

## Direct properties — rules.tls_fingerprint / dce1bed56618 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3faefd9605a3aa7c8fadccfc86cdb3ec695bfd437aad653c5fc47e40d8bdcf7e"></a>

## Next pages — rules.tls_fingerprint / dce1bed56618 / 4

- [rules](resources--user_identification--reference--group-001.md#canonical-46978c4ef888af5e964b2f4d9bf6efaebae1f1057124db03a28413825ee1f023)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)

<a id="canonical-2214772ec6b124fd5f11279c689f3717409c0a48fdd1a5372b3ee45bdacaad3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56be3d65c596457cd14247c5aaee2466aed94b6095918f0693ddd257bcdb14a7"></a>

## timeouts — timeouts / dffd2cbcc7cd / 2

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- timeouts

<a id="canonical-a6db2fe5aeeed2eeb6a5744f604c3c4b26a10836b566a6405c2f76f1d5121884"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-93e1e33924f71a02ff6891d368141b76160932bec03cff3c3b968a1d3b169282"></a>

## Direct properties — timeouts / dffd2cbcc7cd / 3

<a id="canonical-6c16fed40df1cdc01e2e9a67bc783502bd9080e558284413339b240dd1b1f078"></a>

<a id="canonical-5e06fe9e32ad548696fc378ebc3cb761b87fdf8b423d07899d1938003dc11c13"></a>

## create property — timeouts / dffd2cbcc7cd / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d7f4590eb4c6ac9f8bd754d867b7c18dcf293b9eb8765892d19d50bc467b1542"></a>

<a id="canonical-45339ae82878d132b40944de0190bf340b4fac80974ed8ace5e9a7a5f042e3a1"></a>

## delete property — timeouts / dffd2cbcc7cd / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-d59149c4678ee58e9a36eb9d61f563bc2fdc92af49de7a60c06b62b9e7fde06b"></a>

<a id="canonical-a3ed1ea85eb81a04648566a8e267ab6de031ccc28d39784593544e487cd17ea4"></a>

## read property — timeouts / dffd2cbcc7cd / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-b289ede587ce59365af89eef7ac8209863478b84219630e9c948e1f40a113a44"></a>

<a id="canonical-336cf83ca5ea823e49dec02e65ae9769427dec70edd66685e089f579a52a0394"></a>

## update property — timeouts / dffd2cbcc7cd / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-8e17ae1b610d0e0461e637f102fdffa1782ea63306672e32ecfd22c308460c48"></a>

## Next pages — timeouts / dffd2cbcc7cd / 8

- [Property reference](resources--user_identification--reference--group-001.md#canonical-328d4b15110df9e0173252e924887fa57f78e038edb0b55fb6fa3a98c5b5fbb5)
- [xcsh_user_identification](../resources/user_identification.md#canonical-21b6255ba6c8bf41d1721c4ac5c61d5955733e840d9ab4bbf2661d87b6c35e16)
