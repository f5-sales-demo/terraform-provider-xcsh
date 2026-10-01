---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-60befd67ad54fe1683b180b53025ff1e9e8df277126a5f487fdfdb2415f9c396"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / 52914c4fbbba / 4

- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b14e74d4a3f596f591fd0637812bebaa2f2933dba3da6f2ad5048edf88b112bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c0e05c7481b74860c9302240d8747b5ba6ba3a71010389e28ae5c19c739c228"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / 3051f5267fe0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-9f4ca8c38c716850aa03e2e30a27f0be5f4c1aabeba859b37963c2da74f7b770"></a>

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
oversized_body_skip_validation = {}
```

<a id="canonical-95115187e6b2869b5b0c0e774a700a8e7d7d1be4c63d139cf460df4d12754e48"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / 3051f5267fe0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6515793c33b4eef453279ef7a53bf2af20cce3c50a888913cd7640d4e19f5990"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / 3051f5267fe0 / 4

- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f764aee52452acb31cf4503aae35affbf273da93db1d51f138929d61a7a98e5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7061873fb29f1958ebe0f93e9faa8ac831d3f9824464e1dc50d8e8ee703d71b"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom — api_specification.validation_all_spec_endpoints.settings.property_validation_set / b969017af973 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-974a2a04eb216ef15e7b9852a8a0ecea46c065d6e46283df6a1b6c91175d701f"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
property_validation_settings_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-f7cfcc884b12487da5e0b5f214bdaf3e2f9edfaaa7669e556b4d40a8220bb599"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / b969017af973 / 3

- [query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-7b32bbd4d23eca8794ae734a72aea10081147e9c7b5c4b76ee337da7a9995b1b): complete subsection reference.

<a id="canonical-d2796f0fd70cdc242f8f5071cbba556460e97dfab13c9bd886c8a361e8952f41"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / b969017af973 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-7b32bbd4d23eca8794ae734a72aea10081147e9c7b5c4b76ee337da7a9995b1b)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7b32bbd4d23eca8794ae734a72aea10081147e9c7b5c4b76ee337da7a9995b1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-926220e41ca250328d1178402b8e14dda7eb3b7cabfbe6e523aff78c036742ce"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / dca517d521c8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-f764aee52452acb31cf4503aae35affbf273da93db1d51f138929d61a7a98e5b)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-8895ed79203303e401d7bdb0836e138f5eb0ec02a5f1bbb92d536a325e33bb62"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow_additional_parameters",
    "disallow_additional_parameters")}
```

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-919b9a6e439449ac624175d0e7804035399e7d1dd693fbcfd9e2de3e5052eed1"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / dca517d521c8 / 3

- [allow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-472b5c8e959e1edccafbcf7c19852cbe9ff5940acc0b46da1397379f0f7f617a): complete subsection reference.

- [disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-bcfbb49a187cf4a4a7afd8a3defce3f2ee6e2079d6de0e05cf97a1fc2b7d18c7): complete subsection reference.

<a id="canonical-0b822f3c7983e76aa0ef0ad177124dcb6c8739932b9499af0bca647ba62f0ba4"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / dca517d521c8 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-472b5c8e959e1edccafbcf7c19852cbe9ff5940acc0b46da1397379f0f7f617a)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-bcfbb49a187cf4a4a7afd8a3defce3f2ee6e2079d6de0e05cf97a1fc2b7d18c7)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-f764aee52452acb31cf4503aae35affbf273da93db1d51f138929d61a7a98e5b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-472b5c8e959e1edccafbcf7c19852cbe9ff5940acc0b46da1397379f0f7f617a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e1915f32a8f7b43c3c33b90ac8b6b1e9f86a682ee66097ee7ff563470e6bb71"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / eeb0d10e7026 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-f764aee52452acb31cf4503aae35affbf273da93db1d51f138929d61a7a98e5b)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-7b32bbd4d23eca8794ae734a72aea10081147e9c7b5c4b76ee337da7a9995b1b)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-c28579ce932cebb93a7c00658eab87403be55cd6755cf206908b97fe93c7d4c4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow additional parameters.

Terraform syntax:

```terraform
allow_additional_parameters = {}
```

<a id="canonical-d7cab4f1a939e0818d92b0d5d93eea5cdff933bd541197aadbe6788861fc4f32"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / eeb0d10e7026 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14f9f5807275a985fdf1c4794e2186e348e89e949085067a8f2e2ee6991c9a25"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / eeb0d10e7026 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-7b32bbd4d23eca8794ae734a72aea10081147e9c7b5c4b76ee337da7a9995b1b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bcfbb49a187cf4a4a7afd8a3defce3f2ee6e2079d6de0e05cf97a1fc2b7d18c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4c2096136546a89086629c8d11d96c4df1a37c08bbc2718b7120bb49691e122"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 7837d8a945dc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-f764aee52452acb31cf4503aae35affbf273da93db1d51f138929d61a7a98e5b)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-7b32bbd4d23eca8794ae734a72aea10081147e9c7b5c4b76ee337da7a9995b1b)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-955f1607e4ac233a68a19488ec8afb294517979436abf9bb1dde7dd52284cae4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disallow additional parameters.

Terraform syntax:

```terraform
disallow_additional_parameters = {}
```

<a id="canonical-4928dd0571a388879b8571aad4169d18169b4fd40896059a49d994256402f54d"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 7837d8a945dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9a0e340b18dc5c31f02546cd2e568cff273eef263768bf73719e56762d5a512"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 7837d8a945dc / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-7b32bbd4d23eca8794ae734a72aea10081147e9c7b5c4b76ee337da7a9995b1b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-dece8263d52bfb4349340b93842859243cb81e5034097e8d8710701237d5061d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d20c6d9102eefbe55bba5b19cb94463b46f8be1f91977166880f5247a509a4aa"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f2857dddb5d3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-1905c5995bd21e42c4def1409aaf4c6b2e90271d68a353fa8589f251efd60939"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
property_validation_settings_default = {}
```

<a id="canonical-c3b230525e8836e48f88eb180ec75077cee364312307f4055f78b453ba676c4b"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f2857dddb5d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c40faee7889eeb5e6b7cdd96b83dc8b4fa4988e9f57d6298360e41ae3509fb06"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f2857dddb5d3 / 4

- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-005.md#canonical-a66a49e08574d71974d3dce003662351a6bf6d897f2b85ae140c3b35e63b68f4)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1272d521275fa0e4674868818219762e10301e3f8e21aaa7040914d8ca4a8b79"></a>

## api_specification.validation_all_spec_endpoints.validation_mode — api_specification.validation_all_spec_endpoints.validation_mode / cb894d6adf92 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="canonical-266c39b6d5ebf0151037b4113f89189f0a336885f21eeda8bae9ca292ea3ccc5"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("response_validation_mode_active",
    "skip_response_validation"),
  validators.ConflictingObjectAttributes("skip_validation",
    "validation_mode_active")}
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
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-a0d5c0d2439c05d95a69a929f9d96fdad9bce36c72ab602c73dd1a6c2fbc5d0f"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode / cb894d6adf92 / 3

- [response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca6b45ea232954a9e240606ddf01f1290a2d90878306ada53ce781b274acd12f): complete subsection reference.

- [skip_response_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-72e43a0a9e17e9c37af4b8e920bc53b9ed476314b13a5d53ff6d7b0a6d8dec91): complete subsection reference.

- [skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-160a43528a56c16cd851974093771b2c6a9fed6e661eb3df98d8bc550ac58d23): complete subsection reference.

- [validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3adb166aec589e3abe1a1bd98db790608c8656e5d21bbc89e1d0f8dc73c34ea6): complete subsection reference.

<a id="canonical-d234933f4c5ea7ef4cc764a79422b988d3ee8977a95f1f45bd74b39e57e05d62"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode / cb894d6adf92 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca6b45ea232954a9e240606ddf01f1290a2d90878306ada53ce781b274acd12f)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-72e43a0a9e17e9c37af4b8e920bc53b9ed476314b13a5d53ff6d7b0a6d8dec91)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-160a43528a56c16cd851974093771b2c6a9fed6e661eb3df98d8bc550ac58d23)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3adb166aec589e3abe1a1bd98db790608c8656e5d21bbc89e1d0f8dc73c34ea6)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ca6b45ea232954a9e240606ddf01f1290a2d90878306ada53ce781b274acd12f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5d9414e04113feb9e57fc96fbf5919313026ddda1672b19d5e8473711a46d22"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 4e159095bb71 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active

<a id="canonical-375074847f2d8d88d3f012491bd0c608d7ae9372e8b927e4e2a1826393df348f"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
response_validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-b52649a76bfda98b1ebb823788fd9cb7f02fbcf775651787a4f216fb4b910318"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 4e159095bb71 / 3

- [enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-24cbfd346dd66a86dfe074695996c6015cd64ed60f6b77f7bbd2be47cf7b3738): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-e364e3fa3d9287967d45dd3e4753ca1d2da402eaf413c5adeb93d20429601279): complete subsection reference.

<a id="canonical-6b595837a2235dc2339995a11a6a332098761f0be7ec6587180f1124c08fd445"></a>

<a id="canonical-8056a18356c99efdd1644d179331a096c61028c029e72d6981f3506cff8fccab"></a>

## response_validation_properties property — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 4e159095bb71 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

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

<a id="canonical-427955de60a0b3bd82adc5b576f3419e085139678359093bc715feb4f7d5f5ac"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 4e159095bb71 / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-24cbfd346dd66a86dfe074695996c6015cd64ed60f6b77f7bbd2be47cf7b3738)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-e364e3fa3d9287967d45dd3e4753ca1d2da402eaf413c5adeb93d20429601279)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-24cbfd346dd66a86dfe074695996c6015cd64ed60f6b77f7bbd2be47cf7b3738"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6b88549ac909da374089cc58bc2d7fbf18d3804043b67c1f90acf9bb32097c6"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 2cd92dd3c89a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca6b45ea232954a9e240606ddf01f1290a2d90878306ada53ce781b274acd12f)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-4fa2ca75a494cbf211fc522f6e710acf7336cfe00cc13ccc263b9a84b3ce7698"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_block = {}
```

