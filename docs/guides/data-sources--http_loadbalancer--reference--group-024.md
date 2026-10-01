---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-ec5c075c6b9e0ad63ec222132ecd9861d715aeef61a582497c1b2c62ef84f900"></a>

## store_provider property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / 233f051eb79c / 6

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

<a id="canonical-b9a12580f45402fe6609be2e54aecfc7d05cd7fffa94e9b54dea55f18519d467"></a>

## Next pages — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / 233f051eb79c / 7

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-023.md#canonical-60ac953e6fa03b699f98b4339db942385909c5a499261287d7eceba236c45b09)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-390d69dcd8d4a77f3e202a8691d36a9b88f36e46e52c49adad332a3cfbdd4251"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-442ef7b5b33428767e61558b45da684eec94fc6830c5129ef6c0bccc05a52894"></a>

## routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / fe4b94d3a036 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.request_cookies_to_add](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f3a8a13d10671a490d330b6c70591cbd941c941b58f82b19df0fe4c42e76852)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-023.md#canonical-60ac953e6fa03b699f98b4339db942385909c5a499261287d7eceba236c45b09)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-9becac5ce6564a56e49d81a0b2eab5d3c1878c27b89b2dc4d9e047204ca392de"></a>

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

<a id="canonical-433e78c638d856ef98f6edcb73917d42926d8e9b8d607b33bab20aa98b429979"></a>

## Direct properties — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / fe4b94d3a036 / 3

<a id="canonical-229490e395e6ec6aeaf6e7f92ae7598a04459810f9ca8a55920331133ca2a24c"></a>

<a id="canonical-a9e749a3c104001fb140e18890a704a69cbfc13d05a792855fe6f8aac5010ce0"></a>

## provider_ref property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / fe4b94d3a036 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-699a4e9f04c0188b772cbad48f776f64953d7c40856159b3ebc0ce2d0a6c927b"></a>

<a id="canonical-9f6de54af34d788d4a962e3f8bf57e65b57164c40456412306b5ab5592d0728b"></a>

## url property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / fe4b94d3a036 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-5727aa7b686e9ffc99574d6c937636f71ee8bbead4c06105b3e026dba8011133"></a>

## Next pages — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_s / fe4b94d3a036 / 6

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-023.md#canonical-60ac953e6fa03b699f98b4339db942385909c5a499261287d7eceba236c45b09)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9be9557dde347d2723c87163cf37b8d7322baf0f5e2adb7a120366f63e6a90a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0eb8e6b8c1377e7c0b6eda36dc8a75b680feabe2d3b5f23239eb6fe82f10318b"></a>

## routes.simple_route.advanced_options.request_headers_to_add — routes.simple_route.advanced_options.request_headers_to_add / 211ef3ff773e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.request_headers_to_add

<a id="canonical-d45554cd6ca6a349c743134f5c3a6952fd71d73a929955e9e97f523bf5fa4e3b"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP request being routed towards upstream.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-19254952e38bf00cc5a3453099b6f61db07760d8cc6ed377b45b022d8e12eaca"></a>

## Direct properties — routes.simple_route.advanced_options.request_headers_to_add / 211ef3ff773e / 3

<a id="canonical-2cce7bc17539e301ff1f9635f68fb90e6fb713aa4316ca75aff5782034261caf"></a>

<a id="canonical-181739f46e3c231252faf5f57b0bf039f0935f4fcb53e132ca83157e7fbfce89"></a>

## append property — routes.simple_route.advanced_options.request_headers_to_add / 211ef3ff773e / 4

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-124511bd6eef40c235b0faae228e2de23d63ccbe6bca350ac9d364a705f76f62"></a>

<a id="canonical-91393b92bf1826526a90422bdaee4b8bec124e80775e8094858868fb0cebb5ce"></a>

## name property — routes.simple_route.advanced_options.request_headers_to_add / 211ef3ff773e / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-6b9d2e2292f2f72874f434d8d4889c8df05c12a3e51dcf3ee29ba78108fbc386): complete subsection reference.

<a id="canonical-6a2bdb482783fc174e10f3f33a26e42b7f501ce38b2839be84de392452fcb6ed"></a>

<a id="canonical-0bcfa7f17df6425353a37fa182aad37bb4dc0e0ee9054c3e0b3a74a40a0fc238"></a>

## value property — routes.simple_route.advanced_options.request_headers_to_add / 211ef3ff773e / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-7b36cb121e504b45f2f74b0dcaa8e89de14a301a810b1b6ed39151e6773c829d"></a>

## Next pages — routes.simple_route.advanced_options.request_headers_to_add / 211ef3ff773e / 7

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-6b9d2e2292f2f72874f434d8d4889c8df05c12a3e51dcf3ee29ba78108fbc386)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6b9d2e2292f2f72874f434d8d4889c8df05c12a3e51dcf3ee29ba78108fbc386"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55dd1be6ebc04cc8c7a3fcc771720989719b89e85da0e13f8bb50f1916894510"></a>

## routes.simple_route.advanced_options.request_headers_to_add.secret_value — routes.simple_route.advanced_options.request_headers_to_add.secret_value / a37a78e7643a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.request_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-9be9557dde347d2723c87163cf37b8d7322baf0f5e2adb7a120366f63e6a90a6)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value

<a id="canonical-745047388b3ae94f2273802fb79d40f92a67448f590306e1d9b6dc8353ecce91"></a>

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

<a id="canonical-da5582934e6980c016839ad7543b41e9ca73a422bd83250fbc162961519aa051"></a>

## Direct properties — routes.simple_route.advanced_options.request_headers_to_add.secret_value / a37a78e7643a / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-5a176cfba80ca61bec1193f0869f0176ac8d0da0f377559c2359a9c1ed1ea819): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-ccffc91b1fcbfedfc77a65102bf48d33d45ae90a5f62e00ce5f82b619554261b): complete subsection reference.

<a id="canonical-7787cbaf95af0c12e3a2b38e6975c3c6227cff2a2c03bbcf1bb421b2728b78e9"></a>

## Next pages — routes.simple_route.advanced_options.request_headers_to_add.secret_value / a37a78e7643a / 4

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-5a176cfba80ca61bec1193f0869f0176ac8d0da0f377559c2359a9c1ed1ea819)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-ccffc91b1fcbfedfc77a65102bf48d33d45ae90a5f62e00ce5f82b619554261b)
- [routes.simple_route.advanced_options.request_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-9be9557dde347d2723c87163cf37b8d7322baf0f5e2adb7a120366f63e6a90a6)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5a176cfba80ca61bec1193f0869f0176ac8d0da0f377559c2359a9c1ed1ea819"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c0f24b93218064b0aa9483a49becf013e5c40da444d4b9d42bd0b4db179084b"></a>

## routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 5292986458d0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.request_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-9be9557dde347d2723c87163cf37b8d7322baf0f5e2adb7a120366f63e6a90a6)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-6b9d2e2292f2f72874f434d8d4889c8df05c12a3e51dcf3ee29ba78108fbc386)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-8000c75395df19f92cd298c0d9f7837629c82ffa8b21889a9e9461f03520b10c"></a>

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

<a id="canonical-5c1c488cb342a5626f4832c3bfee3307624f6fc24349649a40fed36fcabf905a"></a>

## Direct properties — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 5292986458d0 / 3

<a id="canonical-0233bba4ca65f20ebfd1f564754ee4902a1da2355abd73a92818b3ecdab832ef"></a>

<a id="canonical-80571c1fac8a13be4a59df39c32ba53aa97a87b6c9ab48076b931197bd32bd59"></a>

## decryption_provider property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 5292986458d0 / 4

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

<a id="canonical-7579899129bbeec29a688af7b6c7631d81d157406bb9c97fe2fa10620dbbc6c6"></a>

<a id="canonical-f1910a62e9a967d9468fd17c8300d3aecc9e20a704409122c2f56521d648d3e9"></a>

## location property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 5292986458d0 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-bd32bb193373c41855f745319dee12ff424d5999ef51e7f41639846138625fc7"></a>

<a id="canonical-decd049bca0a5d329e95b95dd03eec444a798ee19b2b773d35c498eebcf4d10e"></a>

## store_provider property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 5292986458d0 / 6

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

<a id="canonical-a29c271e309ea6bfa2a0e1c8da6097af663eb38931ec67e3588ce3f26a3462c8"></a>

## Next pages — routes.simple_route.advanced_options.request_headers_to_add.secret_value.blindfo / 5292986458d0 / 7

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-6b9d2e2292f2f72874f434d8d4889c8df05c12a3e51dcf3ee29ba78108fbc386)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ccffc91b1fcbfedfc77a65102bf48d33d45ae90a5f62e00ce5f82b619554261b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a035c169ae2f19403835b2af8318f47aa2f53a5e5dd459f075b03068465fa8b"></a>

## routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / dbfc808a757c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.request_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-9be9557dde347d2723c87163cf37b8d7322baf0f5e2adb7a120366f63e6a90a6)
- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-6b9d2e2292f2f72874f434d8d4889c8df05c12a3e51dcf3ee29ba78108fbc386)
- routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-73573c879393b06fb41b5221b883ac159fe1e981add31db5a7bd926e894e7f83"></a>

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

<a id="canonical-71e025a5589c1c640ace950c6f37f0e2e42bf8ac57ed54e4e6ca5ab1020afe65"></a>

## Direct properties — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / dbfc808a757c / 3

<a id="canonical-2bb12f884b411c15eea613b66f5803dbb144d36666df2f30f1488008de1e98f4"></a>

<a id="canonical-a0123daf7875d4758d522219e65d83bf0f6329c73bd7628659e542a3df0698f7"></a>

## provider_ref property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / dbfc808a757c / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b7c0330f87f708aba2518ad551c904c2689c5d78e2f425bdf708dc45fc2a72a4"></a>

<a id="canonical-344aea0a07d2d564bdac8723b051f095bfdc90621710aa9fd140b72aed466da9"></a>

## url property — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / dbfc808a757c / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-7c7fbb3dae8cd52b6cf817492445ce2020e61ec350a79ed907c862b0300d46c0"></a>

## Next pages — routes.simple_route.advanced_options.request_headers_to_add.secret_value.clear_s / dbfc808a757c / 6

- [routes.simple_route.advanced_options.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-6b9d2e2292f2f72874f434d8d4889c8df05c12a3e51dcf3ee29ba78108fbc386)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09881c814bb8134741d1608eb86cfaf384f87a74129d9430070fb96759db37ac"></a>

## routes.simple_route.advanced_options.response_cookies_to_add — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.response_cookies_to_add

<a id="canonical-81a526f0fe34b1acae5005834a2ce4d3bdfd49fe11a193d366fa77d21f191e0c"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3b2675a35f7e830f3b820d2580753af2295dd86a4279989b6126bd192d0d386b"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 3

<a id="canonical-1a419b64145461b355a0c0120b02aa67fdc95472564f0df22a0f64b18e9b04a2"></a>

<a id="canonical-0a749bd2142cc46950cb9bda0c6862747d50faf219447aae8af45d184f869087"></a>

## add_domain property — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 4

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

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

<a id="canonical-6baaaa736d65da135c6cb1c80d1b12bed042be0b25651c65a9bfa2dfa56772e2"></a>

<a id="canonical-5559fd1c603bfd897f07bee56f64d7944ace504f70243b8801911d1b0d4db5b4"></a>

## add_expiry property — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 5

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [add_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-9ee65872a4fb99483b84ec08777a737ad5a3e5b3dd6234af65e151e24e2c3593): complete subsection reference.

- [add_partitioned](data-sources--http_loadbalancer--reference--group-024.md#canonical-6b887573d27dde306ff6d8c9feaa9f945c0ade312d3cbc82a62b0740540cd3a0): complete subsection reference.

<a id="canonical-798738d8b54e7b78b539f24a95b2d37c8d79f7f638ebee895d8d03037d45eb24"></a>

<a id="canonical-7be00ec84aea5ae9310e91e78b68b419228c27f604e9571a81ddccdb3c95396f"></a>

## add_path property — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 6

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-340d0fb5f92f79dc102cf29fad282c02d79e2c61c130a3989ebdedecb1cd6315): complete subsection reference.

