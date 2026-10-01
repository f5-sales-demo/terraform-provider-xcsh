---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-e3b3f4a380397a1ae2a50fb44d6b9f58800ab9d9031d308464e205fb2eddf192"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / cc578a3f036d / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-06afe2d77108c161e5b11add33eae536bb73790ddb1e880b3ba2382bbffcb867"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95e53d53f43939f4af7ca0a887a57c34b025969481c8ea3ef2f80fcd29d5d8ab"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fac35a98b27 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-008.md#canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-80834883bb29d02416db66b869a51de6b1563885ac731aacc8498528ba737fd2"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-bc826d307bae6cfede954b14ab954dd40237463b91d81b0a872c8d275fc17bd2"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fac35a98b27 / 3

<a id="canonical-064843f370db8f1660b71b22d999c8136b363757c933b10748316314356bff3e"></a>

<a id="canonical-b82fe47cc60a4783d63da82531402d94be1e529d1ad22925065e6491664eb288"></a>

## methods property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fac35a98b27 / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5207c950b84aa0e92ef9489fb4b0fcd01cfb04d09e28e89b680bc92e61faeeef"></a>

<a id="canonical-9a4e3fee8ecbbe9786363db2d83271f320821631f02c2db0b5cc03b152b2d26d"></a>

## path property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fac35a98b27 / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-0e40607dc6e1b33c14e210d8b5a53004ce52867df937cb28e5b386c099e6672e"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 3fac35a98b27 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c838a2070be0a009813961ecc76c0294317f622fa57a424a00300c9097782711"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-862ec12d7f04e829c764007a5941c2a9cec3643c50922443bdcbebb1dfe9da13"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / bdbf3bea701a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-008.md#canonical-2ca90392cb450837e3fdc3a280257444e64d0846446395f7753467b6781dde1f)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-008.md#canonical-a295ecb477eab1471327ef7fbc10fe3cb7b855e7ce0b2ab23d9344a4bbf3e08b)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-09cb633e3385ebc1ecd84ee1d48f8b2ff053f3589e002d9d88dcce0e7796c79d"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-8dd98fec2a0abe9fb5838575c382d9337f2019424c7f8ac2dee074658d68437e"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / bdbf3bea701a / 3

<a id="canonical-b45368f0bc06b3a62c990382186683cc1d738af4da650d00bd797cd92d609b79"></a>

<a id="canonical-3b9746d6559a5d2061a5722afb78cda3e0dfe6952072a64ec45ff28a1d673515"></a>

## description_spec property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / bdbf3bea701a / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-261f30a4e022ffcbbc34d5cb91762052edb5a44d20805f36ab9f9dcd7656151e"></a>

<a id="canonical-24a3556579ff09a6185408608bb794d39e440287ff9ffc2d8fb54162adabc105"></a>

## name property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / bdbf3bea701a / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-2e51edd10e72fa64d54f06997681a2cc8d25d417dc639b9f7bbb5bf705619b62"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / bdbf3bea701a / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-a85182b7840ea1d7133a59ad426e26ce72187273de49f8f033bf651558c6cc6a)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cbc00ca2880099c07bc1c4a8f6c39807ba80dab43df7e9e7db86c82fab66a16"></a>

## api_specification.validation_all_spec_endpoints.settings — api_specification.validation_all_spec_endpoints.settings / fbcd575755fd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-3f9f3e7537cda436f87599e35c23e0e215ac544e76e5a227ac53014e97298004"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

<a id="canonical-a844e4093fb9999426e58864720258891f6c56b293cb1d601c3b3388bb8dba46"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings / fbcd575755fd / 3

- [oversized_body_fail_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-d7bb28a4ca01b27a6d7153b84526cb85e09e2c967b11e7d521384fe75892609d): complete subsection reference.

- [oversized_body_skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-47b9b1d665784b20405b8e0250f4776129a12e92428e7a72c1dcc029b1d80868): complete subsection reference.

- [property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0972b6f83b0c31ea60ca3233c466d3d942c390bd8d096a11ffd121f8125e2386): complete subsection reference.

- [property_validation_settings_default](data-sources--http_loadbalancer--reference--group-009.md#canonical-e86e1f803d5bbcf191f2bd2332d586debfe139d89810d2d813880ba77ef56533): complete subsection reference.

<a id="canonical-97b00505378d28492cb3c75838f394b2a8c510b3d32e82acf873925b4c5c1782"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings / fbcd575755fd / 4

- [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-d7bb28a4ca01b27a6d7153b84526cb85e09e2c967b11e7d521384fe75892609d)
- [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-47b9b1d665784b20405b8e0250f4776129a12e92428e7a72c1dcc029b1d80868)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0972b6f83b0c31ea60ca3233c466d3d942c390bd8d096a11ffd121f8125e2386)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](data-sources--http_loadbalancer--reference--group-009.md#canonical-e86e1f803d5bbcf191f2bd2332d586debfe139d89810d2d813880ba77ef56533)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d7bb28a4ca01b27a6d7153b84526cb85e09e2c967b11e7d521384fe75892609d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69470a7b05bf6601d36b3338500140884e740b75be7dc2ad1c0a4a0be24cb00d"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / 6654d0bf08f3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-546ab81e1f2e678be2833453a64f2045dfd1ae9e8f52cecb9e4763380be2e4b7"></a>

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

<a id="canonical-c1bc3fbad91495d9ec64c3e302c2fa31872274c5b50737821ca664b2a1bdebc7"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / 6654d0bf08f3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-35f3fa383cc26dcd0c25e42ff78b236e0aa84db099396ea067aa76a4f13dadeb"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / 6654d0bf08f3 / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-47b9b1d665784b20405b8e0250f4776129a12e92428e7a72c1dcc029b1d80868"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b47d2e6afda41e87acff28a47d22bc35866447149d272837e3291b24b8f504d1"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / ddc8e650e0fd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-05937556227c51c4004a25a033c349e333f9ebad1d140ef37246eab91146b60b"></a>

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

<a id="canonical-0b5bab9de83a536b2cb65458e7747e84df8e7c48cf85ddb14f9506bb417fab18"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / ddc8e650e0fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd8ab06a33693399ffbd5f7e6044c632fb3fdec7b38c1fed03ef17888fe6d541"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / ddc8e650e0fd / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0972b6f83b0c31ea60ca3233c466d3d942c390bd8d096a11ffd121f8125e2386"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b51f4ee8e3673a086cd92f41ab8db69c24b37de3923f4679dad67b191918aea"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 6785d77b0803 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-01169c73945d14167cd6bbbdd8bee22e10e01ef0bedc34f42d456cc2b9e248bf"></a>

Type: `"single"`. Computed.

Configuration parameter for property validation settings custom.

Upstream description:

Custom property validation settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cbf72ba2622d0acd8a336a5ee6bf24b3dfb93a7f474a58b93d047e51b97d56dd"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 6785d77b0803 / 3

- [query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-8de4c1a33586077e48f0b1d2932c806bb200e42c329b54effc484accbcec2fbc): complete subsection reference.

<a id="canonical-2ec0c387dc64c4a6a622290122d73e6d6e29b3fd683b39462bac8506542fdc87"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 6785d77b0803 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-8de4c1a33586077e48f0b1d2932c806bb200e42c329b54effc484accbcec2fbc)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8de4c1a33586077e48f0b1d2932c806bb200e42c329b54effc484accbcec2fbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f31140635d1e90f8334590562187376700ebabb351206f8c2ed2d37070da84e2"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / be42c4e192c5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0972b6f83b0c31ea60ca3233c466d3d942c390bd8d096a11ffd121f8125e2386)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-bbf8aa103db82c905469cab482e10be9210fc9f45be1fb5601ab3c7a7e6aa5e8"></a>