<a id="canonical-ebb22e6f8fac9f574dab4cee547d7f5ea7b24f724bf267b922f663f12872f008"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 2cd92dd3c89a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-30d6085f56b33609076ce64f2bda608d55ae8eaf6677dbd94c000ca5389ec59f"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 2cd92dd3c89a / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca6b45ea232954a9e240606ddf01f1290a2d90878306ada53ce781b274acd12f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e364e3fa3d9287967d45dd3e4753ca1d2da402eaf413c5adeb93d20429601279"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57a6e92c3d356f59fe20d448c434c087c1c34ab0f0bef57a0209db72e71c7b9d"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 0efc25fd05b3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca6b45ea232954a9e240606ddf01f1290a2d90878306ada53ce781b274acd12f)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-b43a077e2d61a63c9dbecf113ae1a4298ceff3092b7f7c13fa419e86063dc5b8"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_report = {}
```

<a id="canonical-7112d50ae7352f9c4b23bec9f722cffbb4ac8db97e6556e0883f26d6ba884887"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 0efc25fd05b3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-49ab5d94c6b2c6b63511ba6e437a47564dc32958cd589546aa08259a7795d32c"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 0efc25fd05b3 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca6b45ea232954a9e240606ddf01f1290a2d90878306ada53ce781b274acd12f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-72e43a0a9e17e9c37af4b8e920bc53b9ed476314b13a5d53ff6d7b0a6d8dec91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1964dfe1b4cab6f1ba04706770d224794e53a9b7b2a867bafafec8a13cb2ecfd"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation — api_specification.validation_all_spec_endpoints.validation_mode.skip_response_va / 54433da190cd / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="canonical-88f25ef266a1ecb9408aef924e09f931de895a81c5642bf644caf9eebe2921a7"></a>

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
skip_response_validation = {}
```

<a id="canonical-5c727f87f8c966c985dce2d900575ee4e432c36d55c65ee5d6fc6868d7da5e23"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.skip_response_va / 54433da190cd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-965ac13ed517d0b8f42dedceca748ef02204bc89a60a90c22938cb6005ab9bd4"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.skip_response_va / 54433da190cd / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-160a43528a56c16cd851974093771b2c6a9fed6e661eb3df98d8bc550ac58d23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d14d5a55d3908366b87ac1e8ad4462d15120427ba43c3caa040584069f68958c"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_validation — api_specification.validation_all_spec_endpoints.validation_mode.skip_validation / 9db774ecc599 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_validation

<a id="canonical-c4a4f354a1d54be73dcea7a8547abb0dc87eac234e8ced237f129359b6330e37"></a>

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
skip_validation = {}
```

<a id="canonical-d8814495e0348fb81e725c07f5e856b52f9e633bc822e14f93ee7c37b5a86b57"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.skip_validation / 9db774ecc599 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3f3e0586e06ac054b9af2831fbc770ae390dc625951f76e0a8b32387d5250a78"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.skip_validation / 9db774ecc599 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3adb166aec589e3abe1a1bd98db790608c8656e5d21bbc89e1d0f8dc73c34ea6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-950c3f1f6306173b1677d8c3390fd16498f92a3286148b345880b3b9622832ec"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 36c813b2123d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="canonical-1209db64a1da94b70dc3006cf4bf8f453b58ed53e55413305b16ecb0e6ce390a"></a>

Type: `"object"`. single nested block, Optional.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-e99c451b7ef9493e1bd379183043e728656c62563995d2ae5b1d4f099220a0d2"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 36c813b2123d / 3

- [enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-82472cf4822a6f2320dd8f0d43242e0bf86978b2fc834451e23b2854d1c42a40): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-4e7d8d9916850d9bf8b85d56a0dfce4b5624d487891224a675ac80c85a0d35a1): complete subsection reference.

<a id="canonical-0eadccda1d8ff05ddee06df8b03efbe3c1ddca61db7b27d83851e1c090ceae4a"></a>

<a id="canonical-cc820578f8210fb563689b25b6a6f1f38a912a36410ef549988f853871ae9845"></a>

## request_validation_properties property — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 36c813b2123d / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

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

<a id="canonical-df6e407a9ecfd657c37b00d66d293ae26b993272c4227846f2af0c1feab2d30c"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 36c813b2123d / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-82472cf4822a6f2320dd8f0d43242e0bf86978b2fc834451e23b2854d1c42a40)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-4e7d8d9916850d9bf8b85d56a0dfce4b5624d487891224a675ac80c85a0d35a1)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-82472cf4822a6f2320dd8f0d43242e0bf86978b2fc834451e23b2854d1c42a40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a1ee13c15dfaeaffe99760668ad53fe926650bc61f78539d9854677504c9569"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / d57d13707047 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3adb166aec589e3abe1a1bd98db790608c8656e5d21bbc89e1d0f8dc73c34ea6)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-e1e4615077397389cb2036b6de9c4bf045f34131c3f547403c9aecc0786abca6"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_block = {}
```