- [ignore_domain](data-sources--http_loadbalancer--reference--group-024.md#canonical-166539112312c57d15adcf83038d53c225ab83a2a52dc74354236409f2de23e9): complete subsection reference.

- [ignore_expiry](data-sources--http_loadbalancer--reference--group-024.md#canonical-88a8453f72518f2f73f878808c516cfc59803ad1af9b475d1ddb65d1098ee3e6): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-748e9369f74cf0bb2894e69a4797417967bd1eb8cbd5d31435ac6c3c34ea1e6e): complete subsection reference.

- [ignore_max_age](data-sources--http_loadbalancer--reference--group-024.md#canonical-6d0234382b806ca49e24ac9cf19d97d252642064c0b8d779a29ff7623209b37a): complete subsection reference.

- [ignore_partitioned](data-sources--http_loadbalancer--reference--group-024.md#canonical-5a3a9bcda174109c8d2ce6f22266b1aa081f46e07755a50363fc4901e8ce6fbb): complete subsection reference.

- [ignore_path](data-sources--http_loadbalancer--reference--group-024.md#canonical-c8a9e0a52fcdbf9a2f5e4c0de658cf7349b5a988c838bbb89b134b65249aa7e2): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-024.md#canonical-d2cff4c62a17ade398ca74ca191a897e8ed3a07abe17f316771db8f22d17eab4): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-4f88929bb349f64019a4f1b778f30b145a68c7721ddf805257582cace8eef03a): complete subsection reference.

- [ignore_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-76bb0a9fd556d2fad7545c86e42a641633a4ca6ef1e506352f0cb66c2707e568): complete subsection reference.

<a id="canonical-7b789f55c6f0a67ef92cb7ba7cb5e1d20f3b729ea0d35c75c8420ce3b4da9342"></a>

<a id="canonical-63665d794b781af2d0464afdcf9888d82b7e401aa56ee6c0877c562d7ff0ca88"></a>

## max_age_value property — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 7

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-07cf8c0fb92d45823ecc1f24f1112b3eb074d388d1584b8097f79a944f01a3fa"></a>

<a id="canonical-c1cc512fc076734d6cac94cdba04e42a3154e25251e3d4d8c8e7e07af4b808be"></a>

## name property — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 8

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-09431225e61c6f4cb4a464aec3f10b02e2532f74c3dd05b0b7735b7fc6f65edb"></a>

<a id="canonical-febf3a0d28f253a454bcb0e1ba6165f669b586c02f3507ba2d52166da380009c"></a>

## overwrite property — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 9

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-024.md#canonical-74d282293302a999b3d04648d48f075e5e583363b8beeb638a5e8e77c1937b6f): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-024.md#canonical-93a409efe9b98d09c94640365408f3b22e731b0ac4ae2169309a7909cd5990db): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-024.md#canonical-d86bd971150a1d08147f1f634e191704848602899c4bb37660b2ce664abfd5c6): complete subsection reference.

- [secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-5cb38735a2aac1904ef3b03445be8fb7b4d82c44c758667d5b5edfab547b3d72): complete subsection reference.

<a id="canonical-eeb230d8e496c9e8425c5cffd09951a5d7afbea0beaca9d192e6b5319ee4fa98"></a>

<a id="canonical-409c8665470ed47d47e8b6b5149f750467cfd4ce26800ae76f2cdf686cd1b659"></a>

## value property — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 10

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-978ebc0082b30c490bb2e0ea6c77e59851f7f7fa5e09bf3e062d533532b84bd8"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add / b9cf028eeca7 / 11

- [routes.simple_route.advanced_options.response_cookies_to_add.add_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-9ee65872a4fb99483b84ec08777a737ad5a3e5b3dd6234af65e151e24e2c3593)
- [routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned](data-sources--http_loadbalancer--reference--group-024.md#canonical-6b887573d27dde306ff6d8c9feaa9f945c0ade312d3cbc82a62b0740540cd3a0)
- [routes.simple_route.advanced_options.response_cookies_to_add.add_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-340d0fb5f92f79dc102cf29fad282c02d79e2c61c130a3989ebdedecb1cd6315)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain](data-sources--http_loadbalancer--reference--group-024.md#canonical-166539112312c57d15adcf83038d53c225ab83a2a52dc74354236409f2de23e9)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry](data-sources--http_loadbalancer--reference--group-024.md#canonical-88a8453f72518f2f73f878808c516cfc59803ad1af9b475d1ddb65d1098ee3e6)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-748e9369f74cf0bb2894e69a4797417967bd1eb8cbd5d31435ac6c3c34ea1e6e)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age](data-sources--http_loadbalancer--reference--group-024.md#canonical-6d0234382b806ca49e24ac9cf19d97d252642064c0b8d779a29ff7623209b37a)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned](data-sources--http_loadbalancer--reference--group-024.md#canonical-5a3a9bcda174109c8d2ce6f22266b1aa081f46e07755a50363fc4901e8ce6fbb)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_path](data-sources--http_loadbalancer--reference--group-024.md#canonical-c8a9e0a52fcdbf9a2f5e4c0de658cf7349b5a988c838bbb89b134b65249aa7e2)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite](data-sources--http_loadbalancer--reference--group-024.md#canonical-d2cff4c62a17ade398ca74ca191a897e8ed3a07abe17f316771db8f22d17eab4)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-4f88929bb349f64019a4f1b778f30b145a68c7721ddf805257582cace8eef03a)
- [routes.simple_route.advanced_options.response_cookies_to_add.ignore_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-76bb0a9fd556d2fad7545c86e42a641633a4ca6ef1e506352f0cb66c2707e568)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax](data-sources--http_loadbalancer--reference--group-024.md#canonical-74d282293302a999b3d04648d48f075e5e583363b8beeb638a5e8e77c1937b6f)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_none](data-sources--http_loadbalancer--reference--group-024.md#canonical-93a409efe9b98d09c94640365408f3b22e731b0ac4ae2169309a7909cd5990db)
- [routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict](data-sources--http_loadbalancer--reference--group-024.md#canonical-d86bd971150a1d08147f1f634e191704848602899c4bb37660b2ce664abfd5c6)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-5cb38735a2aac1904ef3b03445be8fb7b4d82c44c758667d5b5edfab547b3d72)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9ee65872a4fb99483b84ec08777a737ad5a3e5b3dd6234af65e151e24e2c3593"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76ab1e42e3900602f3e2e68efe379736cbaf32b1cd32edec63a3ffd64079e27d"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.add_httponly — routes.simple_route.advanced_options.response_cookies_to_add.add_httponly / 487e290b2bfd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.add_httponly

<a id="canonical-2490a7f5266562300ebe447d72ae0c81b245e667bf4e6a2279931d70ce988f22"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-0123cb1768d8ef26f5f4d41d753d735311b19b3796ffb170b5ee82797adbdee8"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.add_httponly / 487e290b2bfd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2460fc42944a608cfc01703277ca3b7575eaaed6717352c8f8c6b9c9008fd35"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.add_httponly / 487e290b2bfd / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6b887573d27dde306ff6d8c9feaa9f945c0ade312d3cbc82a62b0740540cd3a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a812e49a1e3e59d2a3f13ecdba32eb95bb4f60e49a6821a5af8f18ab3cfa848b"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned — routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned / e3fc2e8c1e13 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned

<a id="canonical-6acace8d929cfbda39213295e4769e297a7134c9520ba50d3365da0c7c6d57d0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add partitioned.

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

<a id="canonical-d207703e5cf71c7844790ce225ec97f86288f126381fd029c480dc77c99932f2"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned / e3fc2e8c1e13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a69a2f49fd80c5880b4d2ebc72ba48cc55048a7b2ae88b82f4c92e4a318f7282"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.add_partitioned / e3fc2e8c1e13 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-340d0fb5f92f79dc102cf29fad282c02d79e2c61c130a3989ebdedecb1cd6315"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-806540f33aa868302b02361ec2198c6a001a22d82040186e64c65ec411a9ae62"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.add_secure — routes.simple_route.advanced_options.response_cookies_to_add.add_secure / ca2fac4c433e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.add_secure

<a id="canonical-4f2ed15eedd8751aa1b13864ec311a92ffd3adf0b8efe929abc7f52a26080bc3"></a>

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

<a id="canonical-394f3eca42b5afb51196cc08e44b34614b1bf0b58e4e74bb10d274ab1663a3f9"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.add_secure / ca2fac4c433e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dde68fb39548d21556c51f59a186bd2462bd162a0f265473b523d64ceb39c8b6"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.add_secure / ca2fac4c433e / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-166539112312c57d15adcf83038d53c225ab83a2a52dc74354236409f2de23e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d193f26a5adb64df67728bd88431ed31f910fb0f57236c48d096875b6f7ebc94"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain — routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain / 53f066b52f71 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain

<a id="canonical-4f9e4eaf6ac3430d8407ceefb1b212325066ead4507b84b51d9504c21cd83006"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore domain.

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

<a id="canonical-f0bcc6eed9ea68e5cf90d02eac02f398063e76df31c7a342e710d7621eb8d67b"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain / 53f066b52f71 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0779216795efa8febb238e0bca866f980e400f2398afddfade78f604e8371885"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.ignore_domain / 53f066b52f71 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-88a8453f72518f2f73f878808c516cfc59803ad1af9b475d1ddb65d1098ee3e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-989e62a9b70f9dbd80b8712a7020096098b1af2446264d914ac761b7215f4e2c"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry — routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry / 713d3a07adec / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry

<a id="canonical-977a27822f21cb4046c96a4eb94d83ee1d026d61bc51bf1e861f9dfdbb2abaef"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore expiry.

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

<a id="canonical-904256e8e5e471ca3e872940a55c3f438de5af26dd0ee009732a69e95c449627"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry / 713d3a07adec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f52bf69e25702da15912bd90c651009de842c4a0dcf19a389c745c244401ffc5"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry / 713d3a07adec / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-748e9369f74cf0bb2894e69a4797417967bd1eb8cbd5d31435ac6c3c34ea1e6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-240e9659c226f5a7245c0173323d22faa2f4aeb1e519f41c479661e602e820e1"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly — routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly / b5116dee47c1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly

<a id="canonical-76c61288a3369ca0afcbf77db7a3b1f025ee958fddc35024a20b968c168085c1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-a5f9b7bd3107dc21d17b43f06fb12dfd0a4fe2147e1a7e8063a8320920f448df"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly / b5116dee47c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0beeb1b59d8c932f5db3cdae6ce02882d050e5ae61ea91bd6afade2f29a14c46"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.ignore_httponly / b5116dee47c1 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6d0234382b806ca49e24ac9cf19d97d252642064c0b8d779a29ff7623209b37a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5f942cfa10306acb687c6b560894913ff8d6105c1b5c30ce4c957f640da17aa"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age — routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age / 8ce0b7af885c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age

<a id="canonical-d37952457cd88ef109d1e74223b83416be2e9d182a9e5be86ef0e9af4570ff5f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore max age.

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

<a id="canonical-7c7bc5e5922459e8a7ea01617853106bb096086d9807ac90d0629e247739ea43"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age / 8ce0b7af885c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d7b50417825d6577ffd992d563acc2445da0aa9638a0b09afba2c6ca05f17d0c"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.ignore_max_age / 8ce0b7af885c / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5a3a9bcda174109c8d2ce6f22266b1aa081f46e07755a50363fc4901e8ce6fbb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99206d50dd9f8df255d49e6c6cde679dd530c5a59aa89a3dfeadfd1c87a3fcf6"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned — routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned / 25fdf535954b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned

<a id="canonical-f9f2dcf1e0253270592de2c2e71d59ea6b13dd0d6ad306864aeec8be2d0ecb48"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore partitioned.

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

<a id="canonical-54b4370b56b7af53e547a1d1f6a1b4dbed1d51c6be6c8220183d967bc580a1ce"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned / 25fdf535954b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c5f345117e79656ed82c067cd6706893748f026e0ba7dfe6b5d2d671ebb99249"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.ignore_partitioned / 25fdf535954b / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c8a9e0a52fcdbf9a2f5e4c0de658cf7349b5a988c838bbb89b134b65249aa7e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fae37b62d48b186023449bea505be1d45751c904d561934c6ad124b83c91f678"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_path — routes.simple_route.advanced_options.response_cookies_to_add.ignore_path / faeb6d856f7b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_path

<a id="canonical-b1bfb9d4315199b5e429b38fc511e374eefeaa9e2234701ffb6d2d4b7a8c4e79"></a>

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

<a id="canonical-6eb9beee0394ac5a94beb2112006d6e1a36172539e4bea8692c96f658f190683"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.ignore_path / faeb6d856f7b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b04c5e82df0fa7637df417f7120100170eca56bd330e61e478d57d0fe28d7821"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.ignore_path / faeb6d856f7b / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d2cff4c62a17ade398ca74ca191a897e8ed3a07abe17f316771db8f22d17eab4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98b0c395f4e89aba5fd1a127e30416150a62936f410a9579f915494e1e8b5a66"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite — routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite / 0743731632b9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite

<a id="canonical-7a283ada502423ea8821dd34107e5b7bf4c8182e03aa2151f10d33b517ea03c4"></a>

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

<a id="canonical-5f6a9c8172e815d8b6b41b294ff1b10e8527b056d3f3402fbd48e6900f2bf020"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite / 0743731632b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-79cfca32dbdb15c5f3ee4ff68c53c17b7a8b648e7d85a1153e3f145b920b2c58"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.ignore_samesite / 0743731632b9 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4f88929bb349f64019a4f1b778f30b145a68c7721ddf805257582cace8eef03a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4e948d170fef16f3500534b489a111758860fbbac1f458de4b05ce85359c45a3"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure — routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure / a1a3e6100134 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure

<a id="canonical-2c27b20768fc2620846aca6019442caf0d6cf71af51f6d4978408b0dc9a4d905"></a>

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

<a id="canonical-6f5ae577e30aca464aac3ca0e97231f953330220923f085866ce761ce2730554"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure / a1a3e6100134 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-720c8ec9fd54dc0eb6b55fd2107eb1295120b4543d15bf2a96658b8256a16f56"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.ignore_secure / a1a3e6100134 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-76bb0a9fd556d2fad7545c86e42a641633a4ca6ef1e506352f0cb66c2707e568"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdf10255ed13eaaa9747a86e551ca642e1b3dce459d01157ba5122dbd1523d71"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.ignore_value — routes.simple_route.advanced_options.response_cookies_to_add.ignore_value / 8ceb8ab47310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_value

<a id="canonical-0733e93d187632b867b3906d5b400477813041585edd53e003ab1d22db7813e4"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore value.

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

<a id="canonical-8d0c23322c63ff8dc164007f4a4b6807351bade14a9025c5917b91cb60d386d4"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.ignore_value / 8ceb8ab47310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb5514304e12003c72bae3c78ee80eb6cf82d9eb54e2602c272577431b511658"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.ignore_value / 8ceb8ab47310 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-74d282293302a999b3d04648d48f075e5e583363b8beeb638a5e8e77c1937b6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-375cc00edb27b8593455ed077e1069e764471e465a2ac61e6992ac290a9804fb"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax — routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax / 0c7d875cb255 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax

<a id="canonical-752fe37c9b7079a5f0d8669ef672cd780d5c1027c96016c1a10ee918484bf8e4"></a>

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

<a id="canonical-f6791cdd14dfec4297e8fae3b6a870ca78b852d8becc4b07c033c7111f5bba04"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax / 0c7d875cb255 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16fab024526637cb7e7b13892fa0be2bb49238c4d5fa125bfb29a91d30b5973f"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.samesite_lax / 0c7d875cb255 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-93a409efe9b98d09c94640365408f3b22e731b0ac4ae2169309a7909cd5990db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76a3f352daf713ee72004f1ee80be894e3b0b9226ff2bbe678214326c79e15e4"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.samesite_none — routes.simple_route.advanced_options.response_cookies_to_add.samesite_none / 91b82dc02787 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_none

<a id="canonical-d6c565f43c249b37b5b52d520c1ab088e1d7f1fb29cf5409927d6690a6e3ea43"></a>

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

<a id="canonical-1dd1c031af45c2e003e8340b3cb299b5802e5e5eb9b23e8677cd4b39f70dde9d"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.samesite_none / 91b82dc02787 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c706c203d49cedab3f2f26ab9692e3faacc36d1ef84b12417612630b4fa31034"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.samesite_none / 91b82dc02787 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d86bd971150a1d08147f1f634e191704848602899c4bb37660b2ce664abfd5c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bf17a1a1083a4444ded7ed7dab37df39ebcd53c15faafd7325a6c401e08a364"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict — routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict / a622eca28f17 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict

<a id="canonical-f326d4e0dcfa01e130fc248486d0f7577a964e28f02b15caf4778ddc7bfae7f7"></a>

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

<a id="canonical-4ac383990e5ac73fe7898c4b18c1011527c6e41381258a67e0b772930e8974f9"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict / a622eca28f17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fbad7d5c9bdf330e1af98bd75f3d8c880487e51a490538b0b98fc3622f3767c9"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict / a622eca28f17 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5cb38735a2aac1904ef3b03445be8fb7b4d82c44c758667d5b5edfab547b3d72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fff5db876befa4d4160494e76b98f7fee52ec934003fa85359eba408605f1a0f"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.secret_value — routes.simple_route.advanced_options.response_cookies_to_add.secret_value / 476672d6ab28 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- routes.simple_route.advanced_options.response_cookies_to_add.secret_value

<a id="canonical-736127c4c556ff5efd6faa8debcf33a6944891cea170d471d1539daf9051c759"></a>

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

<a id="canonical-9e37b792af01ea7b87e9d16ba1bc94bcc7371a1beabe7e2844c6e7370c67900b"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.secret_value / 476672d6ab28 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-72fb5a40380a3cd0a994c5a119bacaa4f0d30dd6aa34ab02f006f82ec8813770): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-d5eeda62b5251e880c18bc5a524276214e255afa84edbd6e2a8554e009a585e3): complete subsection reference.

<a id="canonical-586b91c25cd0507225209d6b1fa0654b38ddf594d0d50133dd33768ab002a87d"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.secret_value / 476672d6ab28 / 4

- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-72fb5a40380a3cd0a994c5a119bacaa4f0d30dd6aa34ab02f006f82ec8813770)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-d5eeda62b5251e880c18bc5a524276214e255afa84edbd6e2a8554e009a585e3)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-72fb5a40380a3cd0a994c5a119bacaa4f0d30dd6aa34ab02f006f82ec8813770"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-943319aa233436f959b193c8e5355d13d5b74ca4373c679e4b937641b0bc6151"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindf / e28ba0d2616b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-5cb38735a2aac1904ef3b03445be8fb7b4d82c44c758667d5b5edfab547b3d72)
- routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-f9321bf7c7771b47684c989bf201181d38aaf903e61e2fec3322efdd289b2870"></a>

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

<a id="canonical-d3766ddf7b66011eb24bb639944cc0ababcaf364d8e896d64de344193e3c26f5"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindf / e28ba0d2616b / 3

<a id="canonical-7048d71b9d257250f0c49d846294fbefc335a50ae41899ef9ec06a641b29dec0"></a>

<a id="canonical-a08d96d0e444029a4a02627a15f64ee0f1597b473c8aec67790006a4a753093d"></a>

## decryption_provider property — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindf / e28ba0d2616b / 4

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

<a id="canonical-cf13506891a620963ad27653d5a714ee2380af4936f2aadf52d058ea4fb79860"></a>

<a id="canonical-75897833383cce3584b90e036d4aa35e3df83b2d021e617313d26db7f265b587"></a>

## location property — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindf / e28ba0d2616b / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-649120dcd88dedf6c6f972037d91fd07432fb2ebad3390335d8b28e06d6a5ff7"></a>

<a id="canonical-a06d54b6203f7ae262cb0749fb26d65214bb3c5e2a9c8a223a8d6d90183ece38"></a>

## store_provider property — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindf / e28ba0d2616b / 6

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

<a id="canonical-f4ef10580241a88bdba4bf865f647dfa07a9824b52e6438aebec019a458df714"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.blindf / e28ba0d2616b / 7

- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-5cb38735a2aac1904ef3b03445be8fb7b4d82c44c758667d5b5edfab547b3d72)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d5eeda62b5251e880c18bc5a524276214e255afa84edbd6e2a8554e009a585e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3eac83937ee30e2f620033d67e1818c0d8e3c0643fed130116ecb4ce12c4a82c"></a>

## routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_ / e633869348de / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-5cb38735a2aac1904ef3b03445be8fb7b4d82c44c758667d5b5edfab547b3d72)
- routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-f9e6429281a6fd97bb0f49c53e500c0b3a904d1b5400e678fddbe94ca7238543"></a>

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

<a id="canonical-cb367adf830405898bec4eb8a391578233c3796fcee45d5b7f36438a6129bc27"></a>

## Direct properties — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_ / e633869348de / 3

<a id="canonical-527e3bbabeb94d84849cb99ecb88765890c6f2b402377e1e35316ad42347286b"></a>

<a id="canonical-8fd3bd725922af56bb439d7d096aa2edf079c800fe9ed788ee5e8d1f271b9cf1"></a>

## provider_ref property — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_ / e633869348de / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-b39a782b95606f75be1b68adecdd1dcfba9bb5dd1dffa5491a78b657378a7789"></a>

<a id="canonical-7712c917dd13436da7e5dfee193c9db61d03d7a98ce2ef698c5492197b07e336"></a>

## url property — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_ / e633869348de / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-261f18e671cffb881acc417452d6a9cdd65a39545dedbb6731fea479d3313742"></a>

## Next pages — routes.simple_route.advanced_options.response_cookies_to_add.secret_value.clear_ / e633869348de / 6

- [routes.simple_route.advanced_options.response_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-5cb38735a2aac1904ef3b03445be8fb7b4d82c44c758667d5b5edfab547b3d72)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e30971933d5007464e0b821668a3ae01b4a90a73b358f07cb04984decfa8ca95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6eede1ace66942ae2864239d3abb3e0737fc45c2757b0e788c5f878b1d3c51d"></a>

## routes.simple_route.advanced_options.response_headers_to_add — routes.simple_route.advanced_options.response_headers_to_add / b931f026772b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.response_headers_to_add

<a id="canonical-115adb81e10a2be4aa20b7de79fe6f09720b74f7d9c955b5c3c7c087393fcbfa"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-4fd39d80b32215e38ab0da699e428e9becf14d3af12b7eedcf40a51778a7b78b"></a>

## Direct properties — routes.simple_route.advanced_options.response_headers_to_add / b931f026772b / 3

<a id="canonical-b21be0bb017aa1f1bf4d2d7e4051f23afe42e689c476eb10c4227564f3bf0d80"></a>

<a id="canonical-487b0e9b58bddb4d4ccf82fffea18f3c6a65200edb78372664ac4d0b869efca7"></a>

## append property — routes.simple_route.advanced_options.response_headers_to_add / b931f026772b / 4

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9bdc7614cad82528d76fa2be30a396c9c34a7d3b72a4acda43b2bda32ec5646a"></a>

<a id="canonical-cf7a2367193f79635d71ac0330d00425cb20b1bd53b52cbaa4c5ba0a76a4354d"></a>

## name property — routes.simple_route.advanced_options.response_headers_to_add / b931f026772b / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-368cac61c9e10b18a1830ddae025864f8d9cf951c01c54b3ec978db8bfd4bffa): complete subsection reference.

<a id="canonical-f8779f8a958397ff805d2872c9cd5fc9d403277a12f5a8184a690f8b65ab4369"></a>

<a id="canonical-973b2aba9bf5a64cdbcab9fb3a32fbc59e7bcb3df569f467d064a6ee5fb555fc"></a>

## value property — routes.simple_route.advanced_options.response_headers_to_add / b931f026772b / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-173c5adb789559b10b393eb9aa0c205f8ee650bd2faeb356272c4a4066d6c9a5"></a>

## Next pages — routes.simple_route.advanced_options.response_headers_to_add / b931f026772b / 7

- [routes.simple_route.advanced_options.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-368cac61c9e10b18a1830ddae025864f8d9cf951c01c54b3ec978db8bfd4bffa)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-368cac61c9e10b18a1830ddae025864f8d9cf951c01c54b3ec978db8bfd4bffa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f42b6b31287f36e4a2604e362cb32b830f39c3a12aa6af80f5859a65f543ac4"></a>

## routes.simple_route.advanced_options.response_headers_to_add.secret_value — routes.simple_route.advanced_options.response_headers_to_add.secret_value / 91ad01346bbf / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-e30971933d5007464e0b821668a3ae01b4a90a73b358f07cb04984decfa8ca95)
- routes.simple_route.advanced_options.response_headers_to_add.secret_value

<a id="canonical-f9296315f4747ce263b7604a4501559c91624d449b4b57c40bccc80990333cb0"></a>

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

<a id="canonical-e2b9c16822b7efa5e555b922c1b5d584552af5ed54382e894de834fc44d8a9f7"></a>

## Direct properties — routes.simple_route.advanced_options.response_headers_to_add.secret_value / 91ad01346bbf / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-bd07353b33e91d4c2048b4233db28c38130d23c6a11ea842fe1aa18a0dc0348c): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-49670e1210956510ed9664bbdaaed35235d51be2c6ca348964dd4addbd37251a): complete subsection reference.

<a id="canonical-a2f0d2ec20f78b24b94f6326394d33152dffdb5e795494dc9f8f4d339a2b3b14"></a>

## Next pages — routes.simple_route.advanced_options.response_headers_to_add.secret_value / 91ad01346bbf / 4

- [routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-bd07353b33e91d4c2048b4233db28c38130d23c6a11ea842fe1aa18a0dc0348c)
- [routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-49670e1210956510ed9664bbdaaed35235d51be2c6ca348964dd4addbd37251a)
- [routes.simple_route.advanced_options.response_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-e30971933d5007464e0b821668a3ae01b4a90a73b358f07cb04984decfa8ca95)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bd07353b33e91d4c2048b4233db28c38130d23c6a11ea842fe1aa18a0dc0348c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c19e8cde3663f6662659c1c1eecc3a7a12b173046b87aa1c09c827aa43d5f641"></a>

## routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info — routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindf / bf2db548389a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-e30971933d5007464e0b821668a3ae01b4a90a73b358f07cb04984decfa8ca95)
- [routes.simple_route.advanced_options.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-368cac61c9e10b18a1830ddae025864f8d9cf951c01c54b3ec978db8bfd4bffa)
- routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-9d6ecb5cab859f70e2c8b0f08919c062bd3908e8f5f8f747120d43e4e40b9966"></a>

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

<a id="canonical-e77461602dfc48e4c85d6065244a711808f83fe904d0bfb94a50e2aa30388ce1"></a>

## Direct properties — routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindf / bf2db548389a / 3

<a id="canonical-4c2edb24ccaec400bd62020f056758a8e70d85089f788363885ad2446dde2530"></a>

<a id="canonical-f28be4d8d4f488c6fcc676c8ebeeb4ff635965cb7b2f594d06c74297a00443db"></a>

## decryption_provider property — routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindf / bf2db548389a / 4

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

<a id="canonical-9ce12a8728626de88b41c8b76ed1885c11b29143190aa99c1e138864bb4ab94b"></a>

<a id="canonical-7aeca4287aa44debae094061cb0db4555ecee07864e6c20f8c89fc6cb0e9e2db"></a>

## location property — routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindf / bf2db548389a / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-f9b364463e4d5c51930ce82a107c1beb94fbd606a4d74c90407d8f9cdc01e554"></a>

<a id="canonical-bcf61af8a6d0c63cb4398aab53198fe6d497d3e782bb904721a69ca211082a80"></a>

## store_provider property — routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindf / bf2db548389a / 6

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

<a id="canonical-d40c8288bd2cf1e08607dacc02958fae55bcd6cd2b06fefa543c459f13759b21"></a>

## Next pages — routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindf / bf2db548389a / 7

- [routes.simple_route.advanced_options.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-368cac61c9e10b18a1830ddae025864f8d9cf951c01c54b3ec978db8bfd4bffa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-49670e1210956510ed9664bbdaaed35235d51be2c6ca348964dd4addbd37251a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a61269210aac4909341b348d6e644f42d25b461704e82e34a8a2403cd745ea7f"></a>

## routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info — routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_ / 50298be5c47e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.response_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-e30971933d5007464e0b821668a3ae01b4a90a73b358f07cb04984decfa8ca95)
- [routes.simple_route.advanced_options.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-368cac61c9e10b18a1830ddae025864f8d9cf951c01c54b3ec978db8bfd4bffa)
- routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-20926afbeec639924c7e473245d450ff587688c54e68ff123692a46c5f623357"></a>

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

<a id="canonical-4e581acc6522688262d36470716ebeb2f1f44eef9a183fbf951d4d0c062610ab"></a>

## Direct properties — routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_ / 50298be5c47e / 3

<a id="canonical-37aa6c95ed15e34d3b70f9fd2c4107b63c4217a28e063e8868111fd9c2b0b171"></a>

<a id="canonical-307c43009e209606909a8303b82814abb53402fbfba110c1a44cf3bcfe0e749e"></a>

## provider_ref property — routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_ / 50298be5c47e / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-34d6ce2050b83111a43bfc17c44014cdaa041c78688ea3ab86b155cc200ee4f4"></a>

<a id="canonical-1b8c748dbb546694e8a6e0f383f4d89d3c2ea698d0862fe4d8f1e3184a9416d7"></a>

## url property — routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_ / 50298be5c47e / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-e9f32e85cced8b408cdd0b3558fdacf06265b35ca84165dfb96c82cd1a64334d"></a>

## Next pages — routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_ / 50298be5c47e / 6

- [routes.simple_route.advanced_options.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-024.md#canonical-368cac61c9e10b18a1830ddae025864f8d9cf951c01c54b3ec978db8bfd4bffa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-629d3b101c247273fffdb2a8e94adacd50a906396d477f73469b9afa7539e6b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20e3dfe5d23831657cec0ccfdc0fa93df09a7ebd5cb5743806e7c9f6c81171e9"></a>

## routes.simple_route.advanced_options.retract_cluster — routes.simple_route.advanced_options.retract_cluster / 92b51166c67a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.retract_cluster

<a id="canonical-d4a58523af6536285583bb161977984297f4ba45dbb03026fd94688712e0236a"></a>

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

<a id="canonical-493139d8142ac61714c9e803a572903b7e2195d46d7029bdcaf61a2721b6097e"></a>

## Direct properties — routes.simple_route.advanced_options.retract_cluster / 92b51166c67a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7cc6f242db5bb6a5f119c38029759f978c2cefeff0de6d47e604274f9843ba4e"></a>

## Next pages — routes.simple_route.advanced_options.retract_cluster / 92b51166c67a / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c4fc0ead7e6779144e37321ac5084409dc1876214277674b2e7da8b00f43db70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5d72e173fa5858011b285c4cb583622b5f55160222387a9859c2b046bb2f0dd"></a>

## routes.simple_route.advanced_options.retry_policy — routes.simple_route.advanced_options.retry_policy / ea86b6554117 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.retry_policy

<a id="canonical-4f6026f68255eb9da4e60f61e9e5a095bd27b174257c6e73346266a859f9cc18"></a>

Type: `"single"`. Computed.

Retry policy configuration for route destination.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-81135fd8911dededf5cfbb64920e78c85f9bbc675caf98c4d921c2483031ab99"></a>

## Direct properties — routes.simple_route.advanced_options.retry_policy / ea86b6554117 / 3

- [back_off](data-sources--http_loadbalancer--reference--group-024.md#canonical-11d619861ec7534831eea5550b04e07a20614e1ab560b39133b61dc6aeb7b81e): complete subsection reference.

<a id="canonical-e0d091f76c8e069d0b8664c2f8ab1e0e5d407df49a3f0e046bdf6257099c151e"></a>

<a id="canonical-14933880cae9a2c1512100164dbc62cc0c5591f9da6974ade72e98e1f7ef29c7"></a>

## num_retries property — routes.simple_route.advanced_options.retry_policy / ea86b6554117 / 4

Type: `"number"`. Computed.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Upstream description:

Specifies the allowed number of retries. Defaults to 1. Retries can be done any number of times. An
exponential back-off algorithm is used between each retry.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="canonical-5a59bf86ede857a3f6ea2dd40e8f8c74403db2c1e0cd6b96aa5365279bd4eddd"></a>

<a id="canonical-fa641fe41f84fc34ec52ce6bdff233911c4d556250264e8808fa0291ad4a3a7c"></a>

## per_try_timeout property — routes.simple_route.advanced_options.retry_policy / ea86b6554117 / 5

Type: `"number"`. Computed.

Specifies a non-zero timeout per retry attempt. In milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-160d06b478c228bca119ef210e3d5edc5479f20dcf83be928b33f63ad9a94359"></a>

<a id="canonical-d91fd2a065697b43dab7997cba1309d181918b9f4f47fa75180668022a8d828b"></a>

## retriable_status_codes property — routes.simple_route.advanced_options.retry_policy / ea86b6554117 / 6

Type: `["list", "number"]`. Computed.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a382c70530c47732582245067804ac7d5b42549023f338f07c7041dca47bfbe0"></a>

<a id="canonical-969066d53f2cd11d81f80c3542b1b37e74bd80065b9d9e5698c50a04e6012124"></a>

## retry_condition property — routes.simple_route.advanced_options.retry_policy / ea86b6554117 / 7

Type: `["list", "string"]`. Computed.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Upstream description:

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc

The possible values are

"5xx" : Retry will be done if the upstream server responds with any 5xx response code, or does not
respond at all (disconnect/reset/read timeout).

"gateway-error" : Retry will be done only if the upstream server responds with 502, 503 or 504
responses (Included in 5xx)

"connect-failure" : Retry will be done if the request fails because of a connection failure to the
upstream server (connect timeout, etc.). (Included in 5xx)

"refused-stream" : Retry is done if the upstream server resets the stream with a REFUSED\_STREAM
error code (Included in 5xx)

"retriable-4xx" : Retry is done if the upstream server responds with a retriable 4xx response code.
The only response code in this category is HTTP CONFLICT (409)

"retriable-status-codes" : Retry is done if the upstream server responds with any response code
matching one defined in retriable\_status\_codes field

"reset" : Retry is done if the upstream server does not respond at all (disconnect/reset/read
timeout.)

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 7,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 7,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e92ee761b2540c3c4f8385195ceb4d0f6b8d9cc975c2818660711e1d049256fd"></a>

## Next pages — routes.simple_route.advanced_options.retry_policy / ea86b6554117 / 8

- [routes.simple_route.advanced_options.retry_policy.back_off](data-sources--http_loadbalancer--reference--group-024.md#canonical-11d619861ec7534831eea5550b04e07a20614e1ab560b39133b61dc6aeb7b81e)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-11d619861ec7534831eea5550b04e07a20614e1ab560b39133b61dc6aeb7b81e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac7513f40334e4164cea3913e0f469502cb5f8bf2e085ebebfca48bef365affb"></a>

## routes.simple_route.advanced_options.retry_policy.back_off — routes.simple_route.advanced_options.retry_policy.back_off / 180cf3ed10c7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.retry_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-c4fc0ead7e6779144e37321ac5084409dc1876214277674b2e7da8b00f43db70)
- routes.simple_route.advanced_options.retry_policy.back_off

<a id="canonical-bd43ba3eabe646d4144877b61f33f690ec6c436136a29e504cb37e31a3db2b3c"></a>

Type: `"single"`. Computed.

Specifies parameters that control retry back off.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0c1f80f3089d1cb1c58a7804a3ddfdd37bd578069813fa79ac98a40878cb83af"></a>

## Direct properties — routes.simple_route.advanced_options.retry_policy.back_off / 180cf3ed10c7 / 3

<a id="canonical-623ee58706f11b47483f2e9ead4a41b7c6994a0516f1c19108764ad357de5e55"></a>

<a id="canonical-3a25facfcd084cf5bf8b6018dc9485dbb80ceebdd5c3e25f0f3839f3235cb926"></a>

## base_interval property — routes.simple_route.advanced_options.retry_policy.back_off / 180cf3ed10c7 / 4

Type: `"number"`. Computed.

Specifies the base interval between retries in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-4fac7db14bec5522f96f737cca731920e3ec9d326e81569bf0ca008633f49de2"></a>

<a id="canonical-281b06e9030874b42e2052135375d0fc4daf2137920e46ba42692bedaacd2b88"></a>

## max_interval property — routes.simple_route.advanced_options.retry_policy.back_off / 180cf3ed10c7 / 5

Type: `"number"`. Computed.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Upstream description:

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The default is 10 times the base\_interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b35db41becde773937d8840e3c359ae9b6e526bf18f3f4e0f777c895e5ad771a"></a>

## Next pages — routes.simple_route.advanced_options.retry_policy.back_off / 180cf3ed10c7 / 6

- [routes.simple_route.advanced_options.retry_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-c4fc0ead7e6779144e37321ac5084409dc1876214277674b2e7da8b00f43db70)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb0c47060b6bac6b16c75c56edd7667d41ca54628043a05b16d9277723212682"></a>

## routes.simple_route.advanced_options.specific_hash_policy — routes.simple_route.advanced_options.specific_hash_policy / 115a3ba3c842 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.specific_hash_policy

<a id="canonical-07f490b6430b61d4051665e903c5bfab0e9a6978001f2f1fec22b641a5bcfd6f"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

Upstream description:

List of hash policy rules.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d0317bbd107a59ab22d1adb6b335e2e9a7654066ce405ab4774cef7367b72280"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy / 115a3ba3c842 / 3

- [hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4): complete subsection reference.

<a id="canonical-8de0c965b39a303b75a8abbd0eab1cc64fb55487f13e83822fd054ed12d280fe"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy / 115a3ba3c842 / 4

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-294085d89584b4e5980fa69dfbb19e4abbeb0f6b075b0d1edfc786ef423ec43a"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy — routes.simple_route.advanced_options.specific_hash_policy.hash_policy / 4ee7b9a617a9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy

<a id="canonical-c26baa15240e168bc75a6d8c2941c9e5f49614cfda5d10ed762334005978f59f"></a>

Type: `"list"`. Computed.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Upstream description:

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8dba3832e6257230ac429b167f761c36d3b6044af8188bb55d23bcdab5f9a077"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy / 4ee7b9a617a9 / 3

- [cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2): complete subsection reference.

<a id="canonical-fcd98cb9f56fc3c2d2ff6d4028250543aca5112f934c6fc73c08b0b0097ae78d"></a>

<a id="canonical-15223e40a23afbebffe9ce5170feaf66417909485fbbd1a70ac2bde913b731f4"></a>

## header_name property — routes.simple_route.advanced_options.specific_hash_policy.hash_policy / 4ee7b9a617a9 / 4

Type: `"string"`. Computed.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

Upstream description:

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3afd4b75ad5b527ef005b3fa502d257532b5c248e0aad1f3954a0e6aaaffc26d"></a>

<a id="canonical-1366ce824a61d204fa314a5d54dca20b61003f77b76317993ca75e57609019d0"></a>

## source_ip property — routes.simple_route.advanced_options.specific_hash_policy.hash_policy / 4ee7b9a617a9 / 5

Type: `"bool"`. Computed.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Upstream description:

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-734e4d604ebfc261bc8f1c043ddad9f46b8a41fb65c053101e72b42a2eb54e88"></a>

<a id="canonical-7028a24c68c307217990216c6bedade409df64c625d8db181b6d9d575d8c766b"></a>

## terminal property — routes.simple_route.advanced_options.specific_hash_policy.hash_policy / 4ee7b9a617a9 / 6

Type: `"bool"`. Computed.

Terminal. Specify if its a terminal policy.

Upstream description:

Specify if its a terminal policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d82a4261531c049e29d2f852fb5faae780ec6e410f3dc297e9141281a3bbecd2"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy / 4ee7b9a617a9 / 7

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1620ccd0f3febc03e9dac6ae10517f4240d6792954080280d0bb1f7d5b93fc2"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie / e5701be7f890 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie

<a id="canonical-252ab7ee9acf157f0b1926fc3792286f7b02572ae954e77a10674922143294b0"></a>

Type: `"single"`. Computed.

Two types of cookie affinity: 1. Passive. Takes a cookie that's present in the cookies header and
hashes on its value. 2. Generated. Generates and sets a cookie with an expiration (TTL) on the first
request from the client in its response to the client, based on the endpoint the request gets..

Upstream description:

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

<a id="canonical-e4dc9a3574e1a77b7bd445f351b3c29a043bab7b928be30c429de176c4313af4"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie / e5701be7f890 / 3

- [add_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-5bc5eb041e60ee0d936d0d0df6da897325f5e9bed1729343437772cf4d2daea5): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-30cf49b2fb43583d404537d121004341f7e6c902bd542ac8e5218dfdc65c904b): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-c0405b8f07fb3d2d4eace6fcf2138aed0e188445a3e8389129166080587ec539): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-024.md#canonical-8ccdf9cd4a3d2389c01ce9179709dc83db061e5cda01473fe1babef7907118a7): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-a61daad1ed6bf9453ba0f597aa7f222af44e28a30aad6b794444f796c9096316): complete subsection reference.

<a id="canonical-6d95dc11bca720e204dbaa2d08ec6e6fa17c2bcd21fedd58cf2dabbc7352748a"></a>

<a id="canonical-1685a1c8d48f929b35ae3d70f7b2ab611eb3463f63885750eabecdb254d38922"></a>

## name property — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie / e5701be7f890 / 4

Type: `"string"`. Computed.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Upstream description:

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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

<a id="canonical-26bc3ab6988677e2dbe6c8b4871f72729c7a74ce710ab9481eef568dbcf174cc"></a>

<a id="canonical-8bcf79f7dc0c5127672a37988b5215758a3fca24620780c594f8932b708f67eb"></a>

## path property — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie / e5701be7f890 / 5

Type: `"string"`. Computed.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Upstream description:

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-024.md#canonical-8263f8070775f8b9500a4284e260228495fba7f168d18067afa97cddb1c8d404): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-024.md#canonical-3b5d8ddfbf31e2f9e48af8b76126df62bec316f7d0af3241cbe89b284a62f532): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-024.md#canonical-970bebd15250690408ada8e08b040de10494ef257921e090ec3071438874d57d): complete subsection reference.

<a id="canonical-81330e26b867612e8914bd097ba64ece19e6fb156a1a96359f33706f9da69d97"></a>

<a id="canonical-2af2bced9806b47019c92f3d466730d65bec9ccaa24b84faf58e9e413eec5ae1"></a>

## ttl property — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie / e5701be7f890 / 6

Type: `"number"`. Computed.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Upstream description:

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-a4c1b84fac6d171ca6ef65ac0c69218ea1f81902bf7452382c51731915f5adca"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie / e5701be7f890 / 7

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-5bc5eb041e60ee0d936d0d0df6da897325f5e9bed1729343437772cf4d2daea5)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-30cf49b2fb43583d404537d121004341f7e6c902bd542ac8e5218dfdc65c904b)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly](data-sources--http_loadbalancer--reference--group-024.md#canonical-c0405b8f07fb3d2d4eace6fcf2138aed0e188445a3e8389129166080587ec539)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_samesite](data-sources--http_loadbalancer--reference--group-024.md#canonical-8ccdf9cd4a3d2389c01ce9179709dc83db061e5cda01473fe1babef7907118a7)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_secure](data-sources--http_loadbalancer--reference--group-024.md#canonical-a61daad1ed6bf9453ba0f597aa7f222af44e28a30aad6b794444f796c9096316)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_lax](data-sources--http_loadbalancer--reference--group-024.md#canonical-8263f8070775f8b9500a4284e260228495fba7f168d18067afa97cddb1c8d404)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_none](data-sources--http_loadbalancer--reference--group-024.md#canonical-3b5d8ddfbf31e2f9e48af8b76126df62bec316f7d0af3241cbe89b284a62f532)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_strict](data-sources--http_loadbalancer--reference--group-024.md#canonical-970bebd15250690408ada8e08b040de10494ef257921e090ec3071438874d57d)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5bc5eb041e60ee0d936d0d0df6da897325f5e9bed1729343437772cf4d2daea5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5651ccb48ff5f1a7dbb8643d10341bfccb6e107f40daaaecbd0fbe77b85f47f"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_httponly — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add / 4f54ad33afc0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_httponly

<a id="canonical-bf54618e0656aa6cfd3a566be346aa5094b37da2f0636d2a574bdeab67196b02"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-c12763df3592080eff13570f0aad98c4835d237b6c0ca57a5b53ff483dceac91"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add / 4f54ad33afc0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b94b1e411e0e0e9c9f8d3d0d34433e2b43a28e8bcfab04f3f4dd02cc2ef5043c"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add / 4f54ad33afc0 / 4

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-30cf49b2fb43583d404537d121004341f7e6c902bd542ac8e5218dfdc65c904b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2f745376d2b8234ea76ed8f3edfe30b7003d8f42dbcb95867121052c59d7294"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_secure — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add / 6a1d991e7696 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_secure

<a id="canonical-9da15484c15d58faccce268e46f9f83db693adee455bcd57b5a6e06f54b95c70"></a>

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

<a id="canonical-2d9a493e932e987cfb01b66ddb59419f5d303e82a4b51d0e94130e6123ae9d67"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add / 6a1d991e7696 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc23119feabf9b11e7d14b3f2e1546a2c835ac0fde827f602914cd9f08636403"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add / 6a1d991e7696 / 4

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c0405b8f07fb3d2d4eace6fcf2138aed0e188445a3e8389129166080587ec539"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac40f19b3c54829777cdff05b277634c63351b77c9551da72e13d3cc7c87e276"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ign / ff973b12fbc9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly

<a id="canonical-89b662e5948dfa274d163fab00e79fa50d46cc403641e3a0cda4f274a08ef384"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-a4ad39bc9ef1b6440034828e48470e323951099f39e680c190103bbebb6155ea"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ign / ff973b12fbc9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-573925cd5d4177ce55f2d68748c926cf9bc2221b83216c89d1b49a8948f009d2"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ign / ff973b12fbc9 / 4

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8ccdf9cd4a3d2389c01ce9179709dc83db061e5cda01473fe1babef7907118a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-151c83c8148d0d8c9439b2974f81becbdb468078a7020f348480533432f86bbe"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_samesite — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ign / b40f1086fd95 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_samesite

<a id="canonical-ca97e2f240b9f32487b0a4c77f9509b7a351963308ae99486763d753458e44d1"></a>

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

<a id="canonical-3278564fc4802544443a1efcbfcdc7c21f2140db2522d03d86e6da03bf266931"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ign / b40f1086fd95 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3fa34d3568675a7b6ad6058cbecf4e0eabce3cc580cc6e7367a761c7e95c565e"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ign / b40f1086fd95 / 4

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a61daad1ed6bf9453ba0f597aa7f222af44e28a30aad6b794444f796c9096316"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c081e2032732bce4647f23abb9a00bbca27861473fcf9a69ad35124041a50aae"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_secure — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ign / 5c1fd08a864d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_secure

<a id="canonical-3f2b3c91de6cfd5670058c3a1972c407918c5441556f917d2eb5c55f59c5980e"></a>

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

<a id="canonical-908e5fd0d040ff9eb703a8ef0c2bf1987e5f3f67f97516547635908ddaae339a"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ign / 5c1fd08a864d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b1cb3eff1f6b15666d17e3cb5903627466baeba409916d78bbd29492805ca389"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ign / 5c1fd08a864d / 4

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8263f8070775f8b9500a4284e260228495fba7f168d18067afa97cddb1c8d404"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8368583898dcf276fbc0d57fb98fde062acbf0fe814dbf913dd6d0c9d0982cb1"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_lax — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.sam / 1674c46a70e2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_lax

<a id="canonical-583ffe52229ffa090d8f87973e13b64f82f83a9912ccffd54158e618174aeb9d"></a>

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

<a id="canonical-fa4326f87e9521e635ffd03e599b0aceef3244ee66b71fd1fb03d27df1d8a997"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.sam / 1674c46a70e2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-79feea5c040a1c9d31947e339e445d2821e1e63c7b5b4e3aca58d62e696cb729"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.sam / 1674c46a70e2 / 4

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3b5d8ddfbf31e2f9e48af8b76126df62bec316f7d0af3241cbe89b284a62f532"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd1f60bbc2ed54d7189d243fc0e2eb4718e2107c7ca7f1bc485db59678bd0d7f"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_none — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.sam / aa1153eec3d9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_none

<a id="canonical-cfa8e55c6db2cb2c5286100a8aa9efc66b47a44ca7dfbf71a9b6feb1ee0451ec"></a>

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

<a id="canonical-087cf6a6682080de733bd70e305bb6b68fd54a5f189702bf4f75b697c7867040"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.sam / aa1153eec3d9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-882a826ab0e746afd9b14aaf158cb4b026e1866bb1a88e227084eba421374108"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.sam / aa1153eec3d9 / 4

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-970bebd15250690408ada8e08b040de10494ef257921e090ec3071438874d57d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8e99771d0a68a25e95f6744c446d7bb8ae935bfcfccc5418f113a78bb388f7f"></a>

## routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_strict — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.sam / 8482a78ab97c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-174879adc7ea5b25fff80bf9ceba45812b16a9b7a84b277fadbae3de7d17d7c4)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_strict

<a id="canonical-ea8e9232dcec0eef8e0486b41525473dce6c0beba1eee7c80355d32e64058b4d"></a>

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

<a id="canonical-2b9b614bc469dd2521255677b6fa1ccc0cdcc441fc5934b431e3dd5e8dfaec78"></a>

## Direct properties — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.sam / 8482a78ab97c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3c1381d9b4f1d7cb38e18282fc96dd234f25ba902782845b80198e8475bd79d7"></a>

## Next pages — routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.sam / 8482a78ab97c / 4

- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-024.md#canonical-215f608645afb5d3c8db8a66e004dac0f5609b59250111ed9ee0b493533b5ed2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2e206a249d6baf756f0d86a64ab9201567f3fc42453191a3795a30dee8e4fa96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e03560b829d3a0764c0917186b858cdaa174646b889827297d217116c0a967d4"></a>

## routes.simple_route.advanced_options.waf_exclusion_policy — routes.simple_route.advanced_options.waf_exclusion_policy / 598b72231fd0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.waf_exclusion_policy

<a id="canonical-9ad4c069ed11d76e6883fd1e3a1025dac04c0e2f7f3548eb639650d31bbe9e47"></a>

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

<a id="canonical-16c8cf5990807dcc153a1706eac9666c6e6238488442bacbe87cbf68106d27e3"></a>

## Direct properties — routes.simple_route.advanced_options.waf_exclusion_policy / 598b72231fd0 / 3

<a id="canonical-df8b3838d320aff6fe1ee10d013a2444a8040cf56ad51e0978f7af9efbe38b03"></a>

<a id="canonical-b43323d5c15b86628e60c4db9e388087a22998cddcf4043efee89c0ab687e36c"></a>

## name property — routes.simple_route.advanced_options.waf_exclusion_policy / 598b72231fd0 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-552d1980413c71f3a308d2c334887942c597043ef94ecccd9411da2231701790"></a>

<a id="canonical-e16e887176f887928986b4926e66c735b1ecd04e22668d800ddaee89649e8b9c"></a>

## namespace property — routes.simple_route.advanced_options.waf_exclusion_policy / 598b72231fd0 / 5

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-439d9ff3b4ef5340ecd794a722218ef8fd32b621d7f52090fafb921bda3a1cb7"></a>

<a id="canonical-ae34849bea80a1dceda10d716327d1b81c5560cd3095ca663890ee95c82d065a"></a>

## tenant property — routes.simple_route.advanced_options.waf_exclusion_policy / 598b72231fd0 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-785da79ab0c9055182d40749e14c93653e787102a8e71d5682885a2f0067ff7d"></a>

## Next pages — routes.simple_route.advanced_options.waf_exclusion_policy / 598b72231fd0 / 7

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-da8c7bfa640c00d39cdaa9f4cbbe5e56a6989c13a45e893fffe113fa834c178c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e64f69f444fa910b2cef430e68109bd589fff6b971a4e69c6a82936b12ce9d6a"></a>

## routes.simple_route.advanced_options.web_socket_config — routes.simple_route.advanced_options.web_socket_config / 0ab9eeca3592 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.web_socket_config

<a id="canonical-cc8801d5ab464dd29a7e3eae73d83fe5e84a7f5445694fa26ca7fae67c7b94b6"></a>

Type: `"single"`. Computed.

Configuration to allow Websocket Request headers of such upgrade looks like below 'connection',
'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce
following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'.

Upstream description:

Configuration to allow Websocket

Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'websocket'

With configuration to allow websocket upgrade, ADC will produce following response 'HTTP/1.1 101
Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d3bc9ce771899a27574fea5770bad26337e0567509d6dcc8b1d0cd0c55fbe97d"></a>

## Direct properties — routes.simple_route.advanced_options.web_socket_config / 0ab9eeca3592 / 3

<a id="canonical-0105a719bcaf4ae3a9bbffad84c4a3d04255f47923c85c130cea9a60e1a2db00"></a>

<a id="canonical-f6756fc4a37ca11108d80df3b82747622cc0ad9151225a8386bf2797953a7e39"></a>

## use_websocket property — routes.simple_route.advanced_options.web_socket_config / 0ab9eeca3592 / 4

Type: `"bool"`. Computed.

Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket
connection.

Upstream description:

Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket
connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f26e63e81b686f70458577b83a52948a6f00b3b3df04a46c41fe4bc808384ab4"></a>

## Next pages — routes.simple_route.advanced_options.web_socket_config / 0ab9eeca3592 / 5

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3d492b3bdcca763d60b54e12c22e82299596d8500ba00a401784178986605aee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-456f65ddc2ae635092fccfa0ab5a43c724a1f55fce1ccf6ad489c7b7c0f60740"></a>

## routes.simple_route.auto_host_rewrite — routes.simple_route.auto_host_rewrite / 0d4aeadc4d1d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.auto_host_rewrite

<a id="canonical-9a43682c3c6c2e398045fdd2e550397beb0083d0548b42b4b578aace17260a99"></a>

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

<a id="canonical-1047fefb0f633835064259937a284878e0e097ff656f7b2c6b3d915e5a0e4c8c"></a>

## Direct properties — routes.simple_route.auto_host_rewrite / 0d4aeadc4d1d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e2c8d8f1b4c725901e49877a3aa90820b1b62b843a551aa0abf2fced5426ba6"></a>

## Next pages — routes.simple_route.auto_host_rewrite / 0d4aeadc4d1d / 4

- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5922e2977a3d71bb322b00946c4b30952fb7fdd019e711ad38fe0694ef48aa48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-447cdd810b4c49f4567c8acea8c24359a6bf6d8f2c2a096695e69f7c44586491"></a>

## routes.simple_route.caching_disable — routes.simple_route.caching_disable / 58bd95246c27 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.caching_disable

<a id="canonical-791dd4648874ebb3ae7066e89b1473be46dd31422fd9abf3dc98e7998744eb48"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching disable.

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

<a id="canonical-1e2014fa918f571141d426260c9c48b70991bac7b2613ce3400eec9927d42551"></a>

## Direct properties — routes.simple_route.caching_disable / 58bd95246c27 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-496d94e1b755e43f9e8af0b555b7522d08d61dda0d9e10b7449e186836769f76"></a>

## Next pages — routes.simple_route.caching_disable / 58bd95246c27 / 4

- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-780d60a72d3993fb38d01eae6e6c0f445bab2ba40df7c10859410497869b5fce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e1cf45d34c70fdefb4e60c1038fba0c69f2f40b563e7088809f9b9a0d6426dc"></a>

## routes.simple_route.caching_inherit — routes.simple_route.caching_inherit / dc784bff56cb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.caching_inherit

<a id="canonical-44b63d23d8bb9eadc7eb81a930e30bd462e8fa6e6156e8f842848c8befd236bb"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching inherit.

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

<a id="canonical-1ef9ad2984b5557b54878fab2d5d20219a406e6814a9c5bb35855a4ae70a0a3b"></a>

## Direct properties — routes.simple_route.caching_inherit / dc784bff56cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-30d9b544e8528895a558d4af72df1aae2428e938aed5b503f4108957197e668e"></a>

## Next pages — routes.simple_route.caching_inherit / dc784bff56cb / 4

- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bb64bd3e4bab0522d7301ed7c9bd4204e6ba4350583a5d3fde8eda7b2eaa92b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b41a6d0fc3265e522e272ff3251c377cef265f9fec6e1bb12d6db67000adb46e"></a>

## routes.simple_route.disable_host_rewrite — routes.simple_route.disable_host_rewrite / 7542d974a21c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.disable_host_rewrite

<a id="canonical-d8b29ca3ff0b14a41bbaa06beab58fce37b90d375261de591c382335f28e1042"></a>

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

<a id="canonical-7abd16d28e517d39232b5c25db47306b73cb02a14f25336d1b933ead0a7f0ba9"></a>

## Direct properties — routes.simple_route.disable_host_rewrite / 7542d974a21c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-31f0470abe22ec65b69148c67b1e7394ec4b7d39b2ba2ded548a7dc9d840e634"></a>

## Next pages — routes.simple_route.disable_host_rewrite / 7542d974a21c / 4

- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