Type: `"single"`. Computed.

Custom settings for query parameters validation.

<a id="canonical-d86a676930be750a281c1fb6fd4f706e71156f61ba0d7388afd443de7e499cbc"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / be42c4e192c5 / 3

- [allow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-7f9305d81a8e960cf5e26340ad6849e33b95118c125018e57ff9aa6c6fef1633): complete subsection reference.

- [disallow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-3ce75135df82ba4e7d5693dad2ac45b3986c47f8197ad24402b3c5c53bb17b88): complete subsection reference.

<a id="canonical-f56415b778229a68f4bac7ee58f5389e24e18537b6b7a96f5ee08636bf8b8edb"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / be42c4e192c5 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-7f9305d81a8e960cf5e26340ad6849e33b95118c125018e57ff9aa6c6fef1633)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-3ce75135df82ba4e7d5693dad2ac45b3986c47f8197ad24402b3c5c53bb17b88)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0972b6f83b0c31ea60ca3233c466d3d942c390bd8d096a11ffd121f8125e2386)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7f9305d81a8e960cf5e26340ad6849e33b95118c125018e57ff9aa6c6fef1633"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc1991a15abd3a14a159f2fb7bc29eca053dda66eb4ce707c25e0cd7dd77a55e"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / a545117fec1b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0972b6f83b0c31ea60ca3233c466d3d942c390bd8d096a11ffd121f8125e2386)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-8de4c1a33586077e48f0b1d2932c806bb200e42c329b54effc484accbcec2fbc)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-ad051deacf417f45ae4917588a1e75b043035d724202fdafa2deb087ff899144"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow additional parameters.

<a id="canonical-5a1323a5f2722f7774f33059e4f9c376a4e5f587cdd8abb81433d80d9ee86501"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / a545117fec1b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f41e1449bf9bc6fd72d6d0466c4d453541b1d864451fb28c7316505f1a03ca6"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / a545117fec1b / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-8de4c1a33586077e48f0b1d2932c806bb200e42c329b54effc484accbcec2fbc)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3ce75135df82ba4e7d5693dad2ac45b3986c47f8197ad24402b3c5c53bb17b88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ae4dbbe03d38be85860c873a1cc5f2b56ab8e35d5387f78452bc94e6a10588d"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / e9c8d7216bb7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0972b6f83b0c31ea60ca3233c466d3d942c390bd8d096a11ffd121f8125e2386)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-8de4c1a33586077e48f0b1d2932c806bb200e42c329b54effc484accbcec2fbc)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-e23861cf01685f8b41fadc0c0d5ee5dfc03e5ec7cdde8027ebecaf05ebecb0c0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disallow additional parameters.

<a id="canonical-ff9ed5fc7b06ae5d19cfd89b8c7fc669a76db6194a1d988dfe8eec80f183b50c"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / e9c8d7216bb7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ed0fd86636f42342146cca83089dd18be273a2b2553c8215d83a3cb89d78dc1b"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / e9c8d7216bb7 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-8de4c1a33586077e48f0b1d2932c806bb200e42c329b54effc484accbcec2fbc)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e86e1f803d5bbcf191f2bd2332d586debfe139d89810d2d813880ba77ef56533"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f97a4a4d9af89bf2cc03171e0aaff6f7abd19000a1cd8afacd7b247c8205eff2"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 6a3e5ec7b8bd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-8db2668a11651b2ddc78ef95e6780089884742ebde3e11621a9df71c2a518492"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for property validation settings default.

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

<a id="canonical-c941fc97657abde50ea3d648066917b4e87c916467795a14a147f6ae921fbe15"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 6a3e5ec7b8bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6738d1b95434835b96ec59a3ba070a69fd7f3f184d1216b756fded47943f2d4"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 6a3e5ec7b8bd / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-85a5c1ab7f8574f49028a122e6b70f85fa4110ccd9307d8cf5340f5778df83cf)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bb999e81cd99169c389a3f79a842c9ffa89350b5560f986e4e55abe453903c9"></a>

## api_specification.validation_all_spec_endpoints.validation_mode — api_specification.validation_all_spec_endpoints.validation_mode / 7623ad7e7e5f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="canonical-be7abaedab7b1038ff6455e0edf3051c641d86b9b46a25368794538f7ebdfa0e"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

<a id="canonical-48c9c5ead05689b98ecacc5774bcd573bdfe57954c7ca6d2d394a506e6ee63a0"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode / 7623ad7e7e5f / 3

- [response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-97345efde6917024ffd4a2fcfe991eb2a80b2cd4fb8b55fb2e16098461d03263): complete subsection reference.

- [skip_response_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-de8dac218f3fc1ed05d7fa81feafb0e42b11562c3167139eaf8fac68829f8eb6): complete subsection reference.

- [skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-315d822318bb9eec30ddfa68aa8a05428cfdc00fcc3658fd0358a1c103ca3378): complete subsection reference.

- [validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-af801c78f3acd2823bdd729f762d92c267da2bc5b9fde4de99af37d9199047dd): complete subsection reference.

<a id="canonical-a9e52b4d80f1d94f060ab90f8fd1df89f0eec213ddbb2ab1eb8e74596aa3a26b"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode / 7623ad7e7e5f / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-97345efde6917024ffd4a2fcfe991eb2a80b2cd4fb8b55fb2e16098461d03263)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-de8dac218f3fc1ed05d7fa81feafb0e42b11562c3167139eaf8fac68829f8eb6)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-315d822318bb9eec30ddfa68aa8a05428cfdc00fcc3658fd0358a1c103ca3378)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-af801c78f3acd2823bdd729f762d92c267da2bc5b9fde4de99af37d9199047dd)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-97345efde6917024ffd4a2fcfe991eb2a80b2cd4fb8b55fb2e16098461d03263"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-666c76efadfc20d6215b35733422742b2fb486f8e25a62f2e1942da37b9c04c3"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / bbccab7e0c30 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active

<a id="canonical-a6bbb4e2f7ce61bf1dd67b56067a8c6dfae8891ed934b4613f625cab247ba434"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-f9ced3174db073022744e0f4adc7b27c4bfa15d3b56f900d348332b2350f5e36"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / bbccab7e0c30 / 3

- [enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-dfd99790a8010e0887eb44d72c85c6fc6ca691d676bc6db6707b258d9d0c9a1f): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-3e1ac102e7338ece7d8297b7be9b5a713912a359028b8dc23e76a4d778320369): complete subsection reference.

<a id="canonical-b079b4980ce6336dc5f3efeb0c3b3a45ed4f84faf0c97be00e4112baf413af02"></a>

<a id="canonical-2727f3e0e81fc20ef05424d7f1f3b7704962a0d8602c6eed4e561945e5d7e2be"></a>