<a id="canonical-74ea1f14db4a695311a4362774a0c1f743700deed53dcf95f7be2b7cc208d427"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / d57d13707047 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-271178854dd8e3fcc7068c52a71253788cd70858f6a2745da34a0ccf0783bb88"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / d57d13707047 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3adb166aec589e3abe1a1bd98db790608c8656e5d21bbc89e1d0f8dc73c34ea6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4e7d8d9916850d9bf8b85d56a0dfce4b5624d487891224a675ac80c85a0d35a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77c3c8b9018aca7dbfd8c535f57efeceec9636b64aab9fefc79a86200bee76dd"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 52cc73cbde8b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-fb8c3e0c6edded09135946ecd8e68aca48508a4004f293de1c051d2a2d12f5bf)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0e02e88b49ceb8eeb5b6a7013153c2a07a7c8b4d1a06d15dde773cd3479ac01d)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3adb166aec589e3abe1a1bd98db790608c8656e5d21bbc89e1d0f8dc73c34ea6)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-9daacdfcc13b81470bcd49fac4e4a295b0857bb60b9c269c2c863f5824978ce0"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_report = {}
```

<a id="canonical-0f996deb07467a7bb0e5ac95bd31701829c7274dfad262c89bef159ef0326c69"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 52cc73cbde8b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-960f1d6cc0c7b86c5667980c5ae6a9976a929871c06a449f6a1b430e15d39599"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 52cc73cbde8b / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3adb166aec589e3abe1a1bd98db790608c8656e5d21bbc89e1d0f8dc73c34ea6)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ff4bf5fe1f6567ba88e93d9b8e9dcb9f3d66dc121f23fd73221a2fe25e7578b"></a>

## api_specification.validation_custom_list — api_specification.validation_custom_list / cc8632830bab / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- api_specification.validation_custom_list

<a id="canonical-cac995b862b3add20094256fc4313ed9e07b9ae0ff179b33db2e63be168d919d"></a>

Type: `"object"`. single nested block, Optional.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Upstream description:

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to "Fall Through Mode".

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

Terraform syntax:

```terraform
validation_custom_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-87f306f6aab42351689cc023dc113b4c7ff8ed08d85d1149e1a58236e6596ca0"></a>

## Direct properties — api_specification.validation_custom_list / cc8632830bab / 3

- [fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac): complete subsection reference.

- [open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd): complete subsection reference.

- [settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def): complete subsection reference.

<a id="canonical-e40e30eaf11ce0eb8cc077f2e1de5dd539b2782f4324c0c79560006df6027bcc"></a>

## Next pages — api_specification.validation_custom_list / cc8632830bab / 4

- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd2607daa04701c37b07cc85e1486c4fd7e525cddecce5b55523078ce7bed2d9"></a>

## api_specification.validation_custom_list.fall_through_mode — api_specification.validation_custom_list.fall_through_mode / 3c8c4bd42644 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- api_specification.validation_custom_list.fall_through_mode

<a id="canonical-e8e157075de3352c6ebe4089fd34847115fdab0bc698d7f32435115d52785b76"></a>

Type: `"object"`. single nested block, Optional.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("fall_through_mode_allow",
    "fall_through_mode_custom")}
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
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-93ab001b99fe7eb8101e79984f475b716a2c12365677de89e37195bf95f317c8"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode / 3c8c4bd42644 / 3