## response_validation_properties property — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / bbccab7e0c30 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-22686c9939070bfc2f59a5f8f82dda839d60c76f5da5451a46f2b4789d01bce7"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / bbccab7e0c30 / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-dfd99790a8010e0887eb44d72c85c6fc6ca691d676bc6db6707b258d9d0c9a1f)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-3e1ac102e7338ece7d8297b7be9b5a713912a359028b8dc23e76a4d778320369)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dfd99790a8010e0887eb44d72c85c6fc6ca691d676bc6db6707b258d9d0c9a1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78e44d2dd383d76aad19bdda83bd0f6e0163180d8766effb77fd43f41c70b5b3"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / b1566634c597 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-97345efde6917024ffd4a2fcfe991eb2a80b2cd4fb8b55fb2e16098461d03263)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-6784bda6967a36fc37ef969b80f1832207fea7bdf77e79b2d2c447bd87df3d4e"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-dea4f2f1edc4cfc3648a1bad63f37fcfcead6e6167a8a35553da2cbb20708efa"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / b1566634c597 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fae350377e5ee796e492fa2731dd5b3080b77112c5fdf1e91330c1bcabe5678c"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / b1566634c597 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-97345efde6917024ffd4a2fcfe991eb2a80b2cd4fb8b55fb2e16098461d03263)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3e1ac102e7338ece7d8297b7be9b5a713912a359028b8dc23e76a4d778320369"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cf236990b6147fcb16f4699e2eb182b74f7bbba4eb9cbe77b1ceacd4f9f0931"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 647ed522512b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-97345efde6917024ffd4a2fcfe991eb2a80b2cd4fb8b55fb2e16098461d03263)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-291a7540e6373ea763e6ffe1fb83f8c4771c1129b73d7e87b8c32ddd83b28223"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c1793ca4bec5beb29a1c4e0d77972f87d2d517b9e0d95857dd5105238e07362f"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 647ed522512b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f521c2dde3ca7a8a39ae86c75713a45b6845ee7573befd43d8405bac87075eea"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 647ed522512b / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-97345efde6917024ffd4a2fcfe991eb2a80b2cd4fb8b55fb2e16098461d03263)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-de8dac218f3fc1ed05d7fa81feafb0e42b11562c3167139eaf8fac68829f8eb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-800769457cf76dd2b7d02ac4b347906138549ebd437110b9db1d113de640a3e0"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation — api_specification.validation_all_spec_endpoints.validation_mode.skip_response_va / b296107ca50d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="canonical-792856cd4fe4d97ec3656aea3fae63370c2aef7dc80bbae2b41fedfbc55cc42a"></a>

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

<a id="canonical-0aa7d703fe87644da33a38a18864659d03015f8a21496eef7e150457a93bb7fd"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.skip_response_va / b296107ca50d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-32dbf202de7cf91c2f5fd557a030d4023b01c573515f07f1ad62337516dae0a5"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.skip_response_va / b296107ca50d / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-315d822318bb9eec30ddfa68aa8a05428cfdc00fcc3658fd0358a1c103ca3378"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7d60f5b45a2f2ee1b478623fe18969e2ab08972055048d2d6fdbd0e0fc47282"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_validation — api_specification.validation_all_spec_endpoints.validation_mode.skip_validation / 775e62dd4d58 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_validation

<a id="canonical-fab7a9f148e75cab4a297429d8a1baad13a4d7af085ba9dc88f5b9016f1a9c36"></a>

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

<a id="canonical-77437bb2a1c87c479ea463111f93d6c15b2df52b85049ef8e26aa1fdd23a47e4"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.skip_validation / 775e62dd4d58 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a2d59f4e99bc5ff03ea6510f478be33c629b36d9fb44cb32d393a689f066c961"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.skip_validation / 775e62dd4d58 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-af801c78f3acd2823bdd729f762d92c267da2bc5b9fde4de99af37d9199047dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6814fe467dfac242992f15692adda186f95933871afdbe5671241ed01ee996c9"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 1e37db754afd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="canonical-6ab2041f2cfa52fa309cd1659b16e596d6c9641994130a125a3513f2e42b0349"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-691f283f22f2d81ce72f6ad634dbe2c8117e60fc3d279609f2bde82006e480d9"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 1e37db754afd / 3

- [enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-c0b8fdaae82bab79a6482a2ae5855aba1fbf0b829fff5dbd5a1ce49ba4d892c1): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-fd4e403545e4bdc51e868635840b03dcb264a6ea227bd022b8864a1c272c6b67): complete subsection reference.

<a id="canonical-07c4b2e6dfb8e48910fc5ca0d23897dea501c33c3d74665d4eba7b8e91bcbd1f"></a>

<a id="canonical-2fb35fb0dd7fdc1c5b878a9348bbdaf6e1a9473be76f00fafaa3eafcfa310f32"></a>

## request_validation_properties property — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 1e37db754afd / 4

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-128c2f183db9d408e8cb6193be63729b75bdf0d70b983713083e97260d18b922"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 1e37db754afd / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-c0b8fdaae82bab79a6482a2ae5855aba1fbf0b829fff5dbd5a1ce49ba4d892c1)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-fd4e403545e4bdc51e868635840b03dcb264a6ea227bd022b8864a1c272c6b67)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c0b8fdaae82bab79a6482a2ae5855aba1fbf0b829fff5dbd5a1ce49ba4d892c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f0a06af90ea99a23171024baea9a13815445ec80c39507bf748de7d773e565b"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 065836c9e5e2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-af801c78f3acd2823bdd729f762d92c267da2bc5b9fde4de99af37d9199047dd)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-0f3ded040f4fdd060bb89f1febe9fee5d6a2a54e8a4e2d3139326f1c5c202267"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-6189b01581bc034c6bdc3f962a3ede943c4f9716c70b57a26cbeef0496806802"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 065836c9e5e2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15a9eb484d680b623c49edd5a4685534e6828f013ef8210c002fbbfa6abbd673"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 065836c9e5e2 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-af801c78f3acd2823bdd729f762d92c267da2bc5b9fde4de99af37d9199047dd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fd4e403545e4bdc51e868635840b03dcb264a6ea227bd022b8864a1c272c6b67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1f94bc315d433f794faacd645886d1ba58dce62a43c56227981b9f49248d17e"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 5c912f1eeff8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-011e37067ecac1c07021dfb1254a8073668adaafb9eabaf3dc63be67841ec2f9)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-5f8b6862dbd5a391dd2569966c77717d91f2576d4b1fe80ac6cfce91b8f0aa4c)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-af801c78f3acd2823bdd729f762d92c267da2bc5b9fde4de99af37d9199047dd)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-5ed1abeb70ed9c508848af31355165c426a31cc0ee2e22864372161ef036483c"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0a8c5566873271b204e206b8dd859b9aba52903ef61ee81b6afe1dff5223a021"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 5c912f1eeff8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb1b384a42add218017a00aaa3f5534f4b5de70cb978ff3db57ed62b45b9d1af"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 5c912f1eeff8 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-af801c78f3acd2823bdd729f762d92c267da2bc5b9fde4de99af37d9199047dd)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf2fb9a72d827072c8d735852a9f14dd98734574671c62757427a0371b45f001"></a>

## api_specification.validation_custom_list — api_specification.validation_custom_list / 4abc7a27990c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- api_specification.validation_custom_list

<a id="canonical-d17b6b79ebec8c856939555186fef2da7ff5cbbbc76a2fe612ec40300b009dab"></a>

Type: `"single"`. Computed.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Upstream description:

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to "Fall Through Mode".

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

<a id="canonical-1bf65c2dc054efc1326db3a5af2417d9b0d6976690ff1a33633891367003245f"></a>

## Direct properties — api_specification.validation_custom_list / 4abc7a27990c / 3

- [fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715): complete subsection reference.

- [open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad): complete subsection reference.

- [settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2): complete subsection reference.

<a id="canonical-c1019ad71f32b02a85fe81416596023a49af7f5e22c3c90f4df5a45b642fc6e4"></a>

## Next pages — api_specification.validation_custom_list / 4abc7a27990c / 4

- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7a2f8830401646f816a7c7b13f5f99df1f59acee008a91a5b668e49e88547af"></a>

## api_specification.validation_custom_list.fall_through_mode — api_specification.validation_custom_list.fall_through_mode / f05661577d74 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- api_specification.validation_custom_list.fall_through_mode

<a id="canonical-3dc0bcd6570bd860af37c2eb0a9d2fae4fe04204d9590cc8742f3c085b2a64fa"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

<a id="canonical-e8771b9386bc29652f2c86a995b526d693de34cc9bbcde91ed63a7630c75929f"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode / f05661577d74 / 3

- [fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-009.md#canonical-8908ad78f6fcd360ac2e481ad4a4d5245e7e50188fef7b418af65edcbe308098): complete subsection reference.

- [fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af): complete subsection reference.

<a id="canonical-ddbb556febe04c011d86e601484dc721ef963ca27e0091ad6b204e285c6e9825"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode / f05661577d74 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-009.md#canonical-8908ad78f6fcd360ac2e481ad4a4d5245e7e50188fef7b418af65edcbe308098)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8908ad78f6fcd360ac2e481ad4a4d5245e7e50188fef7b418af65edcbe308098"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf764670732a9286cea87ab0c399c28fb8b823d447d4106b2dbfde5d045d3294"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_all / ec54453a68e3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow

<a id="canonical-8334cbc637c4a2f9d0012f0e875eb0b3033d380a489a16d0478680847ba1ff85"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-51fe24cfbb314b75840027575940c66fd19ad26eb611baeb80ec0788a78907b4"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_all / ec54453a68e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9fd43bb7753946a3e69eb0750f4ff0c02a140f5b3f99ccabe802501a95c3f205"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_all / ec54453a68e3 / 4

- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3207dd4e7ca926fed01a91873ad7a6ae348de97f027889e1fc5af015da479404"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 45c70113eeeb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

<a id="canonical-b8b4f5cf3efea49beadc03ae419c52a804bef09b021e47ee05a1707a9680f68d"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-800bc335ba28c0ca7bbe92ce827920eba89c6b63a7a477f0a5c864324ccfa850"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 45c70113eeeb / 3

- [open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae): complete subsection reference.

<a id="canonical-ebe8cf8b94ec5f4b903455fe4e5d8e77cfeb9d563473fdfc6f9e1f4484e8f227"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 45c70113eeeb / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f31643ce1921da427a02f0a6cf1b33f5480b832a4962b98612c5d7abcdc5f911"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / cb1721f58bff / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-3cb7166b0aa8bf6a01a22eab70b94f8cdd5b28d0b10549f8f078b3e2e46ab001"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2e525cb1beec243e0c29772fa1b8ba0fa6f5545dd705d716aa226c390d62c988"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / cb1721f58bff / 3

- [action_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-15443b91dd43422f84dad994b709711b7fe8f45cd9570d9cba1deaf5d57f6efb): complete subsection reference.

- [action_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-548cdd30697ba3d4eee7bdad1cbfaef1991292fbe01ba020a72e16495e899a01): complete subsection reference.

- [action_skip](data-sources--http_loadbalancer--reference--group-009.md#canonical-3f8cab7ddb190341eac5ad14a1a3f91a1966f70ee3f50f2d8044f79f502ac9c4): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-69de592c7ee3046e2159c6ee761eb700d298656a49ba08c81d487669f35c39b2): complete subsection reference.

<a id="canonical-f3848a24c832579172ba97f829270bc55a0b46c615667a99517df961ddd4f950"></a>

<a id="canonical-ebb8fe71792cb71556805e9c7d60102a00040a4040951fab7054f4520064751b"></a>

## api_group property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / cb1721f58bff / 4

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-566181f6febb902737605be3b36f8d2aeecdf42dd1cbf06e0e83b4ff34c39ee6"></a>

<a id="canonical-0c9feed4b07f182367080d9d940bcff7c4960d5cc2bacbee7ae2619fd2bab5d8"></a>

## base_path property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / cb1721f58bff / 5

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-eccd43966fa0a3b62c40d2b9e8e60402d8f3ea37cee166ad68ce9d3197e8c5a8): complete subsection reference.

<a id="canonical-2664be9f75e58599766d60a09ad5b432a618586f1dae26621aff2a6672aeea90"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / cb1721f58bff / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-15443b91dd43422f84dad994b709711b7fe8f45cd9570d9cba1deaf5d57f6efb)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-548cdd30697ba3d4eee7bdad1cbfaef1991292fbe01ba020a72e16495e899a01)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](data-sources--http_loadbalancer--reference--group-009.md#canonical-3f8cab7ddb190341eac5ad14a1a3f91a1966f70ee3f50f2d8044f79f502ac9c4)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-69de592c7ee3046e2159c6ee761eb700d298656a49ba08c81d487669f35c39b2)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-eccd43966fa0a3b62c40d2b9e8e60402d8f3ea37cee166ad68ce9d3197e8c5a8)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-15443b91dd43422f84dad994b709711b7fe8f45cd9570d9cba1deaf5d57f6efb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee1d93b0764b781401719c2aa025a2e189f88332ecdd414b182ba6a6f6cdd0aa"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 3b586df4587e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-c0bb3517b9142fc727f890fed8140fd05e6a77a2180694db19b197efcebfad95"></a>

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

<a id="canonical-b379a517db4137523b5bbb81989bc0a0e21d26e149cd2f3906ac5815a281339f"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 3b586df4587e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9626d3156ce7b0ef98ca16c08ca198bcdf474e0752c4f2de034e34877e2a996"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 3b586df4587e / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-548cdd30697ba3d4eee7bdad1cbfaef1991292fbe01ba020a72e16495e899a01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a27845fe7b5e46f54e916028e362674a8d3c643827a8017215587bf4f59cea2"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / d04c91ff9498 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-c948df46349c6fce425ec12cbcba7dfdaf1f8b335a849711538c2485598b2990"></a>

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

<a id="canonical-b8fc8b34d12f83cefd7a936fff52fc205bf201523480d1817a198e6ce2925d1e"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / d04c91ff9498 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06646493d82972766cad4a826262cefaa0453594e4dff6876485b0b930133af6"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / d04c91ff9498 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3f8cab7ddb190341eac5ad14a1a3f91a1966f70ee3f50f2d8044f79f502ac9c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e09d01a739af0f4b2488343f607dbc68327397c53f4e5bbb440e8da38ce552af"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 244198dbe9e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-ebcbd712cc92755c80199a07671620c53e2013db133ca506fbef755d0c448c6c"></a>

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

<a id="canonical-81ab323ca601dffe016fced415e4b5af68b0b601a6081dd001485c7ce01dde5b"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 244198dbe9e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d46cefbb4b603f75391268e87389a10a71eb5e337a20805c79fd4ce2c81df500"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 244198dbe9e4 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-69de592c7ee3046e2159c6ee761eb700d298656a49ba08c81d487669f35c39b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1eacdad7cf864b9e615d28b0c40c6145615938d5b65a9b994db0c9572a6d2307"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / af40d13ce1bb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-200b3de89df9e77524c37b848ff84dfa3ac913def44d31f4ec3a9b59b5ce5d79"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b263817823e45a689a2f1020096c54e656751f83632e707df84ca5607394a1b4"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / af40d13ce1bb / 3

<a id="canonical-fdd842d21564b41f9e63f57014f8e9d8cd194ba254c2fc170bbad7774db4d0b7"></a>

<a id="canonical-f81c15d526513707a78ab6f9b9d61a511aadc64c7337d4b596e3452d341e7fd1"></a>

## methods property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / af40d13ce1bb / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f8b1257f9d515e237c38418ed381da516d76f7d9cf02f284ced672604a8c6145"></a>

<a id="canonical-e8d7e2ae574ca769f1689915e6adcfaafe3ef0ad238f6964e5dbecfbcb8302ba"></a>

## path property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / af40d13ce1bb / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-6dae734fa6e288f2ab3dd85f60f297ea49f39b126b61df4f1ca1d766b79d2a4f"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / af40d13ce1bb / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-eccd43966fa0a3b62c40d2b9e8e60402d8f3ea37cee166ad68ce9d3197e8c5a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec6b6ec104602957a4b3e59e7cad41068a87a9ed8ee4c7c009c22978174492c8"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 11fc0f939e85 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-de13675f0c449d978bbb1581fe7c25f44f08fc7d45fde2987f2a17e755051715)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-3a02f130075c496748bc0c908d53ad16cc62e469aa655d523fa069d755b873af)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-41a890000f7310d0f5fd0dae4a8f16f29073d41c5637e345385ba630689412bd"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-ec7517d01453fda74c295b700c77e87693f15f8f3b52c7f0c3f023b6b027544b"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 11fc0f939e85 / 3

<a id="canonical-f5062feda6d7b99a42b88ff3a170d25b3e93a488fa3954b6fd50bec29f297401"></a>

<a id="canonical-e7490379fc2738d2f6f6d5a99a974c915f3bca5a31c5ad87b1ce346431b5fb66"></a>

## description_spec property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 11fc0f939e85 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-4adf0755b50581c110aac2359c89c9dbd7d458f37564a310b623d1d3df7176dd"></a>

<a id="canonical-b86fdb4cbb0787190f92fac123b6288086f697be7918c4f106487f0d9d1600ed"></a>

## name property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 11fc0f939e85 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-1a7d675d1019f763d4707c312baccf77f8c5df948d5df357ae249fb01979dd9b"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 11fc0f939e85 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-6d12da5a869c307475b1e26457ee0455f9fb80e49f8cbc28a680dc82921db1ae)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5bf836722bc5e0976c09c0599fef9da5f83f2cfbe645fdddeca4da4d1397bdc"></a>

## api_specification.validation_custom_list.open_api_validation_rules — api_specification.validation_custom_list.open_api_validation_rules / 58a9e1256c8e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="canonical-54071b169bf3e98d225e1d033b17bd32846b3a590091242cb55a1f28325494f9"></a>

Type: `"list"`. Computed.

Validation List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-a0f25ace04ed2a061f3559ac1d02ad3eed4c2b41060b1dc925ed19eeaf5730c7"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules / 58a9e1256c8e / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-009.md#canonical-7990820be13f0128f6aff3b85f63a8e5ed5bf064f9fd6055f75872b93c4fb91c): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-d5775e4f3b0f9313489332872a20307adb38fc7eb24c3874e830fee0b3054d77): complete subsection reference.

<a id="canonical-7f1f7f1b8b156636989d5bc37a6bd97a4c0c9efbcfaf1420b6d23208095a5461"></a>

<a id="canonical-3f8e2454cead7929f29f3829e615f9e33e2ced7fa3b665111db1e43227ee2f70"></a>

## api_group property — api_specification.validation_custom_list.open_api_validation_rules / 58a9e1256c8e / 4

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-bed4adebc42f0c2261685327880137f67ec981113e65aef0bf94482b83eb244a"></a>

<a id="canonical-aa9cbd2a69c5d62336039a3548b3076ed2626970a62a10c368d3125397373b27"></a>

## base_path property — api_specification.validation_custom_list.open_api_validation_rules / 58a9e1256c8e / 5

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-b23eec20ae3f17d1ed9959f4b5f98a9335fd50b9b0224dec4c35cd5935eb1c71): complete subsection reference.