- [fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-006.md#canonical-15f00a43a3117ec5bd22227265271f10529c9ec1068dd9edbae10fabc3734e6f): complete subsection reference.

- [fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225): complete subsection reference.

<a id="canonical-b2de56ee7abaecf2b3f57d62e2e6e1f10e7b6b8bb3f9542551cb6e946feaafba"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode / 3c8c4bd42644 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-006.md#canonical-15f00a43a3117ec5bd22227265271f10529c9ec1068dd9edbae10fabc3734e6f)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-15f00a43a3117ec5bd22227265271f10529c9ec1068dd9edbae10fabc3734e6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90d0c052fe2da0e878c599f47c3f66d8dccedaa0b94c56e3e8ed314acd79aaaa"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_all / 620ff4354d46 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow

<a id="canonical-e18d487aa2be8adb9f21513593883eefa2c83220549816a72943e7e6ae89442b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
fall_through_mode_allow = {}
```

<a id="canonical-cea8011dc902be698a69c690aeacb566d8bde843fecc6386636364506eb41aa6"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_all / 620ff4354d46 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-168e949e80c862b977380a6a7984f96f8ac0ee84c009879434bdeffee3ac327b"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_all / 620ff4354d46 / 4

- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78434b262a7494f20eb2bce32d869ef8b2937a727d74e7fdbc9a96b8606f2afd"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 670d9d483caf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

<a id="canonical-12cd54488977f8ba1b828055e45a7eae6087f2a298eb5fc00dd51f0a79923cd0"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-802228b8afa31149f8f35a94270982f075e86f05b5a312833c4740f6c3725ff8"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 670d9d483caf / 3

- [open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f): complete subsection reference.

<a id="canonical-14be60bbba0fa7145b1066e89211cab29618e2673ca37c4da6e391b0b11f2dce"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 670d9d483caf / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02bfd5c8204e9a4d154dca0fc7a668908b960596241a94e0d66654ef56b89e86"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 42c3d9170ae4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-a697cbe9fe1fb595756edc58a539cffa37cdece8ba6f30263e87ca1b366d26d0"></a>

Type: `"object"`. list nested block, Optional.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("action_block",
    "action_report"),
  validators.ConflictingListObjectAttributes("action_block",
    "action_skip"),
  validators.ConflictingListObjectAttributes("action_report",
    "action_skip"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

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

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-75bdb95b1473eb9c7e74373e214fcbbb19c4d3585161bcd31fb2191a2785a5ec"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 42c3d9170ae4 / 3

- [action_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-3a1e57b25f52f15d9acbfed970a417df53b959f472a8d29d864cef40dc078288): complete subsection reference.

- [action_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-c23a2c3a693f4fd9c0789acdfcb4daf8a2cd960dd5c2b4719a3348b05c064dd9): complete subsection reference.

- [action_skip](resources--cdn_loadbalancer--reference--group-006.md#canonical-63e8ee15c5bf2f2a5df8a190464078ffb7e97ede0865cb174057a3b04f08edd4): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-6f1761e3bf021cd742447f098e2467e72c67497e14b399c4be77c6748d790635): complete subsection reference.

<a id="canonical-8538ed2e58648fef35ce6e9f41d2babc090b437af325d9b07e80e2e7851a246a"></a>

<a id="canonical-4945bb79a444a8dd7cbd921d83e218fce412891783a7efef902fd6d1b409dd0e"></a>

## api_group property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 42c3d9170ae4 / 4

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

<a id="canonical-6adf8a51b8fa11bff0b2c4500a7129c57f7f7f5c85dc99900486be60050cc083"></a>

<a id="canonical-6317cd1468e615bf838daeff6c6b65b7e16644d3965b41063191ca9fb642ce92"></a>

## base_path property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 42c3d9170ae4 / 5

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

- [metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-a741c294a36999450e979f4592ddb1f9ba18a91e56940e1d6bcb9628c519869d): complete subsection reference.

<a id="canonical-92c7f96f40e9ab6bd329c4dddeb4420ea99766c5642167db0808cbf9bf992b07"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 42c3d9170ae4 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-3a1e57b25f52f15d9acbfed970a417df53b959f472a8d29d864cef40dc078288)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-c23a2c3a693f4fd9c0789acdfcb4daf8a2cd960dd5c2b4719a3348b05c064dd9)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](resources--cdn_loadbalancer--reference--group-006.md#canonical-63e8ee15c5bf2f2a5df8a190464078ffb7e97ede0865cb174057a3b04f08edd4)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-6f1761e3bf021cd742447f098e2467e72c67497e14b399c4be77c6748d790635)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-a741c294a36999450e979f4592ddb1f9ba18a91e56940e1d6bcb9628c519869d)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3a1e57b25f52f15d9acbfed970a417df53b959f472a8d29d864cef40dc078288"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51173d6a8fe11b30ecba05d435254d34ad289b81319ee6965600aa16c87f0b37"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 4a0612c644ad / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-c27a64a04574d4c9a4aa36eeedfd8b11b89205f5bc3ee42de73dad6932c09d83"></a>

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
action_block = {}
```

<a id="canonical-542defcd95a1fe00a785fd959e01c6edff50f6ebea4234e78d9bee9dcf907b83"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 4a0612c644ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8bc7491a9bd4e87768aead21da30be35bf9e6d5b86e381d95dc2bedc24a05364"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 4a0612c644ad / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c23a2c3a693f4fd9c0789acdfcb4daf8a2cd960dd5c2b4719a3348b05c064dd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bd2f049f6d0c9c41a46b51c85e2f8aae0eb5b6b2a322325ac433fd788a3f4e5"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / e6dcc5e28d88 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-1136e70fefe68db01282156f463d45d44b74e70be11e31c9b48d3795b9be6440"></a>

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
action_report = {}
```

<a id="canonical-9aba1ec10266776f6f889db4c997f2c0ef336f53779fc2806f434783812c4ed6"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / e6dcc5e28d88 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f541e6243e9a94401224f5e1da07bc64da77965935f5dd854d02ae17bc74c64"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / e6dcc5e28d88 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-63e8ee15c5bf2f2a5df8a190464078ffb7e97ede0865cb174057a3b04f08edd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76bdc9104610337949ba9a0e52be31e07e719a16f30e71e7eed6d0705168fb4d"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 0c1d746aba1b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-c44a75febadac4a9c738e7aced302b4964e7e04322a71f3524694b39a718da6b"></a>

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
action_skip = {}
```

<a id="canonical-0f8a61b153ed2fd0ff090151823c0605071d8e8599d86215f4be85966ee0c678"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 0c1d746aba1b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9f8cb9883757d90ac9ee11a65662eae25cead65f54376b7d64e1ebf30e916f5"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 0c1d746aba1b / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6f1761e3bf021cd742447f098e2467e72c67497e14b399c4be77c6748d790635"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6013f74c6c82181ce0d5e5d8a58afacbaf7ebc79c7a789e16afdc5035e685275"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 291c3a013332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-71a39d343eb490334cf63cbd6ef3dffc65fbe813d4cac171365226dcd65e3bf1"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-895319b1b4fde40adcad1d05d7f216b53531c92af980c5379f5c73865e715e79"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 291c3a013332 / 3

<a id="canonical-79b9ab962f8bb7ca78bdf973116799e70c1f62ecc9061a46d67e478d78cb2da0"></a>

<a id="canonical-797c109b7f4e4b07886a74bd3a9692ecb769c2072f851e0477d714c0bbbac9a2"></a>

## methods property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 291c3a013332 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-a4e2deecc46dce525c2e0f38046ec55214f73b9c6237ad03140f7c01eaba7223"></a>

<a id="canonical-ff73ff2f1e8e588f2442c0d795ef9d6fb7fb49da69a27ae244e281cba09a7e60"></a>

## path property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 291c3a013332 / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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

<a id="canonical-7e2284dc76f4bc661ca3c4ed0dd17cf987decb9a56d0845080af6eeb3178485a"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 291c3a013332 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a741c294a36999450e979f4592ddb1f9ba18a91e56940e1d6bcb9628c519869d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-882e483ad5fdfcfd9408821038ec6a4ddb3653ede4c626e81f0a1523a555feda"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / e96252e23825 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-7f00550a33a3df516d30b0cff422143a5eb9a63db5aa0b7a673e8908097cdcac)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-92b01a09332cc5ec8b6b8ee8be2f81027282796d8efc2ab0f73961c451703225)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-b9b40374f55c7a8bd1e46f75daafbab8e6e634a4a2ffe9dd9d5736d536fb4c54"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-467a9ed9a75c34eb8e9af767b592fd88a036a2f13658113bb5156a0b73ee5946"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / e96252e23825 / 3

<a id="canonical-05fb56988e633f32e269281fe3d3cf0f7bbc83a218552bd1f35a9b8289230ba5"></a>

<a id="canonical-6997291e1930f568962a3dbc2d1b692b100239d7ee49909f576ffaa26f4ca0b8"></a>

## description_spec property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / e96252e23825 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-fcc04d1fdc5ff9fda5b54cd16be20278f95f64305483145b5813980053ead882"></a>

<a id="canonical-9289ed0957971decb64fd8a9729d9699d5bfaecf4bf24a4cc68eb2084bf0ae76"></a>

## name property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / e96252e23825 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-fcefc92ef0928c1b54916a14bfbed487b7accd10428ed9e236ec6c6c3dd7241b"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / e96252e23825 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-a6dec54c1fe3aef8cfa6205288db4add0dd78263fecd9c337b6b68eee675130f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-670eee78cb1c09ce4f3e0e8063e841f734f82a199ce808a951d2309d18099290"></a>

## api_specification.validation_custom_list.open_api_validation_rules — api_specification.validation_custom_list.open_api_validation_rules / 567b1c8780d0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="canonical-40fe9283e334b8bb2716f092dfd33e604fded1e7201f07c148ae99c53da21d4a"></a>

Type: `"object"`. list nested block, Optional.

Validation List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

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

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-9519bda1471f0c65efa886b2c88613cf4ef41dada2ab6e987d28047920d93ed7"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules / 567b1c8780d0 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-006.md#canonical-0165d0e681310784989cfb76ad2c7b2f89a124e335a7a7dc55822f538ac40c52): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-39ea0691dab16edabab886bcc9cf8a1c35e6bfd5e5a048b3fa5046b08bf72583): complete subsection reference.

<a id="canonical-68f198b661aac31b295a848507ab086327f95024a749ce18eb103ad8a9fb1fc3"></a>

<a id="canonical-563d9817b78f72c079745881243edfa50caac29ab71c7f65f008b622f39eceae"></a>

## api_group property — api_specification.validation_custom_list.open_api_validation_rules / 567b1c8780d0 / 4

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

<a id="canonical-ae0830b3f2fd0abf7c0acd57bd71c8fae9a622230f2a5818f99732bd1681138b"></a>

<a id="canonical-75137b56468346876505e9d3e3bbe7987e103d4755d0e33fa81f9de33934d28b"></a>

## base_path property — api_specification.validation_custom_list.open_api_validation_rules / 567b1c8780d0 / 5

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

- [metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-cd33651c770986a99b302068d633761320f2ee9bb6ea11d4201b378f409881b0): complete subsection reference.

<a id="canonical-f54f87a92d8afef8196d4e547d375d12359f0be72b02158c0cdf851af27926f8"></a>

<a id="canonical-68171f1eab6e981178e1b6b852b78f41f0ff323411c2d1b5e6bf2f5a61c3398a"></a>

## specific_domain property — api_specification.validation_custom_list.open_api_validation_rules / 567b1c8780d0 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

- [validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e): complete subsection reference.

<a id="canonical-25f7ec05cfea039c9814e98369db9a6c1f34a0548e7ef6135ef37a42c9ae874b"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules / 567b1c8780d0 / 7

- [api_specification.validation_custom_list.open_api_validation_rules.any_domain](resources--cdn_loadbalancer--reference--group-006.md#canonical-0165d0e681310784989cfb76ad2c7b2f89a124e335a7a7dc55822f538ac40c52)
- [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-39ea0691dab16edabab886bcc9cf8a1c35e6bfd5e5a048b3fa5046b08bf72583)
- [api_specification.validation_custom_list.open_api_validation_rules.metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-cd33651c770986a99b302068d633761320f2ee9bb6ea11d4201b378f409881b0)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0165d0e681310784989cfb76ad2c7b2f89a124e335a7a7dc55822f538ac40c52"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be904e643825f2f7083dbc3e4a3b341237e451c3e09ca7ca97fe1c19e68696cd"></a>

## api_specification.validation_custom_list.open_api_validation_rules.any_domain — api_specification.validation_custom_list.open_api_validation_rules.any_domain / a7c770f7d256 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- api_specification.validation_custom_list.open_api_validation_rules.any_domain

<a id="canonical-9f5a3948bf5a435b21c571547a57494fc32e2783815b5ce13d4351b1bec1b7ce"></a>

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
any_domain = {}
```

<a id="canonical-b652e4a8464516062f3c3e56626424c8f114d60d9f714df7f87f34e7bdfd37dd"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.any_domain / a7c770f7d256 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4a1d7abf002fe9cbeb8b85baf159681d126903134e596b4bf6d854cb18d5b24c"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.any_domain / a7c770f7d256 / 4

- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-39ea0691dab16edabab886bcc9cf8a1c35e6bfd5e5a048b3fa5046b08bf72583"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f361abce04f6b252d4ff302f3fd4372fb77a566ae76d316da4c052b88329619d"></a>

## api_specification.validation_custom_list.open_api_validation_rules.api_endpoint — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 5c59944a949e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="canonical-f5f47aea27a7bd0a3e4009d77ca19dab1b78a38699b594f5d072841820d3b01e"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-0a71e7d16b87e4999ac226ef12c49215a1ef187298a6861e0cd11a937d7633de"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 5c59944a949e / 3

<a id="canonical-c8661b360c7804a98fe47457450f48aa1bf53c91b3ab357ec94f2182b71d817e"></a>

<a id="canonical-a10450254dd0f7cbbcfcc3526959d0d563cdfccc90443303f22e09b6214b28bf"></a>

## methods property — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 5c59944a949e / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-4f83fa30a2564ad731e2a49dfda03c264c49c29b809879f0d4390fe26f04b8ae"></a>

<a id="canonical-f722efce76703fb3e35dbc69fa22c1c05fcc7b6555033d98882e19dc3e27a7a5"></a>

## path property — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 5c59944a949e / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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

<a id="canonical-266d4411244b50cd1dd492cfd1dc94168ec1f3d534722bd713fa1591795b4486"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 5c59944a949e / 6

- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-cd33651c770986a99b302068d633761320f2ee9bb6ea11d4201b378f409881b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a62e46a84402d0c6a44dfafea4d9f1b2a1fe92d909f1f6db596de2bf9c6368b0"></a>

## api_specification.validation_custom_list.open_api_validation_rules.metadata — api_specification.validation_custom_list.open_api_validation_rules.metadata / 69cb3d93e203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- api_specification.validation_custom_list.open_api_validation_rules.metadata

<a id="canonical-5fe59cd7a350e17cdbfafa002172448596e0c4f7556de9dcea02147b1bd33780"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-7e60fb4ab93cb33bbef2f08c294c73562552be662cc03685b68719328c41dc7a"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.metadata / 69cb3d93e203 / 3

<a id="canonical-cc771636a3c23cc1639c7dd810abbae64897fbc022b82540674dfa44248124a4"></a>

<a id="canonical-d55e003578ea73a977641b7aedd6a3a341fcf9870b829a67472fd99a88aa1547"></a>

## description_spec property — api_specification.validation_custom_list.open_api_validation_rules.metadata / 69cb3d93e203 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-bcea6b9a8731fee217e6c259a866e0ca6bbeea4fc9248626778ef1cf32394fc5"></a>

<a id="canonical-e4618a06502b54499a5372869c333b430703708ac7a994f05ec988f6a7d7946b"></a>

## name property — api_specification.validation_custom_list.open_api_validation_rules.metadata / 69cb3d93e203 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-b29d2f2f839e9bea5fde21197411f94de578769c19542f57eed132b7881605c0"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.metadata / 69cb3d93e203 / 6

- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71cee32dfe6c355d8231dd0e6a85628ee0541e8845277c590c90475e48204dcb"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 525addadc222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="canonical-630a0ffaf0f6de0b6ab5d497381557787080169f545237e092a2a1496b5e6c1b"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("response_validation_mode_active",
    "skip_response_validation"),
  validators.ConflictingObjectAttributes("skip_validation",
    "validation_mode_active")}
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
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c62ef67a392027e62209b32c70a055ae1a98a6adb6c8b254c91821ae2a29370"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 525addadc222 / 3

- [response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-192b1e08fd4dd24c38799b2e04956c51911325548ae02ca7e82f254271e3da22): complete subsection reference.

- [skip_response_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-95cabc4777ffe1c78a07ed841c3fe79e00dabdb41198ce3aeb04d15b4fb81e20): complete subsection reference.

- [skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-1d01783e5a3e15401c126de68702a5e29d88f560f859d4ca1e0dd1c989fed50c): complete subsection reference.

- [validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-79ba54fda099fd52b9ffeea15aaf488610aa15ec482e6c779dc610ae7b42d0cf): complete subsection reference.

<a id="canonical-23fd7321d86f25e4243ec28d54d835bc5d7991f9ea72d5f3f919bd6fdaa11a57"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 525addadc222 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-192b1e08fd4dd24c38799b2e04956c51911325548ae02ca7e82f254271e3da22)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-95cabc4777ffe1c78a07ed841c3fe79e00dabdb41198ce3aeb04d15b4fb81e20)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-1d01783e5a3e15401c126de68702a5e29d88f560f859d4ca1e0dd1c989fed50c)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-79ba54fda099fd52b9ffeea15aaf488610aa15ec482e6c779dc610ae7b42d0cf)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-192b1e08fd4dd24c38799b2e04956c51911325548ae02ca7e82f254271e3da22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a24469d4da910142fc2276d5bb7008f650db9189c9494379813b2994463c62c5"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / bb206a1bf9db / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="canonical-59e7cb0419a1bd4ee867ba695397973b8fb10d9556ae0eac5a398b72373cce85"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
response_validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-0df3e3392250e771235e7671f81ba7ebdac1894b0d3a7f98880c9b3396a4b5c7"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / bb206a1bf9db / 3

- [enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-22c292391292e060976cc94ba3e101a5a890a03e70f233b38fa01352ab74bce7): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-d26023ff729414944340cf1d4923cf00af1c2b6dc7226706fa854b0f8c09ee82): complete subsection reference.

<a id="canonical-a06123508c68e24557d782eacfa949830cea10eadc568f6100ffb9642599078b"></a>

<a id="canonical-058a3d7eb327b49f1435d1735323b425b9d9a0c26f21608d5ae3badf574f839e"></a>

## response_validation_properties property — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / bb206a1bf9db / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

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

<a id="canonical-edd39e95b7ca0991b1c80e8ba1ccdf7dd0f0b46d1c586ab32973c5bc5c4d3389"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / bb206a1bf9db / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-22c292391292e060976cc94ba3e101a5a890a03e70f233b38fa01352ab74bce7)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-d26023ff729414944340cf1d4923cf00af1c2b6dc7226706fa854b0f8c09ee82)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-22c292391292e060976cc94ba3e101a5a890a03e70f233b38fa01352ab74bce7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-483459db8205630bb408770c51e177aa235bb6a053716411486ac0f6a1a427c3"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 67b7b52e68ae / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-192b1e08fd4dd24c38799b2e04956c51911325548ae02ca7e82f254271e3da22)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-b2a6bcceb4393c431b3ee803fb3cab3917db1c479089c91bcd77ff50c5e880e7"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_block = {}
```

<a id="canonical-443368af6a0b094482e5ebcf5dfd8371165faf71253fef294acc9e502dd8ed49"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 67b7b52e68ae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da93a5b2aca80c587b6b79c714d6b02689756fd4aa288918e77144437c4c6373"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 67b7b52e68ae / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-192b1e08fd4dd24c38799b2e04956c51911325548ae02ca7e82f254271e3da22)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d26023ff729414944340cf1d4923cf00af1c2b6dc7226706fa854b0f8c09ee82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-819b4638e4fc92e9196f9389e69a26e1feb81558debe42ec7005ceeb503e4198"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 0899f7f88457 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-192b1e08fd4dd24c38799b2e04956c51911325548ae02ca7e82f254271e3da22)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-80983664b4314c40047d57ca766b7a703392cdba03ac51c3690483cfbbd9a1d8"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_report = {}
```

<a id="canonical-e4a8ca957594ce8077c3996da6494fa25efda9ce508410bedb3c0f8103a3d9c9"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 0899f7f88457 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eabc18f664a231a7289a0c7fea743e4612ae00d71acf464d557362683cfc2b99"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 0899f7f88457 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-192b1e08fd4dd24c38799b2e04956c51911325548ae02ca7e82f254271e3da22)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-95cabc4777ffe1c78a07ed841c3fe79e00dabdb41198ce3aeb04d15b4fb81e20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-188ebc83c1415be7771b3948a048b97d0a0e7d84c22e5e91d4e60a83ee4700b8"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 626a5c01e56a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

<a id="canonical-2e2f39eff28cf4da020333fb14ef4880fe000b48224445a5c1f7db52214da374"></a>

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
skip_response_validation = {}
```

<a id="canonical-9ea9dc23421a738ba9b187b56b2e4d8d6e8eccd6e9d65e9babe95023b46df616"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 626a5c01e56a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-84aa077abd5d9771d007603abd792ef1ae2bc43d96eaf77026a484688f12a9bf"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 626a5c01e56a / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1d01783e5a3e15401c126de68702a5e29d88f560f859d4ca1e0dd1c989fed50c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d298c781c1d2a32a2fd59afc9546c19521786942d90638c902de4f6fa38826b1"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / b655d8e3e2d7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation

<a id="canonical-b4adfaf4da6c8e5597263e3a20b569b8b57276f3dc0256fd44ad196fb0aa54b9"></a>

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
skip_validation = {}
```

<a id="canonical-1d570b8bb88a4f7364a605507c58001015ee6869a6d79c8976764d23582f94d0"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / b655d8e3e2d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a95d5c0df9f50ff4d7a562f4dafa21cf9c3bf085fa515f14a7abc6f7f56da7cd"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / b655d8e3e2d7 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-79ba54fda099fd52b9ffeea15aaf488610aa15ec482e6c779dc610ae7b42d0cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9df8d830148095c2b7ffc7f1d6bfef8c166f84c42d07a6a12fa11929b9caed54"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 5e45fed62995 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

<a id="canonical-96c4b4a6ba6336ce63ccad66af734c1b65b6b665086693b2fbfdbb040a04a029"></a>

Type: `"object"`. single nested block, Optional.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-d350cd7d1b4865e93e4a6b709acd4a58dcfa1dd44216baff993faf26e7c63194"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 5e45fed62995 / 3

- [enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-bdd4f75e4ce82fc1cdec583c36712b49150009c42f021402239b560ded7df727): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-9dbac746fde694bcd3c4c72f57e6713791de5f5f31a1db6e9ddcbc0f1dd36159): complete subsection reference.

<a id="canonical-88830a1cce5ea502d6285410792cb1587c9c4a84ef3d5432a892d66fb9b63d44"></a>

<a id="canonical-55a9ec63ee3d41fd3335eb93dbeb0b2d47298bfe52f34c943409ef604d1e9ba0"></a>

## request_validation_properties property — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 5e45fed62995 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

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

<a id="canonical-6a622823ca7141d9c05f8b084de65cb78642d34c6f2d0d02ca07a811a73324d5"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 5e45fed62995 / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-bdd4f75e4ce82fc1cdec583c36712b49150009c42f021402239b560ded7df727)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-9dbac746fde694bcd3c4c72f57e6713791de5f5f31a1db6e9ddcbc0f1dd36159)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bdd4f75e4ce82fc1cdec583c36712b49150009c42f021402239b560ded7df727"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e67c6cd9a8132cbe69427af8b134b0835525dce0ba60ece2f4751260a60d73a"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 6c13f3afc7d7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-79ba54fda099fd52b9ffeea15aaf488610aa15ec482e6c779dc610ae7b42d0cf)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-438351364c230c3730478d0605b06cc87831a352f3656d8f9c1b2b14894e9628"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_block = {}
```

<a id="canonical-977de8daa8afc947c332c840068631950cc11f03dc0f17cb8c3069fde1b75782"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 6c13f3afc7d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0469e7f72fac8ec428881a0206ca009f93606c3c3216e2e45f28c3097abce229"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 6c13f3afc7d7 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-79ba54fda099fd52b9ffeea15aaf488610aa15ec482e6c779dc610ae7b42d0cf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9dbac746fde694bcd3c4c72f57e6713791de5f5f31a1db6e9ddcbc0f1dd36159"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1bd96d1b53668af1ea031437c8b0673ab8ab5cba2bd75d81d52b9775b9c1edc"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 80cd698cf4a5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-18f79b4fa66b7c532b7694710a6944f2c9dba840d749f8f25c10e2dc287778cd)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-d7e6f56f958b4a2ea8366e3109d0593adcbfef9afd3e22662a7222b94c0bfd4e)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-79ba54fda099fd52b9ffeea15aaf488610aa15ec482e6c779dc610ae7b42d0cf)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-599cb041614a5a6909c69e786bffdc8d075dc66c2761c04ee46dec471fc9de79"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enforcement_report = {}
```

<a id="canonical-4aaeef8be09786c26679b3e8219cb53ce4bb252648c5585c1fe9df867ff91b3c"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 80cd698cf4a5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b458bb48a174180d56f1036f0f22742153872099e1586d8559faafc76e8748b"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 80cd698cf4a5 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-79ba54fda099fd52b9ffeea15aaf488610aa15ec482e6c779dc610ae7b42d0cf)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f491bcd391b5ffde59665aff9eab501068729959771574af59a5f1362bd1e2a0"></a>

## api_specification.validation_custom_list.settings — api_specification.validation_custom_list.settings / 1106ef928bc9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- api_specification.validation_custom_list.settings

<a id="canonical-3bdffec60e36dfdc8b0d9bc3ae52c58497af99653a60a6bd3ca01e046c29a07c"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("oversized_body_fail_validation",
    "oversized_body_skip_validation"),
  validators.ConflictingObjectAttributes("property_validation_settings_custom",
    "property_validation_settings_default")}
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
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-5557b081b8ee95ab4e9bdbd19155d2d85bee1af8ec9867ed69381501e63ffb07"></a>

## Direct properties — api_specification.validation_custom_list.settings / 1106ef928bc9 / 3

- [oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-b8509e3fbf50ea584da69ac614d09cd8a08ea50ee355cd7cb62457986dcf88dd): complete subsection reference.

- [oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-38e8b02a494cd4897fac7d9ef7d3b384f2a24ee107bb6778f4b8f59e64db2c9e): complete subsection reference.

- [property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca373cf3c4de1faa471415429530614136db90dbd481112380775ed7fde867ba): complete subsection reference.

- [property_validation_settings_default](resources--cdn_loadbalancer--reference--group-006.md#canonical-2b5139341ce321da885559f3072f28808dc43b31a0022878315324a4a43336ce): complete subsection reference.

<a id="canonical-19fc2c6f2fd495729cdd7f1cef804e575077213ba1597863e16bdeb185cfccb5"></a>

## Next pages — api_specification.validation_custom_list.settings / 1106ef928bc9 / 4

- [api_specification.validation_custom_list.settings.oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-b8509e3fbf50ea584da69ac614d09cd8a08ea50ee355cd7cb62457986dcf88dd)
- [api_specification.validation_custom_list.settings.oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-38e8b02a494cd4897fac7d9ef7d3b384f2a24ee107bb6778f4b8f59e64db2c9e)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca373cf3c4de1faa471415429530614136db90dbd481112380775ed7fde867ba)
- [api_specification.validation_custom_list.settings.property_validation_settings_default](resources--cdn_loadbalancer--reference--group-006.md#canonical-2b5139341ce321da885559f3072f28808dc43b31a0022878315324a4a43336ce)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b8509e3fbf50ea584da69ac614d09cd8a08ea50ee355cd7cb62457986dcf88dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-674e9d295186c459d9a8f62f1dd118a1244d0c222c6802c097240ec12ff5e52d"></a>

## api_specification.validation_custom_list.settings.oversized_body_fail_validation — api_specification.validation_custom_list.settings.oversized_body_fail_validation / abdf5c936e9b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- api_specification.validation_custom_list.settings.oversized_body_fail_validation

<a id="canonical-42c07d9b5f6c8fc71bf4b0495573118b066658ebb96e06f379eeeda8022c7cb0"></a>

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
oversized_body_fail_validation = {}
```

<a id="canonical-5fde12d0aff8e97ab596b0f1f0a7b9246c0ed90b63393ac95f279c14fb3bffb4"></a>

## Direct properties — api_specification.validation_custom_list.settings.oversized_body_fail_validation / abdf5c936e9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-33341529e6eb7230a4428581353f2f59032573a5835d7962420864d3a0a5b951"></a>

## Next pages — api_specification.validation_custom_list.settings.oversized_body_fail_validation / abdf5c936e9b / 4

- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-38e8b02a494cd4897fac7d9ef7d3b384f2a24ee107bb6778f4b8f59e64db2c9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3711f126b69e39224bc6a0dd64e7145b185f97a6200e36f62566c88a7136eb49"></a>

## api_specification.validation_custom_list.settings.oversized_body_skip_validation — api_specification.validation_custom_list.settings.oversized_body_skip_validation / 78acdab1693c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- api_specification.validation_custom_list.settings.oversized_body_skip_validation

<a id="canonical-105a307747514ad987f105783bd177396bb9c70685d47e8b0318345469632677"></a>

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
oversized_body_skip_validation = {}
```

<a id="canonical-dcff0959088a93be59887c97b885cc4a8d414fb49e2ae25c885a4e0d12978b83"></a>

## Direct properties — api_specification.validation_custom_list.settings.oversized_body_skip_validation / 78acdab1693c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b632e52d8797943c42f4354fedf6257d96d45c4a09398927cfa27c64067977e9"></a>

## Next pages — api_specification.validation_custom_list.settings.oversized_body_skip_validation / 78acdab1693c / 4

- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ca373cf3c4de1faa471415429530614136db90dbd481112380775ed7fde867ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f29e78673f5236cb2a8b589e5aabf34b5691cbe60fa29316673856a750af3a70"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom — api_specification.validation_custom_list.settings.property_validation_settings_c / 26b37bda45e2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- api_specification.validation_custom_list.settings.property_validation_settings_custom

<a id="canonical-3e8ddf617419878e43d4a3c8cc6b4ddcb7529902dbdfad22194f4e40ee3166f2"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
property_validation_settings_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-80f0eaf5332172abcf08a48a18c81744a16fc6b8f734789bbb668e0b2a43eee9"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / 26b37bda45e2 / 3

- [query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-408171b667becd03fd11b94db6171f95e25092badf19ce915f059ea6db9acd14): complete subsection reference.

<a id="canonical-7890de93d17fd304ba6c5490c88d7f0d148710fe7963423a7c506658e238a251"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / 26b37bda45e2 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-408171b667becd03fd11b94db6171f95e25092badf19ce915f059ea6db9acd14)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-408171b667becd03fd11b94db6171f95e25092badf19ce915f059ea6db9acd14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d87e1a301123a82041a4905d1cd8a48844e9e36cff7b93d4311e37f2af3e12d1"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters — api_specification.validation_custom_list.settings.property_validation_settings_c / 9030f8ca436d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca373cf3c4de1faa471415429530614136db90dbd481112380775ed7fde867ba)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

<a id="canonical-4e67977439d7d2fd0d4c4699741945ed33b42d8604a3f78ede52c77638b3ad38"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow_additional_parameters",
    "disallow_additional_parameters")}
```

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-a0e30021cdafad90b9f7b681cc450c5ab18d0be71ba44e9ebff745f106c4e80f"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / 9030f8ca436d / 3

- [allow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-b61daabf685fc4bfeca2bf17e8797b3bdceb400d073e9314bb42725fa9953ee7): complete subsection reference.

- [disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-16b4214a2c2f532fe5f2af533b9669869203769cce0b8f70ea21067ade21851f): complete subsection reference.

<a id="canonical-cbffc3d416b51924507f5374cac5192dd160425dc447e9eba27df3780c70b8c1"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / 9030f8ca436d / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-b61daabf685fc4bfeca2bf17e8797b3bdceb400d073e9314bb42725fa9953ee7)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-16b4214a2c2f532fe5f2af533b9669869203769cce0b8f70ea21067ade21851f)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca373cf3c4de1faa471415429530614136db90dbd481112380775ed7fde867ba)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b61daabf685fc4bfeca2bf17e8797b3bdceb400d073e9314bb42725fa9953ee7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cafd9edd9430f0f6c42e31f25a4767b3aae37e8bdce29c0876dc0110a93b0ba1"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters — api_specification.validation_custom_list.settings.property_validation_settings_c / 4fc04aa05062 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca373cf3c4de1faa471415429530614136db90dbd481112380775ed7fde867ba)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-408171b667becd03fd11b94db6171f95e25092badf19ce915f059ea6db9acd14)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-c74a7e76806c2798b5ad4b19b3548b9d4f4b654a3289815e31109d329f7fa5c3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow additional parameters.