<a id="canonical-f4487f8104e0d4356b0346926ee570380dd02dc5cc33098db12369abb2c2e121"></a>

<a id="canonical-266bcf73dc0c7e333d8fc32c2ba3c2005734edcd3a7269a2dde47d0e9c4211c3"></a>

## specific_domain property — api_specification.validation_custom_list.open_api_validation_rules / 58a9e1256c8e / 6

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c): complete subsection reference.

<a id="canonical-c6e28561ab76ece836afa4e84fe31da1bd62327d5922d0883f9f8448be930034"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules / 58a9e1256c8e / 7

- [api_specification.validation_custom_list.open_api_validation_rules.any_domain](data-sources--http_loadbalancer--reference--group-009.md#canonical-7990820be13f0128f6aff3b85f63a8e5ed5bf064f9fd6055f75872b93c4fb91c)
- [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-d5775e4f3b0f9313489332872a20307adb38fc7eb24c3874e830fee0b3054d77)
- [api_specification.validation_custom_list.open_api_validation_rules.metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-b23eec20ae3f17d1ed9959f4b5f98a9335fd50b9b0224dec4c35cd5935eb1c71)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7990820be13f0128f6aff3b85f63a8e5ed5bf064f9fd6055f75872b93c4fb91c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19a88bd1ea58f6b3f338b8ac1f765db76ca820629d0733bda821daff45710a35"></a>

## api_specification.validation_custom_list.open_api_validation_rules.any_domain — api_specification.validation_custom_list.open_api_validation_rules.any_domain / 5c3782b3dce9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- api_specification.validation_custom_list.open_api_validation_rules.any_domain

<a id="canonical-60c8997f362e781be2a6cfb269cc233b1b146aff21a0cf2acd2ef79b23daf920"></a>

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

<a id="canonical-31c141f10769a59c7e732fadccf1bfceca52c243fb7e769ae76baf5f2a995681"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.any_domain / 5c3782b3dce9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a4ff6fed60f38fc03aae624ce09b745d10b0bc2aac1b9337edf1ec1b741203a1"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.any_domain / 5c3782b3dce9 / 4

- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d5775e4f3b0f9313489332872a20307adb38fc7eb24c3874e830fee0b3054d77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b429182acb8d60d17059ff19a5f3563c0a3d4d7db8059c38d6af7ffcb670aff2"></a>

## api_specification.validation_custom_list.open_api_validation_rules.api_endpoint — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 6763535029c1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="canonical-2652b402fe29c4343bdc66089996bed34b0666a7c595648d72056bb1e8acf570"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-7103b2a1579da6e646a0556eee5a203b87f9827f3c7fd4b705022e0b33c474ed"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 6763535029c1 / 3

<a id="canonical-200887ec4e3ee3323e9c0f2711a60bd009c231b3c3f718f6f42c1f517d25478f"></a>

<a id="canonical-a63f60079727956fde7ff85ecc2e06e6e2ec57439f1ec97f8f642bc294e79bad"></a>

## methods property — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 6763535029c1 / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-71a8080043a4dd965da4e7b5e1eba520316ec4b6555a6665921e01da8ac60a94"></a>

<a id="canonical-3441301dd3750d9f7400ba562b8446f62be162d0cbc5bfca0165839a914074ea"></a>

## path property — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 6763535029c1 / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-33f94e785cbf5789f2a14f6e7f3bf9a78e719209599a1d36e17640f34d6705b7"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 6763535029c1 / 6

- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b23eec20ae3f17d1ed9959f4b5f98a9335fd50b9b0224dec4c35cd5935eb1c71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ab0db91498d82785d19f36a90653ffd9aba962ae39849389e3b9facfc70cac5"></a>

## api_specification.validation_custom_list.open_api_validation_rules.metadata — api_specification.validation_custom_list.open_api_validation_rules.metadata / da32a8a09b34 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- api_specification.validation_custom_list.open_api_validation_rules.metadata

<a id="canonical-b9b7af0ee432e011d33903175e43a13ea96144140355936de8178d13d3b8b8a6"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-aaf7bd9b34b34b45282e6f542be5c317f9027a7fa7b241863245d841f5df6df7"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.metadata / da32a8a09b34 / 3

<a id="canonical-6eefd9af05f26de0421ba3634f0cba55a4314e2b32af67d3baac37bf21e1a4ce"></a>

<a id="canonical-0f71751c3f0d4ff32100aa0bd71c33297e4bb8dc4980d8813564fc945f7a7400"></a>

## description_spec property — api_specification.validation_custom_list.open_api_validation_rules.metadata / da32a8a09b34 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-46d5ef0e691c28ef782c4ce280fc912f50a05fb2f4d4a9ad98e77c27612aa714"></a>

<a id="canonical-d8bbdd9344c47c2cf85cfa4349fbe09023fd874e61a01a6ca32085c78c149285"></a>

## name property — api_specification.validation_custom_list.open_api_validation_rules.metadata / da32a8a09b34 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-11dac8f94238e865a383e733ac707a18236e06b9a316adf9fd86a0fa2ee98bda"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.metadata / da32a8a09b34 / 6

- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-550996f640a69a36a934a67e7fbccdded8e8c2c594200a2d890506b9699f3300"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 5f29e46226d8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="canonical-3285ebb46920c121dc63823797ec9d389b486c55304807d7606e51ab5659c5fa"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

<a id="canonical-a34c13241c14924965b5c54851d112c123b66acd1b0316311efae3c91292b468"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 5f29e46226d8 / 3

- [response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-182138e21e9aece2eaa873353601eaf750c9f78ae92682bf9ae58f0174af368d): complete subsection reference.

- [skip_response_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-72805b8407a9f0acd0b741c8b9141abee1356ac9a1cd66d5662d59ee8b43e400): complete subsection reference.

- [skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-b4d497123b8b62ff8af1bbe9f12d7cdda5054bf806579cdb2146500b5220a8b5): complete subsection reference.

- [validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-8b57227d57f8b17f66c3f50189c69814be4ba8b9ca08ebb69e803ca9f06171c7): complete subsection reference.

<a id="canonical-0e11a5d813fefba0de647e6dd63e1c27b7a90e2be7f6e1d054d1a366a104c72c"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 5f29e46226d8 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-182138e21e9aece2eaa873353601eaf750c9f78ae92682bf9ae58f0174af368d)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-72805b8407a9f0acd0b741c8b9141abee1356ac9a1cd66d5662d59ee8b43e400)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-b4d497123b8b62ff8af1bbe9f12d7cdda5054bf806579cdb2146500b5220a8b5)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-8b57227d57f8b17f66c3f50189c69814be4ba8b9ca08ebb69e803ca9f06171c7)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-182138e21e9aece2eaa873353601eaf750c9f78ae92682bf9ae58f0174af368d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c13996a9b6bc8998c592300527069ef8194bc296f27cc91b069f76a46619d53e"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 6e0ef7a73305 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="canonical-cd24c7a8fc0738c01e1753df6e01a4b1c120f1abe1b033b5b03c401efbc4a6f1"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-9896229a43665e49fb98aa7a33079c2a5f6bdfa4a971a9c54f07b3d937ddd8ae"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 6e0ef7a73305 / 3