Terraform syntax:

```terraform
allow_additional_parameters = {}
```

<a id="canonical-1bf4fa68ee2a3b90395d3cd83eb5c79602c1138b3b36559bd62587aef4ea36fd"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / 4fc04aa05062 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3883841521a0f45fa2be59f9415ed402d347a75be8686cb6e5dd20ee45c8c348"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / 4fc04aa05062 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-408171b667becd03fd11b94db6171f95e25092badf19ce915f059ea6db9acd14)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-16b4214a2c2f532fe5f2af533b9669869203769cce0b8f70ea21067ade21851f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ee16a9aaef3423fabd003b1fb6d6086fc24e6707a83a6bdac46b66de08911d2"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters — api_specification.validation_custom_list.settings.property_validation_settings_c / 7b2ebd139749 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-ca373cf3c4de1faa471415429530614136db90dbd481112380775ed7fde867ba)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-408171b667becd03fd11b94db6171f95e25092badf19ce915f059ea6db9acd14)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-d84fe417ae56e6aefa1c430eb9bd5cbfcdc8986825fa2a73e465f48a8c3c9c84"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disallow additional parameters.

Terraform syntax:

```terraform
disallow_additional_parameters = {}
```

<a id="canonical-8178830c6e4cff358f94cdb9778706ba28d3e686b16c9aa43742c1910cb8dff1"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_c / 7b2ebd139749 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e60d374fb523070966b31d8eed000ef82cc51e75f5ac8ee100fc8948c7f4b411"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_c / 7b2ebd139749 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-408171b667becd03fd11b94db6171f95e25092badf19ce915f059ea6db9acd14)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2b5139341ce321da885559f3072f28808dc43b31a0022878315324a4a43336ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8523f8176859a86c52025654aa8a1289c18aabc46e0c25edd5657bba5fc247ff"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_default — api_specification.validation_custom_list.settings.property_validation_settings_d / ef235eea4298 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-2e9a5006f5dd7517bc974b33a5878d0bbc7568561aed1e961a9f0dd5ea96b55f)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- api_specification.validation_custom_list.settings.property_validation_settings_default

<a id="canonical-52ab9e7b3cea08b0e0a86825a1b0312797157872094eb6763f723f7622a440b2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
property_validation_settings_default = {}
```

<a id="canonical-55224d0c42b17bcc41d1b5610263efa7418668b5f815b8bd0ba4caafa6265e30"></a>

## Direct properties — api_specification.validation_custom_list.settings.property_validation_settings_d / ef235eea4298 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-93a80650a67c913292bd277cc8d09861091803c4c3973d9fa27686fa56319748"></a>

## Next pages — api_specification.validation_custom_list.settings.property_validation_settings_d / ef235eea4298 / 4

- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-b856dab987d3b8da64137bc176eaaec8954230e320afbc9ef8496ae51bd04def)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e1722fb09b74ea9433bd14a1bc5b67040e41a34915ba93990c87f0a463655eee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e20b6d91573f008fcf4914ea4b0f5ff6879adf03dacd555d1e7391bf3d9f380"></a>

## api_specification.validation_disabled — api_specification.validation_disabled / a83ff45bb454 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- api_specification.validation_disabled

<a id="canonical-dd78c3d99010b24b6ccbb1ff0bb28b749edd657d9253e9d0b727c5504b6bf774"></a>

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
validation_disabled = {}
```

<a id="canonical-027d648ecdd82d9b1bd19fbe99c9581637ef4d76bbbf732dd08bc7f74a301fe1"></a>