- [enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-f871bb22edd6306df312b4f0dcca6b437fd43026302c8187fe77154d7abe898d): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-d6afe108e76ce09e6b38b178d56e7d4d74d9008e19b40c32d5b79fa9b0afa1fc): complete subsection reference.

<a id="canonical-6050c5ff423aaabd6eb00ebc9abba0eef876bff390a6e3f846e95075675a5569"></a>

<a id="canonical-db8b3775d7caddd5a0f3727926c92b73f41c7b1bf5511b7dbb4b030d752a0583"></a>

## response_validation_properties property — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 6e0ef7a73305 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b32d4f7e5734c0f67978f3bd0330c3f47d9caf83884ac581e2d6079bad1bf4df"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 6e0ef7a73305 / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-f871bb22edd6306df312b4f0dcca6b437fd43026302c8187fe77154d7abe898d)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-d6afe108e76ce09e6b38b178d56e7d4d74d9008e19b40c32d5b79fa9b0afa1fc)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f871bb22edd6306df312b4f0dcca6b437fd43026302c8187fe77154d7abe898d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af008f90c453b493dbafe0bf813507865ddd92f7cc2660f6a7438952e172b834"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 03f9d265f25c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-182138e21e9aece2eaa873353601eaf750c9f78ae92682bf9ae58f0174af368d)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-1f0a07507b40e47f58d1ce85f52a994b95c2313bf61de491397a8f9131e5e87a"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-40d2559eec519965fa971c432a1f5a519b1cb85c490f2a5fd996f41f2877d4bd"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 03f9d265f25c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-779e4b0c9747b4f3ee6149e7239630ba818fcd61e15b4f90909f9e17c6c75870"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 03f9d265f25c / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-182138e21e9aece2eaa873353601eaf750c9f78ae92682bf9ae58f0174af368d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d6afe108e76ce09e6b38b178d56e7d4d74d9008e19b40c32d5b79fa9b0afa1fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed97ba6ca2f85f2cd502e8496fe39be671729fbf37cd4cbd5cb378665bb2f08c"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 713cfe560c17 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-182138e21e9aece2eaa873353601eaf750c9f78ae92682bf9ae58f0174af368d)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-443203754da7fa6d21961c0e14673f9d246ff99f2e0aeb83c1e71ba8e4359708"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3ea84a2799a08d9ab4b94395674e98b7796c7d7a2ccc952ac60cb6232928743e"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 713cfe560c17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b8c210083881f0e4f8ba9d408493f3ad166a4cef672db354fd2cbb20f05a2ba4"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 713cfe560c17 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-182138e21e9aece2eaa873353601eaf750c9f78ae92682bf9ae58f0174af368d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-72805b8407a9f0acd0b741c8b9141abee1356ac9a1cd66d5662d59ee8b43e400"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9937c3485ee769973dcb52ff3391f8922ca56c7933883231a2567b22aa2781ea"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / f326ba448899 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

<a id="canonical-f1e5d74a414b0ff2f77097a800e57adebe713442e6c8c34e912e20a4f9d1f634"></a>

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

<a id="canonical-672459063e5c28a0d7dc4b96a1d2cbe63bf291c4a2081aa5d0c09ef7c6252f7c"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / f326ba448899 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f81dbb181b00bab35d5a708d1ab5b11b9555f4a0b033913949f514df6d798c4d"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / f326ba448899 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b4d497123b8b62ff8af1bbe9f12d7cdda5054bf806579cdb2146500b5220a8b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c1a57ce11631d86f1cbed109928ce307faffa4270bed7c119c423c505bcd474"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / dd6e69952b8a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation

<a id="canonical-a4ea129333f24e7ef5ca129c96d63053397317d46b5ccb21a6d9cff9d6f76b3a"></a>

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

<a id="canonical-2205d0755395e57082a2cd88a1f5a798f93052f4d043f9b46d3c7b62731a26ea"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / dd6e69952b8a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cc4211b023166d02952e3f80841c7c516c885ccc52190c18eb6e560e4e49c1a1"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / dd6e69952b8a / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8b57227d57f8b17f66c3f50189c69814be4ba8b9ca08ebb69e803ca9f06171c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5df1733327d1814b8a51a8bd9427b4de38041a78a5a923b888967e21496280ce"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 09a6e33cfca6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

<a id="canonical-ddc1e1e343ada57c3d4895d803dbe99e4f789bffc902d89f3255b31782719651"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-1c82e7d4b2fbbcc15b7560801f985f8ec33af5747809602f5ab93227374dec6b"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 09a6e33cfca6 / 3

- [enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-24d41cf741db5f4e3ef1c71eb2c63548a96d9e5bdd1d6a3a2f2caae43ba4bee5): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-37548aaedfe9dd1a091ee15aadabcbf329356b24157b41a52a6ef6fc23af4235): complete subsection reference.

<a id="canonical-7411c7789dd80afe33907f89a63124fc4ba55d0fb1519ff48588743a8de3c684"></a>

<a id="canonical-dc613aad5fe9d68c3156e83a11ccc99a4214a6fb8a1cb9a3ef122bb3691607d7"></a>

## request_validation_properties property — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 09a6e33cfca6 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2246c0b1dfe7c77fdaaf7f3a6edae32a430d9d467b69e4ccced3cd11d487029a"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 09a6e33cfca6 / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-24d41cf741db5f4e3ef1c71eb2c63548a96d9e5bdd1d6a3a2f2caae43ba4bee5)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-37548aaedfe9dd1a091ee15aadabcbf329356b24157b41a52a6ef6fc23af4235)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-24d41cf741db5f4e3ef1c71eb2c63548a96d9e5bdd1d6a3a2f2caae43ba4bee5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c71e28b21906d860fcc4d0fcf8cd148eba87ef7a32f65acfc96209641717c690"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / b209b5206d0a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-8b57227d57f8b17f66c3f50189c69814be4ba8b9ca08ebb69e803ca9f06171c7)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-21d22f2caaa29c5ecc2ad901792c13530f253dcb52ba7409aa19d748cc535c5b"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9f8c86e246c8c6c11e7d1ddbc5c20ec6aedcccabb5c92d4f92eb5cf4c7f3b263"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / b209b5206d0a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-63dc0171c8128985eea08788f0b2cebeb926c6df6d65b33cbf59f1037345ec46"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / b209b5206d0a / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-8b57227d57f8b17f66c3f50189c69814be4ba8b9ca08ebb69e803ca9f06171c7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-37548aaedfe9dd1a091ee15aadabcbf329356b24157b41a52a6ef6fc23af4235"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b657eea297980f748733405e3abe389891ddc69a87c663542d014677c9be7db7"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / fbaca8c56ff7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-d11c41f95cc96407bb464f0daefe1adabb6bcf5f558dbb26dd97716126ea3cad)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-70cad725a6ee4571233b10976877c9d6b3d3a8df68ff7e572c23f44eca33b15c)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-8b57227d57f8b17f66c3f50189c69814be4ba8b9ca08ebb69e803ca9f06171c7)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-af2cf75390b66c56af1ad9acb09290cfd4658042eeba4da4dae719302fe9dfb9"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-48bbf85896be7fac9e2d3803ce0a5be31a38285cd4d3d3147975693d0d56d994"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / fbaca8c56ff7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a266fbc1f3b386b2c1c17a723302add400f94d3f3e821954f619d5c2346558de"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / fbaca8c56ff7 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-8b57227d57f8b17f66c3f50189c69814be4ba8b9ca08ebb69e803ca9f06171c7)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14ef4ddba23cf18a4472bec60d76285ac94709ba15153a1131cb30da3b79128c"></a>

## api_specification.validation_custom_list.settings — api_specification.validation_custom_list.settings / 26c064f59b91 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- api_specification.validation_custom_list.settings

<a id="canonical-109542e6fdd6e6feac4aa06ae92bd279425a94df29562fbc1df2e7e3b9a3bc0c"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

<a id="canonical-1b6ed56cfc9ad21c6b8fc236301079670fb5288d23faf30ed6d6693b54fb4fa2"></a>

## Direct properties — api_specification.validation_custom_list.settings / 26c064f59b91 / 3

- [oversized_body_fail_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-17b9510537329f2f74896ccd27b2e9df329ff245f93a315041e2a2807be7f90d): complete subsection reference.

- [oversized_body_skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-1c54bc27435e7bd4260f18c4283fe8c3f108be6b4a701ae0ffae03625630849c): complete subsection reference.

- [property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-4b4f6a3847c70ede233624be771e13045e5019db10cab805633f1ad46b8aa1cf): complete subsection reference.