## Direct properties — api_specification.validation_disabled / a83ff45bb454 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-98bcdc4758bd318713aef9b1e5487ff8b9ba2beb405baf559859c2990ef2961b"></a>

## Next pages — api_specification.validation_disabled / a83ff45bb454 / 4

- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-c4795dc28ae09e65a442a8b65bbdde5c2803918b465345fcbd9ff2ee5b1a2936)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3897308533814e8cc81ad5c3f6a6b8078cda368c7632d9c91a66707448652ed8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4ca0240bd5a696056740ebd924418bec5f0dc4b22e3f29920f64e5526a0dd8f"></a>

## app_firewall — app_firewall / 9dba38b31111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- app_firewall

<a id="canonical-7da195e2e6bc764f813523a12839b646dc821f851391e60df3847dd3ce3f0b4c"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: app\_firewall, disable\_waf; Default: disable\_waf\] Type establishes a direct reference
from one object(the referrer) to another(the referred). Such a reference is in form of
tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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

OneOf alternatives in this subsection:

- [app_firewall](resources--cdn_loadbalancer--reference--group-006.md#canonical-7da195e2e6bc764f813523a12839b646dc821f851391e60df3847dd3ce3f0b4c)
- [disable_waf](resources--cdn_loadbalancer--reference--group-010.md#canonical-3ce5858609fe6b57ae202e3520ad5f49645e7482910e5af32be7e23444970025)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-9762a6a8b9f7c9c995c3cb20b9ff8d7054ea1556d23cab96dc6e4e88e4c8a027"></a>

## Direct properties — app_firewall / 9dba38b31111 / 3

<a id="canonical-eb13036419280ae930e2a0cf4fe90180380217366a4eb67352737834e69c0ff2"></a>

<a id="canonical-34f39f6d036a599eb1475b1abf6e3128a5a806d86ab3a6cef1e20bb0264d7f6f"></a>

## name property — app_firewall / 9dba38b31111 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-e02d08fe1551f7781877b26d894d08382e14bb4c19952cd168d1963574c47473"></a>

<a id="canonical-aa30ae295eb4584d24594a24f60b81649d5cc10e5687471e2f4206e2395a02a3"></a>

## namespace property — app_firewall / 9dba38b31111 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-d8b0b192ef11c9b509eb72278497ab05e2853b25c3bfb3b519e098b65e2be649"></a>

<a id="canonical-aa62ac7a6d822ab71ecf614b6b09b18c60c8a07a67ec612c19a648b2889f6053"></a>

## tenant property — app_firewall / 9dba38b31111 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-968c2b4d33e011c553b5c44b71139ede5b8ed42c9ce75922cbdda79eaeb64a44"></a>

## Next pages — app_firewall / 9dba38b31111 / 7

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-118168ffb743afdb5fc4125b1eb1ed93aa625aae7cdbf8c967248821b6a03d2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-598c6e26d0d3f9e90df6eb9f5bc4c98af4cbe324dc64a9e6dd3144b84165b848"></a>

## blocked_clients — blocked_clients / d242ae1a4cdf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- blocked_clients

<a id="canonical-5b179db293d1e5a162daf36869eaee70528e15af910141229f8f45e097f92da5"></a>

Type: `"object"`. list nested block, Optional.

Define rules to block IP Prefixes or AS numbers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("actions"),
  validators.ConflictingListObjectAttributes("as_number",
    "http_header"),
  validators.ConflictingListObjectAttributes("as_number",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "skip_processing"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ipv6_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("skip_processing",
    "waf_skip_processing")}
```

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
blocked_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-46e1587042c13066d687ae4fe5b1b8e7a7a7bc30ff4f299a9e9dbe46271e83a0"></a>

## Direct properties — blocked_clients / d242ae1a4cdf / 3

<a id="canonical-e0c3122f20795070664ff3f65bb07b03f907289b316777a3ebb75b662b1181b6"></a>

<a id="canonical-04b0b531c411d73446d7c9a3d407ec2a9fab316bc1d11e35ce9d9797bc19ec7b"></a>

## actions property — blocked_clients / d242ae1a4cdf / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Upstream description:

Actions that should be taken when client identifier matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-82ef053441ca89a626c8e3b7b088e5acb066de5f645ba1de61d91edafe30055a"></a>

<a id="canonical-a04c962bcd872e73863c44d82098d14b07dd660c832927437362e2423c145a12"></a>

## as_number property — blocked_clients / d242ae1a4cdf / 5

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix ipv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 401308),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](resources--cdn_loadbalancer--reference--group-007.md#canonical-a50a91f48f5009c57fe80afca029ee6a6f73abf255da1b37f018a311496e98e5): complete subsection reference.

<a id="canonical-f2a1da1a2c3720d4a256d55e7bfcc13e6159cad3cd13b437832078da82c65ae9"></a>

<a id="canonical-c0e58c33497f995e8a93a5dad6b10f019cd361c79006f90a83411c48b42efbc7"></a>

## expiration_timestamp property — blocked_clients / d242ae1a4cdf / 6

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](resources--cdn_loadbalancer--reference--group-007.md#canonical-017e08a116e7b1d2ab78e4d3f0bd5dc5fd9de77a221c8c74f81ddf180a3e9b8f): complete subsection reference.

<a id="canonical-95fa7cf645d84b093935f278958ca60b1d5bf68f5dbf53a0e605664e40c07e4f"></a>

<a id="canonical-dc87cb1b53f4998b5f659d25c3ca1baa66d3b42c6a9a0110109d0d45e63ceb4b"></a>

## ip_prefix property — blocked_clients / d242ae1a4cdf / 7

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ipv6\_prefix user\_identifier\] IPv4 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-9955fb696c592b83522664bdd2f11acf049e2c11c6f6cf8a78e17ca068cc2f45"></a>

<a id="canonical-7c88309800dc9a9bdfd939cdf602bb1affaed74f1d99fe231a48c07cf365bfd1"></a>

## ipv6_prefix property — blocked_clients / d242ae1a4cdf / 8

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-a0800c1e7ec633a290825977d56239e8fc6884a8764a77c7e044aceb6015db37): complete subsection reference.

- [skip_processing](resources--cdn_loadbalancer--reference--group-007.md#canonical-a2464a9d2e7a6d789619971e9f65e9165e3cd233adb79fdd2537c5a0a9726037): complete subsection reference.

<a id="canonical-46ef175c16b452895f22ba210b79f7c37be7959a2dcf2efd486ead8a2268c4c1"></a>

<a id="canonical-e03a59e8e4ddf9fd2584f5e9d9889e2e5461b8e130a46f6cd8351e32fc96cb9b"></a>

## user_identifier property — blocked_clients / d242ae1a4cdf / 9

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix ipv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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

- [waf_skip_processing](resources--cdn_loadbalancer--reference--group-007.md#canonical-7835e37f3afa0217718dfbbeeeeeb60e407d642aafeff7fe43d0ef55c61d8fdf): complete subsection reference.