- [property_validation_settings_default](data-sources--http_loadbalancer--reference--group-009.md#canonical-af320b398a6c6bce45cd3940f73f1a46881033948dff66fed37171f8dd17e0ff): complete subsection reference.

<a id="canonical-9e3ca1101f30c3f12379fd47f617d34c9a92636590b19d49879d9abf52d787ef"></a>

## Next pages — api_specification.validation_custom_list.settings / 26c064f59b91 / 4

- [api_specification.validation_custom_list.settings.oversized_body_fail_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-17b9510537329f2f74896ccd27b2e9df329ff245f93a315041e2a2807be7f90d)
- [api_specification.validation_custom_list.settings.oversized_body_skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-1c54bc27435e7bd4260f18c4283fe8c3f108be6b4a701ae0ffae03625630849c)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-4b4f6a3847c70ede233624be771e13045e5019db10cab805633f1ad46b8aa1cf)
- [api_specification.validation_custom_list.settings.property_validation_settings_default](data-sources--http_loadbalancer--reference--group-009.md#canonical-af320b398a6c6bce45cd3940f73f1a46881033948dff66fed37171f8dd17e0ff)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-17b9510537329f2f74896ccd27b2e9df329ff245f93a315041e2a2807be7f90d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9acdb570614868247ceab57fec23684f33c272b3ecb1e70cb9892a06279732ae"></a>

## api_specification.validation_custom_list.settings.oversized_body_fail_validation — api_specification.validation_custom_list.settings.oversized_body_fail_validation / b9eebf1e5de6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- api_specification.validation_custom_list.settings.oversized_body_fail_validation

<a id="canonical-216694fa0cbaef3e4514ac1161899c69b482f5165a7f240ce75c86b2cdee751e"></a>

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

<a id="canonical-517c7c3f65d93c3a8f7a4e7b58900cf84b5959ba0ade9ea20a89e53c843bac7f"></a>

## Direct properties — api_specification.validation_custom_list.settings.oversized_body_fail_validation / b9eebf1e5de6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-103989e9a6b52096cb8f36c236cffe9cd9ce3c3431cb68c5932c1755af0db038"></a>

## Next pages — api_specification.validation_custom_list.settings.oversized_body_fail_validation / b9eebf1e5de6 / 4

- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1c54bc27435e7bd4260f18c4283fe8c3f108be6b4a701ae0ffae03625630849c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9203b1a87fb2558c87e5b7c44702f0183578663d23ee4e73669aa8c8c2915b2"></a>

## api_specification.validation_custom_list.settings.oversized_body_skip_validation — api_specification.validation_custom_list.settings.oversized_body_skip_validation / e9e7752bbf3d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- api_specification.validation_custom_list.settings.oversized_body_skip_validation

<a id="canonical-fd2fc581e67ec09f3496501dfd94390f259c4f13ac5bd0d07ce17ddad3895c3f"></a>

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

<a id="canonical-e9bd230f9c43f4e474c393b295d1174370ad8d3bea4200655d595f48033ac975"></a>

## Direct properties — api_specification.validation_custom_list.settings.oversized_body_skip_validation / e9e7752bbf3d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-233a375cbff9dab1b3ba94adb5dd7a32eb7ffb9d823214aad9d453ce464329eb"></a>

## Next pages — api_specification.validation_custom_list.settings.oversized_body_skip_validation / e9e7752bbf3d / 4

- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4b4f6a3847c70ede233624be771e13045e5019db10cab805633f1ad46b8aa1cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43b4df0203b3a15a2104c41a402648a24b4df05ceee7cca6ab05667411393f7f"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom — api_specification.validation_custom_list.settings.property_validation_settings_c / 5d6fa1b742b7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- api_specification.validation_custom_list.settings.property_validation_settings_custom

<a id="canonical-114250943b6aada12bd173bb25067b1c9c01983e8ee0482b2a4482a6ad18d0a1"></a>

Type: `"single"`. Computed.

Configuration parameter for property validation settings custom.

Upstream description:

Custom property validation settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3087dfa34fcec2a3479554c18b1339b2eb9a2ad54f90ec4024e16b7eea2e0c29"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / 5d6fa1b742b7 / 3

- [query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-4d941aef4fcce0005b09e3b8934d365897945931d36681c9f33703f46c71f488): complete subsection reference.

<a id="canonical-278117c2c112b911a6d4a3719d14892a04e613cce326c4e65090944024866b86"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / 5d6fa1b742b7 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-4d941aef4fcce0005b09e3b8934d365897945931d36681c9f33703f46c71f488)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-4d941aef4fcce0005b09e3b8934d365897945931d36681c9f33703f46c71f488"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b1a510c4fba852a412961343880c4cf49017cc2823fe023bbfa97aa98ee3264"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters — api_specification.validation_custom_list.settings.property_validation_settings_c / c0928c9d9fde / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-4b4f6a3847c70ede233624be771e13045e5019db10cab805633f1ad46b8aa1cf)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

<a id="canonical-615816271d89e1a224f00b87e49c6fb12166a8994810096a0e6a79467f16b19f"></a>

Type: `"single"`. Computed.

Custom settings for query parameters validation.

<a id="canonical-a20cb9479b477dc23d0024790c7da64a70bf7070ddd837a7110f42449c18e939"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / c0928c9d9fde / 3

- [allow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-cb4cc92363e2965f2701263b58f4ec74ed5df31b144b51efda2fde81b78b8676): complete subsection reference.

- [disallow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-715974f2530fbdbda377edeefb8958e86c355c6082473157f8b02f659a7db108): complete subsection reference.

<a id="canonical-8c4a4fbf3ad11007c840399cf4db2f28f7aa36f9eca1072c8962d169427150e0"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / c0928c9d9fde / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-cb4cc92363e2965f2701263b58f4ec74ed5df31b144b51efda2fde81b78b8676)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-715974f2530fbdbda377edeefb8958e86c355c6082473157f8b02f659a7db108)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-4b4f6a3847c70ede233624be771e13045e5019db10cab805633f1ad46b8aa1cf)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cb4cc92363e2965f2701263b58f4ec74ed5df31b144b51efda2fde81b78b8676"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5431f608f4d16b49bf1e063ab97be3f59f5078e059c593758961fb3484f1e39b"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters — api_specification.validation_custom_list.settings.property_validation_settings_c / 661343f6d3c1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-4b4f6a3847c70ede233624be771e13045e5019db10cab805633f1ad46b8aa1cf)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-4d941aef4fcce0005b09e3b8934d365897945931d36681c9f33703f46c71f488)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-f0e48b340517e3a7f393bddb12e78e00b74054ec8ec76c45669613f93df573ec"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow additional parameters.

<a id="canonical-a701c9cd6bb1fa1ff9df06e118a6b56f8fac2c2bd99a5a06d88f25a334fe56a4"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / 661343f6d3c1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ba7bd4035a6f30bd9fcfad19c1589fbf36d6a9fabb5537cd368df2488385d74"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / 661343f6d3c1 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-4d941aef4fcce0005b09e3b8934d365897945931d36681c9f33703f46c71f488)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-715974f2530fbdbda377edeefb8958e86c355c6082473157f8b02f659a7db108"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab0894e4b000a24b3441e35d0e8baa4ffc0b05bac98e98af28157bb32230968a"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters — api_specification.validation_custom_list.settings.property_validation_settings_c / bac4db92e5d4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-4b4f6a3847c70ede233624be771e13045e5019db10cab805633f1ad46b8aa1cf)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-4d941aef4fcce0005b09e3b8934d365897945931d36681c9f33703f46c71f488)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-24650e4833cb9f5cb347a0d85eaaa922f5a37927f3234f7a553d3cb2db2329a6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disallow additional parameters.

<a id="canonical-374f7148a1688a9c6c826fa3d3a74edd76d9f2305ccf0264f9c9979e51ad731e"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / bac4db92e5d4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66f452e2e93385d6b6867f6408de8aa62437c65fe985c75e52306b889a0f9c41"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / bac4db92e5d4 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-4d941aef4fcce0005b09e3b8934d365897945931d36681c9f33703f46c71f488)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-af320b398a6c6bce45cd3940f73f1a46881033948dff66fed37171f8dd17e0ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccec5feb259dc2e9c624ea8824ba710d92da895362466202d10cd5121932eac0"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_default — api_specification.validation_custom_list.settings.property_validation_settings_d / 4058969ea6f0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-e7f03d4ed700ce42ee78d7ac35359f0504cd8478bf7eb8c1ffe7b2174aa71b73)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-dc3ee214f76bab280713a67c44a41aa6336b3e7b39632414b5a712c3e621c588)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- api_specification.validation_custom_list.settings.property_validation_settings_default

<a id="canonical-935ddda1ce7250182a936cb71e0b24695893678ff693f91363c06edd3c20b617"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for property validation settings default.

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

<a id="canonical-2de5944e1a12b41e247561d62727e1582933f5286915e103ec28195cb531269d"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_d / 4058969ea6f0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d562b99429362e57ea2013b1e91355f3b861b40f7840c95c5a8f5b53d5b1b054"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_d / 4058969ea6f0 / 4

- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-56832f3c3cece50dff88b022d12312bb8a9be1db1006e5ff7193337dbaa1fef2)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8876fd0662bbfd750e6392c12ef9cfd61476036241f3d647e1ed210220a7cd38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
